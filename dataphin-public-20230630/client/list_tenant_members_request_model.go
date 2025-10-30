// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTenantMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListTenantMembersRequestListQuery) *ListTenantMembersRequest
	GetListQuery() *ListTenantMembersRequestListQuery
	SetOpTenantId(v int64) *ListTenantMembersRequest
	GetOpTenantId() *int64
}

type ListTenantMembersRequest struct {
	// This parameter is required.
	ListQuery *ListTenantMembersRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListTenantMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTenantMembersRequest) GoString() string {
	return s.String()
}

func (s *ListTenantMembersRequest) GetListQuery() *ListTenantMembersRequestListQuery {
	return s.ListQuery
}

func (s *ListTenantMembersRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListTenantMembersRequest) SetListQuery(v *ListTenantMembersRequestListQuery) *ListTenantMembersRequest {
	s.ListQuery = v
	return s
}

func (s *ListTenantMembersRequest) SetOpTenantId(v int64) *ListTenantMembersRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListTenantMembersRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTenantMembersRequestListQuery struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize        *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RoleList        []*string `json:"RoleList,omitempty" xml:"RoleList,omitempty" type:"Repeated"`
	SearchText      *string   `json:"SearchText,omitempty" xml:"SearchText,omitempty"`
	UserGroupIdList []*string `json:"UserGroupIdList,omitempty" xml:"UserGroupIdList,omitempty" type:"Repeated"`
}

func (s ListTenantMembersRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListTenantMembersRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListTenantMembersRequestListQuery) GetPage() *int32 {
	return s.Page
}

func (s *ListTenantMembersRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTenantMembersRequestListQuery) GetRoleList() []*string {
	return s.RoleList
}

func (s *ListTenantMembersRequestListQuery) GetSearchText() *string {
	return s.SearchText
}

func (s *ListTenantMembersRequestListQuery) GetUserGroupIdList() []*string {
	return s.UserGroupIdList
}

func (s *ListTenantMembersRequestListQuery) SetPage(v int32) *ListTenantMembersRequestListQuery {
	s.Page = &v
	return s
}

func (s *ListTenantMembersRequestListQuery) SetPageSize(v int32) *ListTenantMembersRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListTenantMembersRequestListQuery) SetRoleList(v []*string) *ListTenantMembersRequestListQuery {
	s.RoleList = v
	return s
}

func (s *ListTenantMembersRequestListQuery) SetSearchText(v string) *ListTenantMembersRequestListQuery {
	s.SearchText = &v
	return s
}

func (s *ListTenantMembersRequestListQuery) SetUserGroupIdList(v []*string) *ListTenantMembersRequestListQuery {
	s.UserGroupIdList = v
	return s
}

func (s *ListTenantMembersRequestListQuery) Validate() error {
	return dara.Validate(s)
}
