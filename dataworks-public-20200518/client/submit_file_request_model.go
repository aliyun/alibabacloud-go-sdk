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
	// The description of the commit operation.
	//
	// example:
	//
	// Submit a task for the first time
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The file ID. You can call the [ListFiles](https://help.aliyun.com/document_detail/173942.html) operation to query the file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000000
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The DataWorks workspace ID. You can log on to the DataWorks console and go to the Workspace page to query the ID. You must configure either this parameter or the ProjectIdentifier parameter to determine the DataWorks workspace to which the operation is applied.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the DataWorks workspace. You can log on to the DataWorks console and go to the Workspace page to obtain the workspace name. You must configure either this parameter or the ProjectId parameter to determine the DataWorks workspace to which the operation is applied.
	//
	// example:
	//
	// dw_project
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
	// Specifies whether to skip the pre-publish check after the file is committed. Valid values:
	//
	// 	- false: indicates that the pre-publish check is not skipped. After the file is committed, the pre-publish check is automatically triggered. The file can be deployed only after the file passes the check.
	//
	// 	- true: indicates that the pre-publish check is skipped. After the file is submitted, the pre-publish check process is not triggered. You can directly deploy the file.
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
