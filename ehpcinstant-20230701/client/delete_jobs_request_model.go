// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExecutorIds(v []*string) *DeleteJobsRequest
	GetExecutorIds() []*string
	SetJobScheduler(v string) *DeleteJobsRequest
	GetJobScheduler() *string
	SetJobSpec(v []*DeleteJobsRequestJobSpec) *DeleteJobsRequest
	GetJobSpec() []*DeleteJobsRequestJobSpec
}

type DeleteJobsRequest struct {
	ExecutorIds  []*string                   `json:"ExecutorIds,omitempty" xml:"ExecutorIds,omitempty" type:"Repeated"`
	JobScheduler *string                     `json:"JobScheduler,omitempty" xml:"JobScheduler,omitempty"`
	JobSpec      []*DeleteJobsRequestJobSpec `json:"JobSpec,omitempty" xml:"JobSpec,omitempty" type:"Repeated"`
}

func (s DeleteJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteJobsRequest) GoString() string {
	return s.String()
}

func (s *DeleteJobsRequest) GetExecutorIds() []*string {
	return s.ExecutorIds
}

func (s *DeleteJobsRequest) GetJobScheduler() *string {
	return s.JobScheduler
}

func (s *DeleteJobsRequest) GetJobSpec() []*DeleteJobsRequestJobSpec {
	return s.JobSpec
}

func (s *DeleteJobsRequest) SetExecutorIds(v []*string) *DeleteJobsRequest {
	s.ExecutorIds = v
	return s
}

func (s *DeleteJobsRequest) SetJobScheduler(v string) *DeleteJobsRequest {
	s.JobScheduler = &v
	return s
}

func (s *DeleteJobsRequest) SetJobSpec(v []*DeleteJobsRequestJobSpec) *DeleteJobsRequest {
	s.JobSpec = v
	return s
}

func (s *DeleteJobsRequest) Validate() error {
	return dara.Validate(s)
}

type DeleteJobsRequestJobSpec struct {
	// example:
	//
	// job-xxxx
	JobId    *string                             `json:"JobId,omitempty" xml:"JobId,omitempty"`
	TaskSpec []*DeleteJobsRequestJobSpecTaskSpec `json:"TaskSpec,omitempty" xml:"TaskSpec,omitempty" type:"Repeated"`
}

func (s DeleteJobsRequestJobSpec) String() string {
	return dara.Prettify(s)
}

func (s DeleteJobsRequestJobSpec) GoString() string {
	return s.String()
}

func (s *DeleteJobsRequestJobSpec) GetJobId() *string {
	return s.JobId
}

func (s *DeleteJobsRequestJobSpec) GetTaskSpec() []*DeleteJobsRequestJobSpecTaskSpec {
	return s.TaskSpec
}

func (s *DeleteJobsRequestJobSpec) SetJobId(v string) *DeleteJobsRequestJobSpec {
	s.JobId = &v
	return s
}

func (s *DeleteJobsRequestJobSpec) SetTaskSpec(v []*DeleteJobsRequestJobSpecTaskSpec) *DeleteJobsRequestJobSpec {
	s.TaskSpec = v
	return s
}

func (s *DeleteJobsRequestJobSpec) Validate() error {
	return dara.Validate(s)
}

type DeleteJobsRequestJobSpecTaskSpec struct {
	ArrayIndex []*int32 `json:"ArrayIndex,omitempty" xml:"ArrayIndex,omitempty" type:"Repeated"`
	// example:
	//
	// task0
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s DeleteJobsRequestJobSpecTaskSpec) String() string {
	return dara.Prettify(s)
}

func (s DeleteJobsRequestJobSpecTaskSpec) GoString() string {
	return s.String()
}

func (s *DeleteJobsRequestJobSpecTaskSpec) GetArrayIndex() []*int32 {
	return s.ArrayIndex
}

func (s *DeleteJobsRequestJobSpecTaskSpec) GetTaskName() *string {
	return s.TaskName
}

func (s *DeleteJobsRequestJobSpecTaskSpec) SetArrayIndex(v []*int32) *DeleteJobsRequestJobSpecTaskSpec {
	s.ArrayIndex = v
	return s
}

func (s *DeleteJobsRequestJobSpecTaskSpec) SetTaskName(v string) *DeleteJobsRequestJobSpecTaskSpec {
	s.TaskName = &v
	return s
}

func (s *DeleteJobsRequestJobSpecTaskSpec) Validate() error {
	return dara.Validate(s)
}
