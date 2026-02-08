// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBatchJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchJobName(v string) *ModifyBatchJobsRequest
	GetBatchJobName() *string
	SetCallingNumber(v []*string) *ModifyBatchJobsRequest
	GetCallingNumber() []*string
	SetDescription(v string) *ModifyBatchJobsRequest
	GetDescription() *string
	SetInstanceId(v string) *ModifyBatchJobsRequest
	GetInstanceId() *string
	SetJobFilePath(v string) *ModifyBatchJobsRequest
	GetJobFilePath() *string
	SetJobGroupId(v string) *ModifyBatchJobsRequest
	GetJobGroupId() *string
	SetScenarioId(v string) *ModifyBatchJobsRequest
	GetScenarioId() *string
	SetScriptId(v string) *ModifyBatchJobsRequest
	GetScriptId() *string
	SetStrategyJson(v string) *ModifyBatchJobsRequest
	GetStrategyJson() *string
	SetSubmitted(v bool) *ModifyBatchJobsRequest
	GetSubmitted() *bool
}

type ModifyBatchJobsRequest struct {
	// Job name
	//
	// This parameter is required.
	//
	// example:
	//
	// 催收作业组
	BatchJobName *string `json:"BatchJobName,omitempty" xml:"BatchJobName,omitempty"`
	// Calling numbers
	//
	// example:
	//
	// ["95187"]
	CallingNumber []*string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty" type:"Repeated"`
	// Description of the job group
	//
	// example:
	//
	// 这是催收作业组-01
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 174952ab-9825-4cc9-a5e2-de82d7fa4cdd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Download URL of the job Excel file
	//
	// example:
	//
	// 52e80b02-0126-4556-a1e6-ef5b3747ed53/a9a3ddc7-d7d7-48cd-82b5-b31bb5510e71_2a66f8ad-dfbb-4980-9b84-439171295a11.xlsx
	JobFilePath *string `json:"JobFilePath,omitempty" xml:"JobFilePath,omitempty"`
	// Job group ID
	//
	// This parameter is required.
	//
	// example:
	//
	// c6a668d1-3145-4048-9101-cb3678bb8884
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	// Scenario ID
	//
	// example:
	//
	// c6a668d1-3145-4048-9101-cb3678bb8884
	ScenarioId *string `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
	// Script ID
	//
	// example:
	//
	// 7d820242-f4f0-4d2e-ae35-b424c41cbc5b
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// Policy in JSON format
	//
	// example:
	//
	// {"maxAttemptsPerDay":2,"name":"策略名字","workingTime":[{"beginTime":"09:00:00","endTime":"12:00:00"},{"beginTime":"13:00:00","endTime":"18:30:00"}],"minAttemptInterval":60}
	StrategyJson *string `json:"StrategyJson,omitempty" xml:"StrategyJson,omitempty"`
	// Specifies whether to submit the job. If false, the job is submitted. If true, it is saved as a draft.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	Submitted *bool `json:"Submitted,omitempty" xml:"Submitted,omitempty"`
}

func (s ModifyBatchJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBatchJobsRequest) GoString() string {
	return s.String()
}

func (s *ModifyBatchJobsRequest) GetBatchJobName() *string {
	return s.BatchJobName
}

func (s *ModifyBatchJobsRequest) GetCallingNumber() []*string {
	return s.CallingNumber
}

func (s *ModifyBatchJobsRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyBatchJobsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyBatchJobsRequest) GetJobFilePath() *string {
	return s.JobFilePath
}

func (s *ModifyBatchJobsRequest) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *ModifyBatchJobsRequest) GetScenarioId() *string {
	return s.ScenarioId
}

func (s *ModifyBatchJobsRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *ModifyBatchJobsRequest) GetStrategyJson() *string {
	return s.StrategyJson
}

func (s *ModifyBatchJobsRequest) GetSubmitted() *bool {
	return s.Submitted
}

func (s *ModifyBatchJobsRequest) SetBatchJobName(v string) *ModifyBatchJobsRequest {
	s.BatchJobName = &v
	return s
}

func (s *ModifyBatchJobsRequest) SetCallingNumber(v []*string) *ModifyBatchJobsRequest {
	s.CallingNumber = v
	return s
}

func (s *ModifyBatchJobsRequest) SetDescription(v string) *ModifyBatchJobsRequest {
	s.Description = &v
	return s
}

func (s *ModifyBatchJobsRequest) SetInstanceId(v string) *ModifyBatchJobsRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyBatchJobsRequest) SetJobFilePath(v string) *ModifyBatchJobsRequest {
	s.JobFilePath = &v
	return s
}

func (s *ModifyBatchJobsRequest) SetJobGroupId(v string) *ModifyBatchJobsRequest {
	s.JobGroupId = &v
	return s
}

func (s *ModifyBatchJobsRequest) SetScenarioId(v string) *ModifyBatchJobsRequest {
	s.ScenarioId = &v
	return s
}

func (s *ModifyBatchJobsRequest) SetScriptId(v string) *ModifyBatchJobsRequest {
	s.ScriptId = &v
	return s
}

func (s *ModifyBatchJobsRequest) SetStrategyJson(v string) *ModifyBatchJobsRequest {
	s.StrategyJson = &v
	return s
}

func (s *ModifyBatchJobsRequest) SetSubmitted(v bool) *ModifyBatchJobsRequest {
	s.Submitted = &v
	return s
}

func (s *ModifyBatchJobsRequest) Validate() error {
	return dara.Validate(s)
}
