// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEgressProxyIpPoolsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEgressProxyIpPoolsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEgressProxyIpPoolsResponse
	GetStatusCode() *int32
	SetBody(v *ListEgressProxyIpPoolsResponseBody) *ListEgressProxyIpPoolsResponse
	GetBody() *ListEgressProxyIpPoolsResponseBody
}

type ListEgressProxyIpPoolsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEgressProxyIpPoolsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEgressProxyIpPoolsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEgressProxyIpPoolsResponse) GoString() string {
	return s.String()
}

func (s *ListEgressProxyIpPoolsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEgressProxyIpPoolsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEgressProxyIpPoolsResponse) GetBody() *ListEgressProxyIpPoolsResponseBody {
	return s.Body
}

func (s *ListEgressProxyIpPoolsResponse) SetHeaders(v map[string]*string) *ListEgressProxyIpPoolsResponse {
	s.Headers = v
	return s
}

func (s *ListEgressProxyIpPoolsResponse) SetStatusCode(v int32) *ListEgressProxyIpPoolsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEgressProxyIpPoolsResponse) SetBody(v *ListEgressProxyIpPoolsResponseBody) *ListEgressProxyIpPoolsResponse {
	s.Body = v
	return s
}

func (s *ListEgressProxyIpPoolsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
