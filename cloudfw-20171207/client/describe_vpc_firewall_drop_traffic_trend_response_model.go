// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallDropTrafficTrendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpcFirewallDropTrafficTrendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpcFirewallDropTrafficTrendResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpcFirewallDropTrafficTrendResponseBody) *DescribeVpcFirewallDropTrafficTrendResponse
	GetBody() *DescribeVpcFirewallDropTrafficTrendResponseBody
}

type DescribeVpcFirewallDropTrafficTrendResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcFirewallDropTrafficTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcFirewallDropTrafficTrendResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallDropTrafficTrendResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDropTrafficTrendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpcFirewallDropTrafficTrendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpcFirewallDropTrafficTrendResponse) GetBody() *DescribeVpcFirewallDropTrafficTrendResponseBody {
	return s.Body
}

func (s *DescribeVpcFirewallDropTrafficTrendResponse) SetHeaders(v map[string]*string) *DescribeVpcFirewallDropTrafficTrendResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcFirewallDropTrafficTrendResponse) SetStatusCode(v int32) *DescribeVpcFirewallDropTrafficTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcFirewallDropTrafficTrendResponse) SetBody(v *DescribeVpcFirewallDropTrafficTrendResponseBody) *DescribeVpcFirewallDropTrafficTrendResponse {
	s.Body = v
	return s
}

func (s *DescribeVpcFirewallDropTrafficTrendResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
