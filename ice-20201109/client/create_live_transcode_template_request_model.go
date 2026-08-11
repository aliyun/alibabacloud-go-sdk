// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLiveTranscodeTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateLiveTranscodeTemplateRequest
	GetName() *string
	SetTemplateConfig(v *CreateLiveTranscodeTemplateRequestTemplateConfig) *CreateLiveTranscodeTemplateRequest
	GetTemplateConfig() *CreateLiveTranscodeTemplateRequestTemplateConfig
	SetType(v string) *CreateLiveTranscodeTemplateRequest
	GetType() *string
}

type CreateLiveTranscodeTemplateRequest struct {
	// The template name.
	//
	// This parameter is required.
	//
	// example:
	//
	// my template
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The template configuration.
	//
	// > The pass parameter requirements vary based on the templatetype (Type). When Type is set to normal, at least one of the width and height parameters must be specified, and the frame rate and bitrate parameters are required. For other template types, specify the parameters based on your requirements.
	TemplateConfig *CreateLiveTranscodeTemplateRequestTemplateConfig `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty" type:"Struct"`
	// The template type. Valid values:
	//
	// - normal: standard.
	//
	// - narrow-band: narrowband HD.
	//
	// - audio-only: audio only.
	//
	// - origin: original quality.
	//
	// This parameter is required.
	//
	// example:
	//
	// normal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateLiveTranscodeTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveTranscodeTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateLiveTranscodeTemplateRequest) GetName() *string {
	return s.Name
}

func (s *CreateLiveTranscodeTemplateRequest) GetTemplateConfig() *CreateLiveTranscodeTemplateRequestTemplateConfig {
	return s.TemplateConfig
}

func (s *CreateLiveTranscodeTemplateRequest) GetType() *string {
	return s.Type
}

func (s *CreateLiveTranscodeTemplateRequest) SetName(v string) *CreateLiveTranscodeTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateLiveTranscodeTemplateRequest) SetTemplateConfig(v *CreateLiveTranscodeTemplateRequestTemplateConfig) *CreateLiveTranscodeTemplateRequest {
	s.TemplateConfig = v
	return s
}

func (s *CreateLiveTranscodeTemplateRequest) SetType(v string) *CreateLiveTranscodeTemplateRequest {
	s.Type = &v
	return s
}

func (s *CreateLiveTranscodeTemplateRequest) Validate() error {
	if s.TemplateConfig != nil {
		if err := s.TemplateConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateLiveTranscodeTemplateRequestTemplateConfig struct {
	// The audio parameters.
	AudioParams *CreateLiveTranscodeTemplateRequestTemplateConfigAudioParams `json:"AudioParams,omitempty" xml:"AudioParams,omitempty" type:"Struct"`
	// The video parameters.
	VideoParams *CreateLiveTranscodeTemplateRequestTemplateConfigVideoParams `json:"VideoParams,omitempty" xml:"VideoParams,omitempty" type:"Struct"`
}

func (s CreateLiveTranscodeTemplateRequestTemplateConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveTranscodeTemplateRequestTemplateConfig) GoString() string {
	return s.String()
}

func (s *CreateLiveTranscodeTemplateRequestTemplateConfig) GetAudioParams() *CreateLiveTranscodeTemplateRequestTemplateConfigAudioParams {
	return s.AudioParams
}

func (s *CreateLiveTranscodeTemplateRequestTemplateConfig) GetVideoParams() *CreateLiveTranscodeTemplateRequestTemplateConfigVideoParams {
	return s.VideoParams
}

func (s *CreateLiveTranscodeTemplateRequestTemplateConfig) SetAudioParams(v *CreateLiveTranscodeTemplateRequestTemplateConfigAudioParams) *CreateLiveTranscodeTemplateRequestTemplateConfig {
	s.AudioParams = v
	return s
}

func (s *CreateLiveTranscodeTemplateRequestTemplateConfig) SetVideoParams(v *CreateLiveTranscodeTemplateRequestTemplateConfigVideoParams) *CreateLiveTranscodeTemplateRequestTemplateConfig {
	s.VideoParams = v
	return s
}

func (s *CreateLiveTranscodeTemplateRequestTemplateConfig) Validate() error {
	if s.AudioParams != nil {
		if err := s.AudioParams.Validate(); err != nil {
			return err
		}
	}
	if s.VideoParams != nil {
		if err := s.VideoParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateLiveTranscodeTemplateRequestTemplateConfigAudioParams struct {
	// The bitrate of the transcoded audio. Unit: kbps. Valid values: 1 to 1000.
	//
	// example:
	//
	// 100
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The number of audio channels. Valid values:
	//
	// - 1: mono.
	//
	// - 2: stereo.
	//
	// example:
	//
	// 2
	Channels *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	// The audio encoding format. Valid values:
	//
	// - AAC
	//
	// - MP3
	//
	// example:
	//
	// AAC
	Codec *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	// The audio encoding preset. When Codec is set to AAC, valid values:
	//
	// - aac_low
	//
	// - aac_he
	//
	// - aac_he_v2
	//
	// - aac_ld
	//
	// example:
	//
	// aaclow
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The audio sample rate. Valid values: 22050 to 96000.
	//
	// 	Notice: If AudioProfile is set to aac_ld, the sample rate must not exceed 44100.
	//
	// example:
	//
	// 44100
	Samplerate *string `json:"Samplerate,omitempty" xml:"Samplerate,omitempty"`
}

func (s CreateLiveTranscodeTemplateRequestTemplateConfigAudioParams) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveTranscodeTemplateRequestTemplateConfigAudioParams) GoString() string {
	return s.String()
}

func (s *CreateLiveTranscodeTemplateRequestTemplateConfigAudioParams) GetBitrate() *string {
	return s.Bitrate
}

func (s *CreateLiveTranscodeTemplateRequestTemplateConfigAudioParams) GetChannels() *string {
	return s.Channels
}

func (s *CreateLiveTranscodeTemplateRequestTemplateConfigAudioParams) GetCodec() *string {
	return s.Codec
}

func (s *CreateLiveTranscodeTemplateRequestTemplateConfigAudioParams) GetProfile() *string {
	return s.Profile
}

func (s *CreateLiveTranscodeTemplateRequestTemplateConfigAudioParams) GetSamplerate() *string {
	return s.Samplerate
}

func (s *CreateLiveTranscodeTemplateRequestTemplateConfigAudioParams) SetBitrate(v string) *CreateLiveTranscodeTemplateRequestTemplateConfigAudioParams {
	s.Bitrate = &v
	return s
}

func (s *CreateLiveTranscodeTemplateRequestTemplateConfigAudioParams) SetChannels(v string) *CreateLiveTranscodeTemplateRequestTemplateConfigAudioParams {
	s.Channels = &v
	return s
}

func (s *CreateLiveTranscodeTemplateRequestTemplateConfigAudioParams) SetCodec(v string) *CreateLiveTranscodeTemplateRequestTemplateConfigAudioParams {
	s.Codec = &v
	return s
}

func (s *CreateLiveTranscodeTemplateRequestTemplateConfigAudioParams) SetProfile(v string) *CreateLiveTranscodeTemplateRequestTemplateConfigAudioParams {
	s.Profile = &v
	return s
}

func (s *CreateLiveTranscodeTemplateRequestTemplateConfigAudioParams) SetSamplerate(v string) *CreateLiveTranscodeTemplateRequestTemplateConfigAudioParams {
	s.Samplerate = &v
	return s
}

func (s *CreateLiveTranscodeTemplateRequestTemplateConfigAudioParams) Validate() error {
	return dara.Validate(s)
}

type CreateLiveTranscodeTemplateRequestTemplateConfigVideoParams struct {
	// The bitrate of the transcoded video. Unit: kbps. Valid values: 1 to 6000.
	//
	// example:
	//
	// 2500
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The encoding type. Valid values:
	//
	// - H.264
	//
	// - H.265
	//
	// example:
	//
	// H.264
	Codec *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	// The frame rate of the transcoded video. Unit: FPS. Valid values: 1 to 60.
	//
	// example:
	//
	// 25
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The video GOP (Group of Pictures). Unit: frames. Valid values: 1 to 3000.
	//
	// example:
	//
	// 1000
	Gop *string `json:"Gop,omitempty" xml:"Gop,omitempty"`
	// The height of the transcoded video. Valid values:
	//
	// - Height ≥ 128
	//
	// - max(Height, Width) ≤ 2560
	//
	// - min(Height, Width) ≤ 1440
	//
	// 	Notice: For H.265 narrowband HD templates, the resolution must not exceed 1280 × 720.
	//
	// example:
	//
	// 720
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The encoding profile. A set of specific encoding features supported by the video. A higher value generally produces better image quality but consumes more encoding and decoding resources. Valid values:
	//
	// - 1: baseline (suitable for mobile devices).
	//
	// - 2: main (suitable for standard resolution devices).
	//
	// - 3: high (suitable for high resolution devices).
	//
	// example:
	//
	// 2
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The width of the transcoded video. Valid values:
	//
	// - Width ≥ 128
	//
	// - max(Height, Width) ≤ 2560
	//
	// - min(Height, Width) ≤ 1440
	//
	// 	Notice: For H.265 narrowband HD templates, the resolution must not exceed 1280 × 720.
	//
	// example:
	//
	// 1280
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s CreateLiveTranscodeTemplateRequestTemplateConfigVideoParams) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveTranscodeTemplateRequestTemplateConfigVideoParams) GoString() string {
	return s.String()
}

func (s *CreateLiveTranscodeTemplateRequestTemplateConfigVideoParams) GetBitrate() *string {
	return s.Bitrate
}

func (s *CreateLiveTranscodeTemplateRequestTemplateConfigVideoParams) GetCodec() *string {
	return s.Codec
}

func (s *CreateLiveTranscodeTemplateRequestTemplateConfigVideoParams) GetFps() *string {
	return s.Fps
}

func (s *CreateLiveTranscodeTemplateRequestTemplateConfigVideoParams) GetGop() *string {
	return s.Gop
}

func (s *CreateLiveTranscodeTemplateRequestTemplateConfigVideoParams) GetHeight() *string {
	return s.Height
}

func (s *CreateLiveTranscodeTemplateRequestTemplateConfigVideoParams) GetProfile() *string {
	return s.Profile
}

func (s *CreateLiveTranscodeTemplateRequestTemplateConfigVideoParams) GetWidth() *string {
	return s.Width
}

func (s *CreateLiveTranscodeTemplateRequestTemplateConfigVideoParams) SetBitrate(v string) *CreateLiveTranscodeTemplateRequestTemplateConfigVideoParams {
	s.Bitrate = &v
	return s
}

func (s *CreateLiveTranscodeTemplateRequestTemplateConfigVideoParams) SetCodec(v string) *CreateLiveTranscodeTemplateRequestTemplateConfigVideoParams {
	s.Codec = &v
	return s
}

func (s *CreateLiveTranscodeTemplateRequestTemplateConfigVideoParams) SetFps(v string) *CreateLiveTranscodeTemplateRequestTemplateConfigVideoParams {
	s.Fps = &v
	return s
}

func (s *CreateLiveTranscodeTemplateRequestTemplateConfigVideoParams) SetGop(v string) *CreateLiveTranscodeTemplateRequestTemplateConfigVideoParams {
	s.Gop = &v
	return s
}

func (s *CreateLiveTranscodeTemplateRequestTemplateConfigVideoParams) SetHeight(v string) *CreateLiveTranscodeTemplateRequestTemplateConfigVideoParams {
	s.Height = &v
	return s
}

func (s *CreateLiveTranscodeTemplateRequestTemplateConfigVideoParams) SetProfile(v string) *CreateLiveTranscodeTemplateRequestTemplateConfigVideoParams {
	s.Profile = &v
	return s
}

func (s *CreateLiveTranscodeTemplateRequestTemplateConfigVideoParams) SetWidth(v string) *CreateLiveTranscodeTemplateRequestTemplateConfigVideoParams {
	s.Width = &v
	return s
}

func (s *CreateLiveTranscodeTemplateRequestTemplateConfigVideoParams) Validate() error {
	return dara.Validate(s)
}
