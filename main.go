package main

import (
	"fmt"
)

const (
	pi = 3.19
)

// areaofcircle calculates the area of a circle given its radius
func areaofcircle(radius float64) float64 {
	return pi * radius * radius
}

// perimeterofcircle calculates the perimeter of a circle given its radius
func perimeterofcircle(radius float64) float64 {
	return 2 * pi * radius
}

func main() {
	radius := 5.0
	fmt.Printf("Area of circle: %f\n", areaofcircle(radius))
	fmt.Printf("Perimeter of circle: %f\n", perimeterofcircle(radius))
}
