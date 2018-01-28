package main

import (
	"fmt"
)

type greeting string

// Greeter exposed
var Greeter greeting

func (g greeting) Greet() {
	fmt.Println("你好宇宙")
}
