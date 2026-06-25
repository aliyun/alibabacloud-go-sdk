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
	// - besteffort: BestEffort
	//
	// - try-besteffort: BestEffort preferred
	//
	// - default: default
	//
	// example:
	//
	// default
	BestEffortType *string `json:"BestEffortType,omitempty" xml:"BestEffortType,omitempty"`
	// The number of CPU cores required for each instance. Unit: millicores. This value cannot be 0. Only the following defined specifications are supported:
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
	Cpu        *int32  `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	GpuA10     *string `json:"GpuA10,omitempty" xml:"GpuA10,omitempty"`
	GpuPpu810e *string `json:"GpuPpu810e,omitempty" xml:"GpuPpu810e,omitempty"`
	// The amount of memory required for each instance. Unit: MB. This value cannot be 0. The memory size must correspond to the CPU specification. Only the following defined specifications are supported:
	//
	// - **1024**: Corresponds to 500 millicores and 1,000 millicores of CPU.
	//
	// - **2048**: Corresponds to 500, 1,000, and 2,000 millicores of CPU.
	//
	// - **4096**: Corresponds to 1,000, 2,000, and 4,000 millicores of CPU.
	//
	// - **8192**: Corresponds to 2,000, 4,000, and 8,000 millicores of CPU.
	//
	// - **12288**: Corresponds to 12,000 millicores of CPU.
	//
	// - **16384**: Corresponds to 4,000, 8,000, and 16,000 millicores of CPU.
	//
	// - **24576**: Corresponds to 12,000 millicores of CPU.
	//
	// - **32768**: Corresponds to 16,000 millicores of CPU.
	//
	// - **65536**: Corresponds to 8,000, 16,000, and 32,000 millicores of CPU.
	//
	// - **131072**: Corresponds to 32,000 millicores of CPU.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4096
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The application version. Valid values:
	//
	// - lite: Lightweight Edition
	//
	// - std: Standard Edition
	//
	// - pro: Professional Edition
	//
	// example:
	//
	// std
	NewSaeVersion *string `json:"NewSaeVersion,omitempty" xml:"NewSaeVersion,omitempty"`
	// The resource type. Valid values: NULL (default), default, and haiguang (Haiguang server).
	//
	// example:
	//
	// default
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The application scenario. Valid values:
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
