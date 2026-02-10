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
	if s.DomainTranscodeList != nil {
		if err := s.DomainTranscodeList.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.DomainTranscodeInfo != nil {
		for _, item := range s.DomainTranscodeInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfo struct {
	CustomTranscodeParameters *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters `json:"CustomTranscodeParameters,omitempty" xml:"CustomTranscodeParameters,omitempty" type:"Struct"`
	EncryptParameters         *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoEncryptParameters         `json:"EncryptParameters,omitempty" xml:"EncryptParameters,omitempty" type:"Struct"`
	IsLazy                    *bool                                                                                                       `json:"IsLazy,omitempty" xml:"IsLazy,omitempty"`
	TranscodeApp              *string                                                                                                     `json:"TranscodeApp,omitempty" xml:"TranscodeApp,omitempty"`
	TranscodeName             *string                                                                                                     `json:"TranscodeName,omitempty" xml:"TranscodeName,omitempty"`
	TranscodeTemplate         *string                                                                                                     `json:"TranscodeTemplate,omitempty" xml:"TranscodeTemplate,omitempty"`
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
	if s.CustomTranscodeParameters != nil {
		if err := s.CustomTranscodeParameters.Validate(); err != nil {
			return err
		}
	}
	if s.EncryptParameters != nil {
		if err := s.EncryptParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters struct {
	AudioBitrate      *int32                 `json:"AudioBitrate,omitempty" xml:"AudioBitrate,omitempty"`
	AudioChannelNum   *int32                 `json:"AudioChannelNum,omitempty" xml:"AudioChannelNum,omitempty"`
	AudioCodec        *string                `json:"AudioCodec,omitempty" xml:"AudioCodec,omitempty"`
	AudioProfile      *string                `json:"AudioProfile,omitempty" xml:"AudioProfile,omitempty"`
	AudioRate         *int32                 `json:"AudioRate,omitempty" xml:"AudioRate,omitempty"`
	Bframes           *string                `json:"Bframes,omitempty" xml:"Bframes,omitempty"`
	BitrateWithSource map[string]interface{} `json:"BitrateWithSource,omitempty" xml:"BitrateWithSource,omitempty"`
	DeInterlaced      *bool                  `json:"DeInterlaced,omitempty" xml:"DeInterlaced,omitempty"`
	ExtWithSource     map[string]interface{} `json:"ExtWithSource,omitempty" xml:"ExtWithSource,omitempty"`
	FPS               *int32                 `json:"FPS,omitempty" xml:"FPS,omitempty"`
	FpsWithSource     map[string]interface{} `json:"FpsWithSource,omitempty" xml:"FpsWithSource,omitempty"`
	Gop               *string                `json:"Gop,omitempty" xml:"Gop,omitempty"`
	Height            *int32                 `json:"Height,omitempty" xml:"Height,omitempty"`
	ResWithSource     map[string]interface{} `json:"ResWithSource,omitempty" xml:"ResWithSource,omitempty"`
	RtsFlag           *string                `json:"RtsFlag,omitempty" xml:"RtsFlag,omitempty"`
	TemplateType      *string                `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	VideoBitrate      *int32                 `json:"VideoBitrate,omitempty" xml:"VideoBitrate,omitempty"`
	VideoProfile      *string                `json:"VideoProfile,omitempty" xml:"VideoProfile,omitempty"`
	Width             *int32                 `json:"Width,omitempty" xml:"Width,omitempty"`
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

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) GetDeInterlaced() *bool {
	return s.DeInterlaced
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

func (s *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) SetDeInterlaced(v bool) *DescribeLiveStreamTranscodeInfoResponseBodyDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters {
	s.DeInterlaced = &v
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
	EncryptType          *string `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	KmsKeyExpireInterval *string `json:"KmsKeyExpireInterval,omitempty" xml:"KmsKeyExpireInterval,omitempty"`
	KmsKeyID             *string `json:"KmsKeyID,omitempty" xml:"KmsKeyID,omitempty"`
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
