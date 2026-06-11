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
	// The ID of the folder.
	//
	// - If you specify this parameter, all datasets in the folder are returned.
	//
	// example:
	//
	// schemaad8aad00-9c55-4984-a767-b4e0ec60****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The keyword used to search for datasets by name.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number of the dataset list.
	//
	// - Start value: 1
	//
	// - Default value: 1
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page.
	//
	// - Default value: 10
	//
	// - Maximum value: 1000
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether to recursively include datasets in subdirectories. Valid values:
	//
	// - true: Returns all datasets in the folder specified by DirectoryId and its subdirectories.
	//
	// - false: Returns only the datasets in the folder specified by DirectoryId.
	//
	// example:
	//
	// true
	WithChildren *bool `json:"WithChildren,omitempty" xml:"WithChildren,omitempty"`
	// The workspace ID.
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
