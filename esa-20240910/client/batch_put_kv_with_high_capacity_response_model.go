// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchPutKvWithHighCapacityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchPutKvWithHighCapacityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchPutKvWithHighCapacityResponse
	GetStatusCode() *int32
	SetBody(v *BatchPutKvWithHighCapacityResponseBody) *BatchPutKvWithHighCapacityResponse
	GetBody() *BatchPutKvWithHighCapacityResponseBody
}

type BatchPutKvWithHighCapacityResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchPutKvWithHighCapacityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchPutKvWithHighCapacityResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchPutKvWithHighCapacityResponse) GoString() string {
	return s.String()
}

func (s *BatchPutKvWithHighCapacityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchPutKvWithHighCapacityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchPutKvWithHighCapacityResponse) GetBody() *BatchPutKvWithHighCapacityResponseBody {
	return s.Body
}

func (s *BatchPutKvWithHighCapacityResponse) SetHeaders(v map[string]*string) *BatchPutKvWithHighCapacityResponse {
	s.Headers = v
	return s
}

func (s *BatchPutKvWithHighCapacityResponse) SetStatusCode(v int32) *BatchPutKvWithHighCapacityResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchPutKvWithHighCapacityResponse) SetBody(v *BatchPutKvWithHighCapacityResponseBody) *BatchPutKvWithHighCapacityResponse {
	s.Body = v
	return s
}

func (s *BatchPutKvWithHighCapacityResponse) Validate() error {
	return dara.Validate(s)
}
