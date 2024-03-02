// Description: This file contains the option struct and its methods
package option

type enum int

const (
	none enum = iota
	some
)

type Option[T any] struct {
	value *T
	state enum
}

// None returns an Option that contains no value.
func None[T any]() Option[T] {
	return Option[T]{state: none}
}

// Some returns an Option that contains the given value.
func Some[T any](value T) Option[T] {
	return Option[T]{value: &value, state: some}
}

// From returns an Option with a state of Some if the given value is not nil;
// otherwise, it returns None.
func From[T any](value *T) Option[T] {
	if value == nil {
		return None[T]()
	}
	return Some(*value)
}

// New returns an Option that contains the given value if it is
// not nil, otherwise it returns None.
func (self Option[T]) IsNone() bool {
	return self.state == none
}

// New returns an Option that contains the given value if it is
// not nil, otherwise it returns None.
func (self Option[T]) IsSome() bool {
	return self.state == some
}

// Test if the option is some and the value satisfies the given predicate.
func (self Option[T]) IsSomeAnd(f func(T) bool) bool {
	if self.state == none {
		return false
	}
	return f(*self.value)
}

// Returns the value if the state is some. Otherwise, the default value is returned.
func (self Option[T]) And(opt Option[T]) Option[T] {
	if self.state == none {
		return None[T]()
	}
	return opt
}

// Returns the value if the state is some. Otherwise, the result
// of the given function is returned.
func (o Option[T]) AndThen(f func(T) Option[T]) Option[T] {
	if o.state == none {
		return None[T]()
	}
	return f(*o.value)
}

// Map returns an Option that contains the result of applying the given
// function to the value of the original Option, if it has one.
func Map[T any, U any](self Option[T], f func(T) U) Option[U] {
	if self.state == none {
		return None[U]()
	}
	return Some(f(*self.value))
}

// MapOr returns the result of applying the given function to the value of the
// original Option, if it has one. Otherwise, the default value is returned.
func MapOr[T any, U any](self Option[T], def U, f func(T) U) U {
	if self.state == none {
		return def
	}
	return f(*self.value)
}

// MapOrElse returns the result of applying the given function to the value of the
// original Option, if it has one. Otherwise, the result of the given function is returned.
func MapOrElse[T any, U any](self Option[T], def func() U, f func(T) U) U {
	if self.state == none {
		return def()
	}
	return f(*self.value)
}

// Or returns the original Option if it is some, otherwise it returns the given Option.
func (self Option[T]) Or(opt Option[T]) Option[T] {
	if self.state == none {
		return opt
	}
	return self
}

// OrElse returns the original Option if it is some, otherwise it returns the result
// of the given function.
func (self Option[T]) OrElse(f func() Option[T]) Option[T] {
	if self.state == none {
		return f()
	}
	return self
}

// Inspect calls the given function with the value of the Option if it is some.
func (self Option[T]) Inspect(f func(T)) {
	if self.state == some {
		f(*self.value)
	}
}

// Returns the value if the state is some. Otherwise,
// expect will panic with the provided message.
func (self Option[T]) Expect(msg string) T {
	if self.state == none {
		panic(msg)
	}
	return *self.value
}

// If the state is none or the predicate returns false, none is returned.
// Otherwise, the option with the some state is returned.
func (self Option[T]) Filter(f func(T) bool) Option[T] {
	if self.state == none {
		return None[T]()
	}
	if f(*self.value) {
		return self
	}
	return None[T]()
}

// Unwrap returns the value if the state is some. Otherwise, it panics.
func (self Option[T]) Unwrap() T {
	if self.state == none {
		panic("unwrap a none option")
	}
	return *self.value
}

// UnwrapOr returns the value if the state is some. Otherwise, the default value is returned.
func (self Option[T]) UnwrapOr(def T) T {
	if self.state == none {
		return def
	}
	return *self.value
}

// UnwrapOrElse returns the value if the state is some. Otherwise, the result
// of the given function is returned.
func (self Option[T]) UnwrapOrElse(f func() T) T {
	if self.state == none {
		return f()
	}

	return *self.value
}

// UnwrapUnsafe returns the value if the state is some. Otherwise, it panics.
func (self Option[T]) UnwrapUnsafe() T {
	return *self.value
}
