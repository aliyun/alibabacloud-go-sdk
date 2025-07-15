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
	// The new description for the vSwitch.
	//
	// The description must be 1 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// This is my vswitch.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable the IPv6 feature for the vSwitch. Valid values:
	//
	// 	- **true**: enables the IPv6 feature.
	//
	// 	- **false**: disables the IPv6 feature. This is the default value.
	//
	// example:
	//
	// false
	EnableIPv6 *bool `json:"EnableIPv6,omitempty" xml:"EnableIPv6,omitempty"`
	// The last eight bits of the IPv6 CIDR block of the vSwitch. Valid values: **0*	- to **255**.
	//
	// You can set this parameter only when the IPv6 feature is enabled for the virtual private cloud (VPC) to which the vSwitch belongs.
	//
	// example:
	//
	// 10
	Ipv6CidrBlock *int32  `json:"Ipv6CidrBlock,omitempty" xml:"Ipv6CidrBlock,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the vSwitch is deployed. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
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
	// The new name for the vSwitch.
	//
	// The name must be 1 to 128 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// VSwitch-1
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
	// The IPv6 CIDR block of the VPC to which the vSwitch belongs.
	//
	// You can set this parameter only when the IPv6 feature is enabled for the VPC.
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
