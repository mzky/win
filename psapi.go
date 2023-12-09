package win

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

var (
	libpsapi             *windows.LazyDLL
	getModuleFileNameExW *windows.LazyProc
)

func init() {
	libpsapi = windows.NewLazySystemDLL("Psapi.dll")
	getModuleFileNameExW = libpsapi.NewProc("GetModuleFileNameExW")
}

func GetModuleFileNameEx(process, module HANDLE) (string, bool) {
	size := windows.MAX_LONG_PATH
	name := make([]uint16, size)
	ret, _, _ := syscall.Syscall6(getModuleFileNameExW.Addr(), 4,
		uintptr(process),
		uintptr(module),
		uintptr(unsafe.Pointer(&name[0])),
		uintptr(size),
		0,
		0)
	if ret > 0 {
		name = name[:ret]
	}
	return UTF16PtrToString(&name[0]), ret > 0
}
