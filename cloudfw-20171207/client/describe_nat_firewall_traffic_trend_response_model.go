// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatFirewallTrafficTrendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNatFirewallTrafficTrendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNatFirewallTrafficTrendResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNatFirewallTrafficTrendResponseBody) *DescribeNatFirewallTrafficTrendResponse
	GetBody() *DescribeNatFirewallTrafficTrendResponseBody
}

type DescribeNatFirewallTrafficTrendResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNatFirewallTrafficTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNatFirewallTrafficTrendResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatFirewallTrafficTrendResponse) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallTrafficTrendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNatFirewallTrafficTrendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNatFirewallTrafficTrendResponse) GetBody() *DescribeNatFirewallTrafficTrendResponseBody {
	return s.Body
}

func (s *DescribeNatFirewallTrafficTrendResponse) SetHeaders(v map[string]*string) *DescribeNatFirewallTrafficTrendResponse {
	s.Headers = v
	return s
}

func (s *DescribeNatFirewallTrafficTrendResponse) SetStatusCode(v int32) *DescribeNatFirewallTrafficTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNatFirewallTrafficTrendResponse) SetBody(v *DescribeNatFirewallTrafficTrendResponseBody) *DescribeNatFirewallTrafficTrendResponse {
	s.Body = v
	return s
}

func (s *DescribeNatFirewallTrafficTrendResponse) Validate() error {
	return dara.Validate(s)
}
