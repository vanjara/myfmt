package myfmt_test

import (
	"io/ioutil"
	"myfmt"
	"strings"
	"testing"
)

func TestSingle(t *testing.T) {
	input := strings.NewReader("One\n")
	want := "One"
	//input := bytes.NewBufferString("One\n")
	got, err := myfmt.Single(input, ioutil.Discard)
	if err != nil {
		t.Error(err)
	}
	if want != got {
		t.Errorf("Checking: want - %s, got - %s", want, got)
	}
	if input.Len() != 0 {
		t.Errorf("Given input not fully consumed, data still left to consume %d\n", input.Len())
	}
}

func TestSentence(t *testing.T) {
	input := strings.NewReader("This is a sentence\n")
	//input := bufio.NewReader(strings.NewReader("This is a sentence\n"))
	//input := bytes.NewBufferString("This \nis a multi line sentence\n")
	got, err := myfmt.Sentence(input, ioutil.Discard)
	want := "This is a sentence"
	if err != nil {
		t.Error(err)
	}
	if want != got {
		t.Errorf("Checking: want - %s, got - %s", want, got)
	}

	if input.Len() != 0 {
		t.Errorf("Given input not fully consumed - only consumed - %s, data still left to consume %d\n", got, input.Len())
	}
}
