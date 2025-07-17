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
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Business_process/my_first_business_process/MaxCompute
	ParentFolderPath *string `json:"ParentFolderPath,omitempty" xml:"ParentFolderPath,omitempty"`
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
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
