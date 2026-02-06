// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAll(v bool) *CancelJobsRequest
	GetAll() *bool
	SetInstanceId(v string) *CancelJobsRequest
	GetInstanceId() *string
	SetJobGroupId(v string) *CancelJobsRequest
	GetJobGroupId() *string
	SetJobId(v []*string) *CancelJobsRequest
	GetJobId() []*string
	SetJobReferenceId(v []*string) *CancelJobsRequest
	GetJobReferenceId() []*string
	SetScenarioId(v string) *CancelJobsRequest
	GetScenarioId() *string
}

type CancelJobsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// false
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 174952ab-9825-4cc9-a5e2-de82d7fa4cdd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 9f84892d-721a-4069-9975-668c8164d64e
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	// example:
	//
	// edf45790-7200-4cbc-b157-8c0a5f400b75
	JobId []*string `json:"JobId,omitempty" xml:"JobId,omitempty" type:"Repeated"`
	// example:
	//
	// 4a875676-b136-4087-88b4-de67c61fed69
	JobReferenceId []*string `json:"JobReferenceId,omitempty" xml:"JobReferenceId,omitempty" type:"Repeated"`
	// example:
	//
	// 9cef0dd3-b9d6-4748-9a6f-77a8c3402bb1
	ScenarioId *string `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
}

func (s CancelJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelJobsRequest) GoString() string {
	return s.String()
}

func (s *CancelJobsRequest) GetAll() *bool {
	return s.All
}

func (s *CancelJobsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CancelJobsRequest) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *CancelJobsRequest) GetJobId() []*string {
	return s.JobId
}

func (s *CancelJobsRequest) GetJobReferenceId() []*string {
	return s.JobReferenceId
}

func (s *CancelJobsRequest) GetScenarioId() *string {
	return s.ScenarioId
}

func (s *CancelJobsRequest) SetAll(v bool) *CancelJobsRequest {
	s.All = &v
	return s
}

func (s *CancelJobsRequest) SetInstanceId(v string) *CancelJobsRequest {
	s.InstanceId = &v
	return s
}

func (s *CancelJobsRequest) SetJobGroupId(v string) *CancelJobsRequest {
	s.JobGroupId = &v
	return s
}

func (s *CancelJobsRequest) SetJobId(v []*string) *CancelJobsRequest {
	s.JobId = v
	return s
}

func (s *CancelJobsRequest) SetJobReferenceId(v []*string) *CancelJobsRequest {
	s.JobReferenceId = v
	return s
}

func (s *CancelJobsRequest) SetScenarioId(v string) *CancelJobsRequest {
	s.ScenarioId = &v
	return s
}

func (s *CancelJobsRequest) Validate() error {
	return dara.Validate(s)
}
