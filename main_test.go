package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hello Golang"

	got := hello()

	if want != got {
		t.Fatalf("want %s, but got %s", want, got)
	}
}
