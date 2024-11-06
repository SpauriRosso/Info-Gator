package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"syscall"
	"time"
)

func main() {
	sysRep()
}

func sysRep() {
	fmt.Println("========== System Report ==========")
	fmt.Printf("OS : %s\n", runtime.GOOS)
	fmt.Printf("Architecture : %s\n", runtime.GOARCH)
	fmt.Printf("Number of CPU : %d\n", runtime.NumCPU())
	fmt.Printf("Running time : %s\n", timeFormat(time.Since(upTime())))

	var sysInfo syscall.Sysinfo_t
	err := syscall.Sysinfo(&sysInfo)
	if err != nil {
		log.Println("Error getting system memory info:", err)
		return
	}
	totalMem := sysInfo.Totalram * uint64(syscall.Getpagesize()) / 1024 / 1024
	freeMem := sysInfo.Freeram * uint64(syscall.Getpagesize()) / 1024 / 1024

	fmt.Printf("Total Memeory : %v MB\n", totalMem)
	fmt.Printf("Free Memory : %v MB\n", freeMem)
}

func upTime() time.Time {
	info, err := os.Stat("/proc/1")
	if err == nil {
		return info.ModTime()
	}
	return time.Now()
}

func timeFormat(t time.Duration) string {
	hours := int(t.Hours())
	minutes := int(t.Minutes()) % 60
	seconds := int(t.Seconds()) % 60
	return fmt.Sprintf("%dh%dm%ds", hours, minutes, seconds)
}
