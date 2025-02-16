package auth

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected string
		has_err  bool
	}{
		"valid header": {
			input:    "ApiKey 1234567890",
			expected: "1234567890",
			has_err:  false,
		},
		"invalid header": {
			input:    "ApiKey",
			expected: "",
			has_err:  true,
		},
		"valid header 2": {
			input:    "ApiKey 1234567890 1234567890",
			expected: "1234567890",
			has_err:  false,
		},
		"invalid header 2": {
			input:    "HouseKey 1234567890",
			expected: "",
			has_err:  true,
		},
		"empty header": {
			input:    "",
			expected: "",
			has_err:  false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			headers := make(http.Header)
			headers.Set("Authorization", test.input)
			got_string, got_err := GetAPIKey(headers)
			diff := cmp.Diff(test.expected, got_string)
			if diff != "" {
				t.Fatalf("%s", diff)
			}
			if test.has_err != (got_err != nil) {
				t.Fatalf("expected error: %t, got: %t", test.has_err, got_err != nil)
			}
		})
	}
}
