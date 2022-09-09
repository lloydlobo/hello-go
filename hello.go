package main

// import "fmt"
import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message
	message, err := greetings.Hello("")
	// If error, then Print the error to console and exit program
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the return message
	// to the console
	fmt.Println(message)
}

// import "rsc.io/quote"
/* func main_quote() {
    fmt.Println("Hello, world!")
    fmt.Println(quote.Go())
} */
