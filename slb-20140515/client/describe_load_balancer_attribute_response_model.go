// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancerAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLoadBalancerAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLoadBalancerAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLoadBalancerAttributeResponseBody) *DescribeLoadBalancerAttributeResponse
	GetBody() *DescribeLoadBalancerAttributeResponseBody
}

type DescribeLoadBalancerAttributeResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLoadBalancerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLoadBalancerAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLoadBalancerAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLoadBalancerAttributeResponse) GetBody() *DescribeLoadBalancerAttributeResponseBody {
	return s.Body
}

func (s *DescribeLoadBalancerAttributeResponse) SetHeaders(v map[string]*string) *DescribeLoadBalancerAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeLoadBalancerAttributeResponse) SetStatusCode(v int32) *DescribeLoadBalancerAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponse) SetBody(v *DescribeLoadBalancerAttributeResponseBody) *DescribeLoadBalancerAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeLoadBalancerAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
