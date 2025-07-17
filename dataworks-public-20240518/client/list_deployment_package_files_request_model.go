// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeploymentPackageFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessId(v int64) *ListDeploymentPackageFilesRequest
	GetBusinessId() *int64
	SetChangeType(v int32) *ListDeploymentPackageFilesRequest
	GetChangeType() *int32
	SetCommitFrom(v string) *ListDeploymentPackageFilesRequest
	GetCommitFrom() *string
	SetCommitTo(v string) *ListDeploymentPackageFilesRequest
	GetCommitTo() *string
	SetCommitUserId(v string) *ListDeploymentPackageFilesRequest
	GetCommitUserId() *string
	SetFileIds(v []*string) *ListDeploymentPackageFilesRequest
	GetFileIds() []*string
	SetFileName(v string) *ListDeploymentPackageFilesRequest
	GetFileName() *string
	SetFileType(v int32) *ListDeploymentPackageFilesRequest
	GetFileType() *int32
	SetPageNumber(v int32) *ListDeploymentPackageFilesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDeploymentPackageFilesRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListDeploymentPackageFilesRequest
	GetProjectId() *int64
	SetSolutionId(v int64) *ListDeploymentPackageFilesRequest
	GetSolutionId() *int64
}

type ListDeploymentPackageFilesRequest struct {
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
	FileIds []*string `json:"FileIds,omitempty" xml:"FileIds,omitempty" type:"Repeated"`
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

func (s ListDeploymentPackageFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentPackageFilesRequest) GoString() string {
	return s.String()
}

func (s *ListDeploymentPackageFilesRequest) GetBusinessId() *int64 {
	return s.BusinessId
}

func (s *ListDeploymentPackageFilesRequest) GetChangeType() *int32 {
	return s.ChangeType
}

func (s *ListDeploymentPackageFilesRequest) GetCommitFrom() *string {
	return s.CommitFrom
}

func (s *ListDeploymentPackageFilesRequest) GetCommitTo() *string {
	return s.CommitTo
}

func (s *ListDeploymentPackageFilesRequest) GetCommitUserId() *string {
	return s.CommitUserId
}

func (s *ListDeploymentPackageFilesRequest) GetFileIds() []*string {
	return s.FileIds
}

func (s *ListDeploymentPackageFilesRequest) GetFileName() *string {
	return s.FileName
}

func (s *ListDeploymentPackageFilesRequest) GetFileType() *int32 {
	return s.FileType
}

func (s *ListDeploymentPackageFilesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDeploymentPackageFilesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDeploymentPackageFilesRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDeploymentPackageFilesRequest) GetSolutionId() *int64 {
	return s.SolutionId
}

func (s *ListDeploymentPackageFilesRequest) SetBusinessId(v int64) *ListDeploymentPackageFilesRequest {
	s.BusinessId = &v
	return s
}

func (s *ListDeploymentPackageFilesRequest) SetChangeType(v int32) *ListDeploymentPackageFilesRequest {
	s.ChangeType = &v
	return s
}

func (s *ListDeploymentPackageFilesRequest) SetCommitFrom(v string) *ListDeploymentPackageFilesRequest {
	s.CommitFrom = &v
	return s
}

func (s *ListDeploymentPackageFilesRequest) SetCommitTo(v string) *ListDeploymentPackageFilesRequest {
	s.CommitTo = &v
	return s
}

func (s *ListDeploymentPackageFilesRequest) SetCommitUserId(v string) *ListDeploymentPackageFilesRequest {
	s.CommitUserId = &v
	return s
}

func (s *ListDeploymentPackageFilesRequest) SetFileIds(v []*string) *ListDeploymentPackageFilesRequest {
	s.FileIds = v
	return s
}

func (s *ListDeploymentPackageFilesRequest) SetFileName(v string) *ListDeploymentPackageFilesRequest {
	s.FileName = &v
	return s
}

func (s *ListDeploymentPackageFilesRequest) SetFileType(v int32) *ListDeploymentPackageFilesRequest {
	s.FileType = &v
	return s
}

func (s *ListDeploymentPackageFilesRequest) SetPageNumber(v int32) *ListDeploymentPackageFilesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDeploymentPackageFilesRequest) SetPageSize(v int32) *ListDeploymentPackageFilesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDeploymentPackageFilesRequest) SetProjectId(v int64) *ListDeploymentPackageFilesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDeploymentPackageFilesRequest) SetSolutionId(v int64) *ListDeploymentPackageFilesRequest {
	s.SolutionId = &v
	return s
}

func (s *ListDeploymentPackageFilesRequest) Validate() error {
	return dara.Validate(s)
}
