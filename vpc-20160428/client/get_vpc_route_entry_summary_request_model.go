// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVpcRouteEntrySummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *GetVpcRouteEntrySummaryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetVpcRouteEntrySummaryRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetVpcRouteEntrySummaryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GetVpcRouteEntrySummaryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetVpcRouteEntrySummaryRequest
	GetResourceOwnerId() *int64
	SetRouteEntryType(v string) *GetVpcRouteEntrySummaryRequest
	GetRouteEntryType() *string
	SetRouteTableId(v string) *GetVpcRouteEntrySummaryRequest
	GetRouteTableId() *string
	SetVpcId(v string) *GetVpcRouteEntrySummaryRequest
	GetVpcId() *string
}

type GetVpcRouteEntrySummaryRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the route table.
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
	// The type of the route. Valid values:
	//
	// 	- **All**: all route types
	//
	// 	- **Custom**: a custom route
	//
	// 	- **System**: a system route
	//
	// 	- **BGP**: a BGP route
	//
	// 	- **CEN**: a Cloud Enterprise Network (CEN) route
	//
	// This parameter is required.
	//
	// example:
	//
	// Custom
	RouteEntryType *string `json:"RouteEntryType,omitempty" xml:"RouteEntryType,omitempty"`
	// The ID of the route table that you want to query.
	//
	// example:
	//
	// vtb-bp145q7glnuzdvzu2****
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	// The ID of the virtual private cloud (VPC) to which the route table belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp15zckdt37pq72****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetVpcRouteEntrySummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVpcRouteEntrySummaryRequest) GoString() string {
	return s.String()
}

func (s *GetVpcRouteEntrySummaryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetVpcRouteEntrySummaryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetVpcRouteEntrySummaryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetVpcRouteEntrySummaryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetVpcRouteEntrySummaryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetVpcRouteEntrySummaryRequest) GetRouteEntryType() *string {
	return s.RouteEntryType
}

func (s *GetVpcRouteEntrySummaryRequest) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *GetVpcRouteEntrySummaryRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *GetVpcRouteEntrySummaryRequest) SetOwnerAccount(v string) *GetVpcRouteEntrySummaryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetVpcRouteEntrySummaryRequest) SetOwnerId(v int64) *GetVpcRouteEntrySummaryRequest {
	s.OwnerId = &v
	return s
}

func (s *GetVpcRouteEntrySummaryRequest) SetRegionId(v string) *GetVpcRouteEntrySummaryRequest {
	s.RegionId = &v
	return s
}

func (s *GetVpcRouteEntrySummaryRequest) SetResourceOwnerAccount(v string) *GetVpcRouteEntrySummaryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetVpcRouteEntrySummaryRequest) SetResourceOwnerId(v int64) *GetVpcRouteEntrySummaryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetVpcRouteEntrySummaryRequest) SetRouteEntryType(v string) *GetVpcRouteEntrySummaryRequest {
	s.RouteEntryType = &v
	return s
}

func (s *GetVpcRouteEntrySummaryRequest) SetRouteTableId(v string) *GetVpcRouteEntrySummaryRequest {
	s.RouteTableId = &v
	return s
}

func (s *GetVpcRouteEntrySummaryRequest) SetVpcId(v string) *GetVpcRouteEntrySummaryRequest {
	s.VpcId = &v
	return s
}

func (s *GetVpcRouteEntrySummaryRequest) Validate() error {
	return dara.Validate(s)
}
