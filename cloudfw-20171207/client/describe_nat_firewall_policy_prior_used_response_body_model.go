// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatFirewallPolicyPriorUsedResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnd(v int32) *DescribeNatFirewallPolicyPriorUsedResponseBody
	GetEnd() *int32
	SetRequestId(v string) *DescribeNatFirewallPolicyPriorUsedResponseBody
	GetRequestId() *string
	SetStart(v int32) *DescribeNatFirewallPolicyPriorUsedResponseBody
	GetStart() *int32
}

type DescribeNatFirewallPolicyPriorUsedResponseBody struct {
	// The lowest priority for the access control policy.
	//
	// example:
	//
	// 28
	End *int32 `json:"End,omitempty" xml:"End,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BCDF3907-1011-5504-B015-CC7603C0E6B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The highest priority for the access control policy.
	//
	// example:
	//
	// 1
	Start *int32 `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s DescribeNatFirewallPolicyPriorUsedResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatFirewallPolicyPriorUsedResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallPolicyPriorUsedResponseBody) GetEnd() *int32 {
	return s.End
}

func (s *DescribeNatFirewallPolicyPriorUsedResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNatFirewallPolicyPriorUsedResponseBody) GetStart() *int32 {
	return s.Start
}

func (s *DescribeNatFirewallPolicyPriorUsedResponseBody) SetEnd(v int32) *DescribeNatFirewallPolicyPriorUsedResponseBody {
	s.End = &v
	return s
}

func (s *DescribeNatFirewallPolicyPriorUsedResponseBody) SetRequestId(v string) *DescribeNatFirewallPolicyPriorUsedResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNatFirewallPolicyPriorUsedResponseBody) SetStart(v int32) *DescribeNatFirewallPolicyPriorUsedResponseBody {
	s.Start = &v
	return s
}

func (s *DescribeNatFirewallPolicyPriorUsedResponseBody) Validate() error {
	return dara.Validate(s)
}
