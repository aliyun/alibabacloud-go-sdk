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
	SetOpUserId(v string) *ListTenantMembersRequest
	GetOpUserId() *string
}

type ListTenantMembersRequest struct {
	// The request object.
	//
	// This parameter is required.
	ListQuery *ListTenantMembersRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
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

func (s *ListTenantMembersRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListTenantMembersRequest) SetListQuery(v *ListTenantMembersRequestListQuery) *ListTenantMembersRequest {
	s.ListQuery = v
	return s
}

func (s *ListTenantMembersRequest) SetOpTenantId(v int64) *ListTenantMembersRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListTenantMembersRequest) SetOpUserId(v string) *ListTenantMembersRequest {
	s.OpUserId = &v
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
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of records per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The member roles. Valid values:
	//
	// - SUPER_ADMIN: Dataphin super administrator.
	//
	// - SYSTEM_ADMIN: System administrator.
	//
	// - COMMON_USER: Dataphin user.
	//
	// - DATA_ADMIN: Dataphin data administrator.
	//
	// - EXPORT_ADMIN: Export administrator.
	//
	// - SECURITY_ADMIN: Security administrator.
	//
	// - DATASOURCE_MANAGER: Data source administrator.
	//
	// - QUALITY_MANAGER: Asset quality owner.
	//
	// - DATA_STANDARD_MANAGER: Data standard administrator.
	//
	// - LABELS_BUSINESS_PLANNER: Label business planner.
	//
	// - BUSINESS_MEMBER: Regular business user.
	//
	// - DATAPRO_OPERATE_SUPER_ADMIN: Operations super administrator.
	//
	// - DATAPRO_OPERATE_ADMIN: Operations administrator.
	//
	// - DATAPRO_OPERATE_MEMBER: Operations staff.
	//
	// - DATAPRO_BUSINESS_ANALYST: Business analyst.
	//
	// - LABELS_BUSINESS_MEMBER: Label business staff.
	//
	// - DATAPRO_BUSINESS_MEMBER: Datapro regular business user.
	RoleList []*string `json:"RoleList,omitempty" xml:"RoleList,omitempty" type:"Repeated"`
	// The search keyword.
	//
	// example:
	//
	// Test
	SearchText *string `json:"SearchText,omitempty" xml:"SearchText,omitempty"`
	// The IDs of the user groups to which the members belong.
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
