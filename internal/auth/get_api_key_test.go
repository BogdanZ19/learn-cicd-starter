package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headers     http.Header
		expectedKey string
		expectedErr bool
	}{
		{
			name:        "valid API key",
			headers:     http.Header{"Authorization": []string{"ApiKey test-key-123"}},
			expectedKey: "test-key-123",
			expectedErr: false,
		},
		{
			name:        "missing authorization header",
			headers:     http.Header{},
			expectedKey: "",
			expectedErr: true,
		},
		{
			name:        "empty authorization header",
			headers:     http.Header{"Authorization": []string{""}},
			expectedKey: "",
			expectedErr: true,
		},
		{
			name:        "malformed header - no space",
			headers:     http.Header{"Authorization": []string{"ApiKey"}},
			expectedKey: "",
			expectedErr: true,
		},
		{
			name:        "malformed header - wrong prefix",
			headers:     http.Header{"Authorization": []string{"Bearer test-key"}},
			expectedKey: "",
			expectedErr: true,
		},
		{
			name:        "malformed header - case sensitive prefix",
			headers:     http.Header{"Authorization": []string{"apikey test-key"}},
			expectedKey: "",
			expectedErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			key, err := GetAPIKey(tc.headers)

			if (err != nil) != tc.expectedErr {
				t.Errorf("GetAPIKey() error = %v, expectedErr %v", err, tc.expectedErr)
				return
			}

			if key != tc.expectedKey {
				t.Errorf("GetAPIKey() key = %v, want %v", key, tc.expectedKey)
			}
		})
	}
}
