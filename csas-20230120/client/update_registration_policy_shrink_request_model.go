// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRegistrationPolicyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyLimitCountShrink(v string) *UpdateRegistrationPolicyShrinkRequest
	GetCompanyLimitCountShrink() *string
	SetCompanyLimitType(v string) *UpdateRegistrationPolicyShrinkRequest
	GetCompanyLimitType() *string
	SetDescription(v string) *UpdateRegistrationPolicyShrinkRequest
	GetDescription() *string
	SetMatchMode(v string) *UpdateRegistrationPolicyShrinkRequest
	GetMatchMode() *string
	SetName(v string) *UpdateRegistrationPolicyShrinkRequest
	GetName() *string
	SetPersonalLimitCountShrink(v string) *UpdateRegistrationPolicyShrinkRequest
	GetPersonalLimitCountShrink() *string
	SetPersonalLimitType(v string) *UpdateRegistrationPolicyShrinkRequest
	GetPersonalLimitType() *string
	SetPolicyId(v string) *UpdateRegistrationPolicyShrinkRequest
	GetPolicyId() *string
	SetPriority(v int64) *UpdateRegistrationPolicyShrinkRequest
	GetPriority() *int64
	SetStatus(v string) *UpdateRegistrationPolicyShrinkRequest
	GetStatus() *string
	SetUserGroupIds(v []*string) *UpdateRegistrationPolicyShrinkRequest
	GetUserGroupIds() []*string
	SetWhitelist(v []*string) *UpdateRegistrationPolicyShrinkRequest
	GetWhitelist() []*string
}

type UpdateRegistrationPolicyShrinkRequest struct {
	// The registration limit for corporate devices.
	CompanyLimitCountShrink *string `json:"CompanyLimitCount,omitempty" xml:"CompanyLimitCount,omitempty"`
	// The registration limit type for corporate devices. Valid values:
	//
	// - **Unlimited**: No limit.
	//
	// - **LimitAll**: Limits the total number of devices.
	//
	// - **LimitDiff**: Limits devices by terminal type.
	//
	// example:
	//
	// LimitAll
	CompanyLimitType *string `json:"CompanyLimitType,omitempty" xml:"CompanyLimitType,omitempty"`
	// The description of the device registration policy. The description can be 1 to 128 characters long and can contain Chinese characters, uppercase and lowercase letters, digits, periods (.), underscores (_), hyphens (-), and spaces.
	//
	// example:
	//
	// 这是一条设备注册策略
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The matching target type of the policy. Valid values:
	//
	// - **UserGroupAll**: Associates with all users.
	//
	// - **UserGroupNormal**: Associates with specific user groups.
	//
	// example:
	//
	// UserGroupNormal
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// The name of the device registration policy. The name must be 1 to 128 characters in length and can contain Chinese characters, uppercase and lowercase letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// registration_policy_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The registration limit for personal devices.
	PersonalLimitCountShrink *string `json:"PersonalLimitCount,omitempty" xml:"PersonalLimitCount,omitempty"`
	// The registration limit type for personal devices. Valid values:
	//
	// - **Unlimited**: No limit.
	//
	// - **LimitAll**: Limits the total number of devices.
	//
	// - **LimitDiff**: Limits devices by terminal type.
	//
	// example:
	//
	// LimitDiff
	PersonalLimitType *string `json:"PersonalLimitType,omitempty" xml:"PersonalLimitType,omitempty"`
	// The ID of the device registration policy. You can obtain the ID by calling one of the following operations:
	//
	// - [ListRegistrationPolicies](~~ListRegistrationPolicies~~)
	//
	// - [GetRegistrationPolicy](~~GetRegistrationPolicy~~)
	//
	// - [CreateRegistrationPolicy](~~CreateRegistrationPolicy~~)
	//
	// - [UpdateRegistrationPolicy](~~UpdateRegistrationPolicy~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// reg-policy-63b2f1844b86****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The priority of the device registration policy. A smaller value indicates a higher priority. The value 0 indicates the highest priority, and 99 indicates the lowest priority.
	//
	// example:
	//
	// 0
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The status of the device registration policy. Valid values:
	//
	// - **Enabled**
	//
	// - **Disabled**
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The IDs of user groups. This parameter is required when MatchMode is set to **UserGroupNormal**. A policy can be associated with up to 100 user groups.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// The list of whitelisted users for the device registration policy. You can add up to 1,000 usernames.
	Whitelist []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
}

func (s UpdateRegistrationPolicyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRegistrationPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateRegistrationPolicyShrinkRequest) GetCompanyLimitCountShrink() *string {
	return s.CompanyLimitCountShrink
}

func (s *UpdateRegistrationPolicyShrinkRequest) GetCompanyLimitType() *string {
	return s.CompanyLimitType
}

func (s *UpdateRegistrationPolicyShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateRegistrationPolicyShrinkRequest) GetMatchMode() *string {
	return s.MatchMode
}

func (s *UpdateRegistrationPolicyShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateRegistrationPolicyShrinkRequest) GetPersonalLimitCountShrink() *string {
	return s.PersonalLimitCountShrink
}

func (s *UpdateRegistrationPolicyShrinkRequest) GetPersonalLimitType() *string {
	return s.PersonalLimitType
}

func (s *UpdateRegistrationPolicyShrinkRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *UpdateRegistrationPolicyShrinkRequest) GetPriority() *int64 {
	return s.Priority
}

func (s *UpdateRegistrationPolicyShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateRegistrationPolicyShrinkRequest) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *UpdateRegistrationPolicyShrinkRequest) GetWhitelist() []*string {
	return s.Whitelist
}

func (s *UpdateRegistrationPolicyShrinkRequest) SetCompanyLimitCountShrink(v string) *UpdateRegistrationPolicyShrinkRequest {
	s.CompanyLimitCountShrink = &v
	return s
}

func (s *UpdateRegistrationPolicyShrinkRequest) SetCompanyLimitType(v string) *UpdateRegistrationPolicyShrinkRequest {
	s.CompanyLimitType = &v
	return s
}

func (s *UpdateRegistrationPolicyShrinkRequest) SetDescription(v string) *UpdateRegistrationPolicyShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateRegistrationPolicyShrinkRequest) SetMatchMode(v string) *UpdateRegistrationPolicyShrinkRequest {
	s.MatchMode = &v
	return s
}

func (s *UpdateRegistrationPolicyShrinkRequest) SetName(v string) *UpdateRegistrationPolicyShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateRegistrationPolicyShrinkRequest) SetPersonalLimitCountShrink(v string) *UpdateRegistrationPolicyShrinkRequest {
	s.PersonalLimitCountShrink = &v
	return s
}

func (s *UpdateRegistrationPolicyShrinkRequest) SetPersonalLimitType(v string) *UpdateRegistrationPolicyShrinkRequest {
	s.PersonalLimitType = &v
	return s
}

func (s *UpdateRegistrationPolicyShrinkRequest) SetPolicyId(v string) *UpdateRegistrationPolicyShrinkRequest {
	s.PolicyId = &v
	return s
}

func (s *UpdateRegistrationPolicyShrinkRequest) SetPriority(v int64) *UpdateRegistrationPolicyShrinkRequest {
	s.Priority = &v
	return s
}

func (s *UpdateRegistrationPolicyShrinkRequest) SetStatus(v string) *UpdateRegistrationPolicyShrinkRequest {
	s.Status = &v
	return s
}

func (s *UpdateRegistrationPolicyShrinkRequest) SetUserGroupIds(v []*string) *UpdateRegistrationPolicyShrinkRequest {
	s.UserGroupIds = v
	return s
}

func (s *UpdateRegistrationPolicyShrinkRequest) SetWhitelist(v []*string) *UpdateRegistrationPolicyShrinkRequest {
	s.Whitelist = v
	return s
}

func (s *UpdateRegistrationPolicyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
