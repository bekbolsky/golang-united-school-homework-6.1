package golang_united_school_homework

import (
	"math"
)

// Circle is a struct that contains radius and has methods for calculating perimeter and area.
type Circle struct {
	Radius float64
}

func (c *Circle) CalcPerimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c *Circle) CalcArea() float64 {
	return math.Pi * c.Radius * c.Radius
}
