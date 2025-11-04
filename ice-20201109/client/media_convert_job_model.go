// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMediaConvertJob interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *MediaConvertJob
	GetClientToken() *string
	SetCode(v string) *MediaConvertJob
	GetCode() *string
	SetConfig(v *MediaConvertJobConfig) *MediaConvertJob
	GetConfig() *MediaConvertJobConfig
	SetCreateTime(v string) *MediaConvertJob
	GetCreateTime() *string
	SetFinishTime(v string) *MediaConvertJob
	GetFinishTime() *string
	SetJobId(v string) *MediaConvertJob
	GetJobId() *string
	SetMessage(v string) *MediaConvertJob
	GetMessage() *string
	SetOutputDetails(v []*MediaConvertOutputDetail) *MediaConvertJob
	GetOutputDetails() []*MediaConvertOutputDetail
	SetOutputGroupDetails(v []*MediaConvertOutputGroupDetail) *MediaConvertJob
	GetOutputGroupDetails() []*MediaConvertOutputGroupDetail
	SetPercent(v int32) *MediaConvertJob
	GetPercent() *int32
	SetPipelineId(v string) *MediaConvertJob
	GetPipelineId() *string
	SetRequestId(v string) *MediaConvertJob
	GetRequestId() *string
	SetState(v string) *MediaConvertJob
	GetState() *string
	SetUserData(v string) *MediaConvertJob
	GetUserData() *string
}

type MediaConvertJob struct {
	ClientToken        *string                          `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Code               *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Config             *MediaConvertJobConfig           `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	CreateTime         *string                          `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FinishTime         *string                          `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	JobId              *string                          `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Message            *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	OutputDetails      []*MediaConvertOutputDetail      `json:"OutputDetails,omitempty" xml:"OutputDetails,omitempty" type:"Repeated"`
	OutputGroupDetails []*MediaConvertOutputGroupDetail `json:"OutputGroupDetails,omitempty" xml:"OutputGroupDetails,omitempty" type:"Repeated"`
	Percent            *int32                           `json:"Percent,omitempty" xml:"Percent,omitempty"`
	PipelineId         *string                          `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	RequestId          *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	State              *string                          `json:"State,omitempty" xml:"State,omitempty"`
	UserData           *string                          `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s MediaConvertJob) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertJob) GoString() string {
	return s.String()
}

func (s *MediaConvertJob) GetClientToken() *string {
	return s.ClientToken
}

func (s *MediaConvertJob) GetCode() *string {
	return s.Code
}

func (s *MediaConvertJob) GetConfig() *MediaConvertJobConfig {
	return s.Config
}

func (s *MediaConvertJob) GetCreateTime() *string {
	return s.CreateTime
}

func (s *MediaConvertJob) GetFinishTime() *string {
	return s.FinishTime
}

func (s *MediaConvertJob) GetJobId() *string {
	return s.JobId
}

func (s *MediaConvertJob) GetMessage() *string {
	return s.Message
}

func (s *MediaConvertJob) GetOutputDetails() []*MediaConvertOutputDetail {
	return s.OutputDetails
}

func (s *MediaConvertJob) GetOutputGroupDetails() []*MediaConvertOutputGroupDetail {
	return s.OutputGroupDetails
}

func (s *MediaConvertJob) GetPercent() *int32 {
	return s.Percent
}

func (s *MediaConvertJob) GetPipelineId() *string {
	return s.PipelineId
}

func (s *MediaConvertJob) GetRequestId() *string {
	return s.RequestId
}

func (s *MediaConvertJob) GetState() *string {
	return s.State
}

func (s *MediaConvertJob) GetUserData() *string {
	return s.UserData
}

func (s *MediaConvertJob) SetClientToken(v string) *MediaConvertJob {
	s.ClientToken = &v
	return s
}

func (s *MediaConvertJob) SetCode(v string) *MediaConvertJob {
	s.Code = &v
	return s
}

func (s *MediaConvertJob) SetConfig(v *MediaConvertJobConfig) *MediaConvertJob {
	s.Config = v
	return s
}

func (s *MediaConvertJob) SetCreateTime(v string) *MediaConvertJob {
	s.CreateTime = &v
	return s
}

func (s *MediaConvertJob) SetFinishTime(v string) *MediaConvertJob {
	s.FinishTime = &v
	return s
}

func (s *MediaConvertJob) SetJobId(v string) *MediaConvertJob {
	s.JobId = &v
	return s
}

func (s *MediaConvertJob) SetMessage(v string) *MediaConvertJob {
	s.Message = &v
	return s
}

func (s *MediaConvertJob) SetOutputDetails(v []*MediaConvertOutputDetail) *MediaConvertJob {
	s.OutputDetails = v
	return s
}

func (s *MediaConvertJob) SetOutputGroupDetails(v []*MediaConvertOutputGroupDetail) *MediaConvertJob {
	s.OutputGroupDetails = v
	return s
}

func (s *MediaConvertJob) SetPercent(v int32) *MediaConvertJob {
	s.Percent = &v
	return s
}

func (s *MediaConvertJob) SetPipelineId(v string) *MediaConvertJob {
	s.PipelineId = &v
	return s
}

func (s *MediaConvertJob) SetRequestId(v string) *MediaConvertJob {
	s.RequestId = &v
	return s
}

func (s *MediaConvertJob) SetState(v string) *MediaConvertJob {
	s.State = &v
	return s
}

func (s *MediaConvertJob) SetUserData(v string) *MediaConvertJob {
	s.UserData = &v
	return s
}

func (s *MediaConvertJob) Validate() error {
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

type MediaConvertJobConfig struct {
	Inputs       []*MediaConvertInput       `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Repeated"`
	JobName      *string                    `json:"JobName,omitempty" xml:"JobName,omitempty"`
	OutputGroups []*MediaConvertOutputGroup `json:"OutputGroups,omitempty" xml:"OutputGroups,omitempty" type:"Repeated"`
	Outputs      []*MediaConvertOutput      `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
}

func (s MediaConvertJobConfig) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertJobConfig) GoString() string {
	return s.String()
}

func (s *MediaConvertJobConfig) GetInputs() []*MediaConvertInput {
	return s.Inputs
}

func (s *MediaConvertJobConfig) GetJobName() *string {
	return s.JobName
}

func (s *MediaConvertJobConfig) GetOutputGroups() []*MediaConvertOutputGroup {
	return s.OutputGroups
}

func (s *MediaConvertJobConfig) GetOutputs() []*MediaConvertOutput {
	return s.Outputs
}

func (s *MediaConvertJobConfig) SetInputs(v []*MediaConvertInput) *MediaConvertJobConfig {
	s.Inputs = v
	return s
}

func (s *MediaConvertJobConfig) SetJobName(v string) *MediaConvertJobConfig {
	s.JobName = &v
	return s
}

func (s *MediaConvertJobConfig) SetOutputGroups(v []*MediaConvertOutputGroup) *MediaConvertJobConfig {
	s.OutputGroups = v
	return s
}

func (s *MediaConvertJobConfig) SetOutputs(v []*MediaConvertOutput) *MediaConvertJobConfig {
	s.Outputs = v
	return s
}

func (s *MediaConvertJobConfig) Validate() error {
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
