package main

import "testing"

func TestSum(t *testing.T) {
	total := sum(20, 30)

	if total != 50 {
		t.Errorf("sum(20, 30) = %d; want 50", total)
	}
}
