package askname

import "fmt"

func Prompt(phrase string) (name string) {
	fmt.Print(phrase, " ")
	fmt.Scanln(&name)
	return
}
