// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncidentEventForView interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRecoverTime(v int64) *IncidentEventForView
	GetAutoRecoverTime() *int64
	SetContent(v *CmsEventForView) *IncidentEventForView
	GetContent() *CmsEventForView
	SetCount(v int32) *IncidentEventForView
	GetCount() *int32
	SetDimension(v map[string]interface{}) *IncidentEventForView
	GetDimension() map[string]interface{}
	SetEventResource(v *EventResourceForIncidentView) *IncidentEventForView
	GetEventResource() *EventResourceForIncidentView
	SetGroupBy(v map[string]interface{}) *IncidentEventForView
	GetGroupBy() map[string]interface{}
	SetIncidentEventUuid(v string) *IncidentEventForView
	GetIncidentEventUuid() *string
	SetIncidentUuid(v string) *IncidentEventForView
	GetIncidentUuid() *string
	SetLabels(v map[string]interface{}) *IncidentEventForView
	GetLabels() map[string]interface{}
	SetLastTime(v int64) *IncidentEventForView
	GetLastTime() *int64
	SetRecoverTime(v int64) *IncidentEventForView
	GetRecoverTime() *int64
	SetSearchIndex(v []*string) *IncidentEventForView
	GetSearchIndex() []*string
	SetSeverity(v string) *IncidentEventForView
	GetSeverity() *string
	SetSeverityCountMap(v map[string]interface{}) *IncidentEventForView
	GetSeverityCountMap() map[string]interface{}
	SetState(v int32) *IncidentEventForView
	GetState() *int32
	SetTextIndex(v string) *IncidentEventForView
	GetTextIndex() *string
	SetTime(v int64) *IncidentEventForView
	GetTime() *int64
	SetTitle(v string) *IncidentEventForView
	GetTitle() *string
	SetUserId(v string) *IncidentEventForView
	GetUserId() *string
	SetWorkspace(v string) *IncidentEventForView
	GetWorkspace() *string
}

type IncidentEventForView struct {
	// example:
	//
	// 1743876600000
	AutoRecoverTime *int64           `json:"autoRecoverTime,omitempty" xml:"autoRecoverTime,omitempty"`
	Content         *CmsEventForView `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 3
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// example:
	//
	// { "env": "prod", "module": "payment" }
	Dimension     map[string]interface{}        `json:"dimension,omitempty" xml:"dimension,omitempty"`
	EventResource *EventResourceForIncidentView `json:"eventResource,omitempty" xml:"eventResource,omitempty"`
	// example:
	//
	// { "resourceId": "i-bp123456", "severity": "Critical" }
	GroupBy map[string]interface{} `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
	// example:
	//
	// "x1y2z3a4-b5c6-d7e8-f9g0-h1i2j3k4l5m6"
	IncidentEventUuid *string `json:"incidentEventUuid,omitempty" xml:"incidentEventUuid,omitempty"`
	// example:
	//
	// "a1b2c3d4-e5f6-7890-1234-567890abcdef"
	IncidentUuid *string `json:"incidentUuid,omitempty" xml:"incidentUuid,omitempty"`
	// example:
	//
	// { "project": "payment-gateway", "owner": "ops-team" }
	Labels map[string]interface{} `json:"labels,omitempty" xml:"labels,omitempty"`
	// example:
	//
	// 1743876000000
	LastTime *int64 `json:"lastTime,omitempty" xml:"lastTime,omitempty"`
	// example:
	//
	// 1743876600000
	RecoverTime *int64    `json:"recoverTime,omitempty" xml:"recoverTime,omitempty"`
	SearchIndex []*string `json:"searchIndex,omitempty" xml:"searchIndex,omitempty" type:"Repeated"`
	// example:
	//
	// "Critical"
	Severity *string `json:"severity,omitempty" xml:"severity,omitempty"`
	// example:
	//
	// { "Critical": 2, "High": 1 }
	SeverityCountMap map[string]interface{} `json:"severityCountMap,omitempty" xml:"severityCountMap,omitempty"`
	// example:
	//
	// 1
	State     *int32  `json:"state,omitempty" xml:"state,omitempty"`
	TextIndex *string `json:"textIndex,omitempty" xml:"textIndex,omitempty"`
	// example:
	//
	// 1743876000000
	Time  *int64  `json:"time,omitempty" xml:"time,omitempty"`
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// "user-abc123"
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// example:
	//
	// "ws-xyz789"
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s IncidentEventForView) String() string {
	return dara.Prettify(s)
}

func (s IncidentEventForView) GoString() string {
	return s.String()
}

func (s *IncidentEventForView) GetAutoRecoverTime() *int64 {
	return s.AutoRecoverTime
}

func (s *IncidentEventForView) GetContent() *CmsEventForView {
	return s.Content
}

func (s *IncidentEventForView) GetCount() *int32 {
	return s.Count
}

func (s *IncidentEventForView) GetDimension() map[string]interface{} {
	return s.Dimension
}

func (s *IncidentEventForView) GetEventResource() *EventResourceForIncidentView {
	return s.EventResource
}

func (s *IncidentEventForView) GetGroupBy() map[string]interface{} {
	return s.GroupBy
}

func (s *IncidentEventForView) GetIncidentEventUuid() *string {
	return s.IncidentEventUuid
}

func (s *IncidentEventForView) GetIncidentUuid() *string {
	return s.IncidentUuid
}

func (s *IncidentEventForView) GetLabels() map[string]interface{} {
	return s.Labels
}

func (s *IncidentEventForView) GetLastTime() *int64 {
	return s.LastTime
}

func (s *IncidentEventForView) GetRecoverTime() *int64 {
	return s.RecoverTime
}

func (s *IncidentEventForView) GetSearchIndex() []*string {
	return s.SearchIndex
}

func (s *IncidentEventForView) GetSeverity() *string {
	return s.Severity
}

func (s *IncidentEventForView) GetSeverityCountMap() map[string]interface{} {
	return s.SeverityCountMap
}

func (s *IncidentEventForView) GetState() *int32 {
	return s.State
}

func (s *IncidentEventForView) GetTextIndex() *string {
	return s.TextIndex
}

func (s *IncidentEventForView) GetTime() *int64 {
	return s.Time
}

func (s *IncidentEventForView) GetTitle() *string {
	return s.Title
}

func (s *IncidentEventForView) GetUserId() *string {
	return s.UserId
}

func (s *IncidentEventForView) GetWorkspace() *string {
	return s.Workspace
}

func (s *IncidentEventForView) SetAutoRecoverTime(v int64) *IncidentEventForView {
	s.AutoRecoverTime = &v
	return s
}

func (s *IncidentEventForView) SetContent(v *CmsEventForView) *IncidentEventForView {
	s.Content = v
	return s
}

func (s *IncidentEventForView) SetCount(v int32) *IncidentEventForView {
	s.Count = &v
	return s
}

func (s *IncidentEventForView) SetDimension(v map[string]interface{}) *IncidentEventForView {
	s.Dimension = v
	return s
}

func (s *IncidentEventForView) SetEventResource(v *EventResourceForIncidentView) *IncidentEventForView {
	s.EventResource = v
	return s
}

func (s *IncidentEventForView) SetGroupBy(v map[string]interface{}) *IncidentEventForView {
	s.GroupBy = v
	return s
}

func (s *IncidentEventForView) SetIncidentEventUuid(v string) *IncidentEventForView {
	s.IncidentEventUuid = &v
	return s
}

func (s *IncidentEventForView) SetIncidentUuid(v string) *IncidentEventForView {
	s.IncidentUuid = &v
	return s
}

func (s *IncidentEventForView) SetLabels(v map[string]interface{}) *IncidentEventForView {
	s.Labels = v
	return s
}

func (s *IncidentEventForView) SetLastTime(v int64) *IncidentEventForView {
	s.LastTime = &v
	return s
}

func (s *IncidentEventForView) SetRecoverTime(v int64) *IncidentEventForView {
	s.RecoverTime = &v
	return s
}

func (s *IncidentEventForView) SetSearchIndex(v []*string) *IncidentEventForView {
	s.SearchIndex = v
	return s
}

func (s *IncidentEventForView) SetSeverity(v string) *IncidentEventForView {
	s.Severity = &v
	return s
}

func (s *IncidentEventForView) SetSeverityCountMap(v map[string]interface{}) *IncidentEventForView {
	s.SeverityCountMap = v
	return s
}

func (s *IncidentEventForView) SetState(v int32) *IncidentEventForView {
	s.State = &v
	return s
}

func (s *IncidentEventForView) SetTextIndex(v string) *IncidentEventForView {
	s.TextIndex = &v
	return s
}

func (s *IncidentEventForView) SetTime(v int64) *IncidentEventForView {
	s.Time = &v
	return s
}

func (s *IncidentEventForView) SetTitle(v string) *IncidentEventForView {
	s.Title = &v
	return s
}

func (s *IncidentEventForView) SetUserId(v string) *IncidentEventForView {
	s.UserId = &v
	return s
}

func (s *IncidentEventForView) SetWorkspace(v string) *IncidentEventForView {
	s.Workspace = &v
	return s
}

func (s *IncidentEventForView) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	if s.EventResource != nil {
		if err := s.EventResource.Validate(); err != nil {
			return err
		}
	}
	return nil
}
