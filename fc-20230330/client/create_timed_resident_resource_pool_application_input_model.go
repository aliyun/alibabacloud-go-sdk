// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTimedResidentResourcePoolApplicationInput interface {
	dara.Model
	String() string
	GoString() string
	SetAssociatedPoolId(v string) *CreateTimedResidentResourcePoolApplicationInput
	GetAssociatedPoolId() *string
	SetDiskSizeInGB(v int64) *CreateTimedResidentResourcePoolApplicationInput
	GetDiskSizeInGB() *int64
	SetGpuType(v string) *CreateTimedResidentResourcePoolApplicationInput
	GetGpuType() *string
	SetMemorySizeInGB(v int64) *CreateTimedResidentResourcePoolApplicationInput
	GetMemorySizeInGB() *int64
	SetOperationType(v string) *CreateTimedResidentResourcePoolApplicationInput
	GetOperationType() *string
	SetPoolName(v string) *CreateTimedResidentResourcePoolApplicationInput
	GetPoolName() *string
	SetReason(v string) *CreateTimedResidentResourcePoolApplicationInput
	GetReason() *string
	SetTimedConfig(v *TimedPoolConfig) *CreateTimedResidentResourcePoolApplicationInput
	GetTimedConfig() *TimedPoolConfig
	SetTimedPoolId(v string) *CreateTimedResidentResourcePoolApplicationInput
	GetTimedPoolId() *string
	SetTotalGPUCards(v int64) *CreateTimedResidentResourcePoolApplicationInput
	GetTotalGPUCards() *int64
	SetVCpuCores(v int64) *CreateTimedResidentResourcePoolApplicationInput
	GetVCpuCores() *int64
}

type CreateTimedResidentResourcePoolApplicationInput struct {
	AssociatedPoolId *string          `json:"associatedPoolId,omitempty" xml:"associatedPoolId,omitempty"`
	DiskSizeInGB     *int64           `json:"diskSizeInGB,omitempty" xml:"diskSizeInGB,omitempty"`
	GpuType          *string          `json:"gpuType,omitempty" xml:"gpuType,omitempty"`
	MemorySizeInGB   *int64           `json:"memorySizeInGB,omitempty" xml:"memorySizeInGB,omitempty"`
	OperationType    *string          `json:"operationType,omitempty" xml:"operationType,omitempty"`
	PoolName         *string          `json:"poolName,omitempty" xml:"poolName,omitempty"`
	Reason           *string          `json:"reason,omitempty" xml:"reason,omitempty"`
	TimedConfig      *TimedPoolConfig `json:"timedConfig,omitempty" xml:"timedConfig,omitempty"`
	TimedPoolId      *string          `json:"timedPoolId,omitempty" xml:"timedPoolId,omitempty"`
	TotalGPUCards    *int64           `json:"totalGPUCards,omitempty" xml:"totalGPUCards,omitempty"`
	VCpuCores        *int64           `json:"vCpuCores,omitempty" xml:"vCpuCores,omitempty"`
}

func (s CreateTimedResidentResourcePoolApplicationInput) String() string {
	return dara.Prettify(s)
}

func (s CreateTimedResidentResourcePoolApplicationInput) GoString() string {
	return s.String()
}

func (s *CreateTimedResidentResourcePoolApplicationInput) GetAssociatedPoolId() *string {
	return s.AssociatedPoolId
}

func (s *CreateTimedResidentResourcePoolApplicationInput) GetDiskSizeInGB() *int64 {
	return s.DiskSizeInGB
}

func (s *CreateTimedResidentResourcePoolApplicationInput) GetGpuType() *string {
	return s.GpuType
}

func (s *CreateTimedResidentResourcePoolApplicationInput) GetMemorySizeInGB() *int64 {
	return s.MemorySizeInGB
}

func (s *CreateTimedResidentResourcePoolApplicationInput) GetOperationType() *string {
	return s.OperationType
}

func (s *CreateTimedResidentResourcePoolApplicationInput) GetPoolName() *string {
	return s.PoolName
}

func (s *CreateTimedResidentResourcePoolApplicationInput) GetReason() *string {
	return s.Reason
}

func (s *CreateTimedResidentResourcePoolApplicationInput) GetTimedConfig() *TimedPoolConfig {
	return s.TimedConfig
}

func (s *CreateTimedResidentResourcePoolApplicationInput) GetTimedPoolId() *string {
	return s.TimedPoolId
}

func (s *CreateTimedResidentResourcePoolApplicationInput) GetTotalGPUCards() *int64 {
	return s.TotalGPUCards
}

func (s *CreateTimedResidentResourcePoolApplicationInput) GetVCpuCores() *int64 {
	return s.VCpuCores
}

func (s *CreateTimedResidentResourcePoolApplicationInput) SetAssociatedPoolId(v string) *CreateTimedResidentResourcePoolApplicationInput {
	s.AssociatedPoolId = &v
	return s
}

func (s *CreateTimedResidentResourcePoolApplicationInput) SetDiskSizeInGB(v int64) *CreateTimedResidentResourcePoolApplicationInput {
	s.DiskSizeInGB = &v
	return s
}

func (s *CreateTimedResidentResourcePoolApplicationInput) SetGpuType(v string) *CreateTimedResidentResourcePoolApplicationInput {
	s.GpuType = &v
	return s
}

func (s *CreateTimedResidentResourcePoolApplicationInput) SetMemorySizeInGB(v int64) *CreateTimedResidentResourcePoolApplicationInput {
	s.MemorySizeInGB = &v
	return s
}

func (s *CreateTimedResidentResourcePoolApplicationInput) SetOperationType(v string) *CreateTimedResidentResourcePoolApplicationInput {
	s.OperationType = &v
	return s
}

func (s *CreateTimedResidentResourcePoolApplicationInput) SetPoolName(v string) *CreateTimedResidentResourcePoolApplicationInput {
	s.PoolName = &v
	return s
}

func (s *CreateTimedResidentResourcePoolApplicationInput) SetReason(v string) *CreateTimedResidentResourcePoolApplicationInput {
	s.Reason = &v
	return s
}

func (s *CreateTimedResidentResourcePoolApplicationInput) SetTimedConfig(v *TimedPoolConfig) *CreateTimedResidentResourcePoolApplicationInput {
	s.TimedConfig = v
	return s
}

func (s *CreateTimedResidentResourcePoolApplicationInput) SetTimedPoolId(v string) *CreateTimedResidentResourcePoolApplicationInput {
	s.TimedPoolId = &v
	return s
}

func (s *CreateTimedResidentResourcePoolApplicationInput) SetTotalGPUCards(v int64) *CreateTimedResidentResourcePoolApplicationInput {
	s.TotalGPUCards = &v
	return s
}

func (s *CreateTimedResidentResourcePoolApplicationInput) SetVCpuCores(v int64) *CreateTimedResidentResourcePoolApplicationInput {
	s.VCpuCores = &v
	return s
}

func (s *CreateTimedResidentResourcePoolApplicationInput) Validate() error {
	if s.TimedConfig != nil {
		if err := s.TimedConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
