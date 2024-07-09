package process

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Manager struct {
	Pid     int
	User    string
	Command string
	VmSize  string
}

func GetProcess() []Manager {
	process, err := os.ReadDir("/proc")
	if err != nil {
		log.Fatalln(err)
	}

	var ProcessManagerList []Manager

	for _, proc := range process {
		if proc.IsDir() && strings.HasPrefix(proc.Name(), "1") {
			pid, err := strconv.Atoi(proc.Name())
			if err != nil {
				continue
			}

			user, command, vmsize := getProcessInfo(pid)

			news := Manager{Pid: pid, User: user, Command: command, VmSize: vmsize}
			ProcessManagerList = append(ProcessManagerList, news)
		}
	}

	return ProcessManagerList
}

func getProcessInfo(pid int) (string, string, string) {
	statusFile := filepath.Join("/proc", strconv.Itoa(pid), "status")
	statusData, err := os.ReadFile(statusFile)
	if err != nil {
		log.Fatalln(err)
	}

	var user, command, vmsize string

	for _, line := range strings.Split(string(statusData), "\n") {
		if strings.HasPrefix(line, "Uid:") {
			user = strings.Fields(line)[1]
		} else if strings.HasPrefix(line, "Name:") {
			command = strings.TrimPrefix(line, "Name:")
		} else if strings.HasPrefix(line, "VmSize:") {
			vmsize = strings.TrimPrefix(line, "VmSize:")
		}
	}

	return user, command, vmsize
}
