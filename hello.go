package main

import (
	"fmt"

	"github.com/Tomy2e/go-module-test/askname"
)

func main() {
	fmt.Println("Hello World!")

	name := askname.Prompt("What's your name?")

	fmt.Println("Your name is", name)
}
