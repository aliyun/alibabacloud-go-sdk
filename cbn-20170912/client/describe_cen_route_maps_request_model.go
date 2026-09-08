// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCenRouteMapsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *DescribeCenRouteMapsRequest
	GetCenId() *string
	SetCenRegionId(v string) *DescribeCenRouteMapsRequest
	GetCenRegionId() *string
	SetOwnerAccount(v string) *DescribeCenRouteMapsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeCenRouteMapsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeCenRouteMapsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCenRouteMapsRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeCenRouteMapsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeCenRouteMapsRequest
	GetResourceOwnerId() *int64
	SetRouteMapId(v string) *DescribeCenRouteMapsRequest
	GetRouteMapId() *string
	SetTransitRouterRouteTableId(v string) *DescribeCenRouteMapsRequest
	GetTransitRouterRouteTableId() *string
	SetTransmitDirection(v string) *DescribeCenRouteMapsRequest
	GetTransmitDirection() *string
}

type DescribeCenRouteMapsRequest struct {
	// The instance ID of the Cloud Enterprise Network (CEN) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-wx12mmlt17ld82****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The ID of the region where the route map is applied.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query region IDs.
	//
	// example:
	//
	// cn-hangzhou
	CenRegionId  *string `json:"CenRegionId,omitempty" xml:"CenRegionId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number of the list. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in a paged query. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the route map.
	//
	// example:
	//
	// cenrmap-y40mxdvf7joc12****
	RouteMapId *string `json:"RouteMapId,omitempty" xml:"RouteMapId,omitempty"`
	// The ID of the transit router route table associated with the route map.
	//
	// example:
	//
	// vtb-gw8nx3515m1mbd1z1****
	TransitRouterRouteTableId *string `json:"TransitRouterRouteTableId,omitempty" xml:"TransitRouterRouteTableId,omitempty"`
	// The direction in which the route map is applied. Valid values:
	//
	// - **RegionIn**: Routes are advertised to the regional gateway of the CEN instance.
	//
	//  For example, routes are advertised from a network instance in the local region to the local regional gateway, or routes are advertised from a regional gateway in another region to the local regional gateway.
	//
	// - **RegionOut**: Routes are advertised from the regional gateway of the CEN instance.
	//
	//  For example, routes are advertised from the local regional gateway to network instances in the local region, or routes are advertised from the local regional gateway to regional gateways in other regions.
	//
	// example:
	//
	// RegionOut
	TransmitDirection *string `json:"TransmitDirection,omitempty" xml:"TransmitDirection,omitempty"`
}

func (s DescribeCenRouteMapsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenRouteMapsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsRequest) GetCenId() *string {
	return s.CenId
}

func (s *DescribeCenRouteMapsRequest) GetCenRegionId() *string {
	return s.CenRegionId
}

func (s *DescribeCenRouteMapsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeCenRouteMapsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCenRouteMapsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCenRouteMapsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCenRouteMapsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeCenRouteMapsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCenRouteMapsRequest) GetRouteMapId() *string {
	return s.RouteMapId
}

func (s *DescribeCenRouteMapsRequest) GetTransitRouterRouteTableId() *string {
	return s.TransitRouterRouteTableId
}

func (s *DescribeCenRouteMapsRequest) GetTransmitDirection() *string {
	return s.TransmitDirection
}

func (s *DescribeCenRouteMapsRequest) SetCenId(v string) *DescribeCenRouteMapsRequest {
	s.CenId = &v
	return s
}

func (s *DescribeCenRouteMapsRequest) SetCenRegionId(v string) *DescribeCenRouteMapsRequest {
	s.CenRegionId = &v
	return s
}

func (s *DescribeCenRouteMapsRequest) SetOwnerAccount(v string) *DescribeCenRouteMapsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCenRouteMapsRequest) SetOwnerId(v int64) *DescribeCenRouteMapsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCenRouteMapsRequest) SetPageNumber(v int32) *DescribeCenRouteMapsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenRouteMapsRequest) SetPageSize(v int32) *DescribeCenRouteMapsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCenRouteMapsRequest) SetResourceOwnerAccount(v string) *DescribeCenRouteMapsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCenRouteMapsRequest) SetResourceOwnerId(v int64) *DescribeCenRouteMapsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCenRouteMapsRequest) SetRouteMapId(v string) *DescribeCenRouteMapsRequest {
	s.RouteMapId = &v
	return s
}

func (s *DescribeCenRouteMapsRequest) SetTransitRouterRouteTableId(v string) *DescribeCenRouteMapsRequest {
	s.TransitRouterRouteTableId = &v
	return s
}

func (s *DescribeCenRouteMapsRequest) SetTransmitDirection(v string) *DescribeCenRouteMapsRequest {
	s.TransmitDirection = &v
	return s
}

func (s *DescribeCenRouteMapsRequest) Validate() error {
	return dara.Validate(s)
}
