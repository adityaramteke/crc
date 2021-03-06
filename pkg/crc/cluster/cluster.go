package cluster

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/code-ready/crc/pkg/crc/errors"
	"github.com/code-ready/machine/libmachine/drivers"
)

func WaitForSsh(driver drivers.Driver) error {
	checkSshConnectivity := func() error {
		_, err := drivers.RunSSHCommandFromDriver(driver, "exit 0")
		if err != nil {
			return &errors.RetriableError{Err: err}
		}
		return nil
	}

	return errors.RetryAfter(10, checkSshConnectivity, time.Second)
}

// CheckCertsValidity checks if the cluster certs have expired
func CheckCertsValidity(driver drivers.Driver) error {
	certExpiryDateCmd := `date --date="$(sudo openssl x509 -in /var/lib/kubelet/pki/kubelet-client-current.pem -noout -enddate | cut -d= -f 2)" --iso-8601=seconds`
	output, err := drivers.RunSSHCommandFromDriver(driver, certExpiryDateCmd)
	if err != nil {
		return err
	}
	certExpiryDate, err := time.Parse(time.RFC3339, strings.TrimSpace(output))
	if err != nil {
		return err
	}
	if time.Now().After(certExpiryDate) {
		return fmt.Errorf("Certs have expired, they were valid till: %s", certExpiryDate.Format(time.RFC822))
	}
	return nil
}

// Return size of disk, used space in bytes and the mountpoint
func GetRootPartitionUsage(driver drivers.Driver) (int64, int64, error) {
	cmd := "df -B1 --output=size,used,target /sysroot | tail -1"

	out, err := drivers.RunSSHCommandFromDriver(driver, cmd)

	if err != nil {
		return 0, 0, err
	}
	diskDetails := strings.Split(strings.TrimSpace(out), " ")
	diskSize, err := strconv.ParseInt(diskDetails[0], 10, 64)
	if err != nil {
		return 0, 0, err
	}
	diskUsage, err := strconv.ParseInt(diskDetails[1], 10, 64)
	if err != nil {
		return 0, 0, err
	}
	return diskSize, diskUsage, nil
}
