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

	// A slice of name.
	names := []string{"Georgio", "Giovanni", "Fred"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	// If error, then Print the error to console and exit the program
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the return map of
	// messages to the console
	fmt.Println(messages)
}

// import "rsc.io/quote"
/* func main_quote() {
    fmt.Println("Hello, world!")
    fmt.Println(quote.Go())
} */
