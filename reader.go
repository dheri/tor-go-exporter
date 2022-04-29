package main

import (
	"fmt"

	"github.com/reiver/go-telnet"
)

func main() {
	fmt.Println("Hello, world.")

	var handler telnet.Handler = telnet.EchoHandler

	err := telnet.ListenAndServe(":5555", handler)
	if nil != err {
		//@TODO: Handle this error better.
		panic(err)
	}
}
