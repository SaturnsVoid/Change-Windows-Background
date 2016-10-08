package main

import (
	"fmt"
	"syscall"
	"time"
	"unsafe"
)

var (
	user32               = syscall.NewLazyDLL("user32.dll") //user32.dll
	systemParametersInfo = user32.NewProc("SystemParametersInfoW")
)

func main() {
	ImageLocA := `C:\Users\satur\Desktop\a.jpg`
	ImageLocB := `C:\Users\satur\Desktop\b.jpg`
	fmt.Println("Setting BG to: " + ImageLocA)
	ret, _, _ := systemParametersInfo.Call(20, 0, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(ImageLocA))), 2)
	fmt.Println(ret)
	time.Sleep(5 * time.Second)
	fmt.Println("Setting BG to: " + ImageLocB)
	ret, _, _ = systemParametersInfo.Call(20, 0, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(ImageLocB))), 2)
	fmt.Println(ret)
}
