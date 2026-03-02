// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorSlsAlertRule interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *MonitorSlsAlertRule
	GetId() *int64
	SetName(v string) *MonitorSlsAlertRule
	GetName() *string
	SetQueryCondition(v string) *MonitorSlsAlertRule
	GetQueryCondition() *string
	SetTriggerInterval(v int64) *MonitorSlsAlertRule
	GetTriggerInterval() *int64
	SetTriggerIntervalUnit(v string) *MonitorSlsAlertRule
	GetTriggerIntervalUnit() *string
	SetTriggerOperator(v string) *MonitorSlsAlertRule
	GetTriggerOperator() *string
	SetTriggerThreshold(v int64) *MonitorSlsAlertRule
	GetTriggerThreshold() *int64
	SetTriggerThresholdUpper(v int64) *MonitorSlsAlertRule
	GetTriggerThresholdUpper() *int64
}

type MonitorSlsAlertRule struct {
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 规则1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 	- | select status, count(*) as total group by status
	QueryCondition *string `json:"queryCondition,omitempty" xml:"queryCondition,omitempty"`
	// example:
	//
	// 5
	TriggerInterval *int64 `json:"triggerInterval,omitempty" xml:"triggerInterval,omitempty"`
	// example:
	//
	// MINUTE
	TriggerIntervalUnit *string `json:"triggerIntervalUnit,omitempty" xml:"triggerIntervalUnit,omitempty"`
	// example:
	//
	// CURRENT_GTE
	TriggerOperator *string `json:"triggerOperator,omitempty" xml:"triggerOperator,omitempty"`
	// example:
	//
	// 10
	TriggerThreshold *int64 `json:"triggerThreshold,omitempty" xml:"triggerThreshold,omitempty"`
	// example:
	//
	// 100
	TriggerThresholdUpper *int64 `json:"triggerThresholdUpper,omitempty" xml:"triggerThresholdUpper,omitempty"`
}

func (s MonitorSlsAlertRule) String() string {
	return dara.Prettify(s)
}

func (s MonitorSlsAlertRule) GoString() string {
	return s.String()
}

func (s *MonitorSlsAlertRule) GetId() *int64 {
	return s.Id
}

func (s *MonitorSlsAlertRule) GetName() *string {
	return s.Name
}

func (s *MonitorSlsAlertRule) GetQueryCondition() *string {
	return s.QueryCondition
}

func (s *MonitorSlsAlertRule) GetTriggerInterval() *int64 {
	return s.TriggerInterval
}

func (s *MonitorSlsAlertRule) GetTriggerIntervalUnit() *string {
	return s.TriggerIntervalUnit
}

func (s *MonitorSlsAlertRule) GetTriggerOperator() *string {
	return s.TriggerOperator
}

func (s *MonitorSlsAlertRule) GetTriggerThreshold() *int64 {
	return s.TriggerThreshold
}

func (s *MonitorSlsAlertRule) GetTriggerThresholdUpper() *int64 {
	return s.TriggerThresholdUpper
}

func (s *MonitorSlsAlertRule) SetId(v int64) *MonitorSlsAlertRule {
	s.Id = &v
	return s
}

func (s *MonitorSlsAlertRule) SetName(v string) *MonitorSlsAlertRule {
	s.Name = &v
	return s
}

func (s *MonitorSlsAlertRule) SetQueryCondition(v string) *MonitorSlsAlertRule {
	s.QueryCondition = &v
	return s
}

func (s *MonitorSlsAlertRule) SetTriggerInterval(v int64) *MonitorSlsAlertRule {
	s.TriggerInterval = &v
	return s
}

func (s *MonitorSlsAlertRule) SetTriggerIntervalUnit(v string) *MonitorSlsAlertRule {
	s.TriggerIntervalUnit = &v
	return s
}

func (s *MonitorSlsAlertRule) SetTriggerOperator(v string) *MonitorSlsAlertRule {
	s.TriggerOperator = &v
	return s
}

func (s *MonitorSlsAlertRule) SetTriggerThreshold(v int64) *MonitorSlsAlertRule {
	s.TriggerThreshold = &v
	return s
}

func (s *MonitorSlsAlertRule) SetTriggerThresholdUpper(v int64) *MonitorSlsAlertRule {
	s.TriggerThresholdUpper = &v
	return s
}

func (s *MonitorSlsAlertRule) Validate() error {
	return dara.Validate(s)
}
