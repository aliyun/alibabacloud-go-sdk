// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceAuthorizedAppsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListDataServiceAuthorizedAppsRequestListQuery) *ListDataServiceAuthorizedAppsRequest
	GetListQuery() *ListDataServiceAuthorizedAppsRequestListQuery
	SetOpTenantId(v int64) *ListDataServiceAuthorizedAppsRequest
	GetOpTenantId() *int64
	SetProjectId(v int32) *ListDataServiceAuthorizedAppsRequest
	GetProjectId() *int32
}

type ListDataServiceAuthorizedAppsRequest struct {
	ListQuery *ListDataServiceAuthorizedAppsRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
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

func (s ListDataServiceAuthorizedAppsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceAuthorizedAppsRequest) GoString() string {
	return s.String()
}

func (s *ListDataServiceAuthorizedAppsRequest) GetListQuery() *ListDataServiceAuthorizedAppsRequestListQuery {
	return s.ListQuery
}

func (s *ListDataServiceAuthorizedAppsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListDataServiceAuthorizedAppsRequest) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *ListDataServiceAuthorizedAppsRequest) SetListQuery(v *ListDataServiceAuthorizedAppsRequestListQuery) *ListDataServiceAuthorizedAppsRequest {
	s.ListQuery = v
	return s
}

func (s *ListDataServiceAuthorizedAppsRequest) SetOpTenantId(v int64) *ListDataServiceAuthorizedAppsRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListDataServiceAuthorizedAppsRequest) SetProjectId(v int32) *ListDataServiceAuthorizedAppsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDataServiceAuthorizedAppsRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataServiceAuthorizedAppsRequestListQuery struct {
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

func (s ListDataServiceAuthorizedAppsRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceAuthorizedAppsRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListDataServiceAuthorizedAppsRequestListQuery) GetKeyword() *string {
	return s.Keyword
}

func (s *ListDataServiceAuthorizedAppsRequestListQuery) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListDataServiceAuthorizedAppsRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataServiceAuthorizedAppsRequestListQuery) SetKeyword(v string) *ListDataServiceAuthorizedAppsRequestListQuery {
	s.Keyword = &v
	return s
}

func (s *ListDataServiceAuthorizedAppsRequestListQuery) SetPageNo(v int32) *ListDataServiceAuthorizedAppsRequestListQuery {
	s.PageNo = &v
	return s
}

func (s *ListDataServiceAuthorizedAppsRequestListQuery) SetPageSize(v int32) *ListDataServiceAuthorizedAppsRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListDataServiceAuthorizedAppsRequestListQuery) Validate() error {
	return dara.Validate(s)
}
