// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListUserGroupsRequestListQuery) *ListUserGroupsRequest
	GetListQuery() *ListUserGroupsRequestListQuery
	SetOpTenantId(v int64) *ListUserGroupsRequest
	GetOpTenantId() *int64
}

type ListUserGroupsRequest struct {
	// This parameter is required.
	ListQuery *ListUserGroupsRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListUserGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListUserGroupsRequest) GetListQuery() *ListUserGroupsRequestListQuery {
	return s.ListQuery
}

func (s *ListUserGroupsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListUserGroupsRequest) SetListQuery(v *ListUserGroupsRequestListQuery) *ListUserGroupsRequest {
	s.ListQuery = v
	return s
}

func (s *ListUserGroupsRequest) SetOpTenantId(v int64) *ListUserGroupsRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListUserGroupsRequest) Validate() error {
	return dara.Validate(s)
}

type ListUserGroupsRequestListQuery struct {
	// example:
	//
	// true
	Active      *bool     `json:"Active,omitempty" xml:"Active,omitempty"`
	AdminIdList []*string `json:"AdminIdList,omitempty" xml:"AdminIdList,omitempty" type:"Repeated"`
	// example:
	//
	// false
	FilterMine *bool   `json:"FilterMine,omitempty" xml:"FilterMine,omitempty"`
	Keyword    *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListUserGroupsRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListUserGroupsRequestListQuery) GetActive() *bool {
	return s.Active
}

func (s *ListUserGroupsRequestListQuery) GetAdminIdList() []*string {
	return s.AdminIdList
}

func (s *ListUserGroupsRequestListQuery) GetFilterMine() *bool {
	return s.FilterMine
}

func (s *ListUserGroupsRequestListQuery) GetKeyword() *string {
	return s.Keyword
}

func (s *ListUserGroupsRequestListQuery) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListUserGroupsRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUserGroupsRequestListQuery) SetActive(v bool) *ListUserGroupsRequestListQuery {
	s.Active = &v
	return s
}

func (s *ListUserGroupsRequestListQuery) SetAdminIdList(v []*string) *ListUserGroupsRequestListQuery {
	s.AdminIdList = v
	return s
}

func (s *ListUserGroupsRequestListQuery) SetFilterMine(v bool) *ListUserGroupsRequestListQuery {
	s.FilterMine = &v
	return s
}

func (s *ListUserGroupsRequestListQuery) SetKeyword(v string) *ListUserGroupsRequestListQuery {
	s.Keyword = &v
	return s
}

func (s *ListUserGroupsRequestListQuery) SetPageNo(v int32) *ListUserGroupsRequestListQuery {
	s.PageNo = &v
	return s
}

func (s *ListUserGroupsRequestListQuery) SetPageSize(v int32) *ListUserGroupsRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListUserGroupsRequestListQuery) Validate() error {
	return dara.Validate(s)
}
