// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLoadBalancerRegionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLoadBalancerRegionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLoadBalancerRegionsResponse
	GetStatusCode() *int32
	SetBody(v *ListLoadBalancerRegionsResponseBody) *ListLoadBalancerRegionsResponse
	GetBody() *ListLoadBalancerRegionsResponseBody
}

type ListLoadBalancerRegionsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLoadBalancerRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLoadBalancerRegionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLoadBalancerRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListLoadBalancerRegionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLoadBalancerRegionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLoadBalancerRegionsResponse) GetBody() *ListLoadBalancerRegionsResponseBody {
	return s.Body
}

func (s *ListLoadBalancerRegionsResponse) SetHeaders(v map[string]*string) *ListLoadBalancerRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListLoadBalancerRegionsResponse) SetStatusCode(v int32) *ListLoadBalancerRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLoadBalancerRegionsResponse) SetBody(v *ListLoadBalancerRegionsResponseBody) *ListLoadBalancerRegionsResponse {
	s.Body = v
	return s
}

func (s *ListLoadBalancerRegionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
