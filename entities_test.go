package ycfrr

import (
	"encoding/json"
	"testing"
)

type CounterTestEntity struct {
	Value int64  `json:"value"`
	Delta int64  `json:"delta"`
	Name  string `json:"name"`
}

func TestRequest_UnmarshallJSON(t *testing.T) {
	tests := map[string]struct {
		req    Request
		expErr bool
	}{
		"empty body": {
			req: Request{
				Body: nil,
			},
			expErr: true,
		},
		"yandex json encoded body": {
			req: Request{
				Body: json.RawMessage{34, 123, 92, 110, 32, 32, 32, 32, 92, 34, 110, 97, 109, 101, 92, 34, 58, 32, 92, 34, 116, 101, 115, 116, 92, 34, 44, 92, 110, 32, 32, 32, 32, 92, 34, 118, 97, 108, 117, 101, 92, 34, 58, 32, 49, 48, 44, 92, 110, 32, 32, 32, 32, 92, 34, 100, 101, 108, 116, 97, 92, 34, 58, 32, 49, 92, 110, 125, 34},
			},
			expErr: false,
		},
	}
	for n, c := range tests {
		t.Run(n, func(t *testing.T) {
			var v CounterTestEntity
			err := c.req.UnmarshallJSON(&v)
			if c.expErr && err == nil {
				t.Errorf("Expected to receive an error but everything is OK")
			} else if !c.expErr && err != nil {
				t.Errorf("Expected not to receive an error but got: %v", err)
			}
		})
	}
}
