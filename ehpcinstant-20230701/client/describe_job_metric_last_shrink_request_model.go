// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobMetricLastShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArrayIndexShrink(v string) *DescribeJobMetricLastShrinkRequest
	GetArrayIndexShrink() *string
	SetJobId(v string) *DescribeJobMetricLastShrinkRequest
	GetJobId() *string
	SetTaskName(v string) *DescribeJobMetricLastShrinkRequest
	GetTaskName() *string
}

type DescribeJobMetricLastShrinkRequest struct {
	// The list of array job indexes.
	ArrayIndexShrink *string `json:"ArrayIndex,omitempty" xml:"ArrayIndex,omitempty"`
	// The job ID.
	//
	// example:
	//
	// job-xxxxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// Task0
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s DescribeJobMetricLastShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobMetricLastShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeJobMetricLastShrinkRequest) GetArrayIndexShrink() *string {
	return s.ArrayIndexShrink
}

func (s *DescribeJobMetricLastShrinkRequest) GetJobId() *string {
	return s.JobId
}

func (s *DescribeJobMetricLastShrinkRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeJobMetricLastShrinkRequest) SetArrayIndexShrink(v string) *DescribeJobMetricLastShrinkRequest {
	s.ArrayIndexShrink = &v
	return s
}

func (s *DescribeJobMetricLastShrinkRequest) SetJobId(v string) *DescribeJobMetricLastShrinkRequest {
	s.JobId = &v
	return s
}

func (s *DescribeJobMetricLastShrinkRequest) SetTaskName(v string) *DescribeJobMetricLastShrinkRequest {
	s.TaskName = &v
	return s
}

func (s *DescribeJobMetricLastShrinkRequest) Validate() error {
	return dara.Validate(s)
}
