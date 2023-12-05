package main

import (
	"bytes"
	"fyne.io/fyne/v2/test"
	"gold/repository"
	"io"
	"net/http"
	"os"
	"testing"
)

var testApp Config

func TestMain(m *testing.M) {
	a := test.NewApp()
	testApp.App = a
	testApp.MainWindow = a.NewWindow("")
	testApp.HTTPClient = client
	testApp.DB = repository.NewTestRepository()
	os.Exit(m.Run())
}

var jsonToReturn = `
{"ts":1701408947578,"tsj":1701408946249,"date":"Dec 1st 2023, 12:35:46 am NY","items":[{"curr":"AUD","xauPrice":3086.7094,"xagPrice":38.2712,"chgXau":5.5118,"chgXag":0.0939,"pcXau":0.1789,"pcXag":0.246,"xauClose":3081.19744,"xagClose":38.17727}]}
`

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

var client = NewTestClient(func(req *http.Request) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
		Header:     make(http.Header),
	}
})
