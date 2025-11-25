// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatFirewallAclGroupListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNatFirewalls(v []*DescribeNatFirewallAclGroupListResponseBodyNatFirewalls) *DescribeNatFirewallAclGroupListResponseBody
	GetNatFirewalls() []*DescribeNatFirewallAclGroupListResponseBodyNatFirewalls
	SetRequestId(v string) *DescribeNatFirewallAclGroupListResponseBody
	GetRequestId() *string
}

type DescribeNatFirewallAclGroupListResponseBody struct {
	NatFirewalls []*DescribeNatFirewallAclGroupListResponseBodyNatFirewalls `json:"NatFirewalls,omitempty" xml:"NatFirewalls,omitempty" type:"Repeated"`
	// example:
	//
	// F06DE24D-6EB9-5F55-B588-7BB946DF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNatFirewallAclGroupListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatFirewallAclGroupListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallAclGroupListResponseBody) GetNatFirewalls() []*DescribeNatFirewallAclGroupListResponseBodyNatFirewalls {
	return s.NatFirewalls
}

func (s *DescribeNatFirewallAclGroupListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNatFirewallAclGroupListResponseBody) SetNatFirewalls(v []*DescribeNatFirewallAclGroupListResponseBodyNatFirewalls) *DescribeNatFirewallAclGroupListResponseBody {
	s.NatFirewalls = v
	return s
}

func (s *DescribeNatFirewallAclGroupListResponseBody) SetRequestId(v string) *DescribeNatFirewallAclGroupListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNatFirewallAclGroupListResponseBody) Validate() error {
	if s.NatFirewalls != nil {
		for _, item := range s.NatFirewalls {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNatFirewallAclGroupListResponseBodyNatFirewalls struct {
	// example:
	//
	// 32
	AclRuleCount *int32 `json:"AclRuleCount,omitempty" xml:"AclRuleCount,omitempty"`
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// example:
	//
	// ngw-2zed6z6qkd7ogc****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// example:
	//
	// ngw-test
	NatGatewayName *string `json:"NatGatewayName,omitempty" xml:"NatGatewayName,omitempty"`
	// example:
	//
	// cn-beijing
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
}

func (s DescribeNatFirewallAclGroupListResponseBodyNatFirewalls) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatFirewallAclGroupListResponseBodyNatFirewalls) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallAclGroupListResponseBodyNatFirewalls) GetAclRuleCount() *int32 {
	return s.AclRuleCount
}

func (s *DescribeNatFirewallAclGroupListResponseBodyNatFirewalls) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeNatFirewallAclGroupListResponseBodyNatFirewalls) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeNatFirewallAclGroupListResponseBodyNatFirewalls) GetNatGatewayName() *string {
	return s.NatGatewayName
}

func (s *DescribeNatFirewallAclGroupListResponseBodyNatFirewalls) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeNatFirewallAclGroupListResponseBodyNatFirewalls) SetAclRuleCount(v int32) *DescribeNatFirewallAclGroupListResponseBodyNatFirewalls {
	s.AclRuleCount = &v
	return s
}

func (s *DescribeNatFirewallAclGroupListResponseBodyNatFirewalls) SetIsDefault(v bool) *DescribeNatFirewallAclGroupListResponseBodyNatFirewalls {
	s.IsDefault = &v
	return s
}

func (s *DescribeNatFirewallAclGroupListResponseBodyNatFirewalls) SetNatGatewayId(v string) *DescribeNatFirewallAclGroupListResponseBodyNatFirewalls {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeNatFirewallAclGroupListResponseBodyNatFirewalls) SetNatGatewayName(v string) *DescribeNatFirewallAclGroupListResponseBodyNatFirewalls {
	s.NatGatewayName = &v
	return s
}

func (s *DescribeNatFirewallAclGroupListResponseBodyNatFirewalls) SetRegionNo(v string) *DescribeNatFirewallAclGroupListResponseBodyNatFirewalls {
	s.RegionNo = &v
	return s
}

func (s *DescribeNatFirewallAclGroupListResponseBodyNatFirewalls) Validate() error {
	return dara.Validate(s)
}
