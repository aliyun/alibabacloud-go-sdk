// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKgEntityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateKgEntityRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *UpdateKgEntityRequest
	GetOpUserId() *string
	SetUpdateCommand(v *UpdateKgEntityRequestUpdateCommand) *UpdateKgEntityRequest
	GetUpdateCommand() *UpdateKgEntityRequestUpdateCommand
	SetWorkspaceId(v string) *UpdateKgEntityRequest
	GetWorkspaceId() *string
}

type UpdateKgEntityRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The update command.
	//
	// This parameter is required.
	UpdateCommand *UpdateKgEntityRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// f1d4559a4db044158305e2d89bccf81f
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateKgEntityRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKgEntityRequest) GoString() string {
	return s.String()
}

func (s *UpdateKgEntityRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateKgEntityRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *UpdateKgEntityRequest) GetUpdateCommand() *UpdateKgEntityRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateKgEntityRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateKgEntityRequest) SetOpTenantId(v int64) *UpdateKgEntityRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateKgEntityRequest) SetOpUserId(v string) *UpdateKgEntityRequest {
	s.OpUserId = &v
	return s
}

func (s *UpdateKgEntityRequest) SetUpdateCommand(v *UpdateKgEntityRequestUpdateCommand) *UpdateKgEntityRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateKgEntityRequest) SetWorkspaceId(v string) *UpdateKgEntityRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateKgEntityRequest) Validate() error {
	if s.UpdateCommand != nil {
		if err := s.UpdateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateKgEntityRequestUpdateCommand struct {
	// The entity record ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// abc-xxx
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
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
	PropertyList []*UpdateKgEntityRequestUpdateCommandPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Repeated"`
}

func (s UpdateKgEntityRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateKgEntityRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateKgEntityRequestUpdateCommand) GetEntityId() *string {
	return s.EntityId
}

func (s *UpdateKgEntityRequestUpdateCommand) GetEntityType() *string {
	return s.EntityType
}

func (s *UpdateKgEntityRequestUpdateCommand) GetPropertyList() []*UpdateKgEntityRequestUpdateCommandPropertyList {
	return s.PropertyList
}

func (s *UpdateKgEntityRequestUpdateCommand) SetEntityId(v string) *UpdateKgEntityRequestUpdateCommand {
	s.EntityId = &v
	return s
}

func (s *UpdateKgEntityRequestUpdateCommand) SetEntityType(v string) *UpdateKgEntityRequestUpdateCommand {
	s.EntityType = &v
	return s
}

func (s *UpdateKgEntityRequestUpdateCommand) SetPropertyList(v []*UpdateKgEntityRequestUpdateCommandPropertyList) *UpdateKgEntityRequestUpdateCommand {
	s.PropertyList = v
	return s
}

func (s *UpdateKgEntityRequestUpdateCommand) Validate() error {
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

type UpdateKgEntityRequestUpdateCommandPropertyList struct {
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

func (s UpdateKgEntityRequestUpdateCommandPropertyList) String() string {
	return dara.Prettify(s)
}

func (s UpdateKgEntityRequestUpdateCommandPropertyList) GoString() string {
	return s.String()
}

func (s *UpdateKgEntityRequestUpdateCommandPropertyList) GetCode() *string {
	return s.Code
}

func (s *UpdateKgEntityRequestUpdateCommandPropertyList) GetValue() *string {
	return s.Value
}

func (s *UpdateKgEntityRequestUpdateCommandPropertyList) SetCode(v string) *UpdateKgEntityRequestUpdateCommandPropertyList {
	s.Code = &v
	return s
}

func (s *UpdateKgEntityRequestUpdateCommandPropertyList) SetValue(v string) *UpdateKgEntityRequestUpdateCommandPropertyList {
	s.Value = &v
	return s
}

func (s *UpdateKgEntityRequestUpdateCommandPropertyList) Validate() error {
	return dara.Validate(s)
}
