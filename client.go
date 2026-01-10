package adbgo

import (
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
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}

func (a *ADBBridge) Disconnect() error {
	cmd := exec.Command("adb", "disconnect")
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}

func (a *ADBBridge) SetPrivateDNS(dns string) error {
	cmd := exec.Command("adb", "shell", "settings", "put", "global", "private_dns_mode", "hostname")
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	cmd = exec.Command("adb", "shell", "settings", "put", "global", "private_dns_specifier", dns)
	_, err = cmd.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}

func (a *ADBBridge) TurnOffPrivateDNS() error {
	cmd := exec.Command("adb", "shell", "settings", "put", "global", "private_dns_mode", "off")
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}

func (a *ADBBridge) SendText(text string) error {
	cmd := exec.Command("adb", "shell", "input keyboard text '"+text+"'")
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}

func (a *ADBBridge) DeleteChar() error {
	cmd := exec.Command("adb", "shell", "input keyevent 67")
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}
