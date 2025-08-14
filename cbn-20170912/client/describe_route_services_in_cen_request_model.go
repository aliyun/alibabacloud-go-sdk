// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRouteServicesInCenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessRegionId(v string) *DescribeRouteServicesInCenRequest
	GetAccessRegionId() *string
	SetCenId(v string) *DescribeRouteServicesInCenRequest
	GetCenId() *string
	SetHost(v string) *DescribeRouteServicesInCenRequest
	GetHost() *string
	SetHostRegionId(v string) *DescribeRouteServicesInCenRequest
	GetHostRegionId() *string
	SetHostVpcId(v string) *DescribeRouteServicesInCenRequest
	GetHostVpcId() *string
	SetOwnerAccount(v string) *DescribeRouteServicesInCenRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeRouteServicesInCenRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeRouteServicesInCenRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRouteServicesInCenRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeRouteServicesInCenRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeRouteServicesInCenRequest
	GetResourceOwnerId() *int64
}

type DescribeRouteServicesInCenRequest struct {
	// The ID of the region where the cloud service is accessed.
	//
	// example:
	//
	// cn-hangzhou
	AccessRegionId *string `json:"AccessRegionId,omitempty" xml:"AccessRegionId,omitempty"`
	// The ID of the Cloud Enterprise Network (CEN) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-pfa6ugf3xl0qsd****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The service address of the cloud service.
	//
	// You can enter a domain name, an IP address, or a CIDR block.
	//
	// example:
	//
	// 100.118.28.0/24
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The region ID of the cloud service.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	HostRegionId *string `json:"HostRegionId,omitempty" xml:"HostRegionId,omitempty"`
	// The ID of the VPC associated with the cloud service.
	//
	// example:
	//
	// vpc-bp1h8vbrbcgohcju5****
	HostVpcId    *string `json:"HostVpcId,omitempty" xml:"HostVpcId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **10**. Valid values: **1*	- to **50**.
	//
	// example:
	//
	// 10
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeRouteServicesInCenRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteServicesInCenRequest) GoString() string {
	return s.String()
}

func (s *DescribeRouteServicesInCenRequest) GetAccessRegionId() *string {
	return s.AccessRegionId
}

func (s *DescribeRouteServicesInCenRequest) GetCenId() *string {
	return s.CenId
}

func (s *DescribeRouteServicesInCenRequest) GetHost() *string {
	return s.Host
}

func (s *DescribeRouteServicesInCenRequest) GetHostRegionId() *string {
	return s.HostRegionId
}

func (s *DescribeRouteServicesInCenRequest) GetHostVpcId() *string {
	return s.HostVpcId
}

func (s *DescribeRouteServicesInCenRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeRouteServicesInCenRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRouteServicesInCenRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRouteServicesInCenRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRouteServicesInCenRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeRouteServicesInCenRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeRouteServicesInCenRequest) SetAccessRegionId(v string) *DescribeRouteServicesInCenRequest {
	s.AccessRegionId = &v
	return s
}

func (s *DescribeRouteServicesInCenRequest) SetCenId(v string) *DescribeRouteServicesInCenRequest {
	s.CenId = &v
	return s
}

func (s *DescribeRouteServicesInCenRequest) SetHost(v string) *DescribeRouteServicesInCenRequest {
	s.Host = &v
	return s
}

func (s *DescribeRouteServicesInCenRequest) SetHostRegionId(v string) *DescribeRouteServicesInCenRequest {
	s.HostRegionId = &v
	return s
}

func (s *DescribeRouteServicesInCenRequest) SetHostVpcId(v string) *DescribeRouteServicesInCenRequest {
	s.HostVpcId = &v
	return s
}

func (s *DescribeRouteServicesInCenRequest) SetOwnerAccount(v string) *DescribeRouteServicesInCenRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRouteServicesInCenRequest) SetOwnerId(v int64) *DescribeRouteServicesInCenRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRouteServicesInCenRequest) SetPageNumber(v int32) *DescribeRouteServicesInCenRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRouteServicesInCenRequest) SetPageSize(v int32) *DescribeRouteServicesInCenRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRouteServicesInCenRequest) SetResourceOwnerAccount(v string) *DescribeRouteServicesInCenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRouteServicesInCenRequest) SetResourceOwnerId(v int64) *DescribeRouteServicesInCenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRouteServicesInCenRequest) Validate() error {
	return dara.Validate(s)
}
