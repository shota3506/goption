package maybe

type Option[T any] struct {
	v     T
	valid bool
}

func New[T any](v T) Option[T] {
	return Option[T]{v, true}
}

func None[T any]() Option[T] {
	return Option[T]{}
}

func (o Option[T]) Value() (T, bool) {
	if o.valid {
		return o.v, true
	}
	var t T // return zero value
	return t, false
}

func (o Option[T]) Valid() bool {
	return o.valid
}
