package main

import (
	"fmt"
	"runtime"
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
	fmt.Printf("Running time : %s", time.Since(startTime()))
}

func startTime() time.Time {
	return time.Now()
}
