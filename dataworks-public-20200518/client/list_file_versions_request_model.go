// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFileVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v int64) *ListFileVersionsRequest
	GetFileId() *int64
	SetPageNumber(v int32) *ListFileVersionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListFileVersionsRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListFileVersionsRequest
	GetProjectId() *int64
	SetProjectIdentifier(v string) *ListFileVersionsRequest
	GetProjectIdentifier() *string
}

type ListFileVersionsRequest struct {
	// The ID of the file. You can call the [ListFiles](https://help.aliyun.com/document_detail/173942.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100000001
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The DataWorks workspace ID. You can click the Workspace Manage icon in the upper-right corner of the DataStudio page to go to the Workspace page and view the workspace ID.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The unique identifier of the DataWorks workspace. You can click the identifier in the upper-left corner of the DataStudio page to switch to another workspace.
	//
	// You must specify either this parameter or ProjectId to determine the DataWorks workspace to which the operation is applied.
	//
	// example:
	//
	// dw_project
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
}

func (s ListFileVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFileVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListFileVersionsRequest) GetFileId() *int64 {
	return s.FileId
}

func (s *ListFileVersionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListFileVersionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFileVersionsRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListFileVersionsRequest) GetProjectIdentifier() *string {
	return s.ProjectIdentifier
}

func (s *ListFileVersionsRequest) SetFileId(v int64) *ListFileVersionsRequest {
	s.FileId = &v
	return s
}

func (s *ListFileVersionsRequest) SetPageNumber(v int32) *ListFileVersionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFileVersionsRequest) SetPageSize(v int32) *ListFileVersionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListFileVersionsRequest) SetProjectId(v int64) *ListFileVersionsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFileVersionsRequest) SetProjectIdentifier(v string) *ListFileVersionsRequest {
	s.ProjectIdentifier = &v
	return s
}

func (s *ListFileVersionsRequest) Validate() error {
	return dara.Validate(s)
}
