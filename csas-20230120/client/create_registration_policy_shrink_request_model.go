// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRegistrationPolicyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyLimitCountShrink(v string) *CreateRegistrationPolicyShrinkRequest
	GetCompanyLimitCountShrink() *string
	SetCompanyLimitType(v string) *CreateRegistrationPolicyShrinkRequest
	GetCompanyLimitType() *string
	SetDescription(v string) *CreateRegistrationPolicyShrinkRequest
	GetDescription() *string
	SetMatchMode(v string) *CreateRegistrationPolicyShrinkRequest
	GetMatchMode() *string
	SetName(v string) *CreateRegistrationPolicyShrinkRequest
	GetName() *string
	SetPersonalLimitCountShrink(v string) *CreateRegistrationPolicyShrinkRequest
	GetPersonalLimitCountShrink() *string
	SetPersonalLimitType(v string) *CreateRegistrationPolicyShrinkRequest
	GetPersonalLimitType() *string
	SetPriority(v int64) *CreateRegistrationPolicyShrinkRequest
	GetPriority() *int64
	SetStatus(v string) *CreateRegistrationPolicyShrinkRequest
	GetStatus() *string
	SetUserGroupIds(v []*string) *CreateRegistrationPolicyShrinkRequest
	GetUserGroupIds() []*string
	SetWhitelist(v []*string) *CreateRegistrationPolicyShrinkRequest
	GetWhitelist() []*string
}

type CreateRegistrationPolicyShrinkRequest struct {
	// The restriction count for company devices.
	CompanyLimitCountShrink *string `json:"CompanyLimitCount,omitempty" xml:"CompanyLimitCount,omitempty"`
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
	PersonalLimitCountShrink *string `json:"PersonalLimitCount,omitempty" xml:"PersonalLimitCount,omitempty"`
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

func (s CreateRegistrationPolicyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRegistrationPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateRegistrationPolicyShrinkRequest) GetCompanyLimitCountShrink() *string {
	return s.CompanyLimitCountShrink
}

func (s *CreateRegistrationPolicyShrinkRequest) GetCompanyLimitType() *string {
	return s.CompanyLimitType
}

func (s *CreateRegistrationPolicyShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRegistrationPolicyShrinkRequest) GetMatchMode() *string {
	return s.MatchMode
}

func (s *CreateRegistrationPolicyShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateRegistrationPolicyShrinkRequest) GetPersonalLimitCountShrink() *string {
	return s.PersonalLimitCountShrink
}

func (s *CreateRegistrationPolicyShrinkRequest) GetPersonalLimitType() *string {
	return s.PersonalLimitType
}

func (s *CreateRegistrationPolicyShrinkRequest) GetPriority() *int64 {
	return s.Priority
}

func (s *CreateRegistrationPolicyShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *CreateRegistrationPolicyShrinkRequest) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *CreateRegistrationPolicyShrinkRequest) GetWhitelist() []*string {
	return s.Whitelist
}

func (s *CreateRegistrationPolicyShrinkRequest) SetCompanyLimitCountShrink(v string) *CreateRegistrationPolicyShrinkRequest {
	s.CompanyLimitCountShrink = &v
	return s
}

func (s *CreateRegistrationPolicyShrinkRequest) SetCompanyLimitType(v string) *CreateRegistrationPolicyShrinkRequest {
	s.CompanyLimitType = &v
	return s
}

func (s *CreateRegistrationPolicyShrinkRequest) SetDescription(v string) *CreateRegistrationPolicyShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateRegistrationPolicyShrinkRequest) SetMatchMode(v string) *CreateRegistrationPolicyShrinkRequest {
	s.MatchMode = &v
	return s
}

func (s *CreateRegistrationPolicyShrinkRequest) SetName(v string) *CreateRegistrationPolicyShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateRegistrationPolicyShrinkRequest) SetPersonalLimitCountShrink(v string) *CreateRegistrationPolicyShrinkRequest {
	s.PersonalLimitCountShrink = &v
	return s
}

func (s *CreateRegistrationPolicyShrinkRequest) SetPersonalLimitType(v string) *CreateRegistrationPolicyShrinkRequest {
	s.PersonalLimitType = &v
	return s
}

func (s *CreateRegistrationPolicyShrinkRequest) SetPriority(v int64) *CreateRegistrationPolicyShrinkRequest {
	s.Priority = &v
	return s
}

func (s *CreateRegistrationPolicyShrinkRequest) SetStatus(v string) *CreateRegistrationPolicyShrinkRequest {
	s.Status = &v
	return s
}

func (s *CreateRegistrationPolicyShrinkRequest) SetUserGroupIds(v []*string) *CreateRegistrationPolicyShrinkRequest {
	s.UserGroupIds = v
	return s
}

func (s *CreateRegistrationPolicyShrinkRequest) SetWhitelist(v []*string) *CreateRegistrationPolicyShrinkRequest {
	s.Whitelist = v
	return s
}

func (s *CreateRegistrationPolicyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
