package adbgo

import (
	"fmt"
	"os/exec"
)

type ADBBridge struct {
	RemoteAddr string
}

func NewADBBridge(addr string) *ADBBridge {
	return &ADBBridge{
		RemoteAddr: addr,
	}
}

func (a *ADBBridge) Connect() error {
	cmd := exec.Command("adb", "connect", a.RemoteAddr)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println(string(output))
	return nil
}

func (a *ADBBridge) SetPrivateDNS(dns string) error {
	cmd := exec.Command("adb", "shell", "settings", "put", "global", "private_dns_specifier", dns)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println(string(output))
	return nil
}

func (a *ADBBridge) TurnOffPrivateDNS() error {
	cmd := exec.Command("adb", "shell", "settings", "put", "global", "private_dns_mode", "off")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println(string(output))
	return nil
}
