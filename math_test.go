package main

import "testing"

func TestSum(t *testing.T) {

	total := Sum(15, 15)

	if total != 30 {
		t.Errorf("Wrong sum result, expected: %d. got %d", 30, total)
	}
}