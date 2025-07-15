// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDefaultVSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateDefaultVSwitchRequest
	GetClientToken() *string
	SetIpv6CidrBlock(v int32) *CreateDefaultVSwitchRequest
	GetIpv6CidrBlock() *int32
	SetOwnerAccount(v string) *CreateDefaultVSwitchRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateDefaultVSwitchRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateDefaultVSwitchRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateDefaultVSwitchRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDefaultVSwitchRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *CreateDefaultVSwitchRequest
	GetZoneId() *string
}

type CreateDefaultVSwitchRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe63****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The last eight bits of the IPv6 CIDR block of the vSwitch. Valid values: **0*	- to **255**.
	//
	// example:
	//
	// 12
	Ipv6CidrBlock *int32  `json:"Ipv6CidrBlock,omitempty" xml:"Ipv6CidrBlock,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the default vSwitch.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The zone ID of the default vSwitch.
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

func (s CreateDefaultVSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDefaultVSwitchRequest) GoString() string {
	return s.String()
}

func (s *CreateDefaultVSwitchRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDefaultVSwitchRequest) GetIpv6CidrBlock() *int32 {
	return s.Ipv6CidrBlock
}

func (s *CreateDefaultVSwitchRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateDefaultVSwitchRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDefaultVSwitchRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDefaultVSwitchRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDefaultVSwitchRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDefaultVSwitchRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateDefaultVSwitchRequest) SetClientToken(v string) *CreateDefaultVSwitchRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDefaultVSwitchRequest) SetIpv6CidrBlock(v int32) *CreateDefaultVSwitchRequest {
	s.Ipv6CidrBlock = &v
	return s
}

func (s *CreateDefaultVSwitchRequest) SetOwnerAccount(v string) *CreateDefaultVSwitchRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDefaultVSwitchRequest) SetOwnerId(v int64) *CreateDefaultVSwitchRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDefaultVSwitchRequest) SetRegionId(v string) *CreateDefaultVSwitchRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDefaultVSwitchRequest) SetResourceOwnerAccount(v string) *CreateDefaultVSwitchRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDefaultVSwitchRequest) SetResourceOwnerId(v int64) *CreateDefaultVSwitchRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDefaultVSwitchRequest) SetZoneId(v string) *CreateDefaultVSwitchRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateDefaultVSwitchRequest) Validate() error {
	return dara.Validate(s)
}
