// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRegistrationPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyLimitType(v string) *ListRegistrationPoliciesRequest
	GetCompanyLimitType() *string
	SetCurrentPage(v int64) *ListRegistrationPoliciesRequest
	GetCurrentPage() *int64
	SetMatchMode(v string) *ListRegistrationPoliciesRequest
	GetMatchMode() *string
	SetName(v string) *ListRegistrationPoliciesRequest
	GetName() *string
	SetPageSize(v int64) *ListRegistrationPoliciesRequest
	GetPageSize() *int64
	SetPersonalLimitType(v string) *ListRegistrationPoliciesRequest
	GetPersonalLimitType() *string
	SetPolicyIds(v []*string) *ListRegistrationPoliciesRequest
	GetPolicyIds() []*string
	SetStatus(v string) *ListRegistrationPoliciesRequest
	GetStatus() *string
	SetUserGroupId(v string) *ListRegistrationPoliciesRequest
	GetUserGroupId() *string
}

type ListRegistrationPoliciesRequest struct {
	// The type of company device registration limit. Valid values:
	//
	// - **Unlimited**: No limit.
	//
	// - **LimitAll**: Limit by total number.
	//
	// - **LimitDiff**: Limit by device categorization.
	//
	// example:
	//
	// LimitAll
	CompanyLimitType *string `json:"CompanyLimitType,omitempty" xml:"CompanyLimitType,omitempty"`
	// The current page number for paged queries. Values range from 1 to 10000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The policy matching target type. Valid values:
	//
	// - **UserGroupAll**: Associate all users.
	//
	// - **UserGroupNormal**: Associate some user groups.
	//
	// example:
	//
	// UserGroupAll
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// The name of the device registration policy. It can be 1 to 128 characters long. It supports Chinese characters, uppercase and lowercase English letters, numbers, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// registration_policy_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of items per page for paged queries. Values range from 1 to 500.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of personal device registration limit. Valid values:
	//
	// - **Unlimited**: No limit.
	//
	// - **LimitAll**: Limit by total number.
	//
	// - **LimitDiff**: Limit by device categorization.
	//
	// example:
	//
	// LimitDiff
	PersonalLimitType *string `json:"PersonalLimitType,omitempty" xml:"PersonalLimitType,omitempty"`
	// A collection of device registration policy IDs. You can enter up to 100 device registration policy IDs.
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// The status of the device registration policy. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The user group ID. Get this value from:
	//
	// - [ListUserGroups](~~ListUserGroups~~): Batch query user groups.
	//
	// - [CreateUserGroup](~~CreateUserGroup~~): Create user groups.
	//
	// example:
	//
	// usergroup-6f1ef2fc56b6****
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s ListRegistrationPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRegistrationPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListRegistrationPoliciesRequest) GetCompanyLimitType() *string {
	return s.CompanyLimitType
}

func (s *ListRegistrationPoliciesRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListRegistrationPoliciesRequest) GetMatchMode() *string {
	return s.MatchMode
}

func (s *ListRegistrationPoliciesRequest) GetName() *string {
	return s.Name
}

func (s *ListRegistrationPoliciesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListRegistrationPoliciesRequest) GetPersonalLimitType() *string {
	return s.PersonalLimitType
}

func (s *ListRegistrationPoliciesRequest) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *ListRegistrationPoliciesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListRegistrationPoliciesRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *ListRegistrationPoliciesRequest) SetCompanyLimitType(v string) *ListRegistrationPoliciesRequest {
	s.CompanyLimitType = &v
	return s
}

func (s *ListRegistrationPoliciesRequest) SetCurrentPage(v int64) *ListRegistrationPoliciesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListRegistrationPoliciesRequest) SetMatchMode(v string) *ListRegistrationPoliciesRequest {
	s.MatchMode = &v
	return s
}

func (s *ListRegistrationPoliciesRequest) SetName(v string) *ListRegistrationPoliciesRequest {
	s.Name = &v
	return s
}

func (s *ListRegistrationPoliciesRequest) SetPageSize(v int64) *ListRegistrationPoliciesRequest {
	s.PageSize = &v
	return s
}

func (s *ListRegistrationPoliciesRequest) SetPersonalLimitType(v string) *ListRegistrationPoliciesRequest {
	s.PersonalLimitType = &v
	return s
}

func (s *ListRegistrationPoliciesRequest) SetPolicyIds(v []*string) *ListRegistrationPoliciesRequest {
	s.PolicyIds = v
	return s
}

func (s *ListRegistrationPoliciesRequest) SetStatus(v string) *ListRegistrationPoliciesRequest {
	s.Status = &v
	return s
}

func (s *ListRegistrationPoliciesRequest) SetUserGroupId(v string) *ListRegistrationPoliciesRequest {
	s.UserGroupId = &v
	return s
}

func (s *ListRegistrationPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
