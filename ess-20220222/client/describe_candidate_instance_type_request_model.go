// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCandidateInstanceTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowCrossAz(v bool) *DescribeCandidateInstanceTypeRequest
	GetAllowCrossAz() *bool
	SetAllowDifferentGeneration(v bool) *DescribeCandidateInstanceTypeRequest
	GetAllowDifferentGeneration() *bool
	SetDataDiskCategories(v []*string) *DescribeCandidateInstanceTypeRequest
	GetDataDiskCategories() []*string
	SetImageFamily(v string) *DescribeCandidateInstanceTypeRequest
	GetImageFamily() *string
	SetImageId(v string) *DescribeCandidateInstanceTypeRequest
	GetImageId() *string
	SetImageName(v string) *DescribeCandidateInstanceTypeRequest
	GetImageName() *string
	SetInstanceTypes(v []*string) *DescribeCandidateInstanceTypeRequest
	GetInstanceTypes() []*string
	SetIpv6AddressCount(v int32) *DescribeCandidateInstanceTypeRequest
	GetIpv6AddressCount() *int32
	SetMaxPrice(v float32) *DescribeCandidateInstanceTypeRequest
	GetMaxPrice() *float32
	SetOwnerId(v int64) *DescribeCandidateInstanceTypeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeCandidateInstanceTypeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeCandidateInstanceTypeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeCandidateInstanceTypeRequest
	GetResourceOwnerId() *int64
	SetSpotStrategy(v string) *DescribeCandidateInstanceTypeRequest
	GetSpotStrategy() *string
	SetSystemDiskCategories(v []*string) *DescribeCandidateInstanceTypeRequest
	GetSystemDiskCategories() []*string
	SetZoneIds(v []*string) *DescribeCandidateInstanceTypeRequest
	GetZoneIds() []*string
}

type DescribeCandidateInstanceTypeRequest struct {
	// example:
	//
	// true
	AllowCrossAz *bool `json:"AllowCrossAz,omitempty" xml:"AllowCrossAz,omitempty"`
	// example:
	//
	// true
	AllowDifferentGeneration *bool     `json:"AllowDifferentGeneration,omitempty" xml:"AllowDifferentGeneration,omitempty"`
	DataDiskCategories       []*string `json:"DataDiskCategories,omitempty" xml:"DataDiskCategories,omitempty" type:"Repeated"`
	// example:
	//
	// hangzhou-daily-update
	ImageFamily *string `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	// example:
	//
	// centos6u5_64_20G_aliaegis****.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// centos6u5_64_20G_aliaegis_20140703.vhd
	ImageName     *string   `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	InstanceTypes []*string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Ipv6AddressCount *int32 `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	// example:
	//
	// 2
	MaxPrice *float32 `json:"MaxPrice,omitempty" xml:"MaxPrice,omitempty"`
	OwnerId  *int64   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// cn-qingdao
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// NoSpot
	SpotStrategy         *string   `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	SystemDiskCategories []*string `json:"SystemDiskCategories,omitempty" xml:"SystemDiskCategories,omitempty" type:"Repeated"`
	ZoneIds              []*string `json:"ZoneIds,omitempty" xml:"ZoneIds,omitempty" type:"Repeated"`
}

func (s DescribeCandidateInstanceTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCandidateInstanceTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeCandidateInstanceTypeRequest) GetAllowCrossAz() *bool {
	return s.AllowCrossAz
}

func (s *DescribeCandidateInstanceTypeRequest) GetAllowDifferentGeneration() *bool {
	return s.AllowDifferentGeneration
}

func (s *DescribeCandidateInstanceTypeRequest) GetDataDiskCategories() []*string {
	return s.DataDiskCategories
}

func (s *DescribeCandidateInstanceTypeRequest) GetImageFamily() *string {
	return s.ImageFamily
}

func (s *DescribeCandidateInstanceTypeRequest) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeCandidateInstanceTypeRequest) GetImageName() *string {
	return s.ImageName
}

func (s *DescribeCandidateInstanceTypeRequest) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *DescribeCandidateInstanceTypeRequest) GetIpv6AddressCount() *int32 {
	return s.Ipv6AddressCount
}

func (s *DescribeCandidateInstanceTypeRequest) GetMaxPrice() *float32 {
	return s.MaxPrice
}

func (s *DescribeCandidateInstanceTypeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCandidateInstanceTypeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCandidateInstanceTypeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeCandidateInstanceTypeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCandidateInstanceTypeRequest) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *DescribeCandidateInstanceTypeRequest) GetSystemDiskCategories() []*string {
	return s.SystemDiskCategories
}

func (s *DescribeCandidateInstanceTypeRequest) GetZoneIds() []*string {
	return s.ZoneIds
}

func (s *DescribeCandidateInstanceTypeRequest) SetAllowCrossAz(v bool) *DescribeCandidateInstanceTypeRequest {
	s.AllowCrossAz = &v
	return s
}

func (s *DescribeCandidateInstanceTypeRequest) SetAllowDifferentGeneration(v bool) *DescribeCandidateInstanceTypeRequest {
	s.AllowDifferentGeneration = &v
	return s
}

func (s *DescribeCandidateInstanceTypeRequest) SetDataDiskCategories(v []*string) *DescribeCandidateInstanceTypeRequest {
	s.DataDiskCategories = v
	return s
}

func (s *DescribeCandidateInstanceTypeRequest) SetImageFamily(v string) *DescribeCandidateInstanceTypeRequest {
	s.ImageFamily = &v
	return s
}

func (s *DescribeCandidateInstanceTypeRequest) SetImageId(v string) *DescribeCandidateInstanceTypeRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeCandidateInstanceTypeRequest) SetImageName(v string) *DescribeCandidateInstanceTypeRequest {
	s.ImageName = &v
	return s
}

func (s *DescribeCandidateInstanceTypeRequest) SetInstanceTypes(v []*string) *DescribeCandidateInstanceTypeRequest {
	s.InstanceTypes = v
	return s
}

func (s *DescribeCandidateInstanceTypeRequest) SetIpv6AddressCount(v int32) *DescribeCandidateInstanceTypeRequest {
	s.Ipv6AddressCount = &v
	return s
}

func (s *DescribeCandidateInstanceTypeRequest) SetMaxPrice(v float32) *DescribeCandidateInstanceTypeRequest {
	s.MaxPrice = &v
	return s
}

func (s *DescribeCandidateInstanceTypeRequest) SetOwnerId(v int64) *DescribeCandidateInstanceTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCandidateInstanceTypeRequest) SetRegionId(v string) *DescribeCandidateInstanceTypeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCandidateInstanceTypeRequest) SetResourceOwnerAccount(v string) *DescribeCandidateInstanceTypeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCandidateInstanceTypeRequest) SetResourceOwnerId(v int64) *DescribeCandidateInstanceTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCandidateInstanceTypeRequest) SetSpotStrategy(v string) *DescribeCandidateInstanceTypeRequest {
	s.SpotStrategy = &v
	return s
}

func (s *DescribeCandidateInstanceTypeRequest) SetSystemDiskCategories(v []*string) *DescribeCandidateInstanceTypeRequest {
	s.SystemDiskCategories = v
	return s
}

func (s *DescribeCandidateInstanceTypeRequest) SetZoneIds(v []*string) *DescribeCandidateInstanceTypeRequest {
	s.ZoneIds = v
	return s
}

func (s *DescribeCandidateInstanceTypeRequest) Validate() error {
	return dara.Validate(s)
}
