/* 
This application finds a solution to the Eight Queens Puzzle without using brute force. https://en.wikipedia.org/wiki/Eight_queens_puzzle 
Pass the column and row of the first point separated by a space to get the other 7 points. EX: ./queen_solver 4 5
Created by Chandler Mayo - Copyright 2016 Chandler Mayo - http://ChandlerMayo.com 
*/

package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
)

func main() {
	if (len(os.Args) < 3) { // The user did not specify the arguments. 
		fmt.Println("Missing required arguments. Please specify the row and column numbers.")
		fmt.Println("Press 'Enter' to continue...")
		bufio.NewReader(os.Stdin).ReadBytes('\n') 
		os.Exit(1) 
	}
	var err error
	column_s := os.Args[1]
	row_s := os.Args[2]
	fmt.Println("VVVVVVVVVVVVVVVVVVVVVVVVVV\nEight Queens Puzzle Solver\n^^^^^^^^^^^^^^^^^^^^^^^^^^")
	row, err := strconv.Atoi(row_s) // Convert to int.
    if err != nil {
        fmt.Println(err)
        os.Exit(2)
    }
    column, err := strconv.Atoi(column_s)  // Convert to int.
    if err != nil {
        fmt.Println(err)
        os.Exit(2)
    }
    if (column > 8 || column < 1 || row > 8 || row < 1) { // Does the point land on an 8*8 grid?
    	fmt.Println("Value outside expected range. Please enter a value 1-8.")
    	os.Exit(2)
    }
	var qfield [9]int
	qfield[column] = row
	for i := (column-1); i > 0; i-- { // Other points that work as a solution can be determined from the row passed as an argument.
		move := 0
		if (qfield[i+1] == 1) {
			move = -7
		} else if (qfield[i+1] == 2) {
			move = -3
		} else if (qfield[i+1] == 3) {
			move = -3
		} else if (qfield[i+1] == 4) {
			move = -3
		} else if (qfield[i+1] == 5) {
			move = 2
		} else if (qfield[i+1] == 6) {
			move = 2
		} else if (qfield[i+1] == 7) {
			move = 6
		} else if (qfield[i+1] == 8) {
			move = 6
		} 
		qfield[i] = (qfield[i+1] - move)
	}
	for i := (column+1); i < 9; i++ {
		move := 0
		if (qfield[i-1] == 1) {
			move = 6
		} else if (qfield[i-1] == 2) {
			move = 6
		} else if (qfield[i-1] == 3) {
			move = 2
		} else if (qfield[i-1] == 4) {
			move = 2
		} else if (qfield[i-1] == 5) {
			move = -3
		} else if (qfield[i-1] == 6) {
			move = -3
		} else if (qfield[i-1] == 7) {
			move = -3
		} else if (qfield[i-1] == 8) {
			move = -7
		} 
		qfield[i] = (qfield[i-1] + move)
	}
	if ((qfield[1] == 1 || qfield[1] == 8) && (qfield[1] == 1 || qfield[1] == 8)) { // The first solution didn't work so it runs again backwards for the result. 
		for i := (column-1); i > 0; i-- {
			move := 0
			if (qfield[i+1] == 1) {
				move = 6
			} else if (qfield[i+1] == 2) {
				move = 6
			} else if (qfield[i+1] == 3) {
				move = 2
			} else if (qfield[i+1] == 4) {
				move = 2
			} else if (qfield[i+1] == 5) {
				move = -3
			} else if (qfield[i+1] == 6) {
				move = -3
			} else if (qfield[i+1] == 7) {
				move = -3
			} else if (qfield[i+1] == 8) {
				move = -7
			} 
			qfield[i] = (qfield[i+1] + move)
		}
		for i := (column+1); i < 9; i++ {
			move := 0
			if (qfield[i-1] == 1) {
				move = -7 
			} else if (qfield[i-1] == 2) {
				move = -3
			} else if (qfield[i-1] == 3) {
				move = -3 
			} else if (qfield[i-1] == 4) {
				move = -3
			} else if (qfield[i-1] == 5) {
				move = 2
			} else if (qfield[i-1] == 6) {
				move = 2
			} else if (qfield[i-1] == 7) {
				move = 6
			} else if (qfield[i-1] == 8) {
				move = 6
			} 
			qfield[i] = (qfield[i-1] - move)
		}
	}
	fmt.Println("\nResult:\n")
	for i := 1; i < 9; i++ { // Print out the column and row info. 
		fmt.Println(fmt.Sprintf("Column: %d, Row: %d", i, qfield[i]))
	}
	fmt.Println("")
	for i := 1; i < 9; i++ { // Print out an 8*8 grid with the solution. 
		for u := 1; u < 9; u++ {
			if (qfield[u] == i) {
				fmt.Print("■")
			} else {
				fmt.Print("□")
			}
		}
		fmt.Println("")
	}
	os.Exit(0) // Done.
}

