package gomap

import (
	"fmt"
	"testing"
)

var sizes = []int{128, 8192, 32768, 131072}

func BenchmarkGet(b *testing.B) {
	for _, n := range sizes {
		keys := make([]string, 0, n)
		mm := New[string, int64](n)
		stdm := make(map[string]int64, n)

		for i := 0; i < n; i++ {
			k := fmt.Sprintf("key__%d", i)
			mm.Put(k, int64(i)*2)
			stdm[k] = int64(i) * 2
			keys = append(keys, k)
		}

		j := 0
		b.Run(fmt.Sprintf("generic-map %d", n), func(b *testing.B) {
			var got int64
			for i := 0; i < b.N; i++ {
				if j == n {
					j = 0
				}
				got = mm.Get(keys[j])
				j++
			}
			_ = got
		})

		j = 0
		b.Run(fmt.Sprintf("STD-map     %d", n), func(b *testing.B) {
			var got int64
			for i := 0; i < b.N; i++ {
				if j == n {
					j = 0
				}
				got = stdm[keys[j]]
				j++
			}
			_ = got
		})
	}
}

func BenchmarkPut(b *testing.B) {
	for _, n := range sizes {
		keys := make([]string, 0, n)
		for i := 0; i < n; i++ {
			keys = append(keys, fmt.Sprintf("key__%d", i))
		}
		mm := New[string, int64](n)
		j := 0
		multiplier := 1
		b.Run(fmt.Sprintf("generic-map %d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				if j == n {
					j = 0
					multiplier += 1
				}
				mm.Put(keys[j], int64(j*multiplier))
				j++
			}
		})

		j = 0
		multiplier = 1
		stdm := make(map[string]int64, n)
		b.Run(fmt.Sprintf("STD-map     %d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				if j == n {
					j = 0
					multiplier += 1
				}
				stdm[keys[j]] = int64(j * multiplier)
				j++
			}
		})
	}
}

func BenchmarkPutWithOverflow(b *testing.B) {
	startSize := 1_000
	targetSize := []int{10_000, 100_000, 1_000_000, 10_000_000}
	type someStruct struct {
		x string
		y int
	}

	for _, n := range targetSize {
		keys := make([]string, 0, n)
		for i := 0; i < n; i++ {
			keys = append(keys, fmt.Sprintf("key__%d", i))
		}

		mm := New[string, someStruct](startSize)
		j := 0
		multiplier := 1
		b.Run(fmt.Sprintf("gen-map  (string key)%d", n), func(b *testing.B) {
			var key string
			for i := 0; i < b.N; i++ {
				if j == n {
					j = 0
					multiplier += 1
				}
				key = keys[j]
				mm.Put(key, someStruct{x: key, y: j * multiplier})
				j++
			}
		})

		stdm := make(map[string]someStruct, startSize)
		j = 0
		multiplier = 1
		b.Run(fmt.Sprintf("STD      (string key)%d", n), func(b *testing.B) {
			var key string
			for i := 0; i < b.N; i++ {
				if j == n {
					j = 0
					multiplier += 1
				}
				key = keys[j]
				stdm[key] = someStruct{x: key, y: j * multiplier}
				j++
			}
		})
	}
}
