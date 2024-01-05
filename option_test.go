package maybe_test

import (
	"fmt"
	"testing"

	"github.com/shota3506/maybe"
)

type SampleStruct struct {
	x string
}

func (s SampleStruct) String() string {
	return s.x
}

func TestOption(t *testing.T) {
	t.Run("int64", func(t *testing.T) {
		for _, tc := range []struct {
			name string
			opt  maybe.Option[int64]
			v    int64
			ok   bool
		}{
			{
				name: "valid",
				opt:  maybe.New[int64](10),
				v:    10,
				ok:   true,
			},
			{
				name: "none",
				opt:  maybe.None[int64](),
				v:    0,
				ok:   false,
			},
			{
				name: "zero",
				opt:  maybe.Option[int64]{},
				v:    0,
				ok:   false,
			},
		} {
			t.Run(tc.name, func(t *testing.T) {
				v, ok := tc.opt.Value()
				if ok != tc.ok {
					t.Errorf("expected ok is %v, but got %v", tc.ok, ok)
				}
				if v != tc.v {
					t.Errorf("expected value is %v, but got %v", tc.v, v)
				}
			})
		}
	})

	t.Run("string", func(t *testing.T) {
		for _, tc := range []struct {
			name string
			opt  maybe.Option[string]
			v    string
			ok   bool
		}{
			{
				name: "valid",
				opt:  maybe.New[string]("Hello World!"),
				v:    "Hello World!",
				ok:   true,
			},
			{
				name: "none",
				opt:  maybe.None[string](),
				v:    "",
				ok:   false,
			},
			{
				name: "zero",
				opt:  maybe.Option[string]{},
				v:    "",
				ok:   false,
			},
		} {
			t.Run(tc.name, func(t *testing.T) {
				v, ok := tc.opt.Value()
				if ok != tc.ok {
					t.Errorf("expected ok is %v, but got %v", tc.ok, ok)
				}
				if v != tc.v {
					t.Errorf("expected value is %v, but got %v", tc.v, v)
				}
			})
		}
	})

	t.Run("interface", func(t *testing.T) {
		for _, tc := range []struct {
			name string
			opt  maybe.Option[fmt.Stringer]
			v    fmt.Stringer
			ok   bool
		}{
			{
				name: "valid",
				opt:  maybe.New[fmt.Stringer](SampleStruct{x: "sample"}),
				v:    SampleStruct{x: "sample"},
				ok:   true,
			},
			{
				name: "none",
				opt:  maybe.None[fmt.Stringer](),
				v:    nil,
				ok:   false,
			},
			{
				name: "zero",
				opt:  maybe.Option[fmt.Stringer]{},
				v:    nil,
				ok:   false,
			},
		} {
			t.Run(tc.name, func(t *testing.T) {
				v, ok := tc.opt.Value()
				if ok != tc.ok {
					t.Errorf("expected ok is %v, but got %v", tc.ok, ok)
				}
				if v != tc.v {
					t.Errorf("expected value is %v, but got %v", tc.v, v)
				}
			})
		}
	})

	t.Run("struct", func(t *testing.T) {
		for _, tc := range []struct {
			name string
			opt  maybe.Option[SampleStruct]
			v    SampleStruct
			ok   bool
		}{
			{
				name: "valid",
				opt:  maybe.New[SampleStruct](SampleStruct{x: "sample"}),
				v:    SampleStruct{x: "sample"},
				ok:   true,
			},
			{
				name: "none",
				opt:  maybe.None[SampleStruct](),
				v:    SampleStruct{},
				ok:   false,
			},
			{
				name: "zero",
				opt:  maybe.Option[SampleStruct]{},
				v:    SampleStruct{},
				ok:   false,
			},
		} {
			t.Run(tc.name, func(t *testing.T) {
				v, ok := tc.opt.Value()
				if ok != tc.ok {
					t.Errorf("expected ok is %v, but got %v", tc.ok, ok)
				}
				if v != tc.v {
					t.Errorf("expected value is %v, but got %v", tc.v, v)
				}
			})
		}
	})
}
