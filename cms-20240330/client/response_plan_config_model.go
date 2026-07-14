// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResponsePlanConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRecoverSeconds(v int64) *ResponsePlanConfig
	GetAutoRecoverSeconds() *int64
	SetEscalationId(v []*string) *ResponsePlanConfig
	GetEscalationId() []*string
	SetPushingSetting(v *ResponsePlanConfigPushingSetting) *ResponsePlanConfig
	GetPushingSetting() *ResponsePlanConfigPushingSetting
	SetRepeatNotifySetting(v *ResponsePlanConfigRepeatNotifySetting) *ResponsePlanConfig
	GetRepeatNotifySetting() *ResponsePlanConfigRepeatNotifySetting
}

type ResponsePlanConfig struct {
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
	PushingSetting *ResponsePlanConfigPushingSetting `json:"pushingSetting,omitempty" xml:"pushingSetting,omitempty" type:"Struct"`
	// The repeat notification configuration.
	RepeatNotifySetting *ResponsePlanConfigRepeatNotifySetting `json:"repeatNotifySetting,omitempty" xml:"repeatNotifySetting,omitempty" type:"Struct"`
}

func (s ResponsePlanConfig) String() string {
	return dara.Prettify(s)
}

func (s ResponsePlanConfig) GoString() string {
	return s.String()
}

func (s *ResponsePlanConfig) GetAutoRecoverSeconds() *int64 {
	return s.AutoRecoverSeconds
}

func (s *ResponsePlanConfig) GetEscalationId() []*string {
	return s.EscalationId
}

func (s *ResponsePlanConfig) GetPushingSetting() *ResponsePlanConfigPushingSetting {
	return s.PushingSetting
}

func (s *ResponsePlanConfig) GetRepeatNotifySetting() *ResponsePlanConfigRepeatNotifySetting {
	return s.RepeatNotifySetting
}

func (s *ResponsePlanConfig) SetAutoRecoverSeconds(v int64) *ResponsePlanConfig {
	s.AutoRecoverSeconds = &v
	return s
}

func (s *ResponsePlanConfig) SetEscalationId(v []*string) *ResponsePlanConfig {
	s.EscalationId = v
	return s
}

func (s *ResponsePlanConfig) SetPushingSetting(v *ResponsePlanConfigPushingSetting) *ResponsePlanConfig {
	s.PushingSetting = v
	return s
}

func (s *ResponsePlanConfig) SetRepeatNotifySetting(v *ResponsePlanConfigRepeatNotifySetting) *ResponsePlanConfig {
	s.RepeatNotifySetting = v
	return s
}

func (s *ResponsePlanConfig) Validate() error {
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

type ResponsePlanConfigPushingSetting struct {
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

func (s ResponsePlanConfigPushingSetting) String() string {
	return dara.Prettify(s)
}

func (s ResponsePlanConfigPushingSetting) GoString() string {
	return s.String()
}

func (s *ResponsePlanConfigPushingSetting) GetAlertActionIds() []*string {
	return s.AlertActionIds
}

func (s *ResponsePlanConfigPushingSetting) GetRestoreActionIds() []*string {
	return s.RestoreActionIds
}

func (s *ResponsePlanConfigPushingSetting) SetAlertActionIds(v []*string) *ResponsePlanConfigPushingSetting {
	s.AlertActionIds = v
	return s
}

func (s *ResponsePlanConfigPushingSetting) SetRestoreActionIds(v []*string) *ResponsePlanConfigPushingSetting {
	s.RestoreActionIds = v
	return s
}

func (s *ResponsePlanConfigPushingSetting) Validate() error {
	return dara.Validate(s)
}

type ResponsePlanConfigRepeatNotifySetting struct {
	// The target incident status at which repeat notifications stop.
	//
	// example:
	//
	// RECOVERED
	EndIncidentState *string `json:"endIncidentState,omitempty" xml:"endIncidentState,omitempty"`
	// The repeat notification interval, in seconds.
	//
	// example:
	//
	// 300
	RepeatInterval *int32 `json:"repeatInterval,omitempty" xml:"repeatInterval,omitempty"`
}

func (s ResponsePlanConfigRepeatNotifySetting) String() string {
	return dara.Prettify(s)
}

func (s ResponsePlanConfigRepeatNotifySetting) GoString() string {
	return s.String()
}

func (s *ResponsePlanConfigRepeatNotifySetting) GetEndIncidentState() *string {
	return s.EndIncidentState
}

func (s *ResponsePlanConfigRepeatNotifySetting) GetRepeatInterval() *int32 {
	return s.RepeatInterval
}

func (s *ResponsePlanConfigRepeatNotifySetting) SetEndIncidentState(v string) *ResponsePlanConfigRepeatNotifySetting {
	s.EndIncidentState = &v
	return s
}

func (s *ResponsePlanConfigRepeatNotifySetting) SetRepeatInterval(v int32) *ResponsePlanConfigRepeatNotifySetting {
	s.RepeatInterval = &v
	return s
}

func (s *ResponsePlanConfigRepeatNotifySetting) Validate() error {
	return dara.Validate(s)
}
