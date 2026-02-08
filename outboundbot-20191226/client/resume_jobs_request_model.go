// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAll(v bool) *ResumeJobsRequest
	GetAll() *bool
	SetInstanceId(v string) *ResumeJobsRequest
	GetInstanceId() *string
	SetJobGroupId(v string) *ResumeJobsRequest
	GetJobGroupId() *string
	SetJobId(v []*string) *ResumeJobsRequest
	GetJobId() []*string
	SetJobReferenceId(v []*string) *ResumeJobsRequest
	GetJobReferenceId() []*string
	SetScenarioId(v string) *ResumeJobsRequest
	GetScenarioId() *string
}

type ResumeJobsRequest struct {
	// Filter condition indicating whether to restart all jobs
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Filter condition: task group ID
	//
	// example:
	//
	// de48407d-309e-451a-81ec-6fb11f8fdbf3
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	// List of job IDs
	//
	// > This parameter is required when the value of All is false.
	//
	// example:
	//
	// b72425bd-7871-4050-838e-033d80d754b7
	JobId []*string `json:"JobId,omitempty" xml:"JobId,omitempty" type:"Repeated"`
	// Third-party ID of the job
	//
	// > This is the ReferenceId provided by the customer when uploading the outbound call list.
	//
	// example:
	//
	// d5971d98-7312-4f0e-a918-a17d67133e28
	JobReferenceId []*string `json:"JobReferenceId,omitempty" xml:"JobReferenceId,omitempty" type:"Repeated"`
	// Filter condition: scenario ID (historical parameter, deprecated)
	//
	// example:
	//
	// b016fbdb-b81c-4c06-8870-cb36b8783b6d
	ScenarioId *string `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
}

func (s ResumeJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ResumeJobsRequest) GoString() string {
	return s.String()
}

func (s *ResumeJobsRequest) GetAll() *bool {
	return s.All
}

func (s *ResumeJobsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ResumeJobsRequest) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *ResumeJobsRequest) GetJobId() []*string {
	return s.JobId
}

func (s *ResumeJobsRequest) GetJobReferenceId() []*string {
	return s.JobReferenceId
}

func (s *ResumeJobsRequest) GetScenarioId() *string {
	return s.ScenarioId
}

func (s *ResumeJobsRequest) SetAll(v bool) *ResumeJobsRequest {
	s.All = &v
	return s
}

func (s *ResumeJobsRequest) SetInstanceId(v string) *ResumeJobsRequest {
	s.InstanceId = &v
	return s
}

func (s *ResumeJobsRequest) SetJobGroupId(v string) *ResumeJobsRequest {
	s.JobGroupId = &v
	return s
}

func (s *ResumeJobsRequest) SetJobId(v []*string) *ResumeJobsRequest {
	s.JobId = v
	return s
}

func (s *ResumeJobsRequest) SetJobReferenceId(v []*string) *ResumeJobsRequest {
	s.JobReferenceId = v
	return s
}

func (s *ResumeJobsRequest) SetScenarioId(v string) *ResumeJobsRequest {
	s.ScenarioId = &v
	return s
}

func (s *ResumeJobsRequest) Validate() error {
	return dara.Validate(s)
}
