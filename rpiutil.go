// Package rpiutil provides tools for the Raspberry Pi
package rpiutil

import (
	"io/ioutil"
	"strings"
)

const (
	cpuinfo_path   = "/proc/cpuinfo"
	hardware_key   = "Hardware"
	hardware_value = "BCM2708"
	revision_key   = "Revision"
)

// GetPCBRevision read the /proc/cpuinfo and give you the PCB Revision
//
//	-1 : error
//	1 : for revision 0002, 0003
//	2 : for revision 0004, 0005, 0006, 0007, 0008, 0009, 000d, 000e, 000f
//
// If you see a "1000" at the front of the Revision, e.g. 10000002 then it indicates that your Raspberry Pi has been over-volted, and your board revision is simply the last 4 digits (i.e. 0002 in this example).
// more information at http://elinux.org/RPi_HardwareHistory#Board_Revision_History
func GetPCBRevision() int {
	data, _ := ioutil.ReadFile(cpuinfo_path)
	return GetPCBRevisionFrom(string(data))
}

// GetPCBRevisionFrom same as GetPCBRevision but give the PCB Revision by reading the data from the parameters
func GetPCBRevisionFrom(cpuinfo string) int {
	// get the revision string
	revision := ""
	lines := strings.Split(cpuinfo, "\n")
	for _, line := range lines {
		fields := strings.Split(line, "\t: ")
		if fields[0] == hardware_key {
			if len(fields) < 2 || fields[1] != hardware_value {
				break
			}
		}

		if len(fields) > 1 && fields[0] == revision_key {
			revision = fields[1]
			break
		}
	}

	// convert the revision string to int
	// board revision history can be found at http://elinux.org/RPi_HardwareHistory#Board_Revision_History
	switch revision {
	case "":
		return -1
	case "0002", "1000002", "0003", "1000003":
		return 1
	default:
		return 2
	}
}
