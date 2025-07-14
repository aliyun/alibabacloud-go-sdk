// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDatasetListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *QueryDatasetListRequest
	GetDirectoryId() *string
	SetKeyword(v string) *QueryDatasetListRequest
	GetKeyword() *string
	SetPageNum(v int32) *QueryDatasetListRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryDatasetListRequest
	GetPageSize() *int32
	SetWithChildren(v bool) *QueryDatasetListRequest
	GetWithChildren() *bool
	SetWorkspaceId(v string) *QueryDatasetListRequest
	GetWorkspaceId() *string
}

type QueryDatasetListRequest struct {
	// The ID of the request.
	//
	// example:
	//
	// schemaad8aad00-9c55-4984-a767-b4e0ec60****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// Information about the directory where the dataset is located
	//
	// example:
	//
	// Queries the datasets of a specified workspace. The datasets are sorted in descending order by creation time.
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The ID of the workspace.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// Specifies the directory ID.
	//
	// 	- If this field is not empty, all datasets in the directory are obtained.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// true
	WithChildren *bool `json:"WithChildren,omitempty" xml:"WithChildren,omitempty"`
	// The name of the data source.
	//
	// This parameter is required.
	//
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryDatasetListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDatasetListRequest) GoString() string {
	return s.String()
}

func (s *QueryDatasetListRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *QueryDatasetListRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *QueryDatasetListRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryDatasetListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryDatasetListRequest) GetWithChildren() *bool {
	return s.WithChildren
}

func (s *QueryDatasetListRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryDatasetListRequest) SetDirectoryId(v string) *QueryDatasetListRequest {
	s.DirectoryId = &v
	return s
}

func (s *QueryDatasetListRequest) SetKeyword(v string) *QueryDatasetListRequest {
	s.Keyword = &v
	return s
}

func (s *QueryDatasetListRequest) SetPageNum(v int32) *QueryDatasetListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryDatasetListRequest) SetPageSize(v int32) *QueryDatasetListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryDatasetListRequest) SetWithChildren(v bool) *QueryDatasetListRequest {
	s.WithChildren = &v
	return s
}

func (s *QueryDatasetListRequest) SetWorkspaceId(v string) *QueryDatasetListRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryDatasetListRequest) Validate() error {
	return dara.Validate(s)
}
