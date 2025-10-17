// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFirewallRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFirewallRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFirewallRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFirewallRulesResponseBody) *DescribeFirewallRulesResponse
	GetBody() *DescribeFirewallRulesResponseBody
}

type DescribeFirewallRulesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFirewallRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFirewallRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeFirewallRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFirewallRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFirewallRulesResponse) GetBody() *DescribeFirewallRulesResponseBody {
	return s.Body
}

func (s *DescribeFirewallRulesResponse) SetHeaders(v map[string]*string) *DescribeFirewallRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeFirewallRulesResponse) SetStatusCode(v int32) *DescribeFirewallRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFirewallRulesResponse) SetBody(v *DescribeFirewallRulesResponseBody) *DescribeFirewallRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeFirewallRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
