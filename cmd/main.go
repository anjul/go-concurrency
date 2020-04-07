package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Press Enter key at terminal to stop the program!")
	// uncomment one function call at a time to see the working demo
	// demoGoRoutines()
	// demoGoChannels()
	// demoSelect()

	var input string
	fmt.Scanln(&input)
}

func demoGoRoutines() {
	for i := 0; i < 10; i++ {
		go printNumbers(i)
	}
}

// Usage of goroutines
func printNumbers(n int) {

	for i := 0; i < 10; i++ {
		fmt.Println(n, " : ", i)

		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

// Usage of goChannels
func demoGoChannels() {
	var c = make(chan string) // Creating a channel
	go pinger(c)
	go ponger(c)
	go printer(c)
}

// An example of Specifying a direction on a channel, thus this function is restricted to send a string
func pinger(c chan<- string) {
	for {
		c <- "ping"
	}
}

// this function can send & receive a string
func ponger(c chan string) {
	for {
		c <- "pong"
	}
}

// An example of Specifying a direction on a channel, thus this function is restricted to receive a string
func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg) // Or fmt.Println(<- c)
		time.Sleep(time.Second * 1)
	}
}

// Usage of Select
func demoSelect() {

	var c1 = make(chan string)
	var c2 = make(chan string)

	go func() {
		for {
			c1 <- "Message from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "Message from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			select { // Observe how select statement works as switch statement for channels
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			case <-time.After(time.Second): // creating channel without storing it in a variable
				fmt.Println("timeout")
			}
		}
	}()
}
