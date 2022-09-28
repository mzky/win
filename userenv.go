package win

import (
	"golang.org/x/sys/windows"
	"syscall"
)

var (
	libuserenv              *windows.LazyDLL
	createEnvironmentBlock  *windows.LazyProc
	destroyEnvironmentBlock *windows.LazyProc
)

func init() {
	libuserenv = windows.NewLazySystemDLL("userenv.dll")
	createEnvironmentBlock = libuserenv.NewProc("CreateEnvironmentBlock")
	destroyEnvironmentBlock = libuserenv.NewProc("DestroyEnvironmentBlock")
}

func CreateEnvironmentBlock(lpEnvironment, hToken HANDLE, bInherit bool) bool {
	ret, _, _ := syscall.Syscall(createEnvironmentBlock.Addr(), 3,
		uintptr(lpEnvironment),
		uintptr(hToken),
		uintptr(BoolToBOOL(bInherit)))
	return ret == 1
}

func DestroyEnvironmentBlock(lpEnvironment HANDLE) bool {
	ret, _, _ := syscall.Syscall(destroyEnvironmentBlock.Addr(), 1,
		uintptr(lpEnvironment),
		0,
		0)
	return ret == 1
}
