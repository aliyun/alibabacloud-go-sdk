// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncidentResponsePlanForView interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRecoverSeconds(v int64) *IncidentResponsePlanForView
	GetAutoRecoverSeconds() *int64
	SetCreateTime(v string) *IncidentResponsePlanForView
	GetCreateTime() *string
	SetDescription(v string) *IncidentResponsePlanForView
	GetDescription() *string
	SetEnabled(v bool) *IncidentResponsePlanForView
	GetEnabled() *bool
	SetEscalationId(v []*string) *IncidentResponsePlanForView
	GetEscalationId() []*string
	SetMode(v string) *IncidentResponsePlanForView
	GetMode() *string
	SetName(v string) *IncidentResponsePlanForView
	GetName() *string
	SetPushingSetting(v *PushingSetting) *IncidentResponsePlanForView
	GetPushingSetting() *PushingSetting
	SetRepeatNotifySetting(v *RepeatNotifySetting) *IncidentResponsePlanForView
	GetRepeatNotifySetting() *RepeatNotifySetting
	SetSource(v string) *IncidentResponsePlanForView
	GetSource() *string
	SetSyncFromType(v string) *IncidentResponsePlanForView
	GetSyncFromType() *string
	SetType(v string) *IncidentResponsePlanForView
	GetType() *string
	SetUpdateTime(v string) *IncidentResponsePlanForView
	GetUpdateTime() *string
	SetUuid(v string) *IncidentResponsePlanForView
	GetUuid() *string
	SetWorkspace(v string) *IncidentResponsePlanForView
	GetWorkspace() *string
}

type IncidentResponsePlanForView struct {
	AutoRecoverSeconds  *int64               `json:"autoRecoverSeconds,omitempty" xml:"autoRecoverSeconds,omitempty"`
	CreateTime          *string              `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description         *string              `json:"description,omitempty" xml:"description,omitempty"`
	Enabled             *bool                `json:"enabled,omitempty" xml:"enabled,omitempty"`
	EscalationId        []*string            `json:"escalationId,omitempty" xml:"escalationId,omitempty" type:"Repeated"`
	Mode                *string              `json:"mode,omitempty" xml:"mode,omitempty"`
	Name                *string              `json:"name,omitempty" xml:"name,omitempty"`
	PushingSetting      *PushingSetting      `json:"pushingSetting,omitempty" xml:"pushingSetting,omitempty"`
	RepeatNotifySetting *RepeatNotifySetting `json:"repeatNotifySetting,omitempty" xml:"repeatNotifySetting,omitempty"`
	Source              *string              `json:"source,omitempty" xml:"source,omitempty"`
	SyncFromType        *string              `json:"syncFromType,omitempty" xml:"syncFromType,omitempty"`
	Type                *string              `json:"type,omitempty" xml:"type,omitempty"`
	UpdateTime          *string              `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	Uuid                *string              `json:"uuid,omitempty" xml:"uuid,omitempty"`
	Workspace           *string              `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s IncidentResponsePlanForView) String() string {
	return dara.Prettify(s)
}

func (s IncidentResponsePlanForView) GoString() string {
	return s.String()
}

func (s *IncidentResponsePlanForView) GetAutoRecoverSeconds() *int64 {
	return s.AutoRecoverSeconds
}

func (s *IncidentResponsePlanForView) GetCreateTime() *string {
	return s.CreateTime
}

func (s *IncidentResponsePlanForView) GetDescription() *string {
	return s.Description
}

func (s *IncidentResponsePlanForView) GetEnabled() *bool {
	return s.Enabled
}

func (s *IncidentResponsePlanForView) GetEscalationId() []*string {
	return s.EscalationId
}

func (s *IncidentResponsePlanForView) GetMode() *string {
	return s.Mode
}

func (s *IncidentResponsePlanForView) GetName() *string {
	return s.Name
}

func (s *IncidentResponsePlanForView) GetPushingSetting() *PushingSetting {
	return s.PushingSetting
}

func (s *IncidentResponsePlanForView) GetRepeatNotifySetting() *RepeatNotifySetting {
	return s.RepeatNotifySetting
}

func (s *IncidentResponsePlanForView) GetSource() *string {
	return s.Source
}

func (s *IncidentResponsePlanForView) GetSyncFromType() *string {
	return s.SyncFromType
}

func (s *IncidentResponsePlanForView) GetType() *string {
	return s.Type
}

func (s *IncidentResponsePlanForView) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *IncidentResponsePlanForView) GetUuid() *string {
	return s.Uuid
}

func (s *IncidentResponsePlanForView) GetWorkspace() *string {
	return s.Workspace
}

func (s *IncidentResponsePlanForView) SetAutoRecoverSeconds(v int64) *IncidentResponsePlanForView {
	s.AutoRecoverSeconds = &v
	return s
}

func (s *IncidentResponsePlanForView) SetCreateTime(v string) *IncidentResponsePlanForView {
	s.CreateTime = &v
	return s
}

func (s *IncidentResponsePlanForView) SetDescription(v string) *IncidentResponsePlanForView {
	s.Description = &v
	return s
}

func (s *IncidentResponsePlanForView) SetEnabled(v bool) *IncidentResponsePlanForView {
	s.Enabled = &v
	return s
}

func (s *IncidentResponsePlanForView) SetEscalationId(v []*string) *IncidentResponsePlanForView {
	s.EscalationId = v
	return s
}

func (s *IncidentResponsePlanForView) SetMode(v string) *IncidentResponsePlanForView {
	s.Mode = &v
	return s
}

func (s *IncidentResponsePlanForView) SetName(v string) *IncidentResponsePlanForView {
	s.Name = &v
	return s
}

func (s *IncidentResponsePlanForView) SetPushingSetting(v *PushingSetting) *IncidentResponsePlanForView {
	s.PushingSetting = v
	return s
}

func (s *IncidentResponsePlanForView) SetRepeatNotifySetting(v *RepeatNotifySetting) *IncidentResponsePlanForView {
	s.RepeatNotifySetting = v
	return s
}

func (s *IncidentResponsePlanForView) SetSource(v string) *IncidentResponsePlanForView {
	s.Source = &v
	return s
}

func (s *IncidentResponsePlanForView) SetSyncFromType(v string) *IncidentResponsePlanForView {
	s.SyncFromType = &v
	return s
}

func (s *IncidentResponsePlanForView) SetType(v string) *IncidentResponsePlanForView {
	s.Type = &v
	return s
}

func (s *IncidentResponsePlanForView) SetUpdateTime(v string) *IncidentResponsePlanForView {
	s.UpdateTime = &v
	return s
}

func (s *IncidentResponsePlanForView) SetUuid(v string) *IncidentResponsePlanForView {
	s.Uuid = &v
	return s
}

func (s *IncidentResponsePlanForView) SetWorkspace(v string) *IncidentResponsePlanForView {
	s.Workspace = &v
	return s
}

func (s *IncidentResponsePlanForView) Validate() error {
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
