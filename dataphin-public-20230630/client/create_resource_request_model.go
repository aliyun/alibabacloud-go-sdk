// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommand(v *CreateResourceRequestCreateCommand) *CreateResourceRequest
	GetCreateCommand() *CreateResourceRequestCreateCommand
	SetOpTenantId(v int64) *CreateResourceRequest
	GetOpTenantId() *int64
}

type CreateResourceRequest struct {
	// This parameter is required.
	CreateCommand *CreateResourceRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceRequest) GetCreateCommand() *CreateResourceRequestCreateCommand {
	return s.CreateCommand
}

func (s *CreateResourceRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateResourceRequest) SetCreateCommand(v *CreateResourceRequestCreateCommand) *CreateResourceRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateResourceRequest) SetOpTenantId(v int64) *CreateResourceRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateResourceRequest) Validate() error {
	if s.CreateCommand != nil {
		if err := s.CreateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateResourceRequestCreateCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// xx 测试
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MAX_COMPUTE
	ComputeEngineType *string `json:"ComputeEngineType,omitempty" xml:"ComputeEngineType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xx 测试
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
	// udf_sleep.jar
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 711833
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// JAR
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 300011448/711833/cdcd1c44-f1ee-46bb-b318-1d123dbabf6c
	StorageAddress *string `json:"StorageAddress,omitempty" xml:"StorageAddress,omitempty"`
}

func (s CreateResourceRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateResourceRequestCreateCommand) GetComment() *string {
	return s.Comment
}

func (s *CreateResourceRequestCreateCommand) GetComputeEngineType() *string {
	return s.ComputeEngineType
}

func (s *CreateResourceRequestCreateCommand) GetDescription() *string {
	return s.Description
}

func (s *CreateResourceRequestCreateCommand) GetDirectory() *string {
	return s.Directory
}

func (s *CreateResourceRequestCreateCommand) GetName() *string {
	return s.Name
}

func (s *CreateResourceRequestCreateCommand) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateResourceRequestCreateCommand) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateResourceRequestCreateCommand) GetStorageAddress() *string {
	return s.StorageAddress
}

func (s *CreateResourceRequestCreateCommand) SetComment(v string) *CreateResourceRequestCreateCommand {
	s.Comment = &v
	return s
}

func (s *CreateResourceRequestCreateCommand) SetComputeEngineType(v string) *CreateResourceRequestCreateCommand {
	s.ComputeEngineType = &v
	return s
}

func (s *CreateResourceRequestCreateCommand) SetDescription(v string) *CreateResourceRequestCreateCommand {
	s.Description = &v
	return s
}

func (s *CreateResourceRequestCreateCommand) SetDirectory(v string) *CreateResourceRequestCreateCommand {
	s.Directory = &v
	return s
}

func (s *CreateResourceRequestCreateCommand) SetName(v string) *CreateResourceRequestCreateCommand {
	s.Name = &v
	return s
}

func (s *CreateResourceRequestCreateCommand) SetProjectId(v int64) *CreateResourceRequestCreateCommand {
	s.ProjectId = &v
	return s
}

func (s *CreateResourceRequestCreateCommand) SetResourceType(v string) *CreateResourceRequestCreateCommand {
	s.ResourceType = &v
	return s
}

func (s *CreateResourceRequestCreateCommand) SetStorageAddress(v string) *CreateResourceRequestCreateCommand {
	s.StorageAddress = &v
	return s
}

func (s *CreateResourceRequestCreateCommand) Validate() error {
	return dara.Validate(s)
}
