package main

// ArrayOffset function to get new day using an array.
func ArrayOffset(day string, offset int) string {
	days := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	for i, value := range days {
		if value == day {
			offset += i
			break
		}
	}
	return days[(offset%7+7)%7]
}

// This is okra's original solution
// SliceOffset function to get new day using a slice.
func SliceOffset(day string, offset int) string {
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	for i, value := range days {
		if value == day {
			offset += i
			break
		}
	}
	return days[(offset%7+7)%7]
}

// MapsOffset function to get new day using 2 maps.
func MapsOffset(day string, offset int) string {
	daysMap := map[string]int{
		"Monday":    0,
		"Tuesday":   1,
		"Wednesday": 2,
		"Thursday":  3,
		"Friday":    4,
		"Saturday":  5,
		"Sunday":    6,
	}

	offsetsMap := map[int]string{
		0: "Monday",
		1: "Tuesday",
		2: "Wednesday",
		3: "Thursday",
		4: "Friday",
		5: "Saturday",
		6: "Sunday",
	}

	dOffset := daysMap[day]
	newOffset := ((dOffset + offset)%7 + 7) % 7
	return offsetsMap[newOffset]
}

func main() {

}
