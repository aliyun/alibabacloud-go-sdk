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
	// The name of the monitoring metrics used for auto scaling.
	//
	// example:
	//
	// POD_CPU_UTIL
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The duration of the stabilization window for the auto scaling metric, in seconds.
	//
	// example:
	//
	// 60
	StabilizationWindowSeconds *int32 `json:"StabilizationWindowSeconds,omitempty" xml:"StabilizationWindowSeconds,omitempty"`
	// The target value of the monitoring metrics for auto scaling.
	//
	// example:
	//
	// 80
	TargetValue *int32 `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
	// The allowed bias ratio for the target value of the monitoring metrics.
	//
	// example:
	//
	// 0.1
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
