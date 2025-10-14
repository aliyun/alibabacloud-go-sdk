// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancerTCPListenerAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLoadBalancerTCPListenerAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLoadBalancerTCPListenerAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLoadBalancerTCPListenerAttributeResponseBody) *DescribeLoadBalancerTCPListenerAttributeResponse
	GetBody() *DescribeLoadBalancerTCPListenerAttributeResponseBody
}

type DescribeLoadBalancerTCPListenerAttributeResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLoadBalancerTCPListenerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLoadBalancerTCPListenerAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerTCPListenerAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponse) GetBody() *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	return s.Body
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponse) SetHeaders(v map[string]*string) *DescribeLoadBalancerTCPListenerAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponse) SetStatusCode(v int32) *DescribeLoadBalancerTCPListenerAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponse) SetBody(v *DescribeLoadBalancerTCPListenerAttributeResponseBody) *DescribeLoadBalancerTCPListenerAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
