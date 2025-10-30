// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateResourceRequest
	GetOpTenantId() *int64
	SetUpdateCommand(v *UpdateResourceRequestUpdateCommand) *UpdateResourceRequest
	GetUpdateCommand() *UpdateResourceRequestUpdateCommand
}

type UpdateResourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommand *UpdateResourceRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateResourceRequest) GetUpdateCommand() *UpdateResourceRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateResourceRequest) SetOpTenantId(v int64) *UpdateResourceRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateResourceRequest) SetUpdateCommand(v *UpdateResourceRequestUpdateCommand) *UpdateResourceRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateResourceRequest) Validate() error {
	if s.UpdateCommand != nil {
		if err := s.UpdateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateResourceRequestUpdateCommand struct {
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
	// 10112101
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
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
	// 300011448/711833/cdcd1c44-f1ee-46bb-b318-1d123dbabf6c
	StorageAddress *string `json:"StorageAddress,omitempty" xml:"StorageAddress,omitempty"`
}

func (s UpdateResourceRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateResourceRequestUpdateCommand) GetComment() *string {
	return s.Comment
}

func (s *UpdateResourceRequestUpdateCommand) GetComputeEngineType() *string {
	return s.ComputeEngineType
}

func (s *UpdateResourceRequestUpdateCommand) GetDescription() *string {
	return s.Description
}

func (s *UpdateResourceRequestUpdateCommand) GetId() *int64 {
	return s.Id
}

func (s *UpdateResourceRequestUpdateCommand) GetName() *string {
	return s.Name
}

func (s *UpdateResourceRequestUpdateCommand) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateResourceRequestUpdateCommand) GetStorageAddress() *string {
	return s.StorageAddress
}

func (s *UpdateResourceRequestUpdateCommand) SetComment(v string) *UpdateResourceRequestUpdateCommand {
	s.Comment = &v
	return s
}

func (s *UpdateResourceRequestUpdateCommand) SetComputeEngineType(v string) *UpdateResourceRequestUpdateCommand {
	s.ComputeEngineType = &v
	return s
}

func (s *UpdateResourceRequestUpdateCommand) SetDescription(v string) *UpdateResourceRequestUpdateCommand {
	s.Description = &v
	return s
}

func (s *UpdateResourceRequestUpdateCommand) SetId(v int64) *UpdateResourceRequestUpdateCommand {
	s.Id = &v
	return s
}

func (s *UpdateResourceRequestUpdateCommand) SetName(v string) *UpdateResourceRequestUpdateCommand {
	s.Name = &v
	return s
}

func (s *UpdateResourceRequestUpdateCommand) SetProjectId(v int64) *UpdateResourceRequestUpdateCommand {
	s.ProjectId = &v
	return s
}

func (s *UpdateResourceRequestUpdateCommand) SetStorageAddress(v string) *UpdateResourceRequestUpdateCommand {
	s.StorageAddress = &v
	return s
}

func (s *UpdateResourceRequestUpdateCommand) Validate() error {
	return dara.Validate(s)
}
