// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPermission interface {
	dara.Model
	String() string
	GoString() string
	SetActionList(v []*PermissionActionList) *Permission
	GetActionList() []*PermissionActionList
	SetCollection(v string) *Permission
	GetCollection() *string
	SetCondition(v *PermissionCondition) *Permission
	GetCondition() *PermissionCondition
	SetCreatedAt(v int64) *Permission
	GetCreatedAt() *int64
	SetEffect(v string) *Permission
	GetEffect() *string
	SetIdentityId(v string) *Permission
	GetIdentityId() *string
	SetIdentityType(v string) *Permission
	GetIdentityType() *string
	SetResource(v string) *Permission
	GetResource() *string
	SetResourceType(v string) *Permission
	GetResourceType() *string
	SetUpdatedAt(v int64) *Permission
	GetUpdatedAt() *int64
	SetUserTags(v []*string) *Permission
	GetUserTags() []*string
}

type Permission struct {
	// Action list.
	ActionList []*PermissionActionList `json:"action_list,omitempty" xml:"action_list,omitempty" type:"Repeated"`
	// The permission set. Set this parameter to global for global permissions. In other scenarios, this parameter is empty by default.
	//
	// example:
	//
	// global
	Collection *string `json:"collection,omitempty" xml:"collection,omitempty"`
	// Condition
	Condition *PermissionCondition `json:"condition,omitempty" xml:"condition,omitempty"`
	// The creation time in the millisecond timestamp format.
	//
	// example:
	//
	// 1703648502811
	CreatedAt *int64 `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// Effect. Valid values: allow, deny.
	//
	// example:
	//
	// deny
	Effect *string `json:"effect,omitempty" xml:"effect,omitempty"`
	// Identity ID.
	//
	// example:
	//
	// af22***
	IdentityId *string `json:"identity_id,omitempty" xml:"identity_id,omitempty"`
	// Identity type. Valid values: IT_User, IT_Group, IT_Role.
	//
	// example:
	//
	// IT_User
	IdentityType *string `json:"identity_type,omitempty" xml:"identity_type,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// fa212***
	Resource *string `json:"resource,omitempty" xml:"resource,omitempty"`
	// The type of the resource. The file type resource is RT_File.
	//
	// example:
	//
	// RT_File
	ResourceType *string `json:"resource_type,omitempty" xml:"resource_type,omitempty"`
	// The time when the modification was made. The value is a millisecond timestamp.
	//
	// example:
	//
	// 1703648502811
	UpdatedAt *int64 `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// Custom tag.
	UserTags []*string `json:"user_tags,omitempty" xml:"user_tags,omitempty" type:"Repeated"`
}

func (s Permission) String() string {
	return dara.Prettify(s)
}

func (s Permission) GoString() string {
	return s.String()
}

func (s *Permission) GetActionList() []*PermissionActionList {
	return s.ActionList
}

func (s *Permission) GetCollection() *string {
	return s.Collection
}

func (s *Permission) GetCondition() *PermissionCondition {
	return s.Condition
}

func (s *Permission) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *Permission) GetEffect() *string {
	return s.Effect
}

func (s *Permission) GetIdentityId() *string {
	return s.IdentityId
}

func (s *Permission) GetIdentityType() *string {
	return s.IdentityType
}

func (s *Permission) GetResource() *string {
	return s.Resource
}

func (s *Permission) GetResourceType() *string {
	return s.ResourceType
}

func (s *Permission) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *Permission) GetUserTags() []*string {
	return s.UserTags
}

func (s *Permission) SetActionList(v []*PermissionActionList) *Permission {
	s.ActionList = v
	return s
}

func (s *Permission) SetCollection(v string) *Permission {
	s.Collection = &v
	return s
}

func (s *Permission) SetCondition(v *PermissionCondition) *Permission {
	s.Condition = v
	return s
}

func (s *Permission) SetCreatedAt(v int64) *Permission {
	s.CreatedAt = &v
	return s
}

func (s *Permission) SetEffect(v string) *Permission {
	s.Effect = &v
	return s
}

func (s *Permission) SetIdentityId(v string) *Permission {
	s.IdentityId = &v
	return s
}

func (s *Permission) SetIdentityType(v string) *Permission {
	s.IdentityType = &v
	return s
}

func (s *Permission) SetResource(v string) *Permission {
	s.Resource = &v
	return s
}

func (s *Permission) SetResourceType(v string) *Permission {
	s.ResourceType = &v
	return s
}

func (s *Permission) SetUpdatedAt(v int64) *Permission {
	s.UpdatedAt = &v
	return s
}

func (s *Permission) SetUserTags(v []*string) *Permission {
	s.UserTags = v
	return s
}

func (s *Permission) Validate() error {
	if s.ActionList != nil {
		for _, item := range s.ActionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Condition != nil {
		if err := s.Condition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PermissionActionList struct {
	// Specific action, such as FILE.ALL
	//
	// example:
	//
	// FILE.ALL
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
}

func (s PermissionActionList) String() string {
	return dara.Prettify(s)
}

func (s PermissionActionList) GoString() string {
	return s.String()
}

func (s *PermissionActionList) GetAction() *string {
	return s.Action
}

func (s *PermissionActionList) SetAction(v string) *PermissionActionList {
	s.Action = &v
	return s
}

func (s *PermissionActionList) Validate() error {
	return dara.Validate(s)
}
