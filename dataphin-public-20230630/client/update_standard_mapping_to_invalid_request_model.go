// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStandardMappingToInvalidRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateStandardMappingToInvalidRequest
	GetOpTenantId() *int64
	SetUpdateCommand(v *UpdateStandardMappingToInvalidRequestUpdateCommand) *UpdateStandardMappingToInvalidRequest
	GetUpdateCommand() *UpdateStandardMappingToInvalidRequestUpdateCommand
}

type UpdateStandardMappingToInvalidRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommand *UpdateStandardMappingToInvalidRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateStandardMappingToInvalidRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardMappingToInvalidRequest) GoString() string {
	return s.String()
}

func (s *UpdateStandardMappingToInvalidRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateStandardMappingToInvalidRequest) GetUpdateCommand() *UpdateStandardMappingToInvalidRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateStandardMappingToInvalidRequest) SetOpTenantId(v int64) *UpdateStandardMappingToInvalidRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateStandardMappingToInvalidRequest) SetUpdateCommand(v *UpdateStandardMappingToInvalidRequestUpdateCommand) *UpdateStandardMappingToInvalidRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateStandardMappingToInvalidRequest) Validate() error {
	if s.UpdateCommand != nil {
		if err := s.UpdateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateStandardMappingToInvalidRequestUpdateCommand struct {
	BelongGuidList []*string `json:"BelongGuidList,omitempty" xml:"BelongGuidList,omitempty" type:"Repeated"`
	GuidList       []*string `json:"GuidList,omitempty" xml:"GuidList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	StandardId *int64 `json:"StandardId,omitempty" xml:"StandardId,omitempty"`
}

func (s UpdateStandardMappingToInvalidRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardMappingToInvalidRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateStandardMappingToInvalidRequestUpdateCommand) GetBelongGuidList() []*string {
	return s.BelongGuidList
}

func (s *UpdateStandardMappingToInvalidRequestUpdateCommand) GetGuidList() []*string {
	return s.GuidList
}

func (s *UpdateStandardMappingToInvalidRequestUpdateCommand) GetStandardId() *int64 {
	return s.StandardId
}

func (s *UpdateStandardMappingToInvalidRequestUpdateCommand) SetBelongGuidList(v []*string) *UpdateStandardMappingToInvalidRequestUpdateCommand {
	s.BelongGuidList = v
	return s
}

func (s *UpdateStandardMappingToInvalidRequestUpdateCommand) SetGuidList(v []*string) *UpdateStandardMappingToInvalidRequestUpdateCommand {
	s.GuidList = v
	return s
}

func (s *UpdateStandardMappingToInvalidRequestUpdateCommand) SetStandardId(v int64) *UpdateStandardMappingToInvalidRequestUpdateCommand {
	s.StandardId = &v
	return s
}

func (s *UpdateStandardMappingToInvalidRequestUpdateCommand) Validate() error {
	return dara.Validate(s)
}
