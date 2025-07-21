// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirtualHostsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVirtualHostsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVirtualHostsResponse
	GetStatusCode() *int32
	SetBody(v *ListVirtualHostsResponseBody) *ListVirtualHostsResponse
	GetBody() *ListVirtualHostsResponseBody
}

type ListVirtualHostsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVirtualHostsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVirtualHostsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVirtualHostsResponse) GoString() string {
	return s.String()
}

func (s *ListVirtualHostsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVirtualHostsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVirtualHostsResponse) GetBody() *ListVirtualHostsResponseBody {
	return s.Body
}

func (s *ListVirtualHostsResponse) SetHeaders(v map[string]*string) *ListVirtualHostsResponse {
	s.Headers = v
	return s
}

func (s *ListVirtualHostsResponse) SetStatusCode(v int32) *ListVirtualHostsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVirtualHostsResponse) SetBody(v *ListVirtualHostsResponseBody) *ListVirtualHostsResponse {
	s.Body = v
	return s
}

func (s *ListVirtualHostsResponse) Validate() error {
	return dara.Validate(s)
}
