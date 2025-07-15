// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRtsLiveStreamTranscodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApp(v string) *AddRtsLiveStreamTranscodeRequest
	GetApp() *string
	SetAudioBitrate(v int32) *AddRtsLiveStreamTranscodeRequest
	GetAudioBitrate() *int32
	SetAudioChannelNum(v int32) *AddRtsLiveStreamTranscodeRequest
	GetAudioChannelNum() *int32
	SetAudioCodec(v string) *AddRtsLiveStreamTranscodeRequest
	GetAudioCodec() *string
	SetAudioProfile(v string) *AddRtsLiveStreamTranscodeRequest
	GetAudioProfile() *string
	SetAudioRate(v int32) *AddRtsLiveStreamTranscodeRequest
	GetAudioRate() *int32
	SetDeleteBframes(v bool) *AddRtsLiveStreamTranscodeRequest
	GetDeleteBframes() *bool
	SetDomain(v string) *AddRtsLiveStreamTranscodeRequest
	GetDomain() *string
	SetFPS(v int32) *AddRtsLiveStreamTranscodeRequest
	GetFPS() *int32
	SetGop(v string) *AddRtsLiveStreamTranscodeRequest
	GetGop() *string
	SetHeight(v int32) *AddRtsLiveStreamTranscodeRequest
	GetHeight() *int32
	SetLazy(v string) *AddRtsLiveStreamTranscodeRequest
	GetLazy() *string
	SetOpus(v bool) *AddRtsLiveStreamTranscodeRequest
	GetOpus() *bool
	SetOwnerId(v int64) *AddRtsLiveStreamTranscodeRequest
	GetOwnerId() *int64
	SetProfile(v int32) *AddRtsLiveStreamTranscodeRequest
	GetProfile() *int32
	SetRegionId(v string) *AddRtsLiveStreamTranscodeRequest
	GetRegionId() *string
	SetTemplate(v string) *AddRtsLiveStreamTranscodeRequest
	GetTemplate() *string
	SetTemplateType(v string) *AddRtsLiveStreamTranscodeRequest
	GetTemplateType() *string
	SetVideoBitrate(v int32) *AddRtsLiveStreamTranscodeRequest
	GetVideoBitrate() *int32
	SetWidth(v int32) *AddRtsLiveStreamTranscodeRequest
	GetWidth() *int32
}

type AddRtsLiveStreamTranscodeRequest struct {
	// The name of the application to which the live stream belongs. Value requirements:
	//
	// 	- The name can contain letters, digits, underscores (_), and hyphens (-).
	//
	// 	- We recommend that you specify a name that is more than three characters in length. The name must start with a letter or digit.
	//
	// This parameter is required.
	//
	// example:
	//
	// aliyun-test
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// The bitrate of the output audio. Unit: Kbit/s. Valid values: **1*	- to **1000**.
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
	// 	- **aac**
	//
	// 	- **mp3**
	//
	// > If you want to use the Opus encoding format, set the Opus parameter to true.
	//
	// example:
	//
	// aac
	AudioCodec *string `json:"AudioCodec,omitempty" xml:"AudioCodec,omitempty"`
	// The audio codec profile. Valid values:
	//
	// 	- **aac_low**
	//
	// 	- **aac_he**
	//
	// 	- **aac_he_v2**
	//
	// 	- **aac_ld**
	//
	// example:
	//
	// aac_low
	AudioProfile *string `json:"AudioProfile,omitempty" xml:"AudioProfile,omitempty"`
	// The audio sampling rate. Valid values: **22050 to 96000**. The value 44100 is commonly used. Unit: Hz.
	//
	// > If you set the AudioProfile parameter to aac_ld, the audio sampling rate cannot exceed 44,100 Hz.
	//
	// example:
	//
	// 44100
	AudioRate *int32 `json:"AudioRate,omitempty" xml:"AudioRate,omitempty"`
	// Specifies whether to remove B-frames during transcoding. Valid values:
	//
	// >  This parameter is required if you set the TemplateType parameter to h264, h264-nbhd, or h264-origin.
	//
	// 	- **true**: removes B-frames.
	//
	// 	- **false**: retains B-frames. This is the default value.
	//
	// > If you do not specify this parameter, the default value **false*	- is used.
	//
	// example:
	//
	// false
	DeleteBframes *bool `json:"DeleteBframes,omitempty" xml:"DeleteBframes,omitempty"`
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The frame rate of the output video. Unit: FPS. Valid values: **1*	- to **60**.
	//
	// >  This parameter is required if you set the TemplateType parameter to h264, h264-nbhd, or h264-origin.
	//
	// example:
	//
	// 30
	FPS *int32 `json:"FPS,omitempty" xml:"FPS,omitempty"`
	// The group of pictures (GOP) size of the output video. This parameter is used to specify the keyframe interval. Unit: seconds. Valid values: **1*	- to **3**.
	//
	// example:
	//
	// 2
	Gop *string `json:"Gop,omitempty" xml:"Gop,omitempty"`
	// The height of the output video. Unit: pixels.
	//
	// >  This parameter is required if you set the TemplateType parameter to h264, h264-nbhd, or h264-origin.
	//
	// The value must comply with the following rules:****
	//
	// 	- **Height ≥ 100**: The height of the video is no less than 100 pixels.
	//
	// 	- **max(Height,Width) ≤ 2560**: The width or height of the video, whichever is greater, cannot exceed 2,560 pixels.
	//
	// 	- **min(Height,Width) ≤ 1440**: The width or height of the video, whichever is smaller, cannot exceed 1,440 pixels.
	//
	// For example, a resolution of 1560 × 1560 pixels is invalid.
	//
	// > An original quality template needs to retain the source information. Therefore, the video resolution cannot exceed 4K.
	//
	// example:
	//
	// 1280
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// Specifies whether to trigger transcoding only when a stream is pulled. Valid values:
	//
	// 	- **yes**: triggers transcoding only when a stream is pulled.
	//
	// 	- **no**: triggers transcoding whenever a stream is ingested, no matter whether the stream is pulled.
	//
	// example:
	//
	// no
	Lazy *string `json:"Lazy,omitempty" xml:"Lazy,omitempty"`
	// Specifies whether to transcode audio to the Opus format to be compatible with native WebRTC. Valid values:
	//
	// 	- **true**: transcodes audio to the Opus format.
	//
	// 	- **false**: does not transcode audio to the Opus format.
	//
	// > If you do not specify this parameter, the default value **false*	- is used.
	//
	// example:
	//
	// true
	Opus    *bool  `json:"Opus,omitempty" xml:"Opus,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The codec profile. The video codec profile determines how ApsaraVideo Live performs codec on the video. In normal cases, a greater value indicates a higher image quality and requires more codec resources. Valid values:
	//
	// 	- **1**: baseline, which is suitable for mobile devices.
	//
	// 	- **2**: main, which is suitable for standard-definition devices.
	//
	// 	- **3**: high, which is suitable for high-definition devices.
	//
	// example:
	//
	// 2
	Profile  *int32  `json:"Profile,omitempty" xml:"Profile,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the transcoding template. Value requirements:
	//
	// 	- The name can contain letters, digits, underscores (_), and hyphens (-).
	//
	// 	- We recommend that you specify a name that is more than three characters in length. The name must start with a letter or digit.
	//
	// > The name cannot be the same as that of a default transcoding template.
	//
	// This parameter is required.
	//
	// example:
	//
	// ld
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The type of the transcoding template.
	//
	// If you set this parameter to h264, h264-nbhd, or h264-origin, you must also specify the Height, Width, FPS, VideoBitrate, and DeleteBframes parameters. Valid values:
	//
	// 	- **h264**: H.264 standard transcoding template.
	//
	// 	- **h264-nbhd**: H.264 Narrowband HD™ transcoding template.
	//
	// 	- **h264-origin**: H.264 original quality template. If you use this type of template, the same transcoding parameters of the video source are retained by default.
	//
	// 	- **audio**: audio-only transcoding template. If you use this type of template, images are removed from the video source and an audio-only stream is generated. In addition, you must also specify the AudioBitrate parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// h264
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// The bitrate of the output video. Unit: Kbit/s. Valid values: **1*	- to **6000**.
	//
	// >  This parameter is required if you set the TemplateType parameter to h264, h264-nbhd, or h264-origin.
	//
	// > The bitrate of the output video is as close to the value that you specify as possible, but not exactly the same as the value, especially when the value is excessively large or small.
	//
	// example:
	//
	// 2000
	VideoBitrate *int32 `json:"VideoBitrate,omitempty" xml:"VideoBitrate,omitempty"`
	// The width of the output video. Unit: pixels.
	//
	// >  This parameter is required if you set the TemplateType parameter to h264, h264-nbhd, or h264-origin.
	//
	// The value must comply with the following rules:
	//
	// 	- **Width ≥ 100**: The width of the video is no less than 100 pixels.
	//
	// 	- **max(Height,Width) ≤ 2560**: The width or height of the video, whichever is greater, cannot exceed 2,560 pixels.
	//
	// 	- **min(Height,Width) ≤ 1440**: The width or height of the video, whichever is smaller, cannot exceed 1,440 pixels.
	//
	// For example, a resolution of 1560 × 1560 pixels is invalid.
	//
	// > An original quality template needs to retain the source information. Therefore, the video resolution cannot exceed 4K.
	//
	// example:
	//
	// 720
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s AddRtsLiveStreamTranscodeRequest) String() string {
	return dara.Prettify(s)
}

func (s AddRtsLiveStreamTranscodeRequest) GoString() string {
	return s.String()
}

func (s *AddRtsLiveStreamTranscodeRequest) GetApp() *string {
	return s.App
}

func (s *AddRtsLiveStreamTranscodeRequest) GetAudioBitrate() *int32 {
	return s.AudioBitrate
}

func (s *AddRtsLiveStreamTranscodeRequest) GetAudioChannelNum() *int32 {
	return s.AudioChannelNum
}

func (s *AddRtsLiveStreamTranscodeRequest) GetAudioCodec() *string {
	return s.AudioCodec
}

func (s *AddRtsLiveStreamTranscodeRequest) GetAudioProfile() *string {
	return s.AudioProfile
}

func (s *AddRtsLiveStreamTranscodeRequest) GetAudioRate() *int32 {
	return s.AudioRate
}

func (s *AddRtsLiveStreamTranscodeRequest) GetDeleteBframes() *bool {
	return s.DeleteBframes
}

func (s *AddRtsLiveStreamTranscodeRequest) GetDomain() *string {
	return s.Domain
}

func (s *AddRtsLiveStreamTranscodeRequest) GetFPS() *int32 {
	return s.FPS
}

func (s *AddRtsLiveStreamTranscodeRequest) GetGop() *string {
	return s.Gop
}

func (s *AddRtsLiveStreamTranscodeRequest) GetHeight() *int32 {
	return s.Height
}

func (s *AddRtsLiveStreamTranscodeRequest) GetLazy() *string {
	return s.Lazy
}

func (s *AddRtsLiveStreamTranscodeRequest) GetOpus() *bool {
	return s.Opus
}

func (s *AddRtsLiveStreamTranscodeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddRtsLiveStreamTranscodeRequest) GetProfile() *int32 {
	return s.Profile
}

func (s *AddRtsLiveStreamTranscodeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddRtsLiveStreamTranscodeRequest) GetTemplate() *string {
	return s.Template
}

func (s *AddRtsLiveStreamTranscodeRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *AddRtsLiveStreamTranscodeRequest) GetVideoBitrate() *int32 {
	return s.VideoBitrate
}

func (s *AddRtsLiveStreamTranscodeRequest) GetWidth() *int32 {
	return s.Width
}

func (s *AddRtsLiveStreamTranscodeRequest) SetApp(v string) *AddRtsLiveStreamTranscodeRequest {
	s.App = &v
	return s
}

func (s *AddRtsLiveStreamTranscodeRequest) SetAudioBitrate(v int32) *AddRtsLiveStreamTranscodeRequest {
	s.AudioBitrate = &v
	return s
}

func (s *AddRtsLiveStreamTranscodeRequest) SetAudioChannelNum(v int32) *AddRtsLiveStreamTranscodeRequest {
	s.AudioChannelNum = &v
	return s
}

func (s *AddRtsLiveStreamTranscodeRequest) SetAudioCodec(v string) *AddRtsLiveStreamTranscodeRequest {
	s.AudioCodec = &v
	return s
}

func (s *AddRtsLiveStreamTranscodeRequest) SetAudioProfile(v string) *AddRtsLiveStreamTranscodeRequest {
	s.AudioProfile = &v
	return s
}

func (s *AddRtsLiveStreamTranscodeRequest) SetAudioRate(v int32) *AddRtsLiveStreamTranscodeRequest {
	s.AudioRate = &v
	return s
}

func (s *AddRtsLiveStreamTranscodeRequest) SetDeleteBframes(v bool) *AddRtsLiveStreamTranscodeRequest {
	s.DeleteBframes = &v
	return s
}

func (s *AddRtsLiveStreamTranscodeRequest) SetDomain(v string) *AddRtsLiveStreamTranscodeRequest {
	s.Domain = &v
	return s
}

func (s *AddRtsLiveStreamTranscodeRequest) SetFPS(v int32) *AddRtsLiveStreamTranscodeRequest {
	s.FPS = &v
	return s
}

func (s *AddRtsLiveStreamTranscodeRequest) SetGop(v string) *AddRtsLiveStreamTranscodeRequest {
	s.Gop = &v
	return s
}

func (s *AddRtsLiveStreamTranscodeRequest) SetHeight(v int32) *AddRtsLiveStreamTranscodeRequest {
	s.Height = &v
	return s
}

func (s *AddRtsLiveStreamTranscodeRequest) SetLazy(v string) *AddRtsLiveStreamTranscodeRequest {
	s.Lazy = &v
	return s
}

func (s *AddRtsLiveStreamTranscodeRequest) SetOpus(v bool) *AddRtsLiveStreamTranscodeRequest {
	s.Opus = &v
	return s
}

func (s *AddRtsLiveStreamTranscodeRequest) SetOwnerId(v int64) *AddRtsLiveStreamTranscodeRequest {
	s.OwnerId = &v
	return s
}

func (s *AddRtsLiveStreamTranscodeRequest) SetProfile(v int32) *AddRtsLiveStreamTranscodeRequest {
	s.Profile = &v
	return s
}

func (s *AddRtsLiveStreamTranscodeRequest) SetRegionId(v string) *AddRtsLiveStreamTranscodeRequest {
	s.RegionId = &v
	return s
}

func (s *AddRtsLiveStreamTranscodeRequest) SetTemplate(v string) *AddRtsLiveStreamTranscodeRequest {
	s.Template = &v
	return s
}

func (s *AddRtsLiveStreamTranscodeRequest) SetTemplateType(v string) *AddRtsLiveStreamTranscodeRequest {
	s.TemplateType = &v
	return s
}

func (s *AddRtsLiveStreamTranscodeRequest) SetVideoBitrate(v int32) *AddRtsLiveStreamTranscodeRequest {
	s.VideoBitrate = &v
	return s
}

func (s *AddRtsLiveStreamTranscodeRequest) SetWidth(v int32) *AddRtsLiveStreamTranscodeRequest {
	s.Width = &v
	return s
}

func (s *AddRtsLiveStreamTranscodeRequest) Validate() error {
	return dara.Validate(s)
}
