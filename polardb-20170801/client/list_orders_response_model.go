// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOrdersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOrdersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOrdersResponse
	GetStatusCode() *int32
	SetBody(v *ListOrdersResponseBody) *ListOrdersResponse
	GetBody() *ListOrdersResponseBody
}

type ListOrdersResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOrdersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOrdersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOrdersResponse) GoString() string {
	return s.String()
}

func (s *ListOrdersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOrdersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOrdersResponse) GetBody() *ListOrdersResponseBody {
	return s.Body
}

func (s *ListOrdersResponse) SetHeaders(v map[string]*string) *ListOrdersResponse {
	s.Headers = v
	return s
}

func (s *ListOrdersResponse) SetStatusCode(v int32) *ListOrdersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOrdersResponse) SetBody(v *ListOrdersResponseBody) *ListOrdersResponse {
	s.Body = v
	return s
}

func (s *ListOrdersResponse) Validate() error {
	return dara.Validate(s)
}
