// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserGroupsForPrivateAccessPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolices(v []*ListUserGroupsForPrivateAccessPolicyResponseBodyPolices) *ListUserGroupsForPrivateAccessPolicyResponseBody
	GetPolices() []*ListUserGroupsForPrivateAccessPolicyResponseBodyPolices
	SetRequestId(v string) *ListUserGroupsForPrivateAccessPolicyResponseBody
	GetRequestId() *string
}

type ListUserGroupsForPrivateAccessPolicyResponseBody struct {
	// List of private network access policies.
	Polices []*ListUserGroupsForPrivateAccessPolicyResponseBodyPolices `json:"Polices,omitempty" xml:"Polices,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 58D6B23E-E5DA-5418-8F61-51A3B5A30049
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListUserGroupsForPrivateAccessPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsForPrivateAccessPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBody) GetPolices() []*ListUserGroupsForPrivateAccessPolicyResponseBodyPolices {
	return s.Polices
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBody) SetPolices(v []*ListUserGroupsForPrivateAccessPolicyResponseBodyPolices) *ListUserGroupsForPrivateAccessPolicyResponseBody {
	s.Polices = v
	return s
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBody) SetRequestId(v string) *ListUserGroupsForPrivateAccessPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBody) Validate() error {
	if s.Polices != nil {
		for _, item := range s.Polices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserGroupsForPrivateAccessPolicyResponseBodyPolices struct {
	// Private network access policy ID.
	//
	// example:
	//
	// pa-policy-1b0d0e8b4bcf****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// Collection of user groups for the private network access policy.
	UserGroups []*ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroups `json:"UserGroups,omitempty" xml:"UserGroups,omitempty" type:"Repeated"`
}

func (s ListUserGroupsForPrivateAccessPolicyResponseBodyPolices) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsForPrivateAccessPolicyResponseBodyPolices) GoString() string {
	return s.String()
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBodyPolices) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBodyPolices) GetUserGroups() []*ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroups {
	return s.UserGroups
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBodyPolices) SetPolicyId(v string) *ListUserGroupsForPrivateAccessPolicyResponseBodyPolices {
	s.PolicyId = &v
	return s
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBodyPolices) SetUserGroups(v []*ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroups) *ListUserGroupsForPrivateAccessPolicyResponseBodyPolices {
	s.UserGroups = v
	return s
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBodyPolices) Validate() error {
	if s.UserGroups != nil {
		for _, item := range s.UserGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroups struct {
	// Collection of user group properties.
	Attributes []*ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroupsAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// User group creation time.
	//
	// example:
	//
	// 2022-09-27 18:10:25
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// User group description.
	//
	// example:
	//
	// 这是一条被内网访问策略引用的用户组
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// User group name.
	//
	// example:
	//
	// user_group_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// User group ID.
	//
	// example:
	//
	// usergroup-6f1ef2fc56b6****
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroups) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroups) GoString() string {
	return s.String()
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroups) GetAttributes() []*ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroupsAttributes {
	return s.Attributes
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroups) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroups) GetDescription() *string {
	return s.Description
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroups) GetName() *string {
	return s.Name
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroups) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroups) SetAttributes(v []*ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroupsAttributes) *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroups {
	s.Attributes = v
	return s
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroups) SetCreateTime(v string) *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroups {
	s.CreateTime = &v
	return s
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroups) SetDescription(v string) *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroups {
	s.Description = &v
	return s
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroups) SetName(v string) *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroups {
	s.Name = &v
	return s
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroups) SetUserGroupId(v string) *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroups {
	s.UserGroupId = &v
	return s
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroups) Validate() error {
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

type ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroupsAttributes struct {
	// The identity source ID of the user group. This value exists if the custom user group type is **department**.
	//
	// example:
	//
	// 12
	IdpId *int32 `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	// The relationship of the user group. Values:
	//
	// - **Equal**: Equal.
	//
	// - **Unequal**: Unequal.
	//
	// example:
	//
	// Equal
	Relation *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
	// The type of the user group. Values:
	//
	// - **username**: username.
	//
	// - **department**: department.
	//
	// - **email**: mailbox.
	//
	// - **telephone**: telephone.
	//
	// example:
	//
	// department
	UserGroupType *string `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
	// Value of the user group property.
	//
	// - If the user group type is **username**, this indicates the username\\"s value. The length is 1 to 128 characters. It supports Chinese characters and uppercase and lowercase English letters. It can contain numbers, periods (.), underscores (_), and hyphens (-).
	//
	// - If the user group type is **department**, this indicates the department\\"s value, such as OU=Department 1,OU=SASE DingTalk.
	//
	// - If the user group type is **email**, this indicates the mailbox\\"s value, such as username\\@example.com.
	//
	// - If the user group type is **telephone**, this indicates the telephone\\"s value, such as 13900001234.
	//
	// example:
	//
	// OU=部门1,OU=SASE钉钉
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroupsAttributes) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroupsAttributes) GoString() string {
	return s.String()
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroupsAttributes) GetIdpId() *int32 {
	return s.IdpId
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroupsAttributes) GetRelation() *string {
	return s.Relation
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroupsAttributes) GetUserGroupType() *string {
	return s.UserGroupType
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroupsAttributes) GetValue() *string {
	return s.Value
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroupsAttributes) SetIdpId(v int32) *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroupsAttributes {
	s.IdpId = &v
	return s
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroupsAttributes) SetRelation(v string) *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroupsAttributes {
	s.Relation = &v
	return s
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroupsAttributes) SetUserGroupType(v string) *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroupsAttributes {
	s.UserGroupType = &v
	return s
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroupsAttributes) SetValue(v string) *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroupsAttributes {
	s.Value = &v
	return s
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroupsAttributes) Validate() error {
	return dara.Validate(s)
}
