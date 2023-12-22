package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name           string
		headers        http.Header
		expectedAPIKey string
		expectError    bool
	}{
		{
			name: "valid API key",
			headers: http.Header{
				"Authorization": []string{"ApiKey 12345"},
			},
			expectedAPIKey: "12345",
			expectError:    false,
		},
		{
			name: "no auth header",
			headers: http.Header{
				// No Authorization header
			},
			expectedAPIKey: "",
			expectError:    true,
		},
		{
			name: "malformed auth header",
			headers: http.Header{
				"Authorization": []string{"InvalidKey"},
			},
			expectedAPIKey: "",
			expectError:    true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			apiKey, err := GetAPIKey(tc.headers)

            t.Error("Fail on purpose")

			if tc.expectError && err == nil {
				t.Errorf("expected an error but got none")
			}

			if !tc.expectError && err != nil {
				t.Errorf("did not expect an error but got %v", err)
			}

			if apiKey != tc.expectedAPIKey {
				t.Errorf("expected API key to be %v, got %v", tc.expectedAPIKey, apiKey)
			}
		})
	}
}

