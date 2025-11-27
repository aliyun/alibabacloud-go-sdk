// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobMetricLastRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArrayIndex(v []*int32) *DescribeJobMetricLastRequest
	GetArrayIndex() []*int32
	SetJobId(v string) *DescribeJobMetricLastRequest
	GetJobId() *string
	SetTaskName(v string) *DescribeJobMetricLastRequest
	GetTaskName() *string
}

type DescribeJobMetricLastRequest struct {
	// The list of array job indexes.
	ArrayIndex []*int32 `json:"ArrayIndex,omitempty" xml:"ArrayIndex,omitempty" type:"Repeated"`
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

func (s DescribeJobMetricLastRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobMetricLastRequest) GoString() string {
	return s.String()
}

func (s *DescribeJobMetricLastRequest) GetArrayIndex() []*int32 {
	return s.ArrayIndex
}

func (s *DescribeJobMetricLastRequest) GetJobId() *string {
	return s.JobId
}

func (s *DescribeJobMetricLastRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeJobMetricLastRequest) SetArrayIndex(v []*int32) *DescribeJobMetricLastRequest {
	s.ArrayIndex = v
	return s
}

func (s *DescribeJobMetricLastRequest) SetJobId(v string) *DescribeJobMetricLastRequest {
	s.JobId = &v
	return s
}

func (s *DescribeJobMetricLastRequest) SetTaskName(v string) *DescribeJobMetricLastRequest {
	s.TaskName = &v
	return s
}

func (s *DescribeJobMetricLastRequest) Validate() error {
	return dara.Validate(s)
}
