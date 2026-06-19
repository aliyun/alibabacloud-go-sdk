// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceTypesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalAttributes(v []*string) *DescribeInstanceTypesRequest
	GetAdditionalAttributes() []*string
	SetCpuArchitecture(v string) *DescribeInstanceTypesRequest
	GetCpuArchitecture() *string
	SetCpuArchitectures(v []*string) *DescribeInstanceTypesRequest
	GetCpuArchitectures() []*string
	SetGPUSpec(v string) *DescribeInstanceTypesRequest
	GetGPUSpec() *string
	SetGpuSpecs(v []*string) *DescribeInstanceTypesRequest
	GetGpuSpecs() []*string
	SetInstanceCategories(v []*string) *DescribeInstanceTypesRequest
	GetInstanceCategories() []*string
	SetInstanceCategory(v string) *DescribeInstanceTypesRequest
	GetInstanceCategory() *string
	SetInstanceFamilyLevel(v string) *DescribeInstanceTypesRequest
	GetInstanceFamilyLevel() *string
	SetInstanceTypeFamilies(v []*string) *DescribeInstanceTypesRequest
	GetInstanceTypeFamilies() []*string
	SetInstanceTypeFamily(v string) *DescribeInstanceTypesRequest
	GetInstanceTypeFamily() *string
	SetInstanceTypes(v []*string) *DescribeInstanceTypesRequest
	GetInstanceTypes() []*string
	SetLocalStorageCategories(v []*string) *DescribeInstanceTypesRequest
	GetLocalStorageCategories() []*string
	SetLocalStorageCategory(v string) *DescribeInstanceTypesRequest
	GetLocalStorageCategory() *string
	SetMaxResults(v int64) *DescribeInstanceTypesRequest
	GetMaxResults() *int64
	SetMaximumCpuCoreCount(v int32) *DescribeInstanceTypesRequest
	GetMaximumCpuCoreCount() *int32
	SetMaximumCpuSpeedFrequency(v float32) *DescribeInstanceTypesRequest
	GetMaximumCpuSpeedFrequency() *float32
	SetMaximumCpuTurboFrequency(v float32) *DescribeInstanceTypesRequest
	GetMaximumCpuTurboFrequency() *float32
	SetMaximumGPUAmount(v int32) *DescribeInstanceTypesRequest
	GetMaximumGPUAmount() *int32
	SetMaximumMemorySize(v float32) *DescribeInstanceTypesRequest
	GetMaximumMemorySize() *float32
	SetMinimumBaselineCredit(v int32) *DescribeInstanceTypesRequest
	GetMinimumBaselineCredit() *int32
	SetMinimumCpuCoreCount(v int32) *DescribeInstanceTypesRequest
	GetMinimumCpuCoreCount() *int32
	SetMinimumCpuSpeedFrequency(v float32) *DescribeInstanceTypesRequest
	GetMinimumCpuSpeedFrequency() *float32
	SetMinimumCpuTurboFrequency(v float32) *DescribeInstanceTypesRequest
	GetMinimumCpuTurboFrequency() *float32
	SetMinimumDiskQuantity(v int32) *DescribeInstanceTypesRequest
	GetMinimumDiskQuantity() *int32
	SetMinimumEniIpv6AddressQuantity(v int32) *DescribeInstanceTypesRequest
	GetMinimumEniIpv6AddressQuantity() *int32
	SetMinimumEniPrivateIpAddressQuantity(v int32) *DescribeInstanceTypesRequest
	GetMinimumEniPrivateIpAddressQuantity() *int32
	SetMinimumEniQuantity(v int32) *DescribeInstanceTypesRequest
	GetMinimumEniQuantity() *int32
	SetMinimumEriQuantity(v int32) *DescribeInstanceTypesRequest
	GetMinimumEriQuantity() *int32
	SetMinimumGPUAmount(v int32) *DescribeInstanceTypesRequest
	GetMinimumGPUAmount() *int32
	SetMinimumInitialCredit(v int32) *DescribeInstanceTypesRequest
	GetMinimumInitialCredit() *int32
	SetMinimumInstanceBandwidthRx(v int32) *DescribeInstanceTypesRequest
	GetMinimumInstanceBandwidthRx() *int32
	SetMinimumInstanceBandwidthTx(v int32) *DescribeInstanceTypesRequest
	GetMinimumInstanceBandwidthTx() *int32
	SetMinimumInstancePpsRx(v int64) *DescribeInstanceTypesRequest
	GetMinimumInstancePpsRx() *int64
	SetMinimumInstancePpsTx(v int64) *DescribeInstanceTypesRequest
	GetMinimumInstancePpsTx() *int64
	SetMinimumLocalStorageAmount(v int32) *DescribeInstanceTypesRequest
	GetMinimumLocalStorageAmount() *int32
	SetMinimumLocalStorageCapacity(v int64) *DescribeInstanceTypesRequest
	GetMinimumLocalStorageCapacity() *int64
	SetMinimumMemorySize(v float32) *DescribeInstanceTypesRequest
	GetMinimumMemorySize() *float32
	SetMinimumPrimaryEniQueueNumber(v int32) *DescribeInstanceTypesRequest
	GetMinimumPrimaryEniQueueNumber() *int32
	SetMinimumQueuePairNumber(v int32) *DescribeInstanceTypesRequest
	GetMinimumQueuePairNumber() *int32
	SetMinimumSecondaryEniQueueNumber(v int32) *DescribeInstanceTypesRequest
	GetMinimumSecondaryEniQueueNumber() *int32
	SetNextToken(v string) *DescribeInstanceTypesRequest
	GetNextToken() *string
	SetNvmeSupport(v string) *DescribeInstanceTypesRequest
	GetNvmeSupport() *string
	SetOwnerAccount(v string) *DescribeInstanceTypesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeInstanceTypesRequest
	GetOwnerId() *int64
	SetPhysicalProcessorModel(v string) *DescribeInstanceTypesRequest
	GetPhysicalProcessorModel() *string
	SetPhysicalProcessorModels(v []*string) *DescribeInstanceTypesRequest
	GetPhysicalProcessorModels() []*string
	SetResourceOwnerAccount(v string) *DescribeInstanceTypesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeInstanceTypesRequest
	GetResourceOwnerId() *int64
}

type DescribeInstanceTypesRequest struct {
	// The list of advanced features to return for instance types.
	AdditionalAttributes []*string `json:"AdditionalAttributes,omitempty" xml:"AdditionalAttributes,omitempty" type:"Repeated"`
	// The CPU architecture. Valid values:
	//
	// - X86.
	//
	// - ARM.
	//
	// example:
	//
	// X86
	CpuArchitecture *string `json:"CpuArchitecture,omitempty" xml:"CpuArchitecture,omitempty"`
	// The CPU architectures to query. Array length: 1 to 2.
	CpuArchitectures []*string `json:"CpuArchitectures,omitempty" xml:"CpuArchitectures,omitempty" type:"Repeated"`
	// The GPU type.
	//
	// > Fuzzy matching is supported. For example, if the GPU type of an instance type is NVIDIA V100, you can enter NVIDIA to query the instance type.
	//
	// example:
	//
	// NVIDIA V100
	GPUSpec *string `json:"GPUSpec,omitempty" xml:"GPUSpec,omitempty"`
	// The GPU types to query. Array length: 1 to 10.
	GpuSpecs []*string `json:"GpuSpecs,omitempty" xml:"GpuSpecs,omitempty" type:"Repeated"`
	// The instance type categories to query. Array length: 1 to 10.
	InstanceCategories []*string `json:"InstanceCategories,omitempty" xml:"InstanceCategories,omitempty" type:"Repeated"`
	// The category of the instance type. Valid values:
	//
	// - General-purpose: general-purpose instance type.
	//
	// - Compute-optimized: compute-optimized instance type.
	//
	// - Memory-optimized: memory-optimized instance type.
	//
	// - Big data: big data instance type.
	//
	// - Local SSDs: local SSD instance type.
	//
	// - High Clock Speed: high clock speed instance type.
	//
	// - Enhanced: enhanced instance type.
	//
	// - Shared: shared instance type.
	//
	// - Compute-optimized with GPU: GPU-accelerated compute-optimized instance type.
	//
	// - Visual Compute-optimized: visual compute-optimized instance type.
	//
	// - Heterogeneous Service: heterogeneous service instance type.
	//
	// - Compute-optimized with FPGA: FPGA-accelerated compute-optimized instance type.
	//
	// - Compute-optimized with NPU: NPU-accelerated compute-optimized instance type.
	//
	// - ECS Bare Metal: ECS Bare Metal Instance.
	//
	// - Super Computing Cluster: Super Computing Cluster instance type.
	//
	// - High Performance Compute: high-performance computing instance type.
	//
	// example:
	//
	// Big data
	InstanceCategory *string `json:"InstanceCategory,omitempty" xml:"InstanceCategory,omitempty"`
	// The level of the instance family. Valid values:
	//
	// - EntryLevel: entry level (shared).
	//
	// - EnterpriseLevel: enterprise level.
	//
	// - CreditEntryLevel: credit-based entry level.
	//
	// example:
	//
	// EntryLevel
	InstanceFamilyLevel *string `json:"InstanceFamilyLevel,omitempty" xml:"InstanceFamilyLevel,omitempty"`
	// The instance families to query. Array length: 1 to 10.
	InstanceTypeFamilies []*string `json:"InstanceTypeFamilies,omitempty" xml:"InstanceTypeFamilies,omitempty" type:"Repeated"`
	// The instance family that the instance type belongs to. For valid values, see [DescribeInstanceTypeFamilies](https://help.aliyun.com/document_detail/25621.html).
	//
	// For more information about instance families, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// example:
	//
	// ecs.g6
	InstanceTypeFamily *string `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
	// The specified instance types. Array length: 1 to 10. If this parameter is not specified, information about all instance types is queried by default.
	//
	// example:
	//
	// ecs.g6.large
	InstanceTypes []*string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// The local disk categories to query. Array length: 1 to 2.
	LocalStorageCategories []*string `json:"LocalStorageCategories,omitempty" xml:"LocalStorageCategories,omitempty" type:"Repeated"`
	// The category of local disks. For more information, see [Local disks](~~63138#section_n2w_8yc_5u1~~). Valid values:
	//
	// - local_hdd_pro: SATA HDDs used by the d1ne and d1 instance families.
	//
	// - local_ssd_pro: NVMe SSDs used by the i2, i2g, i1, ga1, and gn5 instance families.
	//
	// example:
	//
	// local_ssd_pro
	LocalStorageCategory *string `json:"LocalStorageCategory,omitempty" xml:"LocalStorageCategory,omitempty"`
	// The maximum number of entries per page for a paginated query. Maximum value: 1600.
	//
	// Default value: 1600.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The maximum number of vCPU cores when you query instance types. Valid values: positive integers.
	//
	// > Instance types with more vCPU cores than the specified value are not returned.
	//
	// example:
	//
	// 10
	MaximumCpuCoreCount *int32 `json:"MaximumCpuCoreCount,omitempty" xml:"MaximumCpuCoreCount,omitempty"`
	// The maximum CPU base clock speed when you query instance types.
	//
	// > Instance types with a base clock speed higher than the specified value are not returned.
	//
	// example:
	//
	// 3.2
	MaximumCpuSpeedFrequency *float32 `json:"MaximumCpuSpeedFrequency,omitempty" xml:"MaximumCpuSpeedFrequency,omitempty"`
	// The maximum CPU turbo frequency when you query instance types.
	//
	// > Instance types with a turbo frequency higher than the specified value are not returned.
	//
	// example:
	//
	// 4.1
	MaximumCpuTurboFrequency *float32 `json:"MaximumCpuTurboFrequency,omitempty" xml:"MaximumCpuTurboFrequency,omitempty"`
	// The maximum number of GPUs when you query instance types. Valid values: positive integers.
	//
	// > Instance types with more GPUs than the specified value are not returned.
	//
	// example:
	//
	// 10
	MaximumGPUAmount *int32 `json:"MaximumGPUAmount,omitempty" xml:"MaximumGPUAmount,omitempty"`
	// The maximum memory size when you query instance types. Unit: GiB.
	//
	// > Instance types with a memory size larger than the specified value are not returned.
	//
	// example:
	//
	// 60
	MaximumMemorySize *float32 `json:"MaximumMemorySize,omitempty" xml:"MaximumMemorySize,omitempty"`
	// The minimum baseline vCPU computing performance (total of all vCPUs) for burstable instance types t5 and t6 when you query instance types.
	//
	// > Instance types with a baseline vCPU computing performance lower than the specified value are not returned.
	//
	// example:
	//
	// 12
	MinimumBaselineCredit *int32 `json:"MinimumBaselineCredit,omitempty" xml:"MinimumBaselineCredit,omitempty"`
	// The minimum number of vCPU cores when you query instance types. Valid values: positive integers.
	//
	// > Instance types with fewer vCPU cores than the specified value are not returned.
	//
	// example:
	//
	// 2
	MinimumCpuCoreCount *int32 `json:"MinimumCpuCoreCount,omitempty" xml:"MinimumCpuCoreCount,omitempty"`
	// The minimum CPU base clock speed when you query instance types.
	//
	// > Instance types with a base clock speed lower than the specified value are not returned.
	//
	// example:
	//
	// 2.5
	MinimumCpuSpeedFrequency *float32 `json:"MinimumCpuSpeedFrequency,omitempty" xml:"MinimumCpuSpeedFrequency,omitempty"`
	// The minimum CPU turbo frequency when you query instance types.
	//
	// > Instance types with a turbo frequency lower than the specified value are not returned.
	//
	// example:
	//
	// 3.2
	MinimumCpuTurboFrequency *float32 `json:"MinimumCpuTurboFrequency,omitempty" xml:"MinimumCpuTurboFrequency,omitempty"`
	// The minimum number of cloud disks that can be attached when you query instance types.
	//
	// > Instance types that support fewer cloud disks than the specified value are not returned.
	//
	// example:
	//
	// 4
	MinimumDiskQuantity *int32 `json:"MinimumDiskQuantity,omitempty" xml:"MinimumDiskQuantity,omitempty"`
	// The minimum number of IPv6 addresses per ENI when you query instance types.
	//
	// > Instance types that support fewer IPv6 addresses per ENI than the specified value are not returned.
	//
	// example:
	//
	// 2
	MinimumEniIpv6AddressQuantity *int32 `json:"MinimumEniIpv6AddressQuantity,omitempty" xml:"MinimumEniIpv6AddressQuantity,omitempty"`
	// The minimum number of IPv4 addresses per ENI when you query instance types.
	//
	// > Instance types that support fewer IPv4 addresses per ENI than the specified value are not returned.
	//
	// example:
	//
	// 2
	MinimumEniPrivateIpAddressQuantity *int32 `json:"MinimumEniPrivateIpAddressQuantity,omitempty" xml:"MinimumEniPrivateIpAddressQuantity,omitempty"`
	// The minimum number of ENIs that can be attached when you query instance types.
	//
	// > Instance types that support fewer ENIs than the specified value are not returned.
	//
	// example:
	//
	// 4
	MinimumEniQuantity *int32 `json:"MinimumEniQuantity,omitempty" xml:"MinimumEniQuantity,omitempty"`
	// The minimum number of ERIs when you query instance types.
	//
	// > Instance types with fewer ERIs than the specified value are not returned.
	//
	// example:
	//
	// 0
	MinimumEriQuantity *int32 `json:"MinimumEriQuantity,omitempty" xml:"MinimumEriQuantity,omitempty"`
	// The minimum number of GPUs when you query instance types. Valid values: positive integers.
	//
	// > Instance types with fewer GPUs than the specified value are not returned.
	//
	// example:
	//
	// 3
	MinimumGPUAmount *int32 `json:"MinimumGPUAmount,omitempty" xml:"MinimumGPUAmount,omitempty"`
	// The minimum initial vCPU credits for burstable instance types t5 and t6 when you query instance types.
	//
	// > Instance types with initial vCPU credits lower than the specified value are not returned.
	//
	// example:
	//
	// 12
	MinimumInitialCredit *int32 `json:"MinimumInitialCredit,omitempty" xml:"MinimumInitialCredit,omitempty"`
	// The minimum inbound internal bandwidth when you query instance types. Unit: kbit/s.
	//
	// > Instance types with an inbound internal bandwidth lower than the specified value are not returned.
	//
	// example:
	//
	// 12288
	MinimumInstanceBandwidthRx *int32 `json:"MinimumInstanceBandwidthRx,omitempty" xml:"MinimumInstanceBandwidthRx,omitempty"`
	// The minimum outbound internal bandwidth when you query instance types. Unit: kbit/s.
	//
	// > Instance types with an outbound internal bandwidth lower than the specified value are not returned.
	//
	// example:
	//
	// 12288
	MinimumInstanceBandwidthTx *int32 `json:"MinimumInstanceBandwidthTx,omitempty" xml:"MinimumInstanceBandwidthTx,omitempty"`
	// The minimum inbound packet forwarding rate over the internal network when you query instance types. Unit: pps.
	//
	// > Instance types with an inbound packet forwarding rate lower than the specified value are not returned.
	//
	// example:
	//
	// 15
	MinimumInstancePpsRx *int64 `json:"MinimumInstancePpsRx,omitempty" xml:"MinimumInstancePpsRx,omitempty"`
	// The minimum outbound packet forwarding rate over the internal network when you query instance types. Unit: pps.
	//
	// > Instance types with an outbound packet forwarding rate lower than the specified value are not returned.
	//
	// example:
	//
	// 15
	MinimumInstancePpsTx *int64 `json:"MinimumInstancePpsTx,omitempty" xml:"MinimumInstancePpsTx,omitempty"`
	// The minimum number of local disks attached to the instance when you query instance types.
	//
	// > Instance types with fewer local disks than the specified value are not returned.
	//
	// example:
	//
	// 4
	MinimumLocalStorageAmount *int32 `json:"MinimumLocalStorageAmount,omitempty" xml:"MinimumLocalStorageAmount,omitempty"`
	// The capacity of a single local disk attached to the instance. Unit: GiB.
	//
	// example:
	//
	// 40
	MinimumLocalStorageCapacity *int64 `json:"MinimumLocalStorageCapacity,omitempty" xml:"MinimumLocalStorageCapacity,omitempty"`
	// The minimum memory size when you query instance types. Unit: GiB.
	//
	// > Instance types with a memory size smaller than the specified value are not returned.
	//
	// example:
	//
	// 20
	MinimumMemorySize *float32 `json:"MinimumMemorySize,omitempty" xml:"MinimumMemorySize,omitempty"`
	// The minimum default number of queues supported by the primary ENI when you query instance types.
	//
	// > Instance types with fewer default queues on the primary ENI than the specified value are not returned.
	//
	// example:
	//
	// 8
	MinimumPrimaryEniQueueNumber *int32 `json:"MinimumPrimaryEniQueueNumber,omitempty" xml:"MinimumPrimaryEniQueueNumber,omitempty"`
	// The minimum number of QueuePair (QP) queues per Elastic RDMA Interface (ERI) when you query instance types.
	//
	// > Instance types with fewer QP queues per ERI than the specified value are not returned.
	//
	// example:
	//
	// 8
	MinimumQueuePairNumber *int32 `json:"MinimumQueuePairNumber,omitempty" xml:"MinimumQueuePairNumber,omitempty"`
	// The minimum default number of queues supported by the secondary ENI when you query instance types.
	//
	// > Instance types with fewer default queues on the secondary ENI than the specified value are not returned.
	//
	// example:
	//
	// 4
	MinimumSecondaryEniQueueNumber *int32 `json:"MinimumSecondaryEniQueueNumber,omitempty" xml:"MinimumSecondaryEniQueueNumber,omitempty"`
	// The pagination token. Set this parameter to the NextToken value returned in the previous call to this API operation. You do not need to set this parameter for the first request.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// e71d8a535bd9cc11
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Indicates whether cloud disks attached to the instance type support NVMe. Valid values:
	//
	// - required: supported. Cloud disks are attached in NVMe mode.
	//
	// - unsupported: not supported. Cloud disks are not attached in NVMe mode.
	//
	// example:
	//
	// required
	NvmeSupport  *string `json:"NvmeSupport,omitempty" xml:"NvmeSupport,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The processor model.
	//
	// > Fuzzy matching is supported. For example, if the processor model of an instance type is Intel Xeon(Ice Lake) Platinum 8369B, you can enter Intel to query the instance type.
	//
	// example:
	//
	// Intel Xeon(Ice Lake) Platinum 8369B
	PhysicalProcessorModel *string `json:"PhysicalProcessorModel,omitempty" xml:"PhysicalProcessorModel,omitempty"`
	// The processor models to query. Array length: 1 to 10.
	PhysicalProcessorModels []*string `json:"PhysicalProcessorModels,omitempty" xml:"PhysicalProcessorModels,omitempty" type:"Repeated"`
	ResourceOwnerAccount    *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId         *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeInstanceTypesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTypesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesRequest) GetAdditionalAttributes() []*string {
	return s.AdditionalAttributes
}

func (s *DescribeInstanceTypesRequest) GetCpuArchitecture() *string {
	return s.CpuArchitecture
}

func (s *DescribeInstanceTypesRequest) GetCpuArchitectures() []*string {
	return s.CpuArchitectures
}

func (s *DescribeInstanceTypesRequest) GetGPUSpec() *string {
	return s.GPUSpec
}

func (s *DescribeInstanceTypesRequest) GetGpuSpecs() []*string {
	return s.GpuSpecs
}

func (s *DescribeInstanceTypesRequest) GetInstanceCategories() []*string {
	return s.InstanceCategories
}

func (s *DescribeInstanceTypesRequest) GetInstanceCategory() *string {
	return s.InstanceCategory
}

func (s *DescribeInstanceTypesRequest) GetInstanceFamilyLevel() *string {
	return s.InstanceFamilyLevel
}

func (s *DescribeInstanceTypesRequest) GetInstanceTypeFamilies() []*string {
	return s.InstanceTypeFamilies
}

func (s *DescribeInstanceTypesRequest) GetInstanceTypeFamily() *string {
	return s.InstanceTypeFamily
}

func (s *DescribeInstanceTypesRequest) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *DescribeInstanceTypesRequest) GetLocalStorageCategories() []*string {
	return s.LocalStorageCategories
}

func (s *DescribeInstanceTypesRequest) GetLocalStorageCategory() *string {
	return s.LocalStorageCategory
}

func (s *DescribeInstanceTypesRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *DescribeInstanceTypesRequest) GetMaximumCpuCoreCount() *int32 {
	return s.MaximumCpuCoreCount
}

func (s *DescribeInstanceTypesRequest) GetMaximumCpuSpeedFrequency() *float32 {
	return s.MaximumCpuSpeedFrequency
}

func (s *DescribeInstanceTypesRequest) GetMaximumCpuTurboFrequency() *float32 {
	return s.MaximumCpuTurboFrequency
}

func (s *DescribeInstanceTypesRequest) GetMaximumGPUAmount() *int32 {
	return s.MaximumGPUAmount
}

func (s *DescribeInstanceTypesRequest) GetMaximumMemorySize() *float32 {
	return s.MaximumMemorySize
}

func (s *DescribeInstanceTypesRequest) GetMinimumBaselineCredit() *int32 {
	return s.MinimumBaselineCredit
}

func (s *DescribeInstanceTypesRequest) GetMinimumCpuCoreCount() *int32 {
	return s.MinimumCpuCoreCount
}

func (s *DescribeInstanceTypesRequest) GetMinimumCpuSpeedFrequency() *float32 {
	return s.MinimumCpuSpeedFrequency
}

func (s *DescribeInstanceTypesRequest) GetMinimumCpuTurboFrequency() *float32 {
	return s.MinimumCpuTurboFrequency
}

func (s *DescribeInstanceTypesRequest) GetMinimumDiskQuantity() *int32 {
	return s.MinimumDiskQuantity
}

func (s *DescribeInstanceTypesRequest) GetMinimumEniIpv6AddressQuantity() *int32 {
	return s.MinimumEniIpv6AddressQuantity
}

func (s *DescribeInstanceTypesRequest) GetMinimumEniPrivateIpAddressQuantity() *int32 {
	return s.MinimumEniPrivateIpAddressQuantity
}

func (s *DescribeInstanceTypesRequest) GetMinimumEniQuantity() *int32 {
	return s.MinimumEniQuantity
}

func (s *DescribeInstanceTypesRequest) GetMinimumEriQuantity() *int32 {
	return s.MinimumEriQuantity
}

func (s *DescribeInstanceTypesRequest) GetMinimumGPUAmount() *int32 {
	return s.MinimumGPUAmount
}

func (s *DescribeInstanceTypesRequest) GetMinimumInitialCredit() *int32 {
	return s.MinimumInitialCredit
}

func (s *DescribeInstanceTypesRequest) GetMinimumInstanceBandwidthRx() *int32 {
	return s.MinimumInstanceBandwidthRx
}

func (s *DescribeInstanceTypesRequest) GetMinimumInstanceBandwidthTx() *int32 {
	return s.MinimumInstanceBandwidthTx
}

func (s *DescribeInstanceTypesRequest) GetMinimumInstancePpsRx() *int64 {
	return s.MinimumInstancePpsRx
}

func (s *DescribeInstanceTypesRequest) GetMinimumInstancePpsTx() *int64 {
	return s.MinimumInstancePpsTx
}

func (s *DescribeInstanceTypesRequest) GetMinimumLocalStorageAmount() *int32 {
	return s.MinimumLocalStorageAmount
}

func (s *DescribeInstanceTypesRequest) GetMinimumLocalStorageCapacity() *int64 {
	return s.MinimumLocalStorageCapacity
}

func (s *DescribeInstanceTypesRequest) GetMinimumMemorySize() *float32 {
	return s.MinimumMemorySize
}

func (s *DescribeInstanceTypesRequest) GetMinimumPrimaryEniQueueNumber() *int32 {
	return s.MinimumPrimaryEniQueueNumber
}

func (s *DescribeInstanceTypesRequest) GetMinimumQueuePairNumber() *int32 {
	return s.MinimumQueuePairNumber
}

func (s *DescribeInstanceTypesRequest) GetMinimumSecondaryEniQueueNumber() *int32 {
	return s.MinimumSecondaryEniQueueNumber
}

func (s *DescribeInstanceTypesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeInstanceTypesRequest) GetNvmeSupport() *string {
	return s.NvmeSupport
}

func (s *DescribeInstanceTypesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeInstanceTypesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInstanceTypesRequest) GetPhysicalProcessorModel() *string {
	return s.PhysicalProcessorModel
}

func (s *DescribeInstanceTypesRequest) GetPhysicalProcessorModels() []*string {
	return s.PhysicalProcessorModels
}

func (s *DescribeInstanceTypesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeInstanceTypesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeInstanceTypesRequest) SetAdditionalAttributes(v []*string) *DescribeInstanceTypesRequest {
	s.AdditionalAttributes = v
	return s
}

func (s *DescribeInstanceTypesRequest) SetCpuArchitecture(v string) *DescribeInstanceTypesRequest {
	s.CpuArchitecture = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetCpuArchitectures(v []*string) *DescribeInstanceTypesRequest {
	s.CpuArchitectures = v
	return s
}

func (s *DescribeInstanceTypesRequest) SetGPUSpec(v string) *DescribeInstanceTypesRequest {
	s.GPUSpec = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetGpuSpecs(v []*string) *DescribeInstanceTypesRequest {
	s.GpuSpecs = v
	return s
}

func (s *DescribeInstanceTypesRequest) SetInstanceCategories(v []*string) *DescribeInstanceTypesRequest {
	s.InstanceCategories = v
	return s
}

func (s *DescribeInstanceTypesRequest) SetInstanceCategory(v string) *DescribeInstanceTypesRequest {
	s.InstanceCategory = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetInstanceFamilyLevel(v string) *DescribeInstanceTypesRequest {
	s.InstanceFamilyLevel = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetInstanceTypeFamilies(v []*string) *DescribeInstanceTypesRequest {
	s.InstanceTypeFamilies = v
	return s
}

func (s *DescribeInstanceTypesRequest) SetInstanceTypeFamily(v string) *DescribeInstanceTypesRequest {
	s.InstanceTypeFamily = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetInstanceTypes(v []*string) *DescribeInstanceTypesRequest {
	s.InstanceTypes = v
	return s
}

func (s *DescribeInstanceTypesRequest) SetLocalStorageCategories(v []*string) *DescribeInstanceTypesRequest {
	s.LocalStorageCategories = v
	return s
}

func (s *DescribeInstanceTypesRequest) SetLocalStorageCategory(v string) *DescribeInstanceTypesRequest {
	s.LocalStorageCategory = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMaxResults(v int64) *DescribeInstanceTypesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMaximumCpuCoreCount(v int32) *DescribeInstanceTypesRequest {
	s.MaximumCpuCoreCount = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMaximumCpuSpeedFrequency(v float32) *DescribeInstanceTypesRequest {
	s.MaximumCpuSpeedFrequency = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMaximumCpuTurboFrequency(v float32) *DescribeInstanceTypesRequest {
	s.MaximumCpuTurboFrequency = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMaximumGPUAmount(v int32) *DescribeInstanceTypesRequest {
	s.MaximumGPUAmount = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMaximumMemorySize(v float32) *DescribeInstanceTypesRequest {
	s.MaximumMemorySize = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumBaselineCredit(v int32) *DescribeInstanceTypesRequest {
	s.MinimumBaselineCredit = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumCpuCoreCount(v int32) *DescribeInstanceTypesRequest {
	s.MinimumCpuCoreCount = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumCpuSpeedFrequency(v float32) *DescribeInstanceTypesRequest {
	s.MinimumCpuSpeedFrequency = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumCpuTurboFrequency(v float32) *DescribeInstanceTypesRequest {
	s.MinimumCpuTurboFrequency = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumDiskQuantity(v int32) *DescribeInstanceTypesRequest {
	s.MinimumDiskQuantity = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumEniIpv6AddressQuantity(v int32) *DescribeInstanceTypesRequest {
	s.MinimumEniIpv6AddressQuantity = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumEniPrivateIpAddressQuantity(v int32) *DescribeInstanceTypesRequest {
	s.MinimumEniPrivateIpAddressQuantity = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumEniQuantity(v int32) *DescribeInstanceTypesRequest {
	s.MinimumEniQuantity = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumEriQuantity(v int32) *DescribeInstanceTypesRequest {
	s.MinimumEriQuantity = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumGPUAmount(v int32) *DescribeInstanceTypesRequest {
	s.MinimumGPUAmount = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumInitialCredit(v int32) *DescribeInstanceTypesRequest {
	s.MinimumInitialCredit = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumInstanceBandwidthRx(v int32) *DescribeInstanceTypesRequest {
	s.MinimumInstanceBandwidthRx = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumInstanceBandwidthTx(v int32) *DescribeInstanceTypesRequest {
	s.MinimumInstanceBandwidthTx = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumInstancePpsRx(v int64) *DescribeInstanceTypesRequest {
	s.MinimumInstancePpsRx = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumInstancePpsTx(v int64) *DescribeInstanceTypesRequest {
	s.MinimumInstancePpsTx = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumLocalStorageAmount(v int32) *DescribeInstanceTypesRequest {
	s.MinimumLocalStorageAmount = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumLocalStorageCapacity(v int64) *DescribeInstanceTypesRequest {
	s.MinimumLocalStorageCapacity = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumMemorySize(v float32) *DescribeInstanceTypesRequest {
	s.MinimumMemorySize = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumPrimaryEniQueueNumber(v int32) *DescribeInstanceTypesRequest {
	s.MinimumPrimaryEniQueueNumber = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumQueuePairNumber(v int32) *DescribeInstanceTypesRequest {
	s.MinimumQueuePairNumber = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumSecondaryEniQueueNumber(v int32) *DescribeInstanceTypesRequest {
	s.MinimumSecondaryEniQueueNumber = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetNextToken(v string) *DescribeInstanceTypesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetNvmeSupport(v string) *DescribeInstanceTypesRequest {
	s.NvmeSupport = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetOwnerAccount(v string) *DescribeInstanceTypesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetOwnerId(v int64) *DescribeInstanceTypesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetPhysicalProcessorModel(v string) *DescribeInstanceTypesRequest {
	s.PhysicalProcessorModel = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetPhysicalProcessorModels(v []*string) *DescribeInstanceTypesRequest {
	s.PhysicalProcessorModels = v
	return s
}

func (s *DescribeInstanceTypesRequest) SetResourceOwnerAccount(v string) *DescribeInstanceTypesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetResourceOwnerId(v int64) *DescribeInstanceTypesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInstanceTypesRequest) Validate() error {
	return dara.Validate(s)
}
