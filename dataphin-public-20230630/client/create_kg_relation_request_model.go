// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKgRelationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommand(v *CreateKgRelationRequestCreateCommand) *CreateKgRelationRequest
	GetCreateCommand() *CreateKgRelationRequestCreateCommand
	SetOpTenantId(v int64) *CreateKgRelationRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *CreateKgRelationRequest
	GetOpUserId() *string
	SetWorkspaceId(v string) *CreateKgRelationRequest
	GetWorkspaceId() *string
}

type CreateKgRelationRequest struct {
	// The create command.
	//
	// This parameter is required.
	CreateCommand *CreateKgRelationRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
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
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// f1d4559a4db044158305e2d89bccf81f
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateKgRelationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKgRelationRequest) GoString() string {
	return s.String()
}

func (s *CreateKgRelationRequest) GetCreateCommand() *CreateKgRelationRequestCreateCommand {
	return s.CreateCommand
}

func (s *CreateKgRelationRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateKgRelationRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *CreateKgRelationRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateKgRelationRequest) SetCreateCommand(v *CreateKgRelationRequestCreateCommand) *CreateKgRelationRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateKgRelationRequest) SetOpTenantId(v int64) *CreateKgRelationRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateKgRelationRequest) SetOpUserId(v string) *CreateKgRelationRequest {
	s.OpUserId = &v
	return s
}

func (s *CreateKgRelationRequest) SetWorkspaceId(v string) *CreateKgRelationRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateKgRelationRequest) Validate() error {
	if s.CreateCommand != nil {
		if err := s.CreateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateKgRelationRequestCreateCommand struct {
	// The relationship record property list.
	PropertyList []*CreateKgRelationRequestCreateCommandPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Repeated"`
	// The relationship type code.
	//
	// This parameter is required.
	//
	// example:
	//
	// Company
	RelationType *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
	// The source entity record ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// e1d4559a4db044158305e2d89bccf81f
	SourceEntityId *string `json:"SourceEntityId,omitempty" xml:"SourceEntityId,omitempty"`
	// The target entity record ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// e1d4559a4db044158305e2d89bccf82f
	TargetEntityId *string `json:"TargetEntityId,omitempty" xml:"TargetEntityId,omitempty"`
}

func (s CreateKgRelationRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateKgRelationRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateKgRelationRequestCreateCommand) GetPropertyList() []*CreateKgRelationRequestCreateCommandPropertyList {
	return s.PropertyList
}

func (s *CreateKgRelationRequestCreateCommand) GetRelationType() *string {
	return s.RelationType
}

func (s *CreateKgRelationRequestCreateCommand) GetSourceEntityId() *string {
	return s.SourceEntityId
}

func (s *CreateKgRelationRequestCreateCommand) GetTargetEntityId() *string {
	return s.TargetEntityId
}

func (s *CreateKgRelationRequestCreateCommand) SetPropertyList(v []*CreateKgRelationRequestCreateCommandPropertyList) *CreateKgRelationRequestCreateCommand {
	s.PropertyList = v
	return s
}

func (s *CreateKgRelationRequestCreateCommand) SetRelationType(v string) *CreateKgRelationRequestCreateCommand {
	s.RelationType = &v
	return s
}

func (s *CreateKgRelationRequestCreateCommand) SetSourceEntityId(v string) *CreateKgRelationRequestCreateCommand {
	s.SourceEntityId = &v
	return s
}

func (s *CreateKgRelationRequestCreateCommand) SetTargetEntityId(v string) *CreateKgRelationRequestCreateCommand {
	s.TargetEntityId = &v
	return s
}

func (s *CreateKgRelationRequestCreateCommand) Validate() error {
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

type CreateKgRelationRequestCreateCommandPropertyList struct {
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

func (s CreateKgRelationRequestCreateCommandPropertyList) String() string {
	return dara.Prettify(s)
}

func (s CreateKgRelationRequestCreateCommandPropertyList) GoString() string {
	return s.String()
}

func (s *CreateKgRelationRequestCreateCommandPropertyList) GetCode() *string {
	return s.Code
}

func (s *CreateKgRelationRequestCreateCommandPropertyList) GetValue() *string {
	return s.Value
}

func (s *CreateKgRelationRequestCreateCommandPropertyList) SetCode(v string) *CreateKgRelationRequestCreateCommandPropertyList {
	s.Code = &v
	return s
}

func (s *CreateKgRelationRequestCreateCommandPropertyList) SetValue(v string) *CreateKgRelationRequestCreateCommandPropertyList {
	s.Value = &v
	return s
}

func (s *CreateKgRelationRequestCreateCommandPropertyList) Validate() error {
	return dara.Validate(s)
}
