// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStandardWordRootRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateStandardWordRootRequest
	GetOpTenantId() *int64
	SetUpdateCommand(v *UpdateStandardWordRootRequestUpdateCommand) *UpdateStandardWordRootRequest
	GetUpdateCommand() *UpdateStandardWordRootRequestUpdateCommand
}

type UpdateStandardWordRootRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommand *UpdateStandardWordRootRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateStandardWordRootRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardWordRootRequest) GoString() string {
	return s.String()
}

func (s *UpdateStandardWordRootRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateStandardWordRootRequest) GetUpdateCommand() *UpdateStandardWordRootRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateStandardWordRootRequest) SetOpTenantId(v int64) *UpdateStandardWordRootRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateStandardWordRootRequest) SetUpdateCommand(v *UpdateStandardWordRootRequestUpdateCommand) *UpdateStandardWordRootRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateStandardWordRootRequest) Validate() error {
	if s.UpdateCommand != nil {
		if err := s.UpdateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateStandardWordRootRequestUpdateCommand struct {
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
	// This parameter is required.
	//
	// example:
	//
	// 平均值
	OldName *string `json:"OldName,omitempty" xml:"OldName,omitempty"`
}

func (s UpdateStandardWordRootRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardWordRootRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateStandardWordRootRequestUpdateCommand) GetAbbreviation() *string {
	return s.Abbreviation
}

func (s *UpdateStandardWordRootRequestUpdateCommand) GetDescription() *string {
	return s.Description
}

func (s *UpdateStandardWordRootRequestUpdateCommand) GetFullName() *string {
	return s.FullName
}

func (s *UpdateStandardWordRootRequestUpdateCommand) GetName() *string {
	return s.Name
}

func (s *UpdateStandardWordRootRequestUpdateCommand) GetOldName() *string {
	return s.OldName
}

func (s *UpdateStandardWordRootRequestUpdateCommand) SetAbbreviation(v string) *UpdateStandardWordRootRequestUpdateCommand {
	s.Abbreviation = &v
	return s
}

func (s *UpdateStandardWordRootRequestUpdateCommand) SetDescription(v string) *UpdateStandardWordRootRequestUpdateCommand {
	s.Description = &v
	return s
}

func (s *UpdateStandardWordRootRequestUpdateCommand) SetFullName(v string) *UpdateStandardWordRootRequestUpdateCommand {
	s.FullName = &v
	return s
}

func (s *UpdateStandardWordRootRequestUpdateCommand) SetName(v string) *UpdateStandardWordRootRequestUpdateCommand {
	s.Name = &v
	return s
}

func (s *UpdateStandardWordRootRequestUpdateCommand) SetOldName(v string) *UpdateStandardWordRootRequestUpdateCommand {
	s.OldName = &v
	return s
}

func (s *UpdateStandardWordRootRequestUpdateCommand) Validate() error {
	return dara.Validate(s)
}
