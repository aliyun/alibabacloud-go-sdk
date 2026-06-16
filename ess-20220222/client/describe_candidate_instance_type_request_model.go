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
	// Specifies whether to include vSwitches from other availability zones as candidates.
	//
	// > The instance types remain unchanged. Only new availability zones are added as candidates. If a scaling group fails to scale out in all selected availability zones due to issues such as insufficient inventory, Auto Scaling automatically adds a vSwitch in a new availability zone to the scaling group based on this setting.
	//
	// > For example, if a scaling group is configured for the cn-hangzhou-h and cn-hangzhou-g availability zones and a scale-out fails in both zones, Auto Scaling may create a vSwitch in the cn-hangzhou-k availability zone and add it to the scaling group based on real-time inventory.
	//
	// example:
	//
	// true
	AllowCrossAz *bool `json:"AllowCrossAz,omitempty" xml:"AllowCrossAz,omitempty"`
	// Specifies whether to include instance types from other generations as candidates.
	//
	// - For example, if the current instance type is ecs.c7.large, you can set this parameter to true to include instance types such as ecs.c6.large and ecs.c8.large in the list of candidates.
	//
	// example:
	//
	// true
	AllowDifferentGeneration *bool `json:"AllowDifferentGeneration,omitempty" xml:"AllowDifferentGeneration,omitempty"`
	// The data disk categories, ordered by priority from high to low. If Auto Scaling cannot create a data disk by using a higher-priority category, it tries the next one in the list.
	DataDiskCategories []*string `json:"DataDiskCategories,omitempty" xml:"DataDiskCategories,omitempty" type:"Repeated"`
	// The name of the image family. When specified, the latest image in this family is used to create instances. This parameter cannot be used with ImageId.
	//
	// > If you do not specify the scaling group ID, you must specify at least one of ImageId, ImageName, and ImageFamily.
	//
	// example:
	//
	// hangzhou-daily-update
	ImageFamily *string `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	// The ID of the image used to create instances.
	//
	// > If the specified image contains both a system disk and data disks, any existing data disk information in the scaling configuration is cleared.
	//
	// example:
	//
	// centos6u5_64_20G_aliaegis****.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the image. The name must be unique within a region. You cannot use this parameter to specify an image from Alibaba Cloud Marketplace.
	//
	// > This parameter is an alternative to the `ImageId` parameter. If you specify `ImageId`, `ImageName` is ignored.
	//
	// example:
	//
	// centos6u5_64_20G_aliaegis_20140703.vhd
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The specified ECS instance types.
	InstanceTypes []*string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// The number of IPv6 addresses.
	//
	// example:
	//
	// 1
	Ipv6AddressCount *int32 `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	// The maximum price for a candidate instance type.
	//
	// example:
	//
	// 2
	MaxPrice *float32 `json:"MaxPrice,omitempty" xml:"MaxPrice,omitempty"`
	OwnerId  *int64   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the scaling group is located.
	//
	// example:
	//
	// cn-qingdao
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The bidding strategy for pay-as-you-go instances. Valid values:
	//
	// - NoSpot: a standard pay-as-you-go instance.
	//
	// - SpotWithPriceLimit: a spot instance with a user-defined maximum price.
	//
	// - SpotAsPriceGo: a spot instance where the system automatically bids at the market price.
	//
	// Default value: NoSpot.
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The system disk categories, ordered by priority from high to low. If Auto Scaling cannot create a system disk by using a higher-priority category, it tries the next one in the list.
	SystemDiskCategories []*string `json:"SystemDiskCategories,omitempty" xml:"SystemDiskCategories,omitempty" type:"Repeated"`
	// The specified availability zones.
	ZoneIds []*string `json:"ZoneIds,omitempty" xml:"ZoneIds,omitempty" type:"Repeated"`
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
