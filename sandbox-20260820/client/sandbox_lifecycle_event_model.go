// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSandboxLifecycleEvent interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *SandboxLifecycleEvent
	GetAction() *string
	SetBuildID(v string) *SandboxLifecycleEvent
	GetBuildID() *string
	SetDescription(v string) *SandboxLifecycleEvent
	GetDescription() *string
	SetEventData(v map[string]interface{}) *SandboxLifecycleEvent
	GetEventData() map[string]interface{}
	SetEventID(v string) *SandboxLifecycleEvent
	GetEventID() *string
	SetEventType(v string) *SandboxLifecycleEvent
	GetEventType() *string
	SetLevel(v string) *SandboxLifecycleEvent
	GetLevel() *string
	SetOccurredAt(v string) *SandboxLifecycleEvent
	GetOccurredAt() *string
	SetReason(v string) *SandboxLifecycleEvent
	GetReason() *string
	SetSandboxID(v string) *SandboxLifecycleEvent
	GetSandboxID() *string
	SetTeamID(v string) *SandboxLifecycleEvent
	GetTeamID() *string
	SetTemplateID(v string) *SandboxLifecycleEvent
	GetTemplateID() *string
}

type SandboxLifecycleEvent struct {
	// The action that triggered the event.
	//
	// example:
	//
	// delete
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// The template build ID. This field is not returned if there is no associated build.
	//
	// example:
	//
	// bld-example-001
	BuildID *string `json:"buildID,omitempty" xml:"buildID,omitempty"`
	// The event description.
	//
	// example:
	//
	// Sandbox deleted
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The event data as a JSON object. The specific fields vary depending on the event type.
	EventData map[string]interface{} `json:"eventData" xml:"eventData"`
	// The event ID.
	//
	// example:
	//
	// df3c6750-2c5e-50d1-8b81-7a2248f6ca3b
	EventID *string `json:"eventID,omitempty" xml:"eventID,omitempty"`
	// The event type.
	//
	// example:
	//
	// terminated
	EventType *string `json:"eventType,omitempty" xml:"eventType,omitempty"`
	// The event category. Valid values: INFO and WARNING.
	//
	// example:
	//
	// INFO
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
	// The time when the event occurred. The time is in UTC RFC3339 format and supports fractional seconds.
	//
	// example:
	//
	// 2026-10-08T06:18:38.123Z
	OccurredAt *string `json:"occurredAt,omitempty" xml:"occurredAt,omitempty"`
	// The reason for the event. This field is not returned if there is no reason.
	//
	// example:
	//
	// request
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// sbx-62105527-699e-48df-ad3b-7425001ae4f2
	SandboxID *string `json:"sandboxID,omitempty" xml:"sandboxID,omitempty"`
	// The ID of the team to which the event belongs.
	//
	// example:
	//
	// 70d1c834-0333-58d8-97ac-5336eb9179eb
	TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
	// The template ID.
	//
	// example:
	//
	// 4hrtmghh2nzrsdh6yr65
	TemplateID *string `json:"templateID,omitempty" xml:"templateID,omitempty"`
}

func (s SandboxLifecycleEvent) String() string {
	return dara.Prettify(s)
}

func (s SandboxLifecycleEvent) GoString() string {
	return s.String()
}

func (s *SandboxLifecycleEvent) GetAction() *string {
	return s.Action
}

func (s *SandboxLifecycleEvent) GetBuildID() *string {
	return s.BuildID
}

func (s *SandboxLifecycleEvent) GetDescription() *string {
	return s.Description
}

func (s *SandboxLifecycleEvent) GetEventData() map[string]interface{} {
	return s.EventData
}

func (s *SandboxLifecycleEvent) GetEventID() *string {
	return s.EventID
}

func (s *SandboxLifecycleEvent) GetEventType() *string {
	return s.EventType
}

func (s *SandboxLifecycleEvent) GetLevel() *string {
	return s.Level
}

func (s *SandboxLifecycleEvent) GetOccurredAt() *string {
	return s.OccurredAt
}

func (s *SandboxLifecycleEvent) GetReason() *string {
	return s.Reason
}

func (s *SandboxLifecycleEvent) GetSandboxID() *string {
	return s.SandboxID
}

func (s *SandboxLifecycleEvent) GetTeamID() *string {
	return s.TeamID
}

func (s *SandboxLifecycleEvent) GetTemplateID() *string {
	return s.TemplateID
}

func (s *SandboxLifecycleEvent) SetAction(v string) *SandboxLifecycleEvent {
	s.Action = &v
	return s
}

func (s *SandboxLifecycleEvent) SetBuildID(v string) *SandboxLifecycleEvent {
	s.BuildID = &v
	return s
}

func (s *SandboxLifecycleEvent) SetDescription(v string) *SandboxLifecycleEvent {
	s.Description = &v
	return s
}

func (s *SandboxLifecycleEvent) SetEventData(v map[string]interface{}) *SandboxLifecycleEvent {
	s.EventData = v
	return s
}

func (s *SandboxLifecycleEvent) SetEventID(v string) *SandboxLifecycleEvent {
	s.EventID = &v
	return s
}

func (s *SandboxLifecycleEvent) SetEventType(v string) *SandboxLifecycleEvent {
	s.EventType = &v
	return s
}

func (s *SandboxLifecycleEvent) SetLevel(v string) *SandboxLifecycleEvent {
	s.Level = &v
	return s
}

func (s *SandboxLifecycleEvent) SetOccurredAt(v string) *SandboxLifecycleEvent {
	s.OccurredAt = &v
	return s
}

func (s *SandboxLifecycleEvent) SetReason(v string) *SandboxLifecycleEvent {
	s.Reason = &v
	return s
}

func (s *SandboxLifecycleEvent) SetSandboxID(v string) *SandboxLifecycleEvent {
	s.SandboxID = &v
	return s
}

func (s *SandboxLifecycleEvent) SetTeamID(v string) *SandboxLifecycleEvent {
	s.TeamID = &v
	return s
}

func (s *SandboxLifecycleEvent) SetTemplateID(v string) *SandboxLifecycleEvent {
	s.TemplateID = &v
	return s
}

func (s *SandboxLifecycleEvent) Validate() error {
	return dara.Validate(s)
}
