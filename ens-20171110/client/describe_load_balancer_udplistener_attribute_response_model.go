// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancerUDPListenerAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLoadBalancerUDPListenerAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLoadBalancerUDPListenerAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLoadBalancerUDPListenerAttributeResponseBody) *DescribeLoadBalancerUDPListenerAttributeResponse
	GetBody() *DescribeLoadBalancerUDPListenerAttributeResponseBody
}

type DescribeLoadBalancerUDPListenerAttributeResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLoadBalancerUDPListenerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLoadBalancerUDPListenerAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerUDPListenerAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponse) GetBody() *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	return s.Body
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponse) SetHeaders(v map[string]*string) *DescribeLoadBalancerUDPListenerAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponse) SetStatusCode(v int32) *DescribeLoadBalancerUDPListenerAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponse) SetBody(v *DescribeLoadBalancerUDPListenerAttributeResponseBody) *DescribeLoadBalancerUDPListenerAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
