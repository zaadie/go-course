package main

import "testing"

import (
	"bytes"
	"strconv"
	"testing"
)

func Test(t *testing.T){
	expected := "1
	1
	3"
	actual := Fib(3)
	if expected != actual {
		t.Errorf("Function was incorrect, got %s, want %s", actual, expected)
	}
}

