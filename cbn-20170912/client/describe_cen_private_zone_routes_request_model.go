// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCenPrivateZoneRoutesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessRegionId(v string) *DescribeCenPrivateZoneRoutesRequest
	GetAccessRegionId() *string
	SetCenId(v string) *DescribeCenPrivateZoneRoutesRequest
	GetCenId() *string
	SetHostRegionId(v string) *DescribeCenPrivateZoneRoutesRequest
	GetHostRegionId() *string
	SetPageNumber(v int32) *DescribeCenPrivateZoneRoutesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCenPrivateZoneRoutesRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeCenPrivateZoneRoutesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeCenPrivateZoneRoutesRequest
	GetResourceOwnerId() *int64
}

type DescribeCenPrivateZoneRoutesRequest struct {
	// The ID of the region where PrivateZone is accessed.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query the most recent region list.
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
	// cen-7qthudw0ll6jmc****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The ID of the region where PrivateZone is deployed.
	//
	// example:
	//
	// cn-hangzhou
	HostRegionId *string `json:"HostRegionId,omitempty" xml:"HostRegionId,omitempty"`
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

func (s DescribeCenPrivateZoneRoutesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenPrivateZoneRoutesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCenPrivateZoneRoutesRequest) GetAccessRegionId() *string {
	return s.AccessRegionId
}

func (s *DescribeCenPrivateZoneRoutesRequest) GetCenId() *string {
	return s.CenId
}

func (s *DescribeCenPrivateZoneRoutesRequest) GetHostRegionId() *string {
	return s.HostRegionId
}

func (s *DescribeCenPrivateZoneRoutesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCenPrivateZoneRoutesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCenPrivateZoneRoutesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeCenPrivateZoneRoutesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCenPrivateZoneRoutesRequest) SetAccessRegionId(v string) *DescribeCenPrivateZoneRoutesRequest {
	s.AccessRegionId = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesRequest) SetCenId(v string) *DescribeCenPrivateZoneRoutesRequest {
	s.CenId = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesRequest) SetHostRegionId(v string) *DescribeCenPrivateZoneRoutesRequest {
	s.HostRegionId = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesRequest) SetPageNumber(v int32) *DescribeCenPrivateZoneRoutesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesRequest) SetPageSize(v int32) *DescribeCenPrivateZoneRoutesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesRequest) SetResourceOwnerAccount(v string) *DescribeCenPrivateZoneRoutesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesRequest) SetResourceOwnerId(v int64) *DescribeCenPrivateZoneRoutesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesRequest) Validate() error {
	return dara.Validate(s)
}
