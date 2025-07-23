// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlertRuleAlertMetricInputFilterValue interface {
	dara.Model
	String() string
	GoString() string
	SetDim(v string) *AlertRuleAlertMetricInputFilterValue
	GetDim() *string
	SetOpt(v string) *AlertRuleAlertMetricInputFilterValue
	GetOpt() *string
	SetValue(v string) *AlertRuleAlertMetricInputFilterValue
	GetValue() *string
}

type AlertRuleAlertMetricInputFilterValue struct {
	// This parameter is required.
	Dim *string `json:"dim,omitempty" xml:"dim,omitempty"`
	// This parameter is required.
	Opt   *string `json:"opt,omitempty" xml:"opt,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s AlertRuleAlertMetricInputFilterValue) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleAlertMetricInputFilterValue) GoString() string {
	return s.String()
}

func (s *AlertRuleAlertMetricInputFilterValue) GetDim() *string {
	return s.Dim
}

func (s *AlertRuleAlertMetricInputFilterValue) GetOpt() *string {
	return s.Opt
}

func (s *AlertRuleAlertMetricInputFilterValue) GetValue() *string {
	return s.Value
}

func (s *AlertRuleAlertMetricInputFilterValue) SetDim(v string) *AlertRuleAlertMetricInputFilterValue {
	s.Dim = &v
	return s
}

func (s *AlertRuleAlertMetricInputFilterValue) SetOpt(v string) *AlertRuleAlertMetricInputFilterValue {
	s.Opt = &v
	return s
}

func (s *AlertRuleAlertMetricInputFilterValue) SetValue(v string) *AlertRuleAlertMetricInputFilterValue {
	s.Value = &v
	return s
}

func (s *AlertRuleAlertMetricInputFilterValue) Validate() error {
	return dara.Validate(s)
}
