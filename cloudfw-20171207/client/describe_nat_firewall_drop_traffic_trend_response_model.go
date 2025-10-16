// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatFirewallDropTrafficTrendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNatFirewallDropTrafficTrendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNatFirewallDropTrafficTrendResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNatFirewallDropTrafficTrendResponseBody) *DescribeNatFirewallDropTrafficTrendResponse
	GetBody() *DescribeNatFirewallDropTrafficTrendResponseBody
}

type DescribeNatFirewallDropTrafficTrendResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNatFirewallDropTrafficTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNatFirewallDropTrafficTrendResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatFirewallDropTrafficTrendResponse) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallDropTrafficTrendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNatFirewallDropTrafficTrendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNatFirewallDropTrafficTrendResponse) GetBody() *DescribeNatFirewallDropTrafficTrendResponseBody {
	return s.Body
}

func (s *DescribeNatFirewallDropTrafficTrendResponse) SetHeaders(v map[string]*string) *DescribeNatFirewallDropTrafficTrendResponse {
	s.Headers = v
	return s
}

func (s *DescribeNatFirewallDropTrafficTrendResponse) SetStatusCode(v int32) *DescribeNatFirewallDropTrafficTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNatFirewallDropTrafficTrendResponse) SetBody(v *DescribeNatFirewallDropTrafficTrendResponseBody) *DescribeNatFirewallDropTrafficTrendResponse {
	s.Body = v
	return s
}

func (s *DescribeNatFirewallDropTrafficTrendResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
