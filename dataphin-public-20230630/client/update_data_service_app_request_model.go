// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataServiceAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateDataServiceAppRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *UpdateDataServiceAppRequest
	GetOpUserId() *string
	SetUpdateCommand(v *UpdateDataServiceAppRequestUpdateCommand) *UpdateDataServiceAppRequest
	GetUpdateCommand() *UpdateDataServiceAppRequestUpdateCommand
}

type UpdateDataServiceAppRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The operator ID.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The update command.
	//
	// This parameter is required.
	UpdateCommand *UpdateDataServiceAppRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateDataServiceAppRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataServiceAppRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataServiceAppRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateDataServiceAppRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *UpdateDataServiceAppRequest) GetUpdateCommand() *UpdateDataServiceAppRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateDataServiceAppRequest) SetOpTenantId(v int64) *UpdateDataServiceAppRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateDataServiceAppRequest) SetOpUserId(v string) *UpdateDataServiceAppRequest {
	s.OpUserId = &v
	return s
}

func (s *UpdateDataServiceAppRequest) SetUpdateCommand(v *UpdateDataServiceAppRequestUpdateCommand) *UpdateDataServiceAppRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateDataServiceAppRequest) Validate() error {
	if s.UpdateCommand != nil {
		if err := s.UpdateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDataServiceAppRequestUpdateCommand struct {
	// The application group ID.
	//
	// example:
	//
	// 200000000
	AppGroupId *int32 `json:"AppGroupId,omitempty" xml:"AppGroupId,omitempty"`
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	AppId *int32 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// Default application
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The owner IDs.
	OwnerIds []*string `json:"OwnerIds,omitempty" xml:"OwnerIds,omitempty" type:"Repeated"`
	// The common scenarios.
	//
	// example:
	//
	// Data analytics
	Scenarios *string `json:"Scenarios,omitempty" xml:"Scenarios,omitempty"`
}

func (s UpdateDataServiceAppRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataServiceAppRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateDataServiceAppRequestUpdateCommand) GetAppGroupId() *int32 {
	return s.AppGroupId
}

func (s *UpdateDataServiceAppRequestUpdateCommand) GetAppId() *int32 {
	return s.AppId
}

func (s *UpdateDataServiceAppRequestUpdateCommand) GetAppName() *string {
	return s.AppName
}

func (s *UpdateDataServiceAppRequestUpdateCommand) GetOwnerIds() []*string {
	return s.OwnerIds
}

func (s *UpdateDataServiceAppRequestUpdateCommand) GetScenarios() *string {
	return s.Scenarios
}

func (s *UpdateDataServiceAppRequestUpdateCommand) SetAppGroupId(v int32) *UpdateDataServiceAppRequestUpdateCommand {
	s.AppGroupId = &v
	return s
}

func (s *UpdateDataServiceAppRequestUpdateCommand) SetAppId(v int32) *UpdateDataServiceAppRequestUpdateCommand {
	s.AppId = &v
	return s
}

func (s *UpdateDataServiceAppRequestUpdateCommand) SetAppName(v string) *UpdateDataServiceAppRequestUpdateCommand {
	s.AppName = &v
	return s
}

func (s *UpdateDataServiceAppRequestUpdateCommand) SetOwnerIds(v []*string) *UpdateDataServiceAppRequestUpdateCommand {
	s.OwnerIds = v
	return s
}

func (s *UpdateDataServiceAppRequestUpdateCommand) SetScenarios(v string) *UpdateDataServiceAppRequestUpdateCommand {
	s.Scenarios = &v
	return s
}

func (s *UpdateDataServiceAppRequestUpdateCommand) Validate() error {
	return dara.Validate(s)
}
