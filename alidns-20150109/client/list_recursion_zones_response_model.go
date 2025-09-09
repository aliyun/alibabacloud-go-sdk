// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecursionZonesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRecursionZonesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRecursionZonesResponse
	GetStatusCode() *int32
	SetBody(v *ListRecursionZonesResponseBody) *ListRecursionZonesResponse
	GetBody() *ListRecursionZonesResponseBody
}

type ListRecursionZonesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRecursionZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRecursionZonesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRecursionZonesResponse) GoString() string {
	return s.String()
}

func (s *ListRecursionZonesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRecursionZonesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRecursionZonesResponse) GetBody() *ListRecursionZonesResponseBody {
	return s.Body
}

func (s *ListRecursionZonesResponse) SetHeaders(v map[string]*string) *ListRecursionZonesResponse {
	s.Headers = v
	return s
}

func (s *ListRecursionZonesResponse) SetStatusCode(v int32) *ListRecursionZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRecursionZonesResponse) SetBody(v *ListRecursionZonesResponseBody) *ListRecursionZonesResponse {
	s.Body = v
	return s
}

func (s *ListRecursionZonesResponse) Validate() error {
	return dara.Validate(s)
}
