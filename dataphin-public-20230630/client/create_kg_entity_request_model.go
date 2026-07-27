// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKgEntityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommand(v *CreateKgEntityRequestCreateCommand) *CreateKgEntityRequest
	GetCreateCommand() *CreateKgEntityRequestCreateCommand
	SetOpTenantId(v int64) *CreateKgEntityRequest
	GetOpTenantId() *int64
	SetWorkspaceId(v string) *CreateKgEntityRequest
	GetWorkspaceId() *string
}

type CreateKgEntityRequest struct {
	// The create command.
	//
	// This parameter is required.
	CreateCommand *CreateKgEntityRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
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

func (s CreateKgEntityRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKgEntityRequest) GoString() string {
	return s.String()
}

func (s *CreateKgEntityRequest) GetCreateCommand() *CreateKgEntityRequestCreateCommand {
	return s.CreateCommand
}

func (s *CreateKgEntityRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateKgEntityRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateKgEntityRequest) SetCreateCommand(v *CreateKgEntityRequestCreateCommand) *CreateKgEntityRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateKgEntityRequest) SetOpTenantId(v int64) *CreateKgEntityRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateKgEntityRequest) SetWorkspaceId(v string) *CreateKgEntityRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateKgEntityRequest) Validate() error {
	if s.CreateCommand != nil {
		if err := s.CreateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateKgEntityRequestCreateCommand struct {
	// The entity type code.
	//
	// This parameter is required.
	//
	// example:
	//
	// Company
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The entity record property list.
	//
	// This parameter is required.
	PropertyList []*CreateKgEntityRequestCreateCommandPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Repeated"`
}

func (s CreateKgEntityRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateKgEntityRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateKgEntityRequestCreateCommand) GetEntityType() *string {
	return s.EntityType
}

func (s *CreateKgEntityRequestCreateCommand) GetPropertyList() []*CreateKgEntityRequestCreateCommandPropertyList {
	return s.PropertyList
}

func (s *CreateKgEntityRequestCreateCommand) SetEntityType(v string) *CreateKgEntityRequestCreateCommand {
	s.EntityType = &v
	return s
}

func (s *CreateKgEntityRequestCreateCommand) SetPropertyList(v []*CreateKgEntityRequestCreateCommandPropertyList) *CreateKgEntityRequestCreateCommand {
	s.PropertyList = v
	return s
}

func (s *CreateKgEntityRequestCreateCommand) Validate() error {
	if s.PropertyList != nil {
		for _, item := range s.PropertyList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateKgEntityRequestCreateCommandPropertyList struct {
	// The property code.
	//
	// This parameter is required.
	//
	// example:
	//
	// company_name
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The property value.
	//
	// This parameter is required.
	//
	// example:
	//
	// Alibaba
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateKgEntityRequestCreateCommandPropertyList) String() string {
	return dara.Prettify(s)
}

func (s CreateKgEntityRequestCreateCommandPropertyList) GoString() string {
	return s.String()
}

func (s *CreateKgEntityRequestCreateCommandPropertyList) GetCode() *string {
	return s.Code
}

func (s *CreateKgEntityRequestCreateCommandPropertyList) GetValue() *string {
	return s.Value
}

func (s *CreateKgEntityRequestCreateCommandPropertyList) SetCode(v string) *CreateKgEntityRequestCreateCommandPropertyList {
	s.Code = &v
	return s
}

func (s *CreateKgEntityRequestCreateCommandPropertyList) SetValue(v string) *CreateKgEntityRequestCreateCommandPropertyList {
	s.Value = &v
	return s
}

func (s *CreateKgEntityRequestCreateCommandPropertyList) Validate() error {
	return dara.Validate(s)
}
