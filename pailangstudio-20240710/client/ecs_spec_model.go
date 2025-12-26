// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEcsSpec interface {
  dara.Model
  String() string
  GoString() string
  SetCPU(v int32) *EcsSpec
  GetCPU() *int32 
  SetDriver(v string) *EcsSpec
  GetDriver() *string 
  SetGPU(v int32) *EcsSpec
  GetGPU() *int32 
  SetGPUType(v string) *EcsSpec
  GetGPUType() *string 
  SetInstanceType(v string) *EcsSpec
  GetInstanceType() *string 
  SetMemory(v int32) *EcsSpec
  GetMemory() *int32 
  SetPodCount(v int32) *EcsSpec
  GetPodCount() *int32 
  SetSharedMemory(v int32) *EcsSpec
  GetSharedMemory() *int32 
  SetType(v string) *EcsSpec
  GetType() *string 
}

type EcsSpec struct {
  // CPU核数
  CPU *int32 `json:"CPU,omitempty" xml:"CPU,omitempty"`
  // 驱动版本
  Driver *string `json:"Driver,omitempty" xml:"Driver,omitempty"`
  // GPU卡数
  GPU *int32 `json:"GPU,omitempty" xml:"GPU,omitempty"`
  // GPU类型
  GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
  // 机型名称
  InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
  // 内存大小
  Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
  // 副本数量
  PodCount *int32 `json:"PodCount,omitempty" xml:"PodCount,omitempty"`
  // 共享内存容量
  SharedMemory *int32 `json:"SharedMemory,omitempty" xml:"SharedMemory,omitempty"`
  // 节点类型
  Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s EcsSpec) String() string {
  return dara.Prettify(s)
}

func (s EcsSpec) GoString() string {
  return s.String()
}

func (s *EcsSpec) GetCPU() *int32  {
  return s.CPU
}

func (s *EcsSpec) GetDriver() *string  {
  return s.Driver
}

func (s *EcsSpec) GetGPU() *int32  {
  return s.GPU
}

func (s *EcsSpec) GetGPUType() *string  {
  return s.GPUType
}

func (s *EcsSpec) GetInstanceType() *string  {
  return s.InstanceType
}

func (s *EcsSpec) GetMemory() *int32  {
  return s.Memory
}

func (s *EcsSpec) GetPodCount() *int32  {
  return s.PodCount
}

func (s *EcsSpec) GetSharedMemory() *int32  {
  return s.SharedMemory
}

func (s *EcsSpec) GetType() *string  {
  return s.Type
}

func (s *EcsSpec) SetCPU(v int32) *EcsSpec {
  s.CPU = &v
  return s
}

func (s *EcsSpec) SetDriver(v string) *EcsSpec {
  s.Driver = &v
  return s
}

func (s *EcsSpec) SetGPU(v int32) *EcsSpec {
  s.GPU = &v
  return s
}

func (s *EcsSpec) SetGPUType(v string) *EcsSpec {
  s.GPUType = &v
  return s
}

func (s *EcsSpec) SetInstanceType(v string) *EcsSpec {
  s.InstanceType = &v
  return s
}

func (s *EcsSpec) SetMemory(v int32) *EcsSpec {
  s.Memory = &v
  return s
}

func (s *EcsSpec) SetPodCount(v int32) *EcsSpec {
  s.PodCount = &v
  return s
}

func (s *EcsSpec) SetSharedMemory(v int32) *EcsSpec {
  s.SharedMemory = &v
  return s
}

func (s *EcsSpec) SetType(v string) *EcsSpec {
  s.Type = &v
  return s
}

func (s *EcsSpec) Validate() error {
  return dara.Validate(s)
}

