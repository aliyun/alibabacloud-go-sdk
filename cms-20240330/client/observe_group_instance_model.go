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
	// The entity category (adapted by entityType).
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// The key-value pairs of monitoring dimensions.
	Dimension map[string]*string `json:"dimension,omitempty" xml:"dimension,omitempty"`
	// The entity domain (such as acs).
	EntityDomain *string `json:"entityDomain,omitempty" xml:"entityDomain,omitempty"`
	// The entity type (such as acs.ecs.instance).
	EntityType *string `json:"entityType,omitempty" xml:"entityType,omitempty"`
	// The ID of the observation group to which the entity belongs.
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The instance ID.
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The instance name.
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// The region ID of the instance.
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The resource group ID of the instance.
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The key-value pairs of instance tags.
	Tags map[string]*string `json:"tags,omitempty" xml:"tags,omitempty"`
	// The UID of the user to which the instance belongs.
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
