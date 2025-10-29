// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomLiveStreamTranscodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApp(v string) *UpdateCustomLiveStreamTranscodeRequest
	GetApp() *string
	SetAudioBitrate(v int32) *UpdateCustomLiveStreamTranscodeRequest
	GetAudioBitrate() *int32
	SetAudioChannelNum(v int32) *UpdateCustomLiveStreamTranscodeRequest
	GetAudioChannelNum() *int32
	SetAudioCodec(v string) *UpdateCustomLiveStreamTranscodeRequest
	GetAudioCodec() *string
	SetAudioProfile(v string) *UpdateCustomLiveStreamTranscodeRequest
	GetAudioProfile() *string
	SetAudioRate(v int32) *UpdateCustomLiveStreamTranscodeRequest
	GetAudioRate() *int32
	SetBitrateWithSource(v string) *UpdateCustomLiveStreamTranscodeRequest
	GetBitrateWithSource() *string
	SetDeInterlaced(v bool) *UpdateCustomLiveStreamTranscodeRequest
	GetDeInterlaced() *bool
	SetDomain(v string) *UpdateCustomLiveStreamTranscodeRequest
	GetDomain() *string
	SetEncryptParameters(v string) *UpdateCustomLiveStreamTranscodeRequest
	GetEncryptParameters() *string
	SetExtWithSource(v string) *UpdateCustomLiveStreamTranscodeRequest
	GetExtWithSource() *string
	SetFPS(v int32) *UpdateCustomLiveStreamTranscodeRequest
	GetFPS() *int32
	SetFpsWithSource(v string) *UpdateCustomLiveStreamTranscodeRequest
	GetFpsWithSource() *string
	SetGop(v string) *UpdateCustomLiveStreamTranscodeRequest
	GetGop() *string
	SetHeight(v int32) *UpdateCustomLiveStreamTranscodeRequest
	GetHeight() *int32
	SetLazy(v string) *UpdateCustomLiveStreamTranscodeRequest
	GetLazy() *string
	SetOwnerId(v int64) *UpdateCustomLiveStreamTranscodeRequest
	GetOwnerId() *int64
	SetProfile(v int32) *UpdateCustomLiveStreamTranscodeRequest
	GetProfile() *int32
	SetRegionId(v string) *UpdateCustomLiveStreamTranscodeRequest
	GetRegionId() *string
	SetResWithSource(v string) *UpdateCustomLiveStreamTranscodeRequest
	GetResWithSource() *string
	SetTemplate(v string) *UpdateCustomLiveStreamTranscodeRequest
	GetTemplate() *string
	SetTemplateType(v string) *UpdateCustomLiveStreamTranscodeRequest
	GetTemplateType() *string
	SetVideoBitrate(v int32) *UpdateCustomLiveStreamTranscodeRequest
	GetVideoBitrate() *int32
	SetWidth(v int32) *UpdateCustomLiveStreamTranscodeRequest
	GetWidth() *int32
}

type UpdateCustomLiveStreamTranscodeRequest struct {
	// The name of the application to which the stream belongs, and it cannot be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// Audio transcoding bitrate. Unit: kbps, value range: 1 to 1000.
	//
	// example:
	//
	// 512
	AudioBitrate *int32 `json:"AudioBitrate,omitempty" xml:"AudioBitrate,omitempty"`
	// Number of audio channels. Values:
	//
	// - 1: Mono.
	//
	// - 2: Stereo.
	//
	// example:
	//
	// 2
	AudioChannelNum *int32 `json:"AudioChannelNum,omitempty" xml:"AudioChannelNum,omitempty"`
	// Audio encoding format. Values:
	//
	// - aac
	//
	// - mp3
	//
	// example:
	//
	// aac
	AudioCodec *string `json:"AudioCodec,omitempty" xml:"AudioCodec,omitempty"`
	// Audio encoding. Values:
	//
	// - aac_low
	//
	//  - aac_he
	//
	//  - aac_he_v2
	//
	// - aac_ld
	//
	// example:
	//
	// aac_low
	AudioProfile *string `json:"AudioProfile,omitempty" xml:"AudioProfile,omitempty"`
	// Audio sampling rate. Values: 22050, 32000, 44100, 48000, 96000. Unit: Hz.
	//
	// > If **AudioProfile*	- is set to **aac_ld**, the sampling rate must not exceed 44100.
	//
	// example:
	//
	// 96000
	AudioRate *int32 `json:"AudioRate,omitempty" xml:"AudioRate,omitempty"`
	// The source-based bitrate settings. This parameter takes precedence over other bitrate settings. The following fields must be included:
	//
	// 	- **UpLimit**: the maximum bitrate. Set this field to an integer from 128 to 10000. The value must be greater than the minimum bitrate.
	//
	// 	- **LowerLimit**: the minimum bitrate. Set this field to an integer from 128 to 10000. The value must be smaller than the maximum bitrate.
	//
	// 	- **Factor**: the ratio of the output bitrate to the source bitrate. Valid values: 0.1 to 1. The value is accurate to one decimal place. A value of 1 indicates that the output video has the same bitrate as the source video.
	//
	// example:
	//
	// {\\"UpLimit\\":2500,\\"LowerLimit\\":800,\\"Factor\\":1}
	BitrateWithSource *string `json:"BitrateWithSource,omitempty" xml:"BitrateWithSource,omitempty"`
	DeInterlaced      *bool   `json:"DeInterlaced,omitempty" xml:"DeInterlaced,omitempty"`
	// Streamer domain name, unmodifiable.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Encryption configuration. JSON format, with the following fields:
	//
	// -  EncryptType: Type of encryption. Fixed value is aliyun.
	//
	//  -  KmsKeyID: User\\"s KMS master key ID.
	//
	// -  KmsKeyExpireInterval: Key rotation period. Value range: 60~3600, unit: seconds.
	//
	// > When using DRM encryption, KmsKeyID cannot be modified.
	//
	// example:
	//
	// {"EncryptType": "aliyun", "KmsKeyID":"afce5722-81d2-43c3-9930-7601da11****","KmsKeyExpireInterval":"3600"}
	EncryptParameters *string `json:"EncryptParameters,omitempty" xml:"EncryptParameters,omitempty"`
	// Other source-based settings. The following fields are included:
	//
	// 	- **KeyFrameOpen**: Valid values: yes and no.
	//
	// 	- **Copyts**: Valid values: yes and no.
	//
	// 	- **SeiMode**: Valid values: 0, 1, and 2. 0 specifies that no supplemental enhancement information (SEI) messages are passed through, 1 specifies that part of SEI messages are passed through, and 2 specifies that all SEI messages are passed through.
	//
	// example:
	//
	// {\\"KeyFrameOpen\\":\\"yes\\",\\"Copyts\\":\\"yes\\",\\"SeiMode\\":1}
	ExtWithSource *string `json:"ExtWithSource,omitempty" xml:"ExtWithSource,omitempty"`
	// Transcode video frame rate. Unit: FPS, value range: 1 to 60.
	//
	// example:
	//
	// 30
	FPS *int32 `json:"FPS,omitempty" xml:"FPS,omitempty"`
	// The source-based frame rate settings. This parameter takes precedence over other frame rate settings. The following fields must be included:
	//
	// 	- **UpLimit**: the maximum frame rate. Set this field to an integer from 1 to 60. The value must be greater than the minimum frame rate.
	//
	// 	- **LowerLimit**: the minimum frame rate. Set this field to an integer from 1 to 60. The value must be smaller than the maximum frame rate.
	//
	// example:
	//
	// {\\"UpLimit\\":60,\\"LowerLimit\\":1}
	FpsWithSource *string `json:"FpsWithSource,omitempty" xml:"FpsWithSource,omitempty"`
	// Video GOP (Group of Pictures), supports units in frames or seconds. When the unit is frames, the value should be {number}; when the unit is seconds, the value should be {number}s.
	//
	// - For frames, the range is 1 to 3000.
	//
	// - For seconds, the range is 1 to 20s.
	//
	// example:
	//
	// 1
	Gop *string `json:"Gop,omitempty" xml:"Gop,omitempty"`
	// Video transcoding height. Unit: pixels. The value must meet the following three conditions:
	//
	//  - Height ≥ 100: The height of the video must be no less than 100 pixels.
	//
	//  - max(Height, Width) ≤ 2560: The larger of the video\\"s width and height cannot exceed 2560.
	//
	//  - min(Height, Width) ≤ 1440: The smaller of the video\\"s width and height cannot exceed 1440.
	//
	//  > For 265 narrowband HD templates, the maximum resolution is 1280×720.
	//
	// example:
	//
	// 720
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// Specifies whether to enable triggered transcoding. Valid values:
	//
	// 	- **yes**: enables triggered transcoding.
	//
	// 	- **no**: disables triggered transcoding.
	//
	// example:
	//
	// yes
	Lazy    *string `json:"Lazy,omitempty" xml:"Lazy,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Encoding level. A set of specific encoding features supported by the video, generally, the higher the value, the better the picture quality, but also the higher the resources consumed for encoding and decoding. Values:
	//
	// - 1: baseline (suitable for mobile devices).
	//
	//  - 2: main (suitable for standard resolution devices).
	//
	// - 3: high (suitable for high-resolution devices).
	//
	// example:
	//
	// 2
	Profile  *int32  `json:"Profile,omitempty" xml:"Profile,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The source-based resolution settings. This parameter takes precedence over other resolution settings. The following fields must be included:
	//
	// 	- **Type**: You can set this field to short, long, or screen. short specifies that the resolution of the output video is adapted to the shorter side, long specifies that the resolution of the output video is adapted to the longer side, and screen specifies that the output video has an adaptive resolution.
	//
	// 	- **Value**:
	//
	//     	- Set this field to 360, 480, 540, 720, or 1080 if the Type field is set to short.
	//
	//     	- Set this field to 640, 848, 960, 1280, or 1920 if the Type field is set to long.
	//
	//     	- Set this field to 640\\*360, 848\\*480, 960\\*540, 1280\\*720, or 1920\\*1080 if the Type field is set to screen.
	//
	// example:
	//
	// {\\"Type\\":\\"short\\",\\"Value\\":\\"1080\\"}
	ResWithSource *string `json:"ResWithSource,omitempty" xml:"ResWithSource,omitempty"`
	// Custom name of the transcoding template, not modifiable.
	//
	// This parameter is required.
	//
	// example:
	//
	// LiveCusTranscode****
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// Custom transcoding template type, unmodifiable.
	//
	// This parameter is required.
	//
	// example:
	//
	// h264
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// Video transcoding bitrate. Unit: kbps, value range: 1 to 6000.
	//
	// > The actual bitrate of the transcoded video will try to be as close as possible to the one you set, but it cannot be guaranteed to be exactly the same, especially when the set bitrate is too high or too low.
	//
	// example:
	//
	// 720
	VideoBitrate *int32 `json:"VideoBitrate,omitempty" xml:"VideoBitrate,omitempty"`
	// Video transcoding width. Unit: pixels. The value must meet the following three conditions:
	//
	// - Width ≥ 100: The video width must be no less than 100 pixels.
	//
	// - max(Height, Width) ≤ 2560: The larger of the video\\"s height and width cannot exceed 2560.
	//
	// - min(Height, Width) ≤ 1440: The smaller of the video\\"s height and width cannot exceed 1440.
	//
	// > For 265 narrowband HD templates, the maximum resolution is 1280×720.
	//
	// example:
	//
	// 576
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s UpdateCustomLiveStreamTranscodeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomLiveStreamTranscodeRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomLiveStreamTranscodeRequest) GetApp() *string {
	return s.App
}

func (s *UpdateCustomLiveStreamTranscodeRequest) GetAudioBitrate() *int32 {
	return s.AudioBitrate
}

func (s *UpdateCustomLiveStreamTranscodeRequest) GetAudioChannelNum() *int32 {
	return s.AudioChannelNum
}

func (s *UpdateCustomLiveStreamTranscodeRequest) GetAudioCodec() *string {
	return s.AudioCodec
}

func (s *UpdateCustomLiveStreamTranscodeRequest) GetAudioProfile() *string {
	return s.AudioProfile
}

func (s *UpdateCustomLiveStreamTranscodeRequest) GetAudioRate() *int32 {
	return s.AudioRate
}

func (s *UpdateCustomLiveStreamTranscodeRequest) GetBitrateWithSource() *string {
	return s.BitrateWithSource
}

func (s *UpdateCustomLiveStreamTranscodeRequest) GetDeInterlaced() *bool {
	return s.DeInterlaced
}

func (s *UpdateCustomLiveStreamTranscodeRequest) GetDomain() *string {
	return s.Domain
}

func (s *UpdateCustomLiveStreamTranscodeRequest) GetEncryptParameters() *string {
	return s.EncryptParameters
}

func (s *UpdateCustomLiveStreamTranscodeRequest) GetExtWithSource() *string {
	return s.ExtWithSource
}

func (s *UpdateCustomLiveStreamTranscodeRequest) GetFPS() *int32 {
	return s.FPS
}

func (s *UpdateCustomLiveStreamTranscodeRequest) GetFpsWithSource() *string {
	return s.FpsWithSource
}

func (s *UpdateCustomLiveStreamTranscodeRequest) GetGop() *string {
	return s.Gop
}

func (s *UpdateCustomLiveStreamTranscodeRequest) GetHeight() *int32 {
	return s.Height
}

func (s *UpdateCustomLiveStreamTranscodeRequest) GetLazy() *string {
	return s.Lazy
}

func (s *UpdateCustomLiveStreamTranscodeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateCustomLiveStreamTranscodeRequest) GetProfile() *int32 {
	return s.Profile
}

func (s *UpdateCustomLiveStreamTranscodeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateCustomLiveStreamTranscodeRequest) GetResWithSource() *string {
	return s.ResWithSource
}

func (s *UpdateCustomLiveStreamTranscodeRequest) GetTemplate() *string {
	return s.Template
}

func (s *UpdateCustomLiveStreamTranscodeRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *UpdateCustomLiveStreamTranscodeRequest) GetVideoBitrate() *int32 {
	return s.VideoBitrate
}

func (s *UpdateCustomLiveStreamTranscodeRequest) GetWidth() *int32 {
	return s.Width
}

func (s *UpdateCustomLiveStreamTranscodeRequest) SetApp(v string) *UpdateCustomLiveStreamTranscodeRequest {
	s.App = &v
	return s
}

func (s *UpdateCustomLiveStreamTranscodeRequest) SetAudioBitrate(v int32) *UpdateCustomLiveStreamTranscodeRequest {
	s.AudioBitrate = &v
	return s
}

func (s *UpdateCustomLiveStreamTranscodeRequest) SetAudioChannelNum(v int32) *UpdateCustomLiveStreamTranscodeRequest {
	s.AudioChannelNum = &v
	return s
}

func (s *UpdateCustomLiveStreamTranscodeRequest) SetAudioCodec(v string) *UpdateCustomLiveStreamTranscodeRequest {
	s.AudioCodec = &v
	return s
}

func (s *UpdateCustomLiveStreamTranscodeRequest) SetAudioProfile(v string) *UpdateCustomLiveStreamTranscodeRequest {
	s.AudioProfile = &v
	return s
}

func (s *UpdateCustomLiveStreamTranscodeRequest) SetAudioRate(v int32) *UpdateCustomLiveStreamTranscodeRequest {
	s.AudioRate = &v
	return s
}

func (s *UpdateCustomLiveStreamTranscodeRequest) SetBitrateWithSource(v string) *UpdateCustomLiveStreamTranscodeRequest {
	s.BitrateWithSource = &v
	return s
}

func (s *UpdateCustomLiveStreamTranscodeRequest) SetDeInterlaced(v bool) *UpdateCustomLiveStreamTranscodeRequest {
	s.DeInterlaced = &v
	return s
}

func (s *UpdateCustomLiveStreamTranscodeRequest) SetDomain(v string) *UpdateCustomLiveStreamTranscodeRequest {
	s.Domain = &v
	return s
}

func (s *UpdateCustomLiveStreamTranscodeRequest) SetEncryptParameters(v string) *UpdateCustomLiveStreamTranscodeRequest {
	s.EncryptParameters = &v
	return s
}

func (s *UpdateCustomLiveStreamTranscodeRequest) SetExtWithSource(v string) *UpdateCustomLiveStreamTranscodeRequest {
	s.ExtWithSource = &v
	return s
}

func (s *UpdateCustomLiveStreamTranscodeRequest) SetFPS(v int32) *UpdateCustomLiveStreamTranscodeRequest {
	s.FPS = &v
	return s
}

func (s *UpdateCustomLiveStreamTranscodeRequest) SetFpsWithSource(v string) *UpdateCustomLiveStreamTranscodeRequest {
	s.FpsWithSource = &v
	return s
}

func (s *UpdateCustomLiveStreamTranscodeRequest) SetGop(v string) *UpdateCustomLiveStreamTranscodeRequest {
	s.Gop = &v
	return s
}

func (s *UpdateCustomLiveStreamTranscodeRequest) SetHeight(v int32) *UpdateCustomLiveStreamTranscodeRequest {
	s.Height = &v
	return s
}

func (s *UpdateCustomLiveStreamTranscodeRequest) SetLazy(v string) *UpdateCustomLiveStreamTranscodeRequest {
	s.Lazy = &v
	return s
}

func (s *UpdateCustomLiveStreamTranscodeRequest) SetOwnerId(v int64) *UpdateCustomLiveStreamTranscodeRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateCustomLiveStreamTranscodeRequest) SetProfile(v int32) *UpdateCustomLiveStreamTranscodeRequest {
	s.Profile = &v
	return s
}

func (s *UpdateCustomLiveStreamTranscodeRequest) SetRegionId(v string) *UpdateCustomLiveStreamTranscodeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateCustomLiveStreamTranscodeRequest) SetResWithSource(v string) *UpdateCustomLiveStreamTranscodeRequest {
	s.ResWithSource = &v
	return s
}

func (s *UpdateCustomLiveStreamTranscodeRequest) SetTemplate(v string) *UpdateCustomLiveStreamTranscodeRequest {
	s.Template = &v
	return s
}

func (s *UpdateCustomLiveStreamTranscodeRequest) SetTemplateType(v string) *UpdateCustomLiveStreamTranscodeRequest {
	s.TemplateType = &v
	return s
}

func (s *UpdateCustomLiveStreamTranscodeRequest) SetVideoBitrate(v int32) *UpdateCustomLiveStreamTranscodeRequest {
	s.VideoBitrate = &v
	return s
}

func (s *UpdateCustomLiveStreamTranscodeRequest) SetWidth(v int32) *UpdateCustomLiveStreamTranscodeRequest {
	s.Width = &v
	return s
}

func (s *UpdateCustomLiveStreamTranscodeRequest) Validate() error {
	return dara.Validate(s)
}
