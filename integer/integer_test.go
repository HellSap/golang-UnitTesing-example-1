package main

import "testing"

func TestSimpleIntrest(t *testing.T) {
    t.Run("taking all zero as input", func(t *testing.T) {
	got := SimpleIntrest(1000, 5, 2)
	want := 100
	assertCorrectMessage(t, got, want)
    })
	t.Run("taking all zero as input", func(t *testing.T) {
		got := SimpleIntrest(0, 0, 0)
		want := 0
		assertCorrectMessage(t, got, want)
	})
	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := SimpleIntrest(-100, 2, 4)
		want := 0
		assertCorrectMessage(t, got, want)
	})
	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := SimpleIntrest(400, -1, 0)
		want := 0
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
