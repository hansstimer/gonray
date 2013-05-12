gonray
===

Go package provides a single function for growing an n-dimensional slice.

Usage
---

Grow grows an n-dimensional slice a non-uniformilly. This gives you a
sparse array that is only as large as required to hold the needed
dimensions. Any items between the needed dimension and the last
item are set to 0.

For instance:
```go
a := [][]int{{0, 1, 2, 3}}
 gonray.Grow(&a, 2, 2) // {{0, 1, 2, 3}, {0, 0}}
```

will grow array a to [2][2], but doesn't shrink existing slices.
The new values are set to 0.

Notes
---
* gonray can't add dimensions
* gonray is slow because it uses the reflection library