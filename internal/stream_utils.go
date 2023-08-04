package internal

// MapNotNullTo Convert a slice from one type to another.
// The result slice will contain all the non-nil returned values of the given func
func MapNotNullTo[I any, O any](data []I, f func(I) *O) []O {
	var mapped []O
	for _, e := range data {
		m := f(e)
		if m != nil {
			mapped = append(mapped, *m)
		}
	}
	return mapped
}
