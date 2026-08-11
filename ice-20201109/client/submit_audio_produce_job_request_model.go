// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAudioProduceJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *SubmitAudioProduceJobRequest
	GetDescription() *string
	SetEditingConfig(v string) *SubmitAudioProduceJobRequest
	GetEditingConfig() *string
	SetInputConfig(v string) *SubmitAudioProduceJobRequest
	GetInputConfig() *string
	SetOutputConfig(v string) *SubmitAudioProduceJobRequest
	GetOutputConfig() *string
	SetOverwrite(v bool) *SubmitAudioProduceJobRequest
	GetOverwrite() *bool
	SetTitle(v string) *SubmitAudioProduceJobRequest
	GetTitle() *string
	SetUserData(v string) *SubmitAudioProduceJobRequest
	GetUserData() *string
}

type SubmitAudioProduceJobRequest struct {
	// The task description:
	//
	// - Maximum length: 1024 bytes.
	//
	// - UTF-8 encoding.
	//
	// example:
	//
	// Task description, max 1024 bytes, UTF-8 encoded
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The audio production configuration:
	//
	// - voice: the [voice type](https://help.aliyun.com/document_detail/449563.html).
	//
	// - customizedVoice: the VoiceId for voice cloning.
	//
	// - format: the output file format. Valid values: PCM, WAV, and MP3.
	//
	// - volume: the volume. Valid values: 0 to 100. Default value: 50.
	//
	// - speech_rate: the speech rate. Valid values: -500 to 500. Default value: 0.
	//
	//     - [-500, 0, 500] corresponds to the speed multiplier range of [0.5, 1.0, 2.0].
	//
	//     - The calculation method is as follows:
	//
	//         - 0.8x speed: (1-1/0.8)/0.002 = -125
	//
	//         - 1.2x speed: (1-1/1.2)/0.001 = 166
	//
	//         - For speeds less than 1x, use the 0.002 coefficient.
	//
	//         - For speeds greater than 1x, use the 0.001 coefficient.
	//
	// - pitch_rate: the pitch. Valid values: -500 to 500. Default value: 0.
	//
	// <notice>If both voice and customizedVoice are specified, customizedVoice takes precedence.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"voice":"Siqi","format":"MP3","volume":50}
	EditingConfig *string `json:"EditingConfig,omitempty" xml:"EditingConfig,omitempty"`
	// The text content. A maximum of 10,000 Chinese characters is supported. [SSML markup language](https://help.aliyun.com/document_detail/2672807.html) is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// Audio production task
	InputConfig *string `json:"InputConfig,omitempty" xml:"InputConfig,omitempty"`
	// The audio output configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// For example, to store the output audio at http://my_bucket.oss-cn-shanghai.aliyuncs.com/target_audio.mp3, configure this parameter as:
	//
	// {
	//
	//       "bucket": "my_bucket",
	//
	//       "object": "target_audio"
	//
	// }
	OutputConfig *string `json:"OutputConfig,omitempty" xml:"OutputConfig,omitempty"`
	// Specifies whether to overwrite existing OSS files.
	//
	// example:
	//
	// true
	Overwrite *bool `json:"Overwrite,omitempty" xml:"Overwrite,omitempty"`
	// The task title. If not provided, a default title is automatically generated based on the date.
	//
	// - Maximum length: 128 bytes.
	//
	// - UTF-8 encoding.
	//
	// example:
	//
	// China Regional Daily News
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The custom settings in JSON format. Maximum length: 512 bytes. [Custom callback URL configuration](https://help.aliyun.com/document_detail/451631.html) is supported.
	//
	// example:
	//
	// {"NotifyAddress":"http://xx.xx.xxx"} or {"NotifyAddress":"https://xx.xx.xxx"} or {"NotifyAddress":"ice-callback-demo"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitAudioProduceJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitAudioProduceJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitAudioProduceJobRequest) GetDescription() *string {
	return s.Description
}

func (s *SubmitAudioProduceJobRequest) GetEditingConfig() *string {
	return s.EditingConfig
}

func (s *SubmitAudioProduceJobRequest) GetInputConfig() *string {
	return s.InputConfig
}

func (s *SubmitAudioProduceJobRequest) GetOutputConfig() *string {
	return s.OutputConfig
}

func (s *SubmitAudioProduceJobRequest) GetOverwrite() *bool {
	return s.Overwrite
}

func (s *SubmitAudioProduceJobRequest) GetTitle() *string {
	return s.Title
}

func (s *SubmitAudioProduceJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitAudioProduceJobRequest) SetDescription(v string) *SubmitAudioProduceJobRequest {
	s.Description = &v
	return s
}

func (s *SubmitAudioProduceJobRequest) SetEditingConfig(v string) *SubmitAudioProduceJobRequest {
	s.EditingConfig = &v
	return s
}

func (s *SubmitAudioProduceJobRequest) SetInputConfig(v string) *SubmitAudioProduceJobRequest {
	s.InputConfig = &v
	return s
}

func (s *SubmitAudioProduceJobRequest) SetOutputConfig(v string) *SubmitAudioProduceJobRequest {
	s.OutputConfig = &v
	return s
}

func (s *SubmitAudioProduceJobRequest) SetOverwrite(v bool) *SubmitAudioProduceJobRequest {
	s.Overwrite = &v
	return s
}

func (s *SubmitAudioProduceJobRequest) SetTitle(v string) *SubmitAudioProduceJobRequest {
	s.Title = &v
	return s
}

func (s *SubmitAudioProduceJobRequest) SetUserData(v string) *SubmitAudioProduceJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitAudioProduceJobRequest) Validate() error {
	return dara.Validate(s)
}
