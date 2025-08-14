// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCenInterRegionBandwidthLimitsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *DescribeCenInterRegionBandwidthLimitsRequest
	GetCenId() *string
	SetOwnerAccount(v string) *DescribeCenInterRegionBandwidthLimitsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeCenInterRegionBandwidthLimitsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeCenInterRegionBandwidthLimitsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCenInterRegionBandwidthLimitsRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeCenInterRegionBandwidthLimitsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeCenInterRegionBandwidthLimitsRequest
	GetResourceOwnerId() *int64
	SetTrRegionId(v string) *DescribeCenInterRegionBandwidthLimitsRequest
	GetTrRegionId() *string
}

type DescribeCenInterRegionBandwidthLimitsRequest struct {
	// The ID of the Cloud Enterprise Network (CEN) instance.
	//
	// example:
	//
	// cen-pfa6ugf3xl0qsd****
	CenId        *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// The region ID of the transit router.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query the most recent region list.
	//
	// example:
	//
	// ccn-cn-shanghai
	TrRegionId *string `json:"TrRegionId,omitempty" xml:"TrRegionId,omitempty"`
}

func (s DescribeCenInterRegionBandwidthLimitsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenInterRegionBandwidthLimitsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCenInterRegionBandwidthLimitsRequest) GetCenId() *string {
	return s.CenId
}

func (s *DescribeCenInterRegionBandwidthLimitsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeCenInterRegionBandwidthLimitsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCenInterRegionBandwidthLimitsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCenInterRegionBandwidthLimitsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCenInterRegionBandwidthLimitsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeCenInterRegionBandwidthLimitsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCenInterRegionBandwidthLimitsRequest) GetTrRegionId() *string {
	return s.TrRegionId
}

func (s *DescribeCenInterRegionBandwidthLimitsRequest) SetCenId(v string) *DescribeCenInterRegionBandwidthLimitsRequest {
	s.CenId = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsRequest) SetOwnerAccount(v string) *DescribeCenInterRegionBandwidthLimitsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsRequest) SetOwnerId(v int64) *DescribeCenInterRegionBandwidthLimitsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsRequest) SetPageNumber(v int32) *DescribeCenInterRegionBandwidthLimitsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsRequest) SetPageSize(v int32) *DescribeCenInterRegionBandwidthLimitsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsRequest) SetResourceOwnerAccount(v string) *DescribeCenInterRegionBandwidthLimitsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsRequest) SetResourceOwnerId(v int64) *DescribeCenInterRegionBandwidthLimitsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsRequest) SetTrRegionId(v string) *DescribeCenInterRegionBandwidthLimitsRequest {
	s.TrRegionId = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsRequest) Validate() error {
	return dara.Validate(s)
}
