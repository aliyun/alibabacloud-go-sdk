// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRouteTableAttributesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyRouteTableAttributesRequest
	GetDescription() *string
	SetOwnerAccount(v string) *ModifyRouteTableAttributesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyRouteTableAttributesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyRouteTableAttributesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyRouteTableAttributesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyRouteTableAttributesRequest
	GetResourceOwnerId() *int64
	SetRoutePropagationEnable(v bool) *ModifyRouteTableAttributesRequest
	GetRoutePropagationEnable() *bool
	SetRouteTableId(v string) *ModifyRouteTableAttributesRequest
	GetRouteTableId() *string
	SetRouteTableName(v string) *ModifyRouteTableAttributesRequest
	GetRouteTableName() *string
}

type ModifyRouteTableAttributesRequest struct {
	// The description of the route table.
	//
	// The description must be 1 to 256 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the virtual private cloud (VPC) to which the custom route table belongs.
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
	// Indicates whether to enable route propagation to receive dynamic routes. Valid values:
	//
	// - **true*	- (default): enables route propagation.
	//
	// - **false**: disables route propagation.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// true
	RoutePropagationEnable *bool `json:"RoutePropagationEnable,omitempty" xml:"RoutePropagationEnable,omitempty"`
	// The ID of the route table.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-bp145q7glnuzdvzu2****
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	// The name of the route table.
	//
	// The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// doctest
	RouteTableName *string `json:"RouteTableName,omitempty" xml:"RouteTableName,omitempty"`
}

func (s ModifyRouteTableAttributesRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRouteTableAttributesRequest) GoString() string {
	return s.String()
}

func (s *ModifyRouteTableAttributesRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyRouteTableAttributesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyRouteTableAttributesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyRouteTableAttributesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyRouteTableAttributesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyRouteTableAttributesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyRouteTableAttributesRequest) GetRoutePropagationEnable() *bool {
	return s.RoutePropagationEnable
}

func (s *ModifyRouteTableAttributesRequest) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *ModifyRouteTableAttributesRequest) GetRouteTableName() *string {
	return s.RouteTableName
}

func (s *ModifyRouteTableAttributesRequest) SetDescription(v string) *ModifyRouteTableAttributesRequest {
	s.Description = &v
	return s
}

func (s *ModifyRouteTableAttributesRequest) SetOwnerAccount(v string) *ModifyRouteTableAttributesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyRouteTableAttributesRequest) SetOwnerId(v int64) *ModifyRouteTableAttributesRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyRouteTableAttributesRequest) SetRegionId(v string) *ModifyRouteTableAttributesRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyRouteTableAttributesRequest) SetResourceOwnerAccount(v string) *ModifyRouteTableAttributesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyRouteTableAttributesRequest) SetResourceOwnerId(v int64) *ModifyRouteTableAttributesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyRouteTableAttributesRequest) SetRoutePropagationEnable(v bool) *ModifyRouteTableAttributesRequest {
	s.RoutePropagationEnable = &v
	return s
}

func (s *ModifyRouteTableAttributesRequest) SetRouteTableId(v string) *ModifyRouteTableAttributesRequest {
	s.RouteTableId = &v
	return s
}

func (s *ModifyRouteTableAttributesRequest) SetRouteTableName(v string) *ModifyRouteTableAttributesRequest {
	s.RouteTableName = &v
	return s
}

func (s *ModifyRouteTableAttributesRequest) Validate() error {
	return dara.Validate(s)
}
