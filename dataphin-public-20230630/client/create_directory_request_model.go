// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommand(v *CreateDirectoryRequestCreateCommand) *CreateDirectoryRequest
	GetCreateCommand() *CreateDirectoryRequestCreateCommand
	SetOpTenantId(v int64) *CreateDirectoryRequest
	GetOpTenantId() *int64
}

type CreateDirectoryRequest struct {
	// This parameter is required.
	CreateCommand *CreateDirectoryRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDirectoryRequest) GoString() string {
	return s.String()
}

func (s *CreateDirectoryRequest) GetCreateCommand() *CreateDirectoryRequestCreateCommand {
	return s.CreateCommand
}

func (s *CreateDirectoryRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateDirectoryRequest) SetCreateCommand(v *CreateDirectoryRequestCreateCommand) *CreateDirectoryRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateDirectoryRequest) SetOpTenantId(v int64) *CreateDirectoryRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateDirectoryRequest) Validate() error {
	if s.CreateCommand != nil {
		if err := s.CreateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDirectoryRequestCreateCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// tempCode
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// /
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
	// 1212344
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateDirectoryRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateDirectoryRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateDirectoryRequestCreateCommand) GetCategory() *string {
	return s.Category
}

func (s *CreateDirectoryRequestCreateCommand) GetDirectory() *string {
	return s.Directory
}

func (s *CreateDirectoryRequestCreateCommand) GetName() *string {
	return s.Name
}

func (s *CreateDirectoryRequestCreateCommand) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateDirectoryRequestCreateCommand) SetCategory(v string) *CreateDirectoryRequestCreateCommand {
	s.Category = &v
	return s
}

func (s *CreateDirectoryRequestCreateCommand) SetDirectory(v string) *CreateDirectoryRequestCreateCommand {
	s.Directory = &v
	return s
}

func (s *CreateDirectoryRequestCreateCommand) SetName(v string) *CreateDirectoryRequestCreateCommand {
	s.Name = &v
	return s
}

func (s *CreateDirectoryRequestCreateCommand) SetProjectId(v int64) *CreateDirectoryRequestCreateCommand {
	s.ProjectId = &v
	return s
}

func (s *CreateDirectoryRequestCreateCommand) Validate() error {
	return dara.Validate(s)
}
