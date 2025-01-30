package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey4(t *testing.T) {
	// Arrange: Set up necessary input
	input := http.Header{}
	//input.Add("Content-Type", "application/json")
	expected := ""

	// Act: Call the function being tested
	result, status := GetAPIKey(input)

	// Assert: Compare the result to the expected output
	if result != expected {

		t.Errorf("GetAPIKey(%v) = [%v]; want %v,status %v", input, result, expected, status)
	}
}
