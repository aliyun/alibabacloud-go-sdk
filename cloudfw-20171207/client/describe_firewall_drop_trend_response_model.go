// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFirewallDropTrendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFirewallDropTrendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFirewallDropTrendResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFirewallDropTrendResponseBody) *DescribeFirewallDropTrendResponse
	GetBody() *DescribeFirewallDropTrendResponseBody
}

type DescribeFirewallDropTrendResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFirewallDropTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFirewallDropTrendResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallDropTrendResponse) GoString() string {
	return s.String()
}

func (s *DescribeFirewallDropTrendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFirewallDropTrendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFirewallDropTrendResponse) GetBody() *DescribeFirewallDropTrendResponseBody {
	return s.Body
}

func (s *DescribeFirewallDropTrendResponse) SetHeaders(v map[string]*string) *DescribeFirewallDropTrendResponse {
	s.Headers = v
	return s
}

func (s *DescribeFirewallDropTrendResponse) SetStatusCode(v int32) *DescribeFirewallDropTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFirewallDropTrendResponse) SetBody(v *DescribeFirewallDropTrendResponseBody) *DescribeFirewallDropTrendResponse {
	s.Body = v
	return s
}

func (s *DescribeFirewallDropTrendResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
