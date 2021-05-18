This package contains requets and response structures that Yandex Cloud Function environment passes to your function and accept as a result.

# Installation

```bash
go get github.com/iamsalnikov/ycfrr
```

# Usage

```go
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/iamsalnikov/ycfrr"
)

func Handler(req []byte) (*Response, error) {
	request := &ycfrr.Request{}
	// Массив байтов, содержащий тело запроса, преобразуется в соответствующий объект
	err := json.Unmarshal(req, &request)
	if err != nil {
		return nil, fmt.Errorf("an error has occurred when parsing request: %v", err)
	}

	if req == nil {
		return &ycfrr.Response{
			StatusCode: http.StatusBadRequest,
		}, errors.New("empty request")
	}

	resp := &ycfrr.Response{StatusCode: http.StatusOK, Body: *request}
	resp.Headers = url.Values{}
	resp.Headers.Add("Content-Type", "application/json")
	return resp, nil
}
```