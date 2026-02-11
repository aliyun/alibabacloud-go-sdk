// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrometheusAlertRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPrometheusAlertRules(v []*ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) *ListPrometheusAlertRulesResponseBody
	GetPrometheusAlertRules() []*ListPrometheusAlertRulesResponseBodyPrometheusAlertRules
	SetRequestId(v string) *ListPrometheusAlertRulesResponseBody
	GetRequestId() *string
}

type ListPrometheusAlertRulesResponseBody struct {
	PrometheusAlertRules []*ListPrometheusAlertRulesResponseBodyPrometheusAlertRules `json:"PrometheusAlertRules,omitempty" xml:"PrometheusAlertRules,omitempty" type:"Repeated"`
	RequestId            *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPrometheusAlertRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusAlertRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertRulesResponseBody) GetPrometheusAlertRules() []*ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	return s.PrometheusAlertRules
}

func (s *ListPrometheusAlertRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPrometheusAlertRulesResponseBody) SetPrometheusAlertRules(v []*ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) *ListPrometheusAlertRulesResponseBody {
	s.PrometheusAlertRules = v
	return s
}

func (s *ListPrometheusAlertRulesResponseBody) SetRequestId(v string) *ListPrometheusAlertRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBody) Validate() error {
	if s.PrometheusAlertRules != nil {
		for _, item := range s.PrometheusAlertRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPrometheusAlertRulesResponseBodyPrometheusAlertRules struct {
	AlertId        *int64                                                                 `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	AlertName      *string                                                                `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	Annotations    []*ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations `json:"Annotations,omitempty" xml:"Annotations,omitempty" type:"Repeated"`
	ClusterId      *string                                                                `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DispatchRuleId *int64                                                                 `json:"DispatchRuleId,omitempty" xml:"DispatchRuleId,omitempty"`
	Duration       *string                                                                `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Expression     *string                                                                `json:"Expression,omitempty" xml:"Expression,omitempty"`
	Labels         []*ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels      `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Message        *string                                                                `json:"Message,omitempty" xml:"Message,omitempty"`
	NotifyType     *string                                                                `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	Status         *int32                                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	Type           *string                                                                `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) GetAlertId() *int64 {
	return s.AlertId
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) GetAlertName() *string {
	return s.AlertName
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) GetAnnotations() []*ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations {
	return s.Annotations
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) GetDispatchRuleId() *int64 {
	return s.DispatchRuleId
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) GetDuration() *string {
	return s.Duration
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) GetExpression() *string {
	return s.Expression
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) GetLabels() []*ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels {
	return s.Labels
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) GetMessage() *string {
	return s.Message
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) GetNotifyType() *string {
	return s.NotifyType
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) GetStatus() *int32 {
	return s.Status
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) GetType() *string {
	return s.Type
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetAlertId(v int64) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.AlertId = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetAlertName(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.AlertName = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetAnnotations(v []*ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.Annotations = v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetClusterId(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.ClusterId = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetDispatchRuleId(v int64) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.DispatchRuleId = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetDuration(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.Duration = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetExpression(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.Expression = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetLabels(v []*ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.Labels = v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetMessage(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.Message = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetNotifyType(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.NotifyType = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetStatus(v int32) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.Status = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetType(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.Type = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) Validate() error {
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

type ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations) GetName() *string {
	return s.Name
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations) GetValue() *string {
	return s.Value
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations) SetName(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations {
	s.Name = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations) SetValue(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations {
	s.Value = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations) Validate() error {
	return dara.Validate(s)
}

type ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels) GetName() *string {
	return s.Name
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels) GetValue() *string {
	return s.Value
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels) SetName(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels {
	s.Name = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels) SetValue(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels {
	s.Value = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels) Validate() error {
	return dara.Validate(s)
}
