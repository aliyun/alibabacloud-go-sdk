// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitMediaInfoJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInput(v *SubmitMediaInfoJobRequestInput) *SubmitMediaInfoJobRequest
	GetInput() *SubmitMediaInfoJobRequestInput
	SetName(v string) *SubmitMediaInfoJobRequest
	GetName() *string
	SetScheduleConfig(v *SubmitMediaInfoJobRequestScheduleConfig) *SubmitMediaInfoJobRequest
	GetScheduleConfig() *SubmitMediaInfoJobRequestScheduleConfig
	SetUserData(v string) *SubmitMediaInfoJobRequest
	GetUserData() *string
}

type SubmitMediaInfoJobRequest struct {
	// The input for the job.
	//
	// This parameter is required.
	//
	// example:
	//
	// job-name
	Input *SubmitMediaInfoJobRequestInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The job name.
	//
	// example:
	//
	// job-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The scheduling settings.
	//
	// example:
	//
	// user-data
	ScheduleConfig *SubmitMediaInfoJobRequestScheduleConfig `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty" type:"Struct"`
	// The custom user data.
	//
	// example:
	//
	// user-data
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitMediaInfoJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaInfoJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitMediaInfoJobRequest) GetInput() *SubmitMediaInfoJobRequestInput {
	return s.Input
}

func (s *SubmitMediaInfoJobRequest) GetName() *string {
	return s.Name
}

func (s *SubmitMediaInfoJobRequest) GetScheduleConfig() *SubmitMediaInfoJobRequestScheduleConfig {
	return s.ScheduleConfig
}

func (s *SubmitMediaInfoJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitMediaInfoJobRequest) SetInput(v *SubmitMediaInfoJobRequestInput) *SubmitMediaInfoJobRequest {
	s.Input = v
	return s
}

func (s *SubmitMediaInfoJobRequest) SetName(v string) *SubmitMediaInfoJobRequest {
	s.Name = &v
	return s
}

func (s *SubmitMediaInfoJobRequest) SetScheduleConfig(v *SubmitMediaInfoJobRequestScheduleConfig) *SubmitMediaInfoJobRequest {
	s.ScheduleConfig = v
	return s
}

func (s *SubmitMediaInfoJobRequest) SetUserData(v string) *SubmitMediaInfoJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitMediaInfoJobRequest) Validate() error {
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
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

type SubmitMediaInfoJobRequestInput struct {
	// The source of the input media:
	//
	// - If `Type` is `OSS`, set this parameter to the URL of the input file. You can use OSS (`oss://`), HTTP, or HTTPS URLs.
	//
	// > You must first add the OSS bucket specified in the URL to Intelligent Media Management Service (IMS) by using [Storage Management](https://help.aliyun.com/document_detail/609918.html).
	//
	// - If `Type` is `Media`, set this parameter to the media asset ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of the input media.
	//
	// - `OSS`: The input is an OSS file.
	//
	// - `Media`: The input is a media asset ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitMediaInfoJobRequestInput) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaInfoJobRequestInput) GoString() string {
	return s.String()
}

func (s *SubmitMediaInfoJobRequestInput) GetMedia() *string {
	return s.Media
}

func (s *SubmitMediaInfoJobRequestInput) GetType() *string {
	return s.Type
}

func (s *SubmitMediaInfoJobRequestInput) SetMedia(v string) *SubmitMediaInfoJobRequestInput {
	s.Media = &v
	return s
}

func (s *SubmitMediaInfoJobRequestInput) SetType(v string) *SubmitMediaInfoJobRequestInput {
	s.Type = &v
	return s
}

func (s *SubmitMediaInfoJobRequestInput) Validate() error {
	return dara.Validate(s)
}

type SubmitMediaInfoJobRequestScheduleConfig struct {
	// The pipeline ID.
	//
	// example:
	//
	// e37ebee5d98b4781897f6086e89f9c56
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The job priority. A higher value means a higher priority. Valid values range from 1 to 10.
	//
	// example:
	//
	// 5
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s SubmitMediaInfoJobRequestScheduleConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaInfoJobRequestScheduleConfig) GoString() string {
	return s.String()
}

func (s *SubmitMediaInfoJobRequestScheduleConfig) GetPipelineId() *string {
	return s.PipelineId
}

func (s *SubmitMediaInfoJobRequestScheduleConfig) GetPriority() *int32 {
	return s.Priority
}

func (s *SubmitMediaInfoJobRequestScheduleConfig) SetPipelineId(v string) *SubmitMediaInfoJobRequestScheduleConfig {
	s.PipelineId = &v
	return s
}

func (s *SubmitMediaInfoJobRequestScheduleConfig) SetPriority(v int32) *SubmitMediaInfoJobRequestScheduleConfig {
	s.Priority = &v
	return s
}

func (s *SubmitMediaInfoJobRequestScheduleConfig) Validate() error {
	return dara.Validate(s)
}
