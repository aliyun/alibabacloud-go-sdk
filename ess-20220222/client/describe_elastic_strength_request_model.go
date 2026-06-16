// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticStrengthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataDiskCategories(v []*string) *DescribeElasticStrengthRequest
	GetDataDiskCategories() []*string
	SetImageFamily(v string) *DescribeElasticStrengthRequest
	GetImageFamily() *string
	SetImageId(v string) *DescribeElasticStrengthRequest
	GetImageId() *string
	SetImageName(v string) *DescribeElasticStrengthRequest
	GetImageName() *string
	SetInstanceTypes(v []*string) *DescribeElasticStrengthRequest
	GetInstanceTypes() []*string
	SetIpv6AddressCount(v int32) *DescribeElasticStrengthRequest
	GetIpv6AddressCount() *int32
	SetPriorityStrategy(v string) *DescribeElasticStrengthRequest
	GetPriorityStrategy() *string
	SetRegionId(v string) *DescribeElasticStrengthRequest
	GetRegionId() *string
	SetScalingGroupId(v string) *DescribeElasticStrengthRequest
	GetScalingGroupId() *string
	SetScalingGroupIds(v []*string) *DescribeElasticStrengthRequest
	GetScalingGroupIds() []*string
	SetSpotStrategy(v string) *DescribeElasticStrengthRequest
	GetSpotStrategy() *string
	SetSystemDiskCategories(v []*string) *DescribeElasticStrengthRequest
	GetSystemDiskCategories() []*string
	SetVSwitchIds(v []*string) *DescribeElasticStrengthRequest
	GetVSwitchIds() []*string
}

type DescribeElasticStrengthRequest struct {
	// A list of data disk categories used to evaluate elastic strength. If a category is incompatible, the response identifies the specific mismatched category.
	//
	// > You can specify this parameter if `ScalingGroupId` is not specified.
	DataDiskCategories []*string `json:"DataDiskCategories,omitempty" xml:"DataDiskCategories,omitempty" type:"Repeated"`
	// The name of the image family. You can set this parameter to use the latest available image from the specified image family to create instances. If you specify ImageId, this parameter is ignored.
	//
	// > If `ScalingGroupId` is not specified, you must specify at least one of `ImageId`, `ImageName`, or `ImageFamily`.
	//
	// example:
	//
	// CentOS7
	ImageFamily *string `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	// The ID of the image used to create instances.
	//
	// > If `ScalingGroupId` is not specified, you must specify at least one of `ImageId`, `ImageName`, or `ImageFamily`.
	//
	// example:
	//
	// centos6u5_64_20G_aliaegis****.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the image. The name must be unique within a region. If you specify `ImageId`, this parameter is ignored.
	//
	// You cannot use this parameter to specify a Marketplace image.
	//
	// > If `ScalingGroupId` is not specified, you must specify at least one of `ImageId`, `ImageName`, or `ImageFamily`.
	//
	// example:
	//
	// ubuntu_18_04_x64_20G_alibase_20231225.vhd
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// A list of ECS instance types. If specified, this parameter overrides the instance types in the scaling configuration.
	InstanceTypes []*string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// The number of IPv6 addresses to be configured for each instance. The elastic strength is lowered for instance types that do not support the specified number of IPv6 addresses.
	//
	// > You can specify this parameter if `ScalingGroupId` is not specified.
	//
	// example:
	//
	// 1
	Ipv6AddressCount *int32 `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	// 	Warning: This parameter is deprecated. Use `SpotStrategy` instead.
	//
	// The spot strategy for pay-as-you-go instances. If specified, this parameter overrides the spot strategy in the scaling configuration. Valid values:
	//
	// - `NoSpot`: A regular pay-as-you-go instance.
	//
	// - `SpotWithPriceLimit`: A spot instance with a specified maximum price.
	//
	// - `SpotAsPriceGo`: A spot instance where the system automatically bids at the current market price.
	//
	// Default value: `NoSpot`.
	//
	// example:
	//
	// NoSpot
	PriorityStrategy *string `json:"PriorityStrategy,omitempty" xml:"PriorityStrategy,omitempty"`
	// The ID of the region where the scaling group is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the scaling group.
	//
	// example:
	//
	// asg-bp18p2yfxow2dloq****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// The IDs of one or more scaling groups to query in a batch operation.
	ScalingGroupIds []*string `json:"ScalingGroupIds,omitempty" xml:"ScalingGroupIds,omitempty" type:"Repeated"`
	// The spot strategy for instances. Valid values:
	//
	// - `NoSpot`: A regular pay-as-you-go instance.
	//
	// - `SpotWithPriceLimit`: A spot instance with a specified maximum price.
	//
	// - `SpotAsPriceGo`: A spot instance where the system automatically bids at the current market price.
	//
	// Default value: `NoSpot`.
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// A list of system disk categories. If specified, this parameter overrides the system disk categories in the scaling configuration. Valid values:
	//
	// - `cloud`: Basic Cloud Disk.
	//
	// - `cloud_efficiency`: Ultra Cloud Disk.
	//
	// - `cloud_ssd`: Standard SSD.
	//
	// - `cloud_essd`: ESSD.
	//
	// > This parameter is required if `ScalingGroupId` is not specified.
	SystemDiskCategories []*string `json:"SystemDiskCategories,omitempty" xml:"SystemDiskCategories,omitempty" type:"Repeated"`
	// A list of VSwitch IDs.
	//
	// > This parameter is required if `ScalingGroupId` is not specified.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
}

func (s DescribeElasticStrengthRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticStrengthRequest) GoString() string {
	return s.String()
}

func (s *DescribeElasticStrengthRequest) GetDataDiskCategories() []*string {
	return s.DataDiskCategories
}

func (s *DescribeElasticStrengthRequest) GetImageFamily() *string {
	return s.ImageFamily
}

func (s *DescribeElasticStrengthRequest) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeElasticStrengthRequest) GetImageName() *string {
	return s.ImageName
}

func (s *DescribeElasticStrengthRequest) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *DescribeElasticStrengthRequest) GetIpv6AddressCount() *int32 {
	return s.Ipv6AddressCount
}

func (s *DescribeElasticStrengthRequest) GetPriorityStrategy() *string {
	return s.PriorityStrategy
}

func (s *DescribeElasticStrengthRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeElasticStrengthRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DescribeElasticStrengthRequest) GetScalingGroupIds() []*string {
	return s.ScalingGroupIds
}

func (s *DescribeElasticStrengthRequest) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *DescribeElasticStrengthRequest) GetSystemDiskCategories() []*string {
	return s.SystemDiskCategories
}

func (s *DescribeElasticStrengthRequest) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *DescribeElasticStrengthRequest) SetDataDiskCategories(v []*string) *DescribeElasticStrengthRequest {
	s.DataDiskCategories = v
	return s
}

func (s *DescribeElasticStrengthRequest) SetImageFamily(v string) *DescribeElasticStrengthRequest {
	s.ImageFamily = &v
	return s
}

func (s *DescribeElasticStrengthRequest) SetImageId(v string) *DescribeElasticStrengthRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeElasticStrengthRequest) SetImageName(v string) *DescribeElasticStrengthRequest {
	s.ImageName = &v
	return s
}

func (s *DescribeElasticStrengthRequest) SetInstanceTypes(v []*string) *DescribeElasticStrengthRequest {
	s.InstanceTypes = v
	return s
}

func (s *DescribeElasticStrengthRequest) SetIpv6AddressCount(v int32) *DescribeElasticStrengthRequest {
	s.Ipv6AddressCount = &v
	return s
}

func (s *DescribeElasticStrengthRequest) SetPriorityStrategy(v string) *DescribeElasticStrengthRequest {
	s.PriorityStrategy = &v
	return s
}

func (s *DescribeElasticStrengthRequest) SetRegionId(v string) *DescribeElasticStrengthRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeElasticStrengthRequest) SetScalingGroupId(v string) *DescribeElasticStrengthRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeElasticStrengthRequest) SetScalingGroupIds(v []*string) *DescribeElasticStrengthRequest {
	s.ScalingGroupIds = v
	return s
}

func (s *DescribeElasticStrengthRequest) SetSpotStrategy(v string) *DescribeElasticStrengthRequest {
	s.SpotStrategy = &v
	return s
}

func (s *DescribeElasticStrengthRequest) SetSystemDiskCategories(v []*string) *DescribeElasticStrengthRequest {
	s.SystemDiskCategories = v
	return s
}

func (s *DescribeElasticStrengthRequest) SetVSwitchIds(v []*string) *DescribeElasticStrengthRequest {
	s.VSwitchIds = v
	return s
}

func (s *DescribeElasticStrengthRequest) Validate() error {
	return dara.Validate(s)
}
