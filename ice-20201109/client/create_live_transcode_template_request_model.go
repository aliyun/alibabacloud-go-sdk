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
	// The name of the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// my template
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The configuration of the template.
	TemplateConfig *CreateLiveTranscodeTemplateRequestTemplateConfig `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty" type:"Struct"`
	// The type of the template. Valid values:
	//
	// 	- normal
	//
	// 	- narrow-band
	//
	// 	- audio-only
	//
	// 	- origin
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
	// The audio codec. Valid values:
	//
	// 	- AAC
	//
	// 	- MP3
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
	// aaclow
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The audio sampling rate. Valid values: 22050 to 96000.
	//
	// Note: If you set AudioProfile to aac_ld, the audio sampling rate cannot exceed 44,100.
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
	// 25
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The group of pictures (GOP) of the output video. Unit: frame. Valid values: 1 to 3000.
	//
	// example:
	//
	// 1000
	Gop *string `json:"Gop,omitempty" xml:"Gop,omitempty"`
	// The height of the output video. Valid values: Height ≥ 128 max (Height,Width) ≤ 2560 min (Height,Width) ≤ 1440
	//
	// Note: The resolution of the output video that is transcoded by using the H.265 Narrowband HD transcoding template cannot exceed 1280 × 720 pixels.
	//
	// example:
	//
	// 720
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The encoding profile. The profile determines how a video is encoded. In most cases, a greater value indicates better image quality and higher resource consumption. Valid values: 1: baseline. This value is suitable for mobile devices. 2: main. This value is suitable for standard-definition devices. 3: high. This value is suitable for high-definition devices.
	//
	// example:
	//
	// 2
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The width of the output video. Valid values: Width ≥ 128 max (Height,Width) ≤ 2560 min (Height,Width) ≤ 1440
	//
	// Note: The resolution of the output video that is transcoded by using the H.265 Narrowband HD transcoding template cannot exceed 1280 × 720 pixels.
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
