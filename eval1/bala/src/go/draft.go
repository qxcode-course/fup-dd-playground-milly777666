package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2 float64
	
	// Read the coordinates
	fmt.Scan(&x1, &y1, &x2, &y2)
	
	// Calculate distance: d = sqrt((x2-x1)^2 + (y2-y1)^2)
	dx := x2 - x1
	dy := y2 - y1
	distance := math.Sqrt(dx*dx + dy*dy)
	
	// Print with 2 decimal places
	fmt.Printf("%.2f\n", distance)
}
