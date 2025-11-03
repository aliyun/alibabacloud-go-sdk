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
	// The CIDR block of the vSwitch. Take note of the following limits:
	//
	// 	- The subnet mask of the CIDR block must be 16 to 29 bits in length.
	//
	// 	- The CIDR block of the vSwitch must fall within the CIDR block of the VPC to which the vSwitch belongs.
	//
	// 	- The CIDR block of a vSwitch cannot be the same as the destination CIDR block in a route entry of the VPC. However, it can be a subset of the destination CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// 172.16.0.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among all requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, **ClientToken*	- is set to the value of **RequestId**. The value of **RequestId*	- may be different for each API request.
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
	// The last eight bits of the IPv6 CIDR block of the vSwitch. Valid values: **0*	- to **255**.
	//
	// example:
	//
	// 12
	Ipv6CidrBlock *int32  `json:"Ipv6CidrBlock,omitempty" xml:"Ipv6CidrBlock,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the vSwitch.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tag of the resource.
	Tag []*CreateVSwitchRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The name of the vSwitch.
	//
	// The name must be 1 to 128 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// vSwitch-1
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
	// The ID of the VPC where you want to create the vSwitch.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-257gqcdfvx6n****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The IPv6 CIDR block of the VPC.
	//
	// example:
	//
	// 2408:XXXX:0:6a::/56
	VpcIpv6CidrBlock *string `json:"VpcIpv6CidrBlock,omitempty" xml:"VpcIpv6CidrBlock,omitempty"`
	// The zone ID of the vSwitch.
	//
	// You can call the [DescribeZones](https://help.aliyun.com/document_detail/36064.html) operation to query the most recent zone list.
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
	// The tag key. You can specify at most 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be at most 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify at most 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length, but cannot contain `http://` or `https://`. The tag value cannot start with `aliyun` or `acs:`.
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
