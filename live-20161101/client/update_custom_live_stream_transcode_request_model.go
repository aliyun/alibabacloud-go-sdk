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
	// The AppName of the live stream. This parameter cannot be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// The output audio bitrate. Unit: kbps. Valid values: 1 to **1000**.
	//
	// example:
	//
	// 512
	AudioBitrate *int32 `json:"AudioBitrate,omitempty" xml:"AudioBitrate,omitempty"`
	// The number of audio channels. Valid values:
	//
	// - 1: mono.
	//
	// - 2: stereo.
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
	// example:
	//
	// aac
	AudioCodec *string `json:"AudioCodec,omitempty" xml:"AudioCodec,omitempty"`
	// The audio profile. Valid values:
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
	// The audio sample rate. Valid values: 22050, 32000, 44100, 48000, and 96000. Unit: Hz.
	//
	// > If you set **AudioProfile*	- to **aac_ld**, the sample rate cannot exceed 44100.
	//
	// example:
	//
	// 96000
	AudioRate *int32 `json:"AudioRate,omitempty" xml:"AudioRate,omitempty"`
	// The adaptive bitrate settings. If specified, it overrides the VideoBitrate parameter. Fields:
	//
	// - **UpLimit (integer):*	- Required. The upper limit of the bitrate. This must be an integer from 128 to 10000 and greater than the lower limit.
	//
	// - **LowerLimit (integer):*	- Required. The lower limit of the bitrate. This must be an integer from 128 to 10000 and less than the upper limit.
	//
	// - **Factor (float):*	- Required: The factor by which the source bitrate is multiplied to calculate the output bitrate. Valid values: 0.1 to 1. The value can be accurate to one decimal place. A value of 1 indicates that the output bitrate is the same as the source bitrate.
	//
	// example:
	//
	// {"UpLimit":2500,"LowerLimit":800,"Factor":1}
	BitrateWithSource *string `json:"BitrateWithSource,omitempty" xml:"BitrateWithSource,omitempty"`
	// Specifies whether to automatically detect and remove interlacing during transcoding. Deinterlacing converts interlaced video into progressive video.
	//
	// - true: enables deinterlacing.
	//
	// - false: keeps the source format. This is the default value.
	DeInterlaced *bool `json:"DeInterlaced,omitempty" xml:"DeInterlaced,omitempty"`
	// The streaming domain. This parameter cannot be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The encryption settings, formatted as a JSON string.
	//
	// - **EncryptType**: The encryption type. Set the value to aliyun.
	//
	// - **KmsKeyID**: The ID of the customer master key (CMK) in Key Management Service (KMS).
	//
	// - **KmsKeyExpireInterval**: The key rotation period. Unit: seconds. Valid values: **60 to 3600.**
	//
	// > When you use Digital Rights Management (DRM) encryption, you cannot modify KmsKeyID.
	//
	// example:
	//
	// {"EncryptType": "aliyun", "KmsKeyID":"afce5722-81d2-43c3-9930-7601da11****","KmsKeyExpireInterval":"3600"}
	EncryptParameters *string `json:"EncryptParameters,omitempty" xml:"EncryptParameters,omitempty"`
	// Other adaptive settings that align the transcoded stream with the source stream. Fields:
	//
	// - **KeyFrameOpen**: Specifies whether to align keyframes with the source stream. Valid values: yes and no.
	//
	// - **Copyts (string)**: Specifies whether to align the presentation timestamp (PTS) with the source stream. Valid values: yes and no.
	//
	// - **SeiMode**: The pass-through mode for Supplemental Enhancement Information (SEI). Valid values: 0 (disabled), 1 (pass through partial parameters), and 2 (pass through all).
	//
	// example:
	//
	// {"KeyFrameOpen":"yes","Copyts":"yes","SeiMode":1}
	ExtWithSource *string `json:"ExtWithSource,omitempty" xml:"ExtWithSource,omitempty"`
	// The frame rate of the output video. Unit: frames per second (FPS). Valid values: 1 to **60**.
	//
	// example:
	//
	// 30
	FPS *int32 `json:"FPS,omitempty" xml:"FPS,omitempty"`
	// Adapts the output frame rate based on the source\\"s frame rate, while keeping it within a specified range. If specified, it overrides the FPS parameter. Fields:
	//
	// - **UpLimit (integer):*	- Required. The upper limit of the frame rate. This must be an integer from 1 to 60 and greater than the lower limit.
	//
	// - **LowerLimit (integer):*	- Required. The lower limit of the frame rate. This must be an integer from 1 to 60 and less than the upper limit.
	//
	// example:
	//
	// {"UpLimit":60,"LowerLimit":1}
	FpsWithSource *string `json:"FpsWithSource,omitempty" xml:"FpsWithSource,omitempty"`
	// The Group of Pictures (GOP) size. The unit can be frame or second. Valid values:
	//
	// - By frames: 1 to 3000.
	//
	// - By seconds: 1s to 20s.
	//
	// example:
	//
	// 1
	Gop *string `json:"Gop,omitempty" xml:"Gop,omitempty"`
	// Output video height in pixels. Requirements:
	//
	// - **Height ≥ 100**
	//
	// - **max(Height, Width) ≤ 2560**
	//
	// - **min(Height, Width) ≤ 1440**
	//
	// > For h265-nbhd templates, it cannot exceed 720.
	//
	// example:
	//
	// 720
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// Specifies whether to enable on-demand transcoding. Valid values:
	//
	// - **yes**: Transcoding only starts when the first viewer requests this transcoded stream.
	//
	// - **no**: Transcoding starts immediately after the stream is published.
	//
	// example:
	//
	// yes
	Lazy    *string `json:"Lazy,omitempty" xml:"Lazy,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// The adaptive resolution settings. If specified, it overrides the Height and Width parameters. Fieds:
	//
	// - **Type (string):*	- Required. Valid values:
	//
	//   - **short**: sets the shorter edge of the video to the specified value and scales the other edge to maintain the original aspect ratio.
	//
	//   - **long**: sets the longer edge of the video to the specified value and scales the other edge to maintain the original aspect ratio.
	//
	//   - **screen**: Matches the output to a standard resolution, automatically flipping the dimensions based on the source\\"s orientation.
	//
	// - **Value (string):*	- Required. Valid values:
	//
	//   - For short: 360, 480, 540, 720, and 1080.
	//
	//   - For long: 640, 848, 960, 1280, and 1920.
	//
	//   - For screen: 640×360, 848×480, 960×540, 1280×720, and 1920×1080.
	//
	// example:
	//
	// {"Type":"short","Value":"1080"}
	ResWithSource *string `json:"ResWithSource,omitempty" xml:"ResWithSource,omitempty"`
	// The custom name of the transcoding template. This parameter cannot be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// LiveCusTranscode****
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
	// > The system tries to transcode the video at the specified bitrate. However, the actual bitrate may not be the same as the specified value, especially when the specified value is too high or too low.
	//
	// example:
	//
	// 720
	VideoBitrate *int32 `json:"VideoBitrate,omitempty" xml:"VideoBitrate,omitempty"`
	// Output video width in pixels.
	//
	// Requirements:
	//
	// - **Width ≥ 100**
	//
	// - **max(Height, Width) ≤ 2560**
	//
	// - **min(Height, Width) ≤ 1440**
	//
	// > For h265-nbhd templates, it cannot exceed 1280.
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
