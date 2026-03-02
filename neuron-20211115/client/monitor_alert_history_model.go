// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorAlertHistory interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *MonitorAlertHistory
	GetEndTime() *string
	SetEventState(v string) *MonitorAlertHistory
	GetEventState() *string
	SetId(v int64) *MonitorAlertHistory
	GetId() *int64
	SetName(v string) *MonitorAlertHistory
	GetName() *string
	SetRuleDesc(v string) *MonitorAlertHistory
	GetRuleDesc() *string
	SetRuleName(v string) *MonitorAlertHistory
	GetRuleName() *string
	SetStartTime(v string) *MonitorAlertHistory
	GetStartTime() *string
	SetType(v string) *MonitorAlertHistory
	GetType() *string
	SetUuid(v string) *MonitorAlertHistory
	GetUuid() *string
}

type MonitorAlertHistory struct {
	// example:
	//
	// 2022-10-13 13:58:42
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 已解决
	EventState *string `json:"eventState,omitempty" xml:"eventState,omitempty"`
	// example:
	//
	// 1
	Id       *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	RuleDesc *string `json:"ruleDesc,omitempty" xml:"ruleDesc,omitempty"`
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	// example:
	//
	// 2022-10-13 12:18:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 日志
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s MonitorAlertHistory) String() string {
	return dara.Prettify(s)
}

func (s MonitorAlertHistory) GoString() string {
	return s.String()
}

func (s *MonitorAlertHistory) GetEndTime() *string {
	return s.EndTime
}

func (s *MonitorAlertHistory) GetEventState() *string {
	return s.EventState
}

func (s *MonitorAlertHistory) GetId() *int64 {
	return s.Id
}

func (s *MonitorAlertHistory) GetName() *string {
	return s.Name
}

func (s *MonitorAlertHistory) GetRuleDesc() *string {
	return s.RuleDesc
}

func (s *MonitorAlertHistory) GetRuleName() *string {
	return s.RuleName
}

func (s *MonitorAlertHistory) GetStartTime() *string {
	return s.StartTime
}

func (s *MonitorAlertHistory) GetType() *string {
	return s.Type
}

func (s *MonitorAlertHistory) GetUuid() *string {
	return s.Uuid
}

func (s *MonitorAlertHistory) SetEndTime(v string) *MonitorAlertHistory {
	s.EndTime = &v
	return s
}

func (s *MonitorAlertHistory) SetEventState(v string) *MonitorAlertHistory {
	s.EventState = &v
	return s
}

func (s *MonitorAlertHistory) SetId(v int64) *MonitorAlertHistory {
	s.Id = &v
	return s
}

func (s *MonitorAlertHistory) SetName(v string) *MonitorAlertHistory {
	s.Name = &v
	return s
}

func (s *MonitorAlertHistory) SetRuleDesc(v string) *MonitorAlertHistory {
	s.RuleDesc = &v
	return s
}

func (s *MonitorAlertHistory) SetRuleName(v string) *MonitorAlertHistory {
	s.RuleName = &v
	return s
}

func (s *MonitorAlertHistory) SetStartTime(v string) *MonitorAlertHistory {
	s.StartTime = &v
	return s
}

func (s *MonitorAlertHistory) SetType(v string) *MonitorAlertHistory {
	s.Type = &v
	return s
}

func (s *MonitorAlertHistory) SetUuid(v string) *MonitorAlertHistory {
	s.Uuid = &v
	return s
}

func (s *MonitorAlertHistory) Validate() error {
	return dara.Validate(s)
}
