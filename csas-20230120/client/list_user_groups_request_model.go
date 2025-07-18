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
	// example:
	//
	// username
	AttributeValue *string `json:"AttributeValue,omitempty" xml:"AttributeValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 用户组名称。长度为1~128个字符，支持中文和大小写英文字母，可包含数字、半角句号（.）、下划线（_）和短划线（-）。
	//
	// example:
	//
	// user_group_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// pa-policy-54a7838a48bf****
	PAPolicyId *string `json:"PAPolicyId,omitempty" xml:"PAPolicyId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize     *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
