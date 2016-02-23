package main

import "testing"

func TestSomething(t *testing.T) {
	if "foo" != "foo" {
		t.Error("Expected foo == foo")
	}
}
