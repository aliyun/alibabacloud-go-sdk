// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateKgEntityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommand(v *BatchCreateKgEntityRequestCreateCommand) *BatchCreateKgEntityRequest
	GetCreateCommand() *BatchCreateKgEntityRequestCreateCommand
	SetOpTenantId(v int64) *BatchCreateKgEntityRequest
	GetOpTenantId() *int64
	SetWorkspaceId(v string) *BatchCreateKgEntityRequest
	GetWorkspaceId() *string
}

type BatchCreateKgEntityRequest struct {
	// The create instruction.
	//
	// This parameter is required.
	CreateCommand *BatchCreateKgEntityRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
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

func (s BatchCreateKgEntityRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateKgEntityRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateKgEntityRequest) GetCreateCommand() *BatchCreateKgEntityRequestCreateCommand {
	return s.CreateCommand
}

func (s *BatchCreateKgEntityRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *BatchCreateKgEntityRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *BatchCreateKgEntityRequest) SetCreateCommand(v *BatchCreateKgEntityRequestCreateCommand) *BatchCreateKgEntityRequest {
	s.CreateCommand = v
	return s
}

func (s *BatchCreateKgEntityRequest) SetOpTenantId(v int64) *BatchCreateKgEntityRequest {
	s.OpTenantId = &v
	return s
}

func (s *BatchCreateKgEntityRequest) SetWorkspaceId(v string) *BatchCreateKgEntityRequest {
	s.WorkspaceId = &v
	return s
}

func (s *BatchCreateKgEntityRequest) Validate() error {
	if s.CreateCommand != nil {
		if err := s.CreateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchCreateKgEntityRequestCreateCommand struct {
	// The list of entity records.
	//
	// This parameter is required.
	EntityList []*BatchCreateKgEntityRequestCreateCommandEntityList `json:"EntityList,omitempty" xml:"EntityList,omitempty" type:"Repeated"`
}

func (s BatchCreateKgEntityRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateKgEntityRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *BatchCreateKgEntityRequestCreateCommand) GetEntityList() []*BatchCreateKgEntityRequestCreateCommandEntityList {
	return s.EntityList
}

func (s *BatchCreateKgEntityRequestCreateCommand) SetEntityList(v []*BatchCreateKgEntityRequestCreateCommandEntityList) *BatchCreateKgEntityRequestCreateCommand {
	s.EntityList = v
	return s
}

func (s *BatchCreateKgEntityRequestCreateCommand) Validate() error {
	if s.EntityList != nil {
		for _, item := range s.EntityList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchCreateKgEntityRequestCreateCommandEntityList struct {
	// The entity type code.
	//
	// This parameter is required.
	//
	// example:
	//
	// Company
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The list of entity record properties.
	//
	// This parameter is required.
	PropertyList []*BatchCreateKgEntityRequestCreateCommandEntityListPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Repeated"`
}

func (s BatchCreateKgEntityRequestCreateCommandEntityList) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateKgEntityRequestCreateCommandEntityList) GoString() string {
	return s.String()
}

func (s *BatchCreateKgEntityRequestCreateCommandEntityList) GetEntityType() *string {
	return s.EntityType
}

func (s *BatchCreateKgEntityRequestCreateCommandEntityList) GetPropertyList() []*BatchCreateKgEntityRequestCreateCommandEntityListPropertyList {
	return s.PropertyList
}

func (s *BatchCreateKgEntityRequestCreateCommandEntityList) SetEntityType(v string) *BatchCreateKgEntityRequestCreateCommandEntityList {
	s.EntityType = &v
	return s
}

func (s *BatchCreateKgEntityRequestCreateCommandEntityList) SetPropertyList(v []*BatchCreateKgEntityRequestCreateCommandEntityListPropertyList) *BatchCreateKgEntityRequestCreateCommandEntityList {
	s.PropertyList = v
	return s
}

func (s *BatchCreateKgEntityRequestCreateCommandEntityList) Validate() error {
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

type BatchCreateKgEntityRequestCreateCommandEntityListPropertyList struct {
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

func (s BatchCreateKgEntityRequestCreateCommandEntityListPropertyList) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateKgEntityRequestCreateCommandEntityListPropertyList) GoString() string {
	return s.String()
}

func (s *BatchCreateKgEntityRequestCreateCommandEntityListPropertyList) GetCode() *string {
	return s.Code
}

func (s *BatchCreateKgEntityRequestCreateCommandEntityListPropertyList) GetValue() *string {
	return s.Value
}

func (s *BatchCreateKgEntityRequestCreateCommandEntityListPropertyList) SetCode(v string) *BatchCreateKgEntityRequestCreateCommandEntityListPropertyList {
	s.Code = &v
	return s
}

func (s *BatchCreateKgEntityRequestCreateCommandEntityListPropertyList) SetValue(v string) *BatchCreateKgEntityRequestCreateCommandEntityListPropertyList {
	s.Value = &v
	return s
}

func (s *BatchCreateKgEntityRequestCreateCommandEntityListPropertyList) Validate() error {
	return dara.Validate(s)
}
