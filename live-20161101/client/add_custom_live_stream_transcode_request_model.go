// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomLiveStreamTranscodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApp(v string) *AddCustomLiveStreamTranscodeRequest
	GetApp() *string
	SetAudioBitrate(v int32) *AddCustomLiveStreamTranscodeRequest
	GetAudioBitrate() *int32
	SetAudioChannelNum(v int32) *AddCustomLiveStreamTranscodeRequest
	GetAudioChannelNum() *int32
	SetAudioCodec(v string) *AddCustomLiveStreamTranscodeRequest
	GetAudioCodec() *string
	SetAudioProfile(v string) *AddCustomLiveStreamTranscodeRequest
	GetAudioProfile() *string
	SetAudioRate(v int32) *AddCustomLiveStreamTranscodeRequest
	GetAudioRate() *int32
	SetBitrateWithSource(v string) *AddCustomLiveStreamTranscodeRequest
	GetBitrateWithSource() *string
	SetDeInterlaced(v bool) *AddCustomLiveStreamTranscodeRequest
	GetDeInterlaced() *bool
	SetDomain(v string) *AddCustomLiveStreamTranscodeRequest
	GetDomain() *string
	SetEncryptParameters(v string) *AddCustomLiveStreamTranscodeRequest
	GetEncryptParameters() *string
	SetExtWithSource(v string) *AddCustomLiveStreamTranscodeRequest
	GetExtWithSource() *string
	SetFPS(v int32) *AddCustomLiveStreamTranscodeRequest
	GetFPS() *int32
	SetFpsWithSource(v string) *AddCustomLiveStreamTranscodeRequest
	GetFpsWithSource() *string
	SetGop(v string) *AddCustomLiveStreamTranscodeRequest
	GetGop() *string
	SetHeight(v int32) *AddCustomLiveStreamTranscodeRequest
	GetHeight() *int32
	SetKmsKeyExpireInterval(v string) *AddCustomLiveStreamTranscodeRequest
	GetKmsKeyExpireInterval() *string
	SetKmsKeyID(v string) *AddCustomLiveStreamTranscodeRequest
	GetKmsKeyID() *string
	SetKmsUID(v string) *AddCustomLiveStreamTranscodeRequest
	GetKmsUID() *string
	SetLazy(v string) *AddCustomLiveStreamTranscodeRequest
	GetLazy() *string
	SetOwnerId(v int64) *AddCustomLiveStreamTranscodeRequest
	GetOwnerId() *int64
	SetProfile(v int32) *AddCustomLiveStreamTranscodeRequest
	GetProfile() *int32
	SetRegionId(v string) *AddCustomLiveStreamTranscodeRequest
	GetRegionId() *string
	SetResWithSource(v string) *AddCustomLiveStreamTranscodeRequest
	GetResWithSource() *string
	SetTemplate(v string) *AddCustomLiveStreamTranscodeRequest
	GetTemplate() *string
	SetTemplateType(v string) *AddCustomLiveStreamTranscodeRequest
	GetTemplateType() *string
	SetVideoBitrate(v int32) *AddCustomLiveStreamTranscodeRequest
	GetVideoBitrate() *int32
	SetWidth(v int32) *AddCustomLiveStreamTranscodeRequest
	GetWidth() *int32
}

type AddCustomLiveStreamTranscodeRequest struct {
	// The AppName of the live stream.
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
	// - **AAC**
	//
	// - **MP3**
	//
	// example:
	//
	// AAC
	AudioCodec *string `json:"AudioCodec,omitempty" xml:"AudioCodec,omitempty"`
	// The audio profile. Valid values:
	//
	// - **aac_low**
	//
	// - **aac_he**
	//
	// - **aac_he_v2**
	//
	// - **aac_ld**
	//
	// example:
	//
	// aac_low
	AudioProfile *string `json:"AudioProfile,omitempty" xml:"AudioProfile,omitempty"`
	// The audio sample rate. Valid values: **22050*	- to **96000**.
	//
	// 	Notice:
	//
	// If you set AudioProfile to **aac_ld**, the sample rate cannot exceed 44100.
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
	//
	// example:
	//
	// false
	DeInterlaced *bool `json:"DeInterlaced,omitempty" xml:"DeInterlaced,omitempty"`
	// The streaming domain.
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
	// > If set, its internal fields cannot be empty.
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
	// - **SeiMode**: The pass-through mode for Supplemental Enhancement Information (SEI). Valid values: 0 (disabled) and 1 (enabled).
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
	// > For h265-nbhd, it cannot exceed 720.
	//
	// example:
	//
	// 720
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The key rotation period. Unit: seconds. Valid values: 60 to 3600.
	//
	// example:
	//
	// 3600
	KmsKeyExpireInterval *string `json:"KmsKeyExpireInterval,omitempty" xml:"KmsKeyExpireInterval,omitempty"`
	// The ID of the customer master key (CMK) in Key Management Service (KMS).
	//
	// example:
	//
	// afce5722-81d2-43c3-9930-7601da11****
	KmsKeyID *string `json:"KmsKeyID,omitempty" xml:"KmsKeyID,omitempty"`
	// The ID of the KMS account.
	//
	// example:
	//
	// 25346073170691****
	KmsUID *string `json:"KmsUID,omitempty" xml:"KmsUID,omitempty"`
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
	// The custom name of the transcoding template.
	//
	// > The name can contain digits, letters, and hyphens (-). It must start with a digit or a letter. It cannot be the same as the name of a standard transcoding template.
	//
	// This parameter is required.
	//
	// example:
	//
	// LiveCusTranscode****
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The type of the custom transcoding template. Valid values:
	//
	// - **h264**: custom H.264 standard transcoding.
	//
	// - **h264-nbhd**: custom H.264 Narrowband HD™ transcoding.
	//
	// - **h265**: custom H.265 standard transcoding.
	//
	// - **h265-nbhd**: custom H.265 Narrowband HD™ transcoding.
	//
	// - **audio**: audio-only transcoding.
	//
	// > For video types, Height, Width, FPS, and VideoBitrate are required.
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
	// > For h265-nbhd, it cannot exceed 1280.
	//
	// example:
	//
	// 576
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s AddCustomLiveStreamTranscodeRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCustomLiveStreamTranscodeRequest) GoString() string {
	return s.String()
}

func (s *AddCustomLiveStreamTranscodeRequest) GetApp() *string {
	return s.App
}

func (s *AddCustomLiveStreamTranscodeRequest) GetAudioBitrate() *int32 {
	return s.AudioBitrate
}

func (s *AddCustomLiveStreamTranscodeRequest) GetAudioChannelNum() *int32 {
	return s.AudioChannelNum
}

func (s *AddCustomLiveStreamTranscodeRequest) GetAudioCodec() *string {
	return s.AudioCodec
}

func (s *AddCustomLiveStreamTranscodeRequest) GetAudioProfile() *string {
	return s.AudioProfile
}

func (s *AddCustomLiveStreamTranscodeRequest) GetAudioRate() *int32 {
	return s.AudioRate
}

func (s *AddCustomLiveStreamTranscodeRequest) GetBitrateWithSource() *string {
	return s.BitrateWithSource
}

func (s *AddCustomLiveStreamTranscodeRequest) GetDeInterlaced() *bool {
	return s.DeInterlaced
}

func (s *AddCustomLiveStreamTranscodeRequest) GetDomain() *string {
	return s.Domain
}

func (s *AddCustomLiveStreamTranscodeRequest) GetEncryptParameters() *string {
	return s.EncryptParameters
}

func (s *AddCustomLiveStreamTranscodeRequest) GetExtWithSource() *string {
	return s.ExtWithSource
}

func (s *AddCustomLiveStreamTranscodeRequest) GetFPS() *int32 {
	return s.FPS
}

func (s *AddCustomLiveStreamTranscodeRequest) GetFpsWithSource() *string {
	return s.FpsWithSource
}

func (s *AddCustomLiveStreamTranscodeRequest) GetGop() *string {
	return s.Gop
}

func (s *AddCustomLiveStreamTranscodeRequest) GetHeight() *int32 {
	return s.Height
}

func (s *AddCustomLiveStreamTranscodeRequest) GetKmsKeyExpireInterval() *string {
	return s.KmsKeyExpireInterval
}

func (s *AddCustomLiveStreamTranscodeRequest) GetKmsKeyID() *string {
	return s.KmsKeyID
}

func (s *AddCustomLiveStreamTranscodeRequest) GetKmsUID() *string {
	return s.KmsUID
}

func (s *AddCustomLiveStreamTranscodeRequest) GetLazy() *string {
	return s.Lazy
}

func (s *AddCustomLiveStreamTranscodeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddCustomLiveStreamTranscodeRequest) GetProfile() *int32 {
	return s.Profile
}

func (s *AddCustomLiveStreamTranscodeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddCustomLiveStreamTranscodeRequest) GetResWithSource() *string {
	return s.ResWithSource
}

func (s *AddCustomLiveStreamTranscodeRequest) GetTemplate() *string {
	return s.Template
}

func (s *AddCustomLiveStreamTranscodeRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *AddCustomLiveStreamTranscodeRequest) GetVideoBitrate() *int32 {
	return s.VideoBitrate
}

func (s *AddCustomLiveStreamTranscodeRequest) GetWidth() *int32 {
	return s.Width
}

func (s *AddCustomLiveStreamTranscodeRequest) SetApp(v string) *AddCustomLiveStreamTranscodeRequest {
	s.App = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetAudioBitrate(v int32) *AddCustomLiveStreamTranscodeRequest {
	s.AudioBitrate = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetAudioChannelNum(v int32) *AddCustomLiveStreamTranscodeRequest {
	s.AudioChannelNum = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetAudioCodec(v string) *AddCustomLiveStreamTranscodeRequest {
	s.AudioCodec = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetAudioProfile(v string) *AddCustomLiveStreamTranscodeRequest {
	s.AudioProfile = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetAudioRate(v int32) *AddCustomLiveStreamTranscodeRequest {
	s.AudioRate = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetBitrateWithSource(v string) *AddCustomLiveStreamTranscodeRequest {
	s.BitrateWithSource = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetDeInterlaced(v bool) *AddCustomLiveStreamTranscodeRequest {
	s.DeInterlaced = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetDomain(v string) *AddCustomLiveStreamTranscodeRequest {
	s.Domain = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetEncryptParameters(v string) *AddCustomLiveStreamTranscodeRequest {
	s.EncryptParameters = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetExtWithSource(v string) *AddCustomLiveStreamTranscodeRequest {
	s.ExtWithSource = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetFPS(v int32) *AddCustomLiveStreamTranscodeRequest {
	s.FPS = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetFpsWithSource(v string) *AddCustomLiveStreamTranscodeRequest {
	s.FpsWithSource = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetGop(v string) *AddCustomLiveStreamTranscodeRequest {
	s.Gop = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetHeight(v int32) *AddCustomLiveStreamTranscodeRequest {
	s.Height = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetKmsKeyExpireInterval(v string) *AddCustomLiveStreamTranscodeRequest {
	s.KmsKeyExpireInterval = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetKmsKeyID(v string) *AddCustomLiveStreamTranscodeRequest {
	s.KmsKeyID = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetKmsUID(v string) *AddCustomLiveStreamTranscodeRequest {
	s.KmsUID = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetLazy(v string) *AddCustomLiveStreamTranscodeRequest {
	s.Lazy = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetOwnerId(v int64) *AddCustomLiveStreamTranscodeRequest {
	s.OwnerId = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetProfile(v int32) *AddCustomLiveStreamTranscodeRequest {
	s.Profile = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetRegionId(v string) *AddCustomLiveStreamTranscodeRequest {
	s.RegionId = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetResWithSource(v string) *AddCustomLiveStreamTranscodeRequest {
	s.ResWithSource = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetTemplate(v string) *AddCustomLiveStreamTranscodeRequest {
	s.Template = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetTemplateType(v string) *AddCustomLiveStreamTranscodeRequest {
	s.TemplateType = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetVideoBitrate(v int32) *AddCustomLiveStreamTranscodeRequest {
	s.VideoBitrate = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetWidth(v int32) *AddCustomLiveStreamTranscodeRequest {
	s.Width = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) Validate() error {
	return dara.Validate(s)
}
