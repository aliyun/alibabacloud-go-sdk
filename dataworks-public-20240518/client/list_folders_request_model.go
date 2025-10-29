// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFoldersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListFoldersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListFoldersRequest
	GetPageSize() *int32
	SetParentFolderPath(v string) *ListFoldersRequest
	GetParentFolderPath() *string
	SetProjectId(v int64) *ListFoldersRequest
	GetProjectId() *int64
	SetProjectIdentifier(v string) *ListFoldersRequest
	GetProjectIdentifier() *string
}

type ListFoldersRequest struct {
	// The page number of the request. Used for pagination.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page. Default value: 10. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The path of the parent folder.
	//
	// This parameter is required.
	//
	// example:
	//
	// Business_process/my_first_business_process/MaxCompute
	ParentFolderPath *string `json:"ParentFolderPath,omitempty" xml:"ParentFolderPath,omitempty"`
	// The DataWorks workspace ID. You can log on to the DataWorks console and go to the Workspace page to query the ID. You must specify either this parameter or the ProjectIdentifier parameter to identify the DataWorks workspace when you call this operation.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the DataWorks workspace. You can log on to the DataWorks console and go to the Workspace page to query the workspace name. You must specify either this parameter or the ProjectId parameter to identify the DataWorks workspace when you call this operation.
	//
	// example:
	//
	// dw_project
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
}

func (s ListFoldersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFoldersRequest) GoString() string {
	return s.String()
}

func (s *ListFoldersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListFoldersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFoldersRequest) GetParentFolderPath() *string {
	return s.ParentFolderPath
}

func (s *ListFoldersRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListFoldersRequest) GetProjectIdentifier() *string {
	return s.ProjectIdentifier
}

func (s *ListFoldersRequest) SetPageNumber(v int32) *ListFoldersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFoldersRequest) SetPageSize(v int32) *ListFoldersRequest {
	s.PageSize = &v
	return s
}

func (s *ListFoldersRequest) SetParentFolderPath(v string) *ListFoldersRequest {
	s.ParentFolderPath = &v
	return s
}

func (s *ListFoldersRequest) SetProjectId(v int64) *ListFoldersRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFoldersRequest) SetProjectIdentifier(v string) *ListFoldersRequest {
	s.ProjectIdentifier = &v
	return s
}

func (s *ListFoldersRequest) Validate() error {
	return dara.Validate(s)
}
