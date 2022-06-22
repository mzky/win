package win

import (
	"syscall"
	"unsafe"
)

type SHCONTF int

const (
	SHCONTF_CHECKING_FOR_CHILDREN SHCONTF = 0x10
	SHCONTF_FOLDERS               SHCONTF = 0x20
	SHCONTF_NONFOLDERS            SHCONTF = 0x40
	SHCONTF_INCLUDEHIDDEN         SHCONTF = 0x80
	SHCONTF_INIT_ON_FIRST_NEXT    SHCONTF = 0x100
	SHCONTF_NETPRINTERSRCH        SHCONTF = 0x200
	SHCONTF_SHAREABLE             SHCONTF = 0x400
	SHCONTF_STORAGE               SHCONTF = 0x800
	SHCONTF_NAVIGATION_ENUM       SHCONTF = 0x1000
	SHCONTF_FASTITEMS             SHCONTF = 0x2000
	SHCONTF_FLATLIST              SHCONTF = 0x4000
	SHCONTF_ENABLE_ASYNC          SHCONTF = 0x8000
	SHCONTF_INCLUDESUPERHIDDEN    SHCONTF = 0x10000
)

var IID_IShellFolder = IID{0x000214E6, 0x0000, 0x0000, [8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

type IShellFolderVtbl struct {
	IUnknownVtbl
	ParseDisplayName uintptr
	EnumObjects      uintptr
	BindToObject     uintptr
	BindToStorage    uintptr
	CompareIDs       uintptr
	CreateViewObject uintptr
	GetAttributesOf  uintptr
	GetUIObjectOf    uintptr
	GetDisplayNameOf uintptr
	SetNameOf        uintptr
}

type IShellFolder struct {
	LpVtbl *IShellFolderVtbl
}

func (is *IShellFolder) Release() HRESULT {
	ret, _, _ := syscall.Syscall(is.LpVtbl.Release, 1,
		uintptr(unsafe.Pointer(is)),
		0,
		0)

	return HRESULT(ret)
}

func (is *IShellFolder) EnumObjects(h HWND, flags SHCONTF) (IEnumIDList, HRESULT) {
	var lpData IEnumIDList
	ret, _, _ := syscall.Syscall(is.LpVtbl.EnumObjects, 3,
		uintptr(h),
		uintptr(flags),
		uintptr(unsafe.Pointer(&lpData)))

	return lpData, HRESULT(ret)
}

var IID_IShellFolder2 = IID{0x93f2f68c, 0x1d1b, 0x11d3, [8]byte{0xa3, 0xe, 0x0, 0xc0, 0x4f, 0x79, 0xab, 0xd1}}

type IShellFolder2Vtbl struct {
	IUnknownVtbl
	GetDefaultSearchGUID  uintptr
	EnumSearches          uintptr
	GetDefaultColumn      uintptr
	GetDefaultColumnState uintptr
	GetDetailsEx          uintptr
	GetDetailsOf          uintptr
	MapColumnToSCID       uintptr
}

type IShellFolder2 struct {
	LpVtbl *IShellFolder2Vtbl
}

var IID_IEnumIDList = IID{0x000214F2, 0x0000, 0x0000, [8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

type IEnumIDListVtbl struct {
	IUnknownVtbl
	Next  uintptr
	Skip  uintptr
	Reset uintptr
	Clone uintptr
}

type IEnumIDList struct {
	LpVtbl *IEnumIDListVtbl
}

func (is *IEnumIDList) Next(celt uint32) (*ITEMIDLIST, uint32, HRESULT) {
	var lpData *ITEMIDLIST
	var cIds uint32
	ret, _, _ := syscall.Syscall(is.LpVtbl.Next, 3,
		uintptr(celt),
		uintptr(unsafe.Pointer(&lpData)),
		uintptr(unsafe.Pointer(&cIds)))

	return lpData, cIds, HRESULT(ret)
}

type SHITEMID struct {
	cb   uint16
	abID [1]byte
}
type ITEMIDLIST struct {
	mkid SHITEMID
}
