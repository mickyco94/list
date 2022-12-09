# List

List is an implementation of C#'s `List<T>` type in Go.

This is not something that should be used practically and only stands here as an exercise in data structures.

As shown by the benchmarks below the builtin `append()` function on slices works more efficiently than this implementation. The implementations are similar in principle, except for the Go specification optimisations found in `append()` and its lack of use of generics.

## Benches

```sh
goos: linux
goarch: amd64
pkg: github.com/mickyco94/list
cpu: Intel(R) Core(TM) i7-7700K CPU @ 4.20GHz
BenchmarkAdd-8      	41773298	        37.46 ns/op	      50 B/op	       0 allocs/op
BenchmarkAppend-8   	73854888	        34.68 ns/op	      82 B/op	       0 allocs/op
```
