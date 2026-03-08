package auth

import (
	"testing"
	"reflect"
	"net/http"
)

func TestGetAPIKeyEmpty(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authorization", "")
	_, got := GetAPIKey(headers)
	want := "no authorization header included"
	if !reflect.DeepEqual(want, got.Error()) {
		t.Fatalf("expected: %v, got %v", want, got)
	}
}

func TestGetAPIKeyValid(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authorization", "ApiKey testKey")
	got, _ := GetAPIKey(headers)
	want := "testKey"
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got %v", want, got)
	}
}

func TestGetAPIKeyInvalid(t *testing.T) {
	headers := http.Header()
	headers.Add("Authorization", "testKey")
	_, got := GetAPIKey(headers)
	want := "malformed authorization header"
	if !reflect.DeepEqual(want, got.Error()) {
		t.Fatalf("expected: %v, got %v", want, got)
	}
}
