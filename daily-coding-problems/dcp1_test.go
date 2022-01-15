package dcp

import (
	"testing"
)

func Test_Dcp1(t *testing.T) {
	t.Parallel()
	exp := 19
	act := dcp1(1)
	if act != exp {
		t.Fatalf("expected %v, got %v", exp, act)
	}
}

func Test_Dcp1_TDT(t *testing.T) {
	t.Parallel()
	table := []struct {
		name   string
		input  int
		answer int
		err    error
	}{
		{name: "test1", input: 1, answer: 19},
		{name: "test2", input: 2, answer: 28},
		{name: "test3", input: 3, answer: 37},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			act := dcp1(tt.input)

			if tt.answer != act {
				t.Fatalf("expect %v but got %v instead", tt.err, act)
			}
		})
	}
}
