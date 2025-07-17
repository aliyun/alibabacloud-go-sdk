// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeploymentPackageFilesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessId(v int64) *ListDeploymentPackageFilesShrinkRequest
	GetBusinessId() *int64
	SetChangeType(v int32) *ListDeploymentPackageFilesShrinkRequest
	GetChangeType() *int32
	SetCommitFrom(v string) *ListDeploymentPackageFilesShrinkRequest
	GetCommitFrom() *string
	SetCommitTo(v string) *ListDeploymentPackageFilesShrinkRequest
	GetCommitTo() *string
	SetCommitUserId(v string) *ListDeploymentPackageFilesShrinkRequest
	GetCommitUserId() *string
	SetFileIdsShrink(v string) *ListDeploymentPackageFilesShrinkRequest
	GetFileIdsShrink() *string
	SetFileName(v string) *ListDeploymentPackageFilesShrinkRequest
	GetFileName() *string
	SetFileType(v int32) *ListDeploymentPackageFilesShrinkRequest
	GetFileType() *int32
	SetPageNumber(v int32) *ListDeploymentPackageFilesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDeploymentPackageFilesShrinkRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListDeploymentPackageFilesShrinkRequest
	GetProjectId() *int64
	SetSolutionId(v int64) *ListDeploymentPackageFilesShrinkRequest
	GetSolutionId() *int64
}

type ListDeploymentPackageFilesShrinkRequest struct {
	// The workflow ID. You can call the [ListBusiness](https://help.aliyun.com/document_detail/173945.html) operation to query the workflow ID by name.
	//
	// example:
	//
	// 100001
	BusinessId *int64 `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	// The change type. Valid values:
	//
	// 	- 0: addition
	//
	// 	- 1: update
	//
	// 	- 2: deletion
	//
	// example:
	//
	// 0
	ChangeType *int32 `json:"ChangeType,omitempty" xml:"ChangeType,omitempty"`
	// The start date for committing. Specify the date in the yyyy-MM-dd format.
	//
	// example:
	//
	// 2025-01-01
	CommitFrom *string `json:"CommitFrom,omitempty" xml:"CommitFrom,omitempty"`
	// The end date (included) for committing. Specify the date in the yyyy-MM-dd format.
	//
	// example:
	//
	// 2025-01-31
	CommitTo *string `json:"CommitTo,omitempty" xml:"CommitTo,omitempty"`
	// The ID of the user who commits the file.
	//
	// example:
	//
	// 2003****
	CommitUserId *string `json:"CommitUserId,omitempty" xml:"CommitUserId,omitempty"`
	// The IDs of the files to be queried.
	FileIdsShrink *string `json:"FileIds,omitempty" xml:"FileIds,omitempty"`
	// The name of the file.
	//
	// example:
	//
	// Filename
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The type of the code for the file.
	//
	// The code for files varies based on the file type. For more information, see [DataWorks nodes](https://help.aliyun.com/document_detail/600169.html). You can call the [ListFileType](https://help.aliyun.com/document_detail/212428.html) operation to query the type of the code for the file.
	//
	// example:
	//
	// 10
	FileType *int32 `json:"FileType,omitempty" xml:"FileType,omitempty"`
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
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The solution ID.
	//
	// example:
	//
	// 8065
	SolutionId *int64 `json:"SolutionId,omitempty" xml:"SolutionId,omitempty"`
}

func (s ListDeploymentPackageFilesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentPackageFilesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDeploymentPackageFilesShrinkRequest) GetBusinessId() *int64 {
	return s.BusinessId
}

func (s *ListDeploymentPackageFilesShrinkRequest) GetChangeType() *int32 {
	return s.ChangeType
}

func (s *ListDeploymentPackageFilesShrinkRequest) GetCommitFrom() *string {
	return s.CommitFrom
}

func (s *ListDeploymentPackageFilesShrinkRequest) GetCommitTo() *string {
	return s.CommitTo
}

func (s *ListDeploymentPackageFilesShrinkRequest) GetCommitUserId() *string {
	return s.CommitUserId
}

func (s *ListDeploymentPackageFilesShrinkRequest) GetFileIdsShrink() *string {
	return s.FileIdsShrink
}

func (s *ListDeploymentPackageFilesShrinkRequest) GetFileName() *string {
	return s.FileName
}

func (s *ListDeploymentPackageFilesShrinkRequest) GetFileType() *int32 {
	return s.FileType
}

func (s *ListDeploymentPackageFilesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDeploymentPackageFilesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDeploymentPackageFilesShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDeploymentPackageFilesShrinkRequest) GetSolutionId() *int64 {
	return s.SolutionId
}

func (s *ListDeploymentPackageFilesShrinkRequest) SetBusinessId(v int64) *ListDeploymentPackageFilesShrinkRequest {
	s.BusinessId = &v
	return s
}

func (s *ListDeploymentPackageFilesShrinkRequest) SetChangeType(v int32) *ListDeploymentPackageFilesShrinkRequest {
	s.ChangeType = &v
	return s
}

func (s *ListDeploymentPackageFilesShrinkRequest) SetCommitFrom(v string) *ListDeploymentPackageFilesShrinkRequest {
	s.CommitFrom = &v
	return s
}

func (s *ListDeploymentPackageFilesShrinkRequest) SetCommitTo(v string) *ListDeploymentPackageFilesShrinkRequest {
	s.CommitTo = &v
	return s
}

func (s *ListDeploymentPackageFilesShrinkRequest) SetCommitUserId(v string) *ListDeploymentPackageFilesShrinkRequest {
	s.CommitUserId = &v
	return s
}

func (s *ListDeploymentPackageFilesShrinkRequest) SetFileIdsShrink(v string) *ListDeploymentPackageFilesShrinkRequest {
	s.FileIdsShrink = &v
	return s
}

func (s *ListDeploymentPackageFilesShrinkRequest) SetFileName(v string) *ListDeploymentPackageFilesShrinkRequest {
	s.FileName = &v
	return s
}

func (s *ListDeploymentPackageFilesShrinkRequest) SetFileType(v int32) *ListDeploymentPackageFilesShrinkRequest {
	s.FileType = &v
	return s
}

func (s *ListDeploymentPackageFilesShrinkRequest) SetPageNumber(v int32) *ListDeploymentPackageFilesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDeploymentPackageFilesShrinkRequest) SetPageSize(v int32) *ListDeploymentPackageFilesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListDeploymentPackageFilesShrinkRequest) SetProjectId(v int64) *ListDeploymentPackageFilesShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDeploymentPackageFilesShrinkRequest) SetSolutionId(v int64) *ListDeploymentPackageFilesShrinkRequest {
	s.SolutionId = &v
	return s
}

func (s *ListDeploymentPackageFilesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
