// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHostsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHostsResponse
	GetStatusCode() *int32
	SetBody(v *ListHostsResponseBody) *ListHostsResponse
	GetBody() *ListHostsResponseBody
}

type ListHostsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHostsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHostsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHostsResponse) GoString() string {
	return s.String()
}

func (s *ListHostsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHostsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHostsResponse) GetBody() *ListHostsResponseBody {
	return s.Body
}

func (s *ListHostsResponse) SetHeaders(v map[string]*string) *ListHostsResponse {
	s.Headers = v
	return s
}

func (s *ListHostsResponse) SetStatusCode(v int32) *ListHostsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHostsResponse) SetBody(v *ListHostsResponseBody) *ListHostsResponse {
	s.Body = v
	return s
}

func (s *ListHostsResponse) Validate() error {
	return dara.Validate(s)
}
