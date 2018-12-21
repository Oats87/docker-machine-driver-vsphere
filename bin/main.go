package main

import (
	"github.com/docker/machine/libmachine/drivers/plugin"
	vsphere "github.com/oats87/docker-machine-driver-vsphere"
)

func main() {
	plugin.RegisterDriver(vsphere.NewDriver("", ""))
}
