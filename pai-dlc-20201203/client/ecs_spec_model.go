// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEcsSpec interface {
  dara.Model
  String() string
  GoString() string
  SetAcceleratorType(v string) *EcsSpec
  GetAcceleratorType() *string 
  SetCpu(v int32) *EcsSpec
  GetCpu() *int32 
  SetDefaultGPUDriver(v string) *EcsSpec
  GetDefaultGPUDriver() *string 
  SetGpu(v int32) *EcsSpec
  GetGpu() *int32 
  SetGpuMemory(v int32) *EcsSpec
  GetGpuMemory() *int32 
  SetGpuType(v string) *EcsSpec
  GetGpuType() *string 
  SetInstanceType(v string) *EcsSpec
  GetInstanceType() *string 
  SetIsAvailable(v bool) *EcsSpec
  GetIsAvailable() *bool 
  SetMemory(v int32) *EcsSpec
  GetMemory() *int32 
  SetNonProtectSpotDiscount(v float32) *EcsSpec
  GetNonProtectSpotDiscount() *float32 
  SetPaymentTypes(v []*string) *EcsSpec
  GetPaymentTypes() []*string 
  SetResourceType(v string) *EcsSpec
  GetResourceType() *string 
  SetSpotStockStatus(v string) *EcsSpec
  GetSpotStockStatus() *string 
  SetSupportedGPUDrivers(v []*string) *EcsSpec
  GetSupportedGPUDrivers() []*string 
}

type EcsSpec struct {
  // The accelerator type. Valid values:
  // 
  // 	- CPU
  // 
  // 	- GPU
  // 
  // example:
  // 
  // GPU
  AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
  // The number of CPU cores.
  // 
  // example:
  // 
  // 12
  Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
  // The default GPU driver version.
  // 
  // example:
  // 
  // 470.199.02
  DefaultGPUDriver *string `json:"DefaultGPUDriver,omitempty" xml:"DefaultGPUDriver,omitempty"`
  // The number of GPUs.
  // 
  // example:
  // 
  // 1
  Gpu *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
  // The GPU memory size.
  // 
  // example:
  // 
  // 80
  GpuMemory *int32 `json:"GpuMemory,omitempty" xml:"GpuMemory,omitempty"`
  // The GPU type.
  // 
  // example:
  // 
  // NVIDIA v100
  GpuType *string `json:"GpuType,omitempty" xml:"GpuType,omitempty"`
  // The instance type.
  // 
  // example:
  // 
  // ecs.gn6e-c12g1.3xlarge
  InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
  // Indicates whether the instance type is available. Valid values:
  // 
  // 	- true
  // 
  // 	- false
  // 
  // example:
  // 
  // true
  IsAvailable *bool `json:"IsAvailable,omitempty" xml:"IsAvailable,omitempty"`
  // The memory size. Unit: MiB or GiB.
  // 
  // example:
  // 
  // 92
  Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
  // The discount on the current price of the preemptible instance.
  // 
  // example:
  // 
  // 0.1
  NonProtectSpotDiscount *float32 `json:"NonProtectSpotDiscount,omitempty" xml:"NonProtectSpotDiscount,omitempty"`
  // The billing methods.
  PaymentTypes []*string `json:"PaymentTypes,omitempty" xml:"PaymentTypes,omitempty" type:"Repeated"`
  // The resource type. Valid values:
  // 
  // 	- ECS
  // 
  // 	- Lingjun
  // 
  // example:
  // 
  // ECS
  ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
  // The inventory status of preemptible instance.
  // 
  // example:
  // 
  // WithStock
  SpotStockStatus *string `json:"SpotStockStatus,omitempty" xml:"SpotStockStatus,omitempty"`
  // The GPU driver versions.
  SupportedGPUDrivers []*string `json:"SupportedGPUDrivers,omitempty" xml:"SupportedGPUDrivers,omitempty" type:"Repeated"`
}

func (s EcsSpec) String() string {
  return dara.Prettify(s)
}

func (s EcsSpec) GoString() string {
  return s.String()
}

func (s *EcsSpec) GetAcceleratorType() *string  {
  return s.AcceleratorType
}

func (s *EcsSpec) GetCpu() *int32  {
  return s.Cpu
}

func (s *EcsSpec) GetDefaultGPUDriver() *string  {
  return s.DefaultGPUDriver
}

func (s *EcsSpec) GetGpu() *int32  {
  return s.Gpu
}

func (s *EcsSpec) GetGpuMemory() *int32  {
  return s.GpuMemory
}

func (s *EcsSpec) GetGpuType() *string  {
  return s.GpuType
}

func (s *EcsSpec) GetInstanceType() *string  {
  return s.InstanceType
}

func (s *EcsSpec) GetIsAvailable() *bool  {
  return s.IsAvailable
}

func (s *EcsSpec) GetMemory() *int32  {
  return s.Memory
}

func (s *EcsSpec) GetNonProtectSpotDiscount() *float32  {
  return s.NonProtectSpotDiscount
}

func (s *EcsSpec) GetPaymentTypes() []*string  {
  return s.PaymentTypes
}

func (s *EcsSpec) GetResourceType() *string  {
  return s.ResourceType
}

func (s *EcsSpec) GetSpotStockStatus() *string  {
  return s.SpotStockStatus
}

func (s *EcsSpec) GetSupportedGPUDrivers() []*string  {
  return s.SupportedGPUDrivers
}

func (s *EcsSpec) SetAcceleratorType(v string) *EcsSpec {
  s.AcceleratorType = &v
  return s
}

func (s *EcsSpec) SetCpu(v int32) *EcsSpec {
  s.Cpu = &v
  return s
}

func (s *EcsSpec) SetDefaultGPUDriver(v string) *EcsSpec {
  s.DefaultGPUDriver = &v
  return s
}

func (s *EcsSpec) SetGpu(v int32) *EcsSpec {
  s.Gpu = &v
  return s
}

func (s *EcsSpec) SetGpuMemory(v int32) *EcsSpec {
  s.GpuMemory = &v
  return s
}

func (s *EcsSpec) SetGpuType(v string) *EcsSpec {
  s.GpuType = &v
  return s
}

func (s *EcsSpec) SetInstanceType(v string) *EcsSpec {
  s.InstanceType = &v
  return s
}

func (s *EcsSpec) SetIsAvailable(v bool) *EcsSpec {
  s.IsAvailable = &v
  return s
}

func (s *EcsSpec) SetMemory(v int32) *EcsSpec {
  s.Memory = &v
  return s
}

func (s *EcsSpec) SetNonProtectSpotDiscount(v float32) *EcsSpec {
  s.NonProtectSpotDiscount = &v
  return s
}

func (s *EcsSpec) SetPaymentTypes(v []*string) *EcsSpec {
  s.PaymentTypes = v
  return s
}

func (s *EcsSpec) SetResourceType(v string) *EcsSpec {
  s.ResourceType = &v
  return s
}

func (s *EcsSpec) SetSpotStockStatus(v string) *EcsSpec {
  s.SpotStockStatus = &v
  return s
}

func (s *EcsSpec) SetSupportedGPUDrivers(v []*string) *EcsSpec {
  s.SupportedGPUDrivers = v
  return s
}

func (s *EcsSpec) Validate() error {
  return dara.Validate(s)
}

