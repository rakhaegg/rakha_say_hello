package rakha_say_hello

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTableSayHello(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "TEST 1",
			request:  "RakhaElang",
			expected: "RakhaElang",
		},
	}

	a := []Person{
		{
			id:   123,
			name: "RakhaElang",
		},
		{
			id:   12,
			name: "ElamgGunawan",
		},
		{
			id:   12,
			name: "Gunawan",
		},
		{
			id:   123,
			name: "AXEL",
		},
	}
	for _, test := range tests {
		for _, data := range a {

		}
		resultName := a.SayHello(test.expected)
		require.Equal(t, test.expected, resultName)
	}

}

// func TestShowHello(t *testing.T) {
// 	result := SayHello("Rakha")
// 	require.Equal(t, "Hello Rakha", result, "Result is wrong")
// 	fmt.Println("Test Done")
// }
