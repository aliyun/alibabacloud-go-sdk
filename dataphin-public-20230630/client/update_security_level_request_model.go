// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecurityLevelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateSecurityLevelRequest
	GetOpTenantId() *int64
	SetUpdateCommand(v *UpdateSecurityLevelRequestUpdateCommand) *UpdateSecurityLevelRequest
	GetUpdateCommand() *UpdateSecurityLevelRequestUpdateCommand
}

type UpdateSecurityLevelRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommand *UpdateSecurityLevelRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateSecurityLevelRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecurityLevelRequest) GoString() string {
	return s.String()
}

func (s *UpdateSecurityLevelRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateSecurityLevelRequest) GetUpdateCommand() *UpdateSecurityLevelRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateSecurityLevelRequest) SetOpTenantId(v int64) *UpdateSecurityLevelRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateSecurityLevelRequest) SetUpdateCommand(v *UpdateSecurityLevelRequestUpdateCommand) *UpdateSecurityLevelRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateSecurityLevelRequest) Validate() error {
	if s.UpdateCommand != nil {
		if err := s.UpdateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateSecurityLevelRequestUpdateCommand struct {
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateSecurityLevelRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecurityLevelRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateSecurityLevelRequestUpdateCommand) GetDescription() *string {
	return s.Description
}

func (s *UpdateSecurityLevelRequestUpdateCommand) GetIndex() *int64 {
	return s.Index
}

func (s *UpdateSecurityLevelRequestUpdateCommand) GetName() *string {
	return s.Name
}

func (s *UpdateSecurityLevelRequestUpdateCommand) SetDescription(v string) *UpdateSecurityLevelRequestUpdateCommand {
	s.Description = &v
	return s
}

func (s *UpdateSecurityLevelRequestUpdateCommand) SetIndex(v int64) *UpdateSecurityLevelRequestUpdateCommand {
	s.Index = &v
	return s
}

func (s *UpdateSecurityLevelRequestUpdateCommand) SetName(v string) *UpdateSecurityLevelRequestUpdateCommand {
	s.Name = &v
	return s
}

func (s *UpdateSecurityLevelRequestUpdateCommand) Validate() error {
	return dara.Validate(s)
}
