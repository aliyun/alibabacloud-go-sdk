// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancerListenMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLoadBalancerListenMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLoadBalancerListenMonitorResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLoadBalancerListenMonitorResponseBody) *DescribeLoadBalancerListenMonitorResponse
	GetBody() *DescribeLoadBalancerListenMonitorResponseBody
}

type DescribeLoadBalancerListenMonitorResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLoadBalancerListenMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLoadBalancerListenMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerListenMonitorResponse) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerListenMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLoadBalancerListenMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLoadBalancerListenMonitorResponse) GetBody() *DescribeLoadBalancerListenMonitorResponseBody {
	return s.Body
}

func (s *DescribeLoadBalancerListenMonitorResponse) SetHeaders(v map[string]*string) *DescribeLoadBalancerListenMonitorResponse {
	s.Headers = v
	return s
}

func (s *DescribeLoadBalancerListenMonitorResponse) SetStatusCode(v int32) *DescribeLoadBalancerListenMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLoadBalancerListenMonitorResponse) SetBody(v *DescribeLoadBalancerListenMonitorResponseBody) *DescribeLoadBalancerListenMonitorResponse {
	s.Body = v
	return s
}

func (s *DescribeLoadBalancerListenMonitorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
