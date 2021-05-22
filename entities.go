package ycfrr

import (
	"encoding/base64"
	"encoding/json"
	"net/url"
	"strconv"
)

// Request represents full request data that yandex pass into cloud function
//
// https://cloud.yandex.ru/docs/functions/concepts/function-invoke#request
type Request struct {
	HttpMethod      string          `json:"httpMethod"`
	Headers         url.Values      `json:"multiValueHeaders"`
	URL             string          `json:"url"`
	Params          url.Values      `json:"multiValueParams"`
	Query           url.Values      `json:"multiValueQueryStringParameters"`
	Body            json.RawMessage `json:"body"`
	IsBase64Encoded bool            `json:"isBase64Encoded"`
	RequestContext  struct {
		Identity struct {
			SourceIP  string `json:"sourceIp"`
			UserAgent string `json:"userAgent"`
		} `json:"identity"`
		HTTPMethod       string `json:"httpMethod"`
		RequestID        string `json:"requestId"`
		RequestTime      string `json:"requestTime"`
		RequestTimeEpoch uint64 `json:"requestTimeEpoch"`
	} `json:"requestContext"`
}

// UnmarshallJSON decodes json body to the destination (dst)
func (r Request) UnmarshallJSON(dst interface{}) error {
	var err error
	data := r.Body
	if r.IsBase64Encoded {
		data, err = base64.StdEncoding.DecodeString(string(r.Body))
		if err != nil {
			return err
		}
	}

	strData, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(strData), dst)
}

// Response contains data to be sent
//
// https://cloud.yandex.ru/docs/functions/concepts/function-invoke#response
type Response struct {
	StatusCode      int         `json:"statusCode"`
	Body            interface{} `json:"body"`
	Headers         url.Values  `json:"multiValueHeaders"`
	IsBase64Encoded bool        `json:"isBase64Encoded"`
}
