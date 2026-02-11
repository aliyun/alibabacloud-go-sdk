// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrometheusAlertRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPrometheusAlertRule(v *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) *DescribePrometheusAlertRuleResponseBody
	GetPrometheusAlertRule() *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule
	SetRequestId(v string) *DescribePrometheusAlertRuleResponseBody
	GetRequestId() *string
}

type DescribePrometheusAlertRuleResponseBody struct {
	PrometheusAlertRule *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule `json:"PrometheusAlertRule,omitempty" xml:"PrometheusAlertRule,omitempty" type:"Struct"`
	RequestId           *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePrometheusAlertRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePrometheusAlertRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePrometheusAlertRuleResponseBody) GetPrometheusAlertRule() *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	return s.PrometheusAlertRule
}

func (s *DescribePrometheusAlertRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePrometheusAlertRuleResponseBody) SetPrometheusAlertRule(v *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) *DescribePrometheusAlertRuleResponseBody {
	s.PrometheusAlertRule = v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBody) SetRequestId(v string) *DescribePrometheusAlertRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBody) Validate() error {
	if s.PrometheusAlertRule != nil {
		if err := s.PrometheusAlertRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule struct {
	AlertId        *int64                                                                   `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	AlertName      *string                                                                  `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	Annotations    []*DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations `json:"Annotations,omitempty" xml:"Annotations,omitempty" type:"Repeated"`
	ClusterId      *string                                                                  `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DispatchRuleId *int64                                                                   `json:"DispatchRuleId,omitempty" xml:"DispatchRuleId,omitempty"`
	Duration       *string                                                                  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Expression     *string                                                                  `json:"Expression,omitempty" xml:"Expression,omitempty"`
	Labels         []*DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels      `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Message        *string                                                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	NotifyType     *string                                                                  `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	Status         *int32                                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	Type           *string                                                                  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) String() string {
	return dara.Prettify(s)
}

func (s DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) GoString() string {
	return s.String()
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetAlertId() *int64 {
	return s.AlertId
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetAlertName() *string {
	return s.AlertName
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetAnnotations() []*DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations {
	return s.Annotations
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetDispatchRuleId() *int64 {
	return s.DispatchRuleId
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetDuration() *string {
	return s.Duration
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetExpression() *string {
	return s.Expression
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetLabels() []*DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels {
	return s.Labels
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetMessage() *string {
	return s.Message
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetNotifyType() *string {
	return s.NotifyType
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetStatus() *int32 {
	return s.Status
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetType() *string {
	return s.Type
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetAlertId(v int64) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.AlertId = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetAlertName(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.AlertName = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetAnnotations(v []*DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Annotations = v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetClusterId(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.ClusterId = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetDispatchRuleId(v int64) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.DispatchRuleId = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetDuration(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Duration = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetExpression(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Expression = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetLabels(v []*DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Labels = v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetMessage(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Message = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetNotifyType(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.NotifyType = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetStatus(v int32) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Status = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetType(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Type = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) Validate() error {
	if s.Annotations != nil {
		for _, item := range s.Annotations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) String() string {
	return dara.Prettify(s)
}

func (s DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) GoString() string {
	return s.String()
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) GetName() *string {
	return s.Name
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) GetValue() *string {
	return s.Value
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) SetName(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations {
	s.Name = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) SetValue(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations {
	s.Value = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) Validate() error {
	return dara.Validate(s)
}

type DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) String() string {
	return dara.Prettify(s)
}

func (s DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) GoString() string {
	return s.String()
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) GetName() *string {
	return s.Name
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) GetValue() *string {
	return s.Value
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) SetName(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels {
	s.Name = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) SetValue(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels {
	s.Value = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) Validate() error {
	return dara.Validate(s)
}
