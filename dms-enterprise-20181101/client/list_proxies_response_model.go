// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProxiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProxiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProxiesResponse
	GetStatusCode() *int32
	SetBody(v *ListProxiesResponseBody) *ListProxiesResponse
	GetBody() *ListProxiesResponseBody
}

type ListProxiesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProxiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProxiesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProxiesResponse) GoString() string {
	return s.String()
}

func (s *ListProxiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProxiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProxiesResponse) GetBody() *ListProxiesResponseBody {
	return s.Body
}

func (s *ListProxiesResponse) SetHeaders(v map[string]*string) *ListProxiesResponse {
	s.Headers = v
	return s
}

func (s *ListProxiesResponse) SetStatusCode(v int32) *ListProxiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProxiesResponse) SetBody(v *ListProxiesResponseBody) *ListProxiesResponse {
	s.Body = v
	return s
}

func (s *ListProxiesResponse) Validate() error {
	return dara.Validate(s)
}
