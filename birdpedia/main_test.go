package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"io/ioutil"
)

func TestHandler(t *testing.T) {

	r := newRouter()

	mockServer := httptest.NewServer(r)

	resp, err := http.Get(mockServer.URL + "/hello")

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be ok, got %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	respString := string(b)
	expected := "Hello World!"

	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}
}

func TestRouterFornotAllowedRequest(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)

	// request to an undefined route, like `POST /hello`
	resp, err := http.Post(mockServer.URL + "/hello", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status should be 405, got %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	respString := string(b)
	expected := ""

	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}
}

func TestRouterForNonExistentRoute(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)

	resp, err := http.Get(mockServer.URL + "/")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Status should be 404, got %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	respString := string(b)
	expected := "404 page not found\n"

	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}
}

func TestStaticFileServer(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer()
	 resp, err := http.Get(mockServer.URL + "/assets/")
	 if err != nil {
		 t.Fatal(err)
	 }

	 if resp.StatusCode != http.StatusOK {
		 t.Errorf("sStatus should be 200, got%d", resp.StatusCode)
	 }

	 contentType := resp.Header.Get("Content-Type")
	 expectedContentType := "text/html; charset=utf-8"
	 if expectedContentType != contentType {
		 t.Errorf("Wrong content type, expected %s, got %s", expectedContentType, contentType)
	 }
}