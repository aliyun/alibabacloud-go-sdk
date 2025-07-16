// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMachineSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceMetas(v []*DescribeMachineSpecResponseBodyInstanceMetas) *DescribeMachineSpecResponseBody
	GetInstanceMetas() []*DescribeMachineSpecResponseBodyInstanceMetas
	SetRequestId(v string) *DescribeMachineSpecResponseBody
	GetRequestId() *string
	SetTypes(v []*DescribeMachineSpecResponseBodyTypes) *DescribeMachineSpecResponseBody
	GetTypes() []*DescribeMachineSpecResponseBodyTypes
}

type DescribeMachineSpecResponseBody struct {
	// The instance types when the resources are specified.
	InstanceMetas []*DescribeMachineSpecResponseBodyInstanceMetas `json:"InstanceMetas,omitempty" xml:"InstanceMetas,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The values that can be supported when the number of CPUs and memory size are specified for deployment.
	Types []*DescribeMachineSpecResponseBodyTypes `json:"Types,omitempty" xml:"Types,omitempty" type:"Repeated"`
}

func (s DescribeMachineSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMachineSpecResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMachineSpecResponseBody) GetInstanceMetas() []*DescribeMachineSpecResponseBodyInstanceMetas {
	return s.InstanceMetas
}

func (s *DescribeMachineSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMachineSpecResponseBody) GetTypes() []*DescribeMachineSpecResponseBodyTypes {
	return s.Types
}

func (s *DescribeMachineSpecResponseBody) SetInstanceMetas(v []*DescribeMachineSpecResponseBodyInstanceMetas) *DescribeMachineSpecResponseBody {
	s.InstanceMetas = v
	return s
}

func (s *DescribeMachineSpecResponseBody) SetRequestId(v string) *DescribeMachineSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMachineSpecResponseBody) SetTypes(v []*DescribeMachineSpecResponseBodyTypes) *DescribeMachineSpecResponseBody {
	s.Types = v
	return s
}

func (s *DescribeMachineSpecResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeMachineSpecResponseBodyInstanceMetas struct {
	// The number of CPU cores in the instance type.
	//
	// example:
	//
	// 32
	CPU *int32 `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// The GPU type in the instance type. If the instance type is not a GPU-based instance type, this parameter does not exist.
	//
	// example:
	//
	// GU30
	GPU *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	// The number of GPUs in the instance type.
	//
	// example:
	//
	// 1
	GPUAmount *int32 `json:"GPUAmount,omitempty" xml:"GPUAmount,omitempty"`
	// The GPU memory in the instance type. Unit: GB.
	//
	// example:
	//
	// 24
	GPUMemory *float32 `json:"GPUMemory,omitempty" xml:"GPUMemory,omitempty"`
	// The name of the instance type.
	//
	// example:
	//
	// ml.gu7i.c32m188.1-gu30
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Indicates whether the instance type is available.
	//
	// example:
	//
	// true
	IsAvailable *bool `json:"IsAvailable,omitempty" xml:"IsAvailable,omitempty"`
	// The memory size in the instance type. Unit: GB.
	//
	// example:
	//
	// 188
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The minimum discount that can be accepted when the preemptible instance type does not include a usage duration. 0.1 indicates one fold. If this parameter is not returned, the bidding feature is not supported.
	//
	// example:
	//
	// 0.1
	NonProtectSpotDiscount *float32 `json:"NonProtectSpotDiscount,omitempty" xml:"NonProtectSpotDiscount,omitempty"`
	// The minimum discount that can be accepted when the preemptible instance type has the 1-hour protection duration. 0.1 indicates one fold. If this parameter is not returned, the bidding feature is not supported.
	//
	// example:
	//
	// 0.12
	SpotDiscount *float32 `json:"SpotDiscount,omitempty" xml:"SpotDiscount,omitempty"`
	// The inventory status of the instance type.
	//
	// Valid values:
	//
	// 	- WithStock
	//
	// 	- ClosedWithStock
	//
	// 	- NoStock
	//
	// example:
	//
	// WithStock
	StockStatus *string `json:"StockStatus,omitempty" xml:"StockStatus,omitempty"`
	// The source of the instance type.
	//
	// Valid values:
	//
	// 	- ECS
	//
	// 	- BareMetal
	//
	// 	- Lingjun
	//
	// example:
	//
	// ECS
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s DescribeMachineSpecResponseBodyInstanceMetas) String() string {
	return dara.Prettify(s)
}

func (s DescribeMachineSpecResponseBodyInstanceMetas) GoString() string {
	return s.String()
}

func (s *DescribeMachineSpecResponseBodyInstanceMetas) GetCPU() *int32 {
	return s.CPU
}

func (s *DescribeMachineSpecResponseBodyInstanceMetas) GetGPU() *string {
	return s.GPU
}

func (s *DescribeMachineSpecResponseBodyInstanceMetas) GetGPUAmount() *int32 {
	return s.GPUAmount
}

func (s *DescribeMachineSpecResponseBodyInstanceMetas) GetGPUMemory() *float32 {
	return s.GPUMemory
}

func (s *DescribeMachineSpecResponseBodyInstanceMetas) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeMachineSpecResponseBodyInstanceMetas) GetIsAvailable() *bool {
	return s.IsAvailable
}

func (s *DescribeMachineSpecResponseBodyInstanceMetas) GetMemory() *float32 {
	return s.Memory
}

func (s *DescribeMachineSpecResponseBodyInstanceMetas) GetNonProtectSpotDiscount() *float32 {
	return s.NonProtectSpotDiscount
}

func (s *DescribeMachineSpecResponseBodyInstanceMetas) GetSpotDiscount() *float32 {
	return s.SpotDiscount
}

func (s *DescribeMachineSpecResponseBodyInstanceMetas) GetStockStatus() *string {
	return s.StockStatus
}

func (s *DescribeMachineSpecResponseBodyInstanceMetas) GetVendor() *string {
	return s.Vendor
}

func (s *DescribeMachineSpecResponseBodyInstanceMetas) SetCPU(v int32) *DescribeMachineSpecResponseBodyInstanceMetas {
	s.CPU = &v
	return s
}

func (s *DescribeMachineSpecResponseBodyInstanceMetas) SetGPU(v string) *DescribeMachineSpecResponseBodyInstanceMetas {
	s.GPU = &v
	return s
}

func (s *DescribeMachineSpecResponseBodyInstanceMetas) SetGPUAmount(v int32) *DescribeMachineSpecResponseBodyInstanceMetas {
	s.GPUAmount = &v
	return s
}

func (s *DescribeMachineSpecResponseBodyInstanceMetas) SetGPUMemory(v float32) *DescribeMachineSpecResponseBodyInstanceMetas {
	s.GPUMemory = &v
	return s
}

func (s *DescribeMachineSpecResponseBodyInstanceMetas) SetInstanceType(v string) *DescribeMachineSpecResponseBodyInstanceMetas {
	s.InstanceType = &v
	return s
}

func (s *DescribeMachineSpecResponseBodyInstanceMetas) SetIsAvailable(v bool) *DescribeMachineSpecResponseBodyInstanceMetas {
	s.IsAvailable = &v
	return s
}

func (s *DescribeMachineSpecResponseBodyInstanceMetas) SetMemory(v float32) *DescribeMachineSpecResponseBodyInstanceMetas {
	s.Memory = &v
	return s
}

func (s *DescribeMachineSpecResponseBodyInstanceMetas) SetNonProtectSpotDiscount(v float32) *DescribeMachineSpecResponseBodyInstanceMetas {
	s.NonProtectSpotDiscount = &v
	return s
}

func (s *DescribeMachineSpecResponseBodyInstanceMetas) SetSpotDiscount(v float32) *DescribeMachineSpecResponseBodyInstanceMetas {
	s.SpotDiscount = &v
	return s
}

func (s *DescribeMachineSpecResponseBodyInstanceMetas) SetStockStatus(v string) *DescribeMachineSpecResponseBodyInstanceMetas {
	s.StockStatus = &v
	return s
}

func (s *DescribeMachineSpecResponseBodyInstanceMetas) SetVendor(v string) *DescribeMachineSpecResponseBodyInstanceMetas {
	s.Vendor = &v
	return s
}

func (s *DescribeMachineSpecResponseBodyInstanceMetas) Validate() error {
	return dara.Validate(s)
}

type DescribeMachineSpecResponseBodyTypes struct {
	// Valid values:
	//
	// example:
	//
	// 1
	CPU *int32 `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// The optional values for memory when CPU is set to a specific value as above.
	Memory []*int32 `json:"Memory,omitempty" xml:"Memory,omitempty" type:"Repeated"`
}

func (s DescribeMachineSpecResponseBodyTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeMachineSpecResponseBodyTypes) GoString() string {
	return s.String()
}

func (s *DescribeMachineSpecResponseBodyTypes) GetCPU() *int32 {
	return s.CPU
}

func (s *DescribeMachineSpecResponseBodyTypes) GetMemory() []*int32 {
	return s.Memory
}

func (s *DescribeMachineSpecResponseBodyTypes) SetCPU(v int32) *DescribeMachineSpecResponseBodyTypes {
	s.CPU = &v
	return s
}

func (s *DescribeMachineSpecResponseBodyTypes) SetMemory(v []*int32) *DescribeMachineSpecResponseBodyTypes {
	s.Memory = v
	return s
}

func (s *DescribeMachineSpecResponseBodyTypes) Validate() error {
	return dara.Validate(s)
}
