// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutDcdnKvWithHighCapacityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutDcdnKvWithHighCapacityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutDcdnKvWithHighCapacityResponse
	GetStatusCode() *int32
	SetBody(v *PutDcdnKvWithHighCapacityResponseBody) *PutDcdnKvWithHighCapacityResponse
	GetBody() *PutDcdnKvWithHighCapacityResponseBody
}

type PutDcdnKvWithHighCapacityResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutDcdnKvWithHighCapacityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutDcdnKvWithHighCapacityResponse) String() string {
	return dara.Prettify(s)
}

func (s PutDcdnKvWithHighCapacityResponse) GoString() string {
	return s.String()
}

func (s *PutDcdnKvWithHighCapacityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutDcdnKvWithHighCapacityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutDcdnKvWithHighCapacityResponse) GetBody() *PutDcdnKvWithHighCapacityResponseBody {
	return s.Body
}

func (s *PutDcdnKvWithHighCapacityResponse) SetHeaders(v map[string]*string) *PutDcdnKvWithHighCapacityResponse {
	s.Headers = v
	return s
}

func (s *PutDcdnKvWithHighCapacityResponse) SetStatusCode(v int32) *PutDcdnKvWithHighCapacityResponse {
	s.StatusCode = &v
	return s
}

func (s *PutDcdnKvWithHighCapacityResponse) SetBody(v *PutDcdnKvWithHighCapacityResponseBody) *PutDcdnKvWithHighCapacityResponse {
	s.Body = v
	return s
}

func (s *PutDcdnKvWithHighCapacityResponse) Validate() error {
	return dara.Validate(s)
}
