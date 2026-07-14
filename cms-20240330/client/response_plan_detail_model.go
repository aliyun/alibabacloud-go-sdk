// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResponsePlanDetail interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRecoverSeconds(v int64) *ResponsePlanDetail
	GetAutoRecoverSeconds() *int64
	SetEscalationId(v []*string) *ResponsePlanDetail
	GetEscalationId() []*string
	SetPushingSetting(v *ResponsePlanDetailPushingSetting) *ResponsePlanDetail
	GetPushingSetting() *ResponsePlanDetailPushingSetting
	SetRepeatNotifySetting(v *ResponsePlanDetailRepeatNotifySetting) *ResponsePlanDetail
	GetRepeatNotifySetting() *ResponsePlanDetailRepeatNotifySetting
}

type ResponsePlanDetail struct {
	// The number of seconds for automatic recovery. If no new trigger occurs within this duration, the event is automatically recovered.
	//
	// example:
	//
	// 300
	AutoRecoverSeconds *int64 `json:"autoRecoverSeconds,omitempty" xml:"autoRecoverSeconds,omitempty"`
	// The list of escalation plan IDs.
	//
	// example:
	//
	// ["esc-uuid-xxx"]
	EscalationId []*string `json:"escalationId,omitempty" xml:"escalationId,omitempty" type:"Repeated"`
	// The action integration push settings.
	PushingSetting *ResponsePlanDetailPushingSetting `json:"pushingSetting,omitempty" xml:"pushingSetting,omitempty" type:"Struct"`
	// The repeat notification configuration.
	RepeatNotifySetting *ResponsePlanDetailRepeatNotifySetting `json:"repeatNotifySetting,omitempty" xml:"repeatNotifySetting,omitempty" type:"Struct"`
}

func (s ResponsePlanDetail) String() string {
	return dara.Prettify(s)
}

func (s ResponsePlanDetail) GoString() string {
	return s.String()
}

func (s *ResponsePlanDetail) GetAutoRecoverSeconds() *int64 {
	return s.AutoRecoverSeconds
}

func (s *ResponsePlanDetail) GetEscalationId() []*string {
	return s.EscalationId
}

func (s *ResponsePlanDetail) GetPushingSetting() *ResponsePlanDetailPushingSetting {
	return s.PushingSetting
}

func (s *ResponsePlanDetail) GetRepeatNotifySetting() *ResponsePlanDetailRepeatNotifySetting {
	return s.RepeatNotifySetting
}

func (s *ResponsePlanDetail) SetAutoRecoverSeconds(v int64) *ResponsePlanDetail {
	s.AutoRecoverSeconds = &v
	return s
}

func (s *ResponsePlanDetail) SetEscalationId(v []*string) *ResponsePlanDetail {
	s.EscalationId = v
	return s
}

func (s *ResponsePlanDetail) SetPushingSetting(v *ResponsePlanDetailPushingSetting) *ResponsePlanDetail {
	s.PushingSetting = v
	return s
}

func (s *ResponsePlanDetail) SetRepeatNotifySetting(v *ResponsePlanDetailRepeatNotifySetting) *ResponsePlanDetail {
	s.RepeatNotifySetting = v
	return s
}

func (s *ResponsePlanDetail) Validate() error {
	if s.PushingSetting != nil {
		if err := s.PushingSetting.Validate(); err != nil {
			return err
		}
	}
	if s.RepeatNotifySetting != nil {
		if err := s.RepeatNotifySetting.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ResponsePlanDetailPushingSetting struct {
	// The list of action integration IDs triggered by alerts.
	//
	// example:
	//
	// ["action-uuid-xxx"]
	AlertActionIds []*string `json:"alertActionIds,omitempty" xml:"alertActionIds,omitempty" type:"Repeated"`
	// The list of action integration IDs triggered by recovery.
	//
	// example:
	//
	// ["action-uuid-yyy"]
	RestoreActionIds []*string `json:"restoreActionIds,omitempty" xml:"restoreActionIds,omitempty" type:"Repeated"`
}

func (s ResponsePlanDetailPushingSetting) String() string {
	return dara.Prettify(s)
}

func (s ResponsePlanDetailPushingSetting) GoString() string {
	return s.String()
}

func (s *ResponsePlanDetailPushingSetting) GetAlertActionIds() []*string {
	return s.AlertActionIds
}

func (s *ResponsePlanDetailPushingSetting) GetRestoreActionIds() []*string {
	return s.RestoreActionIds
}

func (s *ResponsePlanDetailPushingSetting) SetAlertActionIds(v []*string) *ResponsePlanDetailPushingSetting {
	s.AlertActionIds = v
	return s
}

func (s *ResponsePlanDetailPushingSetting) SetRestoreActionIds(v []*string) *ResponsePlanDetailPushingSetting {
	s.RestoreActionIds = v
	return s
}

func (s *ResponsePlanDetailPushingSetting) Validate() error {
	return dara.Validate(s)
}

type ResponsePlanDetailRepeatNotifySetting struct {
	// The target event status at which repeated notifications stop.
	//
	// example:
	//
	// RECOVERED
	EndIncidentState *string `json:"endIncidentState,omitempty" xml:"endIncidentState,omitempty"`
	// The interval for repeated notifications, in seconds.
	//
	// example:
	//
	// 300
	RepeatInterval *int32 `json:"repeatInterval,omitempty" xml:"repeatInterval,omitempty"`
}

func (s ResponsePlanDetailRepeatNotifySetting) String() string {
	return dara.Prettify(s)
}

func (s ResponsePlanDetailRepeatNotifySetting) GoString() string {
	return s.String()
}

func (s *ResponsePlanDetailRepeatNotifySetting) GetEndIncidentState() *string {
	return s.EndIncidentState
}

func (s *ResponsePlanDetailRepeatNotifySetting) GetRepeatInterval() *int32 {
	return s.RepeatInterval
}

func (s *ResponsePlanDetailRepeatNotifySetting) SetEndIncidentState(v string) *ResponsePlanDetailRepeatNotifySetting {
	s.EndIncidentState = &v
	return s
}

func (s *ResponsePlanDetailRepeatNotifySetting) SetRepeatInterval(v int32) *ResponsePlanDetailRepeatNotifySetting {
	s.RepeatInterval = &v
	return s
}

func (s *ResponsePlanDetailRepeatNotifySetting) Validate() error {
	return dara.Validate(s)
}
