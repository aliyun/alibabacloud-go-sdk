// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCores(v int32) *DescribeAvailableResourceRequest
	GetCores() *int32
	SetDataDiskCategory(v string) *DescribeAvailableResourceRequest
	GetDataDiskCategory() *string
	SetDedicatedHostId(v string) *DescribeAvailableResourceRequest
	GetDedicatedHostId() *string
	SetDestinationResource(v string) *DescribeAvailableResourceRequest
	GetDestinationResource() *string
	SetInstanceChargeType(v string) *DescribeAvailableResourceRequest
	GetInstanceChargeType() *string
	SetInstanceType(v string) *DescribeAvailableResourceRequest
	GetInstanceType() *string
	SetIoOptimized(v string) *DescribeAvailableResourceRequest
	GetIoOptimized() *string
	SetMemory(v float32) *DescribeAvailableResourceRequest
	GetMemory() *float32
	SetNetworkCategory(v string) *DescribeAvailableResourceRequest
	GetNetworkCategory() *string
	SetOwnerAccount(v string) *DescribeAvailableResourceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAvailableResourceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeAvailableResourceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeAvailableResourceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAvailableResourceRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *DescribeAvailableResourceRequest
	GetResourceType() *string
	SetScope(v string) *DescribeAvailableResourceRequest
	GetScope() *string
	SetSpotDuration(v int32) *DescribeAvailableResourceRequest
	GetSpotDuration() *int32
	SetSpotStrategy(v string) *DescribeAvailableResourceRequest
	GetSpotStrategy() *string
	SetSystemDiskCategory(v string) *DescribeAvailableResourceRequest
	GetSystemDiskCategory() *string
	SetZoneId(v string) *DescribeAvailableResourceRequest
	GetZoneId() *string
}

type DescribeAvailableResourceRequest struct {
	// The number of vCPU cores for the instance type. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// This parameter applies only when `DestinationResource` is set to `InstanceType`.
	//
	// example:
	//
	// 2
	Cores *int32 `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// The category of the data disk. Valid values:
	//
	// - cloud: basic cloud disk.
	//
	// - cloud_efficiency: ultra cloud disk.
	//
	// - cloud_ssd: SSD cloud disk.
	//
	// - ephemeral_ssd: local SSD disk.
	//
	// - cloud_essd: ESSD cloud disk.
	//
	// - cloud_auto: ESSD AutoPL cloud disk.
	//
	// <props="china">
	//
	// - cloud_essd_entry: ESSD Entry cloud disk.
	//
	// example:
	//
	// cloud_ssd
	DataDiskCategory *string `json:"DataDiskCategory,omitempty" xml:"DataDiskCategory,omitempty"`
	// The ID of the dedicated host.
	//
	// example:
	//
	// dh-bp165p6xk2tlw61e****
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	// The type of resource to query. Valid values:
	//
	// - Zone: availability zone.
	//
	// - IoOptimized: I/O optimized.
	//
	// - InstanceType: instance type.
	//
	// - Network: network type.
	//
	// - ddh: dedicated host.
	//
	// - SystemDisk: system disk.
	//
	// - DataDisk: data disk.
	//
	// > When `DestinationResource` is set to `SystemDisk`, you must specify `InstanceType` because available system disks depend on the instance type.
	//
	// For details on how to specify this parameter, see the **API description*	- section.
	//
	// This parameter is required.
	//
	// example:
	//
	// InstanceType
	DestinationResource *string `json:"DestinationResource,omitempty" xml:"DestinationResource,omitempty"`
	// The billing method of the resource. For more information, see [Billing overview](https://help.aliyun.com/document_detail/25398.html). Valid values:
	//
	// - PrePaid: The subscription billing method.
	//
	// - PostPaid: The pay-as-you-go billing method.
	//
	// Default value: PostPaid.
	//
	// example:
	//
	// PrePaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The instance type. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html). You can also call the [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) operation to get the latest list of instance types.
	//
	// For details on how to specify this parameter, see the **API description*	- section.
	//
	// example:
	//
	// ecs.g5.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Specifies whether the instance is I/O optimized. Valid values:
	//
	// - none: The instance is not I/O optimized.
	//
	// - optimized: The instance is I/O optimized.
	//
	// Default value: optimized.
	//
	// example:
	//
	// optimized
	IoOptimized *string `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	// The memory size for the instance type, in GiB. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// This parameter applies only when `DestinationResource` is set to `InstanceType`.
	//
	// example:
	//
	// 8.0
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The network type. Valid values:
	//
	// - vpc: VPC.
	//
	// - classic: classic network.
	//
	// example:
	//
	// vpc
	NetworkCategory *string `json:"NetworkCategory,omitempty" xml:"NetworkCategory,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to get the latest list of Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The resource type. Valid values:
	//
	// - instance: ECS instance.
	//
	// - disk: cloud disk.
	//
	// - reservedinstance: reserved instance.
	//
	// - ddh: dedicated host.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The scope of the reserved instance. Valid values:
	//
	// - Region: The reserved instance is scoped to a region.
	//
	// - Zone: The reserved instance is scoped to an availability zone.
	//
	// example:
	//
	// Region
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The protection period for the spot instance. Unit: hours. Default value: 1. Valid values:
	//
	// - 1: Alibaba Cloud guarantees that the instance will not be automatically reclaimed within 1 hour of creation. After the 1-hour protection period ends, the system compares the bid price with the market price and checks the resource inventory to determine whether to retain or reclaim the instance.
	//
	// - 0: Alibaba Cloud does not guarantee that the instance runs for 1 hour. The system compares the bid price with the market price and checks the resource inventory to determine whether to retain or reclaim the instance.
	//
	// Alibaba Cloud sends a notification through ECS system events 5 minutes before reclaiming an instance. Spot instances are billed by the second. We recommend that you select a protection period based on the expected runtime of your tasks.
	//
	// > This parameter applies only when `InstanceChargeType` is set to `PostPaid` and `SpotStrategy` is set to `SpotWithPriceLimit` or `SpotAsPriceGo`.
	//
	// example:
	//
	// 1
	SpotDuration *int32 `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	// The bidding strategy for pay-as-you-go instances. Valid values:
	//
	// - NoSpot: A standard pay-as-you-go instance.
	//
	// - SpotWithPriceLimit: A spot instance for which you specify a maximum hourly price.
	//
	// - SpotAsPriceGo: A spot instance for which the system automatically bids based on the current market price, up to the pay-as-you-go price.
	//
	// Default value: NoSpot.
	//
	// This parameter applies only when `InstanceChargeType` is set to `PostPaid`.
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The category of the system disk. Valid values:
	//
	// - cloud: basic cloud disk.
	//
	// - cloud_efficiency: ultra cloud disk.
	//
	// - cloud_ssd: SSD cloud disk.
	//
	// - ephemeral_ssd: local SSD disk.
	//
	// - cloud_essd: ESSD cloud disk.
	//
	// - cloud_auto: ESSD AutoPL cloud disk.
	//
	// <props="china">
	//
	// - cloud_essd_entry: ESSD Entry cloud disk.
	//
	//
	//
	//
	// About the default value:
	//
	// - If `InstanceType` specifies a discontinued instance type, the default value is `cloud`.
	//
	// - Otherwise, the default value is `cloud_efficiency`. <props="china">After January 30, 2026, for instance types that support only ESSD cloud disks, the default value changes from cloud_efficiency to cloud_essd PL0. For more information, see the [official announcement](https://www.aliyun.com/notice/117844).
	//
	// > When `ResourceType` is set to `instance` and `DestinationResource` is set to `DataDisk`, this parameter is required.
	//
	// example:
	//
	// cloud_ssd
	SystemDiskCategory *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	// The ID of the availability zone.
	//
	// This parameter has no default value. If you do not specify this parameter, the operation returns resources that meet the query conditions in all availability zones within the specified region.
	//
	// example:
	//
	// cn-hangzhou-e
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAvailableResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceRequest) GetCores() *int32 {
	return s.Cores
}

func (s *DescribeAvailableResourceRequest) GetDataDiskCategory() *string {
	return s.DataDiskCategory
}

func (s *DescribeAvailableResourceRequest) GetDedicatedHostId() *string {
	return s.DedicatedHostId
}

func (s *DescribeAvailableResourceRequest) GetDestinationResource() *string {
	return s.DestinationResource
}

func (s *DescribeAvailableResourceRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeAvailableResourceRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeAvailableResourceRequest) GetIoOptimized() *string {
	return s.IoOptimized
}

func (s *DescribeAvailableResourceRequest) GetMemory() *float32 {
	return s.Memory
}

func (s *DescribeAvailableResourceRequest) GetNetworkCategory() *string {
	return s.NetworkCategory
}

func (s *DescribeAvailableResourceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAvailableResourceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAvailableResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAvailableResourceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAvailableResourceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAvailableResourceRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeAvailableResourceRequest) GetScope() *string {
	return s.Scope
}

func (s *DescribeAvailableResourceRequest) GetSpotDuration() *int32 {
	return s.SpotDuration
}

func (s *DescribeAvailableResourceRequest) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *DescribeAvailableResourceRequest) GetSystemDiskCategory() *string {
	return s.SystemDiskCategory
}

func (s *DescribeAvailableResourceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeAvailableResourceRequest) SetCores(v int32) *DescribeAvailableResourceRequest {
	s.Cores = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetDataDiskCategory(v string) *DescribeAvailableResourceRequest {
	s.DataDiskCategory = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetDedicatedHostId(v string) *DescribeAvailableResourceRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetDestinationResource(v string) *DescribeAvailableResourceRequest {
	s.DestinationResource = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetInstanceChargeType(v string) *DescribeAvailableResourceRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetInstanceType(v string) *DescribeAvailableResourceRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetIoOptimized(v string) *DescribeAvailableResourceRequest {
	s.IoOptimized = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetMemory(v float32) *DescribeAvailableResourceRequest {
	s.Memory = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetNetworkCategory(v string) *DescribeAvailableResourceRequest {
	s.NetworkCategory = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetOwnerAccount(v string) *DescribeAvailableResourceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetOwnerId(v int64) *DescribeAvailableResourceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetRegionId(v string) *DescribeAvailableResourceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetResourceOwnerAccount(v string) *DescribeAvailableResourceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetResourceOwnerId(v int64) *DescribeAvailableResourceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetResourceType(v string) *DescribeAvailableResourceRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetScope(v string) *DescribeAvailableResourceRequest {
	s.Scope = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetSpotDuration(v int32) *DescribeAvailableResourceRequest {
	s.SpotDuration = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetSpotStrategy(v string) *DescribeAvailableResourceRequest {
	s.SpotStrategy = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetSystemDiskCategory(v string) *DescribeAvailableResourceRequest {
	s.SystemDiskCategory = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetZoneId(v string) *DescribeAvailableResourceRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) Validate() error {
	return dara.Validate(s)
}
