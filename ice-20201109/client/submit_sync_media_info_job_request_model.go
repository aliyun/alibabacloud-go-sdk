// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSyncMediaInfoJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInput(v *SubmitSyncMediaInfoJobRequestInput) *SubmitSyncMediaInfoJobRequest
	GetInput() *SubmitSyncMediaInfoJobRequestInput
	SetName(v string) *SubmitSyncMediaInfoJobRequest
	GetName() *string
	SetScheduleConfig(v *SubmitSyncMediaInfoJobRequestScheduleConfig) *SubmitSyncMediaInfoJobRequest
	GetScheduleConfig() *SubmitSyncMediaInfoJobRequestScheduleConfig
	SetUserData(v string) *SubmitSyncMediaInfoJobRequest
	GetUserData() *string
}

type SubmitSyncMediaInfoJobRequest struct {
	// The input of the job.
	//
	// This parameter is required.
	Input *SubmitSyncMediaInfoJobRequestInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The job name.
	//
	// example:
	//
	// job-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The scheduling parameters. This parameter is optional.
	ScheduleConfig *SubmitSyncMediaInfoJobRequestScheduleConfig `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty" type:"Struct"`
	// The user data.
	//
	// example:
	//
	// user-data
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitSyncMediaInfoJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitSyncMediaInfoJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitSyncMediaInfoJobRequest) GetInput() *SubmitSyncMediaInfoJobRequestInput {
	return s.Input
}

func (s *SubmitSyncMediaInfoJobRequest) GetName() *string {
	return s.Name
}

func (s *SubmitSyncMediaInfoJobRequest) GetScheduleConfig() *SubmitSyncMediaInfoJobRequestScheduleConfig {
	return s.ScheduleConfig
}

func (s *SubmitSyncMediaInfoJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitSyncMediaInfoJobRequest) SetInput(v *SubmitSyncMediaInfoJobRequestInput) *SubmitSyncMediaInfoJobRequest {
	s.Input = v
	return s
}

func (s *SubmitSyncMediaInfoJobRequest) SetName(v string) *SubmitSyncMediaInfoJobRequest {
	s.Name = &v
	return s
}

func (s *SubmitSyncMediaInfoJobRequest) SetScheduleConfig(v *SubmitSyncMediaInfoJobRequestScheduleConfig) *SubmitSyncMediaInfoJobRequest {
	s.ScheduleConfig = v
	return s
}

func (s *SubmitSyncMediaInfoJobRequest) SetUserData(v string) *SubmitSyncMediaInfoJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitSyncMediaInfoJobRequest) Validate() error {
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

type SubmitSyncMediaInfoJobRequestInput struct {
	// The media object.
	//
	// 	- If Type is set to OSS, set this parameter to the URL of an OSS object. Both the OSS and HTTP protocols are supported.
	//
	// >  Before you use the OSS bucket in the URL, you must add the bucket on the [Storage Management](https://help.aliyun.com/document_detail/609918.html) page of the Intelligent Media Services (IMS) console.
	//
	// 	- If Type is set to Media, set this parameter to the ID of a media asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of the media object.
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

func (s SubmitSyncMediaInfoJobRequestInput) String() string {
	return dara.Prettify(s)
}

func (s SubmitSyncMediaInfoJobRequestInput) GoString() string {
	return s.String()
}

func (s *SubmitSyncMediaInfoJobRequestInput) GetMedia() *string {
	return s.Media
}

func (s *SubmitSyncMediaInfoJobRequestInput) GetType() *string {
	return s.Type
}

func (s *SubmitSyncMediaInfoJobRequestInput) SetMedia(v string) *SubmitSyncMediaInfoJobRequestInput {
	s.Media = &v
	return s
}

func (s *SubmitSyncMediaInfoJobRequestInput) SetType(v string) *SubmitSyncMediaInfoJobRequestInput {
	s.Type = &v
	return s
}

func (s *SubmitSyncMediaInfoJobRequestInput) Validate() error {
	return dara.Validate(s)
}

type SubmitSyncMediaInfoJobRequestScheduleConfig struct {
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

func (s SubmitSyncMediaInfoJobRequestScheduleConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitSyncMediaInfoJobRequestScheduleConfig) GoString() string {
	return s.String()
}

func (s *SubmitSyncMediaInfoJobRequestScheduleConfig) GetPipelineId() *string {
	return s.PipelineId
}

func (s *SubmitSyncMediaInfoJobRequestScheduleConfig) GetPriority() *int32 {
	return s.Priority
}

func (s *SubmitSyncMediaInfoJobRequestScheduleConfig) SetPipelineId(v string) *SubmitSyncMediaInfoJobRequestScheduleConfig {
	s.PipelineId = &v
	return s
}

func (s *SubmitSyncMediaInfoJobRequestScheduleConfig) SetPriority(v int32) *SubmitSyncMediaInfoJobRequestScheduleConfig {
	s.Priority = &v
	return s
}

func (s *SubmitSyncMediaInfoJobRequestScheduleConfig) Validate() error {
	return dara.Validate(s)
}
