// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQualityWatchSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateQualityWatchSwitchRequest
	GetOpTenantId() *int64
	SetUpdateCommand(v *UpdateQualityWatchSwitchRequestUpdateCommand) *UpdateQualityWatchSwitchRequest
	GetUpdateCommand() *UpdateQualityWatchSwitchRequestUpdateCommand
}

type UpdateQualityWatchSwitchRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The update instruction.
	//
	// This parameter is required.
	UpdateCommand *UpdateQualityWatchSwitchRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateQualityWatchSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateQualityWatchSwitchRequest) GoString() string {
	return s.String()
}

func (s *UpdateQualityWatchSwitchRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateQualityWatchSwitchRequest) GetUpdateCommand() *UpdateQualityWatchSwitchRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateQualityWatchSwitchRequest) SetOpTenantId(v int64) *UpdateQualityWatchSwitchRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateQualityWatchSwitchRequest) SetUpdateCommand(v *UpdateQualityWatchSwitchRequestUpdateCommand) *UpdateQualityWatchSwitchRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateQualityWatchSwitchRequest) Validate() error {
	if s.UpdateCommand != nil {
		if err := s.UpdateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateQualityWatchSwitchRequestUpdateCommand struct {
	// Specifies whether to enable the monitored object.
	//
	// This parameter is required.
	Open *bool `json:"Open,omitempty" xml:"Open,omitempty"`
	// The list of monitoring IDs.
	//
	// This parameter is required.
	WatchIdList []*int64 `json:"WatchIdList,omitempty" xml:"WatchIdList,omitempty" type:"Repeated"`
}

func (s UpdateQualityWatchSwitchRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateQualityWatchSwitchRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateQualityWatchSwitchRequestUpdateCommand) GetOpen() *bool {
	return s.Open
}

func (s *UpdateQualityWatchSwitchRequestUpdateCommand) GetWatchIdList() []*int64 {
	return s.WatchIdList
}

func (s *UpdateQualityWatchSwitchRequestUpdateCommand) SetOpen(v bool) *UpdateQualityWatchSwitchRequestUpdateCommand {
	s.Open = &v
	return s
}

func (s *UpdateQualityWatchSwitchRequestUpdateCommand) SetWatchIdList(v []*int64) *UpdateQualityWatchSwitchRequestUpdateCommand {
	s.WatchIdList = v
	return s
}

func (s *UpdateQualityWatchSwitchRequestUpdateCommand) Validate() error {
	return dara.Validate(s)
}
