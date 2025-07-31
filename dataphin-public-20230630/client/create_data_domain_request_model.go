// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommand(v *CreateDataDomainRequestCreateCommand) *CreateDataDomainRequest
	GetCreateCommand() *CreateDataDomainRequestCreateCommand
	SetOpTenantId(v int64) *CreateDataDomainRequest
	GetOpTenantId() *int64
}

type CreateDataDomainRequest struct {
	// This parameter is required.
	CreateCommand *CreateDataDomainRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateDataDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataDomainRequest) GoString() string {
	return s.String()
}

func (s *CreateDataDomainRequest) GetCreateCommand() *CreateDataDomainRequestCreateCommand {
	return s.CreateCommand
}

func (s *CreateDataDomainRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateDataDomainRequest) SetCreateCommand(v *CreateDataDomainRequestCreateCommand) *CreateDataDomainRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateDataDomainRequest) SetOpTenantId(v int64) *CreateDataDomainRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateDataDomainRequest) Validate() error {
	return dara.Validate(s)
}

type CreateDataDomainRequestCreateCommand struct {
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

func (s CreateDataDomainRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateDataDomainRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateDataDomainRequestCreateCommand) GetAbbreviation() *string {
	return s.Abbreviation
}

func (s *CreateDataDomainRequestCreateCommand) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *CreateDataDomainRequestCreateCommand) GetDescription() *string {
	return s.Description
}

func (s *CreateDataDomainRequestCreateCommand) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateDataDomainRequestCreateCommand) GetName() *string {
	return s.Name
}

func (s *CreateDataDomainRequestCreateCommand) GetParentId() *int64 {
	return s.ParentId
}

func (s *CreateDataDomainRequestCreateCommand) SetAbbreviation(v string) *CreateDataDomainRequestCreateCommand {
	s.Abbreviation = &v
	return s
}

func (s *CreateDataDomainRequestCreateCommand) SetBizUnitId(v int64) *CreateDataDomainRequestCreateCommand {
	s.BizUnitId = &v
	return s
}

func (s *CreateDataDomainRequestCreateCommand) SetDescription(v string) *CreateDataDomainRequestCreateCommand {
	s.Description = &v
	return s
}

func (s *CreateDataDomainRequestCreateCommand) SetDisplayName(v string) *CreateDataDomainRequestCreateCommand {
	s.DisplayName = &v
	return s
}

func (s *CreateDataDomainRequestCreateCommand) SetName(v string) *CreateDataDomainRequestCreateCommand {
	s.Name = &v
	return s
}

func (s *CreateDataDomainRequestCreateCommand) SetParentId(v int64) *CreateDataDomainRequestCreateCommand {
	s.ParentId = &v
	return s
}

func (s *CreateDataDomainRequestCreateCommand) Validate() error {
	return dara.Validate(s)
}
