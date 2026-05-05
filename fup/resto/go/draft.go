package main

import "fmt"

func main() {
    var n, d int
    if _, err := fmt.Scanln(&n); err != nil {
        return
    }
    if _, err := fmt.Scanln(&d); err != nil {
        return
    }

    q := n / d
    r := n % d
    fmt.Printf("%d %d\n", q, r)
}
