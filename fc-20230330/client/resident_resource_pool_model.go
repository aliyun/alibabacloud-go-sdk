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
	SetCreatedTime(v string) *ResidentResourcePool
	GetCreatedTime() *string
	SetExpireTime(v string) *ResidentResourcePool
	GetExpireTime() *string
	SetLastModifiedTime(v string) *ResidentResourcePool
	GetLastModifiedTime() *string
	SetResidentResourcePoolId(v string) *ResidentResourcePool
	GetResidentResourcePoolId() *string
	SetResidentResourcePoolName(v string) *ResidentResourcePool
	GetResidentResourcePoolName() *string
	SetResourcePoolCapacity(v *ResidentResourceCapacity) *ResidentResourcePool
	GetResourcePoolCapacity() *ResidentResourceCapacity
	SetResourcePoolConfig(v *ResidentResourceCapacity) *ResidentResourcePool
	GetResourcePoolConfig() *ResidentResourceCapacity
}

type ResidentResourcePool struct {
	// 资源池实时分配情况，包含每个函数的具体分配情况
	AllocationStatus *ResidentResourceAllocationStatus `json:"allocationStatus,omitempty" xml:"allocationStatus,omitempty"`
	// 代表创建时间的资源属性字段
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// 资源池过期时间
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// 上次修改时间，包含扩容、续费、更名等操作
	LastModifiedTime       *string `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	ResidentResourcePoolId *string `json:"residentResourcePoolId,omitempty" xml:"residentResourcePoolId,omitempty"`
	// 代表资源名称的资源属性字段
	ResidentResourcePoolName *string `json:"residentResourcePoolName,omitempty" xml:"residentResourcePoolName,omitempty"`
	// 资源池总体规格
	ResourcePoolCapacity *ResidentResourceCapacity `json:"resourcePoolCapacity,omitempty" xml:"resourcePoolCapacity,omitempty"`
	ResourcePoolConfig   *ResidentResourceCapacity `json:"resourcePoolConfig,omitempty" xml:"resourcePoolConfig,omitempty"`
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

func (s *ResidentResourcePool) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *ResidentResourcePool) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ResidentResourcePool) GetLastModifiedTime() *string {
	return s.LastModifiedTime
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

func (s *ResidentResourcePool) SetAllocationStatus(v *ResidentResourceAllocationStatus) *ResidentResourcePool {
	s.AllocationStatus = v
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
	return nil
}
