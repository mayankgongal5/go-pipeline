package main

import (
	"math"
	"testing"
)

// TestAreaOfCircle tests the areaofcircle function
func TestAreaOfCircle(t *testing.T) {
	tests := []struct {
		radius float64
		want   float64
	}{
		{0, 0},
		{1, 3.14},
		{2, 12.56},
		{5, 78.5},
	}

	tolerance := 0.001

	for _, tt := range tests {
		got := areaofcircle(tt.radius)
		if math.Abs(got-tt.want) > tolerance {
			t.Errorf("areaofcircle(%f) = %f; want %f", tt.radius, got, tt.want)
		}
	}
}

// TestPerimeterOfCircle tests the perimeterofcircle function
func TestPerimeterOfCircle(t *testing.T) {
	tests := []struct {
		radius float64
		want   float64
	}{
		{0, 0},
		{1, 6.28},
		{2, 12.56},
		{5, 31.4},
	}

	tolerance := 0.001

	for _, tt := range tests {
		got := perimeterofcircle(tt.radius)
		if math.Abs(got-tt.want) > tolerance {
			t.Errorf("perimeterofcircle(%f) = %f; want %f", tt.radius, got, tt.want)
		}
	}
}
