package System

import (
	"fmt"
	"runtime"
)

var OSType string

func InitOS() error {
	switch runtime.GOOS {
	case "linux", "windows":
		OSType = runtime.GOOS
		return nil
	default:
		return fmt.Errorf("unsupported os: %s", runtime.GOOS)
	}
}
