// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLayer7CustomPortsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLayer7CustomPortsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLayer7CustomPortsResponse
	GetStatusCode() *int32
	SetBody(v *ListLayer7CustomPortsResponseBody) *ListLayer7CustomPortsResponse
	GetBody() *ListLayer7CustomPortsResponseBody
}

type ListLayer7CustomPortsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLayer7CustomPortsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLayer7CustomPortsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLayer7CustomPortsResponse) GoString() string {
	return s.String()
}

func (s *ListLayer7CustomPortsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLayer7CustomPortsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLayer7CustomPortsResponse) GetBody() *ListLayer7CustomPortsResponseBody {
	return s.Body
}

func (s *ListLayer7CustomPortsResponse) SetHeaders(v map[string]*string) *ListLayer7CustomPortsResponse {
	s.Headers = v
	return s
}

func (s *ListLayer7CustomPortsResponse) SetStatusCode(v int32) *ListLayer7CustomPortsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLayer7CustomPortsResponse) SetBody(v *ListLayer7CustomPortsResponseBody) *ListLayer7CustomPortsResponse {
	s.Body = v
	return s
}

func (s *ListLayer7CustomPortsResponse) Validate() error {
	return dara.Validate(s)
}
