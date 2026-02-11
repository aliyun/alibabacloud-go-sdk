// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrometheusAlertRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPrometheusAlertRule(v *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) *CreatePrometheusAlertRuleResponseBody
	GetPrometheusAlertRule() *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule
	SetRequestId(v string) *CreatePrometheusAlertRuleResponseBody
	GetRequestId() *string
}

type CreatePrometheusAlertRuleResponseBody struct {
	PrometheusAlertRule *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule `json:"PrometheusAlertRule,omitempty" xml:"PrometheusAlertRule,omitempty" type:"Struct"`
	RequestId           *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePrometheusAlertRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePrometheusAlertRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePrometheusAlertRuleResponseBody) GetPrometheusAlertRule() *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	return s.PrometheusAlertRule
}

func (s *CreatePrometheusAlertRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePrometheusAlertRuleResponseBody) SetPrometheusAlertRule(v *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) *CreatePrometheusAlertRuleResponseBody {
	s.PrometheusAlertRule = v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBody) SetRequestId(v string) *CreatePrometheusAlertRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBody) Validate() error {
	if s.PrometheusAlertRule != nil {
		if err := s.PrometheusAlertRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule struct {
	AlertId        *int64                                                                 `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	AlertName      *string                                                                `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	Annotations    []*CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations `json:"Annotations,omitempty" xml:"Annotations,omitempty" type:"Repeated"`
	ClusterId      *string                                                                `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DispatchRuleId *int64                                                                 `json:"DispatchRuleId,omitempty" xml:"DispatchRuleId,omitempty"`
	Duration       *string                                                                `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Expression     *string                                                                `json:"Expression,omitempty" xml:"Expression,omitempty"`
	Labels         []*CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels      `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Message        *string                                                                `json:"Message,omitempty" xml:"Message,omitempty"`
	NotifyType     *string                                                                `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	Status         *int32                                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	Type           *string                                                                `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) String() string {
	return dara.Prettify(s)
}

func (s CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GoString() string {
	return s.String()
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetAlertId() *int64 {
	return s.AlertId
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetAlertName() *string {
	return s.AlertName
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetAnnotations() []*CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations {
	return s.Annotations
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetDispatchRuleId() *int64 {
	return s.DispatchRuleId
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetDuration() *string {
	return s.Duration
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetExpression() *string {
	return s.Expression
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetLabels() []*CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels {
	return s.Labels
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetMessage() *string {
	return s.Message
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetNotifyType() *string {
	return s.NotifyType
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetStatus() *int32 {
	return s.Status
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetType() *string {
	return s.Type
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetAlertId(v int64) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.AlertId = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetAlertName(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.AlertName = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetAnnotations(v []*CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Annotations = v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetClusterId(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.ClusterId = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetDispatchRuleId(v int64) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.DispatchRuleId = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetDuration(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Duration = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetExpression(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Expression = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetLabels(v []*CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Labels = v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetMessage(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Message = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetNotifyType(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.NotifyType = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetStatus(v int32) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Status = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetType(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Type = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) Validate() error {
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

type CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) String() string {
	return dara.Prettify(s)
}

func (s CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) GoString() string {
	return s.String()
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) GetName() *string {
	return s.Name
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) GetValue() *string {
	return s.Value
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) SetName(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations {
	s.Name = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) SetValue(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations {
	s.Value = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) Validate() error {
	return dara.Validate(s)
}

type CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) String() string {
	return dara.Prettify(s)
}

func (s CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) GoString() string {
	return s.String()
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) GetName() *string {
	return s.Name
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) GetValue() *string {
	return s.Value
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) SetName(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels {
	s.Name = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) SetValue(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels {
	s.Value = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) Validate() error {
	return dara.Validate(s)
}
