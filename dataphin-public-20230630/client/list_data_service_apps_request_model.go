// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceAppsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListDataServiceAppsRequestListQuery) *ListDataServiceAppsRequest
	GetListQuery() *ListDataServiceAppsRequestListQuery
	SetOpTenantId(v int64) *ListDataServiceAppsRequest
	GetOpTenantId() *int64
}

type ListDataServiceAppsRequest struct {
	// This parameter is required.
	ListQuery *ListDataServiceAppsRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListDataServiceAppsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceAppsRequest) GoString() string {
	return s.String()
}

func (s *ListDataServiceAppsRequest) GetListQuery() *ListDataServiceAppsRequestListQuery {
	return s.ListQuery
}

func (s *ListDataServiceAppsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListDataServiceAppsRequest) SetListQuery(v *ListDataServiceAppsRequestListQuery) *ListDataServiceAppsRequest {
	s.ListQuery = v
	return s
}

func (s *ListDataServiceAppsRequest) SetOpTenantId(v int64) *ListDataServiceAppsRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListDataServiceAppsRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataServiceAppsRequestListQuery struct {
	// example:
	//
	// 12345
	AppGroupId *int32 `json:"AppGroupId,omitempty" xml:"AppGroupId,omitempty"`
	// example:
	//
	// 营销看板
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

func (s ListDataServiceAppsRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceAppsRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListDataServiceAppsRequestListQuery) GetAppGroupId() *int32 {
	return s.AppGroupId
}

func (s *ListDataServiceAppsRequestListQuery) GetKeyword() *string {
	return s.Keyword
}

func (s *ListDataServiceAppsRequestListQuery) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListDataServiceAppsRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataServiceAppsRequestListQuery) SetAppGroupId(v int32) *ListDataServiceAppsRequestListQuery {
	s.AppGroupId = &v
	return s
}

func (s *ListDataServiceAppsRequestListQuery) SetKeyword(v string) *ListDataServiceAppsRequestListQuery {
	s.Keyword = &v
	return s
}

func (s *ListDataServiceAppsRequestListQuery) SetPageNo(v int32) *ListDataServiceAppsRequestListQuery {
	s.PageNo = &v
	return s
}

func (s *ListDataServiceAppsRequestListQuery) SetPageSize(v int32) *ListDataServiceAppsRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListDataServiceAppsRequestListQuery) Validate() error {
	return dara.Validate(s)
}
