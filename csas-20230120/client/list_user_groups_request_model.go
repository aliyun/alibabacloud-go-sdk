// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttributeValue(v string) *ListUserGroupsRequest
	GetAttributeValue() *string
	SetCurrentPage(v int32) *ListUserGroupsRequest
	GetCurrentPage() *int32
	SetName(v string) *ListUserGroupsRequest
	GetName() *string
	SetPAPolicyId(v string) *ListUserGroupsRequest
	GetPAPolicyId() *string
	SetPageSize(v int32) *ListUserGroupsRequest
	GetPageSize() *int32
	SetUserGroupIds(v []*string) *ListUserGroupsRequest
	GetUserGroupIds() []*string
}

type ListUserGroupsRequest struct {
	// The value of a user group property. The value must be 1 to 128 characters in length.
	//
	// example:
	//
	// username
	AttributeValue *string `json:"AttributeValue,omitempty" xml:"AttributeValue,omitempty"`
	// The page number of the current page in a paged query. Valid values: 1 to 10000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The name of the user group. The name must be 1 to 128 characters in length. It can contain letters, digits, periods (.), underscores (_), and hyphens (-). It supports both uppercase and lowercase letters and Chinese characters.
	//
	// example:
	//
	// user_group_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of an internal network access policy. You can get this value from:
	//
	// - [ListPrivateAccessPolices](~~ListPrivateAccessPolices~~): List internal network access policies.
	//
	// - [CreatePrivateAccessPolicy](~~CreatePrivateAccessPolicy~~): Create an internal network access policy.
	//
	// example:
	//
	// pa-policy-54a7838a48bf****
	PAPolicyId *string `json:"PAPolicyId,omitempty" xml:"PAPolicyId,omitempty"`
	// The number of entries per page in a paged query. Valid values: 1 to 1000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// A collection of user group IDs. You can specify up to 100 IDs.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
}

func (s ListUserGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListUserGroupsRequest) GetAttributeValue() *string {
	return s.AttributeValue
}

func (s *ListUserGroupsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListUserGroupsRequest) GetName() *string {
	return s.Name
}

func (s *ListUserGroupsRequest) GetPAPolicyId() *string {
	return s.PAPolicyId
}

func (s *ListUserGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUserGroupsRequest) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *ListUserGroupsRequest) SetAttributeValue(v string) *ListUserGroupsRequest {
	s.AttributeValue = &v
	return s
}

func (s *ListUserGroupsRequest) SetCurrentPage(v int32) *ListUserGroupsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListUserGroupsRequest) SetName(v string) *ListUserGroupsRequest {
	s.Name = &v
	return s
}

func (s *ListUserGroupsRequest) SetPAPolicyId(v string) *ListUserGroupsRequest {
	s.PAPolicyId = &v
	return s
}

func (s *ListUserGroupsRequest) SetPageSize(v int32) *ListUserGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserGroupsRequest) SetUserGroupIds(v []*string) *ListUserGroupsRequest {
	s.UserGroupIds = v
	return s
}

func (s *ListUserGroupsRequest) Validate() error {
	return dara.Validate(s)
}
