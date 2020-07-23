package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const greet = `Welcome To ProbaNum :)
A Game Where You Choose A Number And Check Its Probability Among The Go Num Generator
Now Please Enter A Number As An Argument`

func main() {
	if len(os.Args) < 2 {
		fmt.Println(greet)
		return
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Plese Enter Positive Numbers Only")
		return
	}
	if num < 0 {
		fmt.Println("Please Enter Positive Numbers Only, Negative Numbers Does Not Work")
		return
	}
	var track int
	rand.Seed(time.Now().UnixNano())
	fmt.Println("The Number Chosen Is:", num)
	duration := time.Now().UnixNano()
	for {
		rnd := rand.Intn(num + 1)
		track++
		if num == rnd {
			duration = time.Now().UnixNano() - duration
			fmt.Printf("Found After %d Gesses", track)
			fmt.Println("\nThe Time It Took The Program To Guess The Number Is:", float64(duration)/1000000, "Millieconds")
			return
		}
	}
}
