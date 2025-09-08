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
	BestEffortType *string `json:"BestEffortType,omitempty" xml:"BestEffortType,omitempty"`
	// The CPU specifications that are required for each instance. Unit: millicores. This parameter cannot be set to 0. Valid values:
	//
	// 	- **500**
	//
	// 	- **1000**
	//
	// 	- **2000**
	//
	// 	- **4000**
	//
	// 	- **8000**
	//
	// 	- **12000**
	//
	// 	- **16000**
	//
	// 	- **32000**
	//
	// This parameter is required.
	//
	// example:
	//
	// 2000
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The memory size that is required by each instance. Unit: MB. This parameter cannot be set to 0. The values of this parameter correspond to the values of the Cpu parameter:
	//
	// 	- This parameter is set to **1024*	- if the Cpu parameter is set to 500 or 1000.
	//
	// 	- This parameter is set to **2048*	- if the Cpu parameter is set to 500, 1000, or 2000.
	//
	// 	- This parameter is set to **4096*	- if the Cpu parameter is set to 1000, 2000, or 4000.
	//
	// 	- This parameter is set to **8192*	- if the Cpu parameter is set to 2000, 4000, or 8,000.
	//
	// 	- This parameter is set to **12288*	- if the Cpu parameter is set to 12000.
	//
	// 	- This parameter is set to **16384*	- if the Cpu parameter is set to 4000, 8000, or 16000.
	//
	// 	- This parameter is set to **24576*	- if the Cpu parameter is set to 12000.
	//
	// 	- This parameter is set to **32768*	- if the Cpu parameter is set to 16000.
	//
	// 	- This parameter is set to **65536*	- if the Cpu parameter is set to 8000, 16000, or 32000.
	//
	// 	- This parameter is set to **131072*	- if the Cpu parameter is set to 32000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4096
	Memory        *int32  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	NewSaeVersion *string `json:"NewSaeVersion,omitempty" xml:"NewSaeVersion,omitempty"`
	ResourceType  *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// Scenarios:
	//
	// 	- Web
	//
	// 	- micro_service
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
