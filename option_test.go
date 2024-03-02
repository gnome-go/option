package option_test

import (
	"testing"

	"github.com/gnome-go/option"
)

type TestStruct struct {
	Value int
}

func TestIsSome(t *testing.T) {
	s := option.Some(10)
	if !s.IsSome() {
		t.Error("Some(10) must be some")
	}

	s = option.None[int]()
	if s.IsSome() {
		t.Error("None() must not be some")
	}
}

func TestIsSomeAnd(t *testing.T) {
	s := option.Some(10)
	f := func(i int) bool {
		return i > 5
	}

	if !s.IsSomeAnd(f) {
		t.Error("Some(10) must be some and i > 5")
	}

	s = option.Some(3)
	if s.IsSomeAnd(f) {
		t.Error("Some(3) must not be some and i > 5")
	}

	s = option.None[int]()
	if s.IsSomeAnd(f) {
		t.Error("None() must not be some and i > 5")
	}
}

func TestIsNone(t *testing.T) {
	s := option.Some(10)
	if s.IsNone() {
		t.Error("Some(10) must not be none")
	}

	s = option.None[int]()
	if !s.IsNone() {
		t.Error("None() must be none")
	}
}

func TestSomeAndUnwrap(t *testing.T) {
	s := option.Some(10)
	v := s.Unwrap()

	if s.IsNone() {
		t.Error("Some(10) must not be none")
	}

	if v != 10 {
		t.Error("Some(10).Unwrap() must be 10")
	}
}

func TestFromAndUnwrap(t *testing.T) {
	s := option.From(&TestStruct{Value: 10})
	if s.IsNone() {
		t.Error("Some(10) must not be none")
	}

	v := s.Unwrap()

	if v.Value != 10 {
		t.Error("Unwrap() for TestStruct.Value must be 10")
	}
}

func TestNoneAndUnwrap(t *testing.T) {
	s := option.None[int]()
	if !s.IsNone() {
		t.Error("None() must be none")
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("Unwrap() for None() must panic")
		}
	}()

	s.Unwrap()
}

func TestUnwrap(t *testing.T) {
	s := option.Some(10)
	v := s.Unwrap()

	if v != 10 {
		t.Error("Unwrap() for Some(10) must be 10")
	}

	s = option.None[int]()
	defer func() {
		if r := recover(); r == nil {
			t.Error("Unwrap() for None() must panic")
		}
	}()

	s.Unwrap()
}

func TestUnwrapOr(t *testing.T) {
	s := option.Some(10)
	v := s.UnwrapOr(20)

	if v != 10 {
		t.Error("UnwrapOr(20) for Some(10) must be 10")
	}

	s = option.None[int]()
	v = s.UnwrapOr(20)

	if v != 20 {
		t.Error("UnwrapOr(20) for None() must be 20")
	}
}

func TestUnwrapOrElse(t *testing.T) {
	s := option.Some(10)
	v := s.UnwrapOrElse(func() int { return 20 })

	if v != 10 {
		t.Error("UnwrapOrElse(20) for Some(10) must be 10")
	}

	s = option.None[int]()
	v = s.UnwrapOrElse(func() int { return 20 })

	if v != 20 {
		t.Error("UnwrapOrElse(20) for None() must be 20")
	}
}

func TestUnwrapUnsafe(t *testing.T) {
	s := option.Some(10)
	v := s.UnwrapUnsafe()

	if v != 10 {
		t.Error("UnwrapUnsafe() for Some(10) must be 10")
	}

	s = option.None[int]()
	defer func() {
		if r := recover(); r == nil {
			t.Error("UnwrapUnsafe() for None() must panic")
		}
	}()

	s.UnwrapUnsafe()
}

func TestMap(t *testing.T) {
	s := option.Some(10)
	v := option.Map(s, func(i int) int { return i * 2 })

	if v.Unwrap() != 20 {
		t.Error("Map(func(i int) int { return i * 2 }) for Some(10) must be 20")
	}

	s = option.None[int]()
	v = option.Map(s, func(i int) int { return i * 2 })

	if !v.IsNone() {
		t.Error("Map(func(i int) int { return i * 2 }) for None() must be none")
	}
}

func TestMapOr(t *testing.T) {
	s := option.Some(10)
	v := option.MapOr(s, 30, func(i int) int { return i * 2 })

	if v != 20 {
		t.Error("MapOr(func(i int) int { return i * 2 }, 20) for Some(10) must be 20")
	}

	s = option.None[int]()
	v = option.MapOr(s, 30, func(i int) int { return i * 2 })

	if v != 30 {
		t.Error("MapOr(func(i int) int { return i * 2 }, 20) for None() must be 20")
	}
}

func TestMapOrElse(t *testing.T) {
	s := option.Some(10)
	v := option.MapOrElse(s, func() int { return 30 }, func(i int) int { return i * 2 })

	if v != 20 {
		t.Error("MapOrElse(func(i int) int { return i * 2 }, 20) for Some(10) must be 20")
	}

	s = option.None[int]()
	v = option.MapOrElse(s, func() int { return 30 }, func(i int) int { return i * 2 })

	if v != 30 {
		t.Error("MapOrElse(func(i int) int { return i * 2 }, 20) for None() must be 20")
	}
}

func TestAnd(t *testing.T) {
	s1 := option.Some(10)
	s2 := option.Some(20)
	s3 := option.None[int]()

	v := s1.And(s2)
	if v.Unwrap() != 20 {
		t.Error("And(Some(10), Some(20)) must be Some(20)")
	}

	v = s1.And(s3)
	if !v.IsNone() {
		t.Error("And(Some(10), None()) must be none")
	}

	v = s3.And(s2)
	if !v.IsNone() {
		t.Error("And(None(), Some(20)) must be none")
	}

	v = s3.And(s3)
	if !v.IsNone() {
		t.Error("And(None(), None()) must be none")
	}
}

func TestAndThen(t *testing.T) {
	s := option.Some(10)
	f := func(i int) option.Option[int] {
		return option.Some(i * 2)
	}

	v := s.AndThen(f)
	if v.Unwrap() != 20 {
		t.Error("AndThen(func(i int) Option[int] { return Some(i * 2) }) for Some(10) must be Some(20)")
	}

	s = option.None[int]()
	v = s.AndThen(f)
	if !v.IsNone() {
		t.Error("AndThen(func(i int) Option[int] { return Some(i * 2) }) for None() must be none")
	}
}

func TestOr(t *testing.T) {
	s1 := option.Some(10)
	s2 := option.Some(20)
	s3 := option.None[int]()

	v := s1.Or(s2)
	if v.Unwrap() != 10 {
		t.Error("Or(Some(10), Some(20)) must be Some(10)")
	}

	v = s1.Or(s3)
	if v.Unwrap() != 10 {
		t.Error("Or(Some(10), None()) must be Some(10)")
	}

	v = s3.Or(s2)
	if v.Unwrap() != 20 {
		t.Error("Or(None(), Some(20)) must be Some(20)")
	}

	v = s3.Or(s3)
	if !v.IsNone() {
		t.Error("Or(None(), None()) must be none")
	}
}

func TestOrElse(t *testing.T) {
	s := option.Some(10)
	f := func() option.Option[int] {
		return option.Some(20)
	}

	v := s.OrElse(f)
	if v.Unwrap() != 10 {
		t.Error("OrElse(func() Option[int] { return Some(20) }) for Some(10) must be Some(10)")
	}

	s = option.None[int]()
	v = s.OrElse(f)
	if v.Unwrap() != 20 {
		t.Error("OrElse(func() Option[int] { return Some(20) }) for None() must be Some(20)")
	}
}

func TestFilter(t *testing.T) {

	s := option.Some(10)
	f := func(i int) bool {
		return i > 5
	}

	v := s.Filter(f)
	if v.Unwrap() != 10 {
		t.Error("Filter(func(i int) bool { return i > 5 }) for Some(10) must be Some(10)")
	}

	s = option.Some(3)
	v = s.Filter(f)
	if !v.IsNone() {
		t.Error("Filter(func(i int) bool { return i > 5 }) for Some(3) must be none")
	}

	s = option.None[int]()
	v = s.Filter(f)
	if !v.IsNone() {
		t.Error("Filter(func(i int) bool { return i > 5 }) for None() must be none")
	}
}

func TestExpect(t *testing.T) {
	s := option.Some(10)
	v := s.Expect("panic")

	if v != 10 {
		t.Error("Expect(\"panic\") for Some(10) must be 10")
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expect(\"panic\") for None() must panic")
		}
	}()

	s = option.None[int]()
	s.Expect("panic")
}
