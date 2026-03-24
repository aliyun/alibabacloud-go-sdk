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
	// 自动恢复时间。
	//
	// example:
	//
	// 1741234567890
	AutoRecoverTime *int64 `json:"autoRecoverTime,omitempty" xml:"autoRecoverTime,omitempty"`
	// 内容。
	//
	// example:
	//
	// 检测到 user-service 实例 i-abc123 的 CPU 使用率达到 95%，已触发告警。
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 统计数量。
	//
	// example:
	//
	// 6
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// 维度。
	Dimension map[string]*string `json:"dimension,omitempty" xml:"dimension,omitempty"`
	// 用于分组维度。
	GroupBy map[string]*string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
	// 事件事件ID。
	//
	// example:
	//
	// event-001
	IncidentEventId *string `json:"incidentEventId,omitempty" xml:"incidentEventId,omitempty"`
	// 事件ID。
	//
	// example:
	//
	// incident-001
	IncidentId *string `json:"incidentId,omitempty" xml:"incidentId,omitempty"`
	// 最后时间。
	//
	// example:
	//
	// 1741234567890
	LastTime *int64 `json:"lastTime,omitempty" xml:"lastTime,omitempty"`
	// 手动恢复时间。
	//
	// example:
	//
	// 1741234567890
	RecoverTime *int64 `json:"recoverTime,omitempty" xml:"recoverTime,omitempty"`
	// 描述事件关联的资源信息。
	Resource map[string]*string `json:"resource,omitempty" xml:"resource,omitempty"`
	// 状态。
	//
	// example:
	//
	// Running
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 时间。
	//
	// example:
	//
	// 2025-03-11T08:21:58Z
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
	// 标题。
	//
	// example:
	//
	// 用户服务 CPU 使用率过高
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 用户ID。
	//
	// example:
	//
	// user-12345
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
