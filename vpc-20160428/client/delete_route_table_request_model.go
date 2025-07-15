// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouteTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DeleteRouteTableRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteRouteTableRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteRouteTableRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteRouteTableRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteRouteTableRequest
	GetResourceOwnerId() *int64
	SetRouteTableId(v string) *DeleteRouteTableRequest
	GetRouteTableId() *string
}

type DeleteRouteTableRequest struct {
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
	// The ID of the custom route table.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-bp145q7glnuzdvzu2****
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s DeleteRouteTableRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteTableRequest) GoString() string {
	return s.String()
}

func (s *DeleteRouteTableRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteRouteTableRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteRouteTableRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteRouteTableRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteRouteTableRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteRouteTableRequest) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DeleteRouteTableRequest) SetOwnerAccount(v string) *DeleteRouteTableRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteRouteTableRequest) SetOwnerId(v int64) *DeleteRouteTableRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteRouteTableRequest) SetRegionId(v string) *DeleteRouteTableRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteRouteTableRequest) SetResourceOwnerAccount(v string) *DeleteRouteTableRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteRouteTableRequest) SetResourceOwnerId(v int64) *DeleteRouteTableRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteRouteTableRequest) SetRouteTableId(v string) *DeleteRouteTableRequest {
	s.RouteTableId = &v
	return s
}

func (s *DeleteRouteTableRequest) Validate() error {
	return dara.Validate(s)
}
