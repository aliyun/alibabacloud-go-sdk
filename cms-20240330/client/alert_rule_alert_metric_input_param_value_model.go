// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlertRuleAlertMetricInputParamValue interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *AlertRuleAlertMetricInputParamValue
	GetName() *string
	SetValue(v string) *AlertRuleAlertMetricInputParamValue
	GetValue() *string
}

type AlertRuleAlertMetricInputParamValue struct {
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s AlertRuleAlertMetricInputParamValue) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleAlertMetricInputParamValue) GoString() string {
	return s.String()
}

func (s *AlertRuleAlertMetricInputParamValue) GetName() *string {
	return s.Name
}

func (s *AlertRuleAlertMetricInputParamValue) GetValue() *string {
	return s.Value
}

func (s *AlertRuleAlertMetricInputParamValue) SetName(v string) *AlertRuleAlertMetricInputParamValue {
	s.Name = &v
	return s
}

func (s *AlertRuleAlertMetricInputParamValue) SetValue(v string) *AlertRuleAlertMetricInputParamValue {
	s.Value = &v
	return s
}

func (s *AlertRuleAlertMetricInputParamValue) Validate() error {
	return dara.Validate(s)
}
