// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlertEvent interface {
	dara.Model
	String() string
	GoString() string
	SetAlertName(v string) *AlertEvent
	GetAlertName() *string
	SetAlertStatus(v string) *AlertEvent
	GetAlertStatus() *string
	SetArn(v string) *AlertEvent
	GetArn() *string
	SetContent(v string) *AlertEvent
	GetContent() *string
	SetCustomLabels(v map[string]interface{}) *AlertEvent
	GetCustomLabels() map[string]interface{}
	SetDeDupId(v string) *AlertEvent
	GetDeDupId() *string
	SetDetails(v string) *AlertEvent
	GetDetails() *string
	SetEventName(v string) *AlertEvent
	GetEventName() *string
	SetEventType(v string) *AlertEvent
	GetEventType() *string
	SetExpression(v string) *AlertEvent
	GetExpression() *string
	SetMetrics(v []*AlertEventMetrics) *AlertEvent
	GetMetrics() []*AlertEventMetrics
	SetProduct(v string) *AlertEvent
	GetProduct() *string
	SetResourceInfo(v map[string]interface{}) *AlertEvent
	GetResourceInfo() map[string]interface{}
	SetRuleName(v string) *AlertEvent
	GetRuleName() *string
	SetSeverity(v string) *AlertEvent
	GetSeverity() *string
	SetSource(v string) *AlertEvent
	GetSource() *string
	SetSummary(v string) *AlertEvent
	GetSummary() *string
	SetTimestamp(v int64) *AlertEvent
	GetTimestamp() *int64
	SetTraceId(v string) *AlertEvent
	GetTraceId() *string
	SetUserId(v string) *AlertEvent
	GetUserId() *string
}

type AlertEvent struct {
	AlertName    *string                `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	AlertStatus  *string                `json:"AlertStatus,omitempty" xml:"AlertStatus,omitempty"`
	Arn          *string                `json:"Arn,omitempty" xml:"Arn,omitempty"`
	Content      *string                `json:"Content,omitempty" xml:"Content,omitempty"`
	CustomLabels map[string]interface{} `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	DeDupId      *string                `json:"DeDupId,omitempty" xml:"DeDupId,omitempty"`
	Details      *string                `json:"Details,omitempty" xml:"Details,omitempty"`
	EventName    *string                `json:"EventName,omitempty" xml:"EventName,omitempty"`
	EventType    *string                `json:"EventType,omitempty" xml:"EventType,omitempty"`
	Expression   *string                `json:"Expression,omitempty" xml:"Expression,omitempty"`
	Metrics      []*AlertEventMetrics   `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	Product      *string                `json:"Product,omitempty" xml:"Product,omitempty"`
	ResourceInfo map[string]interface{} `json:"ResourceInfo,omitempty" xml:"ResourceInfo,omitempty"`
	RuleName     *string                `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Severity     *string                `json:"Severity,omitempty" xml:"Severity,omitempty"`
	Source       *string                `json:"Source,omitempty" xml:"Source,omitempty"`
	Summary      *string                `json:"Summary,omitempty" xml:"Summary,omitempty"`
	Timestamp    *int64                 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	TraceId      *string                `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
	UserId       *string                `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AlertEvent) String() string {
	return dara.Prettify(s)
}

func (s AlertEvent) GoString() string {
	return s.String()
}

func (s *AlertEvent) GetAlertName() *string {
	return s.AlertName
}

func (s *AlertEvent) GetAlertStatus() *string {
	return s.AlertStatus
}

func (s *AlertEvent) GetArn() *string {
	return s.Arn
}

func (s *AlertEvent) GetContent() *string {
	return s.Content
}

func (s *AlertEvent) GetCustomLabels() map[string]interface{} {
	return s.CustomLabels
}

func (s *AlertEvent) GetDeDupId() *string {
	return s.DeDupId
}

func (s *AlertEvent) GetDetails() *string {
	return s.Details
}

func (s *AlertEvent) GetEventName() *string {
	return s.EventName
}

func (s *AlertEvent) GetEventType() *string {
	return s.EventType
}

func (s *AlertEvent) GetExpression() *string {
	return s.Expression
}

func (s *AlertEvent) GetMetrics() []*AlertEventMetrics {
	return s.Metrics
}

func (s *AlertEvent) GetProduct() *string {
	return s.Product
}

func (s *AlertEvent) GetResourceInfo() map[string]interface{} {
	return s.ResourceInfo
}

func (s *AlertEvent) GetRuleName() *string {
	return s.RuleName
}

func (s *AlertEvent) GetSeverity() *string {
	return s.Severity
}

func (s *AlertEvent) GetSource() *string {
	return s.Source
}

func (s *AlertEvent) GetSummary() *string {
	return s.Summary
}

func (s *AlertEvent) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *AlertEvent) GetTraceId() *string {
	return s.TraceId
}

func (s *AlertEvent) GetUserId() *string {
	return s.UserId
}

func (s *AlertEvent) SetAlertName(v string) *AlertEvent {
	s.AlertName = &v
	return s
}

func (s *AlertEvent) SetAlertStatus(v string) *AlertEvent {
	s.AlertStatus = &v
	return s
}

func (s *AlertEvent) SetArn(v string) *AlertEvent {
	s.Arn = &v
	return s
}

func (s *AlertEvent) SetContent(v string) *AlertEvent {
	s.Content = &v
	return s
}

func (s *AlertEvent) SetCustomLabels(v map[string]interface{}) *AlertEvent {
	s.CustomLabels = v
	return s
}

func (s *AlertEvent) SetDeDupId(v string) *AlertEvent {
	s.DeDupId = &v
	return s
}

func (s *AlertEvent) SetDetails(v string) *AlertEvent {
	s.Details = &v
	return s
}

func (s *AlertEvent) SetEventName(v string) *AlertEvent {
	s.EventName = &v
	return s
}

func (s *AlertEvent) SetEventType(v string) *AlertEvent {
	s.EventType = &v
	return s
}

func (s *AlertEvent) SetExpression(v string) *AlertEvent {
	s.Expression = &v
	return s
}

func (s *AlertEvent) SetMetrics(v []*AlertEventMetrics) *AlertEvent {
	s.Metrics = v
	return s
}

func (s *AlertEvent) SetProduct(v string) *AlertEvent {
	s.Product = &v
	return s
}

func (s *AlertEvent) SetResourceInfo(v map[string]interface{}) *AlertEvent {
	s.ResourceInfo = v
	return s
}

func (s *AlertEvent) SetRuleName(v string) *AlertEvent {
	s.RuleName = &v
	return s
}

func (s *AlertEvent) SetSeverity(v string) *AlertEvent {
	s.Severity = &v
	return s
}

func (s *AlertEvent) SetSource(v string) *AlertEvent {
	s.Source = &v
	return s
}

func (s *AlertEvent) SetSummary(v string) *AlertEvent {
	s.Summary = &v
	return s
}

func (s *AlertEvent) SetTimestamp(v int64) *AlertEvent {
	s.Timestamp = &v
	return s
}

func (s *AlertEvent) SetTraceId(v string) *AlertEvent {
	s.TraceId = &v
	return s
}

func (s *AlertEvent) SetUserId(v string) *AlertEvent {
	s.UserId = &v
	return s
}

func (s *AlertEvent) Validate() error {
	if s.Metrics != nil {
		for _, item := range s.Metrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AlertEventMetrics struct {
	CurValue     *string  `json:"CurValue,omitempty" xml:"CurValue,omitempty"`
	MetricName   *string  `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	MetricNameEn *string  `json:"MetricNameEn,omitempty" xml:"MetricNameEn,omitempty"`
	MetricNameZh *string  `json:"MetricNameZh,omitempty" xml:"MetricNameZh,omitempty"`
	Operator     *string  `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Statistics   *string  `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold    *string  `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Unit         *string  `json:"Unit,omitempty" xml:"Unit,omitempty"`
	UnitFactor   *float32 `json:"UnitFactor,omitempty" xml:"UnitFactor,omitempty"`
}

func (s AlertEventMetrics) String() string {
	return dara.Prettify(s)
}

func (s AlertEventMetrics) GoString() string {
	return s.String()
}

func (s *AlertEventMetrics) GetCurValue() *string {
	return s.CurValue
}

func (s *AlertEventMetrics) GetMetricName() *string {
	return s.MetricName
}

func (s *AlertEventMetrics) GetMetricNameEn() *string {
	return s.MetricNameEn
}

func (s *AlertEventMetrics) GetMetricNameZh() *string {
	return s.MetricNameZh
}

func (s *AlertEventMetrics) GetOperator() *string {
	return s.Operator
}

func (s *AlertEventMetrics) GetStatistics() *string {
	return s.Statistics
}

func (s *AlertEventMetrics) GetThreshold() *string {
	return s.Threshold
}

func (s *AlertEventMetrics) GetUnit() *string {
	return s.Unit
}

func (s *AlertEventMetrics) GetUnitFactor() *float32 {
	return s.UnitFactor
}

func (s *AlertEventMetrics) SetCurValue(v string) *AlertEventMetrics {
	s.CurValue = &v
	return s
}

func (s *AlertEventMetrics) SetMetricName(v string) *AlertEventMetrics {
	s.MetricName = &v
	return s
}

func (s *AlertEventMetrics) SetMetricNameEn(v string) *AlertEventMetrics {
	s.MetricNameEn = &v
	return s
}

func (s *AlertEventMetrics) SetMetricNameZh(v string) *AlertEventMetrics {
	s.MetricNameZh = &v
	return s
}

func (s *AlertEventMetrics) SetOperator(v string) *AlertEventMetrics {
	s.Operator = &v
	return s
}

func (s *AlertEventMetrics) SetStatistics(v string) *AlertEventMetrics {
	s.Statistics = &v
	return s
}

func (s *AlertEventMetrics) SetThreshold(v string) *AlertEventMetrics {
	s.Threshold = &v
	return s
}

func (s *AlertEventMetrics) SetUnit(v string) *AlertEventMetrics {
	s.Unit = &v
	return s
}

func (s *AlertEventMetrics) SetUnitFactor(v float32) *AlertEventMetrics {
	s.UnitFactor = &v
	return s
}

func (s *AlertEventMetrics) Validate() error {
	return dara.Validate(s)
}
