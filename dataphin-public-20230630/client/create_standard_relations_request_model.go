// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStandardRelationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommand(v *CreateStandardRelationsRequestCreateCommand) *CreateStandardRelationsRequest
	GetCreateCommand() *CreateStandardRelationsRequestCreateCommand
	SetOpTenantId(v int64) *CreateStandardRelationsRequest
	GetOpTenantId() *int64
}

type CreateStandardRelationsRequest struct {
	// This parameter is required.
	CreateCommand *CreateStandardRelationsRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateStandardRelationsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardRelationsRequest) GoString() string {
	return s.String()
}

func (s *CreateStandardRelationsRequest) GetCreateCommand() *CreateStandardRelationsRequestCreateCommand {
	return s.CreateCommand
}

func (s *CreateStandardRelationsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateStandardRelationsRequest) SetCreateCommand(v *CreateStandardRelationsRequestCreateCommand) *CreateStandardRelationsRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateStandardRelationsRequest) SetOpTenantId(v int64) *CreateStandardRelationsRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateStandardRelationsRequest) Validate() error {
	if s.CreateCommand != nil {
		if err := s.CreateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateStandardRelationsRequestCreateCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// RELATIVE
	RelationType *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	StandardId *int64 `json:"StandardId,omitempty" xml:"StandardId,omitempty"`
	// This parameter is required.
	StandardRefList []*CreateStandardRelationsRequestCreateCommandStandardRefList `json:"StandardRefList,omitempty" xml:"StandardRefList,omitempty" type:"Repeated"`
}

func (s CreateStandardRelationsRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardRelationsRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateStandardRelationsRequestCreateCommand) GetRelationType() *string {
	return s.RelationType
}

func (s *CreateStandardRelationsRequestCreateCommand) GetStandardId() *int64 {
	return s.StandardId
}

func (s *CreateStandardRelationsRequestCreateCommand) GetStandardRefList() []*CreateStandardRelationsRequestCreateCommandStandardRefList {
	return s.StandardRefList
}

func (s *CreateStandardRelationsRequestCreateCommand) SetRelationType(v string) *CreateStandardRelationsRequestCreateCommand {
	s.RelationType = &v
	return s
}

func (s *CreateStandardRelationsRequestCreateCommand) SetStandardId(v int64) *CreateStandardRelationsRequestCreateCommand {
	s.StandardId = &v
	return s
}

func (s *CreateStandardRelationsRequestCreateCommand) SetStandardRefList(v []*CreateStandardRelationsRequestCreateCommandStandardRefList) *CreateStandardRelationsRequestCreateCommand {
	s.StandardRefList = v
	return s
}

func (s *CreateStandardRelationsRequestCreateCommand) Validate() error {
	if s.StandardRefList != nil {
		for _, item := range s.StandardRefList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateStandardRelationsRequestCreateCommandStandardRefList struct {
	// This parameter is required.
	//
	// example:
	//
	// 2
	StandardId *int64 `json:"StandardId,omitempty" xml:"StandardId,omitempty"`
}

func (s CreateStandardRelationsRequestCreateCommandStandardRefList) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardRelationsRequestCreateCommandStandardRefList) GoString() string {
	return s.String()
}

func (s *CreateStandardRelationsRequestCreateCommandStandardRefList) GetStandardId() *int64 {
	return s.StandardId
}

func (s *CreateStandardRelationsRequestCreateCommandStandardRefList) SetStandardId(v int64) *CreateStandardRelationsRequestCreateCommandStandardRefList {
	s.StandardId = &v
	return s
}

func (s *CreateStandardRelationsRequestCreateCommandStandardRefList) Validate() error {
	return dara.Validate(s)
}
