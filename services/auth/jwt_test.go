package auth

import "testing"

func TestGenerateJWT(t *testing.T) {
	token, err := GenerateToken(1)
	if err != nil {
		t.Errorf("error creating JWT: %v", err)
	}

	if token == "" {
		t.Errorf("expected token to be not empty")
	}
}
