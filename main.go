package main

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	NULL    = 0
	SRCCOPY = 0x00CC0020
)

type (
	HDC uintptr
)

var (
	user32                 = syscall.NewLazyDLL("User32.dll")
	procGetDC              = user32.NewProc("GetDC")
	procLoadImageW         = user32.NewProc("LoadImageW")
	gdi32                  = syscall.NewLazyDLL("Gdi32.dll")
	procCreateCompatibleDC = gdi32.NewProc("CreateCompatibleDC")
	procSelectObject       = gdi32.NewProc("SelectObject")
	procBitBlt             = gdi32.NewProc("BitBlt")
)

func GetDC(hWnd windows.HWND) (hDC HDC, err error) {
	ptr, _, r3 := syscall.SyscallN(procGetDC.Addr(), uintptr(hWnd))
	if r3 != 0 {
		err = syscall.Errno(r3)
	}

	hDC = HDC(ptr)

	return
}

func CreateCompatibleDC(hDC HDC) (hdcMem HDC, err error) {
	ptr, _, r3 := syscall.SyscallN(procCreateCompatibleDC.Addr(), uintptr(hDC))
	if r3 != 0 {
		err = syscall.Errno(r3)
	}

	hdcMem = HDC(ptr)

	return
}

func LoadImage(hInst windows.Handle, name *uint16, iType uint32, cx int, cy int, fuLoad uint32) (h windows.Handle, err error) {
	ptr, _, r3 := syscall.SyscallN(procLoadImageW.Addr(), uintptr(hInst), uintptr(unsafe.Pointer(name)), uintptr(iType), uintptr(cx), uintptr(cy), uintptr(fuLoad))
	if r3 != 0 {
		err = syscall.Errno(r3)
	}

	h = windows.Handle(ptr)

	return
}

func SelectObject(hdc HDC, h windows.Handle) uintptr {
	ptr, _, _ := syscall.SyscallN(procSelectObject.Addr(), uintptr(hdc), uintptr(h))
	return ptr
}

func BitBlt(hdc HDC, x int, y int, cx int, cy int, hdcSrc HDC, x1 int, y1 int, rop uint32) bool {
	r1, _, _ := syscall.SyscallN(procBitBlt.Addr(), uintptr(hdc), uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(hdcSrc), uintptr(x1), uintptr(y1), uintptr(rop))
	_ = r1
	return true
}

func main() {
	fmt.Println("hello world")
	size := 1
	n := 1000
	offset := 50

	whdc, err := GetDC(NULL)
	if err != nil {
		panic(err)
	}
	hdcMem, err := CreateCompatibleDC(whdc)
	if err != nil {
		panic(err)
	}

	name := "dead.bmp"
	ptr, err := windows.UTF16PtrFromString(name)
	if err != nil {
		panic(err)
	}
	bi, err := LoadImage(0, ptr, 0, size, size, 0x10) // i did 10 instead of hex lol
	if err != nil {
		// panic(err)
	}
	SelectObject(hdcMem, bi)

	for {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				BitBlt(whdc, offset+i*size, offset+j*size, size, size, hdcMem, 0, 0, SRCCOPY)
			}
		}
	}

	// // dont need other stuff
	// refresh := time.NewTicker(1 * time.Millisecond) // ~30hz ~30fps
	// for {
	// 	select {
	// 	case <-refresh.C:
	// 		for i := 0; i < 1000; i++ {
	// 			for j := 0; j < 1000; j++ {
	// 				BitBlt(whdc, i*size, j*size, size, size, hdcMem, 0, 0, SRCCOPY)
	// 			}
	// 		}
	// 	}
	// }
}
