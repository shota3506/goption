/*
Package maybe provides a generic Option type that can hold a value or not.

Option[T] is a runtime safe alternative to using nil for optional values.
It enforces the user to check if the value is valid before using it
because Value() method is an only way to get the value.
*/

package maybe

// Option is a type that can either hold a value or not.
type Option[T any] struct {
	v     T
	valid bool
}

// New returns an Option that holds the given value.
func New[T any](v T) Option[T] {
	return Option[T]{v, true}
}

// None returns an Option that holds no value.
// It is equivalent to zero value of Option like Option[T]{}.
func None[T any]() Option[T] {
	return Option[T]{}
}

// Value returns the value held by the Option if it is valid.
// Otherwise it returns zero value of the type T and false.
func (o Option[T]) Value() (T, bool) {
	if o.valid {
		return o.v, true
	}
	var t T // return zero value
	return t, false
}

// Valid returns true if the Option holds a value.
func (o Option[T]) Valid() bool {
	return o.valid
}
