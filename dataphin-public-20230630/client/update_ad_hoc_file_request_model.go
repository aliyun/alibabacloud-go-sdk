// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAdHocFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateAdHocFileRequest
	GetOpTenantId() *int64
	SetUpdateCommand(v *UpdateAdHocFileRequestUpdateCommand) *UpdateAdHocFileRequest
	GetUpdateCommand() *UpdateAdHocFileRequestUpdateCommand
}

type UpdateAdHocFileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommand *UpdateAdHocFileRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateAdHocFileRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAdHocFileRequest) GoString() string {
	return s.String()
}

func (s *UpdateAdHocFileRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateAdHocFileRequest) GetUpdateCommand() *UpdateAdHocFileRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateAdHocFileRequest) SetOpTenantId(v int64) *UpdateAdHocFileRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateAdHocFileRequest) SetUpdateCommand(v *UpdateAdHocFileRequestUpdateCommand) *UpdateAdHocFileRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateAdHocFileRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateAdHocFileRequestUpdateCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// select 1;
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2311113
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1212313
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s UpdateAdHocFileRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateAdHocFileRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateAdHocFileRequestUpdateCommand) GetContent() *string {
	return s.Content
}

func (s *UpdateAdHocFileRequestUpdateCommand) GetFileId() *int64 {
	return s.FileId
}

func (s *UpdateAdHocFileRequestUpdateCommand) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateAdHocFileRequestUpdateCommand) SetContent(v string) *UpdateAdHocFileRequestUpdateCommand {
	s.Content = &v
	return s
}

func (s *UpdateAdHocFileRequestUpdateCommand) SetFileId(v int64) *UpdateAdHocFileRequestUpdateCommand {
	s.FileId = &v
	return s
}

func (s *UpdateAdHocFileRequestUpdateCommand) SetProjectId(v int64) *UpdateAdHocFileRequestUpdateCommand {
	s.ProjectId = &v
	return s
}

func (s *UpdateAdHocFileRequestUpdateCommand) Validate() error {
	return dara.Validate(s)
}
