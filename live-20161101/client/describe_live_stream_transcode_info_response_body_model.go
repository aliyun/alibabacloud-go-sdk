// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamTranscodeInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainTranscodeList(v *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeList) *DescribeLiveStreamTranscodeInfoResponseBody
	GetDomainTranscodeList() *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeList
	SetRequestId(v string) *DescribeLiveStreamTranscodeInfoResponseBody
	GetRequestId() *string
}

type DescribeLiveStreamTranscodeInfoResponseBody struct {
	// The transcoding configurations.
	DomainTranscodeList *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeList `json:"DomainTranscodeList,omitempty" xml:"DomainTranscodeList,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 62136AE6-7793-45ED-B14A-60D19A9486D3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveStreamTranscodeInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamTranscodeInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamTranscodeInfoResponseBody) GetDomainTranscodeList() *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeList {
	return s.DomainTranscodeList
}

func (s *DescribeLiveStreamTranscodeInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveStreamTranscodeInfoResponseBody) SetDomainTranscodeList(v *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeList) *DescribeLiveStreamTranscodeInfoResponseBody {
	s.DomainTranscodeList = v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseBody) SetRequestId(v string) *DescribeLiveStreamTranscodeInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeList struct {
	DomainTranscodeInfo []*DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfo `json:"DomainTranscodeInfo,omitempty" xml:"DomainTranscodeInfo,omitempty" type:"Repeated"`
}

func (s DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeList) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeList) GetDomainTranscodeInfo() []*DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfo {
	return s.DomainTranscodeInfo
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeList) SetDomainTranscodeInfo(v []*DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfo) *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeList {
	s.DomainTranscodeInfo = v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeList) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfo struct {
	// The custom transcoding configuration.
	CustomTranscodeParameters *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters `json:"CustomTranscodeParameters,omitempty" xml:"CustomTranscodeParameters,omitempty" type:"Struct"`
	// The encryption settings.
	EncryptParameters *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoEncryptParameters `json:"EncryptParameters,omitempty" xml:"EncryptParameters,omitempty" type:"Struct"`
	// Indicates whether forcible transcoding is used. Valid values:
	//
	// 	- **true**: Delayed transcoding is used.
	//
	// 	- **false**: Forcible transcoding is used.
	//
	// example:
	//
	// true
	IsLazy *bool `json:"IsLazy,omitempty" xml:"IsLazy,omitempty"`
	// The application name.
	//
	// example:
	//
	// liveApp****
	TranscodeApp *string `json:"TranscodeApp,omitempty" xml:"TranscodeApp,omitempty"`
	// The main streaming domain.
	//
	// example:
	//
	// example.com
	TranscodeName *string `json:"TranscodeName,omitempty" xml:"TranscodeName,omitempty"`
	// The transcoding template ID. Valid values:
	//
	// 	- **Standard transcoding**:
	//
	//     	- **lld**: low definition
	//
	//     	- **lsd**: standard definition
	//
	//     	- **lhd**: high definition
	//
	//     	- **lud**: ultra-high definition
	//
	// 	- **Narrowband HD™ transcoding**:
	//
	//     	- **ld**: low definition
	//
	//     	- **sd**: standard definition
	//
	//     	- **hd**: high definition
	//
	//     	- **ud**: ultra-high definition
	//
	// example:
	//
	// lld
	TranscodeTemplate *string `json:"TranscodeTemplate,omitempty" xml:"TranscodeTemplate,omitempty"`
}

func (s DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfo) GetCustomTranscodeParameters() *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters {
	return s.CustomTranscodeParameters
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfo) GetEncryptParameters() *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoEncryptParameters {
	return s.EncryptParameters
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfo) GetIsLazy() *bool {
	return s.IsLazy
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfo) GetTranscodeApp() *string {
	return s.TranscodeApp
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfo) GetTranscodeName() *string {
	return s.TranscodeName
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfo) GetTranscodeTemplate() *string {
	return s.TranscodeTemplate
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfo) SetCustomTranscodeParameters(v *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfo {
	s.CustomTranscodeParameters = v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfo) SetEncryptParameters(v *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoEncryptParameters) *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfo {
	s.EncryptParameters = v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfo) SetIsLazy(v bool) *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfo {
	s.IsLazy = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfo) SetTranscodeApp(v string) *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfo {
	s.TranscodeApp = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfo) SetTranscodeName(v string) *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfo {
	s.TranscodeName = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfo) SetTranscodeTemplate(v string) *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfo {
	s.TranscodeTemplate = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters struct {
	// The bitrate of the output audio. Unit: Kbit/s. Valid values: **1 to 1000**.
	//
	// example:
	//
	// 64
	AudioBitrate *int32 `json:"AudioBitrate,omitempty" xml:"AudioBitrate,omitempty"`
	// The number of sound channels. Valid values:
	//
	// 	- **1**: mono
	//
	// 	- **2**: binaural
	//
	// example:
	//
	// 2
	AudioChannelNum *int32 `json:"AudioChannelNum,omitempty" xml:"AudioChannelNum,omitempty"`
	// The audio encoding format.
	//
	// example:
	//
	// ACC
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
	// >  If the value of AudioProfile is **aac_ld**, the audio sampling rate cannot exceed 44100.
	//
	// example:
	//
	// 44100
	AudioRate *int32 `json:"AudioRate,omitempty" xml:"AudioRate,omitempty"`
	// Indicates whether B-frame removal is enabled. Fixed value: **0**.
	//
	// example:
	//
	// 0
	Bframes *string `json:"Bframes,omitempty" xml:"Bframes,omitempty"`
	// The source-based bitrate settings.
	//
	// example:
	//
	// {\\"UpLimit\\":2500,\\"LowerLimit\\":800,\\"Factor\\":1}
	BitrateWithSource map[string]interface{} `json:"BitrateWithSource,omitempty" xml:"BitrateWithSource,omitempty"`
	// Other source-based settings.
	//
	// example:
	//
	// {\\"KeyFrameOpen\\":\\"yes\\",\\"Copyts\\":\\"yes\\",\\"SeiMode\\":1}
	ExtWithSource map[string]interface{} `json:"ExtWithSource,omitempty" xml:"ExtWithSource,omitempty"`
	// The frame rate of the output video. Unit: frames per second (FPS).
	//
	// example:
	//
	// 15
	FPS *int32 `json:"FPS,omitempty" xml:"FPS,omitempty"`
	// The source-based frame rate settings.
	//
	// example:
	//
	// {\\"UpLimit\\":60,\\"LowerLimit\\":1}
	FpsWithSource map[string]interface{} `json:"FpsWithSource,omitempty" xml:"FpsWithSource,omitempty"`
	// The group of pictures (GOP) size of the output video. Unit: frames. Valid values: **1 to 3000**.
	//
	// example:
	//
	// 10
	Gop *string `json:"Gop,omitempty" xml:"Gop,omitempty"`
	// The height of the output video.
	//
	// example:
	//
	// 1200
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The source-based resolution settings.
	//
	// example:
	//
	// {\\"Type\\":\\"short\\",\\"Value\\":\\"1080\\"}
	ResWithSource map[string]interface{} `json:"ResWithSource,omitempty" xml:"ResWithSource,omitempty"`
	// The Real-Time Transcoding (RTS) flag. Fixed value: **true**.
	//
	// >  This parameter is returned only if RTS is used for transcoding.
	//
	// example:
	//
	// true
	RtsFlag *string `json:"RtsFlag,omitempty" xml:"RtsFlag,omitempty"`
	// The type of the custom transcoding template. Valid values:
	//
	// 	- **h264**: custom H.264 standard transcoding
	//
	// 	- **h264-nbhd**: custom H.264 Narrowband HD™ transcoding
	//
	// 	- **h265**: custom H.265 standard transcoding
	//
	// 	- **h265-nbhd**: custom H.265 Narrowband HD™ transcoding
	//
	// 	- **audio**: audio-only transcoding
	//
	// example:
	//
	// h264
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// The bitrate of the output video. Unit: Kbit/s.
	//
	// example:
	//
	// 3000
	VideoBitrate *int32 `json:"VideoBitrate,omitempty" xml:"VideoBitrate,omitempty"`
	// The video encoding profile. Valid values:
	//
	// 	- **baseline**: suitable for mobile devices.
	//
	// 	- **main**: suitable for standard-definition devices.
	//
	// 	- **high**: suitable for high-definition devices.
	//
	// example:
	//
	// high
	VideoProfile *string `json:"VideoProfile,omitempty" xml:"VideoProfile,omitempty"`
	// The width of the output video.
	//
	// example:
	//
	// 1000
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) GetAudioBitrate() *int32 {
	return s.AudioBitrate
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) GetAudioChannelNum() *int32 {
	return s.AudioChannelNum
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) GetAudioCodec() *string {
	return s.AudioCodec
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) GetAudioProfile() *string {
	return s.AudioProfile
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) GetAudioRate() *int32 {
	return s.AudioRate
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) GetBframes() *string {
	return s.Bframes
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) GetBitrateWithSource() map[string]interface{} {
	return s.BitrateWithSource
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) GetExtWithSource() map[string]interface{} {
	return s.ExtWithSource
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) GetFPS() *int32 {
	return s.FPS
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) GetFpsWithSource() map[string]interface{} {
	return s.FpsWithSource
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) GetGop() *string {
	return s.Gop
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) GetHeight() *int32 {
	return s.Height
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) GetResWithSource() map[string]interface{} {
	return s.ResWithSource
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) GetRtsFlag() *string {
	return s.RtsFlag
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) GetTemplateType() *string {
	return s.TemplateType
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) GetVideoBitrate() *int32 {
	return s.VideoBitrate
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) GetVideoProfile() *string {
	return s.VideoProfile
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) GetWidth() *int32 {
	return s.Width
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) SetAudioBitrate(v int32) *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters {
	s.AudioBitrate = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) SetAudioChannelNum(v int32) *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters {
	s.AudioChannelNum = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) SetAudioCodec(v string) *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters {
	s.AudioCodec = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) SetAudioProfile(v string) *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters {
	s.AudioProfile = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) SetAudioRate(v int32) *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters {
	s.AudioRate = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) SetBframes(v string) *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters {
	s.Bframes = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) SetBitrateWithSource(v map[string]interface{}) *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters {
	s.BitrateWithSource = v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) SetExtWithSource(v map[string]interface{}) *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters {
	s.ExtWithSource = v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) SetFPS(v int32) *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters {
	s.FPS = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) SetFpsWithSource(v map[string]interface{}) *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters {
	s.FpsWithSource = v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) SetGop(v string) *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters {
	s.Gop = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) SetHeight(v int32) *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters {
	s.Height = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) SetResWithSource(v map[string]interface{}) *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters {
	s.ResWithSource = v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) SetRtsFlag(v string) *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters {
	s.RtsFlag = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) SetTemplateType(v string) *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters {
	s.TemplateType = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) SetVideoBitrate(v int32) *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters {
	s.VideoBitrate = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) SetVideoProfile(v string) *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters {
	s.VideoProfile = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) SetWidth(v int32) *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters {
	s.Width = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoEncryptParameters struct {
	// The type of encryption. Fixed value: **aliyun**.
	//
	// example:
	//
	// aliyun
	EncryptType *string `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	// The rotation period of the CMK. Valid values: **60 to 3600**. Unit: seconds.
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
}

func (s DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoEncryptParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoEncryptParameters) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoEncryptParameters) GetEncryptType() *string {
	return s.EncryptType
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoEncryptParameters) GetKmsKeyExpireInterval() *string {
	return s.KmsKeyExpireInterval
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoEncryptParameters) GetKmsKeyID() *string {
	return s.KmsKeyID
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoEncryptParameters) SetEncryptType(v string) *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoEncryptParameters {
	s.EncryptType = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoEncryptParameters) SetKmsKeyExpireInterval(v string) *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoEncryptParameters {
	s.KmsKeyExpireInterval = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoEncryptParameters) SetKmsKeyID(v string) *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoEncryptParameters {
	s.KmsKeyID = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoEncryptParameters) Validate() error {
	return dara.Validate(s)
}
