// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRdsVSwitchsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeRdsVSwitchsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeRdsVSwitchsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeRdsVSwitchsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeRdsVSwitchsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeRdsVSwitchsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeRdsVSwitchsRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeRdsVSwitchsRequest
	GetSecurityToken() *string
	SetVpcId(v string) *DescribeRdsVSwitchsRequest
	GetVpcId() *string
	SetZoneId(v string) *DescribeRdsVSwitchsRequest
	GetZoneId() *string
}

type DescribeRdsVSwitchsRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query the most recent region list and zone list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is no longer used.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The virtual private cloud (VPC) ID of the instance.
	//
	// >
	//
	// 	- You can call the [DescribeRdsVpcs](https://help.aliyun.com/document_detail/208327.html) operation to query the available VPC IDs.
	//
	// 	- This parameter must be specified.
	//
	// example:
	//
	// vpc-bp*******************
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the zone.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query the most recent region list and zone list.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRdsVSwitchsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsVSwitchsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRdsVSwitchsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeRdsVSwitchsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRdsVSwitchsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRdsVSwitchsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeRdsVSwitchsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeRdsVSwitchsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeRdsVSwitchsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeRdsVSwitchsRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeRdsVSwitchsRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeRdsVSwitchsRequest) SetOwnerAccount(v string) *DescribeRdsVSwitchsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) SetOwnerId(v int64) *DescribeRdsVSwitchsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) SetRegionId(v string) *DescribeRdsVSwitchsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) SetResourceGroupId(v string) *DescribeRdsVSwitchsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) SetResourceOwnerAccount(v string) *DescribeRdsVSwitchsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) SetResourceOwnerId(v int64) *DescribeRdsVSwitchsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) SetSecurityToken(v string) *DescribeRdsVSwitchsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) SetVpcId(v string) *DescribeRdsVSwitchsRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) SetZoneId(v string) *DescribeRdsVSwitchsRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) Validate() error {
	return dara.Validate(s)
}
