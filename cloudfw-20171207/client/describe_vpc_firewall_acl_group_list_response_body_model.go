// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallAclGroupListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAclGroupList(v []*DescribeVpcFirewallAclGroupListResponseBodyAclGroupList) *DescribeVpcFirewallAclGroupListResponseBody
	GetAclGroupList() []*DescribeVpcFirewallAclGroupListResponseBodyAclGroupList
	SetRequestId(v string) *DescribeVpcFirewallAclGroupListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVpcFirewallAclGroupListResponseBody
	GetTotalCount() *int32
}

type DescribeVpcFirewallAclGroupListResponseBody struct {
	// The information about the access control policy groups of the VPC firewall.
	AclGroupList []*DescribeVpcFirewallAclGroupListResponseBodyAclGroupList `json:"AclGroupList,omitempty" xml:"AclGroupList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of access control policy groups of the VPC firewall.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeVpcFirewallAclGroupListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallAclGroupListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallAclGroupListResponseBody) GetAclGroupList() []*DescribeVpcFirewallAclGroupListResponseBodyAclGroupList {
	return s.AclGroupList
}

func (s *DescribeVpcFirewallAclGroupListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpcFirewallAclGroupListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVpcFirewallAclGroupListResponseBody) SetAclGroupList(v []*DescribeVpcFirewallAclGroupListResponseBodyAclGroupList) *DescribeVpcFirewallAclGroupListResponseBody {
	s.AclGroupList = v
	return s
}

func (s *DescribeVpcFirewallAclGroupListResponseBody) SetRequestId(v string) *DescribeVpcFirewallAclGroupListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallAclGroupListResponseBody) SetTotalCount(v int32) *DescribeVpcFirewallAclGroupListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVpcFirewallAclGroupListResponseBody) Validate() error {
	if s.AclGroupList != nil {
		for _, item := range s.AclGroupList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpcFirewallAclGroupListResponseBodyAclGroupList struct {
	// The mode of the access control engine.
	AclConfig *DescribeVpcFirewallAclGroupListResponseBodyAclGroupListAclConfig `json:"AclConfig,omitempty" xml:"AclConfig,omitempty" type:"Struct"`
	// The access control policy group ID of the virtual private cloud (VPC) firewalls at the border. Valid values:
	//
	// - If the virtual private cloud (VPC) firewalls at the border protect a Cloud Enterprise Network (CEN) instance, the policy group ID is the instance ID of the CEN instance.
	//
	//   Example: cen-ervw0g12b5jbw*\\*\\*\\*
	//
	// - If the virtual private cloud (VPC) firewalls at the border protect an Express Connect circuit, the policy group ID is the instance ID of the virtual private cloud (VPC) firewalls at the border.
	//
	//   Example: vfw-a42bbb7b887148c9*\\*\\*\\*
	//
	// example:
	//
	// vfw-a42bbb7b887148c9****
	AclGroupId *string `json:"AclGroupId,omitempty" xml:"AclGroupId,omitempty"`
	// The access control policy group name of the virtual private cloud (VPC) firewalls at the border.
	//
	// - If the virtual private cloud (VPC) firewalls at the border protect a CEN instance, the policy group name is the name of the CEN instance.
	//
	// - If the virtual private cloud (VPC) firewalls at the border protect an Express Connect circuit, the policy group name is the name of the virtual private cloud (VPC) firewalls at the border.
	//
	// example:
	//
	// group_test
	AclGroupName *string `json:"AclGroupName,omitempty" xml:"AclGroupName,omitempty"`
	// The number of policies in the access control policy group of the virtual private cloud (VPC) firewalls at the border.
	//
	// example:
	//
	// 9
	AclRuleCount *int32 `json:"AclRuleCount,omitempty" xml:"AclRuleCount,omitempty"`
	// Specifies whether the firewall is the default firewall. Valid values:
	//
	// - **true**: The firewall is the default firewall.
	//
	// - **false**: The firewall is not the default firewall.
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The ID of the member accounts of the current Alibaba Cloud account.
	//
	// example:
	//
	// 258039427902****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
}

func (s DescribeVpcFirewallAclGroupListResponseBodyAclGroupList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallAclGroupListResponseBodyAclGroupList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallAclGroupListResponseBodyAclGroupList) GetAclConfig() *DescribeVpcFirewallAclGroupListResponseBodyAclGroupListAclConfig {
	return s.AclConfig
}

func (s *DescribeVpcFirewallAclGroupListResponseBodyAclGroupList) GetAclGroupId() *string {
	return s.AclGroupId
}

func (s *DescribeVpcFirewallAclGroupListResponseBodyAclGroupList) GetAclGroupName() *string {
	return s.AclGroupName
}

func (s *DescribeVpcFirewallAclGroupListResponseBodyAclGroupList) GetAclRuleCount() *int32 {
	return s.AclRuleCount
}

func (s *DescribeVpcFirewallAclGroupListResponseBodyAclGroupList) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeVpcFirewallAclGroupListResponseBodyAclGroupList) GetMemberUid() *string {
	return s.MemberUid
}

func (s *DescribeVpcFirewallAclGroupListResponseBodyAclGroupList) SetAclConfig(v *DescribeVpcFirewallAclGroupListResponseBodyAclGroupListAclConfig) *DescribeVpcFirewallAclGroupListResponseBodyAclGroupList {
	s.AclConfig = v
	return s
}

func (s *DescribeVpcFirewallAclGroupListResponseBodyAclGroupList) SetAclGroupId(v string) *DescribeVpcFirewallAclGroupListResponseBodyAclGroupList {
	s.AclGroupId = &v
	return s
}

func (s *DescribeVpcFirewallAclGroupListResponseBodyAclGroupList) SetAclGroupName(v string) *DescribeVpcFirewallAclGroupListResponseBodyAclGroupList {
	s.AclGroupName = &v
	return s
}

func (s *DescribeVpcFirewallAclGroupListResponseBodyAclGroupList) SetAclRuleCount(v int32) *DescribeVpcFirewallAclGroupListResponseBodyAclGroupList {
	s.AclRuleCount = &v
	return s
}

func (s *DescribeVpcFirewallAclGroupListResponseBodyAclGroupList) SetIsDefault(v bool) *DescribeVpcFirewallAclGroupListResponseBodyAclGroupList {
	s.IsDefault = &v
	return s
}

func (s *DescribeVpcFirewallAclGroupListResponseBodyAclGroupList) SetMemberUid(v string) *DescribeVpcFirewallAclGroupListResponseBodyAclGroupList {
	s.MemberUid = &v
	return s
}

func (s *DescribeVpcFirewallAclGroupListResponseBodyAclGroupList) Validate() error {
	if s.AclConfig != nil {
		if err := s.AclConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVpcFirewallAclGroupListResponseBodyAclGroupListAclConfig struct {
	// Specifies whether strict mode is enabled. Valid values:
	//
	// - 1: Strict mode is enabled.
	//
	// - 0: Strict mode is disabled.
	//
	// example:
	//
	// 1
	StrictMode *int32 `json:"StrictMode,omitempty" xml:"StrictMode,omitempty"`
}

func (s DescribeVpcFirewallAclGroupListResponseBodyAclGroupListAclConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallAclGroupListResponseBodyAclGroupListAclConfig) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallAclGroupListResponseBodyAclGroupListAclConfig) GetStrictMode() *int32 {
	return s.StrictMode
}

func (s *DescribeVpcFirewallAclGroupListResponseBodyAclGroupListAclConfig) SetStrictMode(v int32) *DescribeVpcFirewallAclGroupListResponseBodyAclGroupListAclConfig {
	s.StrictMode = &v
	return s
}

func (s *DescribeVpcFirewallAclGroupListResponseBodyAclGroupListAclConfig) Validate() error {
	return dara.Validate(s)
}
