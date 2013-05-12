package gonray

import (
	"fmt"
	"reflect"
	"testing"
)

func assertEquality(t *testing.T, a, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		t.Fatalf("Grow shape incorrect.\ngot: %#v\nwant:\n%s", a, b)
	}
}

func Test1Empty(t *testing.T) {
	a := []int{}
	Grow(&a, 0)

	b := []int{}
	assertEquality(t, a, b)
}

func Test2Empty(t *testing.T) {
	a := [][]int{}
	Grow(&a, 0)

	b := [][]int{}
	assertEquality(t, a, b)
}

func Test3Empty(t *testing.T) {
	a := [][][]int{}
	Grow(&a, 0)

	b := [][][]int{}
	assertEquality(t, a, b)
}

func Test1EmptyGrow(t *testing.T) {
	a := []int{}
	Grow(&a, 5)

	b := []int{0, 0, 0, 0, 0}
	assertEquality(t, a, b)
}

func Test2EmptyGrow(t *testing.T) {
	a := [][]int{}
	Grow(&a, 3, 2)

	b := [][]int{nil, nil, {0, 0}}
	assertEquality(t, a, b)
}

func Test3EmptyGrow(t *testing.T) {
	a := [][][]int{}
	Grow(&a, 3, 2, 1)

	b := [][][]int{nil, nil, {nil, {0}}}
	assertEquality(t, a, b)
}

func Test1NoGrow(t *testing.T) {
	a := []int{0, 1, 2}
	Grow(&a, 3)

	b := []int{0, 1, 2}
	assertEquality(t, a, b)
}

func Test2NoGrow(t *testing.T) {
	a := [][]int{{0, 1, 2, 3}}
	Grow(&a, 1, 4)

	b := [][]int{{0, 1, 2, 3}}
	assertEquality(t, a, b)
}

func Test3NoGrow(t *testing.T) {
	a := [][][]int{nil,
		{{0, 1, 2, 3, 4}, nil, nil},
		{nil, {0, 1, 2, 3}, {0, 1, 2}},
	}
	Grow(&a, 3, 3, 3)

	b := [][][]int{nil,
		{{0, 1, 2, 3, 4}, nil, nil},
		{nil, {0, 1, 2, 3}, {0, 1, 2}},
	}
	assertEquality(t, a, b)
}

func Test1Grow(t *testing.T) {
	a := []int{0, 1, 2}
	Grow(&a, 5)

	b := []int{0, 1, 2, 0, 0}
	assertEquality(t, a, b)
}

func Test2Grow(t *testing.T) {
	a := [][]int{{0, 1, 2, 3}}
	Grow(&a, 3, 2)

	b := [][]int{{0, 1, 2, 3}, nil, {0, 0}}
	assertEquality(t, a, b)
}

func Test3Grow(t *testing.T) {
	a := [][][]int{nil,
		{{0, 1, 2, 3, 4}, nil, nil},
		{nil, {0, 1, 2, 3}, {0, 1, 2}},
	}
	Grow(&a, 3, 2, 5)

	b := [][][]int{nil,
		{{0, 1, 2, 3, 4}, nil, nil},
		{nil, {0, 1, 2, 3, 0}, {0, 1, 2}},
	}
	assertEquality(t, a, b)
}

// This example show how slices are grown: slices get nil, and other types get 0.
func ExampleGrow_Simple() {
	a := [][]int{{0, 1, 2, 3}}
	Grow(&a, 3, 2)

	fmt.Printf("%#v\n", a)
	// Output: [][]int{[]int{0, 1, 2, 3}, []int(nil), []int{0, 0}}
}

// This example show how Grow should be used before assignments.
func ExampleGrow_Complete() {
	a := [][]int{}

	Grow(&a, 3, 2)
	a[2][1] = 23
	Grow(&a, 2, 5)
	a[1][4] = 47

	fmt.Printf("%#v\n", a)
	// Output: [][]int{[]int(nil), []int{0, 0, 0, 0, 47}, []int{0, 23}}
}
