// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListTablesRequestListQuery) *ListTablesRequest
	GetListQuery() *ListTablesRequestListQuery
	SetOpTenantId(v int64) *ListTablesRequest
	GetOpTenantId() *int64
}

type ListTablesRequest struct {
	ListQuery *ListTablesRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTablesRequest) GoString() string {
	return s.String()
}

func (s *ListTablesRequest) GetListQuery() *ListTablesRequestListQuery {
	return s.ListQuery
}

func (s *ListTablesRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListTablesRequest) SetListQuery(v *ListTablesRequestListQuery) *ListTablesRequest {
	s.ListQuery = v
	return s
}

func (s *ListTablesRequest) SetOpTenantId(v int64) *ListTablesRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListTablesRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTablesRequestListQuery struct {
	// This parameter is required.
	//
	// example:
	//
	// LD_test01_dev
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
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

func (s ListTablesRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListTablesRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListTablesRequestListQuery) GetCatalog() *string {
	return s.Catalog
}

func (s *ListTablesRequestListQuery) GetKeyword() *string {
	return s.Keyword
}

func (s *ListTablesRequestListQuery) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListTablesRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTablesRequestListQuery) SetCatalog(v string) *ListTablesRequestListQuery {
	s.Catalog = &v
	return s
}

func (s *ListTablesRequestListQuery) SetKeyword(v string) *ListTablesRequestListQuery {
	s.Keyword = &v
	return s
}

func (s *ListTablesRequestListQuery) SetPageNo(v int32) *ListTablesRequestListQuery {
	s.PageNo = &v
	return s
}

func (s *ListTablesRequestListQuery) SetPageSize(v int32) *ListTablesRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListTablesRequestListQuery) Validate() error {
	return dara.Validate(s)
}
