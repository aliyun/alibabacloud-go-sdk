// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancerHTTPSListenerAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLoadBalancerHTTPSListenerAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) *DescribeLoadBalancerHTTPSListenerAttributeResponse
	GetBody() *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
}

type DescribeLoadBalancerHTTPSListenerAttributeResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLoadBalancerHTTPSListenerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLoadBalancerHTTPSListenerAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerHTTPSListenerAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponse) GetBody() *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	return s.Body
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponse) SetHeaders(v map[string]*string) *DescribeLoadBalancerHTTPSListenerAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponse) SetStatusCode(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponse) SetBody(v *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) *DescribeLoadBalancerHTTPSListenerAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
