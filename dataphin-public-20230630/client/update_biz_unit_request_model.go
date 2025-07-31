// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBizUnitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateBizUnitRequest
	GetOpTenantId() *int64
	SetUpdateCommand(v *UpdateBizUnitRequestUpdateCommand) *UpdateBizUnitRequest
	GetUpdateCommand() *UpdateBizUnitRequestUpdateCommand
}

type UpdateBizUnitRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommand *UpdateBizUnitRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateBizUnitRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBizUnitRequest) GoString() string {
	return s.String()
}

func (s *UpdateBizUnitRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateBizUnitRequest) GetUpdateCommand() *UpdateBizUnitRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateBizUnitRequest) SetOpTenantId(v int64) *UpdateBizUnitRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateBizUnitRequest) SetUpdateCommand(v *UpdateBizUnitRequestUpdateCommand) *UpdateBizUnitRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateBizUnitRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateBizUnitRequestUpdateCommand struct {
	// This parameter is required.
	BizUnitAccountList []*UpdateBizUnitRequestUpdateCommandBizUnitAccountList `json:"BizUnitAccountList,omitempty" xml:"BizUnitAccountList,omitempty" type:"Repeated"`
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
	// 测试
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// icon-environment
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// bz_code_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateBizUnitRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateBizUnitRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateBizUnitRequestUpdateCommand) GetBizUnitAccountList() []*UpdateBizUnitRequestUpdateCommandBizUnitAccountList {
	return s.BizUnitAccountList
}

func (s *UpdateBizUnitRequestUpdateCommand) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *UpdateBizUnitRequestUpdateCommand) GetDescription() *string {
	return s.Description
}

func (s *UpdateBizUnitRequestUpdateCommand) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateBizUnitRequestUpdateCommand) GetIcon() *string {
	return s.Icon
}

func (s *UpdateBizUnitRequestUpdateCommand) GetName() *string {
	return s.Name
}

func (s *UpdateBizUnitRequestUpdateCommand) SetBizUnitAccountList(v []*UpdateBizUnitRequestUpdateCommandBizUnitAccountList) *UpdateBizUnitRequestUpdateCommand {
	s.BizUnitAccountList = v
	return s
}

func (s *UpdateBizUnitRequestUpdateCommand) SetBizUnitId(v int64) *UpdateBizUnitRequestUpdateCommand {
	s.BizUnitId = &v
	return s
}

func (s *UpdateBizUnitRequestUpdateCommand) SetDescription(v string) *UpdateBizUnitRequestUpdateCommand {
	s.Description = &v
	return s
}

func (s *UpdateBizUnitRequestUpdateCommand) SetDisplayName(v string) *UpdateBizUnitRequestUpdateCommand {
	s.DisplayName = &v
	return s
}

func (s *UpdateBizUnitRequestUpdateCommand) SetIcon(v string) *UpdateBizUnitRequestUpdateCommand {
	s.Icon = &v
	return s
}

func (s *UpdateBizUnitRequestUpdateCommand) SetName(v string) *UpdateBizUnitRequestUpdateCommand {
	s.Name = &v
	return s
}

func (s *UpdateBizUnitRequestUpdateCommand) Validate() error {
	return dara.Validate(s)
}

type UpdateBizUnitRequestUpdateCommandBizUnitAccountList struct {
	// This parameter is required.
	//
	// example:
	//
	// 20001201
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateBizUnitRequestUpdateCommandBizUnitAccountList) String() string {
	return dara.Prettify(s)
}

func (s UpdateBizUnitRequestUpdateCommandBizUnitAccountList) GoString() string {
	return s.String()
}

func (s *UpdateBizUnitRequestUpdateCommandBizUnitAccountList) GetUserId() *string {
	return s.UserId
}

func (s *UpdateBizUnitRequestUpdateCommandBizUnitAccountList) SetUserId(v string) *UpdateBizUnitRequestUpdateCommandBizUnitAccountList {
	s.UserId = &v
	return s
}

func (s *UpdateBizUnitRequestUpdateCommandBizUnitAccountList) Validate() error {
	return dara.Validate(s)
}
