// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetUserGroupResponseBody
	GetRequestId() *string
	SetUserGroup(v *GetUserGroupResponseBodyUserGroup) *GetUserGroupResponseBody
	GetUserGroup() *GetUserGroupResponseBodyUserGroup
}

type GetUserGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1310DBC7-7E1F-55D3-B4B4-E4BE912517FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The user group.
	UserGroup *GetUserGroupResponseBodyUserGroup `json:"UserGroup,omitempty" xml:"UserGroup,omitempty" type:"Struct"`
}

func (s GetUserGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserGroupResponseBody) GetUserGroup() *GetUserGroupResponseBodyUserGroup {
	return s.UserGroup
}

func (s *GetUserGroupResponseBody) SetRequestId(v string) *GetUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserGroupResponseBody) SetUserGroup(v *GetUserGroupResponseBodyUserGroup) *GetUserGroupResponseBody {
	s.UserGroup = v
	return s
}

func (s *GetUserGroupResponseBody) Validate() error {
	if s.UserGroup != nil {
		if err := s.UserGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserGroupResponseBodyUserGroup struct {
	// The collection of user group properties. The properties are evaluated with a logical OR operator.
	Attributes []*GetUserGroupResponseBodyUserGroupAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// The time when the user group was created.
	//
	// example:
	//
	// 2022-10-10 11:39:22
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the user group.
	//
	// example:
	//
	// 这是一条用户组
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the user group.
	//
	// example:
	//
	// user_group_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the user group.
	//
	// example:
	//
	// usergroup-6f1ef2fc56b6****
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s GetUserGroupResponseBodyUserGroup) String() string {
	return dara.Prettify(s)
}

func (s GetUserGroupResponseBodyUserGroup) GoString() string {
	return s.String()
}

func (s *GetUserGroupResponseBodyUserGroup) GetAttributes() []*GetUserGroupResponseBodyUserGroupAttributes {
	return s.Attributes
}

func (s *GetUserGroupResponseBodyUserGroup) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetUserGroupResponseBodyUserGroup) GetDescription() *string {
	return s.Description
}

func (s *GetUserGroupResponseBodyUserGroup) GetName() *string {
	return s.Name
}

func (s *GetUserGroupResponseBodyUserGroup) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *GetUserGroupResponseBodyUserGroup) SetAttributes(v []*GetUserGroupResponseBodyUserGroupAttributes) *GetUserGroupResponseBodyUserGroup {
	s.Attributes = v
	return s
}

func (s *GetUserGroupResponseBodyUserGroup) SetCreateTime(v string) *GetUserGroupResponseBodyUserGroup {
	s.CreateTime = &v
	return s
}

func (s *GetUserGroupResponseBodyUserGroup) SetDescription(v string) *GetUserGroupResponseBodyUserGroup {
	s.Description = &v
	return s
}

func (s *GetUserGroupResponseBodyUserGroup) SetName(v string) *GetUserGroupResponseBodyUserGroup {
	s.Name = &v
	return s
}

func (s *GetUserGroupResponseBodyUserGroup) SetUserGroupId(v string) *GetUserGroupResponseBodyUserGroup {
	s.UserGroupId = &v
	return s
}

func (s *GetUserGroupResponseBodyUserGroup) Validate() error {
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

type GetUserGroupResponseBodyUserGroupAttributes struct {
	// The ID of the identity provider (IdP) for the user group. This parameter is returned when UserGroupType is set to **department**.
	//
	// example:
	//
	// 12
	IdpId *int32 `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	// The relationship of the user group. Valid values:
	//
	// - **Equal**: Equal to.
	//
	// - **Unequal**: Not equal to.
	//
	// example:
	//
	// Equal
	Relation *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
	// The type of the user group. Valid values:
	//
	// - **username**: The username.
	//
	// - **department**: The department.
	//
	// - **email**: The email address.
	//
	// - **telephone**: The mobile phone number.
	//
	// example:
	//
	// department
	UserGroupType *string `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
	// The value of the user group property.
	//
	// - If UserGroupType is set to **username**, this parameter specifies the value of the username. The value can be 1 to 128 characters in length and can contain Chinese characters, letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// - If UserGroupType is set to **department**, this parameter specifies the value of the department. Example: OU=Department 1,OU=SASE DingTalk.
	//
	// - If UserGroupType is set to **email**, this parameter specifies the value of the email address. Example: username\\@example.com.
	//
	// - If UserGroupType is set to **telephone**, this parameter specifies the value of the mobile phone number. Example: 13900001234.
	//
	// example:
	//
	// OU=部门1,OU=SASE钉钉
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetUserGroupResponseBodyUserGroupAttributes) String() string {
	return dara.Prettify(s)
}

func (s GetUserGroupResponseBodyUserGroupAttributes) GoString() string {
	return s.String()
}

func (s *GetUserGroupResponseBodyUserGroupAttributes) GetIdpId() *int32 {
	return s.IdpId
}

func (s *GetUserGroupResponseBodyUserGroupAttributes) GetRelation() *string {
	return s.Relation
}

func (s *GetUserGroupResponseBodyUserGroupAttributes) GetUserGroupType() *string {
	return s.UserGroupType
}

func (s *GetUserGroupResponseBodyUserGroupAttributes) GetValue() *string {
	return s.Value
}

func (s *GetUserGroupResponseBodyUserGroupAttributes) SetIdpId(v int32) *GetUserGroupResponseBodyUserGroupAttributes {
	s.IdpId = &v
	return s
}

func (s *GetUserGroupResponseBodyUserGroupAttributes) SetRelation(v string) *GetUserGroupResponseBodyUserGroupAttributes {
	s.Relation = &v
	return s
}

func (s *GetUserGroupResponseBodyUserGroupAttributes) SetUserGroupType(v string) *GetUserGroupResponseBodyUserGroupAttributes {
	s.UserGroupType = &v
	return s
}

func (s *GetUserGroupResponseBodyUserGroupAttributes) SetValue(v string) *GetUserGroupResponseBodyUserGroupAttributes {
	s.Value = &v
	return s
}

func (s *GetUserGroupResponseBodyUserGroupAttributes) Validate() error {
	return dara.Validate(s)
}
