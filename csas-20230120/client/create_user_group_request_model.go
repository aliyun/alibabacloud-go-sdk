// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttributes(v []*CreateUserGroupRequestAttributes) *CreateUserGroupRequest
	GetAttributes() []*CreateUserGroupRequestAttributes
	SetDescription(v string) *CreateUserGroupRequest
	GetDescription() *string
	SetName(v string) *CreateUserGroupRequest
	GetName() *string
}

type CreateUserGroupRequest struct {
	// This parameter is required.
	Attributes  []*CreateUserGroupRequestAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	Description *string                             `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user_group_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateUserGroupRequest) GetAttributes() []*CreateUserGroupRequestAttributes {
	return s.Attributes
}

func (s *CreateUserGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateUserGroupRequest) GetName() *string {
	return s.Name
}

func (s *CreateUserGroupRequest) SetAttributes(v []*CreateUserGroupRequestAttributes) *CreateUserGroupRequest {
	s.Attributes = v
	return s
}

func (s *CreateUserGroupRequest) SetDescription(v string) *CreateUserGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateUserGroupRequest) SetName(v string) *CreateUserGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateUserGroupRequest) Validate() error {
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

type CreateUserGroupRequestAttributes struct {
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

func (s CreateUserGroupRequestAttributes) String() string {
	return dara.Prettify(s)
}

func (s CreateUserGroupRequestAttributes) GoString() string {
	return s.String()
}

func (s *CreateUserGroupRequestAttributes) GetIdpId() *int32 {
	return s.IdpId
}

func (s *CreateUserGroupRequestAttributes) GetRelation() *string {
	return s.Relation
}

func (s *CreateUserGroupRequestAttributes) GetUserGroupType() *string {
	return s.UserGroupType
}

func (s *CreateUserGroupRequestAttributes) GetValue() *string {
	return s.Value
}

func (s *CreateUserGroupRequestAttributes) SetIdpId(v int32) *CreateUserGroupRequestAttributes {
	s.IdpId = &v
	return s
}

func (s *CreateUserGroupRequestAttributes) SetRelation(v string) *CreateUserGroupRequestAttributes {
	s.Relation = &v
	return s
}

func (s *CreateUserGroupRequestAttributes) SetUserGroupType(v string) *CreateUserGroupRequestAttributes {
	s.UserGroupType = &v
	return s
}

func (s *CreateUserGroupRequestAttributes) SetValue(v string) *CreateUserGroupRequestAttributes {
	s.Value = &v
	return s
}

func (s *CreateUserGroupRequestAttributes) Validate() error {
	return dara.Validate(s)
}
