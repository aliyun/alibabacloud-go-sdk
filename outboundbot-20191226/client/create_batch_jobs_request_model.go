// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBatchJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchJobDescription(v string) *CreateBatchJobsRequest
	GetBatchJobDescription() *string
	SetBatchJobName(v string) *CreateBatchJobsRequest
	GetBatchJobName() *string
	SetCallingNumber(v []*string) *CreateBatchJobsRequest
	GetCallingNumber() []*string
	SetInstanceId(v string) *CreateBatchJobsRequest
	GetInstanceId() *string
	SetJobFilePath(v string) *CreateBatchJobsRequest
	GetJobFilePath() *string
	SetScenarioId(v string) *CreateBatchJobsRequest
	GetScenarioId() *string
	SetScriptId(v string) *CreateBatchJobsRequest
	GetScriptId() *string
	SetStrategyJson(v string) *CreateBatchJobsRequest
	GetStrategyJson() *string
	SetSubmitted(v bool) *CreateBatchJobsRequest
	GetSubmitted() *bool
}

type CreateBatchJobsRequest struct {
	BatchJobDescription *string `json:"BatchJobDescription,omitempty" xml:"BatchJobDescription,omitempty"`
	// This parameter is required.
	BatchJobName *string `json:"BatchJobName,omitempty" xml:"BatchJobName,omitempty"`
	// example:
	//
	// ["95187"]
	CallingNumber []*string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 174952ab-9825-4cc9-a5e2-de82d7fa4cdd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 52e80b02-0126-4556-a1e6-ef5b3747ed53/a9a3ddc7-d7d7-48cd-82b5-b31bb5510e71_2a66f8ad-dfbb-4980-9b84-439171295a11.xlsx
	JobFilePath *string `json:"JobFilePath,omitempty" xml:"JobFilePath,omitempty"`
	// example:
	//
	// c6a668d1-3145-4048-9101-cb3678bb8884
	ScenarioId *string `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
	// example:
	//
	// b9ff4e88-65f9-4eb3-987c-11ba51f3f24d
	ScriptId     *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	StrategyJson *string `json:"StrategyJson,omitempty" xml:"StrategyJson,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	Submitted *bool `json:"Submitted,omitempty" xml:"Submitted,omitempty"`
}

func (s CreateBatchJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchJobsRequest) GoString() string {
	return s.String()
}

func (s *CreateBatchJobsRequest) GetBatchJobDescription() *string {
	return s.BatchJobDescription
}

func (s *CreateBatchJobsRequest) GetBatchJobName() *string {
	return s.BatchJobName
}

func (s *CreateBatchJobsRequest) GetCallingNumber() []*string {
	return s.CallingNumber
}

func (s *CreateBatchJobsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateBatchJobsRequest) GetJobFilePath() *string {
	return s.JobFilePath
}

func (s *CreateBatchJobsRequest) GetScenarioId() *string {
	return s.ScenarioId
}

func (s *CreateBatchJobsRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *CreateBatchJobsRequest) GetStrategyJson() *string {
	return s.StrategyJson
}

func (s *CreateBatchJobsRequest) GetSubmitted() *bool {
	return s.Submitted
}

func (s *CreateBatchJobsRequest) SetBatchJobDescription(v string) *CreateBatchJobsRequest {
	s.BatchJobDescription = &v
	return s
}

func (s *CreateBatchJobsRequest) SetBatchJobName(v string) *CreateBatchJobsRequest {
	s.BatchJobName = &v
	return s
}

func (s *CreateBatchJobsRequest) SetCallingNumber(v []*string) *CreateBatchJobsRequest {
	s.CallingNumber = v
	return s
}

func (s *CreateBatchJobsRequest) SetInstanceId(v string) *CreateBatchJobsRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateBatchJobsRequest) SetJobFilePath(v string) *CreateBatchJobsRequest {
	s.JobFilePath = &v
	return s
}

func (s *CreateBatchJobsRequest) SetScenarioId(v string) *CreateBatchJobsRequest {
	s.ScenarioId = &v
	return s
}

func (s *CreateBatchJobsRequest) SetScriptId(v string) *CreateBatchJobsRequest {
	s.ScriptId = &v
	return s
}

func (s *CreateBatchJobsRequest) SetStrategyJson(v string) *CreateBatchJobsRequest {
	s.StrategyJson = &v
	return s
}

func (s *CreateBatchJobsRequest) SetSubmitted(v bool) *CreateBatchJobsRequest {
	s.Submitted = &v
	return s
}

func (s *CreateBatchJobsRequest) Validate() error {
	return dara.Validate(s)
}
