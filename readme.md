# Testing and benchmark results

SliceOffset is okra's implementation.

```
goos: darwin
goarch: amd64
BenchmarkArrayOffset-4   	57465946	        19.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceOffset-4   	91317076	        13.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkMapsOffset-4    	 2813293	       396 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	_/Users/javorszky/goSites/dayoffset	4.891s
```
