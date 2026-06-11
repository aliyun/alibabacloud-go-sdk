// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncidentResponsePlanForSNSView interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRecoverSeconds(v int64) *IncidentResponsePlanForSNSView
	GetAutoRecoverSeconds() *int64
	SetCreateTime(v string) *IncidentResponsePlanForSNSView
	GetCreateTime() *string
	SetEnable(v bool) *IncidentResponsePlanForSNSView
	GetEnable() *bool
	SetEscalationId(v []*string) *IncidentResponsePlanForSNSView
	GetEscalationId() []*string
	SetMode(v string) *IncidentResponsePlanForSNSView
	GetMode() *string
	SetName(v string) *IncidentResponsePlanForSNSView
	GetName() *string
	SetPushingSetting(v *IncidentResponsePlanForSNSViewPushingSetting) *IncidentResponsePlanForSNSView
	GetPushingSetting() *IncidentResponsePlanForSNSViewPushingSetting
	SetRepeatNotifySetting(v *IncidentResponsePlanForSNSViewRepeatNotifySetting) *IncidentResponsePlanForSNSView
	GetRepeatNotifySetting() *IncidentResponsePlanForSNSViewRepeatNotifySetting
	SetSource(v string) *IncidentResponsePlanForSNSView
	GetSource() *string
	SetSyncFromType(v string) *IncidentResponsePlanForSNSView
	GetSyncFromType() *string
	SetType(v string) *IncidentResponsePlanForSNSView
	GetType() *string
	SetUpdateTime(v string) *IncidentResponsePlanForSNSView
	GetUpdateTime() *string
	SetUuid(v string) *IncidentResponsePlanForSNSView
	GetUuid() *string
}

type IncidentResponsePlanForSNSView struct {
	// The duration, in seconds, after which an incident without new alerts is automatically resolved.
	AutoRecoverSeconds *int64 `json:"autoRecoverSeconds,omitempty" xml:"autoRecoverSeconds,omitempty"`
	// The time when the incident response plan was created, in UTC and in the `YYYY-MM-DDThh:mm:ssZ` format.
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// Indicates if the incident response plan is enabled. Valid values: `true` and `false`.
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The IDs of the escalation policies.
	EscalationId []*string `json:"escalationId,omitempty" xml:"escalationId,omitempty" type:"Repeated"`
	// The mode of the incident response plan. Valid values: `AUTO` and `MANUAL`.
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// The name of the incident response plan.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The settings for sending notifications.
	PushingSetting *IncidentResponsePlanForSNSViewPushingSetting `json:"pushingSetting,omitempty" xml:"pushingSetting,omitempty" type:"Struct"`
	// The settings for repeated notifications.
	RepeatNotifySetting *IncidentResponsePlanForSNSViewRepeatNotifySetting `json:"repeatNotifySetting,omitempty" xml:"repeatNotifySetting,omitempty" type:"Struct"`
	// The source of the incident.
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The synchronization source type.
	SyncFromType *string `json:"syncFromType,omitempty" xml:"syncFromType,omitempty"`
	// The type of the incident response plan.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The time when the incident response plan was last updated, in UTC and in the `YYYY-MM-DDThh:mm:ssZ` format.
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The unique ID of the incident response plan.
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s IncidentResponsePlanForSNSView) String() string {
	return dara.Prettify(s)
}

func (s IncidentResponsePlanForSNSView) GoString() string {
	return s.String()
}

func (s *IncidentResponsePlanForSNSView) GetAutoRecoverSeconds() *int64 {
	return s.AutoRecoverSeconds
}

func (s *IncidentResponsePlanForSNSView) GetCreateTime() *string {
	return s.CreateTime
}

func (s *IncidentResponsePlanForSNSView) GetEnable() *bool {
	return s.Enable
}

func (s *IncidentResponsePlanForSNSView) GetEscalationId() []*string {
	return s.EscalationId
}

func (s *IncidentResponsePlanForSNSView) GetMode() *string {
	return s.Mode
}

func (s *IncidentResponsePlanForSNSView) GetName() *string {
	return s.Name
}

func (s *IncidentResponsePlanForSNSView) GetPushingSetting() *IncidentResponsePlanForSNSViewPushingSetting {
	return s.PushingSetting
}

func (s *IncidentResponsePlanForSNSView) GetRepeatNotifySetting() *IncidentResponsePlanForSNSViewRepeatNotifySetting {
	return s.RepeatNotifySetting
}

func (s *IncidentResponsePlanForSNSView) GetSource() *string {
	return s.Source
}

func (s *IncidentResponsePlanForSNSView) GetSyncFromType() *string {
	return s.SyncFromType
}

func (s *IncidentResponsePlanForSNSView) GetType() *string {
	return s.Type
}

func (s *IncidentResponsePlanForSNSView) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *IncidentResponsePlanForSNSView) GetUuid() *string {
	return s.Uuid
}

func (s *IncidentResponsePlanForSNSView) SetAutoRecoverSeconds(v int64) *IncidentResponsePlanForSNSView {
	s.AutoRecoverSeconds = &v
	return s
}

func (s *IncidentResponsePlanForSNSView) SetCreateTime(v string) *IncidentResponsePlanForSNSView {
	s.CreateTime = &v
	return s
}

func (s *IncidentResponsePlanForSNSView) SetEnable(v bool) *IncidentResponsePlanForSNSView {
	s.Enable = &v
	return s
}

func (s *IncidentResponsePlanForSNSView) SetEscalationId(v []*string) *IncidentResponsePlanForSNSView {
	s.EscalationId = v
	return s
}

func (s *IncidentResponsePlanForSNSView) SetMode(v string) *IncidentResponsePlanForSNSView {
	s.Mode = &v
	return s
}

func (s *IncidentResponsePlanForSNSView) SetName(v string) *IncidentResponsePlanForSNSView {
	s.Name = &v
	return s
}

func (s *IncidentResponsePlanForSNSView) SetPushingSetting(v *IncidentResponsePlanForSNSViewPushingSetting) *IncidentResponsePlanForSNSView {
	s.PushingSetting = v
	return s
}

func (s *IncidentResponsePlanForSNSView) SetRepeatNotifySetting(v *IncidentResponsePlanForSNSViewRepeatNotifySetting) *IncidentResponsePlanForSNSView {
	s.RepeatNotifySetting = v
	return s
}

func (s *IncidentResponsePlanForSNSView) SetSource(v string) *IncidentResponsePlanForSNSView {
	s.Source = &v
	return s
}

func (s *IncidentResponsePlanForSNSView) SetSyncFromType(v string) *IncidentResponsePlanForSNSView {
	s.SyncFromType = &v
	return s
}

func (s *IncidentResponsePlanForSNSView) SetType(v string) *IncidentResponsePlanForSNSView {
	s.Type = &v
	return s
}

func (s *IncidentResponsePlanForSNSView) SetUpdateTime(v string) *IncidentResponsePlanForSNSView {
	s.UpdateTime = &v
	return s
}

func (s *IncidentResponsePlanForSNSView) SetUuid(v string) *IncidentResponsePlanForSNSView {
	s.Uuid = &v
	return s
}

func (s *IncidentResponsePlanForSNSView) Validate() error {
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

type IncidentResponsePlanForSNSViewPushingSetting struct {
	// The IDs of actions to run when an alert is triggered.
	AlertActionIds []*string `json:"alertActionIds,omitempty" xml:"alertActionIds,omitempty" type:"Repeated"`
	// The IDs of actions to run when the incident is resolved.
	RestoreActionIds []*string `json:"restoreActionIds,omitempty" xml:"restoreActionIds,omitempty" type:"Repeated"`
	// The ID of the notification template.
	TemplateUuid *string `json:"templateUuid,omitempty" xml:"templateUuid,omitempty"`
}

func (s IncidentResponsePlanForSNSViewPushingSetting) String() string {
	return dara.Prettify(s)
}

func (s IncidentResponsePlanForSNSViewPushingSetting) GoString() string {
	return s.String()
}

func (s *IncidentResponsePlanForSNSViewPushingSetting) GetAlertActionIds() []*string {
	return s.AlertActionIds
}

func (s *IncidentResponsePlanForSNSViewPushingSetting) GetRestoreActionIds() []*string {
	return s.RestoreActionIds
}

func (s *IncidentResponsePlanForSNSViewPushingSetting) GetTemplateUuid() *string {
	return s.TemplateUuid
}

func (s *IncidentResponsePlanForSNSViewPushingSetting) SetAlertActionIds(v []*string) *IncidentResponsePlanForSNSViewPushingSetting {
	s.AlertActionIds = v
	return s
}

func (s *IncidentResponsePlanForSNSViewPushingSetting) SetRestoreActionIds(v []*string) *IncidentResponsePlanForSNSViewPushingSetting {
	s.RestoreActionIds = v
	return s
}

func (s *IncidentResponsePlanForSNSViewPushingSetting) SetTemplateUuid(v string) *IncidentResponsePlanForSNSViewPushingSetting {
	s.TemplateUuid = &v
	return s
}

func (s *IncidentResponsePlanForSNSViewPushingSetting) Validate() error {
	return dara.Validate(s)
}

type IncidentResponsePlanForSNSViewRepeatNotifySetting struct {
	// The incident state at which repeated notifications stop. For example, `resolved`.
	EndIncidentState *string `json:"endIncidentState,omitempty" xml:"endIncidentState,omitempty"`
	// The interval, in seconds, for repeated notifications.
	RepeatInterval *int32 `json:"repeatInterval,omitempty" xml:"repeatInterval,omitempty"`
}

func (s IncidentResponsePlanForSNSViewRepeatNotifySetting) String() string {
	return dara.Prettify(s)
}

func (s IncidentResponsePlanForSNSViewRepeatNotifySetting) GoString() string {
	return s.String()
}

func (s *IncidentResponsePlanForSNSViewRepeatNotifySetting) GetEndIncidentState() *string {
	return s.EndIncidentState
}

func (s *IncidentResponsePlanForSNSViewRepeatNotifySetting) GetRepeatInterval() *int32 {
	return s.RepeatInterval
}

func (s *IncidentResponsePlanForSNSViewRepeatNotifySetting) SetEndIncidentState(v string) *IncidentResponsePlanForSNSViewRepeatNotifySetting {
	s.EndIncidentState = &v
	return s
}

func (s *IncidentResponsePlanForSNSViewRepeatNotifySetting) SetRepeatInterval(v int32) *IncidentResponsePlanForSNSViewRepeatNotifySetting {
	s.RepeatInterval = &v
	return s
}

func (s *IncidentResponsePlanForSNSViewRepeatNotifySetting) Validate() error {
	return dara.Validate(s)
}
