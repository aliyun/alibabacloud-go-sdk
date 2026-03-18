// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPaidResourceDTO interface {
	dara.Model
	String() string
	GoString() string
	SetEffectiveTime(v string) *PaidResourceDTO
	GetEffectiveTime() *string
	SetExpireTime(v string) *PaidResourceDTO
	GetExpireTime() *string
	SetInstanceId(v string) *PaidResourceDTO
	GetInstanceId() *string
	SetQuantity(v int32) *PaidResourceDTO
	GetQuantity() *int32
	SetRemainQuantity(v int32) *PaidResourceDTO
	GetRemainQuantity() *int32
	SetResourceCatalogCode(v string) *PaidResourceDTO
	GetResourceCatalogCode() *string
	SetResourceCatalogName(v string) *PaidResourceDTO
	GetResourceCatalogName() *string
	SetResourcePackageCode(v string) *PaidResourceDTO
	GetResourcePackageCode() *string
	SetResourcePackageName(v string) *PaidResourceDTO
	GetResourcePackageName() *string
	SetResourceStatus(v string) *PaidResourceDTO
	GetResourceStatus() *string
}

type PaidResourceDTO struct {
	// example:
	//
	// 2025-09-01 00:00:00
	EffectiveTime *string `json:"effectiveTime,omitempty" xml:"effectiveTime,omitempty"`
	// example:
	//
	// 2025-09-10 00:00:00
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// example:
	//
	// airec-cn-fou41hse8001
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// 10000
	Quantity *int32 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// example:
	//
	// 1000
	RemainQuantity *int32 `json:"remainQuantity,omitempty" xml:"remainQuantity,omitempty"`
	// example:
	//
	// AI_ORAL
	ResourceCatalogCode *string `json:"resourceCatalogCode,omitempty" xml:"resourceCatalogCode,omitempty"`
	ResourceCatalogName *string `json:"resourceCatalogName,omitempty" xml:"resourceCatalogName,omitempty"`
	// example:
	//
	// PRE_PAID_RECOURSE_PACKAGE
	ResourcePackageCode *string `json:"resourcePackageCode,omitempty" xml:"resourcePackageCode,omitempty"`
	ResourcePackageName *string `json:"resourcePackageName,omitempty" xml:"resourcePackageName,omitempty"`
	// example:
	//
	// ACTIVE
	ResourceStatus *string `json:"resourceStatus,omitempty" xml:"resourceStatus,omitempty"`
}

func (s PaidResourceDTO) String() string {
	return dara.Prettify(s)
}

func (s PaidResourceDTO) GoString() string {
	return s.String()
}

func (s *PaidResourceDTO) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *PaidResourceDTO) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *PaidResourceDTO) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PaidResourceDTO) GetQuantity() *int32 {
	return s.Quantity
}

func (s *PaidResourceDTO) GetRemainQuantity() *int32 {
	return s.RemainQuantity
}

func (s *PaidResourceDTO) GetResourceCatalogCode() *string {
	return s.ResourceCatalogCode
}

func (s *PaidResourceDTO) GetResourceCatalogName() *string {
	return s.ResourceCatalogName
}

func (s *PaidResourceDTO) GetResourcePackageCode() *string {
	return s.ResourcePackageCode
}

func (s *PaidResourceDTO) GetResourcePackageName() *string {
	return s.ResourcePackageName
}

func (s *PaidResourceDTO) GetResourceStatus() *string {
	return s.ResourceStatus
}

func (s *PaidResourceDTO) SetEffectiveTime(v string) *PaidResourceDTO {
	s.EffectiveTime = &v
	return s
}

func (s *PaidResourceDTO) SetExpireTime(v string) *PaidResourceDTO {
	s.ExpireTime = &v
	return s
}

func (s *PaidResourceDTO) SetInstanceId(v string) *PaidResourceDTO {
	s.InstanceId = &v
	return s
}

func (s *PaidResourceDTO) SetQuantity(v int32) *PaidResourceDTO {
	s.Quantity = &v
	return s
}

func (s *PaidResourceDTO) SetRemainQuantity(v int32) *PaidResourceDTO {
	s.RemainQuantity = &v
	return s
}

func (s *PaidResourceDTO) SetResourceCatalogCode(v string) *PaidResourceDTO {
	s.ResourceCatalogCode = &v
	return s
}

func (s *PaidResourceDTO) SetResourceCatalogName(v string) *PaidResourceDTO {
	s.ResourceCatalogName = &v
	return s
}

func (s *PaidResourceDTO) SetResourcePackageCode(v string) *PaidResourceDTO {
	s.ResourcePackageCode = &v
	return s
}

func (s *PaidResourceDTO) SetResourcePackageName(v string) *PaidResourceDTO {
	s.ResourcePackageName = &v
	return s
}

func (s *PaidResourceDTO) SetResourceStatus(v string) *PaidResourceDTO {
	s.ResourceStatus = &v
	return s
}

func (s *PaidResourceDTO) Validate() error {
	return dara.Validate(s)
}
