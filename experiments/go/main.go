package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
	"regexp"
	"strconv"
)

type Board struct {
	empty string

	height int
	width int

	board []string
	row0 []string
	row1 []string
	row2 []string

	mark1 string
	mark2 string
}


func main() {

	ticTacBoard := createTicTacToeBoard()
	printBoard(ticTacBoard)
	fmt.Println()

	runBoardEngine(ticTacBoard)

}

func runBoardEngine(board Board) {
	reader := bufio.NewReader(os.Stdin)
	currentPlayer := 1
	validMoves := 0
	exitFlag := "Q"

	for {
		fmt.Printf("Player %d Enter row and col to place your mark: ", currentPlayer)


		playerMoveInput, playerMoveInputErr := reader.ReadString('\n')

		if playerMoveInputErr != nil {

			break

		}

		cleanInput := cleanPlayerInput(playerMoveInput)


		if len(cleanInput) < 2 {

			if cleanInput[0] == exitFlag {

				printMessage("Thanks for playing!")
				break

			}

			printMessage("Invalid row col pair.")

		} else {

			row, rowErr := strconv.Atoi(cleanInput[0])
			col, colErr := strconv.Atoi(cleanInput[1])


			if rowErr != nil || colErr != nil {

				break

			} else {

				if row >= board.width || col >= board.height {

					printMessage("Invalid range for row col")

				} else {

					if playerMove(board, row, col, currentPlayer) {

						validMoves++
						currentPlayer = switchPlayer(currentPlayer)

						if validMoves > 4 {
							printMessage("Time to check for victory!")
						}

					}

				}
			}

		}
	}
}

func switchPlayer (player int) int {

	if player == 1 {

		return 2

	} else {

		 return 1

	}
}

func cleanPlayerInput (input string) []string {

	numOnlyRegEx := regexp.MustCompile("^Q|[0-9]+$")
	dirtyInput := strings.Split(input, " ")
	cleanString := ""

	for i := 0; i < len(dirtyInput); i++ {
		if numOnlyRegEx.MatchString(strings.TrimSpace(dirtyInput[i])) {
			cleanString += dirtyInput[i] + " "
		}
	}

	cleanString = strings.TrimSpace(cleanString)
	return strings.Split(cleanString, " ")
}

func createTicTacToeBoard() (Board) {
	board := Board{}
	board.empty = "_"
	board.width = 3
	board.height = 3
	board.board = make([]string, board.width * board.height)

	for i := 0; i < len(board.board); i++ {
		board.board[i] = board.empty
	}

	board.row0 = board.board[:3]
	board.row1 = board.board[3:6]
	board.row2 = board.board[6:]

	board.mark1 = "X"
	board.mark2 = "O"

	return board
}

func printBoard(board Board) {
	printBoardRow(board.row0)
	printBoardRow(board.row1)
	printBoardRow(board.row2)
}

func printBoardRow(row []string) {
	for i := 0; i < len(row); i++ {
		fmt.Print(row[i], " ")
	}
	fmt.Println()
}

func playerMove(board Board, row int, col int, player int) bool {
	location := (row * board.width) + col
	if iEmptyLocation(board, location) {
		placeMark(board, location, player)
		return true
	} else {
		return false
	}
}

func iEmptyLocation(board Board, location int) bool {
	return board.board[location] == board.empty
}

func placeMark(board Board, location int, player int)  {

	var mark string
	if player == 1 {
		mark = board.mark1
	} else {
		mark = board.mark2
	}

	board.board[location] = mark

	fmt.Println()
	printBoard(board)
	fmt.Println()
}

func printMessage(message string) {
	fmt.Println()
	fmt.Println(message)
	fmt.Println()
}