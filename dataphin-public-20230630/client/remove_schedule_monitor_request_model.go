// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveScheduleMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *RemoveScheduleMonitorRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *RemoveScheduleMonitorRequest
	GetOpUserId() *string
	SetRemoveCommand(v *RemoveScheduleMonitorRequestRemoveCommand) *RemoveScheduleMonitorRequest
	GetRemoveCommand() *RemoveScheduleMonitorRequestRemoveCommand
}

type RemoveScheduleMonitorRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The request body for removing scheduling monitoring.
	//
	// This parameter is required.
	RemoveCommand *RemoveScheduleMonitorRequestRemoveCommand `json:"RemoveCommand,omitempty" xml:"RemoveCommand,omitempty" type:"Struct"`
}

func (s RemoveScheduleMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveScheduleMonitorRequest) GoString() string {
	return s.String()
}

func (s *RemoveScheduleMonitorRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *RemoveScheduleMonitorRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *RemoveScheduleMonitorRequest) GetRemoveCommand() *RemoveScheduleMonitorRequestRemoveCommand {
	return s.RemoveCommand
}

func (s *RemoveScheduleMonitorRequest) SetOpTenantId(v int64) *RemoveScheduleMonitorRequest {
	s.OpTenantId = &v
	return s
}

func (s *RemoveScheduleMonitorRequest) SetOpUserId(v string) *RemoveScheduleMonitorRequest {
	s.OpUserId = &v
	return s
}

func (s *RemoveScheduleMonitorRequest) SetRemoveCommand(v *RemoveScheduleMonitorRequestRemoveCommand) *RemoveScheduleMonitorRequest {
	s.RemoveCommand = v
	return s
}

func (s *RemoveScheduleMonitorRequest) Validate() error {
	if s.RemoveCommand != nil {
		if err := s.RemoveCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RemoveScheduleMonitorRequestRemoveCommand struct {
	// The ID of the monitored item.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6912345678901234
	MonitoredItemId *int64 `json:"MonitoredItemId,omitempty" xml:"MonitoredItemId,omitempty"`
	// The type of the monitored item. Valid values: LOGICAL_FIELD, OND_ID_NODE, SCHEDULE_PHYSICAL_NODE, MANUAL_PHYSICAL_NODE, ODM_NODE, LOGICAL_TABLE_NODE, and PROJECT.
	//
	// This parameter is required.
	//
	// example:
	//
	// SCHEDULE_PHYSICAL_NODE
	MonitoredItemType *string `json:"MonitoredItemType,omitempty" xml:"MonitoredItemType,omitempty"`
}

func (s RemoveScheduleMonitorRequestRemoveCommand) String() string {
	return dara.Prettify(s)
}

func (s RemoveScheduleMonitorRequestRemoveCommand) GoString() string {
	return s.String()
}

func (s *RemoveScheduleMonitorRequestRemoveCommand) GetMonitoredItemId() *int64 {
	return s.MonitoredItemId
}

func (s *RemoveScheduleMonitorRequestRemoveCommand) GetMonitoredItemType() *string {
	return s.MonitoredItemType
}

func (s *RemoveScheduleMonitorRequestRemoveCommand) SetMonitoredItemId(v int64) *RemoveScheduleMonitorRequestRemoveCommand {
	s.MonitoredItemId = &v
	return s
}

func (s *RemoveScheduleMonitorRequestRemoveCommand) SetMonitoredItemType(v string) *RemoveScheduleMonitorRequestRemoveCommand {
	s.MonitoredItemType = &v
	return s
}

func (s *RemoveScheduleMonitorRequestRemoveCommand) Validate() error {
	return dara.Validate(s)
}
