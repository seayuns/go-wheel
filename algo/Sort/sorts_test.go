package sorts

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	a := []int{1, 3, 2, 5, 4}

	arrs, n := BubbleSort(a)

	fmt.Printf("new arrs is %v , sort nums is %v \n", arrs, n)
}
