// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *SubmitFileRequest
	GetComment() *string
	SetFileId(v int64) *SubmitFileRequest
	GetFileId() *int64
	SetProjectId(v int64) *SubmitFileRequest
	GetProjectId() *int64
	SetProjectIdentifier(v string) *SubmitFileRequest
	GetProjectIdentifier() *string
	SetSkipAllDeployFileExtensions(v bool) *SubmitFileRequest
	GetSkipAllDeployFileExtensions() *bool
}

type SubmitFileRequest struct {
	// The comment for the submission.
	//
	// example:
	//
	// Submit a task for the first time
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the file. Obtain this ID by calling the [ListFiles](https://help.aliyun.com/document_detail/173942.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000000
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the DataWorks Console and go to the Workspace Configurations page to obtain the workspace ID. Specify either this parameter or `ProjectIdentifier` to identify the DataWorks workspace.
	//
	// example:
	//
	// 100001
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the DataWorks workspace. You can log on to the DataWorks Console and go to the Workspace Configurations page to obtain the workspace name. Specify either this parameter or `ProjectId` to identify the DataWorks workspace.
	//
	// example:
	//
	// dw_project
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
	// Specifies whether to skip the pre-deployment check after the file is submitted.
	//
	// - false: Do not skip. After the file is submitted, the pre-deployment check process is automatically triggered. The file becomes deployable only after it passes the check.
	//
	// - true: Skip. The pre-deployment check process is not triggered after the file is submitted. You can proceed directly with the deployment process.
	//
	// example:
	//
	// false
	SkipAllDeployFileExtensions *bool `json:"SkipAllDeployFileExtensions,omitempty" xml:"SkipAllDeployFileExtensions,omitempty"`
}

func (s SubmitFileRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitFileRequest) GoString() string {
	return s.String()
}

func (s *SubmitFileRequest) GetComment() *string {
	return s.Comment
}

func (s *SubmitFileRequest) GetFileId() *int64 {
	return s.FileId
}

func (s *SubmitFileRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *SubmitFileRequest) GetProjectIdentifier() *string {
	return s.ProjectIdentifier
}

func (s *SubmitFileRequest) GetSkipAllDeployFileExtensions() *bool {
	return s.SkipAllDeployFileExtensions
}

func (s *SubmitFileRequest) SetComment(v string) *SubmitFileRequest {
	s.Comment = &v
	return s
}

func (s *SubmitFileRequest) SetFileId(v int64) *SubmitFileRequest {
	s.FileId = &v
	return s
}

func (s *SubmitFileRequest) SetProjectId(v int64) *SubmitFileRequest {
	s.ProjectId = &v
	return s
}

func (s *SubmitFileRequest) SetProjectIdentifier(v string) *SubmitFileRequest {
	s.ProjectIdentifier = &v
	return s
}

func (s *SubmitFileRequest) SetSkipAllDeployFileExtensions(v bool) *SubmitFileRequest {
	s.SkipAllDeployFileExtensions = &v
	return s
}

func (s *SubmitFileRequest) Validate() error {
	return dara.Validate(s)
}
