// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataAssetsGovernObjectStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateDataAssetsGovernObjectStatusRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *UpdateDataAssetsGovernObjectStatusRequest
	GetOpUserId() *string
	SetUpdateCommand(v *UpdateDataAssetsGovernObjectStatusRequestUpdateCommand) *UpdateDataAssetsGovernObjectStatusRequest
	GetUpdateCommand() *UpdateDataAssetsGovernObjectStatusRequestUpdateCommand
}

type UpdateDataAssetsGovernObjectStatusRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator user.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The update command.
	//
	// This parameter is required.
	UpdateCommand *UpdateDataAssetsGovernObjectStatusRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateDataAssetsGovernObjectStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataAssetsGovernObjectStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataAssetsGovernObjectStatusRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateDataAssetsGovernObjectStatusRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *UpdateDataAssetsGovernObjectStatusRequest) GetUpdateCommand() *UpdateDataAssetsGovernObjectStatusRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateDataAssetsGovernObjectStatusRequest) SetOpTenantId(v int64) *UpdateDataAssetsGovernObjectStatusRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateDataAssetsGovernObjectStatusRequest) SetOpUserId(v string) *UpdateDataAssetsGovernObjectStatusRequest {
	s.OpUserId = &v
	return s
}

func (s *UpdateDataAssetsGovernObjectStatusRequest) SetUpdateCommand(v *UpdateDataAssetsGovernObjectStatusRequestUpdateCommand) *UpdateDataAssetsGovernObjectStatusRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateDataAssetsGovernObjectStatusRequest) Validate() error {
	if s.UpdateCommand != nil {
		if err := s.UpdateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDataAssetsGovernObjectStatusRequestUpdateCommand struct {
	// Specifies whether to notify the owner. This parameter takes effect only when the status is reverted to NEW.
	AlertOwners *bool `json:"AlertOwners,omitempty" xml:"AlertOwners,omitempty"`
	// The list of governance object IDs.
	//
	// This parameter is required.
	GovernObjectIds []*int64 `json:"GovernObjectIds,omitempty" xml:"GovernObjectIds,omitempty" type:"Repeated"`
	// The operation description.
	//
	// example:
	//
	// Issue fixed
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The target status. Valid values: FINISHED / NEW / IGNORE / CANCEL_IGNORE.
	//
	// This parameter is required.
	//
	// example:
	//
	// FINISHED
	TargetStatus *string `json:"TargetStatus,omitempty" xml:"TargetStatus,omitempty"`
}

func (s UpdateDataAssetsGovernObjectStatusRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataAssetsGovernObjectStatusRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateDataAssetsGovernObjectStatusRequestUpdateCommand) GetAlertOwners() *bool {
	return s.AlertOwners
}

func (s *UpdateDataAssetsGovernObjectStatusRequestUpdateCommand) GetGovernObjectIds() []*int64 {
	return s.GovernObjectIds
}

func (s *UpdateDataAssetsGovernObjectStatusRequestUpdateCommand) GetRemark() *string {
	return s.Remark
}

func (s *UpdateDataAssetsGovernObjectStatusRequestUpdateCommand) GetTargetStatus() *string {
	return s.TargetStatus
}

func (s *UpdateDataAssetsGovernObjectStatusRequestUpdateCommand) SetAlertOwners(v bool) *UpdateDataAssetsGovernObjectStatusRequestUpdateCommand {
	s.AlertOwners = &v
	return s
}

func (s *UpdateDataAssetsGovernObjectStatusRequestUpdateCommand) SetGovernObjectIds(v []*int64) *UpdateDataAssetsGovernObjectStatusRequestUpdateCommand {
	s.GovernObjectIds = v
	return s
}

func (s *UpdateDataAssetsGovernObjectStatusRequestUpdateCommand) SetRemark(v string) *UpdateDataAssetsGovernObjectStatusRequestUpdateCommand {
	s.Remark = &v
	return s
}

func (s *UpdateDataAssetsGovernObjectStatusRequestUpdateCommand) SetTargetStatus(v string) *UpdateDataAssetsGovernObjectStatusRequestUpdateCommand {
	s.TargetStatus = &v
	return s
}

func (s *UpdateDataAssetsGovernObjectStatusRequestUpdateCommand) Validate() error {
	return dara.Validate(s)
}
