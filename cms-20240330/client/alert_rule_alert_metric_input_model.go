// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlertRuleAlertMetricInput interface {
	dara.Model
	String() string
	GoString() string
	SetFilterValues(v []*AlertRuleAlertMetricInputFilterValue) *AlertRuleAlertMetricInput
	GetFilterValues() []*AlertRuleAlertMetricInputFilterValue
	SetGroupId(v string) *AlertRuleAlertMetricInput
	GetGroupId() *string
	SetMetricId(v string) *AlertRuleAlertMetricInput
	GetMetricId() *string
	SetParamValues(v []*AlertRuleAlertMetricInputParamValue) *AlertRuleAlertMetricInput
	GetParamValues() []*AlertRuleAlertMetricInputParamValue
}

type AlertRuleAlertMetricInput struct {
	// The list of filter conditions entered by the user. You can call ListAlertMetrics to query the parameters and filter conditions supported by the metric.
	FilterValues []*AlertRuleAlertMetricInputFilterValue `json:"filterValues,omitempty" xml:"filterValues,omitempty" type:"Repeated"`
	// The key of the metric group selected by the user.
	//
	// example:
	//
	// apm.jvm
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The key of the predefined metric selected by the user.
	//
	// example:
	//
	// appstat.jvm.GcPsMarkSweepCount
	MetricId *string `json:"metricId,omitempty" xml:"metricId,omitempty"`
	// The list of input parameters. You can call ListAlertMetrics to query the parameters and filter conditions supported by the metric.
	ParamValues []*AlertRuleAlertMetricInputParamValue `json:"paramValues,omitempty" xml:"paramValues,omitempty" type:"Repeated"`
}

func (s AlertRuleAlertMetricInput) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleAlertMetricInput) GoString() string {
	return s.String()
}

func (s *AlertRuleAlertMetricInput) GetFilterValues() []*AlertRuleAlertMetricInputFilterValue {
	return s.FilterValues
}

func (s *AlertRuleAlertMetricInput) GetGroupId() *string {
	return s.GroupId
}

func (s *AlertRuleAlertMetricInput) GetMetricId() *string {
	return s.MetricId
}

func (s *AlertRuleAlertMetricInput) GetParamValues() []*AlertRuleAlertMetricInputParamValue {
	return s.ParamValues
}

func (s *AlertRuleAlertMetricInput) SetFilterValues(v []*AlertRuleAlertMetricInputFilterValue) *AlertRuleAlertMetricInput {
	s.FilterValues = v
	return s
}

func (s *AlertRuleAlertMetricInput) SetGroupId(v string) *AlertRuleAlertMetricInput {
	s.GroupId = &v
	return s
}

func (s *AlertRuleAlertMetricInput) SetMetricId(v string) *AlertRuleAlertMetricInput {
	s.MetricId = &v
	return s
}

func (s *AlertRuleAlertMetricInput) SetParamValues(v []*AlertRuleAlertMetricInputParamValue) *AlertRuleAlertMetricInput {
	s.ParamValues = v
	return s
}

func (s *AlertRuleAlertMetricInput) Validate() error {
	if s.FilterValues != nil {
		for _, item := range s.FilterValues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ParamValues != nil {
		for _, item := range s.ParamValues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
