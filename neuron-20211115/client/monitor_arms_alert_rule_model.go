// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorArmsAlertRule interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *MonitorArmsAlertRule
	GetId() *int64
	SetMetricKey(v string) *MonitorArmsAlertRule
	GetMetricKey() *string
	SetMetricType(v string) *MonitorArmsAlertRule
	GetMetricType() *string
	SetName(v string) *MonitorArmsAlertRule
	GetName() *string
	SetTriggerAggregate(v string) *MonitorArmsAlertRule
	GetTriggerAggregate() *string
	SetTriggerInterval(v int64) *MonitorArmsAlertRule
	GetTriggerInterval() *int64
	SetTriggerOperator(v string) *MonitorArmsAlertRule
	GetTriggerOperator() *string
	SetTriggerThreshold(v int64) *MonitorArmsAlertRule
	GetTriggerThreshold() *int64
}

type MonitorArmsAlertRule struct {
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// appstat.jvm.gc.oldgccountinstant
	MetricKey *string `json:"metricKey,omitempty" xml:"metricKey,omitempty"`
	// example:
	//
	// JVM
	MetricType *string `json:"metricType,omitempty" xml:"metricType,omitempty"`
	// example:
	//
	// 规则1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// AVG
	TriggerAggregate *string `json:"triggerAggregate,omitempty" xml:"triggerAggregate,omitempty"`
	// example:
	//
	// 5
	TriggerInterval *int64 `json:"triggerInterval,omitempty" xml:"triggerInterval,omitempty"`
	// example:
	//
	// CURRENT_GTE
	TriggerOperator *string `json:"triggerOperator,omitempty" xml:"triggerOperator,omitempty"`
	// example:
	//
	// 10
	TriggerThreshold *int64 `json:"triggerThreshold,omitempty" xml:"triggerThreshold,omitempty"`
}

func (s MonitorArmsAlertRule) String() string {
	return dara.Prettify(s)
}

func (s MonitorArmsAlertRule) GoString() string {
	return s.String()
}

func (s *MonitorArmsAlertRule) GetId() *int64 {
	return s.Id
}

func (s *MonitorArmsAlertRule) GetMetricKey() *string {
	return s.MetricKey
}

func (s *MonitorArmsAlertRule) GetMetricType() *string {
	return s.MetricType
}

func (s *MonitorArmsAlertRule) GetName() *string {
	return s.Name
}

func (s *MonitorArmsAlertRule) GetTriggerAggregate() *string {
	return s.TriggerAggregate
}

func (s *MonitorArmsAlertRule) GetTriggerInterval() *int64 {
	return s.TriggerInterval
}

func (s *MonitorArmsAlertRule) GetTriggerOperator() *string {
	return s.TriggerOperator
}

func (s *MonitorArmsAlertRule) GetTriggerThreshold() *int64 {
	return s.TriggerThreshold
}

func (s *MonitorArmsAlertRule) SetId(v int64) *MonitorArmsAlertRule {
	s.Id = &v
	return s
}

func (s *MonitorArmsAlertRule) SetMetricKey(v string) *MonitorArmsAlertRule {
	s.MetricKey = &v
	return s
}

func (s *MonitorArmsAlertRule) SetMetricType(v string) *MonitorArmsAlertRule {
	s.MetricType = &v
	return s
}

func (s *MonitorArmsAlertRule) SetName(v string) *MonitorArmsAlertRule {
	s.Name = &v
	return s
}

func (s *MonitorArmsAlertRule) SetTriggerAggregate(v string) *MonitorArmsAlertRule {
	s.TriggerAggregate = &v
	return s
}

func (s *MonitorArmsAlertRule) SetTriggerInterval(v int64) *MonitorArmsAlertRule {
	s.TriggerInterval = &v
	return s
}

func (s *MonitorArmsAlertRule) SetTriggerOperator(v string) *MonitorArmsAlertRule {
	s.TriggerOperator = &v
	return s
}

func (s *MonitorArmsAlertRule) SetTriggerThreshold(v int64) *MonitorArmsAlertRule {
	s.TriggerThreshold = &v
	return s
}

func (s *MonitorArmsAlertRule) Validate() error {
	return dara.Validate(s)
}
