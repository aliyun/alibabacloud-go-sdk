// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpamResourceDiscoveriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIpamResourceDiscoveriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIpamResourceDiscoveriesResponse
	GetStatusCode() *int32
	SetBody(v *ListIpamResourceDiscoveriesResponseBody) *ListIpamResourceDiscoveriesResponse
	GetBody() *ListIpamResourceDiscoveriesResponseBody
}

type ListIpamResourceDiscoveriesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIpamResourceDiscoveriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIpamResourceDiscoveriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIpamResourceDiscoveriesResponse) GoString() string {
	return s.String()
}

func (s *ListIpamResourceDiscoveriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIpamResourceDiscoveriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIpamResourceDiscoveriesResponse) GetBody() *ListIpamResourceDiscoveriesResponseBody {
	return s.Body
}

func (s *ListIpamResourceDiscoveriesResponse) SetHeaders(v map[string]*string) *ListIpamResourceDiscoveriesResponse {
	s.Headers = v
	return s
}

func (s *ListIpamResourceDiscoveriesResponse) SetStatusCode(v int32) *ListIpamResourceDiscoveriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIpamResourceDiscoveriesResponse) SetBody(v *ListIpamResourceDiscoveriesResponseBody) *ListIpamResourceDiscoveriesResponse {
	s.Body = v
	return s
}

func (s *ListIpamResourceDiscoveriesResponse) Validate() error {
	return dara.Validate(s)
}
