// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancerListenersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLoadBalancerListenersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLoadBalancerListenersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLoadBalancerListenersResponseBody) *DescribeLoadBalancerListenersResponse
	GetBody() *DescribeLoadBalancerListenersResponseBody
}

type DescribeLoadBalancerListenersResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLoadBalancerListenersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLoadBalancerListenersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerListenersResponse) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerListenersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLoadBalancerListenersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLoadBalancerListenersResponse) GetBody() *DescribeLoadBalancerListenersResponseBody {
	return s.Body
}

func (s *DescribeLoadBalancerListenersResponse) SetHeaders(v map[string]*string) *DescribeLoadBalancerListenersResponse {
	s.Headers = v
	return s
}

func (s *DescribeLoadBalancerListenersResponse) SetStatusCode(v int32) *DescribeLoadBalancerListenersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponse) SetBody(v *DescribeLoadBalancerListenersResponseBody) *DescribeLoadBalancerListenersResponse {
	s.Body = v
	return s
}

func (s *DescribeLoadBalancerListenersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
