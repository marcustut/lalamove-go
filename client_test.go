package lalamove

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewClientCreation(t *testing.T) {
	errMsg := "client is able to be created without"

	tests := []struct {
		options     []ClientOption
		description string
		hasError    bool
	}{
		{options: []ClientOption{WithAPIKey("API_KEY")}, description: "%s `secret` and `baseURL`", hasError: true},
		{options: []ClientOption{WithSecret("SECRET_KEY")}, description: "%s `apiKey` and `baseURL`", hasError: true},
		{options: []ClientOption{WithBaseURL("https://sandbox-rest.lalamove.com")}, description: "%s `apiKey` and `secret`", hasError: true},
		{options: []ClientOption{WithAPIKey("API_KEY"), WithSecret("SECRET_KEY")}, description: "%s `baseURL`", hasError: true},
		{options: []ClientOption{WithAPIKey("API_KEY"), WithBaseURL("https://sandbox-rest.lalamove.com")}, description: "%s `secret`", hasError: true},
		{options: []ClientOption{WithSecret("SECRET_KEY"), WithBaseURL("https://sandbox-rest.lalamove.com")}, description: "%s `apiKey`", hasError: true},
		{options: []ClientOption{WithAPIKey("API_KEY"), WithSecret("SECRET_KEY"), WithBaseURL("https://sandbox-rest.lalamove.com")}, description: "client cannot be created even if `apiKey`, `secret` and `baseURL` are provided", hasError: false},
		{options: []ClientOption{WithHTTPClient(http.DefaultClient), WithAPIKey("API_KEY"), WithSecret("SECRET_KEY"), WithBaseURL("https://sandbox-rest.lalamove.com")}, description: "client cannot be created even if `apiKey`, `secret` and `baseURL` are provided", hasError: false},
	}

	for _, test := range tests {
		_, err := NewClient(test.options...)
		if test.hasError {
			assert.NotNilf(t, err, test.description, errMsg)
		} else {
			assert.Nilf(t, err, test.description)
		}
	}
}
