package main

import "fmt"

// Capture flips every 'O' region fully surrounded by 'X' to 'X', in place.
// Inverted reasoning: a region survives iff it touches the border, so mark
// border-connected Os safe first and flip everything else.
func Capture(board [][]byte) {
	rows := len(board)
	if rows == 0 {
		return
	}
	var save func(r, c int)
	save = func(r, c int) {
		if r < 0 || r >= rows || c < 0 || c >= len(board[r]) || board[r][c] != 'O' {
			return
		}
		board[r][c] = 'S' // safe: connected to the border
		save(r+1, c)
		save(r-1, c)
		save(r, c+1)
		save(r, c-1)
	}
	for r := 0; r < rows; r++ {
		save(r, 0)
		save(r, len(board[r])-1)
	}
	for c := 0; c < len(board[0]); c++ {
		save(0, c)
		save(rows-1, c)
	}
	for r := range board {
		for c := range board[r] {
			switch board[r][c] {
			case 'O':
				board[r][c] = 'X' // surrounded: captured
			case 'S':
				board[r][c] = 'O' // safe: restored
			}
		}
	}
}

func main() {
	board := [][]byte{
		[]byte("XXXX"),
		[]byte("XOOX"),
		[]byte("XXOX"),
		[]byte("XOXX"),
	}
	Capture(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
