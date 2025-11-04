// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitPackageJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInputs(v []*SubmitPackageJobRequestInputs) *SubmitPackageJobRequest
	GetInputs() []*SubmitPackageJobRequestInputs
	SetName(v string) *SubmitPackageJobRequest
	GetName() *string
	SetOutput(v *SubmitPackageJobRequestOutput) *SubmitPackageJobRequest
	GetOutput() *SubmitPackageJobRequestOutput
	SetScheduleConfig(v *SubmitPackageJobRequestScheduleConfig) *SubmitPackageJobRequest
	GetScheduleConfig() *SubmitPackageJobRequestScheduleConfig
	SetUserData(v string) *SubmitPackageJobRequest
	GetUserData() *string
}

type SubmitPackageJobRequest struct {
	// The input of the job.
	//
	// This parameter is required.
	Inputs []*SubmitPackageJobRequestInputs `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Repeated"`
	// The name of the job.
	//
	// example:
	//
	// job-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output of the job.
	//
	// This parameter is required.
	Output *SubmitPackageJobRequestOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// The scheduling settings.
	ScheduleConfig *SubmitPackageJobRequestScheduleConfig `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty" type:"Struct"`
	// The user-defined data.
	//
	// example:
	//
	// {"param": "value"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitPackageJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitPackageJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitPackageJobRequest) GetInputs() []*SubmitPackageJobRequestInputs {
	return s.Inputs
}

func (s *SubmitPackageJobRequest) GetName() *string {
	return s.Name
}

func (s *SubmitPackageJobRequest) GetOutput() *SubmitPackageJobRequestOutput {
	return s.Output
}

func (s *SubmitPackageJobRequest) GetScheduleConfig() *SubmitPackageJobRequestScheduleConfig {
	return s.ScheduleConfig
}

func (s *SubmitPackageJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitPackageJobRequest) SetInputs(v []*SubmitPackageJobRequestInputs) *SubmitPackageJobRequest {
	s.Inputs = v
	return s
}

func (s *SubmitPackageJobRequest) SetName(v string) *SubmitPackageJobRequest {
	s.Name = &v
	return s
}

func (s *SubmitPackageJobRequest) SetOutput(v *SubmitPackageJobRequestOutput) *SubmitPackageJobRequest {
	s.Output = v
	return s
}

func (s *SubmitPackageJobRequest) SetScheduleConfig(v *SubmitPackageJobRequestScheduleConfig) *SubmitPackageJobRequest {
	s.ScheduleConfig = v
	return s
}

func (s *SubmitPackageJobRequest) SetUserData(v string) *SubmitPackageJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitPackageJobRequest) Validate() error {
	if s.Inputs != nil {
		for _, item := range s.Inputs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Output != nil {
		if err := s.Output.Validate(); err != nil {
			return err
		}
	}
	if s.ScheduleConfig != nil {
		if err := s.ScheduleConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitPackageJobRequestInputs struct {
	// The information about the input stream file.
	//
	// This parameter is required.
	Input *SubmitPackageJobRequestInputsInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
}

func (s SubmitPackageJobRequestInputs) String() string {
	return dara.Prettify(s)
}

func (s SubmitPackageJobRequestInputs) GoString() string {
	return s.String()
}

func (s *SubmitPackageJobRequestInputs) GetInput() *SubmitPackageJobRequestInputsInput {
	return s.Input
}

func (s *SubmitPackageJobRequestInputs) SetInput(v *SubmitPackageJobRequestInputsInput) *SubmitPackageJobRequestInputs {
	s.Input = v
	return s
}

func (s *SubmitPackageJobRequestInputs) Validate() error {
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitPackageJobRequestInputsInput struct {
	// The media object.
	//
	// 	- If Type is set to OSS, set this parameter to the URL of an OSS object. Both the OSS and HTTP protocols are supported.
	//
	// 	- If Type is set to Media, set this parameter to the ID of a media asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of the media object. Valid values:
	//
	// 	- OSS: an Object Storage Service (OSS) object.
	//
	// 	- Media: a media asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitPackageJobRequestInputsInput) String() string {
	return dara.Prettify(s)
}

func (s SubmitPackageJobRequestInputsInput) GoString() string {
	return s.String()
}

func (s *SubmitPackageJobRequestInputsInput) GetMedia() *string {
	return s.Media
}

func (s *SubmitPackageJobRequestInputsInput) GetType() *string {
	return s.Type
}

func (s *SubmitPackageJobRequestInputsInput) SetMedia(v string) *SubmitPackageJobRequestInputsInput {
	s.Media = &v
	return s
}

func (s *SubmitPackageJobRequestInputsInput) SetType(v string) *SubmitPackageJobRequestInputsInput {
	s.Type = &v
	return s
}

func (s *SubmitPackageJobRequestInputsInput) Validate() error {
	return dara.Validate(s)
}

type SubmitPackageJobRequestOutput struct {
	// The media object. If Type is set to OSS, set this parameter to the URL of an OSS object. Both the OSS and HTTP protocols are supported. If Type is set to Media, set this parameter to the ID of a media asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of the media object. Valid values:
	//
	// 	- OSS: an OSS object.
	//
	// 	- Media: a media asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitPackageJobRequestOutput) String() string {
	return dara.Prettify(s)
}

func (s SubmitPackageJobRequestOutput) GoString() string {
	return s.String()
}

func (s *SubmitPackageJobRequestOutput) GetMedia() *string {
	return s.Media
}

func (s *SubmitPackageJobRequestOutput) GetType() *string {
	return s.Type
}

func (s *SubmitPackageJobRequestOutput) SetMedia(v string) *SubmitPackageJobRequestOutput {
	s.Media = &v
	return s
}

func (s *SubmitPackageJobRequestOutput) SetType(v string) *SubmitPackageJobRequestOutput {
	s.Type = &v
	return s
}

func (s *SubmitPackageJobRequestOutput) Validate() error {
	return dara.Validate(s)
}

type SubmitPackageJobRequestScheduleConfig struct {
	// The ID of the MPS queue to which the job was submitted.
	//
	// example:
	//
	// e37ebee5d98b4781897f6086e89f9c56
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The priority of the job. Valid values: 1 to 10. The greater the value, the higher the priority.
	//
	// example:
	//
	// 5
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s SubmitPackageJobRequestScheduleConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitPackageJobRequestScheduleConfig) GoString() string {
	return s.String()
}

func (s *SubmitPackageJobRequestScheduleConfig) GetPipelineId() *string {
	return s.PipelineId
}

func (s *SubmitPackageJobRequestScheduleConfig) GetPriority() *int32 {
	return s.Priority
}

func (s *SubmitPackageJobRequestScheduleConfig) SetPipelineId(v string) *SubmitPackageJobRequestScheduleConfig {
	s.PipelineId = &v
	return s
}

func (s *SubmitPackageJobRequestScheduleConfig) SetPriority(v int32) *SubmitPackageJobRequestScheduleConfig {
	s.Priority = &v
	return s
}

func (s *SubmitPackageJobRequestScheduleConfig) Validate() error {
	return dara.Validate(s)
}
