// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAdHocFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommand(v *CreateAdHocFileRequestCreateCommand) *CreateAdHocFileRequest
	GetCreateCommand() *CreateAdHocFileRequestCreateCommand
	SetOpTenantId(v int64) *CreateAdHocFileRequest
	GetOpTenantId() *int64
}

type CreateAdHocFileRequest struct {
	// This parameter is required.
	CreateCommand *CreateAdHocFileRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateAdHocFileRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAdHocFileRequest) GoString() string {
	return s.String()
}

func (s *CreateAdHocFileRequest) GetCreateCommand() *CreateAdHocFileRequestCreateCommand {
	return s.CreateCommand
}

func (s *CreateAdHocFileRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateAdHocFileRequest) SetCreateCommand(v *CreateAdHocFileRequestCreateCommand) *CreateAdHocFileRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateAdHocFileRequest) SetOpTenantId(v int64) *CreateAdHocFileRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateAdHocFileRequest) Validate() error {
	return dara.Validate(s)
}

type CreateAdHocFileRequestCreateCommand struct {
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
	// /xx1/xx2/
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_xx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 11212133
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateAdHocFileRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateAdHocFileRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateAdHocFileRequestCreateCommand) GetContent() *string {
	return s.Content
}

func (s *CreateAdHocFileRequestCreateCommand) GetDirectory() *string {
	return s.Directory
}

func (s *CreateAdHocFileRequestCreateCommand) GetName() *string {
	return s.Name
}

func (s *CreateAdHocFileRequestCreateCommand) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateAdHocFileRequestCreateCommand) SetContent(v string) *CreateAdHocFileRequestCreateCommand {
	s.Content = &v
	return s
}

func (s *CreateAdHocFileRequestCreateCommand) SetDirectory(v string) *CreateAdHocFileRequestCreateCommand {
	s.Directory = &v
	return s
}

func (s *CreateAdHocFileRequestCreateCommand) SetName(v string) *CreateAdHocFileRequestCreateCommand {
	s.Name = &v
	return s
}

func (s *CreateAdHocFileRequestCreateCommand) SetProjectId(v int64) *CreateAdHocFileRequestCreateCommand {
	s.ProjectId = &v
	return s
}

func (s *CreateAdHocFileRequestCreateCommand) Validate() error {
	return dara.Validate(s)
}
