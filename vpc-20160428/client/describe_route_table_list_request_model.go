// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRouteTableListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeRouteTableListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeRouteTableListRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeRouteTableListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRouteTableListRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeRouteTableListRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeRouteTableListRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeRouteTableListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeRouteTableListRequest
	GetResourceOwnerId() *int64
	SetRouteTableId(v string) *DescribeRouteTableListRequest
	GetRouteTableId() *string
	SetRouteTableName(v string) *DescribeRouteTableListRequest
	GetRouteTableName() *string
	SetRouteTableType(v string) *DescribeRouteTableListRequest
	GetRouteTableType() *string
	SetRouterId(v string) *DescribeRouteTableListRequest
	GetRouterId() *string
	SetRouterType(v string) *DescribeRouteTableListRequest
	GetRouterType() *string
	SetTag(v []*DescribeRouteTableListRequestTag) *DescribeRouteTableListRequest
	GetTag() []*DescribeRouteTableListRequestTag
	SetVpcId(v string) *DescribeRouteTableListRequest
	GetVpcId() *string
}

type DescribeRouteTableListRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the returned page. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the VPC to which the route table belongs.
	//
	// You can call [DescribeRegions](https://www.alibabacloud.com/help/vpc/developer-reference/api-vpc-2016-04-28-describeregions) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// ap-southeast-6
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the route table belongs.
	//
	// example:
	//
	// rg-acfmxazb4ph****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the route table.
	//
	// example:
	//
	// vtb-bp145q7glnuzdvzu2****
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	// The name of the route table.
	//
	// example:
	//
	// doctest
	RouteTableName *string `json:"RouteTableName,omitempty" xml:"RouteTableName,omitempty"`
	// The type of the route table.
	//
	// 	- **System**
	//
	// 	- **Custom**
	//
	// example:
	//
	// System
	RouteTableType *string `json:"RouteTableType,omitempty" xml:"RouteTableType,omitempty"`
	// The ID of vRouter to which the route table belongs.
	//
	// example:
	//
	// vrt-bp1lhl0taikrteen8****
	RouterId *string `json:"RouterId,omitempty" xml:"RouterId,omitempty"`
	// The type of the router to which the route table belongs. Valid value:
	//
	// 	- **VRouter*	- (default): a vRouter
	//
	// 	- **VBR**: a VBR
	//
	// example:
	//
	// VRouter
	RouterType *string `json:"RouterType,omitempty" xml:"RouterType,omitempty"`
	// The tags of the resource.
	Tag []*DescribeRouteTableListRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the VPC to which the route table belongs.
	//
	// When this parameter is set, the value of **RouterType*	- is automatically assigned to **VRouter**.
	//
	// example:
	//
	// vpc-bp15zckdt37pq72****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeRouteTableListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteTableListRequest) GoString() string {
	return s.String()
}

func (s *DescribeRouteTableListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeRouteTableListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRouteTableListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRouteTableListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRouteTableListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRouteTableListRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeRouteTableListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeRouteTableListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeRouteTableListRequest) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DescribeRouteTableListRequest) GetRouteTableName() *string {
	return s.RouteTableName
}

func (s *DescribeRouteTableListRequest) GetRouteTableType() *string {
	return s.RouteTableType
}

func (s *DescribeRouteTableListRequest) GetRouterId() *string {
	return s.RouterId
}

func (s *DescribeRouteTableListRequest) GetRouterType() *string {
	return s.RouterType
}

func (s *DescribeRouteTableListRequest) GetTag() []*DescribeRouteTableListRequestTag {
	return s.Tag
}

func (s *DescribeRouteTableListRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeRouteTableListRequest) SetOwnerAccount(v string) *DescribeRouteTableListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRouteTableListRequest) SetOwnerId(v int64) *DescribeRouteTableListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRouteTableListRequest) SetPageNumber(v int32) *DescribeRouteTableListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRouteTableListRequest) SetPageSize(v int32) *DescribeRouteTableListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRouteTableListRequest) SetRegionId(v string) *DescribeRouteTableListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRouteTableListRequest) SetResourceGroupId(v string) *DescribeRouteTableListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeRouteTableListRequest) SetResourceOwnerAccount(v string) *DescribeRouteTableListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRouteTableListRequest) SetResourceOwnerId(v int64) *DescribeRouteTableListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRouteTableListRequest) SetRouteTableId(v string) *DescribeRouteTableListRequest {
	s.RouteTableId = &v
	return s
}

func (s *DescribeRouteTableListRequest) SetRouteTableName(v string) *DescribeRouteTableListRequest {
	s.RouteTableName = &v
	return s
}

func (s *DescribeRouteTableListRequest) SetRouteTableType(v string) *DescribeRouteTableListRequest {
	s.RouteTableType = &v
	return s
}

func (s *DescribeRouteTableListRequest) SetRouterId(v string) *DescribeRouteTableListRequest {
	s.RouterId = &v
	return s
}

func (s *DescribeRouteTableListRequest) SetRouterType(v string) *DescribeRouteTableListRequest {
	s.RouterType = &v
	return s
}

func (s *DescribeRouteTableListRequest) SetTag(v []*DescribeRouteTableListRequestTag) *DescribeRouteTableListRequest {
	s.Tag = v
	return s
}

func (s *DescribeRouteTableListRequest) SetVpcId(v string) *DescribeRouteTableListRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeRouteTableListRequest) Validate() error {
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

type DescribeRouteTableListRequestTag struct {
	// The value of tag N to add to the resource. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length. It cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the resource. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length. It cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeRouteTableListRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteTableListRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeRouteTableListRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeRouteTableListRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeRouteTableListRequestTag) SetKey(v string) *DescribeRouteTableListRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeRouteTableListRequestTag) SetValue(v string) *DescribeRouteTableListRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeRouteTableListRequestTag) Validate() error {
	return dara.Validate(s)
}
