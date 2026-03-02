// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorMqAlertRule interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *MonitorMqAlertRule
	GetId() *int64
	SetMetricDimension(v string) *MonitorMqAlertRule
	GetMetricDimension() *string
	SetMetricKey(v string) *MonitorMqAlertRule
	GetMetricKey() *string
	SetMetricType(v string) *MonitorMqAlertRule
	GetMetricType() *string
	SetName(v string) *MonitorMqAlertRule
	GetName() *string
	SetTriggerInterval(v int64) *MonitorMqAlertRule
	GetTriggerInterval() *int64
	SetTriggerOperator(v string) *MonitorMqAlertRule
	GetTriggerOperator() *string
	SetTriggerThreshold(v int64) *MonitorMqAlertRule
	GetTriggerThreshold() *int64
}

type MonitorMqAlertRule struct {
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// PBC_Topic1
	MetricDimension *string `json:"metricDimension,omitempty" xml:"metricDimension,omitempty"`
	// example:
	//
	// appstat.jvm.gc.oldgccountinstant
	MetricKey *string `json:"metricKey,omitempty" xml:"metricKey,omitempty"`
	// example:
	//
	// TOPIC
	MetricType *string `json:"metricType,omitempty" xml:"metricType,omitempty"`
	// example:
	//
	// 规则1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 5
	TriggerInterval *int64 `json:"triggerInterval,omitempty" xml:"triggerInterval,omitempty"`
	// example:
	//
	// GreaterThanOrEqualToThreshold
	TriggerOperator *string `json:"triggerOperator,omitempty" xml:"triggerOperator,omitempty"`
	// example:
	//
	// 10
	TriggerThreshold *int64 `json:"triggerThreshold,omitempty" xml:"triggerThreshold,omitempty"`
}

func (s MonitorMqAlertRule) String() string {
	return dara.Prettify(s)
}

func (s MonitorMqAlertRule) GoString() string {
	return s.String()
}

func (s *MonitorMqAlertRule) GetId() *int64 {
	return s.Id
}

func (s *MonitorMqAlertRule) GetMetricDimension() *string {
	return s.MetricDimension
}

func (s *MonitorMqAlertRule) GetMetricKey() *string {
	return s.MetricKey
}

func (s *MonitorMqAlertRule) GetMetricType() *string {
	return s.MetricType
}

func (s *MonitorMqAlertRule) GetName() *string {
	return s.Name
}

func (s *MonitorMqAlertRule) GetTriggerInterval() *int64 {
	return s.TriggerInterval
}

func (s *MonitorMqAlertRule) GetTriggerOperator() *string {
	return s.TriggerOperator
}

func (s *MonitorMqAlertRule) GetTriggerThreshold() *int64 {
	return s.TriggerThreshold
}

func (s *MonitorMqAlertRule) SetId(v int64) *MonitorMqAlertRule {
	s.Id = &v
	return s
}

func (s *MonitorMqAlertRule) SetMetricDimension(v string) *MonitorMqAlertRule {
	s.MetricDimension = &v
	return s
}

func (s *MonitorMqAlertRule) SetMetricKey(v string) *MonitorMqAlertRule {
	s.MetricKey = &v
	return s
}

func (s *MonitorMqAlertRule) SetMetricType(v string) *MonitorMqAlertRule {
	s.MetricType = &v
	return s
}

func (s *MonitorMqAlertRule) SetName(v string) *MonitorMqAlertRule {
	s.Name = &v
	return s
}

func (s *MonitorMqAlertRule) SetTriggerInterval(v int64) *MonitorMqAlertRule {
	s.TriggerInterval = &v
	return s
}

func (s *MonitorMqAlertRule) SetTriggerOperator(v string) *MonitorMqAlertRule {
	s.TriggerOperator = &v
	return s
}

func (s *MonitorMqAlertRule) SetTriggerThreshold(v int64) *MonitorMqAlertRule {
	s.TriggerThreshold = &v
	return s
}

func (s *MonitorMqAlertRule) Validate() error {
	return dara.Validate(s)
}
