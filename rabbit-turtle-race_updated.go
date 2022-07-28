package main

import (
	"fmt"
	"time"
)

// animal speeds
const rabbitSpeed = 100
const turtleSpeed = 300
const talkSpeed = 500

// Durations of talk
const talkFrames = 16

// Time at which rabbit sleeps
const sleepAtFrame = 60

// Duration of race
const raceFrames = 100

func rabbitTurtle() {

	// Make position reporting channels
	r := make(chan string)
	t := make(chan string)

	// Rabbit starts running
	go func() {
		// Rabbit speaks
		for i := 1; i <= 4; i++ {
			s := getPosition("-", "🐰", 0, " You are slow")
			r <- s
			time.Sleep(time.Millisecond * talkSpeed)
		}
		// Rabbit listen
		for i := 1; i <= 4; i++ {
			s := getPosition("-", "🐰", 0, "")
			r <- s
			time.Sleep(time.Millisecond * talkSpeed)
		}
		// Rabbit speaks
		for i := 1; i <= 4; i++ {
			s := getPosition("-", "🐰", 0, " Lets race !")
			r <- s
			time.Sleep(time.Millisecond * talkSpeed)
		}
		// Rabbit speaks
		for i := 1; i <= 4; i++ {
			s := getPosition("-", "🐰", 0, " Ready, go!")
			r <- s
			time.Sleep(time.Millisecond * talkSpeed)
		}
		// Rabbit runs
		for i := 1; i <= sleepAtFrame; i++ {
			s := getPosition("-", "🐰", i, "")
			r <- s
			time.Sleep(time.Millisecond * rabbitSpeed)
		}
		// Rabbit takes nap
		for i := 1; i <= raceFrames-sleepAtFrame-talkFrames; i++ {
			r <- getPosition("-", "🐰", sleepAtFrame, "i'ma take nap")
			time.Sleep(time.Millisecond * rabbitSpeed)
		}
	}()

	// Turtle starts running
	go func() {
		// Turtle listens
		for i := 1; i <= 4; i++ {
			s := getPosition("-", "🐢", 0, "")
			t <- s
			time.Sleep(time.Millisecond * talkSpeed)
		}
		// Turtle speaks
		for i := 1; i <= 4; i++ {
			s := getPosition("-", "🐢", 0, " I am not")
			t <- s
			time.Sleep(time.Millisecond * talkSpeed)
		}
		// Turtle listens
		for i := 1; i <= 4; i++ {
			s := getPosition("-", "🐢", 0, "")
			t <- s
			time.Sleep(time.Millisecond * talkSpeed)
		}
		// Turtle speaks
		for i := 1; i <= 4; i++ {
			s := getPosition("-", "🐢", 0, " Ready, go!")
			t <- s
			time.Sleep(time.Millisecond * talkSpeed)
		}

		// Turtle walks slowly
		for i := 1; i <= raceFrames-talkFrames; i++ {
			s := getPosition("-", "🐢", i, "")
			t <- s
			time.Sleep(time.Millisecond * turtleSpeed)
		}
	}()

	rmsg := ""
	tmsg := ""

	// Check positions
	for i := 1; i < raceFrames*2; i++ {
		select {
		case R := <-r:
			rmsg = R
		case T := <-t:
			tmsg = T
		}
		// Clear screen and live locations
		fmt.Print("\033[H\033[2J")
		fmt.Printf("%s \n%s", rmsg, tmsg)
	}

	// Closing the channels is very important
	close(r)
	close(t)

	fmt.Println("🏁 I win")
}

func getPosition(st string, icon string, count int, msg string) string {
	out := ""
	for i := 0; i < count; i++ {
		out = out + st
	}
	return out + icon + msg
}
