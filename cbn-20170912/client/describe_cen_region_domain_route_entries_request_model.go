// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCenRegionDomainRouteEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *DescribeCenRegionDomainRouteEntriesRequest
	GetCenId() *string
	SetCenRegionId(v string) *DescribeCenRegionDomainRouteEntriesRequest
	GetCenRegionId() *string
	SetOwnerAccount(v string) *DescribeCenRegionDomainRouteEntriesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeCenRegionDomainRouteEntriesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeCenRegionDomainRouteEntriesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCenRegionDomainRouteEntriesRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeCenRegionDomainRouteEntriesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeCenRegionDomainRouteEntriesRequest
	GetResourceOwnerId() *int64
	SetStatus(v string) *DescribeCenRegionDomainRouteEntriesRequest
	GetStatus() *string
}

type DescribeCenRegionDomainRouteEntriesRequest struct {
	// The CEN instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-7qthudw0ll6j****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The region ID.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	CenRegionId  *string `json:"CenRegionId,omitempty" xml:"CenRegionId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **1*	- to **500**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The route status. Valid values:
	//
	// 	- **Active*	- (default): available
	//
	// 	- **Candidate**: standby
	//
	// 	- **Rejected**: rejected
	//
	// 	- **Prohibited**: prohibited
	//
	// 	- **All*	- (default value): all routes
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCenRegionDomainRouteEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenRegionDomainRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCenRegionDomainRouteEntriesRequest) GetCenId() *string {
	return s.CenId
}

func (s *DescribeCenRegionDomainRouteEntriesRequest) GetCenRegionId() *string {
	return s.CenRegionId
}

func (s *DescribeCenRegionDomainRouteEntriesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeCenRegionDomainRouteEntriesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCenRegionDomainRouteEntriesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCenRegionDomainRouteEntriesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCenRegionDomainRouteEntriesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeCenRegionDomainRouteEntriesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCenRegionDomainRouteEntriesRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeCenRegionDomainRouteEntriesRequest) SetCenId(v string) *DescribeCenRegionDomainRouteEntriesRequest {
	s.CenId = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesRequest) SetCenRegionId(v string) *DescribeCenRegionDomainRouteEntriesRequest {
	s.CenRegionId = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesRequest) SetOwnerAccount(v string) *DescribeCenRegionDomainRouteEntriesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesRequest) SetOwnerId(v int64) *DescribeCenRegionDomainRouteEntriesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesRequest) SetPageNumber(v int32) *DescribeCenRegionDomainRouteEntriesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesRequest) SetPageSize(v int32) *DescribeCenRegionDomainRouteEntriesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesRequest) SetResourceOwnerAccount(v string) *DescribeCenRegionDomainRouteEntriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesRequest) SetResourceOwnerId(v int64) *DescribeCenRegionDomainRouteEntriesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesRequest) SetStatus(v string) *DescribeCenRegionDomainRouteEntriesRequest {
	s.Status = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesRequest) Validate() error {
	return dara.Validate(s)
}
