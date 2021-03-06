package main

import (
	"bytes"
	"testing"
)

func TestExec_exec(t *testing.T) {
	src := "-4^7*2+1"
	bufo := bytes.NewBuffer([]byte{})
	bufe := bytes.NewBuffer([]byte{})
	cli := NewCli(bufo, bufe, func(int int) {})
	cli.exec([]string{"", src})
	actual := bufo.String()
	expected := `Expression: (((-4 ^ 7) * 2) + 1)
tvals: 4
l0: Op: SUBTRACT, Arg1: {}, Arg2: { Val: 4, INTEGER }, Ret: { TV1 }
l1: Op: POW, Arg1: { TV1 }, Arg2: { Val: 7, INTEGER }, Ret: { TV2 }
l2: Op: MULTIPLY, Arg1: { TV2 }, Arg2: { Val: 2, INTEGER }, Ret: { TV3 }
l3: Op: ADD, Arg1: { TV3 }, Arg2: { Val: 1, INTEGER }, Ret: { TV4 }

`
	if actual != expected {
		t.Errorf("\nexpected\n%q \ngot\n%q", expected, actual)
	}
}

func TestExec_missing_argument(t *testing.T) {
	expected := `example usage: ./compile "1 + 1"
`
	bufo := bytes.NewBuffer([]byte{})
	bufe := bytes.NewBuffer([]byte{})
	cli := NewCli(bufo, bufe, func(int int) {})
	cli.exec([]string{})
	actual := bufe.String()
	if actual != expected {
		t.Errorf("expected %v got %v", expected, actual)
	}
}
