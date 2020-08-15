package main

import (
	"fmt"
	"time"
)

func speak(people string, pharse string, times int, c chan<- string) {
	for i := 0; i < times; i++ {
		str := fmt.Sprintf("%s : %s (%d)\n", people, pharse, i+1)
		c <- str
	}
}

func main() {
	channel := make(chan string, 1)
	go speak("Maria", "Ola, como está?", 1, channel)
	go speak("João", "Estou bem e você?", 2, channel)
	go speak("Maria", "Estou Ótima!", 1, channel)

	time.Sleep(time.Second * 3)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
