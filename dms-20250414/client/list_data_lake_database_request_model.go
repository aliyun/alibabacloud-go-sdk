// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakeDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *ListDataLakeDatabaseRequest
	GetCatalogName() *string
	SetMaxResults(v int32) *ListDataLakeDatabaseRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataLakeDatabaseRequest
	GetNextToken() *string
	SetSearchKey(v string) *ListDataLakeDatabaseRequest
	GetSearchKey() *string
	SetTid(v int64) *ListDataLakeDatabaseRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *ListDataLakeDatabaseRequest
	GetWorkspaceId() *int64
}

type ListDataLakeDatabaseRequest struct {
	// The name of the catalog. To find the catelog name, go to the [Data Lake console](https://dlf.console.aliyun.com/cn-hangzhou/metadata/catalog?spm=a2c4g.11186623.0.0.5a225658pT4Dkr).
	//
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// hive
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// The number of entries to return on each page. The maximum value is 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results. Valid values:
	//
	// - If **NextToken*	- is empty, no more results are available.
	//
	// - If **NextToken*	- has a value, it is the token for the next query.
	//
	// example:
	//
	// f056501ada12****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// A keyword to search for database names.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// default
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// The tenant ID.
	//
	// > Hover over your profile picture in the upper-right corner of the DMS console to obtain the tenant ID. For details, see [View tenant information](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3****
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 12****
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListDataLakeDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakeDatabaseRequest) GoString() string {
	return s.String()
}

func (s *ListDataLakeDatabaseRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *ListDataLakeDatabaseRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataLakeDatabaseRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataLakeDatabaseRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *ListDataLakeDatabaseRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListDataLakeDatabaseRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *ListDataLakeDatabaseRequest) SetCatalogName(v string) *ListDataLakeDatabaseRequest {
	s.CatalogName = &v
	return s
}

func (s *ListDataLakeDatabaseRequest) SetMaxResults(v int32) *ListDataLakeDatabaseRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDataLakeDatabaseRequest) SetNextToken(v string) *ListDataLakeDatabaseRequest {
	s.NextToken = &v
	return s
}

func (s *ListDataLakeDatabaseRequest) SetSearchKey(v string) *ListDataLakeDatabaseRequest {
	s.SearchKey = &v
	return s
}

func (s *ListDataLakeDatabaseRequest) SetTid(v int64) *ListDataLakeDatabaseRequest {
	s.Tid = &v
	return s
}

func (s *ListDataLakeDatabaseRequest) SetWorkspaceId(v int64) *ListDataLakeDatabaseRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListDataLakeDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
