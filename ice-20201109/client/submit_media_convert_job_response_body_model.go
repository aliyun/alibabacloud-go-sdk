// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitMediaConvertJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJob(v *SubmitMediaConvertJobResponseBodyJob) *SubmitMediaConvertJobResponseBody
	GetJob() *SubmitMediaConvertJobResponseBodyJob
	SetRequestId(v string) *SubmitMediaConvertJobResponseBody
	GetRequestId() *string
}

type SubmitMediaConvertJobResponseBody struct {
	// The transcoding task.
	Job *SubmitMediaConvertJobResponseBodyJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitMediaConvertJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaConvertJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitMediaConvertJobResponseBody) GetJob() *SubmitMediaConvertJobResponseBodyJob {
	return s.Job
}

func (s *SubmitMediaConvertJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitMediaConvertJobResponseBody) SetJob(v *SubmitMediaConvertJobResponseBodyJob) *SubmitMediaConvertJobResponseBody {
	s.Job = v
	return s
}

func (s *SubmitMediaConvertJobResponseBody) SetRequestId(v string) *SubmitMediaConvertJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitMediaConvertJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitMediaConvertJobResponseBodyJob struct {
	// The idempotency key of the request for creating the transcoding task.
	//
	// example:
	//
	// FB7F25E9-AD9B-1603-8AF6-F1E42DF2E706
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The error code returned when the transcoding task failed.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The configurations of the transcoding task.
	Config *SubmitMediaConvertJobResponseBodyJobConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// The ID of the transcoding task.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The error message returned when the transcoding task failed.
	//
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The details of the transcoded outputs.
	OutputDetails []*MediaConvertOutputDetail `json:"OutputDetails,omitempty" xml:"OutputDetails,omitempty" type:"Repeated"`
	// The details of the output groups.
	OutputGroupDetails []*MediaConvertOutputGroupDetail `json:"OutputGroupDetails,omitempty" xml:"OutputGroupDetails,omitempty" type:"Repeated"`
	// The ID of the queue.
	//
	// example:
	//
	// 3780049
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A2129C9F-CE95-58B5-B8C1-07758FF6C86F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the transcoding task. Valid values:
	//
	// 	- Inited: The task is initialized.
	//
	// 	- Running
	//
	// 	- Success
	//
	// 	- Failed
	//
	// 	- Cancelled
	//
	// example:
	//
	// Created
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The user data.
	//
	// example:
	//
	// {"videoId":"abcd"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitMediaConvertJobResponseBodyJob) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaConvertJobResponseBodyJob) GoString() string {
	return s.String()
}

func (s *SubmitMediaConvertJobResponseBodyJob) GetClientToken() *string {
	return s.ClientToken
}

func (s *SubmitMediaConvertJobResponseBodyJob) GetCode() *string {
	return s.Code
}

func (s *SubmitMediaConvertJobResponseBodyJob) GetConfig() *SubmitMediaConvertJobResponseBodyJobConfig {
	return s.Config
}

func (s *SubmitMediaConvertJobResponseBodyJob) GetJobId() *string {
	return s.JobId
}

func (s *SubmitMediaConvertJobResponseBodyJob) GetMessage() *string {
	return s.Message
}

func (s *SubmitMediaConvertJobResponseBodyJob) GetOutputDetails() []*MediaConvertOutputDetail {
	return s.OutputDetails
}

func (s *SubmitMediaConvertJobResponseBodyJob) GetOutputGroupDetails() []*MediaConvertOutputGroupDetail {
	return s.OutputGroupDetails
}

func (s *SubmitMediaConvertJobResponseBodyJob) GetPipelineId() *string {
	return s.PipelineId
}

func (s *SubmitMediaConvertJobResponseBodyJob) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitMediaConvertJobResponseBodyJob) GetState() *string {
	return s.State
}

func (s *SubmitMediaConvertJobResponseBodyJob) GetUserData() *string {
	return s.UserData
}

func (s *SubmitMediaConvertJobResponseBodyJob) SetClientToken(v string) *SubmitMediaConvertJobResponseBodyJob {
	s.ClientToken = &v
	return s
}

func (s *SubmitMediaConvertJobResponseBodyJob) SetCode(v string) *SubmitMediaConvertJobResponseBodyJob {
	s.Code = &v
	return s
}

func (s *SubmitMediaConvertJobResponseBodyJob) SetConfig(v *SubmitMediaConvertJobResponseBodyJobConfig) *SubmitMediaConvertJobResponseBodyJob {
	s.Config = v
	return s
}

func (s *SubmitMediaConvertJobResponseBodyJob) SetJobId(v string) *SubmitMediaConvertJobResponseBodyJob {
	s.JobId = &v
	return s
}

func (s *SubmitMediaConvertJobResponseBodyJob) SetMessage(v string) *SubmitMediaConvertJobResponseBodyJob {
	s.Message = &v
	return s
}

func (s *SubmitMediaConvertJobResponseBodyJob) SetOutputDetails(v []*MediaConvertOutputDetail) *SubmitMediaConvertJobResponseBodyJob {
	s.OutputDetails = v
	return s
}

func (s *SubmitMediaConvertJobResponseBodyJob) SetOutputGroupDetails(v []*MediaConvertOutputGroupDetail) *SubmitMediaConvertJobResponseBodyJob {
	s.OutputGroupDetails = v
	return s
}

func (s *SubmitMediaConvertJobResponseBodyJob) SetPipelineId(v string) *SubmitMediaConvertJobResponseBodyJob {
	s.PipelineId = &v
	return s
}

func (s *SubmitMediaConvertJobResponseBodyJob) SetRequestId(v string) *SubmitMediaConvertJobResponseBodyJob {
	s.RequestId = &v
	return s
}

func (s *SubmitMediaConvertJobResponseBodyJob) SetState(v string) *SubmitMediaConvertJobResponseBodyJob {
	s.State = &v
	return s
}

func (s *SubmitMediaConvertJobResponseBodyJob) SetUserData(v string) *SubmitMediaConvertJobResponseBodyJob {
	s.UserData = &v
	return s
}

func (s *SubmitMediaConvertJobResponseBodyJob) Validate() error {
	return dara.Validate(s)
}

type SubmitMediaConvertJobResponseBodyJobConfig struct {
	// The inputs of the transcoding task.
	Inputs  []*MediaConvertInput `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Repeated"`
	JobName *string              `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The output group configurations.
	OutputGroups []*MediaConvertOutputGroup `json:"OutputGroups,omitempty" xml:"OutputGroups,omitempty" type:"Repeated"`
	// The output configurations.
	Outputs []*MediaConvertOutput `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
}

func (s SubmitMediaConvertJobResponseBodyJobConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaConvertJobResponseBodyJobConfig) GoString() string {
	return s.String()
}

func (s *SubmitMediaConvertJobResponseBodyJobConfig) GetInputs() []*MediaConvertInput {
	return s.Inputs
}

func (s *SubmitMediaConvertJobResponseBodyJobConfig) GetJobName() *string {
	return s.JobName
}

func (s *SubmitMediaConvertJobResponseBodyJobConfig) GetOutputGroups() []*MediaConvertOutputGroup {
	return s.OutputGroups
}

func (s *SubmitMediaConvertJobResponseBodyJobConfig) GetOutputs() []*MediaConvertOutput {
	return s.Outputs
}

func (s *SubmitMediaConvertJobResponseBodyJobConfig) SetInputs(v []*MediaConvertInput) *SubmitMediaConvertJobResponseBodyJobConfig {
	s.Inputs = v
	return s
}

func (s *SubmitMediaConvertJobResponseBodyJobConfig) SetJobName(v string) *SubmitMediaConvertJobResponseBodyJobConfig {
	s.JobName = &v
	return s
}

func (s *SubmitMediaConvertJobResponseBodyJobConfig) SetOutputGroups(v []*MediaConvertOutputGroup) *SubmitMediaConvertJobResponseBodyJobConfig {
	s.OutputGroups = v
	return s
}

func (s *SubmitMediaConvertJobResponseBodyJobConfig) SetOutputs(v []*MediaConvertOutput) *SubmitMediaConvertJobResponseBodyJobConfig {
	s.Outputs = v
	return s
}

func (s *SubmitMediaConvertJobResponseBodyJobConfig) Validate() error {
	return dara.Validate(s)
}
