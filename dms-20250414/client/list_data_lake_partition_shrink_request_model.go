// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakePartitionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *ListDataLakePartitionShrinkRequest
	GetCatalogName() *string
	SetDbName(v string) *ListDataLakePartitionShrinkRequest
	GetDbName() *string
	SetMaxResults(v int32) *ListDataLakePartitionShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataLakePartitionShrinkRequest
	GetNextToken() *string
	SetPartNamesShrink(v string) *ListDataLakePartitionShrinkRequest
	GetPartNamesShrink() *string
	SetTableName(v string) *ListDataLakePartitionShrinkRequest
	GetTableName() *string
	SetTid(v int64) *ListDataLakePartitionShrinkRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *ListDataLakePartitionShrinkRequest
	GetWorkspaceId() *int64
}

type ListDataLakePartitionShrinkRequest struct {
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
	// The number of entries per page. Valid values: 0 to 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results. Set this parameter to the value of NextToken that is returned in the last response.
	//
	// - If **NextToken*	- is empty, no subsequent query is needed.
	//
	// - If **NextToken*	- has a value, that value is the token to start the next query.
	//
	// example:
	//
	// f056501ada12****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of partition names.
	PartNamesShrink *string `json:"PartNames,omitempty" xml:"PartNames,omitempty"`
	// The name of the table.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The ID of the tenant.
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

func (s ListDataLakePartitionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakePartitionShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDataLakePartitionShrinkRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *ListDataLakePartitionShrinkRequest) GetDbName() *string {
	return s.DbName
}

func (s *ListDataLakePartitionShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataLakePartitionShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataLakePartitionShrinkRequest) GetPartNamesShrink() *string {
	return s.PartNamesShrink
}

func (s *ListDataLakePartitionShrinkRequest) GetTableName() *string {
	return s.TableName
}

func (s *ListDataLakePartitionShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListDataLakePartitionShrinkRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *ListDataLakePartitionShrinkRequest) SetCatalogName(v string) *ListDataLakePartitionShrinkRequest {
	s.CatalogName = &v
	return s
}

func (s *ListDataLakePartitionShrinkRequest) SetDbName(v string) *ListDataLakePartitionShrinkRequest {
	s.DbName = &v
	return s
}

func (s *ListDataLakePartitionShrinkRequest) SetMaxResults(v int32) *ListDataLakePartitionShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDataLakePartitionShrinkRequest) SetNextToken(v string) *ListDataLakePartitionShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListDataLakePartitionShrinkRequest) SetPartNamesShrink(v string) *ListDataLakePartitionShrinkRequest {
	s.PartNamesShrink = &v
	return s
}

func (s *ListDataLakePartitionShrinkRequest) SetTableName(v string) *ListDataLakePartitionShrinkRequest {
	s.TableName = &v
	return s
}

func (s *ListDataLakePartitionShrinkRequest) SetTid(v int64) *ListDataLakePartitionShrinkRequest {
	s.Tid = &v
	return s
}

func (s *ListDataLakePartitionShrinkRequest) SetWorkspaceId(v int64) *ListDataLakePartitionShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListDataLakePartitionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
