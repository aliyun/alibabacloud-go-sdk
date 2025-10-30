// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttributes(v []*UpdateUserGroupRequestAttributes) *UpdateUserGroupRequest
	GetAttributes() []*UpdateUserGroupRequestAttributes
	SetDescription(v string) *UpdateUserGroupRequest
	GetDescription() *string
	SetModifyType(v string) *UpdateUserGroupRequest
	GetModifyType() *string
	SetUserGroupId(v string) *UpdateUserGroupRequest
	GetUserGroupId() *string
}

type UpdateUserGroupRequest struct {
	Attributes []*UpdateUserGroupRequestAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// if can be null:
	// true
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// Cover
	ModifyType *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// usergroup-6f1ef2fc56b6****
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s UpdateUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserGroupRequest) GetAttributes() []*UpdateUserGroupRequestAttributes {
	return s.Attributes
}

func (s *UpdateUserGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateUserGroupRequest) GetModifyType() *string {
	return s.ModifyType
}

func (s *UpdateUserGroupRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *UpdateUserGroupRequest) SetAttributes(v []*UpdateUserGroupRequestAttributes) *UpdateUserGroupRequest {
	s.Attributes = v
	return s
}

func (s *UpdateUserGroupRequest) SetDescription(v string) *UpdateUserGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateUserGroupRequest) SetModifyType(v string) *UpdateUserGroupRequest {
	s.ModifyType = &v
	return s
}

func (s *UpdateUserGroupRequest) SetUserGroupId(v string) *UpdateUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *UpdateUserGroupRequest) Validate() error {
	if s.Attributes != nil {
		for _, item := range s.Attributes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateUserGroupRequestAttributes struct {
	// example:
	//
	// 12
	IdpId *int32 `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Equal
	Relation *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// department
	UserGroupType *string `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
	// This parameter is required.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateUserGroupRequestAttributes) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserGroupRequestAttributes) GoString() string {
	return s.String()
}

func (s *UpdateUserGroupRequestAttributes) GetIdpId() *int32 {
	return s.IdpId
}

func (s *UpdateUserGroupRequestAttributes) GetRelation() *string {
	return s.Relation
}

func (s *UpdateUserGroupRequestAttributes) GetUserGroupType() *string {
	return s.UserGroupType
}

func (s *UpdateUserGroupRequestAttributes) GetValue() *string {
	return s.Value
}

func (s *UpdateUserGroupRequestAttributes) SetIdpId(v int32) *UpdateUserGroupRequestAttributes {
	s.IdpId = &v
	return s
}

func (s *UpdateUserGroupRequestAttributes) SetRelation(v string) *UpdateUserGroupRequestAttributes {
	s.Relation = &v
	return s
}

func (s *UpdateUserGroupRequestAttributes) SetUserGroupType(v string) *UpdateUserGroupRequestAttributes {
	s.UserGroupType = &v
	return s
}

func (s *UpdateUserGroupRequestAttributes) SetValue(v string) *UpdateUserGroupRequestAttributes {
	s.Value = &v
	return s
}

func (s *UpdateUserGroupRequestAttributes) Validate() error {
	return dara.Validate(s)
}
