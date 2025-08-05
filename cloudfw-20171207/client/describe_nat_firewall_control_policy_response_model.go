// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatFirewallControlPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNatFirewallControlPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNatFirewallControlPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNatFirewallControlPolicyResponseBody) *DescribeNatFirewallControlPolicyResponse
	GetBody() *DescribeNatFirewallControlPolicyResponseBody
}

type DescribeNatFirewallControlPolicyResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNatFirewallControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNatFirewallControlPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatFirewallControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallControlPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNatFirewallControlPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNatFirewallControlPolicyResponse) GetBody() *DescribeNatFirewallControlPolicyResponseBody {
	return s.Body
}

func (s *DescribeNatFirewallControlPolicyResponse) SetHeaders(v map[string]*string) *DescribeNatFirewallControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponse) SetStatusCode(v int32) *DescribeNatFirewallControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponse) SetBody(v *DescribeNatFirewallControlPolicyResponseBody) *DescribeNatFirewallControlPolicyResponse {
	s.Body = v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponse) Validate() error {
	return dara.Validate(s)
}
