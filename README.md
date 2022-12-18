# gt: Generics based Test library for Go [![Go Reference](https://pkg.go.dev/badge/github.com/m-mizutani/gt.svg)](https://pkg.go.dev/github.com/m-mizutani/gt) [![test](https://github.com/m-mizutani/gt/actions/workflows/test.yml/badge.svg)](https://github.com/m-mizutani/gt/actions/workflows/test.yml) [![gosec](https://github.com/m-mizutani/gt/actions/workflows/gosec.yml/badge.svg)](https://github.com/m-mizutani/gt/actions/workflows/gosec.yml) [![lint](https://github.com/m-mizutani/gt/actions/workflows/lint.yml/badge.svg)](https://github.com/m-mizutani/gt/actions/workflows/lint.yml)

`gt` is test library leveraging Go generics to check variable type in IDE and compiler.

```go
// gt.Array(t, "123") <- Error before running the test

// gt.Array() accepts only slice and array
unorderedUsers, _ := GetUsers(ctx)
gt.Array(t, unorderedUsers).
    Contain(&user{
        ID:   1000,
        Name: "Alice",
    }).
    Contain(&user{
        ID:   1024,
        Name: "Bob",
    }).
    NotContain(&user{
        ID:   9999,
        Name: "TestUser",
    }).
    Length(3)
```

## Motivation

Existing test libraries in Go such as [testify](https://github.com/stretchr/testify) strongly support writing unit test. Many test libraries uses `reflect` package to identify and compare variable type and value and test functions of the libraries accept any type by `interface{}`. However the approach has two problems:

- Variable types mismatch between _expected_ and _actual_ can not be detected before running the test.
- IDE can not support variable completion because types can not be determined before running the test.

On the other hand, Go started to provide [Generics](https://go.dev/doc/tutorial/generics) feature by version 1.18. It can be leveraged to support type completion and checking types before running a test.

## Usage

In many cases, a developer does not care Go generics in using `gt`. However, a developer need to specify generic type (`Value`, `Array`, `Map`, `Error`, etc.) explicitly to use specific test functions for each types.

### Array case
```go
a1 := []int{1, 2, 3}
gt.Value(t, a1).Equal([]int{1, 2, 3}) // <- OK
// gt.Value(t, a1).Contain(1) <- Compile error
gt.Array(t, a1).Contains(1) // <- OK
```

### Struct case
```go
type user struct {
    Name string
}
u1 := user{Name: "blue"}

gt.Value(u1).Equal(user{Name:"blue"}) // <- OK
// gt.Value(u1).Equal(&user{Name:"blue"}) // <- Compile error
```

## License

Apache License 2.0