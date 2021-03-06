package config

import (
	validations "github.com/code-ready/crc/pkg/crc/config"
)

var (
	// Preflight checks
	SkipCheckVirtEnabled             = createSetting("skip-check-virt-enabled", nil, []validationFnType{validations.ValidateBool})
	WarnCheckVirtEnabled             = createSetting("warn-check-virt-enabled", nil, []validationFnType{validations.ValidateBool})
	SkipCheckKvmEnabled              = createSetting("skip-check-kvm-enabled", nil, []validationFnType{validations.ValidateBool})
	WarnCheckKvmEnabled              = createSetting("warn-check-kvm-enabled", nil, []validationFnType{validations.ValidateBool})
	SkipCheckLibvirtInstalled        = createSetting("skip-check-libvirt-installed", nil, []validationFnType{validations.ValidateBool})
	WarnCheckLibvirtInstalled        = createSetting("warn-check-libvirt-installed", nil, []validationFnType{validations.ValidateBool})
	SkipCheckLibvirtEnabled          = createSetting("skip-check-libvirt-enabled", nil, []validationFnType{validations.ValidateBool})
	WarnCheckLibvirtEnabled          = createSetting("warn-check-libvirt-enabled", nil, []validationFnType{validations.ValidateBool})
	SkipCheckLibvirtRunning          = createSetting("skip-check-libvirt-running", nil, []validationFnType{validations.ValidateBool})
	WarnCheckLibvirtRunning          = createSetting("warn-check-libvirt-running", nil, []validationFnType{validations.ValidateBool})
	SkipCheckUserInLibvirtGroup      = createSetting("skip-check-user-in-libvirt-group", nil, []validationFnType{validations.ValidateBool})
	WarnCheckUserInLibvirtGroup      = createSetting("warn-check-user-in-libvirt-group", nil, []validationFnType{validations.ValidateBool})
	SkipCheckIPForwarding            = createSetting("skip-check-ip-forwarding", nil, []validationFnType{validations.ValidateBool})
	WarnCheckIPForwarding            = createSetting("warn-check-ip-forwarding", nil, []validationFnType{validations.ValidateBool})
	SkipCheckLibvirtDriver           = createSetting("skip-check-libvirt-driver", nil, []validationFnType{validations.ValidateBool})
	WarnCheckLibvirtDriver           = createSetting("warn-check-libvirt-driver", nil, []validationFnType{validations.ValidateBool})
	SkipCheckCrcNetwork              = createSetting("skip-check-crc-network", nil, []validationFnType{validations.ValidateBool})
	WarnCheckCrcNetwork              = createSetting("warn-check-crc-network", nil, []validationFnType{validations.ValidateBool})
	SkipCheckCrcNetworkActive        = createSetting("skip-check-crc-network-active", nil, []validationFnType{validations.ValidateBool})
	WarnCheckCrcNetworkActive        = createSetting("warn-check-crc-network-active", nil, []validationFnType{validations.ValidateBool})
	SkipCheckCrcDnsmasqFile          = createSetting("skip-check-crc-dnsmasq-file", nil, []validationFnType{validations.ValidateBool})
	WarnCheckCrcDnsmasqFile          = createSetting("warn-check-crc-dnsmasq-file", nil, []validationFnType{validations.ValidateBool})
	SkipCheckCrcNetworkManagerConfig = createSetting("skip-check-network-manager-config", nil, []validationFnType{validations.ValidateBool})
	WarnCheckCrcNetworkManagerConfig = createSetting("warn-check-network-manager-config", nil, []validationFnType{validations.ValidateBool})
)
