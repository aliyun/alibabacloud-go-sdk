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
	// The job description.
	//
	// 	- The job description can be up to 1,024 bytes in length.
	//
	// 	- The value must be encoded in UTF-8.
	//
	// example:
	//
	// 任务描述  长度不超过1024字节  UTF8编码
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The audio editing configurations.
	//
	// 	- voice: the [voice type](https://help.aliyun.com/document_detail/449563.html).
	//
	// 	- customizedVoice: the ID of the personalized human voice.
	//
	// 	- format: the format of the output file. Valid values: PCM, WAV, and MP3.
	//
	// 	- volume: the volume. Default value: 50. Valid values: 0 to 100.
	//
	// 	- speech_rate: the speech tempo. Default value: 0. Value range: -500 to 500.
	//
	// 	- pitch_rate: the intonation. Default value: 0. Value range: -500 to 500.
	//
	// >  If you specify both voice and customizedVoice, customizedVoice takes precedence over voice.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"voice":"Siqi","format":"MP3","volume":50}
	EditingConfig *string `json:"EditingConfig,omitempty" xml:"EditingConfig,omitempty"`
	// The text content. A maximum of 2,000 characters are supported. The [Speech Synthesis Markup Language (SSML)](https://help.aliyun.com/document_detail/2672807.html) is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// 测试文本
	InputConfig *string `json:"InputConfig,omitempty" xml:"InputConfig,omitempty"`
	// The output audio configurations.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"bucket":"bucket","object":"objeck"}
	OutputConfig *string `json:"OutputConfig,omitempty" xml:"OutputConfig,omitempty"`
	// Specifies whether to overwrite the existing Object Storage Service (OSS) object.
	//
	// example:
	//
	// true
	Overwrite *bool `json:"Overwrite,omitempty" xml:"Overwrite,omitempty"`
	// The job title. If you do not specify this parameter, the system generates a title based on the current date.
	//
	// 	- The job title can be up to 128 bytes in length.
	//
	// 	- The value must be encoded in UTF-8.
	//
	// example:
	//
	// 任务标题。若不提供，根据日期自动生成默认title  长度不超过128字节  UTF8编码
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The user-defined data in the JSON format, which can be up to 512 bytes in length. You can specify a custom callback URL. For more information, see [Configure a callback upon editing completion](https://help.aliyun.com/document_detail/451631.html).
	//
	// example:
	//
	// {"user":"data"}
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
