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
	// The description of the job.
	//
	// - Cannot exceed 1,024 bytes.
	//
	// - Must be UTF-8 encoded.
	//
	// example:
	//
	// Task description, max 1024 bytes, UTF-8 encoded
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The audio production configuration:
	//
	// - `voice`: The [voice type](https://help.aliyun.com/document_detail/449563.html).
	//
	// - `customizedVoice`: The ID of the custom voice for voice cloning.
	//
	// - `format`: The output file format. Supported formats: `PCM`, `WAV`, and `MP3`.
	//
	// - `volume`: The volume. The value ranges from 0 to 100. Default: 50.
	//
	// - `speech_rate`: The speech rate. The value ranges from -500 to 500. Default: 0.
	//
	//   - Values of -500, 0, and 500 correspond to 0.5x, 1.0x, and 2.0x speed, respectively.
	//
	//   - Calculation method:
	//
	//     - For a 0.8x speed multiplier: (1 - 1/0.8) / 0.002 = -125.
	//
	//     - For a 1.2x speed multiplier: (1 - 1/1.2) / 0.001 = 166.
	//
	//     - For speed multipliers less than 1, use a factor of 0.002.
	//
	//     - For speed multipliers greater than 1, use a factor of 0.001.
	//
	// - `pitch_rate`: The pitch rate. The value ranges from -500 to 500. Default: 0.
	//
	//
	//   	Notice:
	//
	//   If you provide both `voice` and `customizedVoice`, `customizedVoice` takes precedence.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"voice":"Siqi","format":"MP3","volume":50}
	EditingConfig *string `json:"EditingConfig,omitempty" xml:"EditingConfig,omitempty"`
	// The text to synthesize. The maximum length is 10,000 characters. Supports [SSML](https://help.aliyun.com/document_detail/2672807.html).
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
	// Specifies whether to overwrite an existing OSS file.
	//
	// example:
	//
	// true
	Overwrite *bool `json:"Overwrite,omitempty" xml:"Overwrite,omitempty"`
	// The title of the job. If you do not provide a title, the system automatically generates one based on the current date.
	//
	// - Cannot exceed 128 bytes.
	//
	// - Must be UTF-8 encoded.
	//
	// example:
	//
	// China Regional Daily News
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// Custom settings in JSON format. The maximum length is 512 bytes. This parameter supports [custom callback address configuration](https://help.aliyun.com/document_detail/451631.html).
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
