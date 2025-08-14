// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouteServiceInCenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessRegionId(v string) *DeleteRouteServiceInCenRequest
	GetAccessRegionId() *string
	SetCenId(v string) *DeleteRouteServiceInCenRequest
	GetCenId() *string
	SetHost(v string) *DeleteRouteServiceInCenRequest
	GetHost() *string
	SetHostRegionId(v string) *DeleteRouteServiceInCenRequest
	GetHostRegionId() *string
	SetHostVpcId(v string) *DeleteRouteServiceInCenRequest
	GetHostVpcId() *string
	SetOwnerAccount(v string) *DeleteRouteServiceInCenRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteRouteServiceInCenRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteRouteServiceInCenRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteRouteServiceInCenRequest
	GetResourceOwnerId() *int64
}

type DeleteRouteServiceInCenRequest struct {
	// The ID of the region where the cloud service is accessed.
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
	// The IP addresses or CIDR blocks of the cloud service.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100.118.28.0/24
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The region ID of the cloud service.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	HostRegionId *string `json:"HostRegionId,omitempty" xml:"HostRegionId,omitempty"`
	// The ID of the virtual private cloud (VPC) that is associated with the cloud service.
	//
	// example:
	//
	// vpc-bp1t36rn9l53iwbsf****
	HostVpcId            *string `json:"HostVpcId,omitempty" xml:"HostVpcId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteRouteServiceInCenRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteServiceInCenRequest) GoString() string {
	return s.String()
}

func (s *DeleteRouteServiceInCenRequest) GetAccessRegionId() *string {
	return s.AccessRegionId
}

func (s *DeleteRouteServiceInCenRequest) GetCenId() *string {
	return s.CenId
}

func (s *DeleteRouteServiceInCenRequest) GetHost() *string {
	return s.Host
}

func (s *DeleteRouteServiceInCenRequest) GetHostRegionId() *string {
	return s.HostRegionId
}

func (s *DeleteRouteServiceInCenRequest) GetHostVpcId() *string {
	return s.HostVpcId
}

func (s *DeleteRouteServiceInCenRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteRouteServiceInCenRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteRouteServiceInCenRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteRouteServiceInCenRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteRouteServiceInCenRequest) SetAccessRegionId(v string) *DeleteRouteServiceInCenRequest {
	s.AccessRegionId = &v
	return s
}

func (s *DeleteRouteServiceInCenRequest) SetCenId(v string) *DeleteRouteServiceInCenRequest {
	s.CenId = &v
	return s
}

func (s *DeleteRouteServiceInCenRequest) SetHost(v string) *DeleteRouteServiceInCenRequest {
	s.Host = &v
	return s
}

func (s *DeleteRouteServiceInCenRequest) SetHostRegionId(v string) *DeleteRouteServiceInCenRequest {
	s.HostRegionId = &v
	return s
}

func (s *DeleteRouteServiceInCenRequest) SetHostVpcId(v string) *DeleteRouteServiceInCenRequest {
	s.HostVpcId = &v
	return s
}

func (s *DeleteRouteServiceInCenRequest) SetOwnerAccount(v string) *DeleteRouteServiceInCenRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteRouteServiceInCenRequest) SetOwnerId(v int64) *DeleteRouteServiceInCenRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteRouteServiceInCenRequest) SetResourceOwnerAccount(v string) *DeleteRouteServiceInCenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteRouteServiceInCenRequest) SetResourceOwnerId(v int64) *DeleteRouteServiceInCenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteRouteServiceInCenRequest) Validate() error {
	return dara.Validate(s)
}
