package ycfrr

import (
	"encoding/json"
	"net/url"
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

// Response contains data to be sent
//
// https://cloud.yandex.ru/docs/functions/concepts/function-invoke#response
type Response struct {
	StatusCode      int        `json:"statusCode"`
	Body            Request    `json:"body"`
	Headers         url.Values `json:"multiValueHeaders"`
	IsBase64Encoded bool       `json:"isBase64Encoded"`
}
