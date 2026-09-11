// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncidentResponsePlanForSNSModify interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRecoverSeconds(v int64) *IncidentResponsePlanForSNSModify
	GetAutoRecoverSeconds() *int64
	SetEscalationId(v []*string) *IncidentResponsePlanForSNSModify
	GetEscalationId() []*string
	SetPushingSetting(v *IncidentResponsePlanForSNSModifyPushingSetting) *IncidentResponsePlanForSNSModify
	GetPushingSetting() *IncidentResponsePlanForSNSModifyPushingSetting
	SetRepeatNotifySetting(v *IncidentResponsePlanForSNSModifyRepeatNotifySetting) *IncidentResponsePlanForSNSModify
	GetRepeatNotifySetting() *IncidentResponsePlanForSNSModifyRepeatNotifySetting
}

type IncidentResponsePlanForSNSModify struct {
	// The auto-recovery time. Unit: seconds. After this is configured, if no new events are generated for the incident within this period, the incident is automatically marked as resolved.
	//
	// example:
	//
	// 3600
	AutoRecoverSeconds *int64 `json:"autoRecoverSeconds,omitempty" xml:"autoRecoverSeconds,omitempty"`
	// The list of escalation policy IDs. Associates with IncidentEscalationPolicy to define step-by-step escalation rules when an incident is not handled as expected, such as notifying a supervisor if the incident is not acknowledged within 30 minutes.
	EscalationId []*string `json:"escalationId,omitempty" xml:"escalationId,omitempty" type:"Repeated"`
	// The action integration execution configuration that defines automated actions to trigger when an incident occurs and when it is recovered.
	PushingSetting *IncidentResponsePlanForSNSModifyPushingSetting `json:"pushingSetting,omitempty" xml:"pushingSetting,omitempty" type:"Struct"`
	// The repeat notification settings. When an incident remains unresolved, notifications are sent repeatedly at a fixed interval.
	RepeatNotifySetting *IncidentResponsePlanForSNSModifyRepeatNotifySetting `json:"repeatNotifySetting,omitempty" xml:"repeatNotifySetting,omitempty" type:"Struct"`
}

func (s IncidentResponsePlanForSNSModify) String() string {
	return dara.Prettify(s)
}

func (s IncidentResponsePlanForSNSModify) GoString() string {
	return s.String()
}

func (s *IncidentResponsePlanForSNSModify) GetAutoRecoverSeconds() *int64 {
	return s.AutoRecoverSeconds
}

func (s *IncidentResponsePlanForSNSModify) GetEscalationId() []*string {
	return s.EscalationId
}

func (s *IncidentResponsePlanForSNSModify) GetPushingSetting() *IncidentResponsePlanForSNSModifyPushingSetting {
	return s.PushingSetting
}

func (s *IncidentResponsePlanForSNSModify) GetRepeatNotifySetting() *IncidentResponsePlanForSNSModifyRepeatNotifySetting {
	return s.RepeatNotifySetting
}

func (s *IncidentResponsePlanForSNSModify) SetAutoRecoverSeconds(v int64) *IncidentResponsePlanForSNSModify {
	s.AutoRecoverSeconds = &v
	return s
}

func (s *IncidentResponsePlanForSNSModify) SetEscalationId(v []*string) *IncidentResponsePlanForSNSModify {
	s.EscalationId = v
	return s
}

func (s *IncidentResponsePlanForSNSModify) SetPushingSetting(v *IncidentResponsePlanForSNSModifyPushingSetting) *IncidentResponsePlanForSNSModify {
	s.PushingSetting = v
	return s
}

func (s *IncidentResponsePlanForSNSModify) SetRepeatNotifySetting(v *IncidentResponsePlanForSNSModifyRepeatNotifySetting) *IncidentResponsePlanForSNSModify {
	s.RepeatNotifySetting = v
	return s
}

func (s *IncidentResponsePlanForSNSModify) Validate() error {
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

type IncidentResponsePlanForSNSModifyPushingSetting struct {
	// The list of action IDs to execute when an event is triggered. Actions must be created in advance by calling CreateAlertAction.
	AlertActionIds []*string `json:"alertActionIds,omitempty" xml:"alertActionIds,omitempty" type:"Repeated"`
	// The list of action IDs to execute when an event is recovered.
	RestoreActionIds []*string `json:"restoreActionIds,omitempty" xml:"restoreActionIds,omitempty" type:"Repeated"`
	// Deprecated. This parameter does not take effect even if a value is passed in.
	//
	// example:
	//
	// uuid
	TemplateUuid *string `json:"templateUuid,omitempty" xml:"templateUuid,omitempty"`
}

func (s IncidentResponsePlanForSNSModifyPushingSetting) String() string {
	return dara.Prettify(s)
}

func (s IncidentResponsePlanForSNSModifyPushingSetting) GoString() string {
	return s.String()
}

func (s *IncidentResponsePlanForSNSModifyPushingSetting) GetAlertActionIds() []*string {
	return s.AlertActionIds
}

func (s *IncidentResponsePlanForSNSModifyPushingSetting) GetRestoreActionIds() []*string {
	return s.RestoreActionIds
}

func (s *IncidentResponsePlanForSNSModifyPushingSetting) GetTemplateUuid() *string {
	return s.TemplateUuid
}

func (s *IncidentResponsePlanForSNSModifyPushingSetting) SetAlertActionIds(v []*string) *IncidentResponsePlanForSNSModifyPushingSetting {
	s.AlertActionIds = v
	return s
}

func (s *IncidentResponsePlanForSNSModifyPushingSetting) SetRestoreActionIds(v []*string) *IncidentResponsePlanForSNSModifyPushingSetting {
	s.RestoreActionIds = v
	return s
}

func (s *IncidentResponsePlanForSNSModifyPushingSetting) SetTemplateUuid(v string) *IncidentResponsePlanForSNSModifyPushingSetting {
	s.TemplateUuid = &v
	return s
}

func (s *IncidentResponsePlanForSNSModifyPushingSetting) Validate() error {
	return dara.Validate(s)
}

type IncidentResponsePlanForSNSModifyRepeatNotifySetting struct {
	// The incident status at which repeat notifications stop. Repeat notifications are no longer sent after the incident reaches this status.
	//
	// example:
	//
	// resolved
	EndIncidentState *string `json:"endIncidentState,omitempty" xml:"endIncidentState,omitempty"`
	// The repeat notification interval. Unit: seconds.
	//
	// example:
	//
	// 300
	RepeatInterval *int32 `json:"repeatInterval,omitempty" xml:"repeatInterval,omitempty"`
}

func (s IncidentResponsePlanForSNSModifyRepeatNotifySetting) String() string {
	return dara.Prettify(s)
}

func (s IncidentResponsePlanForSNSModifyRepeatNotifySetting) GoString() string {
	return s.String()
}

func (s *IncidentResponsePlanForSNSModifyRepeatNotifySetting) GetEndIncidentState() *string {
	return s.EndIncidentState
}

func (s *IncidentResponsePlanForSNSModifyRepeatNotifySetting) GetRepeatInterval() *int32 {
	return s.RepeatInterval
}

func (s *IncidentResponsePlanForSNSModifyRepeatNotifySetting) SetEndIncidentState(v string) *IncidentResponsePlanForSNSModifyRepeatNotifySetting {
	s.EndIncidentState = &v
	return s
}

func (s *IncidentResponsePlanForSNSModifyRepeatNotifySetting) SetRepeatInterval(v int32) *IncidentResponsePlanForSNSModifyRepeatNotifySetting {
	s.RepeatInterval = &v
	return s
}

func (s *IncidentResponsePlanForSNSModifyRepeatNotifySetting) Validate() error {
	return dara.Validate(s)
}
