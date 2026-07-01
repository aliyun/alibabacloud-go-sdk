// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTranscodeJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *SubmitTranscodeJobRequest
	GetClientToken() *string
	SetInputGroup(v []*SubmitTranscodeJobRequestInputGroup) *SubmitTranscodeJobRequest
	GetInputGroup() []*SubmitTranscodeJobRequestInputGroup
	SetName(v string) *SubmitTranscodeJobRequest
	GetName() *string
	SetOutputGroup(v []*SubmitTranscodeJobRequestOutputGroup) *SubmitTranscodeJobRequest
	GetOutputGroup() []*SubmitTranscodeJobRequestOutputGroup
	SetScheduleConfig(v *SubmitTranscodeJobRequestScheduleConfig) *SubmitTranscodeJobRequest
	GetScheduleConfig() *SubmitTranscodeJobRequestScheduleConfig
	SetUserData(v string) *SubmitTranscodeJobRequest
	GetUserData() *string
}

type SubmitTranscodeJobRequest struct {
	// The idempotence key. Ensures request idempotence.
	//
	// example:
	//
	// ****12e8864746a0a398****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The input group for the job. A single input creates a transcoding job. Multiple inputs create a media merging job.
	//
	// This parameter is required.
	//
	// example:
	//
	// job-name
	InputGroup []*SubmitTranscodeJobRequestInputGroup `json:"InputGroup,omitempty" xml:"InputGroup,omitempty" type:"Repeated"`
	// The job name.
	//
	// example:
	//
	// job-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output group for the job.
	//
	// This parameter is required.
	//
	// example:
	//
	// user-data
	OutputGroup []*SubmitTranscodeJobRequestOutputGroup `json:"OutputGroup,omitempty" xml:"OutputGroup,omitempty" type:"Repeated"`
	// The job scheduling information.
	//
	// example:
	//
	// job-name
	ScheduleConfig *SubmitTranscodeJobRequestScheduleConfig `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty" type:"Struct"`
	// Custom settings in JSON format. The length is limited to 512 bytes. Supports [custom webhook address configuration](https://help.aliyun.com/document_detail/451631.html).
	//
	// example:
	//
	// user-data
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitTranscodeJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SubmitTranscodeJobRequest) GetInputGroup() []*SubmitTranscodeJobRequestInputGroup {
	return s.InputGroup
}

func (s *SubmitTranscodeJobRequest) GetName() *string {
	return s.Name
}

func (s *SubmitTranscodeJobRequest) GetOutputGroup() []*SubmitTranscodeJobRequestOutputGroup {
	return s.OutputGroup
}

func (s *SubmitTranscodeJobRequest) GetScheduleConfig() *SubmitTranscodeJobRequestScheduleConfig {
	return s.ScheduleConfig
}

func (s *SubmitTranscodeJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitTranscodeJobRequest) SetClientToken(v string) *SubmitTranscodeJobRequest {
	s.ClientToken = &v
	return s
}

func (s *SubmitTranscodeJobRequest) SetInputGroup(v []*SubmitTranscodeJobRequestInputGroup) *SubmitTranscodeJobRequest {
	s.InputGroup = v
	return s
}

func (s *SubmitTranscodeJobRequest) SetName(v string) *SubmitTranscodeJobRequest {
	s.Name = &v
	return s
}

func (s *SubmitTranscodeJobRequest) SetOutputGroup(v []*SubmitTranscodeJobRequestOutputGroup) *SubmitTranscodeJobRequest {
	s.OutputGroup = v
	return s
}

func (s *SubmitTranscodeJobRequest) SetScheduleConfig(v *SubmitTranscodeJobRequestScheduleConfig) *SubmitTranscodeJobRequest {
	s.ScheduleConfig = v
	return s
}

func (s *SubmitTranscodeJobRequest) SetUserData(v string) *SubmitTranscodeJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitTranscodeJobRequest) Validate() error {
	if s.InputGroup != nil {
		for _, item := range s.InputGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OutputGroup != nil {
		for _, item := range s.OutputGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ScheduleConfig != nil {
		if err := s.ScheduleConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitTranscodeJobRequestInputGroup struct {
	// The input stream path:
	//
	// - This parameter takes effect only when Type is Media. It lets you select a specific file from the media asset as the input.
	//
	// - The system checks whether the input URL exists in the media asset.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
	InputUrl *string `json:"InputUrl,omitempty" xml:"InputUrl,omitempty"`
	// The media value:
	//
	// - If Type is OSS, this is a URL that supports the OSS or HTTP protocol.
	//
	// > You must add the OSS bucket in the URL to IMS [storage management](https://help.aliyun.com/document_detail/609918.html) before you use it.
	//
	// - If Type is Media, this is a media asset ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The media object type. Valid values:
	//
	// - OSS: An OSS file.
	//
	// - Media: A media asset ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitTranscodeJobRequestInputGroup) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobRequestInputGroup) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobRequestInputGroup) GetInputUrl() *string {
	return s.InputUrl
}

func (s *SubmitTranscodeJobRequestInputGroup) GetMedia() *string {
	return s.Media
}

func (s *SubmitTranscodeJobRequestInputGroup) GetType() *string {
	return s.Type
}

func (s *SubmitTranscodeJobRequestInputGroup) SetInputUrl(v string) *SubmitTranscodeJobRequestInputGroup {
	s.InputUrl = &v
	return s
}

func (s *SubmitTranscodeJobRequestInputGroup) SetMedia(v string) *SubmitTranscodeJobRequestInputGroup {
	s.Media = &v
	return s
}

func (s *SubmitTranscodeJobRequestInputGroup) SetType(v string) *SubmitTranscodeJobRequestInputGroup {
	s.Type = &v
	return s
}

func (s *SubmitTranscodeJobRequestInputGroup) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobRequestOutputGroup struct {
	// The output media configuration.
	//
	// This parameter is required.
	Output *SubmitTranscodeJobRequestOutputGroupOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// The job processing configuration.
	//
	// This parameter is required.
	ProcessConfig *SubmitTranscodeJobRequestOutputGroupProcessConfig `json:"ProcessConfig,omitempty" xml:"ProcessConfig,omitempty" type:"Struct"`
}

func (s SubmitTranscodeJobRequestOutputGroup) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobRequestOutputGroup) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobRequestOutputGroup) GetOutput() *SubmitTranscodeJobRequestOutputGroupOutput {
	return s.Output
}

func (s *SubmitTranscodeJobRequestOutputGroup) GetProcessConfig() *SubmitTranscodeJobRequestOutputGroupProcessConfig {
	return s.ProcessConfig
}

func (s *SubmitTranscodeJobRequestOutputGroup) SetOutput(v *SubmitTranscodeJobRequestOutputGroupOutput) *SubmitTranscodeJobRequestOutputGroup {
	s.Output = v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroup) SetProcessConfig(v *SubmitTranscodeJobRequestOutputGroupProcessConfig) *SubmitTranscodeJobRequestOutputGroup {
	s.ProcessConfig = v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroup) Validate() error {
	if s.Output != nil {
		if err := s.Output.Validate(); err != nil {
			return err
		}
	}
	if s.ProcessConfig != nil {
		if err := s.ProcessConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitTranscodeJobRequestOutputGroupOutput struct {
	// The media value:
	//
	// - If Type is OSS, this is a URL that supports the OSS or HTTP protocol.
	//
	// > You must add the OSS bucket in the URL to IMS [storage management](https://help.aliyun.com/document_detail/609918.html) before you use it.
	//
	// - If Type is Media, this is a media asset ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The output stream path:<br>
	//
	// This parameter takes effect only when Type is Media. It lets you select a specific file from the media asset as the output.<br>
	//
	// Valid placeholders:<br><br>
	//
	// - {MediaId}: The media asset ID.
	//
	// - {JobId}: The sub-job ID.
	//
	// - {MediaBucket}: The bucket where the media asset resides.
	//
	// - {ExtName}: The file extension. This is the output format specified in the transcoding template.
	//
	// - {DestMd5}: The MD5 hash of the transcoded output file.
	//
	// > 1. You must include both {MediaId} and {JobId} in this parameter.
	//
	// >
	//
	// > 2. The output bucket must be the same as the bucket where the media asset resides.
	//
	// example:
	//
	// oss://bucket/path/to/{MediaId}/{JobId}.mp4
	OutputUrl *string `json:"OutputUrl,omitempty" xml:"OutputUrl,omitempty"`
	// The media object type. Valid values:
	//
	// - OSS: An OSS file.
	//
	// - Media: A media asset ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitTranscodeJobRequestOutputGroupOutput) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobRequestOutputGroupOutput) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobRequestOutputGroupOutput) GetMedia() *string {
	return s.Media
}

func (s *SubmitTranscodeJobRequestOutputGroupOutput) GetOutputUrl() *string {
	return s.OutputUrl
}

func (s *SubmitTranscodeJobRequestOutputGroupOutput) GetType() *string {
	return s.Type
}

func (s *SubmitTranscodeJobRequestOutputGroupOutput) SetMedia(v string) *SubmitTranscodeJobRequestOutputGroupOutput {
	s.Media = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupOutput) SetOutputUrl(v string) *SubmitTranscodeJobRequestOutputGroupOutput {
	s.OutputUrl = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupOutput) SetType(v string) *SubmitTranscodeJobRequestOutputGroupOutput {
	s.Type = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupOutput) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobRequestOutputGroupProcessConfig struct {
	// The configuration for merging multiple inputs.
	CombineConfigs []*SubmitTranscodeJobRequestOutputGroupProcessConfigCombineConfigs `json:"CombineConfigs,omitempty" xml:"CombineConfigs,omitempty" type:"Repeated"`
	// The encryption configuration.
	Encryption *SubmitTranscodeJobRequestOutputGroupProcessConfigEncryption `json:"Encryption,omitempty" xml:"Encryption,omitempty" type:"Struct"`
	// The image watermark configuration.
	ImageWatermarks []*SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarks `json:"ImageWatermarks,omitempty" xml:"ImageWatermarks,omitempty" type:"Repeated"`
	// The subtitle burn-in configuration.
	Subtitles []*SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitles `json:"Subtitles,omitempty" xml:"Subtitles,omitempty" type:"Repeated"`
	// The text watermark configuration.
	TextWatermarks []*SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarks `json:"TextWatermarks,omitempty" xml:"TextWatermarks,omitempty" type:"Repeated"`
	// The transcoding configuration.
	//
	// This parameter is required.
	Transcode *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscode `json:"Transcode,omitempty" xml:"Transcode,omitempty" type:"Struct"`
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfig) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfig) GetCombineConfigs() []*SubmitTranscodeJobRequestOutputGroupProcessConfigCombineConfigs {
	return s.CombineConfigs
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfig) GetEncryption() *SubmitTranscodeJobRequestOutputGroupProcessConfigEncryption {
	return s.Encryption
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfig) GetImageWatermarks() []*SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarks {
	return s.ImageWatermarks
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfig) GetSubtitles() []*SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitles {
	return s.Subtitles
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfig) GetTextWatermarks() []*SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarks {
	return s.TextWatermarks
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfig) GetTranscode() *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscode {
	return s.Transcode
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfig) SetCombineConfigs(v []*SubmitTranscodeJobRequestOutputGroupProcessConfigCombineConfigs) *SubmitTranscodeJobRequestOutputGroupProcessConfig {
	s.CombineConfigs = v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfig) SetEncryption(v *SubmitTranscodeJobRequestOutputGroupProcessConfigEncryption) *SubmitTranscodeJobRequestOutputGroupProcessConfig {
	s.Encryption = v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfig) SetImageWatermarks(v []*SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarks) *SubmitTranscodeJobRequestOutputGroupProcessConfig {
	s.ImageWatermarks = v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfig) SetSubtitles(v []*SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitles) *SubmitTranscodeJobRequestOutputGroupProcessConfig {
	s.Subtitles = v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfig) SetTextWatermarks(v []*SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarks) *SubmitTranscodeJobRequestOutputGroupProcessConfig {
	s.TextWatermarks = v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfig) SetTranscode(v *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscode) *SubmitTranscodeJobRequestOutputGroupProcessConfig {
	s.Transcode = v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfig) Validate() error {
	if s.CombineConfigs != nil {
		for _, item := range s.CombineConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Encryption != nil {
		if err := s.Encryption.Validate(); err != nil {
			return err
		}
	}
	if s.ImageWatermarks != nil {
		for _, item := range s.ImageWatermarks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Subtitles != nil {
		for _, item := range s.Subtitles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TextWatermarks != nil {
		for _, item := range s.TextWatermarks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Transcode != nil {
		if err := s.Transcode.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitTranscodeJobRequestOutputGroupProcessConfigCombineConfigs struct {
	// The audio stream index.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0 or exclude
	AudioIndex *string `json:"AudioIndex,omitempty" xml:"AudioIndex,omitempty"`
	// The duration of the input stream. Default: The video duration.
	//
	// example:
	//
	// 20.0
	Duration *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The start time of the input stream. Default: 0.
	//
	// example:
	//
	// 0.0
	Start *float64 `json:"Start,omitempty" xml:"Start,omitempty"`
	// The video stream index.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0 or exclude
	VideoIndex *string `json:"VideoIndex,omitempty" xml:"VideoIndex,omitempty"`
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigCombineConfigs) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigCombineConfigs) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigCombineConfigs) GetAudioIndex() *string {
	return s.AudioIndex
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigCombineConfigs) GetDuration() *float64 {
	return s.Duration
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigCombineConfigs) GetStart() *float64 {
	return s.Start
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigCombineConfigs) GetVideoIndex() *string {
	return s.VideoIndex
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigCombineConfigs) SetAudioIndex(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigCombineConfigs {
	s.AudioIndex = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigCombineConfigs) SetDuration(v float64) *SubmitTranscodeJobRequestOutputGroupProcessConfigCombineConfigs {
	s.Duration = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigCombineConfigs) SetStart(v float64) *SubmitTranscodeJobRequestOutputGroupProcessConfigCombineConfigs {
	s.Start = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigCombineConfigs) SetVideoIndex(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigCombineConfigs {
	s.VideoIndex = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigCombineConfigs) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobRequestOutputGroupProcessConfigEncryption struct {
	// The ciphertext of the key for standard encryption.
	//
	// example:
	//
	// MTYi00NDU0LTg5O****
	CipherText *string `json:"CipherText,omitempty" xml:"CipherText,omitempty"`
	// The decryption service endpoint for standard encryption.
	//
	// example:
	//
	// https://sample.com/path?CipherText=MTYi00NDU0LTg5O****
	DecryptKeyUri *string `json:"DecryptKeyUri,omitempty" xml:"DecryptKeyUri,omitempty"`
	// The encryption type. Valid values:
	//
	// - PrivateEncryption: Alibaba Cloud proprietary cryptography.
	//
	// - HLSEncryption: HLS standard encryption.
	//
	// example:
	//
	// PrivateEncryption
	EncryptType *string `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	// The key service type for standard encryption. Valid values:
	//
	// - KMS
	//
	// - Base64
	//
	// example:
	//
	// KMS
	KeyServiceType *string `json:"KeyServiceType,omitempty" xml:"KeyServiceType,omitempty"`
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigEncryption) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigEncryption) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigEncryption) GetCipherText() *string {
	return s.CipherText
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigEncryption) GetDecryptKeyUri() *string {
	return s.DecryptKeyUri
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigEncryption) GetEncryptType() *string {
	return s.EncryptType
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigEncryption) GetKeyServiceType() *string {
	return s.KeyServiceType
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigEncryption) SetCipherText(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigEncryption {
	s.CipherText = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigEncryption) SetDecryptKeyUri(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigEncryption {
	s.DecryptKeyUri = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigEncryption) SetEncryptType(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigEncryption {
	s.EncryptType = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigEncryption) SetKeyServiceType(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigEncryption {
	s.KeyServiceType = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigEncryption) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarks struct {
	// Parameters to overwrite. If you specify these, they replace the corresponding parameters in the template.
	OverwriteParams *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParams `json:"OverwriteParams,omitempty" xml:"OverwriteParams,omitempty" type:"Struct"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9547c6ad97cb4f2aaa29683ebd18d410
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarks) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarks) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarks) GetOverwriteParams() *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParams {
	return s.OverwriteParams
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarks) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarks) SetOverwriteParams(v *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParams) *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarks {
	s.OverwriteParams = v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarks) SetTemplateId(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarks {
	s.TemplateId = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarks) Validate() error {
	if s.OverwriteParams != nil {
		if err := s.OverwriteParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParams struct {
	// The horizontal offset of the image watermark relative to the output video. Default value: 0.
	//
	// Values can be one of the following:
	//
	// - Integer: The offset in pixels.
	//
	//   - Valid values: [8, 4096]
	//
	//   - Unit: px
	//
	// - Decimal: The ratio of the horizontal offset to the output video width.
	//
	//   - Valid values: (0, 1)
	//
	//   - Up to four decimal places are supported, such as 0.9999. Extra digits are automatically discarded.
	//
	// example:
	//
	// 10
	Dx *string `json:"Dx,omitempty" xml:"Dx,omitempty"`
	// The vertical offset of the image watermark relative to the output video. Default value: 0.
	//
	// Values can be one of the following:
	//
	// - Integer: The offset in pixels.
	//
	//   - Valid values: [8, 4096]
	//
	//   - Unit: px
	//
	// - Decimal: The ratio of the vertical offset to the output video height.
	//
	//   - Valid values: (0, 1)
	//
	//   - Up to four decimal places are supported, such as 0.9999. Extra digits are automatically discarded.
	//
	// example:
	//
	// 10
	Dy *string `json:"Dy,omitempty" xml:"Dy,omitempty"`
	// The watermark image file.
	File *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParamsFile `json:"File,omitempty" xml:"File,omitempty" type:"Struct"`
	// The height of the image watermark on the output video. Values can be one of the following:
	//
	// - Integer: The pixel height of the watermark image.
	//
	//   - Valid values: [8, 4096]
	//
	//   - Unit: px
	//
	// - Decimal: The ratio of the watermark height to the output video height.
	//
	//   - Valid values: (0, 1)
	//
	//   - Up to four decimal places are supported, such as 0.9999. Extra digits are automatically discarded.
	//
	// example:
	//
	// 32
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The position of the watermark.
	//
	// - Valid values: TopRight (top-right), TopLeft (top-left), BottomRight (bottom-right), and BottomLeft (bottom-left).
	//
	// - Default value: TopRight.
	//
	// example:
	//
	// TopLeft
	ReferPos *string `json:"ReferPos,omitempty" xml:"ReferPos,omitempty"`
	// The display time settings for a dynamic watermark.
	Timeline *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline `json:"Timeline,omitempty" xml:"Timeline,omitempty" type:"Struct"`
	// The width of the image watermark on the output video. Values can be one of the following:
	//
	// - Integer: The pixel width of the watermark image.
	//
	//   - Valid values: [8, 4096]
	//
	//   - Unit: px
	//
	// - Decimal: The ratio of the watermark width to the output video width.
	//
	//   - Valid values: (0, 1)
	//
	//   - Up to four decimal places are supported, such as 0.9999. Extra digits are automatically discarded.
	//
	// example:
	//
	// 32
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParams) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParams) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParams) GetDx() *string {
	return s.Dx
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParams) GetDy() *string {
	return s.Dy
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParams) GetFile() *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParamsFile {
	return s.File
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParams) GetHeight() *string {
	return s.Height
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParams) GetReferPos() *string {
	return s.ReferPos
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParams) GetTimeline() *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline {
	return s.Timeline
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParams) GetWidth() *string {
	return s.Width
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParams) SetDx(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParams {
	s.Dx = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParams) SetDy(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParams {
	s.Dy = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParams) SetFile(v *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParamsFile) *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParams {
	s.File = v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParams) SetHeight(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParams {
	s.Height = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParams) SetReferPos(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParams {
	s.ReferPos = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParams) SetTimeline(v *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline) *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParams {
	s.Timeline = v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParams) SetWidth(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParams {
	s.Width = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParams) Validate() error {
	if s.File != nil {
		if err := s.File.Validate(); err != nil {
			return err
		}
	}
	if s.Timeline != nil {
		if err := s.Timeline.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParamsFile struct {
	// The media value:
	//
	// - If Type is OSS, this is a URL that supports the OSS or HTTP protocol.
	//
	// - If Type is Media, this is a media asset ID.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The media object type. Valid values:
	//
	// - OSS: An OSS file.
	//
	// - Media: A media asset ID.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParamsFile) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParamsFile) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParamsFile) GetMedia() *string {
	return s.Media
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParamsFile) GetType() *string {
	return s.Type
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParamsFile) SetMedia(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParamsFile {
	s.Media = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParamsFile) SetType(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParamsFile {
	s.Type = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParamsFile) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline struct {
	// The duration of the watermark.
	//
	// - Valid values: [number, ToEND]
	//
	// - Default value: ToEND
	//
	// example:
	//
	// ToEND
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The start time of the watermark.
	//
	// - Unit: seconds
	//
	// - Valid values: numbers
	//
	// - Default value: 0
	//
	// example:
	//
	// 00:00:05
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline) GetDuration() *string {
	return s.Duration
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline) GetStart() *string {
	return s.Start
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline) SetDuration(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline {
	s.Duration = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline) SetStart(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline {
	s.Start = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitles struct {
	// Parameters to overwrite. If you specify these, they replace the corresponding parameters in the template.
	OverwriteParams *SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitlesOverwriteParams `json:"OverwriteParams,omitempty" xml:"OverwriteParams,omitempty" type:"Struct"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9547c6ad97cb4f2aaa29683ebd18d410
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitles) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitles) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitles) GetOverwriteParams() *SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitlesOverwriteParams {
	return s.OverwriteParams
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitles) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitles) SetOverwriteParams(v *SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitlesOverwriteParams) *SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitles {
	s.OverwriteParams = v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitles) SetTemplateId(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitles {
	s.TemplateId = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitles) Validate() error {
	if s.OverwriteParams != nil {
		if err := s.OverwriteParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitlesOverwriteParams struct {
	// The file encoding format.
	//
	// example:
	//
	// UTF-8
	CharEnc *string `json:"CharEnc,omitempty" xml:"CharEnc,omitempty"`
	// The subtitle file.
	File *SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitlesOverwriteParamsFile `json:"File,omitempty" xml:"File,omitempty" type:"Struct"`
	// The subtitle file format.
	//
	// example:
	//
	// vtt
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitlesOverwriteParams) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitlesOverwriteParams) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitlesOverwriteParams) GetCharEnc() *string {
	return s.CharEnc
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitlesOverwriteParams) GetFile() *SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitlesOverwriteParamsFile {
	return s.File
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitlesOverwriteParams) GetFormat() *string {
	return s.Format
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitlesOverwriteParams) SetCharEnc(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitlesOverwriteParams {
	s.CharEnc = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitlesOverwriteParams) SetFile(v *SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitlesOverwriteParamsFile) *SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitlesOverwriteParams {
	s.File = v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitlesOverwriteParams) SetFormat(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitlesOverwriteParams {
	s.Format = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitlesOverwriteParams) Validate() error {
	if s.File != nil {
		if err := s.File.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitlesOverwriteParamsFile struct {
	// The media value:
	//
	// - If Type is OSS, this is a URL that supports the OSS or HTTP protocol.
	//
	// - If Type is Media, this is a media asset ID.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The media object type. Valid values:
	//
	// - OSS: An OSS file.
	//
	// - Media: A media asset ID.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitlesOverwriteParamsFile) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitlesOverwriteParamsFile) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitlesOverwriteParamsFile) GetMedia() *string {
	return s.Media
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitlesOverwriteParamsFile) GetType() *string {
	return s.Type
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitlesOverwriteParamsFile) SetMedia(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitlesOverwriteParamsFile {
	s.Media = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitlesOverwriteParamsFile) SetType(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitlesOverwriteParamsFile {
	s.Type = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigSubtitlesOverwriteParamsFile) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarks struct {
	// Parameters to overwrite. If you specify these, they replace the corresponding parameters in the template.
	OverwriteParams *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarksOverwriteParams `json:"OverwriteParams,omitempty" xml:"OverwriteParams,omitempty" type:"Struct"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9547c6ad97cb4f2aaa29683ebd18d410
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarks) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarks) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarks) GetOverwriteParams() *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarksOverwriteParams {
	return s.OverwriteParams
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarks) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarks) SetOverwriteParams(v *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarksOverwriteParams) *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarks {
	s.OverwriteParams = v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarks) SetTemplateId(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarks {
	s.TemplateId = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarks) Validate() error {
	if s.OverwriteParams != nil {
		if err := s.OverwriteParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarksOverwriteParams struct {
	// Adjusts the font size based on the output video size. Valid values: true or false. Default: false.
	//
	// example:
	//
	// false
	Adaptive *string `json:"Adaptive,omitempty" xml:"Adaptive,omitempty"`
	// The outline color. Default: Black. For more values, see BorderColor.
	//
	// example:
	//
	// #006400
	BorderColor *string `json:"BorderColor,omitempty" xml:"BorderColor,omitempty"`
	// The outline width.
	//
	// - Default: 0
	//
	// - Valid values: (0, 4096]
	//
	// example:
	//
	// 0
	BorderWidth *int32 `json:"BorderWidth,omitempty" xml:"BorderWidth,omitempty"`
	// The watermark text. It does not need to be Base64 encoded. The string must be UTF-8 encoded.
	//
	// example:
	//
	// Test watermark
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The font transparency.
	//
	// - Valid values: (0, 1]
	//
	// - Default: 1.0
	//
	// example:
	//
	// 1.0
	FontAlpha *string `json:"FontAlpha,omitempty" xml:"FontAlpha,omitempty"`
	// The color.
	//
	// example:
	//
	// #006400
	FontColor *string `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	// The font. Default: SimSun.
	//
	// example:
	//
	// SimSun
	FontName *string `json:"FontName,omitempty" xml:"FontName,omitempty"`
	// The font size.
	//
	// - Default value: 16
	//
	// - Valid values: (4, 120)
	//
	// example:
	//
	// 16
	FontSize *int32 `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	// The left margin of the text.
	//
	// - Default: 0
	//
	// - Valid values: [0, 4096]
	//
	// example:
	//
	// 10
	Left *string `json:"Left,omitempty" xml:"Left,omitempty"`
	// The top margin of the text.
	//
	// - Default: 0
	//
	// - Valid values: [0, 4096]
	//
	// example:
	//
	// 10
	Top *string `json:"Top,omitempty" xml:"Top,omitempty"`
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarksOverwriteParams) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarksOverwriteParams) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarksOverwriteParams) GetAdaptive() *string {
	return s.Adaptive
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarksOverwriteParams) GetBorderColor() *string {
	return s.BorderColor
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarksOverwriteParams) GetBorderWidth() *int32 {
	return s.BorderWidth
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarksOverwriteParams) GetContent() *string {
	return s.Content
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarksOverwriteParams) GetFontAlpha() *string {
	return s.FontAlpha
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarksOverwriteParams) GetFontColor() *string {
	return s.FontColor
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarksOverwriteParams) GetFontName() *string {
	return s.FontName
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarksOverwriteParams) GetFontSize() *int32 {
	return s.FontSize
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarksOverwriteParams) GetLeft() *string {
	return s.Left
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarksOverwriteParams) GetTop() *string {
	return s.Top
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarksOverwriteParams) SetAdaptive(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.Adaptive = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarksOverwriteParams) SetBorderColor(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.BorderColor = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarksOverwriteParams) SetBorderWidth(v int32) *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.BorderWidth = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarksOverwriteParams) SetContent(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.Content = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarksOverwriteParams) SetFontAlpha(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.FontAlpha = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarksOverwriteParams) SetFontColor(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.FontColor = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarksOverwriteParams) SetFontName(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.FontName = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarksOverwriteParams) SetFontSize(v int32) *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.FontSize = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarksOverwriteParams) SetLeft(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.Left = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarksOverwriteParams) SetTop(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.Top = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTextWatermarksOverwriteParams) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobRequestOutputGroupProcessConfigTranscode struct {
	// Parameters to overwrite. If you specify these, they replace the corresponding parameters in the template.
	OverwriteParams *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParams `json:"OverwriteParams,omitempty" xml:"OverwriteParams,omitempty" type:"Struct"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9547c6ad97cb4f2aaa29683ebd18d410
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigTranscode) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigTranscode) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscode) GetOverwriteParams() *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParams {
	return s.OverwriteParams
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscode) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscode) SetOverwriteParams(v *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParams) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscode {
	s.OverwriteParams = v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscode) SetTemplateId(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscode {
	s.TemplateId = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscode) Validate() error {
	if s.OverwriteParams != nil {
		if err := s.OverwriteParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParams struct {
	// The audio settings.
	Audio *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudio `json:"Audio,omitempty" xml:"Audio,omitempty" type:"Struct"`
	// The container format settings.
	Container *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsContainer `json:"Container,omitempty" xml:"Container,omitempty" type:"Struct"`
	// The multiplexing settings.
	MuxConfig *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig `json:"MuxConfig,omitempty" xml:"MuxConfig,omitempty" type:"Struct"`
	// The conditional transcoding parameters.
	TransConfig *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig `json:"TransConfig,omitempty" xml:"TransConfig,omitempty" type:"Struct"`
	// The video settings.
	Video *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo `json:"Video,omitempty" xml:"Video,omitempty" type:"Struct"`
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParams) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParams) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParams) GetAudio() *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudio {
	return s.Audio
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParams) GetContainer() *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsContainer {
	return s.Container
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParams) GetMuxConfig() *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig {
	return s.MuxConfig
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParams) GetTransConfig() *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig {
	return s.TransConfig
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParams) GetVideo() *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	return s.Video
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParams) SetAudio(v *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudio) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParams {
	s.Audio = v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParams) SetContainer(v *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsContainer) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParams {
	s.Container = v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParams) SetMuxConfig(v *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParams {
	s.MuxConfig = v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParams) SetTransConfig(v *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParams {
	s.TransConfig = v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParams) SetVideo(v *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParams {
	s.Video = v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParams) Validate() error {
	if s.Audio != nil {
		if err := s.Audio.Validate(); err != nil {
			return err
		}
	}
	if s.Container != nil {
		if err := s.Container.Validate(); err != nil {
			return err
		}
	}
	if s.MuxConfig != nil {
		if err := s.MuxConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TransConfig != nil {
		if err := s.TransConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Video != nil {
		if err := s.Video.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudio struct {
	// The audio bitrate of the output file. Valid values: [8, 1000]. Unit: Kbps. Default value: 128.
	//
	// example:
	//
	// 128
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The number of sound channels. Default value: 2.
	//
	// example:
	//
	// 2
	Channels *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	// The audio codec. Valid values: AAC, MP3, VORBIS, and FLAC. Default value: AAC.
	//
	// example:
	//
	// AAC
	Codec *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	// The audio encoding profile. When Codec is AAC, valid values are aac_low, aac_he, aac_he_v2, aac_ld, and aac_eld.
	//
	// example:
	//
	// aac_low
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// Specifies whether to delete the audio stream.
	//
	// example:
	//
	// false
	Remove *string `json:"Remove,omitempty" xml:"Remove,omitempty"`
	// The sample rate. Default value: 44100. Valid values: 22050, 32000, 44100, 48000, and 96000. Unit: Hz.
	//
	// example:
	//
	// 44100
	Samplerate *string `json:"Samplerate,omitempty" xml:"Samplerate,omitempty"`
	// The volume control.
	Volume *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume `json:"Volume,omitempty" xml:"Volume,omitempty" type:"Struct"`
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudio) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudio) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudio) GetBitrate() *string {
	return s.Bitrate
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudio) GetChannels() *string {
	return s.Channels
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudio) GetCodec() *string {
	return s.Codec
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudio) GetProfile() *string {
	return s.Profile
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudio) GetRemove() *string {
	return s.Remove
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudio) GetSamplerate() *string {
	return s.Samplerate
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudio) GetVolume() *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume {
	return s.Volume
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudio) SetBitrate(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudio {
	s.Bitrate = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudio) SetChannels(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudio {
	s.Channels = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudio) SetCodec(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudio {
	s.Codec = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudio) SetProfile(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudio {
	s.Profile = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudio) SetRemove(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudio {
	s.Remove = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudio) SetSamplerate(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudio {
	s.Samplerate = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudio) SetVolume(v *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudio {
	s.Volume = v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudio) Validate() error {
	if s.Volume != nil {
		if err := s.Volume.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume struct {
	// The target loudness level.
	//
	// example:
	//
	// -6
	IntegratedLoudnessTarget *string `json:"IntegratedLoudnessTarget,omitempty" xml:"IntegratedLoudnessTarget,omitempty"`
	// The loudness range.
	//
	// example:
	//
	// 8
	LoudnessRangeTarget *string `json:"LoudnessRangeTarget,omitempty" xml:"LoudnessRangeTarget,omitempty"`
	// The volume adjustment method.
	//
	// example:
	//
	// auto
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The true peak volume.
	//
	// example:
	//
	// -1
	TruePeak *string `json:"TruePeak,omitempty" xml:"TruePeak,omitempty"`
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) GetIntegratedLoudnessTarget() *string {
	return s.IntegratedLoudnessTarget
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) GetLoudnessRangeTarget() *string {
	return s.LoudnessRangeTarget
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) GetMethod() *string {
	return s.Method
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) GetTruePeak() *string {
	return s.TruePeak
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) SetIntegratedLoudnessTarget(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume {
	s.IntegratedLoudnessTarget = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) SetLoudnessRangeTarget(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume {
	s.LoudnessRangeTarget = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) SetMethod(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume {
	s.Method = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) SetTruePeak(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume {
	s.TruePeak = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsContainer struct {
	// The container format.
	//
	// example:
	//
	// mp4
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsContainer) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsContainer) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsContainer) GetFormat() *string {
	return s.Format
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsContainer) SetFormat(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsContainer {
	s.Format = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsContainer) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig struct {
	// The segment settings.
	Segment *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment `json:"Segment,omitempty" xml:"Segment,omitempty" type:"Struct"`
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig) GetSegment() *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment {
	return s.Segment
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig) SetSegment(v *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig {
	s.Segment = v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig) Validate() error {
	if s.Segment != nil {
		if err := s.Segment.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment struct {
	// The segment duration.
	//
	// example:
	//
	// 10
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The forced segment time points.
	//
	// example:
	//
	// 2,3
	ForceSegTime *string `json:"ForceSegTime,omitempty" xml:"ForceSegTime,omitempty"`
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment) GetDuration() *string {
	return s.Duration
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment) GetForceSegTime() *string {
	return s.ForceSegTime
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment) SetDuration(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment {
	s.Duration = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment) SetForceSegTime(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment {
	s.ForceSegTime = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig struct {
	// The method used to adjust the display aspect ratio. This parameter takes effect only when both Width and Height are specified. You can use it with LongShortMode.
	//
	// Valid values: rescale, crop, pad, and none.
	//
	// Default value: none.
	//
	// example:
	//
	// none
	AdjDarMethod *string `json:"AdjDarMethod,omitempty" xml:"AdjDarMethod,omitempty"`
	// Specifies whether to check the audio bitrate. IsCheckAudioBitrate and IsCheckAudioBitrateFail are mutually exclusive. IsCheckAudioBitrateFail has higher priority.
	//
	// - true: Check the bitrate. If the input audio bitrate is lower than the output setting, transcode at the input bitrate.
	//
	// - false: Do not check the bitrate.
	//
	// Default value:
	//
	// - Empty and the codec differs from the input source: false.
	//
	// - Empty and the codec matches the input source: true.
	//
	// example:
	//
	// true
	IsCheckAudioBitrate *string `json:"IsCheckAudioBitrate,omitempty" xml:"IsCheckAudioBitrate,omitempty"`
	// Specifies whether to check the audio bitrate. IsCheckAudioBitrate and IsCheckAudioBitrateFail are mutually exclusive. IsCheckAudioBitrateFail has higher priority.
	//
	// - true: Check the bitrate. If the input audio bitrate is lower than the output setting, return a failure.
	//
	// - false: Do not check the bitrate.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	IsCheckAudioBitrateFail *string `json:"IsCheckAudioBitrateFail,omitempty" xml:"IsCheckAudioBitrateFail,omitempty"`
	// Specifies whether to check the video resolution. IsCheckReso and IsCheckResoFail are mutually exclusive. IsCheckResoFail has higher priority.
	//
	// - true: Check the resolution. If the input video resolution (width or height) is smaller than the output setting, transcode at the input resolution.
	//
	// - false: Do not check the resolution.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	IsCheckReso *string `json:"IsCheckReso,omitempty" xml:"IsCheckReso,omitempty"`
	// Specifies whether to check the video resolution. IsCheckReso and IsCheckResoFail are mutually exclusive. IsCheckResoFail has higher priority.
	//
	// - true: Check the resolution. If the input video resolution (width or height) is smaller than the output setting, return a failure.
	//
	// - false: Do not check the resolution.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	IsCheckResoFail *string `json:"IsCheckResoFail,omitempty" xml:"IsCheckResoFail,omitempty"`
	// Specifies whether to check the video bitrate. IsCheckVideoBitrate and IsCheckVideoBitrateFail are mutually exclusive. IsCheckVideoBitrateFail has higher priority.
	//
	// - true: Check the bitrate. If the input video bitrate is lower than the output setting, transcode at the input bitrate.
	//
	// - false: Do not check the bitrate.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	IsCheckVideoBitrate *string `json:"IsCheckVideoBitrate,omitempty" xml:"IsCheckVideoBitrate,omitempty"`
	// Specifies whether to check the video bitrate. IsCheckVideoBitrate and IsCheckVideoBitrateFail are mutually exclusive. IsCheckVideoBitrateFail has higher priority.
	//
	// - true: Check the bitrate. If the input video bitrate is lower than the output setting, return a failure.
	//
	// - false: Do not check the bitrate.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	IsCheckVideoBitrateFail *string `json:"IsCheckVideoBitrateFail,omitempty" xml:"IsCheckVideoBitrateFail,omitempty"`
	// The video transcoding mode. Valid values:
	//
	// - onepass: Used for adaptive bitrate (ABR) streaming. Encoding is faster than twopass.
	//
	// - twopass: Used for variable bitrate (VBR) streaming. Encoding is slower than onepass.
	//
	// - CBR: Constant bitrate mode.
	//
	// Default value: onepass.
	//
	// example:
	//
	// onepass
	TransMode *string `json:"TransMode,omitempty" xml:"TransMode,omitempty"`
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) GetAdjDarMethod() *string {
	return s.AdjDarMethod
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) GetIsCheckAudioBitrate() *string {
	return s.IsCheckAudioBitrate
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) GetIsCheckAudioBitrateFail() *string {
	return s.IsCheckAudioBitrateFail
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) GetIsCheckReso() *string {
	return s.IsCheckReso
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) GetIsCheckResoFail() *string {
	return s.IsCheckResoFail
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) GetIsCheckVideoBitrate() *string {
	return s.IsCheckVideoBitrate
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) GetIsCheckVideoBitrateFail() *string {
	return s.IsCheckVideoBitrateFail
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) GetTransMode() *string {
	return s.TransMode
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) SetAdjDarMethod(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig {
	s.AdjDarMethod = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) SetIsCheckAudioBitrate(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig {
	s.IsCheckAudioBitrate = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) SetIsCheckAudioBitrateFail(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig {
	s.IsCheckAudioBitrateFail = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) SetIsCheckReso(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig {
	s.IsCheckReso = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) SetIsCheckResoFail(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig {
	s.IsCheckResoFail = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) SetIsCheckVideoBitrate(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig {
	s.IsCheckVideoBitrate = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) SetIsCheckVideoBitrateFail(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig {
	s.IsCheckVideoBitrateFail = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) SetTransMode(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig {
	s.TransMode = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo struct {
	// The maximum bitrate for adaptive bitrate streaming (ABR). This applies only to narrow-high 1. Valid values: [10, 50000]. Unit: Kbps.
	//
	// example:
	//
	// 6000
	AbrMax *string `json:"AbrMax,omitempty" xml:"AbrMax,omitempty"`
	// The average video bitrate. Valid values: [10, 50000]. Unit: Kbps.
	//
	// example:
	//
	// 3000
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The buffer size. Valid values: [1000, 128000]. Default value: 6000. Unit: Kb.
	//
	// example:
	//
	// 6000
	Bufsize *string `json:"Bufsize,omitempty" xml:"Bufsize,omitempty"`
	// The encoding format.
	//
	// example:
	//
	// H.264
	Codec *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	// The constant rate factor (CRF), which controls the trade-off between quality and bitrate. Valid values: [0, 51]. Default values: 23 for H.264 and 26 for H.265.
	//
	// > If you set Crf, the Bitrate setting is ignored.
	//
	// example:
	//
	// 23
	Crf *string `json:"Crf,omitempty" xml:"Crf,omitempty"`
	// The video cropping method. Two options are available:
	//
	// - Automatically detect and crop black bars. Set this to border.
	//
	// - Custom cropping. Format: width:height:left:top. Example: 1280:800:0:140
	//
	// example:
	//
	// 1280:800:0:140
	Crop *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	// The frame rate. Valid values: (0, 60]. Default value: The frame rate of the input file.
	//
	// > If the frame rate of the input file exceeds 60, the system uses 60.
	//
	// example:
	//
	// 25
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The maximum number of frames between keyframes. Valid values: [1, 1080000]. Default value: 250.
	//
	// example:
	//
	// 250
	Gop *string `json:"Gop,omitempty" xml:"Gop,omitempty"`
	// The height. Valid values: [128, 4096]. Unit: px. Default value: The original video height.
	//
	// example:
	//
	// 1080
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// Specifies whether to enable automatic rotation for portrait or landscape videos (also known as long-side/short-side mode).
	//
	// example:
	//
	// false
	LongShortMode *string `json:"LongShortMode,omitempty" xml:"LongShortMode,omitempty"`
	// The peak video bitrate. Valid values: [10, 50000]. Unit: Kbps.
	//
	// example:
	//
	// 9000
	Maxrate *string `json:"Maxrate,omitempty" xml:"Maxrate,omitempty"`
	// The padding configuration for black bars. Format: width:height:left:top. Example: 1280:800:0:140
	//
	// example:
	//
	// 1280:800:0:140
	Pad *string `json:"Pad,omitempty" xml:"Pad,omitempty"`
	// The video color format. Valid values include yuv420p and yuvj420p.
	//
	// example:
	//
	// yuv420p
	PixFmt *string `json:"PixFmt,omitempty" xml:"PixFmt,omitempty"`
	// The video encoder preset. Only H.264 supports this parameter. Valid values: veryfast, fast, medium, slow, and slower. Default value: medium.
	//
	// example:
	//
	// medium
	Preset *string `json:"Preset,omitempty" xml:"Preset,omitempty"`
	// The encoding profile. Valid values: baseline, main, and high.
	//
	// - baseline: For mobile devices.
	//
	// - main: For standard-resolution devices.
	//
	// - high: For high-resolution devices.
	//
	// Default value: high.
	//
	// example:
	//
	// Main
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// Specifies whether to remove the video.
	//
	// example:
	//
	// false
	Remove *string `json:"Remove,omitempty" xml:"Remove,omitempty"`
	// The scan mode. Valid values: interlaced and progressive.
	//
	// example:
	//
	// progressive
	ScanMode *string `json:"ScanMode,omitempty" xml:"ScanMode,omitempty"`
	// The width. Valid values: [128, 4096]. Unit: px. Default value: The original video width.
	//
	// example:
	//
	// 1920
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetAbrMax() *string {
	return s.AbrMax
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetBitrate() *string {
	return s.Bitrate
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetBufsize() *string {
	return s.Bufsize
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetCodec() *string {
	return s.Codec
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetCrf() *string {
	return s.Crf
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetCrop() *string {
	return s.Crop
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetFps() *string {
	return s.Fps
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetGop() *string {
	return s.Gop
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetHeight() *string {
	return s.Height
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetLongShortMode() *string {
	return s.LongShortMode
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetMaxrate() *string {
	return s.Maxrate
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetPad() *string {
	return s.Pad
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetPixFmt() *string {
	return s.PixFmt
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetPreset() *string {
	return s.Preset
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetProfile() *string {
	return s.Profile
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetRemove() *string {
	return s.Remove
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetScanMode() *string {
	return s.ScanMode
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetWidth() *string {
	return s.Width
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetAbrMax(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.AbrMax = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetBitrate(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Bitrate = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetBufsize(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Bufsize = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetCodec(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Codec = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetCrf(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Crf = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetCrop(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Crop = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetFps(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Fps = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetGop(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Gop = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetHeight(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Height = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetLongShortMode(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.LongShortMode = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetMaxrate(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Maxrate = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetPad(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Pad = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetPixFmt(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.PixFmt = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetPreset(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Preset = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetProfile(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Profile = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetRemove(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Remove = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetScanMode(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.ScanMode = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetWidth(v string) *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Width = &v
	return s
}

func (s *SubmitTranscodeJobRequestOutputGroupProcessConfigTranscodeOverwriteParamsVideo) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobRequestScheduleConfig struct {
	// The MPS queue ID.
	//
	// example:
	//
	// e37ebee5d98b4781897f6086e89f9c56
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The job priority. A larger number indicates a higher priority. Valid values: 1 to 10.
	//
	// example:
	//
	// 5
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s SubmitTranscodeJobRequestScheduleConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobRequestScheduleConfig) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobRequestScheduleConfig) GetPipelineId() *string {
	return s.PipelineId
}

func (s *SubmitTranscodeJobRequestScheduleConfig) GetPriority() *int32 {
	return s.Priority
}

func (s *SubmitTranscodeJobRequestScheduleConfig) SetPipelineId(v string) *SubmitTranscodeJobRequestScheduleConfig {
	s.PipelineId = &v
	return s
}

func (s *SubmitTranscodeJobRequestScheduleConfig) SetPriority(v int32) *SubmitTranscodeJobRequestScheduleConfig {
	s.Priority = &v
	return s
}

func (s *SubmitTranscodeJobRequestScheduleConfig) Validate() error {
	return dara.Validate(s)
}
