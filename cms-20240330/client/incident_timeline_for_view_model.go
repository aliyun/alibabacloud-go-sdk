// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncidentTimelineForView interface {
	dara.Model
	String() string
	GoString() string
	SetChangeType(v string) *IncidentTimelineForView
	GetChangeType() *string
	SetContent(v interface{}) *IncidentTimelineForView
	GetContent() interface{}
	SetIncidentUuid(v string) *IncidentTimelineForView
	GetIncidentUuid() *string
	SetOperator(v *ContactForIncidentView) *IncidentTimelineForView
	GetOperator() *ContactForIncidentView
	SetTime(v int64) *IncidentTimelineForView
	GetTime() *int64
	SetTimelineUuid(v string) *IncidentTimelineForView
	GetTimelineUuid() *string
	SetTitle(v string) *IncidentTimelineForView
	GetTitle() *string
	SetType(v string) *IncidentTimelineForView
	GetType() *string
	SetUserId(v string) *IncidentTimelineForView
	GetUserId() *string
	SetWorkspace(v string) *IncidentTimelineForView
	GetWorkspace() *string
}

type IncidentTimelineForView struct {
	ChangeType   *string                 `json:"changeType,omitempty" xml:"changeType,omitempty"`
	Content      interface{}             `json:"content,omitempty" xml:"content,omitempty"`
	IncidentUuid *string                 `json:"incidentUuid,omitempty" xml:"incidentUuid,omitempty"`
	Operator     *ContactForIncidentView `json:"operator,omitempty" xml:"operator,omitempty"`
	Time         *int64                  `json:"time,omitempty" xml:"time,omitempty"`
	TimelineUuid *string                 `json:"timelineUuid,omitempty" xml:"timelineUuid,omitempty"`
	Title        *string                 `json:"title,omitempty" xml:"title,omitempty"`
	Type         *string                 `json:"type,omitempty" xml:"type,omitempty"`
	UserId       *string                 `json:"userId,omitempty" xml:"userId,omitempty"`
	Workspace    *string                 `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s IncidentTimelineForView) String() string {
	return dara.Prettify(s)
}

func (s IncidentTimelineForView) GoString() string {
	return s.String()
}

func (s *IncidentTimelineForView) GetChangeType() *string {
	return s.ChangeType
}

func (s *IncidentTimelineForView) GetContent() interface{} {
	return s.Content
}

func (s *IncidentTimelineForView) GetIncidentUuid() *string {
	return s.IncidentUuid
}

func (s *IncidentTimelineForView) GetOperator() *ContactForIncidentView {
	return s.Operator
}

func (s *IncidentTimelineForView) GetTime() *int64 {
	return s.Time
}

func (s *IncidentTimelineForView) GetTimelineUuid() *string {
	return s.TimelineUuid
}

func (s *IncidentTimelineForView) GetTitle() *string {
	return s.Title
}

func (s *IncidentTimelineForView) GetType() *string {
	return s.Type
}

func (s *IncidentTimelineForView) GetUserId() *string {
	return s.UserId
}

func (s *IncidentTimelineForView) GetWorkspace() *string {
	return s.Workspace
}

func (s *IncidentTimelineForView) SetChangeType(v string) *IncidentTimelineForView {
	s.ChangeType = &v
	return s
}

func (s *IncidentTimelineForView) SetContent(v interface{}) *IncidentTimelineForView {
	s.Content = v
	return s
}

func (s *IncidentTimelineForView) SetIncidentUuid(v string) *IncidentTimelineForView {
	s.IncidentUuid = &v
	return s
}

func (s *IncidentTimelineForView) SetOperator(v *ContactForIncidentView) *IncidentTimelineForView {
	s.Operator = v
	return s
}

func (s *IncidentTimelineForView) SetTime(v int64) *IncidentTimelineForView {
	s.Time = &v
	return s
}

func (s *IncidentTimelineForView) SetTimelineUuid(v string) *IncidentTimelineForView {
	s.TimelineUuid = &v
	return s
}

func (s *IncidentTimelineForView) SetTitle(v string) *IncidentTimelineForView {
	s.Title = &v
	return s
}

func (s *IncidentTimelineForView) SetType(v string) *IncidentTimelineForView {
	s.Type = &v
	return s
}

func (s *IncidentTimelineForView) SetUserId(v string) *IncidentTimelineForView {
	s.UserId = &v
	return s
}

func (s *IncidentTimelineForView) SetWorkspace(v string) *IncidentTimelineForView {
	s.Workspace = &v
	return s
}

func (s *IncidentTimelineForView) Validate() error {
	if s.Operator != nil {
		if err := s.Operator.Validate(); err != nil {
			return err
		}
	}
	return nil
}
