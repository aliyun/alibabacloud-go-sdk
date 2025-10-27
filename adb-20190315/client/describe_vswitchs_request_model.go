// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVSwitchsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeVSwitchsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeVSwitchsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeVSwitchsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeVSwitchsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeVSwitchsRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeVSwitchsRequest
	GetSecurityToken() *string
	SetVpcId(v string) *DescribeVSwitchsRequest
	GetVpcId() *string
	SetZoneId(v string) *DescribeVSwitchsRequest
	GetZoneId() *string
}

type DescribeVSwitchsRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the VPC ID.
	//
	// example:
	//
	// vpc-2ze1lz7nk4pn4zwy1j7pm
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/129857.html) operation to query the most recent zone list.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeVSwitchsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeVSwitchsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVSwitchsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVSwitchsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeVSwitchsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeVSwitchsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeVSwitchsRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVSwitchsRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeVSwitchsRequest) SetOwnerAccount(v string) *DescribeVSwitchsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeVSwitchsRequest) SetOwnerId(v int64) *DescribeVSwitchsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVSwitchsRequest) SetRegionId(v string) *DescribeVSwitchsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVSwitchsRequest) SetResourceOwnerAccount(v string) *DescribeVSwitchsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVSwitchsRequest) SetResourceOwnerId(v int64) *DescribeVSwitchsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVSwitchsRequest) SetSecurityToken(v string) *DescribeVSwitchsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeVSwitchsRequest) SetVpcId(v string) *DescribeVSwitchsRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeVSwitchsRequest) SetZoneId(v string) *DescribeVSwitchsRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeVSwitchsRequest) Validate() error {
	return dara.Validate(s)
}
