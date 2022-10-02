package gomap

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMapp(t *testing.T) {
	mm := New[int64](8)

	v := mm.Get("123")
	isEqual(t, v, int64(0))

	mm.Put("key1", 10)
	v = mm.Get("key1")
	isEqual(t, v, int64(10))

	mm.Put("key2", 20)
	v = mm.Get("key2")
	isEqual(t, v, int64(20))

	mm.Put("key2", 30)
	v = mm.Get("key2")
	isEqual(t, v, int64(30))

	mm.Put("adsdadw1231", 4423)
	v = mm.Get("adsdadw1231")
	isEqual(t, v, int64(4423))

	mm.Put("", 144)
	isEqual(t, mm.Get(""), int64(144))

	mm.Put(" ", 145)
	isEqual(t, mm.Get(" "), int64(145))
}

func isEqual(t *testing.T, got interface{}, want interface{}) {
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("result is not equal\ngot:  %+v\nwant: %+v\n", got, want)
	}
}

func TestBucketOverflow(t *testing.T) {
	// create map with 1 bucket
	mm := New[int](1)

	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	prefix := "key_"

	for _, v := range values {
		mm.Put(fmt.Sprintf("%s%d", prefix, v), v)
	}

	for _, v := range values {
		got := mm.Get(fmt.Sprintf("%s%d", prefix, v))
		isEqual(t, got, v)
	}
}
