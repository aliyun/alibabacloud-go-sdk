// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRegistrationPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyLimitCount(v *CreateRegistrationPolicyRequestCompanyLimitCount) *CreateRegistrationPolicyRequest
	GetCompanyLimitCount() *CreateRegistrationPolicyRequestCompanyLimitCount
	SetCompanyLimitType(v string) *CreateRegistrationPolicyRequest
	GetCompanyLimitType() *string
	SetDescription(v string) *CreateRegistrationPolicyRequest
	GetDescription() *string
	SetMatchMode(v string) *CreateRegistrationPolicyRequest
	GetMatchMode() *string
	SetName(v string) *CreateRegistrationPolicyRequest
	GetName() *string
	SetPersonalLimitCount(v *CreateRegistrationPolicyRequestPersonalLimitCount) *CreateRegistrationPolicyRequest
	GetPersonalLimitCount() *CreateRegistrationPolicyRequestPersonalLimitCount
	SetPersonalLimitType(v string) *CreateRegistrationPolicyRequest
	GetPersonalLimitType() *string
	SetPriority(v int64) *CreateRegistrationPolicyRequest
	GetPriority() *int64
	SetStatus(v string) *CreateRegistrationPolicyRequest
	GetStatus() *string
	SetUserGroupIds(v []*string) *CreateRegistrationPolicyRequest
	GetUserGroupIds() []*string
	SetWhitelist(v []*string) *CreateRegistrationPolicyRequest
	GetWhitelist() []*string
}

type CreateRegistrationPolicyRequest struct {
	// The restriction count for company devices.
	CompanyLimitCount *CreateRegistrationPolicyRequestCompanyLimitCount `json:"CompanyLimitCount,omitempty" xml:"CompanyLimitCount,omitempty" type:"Struct"`
	// The restriction type for company devices. Valid values:
	//
	// - **Unlimited**: No restrictions.
	//
	// - **LimitAll**: Limit by total count.
	//
	// - **LimitDiff**: Limit by device category.
	//
	// This parameter is required.
	//
	// example:
	//
	// LimitAll
	CompanyLimitType *string `json:"CompanyLimitType,omitempty" xml:"CompanyLimitType,omitempty"`
	// A description of the device registration policy. The description must be 1 to 128 characters in length. It can contain letters, digits, periods (.), underscores (_), hyphens (-), and spaces.
	//
	// example:
	//
	// 这是一条设备注册策略
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The target type for policy matching. Valid values:
	//
	// - **UserGroupAll**: Apply to all users.
	//
	// - **UserGroupNormal**: Apply to selected user groups.
	//
	// This parameter is required.
	//
	// example:
	//
	// UserGroupAll
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// The name of the device registration policy. The name must be 1 to 128 characters in length. It can contain letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// registration_policy_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The restriction count for personal devices.
	PersonalLimitCount *CreateRegistrationPolicyRequestPersonalLimitCount `json:"PersonalLimitCount,omitempty" xml:"PersonalLimitCount,omitempty" type:"Struct"`
	// The restriction type for personal devices. Valid values:
	//
	// - **Unlimited**: No restrictions.
	//
	// - **LimitAll**: Limit by total count.
	//
	// - **LimitDiff**: Limit by device category.
	//
	// This parameter is required.
	//
	// example:
	//
	// LimitDiff
	PersonalLimitType *string `json:"PersonalLimitType,omitempty" xml:"PersonalLimitType,omitempty"`
	// The priority of the device registration policy. A value of 0 indicates the highest priority. A value of 99 indicates the lowest priority.
	//
	// example:
	//
	// 99
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The status of the device registration policy. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The IDs of user groups to which the device registration policy applies. Required if MatchMode is set to **UserGroupNormal**. A maximum of 100 user groups can be specified per policy.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// The list of usernames in the whitelist for the device registration policy. You can specify up to 1,000 usernames.
	Whitelist []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
}

func (s CreateRegistrationPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRegistrationPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateRegistrationPolicyRequest) GetCompanyLimitCount() *CreateRegistrationPolicyRequestCompanyLimitCount {
	return s.CompanyLimitCount
}

func (s *CreateRegistrationPolicyRequest) GetCompanyLimitType() *string {
	return s.CompanyLimitType
}

func (s *CreateRegistrationPolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRegistrationPolicyRequest) GetMatchMode() *string {
	return s.MatchMode
}

func (s *CreateRegistrationPolicyRequest) GetName() *string {
	return s.Name
}

func (s *CreateRegistrationPolicyRequest) GetPersonalLimitCount() *CreateRegistrationPolicyRequestPersonalLimitCount {
	return s.PersonalLimitCount
}

func (s *CreateRegistrationPolicyRequest) GetPersonalLimitType() *string {
	return s.PersonalLimitType
}

func (s *CreateRegistrationPolicyRequest) GetPriority() *int64 {
	return s.Priority
}

func (s *CreateRegistrationPolicyRequest) GetStatus() *string {
	return s.Status
}

func (s *CreateRegistrationPolicyRequest) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *CreateRegistrationPolicyRequest) GetWhitelist() []*string {
	return s.Whitelist
}

func (s *CreateRegistrationPolicyRequest) SetCompanyLimitCount(v *CreateRegistrationPolicyRequestCompanyLimitCount) *CreateRegistrationPolicyRequest {
	s.CompanyLimitCount = v
	return s
}

func (s *CreateRegistrationPolicyRequest) SetCompanyLimitType(v string) *CreateRegistrationPolicyRequest {
	s.CompanyLimitType = &v
	return s
}

func (s *CreateRegistrationPolicyRequest) SetDescription(v string) *CreateRegistrationPolicyRequest {
	s.Description = &v
	return s
}

func (s *CreateRegistrationPolicyRequest) SetMatchMode(v string) *CreateRegistrationPolicyRequest {
	s.MatchMode = &v
	return s
}

func (s *CreateRegistrationPolicyRequest) SetName(v string) *CreateRegistrationPolicyRequest {
	s.Name = &v
	return s
}

func (s *CreateRegistrationPolicyRequest) SetPersonalLimitCount(v *CreateRegistrationPolicyRequestPersonalLimitCount) *CreateRegistrationPolicyRequest {
	s.PersonalLimitCount = v
	return s
}

func (s *CreateRegistrationPolicyRequest) SetPersonalLimitType(v string) *CreateRegistrationPolicyRequest {
	s.PersonalLimitType = &v
	return s
}

func (s *CreateRegistrationPolicyRequest) SetPriority(v int64) *CreateRegistrationPolicyRequest {
	s.Priority = &v
	return s
}

func (s *CreateRegistrationPolicyRequest) SetStatus(v string) *CreateRegistrationPolicyRequest {
	s.Status = &v
	return s
}

func (s *CreateRegistrationPolicyRequest) SetUserGroupIds(v []*string) *CreateRegistrationPolicyRequest {
	s.UserGroupIds = v
	return s
}

func (s *CreateRegistrationPolicyRequest) SetWhitelist(v []*string) *CreateRegistrationPolicyRequest {
	s.Whitelist = v
	return s
}

func (s *CreateRegistrationPolicyRequest) Validate() error {
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

type CreateRegistrationPolicyRequestCompanyLimitCount struct {
	// The total restriction count for company devices. Valid values: 0 to 100. Default value: 0. This parameter takes effect only when CompanyLimitType is set to **LimitAll**.
	//
	// example:
	//
	// 1
	All *int32 `json:"All,omitempty" xml:"All,omitempty"`
	// The restriction count for mobile logins by company devices. Valid values: 0 to 100. Default value: 0. This parameter takes effect only when CompanyLimitType is set to **LimitDiff**.
	//
	// example:
	//
	// 0
	Mobile *int32 `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The restriction count for PC logins by company devices. Valid values: 0 to 100. Default value: 0. This parameter takes effect only when CompanyLimitType is set to **LimitDiff**.
	//
	// example:
	//
	// 0
	PC *int32 `json:"PC,omitempty" xml:"PC,omitempty"`
}

func (s CreateRegistrationPolicyRequestCompanyLimitCount) String() string {
	return dara.Prettify(s)
}

func (s CreateRegistrationPolicyRequestCompanyLimitCount) GoString() string {
	return s.String()
}

func (s *CreateRegistrationPolicyRequestCompanyLimitCount) GetAll() *int32 {
	return s.All
}

func (s *CreateRegistrationPolicyRequestCompanyLimitCount) GetMobile() *int32 {
	return s.Mobile
}

func (s *CreateRegistrationPolicyRequestCompanyLimitCount) GetPC() *int32 {
	return s.PC
}

func (s *CreateRegistrationPolicyRequestCompanyLimitCount) SetAll(v int32) *CreateRegistrationPolicyRequestCompanyLimitCount {
	s.All = &v
	return s
}

func (s *CreateRegistrationPolicyRequestCompanyLimitCount) SetMobile(v int32) *CreateRegistrationPolicyRequestCompanyLimitCount {
	s.Mobile = &v
	return s
}

func (s *CreateRegistrationPolicyRequestCompanyLimitCount) SetPC(v int32) *CreateRegistrationPolicyRequestCompanyLimitCount {
	s.PC = &v
	return s
}

func (s *CreateRegistrationPolicyRequestCompanyLimitCount) Validate() error {
	return dara.Validate(s)
}

type CreateRegistrationPolicyRequestPersonalLimitCount struct {
	// The total restriction count for personal devices. Valid values: 0 to 100. Default value: 0. This parameter takes effect only when PersonalLimitType is set to **LimitAll**.
	//
	// example:
	//
	// 0
	All *int32 `json:"All,omitempty" xml:"All,omitempty"`
	// The restriction count for mobile logins by personal devices. Valid values: 0 to 100. Default value: 0. This parameter takes effect only when PersonalLimitType is set to **LimitDiff**.
	//
	// example:
	//
	// 3
	Mobile *int32 `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The restriction count for PC logins by personal devices. Valid values: 0 to 100. Default value: 0. This parameter takes effect only when PersonalLimitType is set to **LimitDiff**.
	//
	// example:
	//
	// 2
	PC *int32 `json:"PC,omitempty" xml:"PC,omitempty"`
}

func (s CreateRegistrationPolicyRequestPersonalLimitCount) String() string {
	return dara.Prettify(s)
}

func (s CreateRegistrationPolicyRequestPersonalLimitCount) GoString() string {
	return s.String()
}

func (s *CreateRegistrationPolicyRequestPersonalLimitCount) GetAll() *int32 {
	return s.All
}

func (s *CreateRegistrationPolicyRequestPersonalLimitCount) GetMobile() *int32 {
	return s.Mobile
}

func (s *CreateRegistrationPolicyRequestPersonalLimitCount) GetPC() *int32 {
	return s.PC
}

func (s *CreateRegistrationPolicyRequestPersonalLimitCount) SetAll(v int32) *CreateRegistrationPolicyRequestPersonalLimitCount {
	s.All = &v
	return s
}

func (s *CreateRegistrationPolicyRequestPersonalLimitCount) SetMobile(v int32) *CreateRegistrationPolicyRequestPersonalLimitCount {
	s.Mobile = &v
	return s
}

func (s *CreateRegistrationPolicyRequestPersonalLimitCount) SetPC(v int32) *CreateRegistrationPolicyRequestPersonalLimitCount {
	s.PC = &v
	return s
}

func (s *CreateRegistrationPolicyRequestPersonalLimitCount) Validate() error {
	return dara.Validate(s)
}
