package main

import (
	"fmt"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Move struct {
	row, col int
	dir      Direction
}

func (move Move) String() string {
	var dir string
	switch move.dir {
	case Up:
		dir = "Up"
	case Down:
		dir = "Down"
	case Left:
		dir = "Left"
	case Right:
		dir = "Right"
	}
	return fmt.Sprintf("(%d, %d) %s", move.row, move.col, dir)
}

type Hole int

const (
	Empty Hole = iota
	Filled
	Blocked
)

type Board [][]Hole

func (board *Board) print() {
	for _, row := range *board {
		fmt.Print("| ")
		for _, hole := range row {
			switch hole {
			case Empty:
				fmt.Print("   |")
			case Filled:
				fmt.Print(" X |")
			case Blocked:
				fmt.Print(" # |")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

type MoveSequence []Move

func appendSolutions(solutions []MoveSequence, board Board, move Move, history MoveSequence, stoneCount int) []MoveSequence {
	history = append(history, move)
	stoneCount--
	if stoneCount == 1 {
		return []MoveSequence{history}
	}

	newBoard := applyMove(board, move)
	nextMoves := allMoves(newBoard)
	// if stoneCount < 5 {
	// 	fmt.Print("\n")
	// 	fmt.Println("Stones: ", stoneCount, "Moves: ", len(nextMoves), nextMoves)
	// 	newBoard.print()
	// }

	for _, nextMove := range nextMoves {
		solutions = appendSolutions(solutions, newBoard, nextMove, history, stoneCount)
	}
	return solutions
}

// All legal moves that will fill the hole at (row, col)
func fillMoves(board Board, row int, col int) []Move {
	moves := []Move{}
	if col > 1 && board[row][col-1] == Filled && board[row][col-2] == Filled {
		moves = append(moves, Move{row, col - 2, Right})
	}
	if col < len(board[0])-2 && board[row][col+1] == Filled && board[row][col+2] == Filled {
		moves = append(moves, Move{row, col + 2, Left})
	}
	if row > 1 && board[row-1][col] == Filled && board[row-2][col] == Filled {
		moves = append(moves, Move{row - 2, col, Down})
	}
	if row < len(board)-2 && board[row+1][col] == Filled && board[row+2][col] == Filled {
		moves = append(moves, Move{row + 2, col, Up})
	}
	return moves
}

func allMoves(board Board) []Move {
	moves := []Move{}
	for i, row := range board {
		for j, hole := range row {
			if hole == Empty {
				moves = append(moves, fillMoves(board, i, j)...)
			}
		}
	}
	return moves
}

func applyMove(board Board, move Move) Board {
	newBoard := make(Board, len(board))
	for i, row := range board {
		newBoard[i] = make([]Hole, len(row))
		copy(newBoard[i], row)
	}

	if newBoard[move.row][move.col] != Filled {
		panic("Invalid move")
	}
	newBoard[move.row][move.col] = Empty
	switch move.dir {
	case Up:
		if newBoard[move.row-1][move.col] != Filled {
			panic("Invalid move")
		}
		newBoard[move.row-1][move.col] = Empty
		newBoard[move.row-2][move.col] = Filled
	case Down:
		if newBoard[move.row+1][move.col] != Filled {
			panic("Invalid move")
		}
		newBoard[move.row+1][move.col] = Empty
		newBoard[move.row+2][move.col] = Filled
	case Left:
		if newBoard[move.row][move.col-1] != Filled {
			panic("Invalid move")
		}
		newBoard[move.row][move.col-1] = Empty
		newBoard[move.row][move.col-2] = Filled
	case Right:
		if newBoard[move.row][move.col+1] != Filled {
			panic("Invalid move")
		}
		newBoard[move.row][move.col+1] = Empty
		newBoard[move.row][move.col+2] = Filled
	}
	return newBoard
}

func main() {
	board := Board{
		{Blocked, Blocked, Filled, Filled, Filled, Blocked, Blocked},
		{Blocked, Filled, Filled, Filled, Filled, Filled, Blocked},
		{Filled, Filled, Filled, Filled, Filled, Filled, Filled},
		{Filled, Filled, Filled, Empty, Filled, Filled, Filled},
		{Filled, Filled, Filled, Filled, Filled, Filled, Filled},
		{Blocked, Filled, Filled, Filled, Filled, Filled, Blocked},
		{Blocked, Blocked, Filled, Filled, Filled, Blocked, Blocked},
	}

	// moves := fillMoves(board, 3, 3, []Move{})
	// board.print()
	// fmt.Println(moves)
	solutions := appendSolutions([]MoveSequence{}, board, Move{5, 3, Up}, MoveSequence{}, 36)
	fmt.Println()
	fmt.Println(solutions)
}
