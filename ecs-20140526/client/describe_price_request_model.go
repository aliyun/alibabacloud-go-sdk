// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataDisk(v []*DescribePriceRequestDataDisk) *DescribePriceRequest
	GetDataDisk() []*DescribePriceRequestDataDisk
	SetSchedulerOptions(v *DescribePriceRequestSchedulerOptions) *DescribePriceRequest
	GetSchedulerOptions() *DescribePriceRequestSchedulerOptions
	SetSystemDisk(v *DescribePriceRequestSystemDisk) *DescribePriceRequest
	GetSystemDisk() *DescribePriceRequestSystemDisk
	SetAmount(v int32) *DescribePriceRequest
	GetAmount() *int32
	SetAssuranceTimes(v string) *DescribePriceRequest
	GetAssuranceTimes() *string
	SetCapacity(v int32) *DescribePriceRequest
	GetCapacity() *int32
	SetDedicatedHostType(v string) *DescribePriceRequest
	GetDedicatedHostType() *string
	SetImageId(v string) *DescribePriceRequest
	GetImageId() *string
	SetInstanceAmount(v int32) *DescribePriceRequest
	GetInstanceAmount() *int32
	SetInstanceCpuCoreCount(v int32) *DescribePriceRequest
	GetInstanceCpuCoreCount() *int32
	SetInstanceNetworkType(v string) *DescribePriceRequest
	GetInstanceNetworkType() *string
	SetInstanceType(v string) *DescribePriceRequest
	GetInstanceType() *string
	SetInstanceTypeList(v []*string) *DescribePriceRequest
	GetInstanceTypeList() []*string
	SetInternetChargeType(v string) *DescribePriceRequest
	GetInternetChargeType() *string
	SetInternetMaxBandwidthOut(v int32) *DescribePriceRequest
	GetInternetMaxBandwidthOut() *int32
	SetIoOptimized(v string) *DescribePriceRequest
	GetIoOptimized() *string
	SetIsp(v string) *DescribePriceRequest
	GetIsp() *string
	SetOfferingType(v string) *DescribePriceRequest
	GetOfferingType() *string
	SetOwnerAccount(v string) *DescribePriceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribePriceRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *DescribePriceRequest
	GetPeriod() *int32
	SetPlatform(v string) *DescribePriceRequest
	GetPlatform() *string
	SetPriceUnit(v string) *DescribePriceRequest
	GetPriceUnit() *string
	SetRecurrenceRules(v []*DescribePriceRequestRecurrenceRules) *DescribePriceRequest
	GetRecurrenceRules() []*DescribePriceRequestRecurrenceRules
	SetRegionId(v string) *DescribePriceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribePriceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribePriceRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *DescribePriceRequest
	GetResourceType() *string
	SetScope(v string) *DescribePriceRequest
	GetScope() *string
	SetSpotDuration(v int32) *DescribePriceRequest
	GetSpotDuration() *int32
	SetSpotStrategy(v string) *DescribePriceRequest
	GetSpotStrategy() *string
	SetStartTime(v string) *DescribePriceRequest
	GetStartTime() *string
	SetZoneId(v string) *DescribePriceRequest
	GetZoneId() *string
}

type DescribePriceRequest struct {
	DataDisk         []*DescribePriceRequestDataDisk       `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	SchedulerOptions *DescribePriceRequestSchedulerOptions `json:"SchedulerOptions,omitempty" xml:"SchedulerOptions,omitempty" type:"Struct"`
	SystemDisk       *DescribePriceRequestSystemDisk       `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// The number of resources for which to query prices. Valid values: 1–1000.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The number of times the elasticity assurance can be used. Set this to `Unlimited`, which allows the assurance to be used any number of times during its effective period.
	//
	// Default value: `Unlimited`.
	//
	// example:
	//
	// Unlimited
	AssuranceTimes *string `json:"AssuranceTimes,omitempty" xml:"AssuranceTimes,omitempty"`
	// The memory capacity for the elasticity assurance. Unit: GiB.
	//
	// example:
	//
	// 1024
	Capacity *int32 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The dedicated host type. You can call the [DescribeDedicatedHostTypes](https://help.aliyun.com/document_detail/134240.html) operation to query dedicated host types.
	//
	// example:
	//
	// ddh.c5
	DedicatedHostType *string `json:"DedicatedHostType,omitempty" xml:"DedicatedHostType,omitempty"`
	// This parameter is valid only when `ResourceType` is set to `instance`.
	//
	// The ID of the image. The image provides the runtime environment for the instance. You can call the [DescribeImages](https://help.aliyun.com/document_detail/25534.html) operation to query available images. If you do not specify this parameter, the system queries prices for Linux instances by default.
	//
	// example:
	//
	// centos_7_05_64_20G_alibase_20181212.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The number of instances to include in the reserved instance offering.
	//
	// Valid values: 1–1000.
	//
	// example:
	//
	// 100
	InstanceAmount *int32 `json:"InstanceAmount,omitempty" xml:"InstanceAmount,omitempty"`
	// The total number of vCPUs for instances that are covered by the elasticity assurance. When you call this operation, the system calculates the number of supported instances based on the specified `InstanceType` and rounds the value up to the nearest integer.
	//
	// > When you query the price of an elasticity assurance, you can specify only one of the `InstanceCpuCoreCount` and `InstanceAmount` parameters.
	//
	// example:
	//
	// 1024
	InstanceCpuCoreCount *int32 `json:"InstanceCpuCoreCount,omitempty" xml:"InstanceCpuCoreCount,omitempty"`
	// The network type of the instance. Valid values:
	//
	// - `classic`: classic network
	//
	// - `vpc`: VPC
	//
	// Default value: `vpc`.
	//
	// example:
	//
	// vpc
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	// The instance type. This parameter is required when `ResourceType` is set to `instance`. For more information, see [Instance type families](https://help.aliyun.com/document_detail/25378.html) or call the [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) operation to query the instance types.
	//
	// example:
	//
	// ecs.g6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The instance type. You can specify only one instance type for an elasticity assurance of the `Unlimited` type.
	//
	// example:
	//
	// ecs.g6.xlarge
	InstanceTypeList []*string `json:"InstanceTypeList,omitempty" xml:"InstanceTypeList,omitempty" type:"Repeated"`
	// The billing method for network usage. Valid values:
	//
	// - `PayByBandwidth`: pay-by-bandwidth
	//
	// - `PayByTraffic`: pay-by-traffic
	//
	// Default value: `PayByTraffic`.
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The maximum outbound public bandwidth. Unit: Mbit/s. Valid values: 0–100.
	//
	// Default value: 0.
	//
	// example:
	//
	// 5
	InternetMaxBandwidthOut *int32 `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// Specifies whether the instance is I/O optimized. Valid values:
	//
	// - `none`: non-I/O-optimized.
	//
	// - `optimized`: I/O-optimized.
	//
	// For [generation I](https://help.aliyun.com/document_detail/55263.html) instances, the default value is `none`.
	//
	// For other instance types, the default value is `optimized`.
	//
	// example:
	//
	// optimized
	IoOptimized *string `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	// The Internet service provider (ISP). Valid values:
	//
	// - `cmcc`: China Mobile
	//
	// - `telecom`: China Telecom
	//
	// - `unicom`: China Unicom
	//
	// - `multiCarrier`: BGP (Multi-ISP)
	//
	// example:
	//
	// cmcc
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The payment option for the reserved instance. Valid values:
	//
	// - `No Upfront`
	//
	// - `Partial Upfront`
	//
	// - `All Upfront`
	//
	// example:
	//
	// All Upfront
	OfferingType *string `json:"OfferingType,omitempty" xml:"OfferingType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing duration of the resource. This parameter is used with `PriceUnit`. Valid values:
	//
	// <props="china">
	//
	// - If `PriceUnit` is set to `Month`: 1–9.
	//
	// - If `PriceUnit` is set to `Year`: 1–5.
	//
	// - If `PriceUnit` is set to `Hour`: 1.
	//
	// - If `PriceUnit` is set to `Week`: 1–4.
	//
	//
	//
	// <props="intl">
	//
	// - If `PriceUnit` is set to `Month`: 1–9.
	//
	// - If `PriceUnit` is set to `Year`: 1–5.
	//
	// - If `PriceUnit` is set to `Hour`: 1.
	//
	//
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The operating system of the instance. Valid values:
	//
	// - `Windows`: Windows Server
	//
	// - `Linux`: Linux
	//
	// example:
	//
	// Linux
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The billing cycle of the resource. Valid values:
	//
	// <props="china">
	//
	// - `Month`: For monthly pricing.
	//
	// - `Year`: For yearly pricing.
	//
	// - `Hour` (Default): For hourly pricing.
	//
	// - `Week`: For weekly pricing.
	//
	//
	//
	// <props="intl">
	//
	// - `Month`: For monthly pricing.
	//
	// - `Year`: For yearly pricing.
	//
	// - `Hour` (Default): For hourly pricing.
	//
	// example:
	//
	// Year
	PriceUnit *string `json:"PriceUnit,omitempty" xml:"PriceUnit,omitempty"`
	// The list of recurrence rules for the time-based elasticity assurance.
	//
	// <props="china">
	//
	// > The time-based elasticity assurance feature is available only in specific regions and to specific users. To use this feature, [submit a ticket](https://selfservice.console.aliyun.com/ticket/createIndex).
	//
	//
	//
	// <props="intl">
	//
	// > The time-based elasticity assurance feature is available only in specific regions and to specific users. To use this feature, [submit a ticket](https://smartservice.console.aliyun.com/service/create-ticket-intl).
	RecurrenceRules []*DescribePriceRequestRecurrenceRules `json:"RecurrenceRules,omitempty" xml:"RecurrenceRules,omitempty" type:"Repeated"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the list of Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the resource for which you want to query the price. Valid values:
	//
	// - `instance`: Query the prices of ECS instances. If you set this parameter to `instance`, you must also specify the `InstanceType` parameter.
	//
	// - `disk`: Query the prices of cloud disks. If you set this parameter to `disk`, you must also specify the `DataDisk.1.Category` and `DataDisk.1.Size` parameters.
	//
	// - `diskperformance`: Query the prices of the provisioned performance of an ESSD AutoPL cloud disk. You must also specify the `DataDisk.1.Category` and `DataDisk.1.ProvisionedIops` parameters.
	//
	// - `bandwidth`: Query the prices of network bandwidth.
	//
	// - `ddh`: Query the prices of dedicated hosts.
	//
	// - `ElasticityAssurance`: Query the prices of Elasticity Assurance. If you set this parameter to `ElasticityAssurance`, you must also specify the `InstanceType` parameter.
	//
	// - `CapacityReservation`: Query the prices of Capacity Reservation. If you set this parameter to `CapacityReservation`, you must also specify the `InstanceType` parameter.
	//
	// Default value: `instance`.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The scope of the reserved instance. Valid values:
	//
	// - `Region`: region-scoped
	//
	// - `Zone`: zone-scoped
	//
	// Default value: `Region`.
	//
	// example:
	//
	// Zone
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The protection period of the spot instance. Unit: hours. Default value: 1. Valid values:
	//
	// - `1`: Alibaba Cloud does not automatically release the instance within 1 hour. After the 1-hour protection period ends, the system checks the market price and resource inventory to determine whether to retain or release the instance.
	//
	// - `0`: The instance has no protection period. The system checks the market price and resource inventory to determine whether to retain or release the instance.
	//
	// Alibaba Cloud sends you an ECS system event five minutes before the instance is released. Spot instances are billed by the second. Select a protection period based on the time required to complete your task.
	//
	// > This parameter is valid only when `SpotStrategy` is set to `SpotWithPriceLimit` or `SpotAsPriceGo`.
	//
	// example:
	//
	// 1
	SpotDuration *int32 `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	// The preemption policy for the pay-as-you-go instance. Valid values:
	//
	// - `NoSpot`: A regular pay-as-you-go instance.
	//
	// - `SpotWithPriceLimit`: A spot instance for which you specify a maximum hourly price.
	//
	// - `SpotAsPriceGo`: A spot instance where the system automatically bids up to the pay-as-you-go price.
	//
	// Default value: `NoSpot`.
	//
	// > This parameter applies only when you query hourly prices, where `PriceUnit` is `Hour` and `Period` is `1`. Because these are the default values, you do not need to set them when you use `SpotStrategy`.
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The time when the time-based elasticity assurance takes effect. The time must be specified in UTC and formatted as `yyyy-MM-ddTHH:mm:ssZ` in accordance with the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard.
	//
	// example:
	//
	// 2020-10-30T06:32:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the availability zone.
	//
	// > The prices of spot instances may vary by availability zone. When you query the price of a spot instance, specify `ZoneId` to query the price for a specific availability zone.
	//
	// example:
	//
	// cn-hagzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribePriceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceRequest) GoString() string {
	return s.String()
}

func (s *DescribePriceRequest) GetDataDisk() []*DescribePriceRequestDataDisk {
	return s.DataDisk
}

func (s *DescribePriceRequest) GetSchedulerOptions() *DescribePriceRequestSchedulerOptions {
	return s.SchedulerOptions
}

func (s *DescribePriceRequest) GetSystemDisk() *DescribePriceRequestSystemDisk {
	return s.SystemDisk
}

func (s *DescribePriceRequest) GetAmount() *int32 {
	return s.Amount
}

func (s *DescribePriceRequest) GetAssuranceTimes() *string {
	return s.AssuranceTimes
}

func (s *DescribePriceRequest) GetCapacity() *int32 {
	return s.Capacity
}

func (s *DescribePriceRequest) GetDedicatedHostType() *string {
	return s.DedicatedHostType
}

func (s *DescribePriceRequest) GetImageId() *string {
	return s.ImageId
}

func (s *DescribePriceRequest) GetInstanceAmount() *int32 {
	return s.InstanceAmount
}

func (s *DescribePriceRequest) GetInstanceCpuCoreCount() *int32 {
	return s.InstanceCpuCoreCount
}

func (s *DescribePriceRequest) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *DescribePriceRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribePriceRequest) GetInstanceTypeList() []*string {
	return s.InstanceTypeList
}

func (s *DescribePriceRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribePriceRequest) GetInternetMaxBandwidthOut() *int32 {
	return s.InternetMaxBandwidthOut
}

func (s *DescribePriceRequest) GetIoOptimized() *string {
	return s.IoOptimized
}

func (s *DescribePriceRequest) GetIsp() *string {
	return s.Isp
}

func (s *DescribePriceRequest) GetOfferingType() *string {
	return s.OfferingType
}

func (s *DescribePriceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribePriceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePriceRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribePriceRequest) GetPlatform() *string {
	return s.Platform
}

func (s *DescribePriceRequest) GetPriceUnit() *string {
	return s.PriceUnit
}

func (s *DescribePriceRequest) GetRecurrenceRules() []*DescribePriceRequestRecurrenceRules {
	return s.RecurrenceRules
}

func (s *DescribePriceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePriceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribePriceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribePriceRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribePriceRequest) GetScope() *string {
	return s.Scope
}

func (s *DescribePriceRequest) GetSpotDuration() *int32 {
	return s.SpotDuration
}

func (s *DescribePriceRequest) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *DescribePriceRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribePriceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribePriceRequest) SetDataDisk(v []*DescribePriceRequestDataDisk) *DescribePriceRequest {
	s.DataDisk = v
	return s
}

func (s *DescribePriceRequest) SetSchedulerOptions(v *DescribePriceRequestSchedulerOptions) *DescribePriceRequest {
	s.SchedulerOptions = v
	return s
}

func (s *DescribePriceRequest) SetSystemDisk(v *DescribePriceRequestSystemDisk) *DescribePriceRequest {
	s.SystemDisk = v
	return s
}

func (s *DescribePriceRequest) SetAmount(v int32) *DescribePriceRequest {
	s.Amount = &v
	return s
}

func (s *DescribePriceRequest) SetAssuranceTimes(v string) *DescribePriceRequest {
	s.AssuranceTimes = &v
	return s
}

func (s *DescribePriceRequest) SetCapacity(v int32) *DescribePriceRequest {
	s.Capacity = &v
	return s
}

func (s *DescribePriceRequest) SetDedicatedHostType(v string) *DescribePriceRequest {
	s.DedicatedHostType = &v
	return s
}

func (s *DescribePriceRequest) SetImageId(v string) *DescribePriceRequest {
	s.ImageId = &v
	return s
}

func (s *DescribePriceRequest) SetInstanceAmount(v int32) *DescribePriceRequest {
	s.InstanceAmount = &v
	return s
}

func (s *DescribePriceRequest) SetInstanceCpuCoreCount(v int32) *DescribePriceRequest {
	s.InstanceCpuCoreCount = &v
	return s
}

func (s *DescribePriceRequest) SetInstanceNetworkType(v string) *DescribePriceRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribePriceRequest) SetInstanceType(v string) *DescribePriceRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribePriceRequest) SetInstanceTypeList(v []*string) *DescribePriceRequest {
	s.InstanceTypeList = v
	return s
}

func (s *DescribePriceRequest) SetInternetChargeType(v string) *DescribePriceRequest {
	s.InternetChargeType = &v
	return s
}

func (s *DescribePriceRequest) SetInternetMaxBandwidthOut(v int32) *DescribePriceRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *DescribePriceRequest) SetIoOptimized(v string) *DescribePriceRequest {
	s.IoOptimized = &v
	return s
}

func (s *DescribePriceRequest) SetIsp(v string) *DescribePriceRequest {
	s.Isp = &v
	return s
}

func (s *DescribePriceRequest) SetOfferingType(v string) *DescribePriceRequest {
	s.OfferingType = &v
	return s
}

func (s *DescribePriceRequest) SetOwnerAccount(v string) *DescribePriceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribePriceRequest) SetOwnerId(v int64) *DescribePriceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePriceRequest) SetPeriod(v int32) *DescribePriceRequest {
	s.Period = &v
	return s
}

func (s *DescribePriceRequest) SetPlatform(v string) *DescribePriceRequest {
	s.Platform = &v
	return s
}

func (s *DescribePriceRequest) SetPriceUnit(v string) *DescribePriceRequest {
	s.PriceUnit = &v
	return s
}

func (s *DescribePriceRequest) SetRecurrenceRules(v []*DescribePriceRequestRecurrenceRules) *DescribePriceRequest {
	s.RecurrenceRules = v
	return s
}

func (s *DescribePriceRequest) SetRegionId(v string) *DescribePriceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePriceRequest) SetResourceOwnerAccount(v string) *DescribePriceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePriceRequest) SetResourceOwnerId(v int64) *DescribePriceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePriceRequest) SetResourceType(v string) *DescribePriceRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribePriceRequest) SetScope(v string) *DescribePriceRequest {
	s.Scope = &v
	return s
}

func (s *DescribePriceRequest) SetSpotDuration(v int32) *DescribePriceRequest {
	s.SpotDuration = &v
	return s
}

func (s *DescribePriceRequest) SetSpotStrategy(v string) *DescribePriceRequest {
	s.SpotStrategy = &v
	return s
}

func (s *DescribePriceRequest) SetStartTime(v string) *DescribePriceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePriceRequest) SetZoneId(v string) *DescribePriceRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribePriceRequest) Validate() error {
	if s.DataDisk != nil {
		for _, item := range s.DataDisk {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SchedulerOptions != nil {
		if err := s.SchedulerOptions.Validate(); err != nil {
			return err
		}
	}
	if s.SystemDisk != nil {
		if err := s.SystemDisk.Validate(); err != nil {
			return err
		}
	}
	if s.RecurrenceRules != nil {
		for _, item := range s.RecurrenceRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePriceRequestDataDisk struct {
	// The category of data disk N. Valid values:
	//
	// - `cloud`: basic cloud disk
	//
	// - `cloud_efficiency`: efficiency cloud disk
	//
	// - `cloud_ssd`: SSD cloud disk
	//
	// - `ephemeral_ssd`: local SSD
	//
	// - `cloud_essd`: ESSD
	//
	// - `cloud_auto`: ESSD AutoPL
	//
	// <props="china">
	//
	// - `cloud_essd_entry`: ESSD Entry
	//
	//
	//
	//
	// The value of N can be 1–16.
	//
	// example:
	//
	// cloud_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The performance level of data disk N when it is an ESSD. This parameter is valid only when `DataDisk.N.Category` is set to `cloud_essd`. Valid values:
	//
	// - `PL0`
	//
	// - `PL1` (Default)
	//
	// - `PL2`
	//
	// - `PL3`
	//
	// The value of N can be 1–16.
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The size of data disk N. Unit: GiB. Valid values:
	//
	// - `cloud`: 5–2000
	//
	// - `cloud_efficiency`: 20–32768
	//
	// - `cloud_ssd`: 20–32768
	//
	// - `cloud_auto`: 1–32768
	//
	// <props="china">
	//
	// - `cloud_essd_entry`: 10–32768
	//
	//
	//
	//
	// - `cloud_essd`: The value range depends on the `DataDisk.N.PerformanceLevel`.
	//
	//   - PL0: 1–32768
	//
	//   - PL1: 20–32768
	//
	//   - PL2: 461–32768
	//
	//   - PL3: 1261–32768
	//
	// - `ephemeral_ssd`: 5–800
	//
	// The value of N can be 1–16.
	//
	// example:
	//
	// 2000
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The provisioned read/write IOPS for the ESSD AutoPL cloud disk. Valid values: 0–`min{50000, 1000 	- Capacity - Baseline IOPS}`.
	//
	// `Baseline IOPS = min{1800 + 50 	- Capacity, 50000}`.
	//
	// > This parameter is valid only when `DataDisk.N.Category` is set to `cloud_auto`. For more information, see [ESSD AutoPL cloud disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// 40000
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
}

func (s DescribePriceRequestDataDisk) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceRequestDataDisk) GoString() string {
	return s.String()
}

func (s *DescribePriceRequestDataDisk) GetCategory() *string {
	return s.Category
}

func (s *DescribePriceRequestDataDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *DescribePriceRequestDataDisk) GetSize() *int64 {
	return s.Size
}

func (s *DescribePriceRequestDataDisk) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *DescribePriceRequestDataDisk) SetCategory(v string) *DescribePriceRequestDataDisk {
	s.Category = &v
	return s
}

func (s *DescribePriceRequestDataDisk) SetPerformanceLevel(v string) *DescribePriceRequestDataDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribePriceRequestDataDisk) SetSize(v int64) *DescribePriceRequestDataDisk {
	s.Size = &v
	return s
}

func (s *DescribePriceRequestDataDisk) SetProvisionedIops(v int64) *DescribePriceRequestDataDisk {
	s.ProvisionedIops = &v
	return s
}

func (s *DescribePriceRequestDataDisk) Validate() error {
	return dara.Validate(s)
}

type DescribePriceRequestSchedulerOptions struct {
	// This parameter is valid only when `ResourceType` is set to `instance`.
	//
	// The ID of the dedicated host. You can call the [DescribeDedicatedHosts](https://help.aliyun.com/document_detail/134242.html) operation to query dedicated host IDs.
	//
	// example:
	//
	// dh-bp67acfmxazb4p****
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	// The deployment set strategy. Valid values:
	//
	// - `Availability`: high availability
	//
	// - `AvailabilityGroup`: high availability for deployment set groups
	//
	// - `LowLatency`: low latency
	//
	// - `ProximityLooseDispersion`: proximity loose dispersion
	//
	// > Only the `ProximityLooseDispersion` strategy incurs a fee. The API response includes price details for the deployment set (where `Resource` is `deploymentSet`) only when this strategy is used. Other deployment set strategies are free of charge.
	//
	// example:
	//
	// ProximityLooseDispersion
	DeploymentSetStrategy *string `json:"DeploymentSetStrategy,omitempty" xml:"DeploymentSetStrategy,omitempty"`
}

func (s DescribePriceRequestSchedulerOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceRequestSchedulerOptions) GoString() string {
	return s.String()
}

func (s *DescribePriceRequestSchedulerOptions) GetDedicatedHostId() *string {
	return s.DedicatedHostId
}

func (s *DescribePriceRequestSchedulerOptions) GetDeploymentSetStrategy() *string {
	return s.DeploymentSetStrategy
}

func (s *DescribePriceRequestSchedulerOptions) SetDedicatedHostId(v string) *DescribePriceRequestSchedulerOptions {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribePriceRequestSchedulerOptions) SetDeploymentSetStrategy(v string) *DescribePriceRequestSchedulerOptions {
	s.DeploymentSetStrategy = &v
	return s
}

func (s *DescribePriceRequestSchedulerOptions) Validate() error {
	return dara.Validate(s)
}

type DescribePriceRequestSystemDisk struct {
	// The category of the system disk. You must specify `ImageId` when you query the price of a system disk. Valid values:
	//
	// - `cloud`: basic cloud disk
	//
	// - `cloud_efficiency`: efficiency cloud disk
	//
	// - `cloud_ssd`: SSD cloud disk
	//
	// - `ephemeral_ssd`: local SSD
	//
	// - `cloud_essd`: ESSD
	//
	// - `cloud_auto`: ESSD AutoPL
	//
	// <props="china">
	//
	// - `cloud_essd_entry`: ESSD Entry
	//
	//
	//
	//
	// 	- For retired instance types where `IoOptimized` is `none`, the default value is `cloud`.
	//
	// 	- In other cases, the default value is `cloud_efficiency`.<props="china">After January 30, 2026, for instance types that support only ESSDs, the default value will be changed from `cloud_efficiency` to `cloud_essd` at PL0. For more information, see the [change announcement](https://www.aliyun.com/notice/117844).
	//
	// example:
	//
	// cloud_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The performance level of the ESSD when used as a system disk. This parameter is valid only when `SystemDisk.Category` is set to `cloud_essd`. Valid values:
	//
	// `PL0`<br>`PL1` (Default)<br>`PL2`<br>`PL3`<br><br><br>
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The size of the system disk. Unit: GiB. Valid values:
	//
	// - Basic cloud disk: 20–500.
	//
	// - ESSD cloud disk:
	//
	//   - PL0: 1–2048.
	//
	//   - PL1: 20–2048.
	//
	//   - PL2: 461–2048.
	//
	//   - PL3: 1261–2048.
	//
	// - ESSD AutoPL cloud disk: 1–2048.
	//
	// - Other cloud disk categories: 20–2048.
	//
	// Default value: `max{20, ImageSize}`, which is the greater of 20 and the size of the specified image (`ImageId`).
	//
	// example:
	//
	// 80
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribePriceRequestSystemDisk) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *DescribePriceRequestSystemDisk) GetCategory() *string {
	return s.Category
}

func (s *DescribePriceRequestSystemDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *DescribePriceRequestSystemDisk) GetSize() *int32 {
	return s.Size
}

func (s *DescribePriceRequestSystemDisk) SetCategory(v string) *DescribePriceRequestSystemDisk {
	s.Category = &v
	return s
}

func (s *DescribePriceRequestSystemDisk) SetPerformanceLevel(v string) *DescribePriceRequestSystemDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribePriceRequestSystemDisk) SetSize(v int32) *DescribePriceRequestSystemDisk {
	s.Size = &v
	return s
}

func (s *DescribePriceRequestSystemDisk) Validate() error {
	return dara.Validate(s)
}

type DescribePriceRequestRecurrenceRules struct {
	// The end time of the time-based assurance. The value must be on the hour.
	//
	// example:
	//
	// 10
	EndHour *int32 `json:"EndHour,omitempty" xml:"EndHour,omitempty"`
	// The recurrence type of the rule. Valid values:
	//
	// - `Daily`: repeats on a daily basis.
	//
	// - `Weekly`: repeats on a weekly basis.
	//
	// - `Monthly`: repeats on a monthly basis.
	//
	// > You must specify both `RecurrenceType` and `RecurrenceValue`.
	//
	// example:
	//
	// Daily
	RecurrenceType *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	// The recurrence value.
	//
	// - If `RecurrenceType` is set to `Daily`, this parameter takes a single value that specifies the recurrence interval in days. Valid values: 1–31.
	//
	// - If `RecurrenceType` is set to `Weekly`, this parameter can have multiple values separated by commas (,). The values 0, 1, 2, 3, 4, 5, and 6 correspond to Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, and Saturday. For example, `1,2` specifies Monday and Tuesday.
	//
	// - If `RecurrenceType` is set to `Monthly`, the value must be in the `A–B` format. The values of A and B must be between 1 and 31, and B must be greater than or equal to A. For example, `1–5` specifies the first to the fifth day of each month.
	//
	// > You must specify both `RecurrenceType` and `RecurrenceValue`.
	//
	// example:
	//
	// 5
	RecurrenceValue *string `json:"RecurrenceValue,omitempty" xml:"RecurrenceValue,omitempty"`
	// The start time of the time-based assurance. The value must be on the hour.
	//
	// > Both `StartHour` and `EndHour` are required. The interval between them must be at least 4 hours.
	//
	// example:
	//
	// 4
	StartHour *int32 `json:"StartHour,omitempty" xml:"StartHour,omitempty"`
}

func (s DescribePriceRequestRecurrenceRules) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceRequestRecurrenceRules) GoString() string {
	return s.String()
}

func (s *DescribePriceRequestRecurrenceRules) GetEndHour() *int32 {
	return s.EndHour
}

func (s *DescribePriceRequestRecurrenceRules) GetRecurrenceType() *string {
	return s.RecurrenceType
}

func (s *DescribePriceRequestRecurrenceRules) GetRecurrenceValue() *string {
	return s.RecurrenceValue
}

func (s *DescribePriceRequestRecurrenceRules) GetStartHour() *int32 {
	return s.StartHour
}

func (s *DescribePriceRequestRecurrenceRules) SetEndHour(v int32) *DescribePriceRequestRecurrenceRules {
	s.EndHour = &v
	return s
}

func (s *DescribePriceRequestRecurrenceRules) SetRecurrenceType(v string) *DescribePriceRequestRecurrenceRules {
	s.RecurrenceType = &v
	return s
}

func (s *DescribePriceRequestRecurrenceRules) SetRecurrenceValue(v string) *DescribePriceRequestRecurrenceRules {
	s.RecurrenceValue = &v
	return s
}

func (s *DescribePriceRequestRecurrenceRules) SetStartHour(v int32) *DescribePriceRequestRecurrenceRules {
	s.StartHour = &v
	return s
}

func (s *DescribePriceRequestRecurrenceRules) Validate() error {
	return dara.Validate(s)
}
