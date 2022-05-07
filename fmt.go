package fmt

import (
	gofmt "fmt"

	"github.com/go-zoox/encoding/ini"
	"github.com/go-zoox/encoding/json"
	"github.com/go-zoox/encoding/toml"
	"github.com/go-zoox/encoding/yaml"
)

// PrintJSON prints the JSON encoding of the data.
func PrintJSON(v any) {
	if t, err := json.Encode(v); err != nil {
		panic(err)
	} else {
		gofmt.Println(string(t))
	}
}

// PrintYAML prints the YAML encoding of the data.
func PrintYAML(v any) {
	if t, err := yaml.Encode(v); err != nil {
		panic(err)
	} else {
		gofmt.Println(string(t))
	}
}

// PrintTOML prints the TOML encoding of the data.
func PrintTOML(v any) {
	if t, err := toml.Encode(v); err != nil {
		panic(err)
	} else {
		gofmt.Println(string(t))
	}
}

// PrintINI prints the INI encoding of the data.
func PrintINI(v any) {
	if t, err := ini.Encode(v); err != nil {
		panic(err)
	} else {
		gofmt.Println(string(t))
	}
}

// Println prints the data, like fmt.Println.
func Println(v any) {
	gofmt.Println(v)
}

// Printf prints the data, like fmt.Printf.
func Printf(format string, v ...interface{}) {
	gofmt.Printf(format, v...)
}
