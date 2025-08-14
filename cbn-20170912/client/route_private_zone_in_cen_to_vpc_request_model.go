// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRoutePrivateZoneInCenToVpcRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessRegionId(v string) *RoutePrivateZoneInCenToVpcRequest
	GetAccessRegionId() *string
	SetCenId(v string) *RoutePrivateZoneInCenToVpcRequest
	GetCenId() *string
	SetHostRegionId(v string) *RoutePrivateZoneInCenToVpcRequest
	GetHostRegionId() *string
	SetHostVpcId(v string) *RoutePrivateZoneInCenToVpcRequest
	GetHostVpcId() *string
	SetOwnerAccount(v string) *RoutePrivateZoneInCenToVpcRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RoutePrivateZoneInCenToVpcRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *RoutePrivateZoneInCenToVpcRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RoutePrivateZoneInCenToVpcRequest
	GetResourceOwnerId() *int64
}

type RoutePrivateZoneInCenToVpcRequest struct {
	// The ID of the region where PrivateZone is accessed.
	//
	// This region refers to the region in which PrivateZone is accessed by clients.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	AccessRegionId *string `json:"AccessRegionId,omitempty" xml:"AccessRegionId,omitempty"`
	// The ID of the CEN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-7qthudw0ll6jmc****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The ID of the region where PrivateZone is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	HostRegionId *string `json:"HostRegionId,omitempty" xml:"HostRegionId,omitempty"`
	// The ID of the VPC that is associated with PrivateZone.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1h8vbrbcgohcju5****
	HostVpcId            *string `json:"HostVpcId,omitempty" xml:"HostVpcId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RoutePrivateZoneInCenToVpcRequest) String() string {
	return dara.Prettify(s)
}

func (s RoutePrivateZoneInCenToVpcRequest) GoString() string {
	return s.String()
}

func (s *RoutePrivateZoneInCenToVpcRequest) GetAccessRegionId() *string {
	return s.AccessRegionId
}

func (s *RoutePrivateZoneInCenToVpcRequest) GetCenId() *string {
	return s.CenId
}

func (s *RoutePrivateZoneInCenToVpcRequest) GetHostRegionId() *string {
	return s.HostRegionId
}

func (s *RoutePrivateZoneInCenToVpcRequest) GetHostVpcId() *string {
	return s.HostVpcId
}

func (s *RoutePrivateZoneInCenToVpcRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RoutePrivateZoneInCenToVpcRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RoutePrivateZoneInCenToVpcRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RoutePrivateZoneInCenToVpcRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RoutePrivateZoneInCenToVpcRequest) SetAccessRegionId(v string) *RoutePrivateZoneInCenToVpcRequest {
	s.AccessRegionId = &v
	return s
}

func (s *RoutePrivateZoneInCenToVpcRequest) SetCenId(v string) *RoutePrivateZoneInCenToVpcRequest {
	s.CenId = &v
	return s
}

func (s *RoutePrivateZoneInCenToVpcRequest) SetHostRegionId(v string) *RoutePrivateZoneInCenToVpcRequest {
	s.HostRegionId = &v
	return s
}

func (s *RoutePrivateZoneInCenToVpcRequest) SetHostVpcId(v string) *RoutePrivateZoneInCenToVpcRequest {
	s.HostVpcId = &v
	return s
}

func (s *RoutePrivateZoneInCenToVpcRequest) SetOwnerAccount(v string) *RoutePrivateZoneInCenToVpcRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RoutePrivateZoneInCenToVpcRequest) SetOwnerId(v int64) *RoutePrivateZoneInCenToVpcRequest {
	s.OwnerId = &v
	return s
}

func (s *RoutePrivateZoneInCenToVpcRequest) SetResourceOwnerAccount(v string) *RoutePrivateZoneInCenToVpcRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RoutePrivateZoneInCenToVpcRequest) SetResourceOwnerId(v int64) *RoutePrivateZoneInCenToVpcRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RoutePrivateZoneInCenToVpcRequest) Validate() error {
	return dara.Validate(s)
}
