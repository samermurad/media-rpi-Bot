package main

import (
	"fmt"
	"testing"
	"time"

	"samermurad.com/piBot/spinner"
)

func Test1(t *testing.T) {
	sp := spinner.NewTmSpinner(68386493, "Timer Spinner")
	fmt.Println("WTF")
	chaa := make(chan bool)
	result := true
	for i := 0; i < 11; i++ {
		go sp.Progress(1, chaa)
		result = <-chaa
		<-time.After(2 * time.Second)
		if !result {
			t.Error("Bummer")
		}
	}
}
