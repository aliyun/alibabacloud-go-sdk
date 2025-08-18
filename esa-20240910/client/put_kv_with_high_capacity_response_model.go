// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutKvWithHighCapacityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutKvWithHighCapacityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutKvWithHighCapacityResponse
	GetStatusCode() *int32
	SetBody(v *PutKvWithHighCapacityResponseBody) *PutKvWithHighCapacityResponse
	GetBody() *PutKvWithHighCapacityResponseBody
}

type PutKvWithHighCapacityResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutKvWithHighCapacityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutKvWithHighCapacityResponse) String() string {
	return dara.Prettify(s)
}

func (s PutKvWithHighCapacityResponse) GoString() string {
	return s.String()
}

func (s *PutKvWithHighCapacityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutKvWithHighCapacityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutKvWithHighCapacityResponse) GetBody() *PutKvWithHighCapacityResponseBody {
	return s.Body
}

func (s *PutKvWithHighCapacityResponse) SetHeaders(v map[string]*string) *PutKvWithHighCapacityResponse {
	s.Headers = v
	return s
}

func (s *PutKvWithHighCapacityResponse) SetStatusCode(v int32) *PutKvWithHighCapacityResponse {
	s.StatusCode = &v
	return s
}

func (s *PutKvWithHighCapacityResponse) SetBody(v *PutKvWithHighCapacityResponseBody) *PutKvWithHighCapacityResponse {
	s.Body = v
	return s
}

func (s *PutKvWithHighCapacityResponse) Validate() error {
	return dara.Validate(s)
}
