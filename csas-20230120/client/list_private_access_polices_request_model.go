// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivateAccessPolicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ListPrivateAccessPolicesRequest
	GetApplicationId() *string
	SetApplicationName(v string) *ListPrivateAccessPolicesRequest
	GetApplicationName() *string
	SetCurrentPage(v int32) *ListPrivateAccessPolicesRequest
	GetCurrentPage() *int32
	SetName(v string) *ListPrivateAccessPolicesRequest
	GetName() *string
	SetPageSize(v int32) *ListPrivateAccessPolicesRequest
	GetPageSize() *int32
	SetPolicyAction(v string) *ListPrivateAccessPolicesRequest
	GetPolicyAction() *string
	SetPolicyIds(v []*string) *ListPrivateAccessPolicesRequest
	GetPolicyIds() []*string
	SetStatus(v string) *ListPrivateAccessPolicesRequest
	GetStatus() *string
	SetTagId(v string) *ListPrivateAccessPolicesRequest
	GetTagId() *string
	SetTagName(v string) *ListPrivateAccessPolicesRequest
	GetTagName() *string
	SetUserGroupId(v string) *ListPrivateAccessPolicesRequest
	GetUserGroupId() *string
}

type ListPrivateAccessPolicesRequest struct {
	// The ID of the office application. Either the ID or tag of the office application is used for queries. You can obtain the value by calling the following operations:
	//
	// 	- [ListPrivateAccessApplications](~~ListPrivateAccessApplications~~): queries office applications.
	//
	// 	- [CreatePrivateAccessApplication](~~CreatePrivateAccessApplication~~): creates an office application.
	//
	// example:
	//
	// pa-application-e12860ef6c48****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The name of the office application.
	//
	// example:
	//
	// Office
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The page number. Valid values: 1 to 10000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The name of the private access policy. The value must be 1 to 128 characters in length and can contain letters, digits, hyphens (-), underscores (_), and periods (.).
	//
	// example:
	//
	// private_access_policy_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries per page. Valid values: 1 to 1000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The action in the private access policy. Valid values:
	//
	// 	- **Block**
	//
	// 	- **Allow**
	//
	// example:
	//
	// Allow
	PolicyAction *string `json:"PolicyAction,omitempty" xml:"PolicyAction,omitempty"`
	// The IDs of the private access policies. You can enter up to 100 IDs.
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// The status of the private access policy. Valid values:
	//
	// 	- **Enabled**
	//
	// 	- **Disabled**
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the tag for the office application. Either the ID or tag of the office application is used for queries. You can obtain the value by calling the following operations:
	//
	// 	- [ListPrivateAccessTags](~~ListPrivateAccessTags~~): queries tags for office applications.
	//
	// 	- [CreatePrivateAccessTag](~~CreatePrivateAccessTag~~): creates a tag for office applications.
	//
	// example:
	//
	// tag-c0cb77857a99****
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// The name of the tag.
	//
	// example:
	//
	// Cloud service
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// The ID of the user group. You can obtain the value by calling the following operations:
	//
	// 	- [ListUserGroups](~~ListUserGroups~~): queries user groups.
	//
	// 	- [CreateUserGroup](~~CreateUserGroup~~): creates a user group.
	//
	// example:
	//
	// usergroup-6f1ef2fc56b6****
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s ListPrivateAccessPolicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateAccessPolicesRequest) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessPolicesRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListPrivateAccessPolicesRequest) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *ListPrivateAccessPolicesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListPrivateAccessPolicesRequest) GetName() *string {
	return s.Name
}

func (s *ListPrivateAccessPolicesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPrivateAccessPolicesRequest) GetPolicyAction() *string {
	return s.PolicyAction
}

func (s *ListPrivateAccessPolicesRequest) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *ListPrivateAccessPolicesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListPrivateAccessPolicesRequest) GetTagId() *string {
	return s.TagId
}

func (s *ListPrivateAccessPolicesRequest) GetTagName() *string {
	return s.TagName
}

func (s *ListPrivateAccessPolicesRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *ListPrivateAccessPolicesRequest) SetApplicationId(v string) *ListPrivateAccessPolicesRequest {
	s.ApplicationId = &v
	return s
}

func (s *ListPrivateAccessPolicesRequest) SetApplicationName(v string) *ListPrivateAccessPolicesRequest {
	s.ApplicationName = &v
	return s
}

func (s *ListPrivateAccessPolicesRequest) SetCurrentPage(v int32) *ListPrivateAccessPolicesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListPrivateAccessPolicesRequest) SetName(v string) *ListPrivateAccessPolicesRequest {
	s.Name = &v
	return s
}

func (s *ListPrivateAccessPolicesRequest) SetPageSize(v int32) *ListPrivateAccessPolicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListPrivateAccessPolicesRequest) SetPolicyAction(v string) *ListPrivateAccessPolicesRequest {
	s.PolicyAction = &v
	return s
}

func (s *ListPrivateAccessPolicesRequest) SetPolicyIds(v []*string) *ListPrivateAccessPolicesRequest {
	s.PolicyIds = v
	return s
}

func (s *ListPrivateAccessPolicesRequest) SetStatus(v string) *ListPrivateAccessPolicesRequest {
	s.Status = &v
	return s
}

func (s *ListPrivateAccessPolicesRequest) SetTagId(v string) *ListPrivateAccessPolicesRequest {
	s.TagId = &v
	return s
}

func (s *ListPrivateAccessPolicesRequest) SetTagName(v string) *ListPrivateAccessPolicesRequest {
	s.TagName = &v
	return s
}

func (s *ListPrivateAccessPolicesRequest) SetUserGroupId(v string) *ListPrivateAccessPolicesRequest {
	s.UserGroupId = &v
	return s
}

func (s *ListPrivateAccessPolicesRequest) Validate() error {
	return dara.Validate(s)
}
