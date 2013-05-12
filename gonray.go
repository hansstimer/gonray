// Package gonray provides a single function for growing an n-dimensional
// slice.
package gonray

import (
	"reflect"
)

// Grow grows an n-dimensional slice a non-uniformilly. This gives you a
// sparse array that is only as large as required to hold the needed
// dimensions. Any items between the needed dimension and the last
// item are set to 0.
//
// For instance:
//	a := [][]int{{0, 1, 2, 3}}
//  gonray.Grow(&a, 2, 2) // {{0, 1, 2, 3}, {0, 0}}
//
// will grow array a to [2][2], but doesn't shrink existing slices.
// The new values are set to 0.
//
func Grow(a interface{}, pos ...int) interface{} {
	v := reflect.ValueOf(a).Elem()
	p := pos[0]
	if p > v.Len() {
		cap := v.Cap()
		if p >= cap {
			vv := reflect.MakeSlice(v.Type(), p, p*2)
			reflect.Copy(vv, v)
			v.Set(vv)
		}
		v.SetLen(p)
	}

	if len(pos) > 1 {
		vi := v.Index(p - 1)
		inf := Grow(vi.Addr().Interface(), pos[1:]...)
		v.Index(p - 1).Set(reflect.ValueOf(inf))
	}

	return v.Interface()
}
