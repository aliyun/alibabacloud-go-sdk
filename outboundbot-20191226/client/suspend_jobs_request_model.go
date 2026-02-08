// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAll(v bool) *SuspendJobsRequest
	GetAll() *bool
	SetInstanceId(v string) *SuspendJobsRequest
	GetInstanceId() *string
	SetJobGroupId(v string) *SuspendJobsRequest
	GetJobGroupId() *string
	SetJobId(v []*string) *SuspendJobsRequest
	GetJobId() []*string
	SetJobReferenceId(v []*string) *SuspendJobsRequest
	GetJobReferenceId() []*string
	SetScenarioId(v string) *SuspendJobsRequest
	GetScenarioId() *string
}

type SuspendJobsRequest struct {
	// Filter condition indicating whether to pause all jobs.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ef4c09ac-2f5c-49e7-9d33-5d85249cee10
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Filter condition: task group ID.
	//
	// example:
	//
	// 6b3ea2a1-32b3-4041-842b-9bde5de9dda0
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	// List of job IDs.
	//
	// > When All is false, JobId is required.
	//
	// example:
	//
	// 11994321-e6bc-47bb-8b1c-8eef8f2f768b
	JobId []*string `json:"JobId,omitempty" xml:"JobId,omitempty" type:"Repeated"`
	// Third-party ID of the job.
	//
	// > This corresponds to the ReferenceId uploaded by the customer when uploading the outbound calling list.
	//
	// example:
	//
	// de3ab269-6746-477c-b13d-bd49f13202c2
	JobReferenceId []*string `json:"JobReferenceId,omitempty" xml:"JobReferenceId,omitempty" type:"Repeated"`
	// Filter condition: scenario ID (historical parameter, deprecated).
	//
	// example:
	//
	// 4b6dd926-3cc3-4111-a333-15d9b006fe81
	ScenarioId *string `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
}

func (s SuspendJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s SuspendJobsRequest) GoString() string {
	return s.String()
}

func (s *SuspendJobsRequest) GetAll() *bool {
	return s.All
}

func (s *SuspendJobsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SuspendJobsRequest) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *SuspendJobsRequest) GetJobId() []*string {
	return s.JobId
}

func (s *SuspendJobsRequest) GetJobReferenceId() []*string {
	return s.JobReferenceId
}

func (s *SuspendJobsRequest) GetScenarioId() *string {
	return s.ScenarioId
}

func (s *SuspendJobsRequest) SetAll(v bool) *SuspendJobsRequest {
	s.All = &v
	return s
}

func (s *SuspendJobsRequest) SetInstanceId(v string) *SuspendJobsRequest {
	s.InstanceId = &v
	return s
}

func (s *SuspendJobsRequest) SetJobGroupId(v string) *SuspendJobsRequest {
	s.JobGroupId = &v
	return s
}

func (s *SuspendJobsRequest) SetJobId(v []*string) *SuspendJobsRequest {
	s.JobId = v
	return s
}

func (s *SuspendJobsRequest) SetJobReferenceId(v []*string) *SuspendJobsRequest {
	s.JobReferenceId = v
	return s
}

func (s *SuspendJobsRequest) SetScenarioId(v string) *SuspendJobsRequest {
	s.ScenarioId = &v
	return s
}

func (s *SuspendJobsRequest) Validate() error {
	return dara.Validate(s)
}
