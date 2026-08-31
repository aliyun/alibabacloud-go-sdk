// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKgRelationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateKgRelationRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *UpdateKgRelationRequest
	GetOpUserId() *string
	SetUpdateCommand(v *UpdateKgRelationRequestUpdateCommand) *UpdateKgRelationRequest
	GetUpdateCommand() *UpdateKgRelationRequestUpdateCommand
	SetWorkspaceId(v string) *UpdateKgRelationRequest
	GetWorkspaceId() *string
}

type UpdateKgRelationRequest struct {
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
	UpdateCommand *UpdateKgRelationRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// f1d4559a4db044158305e2d89bccf81f
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateKgRelationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKgRelationRequest) GoString() string {
	return s.String()
}

func (s *UpdateKgRelationRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateKgRelationRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *UpdateKgRelationRequest) GetUpdateCommand() *UpdateKgRelationRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateKgRelationRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateKgRelationRequest) SetOpTenantId(v int64) *UpdateKgRelationRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateKgRelationRequest) SetOpUserId(v string) *UpdateKgRelationRequest {
	s.OpUserId = &v
	return s
}

func (s *UpdateKgRelationRequest) SetUpdateCommand(v *UpdateKgRelationRequestUpdateCommand) *UpdateKgRelationRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateKgRelationRequest) SetWorkspaceId(v string) *UpdateKgRelationRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateKgRelationRequest) Validate() error {
	if s.UpdateCommand != nil {
		if err := s.UpdateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateKgRelationRequestUpdateCommand struct {
	// The list of relationship record properties.
	PropertyList []*UpdateKgRelationRequestUpdateCommandPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Repeated"`
	// The relationship record ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// abc-xxx
	RelationId *string `json:"RelationId,omitempty" xml:"RelationId,omitempty"`
	// The relationship type code.
	//
	// This parameter is required.
	//
	// example:
	//
	// Company
	RelationType *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
}

func (s UpdateKgRelationRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateKgRelationRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateKgRelationRequestUpdateCommand) GetPropertyList() []*UpdateKgRelationRequestUpdateCommandPropertyList {
	return s.PropertyList
}

func (s *UpdateKgRelationRequestUpdateCommand) GetRelationId() *string {
	return s.RelationId
}

func (s *UpdateKgRelationRequestUpdateCommand) GetRelationType() *string {
	return s.RelationType
}

func (s *UpdateKgRelationRequestUpdateCommand) SetPropertyList(v []*UpdateKgRelationRequestUpdateCommandPropertyList) *UpdateKgRelationRequestUpdateCommand {
	s.PropertyList = v
	return s
}

func (s *UpdateKgRelationRequestUpdateCommand) SetRelationId(v string) *UpdateKgRelationRequestUpdateCommand {
	s.RelationId = &v
	return s
}

func (s *UpdateKgRelationRequestUpdateCommand) SetRelationType(v string) *UpdateKgRelationRequestUpdateCommand {
	s.RelationType = &v
	return s
}

func (s *UpdateKgRelationRequestUpdateCommand) Validate() error {
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

type UpdateKgRelationRequestUpdateCommandPropertyList struct {
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

func (s UpdateKgRelationRequestUpdateCommandPropertyList) String() string {
	return dara.Prettify(s)
}

func (s UpdateKgRelationRequestUpdateCommandPropertyList) GoString() string {
	return s.String()
}

func (s *UpdateKgRelationRequestUpdateCommandPropertyList) GetCode() *string {
	return s.Code
}

func (s *UpdateKgRelationRequestUpdateCommandPropertyList) GetValue() *string {
	return s.Value
}

func (s *UpdateKgRelationRequestUpdateCommandPropertyList) SetCode(v string) *UpdateKgRelationRequestUpdateCommandPropertyList {
	s.Code = &v
	return s
}

func (s *UpdateKgRelationRequestUpdateCommandPropertyList) SetValue(v string) *UpdateKgRelationRequestUpdateCommandPropertyList {
	s.Value = &v
	return s
}

func (s *UpdateKgRelationRequestUpdateCommandPropertyList) Validate() error {
	return dara.Validate(s)
}
