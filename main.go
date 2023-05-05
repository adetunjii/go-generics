package main

import (
	"fmt"
)

type Ordered interface {
	~int | ~float64 | ~string
}

func main() {

	i := []int{1, 2, 3, 4, 5, 6, 6}
	j := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0}

	fmt.Println(Max(i))
	fmt.Println(Max(j))

}

// ~ means for every underlying type
func Max[T Ordered](values []T) (T, error) {
	if len(values) == 0 {
		var zero T
		return zero, fmt.Errorf("max of empty slice")
	}

	m := values[0]
	for _, v := range values[1:] {
		if m > v {
			m = v
		}
	}

	return m, nil
}
