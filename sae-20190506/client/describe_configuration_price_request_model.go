// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConfigurationPriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBestEffortType(v string) *DescribeConfigurationPriceRequest
	GetBestEffortType() *string
	SetCpu(v int32) *DescribeConfigurationPriceRequest
	GetCpu() *int32
	SetGpuA10(v string) *DescribeConfigurationPriceRequest
	GetGpuA10() *string
	SetGpuPpu810e(v string) *DescribeConfigurationPriceRequest
	GetGpuPpu810e() *string
	SetMemory(v int32) *DescribeConfigurationPriceRequest
	GetMemory() *int32
	SetNewSaeVersion(v string) *DescribeConfigurationPriceRequest
	GetNewSaeVersion() *string
	SetResourceType(v string) *DescribeConfigurationPriceRequest
	GetResourceType() *string
	SetWorkload(v string) *DescribeConfigurationPriceRequest
	GetWorkload() *string
}

type DescribeConfigurationPriceRequest struct {
	// The BestEffort policy. Valid values:
	//
	// - besteffort: BestEffort.
	//
	// - try-besteffort: BestEffort preferred.
	//
	// - default: Default.
	//
	// example:
	//
	// default
	BestEffortType *string `json:"BestEffortType,omitempty" xml:"BestEffortType,omitempty"`
	// The required CPU for each instance, in millicores. The value cannot be 0. Only the following defined specifications are supported:
	//
	// - **500**
	//
	// - **1000**
	//
	// - **2000**
	//
	// - **4000**
	//
	// - **8000**
	//
	// - **12000**
	//
	// - **16000**
	//
	// - **32000**
	//
	// This parameter is required.
	//
	// example:
	//
	// 2000
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The number of A10 GPUs.
	//
	// example:
	//
	// 0
	GpuA10 *string `json:"GpuA10,omitempty" xml:"GpuA10,omitempty"`
	// The number of PPU 810E units.
	//
	// example:
	//
	// 0
	GpuPpu810e *string `json:"GpuPpu810e,omitempty" xml:"GpuPpu810e,omitempty"`
	// The required memory for each instance, in MB. The value cannot be 0. Memory has a one-to-one mapping with CPU. Only the following defined specifications are supported:
	//
	// - **1024**: Corresponds to 500 and 1000 millicores.
	//
	// - **2048**: Corresponds to 500, 1000, and 2000 millicores.
	//
	// - **4096**: Corresponds to 1000, 2000, and 4000 millicores.
	//
	// - **8192**: Corresponds to 2000, 4000, and 8000 millicores.
	//
	// - **12288**: Corresponds to 12000 millicores.
	//
	// - **16384**: Corresponds to 4000, 8000, and 16000 millicores.
	//
	// - **24576**: Corresponds to 12000 millicores.
	//
	// - **32768**: Corresponds to 16000 millicores.
	//
	// - **65536**: Corresponds to 8000, 16000, and 32000 millicores.
	//
	// - **131072**: Corresponds to 32000 millicores.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4096
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The application edition. Valid values:
	//
	// - lite: Lite Edition.
	//
	// - std: Standard Edition.
	//
	// - pro: Professional Edition.
	//
	// example:
	//
	// std
	NewSaeVersion *string `json:"NewSaeVersion,omitempty" xml:"NewSaeVersion,omitempty"`
	// The resource type. Valid values: NULL (default), default (default), and haiguang (Hygon server).
	//
	// example:
	//
	// default
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// Common scenarios. Valid values:
	//
	// - web
	//
	// - micro_service
	//
	// example:
	//
	// Web
	Workload *string `json:"Workload,omitempty" xml:"Workload,omitempty"`
}

func (s DescribeConfigurationPriceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigurationPriceRequest) GoString() string {
	return s.String()
}

func (s *DescribeConfigurationPriceRequest) GetBestEffortType() *string {
	return s.BestEffortType
}

func (s *DescribeConfigurationPriceRequest) GetCpu() *int32 {
	return s.Cpu
}

func (s *DescribeConfigurationPriceRequest) GetGpuA10() *string {
	return s.GpuA10
}

func (s *DescribeConfigurationPriceRequest) GetGpuPpu810e() *string {
	return s.GpuPpu810e
}

func (s *DescribeConfigurationPriceRequest) GetMemory() *int32 {
	return s.Memory
}

func (s *DescribeConfigurationPriceRequest) GetNewSaeVersion() *string {
	return s.NewSaeVersion
}

func (s *DescribeConfigurationPriceRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeConfigurationPriceRequest) GetWorkload() *string {
	return s.Workload
}

func (s *DescribeConfigurationPriceRequest) SetBestEffortType(v string) *DescribeConfigurationPriceRequest {
	s.BestEffortType = &v
	return s
}

func (s *DescribeConfigurationPriceRequest) SetCpu(v int32) *DescribeConfigurationPriceRequest {
	s.Cpu = &v
	return s
}

func (s *DescribeConfigurationPriceRequest) SetGpuA10(v string) *DescribeConfigurationPriceRequest {
	s.GpuA10 = &v
	return s
}

func (s *DescribeConfigurationPriceRequest) SetGpuPpu810e(v string) *DescribeConfigurationPriceRequest {
	s.GpuPpu810e = &v
	return s
}

func (s *DescribeConfigurationPriceRequest) SetMemory(v int32) *DescribeConfigurationPriceRequest {
	s.Memory = &v
	return s
}

func (s *DescribeConfigurationPriceRequest) SetNewSaeVersion(v string) *DescribeConfigurationPriceRequest {
	s.NewSaeVersion = &v
	return s
}

func (s *DescribeConfigurationPriceRequest) SetResourceType(v string) *DescribeConfigurationPriceRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeConfigurationPriceRequest) SetWorkload(v string) *DescribeConfigurationPriceRequest {
	s.Workload = &v
	return s
}

func (s *DescribeConfigurationPriceRequest) Validate() error {
	return dara.Validate(s)
}
