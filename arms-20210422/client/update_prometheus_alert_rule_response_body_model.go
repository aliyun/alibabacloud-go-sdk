// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrometheusAlertRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPrometheusAlertRule(v *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) *UpdatePrometheusAlertRuleResponseBody
	GetPrometheusAlertRule() *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule
	SetRequestId(v string) *UpdatePrometheusAlertRuleResponseBody
	GetRequestId() *string
}

type UpdatePrometheusAlertRuleResponseBody struct {
	PrometheusAlertRule *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule `json:"PrometheusAlertRule,omitempty" xml:"PrometheusAlertRule,omitempty" type:"Struct"`
	RequestId           *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePrometheusAlertRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrometheusAlertRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusAlertRuleResponseBody) GetPrometheusAlertRule() *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	return s.PrometheusAlertRule
}

func (s *UpdatePrometheusAlertRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePrometheusAlertRuleResponseBody) SetPrometheusAlertRule(v *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) *UpdatePrometheusAlertRuleResponseBody {
	s.PrometheusAlertRule = v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBody) SetRequestId(v string) *UpdatePrometheusAlertRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBody) Validate() error {
	if s.PrometheusAlertRule != nil {
		if err := s.PrometheusAlertRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule struct {
	AlertId        *int64                                                                 `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	AlertName      *string                                                                `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	Annotations    []*UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations `json:"Annotations,omitempty" xml:"Annotations,omitempty" type:"Repeated"`
	ClusterId      *string                                                                `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DispatchRuleId *int64                                                                 `json:"DispatchRuleId,omitempty" xml:"DispatchRuleId,omitempty"`
	Duration       *string                                                                `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Expression     *string                                                                `json:"Expression,omitempty" xml:"Expression,omitempty"`
	Labels         []*UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels      `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Message        *string                                                                `json:"Message,omitempty" xml:"Message,omitempty"`
	NotifyType     *string                                                                `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	Status         *int32                                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	Type           *string                                                                `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetAlertId() *int64 {
	return s.AlertId
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetAlertName() *string {
	return s.AlertName
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetAnnotations() []*UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations {
	return s.Annotations
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetDispatchRuleId() *int64 {
	return s.DispatchRuleId
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetDuration() *string {
	return s.Duration
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetExpression() *string {
	return s.Expression
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetLabels() []*UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels {
	return s.Labels
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetMessage() *string {
	return s.Message
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetNotifyType() *string {
	return s.NotifyType
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetStatus() *int32 {
	return s.Status
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetType() *string {
	return s.Type
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetAlertId(v int64) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.AlertId = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetAlertName(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.AlertName = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetAnnotations(v []*UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Annotations = v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetClusterId(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.ClusterId = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetDispatchRuleId(v int64) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.DispatchRuleId = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetDuration(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Duration = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetExpression(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Expression = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetLabels(v []*UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Labels = v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetMessage(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Message = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetNotifyType(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.NotifyType = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetStatus(v int32) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Status = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetType(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Type = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) Validate() error {
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

type UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) GetName() *string {
	return s.Name
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) GetValue() *string {
	return s.Value
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) SetName(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations {
	s.Name = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) SetValue(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations {
	s.Value = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) Validate() error {
	return dara.Validate(s)
}

type UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) GetName() *string {
	return s.Name
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) GetValue() *string {
	return s.Value
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) SetName(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels {
	s.Name = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) SetValue(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels {
	s.Value = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) Validate() error {
	return dara.Validate(s)
}
