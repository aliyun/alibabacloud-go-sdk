// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFirewallTrafficTrendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFirewallTrafficTrendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFirewallTrafficTrendResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFirewallTrafficTrendResponseBody) *DescribeFirewallTrafficTrendResponse
	GetBody() *DescribeFirewallTrafficTrendResponseBody
}

type DescribeFirewallTrafficTrendResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFirewallTrafficTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFirewallTrafficTrendResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallTrafficTrendResponse) GoString() string {
	return s.String()
}

func (s *DescribeFirewallTrafficTrendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFirewallTrafficTrendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFirewallTrafficTrendResponse) GetBody() *DescribeFirewallTrafficTrendResponseBody {
	return s.Body
}

func (s *DescribeFirewallTrafficTrendResponse) SetHeaders(v map[string]*string) *DescribeFirewallTrafficTrendResponse {
	s.Headers = v
	return s
}

func (s *DescribeFirewallTrafficTrendResponse) SetStatusCode(v int32) *DescribeFirewallTrafficTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFirewallTrafficTrendResponse) SetBody(v *DescribeFirewallTrafficTrendResponseBody) *DescribeFirewallTrafficTrendResponse {
	s.Body = v
	return s
}

func (s *DescribeFirewallTrafficTrendResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
