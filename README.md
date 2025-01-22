# JSON Serialization Benchmark with EasyJSON

This project benchmarks JSON serialization and deserialization performance using the [easyjson](https://github.com/mailru/easyjson) library compared to Go's standard `encoding/json` package.

## Features

- **JSON Serialization/Deserialization**:
  - Compare standard `encoding/json` with `easyjson`.
- **Benchmarking**:
  - Benchmarks both serialization and deserialization for precise performance comparison.
- **Scalable Structure**:
  - Easily extendable for other serialization libraries or use cases.

## Prerequisites

- Go 1.20+ installed on your system.
- `easyjson` library installed:
  ```go
  go get github.com/mailru/easyjson && go install github.com/mailru/easyjson/...@latest
  ```
- Add gopath in PATH
  ```bash
  export GOPATH=/Users/<username>/go
  export PATH=$GOPATH/bin:$PATH
  ```

## Installation and Setup

- Clone the repository:

  ```bash
  git clone https://github.com/Akshit-Zatakia/go-easyjson-benchmark

  cd go-easyjson-benchmark
  ```

- Generate the easyjson file
  ```bash
  easyjson -all user.go
  ```

## Benchmarking

To benchmark JSON serialization and deserialization:

- Run the benchmarks:
  ```bash
  go test -bench=. -benchmem
  ```
- Example output:
  ```bash
    goos: darwin
    goarch: amd64
    pkg: go-easyjson-benchmark
    cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz

    BenchmarkSerialization/Standard_JSON_Marshal-16                  2512738               478.6 ns/op           168 B/op          4 allocs/op
    BenchmarkSerialization/easyjson_Marshal-16                      12108747                94.39 ns/op          192 B/op          2 allocs/op
    BenchmarkSerialization/Standard_JSON_Unmarshal-16               13809604                83.57 ns/op          168 B/op          2 allocs/op
    BenchmarkSerialization/easyjson_Unmarshal-16                    1000000000               0.2372 ns/op          0 B/op          0 allocs/op

    PASS
    ok      go-easyjson-benchmark   5.019s
  ```
