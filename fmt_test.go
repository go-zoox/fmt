package fmt

import "testing"

func TestFmt(t *testing.T) {
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

	Println("JSON Format:")
	PrintJSON(user)
	Println("\nYAML Format:")
	PrintYAML(user)
	Println("\nTOML Format:")
	PrintTOML(user)
	// Println("\nINI Format:")
	// PrintINI(user)
	Println("\nRaw Format:")
	Println(user)

}
