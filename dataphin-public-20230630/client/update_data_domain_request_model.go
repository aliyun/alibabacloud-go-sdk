// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateDataDomainRequest
	GetOpTenantId() *int64
	SetUpdateCommand(v *UpdateDataDomainRequestUpdateCommand) *UpdateDataDomainRequest
	GetUpdateCommand() *UpdateDataDomainRequestUpdateCommand
}

type UpdateDataDomainRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommand *UpdateDataDomainRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateDataDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataDomainRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataDomainRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateDataDomainRequest) GetUpdateCommand() *UpdateDataDomainRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateDataDomainRequest) SetOpTenantId(v int64) *UpdateDataDomainRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateDataDomainRequest) SetUpdateCommand(v *UpdateDataDomainRequestUpdateCommand) *UpdateDataDomainRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateDataDomainRequest) Validate() error {
	if s.UpdateCommand != nil {
		if err := s.UpdateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDataDomainRequestUpdateCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// dm_code_name
	Abbreviation *string `json:"Abbreviation,omitempty" xml:"Abbreviation,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 545844456
	BizUnitId *int64 `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1241844456
	DataDomainId *int64 `json:"DataDomainId,omitempty" xml:"DataDomainId,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 主题域测试
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dm_code_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 10232311
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s UpdateDataDomainRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataDomainRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateDataDomainRequestUpdateCommand) GetAbbreviation() *string {
	return s.Abbreviation
}

func (s *UpdateDataDomainRequestUpdateCommand) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *UpdateDataDomainRequestUpdateCommand) GetDataDomainId() *int64 {
	return s.DataDomainId
}

func (s *UpdateDataDomainRequestUpdateCommand) GetDescription() *string {
	return s.Description
}

func (s *UpdateDataDomainRequestUpdateCommand) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateDataDomainRequestUpdateCommand) GetName() *string {
	return s.Name
}

func (s *UpdateDataDomainRequestUpdateCommand) GetParentId() *int64 {
	return s.ParentId
}

func (s *UpdateDataDomainRequestUpdateCommand) SetAbbreviation(v string) *UpdateDataDomainRequestUpdateCommand {
	s.Abbreviation = &v
	return s
}

func (s *UpdateDataDomainRequestUpdateCommand) SetBizUnitId(v int64) *UpdateDataDomainRequestUpdateCommand {
	s.BizUnitId = &v
	return s
}

func (s *UpdateDataDomainRequestUpdateCommand) SetDataDomainId(v int64) *UpdateDataDomainRequestUpdateCommand {
	s.DataDomainId = &v
	return s
}

func (s *UpdateDataDomainRequestUpdateCommand) SetDescription(v string) *UpdateDataDomainRequestUpdateCommand {
	s.Description = &v
	return s
}

func (s *UpdateDataDomainRequestUpdateCommand) SetDisplayName(v string) *UpdateDataDomainRequestUpdateCommand {
	s.DisplayName = &v
	return s
}

func (s *UpdateDataDomainRequestUpdateCommand) SetName(v string) *UpdateDataDomainRequestUpdateCommand {
	s.Name = &v
	return s
}

func (s *UpdateDataDomainRequestUpdateCommand) SetParentId(v int64) *UpdateDataDomainRequestUpdateCommand {
	s.ParentId = &v
	return s
}

func (s *UpdateDataDomainRequestUpdateCommand) Validate() error {
	return dara.Validate(s)
}
