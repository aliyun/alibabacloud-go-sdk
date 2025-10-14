// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEcsSpecsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListEcsSpecsResponseBody
	GetCode() *string
	SetEcsSpecs(v []*ListEcsSpecsResponseBodyEcsSpecs) *ListEcsSpecsResponseBody
	GetEcsSpecs() []*ListEcsSpecsResponseBodyEcsSpecs
	SetHttpStatusCode(v int32) *ListEcsSpecsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListEcsSpecsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListEcsSpecsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListEcsSpecsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListEcsSpecsResponseBody
	GetTotalCount() *int64
}

type ListEcsSpecsResponseBody struct {
	// The status code. Valid values:
	//
	// 	- InternalError: an internal error. All errors, except for parameter validation errors, are classified as internal errors.
	//
	// 	- ValidationError: a parameter validation error.
	//
	// example:
	//
	// null
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The specifications of the ECS instances returned on this page.
	EcsSpecs []*ListEcsSpecsResponseBodyEcsSpecs `json:"EcsSpecs,omitempty" xml:"EcsSpecs,omitempty" type:"Repeated"`
	// The HTTP status code. Valid values:
	//
	// 	- 400
	//
	// 	- 404
	//
	// example:
	//
	// null
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// "XXX"
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E7D55162-4489-1619-AAF5-3F97D5FCA948
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of ECS instances.
	//
	// example:
	//
	// 35
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEcsSpecsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEcsSpecsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEcsSpecsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListEcsSpecsResponseBody) GetEcsSpecs() []*ListEcsSpecsResponseBodyEcsSpecs {
	return s.EcsSpecs
}

func (s *ListEcsSpecsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListEcsSpecsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListEcsSpecsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEcsSpecsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListEcsSpecsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListEcsSpecsResponseBody) SetCode(v string) *ListEcsSpecsResponseBody {
	s.Code = &v
	return s
}

func (s *ListEcsSpecsResponseBody) SetEcsSpecs(v []*ListEcsSpecsResponseBodyEcsSpecs) *ListEcsSpecsResponseBody {
	s.EcsSpecs = v
	return s
}

func (s *ListEcsSpecsResponseBody) SetHttpStatusCode(v int32) *ListEcsSpecsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListEcsSpecsResponseBody) SetMessage(v string) *ListEcsSpecsResponseBody {
	s.Message = &v
	return s
}

func (s *ListEcsSpecsResponseBody) SetRequestId(v string) *ListEcsSpecsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEcsSpecsResponseBody) SetSuccess(v bool) *ListEcsSpecsResponseBody {
	s.Success = &v
	return s
}

func (s *ListEcsSpecsResponseBody) SetTotalCount(v int64) *ListEcsSpecsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListEcsSpecsResponseBody) Validate() error {
	if s.EcsSpecs != nil {
		for _, item := range s.EcsSpecs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEcsSpecsResponseBodyEcsSpecs struct {
	// The accelerator type.
	//
	// example:
	//
	// CPU
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	// The number of vCPUs.
	//
	// example:
	//
	// 32
	CPU *int64 `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// The currency unit.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The number of GPUs.
	//
	// example:
	//
	// 4
	GPU           *int64   `json:"GPU,omitempty" xml:"GPU,omitempty"`
	GPUMemorySize *float32 `json:"GPUMemorySize,omitempty" xml:"GPUMemorySize,omitempty"`
	// The GPU type. Valid values:
	//
	// 	- V100
	//
	// 	- A100
	//
	// 	- A10
	//
	// 	- T4
	//
	// 	- P100
	//
	// example:
	//
	// v100
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// The inbound bandwidth of the instance.
	//
	// example:
	//
	// 5120000
	InstanceBandwidthRx *int64 `json:"InstanceBandwidthRx,omitempty" xml:"InstanceBandwidthRx,omitempty"`
	// The instance type.
	//
	// example:
	//
	// ecs.gn5-c28g1.7xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Indicates whether the resource was available.
	//
	// example:
	//
	// True
	IsAvailable *bool `json:"IsAvailable,omitempty" xml:"IsAvailable,omitempty"`
	// The labels of the ECS specification.
	//
	// example:
	//
	// {\\"foo\\": \\"bar\\"}
	Labels []*ListEcsSpecsResponseBodyEcsSpecsLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The memory size. Unit: GB.
	//
	// example:
	//
	// 32
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The price.
	//
	// example:
	//
	// 22.8
	Price           *float64 `json:"Price,omitempty" xml:"Price,omitempty"`
	SpotStockStatus *string  `json:"SpotStockStatus,omitempty" xml:"SpotStockStatus,omitempty"`
	// The size of the system disk. Unit: GB.
	//
	// example:
	//
	// 500
	SystemDiskCapacity *int64 `json:"SystemDiskCapacity,omitempty" xml:"SystemDiskCapacity,omitempty"`
}

func (s ListEcsSpecsResponseBodyEcsSpecs) String() string {
	return dara.Prettify(s)
}

func (s ListEcsSpecsResponseBodyEcsSpecs) GoString() string {
	return s.String()
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) GetAcceleratorType() *string {
	return s.AcceleratorType
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) GetCPU() *int64 {
	return s.CPU
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) GetCurrency() *string {
	return s.Currency
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) GetGPU() *int64 {
	return s.GPU
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) GetGPUMemorySize() *float32 {
	return s.GPUMemorySize
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) GetGPUType() *string {
	return s.GPUType
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) GetInstanceBandwidthRx() *int64 {
	return s.InstanceBandwidthRx
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) GetIsAvailable() *bool {
	return s.IsAvailable
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) GetLabels() []*ListEcsSpecsResponseBodyEcsSpecsLabels {
	return s.Labels
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) GetMemory() *float32 {
	return s.Memory
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) GetPrice() *float64 {
	return s.Price
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) GetSpotStockStatus() *string {
	return s.SpotStockStatus
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) GetSystemDiskCapacity() *int64 {
	return s.SystemDiskCapacity
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) SetAcceleratorType(v string) *ListEcsSpecsResponseBodyEcsSpecs {
	s.AcceleratorType = &v
	return s
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) SetCPU(v int64) *ListEcsSpecsResponseBodyEcsSpecs {
	s.CPU = &v
	return s
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) SetCurrency(v string) *ListEcsSpecsResponseBodyEcsSpecs {
	s.Currency = &v
	return s
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) SetGPU(v int64) *ListEcsSpecsResponseBodyEcsSpecs {
	s.GPU = &v
	return s
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) SetGPUMemorySize(v float32) *ListEcsSpecsResponseBodyEcsSpecs {
	s.GPUMemorySize = &v
	return s
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) SetGPUType(v string) *ListEcsSpecsResponseBodyEcsSpecs {
	s.GPUType = &v
	return s
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) SetInstanceBandwidthRx(v int64) *ListEcsSpecsResponseBodyEcsSpecs {
	s.InstanceBandwidthRx = &v
	return s
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) SetInstanceType(v string) *ListEcsSpecsResponseBodyEcsSpecs {
	s.InstanceType = &v
	return s
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) SetIsAvailable(v bool) *ListEcsSpecsResponseBodyEcsSpecs {
	s.IsAvailable = &v
	return s
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) SetLabels(v []*ListEcsSpecsResponseBodyEcsSpecsLabels) *ListEcsSpecsResponseBodyEcsSpecs {
	s.Labels = v
	return s
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) SetMemory(v float32) *ListEcsSpecsResponseBodyEcsSpecs {
	s.Memory = &v
	return s
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) SetPrice(v float64) *ListEcsSpecsResponseBodyEcsSpecs {
	s.Price = &v
	return s
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) SetSpotStockStatus(v string) *ListEcsSpecsResponseBodyEcsSpecs {
	s.SpotStockStatus = &v
	return s
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) SetSystemDiskCapacity(v int64) *ListEcsSpecsResponseBodyEcsSpecs {
	s.SystemDiskCapacity = &v
	return s
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEcsSpecsResponseBodyEcsSpecsLabels struct {
	// The label key added to the ECS specification.
	//
	// example:
	//
	// SupportResourcePackDeduction
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The label value added to the ECS specification.
	//
	// example:
	//
	// true
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEcsSpecsResponseBodyEcsSpecsLabels) String() string {
	return dara.Prettify(s)
}

func (s ListEcsSpecsResponseBodyEcsSpecsLabels) GoString() string {
	return s.String()
}

func (s *ListEcsSpecsResponseBodyEcsSpecsLabels) GetKey() *string {
	return s.Key
}

func (s *ListEcsSpecsResponseBodyEcsSpecsLabels) GetValue() *string {
	return s.Value
}

func (s *ListEcsSpecsResponseBodyEcsSpecsLabels) SetKey(v string) *ListEcsSpecsResponseBodyEcsSpecsLabels {
	s.Key = &v
	return s
}

func (s *ListEcsSpecsResponseBodyEcsSpecsLabels) SetValue(v string) *ListEcsSpecsResponseBodyEcsSpecsLabels {
	s.Value = &v
	return s
}

func (s *ListEcsSpecsResponseBodyEcsSpecsLabels) Validate() error {
	return dara.Validate(s)
}
