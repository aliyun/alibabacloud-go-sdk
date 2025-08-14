// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveTranscodeTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *UpdateLiveTranscodeTemplateRequest
	GetName() *string
	SetTemplateConfig(v *UpdateLiveTranscodeTemplateRequestTemplateConfig) *UpdateLiveTranscodeTemplateRequest
	GetTemplateConfig() *UpdateLiveTranscodeTemplateRequestTemplateConfig
	SetTemplateId(v string) *UpdateLiveTranscodeTemplateRequest
	GetTemplateId() *string
}

type UpdateLiveTranscodeTemplateRequest struct {
	// The template name.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The configuration of the template.
	TemplateConfig *UpdateLiveTranscodeTemplateRequestTemplateConfig `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty" type:"Struct"`
	// The template ID. To obtain the template ID, log on to the [Intelligent Media Services (IMS) console](https://ims.console.aliyun.com/summary), choose Real-time Media Processing > Template Management, and then click the Transcoding tab. Alternatively, find the ID from the response parameters of the [CreateLiveTranscodeTemplate](https://help.aliyun.com/document_detail/449217.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****96e8864746a0b6f3****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s UpdateLiveTranscodeTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveTranscodeTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveTranscodeTemplateRequest) GetName() *string {
	return s.Name
}

func (s *UpdateLiveTranscodeTemplateRequest) GetTemplateConfig() *UpdateLiveTranscodeTemplateRequestTemplateConfig {
	return s.TemplateConfig
}

func (s *UpdateLiveTranscodeTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateLiveTranscodeTemplateRequest) SetName(v string) *UpdateLiveTranscodeTemplateRequest {
	s.Name = &v
	return s
}

func (s *UpdateLiveTranscodeTemplateRequest) SetTemplateConfig(v *UpdateLiveTranscodeTemplateRequestTemplateConfig) *UpdateLiveTranscodeTemplateRequest {
	s.TemplateConfig = v
	return s
}

func (s *UpdateLiveTranscodeTemplateRequest) SetTemplateId(v string) *UpdateLiveTranscodeTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateLiveTranscodeTemplateRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateLiveTranscodeTemplateRequestTemplateConfig struct {
	// The audio parameters.
	AudioParams *UpdateLiveTranscodeTemplateRequestTemplateConfigAudioParams `json:"AudioParams,omitempty" xml:"AudioParams,omitempty" type:"Struct"`
	// The video parameters.
	VideoParams *UpdateLiveTranscodeTemplateRequestTemplateConfigVideoParams `json:"VideoParams,omitempty" xml:"VideoParams,omitempty" type:"Struct"`
}

func (s UpdateLiveTranscodeTemplateRequestTemplateConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveTranscodeTemplateRequestTemplateConfig) GoString() string {
	return s.String()
}

func (s *UpdateLiveTranscodeTemplateRequestTemplateConfig) GetAudioParams() *UpdateLiveTranscodeTemplateRequestTemplateConfigAudioParams {
	return s.AudioParams
}

func (s *UpdateLiveTranscodeTemplateRequestTemplateConfig) GetVideoParams() *UpdateLiveTranscodeTemplateRequestTemplateConfigVideoParams {
	return s.VideoParams
}

func (s *UpdateLiveTranscodeTemplateRequestTemplateConfig) SetAudioParams(v *UpdateLiveTranscodeTemplateRequestTemplateConfigAudioParams) *UpdateLiveTranscodeTemplateRequestTemplateConfig {
	s.AudioParams = v
	return s
}

func (s *UpdateLiveTranscodeTemplateRequestTemplateConfig) SetVideoParams(v *UpdateLiveTranscodeTemplateRequestTemplateConfigVideoParams) *UpdateLiveTranscodeTemplateRequestTemplateConfig {
	s.VideoParams = v
	return s
}

func (s *UpdateLiveTranscodeTemplateRequestTemplateConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateLiveTranscodeTemplateRequestTemplateConfigAudioParams struct {
	// The bitrate of the output audio. Unit: Kbit/s. Valid values: 1 to 1000.
	//
	// example:
	//
	// 100
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The number of sound channels. Valid values: 1: mono 2: binaural
	//
	// example:
	//
	// 2
	Channels *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	// The audio codec. Valid values: AAC MP3
	//
	// example:
	//
	// AAC
	Codec *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	// The audio codec profile. Valid values when the Codec parameter is set to AAC:
	//
	// 	- aac_low
	//
	// 	- aac_he
	//
	// 	- aac_he_v2
	//
	// 	- aac_ld
	//
	// example:
	//
	// aac_low
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The audio sampling rate. Valid values: 22050 to 96000.
	//
	// Note If you set AudioProfile to aac_ld, the audio sampling rate cannot exceed 44100.
	//
	// example:
	//
	// 44100
	Samplerate *string `json:"Samplerate,omitempty" xml:"Samplerate,omitempty"`
}

func (s UpdateLiveTranscodeTemplateRequestTemplateConfigAudioParams) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveTranscodeTemplateRequestTemplateConfigAudioParams) GoString() string {
	return s.String()
}

func (s *UpdateLiveTranscodeTemplateRequestTemplateConfigAudioParams) GetBitrate() *string {
	return s.Bitrate
}

func (s *UpdateLiveTranscodeTemplateRequestTemplateConfigAudioParams) GetChannels() *string {
	return s.Channels
}

func (s *UpdateLiveTranscodeTemplateRequestTemplateConfigAudioParams) GetCodec() *string {
	return s.Codec
}

func (s *UpdateLiveTranscodeTemplateRequestTemplateConfigAudioParams) GetProfile() *string {
	return s.Profile
}

func (s *UpdateLiveTranscodeTemplateRequestTemplateConfigAudioParams) GetSamplerate() *string {
	return s.Samplerate
}

func (s *UpdateLiveTranscodeTemplateRequestTemplateConfigAudioParams) SetBitrate(v string) *UpdateLiveTranscodeTemplateRequestTemplateConfigAudioParams {
	s.Bitrate = &v
	return s
}

func (s *UpdateLiveTranscodeTemplateRequestTemplateConfigAudioParams) SetChannels(v string) *UpdateLiveTranscodeTemplateRequestTemplateConfigAudioParams {
	s.Channels = &v
	return s
}

func (s *UpdateLiveTranscodeTemplateRequestTemplateConfigAudioParams) SetCodec(v string) *UpdateLiveTranscodeTemplateRequestTemplateConfigAudioParams {
	s.Codec = &v
	return s
}

func (s *UpdateLiveTranscodeTemplateRequestTemplateConfigAudioParams) SetProfile(v string) *UpdateLiveTranscodeTemplateRequestTemplateConfigAudioParams {
	s.Profile = &v
	return s
}

func (s *UpdateLiveTranscodeTemplateRequestTemplateConfigAudioParams) SetSamplerate(v string) *UpdateLiveTranscodeTemplateRequestTemplateConfigAudioParams {
	s.Samplerate = &v
	return s
}

func (s *UpdateLiveTranscodeTemplateRequestTemplateConfigAudioParams) Validate() error {
	return dara.Validate(s)
}

type UpdateLiveTranscodeTemplateRequestTemplateConfigVideoParams struct {
	// The bitrate of the output video. Unit: Kbit/s. Valid values: 1 to 6000.
	//
	// example:
	//
	// 2500
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The encoding type. Valid values:
	//
	// 	- H.264
	//
	// 	- H.265
	//
	// example:
	//
	// H.264
	Codec *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	// The frame rate of the output video. Unit: frames per second (FPS). Valid values: 1 to 60.
	//
	// example:
	//
	// 30
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The group of pictures (GOP) of the output video. Unit: frame. Valid values: 1 to 3000.
	//
	// example:
	//
	// 1000
	Gop *string `json:"Gop,omitempty" xml:"Gop,omitempty"`
	// The height of the output video. Valid values:
	//
	// 	- Height ≥ 128
	//
	// 	- max (Height,Width) ≤ 2560
	//
	// 	- min（Height,Width）≤ 1440
	//
	// >  The resolution of a video transcoded by using the H.265 Narrowband HD template cannot exceed 1,280 × 720 pixels.
	//
	// example:
	//
	// 720
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The video encoding profile. The profile determines how a video is encoded. In most cases, a greater value indicates better image quality and higher resource consumption. Valid values:
	//
	// 	- 1: baseline. This value is suitable for mobile devices.
	//
	// 	- 2: main. This value is suitable for standard-definition devices.
	//
	// 	- 3: high. This value is suitable for high-definition devices.
	//
	// example:
	//
	// 2
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The width of the output video. Valid values:
	//
	// 	- Width ≥ 128
	//
	// 	- max (Height,Width) ≤ 2560
	//
	// 	- min（Height,Width）≤ 1440
	//
	// >  The resolution of a video transcoded by using the H.265 Narrowband HD template cannot exceed 1,280 × 720 pixels.
	//
	// example:
	//
	// 1280
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s UpdateLiveTranscodeTemplateRequestTemplateConfigVideoParams) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveTranscodeTemplateRequestTemplateConfigVideoParams) GoString() string {
	return s.String()
}

func (s *UpdateLiveTranscodeTemplateRequestTemplateConfigVideoParams) GetBitrate() *string {
	return s.Bitrate
}

func (s *UpdateLiveTranscodeTemplateRequestTemplateConfigVideoParams) GetCodec() *string {
	return s.Codec
}

func (s *UpdateLiveTranscodeTemplateRequestTemplateConfigVideoParams) GetFps() *string {
	return s.Fps
}

func (s *UpdateLiveTranscodeTemplateRequestTemplateConfigVideoParams) GetGop() *string {
	return s.Gop
}

func (s *UpdateLiveTranscodeTemplateRequestTemplateConfigVideoParams) GetHeight() *string {
	return s.Height
}

func (s *UpdateLiveTranscodeTemplateRequestTemplateConfigVideoParams) GetProfile() *string {
	return s.Profile
}

func (s *UpdateLiveTranscodeTemplateRequestTemplateConfigVideoParams) GetWidth() *string {
	return s.Width
}

func (s *UpdateLiveTranscodeTemplateRequestTemplateConfigVideoParams) SetBitrate(v string) *UpdateLiveTranscodeTemplateRequestTemplateConfigVideoParams {
	s.Bitrate = &v
	return s
}

func (s *UpdateLiveTranscodeTemplateRequestTemplateConfigVideoParams) SetCodec(v string) *UpdateLiveTranscodeTemplateRequestTemplateConfigVideoParams {
	s.Codec = &v
	return s
}

func (s *UpdateLiveTranscodeTemplateRequestTemplateConfigVideoParams) SetFps(v string) *UpdateLiveTranscodeTemplateRequestTemplateConfigVideoParams {
	s.Fps = &v
	return s
}

func (s *UpdateLiveTranscodeTemplateRequestTemplateConfigVideoParams) SetGop(v string) *UpdateLiveTranscodeTemplateRequestTemplateConfigVideoParams {
	s.Gop = &v
	return s
}

func (s *UpdateLiveTranscodeTemplateRequestTemplateConfigVideoParams) SetHeight(v string) *UpdateLiveTranscodeTemplateRequestTemplateConfigVideoParams {
	s.Height = &v
	return s
}

func (s *UpdateLiveTranscodeTemplateRequestTemplateConfigVideoParams) SetProfile(v string) *UpdateLiveTranscodeTemplateRequestTemplateConfigVideoParams {
	s.Profile = &v
	return s
}

func (s *UpdateLiveTranscodeTemplateRequestTemplateConfigVideoParams) SetWidth(v string) *UpdateLiveTranscodeTemplateRequestTemplateConfigVideoParams {
	s.Width = &v
	return s
}

func (s *UpdateLiveTranscodeTemplateRequestTemplateConfigVideoParams) Validate() error {
	return dara.Validate(s)
}
