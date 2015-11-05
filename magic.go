package magic

import (
	"reflect"
	"time"
)

// Change represents a change of value in a struct field.
type Change struct {
	Name string      // The field name.
	Old  interface{} // The previous (a) value.
	New  interface{} // The current (b) value.
}

// Diff calculates the changes between structs a and b.
func Diff(a, b interface{}) []Change {
	av := indirect(reflect.ValueOf(a))
	bv := indirect(reflect.ValueOf(b))

	if av.Kind() != reflect.Struct || bv.Kind() != reflect.Struct {
		panic("magic: not struct")
	}

	if av.Type() != bv.Type() {
		panic("magic: different types")
	}

	var changes []Change

	for i := 0; i < av.Type().NumField(); i++ {
		ftype := av.Type().Field(i)
		fav := av.Field(i)
		fbv := bv.Field(i)

		if !fav.CanInterface() {
			continue
		}

		if !equals(fav, fbv) {
			changes = append(changes, Change{
				Name: ftype.Name,
				Old:  fav.Interface(),
				New:  fbv.Interface(),
			})
		}
	}

	return changes
}

func indirect(rv reflect.Value) reflect.Value {
	for rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	return rv
}

var timeType = reflect.TypeOf(time.Time{})

func equals(av, bv reflect.Value) bool {
	aface, bface := av.Interface(), bv.Interface()
	if av.Type() == timeType {
		at := aface.(time.Time)
		bt := bface.(time.Time)
		return at.Equal(bt)
	}
	return reflect.DeepEqual(aface, bface)
}
