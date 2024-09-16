package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	resetTimer := make(chan bool)
	var px, py int

	d := 5 * time.Minute // duration
	timer := time.NewTimer(d)

	go func() {
		rem := d

		for {
			select {
			case <-time.After(1 * time.Second):
				rem -= time.Second
				// time since last movement
				fmt.Printf("\rt: %v", rem.Round(time.Second))
				if rem <= 0 {
					fmt.Println("\n5 mins passed without movement!")
					return
				}
			case <-resetTimer:
				fmt.Printf("\r .......")
				rem = d
			}
		}
	}()

	go func() {
		for {
			x, y := robotgo.Location()

			if x != px || y != py {
				resetTimer <- true
				px, py = x, y

				// Avoid too much polling
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	<-timer.C
}
