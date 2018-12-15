io.Copy vs io.CopyBuffer benchmark; run with:
````
go test -bench .
````
Results:
````
> go test -bench .
goos: windows
goarch: amd64
pkg: iocpb
BenchmarkIoCopy-4                         200000              8295 ns/op
BenchmarkIoCopyBufferWithPool-4          2000000               609 ns/op
BenchmarkIoCopyBuffer-4                  3000000               534 ns/op
PASS
ok      iocpb   5.786s
````