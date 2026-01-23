// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStandardWordRootRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommand(v *CreateStandardWordRootRequestCreateCommand) *CreateStandardWordRootRequest
	GetCreateCommand() *CreateStandardWordRootRequestCreateCommand
	SetOpTenantId(v int64) *CreateStandardWordRootRequest
	GetOpTenantId() *int64
}

type CreateStandardWordRootRequest struct {
	// This parameter is required.
	CreateCommand *CreateStandardWordRootRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateStandardWordRootRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardWordRootRequest) GoString() string {
	return s.String()
}

func (s *CreateStandardWordRootRequest) GetCreateCommand() *CreateStandardWordRootRequestCreateCommand {
	return s.CreateCommand
}

func (s *CreateStandardWordRootRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateStandardWordRootRequest) SetCreateCommand(v *CreateStandardWordRootRequestCreateCommand) *CreateStandardWordRootRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateStandardWordRootRequest) SetOpTenantId(v int64) *CreateStandardWordRootRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateStandardWordRootRequest) Validate() error {
	if s.CreateCommand != nil {
		if err := s.CreateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateStandardWordRootRequestCreateCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// avg
	Abbreviation *string `json:"Abbreviation,omitempty" xml:"Abbreviation,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// average
	FullName *string `json:"FullName,omitempty" xml:"FullName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 平均值
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateStandardWordRootRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardWordRootRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateStandardWordRootRequestCreateCommand) GetAbbreviation() *string {
	return s.Abbreviation
}

func (s *CreateStandardWordRootRequestCreateCommand) GetDescription() *string {
	return s.Description
}

func (s *CreateStandardWordRootRequestCreateCommand) GetFullName() *string {
	return s.FullName
}

func (s *CreateStandardWordRootRequestCreateCommand) GetName() *string {
	return s.Name
}

func (s *CreateStandardWordRootRequestCreateCommand) SetAbbreviation(v string) *CreateStandardWordRootRequestCreateCommand {
	s.Abbreviation = &v
	return s
}

func (s *CreateStandardWordRootRequestCreateCommand) SetDescription(v string) *CreateStandardWordRootRequestCreateCommand {
	s.Description = &v
	return s
}

func (s *CreateStandardWordRootRequestCreateCommand) SetFullName(v string) *CreateStandardWordRootRequestCreateCommand {
	s.FullName = &v
	return s
}

func (s *CreateStandardWordRootRequestCreateCommand) SetName(v string) *CreateStandardWordRootRequestCreateCommand {
	s.Name = &v
	return s
}

func (s *CreateStandardWordRootRequestCreateCommand) Validate() error {
	return dara.Validate(s)
}
