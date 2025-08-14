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
	// The input of the job.
	//
	// This parameter is required.
	Input *SubmitMediaInfoJobRequestInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The job name.
	//
	// example:
	//
	// job-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The scheduling parameters.
	ScheduleConfig *SubmitMediaInfoJobRequestScheduleConfig `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty" type:"Struct"`
	// The user data.
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
	return dara.Validate(s)
}

type SubmitMediaInfoJobRequestInput struct {
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
	// The type of the media object. Valid values: OSS and Media. A value of OSS indicates an Object Storage Service (OSS) object. A value of Media indicates a media asset.
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
	// The ID of the ApsaraVideo Media Processing (MPS) queue that is used to run the job.
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
