// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAvatarVideoJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *SubmitAvatarVideoJobRequest
	GetDescription() *string
	SetEditingConfig(v string) *SubmitAvatarVideoJobRequest
	GetEditingConfig() *string
	SetInputConfig(v string) *SubmitAvatarVideoJobRequest
	GetInputConfig() *string
	SetOutputConfig(v string) *SubmitAvatarVideoJobRequest
	GetOutputConfig() *string
	SetTitle(v string) *SubmitAvatarVideoJobRequest
	GetTitle() *string
	SetUserData(v string) *SubmitAvatarVideoJobRequest
	GetUserData() *string
}

type SubmitAvatarVideoJobRequest struct {
	// The task description. Max length: 128 bytes.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The avatar configurations, including the avatar ID, voice, and speech rate.
	//
	// example:
	//
	// {"AvatarId":"yunqiao"}
	EditingConfig *string `json:"EditingConfig,omitempty" xml:"EditingConfig,omitempty"`
	// The input configurations of the video rendering task for an avatar. You can specify text, the Object Storage Service (OSS) URL of an audio file, or the ID of a media asset. The audio file must be in the MP3 or WAV format.
	//
	// 	Notice: The text must be at least five characters in length.
	//
	// example:
	//
	// {"Text": "To be, or not to be, that is the question."}
	InputConfig *string `json:"InputConfig,omitempty" xml:"InputConfig,omitempty"`
	// The output configurations, including the destination URL for the rendered video.
	//
	// example:
	//
	// {"MediaURL":"https://your-bucket.oss-cn-shanghai.aliyuncs.com/xxx.mp4"}
	OutputConfig *string `json:"OutputConfig,omitempty" xml:"OutputConfig,omitempty"`
	// The task name. Max length: 128 bytes.
	//
	// example:
	//
	// test
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// A user-defined JSON string for passing custom business information, such as environment details or task metadata.
	//
	// example:
	//
	// {"user":"data","env":"prod"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitAvatarVideoJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitAvatarVideoJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitAvatarVideoJobRequest) GetDescription() *string {
	return s.Description
}

func (s *SubmitAvatarVideoJobRequest) GetEditingConfig() *string {
	return s.EditingConfig
}

func (s *SubmitAvatarVideoJobRequest) GetInputConfig() *string {
	return s.InputConfig
}

func (s *SubmitAvatarVideoJobRequest) GetOutputConfig() *string {
	return s.OutputConfig
}

func (s *SubmitAvatarVideoJobRequest) GetTitle() *string {
	return s.Title
}

func (s *SubmitAvatarVideoJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitAvatarVideoJobRequest) SetDescription(v string) *SubmitAvatarVideoJobRequest {
	s.Description = &v
	return s
}

func (s *SubmitAvatarVideoJobRequest) SetEditingConfig(v string) *SubmitAvatarVideoJobRequest {
	s.EditingConfig = &v
	return s
}

func (s *SubmitAvatarVideoJobRequest) SetInputConfig(v string) *SubmitAvatarVideoJobRequest {
	s.InputConfig = &v
	return s
}

func (s *SubmitAvatarVideoJobRequest) SetOutputConfig(v string) *SubmitAvatarVideoJobRequest {
	s.OutputConfig = &v
	return s
}

func (s *SubmitAvatarVideoJobRequest) SetTitle(v string) *SubmitAvatarVideoJobRequest {
	s.Title = &v
	return s
}

func (s *SubmitAvatarVideoJobRequest) SetUserData(v string) *SubmitAvatarVideoJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitAvatarVideoJobRequest) Validate() error {
	return dara.Validate(s)
}
