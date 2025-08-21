// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServerLoadBalancerMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeServerLoadBalancerMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeServerLoadBalancerMonitorResponse
	GetStatusCode() *int32
	SetBody(v *DescribeServerLoadBalancerMonitorResponseBody) *DescribeServerLoadBalancerMonitorResponse
	GetBody() *DescribeServerLoadBalancerMonitorResponseBody
}

type DescribeServerLoadBalancerMonitorResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServerLoadBalancerMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServerLoadBalancerMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeServerLoadBalancerMonitorResponse) GoString() string {
	return s.String()
}

func (s *DescribeServerLoadBalancerMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeServerLoadBalancerMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeServerLoadBalancerMonitorResponse) GetBody() *DescribeServerLoadBalancerMonitorResponseBody {
	return s.Body
}

func (s *DescribeServerLoadBalancerMonitorResponse) SetHeaders(v map[string]*string) *DescribeServerLoadBalancerMonitorResponse {
	s.Headers = v
	return s
}

func (s *DescribeServerLoadBalancerMonitorResponse) SetStatusCode(v int32) *DescribeServerLoadBalancerMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServerLoadBalancerMonitorResponse) SetBody(v *DescribeServerLoadBalancerMonitorResponseBody) *DescribeServerLoadBalancerMonitorResponse {
	s.Body = v
	return s
}

func (s *DescribeServerLoadBalancerMonitorResponse) Validate() error {
	return dara.Validate(s)
}
