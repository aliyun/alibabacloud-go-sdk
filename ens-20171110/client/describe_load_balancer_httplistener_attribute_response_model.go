// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancerHTTPListenerAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLoadBalancerHTTPListenerAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLoadBalancerHTTPListenerAttributeResponseBody) *DescribeLoadBalancerHTTPListenerAttributeResponse
	GetBody() *DescribeLoadBalancerHTTPListenerAttributeResponseBody
}

type DescribeLoadBalancerHTTPListenerAttributeResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLoadBalancerHTTPListenerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLoadBalancerHTTPListenerAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerHTTPListenerAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponse) GetBody() *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	return s.Body
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponse) SetHeaders(v map[string]*string) *DescribeLoadBalancerHTTPListenerAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponse) SetStatusCode(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponse) SetBody(v *DescribeLoadBalancerHTTPListenerAttributeResponseBody) *DescribeLoadBalancerHTTPListenerAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponse) Validate() error {
	return dara.Validate(s)
}
