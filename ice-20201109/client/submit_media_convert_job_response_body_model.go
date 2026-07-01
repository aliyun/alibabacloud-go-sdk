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
	// The media transcoding job.
	Job *SubmitMediaConvertJobResponseBodyJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Struct"`
	// The request ID.
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
	if s.Job != nil {
		if err := s.Job.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitMediaConvertJobResponseBodyJob struct {
	// The idempotency token for the request.
	//
	// example:
	//
	// FB7F25E9-AD9B-1603-8AF6-F1E42DF2E706
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The error code returned if the job fails.
	//
	// example:
	//
	// InvalidParameter.ResourceContentBad
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The job configuration.
	Config *SubmitMediaConvertJobResponseBodyJobConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// The job ID.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The error message returned if the job fails.
	//
	// example:
	//
	// The resource operated InputFile is bad
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The output details.
	OutputDetails []*MediaConvertOutputDetail `json:"OutputDetails,omitempty" xml:"OutputDetails,omitempty" type:"Repeated"`
	// The output group details.
	OutputGroupDetails []*MediaConvertOutputGroupDetail `json:"OutputGroupDetails,omitempty" xml:"OutputGroupDetails,omitempty" type:"Repeated"`
	// The pipeline ID.
	//
	// example:
	//
	// ***48a4edf410b908aecd91fc3b***
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A2129C9F-CE95-58B5-B8C1-07758FF6C86F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The job status. Valid values:
	//
	// - Inited: The job has been initialized.
	//
	// - Running: The job is running.
	//
	// - Success: The job has completed successfully.
	//
	// - Failed: The job has failed.
	//
	// - Canceled: The job has been canceled.
	//
	// example:
	//
	// Success
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
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	if s.OutputDetails != nil {
		for _, item := range s.OutputDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OutputGroupDetails != nil {
		for _, item := range s.OutputGroupDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SubmitMediaConvertJobResponseBodyJobConfig struct {
	// The job inputs.
	Inputs []*MediaConvertInput `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Repeated"`
	// The job name.
	//
	// example:
	//
	// Name
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
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
	if s.Inputs != nil {
		for _, item := range s.Inputs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OutputGroups != nil {
		for _, item := range s.OutputGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Outputs != nil {
		for _, item := range s.Outputs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
