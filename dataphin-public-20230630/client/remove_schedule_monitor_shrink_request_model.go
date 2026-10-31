// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveScheduleMonitorShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *RemoveScheduleMonitorShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *RemoveScheduleMonitorShrinkRequest
	GetOpUserId() *string
	SetRemoveCommandShrink(v string) *RemoveScheduleMonitorShrinkRequest
	GetRemoveCommandShrink() *string
}

type RemoveScheduleMonitorShrinkRequest struct {
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
	RemoveCommandShrink *string `json:"RemoveCommand,omitempty" xml:"RemoveCommand,omitempty"`
}

func (s RemoveScheduleMonitorShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveScheduleMonitorShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveScheduleMonitorShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *RemoveScheduleMonitorShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *RemoveScheduleMonitorShrinkRequest) GetRemoveCommandShrink() *string {
	return s.RemoveCommandShrink
}

func (s *RemoveScheduleMonitorShrinkRequest) SetOpTenantId(v int64) *RemoveScheduleMonitorShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *RemoveScheduleMonitorShrinkRequest) SetOpUserId(v string) *RemoveScheduleMonitorShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *RemoveScheduleMonitorShrinkRequest) SetRemoveCommandShrink(v string) *RemoveScheduleMonitorShrinkRequest {
	s.RemoveCommandShrink = &v
	return s
}

func (s *RemoveScheduleMonitorShrinkRequest) Validate() error {
	return dara.Validate(s)
}
