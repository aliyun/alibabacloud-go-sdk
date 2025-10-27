// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVSwitchesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeVSwitchesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeVSwitchesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeVSwitchesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeVSwitchesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeVSwitchesRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeVSwitchesRequest
	GetSecurityToken() *string
	SetVpcId(v string) *DescribeVSwitchesRequest
	GetVpcId() *string
	SetVswId(v string) *DescribeVSwitchesRequest
	GetVswId() *string
	SetZoneId(v string) *DescribeVSwitchesRequest
	GetZoneId() *string
}

type DescribeVSwitchesRequest struct {
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
	// vpc-bp13h7uzhulpuxvnp****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-8vbpu1teend1qcosfojlz
	VswId *string `json:"VswId,omitempty" xml:"VswId,omitempty"`
	// The zone ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/129857.html) operation to query the most recent zone list.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeVSwitchesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchesRequest) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeVSwitchesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVSwitchesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVSwitchesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeVSwitchesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeVSwitchesRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeVSwitchesRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVSwitchesRequest) GetVswId() *string {
	return s.VswId
}

func (s *DescribeVSwitchesRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeVSwitchesRequest) SetOwnerAccount(v string) *DescribeVSwitchesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetOwnerId(v int64) *DescribeVSwitchesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetRegionId(v string) *DescribeVSwitchesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetResourceOwnerAccount(v string) *DescribeVSwitchesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetResourceOwnerId(v int64) *DescribeVSwitchesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetSecurityToken(v string) *DescribeVSwitchesRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetVpcId(v string) *DescribeVSwitchesRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetVswId(v string) *DescribeVSwitchesRequest {
	s.VswId = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetZoneId(v string) *DescribeVSwitchesRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeVSwitchesRequest) Validate() error {
	return dara.Validate(s)
}
