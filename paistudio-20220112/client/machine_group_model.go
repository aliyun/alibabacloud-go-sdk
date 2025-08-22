// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMachineGroup interface {
	dara.Model
	String() string
	GoString() string
	SetAllocatableCpu(v int64) *MachineGroup
	GetAllocatableCpu() *int64
	SetAllocatableMemory(v int64) *MachineGroup
	GetAllocatableMemory() *int64
	SetCpu(v int64) *MachineGroup
	GetCpu() *int64
	SetCreatorID(v string) *MachineGroup
	GetCreatorID() *string
	SetDefaultDriver(v string) *MachineGroup
	GetDefaultDriver() *string
	SetDiskCapacity(v int64) *MachineGroup
	GetDiskCapacity() *int64
	SetDiskPL(v string) *MachineGroup
	GetDiskPL() *string
	SetEcsCount(v int64) *MachineGroup
	GetEcsCount() *int64
	SetEcsSpec(v string) *MachineGroup
	GetEcsSpec() *string
	SetGmtCreatedTime(v string) *MachineGroup
	GetGmtCreatedTime() *string
	SetGmtExpiredTime(v string) *MachineGroup
	GetGmtExpiredTime() *string
	SetGmtModifiedTime(v string) *MachineGroup
	GetGmtModifiedTime() *string
	SetGmtStartedTime(v string) *MachineGroup
	GetGmtStartedTime() *string
	SetGpu(v int64) *MachineGroup
	GetGpu() *int64
	SetGpuMemory(v int64) *MachineGroup
	GetGpuMemory() *int64
	SetGpuType(v string) *MachineGroup
	GetGpuType() *string
	SetMachineGroupID(v string) *MachineGroup
	GetMachineGroupID() *string
	SetMemory(v int64) *MachineGroup
	GetMemory() *int64
	SetOrderInstanceId(v string) *MachineGroup
	GetOrderInstanceId() *string
	SetPaymentDuration(v string) *MachineGroup
	GetPaymentDuration() *string
	SetPaymentDurationUnit(v string) *MachineGroup
	GetPaymentDurationUnit() *string
	SetPaymentType(v string) *MachineGroup
	GetPaymentType() *string
	SetReasonCode(v string) *MachineGroup
	GetReasonCode() *string
	SetReasonMessage(v string) *MachineGroup
	GetReasonMessage() *string
	SetResourceGroupID(v string) *MachineGroup
	GetResourceGroupID() *string
	SetResourceType(v string) *MachineGroup
	GetResourceType() *string
	SetStatus(v string) *MachineGroup
	GetStatus() *string
	SetSupportedDrivers(v []*string) *MachineGroup
	GetSupportedDrivers() []*string
	SetSystemReservedCpu(v int64) *MachineGroup
	GetSystemReservedCpu() *int64
	SetSystemReservedMemory(v int64) *MachineGroup
	GetSystemReservedMemory() *int64
}

type MachineGroup struct {
	AllocatableCpu    *int64  `json:"AllocatableCpu,omitempty" xml:"AllocatableCpu,omitempty"`
	AllocatableMemory *int64  `json:"AllocatableMemory,omitempty" xml:"AllocatableMemory,omitempty"`
	Cpu               *int64  `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	CreatorID         *string `json:"CreatorID,omitempty" xml:"CreatorID,omitempty"`
	// example:
	//
	// 470.199.02
	DefaultDriver   *string `json:"DefaultDriver,omitempty" xml:"DefaultDriver,omitempty"`
	DiskCapacity    *int64  `json:"DiskCapacity,omitempty" xml:"DiskCapacity,omitempty"`
	DiskPL          *string `json:"DiskPL,omitempty" xml:"DiskPL,omitempty"`
	EcsCount        *int64  `json:"EcsCount,omitempty" xml:"EcsCount,omitempty"`
	EcsSpec         *string `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	GmtCreatedTime  *string `json:"GmtCreatedTime,omitempty" xml:"GmtCreatedTime,omitempty"`
	GmtExpiredTime  *string `json:"GmtExpiredTime,omitempty" xml:"GmtExpiredTime,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	GmtStartedTime  *string `json:"GmtStartedTime,omitempty" xml:"GmtStartedTime,omitempty"`
	Gpu             *int64  `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	GpuMemory       *int64  `json:"GpuMemory,omitempty" xml:"GpuMemory,omitempty"`
	GpuType         *string `json:"GpuType,omitempty" xml:"GpuType,omitempty"`
	// example:
	//
	// mg1
	MachineGroupID       *string   `json:"MachineGroupID,omitempty" xml:"MachineGroupID,omitempty"`
	Memory               *int64    `json:"Memory,omitempty" xml:"Memory,omitempty"`
	OrderInstanceId      *string   `json:"OrderInstanceId,omitempty" xml:"OrderInstanceId,omitempty"`
	PaymentDuration      *string   `json:"PaymentDuration,omitempty" xml:"PaymentDuration,omitempty"`
	PaymentDurationUnit  *string   `json:"PaymentDurationUnit,omitempty" xml:"PaymentDurationUnit,omitempty"`
	PaymentType          *string   `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	ReasonCode           *string   `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	ReasonMessage        *string   `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	ResourceGroupID      *string   `json:"ResourceGroupID,omitempty" xml:"ResourceGroupID,omitempty"`
	ResourceType         *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Status               *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	SupportedDrivers     []*string `json:"SupportedDrivers,omitempty" xml:"SupportedDrivers,omitempty" type:"Repeated"`
	SystemReservedCpu    *int64    `json:"SystemReservedCpu,omitempty" xml:"SystemReservedCpu,omitempty"`
	SystemReservedMemory *int64    `json:"SystemReservedMemory,omitempty" xml:"SystemReservedMemory,omitempty"`
}

func (s MachineGroup) String() string {
	return dara.Prettify(s)
}

func (s MachineGroup) GoString() string {
	return s.String()
}

func (s *MachineGroup) GetAllocatableCpu() *int64 {
	return s.AllocatableCpu
}

func (s *MachineGroup) GetAllocatableMemory() *int64 {
	return s.AllocatableMemory
}

func (s *MachineGroup) GetCpu() *int64 {
	return s.Cpu
}

func (s *MachineGroup) GetCreatorID() *string {
	return s.CreatorID
}

func (s *MachineGroup) GetDefaultDriver() *string {
	return s.DefaultDriver
}

func (s *MachineGroup) GetDiskCapacity() *int64 {
	return s.DiskCapacity
}

func (s *MachineGroup) GetDiskPL() *string {
	return s.DiskPL
}

func (s *MachineGroup) GetEcsCount() *int64 {
	return s.EcsCount
}

func (s *MachineGroup) GetEcsSpec() *string {
	return s.EcsSpec
}

func (s *MachineGroup) GetGmtCreatedTime() *string {
	return s.GmtCreatedTime
}

func (s *MachineGroup) GetGmtExpiredTime() *string {
	return s.GmtExpiredTime
}

func (s *MachineGroup) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *MachineGroup) GetGmtStartedTime() *string {
	return s.GmtStartedTime
}

func (s *MachineGroup) GetGpu() *int64 {
	return s.Gpu
}

func (s *MachineGroup) GetGpuMemory() *int64 {
	return s.GpuMemory
}

func (s *MachineGroup) GetGpuType() *string {
	return s.GpuType
}

func (s *MachineGroup) GetMachineGroupID() *string {
	return s.MachineGroupID
}

func (s *MachineGroup) GetMemory() *int64 {
	return s.Memory
}

func (s *MachineGroup) GetOrderInstanceId() *string {
	return s.OrderInstanceId
}

func (s *MachineGroup) GetPaymentDuration() *string {
	return s.PaymentDuration
}

func (s *MachineGroup) GetPaymentDurationUnit() *string {
	return s.PaymentDurationUnit
}

func (s *MachineGroup) GetPaymentType() *string {
	return s.PaymentType
}

func (s *MachineGroup) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *MachineGroup) GetReasonMessage() *string {
	return s.ReasonMessage
}

func (s *MachineGroup) GetResourceGroupID() *string {
	return s.ResourceGroupID
}

func (s *MachineGroup) GetResourceType() *string {
	return s.ResourceType
}

func (s *MachineGroup) GetStatus() *string {
	return s.Status
}

func (s *MachineGroup) GetSupportedDrivers() []*string {
	return s.SupportedDrivers
}

func (s *MachineGroup) GetSystemReservedCpu() *int64 {
	return s.SystemReservedCpu
}

func (s *MachineGroup) GetSystemReservedMemory() *int64 {
	return s.SystemReservedMemory
}

func (s *MachineGroup) SetAllocatableCpu(v int64) *MachineGroup {
	s.AllocatableCpu = &v
	return s
}

func (s *MachineGroup) SetAllocatableMemory(v int64) *MachineGroup {
	s.AllocatableMemory = &v
	return s
}

func (s *MachineGroup) SetCpu(v int64) *MachineGroup {
	s.Cpu = &v
	return s
}

func (s *MachineGroup) SetCreatorID(v string) *MachineGroup {
	s.CreatorID = &v
	return s
}

func (s *MachineGroup) SetDefaultDriver(v string) *MachineGroup {
	s.DefaultDriver = &v
	return s
}

func (s *MachineGroup) SetDiskCapacity(v int64) *MachineGroup {
	s.DiskCapacity = &v
	return s
}

func (s *MachineGroup) SetDiskPL(v string) *MachineGroup {
	s.DiskPL = &v
	return s
}

func (s *MachineGroup) SetEcsCount(v int64) *MachineGroup {
	s.EcsCount = &v
	return s
}

func (s *MachineGroup) SetEcsSpec(v string) *MachineGroup {
	s.EcsSpec = &v
	return s
}

func (s *MachineGroup) SetGmtCreatedTime(v string) *MachineGroup {
	s.GmtCreatedTime = &v
	return s
}

func (s *MachineGroup) SetGmtExpiredTime(v string) *MachineGroup {
	s.GmtExpiredTime = &v
	return s
}

func (s *MachineGroup) SetGmtModifiedTime(v string) *MachineGroup {
	s.GmtModifiedTime = &v
	return s
}

func (s *MachineGroup) SetGmtStartedTime(v string) *MachineGroup {
	s.GmtStartedTime = &v
	return s
}

func (s *MachineGroup) SetGpu(v int64) *MachineGroup {
	s.Gpu = &v
	return s
}

func (s *MachineGroup) SetGpuMemory(v int64) *MachineGroup {
	s.GpuMemory = &v
	return s
}

func (s *MachineGroup) SetGpuType(v string) *MachineGroup {
	s.GpuType = &v
	return s
}

func (s *MachineGroup) SetMachineGroupID(v string) *MachineGroup {
	s.MachineGroupID = &v
	return s
}

func (s *MachineGroup) SetMemory(v int64) *MachineGroup {
	s.Memory = &v
	return s
}

func (s *MachineGroup) SetOrderInstanceId(v string) *MachineGroup {
	s.OrderInstanceId = &v
	return s
}

func (s *MachineGroup) SetPaymentDuration(v string) *MachineGroup {
	s.PaymentDuration = &v
	return s
}

func (s *MachineGroup) SetPaymentDurationUnit(v string) *MachineGroup {
	s.PaymentDurationUnit = &v
	return s
}

func (s *MachineGroup) SetPaymentType(v string) *MachineGroup {
	s.PaymentType = &v
	return s
}

func (s *MachineGroup) SetReasonCode(v string) *MachineGroup {
	s.ReasonCode = &v
	return s
}

func (s *MachineGroup) SetReasonMessage(v string) *MachineGroup {
	s.ReasonMessage = &v
	return s
}

func (s *MachineGroup) SetResourceGroupID(v string) *MachineGroup {
	s.ResourceGroupID = &v
	return s
}

func (s *MachineGroup) SetResourceType(v string) *MachineGroup {
	s.ResourceType = &v
	return s
}

func (s *MachineGroup) SetStatus(v string) *MachineGroup {
	s.Status = &v
	return s
}

func (s *MachineGroup) SetSupportedDrivers(v []*string) *MachineGroup {
	s.SupportedDrivers = v
	return s
}

func (s *MachineGroup) SetSystemReservedCpu(v int64) *MachineGroup {
	s.SystemReservedCpu = &v
	return s
}

func (s *MachineGroup) SetSystemReservedMemory(v int64) *MachineGroup {
	s.SystemReservedMemory = &v
	return s
}

func (s *MachineGroup) Validate() error {
	return dara.Validate(s)
}
