// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallPolicyPriorUsedResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpcFirewallPolicyPriorUsedResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpcFirewallPolicyPriorUsedResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpcFirewallPolicyPriorUsedResponseBody) *DescribeVpcFirewallPolicyPriorUsedResponse
	GetBody() *DescribeVpcFirewallPolicyPriorUsedResponseBody
}

type DescribeVpcFirewallPolicyPriorUsedResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcFirewallPolicyPriorUsedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcFirewallPolicyPriorUsedResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallPolicyPriorUsedResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallPolicyPriorUsedResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpcFirewallPolicyPriorUsedResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpcFirewallPolicyPriorUsedResponse) GetBody() *DescribeVpcFirewallPolicyPriorUsedResponseBody {
	return s.Body
}

func (s *DescribeVpcFirewallPolicyPriorUsedResponse) SetHeaders(v map[string]*string) *DescribeVpcFirewallPolicyPriorUsedResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcFirewallPolicyPriorUsedResponse) SetStatusCode(v int32) *DescribeVpcFirewallPolicyPriorUsedResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcFirewallPolicyPriorUsedResponse) SetBody(v *DescribeVpcFirewallPolicyPriorUsedResponseBody) *DescribeVpcFirewallPolicyPriorUsedResponse {
	s.Body = v
	return s
}

func (s *DescribeVpcFirewallPolicyPriorUsedResponse) Validate() error {
	return dara.Validate(s)
}
