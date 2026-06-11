// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakeTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *ListDataLakeTableRequest
	GetCatalogName() *string
	SetDbName(v string) *ListDataLakeTableRequest
	GetDbName() *string
	SetMaxResults(v int32) *ListDataLakeTableRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataLakeTableRequest
	GetNextToken() *string
	SetTableNamePattern(v string) *ListDataLakeTableRequest
	GetTableNamePattern() *string
	SetTableType(v string) *ListDataLakeTableRequest
	GetTableType() *string
	SetTid(v int64) *ListDataLakeTableRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *ListDataLakeTableRequest
	GetWorkspaceId() *int64
}

type ListDataLakeTableRequest struct {
	// The name of the data catalog.
	//
	// This parameter is required.
	//
	// example:
	//
	// hive
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The number of entries to return on each page. The maximum value is 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to retrieve the next page of results. To retrieve the next page, set this parameter to the \\`NextToken\\` value from the previous response. If you do not specify this parameter, the first page is returned.
	//
	// - If **NextToken*	- is empty, there are no more queries.
	//
	// - If **NextToken*	- has a value, the value is the token to start the next query.
	//
	// example:
	//
	// f056501ada12****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The name of the metadata table. This parameter supports regular expressions.
	//
	// example:
	//
	// .*
	TableNamePattern *string `json:"TableNamePattern,omitempty" xml:"TableNamePattern,omitempty"`
	// The type of the table. Valid values:
	//
	// - MANAGED_TABLE: Internal table.
	//
	// - EXTERNAL_TABLE: External table.
	//
	// - VIRTUAL_VIEW: Virtual view.
	//
	// - INDEX_TABLE: Index table.
	//
	// - MATERIALIZED_VIEW: Materialized view.
	//
	// example:
	//
	// MANAGED_TABLE
	TableType *string `json:"TableType,omitempty" xml:"TableType,omitempty"`
	// The tenant ID.
	//
	// > Hover over your profile picture in the upper-right corner of the DMS console to obtain the tenant ID. For details, see [View tenant information](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3****
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The ID of the workspace.
	//
	// example:
	//
	// 12****
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListDataLakeTableRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakeTableRequest) GoString() string {
	return s.String()
}

func (s *ListDataLakeTableRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *ListDataLakeTableRequest) GetDbName() *string {
	return s.DbName
}

func (s *ListDataLakeTableRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataLakeTableRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataLakeTableRequest) GetTableNamePattern() *string {
	return s.TableNamePattern
}

func (s *ListDataLakeTableRequest) GetTableType() *string {
	return s.TableType
}

func (s *ListDataLakeTableRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListDataLakeTableRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *ListDataLakeTableRequest) SetCatalogName(v string) *ListDataLakeTableRequest {
	s.CatalogName = &v
	return s
}

func (s *ListDataLakeTableRequest) SetDbName(v string) *ListDataLakeTableRequest {
	s.DbName = &v
	return s
}

func (s *ListDataLakeTableRequest) SetMaxResults(v int32) *ListDataLakeTableRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDataLakeTableRequest) SetNextToken(v string) *ListDataLakeTableRequest {
	s.NextToken = &v
	return s
}

func (s *ListDataLakeTableRequest) SetTableNamePattern(v string) *ListDataLakeTableRequest {
	s.TableNamePattern = &v
	return s
}

func (s *ListDataLakeTableRequest) SetTableType(v string) *ListDataLakeTableRequest {
	s.TableType = &v
	return s
}

func (s *ListDataLakeTableRequest) SetTid(v int64) *ListDataLakeTableRequest {
	s.Tid = &v
	return s
}

func (s *ListDataLakeTableRequest) SetWorkspaceId(v int64) *ListDataLakeTableRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListDataLakeTableRequest) Validate() error {
	return dara.Validate(s)
}
