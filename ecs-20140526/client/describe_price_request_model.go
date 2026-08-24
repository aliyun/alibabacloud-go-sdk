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
	// The number of Elastic Computing Service (ECS) servers that you want to purchase in a batch. Valid values: 1 to 1000.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The total number of times that the elasticity assurance can take effect. Set the value to Unlimited. Only the unlimited mode is supported within the effective period.
	//
	// example:
	//
	// Unlimited
	AssuranceTimes *string `json:"AssuranceTimes,omitempty" xml:"AssuranceTimes,omitempty"`
	// The capacity. Unit: GiB.
	//
	// example:
	//
	// 1024
	Capacity *int32 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The dedicated host type. You can call [DescribeDedicatedHostTypes](https://help.aliyun.com/document_detail/134240.html) to query the most recent list of dedicated host types.
	//
	// example:
	//
	// ddh.c5
	DedicatedHostType *string `json:"DedicatedHostType,omitempty" xml:"DedicatedHostType,omitempty"`
	// This parameter takes effect only when ResourceType is set to instance.
	//
	// example:
	//
	// centos_7_05_64_20G_alibase_20181212.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The total number of instances that you want to reserve within an instance type.
	//
	// example:
	//
	// 100
	InstanceAmount *int32 `json:"InstanceAmount,omitempty" xml:"InstanceAmount,omitempty"`
	// The total number of vCPUs supported by the elasticity assurance. When you call this operation, the system calculates the number of instances that the elasticity assurance needs to support based on the specified InstanceType (rounded up).
	//
	// example:
	//
	// 1024
	InstanceCpuCoreCount *int32 `json:"InstanceCpuCoreCount,omitempty" xml:"InstanceCpuCoreCount,omitempty"`
	// The network type of the instance. Valid values:
	//
	// example:
	//
	// vpc
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	// The instance type. You must specify this parameter when ResourceType is set to `instance`. For more details, see [Instance family](https://help.aliyun.com/document_detail/25378.html). You can also invoke [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) to query the most recent instance type list.
	//
	// example:
	//
	// ecs.g6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The instance type. Only a single instance type is supported for the unlimited elasticity assurance.
	//
	// example:
	//
	// ecs.g6.xlarge
	InstanceTypeList []*string `json:"InstanceTypeList,omitempty" xml:"InstanceTypeList,omitempty" type:"Repeated"`
	// The billing method for network bandwidth. Valid values:
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The maximum outbound public bandwidth. Unit: Mbit/s. Valid values: 0 to 100.
	//
	// Default value: 0.
	//
	// example:
	//
	// 5
	InternetMaxBandwidthOut *int32 `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// Specifies whether the queried instance is an I/O optimized instance. Valid values:
	//
	// - none: non-I/O optimized.
	//
	// - optimized: I/O optimized.
	//
	// If InstanceType is set to an instance type in [Series I](https://help.aliyun.com/document_detail/55263.html), the default value is none.
	//
	// If InstanceType is set to an instance type not in [Series I](https://help.aliyun.com/document_detail/55263.html), the default value is optimized.
	//
	// example:
	//
	// optimized
	IoOptimized *string `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	// The Internet service operation provider. Valid values:
	//
	// example:
	//
	// cmcc
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The payment option of the reserved instance. Valid values:
	//
	// example:
	//
	// All Upfront
	OfferingType *string `json:"OfferingType,omitempty" xml:"OfferingType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing duration of the ECS instance. Valid values:
	//
	// <props="china">
	//
	// - If PriceUnit is set to Month: 1 to 9.
	//
	// - If PriceUnit is set to Year: 1 to 5.
	//
	// - If PriceUnit is set to Hour: 1.
	//
	// - If PriceUnit is set to Week: 1 to 4.
	//
	//
	//
	// <props="intl">
	//
	// - If PriceUnit is set to Month: 1 to 9.
	//
	// - If PriceUnit is set to Year: 1 to 5.
	//
	// - If PriceUnit is set to Hour: 1.
	//
	//
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The operating system of the image used by the instance. Valid values:
	//
	// example:
	//
	// Linux
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// Queries the price of an ECS instance for different billing cycles. Valid values:
	//
	// <props="china">
	//
	// - Month: monthly pricing unit.
	//
	// - Year: yearly pricing unit.
	//
	// - Hour (default): hourly pricing unit.
	//
	// - Week: weekly pricing unit.
	//
	//
	//
	// <props="intl">
	//
	// - Month: monthly pricing unit.
	//
	// - Year: yearly pricing unit.
	//
	// - Hour (default): hourly pricing unit.
	//
	// example:
	//
	// Year
	PriceUnit *string `json:"PriceUnit,omitempty" xml:"PriceUnit,omitempty"`
	// The list of recurrence rules for the time-sharing elasticity assurance.
	RecurrenceRules []*DescribePriceRequestRecurrenceRules `json:"RecurrenceRules,omitempty" xml:"RecurrenceRules,omitempty" type:"Repeated"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the resource. Valid values:
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The scope of the reserved instance. Valid values:
	//
	// example:
	//
	// Zone
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The protection period of the spot instance. Unit: hours. Default value: 1. Valid values:
	//
	// example:
	//
	// 1
	SpotDuration *int32 `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	// The bidding policy for pay-as-you-go instances. Valid values:
	//
	// - NoSpot: a regular pay-as-you-go instance.
	//
	// - SpotWithPriceLimit: a spot instance with a maximum hourly price.
	//
	// - SpotAsPriceGo: a spot instance for which the system automatically bids at up to the pay-as-you-go price.
	//
	// Default value: NoSpot.
	//
	// > This parameter takes effect only when `PriceUnit=Hour` and `Period=1`. Because the default value of `PriceUnit` is `Hour` and the default value of `Period` is `1`, you do not need to set the `PriceUnit` and `Period` parameters when you specify this parameter.
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The effective period of the time-sharing elasticity assurance. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC+0. For more information, see [ISO 8601](https://help.aliyun.com/document_detail/25696.html).
	//
	// example:
	//
	// 2020-10-30T06:32:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The zone ID.
	//
	// > Spot instance prices may vary across zones. When you query spot instance prices, specify ZoneId to query the price in a specific zone.
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
	// example:
	//
	// cloud_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The performance level of the Nth data disk when the disk type is enterprise SSD. This parameter is valid only when `DataDisk.N.Category=cloud_essd`. Valid values:
	//
	// - PL0
	//
	// - PL1 (default)
	//
	// - PL2
	//
	// - PL3
	//
	// Valid values of N: 1 to 16.
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The size of data disk N. Unit: GiB. Valid values:
	//
	// example:
	//
	// 2000
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The provisioned read/write IOPS of the ESSD AutoPL disk. Valid values: 0 to min{50,000, 1000 × Capacity - Baseline Performance}.
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
	// This parameter takes effect only when ResourceType is set to instance.
	//
	// example:
	//
	// dh-bp67acfmxazb4p****
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	// The deployment set strategy. Valid values:
	//
	// - Availability: high availability.
	//
	// - AvailabilityGroup: high availability for deployment set groups.
	//
	// - LowLatency: low network latency.
	//
	// - ProximityLooseDispersion: proximity loose dispersion.
	//
	// >Only when the strategy is set to ProximityLooseDispersion, the API response includes the price details for "Resource": "deploymentSet". Other deployment set strategies are free of charge, so the API response does not include the price information for "Resource": "deploymentSet".
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
	// example:
	//
	// cloud_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The performance level of the system disk when the system disk type is enterprise SSD. This parameter is valid only when `SystemDiskCategory=cloud_essd`. Valid values:
	//
	// PL0.
	//
	// PL1 (default).
	//
	// PL2.
	//
	// PL3.
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The size of the system disk. Unit: GiB. Valid values:
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
	// The end time of the time-sharing assurance. The value must be on the hour.
	//
	// example:
	//
	// 10
	EndHour *int32 `json:"EndHour,omitempty" xml:"EndHour,omitempty"`
	// The type of the recurrence rule policy. Valid values:
	//
	// example:
	//
	// Daily
	RecurrenceType *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	// The value of the recurrence rule.
	//
	// example:
	//
	// 5
	RecurrenceValue *string `json:"RecurrenceValue,omitempty" xml:"RecurrenceValue,omitempty"`
	// The effective period start time of the time-sharing assurance. The value must be on the hour.
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
