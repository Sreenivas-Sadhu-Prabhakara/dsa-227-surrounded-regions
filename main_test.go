package main

import "testing"

func boards(rows ...string) [][]byte {
	g := make([][]byte, len(rows))
	for i, r := range rows {
		g[i] = []byte(r)
	}
	return g
}

func assertBoard(t *testing.T, got [][]byte, want ...string) {
	t.Helper()
	for i := range want {
		if string(got[i]) != want[i] {
			t.Errorf("row %d = %q, want %q", i, got[i], want[i])
		}
	}
}

func TestCapture(t *testing.T) {
	b := boards("XXXX", "XOOX", "XXOX", "XOXX")
	Capture(b)
	assertBoard(t, b, "XXXX", "XXXX", "XXXX", "XOXX") // bottom O touches the border

	// A region connected to the border through a chain survives whole.
	b2 := boards("OXO", "OXO", "OOO")
	Capture(b2)
	assertBoard(t, b2, "OXO", "OXO", "OOO")

	// Fully interior region is captured.
	b3 := boards("XXX", "XOX", "XXX")
	Capture(b3)
	assertBoard(t, b3, "XXX", "XXX", "XXX")

	Capture(nil) // must not panic
	b4 := boards("O")
	Capture(b4)
	assertBoard(t, b4, "O") // border cell survives
}
