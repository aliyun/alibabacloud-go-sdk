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
	// The set of user group attributes. The maximum total number is 3000. Multiple user group attributes have an OR relationship and take effect as a union.
	Attributes []*UpdateUserGroupRequestAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// The description of the user group. The description must be 1 to 128 characters in length, and can contain Chinese characters, uppercase and lowercase letters, digits, periods (.), underscores (_), hyphens (-), and spaces.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 这是一条用户组
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The modification type of the user group. Valid values:
	//
	// - **Cover*	- (default): overwrites the original user group attribute set with the value of the **Attributes*	- parameter.
	//
	// - **Append**: separately appends the values entered in the **Attributes*	- parameter to the user group attribute set.
	//
	// example:
	//
	// Cover
	ModifyType *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	// The ID of the user group. Value sources:
	//
	// - [ListUserGroups](~~ListUserGroups~~): queries user groups in batches.
	//
	// - [CreateUserGroup](~~CreateUserGroup~~): creates a user group.
	//
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
	// The identity provider ID of the user group. This value exists when the custom user group type is **department**.
	//
	// example:
	//
	// 12
	IdpId *int32 `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	// The relation of the user group. Valid values:
	//
	// - **Equal**: equal to.
	//
	// - **Unequal**: not equal to.
	//
	// This parameter is required.
	//
	// example:
	//
	// Equal
	Relation *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
	// The type of the user group. Valid values:
	//
	// - **username**: username.
	//
	// - **department**: department.
	//
	// - **email**: email.
	//
	// - **telephone**: mobile phone.
	//
	// This parameter is required.
	//
	// example:
	//
	// department
	UserGroupType *string `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
	// The value of the user group attribute.
	//
	// - If the user group type is **username**, this parameter indicates the username value. The value must be 1 to 128 characters in length, and can contain Chinese characters, uppercase and lowercase letters, digits, periods (.), underscores (_), hyphens (-), asterisks (*), at signs (@), and spaces.
	//
	// - If the user group type is **department**, this parameter indicates the department value. Example: OU=Department1,OU=SASEDingTalk.
	//
	// - If the user group type is **email**, this parameter indicates the email value. Example: username@example.com.
	//
	// - If the user group type is **telephone**, this parameter indicates the mobile phone value. Example: 13900001234.
	//
	// This parameter is required.
	//
	// example:
	//
	// OU=部门1,OU=SASE钉钉
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
