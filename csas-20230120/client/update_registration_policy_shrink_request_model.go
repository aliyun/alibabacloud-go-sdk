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
	CompanyLimitCountShrink *string `json:"CompanyLimitCount,omitempty" xml:"CompanyLimitCount,omitempty"`
	// example:
	//
	// LimitAll
	CompanyLimitType *string `json:"CompanyLimitType,omitempty" xml:"CompanyLimitType,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// UserGroupNormal
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// example:
	//
	// registration_policy_name
	Name                     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PersonalLimitCountShrink *string `json:"PersonalLimitCount,omitempty" xml:"PersonalLimitCount,omitempty"`
	// example:
	//
	// LimitDiff
	PersonalLimitType *string `json:"PersonalLimitType,omitempty" xml:"PersonalLimitType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// reg-policy-63b2f1844b86****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// example:
	//
	// 0
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// Enabled
	Status       *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	Whitelist    []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
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
