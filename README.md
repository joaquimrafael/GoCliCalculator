# Go CLI Calculator

A simple command-line calculator written in Go. It runs in a loop, reading
expressions from the console until you tell it to stop. Each line takes two
numbers and an operator (`+ - * /`) and prints the result.

This is a learning project focused on Go fundamentals: errors as return values,
multiple return values, `switch` without fallthrough, and reading input with
`bufio.Scanner`.

## Requirements

- [Go](https://go.dev/dl/) 1.26 or newer

## How to run

From the project root:

```bash
go run .
```

Then type an expression as `value1 value2 operator` and press Enter:

```bash
Type: value1 value2 operator ->
10 4 /
Result: 2.50
```

Type `close` to exit the program.

Invalid input (wrong number of fields, non-numeric values, division by zero,
or an unknown operator) prints a warning and keeps the loop running instead of
crashing.

You can also build a standalone binary:

```bash
go build -o calculator .
./calculator
```

## How to test

The core `calculate` function has a table-driven test suite:

```bash
go test ./...        # run all tests
go test -v ./...     # verbose: show each subtest
go test -cover ./...  # show test coverage
```

## Project structure

```Diagram
GoCliCalculator/
├── go.mod                # module definition
├── calculator.go         # main program: input loop + calculate()
├── calculator_test.go    # unit tests for calculate()
└── README.md             # this file
```

- **`calculate(a, b float64, op string) (float64, error)`** — pure function
  that performs one operation and returns an `error` for invalid cases (e.g.
  division by zero) instead of panicking.
- **`main()`** — reads from `os.Stdin` with a `bufio.Scanner`, parses each line,
  calls `calculate`, and prints the result, looping until `close` or EOF.
