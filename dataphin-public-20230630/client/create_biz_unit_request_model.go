// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBizUnitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommand(v *CreateBizUnitRequestCreateCommand) *CreateBizUnitRequest
	GetCreateCommand() *CreateBizUnitRequestCreateCommand
	SetOpTenantId(v int64) *CreateBizUnitRequest
	GetOpTenantId() *int64
}

type CreateBizUnitRequest struct {
	// This parameter is required.
	CreateCommand *CreateBizUnitRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateBizUnitRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBizUnitRequest) GoString() string {
	return s.String()
}

func (s *CreateBizUnitRequest) GetCreateCommand() *CreateBizUnitRequestCreateCommand {
	return s.CreateCommand
}

func (s *CreateBizUnitRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateBizUnitRequest) SetCreateCommand(v *CreateBizUnitRequestCreateCommand) *CreateBizUnitRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateBizUnitRequest) SetOpTenantId(v int64) *CreateBizUnitRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateBizUnitRequest) Validate() error {
	return dara.Validate(s)
}

type CreateBizUnitRequestCreateCommand struct {
	// This parameter is required.
	BizUnitAccountList []*CreateBizUnitRequestCreateCommandBizUnitAccountList `json:"BizUnitAccountList,omitempty" xml:"BizUnitAccountList,omitempty" type:"Repeated"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// create_object_name
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// icon-environment
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// example:
	//
	// DEV_PROD
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// create_object_code_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateBizUnitRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateBizUnitRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateBizUnitRequestCreateCommand) GetBizUnitAccountList() []*CreateBizUnitRequestCreateCommandBizUnitAccountList {
	return s.BizUnitAccountList
}

func (s *CreateBizUnitRequestCreateCommand) GetDescription() *string {
	return s.Description
}

func (s *CreateBizUnitRequestCreateCommand) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateBizUnitRequestCreateCommand) GetIcon() *string {
	return s.Icon
}

func (s *CreateBizUnitRequestCreateCommand) GetMode() *string {
	return s.Mode
}

func (s *CreateBizUnitRequestCreateCommand) GetName() *string {
	return s.Name
}

func (s *CreateBizUnitRequestCreateCommand) SetBizUnitAccountList(v []*CreateBizUnitRequestCreateCommandBizUnitAccountList) *CreateBizUnitRequestCreateCommand {
	s.BizUnitAccountList = v
	return s
}

func (s *CreateBizUnitRequestCreateCommand) SetDescription(v string) *CreateBizUnitRequestCreateCommand {
	s.Description = &v
	return s
}

func (s *CreateBizUnitRequestCreateCommand) SetDisplayName(v string) *CreateBizUnitRequestCreateCommand {
	s.DisplayName = &v
	return s
}

func (s *CreateBizUnitRequestCreateCommand) SetIcon(v string) *CreateBizUnitRequestCreateCommand {
	s.Icon = &v
	return s
}

func (s *CreateBizUnitRequestCreateCommand) SetMode(v string) *CreateBizUnitRequestCreateCommand {
	s.Mode = &v
	return s
}

func (s *CreateBizUnitRequestCreateCommand) SetName(v string) *CreateBizUnitRequestCreateCommand {
	s.Name = &v
	return s
}

func (s *CreateBizUnitRequestCreateCommand) Validate() error {
	return dara.Validate(s)
}

type CreateBizUnitRequestCreateCommandBizUnitAccountList struct {
	// This parameter is required.
	//
	// example:
	//
	// 20001201
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateBizUnitRequestCreateCommandBizUnitAccountList) String() string {
	return dara.Prettify(s)
}

func (s CreateBizUnitRequestCreateCommandBizUnitAccountList) GoString() string {
	return s.String()
}

func (s *CreateBizUnitRequestCreateCommandBizUnitAccountList) GetUserId() *string {
	return s.UserId
}

func (s *CreateBizUnitRequestCreateCommandBizUnitAccountList) SetUserId(v string) *CreateBizUnitRequestCreateCommandBizUnitAccountList {
	s.UserId = &v
	return s
}

func (s *CreateBizUnitRequestCreateCommandBizUnitAccountList) Validate() error {
	return dara.Validate(s)
}
