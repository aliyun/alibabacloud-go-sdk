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
	SetDataRegion(v string) *ListDataLakeDatabaseRequest
	GetDataRegion() *string
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
	// The name of the data catalog. You can view the name of the data catalog in the [DLF console](https://dlf.console.aliyun.com/cn-hangzhou/metadata/catalog?spm=a2c4g.11186623.0.0.5a225658pT4Dkr).
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
	// The region where the data lake resides.
	//
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// cn-hangzhou
	DataRegion *string `json:"DataRegion,omitempty" xml:"DataRegion,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// f056501ada12c1cc
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The keyword that is used to search for databases.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// default
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// The ID of the tenant.
	//
	// > To view the tenant ID, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
	//
	// example:
	//
	// 3
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

func (s *ListDataLakeDatabaseRequest) GetDataRegion() *string {
	return s.DataRegion
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

func (s *ListDataLakeDatabaseRequest) SetDataRegion(v string) *ListDataLakeDatabaseRequest {
	s.DataRegion = &v
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
