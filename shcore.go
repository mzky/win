// +build windows

package win

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

var (
	// Library
	shcore *windows.LazyDLL

	// Functions
	getDpiForMonitor         *windows.LazyProc
	getScaleFactorForMonitor *windows.LazyProc
)

const (
	MDT_EFFECTIVE_DPI = 0
	MDT_ANGULAR_DPI
	MDT_RAW_DPI
)

func init() {
	// Library
	shcore = windows.NewLazySystemDLL("Shcore.dll")

	// Functions
	getDpiForMonitor = shcore.NewProc("GetDpiForMonitor")
	getScaleFactorForMonitor = shcore.NewProc("GetScaleFactorForMonitor")
}
func GetDpiForMonitor(hWnd HMONITOR, flag int) (uint32, uint32, bool) {
	if getDpiForMonitor.Find() != nil {
		return 0, 0, false
	}
	var dpiX, dpiY uint32
	ret, _, _ := syscall.Syscall6(getDpiForMonitor.Addr(), 4,
		uintptr(hWnd),
		uintptr(flag),
		uintptr(unsafe.Pointer(&dpiX)),
		uintptr(unsafe.Pointer(&dpiY)),
		0,
		0)

	return dpiX, dpiY, ret == S_OK
}
func GetScaleFactorForMonitor(hWnd HMONITOR) (uint32, bool) {
	if getScaleFactorForMonitor.Find() != nil {
		return 0, false
	}
	var dpiX uint32
	ret, _, _ := syscall.Syscall(getScaleFactorForMonitor.Addr(), 2,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(&dpiX)),
		0)

	return dpiX, ret == S_OK
}
