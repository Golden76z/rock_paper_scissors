package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	var win, lose, draw int
	fmt.Println("-------------------------------------------------------------------------------")
	fmt.Println(`Welcome to Rock-Paper-Scissors game, type "rules" to get the rules of the game `)
	fmt.Println(`                            Enter 'Stop' to exit								`)
	fmt.Println("-------------------------------------------------------------------------------")
	first := true
	for {
		if !first {
			fmt.Println("Press 'Stop' to stop")
		}
		s := bufio.NewScanner(os.Stdin)
		fmt.Print("Player 1: ")
		s.Scan()
		text := strings.ToLower(s.Text())

		if text == "stop" {
			break
		} else if text == "rules" {
			fmt.Println("-------------------------------------------------------------------------------")
			fmt.Println("              3 Possibles entries: 'Rock', 'Paper' and 'Scissors' 				")
			fmt.Println("                          'Scissors beats 'Paper'")
			fmt.Println("                            'Paper beats 'Rock'")
			fmt.Println("                           'Rock beats 'Scissors'")
			fmt.Println("-------------------------------------------------------------------------------")
		} else {
			// Generate a random move for the ia
			rand.Seed(time.Now().UnixNano())
			randomNumber := rand.Intn(3) + 1
			
			var enemy string
			switch randomNumber {
			case 1:
				enemy = "Scissors"
			case 2:
				enemy = "Rock"
			case 3:
				enemy = "Paper"		
			}

			if enemy == "Scissors" {
				if text == "paper" {
					fmt.Println("                 Enemy played 'Scissors' -> You lose !")
					fmt.Println("-------------------------------------------------------------------------------")
					lose++
				} else if text == "scissors" {
					fmt.Println("                 Enemy played 'Scissors' -> It's a draw")
					fmt.Println("-------------------------------------------------------------------------------")
					draw++
				} else if text == "rock" {
					fmt.Println("                 Enemy played 'Scissors' -> You win !")
					fmt.Println("-------------------------------------------------------------------------------")
					win++
				}
			} else if enemy == "Paper" {
				if text == "paper" {
					fmt.Println("                 Enemy played 'Paper' -> It's a draw")
					fmt.Println("-------------------------------------------------------------------------------")
					draw++
				} else if text == "scissors" {
					fmt.Println("                 Enemy played 'Paper' -> You win !")
					fmt.Println("-------------------------------------------------------------------------------")
					win++
				} else if text == "rock" {
					fmt.Println("                 Enemy played 'Paper' -> You lose !")
					fmt.Println("-------------------------------------------------------------------------------")
					lose++
				}
			} else {
				if text == "paper" {
					fmt.Println("                 Enemy played 'Rock' -> You win !")
					fmt.Println("-------------------------------------------------------------------------------")
					win++
				} else if text == "scissors" {
					fmt.Println("                 Enemy played 'Rock' -> You lose !")
					fmt.Println("-------------------------------------------------------------------------------")
					lose++
				} else if text == "paper" {
					fmt.Println("                 Enemy played 'Rock' -> It's a draw")
					fmt.Println("-------------------------------------------------------------------------------")
					draw++
				}
			}
			if text != "rock" && text != "scissors" && text != "paper" {
				fmt.Println("Wrong input, valid inputs: 'Rock', 'Paper', 'Scissors'")
				fmt.Println("-------------------------------------------------------------------------------")
			}
		}
	}
	fmt.Println("-------------------------------------------------------------------------------")
	fmt.Println("                       Number of wins:", win)
	fmt.Println("                       Number of loses:", lose)
	fmt.Println("                       Number of draw:", draw)
	fmt.Println("-------------------------------------------------------------------------------")
}