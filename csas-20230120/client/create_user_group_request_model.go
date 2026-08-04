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
	// The collection of user group attributes. You can specify a maximum of 3,000 attributes. The attributes are combined using a logical OR.
	//
	// This parameter is required.
	Attributes []*CreateUserGroupRequestAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// The user group description. The description must be 1 to 128 characters long and can contain Chinese characters, letters, digits, periods (.), underscores (_), hyphens (-), and spaces.
	//
	// example:
	//
	// 这是一条用户组
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The user group name. The name must be 1 to 128 characters long and can contain Chinese characters, letters, digits, periods (.), underscores (_), and hyphens (-).
	//
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
	// The ID of the identity provider (IdP) for the user group. This parameter is used when UserGroupType is set to **department**.
	//
	// example:
	//
	// 12
	IdpId *int32 `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	// The relationship for the user group. Valid values:
	//
	// - **Equal**: Equal to.
	//
	// - **Unequal**: Not equal to.
	//
	// This parameter is required.
	//
	// example:
	//
	// Equal
	Relation *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
	// The type of the user group. Valid values:
	//
	// - **username**: Username.
	//
	// - **department**: Department.
	//
	// - **email**: Email.
	//
	// - **telephone**: Mobile phone.
	//
	// This parameter is required.
	//
	// example:
	//
	// department
	UserGroupType *string `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
	// The value of the user group attribute.
	//
	// - If UserGroupType is set to **username**, this parameter specifies the username value. The value must be 1 to 128 characters long. It can contain Chinese characters, letters, digits, periods (.), underscores (_), hyphens (-), asterisks (\\*), at signs (@), and spaces.
	//
	// - If UserGroupType is set to **department**, this parameter specifies the department value. For example: OU=Department 1,OU=SASE DingTalk.
	//
	// - If UserGroupType is set to **email**, this parameter specifies the email address. For example: username\\@example.com.
	//
	// - If UserGroupType is set to **telephone**, this parameter specifies the mobile phone number. For example: 13900001234.
	//
	// This parameter is required.
	//
	// example:
	//
	// OU=部门1,OU=SASE钉钉
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
