// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobMetricDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArrayIndex(v []*int32) *DescribeJobMetricDataRequest
	GetArrayIndex() []*int32
	SetJobId(v string) *DescribeJobMetricDataRequest
	GetJobId() *string
	SetMetricName(v string) *DescribeJobMetricDataRequest
	GetMetricName() *string
	SetTaskName(v string) *DescribeJobMetricDataRequest
	GetTaskName() *string
}

type DescribeJobMetricDataRequest struct {
	// The list of array job indexes.
	ArrayIndex []*int32 `json:"ArrayIndex,omitempty" xml:"ArrayIndex,omitempty" type:"Repeated"`
	// The job ID.
	//
	// example:
	//
	// job-xxxxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The metrics of the job.
	//
	// example:
	//
	// cpu_utilization
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// Task0
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s DescribeJobMetricDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobMetricDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeJobMetricDataRequest) GetArrayIndex() []*int32 {
	return s.ArrayIndex
}

func (s *DescribeJobMetricDataRequest) GetJobId() *string {
	return s.JobId
}

func (s *DescribeJobMetricDataRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeJobMetricDataRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeJobMetricDataRequest) SetArrayIndex(v []*int32) *DescribeJobMetricDataRequest {
	s.ArrayIndex = v
	return s
}

func (s *DescribeJobMetricDataRequest) SetJobId(v string) *DescribeJobMetricDataRequest {
	s.JobId = &v
	return s
}

func (s *DescribeJobMetricDataRequest) SetMetricName(v string) *DescribeJobMetricDataRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeJobMetricDataRequest) SetTaskName(v string) *DescribeJobMetricDataRequest {
	s.TaskName = &v
	return s
}

func (s *DescribeJobMetricDataRequest) Validate() error {
	return dara.Validate(s)
}
