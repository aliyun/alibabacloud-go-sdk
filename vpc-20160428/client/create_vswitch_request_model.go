// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCidrBlock(v string) *CreateVSwitchRequest
	GetCidrBlock() *string
	SetCidrMask(v int32) *CreateVSwitchRequest
	GetCidrMask() *int32
	SetClientToken(v string) *CreateVSwitchRequest
	GetClientToken() *string
	SetDescription(v string) *CreateVSwitchRequest
	GetDescription() *string
	SetIpv6CidrBlock(v int32) *CreateVSwitchRequest
	GetIpv6CidrBlock() *int32
	SetIpv6CidrMask(v int32) *CreateVSwitchRequest
	GetIpv6CidrMask() *int32
	SetIpv6Only(v bool) *CreateVSwitchRequest
	GetIpv6Only() *bool
	SetOwnerAccount(v string) *CreateVSwitchRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateVSwitchRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateVSwitchRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateVSwitchRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateVSwitchRequest
	GetResourceOwnerId() *int64
	SetTag(v []*CreateVSwitchRequestTag) *CreateVSwitchRequest
	GetTag() []*CreateVSwitchRequestTag
	SetVSwitchName(v string) *CreateVSwitchRequest
	GetVSwitchName() *string
	SetVpcId(v string) *CreateVSwitchRequest
	GetVpcId() *string
	SetVpcIpv6CidrBlock(v string) *CreateVSwitchRequest
	GetVpcIpv6CidrBlock() *string
	SetZoneId(v string) *CreateVSwitchRequest
	GetZoneId() *string
}

type CreateVSwitchRequest struct {
	// The CIDR block of the vSwitch. The CIDR block of the vSwitch must meet the following requirements:
	//
	// - The mask length of the vSwitch CIDR block ranges from 16 to 29 bits.
	//
	// - The vSwitch CIDR block must be a subset of the VPC CIDR block.
	//
	// - The vSwitch CIDR block cannot be the same as the destination CIDR block of a route in the VPC. However, it can be a subset of the destination CIDR block.
	//
	// - The vSwitch CIDR block cannot fall within the following reserved address ranges: 100.64.0.0/10, 127.0.0.0/8, 169.254.0.0/16, and 224.0.0.0/4.
	//
	// > After a vSwitch is created, you cannot modify its CIDR block.
	//
	// example:
	//
	// 172.16.0.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The IPv4 CIDR block mask of the vSwitch.
	//
	// > The mask length of the IPv4 CIDR block of the vSwitch ranges from 16 to 29 bits. You must specify at least one of CidrBlock and CidrMask.
	//
	// example:
	//
	// 24
	CidrMask *int32 `json:"CidrMask,omitempty" xml:"CidrMask,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that the value is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe63****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the vSwitch.
	//
	// The description must be 1 to 256 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// vSwitch
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The last 8 bits of the IPv6 CIDR block of the vSwitch. Valid values: **0*	- to **255**.
	//
	// You can specify this parameter to assign an IPv6 CIDR block to the vSwitch only when IPv6 is enabled for the VPC to which the vSwitch belongs. After the IPv6 CIDR block is assigned, you cannot modify it. Make sure that the IPv6 CIDR block does not overlap with the CIDR blocks of other vSwitches in the same VPC.
	//
	// example:
	//
	// 12
	Ipv6CidrBlock *int32 `json:"Ipv6CidrBlock,omitempty" xml:"Ipv6CidrBlock,omitempty"`
	// The subnet mask of the IPv6 CIDR block of the vSwitch. You can specify this parameter to assign an IPv6 CIDR block to the vSwitch only when IPv6 is enabled for the VPC to which the vSwitch belongs.
	//
	// > Only 64 is supported.
	//
	// example:
	//
	// 64
	Ipv6CidrMask *int32 `json:"Ipv6CidrMask,omitempty" xml:"Ipv6CidrMask,omitempty"`
	// Specifies whether to create an IPv6-only vSwitch. Valid values:
	//
	// - false (default): creates a regular vSwitch (not IPv6-only).
	//
	// - true: creates an IPv6-only vSwitch.
	//
	// example:
	//
	// false
	Ipv6Only     *bool   `json:"Ipv6Only,omitempty" xml:"Ipv6Only,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the vSwitch to be created.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags of the resource.
	Tag []*CreateVSwitchRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The name of the vSwitch.
	//
	// The name must be 1 to 128 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// vSwitch-1
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
	// The ID of the VPC to which the vSwitch to be created belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-257gqcdfvx6n****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The IPv6 CIDR block of the VPC. If the VPC to which the vSwitch belongs has multiple IPv6 CIDR blocks, you can specify this parameter to specify the IPv6 CIDR block to which the vSwitch belongs. If you do not specify this parameter, the IPv6 CIDR block assigned when IPv6 is enabled for the VPC is used.
	//
	// example:
	//
	// 2408:XXXX:0:6a::/56
	VpcIpv6CidrBlock *string `json:"VpcIpv6CidrBlock,omitempty" xml:"VpcIpv6CidrBlock,omitempty"`
	// The zone ID of the vSwitch to be created.
	//
	// You can call the [DescribeZones](https://help.aliyun.com/document_detail/36064.html) operation to query zone IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateVSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVSwitchRequest) GoString() string {
	return s.String()
}

func (s *CreateVSwitchRequest) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *CreateVSwitchRequest) GetCidrMask() *int32 {
	return s.CidrMask
}

func (s *CreateVSwitchRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateVSwitchRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateVSwitchRequest) GetIpv6CidrBlock() *int32 {
	return s.Ipv6CidrBlock
}

func (s *CreateVSwitchRequest) GetIpv6CidrMask() *int32 {
	return s.Ipv6CidrMask
}

func (s *CreateVSwitchRequest) GetIpv6Only() *bool {
	return s.Ipv6Only
}

func (s *CreateVSwitchRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateVSwitchRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateVSwitchRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateVSwitchRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateVSwitchRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateVSwitchRequest) GetTag() []*CreateVSwitchRequestTag {
	return s.Tag
}

func (s *CreateVSwitchRequest) GetVSwitchName() *string {
	return s.VSwitchName
}

func (s *CreateVSwitchRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateVSwitchRequest) GetVpcIpv6CidrBlock() *string {
	return s.VpcIpv6CidrBlock
}

func (s *CreateVSwitchRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateVSwitchRequest) SetCidrBlock(v string) *CreateVSwitchRequest {
	s.CidrBlock = &v
	return s
}

func (s *CreateVSwitchRequest) SetCidrMask(v int32) *CreateVSwitchRequest {
	s.CidrMask = &v
	return s
}

func (s *CreateVSwitchRequest) SetClientToken(v string) *CreateVSwitchRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateVSwitchRequest) SetDescription(v string) *CreateVSwitchRequest {
	s.Description = &v
	return s
}

func (s *CreateVSwitchRequest) SetIpv6CidrBlock(v int32) *CreateVSwitchRequest {
	s.Ipv6CidrBlock = &v
	return s
}

func (s *CreateVSwitchRequest) SetIpv6CidrMask(v int32) *CreateVSwitchRequest {
	s.Ipv6CidrMask = &v
	return s
}

func (s *CreateVSwitchRequest) SetIpv6Only(v bool) *CreateVSwitchRequest {
	s.Ipv6Only = &v
	return s
}

func (s *CreateVSwitchRequest) SetOwnerAccount(v string) *CreateVSwitchRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateVSwitchRequest) SetOwnerId(v int64) *CreateVSwitchRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateVSwitchRequest) SetRegionId(v string) *CreateVSwitchRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVSwitchRequest) SetResourceOwnerAccount(v string) *CreateVSwitchRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateVSwitchRequest) SetResourceOwnerId(v int64) *CreateVSwitchRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateVSwitchRequest) SetTag(v []*CreateVSwitchRequestTag) *CreateVSwitchRequest {
	s.Tag = v
	return s
}

func (s *CreateVSwitchRequest) SetVSwitchName(v string) *CreateVSwitchRequest {
	s.VSwitchName = &v
	return s
}

func (s *CreateVSwitchRequest) SetVpcId(v string) *CreateVSwitchRequest {
	s.VpcId = &v
	return s
}

func (s *CreateVSwitchRequest) SetVpcIpv6CidrBlock(v string) *CreateVSwitchRequest {
	s.VpcIpv6CidrBlock = &v
	return s
}

func (s *CreateVSwitchRequest) SetZoneId(v string) *CreateVSwitchRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateVSwitchRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateVSwitchRequestTag struct {
	// The tag key of the resource. You can specify at most 20 tag keys. If you specify this parameter, you cannot specify an empty string.
	//
	// A tag key can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify at most 20 tag values. If you specify this parameter, you can specify an empty string.
	//
	// A tag value can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateVSwitchRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateVSwitchRequestTag) GoString() string {
	return s.String()
}

func (s *CreateVSwitchRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateVSwitchRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateVSwitchRequestTag) SetKey(v string) *CreateVSwitchRequestTag {
	s.Key = &v
	return s
}

func (s *CreateVSwitchRequestTag) SetValue(v string) *CreateVSwitchRequestTag {
	s.Value = &v
	return s
}

func (s *CreateVSwitchRequestTag) Validate() error {
	return dara.Validate(s)
}
