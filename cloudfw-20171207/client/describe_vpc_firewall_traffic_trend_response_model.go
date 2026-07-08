// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallTrafficTrendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpcFirewallTrafficTrendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpcFirewallTrafficTrendResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpcFirewallTrafficTrendResponseBody) *DescribeVpcFirewallTrafficTrendResponse
	GetBody() *DescribeVpcFirewallTrafficTrendResponseBody
}

type DescribeVpcFirewallTrafficTrendResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcFirewallTrafficTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcFirewallTrafficTrendResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallTrafficTrendResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallTrafficTrendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpcFirewallTrafficTrendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpcFirewallTrafficTrendResponse) GetBody() *DescribeVpcFirewallTrafficTrendResponseBody {
	return s.Body
}

func (s *DescribeVpcFirewallTrafficTrendResponse) SetHeaders(v map[string]*string) *DescribeVpcFirewallTrafficTrendResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcFirewallTrafficTrendResponse) SetStatusCode(v int32) *DescribeVpcFirewallTrafficTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcFirewallTrafficTrendResponse) SetBody(v *DescribeVpcFirewallTrafficTrendResponseBody) *DescribeVpcFirewallTrafficTrendResponse {
	s.Body = v
	return s
}

func (s *DescribeVpcFirewallTrafficTrendResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
