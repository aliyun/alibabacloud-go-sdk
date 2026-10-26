// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObserveGroupInstance interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *ObserveGroupInstance
	GetCategory() *string
	SetDimension(v map[string]*string) *ObserveGroupInstance
	GetDimension() map[string]*string
	SetEntityDomain(v string) *ObserveGroupInstance
	GetEntityDomain() *string
	SetEntityType(v string) *ObserveGroupInstance
	GetEntityType() *string
	SetGroupId(v string) *ObserveGroupInstance
	GetGroupId() *string
	SetInstanceId(v string) *ObserveGroupInstance
	GetInstanceId() *string
	SetInstanceName(v string) *ObserveGroupInstance
	GetInstanceName() *string
	SetRegionId(v string) *ObserveGroupInstance
	GetRegionId() *string
	SetResourceGroupId(v string) *ObserveGroupInstance
	GetResourceGroupId() *string
	SetTags(v map[string]*string) *ObserveGroupInstance
	GetTags() map[string]*string
	SetUserId(v string) *ObserveGroupInstance
	GetUserId() *string
}

type ObserveGroupInstance struct {
	// 实体分类（由 entityType 适配）。
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// 监控维度键值对。
	Dimension map[string]*string `json:"dimension,omitempty" xml:"dimension,omitempty"`
	// 实体域（如 acs）。
	EntityDomain *string `json:"entityDomain,omitempty" xml:"entityDomain,omitempty"`
	// 实体类型（如 acs.ecs.instance）。
	EntityType *string `json:"entityType,omitempty" xml:"entityType,omitempty"`
	// 所属观测分组 ID。
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// 实例 ID。
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// 实例名称。
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// 实例所属地域 ID。
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// 实例所属资源组 ID。
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// 实例标签键值对。
	Tags map[string]*string `json:"tags,omitempty" xml:"tags,omitempty"`
	// 实例归属的用户 UID。
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ObserveGroupInstance) String() string {
	return dara.Prettify(s)
}

func (s ObserveGroupInstance) GoString() string {
	return s.String()
}

func (s *ObserveGroupInstance) GetCategory() *string {
	return s.Category
}

func (s *ObserveGroupInstance) GetDimension() map[string]*string {
	return s.Dimension
}

func (s *ObserveGroupInstance) GetEntityDomain() *string {
	return s.EntityDomain
}

func (s *ObserveGroupInstance) GetEntityType() *string {
	return s.EntityType
}

func (s *ObserveGroupInstance) GetGroupId() *string {
	return s.GroupId
}

func (s *ObserveGroupInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ObserveGroupInstance) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ObserveGroupInstance) GetRegionId() *string {
	return s.RegionId
}

func (s *ObserveGroupInstance) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ObserveGroupInstance) GetTags() map[string]*string {
	return s.Tags
}

func (s *ObserveGroupInstance) GetUserId() *string {
	return s.UserId
}

func (s *ObserveGroupInstance) SetCategory(v string) *ObserveGroupInstance {
	s.Category = &v
	return s
}

func (s *ObserveGroupInstance) SetDimension(v map[string]*string) *ObserveGroupInstance {
	s.Dimension = v
	return s
}

func (s *ObserveGroupInstance) SetEntityDomain(v string) *ObserveGroupInstance {
	s.EntityDomain = &v
	return s
}

func (s *ObserveGroupInstance) SetEntityType(v string) *ObserveGroupInstance {
	s.EntityType = &v
	return s
}

func (s *ObserveGroupInstance) SetGroupId(v string) *ObserveGroupInstance {
	s.GroupId = &v
	return s
}

func (s *ObserveGroupInstance) SetInstanceId(v string) *ObserveGroupInstance {
	s.InstanceId = &v
	return s
}

func (s *ObserveGroupInstance) SetInstanceName(v string) *ObserveGroupInstance {
	s.InstanceName = &v
	return s
}

func (s *ObserveGroupInstance) SetRegionId(v string) *ObserveGroupInstance {
	s.RegionId = &v
	return s
}

func (s *ObserveGroupInstance) SetResourceGroupId(v string) *ObserveGroupInstance {
	s.ResourceGroupId = &v
	return s
}

func (s *ObserveGroupInstance) SetTags(v map[string]*string) *ObserveGroupInstance {
	s.Tags = v
	return s
}

func (s *ObserveGroupInstance) SetUserId(v string) *ObserveGroupInstance {
	s.UserId = &v
	return s
}

func (s *ObserveGroupInstance) Validate() error {
	return dara.Validate(s)
}
