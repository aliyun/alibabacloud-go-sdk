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
	Policies []*ListUserGroupsForRegistrationPolicyResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
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
	// example:
	//
	// reg-policy-f25c9e5872e5****
	PolicyId   *string                                                              `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
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
	Attributes []*ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroupsAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// example:
	//
	// 2022-09-27 18:10:25
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// user_group_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// example:
	//
	// 12
	IdpId *int32 `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	// example:
	//
	// Equal
	Relation *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
	// example:
	//
	// department
	UserGroupType *string `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
