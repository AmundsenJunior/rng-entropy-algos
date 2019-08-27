# rng-entropy-algos
*A randomness algorithm that, using various entropy generators, outputs a random 
stream of characters to `STDOUT`, á la calls to `/dev/random` on Linux systems
(e.g., `$ cat /dev/random`)*

## Build and execution

To run the program:

```
$ go build
$ ./rng-entropy-algos
```

Exit program execution via `Ctrl+C`.

See comments in [rng-entropy-algos.go](./rng-entropy-algos.go) for details of the
solution.

## Testing

A number of tests are provided to evaluate the randomness of the characters
stream produced by `rng-entropy-algos` (in-progress). See comments in
[rng-entropy-algos_test.go](./rng-entropy-algos_test.go) for more details on the
statistical tests being used.

If using Go 1.11+, use `go mod download` to pull the 
[gonum](https://godoc.org/gonum.org/v1/gonum/stat) test dependency. Otherwise, 
`go get gonum.org/v1/gonum/stat`.

To run the tests:

```
$ go test
```

## References

A number of online materials were consulted prior to working on this code, as 
I had no prior familiarity to the problem space of random number generators
and system entropy. These materials were of particular help in preparing:

* http://www.chronox.de/jent/doc/CPU-Jitter-NPTRNG.html
* https://medium.com/swlh/random-functions-a4f36b1dfd8f
* https://lwn.net/Articles/283103/
* https://hackaday.com/2017/11/02/what-is-entropy-and-how-do-i-get-more-of-it/
* https://cs.stackexchange.com/a/29881
* https://elixir.bootlin.com/linux/latest/source/drivers/char/random.c
