// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncidentStruct interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *IncidentStruct
	GetContent() *string
	SetEscalations(v []*IncidentEscalationStruct) *IncidentStruct
	GetEscalations() []*IncidentEscalationStruct
	SetIncidentId(v string) *IncidentStruct
	GetIncidentId() *string
	SetIncidentPlan(v *IncidentPlanStruct) *IncidentStruct
	GetIncidentPlan() *IncidentPlanStruct
	SetResource(v *IncidentResourceDetail) *IncidentStruct
	GetResource() *IncidentResourceDetail
	SetSeverity(v string) *IncidentStruct
	GetSeverity() *string
	SetStatus(v string) *IncidentStruct
	GetStatus() *string
	SetTime(v int64) *IncidentStruct
	GetTime() *int64
	SetTitle(v string) *IncidentStruct
	GetTitle() *string
	SetUserId(v string) *IncidentStruct
	GetUserId() *string
}

type IncidentStruct struct {
	// 事件内容。
	//
	// example:
	//
	// 检测到 RDS 实例 rds-bp1234567890abcdef 的连接数达到 1000，已触发告警。
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 升级策略列表。
	Escalations []*IncidentEscalationStruct `json:"escalations,omitempty" xml:"escalations,omitempty" type:"Repeated"`
	// 事件唯一 ID。
	//
	// example:
	//
	// incident-001
	IncidentId *string `json:"incidentId,omitempty" xml:"incidentId,omitempty"`
	// 关联的事件预案。
	IncidentPlan *IncidentPlanStruct `json:"incidentPlan,omitempty" xml:"incidentPlan,omitempty"`
	// 关联的资源详情。
	Resource *IncidentResourceDetail `json:"resource,omitempty" xml:"resource,omitempty"`
	// 事件严重等级。
	//
	// example:
	//
	// P1
	Severity *string `json:"severity,omitempty" xml:"severity,omitempty"`
	// 当前状态。
	//
	// example:
	//
	// OPEN
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 时间戳
	//
	// example:
	//
	// 1741234567890
	Time *int64 `json:"time,omitempty" xml:"time,omitempty"`
	// 事件标题。
	//
	// example:
	//
	// 数据库连接数过高
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 用户 ID。
	//
	// example:
	//
	// user-12345
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s IncidentStruct) String() string {
	return dara.Prettify(s)
}

func (s IncidentStruct) GoString() string {
	return s.String()
}

func (s *IncidentStruct) GetContent() *string {
	return s.Content
}

func (s *IncidentStruct) GetEscalations() []*IncidentEscalationStruct {
	return s.Escalations
}

func (s *IncidentStruct) GetIncidentId() *string {
	return s.IncidentId
}

func (s *IncidentStruct) GetIncidentPlan() *IncidentPlanStruct {
	return s.IncidentPlan
}

func (s *IncidentStruct) GetResource() *IncidentResourceDetail {
	return s.Resource
}

func (s *IncidentStruct) GetSeverity() *string {
	return s.Severity
}

func (s *IncidentStruct) GetStatus() *string {
	return s.Status
}

func (s *IncidentStruct) GetTime() *int64 {
	return s.Time
}

func (s *IncidentStruct) GetTitle() *string {
	return s.Title
}

func (s *IncidentStruct) GetUserId() *string {
	return s.UserId
}

func (s *IncidentStruct) SetContent(v string) *IncidentStruct {
	s.Content = &v
	return s
}

func (s *IncidentStruct) SetEscalations(v []*IncidentEscalationStruct) *IncidentStruct {
	s.Escalations = v
	return s
}

func (s *IncidentStruct) SetIncidentId(v string) *IncidentStruct {
	s.IncidentId = &v
	return s
}

func (s *IncidentStruct) SetIncidentPlan(v *IncidentPlanStruct) *IncidentStruct {
	s.IncidentPlan = v
	return s
}

func (s *IncidentStruct) SetResource(v *IncidentResourceDetail) *IncidentStruct {
	s.Resource = v
	return s
}

func (s *IncidentStruct) SetSeverity(v string) *IncidentStruct {
	s.Severity = &v
	return s
}

func (s *IncidentStruct) SetStatus(v string) *IncidentStruct {
	s.Status = &v
	return s
}

func (s *IncidentStruct) SetTime(v int64) *IncidentStruct {
	s.Time = &v
	return s
}

func (s *IncidentStruct) SetTitle(v string) *IncidentStruct {
	s.Title = &v
	return s
}

func (s *IncidentStruct) SetUserId(v string) *IncidentStruct {
	s.UserId = &v
	return s
}

func (s *IncidentStruct) Validate() error {
	if s.Escalations != nil {
		for _, item := range s.Escalations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.IncidentPlan != nil {
		if err := s.IncidentPlan.Validate(); err != nil {
			return err
		}
	}
	if s.Resource != nil {
		if err := s.Resource.Validate(); err != nil {
			return err
		}
	}
	return nil
}
