// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallControlPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpcFirewallControlPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpcFirewallControlPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpcFirewallControlPolicyResponseBody) *DescribeVpcFirewallControlPolicyResponse
	GetBody() *DescribeVpcFirewallControlPolicyResponseBody
}

type DescribeVpcFirewallControlPolicyResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcFirewallControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcFirewallControlPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallControlPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpcFirewallControlPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpcFirewallControlPolicyResponse) GetBody() *DescribeVpcFirewallControlPolicyResponseBody {
	return s.Body
}

func (s *DescribeVpcFirewallControlPolicyResponse) SetHeaders(v map[string]*string) *DescribeVpcFirewallControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponse) SetStatusCode(v int32) *DescribeVpcFirewallControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponse) SetBody(v *DescribeVpcFirewallControlPolicyResponseBody) *DescribeVpcFirewallControlPolicyResponse {
	s.Body = v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponse) Validate() error {
	return dara.Validate(s)
}
