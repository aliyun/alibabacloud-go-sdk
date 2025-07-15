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
	// The name of the application to which the live stream belongs, which cannot be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// aliyun-test
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// The bitrate of the output audio. Unit: Kbit/s. Valid values: 1 to 1000.
	//
	// >  This parameter is required if you set the TemplateType parameter to audio.
	//
	// example:
	//
	// 128
	AudioBitrate *int32 `json:"AudioBitrate,omitempty" xml:"AudioBitrate,omitempty"`
	// The number of sound channels. Valid values:
	//
	// 	- **1**: mono
	//
	// 	- **2**: stereo
	//
	// example:
	//
	// 2
	AudioChannelNum *int32 `json:"AudioChannelNum,omitempty" xml:"AudioChannelNum,omitempty"`
	// The audio encoder. Valid values:
	//
	// 	- aac
	//
	// 	- mp3
	//
	// >  To use the Opus encoder, you need only to set the Opus parameter to true.
	//
	// example:
	//
	// aac
	AudioCodec *string `json:"AudioCodec,omitempty" xml:"AudioCodec,omitempty"`
	// The audio codec profile. Valid values:
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
	AudioProfile *string `json:"AudioProfile,omitempty" xml:"AudioProfile,omitempty"`
	// The audio sampling rate. Valid values: 22050, 32000, 44100, 48000, and 96000. 44100 is commonly used. Unit: Hz.
	//
	// >  If the value of the AudioProfile parameter is aac_ld, the audio sampling rate cannot exceed 44,100.
	//
	// example:
	//
	// 44100
	AudioRate *int32 `json:"AudioRate,omitempty" xml:"AudioRate,omitempty"`
	// Specifies whether to remove B frames during transcoding. Valid values:
	//
	// 	- true: Remove B frames.
	//
	// 	- false (default): Retain B frames.
	//
	// >
	//
	// 	- This parameter is required when the TemplateType parameter is set to h264, h264-nbhd, or h264-origin.
	//
	// 	- If this parameter is not specified, the default value false is used.
	//
	// example:
	//
	// false
	DeleteBframes *bool `json:"DeleteBframes,omitempty" xml:"DeleteBframes,omitempty"`
	// The main streaming domain, which cannot be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The frame rate of the output video. Unit: frames per second (FPS). Valid values: 1 to 60.
	//
	// >  This parameter is required when the TemplateType parameter is set to h264, h264-nbhd, or h264-origin.
	//
	// example:
	//
	// 30
	FPS *int32 `json:"FPS,omitempty" xml:"FPS,omitempty"`
	// The group of pictures (GOP) of the output video. This parameter is used to specify the keyframe interval. Unit: seconds. Valid values: 1 to 3.
	//
	// example:
	//
	// 2
	Gop *string `json:"Gop,omitempty" xml:"Gop,omitempty"`
	// The height of the output video. Unit: pixel.
	//
	// The value must comply with the following rules:
	//
	// 	- Height ≥ 100: The height of the video is at least 100 pixels.
	//
	// 	- max(Height,Width) ≤ 2560: The larger of the width and height of the video cannot exceed 2,560 pixels.
	//
	// 	- min(Height,Width) ≤ 1440: The smaller of the width and height of the video cannot exceed 1,440 pixels. For example, a resolution of 1560 × 1560 pixels is invalid.
	//
	// >
	//
	// 	- This parameter is required when the TemplateType parameter is set to h264, h264-nbhd, or h264-origin.
	//
	// 	- If the TemplateType parameter is set to h264-origin, the highest resolution supported is 4K.
	//
	// example:
	//
	// 1280
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// Specifies whether to enable triggered transcoding. Valid values:
	//
	// 	- **yes**: enables triggered transcoding.
	//
	// 	- **no**: disables triggered transcoding.
	//
	// example:
	//
	// no
	Lazy *string `json:"Lazy,omitempty" xml:"Lazy,omitempty"`
	// Specifies whether to transcode streams to the Opus format to be compatible with native WebRTC. Valid values:
	//
	// 	- true: Transcode streams to the Opus format.
	//
	// 	- false: Do not transcode streams to the Opus format.
	//
	// >  If this parameter is not specified, the default value false is used.
	//
	// example:
	//
	// true
	Opus    *bool  `json:"Opus,omitempty" xml:"Opus,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The video encoding profile. The profile defines a set of parameters that are used to encode a video. In most cases, a greater value indicates better image quality and higher resource consumption. Valid values:
	//
	// 	- **1**: baseline. This value is suitable for mobile devices.
	//
	// 	- **2**: main. This value is suitable for standard-definition devices.
	//
	// 	- **3**: high. This value is suitable for high-definition devices.
	//
	// example:
	//
	// 2
	Profile  *int32  `json:"Profile,omitempty" xml:"Profile,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the custom transcoding template, which cannot be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// ld
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The type of the custom transcoding template, which cannot be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// h264
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// The bitrate of the output video. Unit: Kbit/s. Valid values: 1 to 6000.
	//
	// >
	//
	// 	- This parameter is required when the TemplateType parameter is set to h264, h264-nbhd, or h264-origin.
	//
	// 	- The bitrate of the output video is as close to the value that you specify as possible, but not the same as the value, especially when the value is excessively large or small.
	//
	// example:
	//
	// 2000
	VideoBitrate *int32 `json:"VideoBitrate,omitempty" xml:"VideoBitrate,omitempty"`
	// The width of the output video. Unit: pixel.
	//
	// The value must comply with the following rules:
	//
	// 	- Width ≥ 100: The width of the video is at least 100 pixels.
	//
	// 	- max(Height,Width) ≤ 2560: The larger of the width and height of the video cannot exceed 2,560 pixels.
	//
	// 	- min(Height,Width) ≤ 1440: The smaller of the width and height of the video cannot exceed 1,440 pixels. For example, a resolution of 1560 × 1560 pixels is invalid.
	//
	// >
	//
	// 	- This parameter is required when the TemplateType parameter is set to h264, h264-nbhd, or h264-origin.
	//
	// 	- If the TemplateType parameter is set to h264-origin, the highest resolution supported is 4K.
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
