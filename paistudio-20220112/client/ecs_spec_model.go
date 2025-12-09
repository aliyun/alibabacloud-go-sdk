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
  SetEcsImageId(v string) *EcsSpec
  GetEcsImageId() *string 
  SetEriQuantity(v int32) *EcsSpec
  GetEriQuantity() *int32 
  SetGpu(v int32) *EcsSpec
  GetGpu() *int32 
  SetGpuGUSpec(v string) *EcsSpec
  GetGpuGUSpec() *string 
  SetGpuMemory(v int32) *EcsSpec
  GetGpuMemory() *int32 
  SetGpuType(v string) *EcsSpec
  GetGpuType() *string 
  SetGpuTypeAlias(v string) *EcsSpec
  GetGpuTypeAlias() *string 
  SetInstanceType(v string) *EcsSpec
  GetInstanceType() *string 
  SetMachineModel(v string) *EcsSpec
  GetMachineModel() *string 
  SetMemory(v int32) *EcsSpec
  GetMemory() *int32 
  SetNetworkMode(v string) *EcsSpec
  GetNetworkMode() *string 
  SetPlannedCpu(v int32) *EcsSpec
  GetPlannedCpu() *int32 
  SetPlannedMemory(v int32) *EcsSpec
  GetPlannedMemory() *int32 
  SetResourceType(v string) *EcsSpec
  GetResourceType() *string 
  SetSupportGPUShare(v bool) *EcsSpec
  GetSupportGPUShare() *bool 
  SetSupportRDMA(v bool) *EcsSpec
  GetSupportRDMA() *bool 
  SetSupportSetNetworkCardIndex(v bool) *EcsSpec
  GetSupportSetNetworkCardIndex() *bool 
}

type EcsSpec struct {
  AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
  Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
  EcsImageId *string `json:"EcsImageId,omitempty" xml:"EcsImageId,omitempty"`
  EriQuantity *int32 `json:"EriQuantity,omitempty" xml:"EriQuantity,omitempty"`
  Gpu *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
  GpuGUSpec *string `json:"GpuGUSpec,omitempty" xml:"GpuGUSpec,omitempty"`
  GpuMemory *int32 `json:"GpuMemory,omitempty" xml:"GpuMemory,omitempty"`
  GpuType *string `json:"GpuType,omitempty" xml:"GpuType,omitempty"`
  GpuTypeAlias *string `json:"GpuTypeAlias,omitempty" xml:"GpuTypeAlias,omitempty"`
  InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
  MachineModel *string `json:"MachineModel,omitempty" xml:"MachineModel,omitempty"`
  Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
  NetworkMode *string `json:"NetworkMode,omitempty" xml:"NetworkMode,omitempty"`
  PlannedCpu *int32 `json:"PlannedCpu,omitempty" xml:"PlannedCpu,omitempty"`
  PlannedMemory *int32 `json:"PlannedMemory,omitempty" xml:"PlannedMemory,omitempty"`
  ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
  SupportGPUShare *bool `json:"SupportGPUShare,omitempty" xml:"SupportGPUShare,omitempty"`
  SupportRDMA *bool `json:"SupportRDMA,omitempty" xml:"SupportRDMA,omitempty"`
  SupportSetNetworkCardIndex *bool `json:"SupportSetNetworkCardIndex,omitempty" xml:"SupportSetNetworkCardIndex,omitempty"`
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

func (s *EcsSpec) GetEcsImageId() *string  {
  return s.EcsImageId
}

func (s *EcsSpec) GetEriQuantity() *int32  {
  return s.EriQuantity
}

func (s *EcsSpec) GetGpu() *int32  {
  return s.Gpu
}

func (s *EcsSpec) GetGpuGUSpec() *string  {
  return s.GpuGUSpec
}

func (s *EcsSpec) GetGpuMemory() *int32  {
  return s.GpuMemory
}

func (s *EcsSpec) GetGpuType() *string  {
  return s.GpuType
}

func (s *EcsSpec) GetGpuTypeAlias() *string  {
  return s.GpuTypeAlias
}

func (s *EcsSpec) GetInstanceType() *string  {
  return s.InstanceType
}

func (s *EcsSpec) GetMachineModel() *string  {
  return s.MachineModel
}

func (s *EcsSpec) GetMemory() *int32  {
  return s.Memory
}

func (s *EcsSpec) GetNetworkMode() *string  {
  return s.NetworkMode
}

func (s *EcsSpec) GetPlannedCpu() *int32  {
  return s.PlannedCpu
}

func (s *EcsSpec) GetPlannedMemory() *int32  {
  return s.PlannedMemory
}

func (s *EcsSpec) GetResourceType() *string  {
  return s.ResourceType
}

func (s *EcsSpec) GetSupportGPUShare() *bool  {
  return s.SupportGPUShare
}

func (s *EcsSpec) GetSupportRDMA() *bool  {
  return s.SupportRDMA
}

func (s *EcsSpec) GetSupportSetNetworkCardIndex() *bool  {
  return s.SupportSetNetworkCardIndex
}

func (s *EcsSpec) SetAcceleratorType(v string) *EcsSpec {
  s.AcceleratorType = &v
  return s
}

func (s *EcsSpec) SetCpu(v int32) *EcsSpec {
  s.Cpu = &v
  return s
}

func (s *EcsSpec) SetEcsImageId(v string) *EcsSpec {
  s.EcsImageId = &v
  return s
}

func (s *EcsSpec) SetEriQuantity(v int32) *EcsSpec {
  s.EriQuantity = &v
  return s
}

func (s *EcsSpec) SetGpu(v int32) *EcsSpec {
  s.Gpu = &v
  return s
}

func (s *EcsSpec) SetGpuGUSpec(v string) *EcsSpec {
  s.GpuGUSpec = &v
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

func (s *EcsSpec) SetGpuTypeAlias(v string) *EcsSpec {
  s.GpuTypeAlias = &v
  return s
}

func (s *EcsSpec) SetInstanceType(v string) *EcsSpec {
  s.InstanceType = &v
  return s
}

func (s *EcsSpec) SetMachineModel(v string) *EcsSpec {
  s.MachineModel = &v
  return s
}

func (s *EcsSpec) SetMemory(v int32) *EcsSpec {
  s.Memory = &v
  return s
}

func (s *EcsSpec) SetNetworkMode(v string) *EcsSpec {
  s.NetworkMode = &v
  return s
}

func (s *EcsSpec) SetPlannedCpu(v int32) *EcsSpec {
  s.PlannedCpu = &v
  return s
}

func (s *EcsSpec) SetPlannedMemory(v int32) *EcsSpec {
  s.PlannedMemory = &v
  return s
}

func (s *EcsSpec) SetResourceType(v string) *EcsSpec {
  s.ResourceType = &v
  return s
}

func (s *EcsSpec) SetSupportGPUShare(v bool) *EcsSpec {
  s.SupportGPUShare = &v
  return s
}

func (s *EcsSpec) SetSupportRDMA(v bool) *EcsSpec {
  s.SupportRDMA = &v
  return s
}

func (s *EcsSpec) SetSupportSetNetworkCardIndex(v bool) *EcsSpec {
  s.SupportSetNetworkCardIndex = &v
  return s
}

func (s *EcsSpec) Validate() error {
  return dara.Validate(s)
}

