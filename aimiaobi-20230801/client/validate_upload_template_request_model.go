// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateUploadTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileKey(v string) *ValidateUploadTemplateRequest
	GetFileKey() *string
	SetTaskType(v string) *ValidateUploadTemplateRequest
	GetTaskType() *string
	SetTemplateType(v string) *ValidateUploadTemplateRequest
	GetTemplateType() *string
	SetWorkspaceId(v string) *ValidateUploadTemplateRequest
	GetWorkspaceId() *string
}

type ValidateUploadTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Content
	FileKey *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	// example:
	//
	// lightAppSass
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Content
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ValidateUploadTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidateUploadTemplateRequest) GoString() string {
	return s.String()
}

func (s *ValidateUploadTemplateRequest) GetFileKey() *string {
	return s.FileKey
}

func (s *ValidateUploadTemplateRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *ValidateUploadTemplateRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *ValidateUploadTemplateRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ValidateUploadTemplateRequest) SetFileKey(v string) *ValidateUploadTemplateRequest {
	s.FileKey = &v
	return s
}

func (s *ValidateUploadTemplateRequest) SetTaskType(v string) *ValidateUploadTemplateRequest {
	s.TaskType = &v
	return s
}

func (s *ValidateUploadTemplateRequest) SetTemplateType(v string) *ValidateUploadTemplateRequest {
	s.TemplateType = &v
	return s
}

func (s *ValidateUploadTemplateRequest) SetWorkspaceId(v string) *ValidateUploadTemplateRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ValidateUploadTemplateRequest) Validate() error {
	return dara.Validate(s)
}
