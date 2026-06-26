// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResidentResourcePool interface {
	dara.Model
	String() string
	GoString() string
	SetAllocationStatus(v *ResidentResourceAllocationStatus) *ResidentResourcePool
	GetAllocationStatus() *ResidentResourceAllocationStatus
	SetAssociatedPoolId(v string) *ResidentResourcePool
	GetAssociatedPoolId() *string
	SetCreatedTime(v string) *ResidentResourcePool
	GetCreatedTime() *string
	SetExpireTime(v string) *ResidentResourcePool
	GetExpireTime() *string
	SetLastModifiedTime(v string) *ResidentResourcePool
	GetLastModifiedTime() *string
	SetPoolType(v string) *ResidentResourcePool
	GetPoolType() *string
	SetResidentResourcePoolId(v string) *ResidentResourcePool
	GetResidentResourcePoolId() *string
	SetResidentResourcePoolName(v string) *ResidentResourcePool
	GetResidentResourcePoolName() *string
	SetResourcePoolCapacity(v *ResidentResourceCapacity) *ResidentResourcePool
	GetResourcePoolCapacity() *ResidentResourceCapacity
	SetResourcePoolConfig(v *ResidentResourceCapacity) *ResidentResourcePool
	GetResourcePoolConfig() *ResidentResourceCapacity
	SetTimedConfig(v *TimedPoolConfig) *ResidentResourcePool
	GetTimedConfig() *TimedPoolConfig
}

type ResidentResourcePool struct {
	// The real-time allocation status of the resource pool, including the specific allocation details for each function.
	AllocationStatus *ResidentResourceAllocationStatus `json:"allocationStatus,omitempty" xml:"allocationStatus,omitempty"`
	AssociatedPoolId *string                           `json:"associatedPoolId,omitempty" xml:"associatedPoolId,omitempty"`
	// The resource property field that represents the creation time.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// The expiration time of the resource pool.
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// The last modification time, including operations such as scaling, renewal, and renaming.
	LastModifiedTime       *string `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	PoolType               *string `json:"poolType,omitempty" xml:"poolType,omitempty"`
	ResidentResourcePoolId *string `json:"residentResourcePoolId,omitempty" xml:"residentResourcePoolId,omitempty"`
	// The resource property field that represents the resource name.
	ResidentResourcePoolName *string `json:"residentResourcePoolName,omitempty" xml:"residentResourcePoolName,omitempty"`
	// The overall specifications of the resource pool.
	ResourcePoolCapacity *ResidentResourceCapacity `json:"resourcePoolCapacity,omitempty" xml:"resourcePoolCapacity,omitempty"`
	ResourcePoolConfig   *ResidentResourceCapacity `json:"resourcePoolConfig,omitempty" xml:"resourcePoolConfig,omitempty"`
	TimedConfig          *TimedPoolConfig          `json:"timedConfig,omitempty" xml:"timedConfig,omitempty"`
}

func (s ResidentResourcePool) String() string {
	return dara.Prettify(s)
}

func (s ResidentResourcePool) GoString() string {
	return s.String()
}

func (s *ResidentResourcePool) GetAllocationStatus() *ResidentResourceAllocationStatus {
	return s.AllocationStatus
}

func (s *ResidentResourcePool) GetAssociatedPoolId() *string {
	return s.AssociatedPoolId
}

func (s *ResidentResourcePool) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *ResidentResourcePool) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ResidentResourcePool) GetLastModifiedTime() *string {
	return s.LastModifiedTime
}

func (s *ResidentResourcePool) GetPoolType() *string {
	return s.PoolType
}

func (s *ResidentResourcePool) GetResidentResourcePoolId() *string {
	return s.ResidentResourcePoolId
}

func (s *ResidentResourcePool) GetResidentResourcePoolName() *string {
	return s.ResidentResourcePoolName
}

func (s *ResidentResourcePool) GetResourcePoolCapacity() *ResidentResourceCapacity {
	return s.ResourcePoolCapacity
}

func (s *ResidentResourcePool) GetResourcePoolConfig() *ResidentResourceCapacity {
	return s.ResourcePoolConfig
}

func (s *ResidentResourcePool) GetTimedConfig() *TimedPoolConfig {
	return s.TimedConfig
}

func (s *ResidentResourcePool) SetAllocationStatus(v *ResidentResourceAllocationStatus) *ResidentResourcePool {
	s.AllocationStatus = v
	return s
}

func (s *ResidentResourcePool) SetAssociatedPoolId(v string) *ResidentResourcePool {
	s.AssociatedPoolId = &v
	return s
}

func (s *ResidentResourcePool) SetCreatedTime(v string) *ResidentResourcePool {
	s.CreatedTime = &v
	return s
}

func (s *ResidentResourcePool) SetExpireTime(v string) *ResidentResourcePool {
	s.ExpireTime = &v
	return s
}

func (s *ResidentResourcePool) SetLastModifiedTime(v string) *ResidentResourcePool {
	s.LastModifiedTime = &v
	return s
}

func (s *ResidentResourcePool) SetPoolType(v string) *ResidentResourcePool {
	s.PoolType = &v
	return s
}

func (s *ResidentResourcePool) SetResidentResourcePoolId(v string) *ResidentResourcePool {
	s.ResidentResourcePoolId = &v
	return s
}

func (s *ResidentResourcePool) SetResidentResourcePoolName(v string) *ResidentResourcePool {
	s.ResidentResourcePoolName = &v
	return s
}

func (s *ResidentResourcePool) SetResourcePoolCapacity(v *ResidentResourceCapacity) *ResidentResourcePool {
	s.ResourcePoolCapacity = v
	return s
}

func (s *ResidentResourcePool) SetResourcePoolConfig(v *ResidentResourceCapacity) *ResidentResourcePool {
	s.ResourcePoolConfig = v
	return s
}

func (s *ResidentResourcePool) SetTimedConfig(v *TimedPoolConfig) *ResidentResourcePool {
	s.TimedConfig = v
	return s
}

func (s *ResidentResourcePool) Validate() error {
	if s.AllocationStatus != nil {
		if err := s.AllocationStatus.Validate(); err != nil {
			return err
		}
	}
	if s.ResourcePoolCapacity != nil {
		if err := s.ResourcePoolCapacity.Validate(); err != nil {
			return err
		}
	}
	if s.ResourcePoolConfig != nil {
		if err := s.ResourcePoolConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TimedConfig != nil {
		if err := s.TimedConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
