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
	Polices []*ListUserGroupsForPrivateAccessPolicyResponseBodyPolices `json:"Polices,omitempty" xml:"Polices,omitempty" type:"Repeated"`
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
	return dara.Validate(s)
}

type ListUserGroupsForPrivateAccessPolicyResponseBodyPolices struct {
	// example:
	//
	// pa-policy-1b0d0e8b4bcf****
	PolicyId   *string                                                              `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
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
	return dara.Validate(s)
}

type ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroups struct {
	Attributes []*ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroupsAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// 用户组创建时间。
	//
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
	return dara.Validate(s)
}

type ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroupsAttributes struct {
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
