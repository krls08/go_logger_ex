package logger

import "fmt"

var Version string = "1.0"

func Log(msg string) {
	fmt.Println("[LOG] " + msg)
}