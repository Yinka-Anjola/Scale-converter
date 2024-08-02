package main

import "fmt"

func main() {
        var kilogram, pound, gram float64
        var conversionType string

        fmt.Println("Enter kilogram: ")
        fmt.Scanln(&kilogram)

        fmt.Println("Enter conversion type (pound or gram): ")
        fmt.Scanln(&conversionType)

        switch conversionType {
        case "pound":
                pound = kilogram * 2.205
                fmt.Println("kilograms is equal to", pound, "pounds")
        case "gram":
                gram = kilogram * 1000
                fmt.Println("kilograms is equal to", gram, "grams")
        default:
                fmt.Println("Invalid conversion type")
        }
}
