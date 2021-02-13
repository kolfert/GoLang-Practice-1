package main

import (
	"fmt"

	"rsc.io/quote"

	"time"
)

func main() {
	//fmt.Println("\nHello beautiful blockchain world!\n")
	fmt.Println("")
	fmt.Println(time.Now(), "---", quote.Go()) // Adding \n here causes complaints
	fmt.Println("")                            // pre-compiler omplalins when adding a new line, so we just Println with nothing.
	//fmt.Println(quote.Go(), "\n")
}
