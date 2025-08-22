// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceGroupMachineGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllocatableCpu(v string) *GetResourceGroupMachineGroupResponseBody
	GetAllocatableCpu() *string
	SetAllocatableMemory(v string) *GetResourceGroupMachineGroupResponseBody
	GetAllocatableMemory() *string
	SetCpu(v string) *GetResourceGroupMachineGroupResponseBody
	GetCpu() *string
	SetDefaultDriver(v string) *GetResourceGroupMachineGroupResponseBody
	GetDefaultDriver() *string
	SetEcsCount(v int64) *GetResourceGroupMachineGroupResponseBody
	GetEcsCount() *int64
	SetEcsSpec(v string) *GetResourceGroupMachineGroupResponseBody
	GetEcsSpec() *string
	SetGmtCreatedTime(v string) *GetResourceGroupMachineGroupResponseBody
	GetGmtCreatedTime() *string
	SetGmtExpiredTime(v string) *GetResourceGroupMachineGroupResponseBody
	GetGmtExpiredTime() *string
	SetGmtModifiedTime(v string) *GetResourceGroupMachineGroupResponseBody
	GetGmtModifiedTime() *string
	SetGmtStartedTime(v string) *GetResourceGroupMachineGroupResponseBody
	GetGmtStartedTime() *string
	SetGpu(v string) *GetResourceGroupMachineGroupResponseBody
	GetGpu() *string
	SetGpuType(v string) *GetResourceGroupMachineGroupResponseBody
	GetGpuType() *string
	SetMachineGroupID(v string) *GetResourceGroupMachineGroupResponseBody
	GetMachineGroupID() *string
	SetMemory(v string) *GetResourceGroupMachineGroupResponseBody
	GetMemory() *string
	SetName(v string) *GetResourceGroupMachineGroupResponseBody
	GetName() *string
	SetPaymentDuration(v string) *GetResourceGroupMachineGroupResponseBody
	GetPaymentDuration() *string
	SetPaymentDurationUnit(v string) *GetResourceGroupMachineGroupResponseBody
	GetPaymentDurationUnit() *string
	SetPaymentType(v string) *GetResourceGroupMachineGroupResponseBody
	GetPaymentType() *string
	SetRequestId(v string) *GetResourceGroupMachineGroupResponseBody
	GetRequestId() *string
	SetResourceGroupID(v string) *GetResourceGroupMachineGroupResponseBody
	GetResourceGroupID() *string
	SetStatus(v string) *GetResourceGroupMachineGroupResponseBody
	GetStatus() *string
	SetSupportedDrivers(v []*string) *GetResourceGroupMachineGroupResponseBody
	GetSupportedDrivers() []*string
	SetSystemReservedCpu(v string) *GetResourceGroupMachineGroupResponseBody
	GetSystemReservedCpu() *string
	SetSystemReservedMemory(v string) *GetResourceGroupMachineGroupResponseBody
	GetSystemReservedMemory() *string
	SetTags(v []*GetResourceGroupMachineGroupResponseBodyTags) *GetResourceGroupMachineGroupResponseBody
	GetTags() []*GetResourceGroupMachineGroupResponseBodyTags
}

type GetResourceGroupMachineGroupResponseBody struct {
	AllocatableCpu    *string `json:"AllocatableCpu,omitempty" xml:"AllocatableCpu,omitempty"`
	AllocatableMemory *string `json:"AllocatableMemory,omitempty" xml:"AllocatableMemory,omitempty"`
	// example:
	//
	// 2
	Cpu *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// example:
	//
	// 535
	DefaultDriver *string `json:"DefaultDriver,omitempty" xml:"DefaultDriver,omitempty"`
	// example:
	//
	// 1
	EcsCount *int64 `json:"EcsCount,omitempty" xml:"EcsCount,omitempty"`
	// example:
	//
	// ecs.c6.large
	EcsSpec *string `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	// example:
	//
	// 2023-06-22T00:00:00Z
	GmtCreatedTime *string `json:"GmtCreatedTime,omitempty" xml:"GmtCreatedTime,omitempty"`
	// example:
	//
	// 2023-06-22T00:00:00Z
	GmtExpiredTime *string `json:"GmtExpiredTime,omitempty" xml:"GmtExpiredTime,omitempty"`
	// example:
	//
	// 2023-06-22T00:00:00Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// 2023-06-22T00:00:00Z
	GmtStartedTime *string `json:"GmtStartedTime,omitempty" xml:"GmtStartedTime,omitempty"`
	// example:
	//
	// 8
	Gpu *string `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// example:
	//
	// A100
	GpuType *string `json:"GpuType,omitempty" xml:"GpuType,omitempty"`
	// example:
	//
	// mgmioirqjgw6c5lg
	MachineGroupID *string `json:"MachineGroupID,omitempty" xml:"MachineGroupID,omitempty"`
	// example:
	//
	// 64
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// example:
	//
	// testMachineGroup
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	PaymentDuration *string `json:"PaymentDuration,omitempty" xml:"PaymentDuration,omitempty"`
	// example:
	//
	// Month
	PaymentDurationUnit *string `json:"PaymentDurationUnit,omitempty" xml:"PaymentDurationUnit,omitempty"`
	// example:
	//
	// PREPAY
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// example:
	//
	// 18D5A1C6-14B8-545E-8408-0A7DDB4C6B5E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// rgf0zhfqn1d4ity2
	ResourceGroupID *string `json:"ResourceGroupID,omitempty" xml:"ResourceGroupID,omitempty"`
	// example:
	//
	// Ready
	Status               *string                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	SupportedDrivers     []*string                                       `json:"SupportedDrivers,omitempty" xml:"SupportedDrivers,omitempty" type:"Repeated"`
	SystemReservedCpu    *string                                         `json:"SystemReservedCpu,omitempty" xml:"SystemReservedCpu,omitempty"`
	SystemReservedMemory *string                                         `json:"SystemReservedMemory,omitempty" xml:"SystemReservedMemory,omitempty"`
	Tags                 []*GetResourceGroupMachineGroupResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s GetResourceGroupMachineGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupMachineGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceGroupMachineGroupResponseBody) GetAllocatableCpu() *string {
	return s.AllocatableCpu
}

func (s *GetResourceGroupMachineGroupResponseBody) GetAllocatableMemory() *string {
	return s.AllocatableMemory
}

func (s *GetResourceGroupMachineGroupResponseBody) GetCpu() *string {
	return s.Cpu
}

func (s *GetResourceGroupMachineGroupResponseBody) GetDefaultDriver() *string {
	return s.DefaultDriver
}

func (s *GetResourceGroupMachineGroupResponseBody) GetEcsCount() *int64 {
	return s.EcsCount
}

func (s *GetResourceGroupMachineGroupResponseBody) GetEcsSpec() *string {
	return s.EcsSpec
}

func (s *GetResourceGroupMachineGroupResponseBody) GetGmtCreatedTime() *string {
	return s.GmtCreatedTime
}

func (s *GetResourceGroupMachineGroupResponseBody) GetGmtExpiredTime() *string {
	return s.GmtExpiredTime
}

func (s *GetResourceGroupMachineGroupResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetResourceGroupMachineGroupResponseBody) GetGmtStartedTime() *string {
	return s.GmtStartedTime
}

func (s *GetResourceGroupMachineGroupResponseBody) GetGpu() *string {
	return s.Gpu
}

func (s *GetResourceGroupMachineGroupResponseBody) GetGpuType() *string {
	return s.GpuType
}

func (s *GetResourceGroupMachineGroupResponseBody) GetMachineGroupID() *string {
	return s.MachineGroupID
}

func (s *GetResourceGroupMachineGroupResponseBody) GetMemory() *string {
	return s.Memory
}

func (s *GetResourceGroupMachineGroupResponseBody) GetName() *string {
	return s.Name
}

func (s *GetResourceGroupMachineGroupResponseBody) GetPaymentDuration() *string {
	return s.PaymentDuration
}

func (s *GetResourceGroupMachineGroupResponseBody) GetPaymentDurationUnit() *string {
	return s.PaymentDurationUnit
}

func (s *GetResourceGroupMachineGroupResponseBody) GetPaymentType() *string {
	return s.PaymentType
}

func (s *GetResourceGroupMachineGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceGroupMachineGroupResponseBody) GetResourceGroupID() *string {
	return s.ResourceGroupID
}

func (s *GetResourceGroupMachineGroupResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetResourceGroupMachineGroupResponseBody) GetSupportedDrivers() []*string {
	return s.SupportedDrivers
}

func (s *GetResourceGroupMachineGroupResponseBody) GetSystemReservedCpu() *string {
	return s.SystemReservedCpu
}

func (s *GetResourceGroupMachineGroupResponseBody) GetSystemReservedMemory() *string {
	return s.SystemReservedMemory
}

func (s *GetResourceGroupMachineGroupResponseBody) GetTags() []*GetResourceGroupMachineGroupResponseBodyTags {
	return s.Tags
}

func (s *GetResourceGroupMachineGroupResponseBody) SetAllocatableCpu(v string) *GetResourceGroupMachineGroupResponseBody {
	s.AllocatableCpu = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetAllocatableMemory(v string) *GetResourceGroupMachineGroupResponseBody {
	s.AllocatableMemory = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetCpu(v string) *GetResourceGroupMachineGroupResponseBody {
	s.Cpu = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetDefaultDriver(v string) *GetResourceGroupMachineGroupResponseBody {
	s.DefaultDriver = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetEcsCount(v int64) *GetResourceGroupMachineGroupResponseBody {
	s.EcsCount = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetEcsSpec(v string) *GetResourceGroupMachineGroupResponseBody {
	s.EcsSpec = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetGmtCreatedTime(v string) *GetResourceGroupMachineGroupResponseBody {
	s.GmtCreatedTime = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetGmtExpiredTime(v string) *GetResourceGroupMachineGroupResponseBody {
	s.GmtExpiredTime = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetGmtModifiedTime(v string) *GetResourceGroupMachineGroupResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetGmtStartedTime(v string) *GetResourceGroupMachineGroupResponseBody {
	s.GmtStartedTime = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetGpu(v string) *GetResourceGroupMachineGroupResponseBody {
	s.Gpu = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetGpuType(v string) *GetResourceGroupMachineGroupResponseBody {
	s.GpuType = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetMachineGroupID(v string) *GetResourceGroupMachineGroupResponseBody {
	s.MachineGroupID = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetMemory(v string) *GetResourceGroupMachineGroupResponseBody {
	s.Memory = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetName(v string) *GetResourceGroupMachineGroupResponseBody {
	s.Name = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetPaymentDuration(v string) *GetResourceGroupMachineGroupResponseBody {
	s.PaymentDuration = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetPaymentDurationUnit(v string) *GetResourceGroupMachineGroupResponseBody {
	s.PaymentDurationUnit = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetPaymentType(v string) *GetResourceGroupMachineGroupResponseBody {
	s.PaymentType = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetRequestId(v string) *GetResourceGroupMachineGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetResourceGroupID(v string) *GetResourceGroupMachineGroupResponseBody {
	s.ResourceGroupID = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetStatus(v string) *GetResourceGroupMachineGroupResponseBody {
	s.Status = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetSupportedDrivers(v []*string) *GetResourceGroupMachineGroupResponseBody {
	s.SupportedDrivers = v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetSystemReservedCpu(v string) *GetResourceGroupMachineGroupResponseBody {
	s.SystemReservedCpu = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetSystemReservedMemory(v string) *GetResourceGroupMachineGroupResponseBody {
	s.SystemReservedMemory = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetTags(v []*GetResourceGroupMachineGroupResponseBodyTags) *GetResourceGroupMachineGroupResponseBody {
	s.Tags = v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetResourceGroupMachineGroupResponseBodyTags struct {
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// example:
	//
	// TestValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetResourceGroupMachineGroupResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupMachineGroupResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetResourceGroupMachineGroupResponseBodyTags) GetTagKey() *string {
	return s.TagKey
}

func (s *GetResourceGroupMachineGroupResponseBodyTags) GetTagValue() *string {
	return s.TagValue
}

func (s *GetResourceGroupMachineGroupResponseBodyTags) SetTagKey(v string) *GetResourceGroupMachineGroupResponseBodyTags {
	s.TagKey = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBodyTags) SetTagValue(v string) *GetResourceGroupMachineGroupResponseBodyTags {
	s.TagValue = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
