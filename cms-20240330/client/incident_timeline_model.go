// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncidentTimeline interface {
	dara.Model
	String() string
	GoString() string
	SetChildType(v string) *IncidentTimeline
	GetChildType() *string
	SetContent(v string) *IncidentTimeline
	GetContent() *string
	SetIncidentId(v string) *IncidentTimeline
	GetIncidentId() *string
	SetIncidentTimelineId(v string) *IncidentTimeline
	GetIncidentTimelineId() *string
	SetTime(v int64) *IncidentTimeline
	GetTime() *int64
	SetTimelineId(v string) *IncidentTimeline
	GetTimelineId() *string
	SetTitle(v string) *IncidentTimeline
	GetTitle() *string
	SetType(v string) *IncidentTimeline
	GetType() *string
	SetUserId(v string) *IncidentTimeline
	GetUserId() *string
}

type IncidentTimeline struct {
	// Subtype.
	//
	// example:
	//
	// comment
	ChildType *string `json:"childType,omitempty" xml:"childType,omitempty"`
	// Content.
	//
	// example:
	//
	// The operations team has been notified to investigate.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// Incident ID.
	//
	// example:
	//
	// incident-001
	IncidentId *string `json:"incidentId,omitempty" xml:"incidentId,omitempty"`
	// Timeline node unique ID.
	//
	// example:
	//
	// timeline-001
	IncidentTimelineId *string `json:"incidentTimelineId,omitempty" xml:"incidentTimelineId,omitempty"`
	// Time.
	//
	// example:
	//
	// 1741234567890
	Time *int64 `json:"time,omitempty" xml:"time,omitempty"`
	// Parent timeline ID.
	//
	// example:
	//
	// timeline-main-001
	TimelineId *string `json:"timelineId,omitempty" xml:"timelineId,omitempty"`
	// Title.
	//
	// example:
	//
	// The operations team has been contacted.
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// Type.
	//
	// example:
	//
	// ACTION
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// User ID.
	//
	// example:
	//
	// user-12345
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s IncidentTimeline) String() string {
	return dara.Prettify(s)
}

func (s IncidentTimeline) GoString() string {
	return s.String()
}

func (s *IncidentTimeline) GetChildType() *string {
	return s.ChildType
}

func (s *IncidentTimeline) GetContent() *string {
	return s.Content
}

func (s *IncidentTimeline) GetIncidentId() *string {
	return s.IncidentId
}

func (s *IncidentTimeline) GetIncidentTimelineId() *string {
	return s.IncidentTimelineId
}

func (s *IncidentTimeline) GetTime() *int64 {
	return s.Time
}

func (s *IncidentTimeline) GetTimelineId() *string {
	return s.TimelineId
}

func (s *IncidentTimeline) GetTitle() *string {
	return s.Title
}

func (s *IncidentTimeline) GetType() *string {
	return s.Type
}

func (s *IncidentTimeline) GetUserId() *string {
	return s.UserId
}

func (s *IncidentTimeline) SetChildType(v string) *IncidentTimeline {
	s.ChildType = &v
	return s
}

func (s *IncidentTimeline) SetContent(v string) *IncidentTimeline {
	s.Content = &v
	return s
}

func (s *IncidentTimeline) SetIncidentId(v string) *IncidentTimeline {
	s.IncidentId = &v
	return s
}

func (s *IncidentTimeline) SetIncidentTimelineId(v string) *IncidentTimeline {
	s.IncidentTimelineId = &v
	return s
}

func (s *IncidentTimeline) SetTime(v int64) *IncidentTimeline {
	s.Time = &v
	return s
}

func (s *IncidentTimeline) SetTimelineId(v string) *IncidentTimeline {
	s.TimelineId = &v
	return s
}

func (s *IncidentTimeline) SetTitle(v string) *IncidentTimeline {
	s.Title = &v
	return s
}

func (s *IncidentTimeline) SetType(v string) *IncidentTimeline {
	s.Type = &v
	return s
}

func (s *IncidentTimeline) SetUserId(v string) *IncidentTimeline {
	s.UserId = &v
	return s
}

func (s *IncidentTimeline) Validate() error {
	return dara.Validate(s)
}
