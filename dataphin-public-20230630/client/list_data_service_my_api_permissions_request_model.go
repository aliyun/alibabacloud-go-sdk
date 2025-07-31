// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceMyApiPermissionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListDataServiceMyApiPermissionsRequestListQuery) *ListDataServiceMyApiPermissionsRequest
	GetListQuery() *ListDataServiceMyApiPermissionsRequestListQuery
	SetOpTenantId(v int64) *ListDataServiceMyApiPermissionsRequest
	GetOpTenantId() *int64
	SetProjectId(v int32) *ListDataServiceMyApiPermissionsRequest
	GetProjectId() *int32
}

type ListDataServiceMyApiPermissionsRequest struct {
	ListQuery *ListDataServiceMyApiPermissionsRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 102102
	ProjectId *int32 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListDataServiceMyApiPermissionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceMyApiPermissionsRequest) GoString() string {
	return s.String()
}

func (s *ListDataServiceMyApiPermissionsRequest) GetListQuery() *ListDataServiceMyApiPermissionsRequestListQuery {
	return s.ListQuery
}

func (s *ListDataServiceMyApiPermissionsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListDataServiceMyApiPermissionsRequest) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *ListDataServiceMyApiPermissionsRequest) SetListQuery(v *ListDataServiceMyApiPermissionsRequestListQuery) *ListDataServiceMyApiPermissionsRequest {
	s.ListQuery = v
	return s
}

func (s *ListDataServiceMyApiPermissionsRequest) SetOpTenantId(v int64) *ListDataServiceMyApiPermissionsRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListDataServiceMyApiPermissionsRequest) SetProjectId(v int32) *ListDataServiceMyApiPermissionsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDataServiceMyApiPermissionsRequest) Validate() error {
	return dara.Validate(s)
}

type ListDataServiceMyApiPermissionsRequestListQuery struct {
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDataServiceMyApiPermissionsRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceMyApiPermissionsRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListDataServiceMyApiPermissionsRequestListQuery) GetKeyword() *string {
	return s.Keyword
}

func (s *ListDataServiceMyApiPermissionsRequestListQuery) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListDataServiceMyApiPermissionsRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataServiceMyApiPermissionsRequestListQuery) SetKeyword(v string) *ListDataServiceMyApiPermissionsRequestListQuery {
	s.Keyword = &v
	return s
}

func (s *ListDataServiceMyApiPermissionsRequestListQuery) SetPageNo(v int32) *ListDataServiceMyApiPermissionsRequestListQuery {
	s.PageNo = &v
	return s
}

func (s *ListDataServiceMyApiPermissionsRequestListQuery) SetPageSize(v int32) *ListDataServiceMyApiPermissionsRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListDataServiceMyApiPermissionsRequestListQuery) Validate() error {
	return dara.Validate(s)
}
