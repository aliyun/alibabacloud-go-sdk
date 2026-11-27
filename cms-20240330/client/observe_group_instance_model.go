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
	SetEntityId(v string) *ObserveGroupInstance
	GetEntityId() *string
	SetEntityRole(v string) *ObserveGroupInstance
	GetEntityRole() *string
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
	// The entity category adapted from entityType.
	//
	// example:
	//
	// ecs
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// The key-value pairs of the monitoring dimensions.
	//
	// example:
	//
	// {"userId":"1234567890123456","instanceId":"i-bp1gj8wyq4fxk28k6iy2"}
	Dimension map[string]*string `json:"dimension,omitempty" xml:"dimension,omitempty"`
	// The entity domain, such as acs.
	//
	// example:
	//
	// acs
	EntityDomain *string `json:"entityDomain,omitempty" xml:"entityDomain,omitempty"`
	// The UModel entity ID, which corresponds to the __entity_id__ field of the entity and identifies a member of the observation group together with entityDomain and entityType.
	//
	// example:
	//
	// i-bp1gj8wyq4fxk28k6iy2
	EntityId *string `json:"entityId,omitempty" xml:"entityId,omitempty"`
	// The role of the entity in the observation group. main indicates the primary entity directly matched by the management rule. related indicates the child entity managed through associations.
	//
	// example:
	//
	// main
	EntityRole *string `json:"entityRole,omitempty" xml:"entityRole,omitempty"`
	// The entity type, such as acs.ecs.instance.
	//
	// example:
	//
	// acs.ecs.instance
	EntityType *string `json:"entityType,omitempty" xml:"entityType,omitempty"`
	// The ID of the observation group to which the entity belongs.
	//
	// example:
	//
	// og-845e0a26455f437c
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-bp1gj8wyq4fxk28k6iy2
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// ecs-prod-01
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// The ID of the region to which the instance belongs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-acfm3gn5i6bigbi
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The key-value pairs of the instance tags.
	//
	// example:
	//
	// {"env":"prod"}
	Tags map[string]*string `json:"tags,omitempty" xml:"tags,omitempty"`
	// The UID of the user that owns the instance.
	//
	// example:
	//
	// 1234567890123456
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

func (s *ObserveGroupInstance) GetEntityId() *string {
	return s.EntityId
}

func (s *ObserveGroupInstance) GetEntityRole() *string {
	return s.EntityRole
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

func (s *ObserveGroupInstance) SetEntityId(v string) *ObserveGroupInstance {
	s.EntityId = &v
	return s
}

func (s *ObserveGroupInstance) SetEntityRole(v string) *ObserveGroupInstance {
	s.EntityRole = &v
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
