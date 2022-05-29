package golang_united_school_homework

// Rectangle is a struct that contains height and weight and has methods for calculating perimeter and area.
type Rectangle struct {
	Height, Weight float64
}

func (r *Rectangle) CalcPerimeter() float64 {
	return 2 * (r.Height + r.Weight)
}

func (r *Rectangle) CalcArea() float64 {
	return r.Height * r.Weight
}
