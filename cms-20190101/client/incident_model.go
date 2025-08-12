// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncident interface {
	dara.Model
	String() string
	GoString() string
	SetActionTime(v int64) *Incident
	GetActionTime() *int64
	SetAlertCount(v int64) *Incident
	GetAlertCount() *int64
	SetEndTime(v int64) *Incident
	GetEndTime() *int64
	SetGroupingData(v map[string]interface{}) *Incident
	GetGroupingData() map[string]interface{}
	SetGroupingId(v string) *Incident
	GetGroupingId() *string
	SetGroupingKey(v string) *Incident
	GetGroupingKey() *string
	SetIncidentId(v string) *Incident
	GetIncidentId() *string
	SetIncidentStatus(v string) *Incident
	GetIncidentStatus() *string
	SetSeverity(v string) *Incident
	GetSeverity() *string
	SetStartTime(v int64) *Incident
	GetStartTime() *int64
	SetStrategyUuid(v string) *Incident
	GetStrategyUuid() *string
	SetUserId(v string) *Incident
	GetUserId() *string
}

type Incident struct {
	ActionTime     *int64                 `json:"ActionTime,omitempty" xml:"ActionTime,omitempty"`
	AlertCount     *int64                 `json:"AlertCount,omitempty" xml:"AlertCount,omitempty"`
	EndTime        *int64                 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	GroupingData   map[string]interface{} `json:"GroupingData,omitempty" xml:"GroupingData,omitempty"`
	GroupingId     *string                `json:"GroupingId,omitempty" xml:"GroupingId,omitempty"`
	GroupingKey    *string                `json:"GroupingKey,omitempty" xml:"GroupingKey,omitempty"`
	IncidentId     *string                `json:"IncidentId,omitempty" xml:"IncidentId,omitempty"`
	IncidentStatus *string                `json:"IncidentStatus,omitempty" xml:"IncidentStatus,omitempty"`
	Severity       *string                `json:"Severity,omitempty" xml:"Severity,omitempty"`
	StartTime      *int64                 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StrategyUuid   *string                `json:"StrategyUuid,omitempty" xml:"StrategyUuid,omitempty"`
	UserId         *string                `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s Incident) String() string {
	return dara.Prettify(s)
}

func (s Incident) GoString() string {
	return s.String()
}

func (s *Incident) GetActionTime() *int64 {
	return s.ActionTime
}

func (s *Incident) GetAlertCount() *int64 {
	return s.AlertCount
}

func (s *Incident) GetEndTime() *int64 {
	return s.EndTime
}

func (s *Incident) GetGroupingData() map[string]interface{} {
	return s.GroupingData
}

func (s *Incident) GetGroupingId() *string {
	return s.GroupingId
}

func (s *Incident) GetGroupingKey() *string {
	return s.GroupingKey
}

func (s *Incident) GetIncidentId() *string {
	return s.IncidentId
}

func (s *Incident) GetIncidentStatus() *string {
	return s.IncidentStatus
}

func (s *Incident) GetSeverity() *string {
	return s.Severity
}

func (s *Incident) GetStartTime() *int64 {
	return s.StartTime
}

func (s *Incident) GetStrategyUuid() *string {
	return s.StrategyUuid
}

func (s *Incident) GetUserId() *string {
	return s.UserId
}

func (s *Incident) SetActionTime(v int64) *Incident {
	s.ActionTime = &v
	return s
}

func (s *Incident) SetAlertCount(v int64) *Incident {
	s.AlertCount = &v
	return s
}

func (s *Incident) SetEndTime(v int64) *Incident {
	s.EndTime = &v
	return s
}

func (s *Incident) SetGroupingData(v map[string]interface{}) *Incident {
	s.GroupingData = v
	return s
}

func (s *Incident) SetGroupingId(v string) *Incident {
	s.GroupingId = &v
	return s
}

func (s *Incident) SetGroupingKey(v string) *Incident {
	s.GroupingKey = &v
	return s
}

func (s *Incident) SetIncidentId(v string) *Incident {
	s.IncidentId = &v
	return s
}

func (s *Incident) SetIncidentStatus(v string) *Incident {
	s.IncidentStatus = &v
	return s
}

func (s *Incident) SetSeverity(v string) *Incident {
	s.Severity = &v
	return s
}

func (s *Incident) SetStartTime(v int64) *Incident {
	s.StartTime = &v
	return s
}

func (s *Incident) SetStrategyUuid(v string) *Incident {
	s.StrategyUuid = &v
	return s
}

func (s *Incident) SetUserId(v string) *Incident {
	s.UserId = &v
	return s
}

func (s *Incident) Validate() error {
	return dara.Validate(s)
}
