package main

import (
	"fmt"
	pluginapi "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
)

func getFpgaCpuDevices(number int) (devs []*pluginapi.Device) {
	for i := 0; i < number; i++ {

		fpgaCpuDevice := &pluginapi.Device{
			ID:     fmt.Sprintf("%s-%d", fpgaCpu, i),
			Health: pluginapi.Healthy,
		}

		devs = append(devs, fpgaCpuDevice)
	}
	return
}

func getFpgaMemDevices(number int) (devs []*pluginapi.Device) {
	for i := 0; i < number; i++ {
		fpgaMemDevice := &pluginapi.Device{
			ID:     fmt.Sprintf("%s-%d-%d", fpgaMem, MemoryBlockSize, i),
			Health: pluginapi.Healthy,
		}
		devs = append(devs, fpgaMemDevice)
	}
	return
}

func deviceExists(devs []*pluginapi.Device, id string) bool {
	for _, d := range devs {
		if d.ID == id {
			return true
		}
	}
	return false
}

// ToDo: need TPU lib Supported
func deviceIsFree(devs []*pluginapi.Device, id string) bool {
	for _, d := range devs {
		if d.ID == id {
			if d.Health == pluginapi.Healthy {
				return true
			} else {
				return false
			}
		}
	}
	return false
}
