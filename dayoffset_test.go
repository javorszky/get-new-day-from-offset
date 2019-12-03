package main

import "testing"

func TestArray(t *testing.T) {
	tables := []struct {
		day string
		offset int
		newDay string
	}{
		{"Monday", 1, "Tuesday"},
		{"Monday", 7, "Monday"},
		{"Monday", 73, "Thursday"},
		{"Monday", -1, "Sunday"},
		{"Monday", -6, "Tuesday"},
		{"Monday", -73, "Friday"},
	}

	for _, table := range tables {
		newDay := ArrayOffset(table.day, table.offset)
		if newDay != table.newDay {
			t.Errorf("Getting the new day with array failed. Day: %s, Offset: %d, Expected: %s, Got: %s", table.day, table.offset, table.newDay, newDay)
		}
	}
}


func TestSlice(t *testing.T) {
	tables := []struct {
		day string
		offset int
		newDay string
	}{
		{"Monday", 1, "Tuesday"},
		{"Monday", 7, "Monday"},
		{"Monday", 73, "Thursday"},
		{"Monday", -1, "Sunday"},
		{"Monday", -6, "Tuesday"},
		{"Monday", -73, "Friday"},
	}

	for _, table := range tables {
		newDay := SliceOffset(table.day, table.offset)
		if newDay != table.newDay {
			t.Errorf("Getting the new day with slice failed. Day: %s, Offset: %d, Expected: %s, Got: %s", table.day, table.offset, table.newDay, newDay)
		}
	}
}


func TestMaps(t *testing.T) {
	tables := []struct {
		day string
		offset int
		newDay string
	}{
		{"Monday", 1, "Tuesday"},
		{"Monday", 7, "Monday"},
		{"Monday", 73, "Thursday"},
		{"Monday", -1, "Sunday"},
		{"Monday", -6, "Tuesday"},
		{"Monday", -73, "Friday"},
	}

	for _, table := range tables {
		newDay := MapsOffset(table.day, table.offset)
		if newDay != table.newDay {
			t.Errorf("Getting the new day with maps failed. Day: %s, Offset: %d, Expected: %s, Got: %s", table.day, table.offset, table.newDay, newDay)
		}
	}
}

func BenchmarkArrayOffset(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		ArrayOffset("Monday", -73)
	}
}

func BenchmarkSliceOffset(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		SliceOffset("Monday", -73)
	}
}

func BenchmarkMapsOffset(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		MapsOffset("Monday", -73)
	}
}
