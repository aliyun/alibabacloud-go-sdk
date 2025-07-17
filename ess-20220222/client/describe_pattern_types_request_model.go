// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePatternTypesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArchitecture(v []*string) *DescribePatternTypesRequest
	GetArchitecture() []*string
	SetBurstablePerformance(v string) *DescribePatternTypesRequest
	GetBurstablePerformance() *string
	SetChannelId(v int64) *DescribePatternTypesRequest
	GetChannelId() *int64
	SetCores(v int32) *DescribePatternTypesRequest
	GetCores() *int32
	SetCoresList(v []*int32) *DescribePatternTypesRequest
	GetCoresList() []*int32
	SetCpuArchitectures(v []*string) *DescribePatternTypesRequest
	GetCpuArchitectures() []*string
	SetExcludedInstanceType(v []*string) *DescribePatternTypesRequest
	GetExcludedInstanceType() []*string
	SetGpuSpecs(v []*string) *DescribePatternTypesRequest
	GetGpuSpecs() []*string
	SetInstanceCategories(v []*string) *DescribePatternTypesRequest
	GetInstanceCategories() []*string
	SetInstanceFamilyLevel(v string) *DescribePatternTypesRequest
	GetInstanceFamilyLevel() *string
	SetInstanceTypeFamilies(v []*string) *DescribePatternTypesRequest
	GetInstanceTypeFamilies() []*string
	SetMaxPrice(v float32) *DescribePatternTypesRequest
	GetMaxPrice() *float32
	SetMaximumCpuCoreCount(v int32) *DescribePatternTypesRequest
	GetMaximumCpuCoreCount() *int32
	SetMaximumGpuAmount(v int32) *DescribePatternTypesRequest
	GetMaximumGpuAmount() *int32
	SetMaximumMemorySize(v float32) *DescribePatternTypesRequest
	GetMaximumMemorySize() *float32
	SetMemory(v float32) *DescribePatternTypesRequest
	GetMemory() *float32
	SetMemoryList(v []*float32) *DescribePatternTypesRequest
	GetMemoryList() []*float32
	SetMinimumBaselineCredit(v int32) *DescribePatternTypesRequest
	GetMinimumBaselineCredit() *int32
	SetMinimumCpuCoreCount(v int32) *DescribePatternTypesRequest
	GetMinimumCpuCoreCount() *int32
	SetMinimumEniIpv6AddressQuantity(v int32) *DescribePatternTypesRequest
	GetMinimumEniIpv6AddressQuantity() *int32
	SetMinimumEniPrivateIpAddressQuantity(v int32) *DescribePatternTypesRequest
	GetMinimumEniPrivateIpAddressQuantity() *int32
	SetMinimumEniQuantity(v int32) *DescribePatternTypesRequest
	GetMinimumEniQuantity() *int32
	SetMinimumGpuAmount(v int32) *DescribePatternTypesRequest
	GetMinimumGpuAmount() *int32
	SetMinimumInitialCredit(v int32) *DescribePatternTypesRequest
	GetMinimumInitialCredit() *int32
	SetMinimumMemorySize(v float32) *DescribePatternTypesRequest
	GetMinimumMemorySize() *float32
	SetPhysicalProcessorModels(v []*string) *DescribePatternTypesRequest
	GetPhysicalProcessorModels() []*string
	SetRegionId(v string) *DescribePatternTypesRequest
	GetRegionId() *string
	SetSpotStrategy(v string) *DescribePatternTypesRequest
	GetSpotStrategy() *string
	SetVSwitchId(v []*string) *DescribePatternTypesRequest
	GetVSwitchId() []*string
	SetZoneId(v []*string) *DescribePatternTypesRequest
	GetZoneId() []*string
}

type DescribePatternTypesRequest struct {
	// The architecture types of the instance types. Valid values:
	//
	// 	- X86: x86 architecture.
	//
	// 	- Heterogeneous: heterogeneous computing, such as GPU-accelerated or FPGA-accelerated.
	//
	// 	- BareMetal: ECS Bare Metal Instance.
	//
	// 	- Arm: Arm.
	//
	// By default, all values are selected.
	Architecture []*string `json:"Architecture,omitempty" xml:"Architecture,omitempty" type:"Repeated"`
	// Specifies whether to include burstable instance types. Valid values:
	//
	// 	- Exclude: does not include burstable instance types.
	//
	// 	- Include: includes burstable instance types.
	//
	// 	- Required: includes only burstable instance types.
	//
	// Default value: Include.
	//
	// example:
	//
	// Include
	BurstablePerformance *string `json:"BurstablePerformance,omitempty" xml:"BurstablePerformance,omitempty"`
	// The channel ID. This parameter is not for public use.
	//
	// example:
	//
	// 79425074
	ChannelId *int64 `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The number of vCPUs that you want to assign to the instance type.
	//
	// example:
	//
	// 2
	Cores *int32 `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// The number of vCPUs that you want to assign to the instance type. You can specify multiple vCPUs.
	CoresList []*int32 `json:"CoresList,omitempty" xml:"CoresList,omitempty" type:"Repeated"`
	// The CPU architectures of the instance types. Valid values:
	//
	// >  You can specify 1 to 2 CPU architectures.
	//
	// 	- x86
	//
	// 	- Arm
	CpuArchitectures     []*string `json:"CpuArchitectures,omitempty" xml:"CpuArchitectures,omitempty" type:"Repeated"`
	ExcludedInstanceType []*string `json:"ExcludedInstanceType,omitempty" xml:"ExcludedInstanceType,omitempty" type:"Repeated"`
	// The GPU models.
	GpuSpecs []*string `json:"GpuSpecs,omitempty" xml:"GpuSpecs,omitempty" type:"Repeated"`
	// The classifications of the instance types. Valid values:
	//
	// 	- General-purpose: general-purpose instance type.
	//
	// 	- Compute-optimized: compute-optimized instance type.
	//
	// 	- Memory-optimized: memory-optimized instance type.
	//
	// 	- Big data: big data instance type.
	//
	// 	- Local SSDs: instance type with local SSDs.
	//
	// 	- High Clock Speed: instance type with high clock speeds.
	//
	// 	- Enhanced: enhanced instance type.
	//
	// 	- Shared: shared instance type.
	//
	// 	- Compute-optimized with GPU: GPU-accelerated compute-optimized instance type.
	//
	// 	- Visual Compute-optimized: visual compute-optimized instance type.
	//
	// 	- Heterogeneous Service: heterogeneous service instance type.
	//
	// 	- Compute-optimized with FPGA: FPGA-accelerated compute-optimized instance type.
	//
	// 	- Compute-optimized with NPU: NPU-accelerated compute-optimized instance type.
	//
	// 	- ECS Bare Metal: ECS Bare Metal Instance type.
	//
	// 	- High Performance Compute: HPC-optimized instance type.
	InstanceCategories []*string `json:"InstanceCategories,omitempty" xml:"InstanceCategories,omitempty" type:"Repeated"`
	// The level of the instance family. Valid values:
	//
	// 	- EntryLevel: entry level
	//
	// 	- EnterpriseLevel: enterprise level
	//
	// 	- CreditEntryLevel: credit-based entry level For more information, see [Burstable instance families](https://help.aliyun.com/document_detail/59977.html).
	//
	// example:
	//
	// EnterpriseLevel
	InstanceFamilyLevel *string `json:"InstanceFamilyLevel,omitempty" xml:"InstanceFamilyLevel,omitempty"`
	// The instance families that you want to query. You can query 1 to 10 instance families in each call.
	InstanceTypeFamilies []*string `json:"InstanceTypeFamilies,omitempty" xml:"InstanceTypeFamilies,omitempty" type:"Repeated"`
	// The maximum hourly price for pay-as-you-go or preemptible instances.
	//
	// example:
	//
	// 2
	MaxPrice *float32 `json:"MaxPrice,omitempty" xml:"MaxPrice,omitempty"`
	// The maximum number of vCPUs per instance type.
	//
	// example:
	//
	// 4
	MaximumCpuCoreCount *int32 `json:"MaximumCpuCoreCount,omitempty" xml:"MaximumCpuCoreCount,omitempty"`
	// The maximum number of GPUs per instance. The value must be a positive integer.
	//
	// example:
	//
	// 2
	MaximumGpuAmount *int32 `json:"MaximumGpuAmount,omitempty" xml:"MaximumGpuAmount,omitempty"`
	// The maximum memory size per instance. Unit: GiB.
	//
	// example:
	//
	// 4
	MaximumMemorySize *float32 `json:"MaximumMemorySize,omitempty" xml:"MaximumMemorySize,omitempty"`
	// The memory size that you want to assign to the instance type. Unit: GiB.
	//
	// example:
	//
	// 4
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The memory size that you want to assign to the instance type. Unit: GiB. You can specify multiple memory sizes.
	MemoryList []*float32 `json:"MemoryList,omitempty" xml:"MemoryList,omitempty" type:"Repeated"`
	// The baseline vCPU computing performance (overall baseline performance of all vCPUs) per t5 or t6 burstable instance.
	//
	// example:
	//
	// 12
	MinimumBaselineCredit *int32 `json:"MinimumBaselineCredit,omitempty" xml:"MinimumBaselineCredit,omitempty"`
	// The minimum number of vCPUs per instance type.
	//
	// example:
	//
	// 2
	MinimumCpuCoreCount *int32 `json:"MinimumCpuCoreCount,omitempty" xml:"MinimumCpuCoreCount,omitempty"`
	// The minimum number of IPv6 addresses per ENI.
	//
	// example:
	//
	// 1
	MinimumEniIpv6AddressQuantity *int32 `json:"MinimumEniIpv6AddressQuantity,omitempty" xml:"MinimumEniIpv6AddressQuantity,omitempty"`
	// The minimum number of IPv4 addresses per ENI.
	//
	// example:
	//
	// 2
	MinimumEniPrivateIpAddressQuantity *int32 `json:"MinimumEniPrivateIpAddressQuantity,omitempty" xml:"MinimumEniPrivateIpAddressQuantity,omitempty"`
	// The minimum number of elastic network interfaces (ENIs) per instance.
	//
	// example:
	//
	// 2
	MinimumEniQuantity *int32 `json:"MinimumEniQuantity,omitempty" xml:"MinimumEniQuantity,omitempty"`
	// The minimum number of GPUs per instance. The value must be a positive integer.
	//
	// example:
	//
	// 2
	MinimumGpuAmount *int32 `json:"MinimumGpuAmount,omitempty" xml:"MinimumGpuAmount,omitempty"`
	// The initial vCPU credits per t5 or t6 burstable instance.
	//
	// example:
	//
	// 12
	MinimumInitialCredit *int32 `json:"MinimumInitialCredit,omitempty" xml:"MinimumInitialCredit,omitempty"`
	// The minimum memory size per instance. Unit: GiB.
	//
	// example:
	//
	// 4
	MinimumMemorySize *float32 `json:"MinimumMemorySize,omitempty" xml:"MinimumMemorySize,omitempty"`
	// The processor models of the instance types. You can specify 1 to 10 processor models.
	PhysicalProcessorModels []*string `json:"PhysicalProcessorModels,omitempty" xml:"PhysicalProcessorModels,omitempty" type:"Repeated"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The preemption policy that you want to apply to pay-as-you-go instances. Valid values:
	//
	// 	- NoSpot: The instances are created as regular pay-as-you-go instances.
	//
	// 	- SpotWithPriceLimit: The instances are created as preemptible instances that have a user-defined maximum hourly price.
	//
	// 	- SpotAsPriceGo: The instances are created as preemptible instances for which the market price at the time of purchase is automatically used as the bidding price.
	//
	// Default value: NoSpot.
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The IDs of the vSwitches.
	VSwitchId []*string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" type:"Repeated"`
	// The zone IDs. If you pass vSwitch IDs to the system, this parameter does not take effect.
	ZoneId []*string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty" type:"Repeated"`
}

func (s DescribePatternTypesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePatternTypesRequest) GoString() string {
	return s.String()
}

func (s *DescribePatternTypesRequest) GetArchitecture() []*string {
	return s.Architecture
}

func (s *DescribePatternTypesRequest) GetBurstablePerformance() *string {
	return s.BurstablePerformance
}

func (s *DescribePatternTypesRequest) GetChannelId() *int64 {
	return s.ChannelId
}

func (s *DescribePatternTypesRequest) GetCores() *int32 {
	return s.Cores
}

func (s *DescribePatternTypesRequest) GetCoresList() []*int32 {
	return s.CoresList
}

func (s *DescribePatternTypesRequest) GetCpuArchitectures() []*string {
	return s.CpuArchitectures
}

func (s *DescribePatternTypesRequest) GetExcludedInstanceType() []*string {
	return s.ExcludedInstanceType
}

func (s *DescribePatternTypesRequest) GetGpuSpecs() []*string {
	return s.GpuSpecs
}

func (s *DescribePatternTypesRequest) GetInstanceCategories() []*string {
	return s.InstanceCategories
}

func (s *DescribePatternTypesRequest) GetInstanceFamilyLevel() *string {
	return s.InstanceFamilyLevel
}

func (s *DescribePatternTypesRequest) GetInstanceTypeFamilies() []*string {
	return s.InstanceTypeFamilies
}

func (s *DescribePatternTypesRequest) GetMaxPrice() *float32 {
	return s.MaxPrice
}

func (s *DescribePatternTypesRequest) GetMaximumCpuCoreCount() *int32 {
	return s.MaximumCpuCoreCount
}

func (s *DescribePatternTypesRequest) GetMaximumGpuAmount() *int32 {
	return s.MaximumGpuAmount
}

func (s *DescribePatternTypesRequest) GetMaximumMemorySize() *float32 {
	return s.MaximumMemorySize
}

func (s *DescribePatternTypesRequest) GetMemory() *float32 {
	return s.Memory
}

func (s *DescribePatternTypesRequest) GetMemoryList() []*float32 {
	return s.MemoryList
}

func (s *DescribePatternTypesRequest) GetMinimumBaselineCredit() *int32 {
	return s.MinimumBaselineCredit
}

func (s *DescribePatternTypesRequest) GetMinimumCpuCoreCount() *int32 {
	return s.MinimumCpuCoreCount
}

func (s *DescribePatternTypesRequest) GetMinimumEniIpv6AddressQuantity() *int32 {
	return s.MinimumEniIpv6AddressQuantity
}

func (s *DescribePatternTypesRequest) GetMinimumEniPrivateIpAddressQuantity() *int32 {
	return s.MinimumEniPrivateIpAddressQuantity
}

func (s *DescribePatternTypesRequest) GetMinimumEniQuantity() *int32 {
	return s.MinimumEniQuantity
}

func (s *DescribePatternTypesRequest) GetMinimumGpuAmount() *int32 {
	return s.MinimumGpuAmount
}

func (s *DescribePatternTypesRequest) GetMinimumInitialCredit() *int32 {
	return s.MinimumInitialCredit
}

func (s *DescribePatternTypesRequest) GetMinimumMemorySize() *float32 {
	return s.MinimumMemorySize
}

func (s *DescribePatternTypesRequest) GetPhysicalProcessorModels() []*string {
	return s.PhysicalProcessorModels
}

func (s *DescribePatternTypesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePatternTypesRequest) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *DescribePatternTypesRequest) GetVSwitchId() []*string {
	return s.VSwitchId
}

func (s *DescribePatternTypesRequest) GetZoneId() []*string {
	return s.ZoneId
}

func (s *DescribePatternTypesRequest) SetArchitecture(v []*string) *DescribePatternTypesRequest {
	s.Architecture = v
	return s
}

func (s *DescribePatternTypesRequest) SetBurstablePerformance(v string) *DescribePatternTypesRequest {
	s.BurstablePerformance = &v
	return s
}

func (s *DescribePatternTypesRequest) SetChannelId(v int64) *DescribePatternTypesRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribePatternTypesRequest) SetCores(v int32) *DescribePatternTypesRequest {
	s.Cores = &v
	return s
}

func (s *DescribePatternTypesRequest) SetCoresList(v []*int32) *DescribePatternTypesRequest {
	s.CoresList = v
	return s
}

func (s *DescribePatternTypesRequest) SetCpuArchitectures(v []*string) *DescribePatternTypesRequest {
	s.CpuArchitectures = v
	return s
}

func (s *DescribePatternTypesRequest) SetExcludedInstanceType(v []*string) *DescribePatternTypesRequest {
	s.ExcludedInstanceType = v
	return s
}

func (s *DescribePatternTypesRequest) SetGpuSpecs(v []*string) *DescribePatternTypesRequest {
	s.GpuSpecs = v
	return s
}

func (s *DescribePatternTypesRequest) SetInstanceCategories(v []*string) *DescribePatternTypesRequest {
	s.InstanceCategories = v
	return s
}

func (s *DescribePatternTypesRequest) SetInstanceFamilyLevel(v string) *DescribePatternTypesRequest {
	s.InstanceFamilyLevel = &v
	return s
}

func (s *DescribePatternTypesRequest) SetInstanceTypeFamilies(v []*string) *DescribePatternTypesRequest {
	s.InstanceTypeFamilies = v
	return s
}

func (s *DescribePatternTypesRequest) SetMaxPrice(v float32) *DescribePatternTypesRequest {
	s.MaxPrice = &v
	return s
}

func (s *DescribePatternTypesRequest) SetMaximumCpuCoreCount(v int32) *DescribePatternTypesRequest {
	s.MaximumCpuCoreCount = &v
	return s
}

func (s *DescribePatternTypesRequest) SetMaximumGpuAmount(v int32) *DescribePatternTypesRequest {
	s.MaximumGpuAmount = &v
	return s
}

func (s *DescribePatternTypesRequest) SetMaximumMemorySize(v float32) *DescribePatternTypesRequest {
	s.MaximumMemorySize = &v
	return s
}

func (s *DescribePatternTypesRequest) SetMemory(v float32) *DescribePatternTypesRequest {
	s.Memory = &v
	return s
}

func (s *DescribePatternTypesRequest) SetMemoryList(v []*float32) *DescribePatternTypesRequest {
	s.MemoryList = v
	return s
}

func (s *DescribePatternTypesRequest) SetMinimumBaselineCredit(v int32) *DescribePatternTypesRequest {
	s.MinimumBaselineCredit = &v
	return s
}

func (s *DescribePatternTypesRequest) SetMinimumCpuCoreCount(v int32) *DescribePatternTypesRequest {
	s.MinimumCpuCoreCount = &v
	return s
}

func (s *DescribePatternTypesRequest) SetMinimumEniIpv6AddressQuantity(v int32) *DescribePatternTypesRequest {
	s.MinimumEniIpv6AddressQuantity = &v
	return s
}

func (s *DescribePatternTypesRequest) SetMinimumEniPrivateIpAddressQuantity(v int32) *DescribePatternTypesRequest {
	s.MinimumEniPrivateIpAddressQuantity = &v
	return s
}

func (s *DescribePatternTypesRequest) SetMinimumEniQuantity(v int32) *DescribePatternTypesRequest {
	s.MinimumEniQuantity = &v
	return s
}

func (s *DescribePatternTypesRequest) SetMinimumGpuAmount(v int32) *DescribePatternTypesRequest {
	s.MinimumGpuAmount = &v
	return s
}

func (s *DescribePatternTypesRequest) SetMinimumInitialCredit(v int32) *DescribePatternTypesRequest {
	s.MinimumInitialCredit = &v
	return s
}

func (s *DescribePatternTypesRequest) SetMinimumMemorySize(v float32) *DescribePatternTypesRequest {
	s.MinimumMemorySize = &v
	return s
}

func (s *DescribePatternTypesRequest) SetPhysicalProcessorModels(v []*string) *DescribePatternTypesRequest {
	s.PhysicalProcessorModels = v
	return s
}

func (s *DescribePatternTypesRequest) SetRegionId(v string) *DescribePatternTypesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePatternTypesRequest) SetSpotStrategy(v string) *DescribePatternTypesRequest {
	s.SpotStrategy = &v
	return s
}

func (s *DescribePatternTypesRequest) SetVSwitchId(v []*string) *DescribePatternTypesRequest {
	s.VSwitchId = v
	return s
}

func (s *DescribePatternTypesRequest) SetZoneId(v []*string) *DescribePatternTypesRequest {
	s.ZoneId = v
	return s
}

func (s *DescribePatternTypesRequest) Validate() error {
	return dara.Validate(s)
}
