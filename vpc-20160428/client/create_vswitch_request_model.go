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
	SetClientToken(v string) *CreateVSwitchRequest
	GetClientToken() *string
	SetDescription(v string) *CreateVSwitchRequest
	GetDescription() *string
	SetIpv6CidrBlock(v int32) *CreateVSwitchRequest
	GetIpv6CidrBlock() *int32
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
	// The CIDR block of the vSwitch. The following requirements apply:
	//
	// - The mask length of the vSwitch CIDR block must be 16 to 29 bits.
	//
	// - The CIDR block of the vSwitch must be a subset of the CIDR block of the VPC to which the vSwitch belongs.
	//
	// - The CIDR block of the vSwitch cannot be the same as the destination CIDR block of a route in the VPC, but can be a subset of the destination CIDR block.
	//
	// - The CIDR block of the vSwitch cannot be within the following reserved address ranges: 100.64.0.0/10, 127.0.0.0/8, 169.254.0.0/16, or 224.0.0.0/4.
	//
	// > After a vSwitch is created, you cannot modify its CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// 172.16.0.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The ClientToken value can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may differ for each API request.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe63****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the vSwitch.
	//
	// The description must be 1 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// vSwitch
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The last 8 bits of the IPv6 CIDR block of the vSwitch. Valid values: **0*	- to **255**.
	//
	// You can specify this parameter only when the VPC to which the vSwitch belongs has IPv6 enabled. This allows you to assign an IPv6 CIDR block to the vSwitch. After the IPv6 CIDR block is allocated, it cannot be changed. Make sure that the CIDR block does not overlap with those of other vSwitches in the VPC.
	//
	// example:
	//
	// 12
	Ipv6CidrBlock *int32  `json:"Ipv6CidrBlock,omitempty" xml:"Ipv6CidrBlock,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the vSwitch that you want to create.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
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
	// The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// vSwitch-1
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
	// The ID of the VPC to which the vSwitch belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-257gqcdfvx6n****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The IPv6 CIDR block of the VPC. If the VPC to which the vSwitch belongs has multiple IPv6 CIDR blocks, you can specify this parameter to determine the IPv6 CIDR block range for the vSwitch. If you do not specify this parameter, the IPv6 CIDR block assigned when IPv6 was enabled for the VPC is used.
	//
	// example:
	//
	// 2408:XXXX:0:6a::/56
	VpcIpv6CidrBlock *string `json:"VpcIpv6CidrBlock,omitempty" xml:"VpcIpv6CidrBlock,omitempty"`
	// The ID of the zone in which to create the vSwitch.
	//
	// You can call the [DescribeZones](https://help.aliyun.com/document_detail/36064.html) operation to query the zone ID.
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

func (s *CreateVSwitchRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateVSwitchRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateVSwitchRequest) GetIpv6CidrBlock() *int32 {
	return s.Ipv6CidrBlock
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
	// The tag key of the resource. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// A tag key can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
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
