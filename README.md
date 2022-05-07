# Fmt - Pretty Data Display, Inlcude JSON/YAML/TOML/INI ...

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/fmt)](https://pkg.go.dev/github.com/go-zoox/fmt)
[![Build Status](https://github.com/go-zoox/fmt/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/fmt/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/fmt)](https://goreportcard.com/report/github.com/go-zoox/fmt)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/fmt/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/fmt?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/fmt.svg)](https://github.com/go-zoox/fmt/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/fmt.svg?label=Release)](https://github.com/go-zoox/fmt/tags)

## Installation
To install the package, run:
```bash
go get github.com/go-zoox/fmt
```

## Getting Started

```go
import (
  "testing"
  "github.com/go-zoox/fmt"
)

func main(t *testing.T) {
		type User struct {
		Name    string
		Age     int
		Address struct {
			City   string
			Street string
		}
		Hobbies []string
	}

	user := User{
		Name: "Zero",
		Age:  18,
		Address: struct {
			City   string
			Street string
		}{
			City:   "Shanghai",
			Street: "Nanjingxi Road Street",
		},
		Hobbies: []string{"sport", "music", "movies"},
	}

	fmt.Println("JSON Format:")
	fmt.PrintJSON(user)
	fmt.Println("\nYAML Format:")
	fmt.PrintYAML(user)
	fmt.Println("\nTOML Format:")
	fmt.PrintTOML(user)
	// Println("\nINI Format:")
	// PrintINI(user)
	fmt.Println("\nRaw Format:")
	fmt.Println(user)
}
```

## License
GoZoox is released under the [MIT License](./LICENSE).
