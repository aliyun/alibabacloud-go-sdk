// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAclRuleCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInternetInAclCount(v int32) *DescribeAclRuleCountResponseBody
	GetInternetInAclCount() *int32
	SetInternetOutAclCount(v int32) *DescribeAclRuleCountResponseBody
	GetInternetOutAclCount() *int32
	SetNatInAclCount(v int32) *DescribeAclRuleCountResponseBody
	GetNatInAclCount() *int32
	SetNatOutAclCount(v int32) *DescribeAclRuleCountResponseBody
	GetNatOutAclCount() *int32
	SetRequestId(v string) *DescribeAclRuleCountResponseBody
	GetRequestId() *string
	SetTotalAclCount(v int32) *DescribeAclRuleCountResponseBody
	GetTotalAclCount() *int32
	SetVpcAclCount(v int32) *DescribeAclRuleCountResponseBody
	GetVpcAclCount() *int32
}

type DescribeAclRuleCountResponseBody struct {
	// The number of inbound access control policies for the Internet firewall.
	//
	// example:
	//
	// 2
	InternetInAclCount *int32 `json:"InternetInAclCount,omitempty" xml:"InternetInAclCount,omitempty"`
	// The number of outbound access control policies that are created for the Internet firewall.
	//
	// example:
	//
	// 3
	InternetOutAclCount *int32 `json:"InternetOutAclCount,omitempty" xml:"InternetOutAclCount,omitempty"`
	// Deprecated
	//
	// The number of inbound access control policies for the NAT firewall.	Notice: This field is deprecated..
	//
	// example:
	//
	// 0
	NatInAclCount *int32 `json:"NatInAclCount,omitempty" xml:"NatInAclCount,omitempty"`
	// The number of internal-to-external access control policies for the NAT firewall.
	//
	// example:
	//
	// 2
	NatOutAclCount *int32 `json:"NatOutAclCount,omitempty" xml:"NatOutAclCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 27936D6C-1B7A-5A5A-B9E4-FBEBBDAA****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of access control policies.
	//
	// example:
	//
	// 8
	TotalAclCount *int32 `json:"TotalAclCount,omitempty" xml:"TotalAclCount,omitempty"`
	// The number of access control policies for the VPC firewall.
	//
	// example:
	//
	// 3
	VpcAclCount *int32 `json:"VpcAclCount,omitempty" xml:"VpcAclCount,omitempty"`
}

func (s DescribeAclRuleCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclRuleCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAclRuleCountResponseBody) GetInternetInAclCount() *int32 {
	return s.InternetInAclCount
}

func (s *DescribeAclRuleCountResponseBody) GetInternetOutAclCount() *int32 {
	return s.InternetOutAclCount
}

func (s *DescribeAclRuleCountResponseBody) GetNatInAclCount() *int32 {
	return s.NatInAclCount
}

func (s *DescribeAclRuleCountResponseBody) GetNatOutAclCount() *int32 {
	return s.NatOutAclCount
}

func (s *DescribeAclRuleCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAclRuleCountResponseBody) GetTotalAclCount() *int32 {
	return s.TotalAclCount
}

func (s *DescribeAclRuleCountResponseBody) GetVpcAclCount() *int32 {
	return s.VpcAclCount
}

func (s *DescribeAclRuleCountResponseBody) SetInternetInAclCount(v int32) *DescribeAclRuleCountResponseBody {
	s.InternetInAclCount = &v
	return s
}

func (s *DescribeAclRuleCountResponseBody) SetInternetOutAclCount(v int32) *DescribeAclRuleCountResponseBody {
	s.InternetOutAclCount = &v
	return s
}

func (s *DescribeAclRuleCountResponseBody) SetNatInAclCount(v int32) *DescribeAclRuleCountResponseBody {
	s.NatInAclCount = &v
	return s
}

func (s *DescribeAclRuleCountResponseBody) SetNatOutAclCount(v int32) *DescribeAclRuleCountResponseBody {
	s.NatOutAclCount = &v
	return s
}

func (s *DescribeAclRuleCountResponseBody) SetRequestId(v string) *DescribeAclRuleCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAclRuleCountResponseBody) SetTotalAclCount(v int32) *DescribeAclRuleCountResponseBody {
	s.TotalAclCount = &v
	return s
}

func (s *DescribeAclRuleCountResponseBody) SetVpcAclCount(v int32) *DescribeAclRuleCountResponseBody {
	s.VpcAclCount = &v
	return s
}

func (s *DescribeAclRuleCountResponseBody) Validate() error {
	return dara.Validate(s)
}
