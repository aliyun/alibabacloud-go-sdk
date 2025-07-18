// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRegistrationPoliciesForUserGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListRegistrationPoliciesForUserGroupResponseBody
	GetRequestId() *string
	SetUserGroups(v []*ListRegistrationPoliciesForUserGroupResponseBodyUserGroups) *ListRegistrationPoliciesForUserGroupResponseBody
	GetUserGroups() []*ListRegistrationPoliciesForUserGroupResponseBodyUserGroups
}

type ListRegistrationPoliciesForUserGroupResponseBody struct {
	// example:
	//
	// 58D6B23E-E5DA-5418-8F61-51A3B5A30049
	RequestId  *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserGroups []*ListRegistrationPoliciesForUserGroupResponseBodyUserGroups `json:"UserGroups,omitempty" xml:"UserGroups,omitempty" type:"Repeated"`
}

func (s ListRegistrationPoliciesForUserGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRegistrationPoliciesForUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegistrationPoliciesForUserGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRegistrationPoliciesForUserGroupResponseBody) GetUserGroups() []*ListRegistrationPoliciesForUserGroupResponseBodyUserGroups {
	return s.UserGroups
}

func (s *ListRegistrationPoliciesForUserGroupResponseBody) SetRequestId(v string) *ListRegistrationPoliciesForUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRegistrationPoliciesForUserGroupResponseBody) SetUserGroups(v []*ListRegistrationPoliciesForUserGroupResponseBodyUserGroups) *ListRegistrationPoliciesForUserGroupResponseBody {
	s.UserGroups = v
	return s
}

func (s *ListRegistrationPoliciesForUserGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListRegistrationPoliciesForUserGroupResponseBodyUserGroups struct {
	Policies []*ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// example:
	//
	// usergroup-6f1ef2fc56b6****
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s ListRegistrationPoliciesForUserGroupResponseBodyUserGroups) String() string {
	return dara.Prettify(s)
}

func (s ListRegistrationPoliciesForUserGroupResponseBodyUserGroups) GoString() string {
	return s.String()
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroups) GetPolicies() []*ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies {
	return s.Policies
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroups) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroups) SetPolicies(v []*ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies) *ListRegistrationPoliciesForUserGroupResponseBodyUserGroups {
	s.Policies = v
	return s
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroups) SetUserGroupId(v string) *ListRegistrationPoliciesForUserGroupResponseBodyUserGroups {
	s.UserGroupId = &v
	return s
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroups) Validate() error {
	return dara.Validate(s)
}

type ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies struct {
	// example:
	//
	// 2023-05-16 17:18:46
	CreateTime  *string                                                                          `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string                                                                          `json:"Description,omitempty" xml:"Description,omitempty"`
	LimitDetail []*ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetail `json:"LimitDetail,omitempty" xml:"LimitDetail,omitempty" type:"Repeated"`
	// example:
	//
	// UserGroupNormal
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// example:
	//
	// registration_policy_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// reg-policy-dcbfd33cb004****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// example:
	//
	// 1
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// Enabled
	Status    *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	Whitelist []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
}

func (s ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies) GoString() string {
	return s.String()
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies) GetDescription() *string {
	return s.Description
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies) GetLimitDetail() []*ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetail {
	return s.LimitDetail
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies) GetMatchMode() *string {
	return s.MatchMode
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies) GetName() *string {
	return s.Name
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies) GetPriority() *int64 {
	return s.Priority
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies) GetStatus() *string {
	return s.Status
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies) GetWhitelist() []*string {
	return s.Whitelist
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies) SetCreateTime(v string) *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies {
	s.CreateTime = &v
	return s
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies) SetDescription(v string) *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies {
	s.Description = &v
	return s
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies) SetLimitDetail(v []*ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetail) *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies {
	s.LimitDetail = v
	return s
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies) SetMatchMode(v string) *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies {
	s.MatchMode = &v
	return s
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies) SetName(v string) *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies {
	s.Name = &v
	return s
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies) SetPolicyId(v string) *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies {
	s.PolicyId = &v
	return s
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies) SetPriority(v int64) *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies {
	s.Priority = &v
	return s
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies) SetStatus(v string) *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies {
	s.Status = &v
	return s
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies) SetWhitelist(v []*string) *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies {
	s.Whitelist = v
	return s
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies) Validate() error {
	return dara.Validate(s)
}

type ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetail struct {
	// example:
	//
	// Company
	DeviceBelong *string                                                                                  `json:"DeviceBelong,omitempty" xml:"DeviceBelong,omitempty"`
	LimitCount   *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetailLimitCount `json:"LimitCount,omitempty" xml:"LimitCount,omitempty" type:"Struct"`
	// example:
	//
	// LimitAll
	LimitType *string `json:"LimitType,omitempty" xml:"LimitType,omitempty"`
}

func (s ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetail) String() string {
	return dara.Prettify(s)
}

func (s ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetail) GoString() string {
	return s.String()
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetail) GetDeviceBelong() *string {
	return s.DeviceBelong
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetail) GetLimitCount() *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetailLimitCount {
	return s.LimitCount
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetail) GetLimitType() *string {
	return s.LimitType
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetail) SetDeviceBelong(v string) *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetail {
	s.DeviceBelong = &v
	return s
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetail) SetLimitCount(v *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetailLimitCount) *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetail {
	s.LimitCount = v
	return s
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetail) SetLimitType(v string) *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetail {
	s.LimitType = &v
	return s
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetail) Validate() error {
	return dara.Validate(s)
}

type ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetailLimitCount struct {
	// example:
	//
	// 3
	All *string `json:"All,omitempty" xml:"All,omitempty"`
	// example:
	//
	// 0
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// example:
	//
	// 0
	PC *string `json:"PC,omitempty" xml:"PC,omitempty"`
}

func (s ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetailLimitCount) String() string {
	return dara.Prettify(s)
}

func (s ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetailLimitCount) GoString() string {
	return s.String()
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetailLimitCount) GetAll() *string {
	return s.All
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetailLimitCount) GetMobile() *string {
	return s.Mobile
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetailLimitCount) GetPC() *string {
	return s.PC
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetailLimitCount) SetAll(v string) *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetailLimitCount {
	s.All = &v
	return s
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetailLimitCount) SetMobile(v string) *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetailLimitCount {
	s.Mobile = &v
	return s
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetailLimitCount) SetPC(v string) *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetailLimitCount {
	s.PC = &v
	return s
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetailLimitCount) Validate() error {
	return dara.Validate(s)
}
