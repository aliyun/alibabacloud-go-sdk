// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallPolicyPriorUsedResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnd(v int32) *DescribeVpcFirewallPolicyPriorUsedResponseBody
	GetEnd() *int32
	SetRequestId(v string) *DescribeVpcFirewallPolicyPriorUsedResponseBody
	GetRequestId() *string
	SetStart(v int32) *DescribeVpcFirewallPolicyPriorUsedResponseBody
	GetStart() *int32
}

type DescribeVpcFirewallPolicyPriorUsedResponseBody struct {
	// The lowest priority for the access control policies.
	//
	// example:
	//
	// 150
	End *int32 `json:"End,omitempty" xml:"End,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The highest priority for the access control policies.
	//
	// example:
	//
	// 1
	Start *int32 `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s DescribeVpcFirewallPolicyPriorUsedResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallPolicyPriorUsedResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallPolicyPriorUsedResponseBody) GetEnd() *int32 {
	return s.End
}

func (s *DescribeVpcFirewallPolicyPriorUsedResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpcFirewallPolicyPriorUsedResponseBody) GetStart() *int32 {
	return s.Start
}

func (s *DescribeVpcFirewallPolicyPriorUsedResponseBody) SetEnd(v int32) *DescribeVpcFirewallPolicyPriorUsedResponseBody {
	s.End = &v
	return s
}

func (s *DescribeVpcFirewallPolicyPriorUsedResponseBody) SetRequestId(v string) *DescribeVpcFirewallPolicyPriorUsedResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallPolicyPriorUsedResponseBody) SetStart(v int32) *DescribeVpcFirewallPolicyPriorUsedResponseBody {
	s.Start = &v
	return s
}

func (s *DescribeVpcFirewallPolicyPriorUsedResponseBody) Validate() error {
	return dara.Validate(s)
}
