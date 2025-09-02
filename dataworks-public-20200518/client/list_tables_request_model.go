// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceType(v string) *ListTablesRequest
	GetDataSourceType() *string
	SetNextToken(v string) *ListTablesRequest
	GetNextToken() *string
	SetPageSize(v int32) *ListTablesRequest
	GetPageSize() *int32
}

type ListTablesRequest struct {
	// The type of the data source. Valid values: ODPS, emr, mysql, holo, analyticdb_for_mysql, oracle, postgresql, sqlserver, clickhouse, and starrocks.
	//
	// This parameter is required.
	//
	// example:
	//
	// odps
	//
	// emr
	//
	// mysql
	//
	// holo
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// Pagination information, which specifies the starting point of this read.
	//
	// example:
	//
	// 12222
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries displayed on each page. The default value is 10 and the maximum value is 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTablesRequest) GoString() string {
	return s.String()
}

func (s *ListTablesRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *ListTablesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTablesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTablesRequest) SetDataSourceType(v string) *ListTablesRequest {
	s.DataSourceType = &v
	return s
}

func (s *ListTablesRequest) SetNextToken(v string) *ListTablesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTablesRequest) SetPageSize(v int32) *ListTablesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTablesRequest) Validate() error {
	return dara.Validate(s)
}
