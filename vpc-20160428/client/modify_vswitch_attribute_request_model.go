// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVSwitchAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyVSwitchAttributeRequest
	GetDescription() *string
	SetEnableIPv6(v bool) *ModifyVSwitchAttributeRequest
	GetEnableIPv6() *bool
	SetIpv6CidrBlock(v int32) *ModifyVSwitchAttributeRequest
	GetIpv6CidrBlock() *int32
	SetOwnerAccount(v string) *ModifyVSwitchAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyVSwitchAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyVSwitchAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyVSwitchAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyVSwitchAttributeRequest
	GetResourceOwnerId() *int64
	SetVSwitchId(v string) *ModifyVSwitchAttributeRequest
	GetVSwitchId() *string
	SetVSwitchName(v string) *ModifyVSwitchAttributeRequest
	GetVSwitchName() *string
	SetVpcIpv6CidrBlock(v string) *ModifyVSwitchAttributeRequest
	GetVpcIpv6CidrBlock() *string
}

type ModifyVSwitchAttributeRequest struct {
	// The new description of the vSwitch.
	//
	// The description must be 1 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// This is my vswitch.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable IPv6 for the vSwitch. Valid values:
	//
	// - **true**: enables IPv6. The VPC to which the vSwitch belongs must have IPv6 enabled. You must also specify Ipv6CidrBlock to assign an IPv6 CIDR block to the vSwitch.
	//
	// - **false*	- (default): disables IPv6. When you disable IPv6 for the vSwitch, make sure that no IPv6 addresses are in use. You cannot specify Ipv6CidrBlock at the same time.
	//
	// example:
	//
	// false
	EnableIPv6 *bool `json:"EnableIPv6,omitempty" xml:"EnableIPv6,omitempty"`
	// The last 8 bits of the IPv6 CIDR block of the vSwitch. Valid values: **0*	- to **255**.
	//
	// You can specify this parameter only when the VPC to which the vSwitch belongs has IPv6 enabled. This parameter is used to assign an IPv6 CIDR block to the vSwitch. After the IPv6 CIDR block is allocated, it cannot be changed to another CIDR block. Make sure that the CIDR block does not overlap with those of other vSwitches in the same VPC.
	//
	// example:
	//
	// 10
	Ipv6CidrBlock *int32  `json:"Ipv6CidrBlock,omitempty" xml:"Ipv6CidrBlock,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the vSwitch. You can call [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the vSwitch.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-25nacdfvue4****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The new name of the vSwitch.
	//
	// The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// VSwitch-1
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
	// The IPv6 CIDR block of the VPC to which the vSwitch belongs.
	//
	// If the VPC has multiple IPv6 CIDR blocks, you can specify this parameter to indicate the IPv6 CIDR block range for the vSwitch. If you do not specify this parameter, the IPv6 CIDR block assigned when IPv6 was enabled for the VPC is used.
	//
	// example:
	//
	// 2408:XXXX:312:3e00::/56
	VpcIpv6CidrBlock *string `json:"VpcIpv6CidrBlock,omitempty" xml:"VpcIpv6CidrBlock,omitempty"`
}

func (s ModifyVSwitchAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVSwitchAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyVSwitchAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyVSwitchAttributeRequest) GetEnableIPv6() *bool {
	return s.EnableIPv6
}

func (s *ModifyVSwitchAttributeRequest) GetIpv6CidrBlock() *int32 {
	return s.Ipv6CidrBlock
}

func (s *ModifyVSwitchAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyVSwitchAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyVSwitchAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyVSwitchAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyVSwitchAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyVSwitchAttributeRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ModifyVSwitchAttributeRequest) GetVSwitchName() *string {
	return s.VSwitchName
}

func (s *ModifyVSwitchAttributeRequest) GetVpcIpv6CidrBlock() *string {
	return s.VpcIpv6CidrBlock
}

func (s *ModifyVSwitchAttributeRequest) SetDescription(v string) *ModifyVSwitchAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyVSwitchAttributeRequest) SetEnableIPv6(v bool) *ModifyVSwitchAttributeRequest {
	s.EnableIPv6 = &v
	return s
}

func (s *ModifyVSwitchAttributeRequest) SetIpv6CidrBlock(v int32) *ModifyVSwitchAttributeRequest {
	s.Ipv6CidrBlock = &v
	return s
}

func (s *ModifyVSwitchAttributeRequest) SetOwnerAccount(v string) *ModifyVSwitchAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyVSwitchAttributeRequest) SetOwnerId(v int64) *ModifyVSwitchAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyVSwitchAttributeRequest) SetRegionId(v string) *ModifyVSwitchAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyVSwitchAttributeRequest) SetResourceOwnerAccount(v string) *ModifyVSwitchAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyVSwitchAttributeRequest) SetResourceOwnerId(v int64) *ModifyVSwitchAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyVSwitchAttributeRequest) SetVSwitchId(v string) *ModifyVSwitchAttributeRequest {
	s.VSwitchId = &v
	return s
}

func (s *ModifyVSwitchAttributeRequest) SetVSwitchName(v string) *ModifyVSwitchAttributeRequest {
	s.VSwitchName = &v
	return s
}

func (s *ModifyVSwitchAttributeRequest) SetVpcIpv6CidrBlock(v string) *ModifyVSwitchAttributeRequest {
	s.VpcIpv6CidrBlock = &v
	return s
}

func (s *ModifyVSwitchAttributeRequest) Validate() error {
	return dara.Validate(s)
}
