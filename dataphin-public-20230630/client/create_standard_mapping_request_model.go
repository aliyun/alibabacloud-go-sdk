// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStandardMappingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommand(v *CreateStandardMappingRequestCreateCommand) *CreateStandardMappingRequest
	GetCreateCommand() *CreateStandardMappingRequestCreateCommand
	SetOpTenantId(v int64) *CreateStandardMappingRequest
	GetOpTenantId() *int64
}

type CreateStandardMappingRequest struct {
	// The create command.
	//
	// This parameter is required.
	CreateCommand *CreateStandardMappingRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateStandardMappingRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardMappingRequest) GoString() string {
	return s.String()
}

func (s *CreateStandardMappingRequest) GetCreateCommand() *CreateStandardMappingRequestCreateCommand {
	return s.CreateCommand
}

func (s *CreateStandardMappingRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateStandardMappingRequest) SetCreateCommand(v *CreateStandardMappingRequestCreateCommand) *CreateStandardMappingRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateStandardMappingRequest) SetOpTenantId(v int64) *CreateStandardMappingRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateStandardMappingRequest) Validate() error {
	if s.CreateCommand != nil {
		if err := s.CreateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateStandardMappingRequestCreateCommand struct {
	// The list of asset GUIDs.
	//
	// This parameter is required.
	AssetGuidList []*string `json:"AssetGuidList,omitempty" xml:"AssetGuidList,omitempty" type:"Repeated"`
	// The processing policy for invalid mappings when importing mapping relationships. Valid values:
	//
	// - SET_INVALID_TO_VALID: sets invalid mappings to valid mappings.
	//
	// - KEEP_INVALID_AND_SKIP: retains invalid mappings and skips them.
	//
	// Default value: SET_INVALID_TO_VALID.
	//
	// example:
	//
	// SET_INVALID_TO_VALID
	InvalidMappingRelationOperationType *string `json:"InvalidMappingRelationOperationType,omitempty" xml:"InvalidMappingRelationOperationType,omitempty"`
	// The mapping relationship type. Valid values:
	//
	// - VALID: valid mapping.
	//
	// - INVALID: invalid mapping.
	//
	// Default value: VALID.
	//
	// example:
	//
	// VALID
	RelationType *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
	// The standard ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	StandardId *int64 `json:"StandardId,omitempty" xml:"StandardId,omitempty"`
}

func (s CreateStandardMappingRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardMappingRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateStandardMappingRequestCreateCommand) GetAssetGuidList() []*string {
	return s.AssetGuidList
}

func (s *CreateStandardMappingRequestCreateCommand) GetInvalidMappingRelationOperationType() *string {
	return s.InvalidMappingRelationOperationType
}

func (s *CreateStandardMappingRequestCreateCommand) GetRelationType() *string {
	return s.RelationType
}

func (s *CreateStandardMappingRequestCreateCommand) GetStandardId() *int64 {
	return s.StandardId
}

func (s *CreateStandardMappingRequestCreateCommand) SetAssetGuidList(v []*string) *CreateStandardMappingRequestCreateCommand {
	s.AssetGuidList = v
	return s
}

func (s *CreateStandardMappingRequestCreateCommand) SetInvalidMappingRelationOperationType(v string) *CreateStandardMappingRequestCreateCommand {
	s.InvalidMappingRelationOperationType = &v
	return s
}

func (s *CreateStandardMappingRequestCreateCommand) SetRelationType(v string) *CreateStandardMappingRequestCreateCommand {
	s.RelationType = &v
	return s
}

func (s *CreateStandardMappingRequestCreateCommand) SetStandardId(v int64) *CreateStandardMappingRequestCreateCommand {
	s.StandardId = &v
	return s
}

func (s *CreateStandardMappingRequestCreateCommand) Validate() error {
	return dara.Validate(s)
}
