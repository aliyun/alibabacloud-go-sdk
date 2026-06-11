// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncidentResponsePlanForModify interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRecoverSeconds(v int64) *IncidentResponsePlanForModify
	GetAutoRecoverSeconds() *int64
	SetDescription(v string) *IncidentResponsePlanForModify
	GetDescription() *string
	SetEnabled(v bool) *IncidentResponsePlanForModify
	GetEnabled() *bool
	SetEscalationId(v []*string) *IncidentResponsePlanForModify
	GetEscalationId() []*string
	SetMode(v string) *IncidentResponsePlanForModify
	GetMode() *string
	SetName(v string) *IncidentResponsePlanForModify
	GetName() *string
	SetPushingSetting(v *PushingSetting) *IncidentResponsePlanForModify
	GetPushingSetting() *PushingSetting
	SetRepeatNotifySetting(v *RepeatNotifySetting) *IncidentResponsePlanForModify
	GetRepeatNotifySetting() *RepeatNotifySetting
	SetSource(v string) *IncidentResponsePlanForModify
	GetSource() *string
	SetSyncFromType(v string) *IncidentResponsePlanForModify
	GetSyncFromType() *string
	SetType(v string) *IncidentResponsePlanForModify
	GetType() *string
	SetUuid(v string) *IncidentResponsePlanForModify
	GetUuid() *string
}

type IncidentResponsePlanForModify struct {
	AutoRecoverSeconds  *int64               `json:"autoRecoverSeconds,omitempty" xml:"autoRecoverSeconds,omitempty"`
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
	// Required for update operations. Optional for create operations; if omitted, the backend generates one. `NotifyStrategy` also uses this UUID.
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s IncidentResponsePlanForModify) String() string {
	return dara.Prettify(s)
}

func (s IncidentResponsePlanForModify) GoString() string {
	return s.String()
}

func (s *IncidentResponsePlanForModify) GetAutoRecoverSeconds() *int64 {
	return s.AutoRecoverSeconds
}

func (s *IncidentResponsePlanForModify) GetDescription() *string {
	return s.Description
}

func (s *IncidentResponsePlanForModify) GetEnabled() *bool {
	return s.Enabled
}

func (s *IncidentResponsePlanForModify) GetEscalationId() []*string {
	return s.EscalationId
}

func (s *IncidentResponsePlanForModify) GetMode() *string {
	return s.Mode
}

func (s *IncidentResponsePlanForModify) GetName() *string {
	return s.Name
}

func (s *IncidentResponsePlanForModify) GetPushingSetting() *PushingSetting {
	return s.PushingSetting
}

func (s *IncidentResponsePlanForModify) GetRepeatNotifySetting() *RepeatNotifySetting {
	return s.RepeatNotifySetting
}

func (s *IncidentResponsePlanForModify) GetSource() *string {
	return s.Source
}

func (s *IncidentResponsePlanForModify) GetSyncFromType() *string {
	return s.SyncFromType
}

func (s *IncidentResponsePlanForModify) GetType() *string {
	return s.Type
}

func (s *IncidentResponsePlanForModify) GetUuid() *string {
	return s.Uuid
}

func (s *IncidentResponsePlanForModify) SetAutoRecoverSeconds(v int64) *IncidentResponsePlanForModify {
	s.AutoRecoverSeconds = &v
	return s
}

func (s *IncidentResponsePlanForModify) SetDescription(v string) *IncidentResponsePlanForModify {
	s.Description = &v
	return s
}

func (s *IncidentResponsePlanForModify) SetEnabled(v bool) *IncidentResponsePlanForModify {
	s.Enabled = &v
	return s
}

func (s *IncidentResponsePlanForModify) SetEscalationId(v []*string) *IncidentResponsePlanForModify {
	s.EscalationId = v
	return s
}

func (s *IncidentResponsePlanForModify) SetMode(v string) *IncidentResponsePlanForModify {
	s.Mode = &v
	return s
}

func (s *IncidentResponsePlanForModify) SetName(v string) *IncidentResponsePlanForModify {
	s.Name = &v
	return s
}

func (s *IncidentResponsePlanForModify) SetPushingSetting(v *PushingSetting) *IncidentResponsePlanForModify {
	s.PushingSetting = v
	return s
}

func (s *IncidentResponsePlanForModify) SetRepeatNotifySetting(v *RepeatNotifySetting) *IncidentResponsePlanForModify {
	s.RepeatNotifySetting = v
	return s
}

func (s *IncidentResponsePlanForModify) SetSource(v string) *IncidentResponsePlanForModify {
	s.Source = &v
	return s
}

func (s *IncidentResponsePlanForModify) SetSyncFromType(v string) *IncidentResponsePlanForModify {
	s.SyncFromType = &v
	return s
}

func (s *IncidentResponsePlanForModify) SetType(v string) *IncidentResponsePlanForModify {
	s.Type = &v
	return s
}

func (s *IncidentResponsePlanForModify) SetUuid(v string) *IncidentResponsePlanForModify {
	s.Uuid = &v
	return s
}

func (s *IncidentResponsePlanForModify) Validate() error {
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
