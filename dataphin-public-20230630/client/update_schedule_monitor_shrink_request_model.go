// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScheduleMonitorShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateScheduleMonitorShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *UpdateScheduleMonitorShrinkRequest
	GetOpUserId() *string
	SetUpdateCommandShrink(v string) *UpdateScheduleMonitorShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdateScheduleMonitorShrinkRequest struct {
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
	// The request body for updating scheduling monitoring.
	//
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateScheduleMonitorShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduleMonitorShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateScheduleMonitorShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateScheduleMonitorShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *UpdateScheduleMonitorShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateScheduleMonitorShrinkRequest) SetOpTenantId(v int64) *UpdateScheduleMonitorShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateScheduleMonitorShrinkRequest) SetOpUserId(v string) *UpdateScheduleMonitorShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *UpdateScheduleMonitorShrinkRequest) SetUpdateCommandShrink(v string) *UpdateScheduleMonitorShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateScheduleMonitorShrinkRequest) Validate() error {
	return dara.Validate(s)
}
