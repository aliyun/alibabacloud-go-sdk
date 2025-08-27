// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScalingPolicy interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ScalingPolicy
	GetEndTime() *string
	SetMaxInstances(v int64) *ScalingPolicy
	GetMaxInstances() *int64
	SetMetricTarget(v float32) *ScalingPolicy
	GetMetricTarget() *float32
	SetMetricType(v string) *ScalingPolicy
	GetMetricType() *string
	SetMinInstances(v int64) *ScalingPolicy
	GetMinInstances() *int64
	SetName(v string) *ScalingPolicy
	GetName() *string
	SetStartTime(v string) *ScalingPolicy
	GetStartTime() *string
	SetTimeZone(v string) *ScalingPolicy
	GetTimeZone() *string
}

type ScalingPolicy struct {
	EndTime      *string  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	MaxInstances *int64   `json:"maxInstances,omitempty" xml:"maxInstances,omitempty"`
	MetricTarget *float32 `json:"metricTarget,omitempty" xml:"metricTarget,omitempty"`
	MetricType   *string  `json:"metricType,omitempty" xml:"metricType,omitempty"`
	MinInstances *int64   `json:"minInstances,omitempty" xml:"minInstances,omitempty"`
	Name         *string  `json:"name,omitempty" xml:"name,omitempty"`
	StartTime    *string  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	TimeZone     *string  `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s ScalingPolicy) String() string {
	return dara.Prettify(s)
}

func (s ScalingPolicy) GoString() string {
	return s.String()
}

func (s *ScalingPolicy) GetEndTime() *string {
	return s.EndTime
}

func (s *ScalingPolicy) GetMaxInstances() *int64 {
	return s.MaxInstances
}

func (s *ScalingPolicy) GetMetricTarget() *float32 {
	return s.MetricTarget
}

func (s *ScalingPolicy) GetMetricType() *string {
	return s.MetricType
}

func (s *ScalingPolicy) GetMinInstances() *int64 {
	return s.MinInstances
}

func (s *ScalingPolicy) GetName() *string {
	return s.Name
}

func (s *ScalingPolicy) GetStartTime() *string {
	return s.StartTime
}

func (s *ScalingPolicy) GetTimeZone() *string {
	return s.TimeZone
}

func (s *ScalingPolicy) SetEndTime(v string) *ScalingPolicy {
	s.EndTime = &v
	return s
}

func (s *ScalingPolicy) SetMaxInstances(v int64) *ScalingPolicy {
	s.MaxInstances = &v
	return s
}

func (s *ScalingPolicy) SetMetricTarget(v float32) *ScalingPolicy {
	s.MetricTarget = &v
	return s
}

func (s *ScalingPolicy) SetMetricType(v string) *ScalingPolicy {
	s.MetricType = &v
	return s
}

func (s *ScalingPolicy) SetMinInstances(v int64) *ScalingPolicy {
	s.MinInstances = &v
	return s
}

func (s *ScalingPolicy) SetName(v string) *ScalingPolicy {
	s.Name = &v
	return s
}

func (s *ScalingPolicy) SetStartTime(v string) *ScalingPolicy {
	s.StartTime = &v
	return s
}

func (s *ScalingPolicy) SetTimeZone(v string) *ScalingPolicy {
	s.TimeZone = &v
	return s
}

func (s *ScalingPolicy) Validate() error {
	return dara.Validate(s)
}
