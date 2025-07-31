// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataSourceBasicInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateDataSourceBasicInfoRequest
	GetOpTenantId() *int64
	SetUpdateCommand(v *UpdateDataSourceBasicInfoRequestUpdateCommand) *UpdateDataSourceBasicInfoRequest
	GetUpdateCommand() *UpdateDataSourceBasicInfoRequestUpdateCommand
}

type UpdateDataSourceBasicInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommand *UpdateDataSourceBasicInfoRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateDataSourceBasicInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataSourceBasicInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceBasicInfoRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateDataSourceBasicInfoRequest) GetUpdateCommand() *UpdateDataSourceBasicInfoRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateDataSourceBasicInfoRequest) SetOpTenantId(v int64) *UpdateDataSourceBasicInfoRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateDataSourceBasicInfoRequest) SetUpdateCommand(v *UpdateDataSourceBasicInfoRequestUpdateCommand) *UpdateDataSourceBasicInfoRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateDataSourceBasicInfoRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateDataSourceBasicInfoRequestUpdateCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// xx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 23231
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateDataSourceBasicInfoRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataSourceBasicInfoRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceBasicInfoRequestUpdateCommand) GetDescription() *string {
	return s.Description
}

func (s *UpdateDataSourceBasicInfoRequestUpdateCommand) GetId() *int64 {
	return s.Id
}

func (s *UpdateDataSourceBasicInfoRequestUpdateCommand) GetName() *string {
	return s.Name
}

func (s *UpdateDataSourceBasicInfoRequestUpdateCommand) SetDescription(v string) *UpdateDataSourceBasicInfoRequestUpdateCommand {
	s.Description = &v
	return s
}

func (s *UpdateDataSourceBasicInfoRequestUpdateCommand) SetId(v int64) *UpdateDataSourceBasicInfoRequestUpdateCommand {
	s.Id = &v
	return s
}

func (s *UpdateDataSourceBasicInfoRequestUpdateCommand) SetName(v string) *UpdateDataSourceBasicInfoRequestUpdateCommand {
	s.Name = &v
	return s
}

func (s *UpdateDataSourceBasicInfoRequestUpdateCommand) Validate() error {
	return dara.Validate(s)
}
