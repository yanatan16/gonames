package gonames

import (
	"testing"
)

func TestRand(t *testing.T) {
	s1 := Rand()
	s2 := Rand()

	if s1 == s2 {
		t.Error("A highly unlikely error occurred!")
	}
}