// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAutoscalingMetricSpec interface {
	dara.Model
	String() string
	GoString() string
	SetMetricName(v string) *AutoscalingMetricSpec
	GetMetricName() *string
	SetTargetValue(v int32) *AutoscalingMetricSpec
	GetTargetValue() *int32
}

type AutoscalingMetricSpec struct {
	MetricName  *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	TargetValue *int32  `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
}

func (s AutoscalingMetricSpec) String() string {
	return dara.Prettify(s)
}

func (s AutoscalingMetricSpec) GoString() string {
	return s.String()
}

func (s *AutoscalingMetricSpec) GetMetricName() *string {
	return s.MetricName
}

func (s *AutoscalingMetricSpec) GetTargetValue() *int32 {
	return s.TargetValue
}

func (s *AutoscalingMetricSpec) SetMetricName(v string) *AutoscalingMetricSpec {
	s.MetricName = &v
	return s
}

func (s *AutoscalingMetricSpec) SetTargetValue(v int32) *AutoscalingMetricSpec {
	s.TargetValue = &v
	return s
}

func (s *AutoscalingMetricSpec) Validate() error {
	return dara.Validate(s)
}
