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
	// The file ID. You can call [ListFiles](https://help.aliyun.com/document_detail/173942.html) to query the file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100000001
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The page number for pagination.
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
	// The DataWorks workspace ID. To find this, click the wrench icon in the upper-right corner and navigate to the workspace management page.
	//
	// example:
	//
	// 100001
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The unique identifier of the DataWorks workspace, which is the identifier at the top of the Data Studio page where you switch workspaces.
	//
	// Either this parameter or ProjectId must be specified to identify the target DataWorks workspace for this API call.
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
