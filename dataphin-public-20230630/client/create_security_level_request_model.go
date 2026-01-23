// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecurityLevelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommand(v *CreateSecurityLevelRequestCreateCommand) *CreateSecurityLevelRequest
	GetCreateCommand() *CreateSecurityLevelRequestCreateCommand
	SetOpTenantId(v int64) *CreateSecurityLevelRequest
	GetOpTenantId() *int64
}

type CreateSecurityLevelRequest struct {
	// This parameter is required.
	CreateCommand *CreateSecurityLevelRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateSecurityLevelRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityLevelRequest) GoString() string {
	return s.String()
}

func (s *CreateSecurityLevelRequest) GetCreateCommand() *CreateSecurityLevelRequestCreateCommand {
	return s.CreateCommand
}

func (s *CreateSecurityLevelRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateSecurityLevelRequest) SetCreateCommand(v *CreateSecurityLevelRequestCreateCommand) *CreateSecurityLevelRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateSecurityLevelRequest) SetOpTenantId(v int64) *CreateSecurityLevelRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateSecurityLevelRequest) Validate() error {
	if s.CreateCommand != nil {
		if err := s.CreateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSecurityLevelRequestCreateCommand struct {
	// example:
	//
	// test
	Abbreviation *string `json:"Abbreviation,omitempty" xml:"Abbreviation,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateSecurityLevelRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityLevelRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateSecurityLevelRequestCreateCommand) GetAbbreviation() *string {
	return s.Abbreviation
}

func (s *CreateSecurityLevelRequestCreateCommand) GetDescription() *string {
	return s.Description
}

func (s *CreateSecurityLevelRequestCreateCommand) GetIndex() *int64 {
	return s.Index
}

func (s *CreateSecurityLevelRequestCreateCommand) GetName() *string {
	return s.Name
}

func (s *CreateSecurityLevelRequestCreateCommand) SetAbbreviation(v string) *CreateSecurityLevelRequestCreateCommand {
	s.Abbreviation = &v
	return s
}

func (s *CreateSecurityLevelRequestCreateCommand) SetDescription(v string) *CreateSecurityLevelRequestCreateCommand {
	s.Description = &v
	return s
}

func (s *CreateSecurityLevelRequestCreateCommand) SetIndex(v int64) *CreateSecurityLevelRequestCreateCommand {
	s.Index = &v
	return s
}

func (s *CreateSecurityLevelRequestCreateCommand) SetName(v string) *CreateSecurityLevelRequestCreateCommand {
	s.Name = &v
	return s
}

func (s *CreateSecurityLevelRequestCreateCommand) Validate() error {
	return dara.Validate(s)
}
