package main

import (
	"fmt"
	"myfmt"
	"os"
)

func main() {
	fmt.Println("This is the main program")
	test, _ := myfmt.Single(os.Stdin, os.Stderr)
	fmt.Println("The single word is", test)
	test2, _ := myfmt.Sentence(os.Stdin, os.Stderr)
	fmt.Println("The sentence entered is", test2)
}
