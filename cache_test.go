package main

import (
	"testing"
)

func TestSet(t *testing.T) {
	g := New(1)
	g.Set("Minimal", "Best")
	g.Set("Cute", "Gopher")

	got, err := g.Get("Cute")
	if err != nil {
		t.Errorf("Couldn't retrieve value %v", err)
	}
	want := "Gopher"

	if got != want {
		t.Errorf("got %q, wanted %q ", got, want)
	}
}

func TestGet(t *testing.T) {
	g := New(1)
	g.Set("Minimal", "Best")
	g.Set("Cute", "Gopher")

	got, err := g.Get("Minimal")
	if err == nil {
		t.Errorf("Size of gache not working")
	}
	notWant := "Best"

	if got == notWant {
		t.Errorf("Cache flushing not working")
	}
}
