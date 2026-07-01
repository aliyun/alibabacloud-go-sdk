// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitASRJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *SubmitASRJobRequest
	GetDescription() *string
	SetDuration(v string) *SubmitASRJobRequest
	GetDuration() *string
	SetEditingConfig(v string) *SubmitASRJobRequest
	GetEditingConfig() *string
	SetInputFile(v string) *SubmitASRJobRequest
	GetInputFile() *string
	SetStartTime(v string) *SubmitASRJobRequest
	GetStartTime() *string
	SetTitle(v string) *SubmitASRJobRequest
	GetTitle() *string
	SetUserData(v string) *SubmitASRJobRequest
	GetUserData() *string
}

type SubmitASRJobRequest struct {
	// The job description. The maximum length is 128 bytes.
	//
	// example:
	//
	// Test description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The duration of the segment to transcribe.
	//
	// example:
	//
	// 00:00:10
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The configuration for speech-to-text conversion:
	//
	// - HotwordLibraryIdList: A list of custom vocabulary IDs. Only one custom vocabulary ID is currently supported. Future releases will support multiple IDs.
	//
	// example:
	//
	// {
	//
	// 	"HotwordLibraryIdList": "******2609a14f54a0636b7e16******"
	//
	// }
	EditingConfig *string `json:"EditingConfig,omitempty" xml:"EditingConfig,omitempty"`
	// The input configuration. Specify either an OSS URL or a media ID from the media library.
	//
	// example:
	//
	// oss://example-bucket.oss-cn-shanghai.aliyuncs.com/example.mp4 or ****20b48fb04483915d4f2cd8ac****
	InputFile *string `json:"InputFile,omitempty" xml:"InputFile,omitempty"`
	// The start time of the segment to be transcribed from the media file.
	//
	// example:
	//
	// 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The job title. The maximum length is 128 bytes.
	//
	// example:
	//
	// Test title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// User-defined settings in JSON format. Use this to pass business-related data, such as the operating environment or other job details.
	//
	// example:
	//
	// {
	//
	//       "user": "data",
	//
	//       "env": "prod"
	//
	// }
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitASRJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitASRJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitASRJobRequest) GetDescription() *string {
	return s.Description
}

func (s *SubmitASRJobRequest) GetDuration() *string {
	return s.Duration
}

func (s *SubmitASRJobRequest) GetEditingConfig() *string {
	return s.EditingConfig
}

func (s *SubmitASRJobRequest) GetInputFile() *string {
	return s.InputFile
}

func (s *SubmitASRJobRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *SubmitASRJobRequest) GetTitle() *string {
	return s.Title
}

func (s *SubmitASRJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitASRJobRequest) SetDescription(v string) *SubmitASRJobRequest {
	s.Description = &v
	return s
}

func (s *SubmitASRJobRequest) SetDuration(v string) *SubmitASRJobRequest {
	s.Duration = &v
	return s
}

func (s *SubmitASRJobRequest) SetEditingConfig(v string) *SubmitASRJobRequest {
	s.EditingConfig = &v
	return s
}

func (s *SubmitASRJobRequest) SetInputFile(v string) *SubmitASRJobRequest {
	s.InputFile = &v
	return s
}

func (s *SubmitASRJobRequest) SetStartTime(v string) *SubmitASRJobRequest {
	s.StartTime = &v
	return s
}

func (s *SubmitASRJobRequest) SetTitle(v string) *SubmitASRJobRequest {
	s.Title = &v
	return s
}

func (s *SubmitASRJobRequest) SetUserData(v string) *SubmitASRJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitASRJobRequest) Validate() error {
	return dara.Validate(s)
}
