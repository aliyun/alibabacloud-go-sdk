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
	// The disk categories of the data disks. The disk categories that do not match the specified criteria are returned after you call this operation.
	//
	// >  If you do not specify the scaling group ID, you must specify this parameter.
	DataDiskCategories []*string `json:"DataDiskCategories,omitempty" xml:"DataDiskCategories,omitempty" type:"Repeated"`
	// The name of the image family. You can specify the ImageFamily request parameter to obtain the most recent available images in the current image family for instance creation. If you specify ImageId, you cannot specify ImageFamily.
	//
	// >  If you do not specify the scaling group ID, you must specify at least one of ImageId, ImageName, and ImageFamily.
	//
	// example:
	//
	// CentOS7
	ImageFamily *string `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	// The ID of the image file that provides the image resource for Auto Scaling to create instances.
	//
	// >  If you do not specify the scaling group ID, you must specify at least one of ImageId, ImageName, and ImageFamily.
	//
	// example:
	//
	// centos6u5_64_20G_aliaegis****.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the image. Each image name must be unique in a region. If you specify ImageId, ImageName is ignored.
	//
	// You cannot use ImageName to specify an Alibaba Cloud Marketplace image.
	//
	// >  If you do not specify the scaling group ID, you must specify at least one of ImageId, ImageName, and ImageFamily.
	//
	// example:
	//
	// ubuntu_18_04_x64_20G_alibase_20231225.vhd
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The instance types. The instance types specified by this parameter overwrite the instance types specified in the scaling configuration.
	InstanceTypes []*string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// The number of IPv6 addresses. If the instance type that you specified does meet the requirement for the number of IPv6 addresses, the scaling strength is weak.
	//
	// >  If you do not specify the scaling group ID, you must specify this parameter.
	//
	// example:
	//
	// 1
	Ipv6AddressCount *int32 `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	// **
	//
	// **Warning*	- This parameter is deprecated. We recommend that you use SpotStrategy.
	//
	// The preemption policy that you want to apply to pay-as-you-go instances. The preemption policy specified by this parameter overwrites the preemption policy specified in the scaling configuration. Valid values:
	//
	// 	- NoSpot: The instances are created as regular pay-as-you-go instances.
	//
	// 	- SpotWithPriceLimit: The instances are created as preemptible instances with a user-defined maximum hourly price.
	//
	// 	- SpotAsPriceGo: The instances are created as preemptible instances for which the market price at the time of purchase is automatically used as the bidding price.
	//
	// Default value: NoSpot.
	//
	// example:
	//
	// NoSpot
	PriorityStrategy *string `json:"PriorityStrategy,omitempty" xml:"PriorityStrategy,omitempty"`
	// The region ID of the scaling group.
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
	// The IDs of the scaling groups that you want to query.
	ScalingGroupIds []*string `json:"ScalingGroupIds,omitempty" xml:"ScalingGroupIds,omitempty" type:"Repeated"`
	// The instance bidding policy. Valid values:
	//
	// 	- NoSpot: The instances are created as pay-as-you-go instances.
	//
	// 	- SpotWithPriceLimit: The instances are created as preemptible instances with a user-defined maximum hourly price.
	//
	// 	- SpotAsPriceGo: The instances are created as preemptible instances for which the market price at the time of purchase is used as the bid price.
	//
	// Default value: NoSpot.
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The categories of the system disks. The categories of the system disks specified by this parameter overwrite the categories of the system disks specified in the scaling configuration. Valid values:
	//
	// 	- cloud: basic disk.
	//
	// 	- cloud_efficiency: ultra disk.
	//
	// 	- cloud_ssd: standard SSD.
	//
	// 	- cloud_essd: Enterprise SSD (ESSD).
	//
	// >  If you do not specify the scaling group ID, you must specify this parameter.
	SystemDiskCategories []*string `json:"SystemDiskCategories,omitempty" xml:"SystemDiskCategories,omitempty" type:"Repeated"`
	// The vSwitch IDs.
	//
	// >  If you do not specify the scaling group ID, you must specify this parameter.
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
