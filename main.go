package main

import (
    "fmt"
    "time"
    "math"
    "math/rand"
)

func abs(x int) int {
    if x <0 {
        return -x
    }
    return x
}

func random(max int) int {
    return rand.Intn(max)
}

func initialSolution(n int) []int{

    board := make([]int, n)
    for i:= 0; i<n; i++ {
        board[i] = i
    }
    return board
}

func printBoard(board []int) {
    board_len := len(board)
    boardMatrix := createMatrix(board_len)
    for i:=0; i < board_len; i++ {
        boardMatrix[i][board[i]] = 1
        fmt.Println(boardMatrix[i])
    }
    //fmt.Println(boardMatrix)
}

func createMatrix(n int) [][]int{
    matrix := make([][]int, n)
    for i:=0; i<n; i++ {
        matrix[i] = make([]int, n)
    }
    return matrix
}

func calculateEnergy(board []int) int {
    board_len := len(board)

    collisions := 0
    for i:=0; i < board_len; i++ {
        fmt.Println(i, " iteration")
        for j:=i+1; j < board_len; j++ {
            diff := j - i
            //if board[j] == abs(diff - board[i]) || board[j] == abs(diff + board[i]) {
            if abs(board[j] - board[i]) == diff {
                //fmt.Printf(" j=%d - i=%d\n", j, i)
                //fmt.Printf("board[i]=%d , board[j]=%d, diff=%d\n", board[i], board[j], diff)
                collisions++
            } else {
                //fmt.Printf("not match i=%d , j=%d, diff=%d\n", board[i], board[j], diff)
            }
        }
    }
    return collisions
}

func generateNewSolution(board []int) []int {
    board_len := len(board)
    newBoard := make([]int, board_len)
    copy(newBoard, board)
    x := rand.Intn(board_len)
    y := rand.Intn(board_len)

    for x == y {
        x = rand.Intn(board_len)
        y = rand.Intn(board_len)
    }


    fmt.Println("rand", x, y)
    newBoard[x], newBoard[y] = newBoard[y], newBoard[x]

    return newBoard
}

func updateTemperature(t float64) float64{
    t = 0.99 * t
    return t
}

func main() {
    rand.Seed(time.Now().Unix())
    board := initialSolution(128)
    //var temperature float64 = float64(calculateEnergy(board))/0.2
    var temperature float64 = float64(35)

    fmt.Println(temperature)
    printBoard(board)

    fmt.Println("collisions: ", calculateEnergy(board))

    for calculateEnergy(board) > 0 {
        fmt.Println("Generate New Solution")

        newBoard := generateNewSolution(board)

        delta := calculateEnergy(board) - calculateEnergy(newBoard)
        probability := math.Exp(float64(delta)/temperature)
        randomNumber := rand.Float64()
        fmt.Println("probability rand ", probability)
        fmt.Println("float64 rand ", randomNumber)
        if calculateEnergy(newBoard) < calculateEnergy(board) {
            board = newBoard
            printBoard(board)

            fmt.Println("collisions: ", calculateEnergy(board))
        } else if(randomNumber <= probability) {
            board = newBoard
            printBoard(board)
        }
        temperature = updateTemperature(temperature)
    }

    printBoard(board)

}
