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
	// The name of the application to which the live stream belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// The bitrate of the output audio. Unit: Kbit/s. Valid values: **1 to 1000**.
	//
	// example:
	//
	// 512
	AudioBitrate *int32 `json:"AudioBitrate,omitempty" xml:"AudioBitrate,omitempty"`
	// The number of sound channels. Valid values:
	//
	// 	- **1**: mono.
	//
	// 	- **2**: binaural.
	//
	// example:
	//
	// 2
	AudioChannelNum *int32 `json:"AudioChannelNum,omitempty" xml:"AudioChannelNum,omitempty"`
	// The audio encoding format. Valid values:
	//
	// 	- **AAC**
	//
	// 	- **MP3**
	//
	// example:
	//
	// AAC
	AudioCodec *string `json:"AudioCodec,omitempty" xml:"AudioCodec,omitempty"`
	// The audio encoding profile. Valid values:
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
	// The audio sampling rate. Valid values: **22050 to 96000**.
	//
	//
	// 	Notice: If you set AudioProfile to **aac_ld**, the audio sampling rate cannot exceed 44100.
	//
	// example:
	//
	// 96000
	AudioRate *int32 `json:"AudioRate,omitempty" xml:"AudioRate,omitempty"`
	// The source-based bitrate settings. This parameter takes precedence over other bitrate settings. The following fields must be included:
	//
	// 	- **UpLimit**: the maximum bitrate limit. Valid values: an integer from 128 to 10000. The value must be greater than the minimum bitrate.
	//
	// 	- **LowerLimit int*	- : the minimum bitrate rate. Valid values: an integer from 128 to 10000. The value must be smaller than the maximum bitrate.
	//
	// 	- **Factor**: The ratio of the output bitrate to the source bitrate. Valid values: 0.1 to 1. The value is accurate to one decimal place. A value of 1 indicates that the output video has the same bitrate as the source video.
	//
	// example:
	//
	// {"UpLimit":2500,"LowerLimit":800,"Factor":1}
	BitrateWithSource *string `json:"BitrateWithSource,omitempty" xml:"BitrateWithSource,omitempty"`
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Encryption configuration. In JSON format, the fields are explained as follows:
	//
	// - **EncryptType**: Encryption type. Fixed value is aliyun.
	//
	// - **KmsKeyID**: User KMS master key ID.
	//
	// - **KmsKeyExpireInterval**: Key rotation period. Range: 60~3600, unit: seconds.
	//
	// > If the EncryptParameters is configured, the KmsKeyID, KmsUID, and KmsKeyExpireInterval parameters cannot be empty
	//
	// example:
	//
	// {"EncryptType": "aliyun", "KmsKeyID":"afce5722-81d2-43c3-9930-7601da11****","KmsKeyExpireInterval":"3600"}
	EncryptParameters *string `json:"EncryptParameters,omitempty" xml:"EncryptParameters,omitempty"`
	// Other source-based settings, including the following fields:
	//
	// 	- **KeyFrameOpen**: specifies whether to use the key frames of the source video. Valid values: yes or no.
	//
	// 	- **Copyts**: specifies whether to use the presentation time stamp (PTS) of the source video. Valid values: yes or no.
	//
	// 	- **SeiMode**: specifies whether to pass through supplemental enhancement information (SEI) messages. Valid values: 0, 1, and 2, where 0 specifies that no SEI messages are passed through, 1 specifies that part of SEI messages are passed through, and 2 specifies that all SEI messages are passed through.
	//
	// example:
	//
	// {"KeyFrameOpen":"yes","Copyts":"yes","SeiMode":1}
	ExtWithSource *string `json:"ExtWithSource,omitempty" xml:"ExtWithSource,omitempty"`
	// The frame rate of the output video. Unit: frames per second (FPS). Valid values: **1 to 60**.
	//
	// example:
	//
	// 30
	FPS *int32 `json:"FPS,omitempty" xml:"FPS,omitempty"`
	// The source-based frame rate settings. This parameter takes precedence over other frame rate settings. The following fields must be included:
	//
	// 	- **UpLimit**: the maximum frame rate. Valid values: an integer from 1 to 60. The value must be greater than the minimum frame rate.
	//
	// 	- **LowerLimit**: the minimum frame rate. Valid values: an integer from 1 to 60. The value must be smaller than the maximum frame rate.
	//
	// example:
	//
	// {"UpLimit":60,"LowerLimit":1}
	FpsWithSource *string `json:"FpsWithSource,omitempty" xml:"FpsWithSource,omitempty"`
	// The Group of Picture (GOP) size of the video. Unit: frames or seconds.
	//
	// 	- Unit: frames. Valid values: **1 to 3000**.
	//
	// 	- Unit: seconds. Valid value: **1 to 20**.
	//
	// example:
	//
	// 1
	Gop *string `json:"Gop,omitempty" xml:"Gop,omitempty"`
	// The height of the output video. Unit: pixel. Valid values:
	//
	// The value must comply with all the following rules:
	//
	// 	- **Height ≥ 100**: The height of the video is greater than or equal to 100 pixels.
	//
	// 	- **max(Height,Width) ≤ 2560**: The width or height of the video, whichever is greater, cannot exceed 2,560 pixels.
	//
	// 	- **min(Height,Width) ≤ 1440**: The width or height of the video, whichever is smaller, cannot exceed 1,440 pixels.
	//
	// > The resolution of the output video that is transcoded by using the H.265 Narrowband HD™ transcoding template cannot exceed 1280 × 720 pixels.
	//
	// example:
	//
	// 720
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The rotation period of the CMK. Valid values: 60 to 3600. Unit: seconds.
	//
	// example:
	//
	// 3600
	KmsKeyExpireInterval *string `json:"KmsKeyExpireInterval,omitempty" xml:"KmsKeyExpireInterval,omitempty"`
	// The ID of the customer master key (CMK) that you created in Key Management Service (KMS).
	//
	// example:
	//
	// afce5722-81d2-43c3-9930-7601da11****
	KmsKeyID *string `json:"KmsKeyID,omitempty" xml:"KmsKeyID,omitempty"`
	// The ID of your KMS account.
	//
	// example:
	//
	// 25346073170691****
	KmsUID *string `json:"KmsUID,omitempty" xml:"KmsUID,omitempty"`
	// Specifies whether to use the load-on-demand mechanism for transcoding. Valid values: yes and no. Default value: **yes**.
	//
	// example:
	//
	// yes
	Lazy    *string `json:"Lazy,omitempty" xml:"Lazy,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// {"Type":"short","Value":"1080"}
	ResWithSource *string `json:"ResWithSource,omitempty" xml:"ResWithSource,omitempty"`
	// The name of the custom transcoding template.
	//
	// > The name can contain digits, letters, and hyphens (-), and must start with a letter or digit. The name must be different from the names of any default transcoding templates.
	//
	// This parameter is required.
	//
	// example:
	//
	// LiveCusTranscode****
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The type of the custom transcoding template. Valid values:
	//
	// 	- **h264**: custom H.264 standard transcoding.
	//
	// 	- **h264-nbhd**: custom H.264 Narrowband HD™ transcoding.
	//
	// 	- **h265**: custom H.265 standard transcoding.
	//
	// 	- **h265-nbhd**: custom H.265 Narrowband HD™ transcoding.
	//
	// 	- **audio**: audio-only transcoding.
	//
	// > If you set **TemplateType*	- to **h264**, **h264-nbhd**, **h265**, or **h265-nbhd**, the **Height**, **Width**, **FPS**, and **VideoBitrate*	- parameters are required.
	//
	// This parameter is required.
	//
	// example:
	//
	// h264
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// The bitrate of the output video. Unit: Kbit/s. Valid values: **1 to 6000**.
	//
	// > The bitrate of the output video may not be the same as the value that you specify, but is as close to the value as possible, especially when the value is excessively large or small.
	//
	// example:
	//
	// 720
	VideoBitrate *int32 `json:"VideoBitrate,omitempty" xml:"VideoBitrate,omitempty"`
	// The width of the output video. Unit: pixel. Valid values:
	//
	// The value must comply with all the following rules:
	//
	// 	- **Width ≥ 100**: The width of the video is greater than or equal to 100 pixels.
	//
	// 	- **max(Height,Width) ≤ 2560**: The width or height of the video, whichever is greater, cannot exceed 2,560 pixels.
	//
	// 	- **min(Height,Width) ≤ 1440**: The width or height of the video, whichever is smaller, cannot exceed 1,440 pixels.
	//
	// > The resolution of the output video that is transcoded by using the H.265 Narrowband HD™ transcoding template cannot exceed 1280 × 720 pixels.
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
