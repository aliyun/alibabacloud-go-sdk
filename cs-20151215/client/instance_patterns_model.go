// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstancePatterns interface {
	dara.Model
	String() string
	GoString() string
	SetArchitectures(v []*string) *InstancePatterns
	GetArchitectures() []*string
	SetBurstPerformanceOption(v string) *InstancePatterns
	GetBurstPerformanceOption() *string
	SetCore(v int64) *InstancePatterns
	GetCore() *int64
	SetCores(v int64) *InstancePatterns
	GetCores() *int64
	SetCpuArchitectures(v []*string) *InstancePatterns
	GetCpuArchitectures() []*string
	SetExcludedInstanceTypes(v []*string) *InstancePatterns
	GetExcludedInstanceTypes() []*string
	SetInstanceCategories(v []*string) *InstancePatterns
	GetInstanceCategories() []*string
	SetInstanceFamilyLevel(v string) *InstancePatterns
	GetInstanceFamilyLevel() *string
	SetInstanceTypeFamilies(v []*string) *InstancePatterns
	GetInstanceTypeFamilies() []*string
	SetMaxCpuCores(v int64) *InstancePatterns
	GetMaxCpuCores() *int64
	SetMaxMemorySize(v float32) *InstancePatterns
	GetMaxMemorySize() *float32
	SetMaxPrice(v float32) *InstancePatterns
	GetMaxPrice() *float32
	SetMaximumGpuAmount(v int64) *InstancePatterns
	GetMaximumGpuAmount() *int64
	SetMemory(v float32) *InstancePatterns
	GetMemory() *float32
	SetMinCpuCores(v int64) *InstancePatterns
	GetMinCpuCores() *int64
	SetMinMemorySize(v float32) *InstancePatterns
	GetMinMemorySize() *float32
	SetMinimumEniIpv6AddressQuantity(v int64) *InstancePatterns
	GetMinimumEniIpv6AddressQuantity() *int64
	SetMinimumEniPrivateIpAddressQuantity(v int64) *InstancePatterns
	GetMinimumEniPrivateIpAddressQuantity() *int64
	SetMinimumEniQuantity(v int64) *InstancePatterns
	GetMinimumEniQuantity() *int64
}

type InstancePatterns struct {
	// Deprecated
	Architectures []*string `json:"architectures,omitempty" xml:"architectures,omitempty" type:"Repeated"`
	// Deprecated
	//
	// example:
	//
	// Exclude
	BurstPerformanceOption *string `json:"burst_performance_option,omitempty" xml:"burst_performance_option,omitempty"`
	// Deprecated
	//
	// example:
	//
	// 4
	Core *int64 `json:"core,omitempty" xml:"core,omitempty"`
	// example:
	//
	// 4
	Cores                 *int64    `json:"cores,omitempty" xml:"cores,omitempty"`
	CpuArchitectures      []*string `json:"cpu_architectures,omitempty" xml:"cpu_architectures,omitempty" type:"Repeated"`
	ExcludedInstanceTypes []*string `json:"excluded_instance_types,omitempty" xml:"excluded_instance_types,omitempty" type:"Repeated"`
	InstanceCategories    []*string `json:"instance_categories,omitempty" xml:"instance_categories,omitempty" type:"Repeated"`
	// example:
	//
	// EnterpriseLevel
	InstanceFamilyLevel  *string   `json:"instance_family_level,omitempty" xml:"instance_family_level,omitempty"`
	InstanceTypeFamilies []*string `json:"instance_type_families,omitempty" xml:"instance_type_families,omitempty" type:"Repeated"`
	// example:
	//
	// 8
	MaxCpuCores *int64 `json:"max_cpu_cores,omitempty" xml:"max_cpu_cores,omitempty"`
	// example:
	//
	// 16
	MaxMemorySize *float32 `json:"max_memory_size,omitempty" xml:"max_memory_size,omitempty"`
	// Deprecated
	//
	// example:
	//
	// 2
	MaxPrice         *float32 `json:"max_price,omitempty" xml:"max_price,omitempty"`
	MaximumGpuAmount *int64   `json:"maximum_gpu_amount,omitempty" xml:"maximum_gpu_amount,omitempty"`
	// example:
	//
	// 8
	Memory *float32 `json:"memory,omitempty" xml:"memory,omitempty"`
	// example:
	//
	// 4
	MinCpuCores *int64 `json:"min_cpu_cores,omitempty" xml:"min_cpu_cores,omitempty"`
	// example:
	//
	// 8
	MinMemorySize                      *float32 `json:"min_memory_size,omitempty" xml:"min_memory_size,omitempty"`
	MinimumEniIpv6AddressQuantity      *int64   `json:"minimum_eni_ipv6_address_quantity,omitempty" xml:"minimum_eni_ipv6_address_quantity,omitempty"`
	MinimumEniPrivateIpAddressQuantity *int64   `json:"minimum_eni_private_ip_address_quantity,omitempty" xml:"minimum_eni_private_ip_address_quantity,omitempty"`
	MinimumEniQuantity                 *int64   `json:"minimum_eni_quantity,omitempty" xml:"minimum_eni_quantity,omitempty"`
}

func (s InstancePatterns) String() string {
	return dara.Prettify(s)
}

func (s InstancePatterns) GoString() string {
	return s.String()
}

func (s *InstancePatterns) GetArchitectures() []*string {
	return s.Architectures
}

func (s *InstancePatterns) GetBurstPerformanceOption() *string {
	return s.BurstPerformanceOption
}

func (s *InstancePatterns) GetCore() *int64 {
	return s.Core
}

func (s *InstancePatterns) GetCores() *int64 {
	return s.Cores
}

func (s *InstancePatterns) GetCpuArchitectures() []*string {
	return s.CpuArchitectures
}

func (s *InstancePatterns) GetExcludedInstanceTypes() []*string {
	return s.ExcludedInstanceTypes
}

func (s *InstancePatterns) GetInstanceCategories() []*string {
	return s.InstanceCategories
}

func (s *InstancePatterns) GetInstanceFamilyLevel() *string {
	return s.InstanceFamilyLevel
}

func (s *InstancePatterns) GetInstanceTypeFamilies() []*string {
	return s.InstanceTypeFamilies
}

func (s *InstancePatterns) GetMaxCpuCores() *int64 {
	return s.MaxCpuCores
}

func (s *InstancePatterns) GetMaxMemorySize() *float32 {
	return s.MaxMemorySize
}

func (s *InstancePatterns) GetMaxPrice() *float32 {
	return s.MaxPrice
}

func (s *InstancePatterns) GetMaximumGpuAmount() *int64 {
	return s.MaximumGpuAmount
}

func (s *InstancePatterns) GetMemory() *float32 {
	return s.Memory
}

func (s *InstancePatterns) GetMinCpuCores() *int64 {
	return s.MinCpuCores
}

func (s *InstancePatterns) GetMinMemorySize() *float32 {
	return s.MinMemorySize
}

func (s *InstancePatterns) GetMinimumEniIpv6AddressQuantity() *int64 {
	return s.MinimumEniIpv6AddressQuantity
}

func (s *InstancePatterns) GetMinimumEniPrivateIpAddressQuantity() *int64 {
	return s.MinimumEniPrivateIpAddressQuantity
}

func (s *InstancePatterns) GetMinimumEniQuantity() *int64 {
	return s.MinimumEniQuantity
}

func (s *InstancePatterns) SetArchitectures(v []*string) *InstancePatterns {
	s.Architectures = v
	return s
}

func (s *InstancePatterns) SetBurstPerformanceOption(v string) *InstancePatterns {
	s.BurstPerformanceOption = &v
	return s
}

func (s *InstancePatterns) SetCore(v int64) *InstancePatterns {
	s.Core = &v
	return s
}

func (s *InstancePatterns) SetCores(v int64) *InstancePatterns {
	s.Cores = &v
	return s
}

func (s *InstancePatterns) SetCpuArchitectures(v []*string) *InstancePatterns {
	s.CpuArchitectures = v
	return s
}

func (s *InstancePatterns) SetExcludedInstanceTypes(v []*string) *InstancePatterns {
	s.ExcludedInstanceTypes = v
	return s
}

func (s *InstancePatterns) SetInstanceCategories(v []*string) *InstancePatterns {
	s.InstanceCategories = v
	return s
}

func (s *InstancePatterns) SetInstanceFamilyLevel(v string) *InstancePatterns {
	s.InstanceFamilyLevel = &v
	return s
}

func (s *InstancePatterns) SetInstanceTypeFamilies(v []*string) *InstancePatterns {
	s.InstanceTypeFamilies = v
	return s
}

func (s *InstancePatterns) SetMaxCpuCores(v int64) *InstancePatterns {
	s.MaxCpuCores = &v
	return s
}

func (s *InstancePatterns) SetMaxMemorySize(v float32) *InstancePatterns {
	s.MaxMemorySize = &v
	return s
}

func (s *InstancePatterns) SetMaxPrice(v float32) *InstancePatterns {
	s.MaxPrice = &v
	return s
}

func (s *InstancePatterns) SetMaximumGpuAmount(v int64) *InstancePatterns {
	s.MaximumGpuAmount = &v
	return s
}

func (s *InstancePatterns) SetMemory(v float32) *InstancePatterns {
	s.Memory = &v
	return s
}

func (s *InstancePatterns) SetMinCpuCores(v int64) *InstancePatterns {
	s.MinCpuCores = &v
	return s
}

func (s *InstancePatterns) SetMinMemorySize(v float32) *InstancePatterns {
	s.MinMemorySize = &v
	return s
}

func (s *InstancePatterns) SetMinimumEniIpv6AddressQuantity(v int64) *InstancePatterns {
	s.MinimumEniIpv6AddressQuantity = &v
	return s
}

func (s *InstancePatterns) SetMinimumEniPrivateIpAddressQuantity(v int64) *InstancePatterns {
	s.MinimumEniPrivateIpAddressQuantity = &v
	return s
}

func (s *InstancePatterns) SetMinimumEniQuantity(v int64) *InstancePatterns {
	s.MinimumEniQuantity = &v
	return s
}

func (s *InstancePatterns) Validate() error {
	return dara.Validate(s)
}
