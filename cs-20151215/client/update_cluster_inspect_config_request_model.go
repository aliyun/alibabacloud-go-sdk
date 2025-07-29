// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClusterInspectConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDisabledCheckItems(v []*string) *UpdateClusterInspectConfigRequest
	GetDisabledCheckItems() []*string
	SetEnabled(v bool) *UpdateClusterInspectConfigRequest
	GetEnabled() *bool
	SetScheduleTime(v string) *UpdateClusterInspectConfigRequest
	GetScheduleTime() *string
}

type UpdateClusterInspectConfigRequest struct {
	// The list of disabled inspection check items.
	DisabledCheckItems []*string `json:"disabledCheckItems,omitempty" xml:"disabledCheckItems,omitempty" type:"Repeated"`
	// Specifies whether to enable cluster inspection.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The inspection period defined using RFC5545 Recurrence Rule. You must specify BYHOUR and BYMINUTE. Only FREQ=DAILY is supported. COUNT or UNTIL is not supported.
	//
	// example:
	//
	// FREQ=DAILY;BYHOUR=10;BYMINUTE=15
	ScheduleTime *string `json:"scheduleTime,omitempty" xml:"scheduleTime,omitempty"`
}

func (s UpdateClusterInspectConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateClusterInspectConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateClusterInspectConfigRequest) GetDisabledCheckItems() []*string {
	return s.DisabledCheckItems
}

func (s *UpdateClusterInspectConfigRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateClusterInspectConfigRequest) GetScheduleTime() *string {
	return s.ScheduleTime
}

func (s *UpdateClusterInspectConfigRequest) SetDisabledCheckItems(v []*string) *UpdateClusterInspectConfigRequest {
	s.DisabledCheckItems = v
	return s
}

func (s *UpdateClusterInspectConfigRequest) SetEnabled(v bool) *UpdateClusterInspectConfigRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateClusterInspectConfigRequest) SetScheduleTime(v string) *UpdateClusterInspectConfigRequest {
	s.ScheduleTime = &v
	return s
}

func (s *UpdateClusterInspectConfigRequest) Validate() error {
	return dara.Validate(s)
}
