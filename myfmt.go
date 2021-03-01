package myfmt

import (
	"bufio"
	"fmt"
	"io"
)

//Single word Testing...
func Single(r io.Reader, w io.Writer) (string, error) {
	var word string
	fmt.Fprintln(w, "Please enter a single word: ")
	_, _ = fmt.Fscanln(r, &word)
	return word, nil
}

//Sentence word Testing...
func Sentence(r io.Reader, w io.Writer) (string, error) {
	fmt.Fprintln(w, "Please enter a sentence: ")
	// var word string
	// _, _ = fmt.Fscanln(r, &word)

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	word := scanner.Text()

	return word, nil
}
