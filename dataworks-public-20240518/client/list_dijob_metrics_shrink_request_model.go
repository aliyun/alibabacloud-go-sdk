// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDIJobMetricsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDIJobId(v int64) *ListDIJobMetricsShrinkRequest
	GetDIJobId() *int64
	SetEndTime(v int64) *ListDIJobMetricsShrinkRequest
	GetEndTime() *int64
	SetMetricNameShrink(v string) *ListDIJobMetricsShrinkRequest
	GetMetricNameShrink() *string
	SetStartTime(v int64) *ListDIJobMetricsShrinkRequest
	GetStartTime() *int64
}

type ListDIJobMetricsShrinkRequest struct {
	// The ID of the synchronization task.
	//
	// example:
	//
	// 11265
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// The end of the time range to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1712205941
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The metrics that you want to query.
	//
	// This parameter is required.
	MetricNameShrink *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The beginning of the time range to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1586509407
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListDIJobMetricsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDIJobMetricsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDIJobMetricsShrinkRequest) GetDIJobId() *int64 {
	return s.DIJobId
}

func (s *ListDIJobMetricsShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListDIJobMetricsShrinkRequest) GetMetricNameShrink() *string {
	return s.MetricNameShrink
}

func (s *ListDIJobMetricsShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListDIJobMetricsShrinkRequest) SetDIJobId(v int64) *ListDIJobMetricsShrinkRequest {
	s.DIJobId = &v
	return s
}

func (s *ListDIJobMetricsShrinkRequest) SetEndTime(v int64) *ListDIJobMetricsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *ListDIJobMetricsShrinkRequest) SetMetricNameShrink(v string) *ListDIJobMetricsShrinkRequest {
	s.MetricNameShrink = &v
	return s
}

func (s *ListDIJobMetricsShrinkRequest) SetStartTime(v int64) *ListDIJobMetricsShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *ListDIJobMetricsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
