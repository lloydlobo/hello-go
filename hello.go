package main

// import "fmt"
import (
    "fmt"
   "example.com/greetings"
)

func main() {
    // Get a greeting message and print it
    message := greetings.Hello("Gladys")

    fmt.Println(message)
}

// import "rsc.io/quote"
/* func main_quote() {
    fmt.Println("Hello, world!")
    fmt.Println(quote.Go())
} */
