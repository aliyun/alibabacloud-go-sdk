// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateKgRelationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommand(v *BatchCreateKgRelationRequestCreateCommand) *BatchCreateKgRelationRequest
	GetCreateCommand() *BatchCreateKgRelationRequestCreateCommand
	SetOpTenantId(v int64) *BatchCreateKgRelationRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *BatchCreateKgRelationRequest
	GetOpUserId() *string
	SetWorkspaceId(v string) *BatchCreateKgRelationRequest
	GetWorkspaceId() *string
}

type BatchCreateKgRelationRequest struct {
	// The create command.
	//
	// This parameter is required.
	CreateCommand *BatchCreateKgRelationRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
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

func (s BatchCreateKgRelationRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateKgRelationRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateKgRelationRequest) GetCreateCommand() *BatchCreateKgRelationRequestCreateCommand {
	return s.CreateCommand
}

func (s *BatchCreateKgRelationRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *BatchCreateKgRelationRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *BatchCreateKgRelationRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *BatchCreateKgRelationRequest) SetCreateCommand(v *BatchCreateKgRelationRequestCreateCommand) *BatchCreateKgRelationRequest {
	s.CreateCommand = v
	return s
}

func (s *BatchCreateKgRelationRequest) SetOpTenantId(v int64) *BatchCreateKgRelationRequest {
	s.OpTenantId = &v
	return s
}

func (s *BatchCreateKgRelationRequest) SetOpUserId(v string) *BatchCreateKgRelationRequest {
	s.OpUserId = &v
	return s
}

func (s *BatchCreateKgRelationRequest) SetWorkspaceId(v string) *BatchCreateKgRelationRequest {
	s.WorkspaceId = &v
	return s
}

func (s *BatchCreateKgRelationRequest) Validate() error {
	if s.CreateCommand != nil {
		if err := s.CreateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchCreateKgRelationRequestCreateCommand struct {
	// The list of relationships.
	//
	// This parameter is required.
	RelationList []*BatchCreateKgRelationRequestCreateCommandRelationList `json:"RelationList,omitempty" xml:"RelationList,omitempty" type:"Repeated"`
}

func (s BatchCreateKgRelationRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateKgRelationRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *BatchCreateKgRelationRequestCreateCommand) GetRelationList() []*BatchCreateKgRelationRequestCreateCommandRelationList {
	return s.RelationList
}

func (s *BatchCreateKgRelationRequestCreateCommand) SetRelationList(v []*BatchCreateKgRelationRequestCreateCommandRelationList) *BatchCreateKgRelationRequestCreateCommand {
	s.RelationList = v
	return s
}

func (s *BatchCreateKgRelationRequestCreateCommand) Validate() error {
	if s.RelationList != nil {
		for _, item := range s.RelationList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchCreateKgRelationRequestCreateCommandRelationList struct {
	// The list of relationship record properties.
	PropertyList []*BatchCreateKgRelationRequestCreateCommandRelationListPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Repeated"`
	// The relationship type code.
	//
	// This parameter is required.
	//
	// example:
	//
	// Company
	RelationType *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
	// The ID of the source entity record.
	//
	// This parameter is required.
	//
	// example:
	//
	// e1d4559a4db044158305e2d89bccf81f
	SourceEntityId *string `json:"SourceEntityId,omitempty" xml:"SourceEntityId,omitempty"`
	// The ID of the target entity record.
	//
	// This parameter is required.
	//
	// example:
	//
	// e1d4559a4db044158305e2d89bccf82f
	TargetEntityId *string `json:"TargetEntityId,omitempty" xml:"TargetEntityId,omitempty"`
}

func (s BatchCreateKgRelationRequestCreateCommandRelationList) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateKgRelationRequestCreateCommandRelationList) GoString() string {
	return s.String()
}

func (s *BatchCreateKgRelationRequestCreateCommandRelationList) GetPropertyList() []*BatchCreateKgRelationRequestCreateCommandRelationListPropertyList {
	return s.PropertyList
}

func (s *BatchCreateKgRelationRequestCreateCommandRelationList) GetRelationType() *string {
	return s.RelationType
}

func (s *BatchCreateKgRelationRequestCreateCommandRelationList) GetSourceEntityId() *string {
	return s.SourceEntityId
}

func (s *BatchCreateKgRelationRequestCreateCommandRelationList) GetTargetEntityId() *string {
	return s.TargetEntityId
}

func (s *BatchCreateKgRelationRequestCreateCommandRelationList) SetPropertyList(v []*BatchCreateKgRelationRequestCreateCommandRelationListPropertyList) *BatchCreateKgRelationRequestCreateCommandRelationList {
	s.PropertyList = v
	return s
}

func (s *BatchCreateKgRelationRequestCreateCommandRelationList) SetRelationType(v string) *BatchCreateKgRelationRequestCreateCommandRelationList {
	s.RelationType = &v
	return s
}

func (s *BatchCreateKgRelationRequestCreateCommandRelationList) SetSourceEntityId(v string) *BatchCreateKgRelationRequestCreateCommandRelationList {
	s.SourceEntityId = &v
	return s
}

func (s *BatchCreateKgRelationRequestCreateCommandRelationList) SetTargetEntityId(v string) *BatchCreateKgRelationRequestCreateCommandRelationList {
	s.TargetEntityId = &v
	return s
}

func (s *BatchCreateKgRelationRequestCreateCommandRelationList) Validate() error {
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

type BatchCreateKgRelationRequestCreateCommandRelationListPropertyList struct {
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

func (s BatchCreateKgRelationRequestCreateCommandRelationListPropertyList) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateKgRelationRequestCreateCommandRelationListPropertyList) GoString() string {
	return s.String()
}

func (s *BatchCreateKgRelationRequestCreateCommandRelationListPropertyList) GetCode() *string {
	return s.Code
}

func (s *BatchCreateKgRelationRequestCreateCommandRelationListPropertyList) GetValue() *string {
	return s.Value
}

func (s *BatchCreateKgRelationRequestCreateCommandRelationListPropertyList) SetCode(v string) *BatchCreateKgRelationRequestCreateCommandRelationListPropertyList {
	s.Code = &v
	return s
}

func (s *BatchCreateKgRelationRequestCreateCommandRelationListPropertyList) SetValue(v string) *BatchCreateKgRelationRequestCreateCommandRelationListPropertyList {
	s.Value = &v
	return s
}

func (s *BatchCreateKgRelationRequestCreateCommandRelationListPropertyList) Validate() error {
	return dara.Validate(s)
}
