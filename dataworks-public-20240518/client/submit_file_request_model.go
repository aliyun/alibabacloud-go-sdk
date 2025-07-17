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
	// example:
	//
	// Submit a task for the first time
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1000000
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// example:
	//
	// 100001
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// dw_project
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
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
