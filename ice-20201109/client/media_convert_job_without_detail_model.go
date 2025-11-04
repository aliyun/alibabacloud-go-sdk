// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMediaConvertJobWithoutDetail interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *MediaConvertJobWithoutDetail
	GetClientToken() *string
	SetCode(v string) *MediaConvertJobWithoutDetail
	GetCode() *string
	SetConfig(v *MediaConvertJobWithoutDetailConfig) *MediaConvertJobWithoutDetail
	GetConfig() *MediaConvertJobWithoutDetailConfig
	SetCreateTime(v string) *MediaConvertJobWithoutDetail
	GetCreateTime() *string
	SetFinishTime(v string) *MediaConvertJobWithoutDetail
	GetFinishTime() *string
	SetJobId(v string) *MediaConvertJobWithoutDetail
	GetJobId() *string
	SetMessage(v string) *MediaConvertJobWithoutDetail
	GetMessage() *string
	SetPipelineId(v string) *MediaConvertJobWithoutDetail
	GetPipelineId() *string
	SetRequestId(v string) *MediaConvertJobWithoutDetail
	GetRequestId() *string
	SetState(v string) *MediaConvertJobWithoutDetail
	GetState() *string
	SetUserData(v string) *MediaConvertJobWithoutDetail
	GetUserData() *string
}

type MediaConvertJobWithoutDetail struct {
	ClientToken *string                             `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Code        *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Config      *MediaConvertJobWithoutDetailConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	CreateTime  *string                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FinishTime  *string                             `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	JobId       *string                             `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Message     *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	PipelineId  *string                             `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	RequestId   *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	State       *string                             `json:"State,omitempty" xml:"State,omitempty"`
	UserData    *string                             `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s MediaConvertJobWithoutDetail) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertJobWithoutDetail) GoString() string {
	return s.String()
}

func (s *MediaConvertJobWithoutDetail) GetClientToken() *string {
	return s.ClientToken
}

func (s *MediaConvertJobWithoutDetail) GetCode() *string {
	return s.Code
}

func (s *MediaConvertJobWithoutDetail) GetConfig() *MediaConvertJobWithoutDetailConfig {
	return s.Config
}

func (s *MediaConvertJobWithoutDetail) GetCreateTime() *string {
	return s.CreateTime
}

func (s *MediaConvertJobWithoutDetail) GetFinishTime() *string {
	return s.FinishTime
}

func (s *MediaConvertJobWithoutDetail) GetJobId() *string {
	return s.JobId
}

func (s *MediaConvertJobWithoutDetail) GetMessage() *string {
	return s.Message
}

func (s *MediaConvertJobWithoutDetail) GetPipelineId() *string {
	return s.PipelineId
}

func (s *MediaConvertJobWithoutDetail) GetRequestId() *string {
	return s.RequestId
}

func (s *MediaConvertJobWithoutDetail) GetState() *string {
	return s.State
}

func (s *MediaConvertJobWithoutDetail) GetUserData() *string {
	return s.UserData
}

func (s *MediaConvertJobWithoutDetail) SetClientToken(v string) *MediaConvertJobWithoutDetail {
	s.ClientToken = &v
	return s
}

func (s *MediaConvertJobWithoutDetail) SetCode(v string) *MediaConvertJobWithoutDetail {
	s.Code = &v
	return s
}

func (s *MediaConvertJobWithoutDetail) SetConfig(v *MediaConvertJobWithoutDetailConfig) *MediaConvertJobWithoutDetail {
	s.Config = v
	return s
}

func (s *MediaConvertJobWithoutDetail) SetCreateTime(v string) *MediaConvertJobWithoutDetail {
	s.CreateTime = &v
	return s
}

func (s *MediaConvertJobWithoutDetail) SetFinishTime(v string) *MediaConvertJobWithoutDetail {
	s.FinishTime = &v
	return s
}

func (s *MediaConvertJobWithoutDetail) SetJobId(v string) *MediaConvertJobWithoutDetail {
	s.JobId = &v
	return s
}

func (s *MediaConvertJobWithoutDetail) SetMessage(v string) *MediaConvertJobWithoutDetail {
	s.Message = &v
	return s
}

func (s *MediaConvertJobWithoutDetail) SetPipelineId(v string) *MediaConvertJobWithoutDetail {
	s.PipelineId = &v
	return s
}

func (s *MediaConvertJobWithoutDetail) SetRequestId(v string) *MediaConvertJobWithoutDetail {
	s.RequestId = &v
	return s
}

func (s *MediaConvertJobWithoutDetail) SetState(v string) *MediaConvertJobWithoutDetail {
	s.State = &v
	return s
}

func (s *MediaConvertJobWithoutDetail) SetUserData(v string) *MediaConvertJobWithoutDetail {
	s.UserData = &v
	return s
}

func (s *MediaConvertJobWithoutDetail) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MediaConvertJobWithoutDetailConfig struct {
	Inputs       []*MediaConvertInput       `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Repeated"`
	JobName      *string                    `json:"JobName,omitempty" xml:"JobName,omitempty"`
	OutputGroups []*MediaConvertOutputGroup `json:"OutputGroups,omitempty" xml:"OutputGroups,omitempty" type:"Repeated"`
	Outputs      []*MediaConvertOutput      `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
}

func (s MediaConvertJobWithoutDetailConfig) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertJobWithoutDetailConfig) GoString() string {
	return s.String()
}

func (s *MediaConvertJobWithoutDetailConfig) GetInputs() []*MediaConvertInput {
	return s.Inputs
}

func (s *MediaConvertJobWithoutDetailConfig) GetJobName() *string {
	return s.JobName
}

func (s *MediaConvertJobWithoutDetailConfig) GetOutputGroups() []*MediaConvertOutputGroup {
	return s.OutputGroups
}

func (s *MediaConvertJobWithoutDetailConfig) GetOutputs() []*MediaConvertOutput {
	return s.Outputs
}

func (s *MediaConvertJobWithoutDetailConfig) SetInputs(v []*MediaConvertInput) *MediaConvertJobWithoutDetailConfig {
	s.Inputs = v
	return s
}

func (s *MediaConvertJobWithoutDetailConfig) SetJobName(v string) *MediaConvertJobWithoutDetailConfig {
	s.JobName = &v
	return s
}

func (s *MediaConvertJobWithoutDetailConfig) SetOutputGroups(v []*MediaConvertOutputGroup) *MediaConvertJobWithoutDetailConfig {
	s.OutputGroups = v
	return s
}

func (s *MediaConvertJobWithoutDetailConfig) SetOutputs(v []*MediaConvertOutput) *MediaConvertJobWithoutDetailConfig {
	s.Outputs = v
	return s
}

func (s *MediaConvertJobWithoutDetailConfig) Validate() error {
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
