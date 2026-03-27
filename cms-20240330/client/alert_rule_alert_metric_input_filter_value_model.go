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
	// Dimension of the filter condition.
	//
	// This parameter is required.
	//
	// example:
	//
	// rootIp
	Dim *string `json:"dim,omitempty" xml:"dim,omitempty"`
	// Filter Condition Operator.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALL
	Opt *string `json:"opt,omitempty" xml:"opt,omitempty"`
	// Filter Condition Value.
	//
	// example:
	//
	// 127.0.0.1
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
