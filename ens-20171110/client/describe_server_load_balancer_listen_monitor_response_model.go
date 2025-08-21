// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServerLoadBalancerListenMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeServerLoadBalancerListenMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeServerLoadBalancerListenMonitorResponse
	GetStatusCode() *int32
	SetBody(v *DescribeServerLoadBalancerListenMonitorResponseBody) *DescribeServerLoadBalancerListenMonitorResponse
	GetBody() *DescribeServerLoadBalancerListenMonitorResponseBody
}

type DescribeServerLoadBalancerListenMonitorResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServerLoadBalancerListenMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServerLoadBalancerListenMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeServerLoadBalancerListenMonitorResponse) GoString() string {
	return s.String()
}

func (s *DescribeServerLoadBalancerListenMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeServerLoadBalancerListenMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeServerLoadBalancerListenMonitorResponse) GetBody() *DescribeServerLoadBalancerListenMonitorResponseBody {
	return s.Body
}

func (s *DescribeServerLoadBalancerListenMonitorResponse) SetHeaders(v map[string]*string) *DescribeServerLoadBalancerListenMonitorResponse {
	s.Headers = v
	return s
}

func (s *DescribeServerLoadBalancerListenMonitorResponse) SetStatusCode(v int32) *DescribeServerLoadBalancerListenMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServerLoadBalancerListenMonitorResponse) SetBody(v *DescribeServerLoadBalancerListenMonitorResponseBody) *DescribeServerLoadBalancerListenMonitorResponse {
	s.Body = v
	return s
}

func (s *DescribeServerLoadBalancerListenMonitorResponse) Validate() error {
	return dara.Validate(s)
}
