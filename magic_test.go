package magic

import (
	"reflect"
	"testing"
)

func TestDiff(t *testing.T) {
	type widget struct {
		ID   int
		Name string
		X    *int
	}
	a := widget{
		ID:   1,
		Name: "A",
		X:    new(int),
	}
	b := widget{
		ID:   1,
		Name: "B",
		X:    new(int),
	}
	diff := Diff(a, b)

	expect := []Change{
		{
			Name: "Name",
			Old:  a.Name,
			New:  b.Name,
		},
	}

	if !reflect.DeepEqual(diff, expect) {
		t.Errorf("bad diff. %#v ≠ %#v", diff, expect)
	}
}
