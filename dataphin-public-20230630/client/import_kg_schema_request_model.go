// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportKgSchemaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImportCommand(v *ImportKgSchemaRequestImportCommand) *ImportKgSchemaRequest
	GetImportCommand() *ImportKgSchemaRequestImportCommand
	SetOpTenantId(v int64) *ImportKgSchemaRequest
	GetOpTenantId() *int64
	SetWorkspaceId(v string) *ImportKgSchemaRequest
	GetWorkspaceId() *string
}

type ImportKgSchemaRequest struct {
	// The instruction for importing the knowledge graph definition.
	//
	// This parameter is required.
	//
	// example:
	//
	// f1d4559a4db044158305e2d89bccf81f
	ImportCommand *ImportKgSchemaRequestImportCommand `json:"ImportCommand,omitempty" xml:"ImportCommand,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// f1d4559a4db044158305e2d89bccf81f
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ImportKgSchemaRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportKgSchemaRequest) GoString() string {
	return s.String()
}

func (s *ImportKgSchemaRequest) GetImportCommand() *ImportKgSchemaRequestImportCommand {
	return s.ImportCommand
}

func (s *ImportKgSchemaRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ImportKgSchemaRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ImportKgSchemaRequest) SetImportCommand(v *ImportKgSchemaRequestImportCommand) *ImportKgSchemaRequest {
	s.ImportCommand = v
	return s
}

func (s *ImportKgSchemaRequest) SetOpTenantId(v int64) *ImportKgSchemaRequest {
	s.OpTenantId = &v
	return s
}

func (s *ImportKgSchemaRequest) SetWorkspaceId(v string) *ImportKgSchemaRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ImportKgSchemaRequest) Validate() error {
	if s.ImportCommand != nil {
		if err := s.ImportCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImportKgSchemaRequestImportCommand struct {
	// The knowledge graph definition content converted based on the specified format.
	//
	// example:
	//
	// name:xxx
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The format of the knowledge graph definition content. Valid values: json and yaml. Default value: yaml.
	//
	// example:
	//
	// yaml
	InputFormat *string `json:"InputFormat,omitempty" xml:"InputFormat,omitempty"`
	// The merge strategy for the knowledge graph definition content. Valid values: replace and merge. Default value: replace.
	//
	// example:
	//
	// replace
	MergeStrategy *string `json:"MergeStrategy,omitempty" xml:"MergeStrategy,omitempty"`
}

func (s ImportKgSchemaRequestImportCommand) String() string {
	return dara.Prettify(s)
}

func (s ImportKgSchemaRequestImportCommand) GoString() string {
	return s.String()
}

func (s *ImportKgSchemaRequestImportCommand) GetContent() *string {
	return s.Content
}

func (s *ImportKgSchemaRequestImportCommand) GetInputFormat() *string {
	return s.InputFormat
}

func (s *ImportKgSchemaRequestImportCommand) GetMergeStrategy() *string {
	return s.MergeStrategy
}

func (s *ImportKgSchemaRequestImportCommand) SetContent(v string) *ImportKgSchemaRequestImportCommand {
	s.Content = &v
	return s
}

func (s *ImportKgSchemaRequestImportCommand) SetInputFormat(v string) *ImportKgSchemaRequestImportCommand {
	s.InputFormat = &v
	return s
}

func (s *ImportKgSchemaRequestImportCommand) SetMergeStrategy(v string) *ImportKgSchemaRequestImportCommand {
	s.MergeStrategy = &v
	return s
}

func (s *ImportKgSchemaRequestImportCommand) Validate() error {
	return dara.Validate(s)
}
