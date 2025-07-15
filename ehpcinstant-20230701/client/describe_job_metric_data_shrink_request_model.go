// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobMetricDataShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArrayIndexShrink(v string) *DescribeJobMetricDataShrinkRequest
	GetArrayIndexShrink() *string
	SetJobId(v string) *DescribeJobMetricDataShrinkRequest
	GetJobId() *string
	SetMetricName(v string) *DescribeJobMetricDataShrinkRequest
	GetMetricName() *string
	SetTaskName(v string) *DescribeJobMetricDataShrinkRequest
	GetTaskName() *string
}

type DescribeJobMetricDataShrinkRequest struct {
	ArrayIndexShrink *string `json:"ArrayIndex,omitempty" xml:"ArrayIndex,omitempty"`
	// example:
	//
	// job-xxxxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// cpu_utilization
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// example:
	//
	// Task0
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s DescribeJobMetricDataShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobMetricDataShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeJobMetricDataShrinkRequest) GetArrayIndexShrink() *string {
	return s.ArrayIndexShrink
}

func (s *DescribeJobMetricDataShrinkRequest) GetJobId() *string {
	return s.JobId
}

func (s *DescribeJobMetricDataShrinkRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeJobMetricDataShrinkRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeJobMetricDataShrinkRequest) SetArrayIndexShrink(v string) *DescribeJobMetricDataShrinkRequest {
	s.ArrayIndexShrink = &v
	return s
}

func (s *DescribeJobMetricDataShrinkRequest) SetJobId(v string) *DescribeJobMetricDataShrinkRequest {
	s.JobId = &v
	return s
}

func (s *DescribeJobMetricDataShrinkRequest) SetMetricName(v string) *DescribeJobMetricDataShrinkRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeJobMetricDataShrinkRequest) SetTaskName(v string) *DescribeJobMetricDataShrinkRequest {
	s.TaskName = &v
	return s
}

func (s *DescribeJobMetricDataShrinkRequest) Validate() error {
	return dara.Validate(s)
}
