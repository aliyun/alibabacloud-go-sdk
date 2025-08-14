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
	// The job description, which can up to 128 bytes in length.
	//
	// example:
	//
	// 测试描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The speech duration.
	//
	// example:
	//
	// 00:00:10
	Duration      *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	EditingConfig *string `json:"EditingConfig,omitempty" xml:"EditingConfig,omitempty"`
	// The input file. You can specify an Object Storage Service (OSS) URL or the ID of a media asset in the media asset library.
	//
	// example:
	//
	// oss://example-bucket.oss-cn-shanghai.aliyuncs.com/example.mp4 或 ****20b48fb04483915d4f2cd8ac****
	InputFile *string `json:"InputFile,omitempty" xml:"InputFile,omitempty"`
	// The start time of the speech to recognize.
	//
	// example:
	//
	// 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The job title, which can be up to 128 bytes in length.
	//
	// example:
	//
	// 测试标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The user-defined data in the JSON format. You can specify your business information, such as the business environment and job information.
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
