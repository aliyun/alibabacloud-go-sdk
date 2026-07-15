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
	SetStabilizationWindowSeconds(v int32) *AutoscalingMetricSpec
	GetStabilizationWindowSeconds() *int32
	SetTargetValue(v int32) *AutoscalingMetricSpec
	GetTargetValue() *int32
	SetTolerance(v string) *AutoscalingMetricSpec
	GetTolerance() *string
}

type AutoscalingMetricSpec struct {
	// The name of the metric for autoscaling. This can be a predefined or a custom metric.
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The cooldown period, in seconds, after a scaling activity. This prevents the service from initiating another scaling action before the effects of the previous one are observable, stabilizing resource fluctuations.
	StabilizationWindowSeconds *int32 `json:"StabilizationWindowSeconds,omitempty" xml:"StabilizationWindowSeconds,omitempty"`
	// The target value for the specified metric. The autoscaling service tries to maintain the metric at or near this value.
	TargetValue *int32 `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
	// The acceptable deviation from the `TargetValue`, specified as a percentage string. A scaling action is triggered only if the metric value moves outside the range defined by the `TargetValue` and this tolerance. This prevents scaling actions based on minor fluctuations.
	Tolerance *string `json:"Tolerance,omitempty" xml:"Tolerance,omitempty"`
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

func (s *AutoscalingMetricSpec) GetStabilizationWindowSeconds() *int32 {
	return s.StabilizationWindowSeconds
}

func (s *AutoscalingMetricSpec) GetTargetValue() *int32 {
	return s.TargetValue
}

func (s *AutoscalingMetricSpec) GetTolerance() *string {
	return s.Tolerance
}

func (s *AutoscalingMetricSpec) SetMetricName(v string) *AutoscalingMetricSpec {
	s.MetricName = &v
	return s
}

func (s *AutoscalingMetricSpec) SetStabilizationWindowSeconds(v int32) *AutoscalingMetricSpec {
	s.StabilizationWindowSeconds = &v
	return s
}

func (s *AutoscalingMetricSpec) SetTargetValue(v int32) *AutoscalingMetricSpec {
	s.TargetValue = &v
	return s
}

func (s *AutoscalingMetricSpec) SetTolerance(v string) *AutoscalingMetricSpec {
	s.Tolerance = &v
	return s
}

func (s *AutoscalingMetricSpec) Validate() error {
	return dara.Validate(s)
}
