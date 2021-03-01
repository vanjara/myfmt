package myfmt_test

import (
	"io/ioutil"
	"myfmt"
	"strings"
	"testing"
)

func TestSingle(t *testing.T) {
	input := strings.NewReader("One\n")
	//input := bytes.NewBufferString("One\n")
	_, err := myfmt.Single(input, ioutil.Discard)
	if err != nil {
		t.Error(err)
	}
	if input.Len() != 0 {
		t.Errorf("Given input not fully consumed, data still left to consume %d\n", input.Len())
	}
}

func TestSentence(t *testing.T) {
	input := strings.NewReader("This is a sentence\n")
	//input := bytes.NewBufferString("This \nis a multi line sentence\n")
	ret_word, err := myfmt.Single(input, ioutil.Discard)
	if err != nil {
		t.Error(err)
	}
	if input.Len() != 0 {
		t.Errorf("Given input not fully consumed - only consumed %s, data still left to consume %d\n", ret_word, input.Len())
	}
}
