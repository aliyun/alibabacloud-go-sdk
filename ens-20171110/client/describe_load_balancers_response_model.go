// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLoadBalancersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLoadBalancersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLoadBalancersResponseBody) *DescribeLoadBalancersResponse
	GetBody() *DescribeLoadBalancersResponseBody
}

type DescribeLoadBalancersResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLoadBalancersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLoadBalancersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancersResponse) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLoadBalancersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLoadBalancersResponse) GetBody() *DescribeLoadBalancersResponseBody {
	return s.Body
}

func (s *DescribeLoadBalancersResponse) SetHeaders(v map[string]*string) *DescribeLoadBalancersResponse {
	s.Headers = v
	return s
}

func (s *DescribeLoadBalancersResponse) SetStatusCode(v int32) *DescribeLoadBalancersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLoadBalancersResponse) SetBody(v *DescribeLoadBalancersResponseBody) *DescribeLoadBalancersResponse {
	s.Body = v
	return s
}

func (s *DescribeLoadBalancersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
