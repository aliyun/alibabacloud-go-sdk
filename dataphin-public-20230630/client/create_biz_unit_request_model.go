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
	// The create request.
	//
	// This parameter is required.
	CreateCommand *CreateBizUnitRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// The tenant ID.
	//
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
	if s.CreateCommand != nil {
		if err := s.CreateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateBizUnitRequestCreateCommand struct {
	// The list of data domain architects.
	//
	// This parameter is required.
	BizUnitAccountList []*CreateBizUnitRequestCreateCommandBizUnitAccountList `json:"BizUnitAccountList,omitempty" xml:"BizUnitAccountList,omitempty" type:"Repeated"`
	// The description of the business object. The description can be up to 128 characters in length.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the business object. The name can be up to 64 characters in length and can contain only Chinese characters, letters, digits, underscores, and hyphens.
	//
	// This parameter is required.
	//
	// example:
	//
	// create_object_name
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The preset icon. Valid values:
	//
	// - icon-e-commerce: E-commerce.
	//
	// - icon-finance: Finance.
	//
	// - con-cloud-computing: Cloud computing.
	//
	// - icon-advertisement: Advertising and marketing.
	//
	// - icon-logistics: Logistics.
	//
	// - icon-entertainment: Entertainment.
	//
	// - icon-traffic: Travel.
	//
	// - icon-health: Health.
	//
	// - icon-social-contact: Social and communication.
	//
	// - con-dining: Dining.
	//
	// - icon-education: Education.
	//
	// - icon-environment: Environment.
	//
	// This parameter is required.
	//
	// example:
	//
	// icon-environment
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// The production mode. Valid values:
	//
	// - BASIC: single-environment mode.
	//
	// - DEV_PROD: development/production dual-environment mode.
	//
	// example:
	//
	// DEV_PROD
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The code name of the business object. The name can be up to 64 characters in length and can contain only letters, digits, and underscores. For ADB_PG engines, the code name can be up to 40 characters in length.
	//
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
	if s.BizUnitAccountList != nil {
		for _, item := range s.BizUnitAccountList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateBizUnitRequestCreateCommandBizUnitAccountList struct {
	// The user ID.
	//
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
