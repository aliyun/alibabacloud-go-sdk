// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRegistrationPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyLimitCount(v *UpdateRegistrationPolicyRequestCompanyLimitCount) *UpdateRegistrationPolicyRequest
	GetCompanyLimitCount() *UpdateRegistrationPolicyRequestCompanyLimitCount
	SetCompanyLimitType(v string) *UpdateRegistrationPolicyRequest
	GetCompanyLimitType() *string
	SetDescription(v string) *UpdateRegistrationPolicyRequest
	GetDescription() *string
	SetMatchMode(v string) *UpdateRegistrationPolicyRequest
	GetMatchMode() *string
	SetName(v string) *UpdateRegistrationPolicyRequest
	GetName() *string
	SetPersonalLimitCount(v *UpdateRegistrationPolicyRequestPersonalLimitCount) *UpdateRegistrationPolicyRequest
	GetPersonalLimitCount() *UpdateRegistrationPolicyRequestPersonalLimitCount
	SetPersonalLimitType(v string) *UpdateRegistrationPolicyRequest
	GetPersonalLimitType() *string
	SetPolicyId(v string) *UpdateRegistrationPolicyRequest
	GetPolicyId() *string
	SetPriority(v int64) *UpdateRegistrationPolicyRequest
	GetPriority() *int64
	SetStatus(v string) *UpdateRegistrationPolicyRequest
	GetStatus() *string
	SetUserGroupIds(v []*string) *UpdateRegistrationPolicyRequest
	GetUserGroupIds() []*string
	SetWhitelist(v []*string) *UpdateRegistrationPolicyRequest
	GetWhitelist() []*string
}

type UpdateRegistrationPolicyRequest struct {
	// The registration limit for corporate devices.
	CompanyLimitCount *UpdateRegistrationPolicyRequestCompanyLimitCount `json:"CompanyLimitCount,omitempty" xml:"CompanyLimitCount,omitempty" type:"Struct"`
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
	PersonalLimitCount *UpdateRegistrationPolicyRequestPersonalLimitCount `json:"PersonalLimitCount,omitempty" xml:"PersonalLimitCount,omitempty" type:"Struct"`
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

func (s UpdateRegistrationPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRegistrationPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateRegistrationPolicyRequest) GetCompanyLimitCount() *UpdateRegistrationPolicyRequestCompanyLimitCount {
	return s.CompanyLimitCount
}

func (s *UpdateRegistrationPolicyRequest) GetCompanyLimitType() *string {
	return s.CompanyLimitType
}

func (s *UpdateRegistrationPolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateRegistrationPolicyRequest) GetMatchMode() *string {
	return s.MatchMode
}

func (s *UpdateRegistrationPolicyRequest) GetName() *string {
	return s.Name
}

func (s *UpdateRegistrationPolicyRequest) GetPersonalLimitCount() *UpdateRegistrationPolicyRequestPersonalLimitCount {
	return s.PersonalLimitCount
}

func (s *UpdateRegistrationPolicyRequest) GetPersonalLimitType() *string {
	return s.PersonalLimitType
}

func (s *UpdateRegistrationPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *UpdateRegistrationPolicyRequest) GetPriority() *int64 {
	return s.Priority
}

func (s *UpdateRegistrationPolicyRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateRegistrationPolicyRequest) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *UpdateRegistrationPolicyRequest) GetWhitelist() []*string {
	return s.Whitelist
}

func (s *UpdateRegistrationPolicyRequest) SetCompanyLimitCount(v *UpdateRegistrationPolicyRequestCompanyLimitCount) *UpdateRegistrationPolicyRequest {
	s.CompanyLimitCount = v
	return s
}

func (s *UpdateRegistrationPolicyRequest) SetCompanyLimitType(v string) *UpdateRegistrationPolicyRequest {
	s.CompanyLimitType = &v
	return s
}

func (s *UpdateRegistrationPolicyRequest) SetDescription(v string) *UpdateRegistrationPolicyRequest {
	s.Description = &v
	return s
}

func (s *UpdateRegistrationPolicyRequest) SetMatchMode(v string) *UpdateRegistrationPolicyRequest {
	s.MatchMode = &v
	return s
}

func (s *UpdateRegistrationPolicyRequest) SetName(v string) *UpdateRegistrationPolicyRequest {
	s.Name = &v
	return s
}

func (s *UpdateRegistrationPolicyRequest) SetPersonalLimitCount(v *UpdateRegistrationPolicyRequestPersonalLimitCount) *UpdateRegistrationPolicyRequest {
	s.PersonalLimitCount = v
	return s
}

func (s *UpdateRegistrationPolicyRequest) SetPersonalLimitType(v string) *UpdateRegistrationPolicyRequest {
	s.PersonalLimitType = &v
	return s
}

func (s *UpdateRegistrationPolicyRequest) SetPolicyId(v string) *UpdateRegistrationPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *UpdateRegistrationPolicyRequest) SetPriority(v int64) *UpdateRegistrationPolicyRequest {
	s.Priority = &v
	return s
}

func (s *UpdateRegistrationPolicyRequest) SetStatus(v string) *UpdateRegistrationPolicyRequest {
	s.Status = &v
	return s
}

func (s *UpdateRegistrationPolicyRequest) SetUserGroupIds(v []*string) *UpdateRegistrationPolicyRequest {
	s.UserGroupIds = v
	return s
}

func (s *UpdateRegistrationPolicyRequest) SetWhitelist(v []*string) *UpdateRegistrationPolicyRequest {
	s.Whitelist = v
	return s
}

func (s *UpdateRegistrationPolicyRequest) Validate() error {
	if s.CompanyLimitCount != nil {
		if err := s.CompanyLimitCount.Validate(); err != nil {
			return err
		}
	}
	if s.PersonalLimitCount != nil {
		if err := s.PersonalLimitCount.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateRegistrationPolicyRequestCompanyLimitCount struct {
	// The total number of corporate devices that can be registered. The value can be from 0 to 100. The default value is 0. This parameter is valid only when CompanyLimitType is set to **LimitAll**.
	//
	// example:
	//
	// 1
	All *int32 `json:"All,omitempty" xml:"All,omitempty"`
	// The number of corporate mobile devices that can be registered. The value can be from 0 to 100. The default value is 0. This parameter is valid only when CompanyLimitType is set to **LimitDiff**.
	//
	// example:
	//
	// 0
	Mobile *int32 `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The number of corporate PCs that can be registered. The value can be from 0 to 100. The default value is 0. This parameter is valid only when CompanyLimitType is set to **LimitDiff**.
	//
	// example:
	//
	// 0
	PC *int32 `json:"PC,omitempty" xml:"PC,omitempty"`
}

func (s UpdateRegistrationPolicyRequestCompanyLimitCount) String() string {
	return dara.Prettify(s)
}

func (s UpdateRegistrationPolicyRequestCompanyLimitCount) GoString() string {
	return s.String()
}

func (s *UpdateRegistrationPolicyRequestCompanyLimitCount) GetAll() *int32 {
	return s.All
}

func (s *UpdateRegistrationPolicyRequestCompanyLimitCount) GetMobile() *int32 {
	return s.Mobile
}

func (s *UpdateRegistrationPolicyRequestCompanyLimitCount) GetPC() *int32 {
	return s.PC
}

func (s *UpdateRegistrationPolicyRequestCompanyLimitCount) SetAll(v int32) *UpdateRegistrationPolicyRequestCompanyLimitCount {
	s.All = &v
	return s
}

func (s *UpdateRegistrationPolicyRequestCompanyLimitCount) SetMobile(v int32) *UpdateRegistrationPolicyRequestCompanyLimitCount {
	s.Mobile = &v
	return s
}

func (s *UpdateRegistrationPolicyRequestCompanyLimitCount) SetPC(v int32) *UpdateRegistrationPolicyRequestCompanyLimitCount {
	s.PC = &v
	return s
}

func (s *UpdateRegistrationPolicyRequestCompanyLimitCount) Validate() error {
	return dara.Validate(s)
}

type UpdateRegistrationPolicyRequestPersonalLimitCount struct {
	// The total number of personal devices that can be registered. The value can be from 0 to 100. The default value is 0. This parameter is valid only when PersonalLimitType is set to **LimitAll**.
	//
	// example:
	//
	// 0
	All *int32 `json:"All,omitempty" xml:"All,omitempty"`
	// The number of personal mobile devices that can be registered. The value can be from 0 to 100. The default value is 0. This parameter is valid only when PersonalLimitType is set to **LimitDiff**.
	//
	// example:
	//
	// 1
	Mobile *int32 `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The number of personal PCs that can be registered. The value can be from 0 to 100. The default value is 0. This parameter is valid only when PersonalLimitType is set to **LimitDiff**.
	//
	// example:
	//
	// 2
	PC *int32 `json:"PC,omitempty" xml:"PC,omitempty"`
}

func (s UpdateRegistrationPolicyRequestPersonalLimitCount) String() string {
	return dara.Prettify(s)
}

func (s UpdateRegistrationPolicyRequestPersonalLimitCount) GoString() string {
	return s.String()
}

func (s *UpdateRegistrationPolicyRequestPersonalLimitCount) GetAll() *int32 {
	return s.All
}

func (s *UpdateRegistrationPolicyRequestPersonalLimitCount) GetMobile() *int32 {
	return s.Mobile
}

func (s *UpdateRegistrationPolicyRequestPersonalLimitCount) GetPC() *int32 {
	return s.PC
}

func (s *UpdateRegistrationPolicyRequestPersonalLimitCount) SetAll(v int32) *UpdateRegistrationPolicyRequestPersonalLimitCount {
	s.All = &v
	return s
}

func (s *UpdateRegistrationPolicyRequestPersonalLimitCount) SetMobile(v int32) *UpdateRegistrationPolicyRequestPersonalLimitCount {
	s.Mobile = &v
	return s
}

func (s *UpdateRegistrationPolicyRequestPersonalLimitCount) SetPC(v int32) *UpdateRegistrationPolicyRequestPersonalLimitCount {
	s.PC = &v
	return s
}

func (s *UpdateRegistrationPolicyRequestPersonalLimitCount) Validate() error {
	return dara.Validate(s)
}
