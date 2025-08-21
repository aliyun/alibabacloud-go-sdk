// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancerSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLoadBalancerSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLoadBalancerSpecResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLoadBalancerSpecResponseBody) *DescribeLoadBalancerSpecResponse
	GetBody() *DescribeLoadBalancerSpecResponseBody
}

type DescribeLoadBalancerSpecResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLoadBalancerSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLoadBalancerSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerSpecResponse) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLoadBalancerSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLoadBalancerSpecResponse) GetBody() *DescribeLoadBalancerSpecResponseBody {
	return s.Body
}

func (s *DescribeLoadBalancerSpecResponse) SetHeaders(v map[string]*string) *DescribeLoadBalancerSpecResponse {
	s.Headers = v
	return s
}

func (s *DescribeLoadBalancerSpecResponse) SetStatusCode(v int32) *DescribeLoadBalancerSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLoadBalancerSpecResponse) SetBody(v *DescribeLoadBalancerSpecResponseBody) *DescribeLoadBalancerSpecResponse {
	s.Body = v
	return s
}

func (s *DescribeLoadBalancerSpecResponse) Validate() error {
	return dara.Validate(s)
}
