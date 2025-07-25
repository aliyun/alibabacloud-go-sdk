// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpamPoolsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIpamPoolsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIpamPoolsResponse
	GetStatusCode() *int32
	SetBody(v *ListIpamPoolsResponseBody) *ListIpamPoolsResponse
	GetBody() *ListIpamPoolsResponseBody
}

type ListIpamPoolsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIpamPoolsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIpamPoolsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIpamPoolsResponse) GoString() string {
	return s.String()
}

func (s *ListIpamPoolsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIpamPoolsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIpamPoolsResponse) GetBody() *ListIpamPoolsResponseBody {
	return s.Body
}

func (s *ListIpamPoolsResponse) SetHeaders(v map[string]*string) *ListIpamPoolsResponse {
	s.Headers = v
	return s
}

func (s *ListIpamPoolsResponse) SetStatusCode(v int32) *ListIpamPoolsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIpamPoolsResponse) SetBody(v *ListIpamPoolsResponseBody) *ListIpamPoolsResponse {
	s.Body = v
	return s
}

func (s *ListIpamPoolsResponse) Validate() error {
	return dara.Validate(s)
}
