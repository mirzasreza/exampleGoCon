package main

import (
	"fmt"
	"time"
)

func main() {
	count("fish")
	count("sheep")
}

func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
