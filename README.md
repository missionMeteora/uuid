# UUID
A simple and fast UUID generating library 

# Benchmarks
```
BenchmarkCurrent-4              20000000		67.9 ns/op		 0 B/op		0 allocs/op
BenchmarkCurrentString-4        10000000		 208 ns/op		64 B/op		2 allocs/op
BenchmarkCurrentParallel-4      10000000		 173 ns/op		 0 B/op		0 allocs/op
BenchmarkCurrentParallelGen-4   50000000		35.4 ns/op		 0 B/op		0 allocs/op

BenchmarkBasic-4                 5000000		 330 ns/op		96 B/op		4 allocs/op
BenchmarkBasicParallel-4        10000000		 214 ns/op		96 B/op		4 allocs/op

BenchmarkPbor-4                  1000000		1683 ns/op		64 B/op		2 allocs/op
BenchmarkPborParallel-4          1000000		1222 ns/op		64 B/op		2 allocs/op
```