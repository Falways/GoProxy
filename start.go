package main

import (
	"os/exec"
)

func main() {
	comand := "cyberexclient.exe --tundev tap0901:tapa:10.10.1.4:10.10.1.0:255.255.255.0 --netif-ipaddr 10.10.1.1 --netif-netmask 255.255.255.0 --socks-server-addr 127.0.0.1:2080 --udpgw-remote-server-addr 127.0.0.1:4446"
	cmd := exec.Command("cmd.exe","/c",comand)
	cmd.Run()
}
