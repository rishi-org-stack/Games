package game

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Game struct {
	board         [][]string
	result        string
	usermoves     [][]int
	computermoves [][]int
}

var Symbols = map[int]string{
	1: "X",
	0: "O",
}

func StartGame() *Game {
	var g = &Game{}
	g.board = [][]string{
		{"|", "|", "|"},
		{"|", "|", "|"},
		{"|", "|", "|"},
	}
	g.result = ""
	g.usermoves = make([][]int, 0)
	g.computermoves = make([][]int, 0)
	return g
}

func (g *Game) UpdateBoard(loc []int) {
	loc1 := SelectARandInt()
	g.usermoves = append(g.usermoves, loc)
	g.computermoves = append(g.computermoves, loc1)
	if len(g.board[loc[0]][loc[1]]) == 1 { // g.board[loc[0]][loc[1]]  != "|O" || g.board[loc[0]][loc[1]] != "|X"{
		g.board[loc[0]][loc[1]] = g.board[loc[0]][loc[1]] + "O"

	} else if len(g.board[loc[0]][loc[1]]) > 1 {
		fmt.Println("Invalid move, place already taken")
		return
	}
	if len(g.board[loc1[0]][loc1[1]]) == 1 { // g.board[loc1[0]][loc1[1]] != "|X" || g.board[loc1[0]][loc1[1]] != "|O" {
		g.board[loc1[0]][loc1[1]] = g.board[loc1[0]][loc1[1]] + "X"

	} else {
		g.UpdateBoard(loc)
	}

}

func (g *Game) PrintBoard() {
	for _, line := range g.board {
		for _, row := range line {
			fmt.Printf("%v  ", row)
		}
		fmt.Println()
	}
}

func (g *Game) SelectWiner() string {
	user := 0
	Computer := 0
	res := "Draw"
	for i := 1; i < 3; i++ {
		if g.usermoves[i][0] == g.usermoves[i-1][0]+1 || g.usermoves[i][1] == g.usermoves[i-1][1]+1 {
			user++
		}
		if g.computermoves[i][0] == g.computermoves[i-1][0]+1 || g.computermoves[i][1] == g.computermoves[i-1][1]+1 {
			Computer++
		}
	}
	if user == 3 {
		res = "user won"
	} else if Computer == 3 {
		res = "computer won"
	} else {

	}
	return res
}

func userinput() []int {
	var s, s2 string
	fmt.Scan(&s)
	fmt.Scan(&s2)
	fmt.Println()
	x, _ := strconv.Atoi(string(s2[0]))
	y, _ := strconv.Atoi(string(s2[2]))
	ar := []int{x, y}
	return ar
}
func state() {
	fmt.Println()
	fmt.Println("tictactoe>")
	fmt.Println()
}
func Play() {
	var what string
	
	result := []int{0,0}
	for {
		fmt.Scan(&what)
		if what == "continue" {
			continue
		}else if what =="End"{
			break
		}else if what == "Scoreboard"{
			fmt.Printf("\nuser -> %v\n",result[0])
			fmt.Printf("\ncomputer-> %v\n",result[1])
			break
		}
		state()
		g := StartGame()
		g.PrintBoard()

		for i := 0; i < 9; i++ {
			state()
			ar := userinput()
			g.UpdateBoard(ar)
			g.PrintBoard()

		}
		fmt.Println(g.SelectWiner())
		if g.SelectWiner() =="user won"{
			result[0]++
		}else if g.SelectWiner() =="computer won" {
			result[1]++
		}
		// if what == "continue" {
		// 	continue
		// }else if what =="End"{
		// 	break
		// }else if what == "Scoreboard"{
		// 	fmt.Printf("\nuser -> %v\n",result[0])
		// 	fmt.Printf("\ncomputer-> %v\n",result[1])
		// 	break
		// }
	}
	

}

func SelectARandInt() []int {
	rand.Seed(time.Now().UTC().UnixNano())
	val := rand.Intn(3)
	ar := []int{val, rand.Intn(3)}
	return ar
}
func selectrand() int {
	rand.Seed(time.Now().UTC().UnixNano())
	val := rand.Intn(2)
	return val
}
