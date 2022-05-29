package golang_united_school_homework

import (
	"math"
)

// Triangle is a struct that contains side and has methods for calculating perimeter and area.
type Triangle struct {
	Side float64
}

func (t *Triangle) CalcPerimeter() float64 {
	return 3 * t.Side
}

func (t *Triangle) CalcArea() float64 {
	return math.Sqrt(3) * t.Side * t.Side / 4
}
