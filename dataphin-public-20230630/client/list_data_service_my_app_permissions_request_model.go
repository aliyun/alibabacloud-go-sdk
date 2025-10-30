// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceMyAppPermissionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListDataServiceMyAppPermissionsRequestListQuery) *ListDataServiceMyAppPermissionsRequest
	GetListQuery() *ListDataServiceMyAppPermissionsRequestListQuery
	SetOpTenantId(v int64) *ListDataServiceMyAppPermissionsRequest
	GetOpTenantId() *int64
	SetProjectId(v int32) *ListDataServiceMyAppPermissionsRequest
	GetProjectId() *int32
}

type ListDataServiceMyAppPermissionsRequest struct {
	// This parameter is required.
	ListQuery *ListDataServiceMyAppPermissionsRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
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

func (s ListDataServiceMyAppPermissionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceMyAppPermissionsRequest) GoString() string {
	return s.String()
}

func (s *ListDataServiceMyAppPermissionsRequest) GetListQuery() *ListDataServiceMyAppPermissionsRequestListQuery {
	return s.ListQuery
}

func (s *ListDataServiceMyAppPermissionsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListDataServiceMyAppPermissionsRequest) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *ListDataServiceMyAppPermissionsRequest) SetListQuery(v *ListDataServiceMyAppPermissionsRequestListQuery) *ListDataServiceMyAppPermissionsRequest {
	s.ListQuery = v
	return s
}

func (s *ListDataServiceMyAppPermissionsRequest) SetOpTenantId(v int64) *ListDataServiceMyAppPermissionsRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListDataServiceMyAppPermissionsRequest) SetProjectId(v int32) *ListDataServiceMyAppPermissionsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDataServiceMyAppPermissionsRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataServiceMyAppPermissionsRequestListQuery struct {
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

func (s ListDataServiceMyAppPermissionsRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceMyAppPermissionsRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListDataServiceMyAppPermissionsRequestListQuery) GetKeyword() *string {
	return s.Keyword
}

func (s *ListDataServiceMyAppPermissionsRequestListQuery) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListDataServiceMyAppPermissionsRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataServiceMyAppPermissionsRequestListQuery) SetKeyword(v string) *ListDataServiceMyAppPermissionsRequestListQuery {
	s.Keyword = &v
	return s
}

func (s *ListDataServiceMyAppPermissionsRequestListQuery) SetPageNo(v int32) *ListDataServiceMyAppPermissionsRequestListQuery {
	s.PageNo = &v
	return s
}

func (s *ListDataServiceMyAppPermissionsRequestListQuery) SetPageSize(v int32) *ListDataServiceMyAppPermissionsRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListDataServiceMyAppPermissionsRequestListQuery) Validate() error {
	return dara.Validate(s)
}
