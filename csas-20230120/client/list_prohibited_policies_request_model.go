// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProhibitedPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *ListProhibitedPoliciesRequest
	GetCurrentPage() *int64
	SetEnabled(v bool) *ListProhibitedPoliciesRequest
	GetEnabled() *bool
	SetMatchMode(v string) *ListProhibitedPoliciesRequest
	GetMatchMode() *string
	SetName(v string) *ListProhibitedPoliciesRequest
	GetName() *string
	SetObjectType(v string) *ListProhibitedPoliciesRequest
	GetObjectType() *string
	SetPageSize(v int64) *ListProhibitedPoliciesRequest
	GetPageSize() *int64
	SetPolicyIds(v []*string) *ListProhibitedPoliciesRequest
	GetPolicyIds() []*string
	SetPolicyType(v string) *ListProhibitedPoliciesRequest
	GetPolicyType() *string
	SetSoftwareId(v *ListProhibitedPoliciesRequestSoftwareId) *ListProhibitedPoliciesRequest
	GetSoftwareId() *ListProhibitedPoliciesRequestSoftwareId
	SetSoftwareName(v string) *ListProhibitedPoliciesRequest
	GetSoftwareName() *string
	SetTagId(v string) *ListProhibitedPoliciesRequest
	GetTagId() *string
	SetTagName(v string) *ListProhibitedPoliciesRequest
	GetTagName() *string
	SetUserGroupId(v string) *ListProhibitedPoliciesRequest
	GetUserGroupId() *string
}

type ListProhibitedPoliciesRequest struct {
	// The page number of the current page in a paged query. Valid values: 1 to 10000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Specifies whether the policy is enabled. Valid values:
	//
	// - **true**: Enabled. The policy is delivered to endpoints and takes effect.
	//
	// - **false**: Disabled. The policy configuration is retained but not delivered to endpoints.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The effective scope. Valid values:
	//
	// - **UserGroupAll**: Applies to all users under the current Alibaba Cloud account. No user group needs to be specified.
	//
	// - **UserGroupNormal**: Applies only to users in the user groups specified by UserGroupIds.
	//
	// example:
	//
	// UserGroupAll
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// Policy Name of the software prohibition policy. Fuzzy match is supported. Policy Name can be up to 128 characters in length and can contain Chinese characters, uppercase and lowercase letters, digits, periods (.), underscores (_), and hyphens (-). Spaces are not supported.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The object type of the controlled target. Valid values:
	//
	// - **App**: Controls by prohibited software. The controlled objects are specified by SoftwareIds.
	//
	// - **Tag**: Controls by prohibited software tag. The controlled objects are specified by TagIds. All prohibited software under the tag is controlled.
	//
	// example:
	//
	// App
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The number of entries per page in a paged query. Valid values: 1 to 500.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The collection of software prohibition policy IDs. Duplicate values are not allowed.
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// The action to take. Valid values:
	//
	// - **Ban**: Blocks the software from running and displays a pop-up notification on the endpoint to alert the user.
	//
	// - **BanSilent**: Blocks the software from running without notifying the user. The blocking is silent.
	//
	// - **Warn**: Only displays a pop-up notification on the endpoint to alert the user without blocking the software from running.
	//
	// example:
	//
	// Ban
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The unique identifier of the prohibited software.
	SoftwareId *ListProhibitedPoliciesRequestSoftwareId `json:"SoftwareId,omitempty" xml:"SoftwareId,omitempty" type:"Struct"`
	// The name of the prohibited software. Fuzzy match is supported. The name can be up to 128 characters in length and can contain Chinese characters, uppercase and lowercase letters, digits, periods (.), underscores (_), and hyphens (-). Spaces are not supported.
	//
	// example:
	//
	// Thunder
	SoftwareName *string `json:"SoftwareName,omitempty" xml:"SoftwareName,omitempty"`
	// The prohibited software tag ID, used to filter policies that reference this tag. You can obtain the value from the following operations:
	//
	// - [ListProhibitedTags](~~ListProhibitedTags~~): Lists prohibited software tags.
	//
	// - [CreateProhibitedTag](~~CreateProhibitedTag~~): Creates a custom prohibited software tag.
	//
	// example:
	//
	// tag-d3f64e8bdd4a****
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// The name of the prohibited software tag. Fuzzy match is supported. The name can be up to 128 characters in length and can contain Chinese characters, uppercase and lowercase letters, digits, periods (.), underscores (_), and hyphens (-). Spaces are not supported.
	//
	// example:
	//
	// CloudProduct
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// The user group ID, used to filter policies whose effective scope includes this user group. You can obtain the value from the following operations:
	//
	// - [ListUserGroups](~~ListUserGroups~~): Lists user groups.
	//
	// - [CreateUserGroup](~~CreateUserGroup~~): Creates a user group.
	//
	// example:
	//
	// usergroup-9d4f2a7b3c1e****
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s ListProhibitedPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProhibitedPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListProhibitedPoliciesRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListProhibitedPoliciesRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListProhibitedPoliciesRequest) GetMatchMode() *string {
	return s.MatchMode
}

func (s *ListProhibitedPoliciesRequest) GetName() *string {
	return s.Name
}

func (s *ListProhibitedPoliciesRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *ListProhibitedPoliciesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListProhibitedPoliciesRequest) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *ListProhibitedPoliciesRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListProhibitedPoliciesRequest) GetSoftwareId() *ListProhibitedPoliciesRequestSoftwareId {
	return s.SoftwareId
}

func (s *ListProhibitedPoliciesRequest) GetSoftwareName() *string {
	return s.SoftwareName
}

func (s *ListProhibitedPoliciesRequest) GetTagId() *string {
	return s.TagId
}

func (s *ListProhibitedPoliciesRequest) GetTagName() *string {
	return s.TagName
}

func (s *ListProhibitedPoliciesRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *ListProhibitedPoliciesRequest) SetCurrentPage(v int64) *ListProhibitedPoliciesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListProhibitedPoliciesRequest) SetEnabled(v bool) *ListProhibitedPoliciesRequest {
	s.Enabled = &v
	return s
}

func (s *ListProhibitedPoliciesRequest) SetMatchMode(v string) *ListProhibitedPoliciesRequest {
	s.MatchMode = &v
	return s
}

func (s *ListProhibitedPoliciesRequest) SetName(v string) *ListProhibitedPoliciesRequest {
	s.Name = &v
	return s
}

func (s *ListProhibitedPoliciesRequest) SetObjectType(v string) *ListProhibitedPoliciesRequest {
	s.ObjectType = &v
	return s
}

func (s *ListProhibitedPoliciesRequest) SetPageSize(v int64) *ListProhibitedPoliciesRequest {
	s.PageSize = &v
	return s
}

func (s *ListProhibitedPoliciesRequest) SetPolicyIds(v []*string) *ListProhibitedPoliciesRequest {
	s.PolicyIds = v
	return s
}

func (s *ListProhibitedPoliciesRequest) SetPolicyType(v string) *ListProhibitedPoliciesRequest {
	s.PolicyType = &v
	return s
}

func (s *ListProhibitedPoliciesRequest) SetSoftwareId(v *ListProhibitedPoliciesRequestSoftwareId) *ListProhibitedPoliciesRequest {
	s.SoftwareId = v
	return s
}

func (s *ListProhibitedPoliciesRequest) SetSoftwareName(v string) *ListProhibitedPoliciesRequest {
	s.SoftwareName = &v
	return s
}

func (s *ListProhibitedPoliciesRequest) SetTagId(v string) *ListProhibitedPoliciesRequest {
	s.TagId = &v
	return s
}

func (s *ListProhibitedPoliciesRequest) SetTagName(v string) *ListProhibitedPoliciesRequest {
	s.TagName = &v
	return s
}

func (s *ListProhibitedPoliciesRequest) SetUserGroupId(v string) *ListProhibitedPoliciesRequest {
	s.UserGroupId = &v
	return s
}

func (s *ListProhibitedPoliciesRequest) Validate() error {
	if s.SoftwareId != nil {
		if err := s.SoftwareId.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListProhibitedPoliciesRequestSoftwareId struct {
	// Specifies whether the prohibited software is a system built-in entry. Valid values:
	//
	// - **true**: A system built-in prohibited software entry shared by all Alibaba Cloud accounts. Modification and deletion are not supported.
	//
	// - **false**: Custom prohibited software under the current Alibaba Cloud account.
	//
	// example:
	//
	// false
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The prohibited software ID. You can obtain the value from the following operations:
	//
	// - [ListProhibitedSoftware](~~ListProhibitedSoftware~~): Lists prohibited software.
	//
	// - [CreateProhibitedSoftware](~~CreateProhibitedSoftware~~): Creates custom prohibited software.
	//
	// example:
	//
	// swb-83995ff2ae38****
	SoftwareId *string `json:"SoftwareId,omitempty" xml:"SoftwareId,omitempty"`
}

func (s ListProhibitedPoliciesRequestSoftwareId) String() string {
	return dara.Prettify(s)
}

func (s ListProhibitedPoliciesRequestSoftwareId) GoString() string {
	return s.String()
}

func (s *ListProhibitedPoliciesRequestSoftwareId) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListProhibitedPoliciesRequestSoftwareId) GetSoftwareId() *string {
	return s.SoftwareId
}

func (s *ListProhibitedPoliciesRequestSoftwareId) SetIsDefault(v bool) *ListProhibitedPoliciesRequestSoftwareId {
	s.IsDefault = &v
	return s
}

func (s *ListProhibitedPoliciesRequestSoftwareId) SetSoftwareId(v string) *ListProhibitedPoliciesRequestSoftwareId {
	s.SoftwareId = &v
	return s
}

func (s *ListProhibitedPoliciesRequestSoftwareId) Validate() error {
	return dara.Validate(s)
}
