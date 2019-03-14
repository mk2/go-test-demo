package calc_test

import (
	"testing"

	"github.com/mk2/go-test-demo"
)

func Test_Plus(t *testing.T) {
	a := 3
	b := 2
	expected := 5
	if calc.Plus(a, b) != expected {
		t.Fatalf("not expected result: " + string(expected))
	}
}

func Test_Minus(t *testing.T) {
	a := 3
	b := 2
	expected := 1
	result := calc.Minus(a, b)
	if result != expected {
		t.Errorf("wanted: %d\nreceived: %d\n", expected, result)
	}
}
