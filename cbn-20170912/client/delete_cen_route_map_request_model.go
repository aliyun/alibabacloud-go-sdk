// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCenRouteMapRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *DeleteCenRouteMapRequest
	GetCenId() *string
	SetCenRegionId(v string) *DeleteCenRouteMapRequest
	GetCenRegionId() *string
	SetOwnerAccount(v string) *DeleteCenRouteMapRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteCenRouteMapRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteCenRouteMapRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteCenRouteMapRequest
	GetResourceOwnerId() *int64
	SetRouteMapId(v string) *DeleteCenRouteMapRequest
	GetRouteMapId() *string
}

type DeleteCenRouteMapRequest struct {
	// The ID of the Cloud Enterprise Network (CEN) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-7qthudw0ll6jm****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The ID of the region in which the routing policy is applied.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	CenRegionId          *string `json:"CenRegionId,omitempty" xml:"CenRegionId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the routing policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// cenrmap-abcdedfghij****
	RouteMapId *string `json:"RouteMapId,omitempty" xml:"RouteMapId,omitempty"`
}

func (s DeleteCenRouteMapRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCenRouteMapRequest) GoString() string {
	return s.String()
}

func (s *DeleteCenRouteMapRequest) GetCenId() *string {
	return s.CenId
}

func (s *DeleteCenRouteMapRequest) GetCenRegionId() *string {
	return s.CenRegionId
}

func (s *DeleteCenRouteMapRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteCenRouteMapRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteCenRouteMapRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteCenRouteMapRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteCenRouteMapRequest) GetRouteMapId() *string {
	return s.RouteMapId
}

func (s *DeleteCenRouteMapRequest) SetCenId(v string) *DeleteCenRouteMapRequest {
	s.CenId = &v
	return s
}

func (s *DeleteCenRouteMapRequest) SetCenRegionId(v string) *DeleteCenRouteMapRequest {
	s.CenRegionId = &v
	return s
}

func (s *DeleteCenRouteMapRequest) SetOwnerAccount(v string) *DeleteCenRouteMapRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteCenRouteMapRequest) SetOwnerId(v int64) *DeleteCenRouteMapRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCenRouteMapRequest) SetResourceOwnerAccount(v string) *DeleteCenRouteMapRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteCenRouteMapRequest) SetResourceOwnerId(v int64) *DeleteCenRouteMapRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteCenRouteMapRequest) SetRouteMapId(v string) *DeleteCenRouteMapRequest {
	s.RouteMapId = &v
	return s
}

func (s *DeleteCenRouteMapRequest) Validate() error {
	return dara.Validate(s)
}
