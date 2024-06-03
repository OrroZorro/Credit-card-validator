package main

import (
	"fmt"
	"unicode"
)


func validateCreditCard(number string) bool {
    var sum int
    var alternate bool

 
    for i := len(number) - 1; i >= 0; i-- {
        n := int(number[i] - '0')


        if alternate {
            n *= 2
            if n > 9 {
                n = (n % 10) + 1 
            }
        }

        sum += n
        alternate = !alternate
    }

    return (sum % 10) == 0
}


func isNumeric(s string) bool {
    for _, c := range s {
        if !unicode.IsDigit(c) {
            return false
        }
    }
    return true
}

func main() {
    var cardNumber string
    fmt.Print("Enter credit card number: ")
    fmt.Scan(&cardNumber)

    if !isNumeric(cardNumber) {
        fmt.Println("Invalid input. Please enter only numeric characters.")
        return
    }

    if validateCreditCard(cardNumber) {
        fmt.Println("Credit card number is valid.")
    } else {
        fmt.Println("Credit card number is invalid.")
    }
}
