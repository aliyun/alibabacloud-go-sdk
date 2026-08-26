// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRtsLiveStreamTranscodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApp(v string) *UpdateRtsLiveStreamTranscodeRequest
	GetApp() *string
	SetAudioBitrate(v int32) *UpdateRtsLiveStreamTranscodeRequest
	GetAudioBitrate() *int32
	SetAudioChannelNum(v int32) *UpdateRtsLiveStreamTranscodeRequest
	GetAudioChannelNum() *int32
	SetAudioCodec(v string) *UpdateRtsLiveStreamTranscodeRequest
	GetAudioCodec() *string
	SetAudioProfile(v string) *UpdateRtsLiveStreamTranscodeRequest
	GetAudioProfile() *string
	SetAudioRate(v int32) *UpdateRtsLiveStreamTranscodeRequest
	GetAudioRate() *int32
	SetDeleteBframes(v bool) *UpdateRtsLiveStreamTranscodeRequest
	GetDeleteBframes() *bool
	SetDomain(v string) *UpdateRtsLiveStreamTranscodeRequest
	GetDomain() *string
	SetFPS(v int32) *UpdateRtsLiveStreamTranscodeRequest
	GetFPS() *int32
	SetGop(v string) *UpdateRtsLiveStreamTranscodeRequest
	GetGop() *string
	SetHeight(v int32) *UpdateRtsLiveStreamTranscodeRequest
	GetHeight() *int32
	SetLazy(v string) *UpdateRtsLiveStreamTranscodeRequest
	GetLazy() *string
	SetOpus(v bool) *UpdateRtsLiveStreamTranscodeRequest
	GetOpus() *bool
	SetOwnerId(v int64) *UpdateRtsLiveStreamTranscodeRequest
	GetOwnerId() *int64
	SetProfile(v int32) *UpdateRtsLiveStreamTranscodeRequest
	GetProfile() *int32
	SetRegionId(v string) *UpdateRtsLiveStreamTranscodeRequest
	GetRegionId() *string
	SetTemplate(v string) *UpdateRtsLiveStreamTranscodeRequest
	GetTemplate() *string
	SetTemplateType(v string) *UpdateRtsLiveStreamTranscodeRequest
	GetTemplateType() *string
	SetVideoBitrate(v int32) *UpdateRtsLiveStreamTranscodeRequest
	GetVideoBitrate() *int32
	SetWidth(v int32) *UpdateRtsLiveStreamTranscodeRequest
	GetWidth() *int32
}

type UpdateRtsLiveStreamTranscodeRequest struct {
	// The AppName of the live stream. This parameter cannot be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// aliyun-test
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// The output audio bitrate. Unit: kbps. Valid values: 1 to **1000**.
	//
	// > Required if you set TemplateType to audio.
	//
	// example:
	//
	// 128
	AudioBitrate *int32 `json:"AudioBitrate,omitempty" xml:"AudioBitrate,omitempty"`
	// The number of audio channels. Valid values:
	//
	// - **1**: mono.
	//
	// - **2**: stereo.
	//
	// example:
	//
	// 2
	AudioChannelNum *int32 `json:"AudioChannelNum,omitempty" xml:"AudioChannelNum,omitempty"`
	// The audio codec. Valid values:
	//
	// - aac
	//
	// - mp3
	//
	// > To use the Opus codec, set the Opus parameter to true.
	//
	// example:
	//
	// aac
	AudioCodec *string `json:"AudioCodec,omitempty" xml:"AudioCodec,omitempty"`
	// The audio codec profile. Valid values:
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
	// aac_low
	AudioProfile *string `json:"AudioProfile,omitempty" xml:"AudioProfile,omitempty"`
	// The audio sample rate. Valid values: 22050, 32000, 44100, 48000, 96000. Recommended: 44100. Unit: Hz.
	//
	// > If AudioProfile is set to aac_ld, the sample rate cannot exceed 44100.
	//
	// example:
	//
	// 44100
	AudioRate *int32 `json:"AudioRate,omitempty" xml:"AudioRate,omitempty"`
	// Controls whether to remove B-frames from the transcoded output video. Valid values:
	//
	// - **true**: The transcoded video has no B-frames.
	//
	// - **false**: The transcoded video contains B-frames. This is the default value.
	//
	// > Required if you set TemplateType to h264, h264-nbhd, or h264-origin.
	//
	// example:
	//
	// false
	DeleteBframes *bool `json:"DeleteBframes,omitempty" xml:"DeleteBframes,omitempty"`
	// The streaming domain. This parameter cannot be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The frame rate of the output video. Unit: frames per second (FPS). Valid values: 1 to **60**.
	//
	// > Required if you set TemplateType to h264, h264-nbhd, or h264-origin.
	//
	// example:
	//
	// 30
	FPS *int32 `json:"FPS,omitempty" xml:"FPS,omitempty"`
	// The Group of Pictures (GOP) size, which specifies the keyframe interval. Unit: seconds. Valid values: **1*	- to **3**.
	//
	// example:
	//
	// 2
	Gop *string `json:"Gop,omitempty" xml:"Gop,omitempty"`
	// Output video height in pixels.
	//
	// Requirements:
	//
	// - Height ≥ 100
	//
	// - max(Height, Width) ≤ 2560
	//
	// - min(Height, Width) ≤ 1440
	//
	// > 	- Required if you set TemplateType to h264, h264-nbhd, or h264-origin.
	//
	// >
	//
	// > 	- For h264-origin templates, the resolution can be up to 4K to retain the information of the source stream.
	//
	// example:
	//
	// 1280
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// Specifies whether to enable on-demand transcoding. Valid values:
	//
	// - **yes**: Transcoding only starts when the first viewer requests this transcoded stream.
	//
	// - **no**: Transcoding starts immediately after the stream is published.
	//
	// example:
	//
	// no
	Lazy *string `json:"Lazy,omitempty" xml:"Lazy,omitempty"`
	// Specifies whether to use the Opus codec for audio transcoding. This is mainly for compatibility with native WebRTC. Valid values:
	//
	// - **true**: Transcodes the audio to the Opus format.
	//
	// - **false**: Does not use the Opus format for transcoding. This is the default value.
	//
	// example:
	//
	// true
	Opus    *bool  `json:"Opus,omitempty" xml:"Opus,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The video codec profile. A larger value indicates better video quality and higher resource consumption for encoding and decoding. Valid values:
	//
	// - **1**: baseline (for mobile devices).
	//
	// - **2**: main (for SD devices).
	//
	// - **3**: high (for HD devices).
	//
	// example:
	//
	// 2
	Profile *int32 `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the custom transcoding template. This parameter cannot be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// ld
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The type of the custom transcoding template. This parameter cannot be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// h264
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// The output video bitrate. Unit: kbps. Valid values: 1 to **6000**.
	//
	// > - Required if you set TemplateType to h264, h264-nbhd, or h264-origin.
	//
	// >
	//
	// > - The system tries to transcode the video at the specified bitrate. However, the actual bitrate may not be the same as the specified value, especially when the specified value is too high or too low.
	//
	// example:
	//
	// 2000
	VideoBitrate *int32 `json:"VideoBitrate,omitempty" xml:"VideoBitrate,omitempty"`
	// Output video width in pixels.
	//
	// Requirements:
	//
	// - Width ≥ 100
	//
	// - max(Height, Width) ≤ 2560
	//
	// - min(Height, Width) ≤ 1440
	//
	// > Required if you set TemplateType to h264, h264-nbhd, or h264-origin.
	//
	// For h264-origin templates, the resolution can be up to 4K to retain the information of the source stream.
	//
	// example:
	//
	// 720
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s UpdateRtsLiveStreamTranscodeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRtsLiveStreamTranscodeRequest) GoString() string {
	return s.String()
}

func (s *UpdateRtsLiveStreamTranscodeRequest) GetApp() *string {
	return s.App
}

func (s *UpdateRtsLiveStreamTranscodeRequest) GetAudioBitrate() *int32 {
	return s.AudioBitrate
}

func (s *UpdateRtsLiveStreamTranscodeRequest) GetAudioChannelNum() *int32 {
	return s.AudioChannelNum
}

func (s *UpdateRtsLiveStreamTranscodeRequest) GetAudioCodec() *string {
	return s.AudioCodec
}

func (s *UpdateRtsLiveStreamTranscodeRequest) GetAudioProfile() *string {
	return s.AudioProfile
}

func (s *UpdateRtsLiveStreamTranscodeRequest) GetAudioRate() *int32 {
	return s.AudioRate
}

func (s *UpdateRtsLiveStreamTranscodeRequest) GetDeleteBframes() *bool {
	return s.DeleteBframes
}

func (s *UpdateRtsLiveStreamTranscodeRequest) GetDomain() *string {
	return s.Domain
}

func (s *UpdateRtsLiveStreamTranscodeRequest) GetFPS() *int32 {
	return s.FPS
}

func (s *UpdateRtsLiveStreamTranscodeRequest) GetGop() *string {
	return s.Gop
}

func (s *UpdateRtsLiveStreamTranscodeRequest) GetHeight() *int32 {
	return s.Height
}

func (s *UpdateRtsLiveStreamTranscodeRequest) GetLazy() *string {
	return s.Lazy
}

func (s *UpdateRtsLiveStreamTranscodeRequest) GetOpus() *bool {
	return s.Opus
}

func (s *UpdateRtsLiveStreamTranscodeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateRtsLiveStreamTranscodeRequest) GetProfile() *int32 {
	return s.Profile
}

func (s *UpdateRtsLiveStreamTranscodeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateRtsLiveStreamTranscodeRequest) GetTemplate() *string {
	return s.Template
}

func (s *UpdateRtsLiveStreamTranscodeRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *UpdateRtsLiveStreamTranscodeRequest) GetVideoBitrate() *int32 {
	return s.VideoBitrate
}

func (s *UpdateRtsLiveStreamTranscodeRequest) GetWidth() *int32 {
	return s.Width
}

func (s *UpdateRtsLiveStreamTranscodeRequest) SetApp(v string) *UpdateRtsLiveStreamTranscodeRequest {
	s.App = &v
	return s
}

func (s *UpdateRtsLiveStreamTranscodeRequest) SetAudioBitrate(v int32) *UpdateRtsLiveStreamTranscodeRequest {
	s.AudioBitrate = &v
	return s
}

func (s *UpdateRtsLiveStreamTranscodeRequest) SetAudioChannelNum(v int32) *UpdateRtsLiveStreamTranscodeRequest {
	s.AudioChannelNum = &v
	return s
}

func (s *UpdateRtsLiveStreamTranscodeRequest) SetAudioCodec(v string) *UpdateRtsLiveStreamTranscodeRequest {
	s.AudioCodec = &v
	return s
}

func (s *UpdateRtsLiveStreamTranscodeRequest) SetAudioProfile(v string) *UpdateRtsLiveStreamTranscodeRequest {
	s.AudioProfile = &v
	return s
}

func (s *UpdateRtsLiveStreamTranscodeRequest) SetAudioRate(v int32) *UpdateRtsLiveStreamTranscodeRequest {
	s.AudioRate = &v
	return s
}

func (s *UpdateRtsLiveStreamTranscodeRequest) SetDeleteBframes(v bool) *UpdateRtsLiveStreamTranscodeRequest {
	s.DeleteBframes = &v
	return s
}

func (s *UpdateRtsLiveStreamTranscodeRequest) SetDomain(v string) *UpdateRtsLiveStreamTranscodeRequest {
	s.Domain = &v
	return s
}

func (s *UpdateRtsLiveStreamTranscodeRequest) SetFPS(v int32) *UpdateRtsLiveStreamTranscodeRequest {
	s.FPS = &v
	return s
}

func (s *UpdateRtsLiveStreamTranscodeRequest) SetGop(v string) *UpdateRtsLiveStreamTranscodeRequest {
	s.Gop = &v
	return s
}

func (s *UpdateRtsLiveStreamTranscodeRequest) SetHeight(v int32) *UpdateRtsLiveStreamTranscodeRequest {
	s.Height = &v
	return s
}

func (s *UpdateRtsLiveStreamTranscodeRequest) SetLazy(v string) *UpdateRtsLiveStreamTranscodeRequest {
	s.Lazy = &v
	return s
}

func (s *UpdateRtsLiveStreamTranscodeRequest) SetOpus(v bool) *UpdateRtsLiveStreamTranscodeRequest {
	s.Opus = &v
	return s
}

func (s *UpdateRtsLiveStreamTranscodeRequest) SetOwnerId(v int64) *UpdateRtsLiveStreamTranscodeRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateRtsLiveStreamTranscodeRequest) SetProfile(v int32) *UpdateRtsLiveStreamTranscodeRequest {
	s.Profile = &v
	return s
}

func (s *UpdateRtsLiveStreamTranscodeRequest) SetRegionId(v string) *UpdateRtsLiveStreamTranscodeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateRtsLiveStreamTranscodeRequest) SetTemplate(v string) *UpdateRtsLiveStreamTranscodeRequest {
	s.Template = &v
	return s
}

func (s *UpdateRtsLiveStreamTranscodeRequest) SetTemplateType(v string) *UpdateRtsLiveStreamTranscodeRequest {
	s.TemplateType = &v
	return s
}

func (s *UpdateRtsLiveStreamTranscodeRequest) SetVideoBitrate(v int32) *UpdateRtsLiveStreamTranscodeRequest {
	s.VideoBitrate = &v
	return s
}

func (s *UpdateRtsLiveStreamTranscodeRequest) SetWidth(v int32) *UpdateRtsLiveStreamTranscodeRequest {
	s.Width = &v
	return s
}

func (s *UpdateRtsLiveStreamTranscodeRequest) Validate() error {
	return dara.Validate(s)
}
