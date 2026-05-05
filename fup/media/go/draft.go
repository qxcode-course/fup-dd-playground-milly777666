package main

import "fmt"

func main() {
    var a, b int
    _, err := fmt.Scanln(&a)
    if err != nil {
        return
    }
    _, err = fmt.Scanln(&b)
    if err != nil {
        return
    }
    avg := float64(a+b) / 2.0
    fmt.Printf("%.1f\n", avg)
}
