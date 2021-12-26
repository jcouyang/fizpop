package maps

import "constraints"

func Filter[K constraints.Ordered, A any](fn func (k K,v A) bool) func (a  map[K]A) map[K]A {
	return func (a map[K]A) (b map[K]A) {
		b = make(map[K]A)
		for i,j := range a {
			if fn(i, j) {
				b[i] = j
			}
		}
		return
	}
}