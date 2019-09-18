package main

import "testing"

func TestMyAdd(t *testing.T) {
	a := 1
	b := 2
	c := 3
	if MyAdd(a, b) != c {
		t.Errorf("%d + %d =, excepted %d, got %d", a, b, c, MyAdd(a, b))
	}
}

func TestMyMutiply(t *testing.T) {
	a := 2
	b := 3
	c := 60
	if MyMutiply(a, b) != c {
		t.Errorf("%d * %d =, excepted %d, got %d", a, b, c, MyMutiply(a, b))
	}
}
