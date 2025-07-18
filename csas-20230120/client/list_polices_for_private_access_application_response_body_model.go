// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolicesForPrivateAccessApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplications(v []*ListPolicesForPrivateAccessApplicationResponseBodyApplications) *ListPolicesForPrivateAccessApplicationResponseBody
	GetApplications() []*ListPolicesForPrivateAccessApplicationResponseBodyApplications
	SetRequestId(v string) *ListPolicesForPrivateAccessApplicationResponseBody
	GetRequestId() *string
}

type ListPolicesForPrivateAccessApplicationResponseBody struct {
	Applications []*ListPolicesForPrivateAccessApplicationResponseBodyApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// example:
	//
	// 4AB972E2-D702-5464-B132-B1911498B8BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPolicesForPrivateAccessApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPolicesForPrivateAccessApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ListPolicesForPrivateAccessApplicationResponseBody) GetApplications() []*ListPolicesForPrivateAccessApplicationResponseBodyApplications {
	return s.Applications
}

func (s *ListPolicesForPrivateAccessApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPolicesForPrivateAccessApplicationResponseBody) SetApplications(v []*ListPolicesForPrivateAccessApplicationResponseBodyApplications) *ListPolicesForPrivateAccessApplicationResponseBody {
	s.Applications = v
	return s
}

func (s *ListPolicesForPrivateAccessApplicationResponseBody) SetRequestId(v string) *ListPolicesForPrivateAccessApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPolicesForPrivateAccessApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListPolicesForPrivateAccessApplicationResponseBodyApplications struct {
	// example:
	//
	// pa-application-b927baf3e592****
	ApplicationId *string                                                                   `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	Policies      []*ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
}

func (s ListPolicesForPrivateAccessApplicationResponseBodyApplications) String() string {
	return dara.Prettify(s)
}

func (s ListPolicesForPrivateAccessApplicationResponseBodyApplications) GoString() string {
	return s.String()
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplications) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplications) GetPolicies() []*ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies {
	return s.Policies
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplications) SetApplicationId(v string) *ListPolicesForPrivateAccessApplicationResponseBodyApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplications) SetPolicies(v []*ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies) *ListPolicesForPrivateAccessApplicationResponseBodyApplications {
	s.Policies = v
	return s
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplications) Validate() error {
	return dara.Validate(s)
}

type ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies struct {
	// example:
	//
	// Application
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// example:
	//
	// 2022-09-27 18:10:25
	CreateTime           *string                                                                                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CustomUserAttributes []*ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPoliciesCustomUserAttributes `json:"CustomUserAttributes,omitempty" xml:"CustomUserAttributes,omitempty" type:"Repeated"`
	Description          *string                                                                                       `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// private_access_policy_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Allow
	PolicyAction *string `json:"PolicyAction,omitempty" xml:"PolicyAction,omitempty"`
	// example:
	//
	// pa-policy-867ef4007c8a****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// Normal
	UserGroupType *string `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
}

func (s ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies) GoString() string {
	return s.String()
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies) GetCustomUserAttributes() []*ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPoliciesCustomUserAttributes {
	return s.CustomUserAttributes
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies) GetDescription() *string {
	return s.Description
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies) GetName() *string {
	return s.Name
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies) GetPolicyAction() *string {
	return s.PolicyAction
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies) GetPriority() *int32 {
	return s.Priority
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies) GetStatus() *string {
	return s.Status
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies) GetUserGroupType() *string {
	return s.UserGroupType
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies) SetApplicationType(v string) *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies {
	s.ApplicationType = &v
	return s
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies) SetCreateTime(v string) *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies {
	s.CreateTime = &v
	return s
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies) SetCustomUserAttributes(v []*ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPoliciesCustomUserAttributes) *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies {
	s.CustomUserAttributes = v
	return s
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies) SetDescription(v string) *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies {
	s.Description = &v
	return s
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies) SetName(v string) *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies {
	s.Name = &v
	return s
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies) SetPolicyAction(v string) *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies {
	s.PolicyAction = &v
	return s
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies) SetPolicyId(v string) *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies {
	s.PolicyId = &v
	return s
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies) SetPriority(v int32) *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies {
	s.Priority = &v
	return s
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies) SetStatus(v string) *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies {
	s.Status = &v
	return s
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies) SetUserGroupType(v string) *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies {
	s.UserGroupType = &v
	return s
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies) Validate() error {
	return dara.Validate(s)
}

type ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPoliciesCustomUserAttributes struct {
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

func (s ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPoliciesCustomUserAttributes) String() string {
	return dara.Prettify(s)
}

func (s ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPoliciesCustomUserAttributes) GoString() string {
	return s.String()
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPoliciesCustomUserAttributes) GetIdpId() *int32 {
	return s.IdpId
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPoliciesCustomUserAttributes) GetRelation() *string {
	return s.Relation
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPoliciesCustomUserAttributes) GetUserGroupType() *string {
	return s.UserGroupType
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPoliciesCustomUserAttributes) GetValue() *string {
	return s.Value
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPoliciesCustomUserAttributes) SetIdpId(v int32) *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPoliciesCustomUserAttributes {
	s.IdpId = &v
	return s
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPoliciesCustomUserAttributes) SetRelation(v string) *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPoliciesCustomUserAttributes {
	s.Relation = &v
	return s
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPoliciesCustomUserAttributes) SetUserGroupType(v string) *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPoliciesCustomUserAttributes {
	s.UserGroupType = &v
	return s
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPoliciesCustomUserAttributes) SetValue(v string) *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPoliciesCustomUserAttributes {
	s.Value = &v
	return s
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPoliciesCustomUserAttributes) Validate() error {
	return dara.Validate(s)
}
