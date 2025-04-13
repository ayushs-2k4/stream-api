package streams

type Stream[T any] struct {
	data []T
}

func New[T any](data []T) *Stream[T] {
	return &Stream[T]{data: data}
}

// Map transforms elements using fn
func (s *Stream[T]) Map(fn func(T) any) *Stream[any] {
	result := make([]any, len(s.data))
	for i, v := range s.data {
		result[i] = fn(v)
	}
	return &Stream[any]{data: result}
}

// Filter filters elements using fn
func (s *Stream[T]) Filter(fn func(T) bool) *Stream[T] {
	var result []T
	for _, v := range s.data {
		if fn(v) {
			result = append(result, v)
		}
	}
	return &Stream[T]{data: result}
}

// Reduce folds all values into one
func (s *Stream[T]) Reduce(fn func(T, T) T, initial T) T {
	acc := initial
	for _, v := range s.data {
		acc = fn(acc, v)
	}
	return acc
}

// ForEach runs a side-effect function
func (s *Stream[T]) ForEach(fn func(T)) {
	for _, v := range s.data {
		fn(v)
	}
}

// Any checks if any element matches the condition
func (s *Stream[T]) Any(fn func(T) bool) bool {
	for _, v := range s.data {
		if fn(v) {
			return true
		}
	}
	return false
}

// All checks if all elements match the condition
func (s *Stream[T]) All(fn func(T) bool) bool {
	for _, v := range s.data {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Find returns the first element matching fn, or false if not found
func (s *Stream[T]) Find(fn func(T) bool) (T, bool) {
	for _, v := range s.data {
		if fn(v) {
			return v, true
		}
	}
	var zero T
	return zero, false
}

// First returns the first element
func (s *Stream[T]) First() (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}
	return s.data[0], true
}

// Last returns the last element
func (s *Stream[T]) Last() (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}
	return s.data[len(s.data)-1], true
}

// Count returns number of elements
func (s *Stream[T]) Count() int {
	return len(s.data)
}

// Reverse returns a reversed slice
func (s *Stream[T]) Reverse() *Stream[T] {
	n := len(s.data)
	result := make([]T, n)
	for i := range s.data {
		result[n-1-i] = s.data[i]
	}
	return &Stream[T]{data: result}
}

// Collect unwraps back to []T
func (s *Stream[T]) Collect() []T {
	return s.data
}
