// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVSwitchesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDedicatedHostGroupId(v string) *DescribeVSwitchesRequest
	GetDedicatedHostGroupId() *string
	SetOwnerAccount(v string) *DescribeVSwitchesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeVSwitchesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeVSwitchesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVSwitchesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeVSwitchesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeVSwitchesRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeVSwitchesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeVSwitchesRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeVSwitchesRequest
	GetSecurityToken() *string
	SetVpcId(v string) *DescribeVSwitchesRequest
	GetVpcId() *string
	SetZoneId(v string) *DescribeVSwitchesRequest
	GetZoneId() *string
}

type DescribeVSwitchesRequest struct {
	// The dedicated cluster ID.
	//
	// >  You must specify at least one of the **VpcId*	- and **DedicatedHostGroupId*	- parameters.
	//
	// example:
	//
	// dhg-4n****
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: 50. The default value is 50.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the vSwitch is deployed.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the vSwitch belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4ph****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The ID of the virtual private type (VPC) to which the vSwitch belongs.
	//
	// >  You must specify at least one of the **VpcId*	- and **DedicatedHostGroupId*	- parameters.
	//
	// example:
	//
	// vpc-25cdvfeq58pl****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the zone to which the vSwitch belongs.
	//
	// example:
	//
	// cn-hangzhou-d
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeVSwitchesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchesRequest) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesRequest) GetDedicatedHostGroupId() *string {
	return s.DedicatedHostGroupId
}

func (s *DescribeVSwitchesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeVSwitchesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVSwitchesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVSwitchesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVSwitchesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVSwitchesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
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

func (s *DescribeVSwitchesRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeVSwitchesRequest) SetDedicatedHostGroupId(v string) *DescribeVSwitchesRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetOwnerAccount(v string) *DescribeVSwitchesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetOwnerId(v int64) *DescribeVSwitchesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetPageNumber(v int32) *DescribeVSwitchesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetPageSize(v int32) *DescribeVSwitchesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetRegionId(v string) *DescribeVSwitchesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetResourceGroupId(v string) *DescribeVSwitchesRequest {
	s.ResourceGroupId = &v
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

func (s *DescribeVSwitchesRequest) SetZoneId(v string) *DescribeVSwitchesRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeVSwitchesRequest) Validate() error {
	return dara.Validate(s)
}
