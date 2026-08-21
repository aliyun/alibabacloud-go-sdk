// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProhibitedPoliciesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *ListProhibitedPoliciesShrinkRequest
	GetCurrentPage() *int64
	SetEnabled(v bool) *ListProhibitedPoliciesShrinkRequest
	GetEnabled() *bool
	SetMatchMode(v string) *ListProhibitedPoliciesShrinkRequest
	GetMatchMode() *string
	SetName(v string) *ListProhibitedPoliciesShrinkRequest
	GetName() *string
	SetObjectType(v string) *ListProhibitedPoliciesShrinkRequest
	GetObjectType() *string
	SetPageSize(v int64) *ListProhibitedPoliciesShrinkRequest
	GetPageSize() *int64
	SetPolicyIds(v []*string) *ListProhibitedPoliciesShrinkRequest
	GetPolicyIds() []*string
	SetPolicyType(v string) *ListProhibitedPoliciesShrinkRequest
	GetPolicyType() *string
	SetSoftwareIdShrink(v string) *ListProhibitedPoliciesShrinkRequest
	GetSoftwareIdShrink() *string
	SetSoftwareName(v string) *ListProhibitedPoliciesShrinkRequest
	GetSoftwareName() *string
	SetTagId(v string) *ListProhibitedPoliciesShrinkRequest
	GetTagId() *string
	SetTagName(v string) *ListProhibitedPoliciesShrinkRequest
	GetTagName() *string
	SetUserGroupId(v string) *ListProhibitedPoliciesShrinkRequest
	GetUserGroupId() *string
}

type ListProhibitedPoliciesShrinkRequest struct {
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
	SoftwareIdShrink *string `json:"SoftwareId,omitempty" xml:"SoftwareId,omitempty"`
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

func (s ListProhibitedPoliciesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProhibitedPoliciesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListProhibitedPoliciesShrinkRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListProhibitedPoliciesShrinkRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListProhibitedPoliciesShrinkRequest) GetMatchMode() *string {
	return s.MatchMode
}

func (s *ListProhibitedPoliciesShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListProhibitedPoliciesShrinkRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *ListProhibitedPoliciesShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListProhibitedPoliciesShrinkRequest) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *ListProhibitedPoliciesShrinkRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListProhibitedPoliciesShrinkRequest) GetSoftwareIdShrink() *string {
	return s.SoftwareIdShrink
}

func (s *ListProhibitedPoliciesShrinkRequest) GetSoftwareName() *string {
	return s.SoftwareName
}

func (s *ListProhibitedPoliciesShrinkRequest) GetTagId() *string {
	return s.TagId
}

func (s *ListProhibitedPoliciesShrinkRequest) GetTagName() *string {
	return s.TagName
}

func (s *ListProhibitedPoliciesShrinkRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *ListProhibitedPoliciesShrinkRequest) SetCurrentPage(v int64) *ListProhibitedPoliciesShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListProhibitedPoliciesShrinkRequest) SetEnabled(v bool) *ListProhibitedPoliciesShrinkRequest {
	s.Enabled = &v
	return s
}

func (s *ListProhibitedPoliciesShrinkRequest) SetMatchMode(v string) *ListProhibitedPoliciesShrinkRequest {
	s.MatchMode = &v
	return s
}

func (s *ListProhibitedPoliciesShrinkRequest) SetName(v string) *ListProhibitedPoliciesShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListProhibitedPoliciesShrinkRequest) SetObjectType(v string) *ListProhibitedPoliciesShrinkRequest {
	s.ObjectType = &v
	return s
}

func (s *ListProhibitedPoliciesShrinkRequest) SetPageSize(v int64) *ListProhibitedPoliciesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListProhibitedPoliciesShrinkRequest) SetPolicyIds(v []*string) *ListProhibitedPoliciesShrinkRequest {
	s.PolicyIds = v
	return s
}

func (s *ListProhibitedPoliciesShrinkRequest) SetPolicyType(v string) *ListProhibitedPoliciesShrinkRequest {
	s.PolicyType = &v
	return s
}

func (s *ListProhibitedPoliciesShrinkRequest) SetSoftwareIdShrink(v string) *ListProhibitedPoliciesShrinkRequest {
	s.SoftwareIdShrink = &v
	return s
}

func (s *ListProhibitedPoliciesShrinkRequest) SetSoftwareName(v string) *ListProhibitedPoliciesShrinkRequest {
	s.SoftwareName = &v
	return s
}

func (s *ListProhibitedPoliciesShrinkRequest) SetTagId(v string) *ListProhibitedPoliciesShrinkRequest {
	s.TagId = &v
	return s
}

func (s *ListProhibitedPoliciesShrinkRequest) SetTagName(v string) *ListProhibitedPoliciesShrinkRequest {
	s.TagName = &v
	return s
}

func (s *ListProhibitedPoliciesShrinkRequest) SetUserGroupId(v string) *ListProhibitedPoliciesShrinkRequest {
	s.UserGroupId = &v
	return s
}

func (s *ListProhibitedPoliciesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
