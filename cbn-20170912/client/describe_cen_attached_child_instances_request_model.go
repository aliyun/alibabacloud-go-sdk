// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCenAttachedChildInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *DescribeCenAttachedChildInstancesRequest
	GetCenId() *string
	SetChildInstanceRegionId(v string) *DescribeCenAttachedChildInstancesRequest
	GetChildInstanceRegionId() *string
	SetChildInstanceType(v string) *DescribeCenAttachedChildInstancesRequest
	GetChildInstanceType() *string
	SetOwnerAccount(v string) *DescribeCenAttachedChildInstancesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeCenAttachedChildInstancesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeCenAttachedChildInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCenAttachedChildInstancesRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeCenAttachedChildInstancesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeCenAttachedChildInstancesRequest
	GetResourceOwnerId() *int64
}

type DescribeCenAttachedChildInstancesRequest struct {
	// The ID of the CEN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-j3jzhw1zpau2km****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The ID of the region where the network instance is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-zhangjiakou
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	// The type of the network instance. Valid values:
	//
	// 	- **VPC**: virtual private cloud (VPC)
	//
	// 	- **VBR**: virtual border router (VBR)
	//
	// 	- **CCN**: Cloud Connect Network (CCN) instance
	//
	// example:
	//
	// VPC
	ChildInstanceType *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	OwnerAccount      *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **10**. Valid values: **1*	- to **50**.
	//
	// example:
	//
	// 10
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeCenAttachedChildInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenAttachedChildInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCenAttachedChildInstancesRequest) GetCenId() *string {
	return s.CenId
}

func (s *DescribeCenAttachedChildInstancesRequest) GetChildInstanceRegionId() *string {
	return s.ChildInstanceRegionId
}

func (s *DescribeCenAttachedChildInstancesRequest) GetChildInstanceType() *string {
	return s.ChildInstanceType
}

func (s *DescribeCenAttachedChildInstancesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeCenAttachedChildInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCenAttachedChildInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCenAttachedChildInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCenAttachedChildInstancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeCenAttachedChildInstancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCenAttachedChildInstancesRequest) SetCenId(v string) *DescribeCenAttachedChildInstancesRequest {
	s.CenId = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesRequest) SetChildInstanceRegionId(v string) *DescribeCenAttachedChildInstancesRequest {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesRequest) SetChildInstanceType(v string) *DescribeCenAttachedChildInstancesRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesRequest) SetOwnerAccount(v string) *DescribeCenAttachedChildInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesRequest) SetOwnerId(v int64) *DescribeCenAttachedChildInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesRequest) SetPageNumber(v int32) *DescribeCenAttachedChildInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesRequest) SetPageSize(v int32) *DescribeCenAttachedChildInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesRequest) SetResourceOwnerAccount(v string) *DescribeCenAttachedChildInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesRequest) SetResourceOwnerId(v int64) *DescribeCenAttachedChildInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesRequest) Validate() error {
	return dara.Validate(s)
}
