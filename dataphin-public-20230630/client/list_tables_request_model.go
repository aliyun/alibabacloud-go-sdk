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
	SetOpUserId(v string) *ListTablesRequest
	GetOpUserId() *string
}

type ListTablesRequest struct {
	// The paged query conditions.
	ListQuery *ListTablesRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
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

func (s *ListTablesRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListTablesRequest) SetListQuery(v *ListTablesRequestListQuery) *ListTablesRequest {
	s.ListQuery = v
	return s
}

func (s *ListTablesRequest) SetOpTenantId(v int64) *ListTablesRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListTablesRequest) SetOpUserId(v string) *ListTablesRequest {
	s.OpUserId = &v
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
	// The asset catalog, such as the project name or business unit name.
	//
	// example:
	//
	// LD_test01_dev
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// The keyword for searching. Table names are supported.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// 30012011
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of records per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SubTypes []*string `json:"SubTypes,omitempty" xml:"SubTypes,omitempty" type:"Repeated"`
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

func (s *ListTablesRequestListQuery) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListTablesRequestListQuery) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListTablesRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTablesRequestListQuery) GetSubTypes() []*string {
	return s.SubTypes
}

func (s *ListTablesRequestListQuery) SetCatalog(v string) *ListTablesRequestListQuery {
	s.Catalog = &v
	return s
}

func (s *ListTablesRequestListQuery) SetKeyword(v string) *ListTablesRequestListQuery {
	s.Keyword = &v
	return s
}

func (s *ListTablesRequestListQuery) SetOwnerId(v string) *ListTablesRequestListQuery {
	s.OwnerId = &v
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

func (s *ListTablesRequestListQuery) SetSubTypes(v []*string) *ListTablesRequestListQuery {
	s.SubTypes = v
	return s
}

func (s *ListTablesRequestListQuery) Validate() error {
	return dara.Validate(s)
}
