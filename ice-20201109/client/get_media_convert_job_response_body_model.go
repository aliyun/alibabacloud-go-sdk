// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaConvertJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJob(v *GetMediaConvertJobResponseBodyJob) *GetMediaConvertJobResponseBody
	GetJob() *GetMediaConvertJobResponseBodyJob
	SetRequestId(v string) *GetMediaConvertJobResponseBody
	GetRequestId() *string
}

type GetMediaConvertJobResponseBody struct {
	// The transcoding task.
	Job *GetMediaConvertJobResponseBodyJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 4BAEA8E8-1C16-5CD3-AC50-CCBA81A53402
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMediaConvertJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMediaConvertJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaConvertJobResponseBody) GetJob() *GetMediaConvertJobResponseBodyJob {
	return s.Job
}

func (s *GetMediaConvertJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMediaConvertJobResponseBody) SetJob(v *GetMediaConvertJobResponseBodyJob) *GetMediaConvertJobResponseBody {
	s.Job = v
	return s
}

func (s *GetMediaConvertJobResponseBody) SetRequestId(v string) *GetMediaConvertJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMediaConvertJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMediaConvertJobResponseBodyJob struct {
	// The idempotency key of the request for creating the transcoding task.
	//
	// example:
	//
	// 780018cb-55ba-466d-8acc-946c0c319a0e
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The error code returned when the transcoding task failed.
	//
	// example:
	//
	// InvalidParameter.ResourceContentBad
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The configurations of the transcoding task.
	Config     *GetMediaConvertJobResponseBodyJobConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	CreateTime *string                                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FinishTime *string                                  `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The ID of the transcoding task, which is a 32-bit string.
	//
	// example:
	//
	// ******4579b5e748b99a27f6d6******
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The error message returned when the transcoding task failed.
	//
	// example:
	//
	// The resource operated InputFile is bad
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The details of the transcoded outputs, each corresponding to an output configuration.
	OutputDetails []*MediaConvertOutputDetail `json:"OutputDetails,omitempty" xml:"OutputDetails,omitempty" type:"Repeated"`
	// The details of the output groups, each corresponding to an output group configuration.
	OutputGroupDetails []*MediaConvertOutputGroupDetail `json:"OutputGroupDetails,omitempty" xml:"OutputGroupDetails,omitempty" type:"Repeated"`
	Percent            *int32                           `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The ID of the queue.
	//
	// example:
	//
	// 83500cb2a3b94fabb0956e38d64bd16d
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The ID of the request for creating the transcoding task.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
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
	// Success
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The user data.
	//
	// example:
	//
	// {"videoId":"ddd333"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetMediaConvertJobResponseBodyJob) String() string {
	return dara.Prettify(s)
}

func (s GetMediaConvertJobResponseBodyJob) GoString() string {
	return s.String()
}

func (s *GetMediaConvertJobResponseBodyJob) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetMediaConvertJobResponseBodyJob) GetCode() *string {
	return s.Code
}

func (s *GetMediaConvertJobResponseBodyJob) GetConfig() *GetMediaConvertJobResponseBodyJobConfig {
	return s.Config
}

func (s *GetMediaConvertJobResponseBodyJob) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetMediaConvertJobResponseBodyJob) GetFinishTime() *string {
	return s.FinishTime
}

func (s *GetMediaConvertJobResponseBodyJob) GetJobId() *string {
	return s.JobId
}

func (s *GetMediaConvertJobResponseBodyJob) GetMessage() *string {
	return s.Message
}

func (s *GetMediaConvertJobResponseBodyJob) GetOutputDetails() []*MediaConvertOutputDetail {
	return s.OutputDetails
}

func (s *GetMediaConvertJobResponseBodyJob) GetOutputGroupDetails() []*MediaConvertOutputGroupDetail {
	return s.OutputGroupDetails
}

func (s *GetMediaConvertJobResponseBodyJob) GetPercent() *int32 {
	return s.Percent
}

func (s *GetMediaConvertJobResponseBodyJob) GetPipelineId() *string {
	return s.PipelineId
}

func (s *GetMediaConvertJobResponseBodyJob) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMediaConvertJobResponseBodyJob) GetState() *string {
	return s.State
}

func (s *GetMediaConvertJobResponseBodyJob) GetUserData() *string {
	return s.UserData
}

func (s *GetMediaConvertJobResponseBodyJob) SetClientToken(v string) *GetMediaConvertJobResponseBodyJob {
	s.ClientToken = &v
	return s
}

func (s *GetMediaConvertJobResponseBodyJob) SetCode(v string) *GetMediaConvertJobResponseBodyJob {
	s.Code = &v
	return s
}

func (s *GetMediaConvertJobResponseBodyJob) SetConfig(v *GetMediaConvertJobResponseBodyJobConfig) *GetMediaConvertJobResponseBodyJob {
	s.Config = v
	return s
}

func (s *GetMediaConvertJobResponseBodyJob) SetCreateTime(v string) *GetMediaConvertJobResponseBodyJob {
	s.CreateTime = &v
	return s
}

func (s *GetMediaConvertJobResponseBodyJob) SetFinishTime(v string) *GetMediaConvertJobResponseBodyJob {
	s.FinishTime = &v
	return s
}

func (s *GetMediaConvertJobResponseBodyJob) SetJobId(v string) *GetMediaConvertJobResponseBodyJob {
	s.JobId = &v
	return s
}

func (s *GetMediaConvertJobResponseBodyJob) SetMessage(v string) *GetMediaConvertJobResponseBodyJob {
	s.Message = &v
	return s
}

func (s *GetMediaConvertJobResponseBodyJob) SetOutputDetails(v []*MediaConvertOutputDetail) *GetMediaConvertJobResponseBodyJob {
	s.OutputDetails = v
	return s
}

func (s *GetMediaConvertJobResponseBodyJob) SetOutputGroupDetails(v []*MediaConvertOutputGroupDetail) *GetMediaConvertJobResponseBodyJob {
	s.OutputGroupDetails = v
	return s
}

func (s *GetMediaConvertJobResponseBodyJob) SetPercent(v int32) *GetMediaConvertJobResponseBodyJob {
	s.Percent = &v
	return s
}

func (s *GetMediaConvertJobResponseBodyJob) SetPipelineId(v string) *GetMediaConvertJobResponseBodyJob {
	s.PipelineId = &v
	return s
}

func (s *GetMediaConvertJobResponseBodyJob) SetRequestId(v string) *GetMediaConvertJobResponseBodyJob {
	s.RequestId = &v
	return s
}

func (s *GetMediaConvertJobResponseBodyJob) SetState(v string) *GetMediaConvertJobResponseBodyJob {
	s.State = &v
	return s
}

func (s *GetMediaConvertJobResponseBodyJob) SetUserData(v string) *GetMediaConvertJobResponseBodyJob {
	s.UserData = &v
	return s
}

func (s *GetMediaConvertJobResponseBodyJob) Validate() error {
	return dara.Validate(s)
}

type GetMediaConvertJobResponseBodyJobConfig struct {
	// The inputs of the transcoding task.
	Inputs  []*MediaConvertInput `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Repeated"`
	JobName *string              `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The output group configurations.
	OutputGroups []*MediaConvertOutputGroup `json:"OutputGroups,omitempty" xml:"OutputGroups,omitempty" type:"Repeated"`
	// The output configurations.
	Outputs []*MediaConvertOutput `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
}

func (s GetMediaConvertJobResponseBodyJobConfig) String() string {
	return dara.Prettify(s)
}

func (s GetMediaConvertJobResponseBodyJobConfig) GoString() string {
	return s.String()
}

func (s *GetMediaConvertJobResponseBodyJobConfig) GetInputs() []*MediaConvertInput {
	return s.Inputs
}

func (s *GetMediaConvertJobResponseBodyJobConfig) GetJobName() *string {
	return s.JobName
}

func (s *GetMediaConvertJobResponseBodyJobConfig) GetOutputGroups() []*MediaConvertOutputGroup {
	return s.OutputGroups
}

func (s *GetMediaConvertJobResponseBodyJobConfig) GetOutputs() []*MediaConvertOutput {
	return s.Outputs
}

func (s *GetMediaConvertJobResponseBodyJobConfig) SetInputs(v []*MediaConvertInput) *GetMediaConvertJobResponseBodyJobConfig {
	s.Inputs = v
	return s
}

func (s *GetMediaConvertJobResponseBodyJobConfig) SetJobName(v string) *GetMediaConvertJobResponseBodyJobConfig {
	s.JobName = &v
	return s
}

func (s *GetMediaConvertJobResponseBodyJobConfig) SetOutputGroups(v []*MediaConvertOutputGroup) *GetMediaConvertJobResponseBodyJobConfig {
	s.OutputGroups = v
	return s
}

func (s *GetMediaConvertJobResponseBodyJobConfig) SetOutputs(v []*MediaConvertOutput) *GetMediaConvertJobResponseBodyJobConfig {
	s.Outputs = v
	return s
}

func (s *GetMediaConvertJobResponseBodyJobConfig) Validate() error {
	return dara.Validate(s)
}
