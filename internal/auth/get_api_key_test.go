package auth // Replace with your actual package name

import (
	"net/http"
	"testing"
)

// Note: Ensure ErrNoAuthHeaderIncluded is accessible in your test file.
// If it's in the same package, it will be automatically available.

func TestGetAPIKey(t *testing.T) {
	// Define the table of test cases
	tests := []struct {
		name          string
		headers       http.Header
		expectedKey   string
		expectedError string // Using string to easily compare error messages
	}{
		{
			name:          "No Authorization Header",
			headers:       http.Header{}, // Empty headers
			expectedKey:   "",
			expectedError: ErrNoAuthHeaderIncluded.Error(),
		},
		{
			name: "Empty Authorization Header",
			headers: http.Header{
				"Authorization": []string{""},
			},
			expectedKey:   "",
			expectedError: ErrNoAuthHeaderIncluded.Error(),
		},
		{
			name: "Malformed Header - Missing Prefix",
			headers: http.Header{
				"Authorization": []string{"123456789"},
			},
			expectedKey:   "",
			expectedError: "malformed authorization header",
		},
		{
			name: "Malformed Header - Wrong Prefix",
			headers: http.Header{
				"Authorization": []string{"Bearer 123456789"},
			},
			expectedKey:   "",
			expectedError: "malformed authorization header",
		},
		{
			name: "Valid API Key",
			headers: http.Header{
				"Authorization": []string{"ApiKey my-super-secret-key"},
			},
			expectedKey:   "my-super-secret-key",
			expectedError: "",
		},
		{
			name: "Valid API Key - Ignores Trailing Data",
			headers: http.Header{
				"Authorization": []string{"ApiKey my-super-secret-key extra-data"},
			},
			expectedKey:   "my-super-secret-key",
			expectedError: "",
		},
	}

	// Iterate over all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.headers)

			// Check if the returned key matches our expectation
			if key != tt.expectedKey {
				t.Errorf("Expected key %q, got %q", tt.expectedKey, key)
			}

			// Check error expectations
			if tt.expectedError != "" {
				if err == nil {
					t.Errorf("Expected error containing %q, got nil", tt.expectedError)
				} else if err.Error() != tt.expectedError {
					t.Errorf("Expected error %q, got %q", tt.expectedError, err.Error())
				}
			} else if err != nil {
				t.Errorf("Expected no error, got %q", err.Error())
			}
		})
	}
}
