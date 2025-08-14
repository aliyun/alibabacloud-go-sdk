// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCenGeographicSpanRemainingBandwidthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *DescribeCenGeographicSpanRemainingBandwidthRequest
	GetCenId() *string
	SetGeographicRegionAId(v string) *DescribeCenGeographicSpanRemainingBandwidthRequest
	GetGeographicRegionAId() *string
	SetGeographicRegionBId(v string) *DescribeCenGeographicSpanRemainingBandwidthRequest
	GetGeographicRegionBId() *string
	SetOwnerAccount(v string) *DescribeCenGeographicSpanRemainingBandwidthRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeCenGeographicSpanRemainingBandwidthRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeCenGeographicSpanRemainingBandwidthRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCenGeographicSpanRemainingBandwidthRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeCenGeographicSpanRemainingBandwidthRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeCenGeographicSpanRemainingBandwidthRequest
	GetResourceOwnerId() *int64
}

type DescribeCenGeographicSpanRemainingBandwidthRequest struct {
	// The ID of the Cloud Enterprise Network (CEN) instance to which the bandwidth plan is associated.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-nh98vzx8gfhlwn****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The ID of one of the connected areas of the bandwidth plan. Valid values:
	//
	// 	- **China**: Chinese mainland
	//
	// 	- **North-America**: North America
	//
	// 	- **Asia-Pacific**: Asia Pacific
	//
	// 	- **Europe**: Europe
	//
	// This parameter is required.
	//
	// example:
	//
	// China
	GeographicRegionAId *string `json:"GeographicRegionAId,omitempty" xml:"GeographicRegionAId,omitempty"`
	// The ID of the other area connected by the bandwidth plan. Valid values:
	//
	// 	- **China**: Chinese mainland
	//
	// 	- **North-America**: North America
	//
	// 	- **Asia-Pacific**: Asia Pacific
	//
	// 	- **Europe**: Europe
	//
	// This parameter is required.
	//
	// example:
	//
	// North-America
	GeographicRegionBId *string `json:"GeographicRegionBId,omitempty" xml:"GeographicRegionBId,omitempty"`
	OwnerAccount        *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s DescribeCenGeographicSpanRemainingBandwidthRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenGeographicSpanRemainingBandwidthRequest) GoString() string {
	return s.String()
}

func (s *DescribeCenGeographicSpanRemainingBandwidthRequest) GetCenId() *string {
	return s.CenId
}

func (s *DescribeCenGeographicSpanRemainingBandwidthRequest) GetGeographicRegionAId() *string {
	return s.GeographicRegionAId
}

func (s *DescribeCenGeographicSpanRemainingBandwidthRequest) GetGeographicRegionBId() *string {
	return s.GeographicRegionBId
}

func (s *DescribeCenGeographicSpanRemainingBandwidthRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeCenGeographicSpanRemainingBandwidthRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCenGeographicSpanRemainingBandwidthRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCenGeographicSpanRemainingBandwidthRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCenGeographicSpanRemainingBandwidthRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeCenGeographicSpanRemainingBandwidthRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCenGeographicSpanRemainingBandwidthRequest) SetCenId(v string) *DescribeCenGeographicSpanRemainingBandwidthRequest {
	s.CenId = &v
	return s
}

func (s *DescribeCenGeographicSpanRemainingBandwidthRequest) SetGeographicRegionAId(v string) *DescribeCenGeographicSpanRemainingBandwidthRequest {
	s.GeographicRegionAId = &v
	return s
}

func (s *DescribeCenGeographicSpanRemainingBandwidthRequest) SetGeographicRegionBId(v string) *DescribeCenGeographicSpanRemainingBandwidthRequest {
	s.GeographicRegionBId = &v
	return s
}

func (s *DescribeCenGeographicSpanRemainingBandwidthRequest) SetOwnerAccount(v string) *DescribeCenGeographicSpanRemainingBandwidthRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCenGeographicSpanRemainingBandwidthRequest) SetOwnerId(v int64) *DescribeCenGeographicSpanRemainingBandwidthRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCenGeographicSpanRemainingBandwidthRequest) SetPageNumber(v int32) *DescribeCenGeographicSpanRemainingBandwidthRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenGeographicSpanRemainingBandwidthRequest) SetPageSize(v int32) *DescribeCenGeographicSpanRemainingBandwidthRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCenGeographicSpanRemainingBandwidthRequest) SetResourceOwnerAccount(v string) *DescribeCenGeographicSpanRemainingBandwidthRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCenGeographicSpanRemainingBandwidthRequest) SetResourceOwnerId(v int64) *DescribeCenGeographicSpanRemainingBandwidthRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCenGeographicSpanRemainingBandwidthRequest) Validate() error {
	return dara.Validate(s)
}
