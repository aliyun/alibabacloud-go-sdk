// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatFirewallPolicyPriorUsedResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNatFirewallPolicyPriorUsedResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNatFirewallPolicyPriorUsedResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNatFirewallPolicyPriorUsedResponseBody) *DescribeNatFirewallPolicyPriorUsedResponse
	GetBody() *DescribeNatFirewallPolicyPriorUsedResponseBody
}

type DescribeNatFirewallPolicyPriorUsedResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNatFirewallPolicyPriorUsedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNatFirewallPolicyPriorUsedResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatFirewallPolicyPriorUsedResponse) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallPolicyPriorUsedResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNatFirewallPolicyPriorUsedResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNatFirewallPolicyPriorUsedResponse) GetBody() *DescribeNatFirewallPolicyPriorUsedResponseBody {
	return s.Body
}

func (s *DescribeNatFirewallPolicyPriorUsedResponse) SetHeaders(v map[string]*string) *DescribeNatFirewallPolicyPriorUsedResponse {
	s.Headers = v
	return s
}

func (s *DescribeNatFirewallPolicyPriorUsedResponse) SetStatusCode(v int32) *DescribeNatFirewallPolicyPriorUsedResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNatFirewallPolicyPriorUsedResponse) SetBody(v *DescribeNatFirewallPolicyPriorUsedResponseBody) *DescribeNatFirewallPolicyPriorUsedResponse {
	s.Body = v
	return s
}

func (s *DescribeNatFirewallPolicyPriorUsedResponse) Validate() error {
	return dara.Validate(s)
}
