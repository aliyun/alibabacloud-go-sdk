// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduleMonitorShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommandShrink(v string) *CreateScheduleMonitorShrinkRequest
	GetCreateCommandShrink() *string
	SetOpTenantId(v int64) *CreateScheduleMonitorShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *CreateScheduleMonitorShrinkRequest
	GetOpUserId() *string
}

type CreateScheduleMonitorShrinkRequest struct {
	// The request body for creating a scheduling monitoring configuration.
	//
	// This parameter is required.
	CreateCommandShrink *string `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The operator user ID.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s CreateScheduleMonitorShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleMonitorShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduleMonitorShrinkRequest) GetCreateCommandShrink() *string {
	return s.CreateCommandShrink
}

func (s *CreateScheduleMonitorShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateScheduleMonitorShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *CreateScheduleMonitorShrinkRequest) SetCreateCommandShrink(v string) *CreateScheduleMonitorShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateScheduleMonitorShrinkRequest) SetOpTenantId(v int64) *CreateScheduleMonitorShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateScheduleMonitorShrinkRequest) SetOpUserId(v string) *CreateScheduleMonitorShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *CreateScheduleMonitorShrinkRequest) Validate() error {
	return dara.Validate(s)
}
