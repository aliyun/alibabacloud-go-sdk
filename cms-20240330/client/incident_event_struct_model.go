// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncidentEventStruct interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRecoverTime(v int64) *IncidentEventStruct
	GetAutoRecoverTime() *int64
	SetContent(v string) *IncidentEventStruct
	GetContent() *string
	SetCount(v int32) *IncidentEventStruct
	GetCount() *int32
	SetDimension(v map[string]*string) *IncidentEventStruct
	GetDimension() map[string]*string
	SetGroupBy(v map[string]*string) *IncidentEventStruct
	GetGroupBy() map[string]*string
	SetIncidentEventId(v string) *IncidentEventStruct
	GetIncidentEventId() *string
	SetIncidentId(v string) *IncidentEventStruct
	GetIncidentId() *string
	SetLastTime(v int64) *IncidentEventStruct
	GetLastTime() *int64
	SetRecoverTime(v int64) *IncidentEventStruct
	GetRecoverTime() *int64
	SetResource(v map[string]*string) *IncidentEventStruct
	GetResource() map[string]*string
	SetStatus(v int64) *IncidentEventStruct
	GetStatus() *int64
	SetTime(v string) *IncidentEventStruct
	GetTime() *string
	SetTitle(v string) *IncidentEventStruct
	GetTitle() *string
	SetUserId(v string) *IncidentEventStruct
	GetUserId() *string
}

type IncidentEventStruct struct {
	AutoRecoverTime *int64             `json:"autoRecoverTime,omitempty" xml:"autoRecoverTime,omitempty"`
	Content         *string            `json:"content,omitempty" xml:"content,omitempty"`
	Count           *int32             `json:"count,omitempty" xml:"count,omitempty"`
	Dimension       map[string]*string `json:"dimension,omitempty" xml:"dimension,omitempty"`
	GroupBy         map[string]*string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
	IncidentEventId *string            `json:"incidentEventId,omitempty" xml:"incidentEventId,omitempty"`
	IncidentId      *string            `json:"incidentId,omitempty" xml:"incidentId,omitempty"`
	LastTime        *int64             `json:"lastTime,omitempty" xml:"lastTime,omitempty"`
	RecoverTime     *int64             `json:"recoverTime,omitempty" xml:"recoverTime,omitempty"`
	Resource        map[string]*string `json:"resource,omitempty" xml:"resource,omitempty"`
	Status          *int64             `json:"status,omitempty" xml:"status,omitempty"`
	Time            *string            `json:"time,omitempty" xml:"time,omitempty"`
	Title           *string            `json:"title,omitempty" xml:"title,omitempty"`
	UserId          *string            `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s IncidentEventStruct) String() string {
	return dara.Prettify(s)
}

func (s IncidentEventStruct) GoString() string {
	return s.String()
}

func (s *IncidentEventStruct) GetAutoRecoverTime() *int64 {
	return s.AutoRecoverTime
}

func (s *IncidentEventStruct) GetContent() *string {
	return s.Content
}

func (s *IncidentEventStruct) GetCount() *int32 {
	return s.Count
}

func (s *IncidentEventStruct) GetDimension() map[string]*string {
	return s.Dimension
}

func (s *IncidentEventStruct) GetGroupBy() map[string]*string {
	return s.GroupBy
}

func (s *IncidentEventStruct) GetIncidentEventId() *string {
	return s.IncidentEventId
}

func (s *IncidentEventStruct) GetIncidentId() *string {
	return s.IncidentId
}

func (s *IncidentEventStruct) GetLastTime() *int64 {
	return s.LastTime
}

func (s *IncidentEventStruct) GetRecoverTime() *int64 {
	return s.RecoverTime
}

func (s *IncidentEventStruct) GetResource() map[string]*string {
	return s.Resource
}

func (s *IncidentEventStruct) GetStatus() *int64 {
	return s.Status
}

func (s *IncidentEventStruct) GetTime() *string {
	return s.Time
}

func (s *IncidentEventStruct) GetTitle() *string {
	return s.Title
}

func (s *IncidentEventStruct) GetUserId() *string {
	return s.UserId
}

func (s *IncidentEventStruct) SetAutoRecoverTime(v int64) *IncidentEventStruct {
	s.AutoRecoverTime = &v
	return s
}

func (s *IncidentEventStruct) SetContent(v string) *IncidentEventStruct {
	s.Content = &v
	return s
}

func (s *IncidentEventStruct) SetCount(v int32) *IncidentEventStruct {
	s.Count = &v
	return s
}

func (s *IncidentEventStruct) SetDimension(v map[string]*string) *IncidentEventStruct {
	s.Dimension = v
	return s
}

func (s *IncidentEventStruct) SetGroupBy(v map[string]*string) *IncidentEventStruct {
	s.GroupBy = v
	return s
}

func (s *IncidentEventStruct) SetIncidentEventId(v string) *IncidentEventStruct {
	s.IncidentEventId = &v
	return s
}

func (s *IncidentEventStruct) SetIncidentId(v string) *IncidentEventStruct {
	s.IncidentId = &v
	return s
}

func (s *IncidentEventStruct) SetLastTime(v int64) *IncidentEventStruct {
	s.LastTime = &v
	return s
}

func (s *IncidentEventStruct) SetRecoverTime(v int64) *IncidentEventStruct {
	s.RecoverTime = &v
	return s
}

func (s *IncidentEventStruct) SetResource(v map[string]*string) *IncidentEventStruct {
	s.Resource = v
	return s
}

func (s *IncidentEventStruct) SetStatus(v int64) *IncidentEventStruct {
	s.Status = &v
	return s
}

func (s *IncidentEventStruct) SetTime(v string) *IncidentEventStruct {
	s.Time = &v
	return s
}

func (s *IncidentEventStruct) SetTitle(v string) *IncidentEventStruct {
	s.Title = &v
	return s
}

func (s *IncidentEventStruct) SetUserId(v string) *IncidentEventStruct {
	s.UserId = &v
	return s
}

func (s *IncidentEventStruct) Validate() error {
	return dara.Validate(s)
}
