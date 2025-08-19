// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTargetTrackingPolicy interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *TargetTrackingPolicy
	GetEndTime() *string
	SetMaxCapacity(v int64) *TargetTrackingPolicy
	GetMaxCapacity() *int64
	SetMetricTarget(v float32) *TargetTrackingPolicy
	GetMetricTarget() *float32
	SetMetricType(v string) *TargetTrackingPolicy
	GetMetricType() *string
	SetMinCapacity(v int64) *TargetTrackingPolicy
	GetMinCapacity() *int64
	SetName(v string) *TargetTrackingPolicy
	GetName() *string
	SetStartTime(v string) *TargetTrackingPolicy
	GetStartTime() *string
	SetTimeZone(v string) *TargetTrackingPolicy
	GetTimeZone() *string
}

type TargetTrackingPolicy struct {
	// example:
	//
	// 2024-03-10T10:10:10
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	MaxCapacity *int64 `json:"maxCapacity,omitempty" xml:"maxCapacity,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.6
	MetricTarget *float32 `json:"metricTarget,omitempty" xml:"metricTarget,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CPUUtilization
	MetricType *string `json:"metricType,omitempty" xml:"metricType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	MinCapacity *int64 `json:"minCapacity,omitempty" xml:"minCapacity,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 2023-03-10T10:10:10
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s TargetTrackingPolicy) String() string {
	return dara.Prettify(s)
}

func (s TargetTrackingPolicy) GoString() string {
	return s.String()
}

func (s *TargetTrackingPolicy) GetEndTime() *string {
	return s.EndTime
}

func (s *TargetTrackingPolicy) GetMaxCapacity() *int64 {
	return s.MaxCapacity
}

func (s *TargetTrackingPolicy) GetMetricTarget() *float32 {
	return s.MetricTarget
}

func (s *TargetTrackingPolicy) GetMetricType() *string {
	return s.MetricType
}

func (s *TargetTrackingPolicy) GetMinCapacity() *int64 {
	return s.MinCapacity
}

func (s *TargetTrackingPolicy) GetName() *string {
	return s.Name
}

func (s *TargetTrackingPolicy) GetStartTime() *string {
	return s.StartTime
}

func (s *TargetTrackingPolicy) GetTimeZone() *string {
	return s.TimeZone
}

func (s *TargetTrackingPolicy) SetEndTime(v string) *TargetTrackingPolicy {
	s.EndTime = &v
	return s
}

func (s *TargetTrackingPolicy) SetMaxCapacity(v int64) *TargetTrackingPolicy {
	s.MaxCapacity = &v
	return s
}

func (s *TargetTrackingPolicy) SetMetricTarget(v float32) *TargetTrackingPolicy {
	s.MetricTarget = &v
	return s
}

func (s *TargetTrackingPolicy) SetMetricType(v string) *TargetTrackingPolicy {
	s.MetricType = &v
	return s
}

func (s *TargetTrackingPolicy) SetMinCapacity(v int64) *TargetTrackingPolicy {
	s.MinCapacity = &v
	return s
}

func (s *TargetTrackingPolicy) SetName(v string) *TargetTrackingPolicy {
	s.Name = &v
	return s
}

func (s *TargetTrackingPolicy) SetStartTime(v string) *TargetTrackingPolicy {
	s.StartTime = &v
	return s
}

func (s *TargetTrackingPolicy) SetTimeZone(v string) *TargetTrackingPolicy {
	s.TimeZone = &v
	return s
}

func (s *TargetTrackingPolicy) Validate() error {
	return dara.Validate(s)
}
