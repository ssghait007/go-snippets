package main

import (
	"fmt"
	"time"
)

const rabbitSpeed = 50
const turtleSpeed = 150

const sleepFrames = 80
const raceFrames = 100

func main() {

	h := make(chan string)
	t := make(chan string)

	// Rabbit starts running
	go func() {
		for i := 1; i < sleepFrames; i++ {
			s := getRabbitPosition("-", i, "")
			h <- s
			time.Sleep(time.Millisecond * rabbitSpeed)
		}
		for i := 1; i <= raceFrames-sleepFrames; i++ {
			h <- getRabbitPosition("-", sleepFrames, "i'ma take nap")
			time.Sleep(time.Millisecond * rabbitSpeed)
		}
	}()

	// Turtle starts running
	go func() {
		t <- "t"
		for i := 1; i < raceFrames; i++ {
			s := getTurtleLocation("-", i)
			t <- s
			time.Sleep(time.Millisecond * turtleSpeed)
		}
	}()

	hmsg := ""
	tmsg := ""

	// Check positions
	for i := 1; i < raceFrames*2; i++ {
		select {
		case H := <-h:
			hmsg = H
		case T := <-t:
			tmsg = T
		}
		// Clear screen and live locations
		fmt.Print("\033[H\033[2J")
		fmt.Printf("%s \n%s", hmsg, tmsg)
	}

	fmt.Println("🏁 Turtle wins")
}

func getRabbitPosition(st string, count int, msg string) string {
	out := ""
	for i := 0; i < count; i++ {
		out = out + st
	}
	return out + "🐰" + msg
}

func getTurtleLocation(st string, count int) string {
	out := ""
	for i := 0; i < count; i++ {
		out = out + st
	}
	return out + "🐢"
}
