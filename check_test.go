package main

import (
	"testing"
	"testing/quick"
)

func add(x, y int) int {
	return x + y
}

func TestAddingEx1(t *testing.T) {
	if add(3, 2) != 5 {
		t.Error("3 plus 2 is 5")
	}
}

func TestAddingZeroMakesNoDifference(t *testing.T) {
	assertion := func(x int) bool {
		return add(x, 0) == x
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error(err)
	}
}

func TestAssociativity(t *testing.T) {
	assertion := func(x int) bool {
		return add(add(x, 1), 1) == add(x, 2)
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error(err)
	}
}

func TestParamOrderDoesNotMatter(t *testing.T) {
	assertion := func(x, y int) bool {
		return add(x, y) == add(y, x)
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error(err)
	}
}
