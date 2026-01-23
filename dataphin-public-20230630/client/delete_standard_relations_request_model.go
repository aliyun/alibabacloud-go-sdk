// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStandardRelationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteCommand(v *DeleteStandardRelationsRequestDeleteCommand) *DeleteStandardRelationsRequest
	GetDeleteCommand() *DeleteStandardRelationsRequestDeleteCommand
	SetOpTenantId(v int64) *DeleteStandardRelationsRequest
	GetOpTenantId() *int64
}

type DeleteStandardRelationsRequest struct {
	// This parameter is required.
	DeleteCommand *DeleteStandardRelationsRequestDeleteCommand `json:"DeleteCommand,omitempty" xml:"DeleteCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s DeleteStandardRelationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteStandardRelationsRequest) GoString() string {
	return s.String()
}

func (s *DeleteStandardRelationsRequest) GetDeleteCommand() *DeleteStandardRelationsRequestDeleteCommand {
	return s.DeleteCommand
}

func (s *DeleteStandardRelationsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteStandardRelationsRequest) SetDeleteCommand(v *DeleteStandardRelationsRequestDeleteCommand) *DeleteStandardRelationsRequest {
	s.DeleteCommand = v
	return s
}

func (s *DeleteStandardRelationsRequest) SetOpTenantId(v int64) *DeleteStandardRelationsRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteStandardRelationsRequest) Validate() error {
	if s.DeleteCommand != nil {
		if err := s.DeleteCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteStandardRelationsRequestDeleteCommand struct {
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
	StandardRefList []*DeleteStandardRelationsRequestDeleteCommandStandardRefList `json:"StandardRefList,omitempty" xml:"StandardRefList,omitempty" type:"Repeated"`
}

func (s DeleteStandardRelationsRequestDeleteCommand) String() string {
	return dara.Prettify(s)
}

func (s DeleteStandardRelationsRequestDeleteCommand) GoString() string {
	return s.String()
}

func (s *DeleteStandardRelationsRequestDeleteCommand) GetRelationType() *string {
	return s.RelationType
}

func (s *DeleteStandardRelationsRequestDeleteCommand) GetStandardId() *int64 {
	return s.StandardId
}

func (s *DeleteStandardRelationsRequestDeleteCommand) GetStandardRefList() []*DeleteStandardRelationsRequestDeleteCommandStandardRefList {
	return s.StandardRefList
}

func (s *DeleteStandardRelationsRequestDeleteCommand) SetRelationType(v string) *DeleteStandardRelationsRequestDeleteCommand {
	s.RelationType = &v
	return s
}

func (s *DeleteStandardRelationsRequestDeleteCommand) SetStandardId(v int64) *DeleteStandardRelationsRequestDeleteCommand {
	s.StandardId = &v
	return s
}

func (s *DeleteStandardRelationsRequestDeleteCommand) SetStandardRefList(v []*DeleteStandardRelationsRequestDeleteCommandStandardRefList) *DeleteStandardRelationsRequestDeleteCommand {
	s.StandardRefList = v
	return s
}

func (s *DeleteStandardRelationsRequestDeleteCommand) Validate() error {
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

type DeleteStandardRelationsRequestDeleteCommandStandardRefList struct {
	// This parameter is required.
	//
	// example:
	//
	// 2
	StandardId *int64 `json:"StandardId,omitempty" xml:"StandardId,omitempty"`
}

func (s DeleteStandardRelationsRequestDeleteCommandStandardRefList) String() string {
	return dara.Prettify(s)
}

func (s DeleteStandardRelationsRequestDeleteCommandStandardRefList) GoString() string {
	return s.String()
}

func (s *DeleteStandardRelationsRequestDeleteCommandStandardRefList) GetStandardId() *int64 {
	return s.StandardId
}

func (s *DeleteStandardRelationsRequestDeleteCommandStandardRefList) SetStandardId(v int64) *DeleteStandardRelationsRequestDeleteCommandStandardRefList {
	s.StandardId = &v
	return s
}

func (s *DeleteStandardRelationsRequestDeleteCommandStandardRefList) Validate() error {
	return dara.Validate(s)
}
