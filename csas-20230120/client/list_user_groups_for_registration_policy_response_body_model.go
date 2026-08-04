// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserGroupsForRegistrationPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicies(v []*ListUserGroupsForRegistrationPolicyResponseBodyPolicies) *ListUserGroupsForRegistrationPolicyResponseBody
	GetPolicies() []*ListUserGroupsForRegistrationPolicyResponseBodyPolicies
	SetRequestId(v string) *ListUserGroupsForRegistrationPolicyResponseBody
	GetRequestId() *string
}

type ListUserGroupsForRegistrationPolicyResponseBody struct {
	// A list of device registration policies.
	Policies []*ListUserGroupsForRegistrationPolicyResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// The ID of this request.
	//
	// example:
	//
	// D89009C7-54C6-51B6-BAE7-3F373920C6BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListUserGroupsForRegistrationPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsForRegistrationPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserGroupsForRegistrationPolicyResponseBody) GetPolicies() []*ListUserGroupsForRegistrationPolicyResponseBodyPolicies {
	return s.Policies
}

func (s *ListUserGroupsForRegistrationPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserGroupsForRegistrationPolicyResponseBody) SetPolicies(v []*ListUserGroupsForRegistrationPolicyResponseBodyPolicies) *ListUserGroupsForRegistrationPolicyResponseBody {
	s.Policies = v
	return s
}

func (s *ListUserGroupsForRegistrationPolicyResponseBody) SetRequestId(v string) *ListUserGroupsForRegistrationPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserGroupsForRegistrationPolicyResponseBody) Validate() error {
	if s.Policies != nil {
		for _, item := range s.Policies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserGroupsForRegistrationPolicyResponseBodyPolicies struct {
	// The ID of the device registration policy.
	//
	// example:
	//
	// reg-policy-f25c9e5872e5****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// A collection of user groups associated with the device registration policy.
	UserGroups []*ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroups `json:"UserGroups,omitempty" xml:"UserGroups,omitempty" type:"Repeated"`
}

func (s ListUserGroupsForRegistrationPolicyResponseBodyPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsForRegistrationPolicyResponseBodyPolicies) GoString() string {
	return s.String()
}

func (s *ListUserGroupsForRegistrationPolicyResponseBodyPolicies) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListUserGroupsForRegistrationPolicyResponseBodyPolicies) GetUserGroups() []*ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroups {
	return s.UserGroups
}

func (s *ListUserGroupsForRegistrationPolicyResponseBodyPolicies) SetPolicyId(v string) *ListUserGroupsForRegistrationPolicyResponseBodyPolicies {
	s.PolicyId = &v
	return s
}

func (s *ListUserGroupsForRegistrationPolicyResponseBodyPolicies) SetUserGroups(v []*ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroups) *ListUserGroupsForRegistrationPolicyResponseBodyPolicies {
	s.UserGroups = v
	return s
}

func (s *ListUserGroupsForRegistrationPolicyResponseBodyPolicies) Validate() error {
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

type ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroups struct {
	// A collection of user group attributes.
	Attributes []*ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroupsAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// The time when the user group was created.
	//
	// example:
	//
	// 2022-09-27 18:10:25
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// A description of the user group.
	//
	// example:
	//
	// 这是一条被设备注册策略引用的用户组。
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

func (s ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroups) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroups) GoString() string {
	return s.String()
}

func (s *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroups) GetAttributes() []*ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroupsAttributes {
	return s.Attributes
}

func (s *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroups) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroups) GetDescription() *string {
	return s.Description
}

func (s *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroups) GetName() *string {
	return s.Name
}

func (s *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroups) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroups) SetAttributes(v []*ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroupsAttributes) *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroups {
	s.Attributes = v
	return s
}

func (s *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroups) SetCreateTime(v string) *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroups {
	s.CreateTime = &v
	return s
}

func (s *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroups) SetDescription(v string) *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroups {
	s.Description = &v
	return s
}

func (s *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroups) SetName(v string) *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroups {
	s.Name = &v
	return s
}

func (s *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroups) SetUserGroupId(v string) *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroups {
	s.UserGroupId = &v
	return s
}

func (s *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroups) Validate() error {
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

type ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroupsAttributes struct {
	// The identity provider ID for the user group. This field appears only when UserGroupType is **department**.
	//
	// example:
	//
	// 12
	IdpId *int32 `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	// The relation for the user group. Valid values:
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
	// - **username**: A username.
	//
	// - **department**: A department.
	//
	// - **email**: An email address.
	//
	// - **telephone**: A phone number.
	//
	// example:
	//
	// department
	UserGroupType *string `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
	// The value of the user group attribute.
	//
	// - If UserGroupType is **username**, this is the username. It must be 1–128 characters long and can contain uppercase and lowercase letters, Chinese characters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// - If UserGroupType is **department**, this is the department name. Example: OU=Department 1,OU=SASE DingTalk.
	//
	// - If UserGroupType is **email**, this is the email address. Example: username\\@example.com.
	//
	// - If UserGroupType is **telephone**, this is the phone number. Example: 13900001234.
	//
	// example:
	//
	// OU=部门1,OU=SASE钉钉
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroupsAttributes) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroupsAttributes) GoString() string {
	return s.String()
}

func (s *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroupsAttributes) GetIdpId() *int32 {
	return s.IdpId
}

func (s *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroupsAttributes) GetRelation() *string {
	return s.Relation
}

func (s *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroupsAttributes) GetUserGroupType() *string {
	return s.UserGroupType
}

func (s *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroupsAttributes) GetValue() *string {
	return s.Value
}

func (s *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroupsAttributes) SetIdpId(v int32) *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroupsAttributes {
	s.IdpId = &v
	return s
}

func (s *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroupsAttributes) SetRelation(v string) *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroupsAttributes {
	s.Relation = &v
	return s
}

func (s *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroupsAttributes) SetUserGroupType(v string) *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroupsAttributes {
	s.UserGroupType = &v
	return s
}

func (s *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroupsAttributes) SetValue(v string) *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroupsAttributes {
	s.Value = &v
	return s
}

func (s *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroupsAttributes) Validate() error {
	return dara.Validate(s)
}
