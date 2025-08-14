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
	// example:
	//
	// 测试描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// {"AvatarId":"yunqiao"}
	EditingConfig *string `json:"EditingConfig,omitempty" xml:"EditingConfig,omitempty"`
	// The input configurations of the video rendering job for an avatar. You can specify text, the Object Storage Service (OSS) URL of an audio file, or the ID of a media asset. The audio file must be in the MP3 or WAV format.
	//
	// >  The text must be at least five words in length.
	InputConfig *string `json:"InputConfig,omitempty" xml:"InputConfig,omitempty"`
	// example:
	//
	// {"MediaURL":"https://your-bucket.oss-cn-shanghai.aliyuncs.com/xxx.mp4","Width":1920,"Height":1080}
	OutputConfig *string `json:"OutputConfig,omitempty" xml:"OutputConfig,omitempty"`
	// example:
	//
	// 测试标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
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
