// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLiveTranscodeTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetLiveTranscodeTemplateResponseBody
	GetRequestId() *string
	SetTemplateContent(v *GetLiveTranscodeTemplateResponseBodyTemplateContent) *GetLiveTranscodeTemplateResponseBody
	GetTemplateContent() *GetLiveTranscodeTemplateResponseBodyTemplateContent
}

type GetLiveTranscodeTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The content of the template.
	TemplateContent *GetLiveTranscodeTemplateResponseBodyTemplateContent `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty" type:"Struct"`
}

func (s GetLiveTranscodeTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLiveTranscodeTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetLiveTranscodeTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLiveTranscodeTemplateResponseBody) GetTemplateContent() *GetLiveTranscodeTemplateResponseBodyTemplateContent {
	return s.TemplateContent
}

func (s *GetLiveTranscodeTemplateResponseBody) SetRequestId(v string) *GetLiveTranscodeTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLiveTranscodeTemplateResponseBody) SetTemplateContent(v *GetLiveTranscodeTemplateResponseBodyTemplateContent) *GetLiveTranscodeTemplateResponseBody {
	s.TemplateContent = v
	return s
}

func (s *GetLiveTranscodeTemplateResponseBody) Validate() error {
	if s.TemplateContent != nil {
		if err := s.TemplateContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLiveTranscodeTemplateResponseBodyTemplateContent struct {
	// The category of the template. Valid values:
	//
	// 	- system
	//
	// 	- customized
	//
	// example:
	//
	// customized
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The time when the template was created.
	//
	// example:
	//
	// 2022-07-25T06:15:14Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// my-template
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The configuration of the template.
	TemplateConfig *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfig `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty" type:"Struct"`
	// The template ID.
	//
	// example:
	//
	// bcfa57950bc649b2abfb476ecd36ea4f
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The type of the template.
	//
	// example:
	//
	// normal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetLiveTranscodeTemplateResponseBodyTemplateContent) String() string {
	return dara.Prettify(s)
}

func (s GetLiveTranscodeTemplateResponseBodyTemplateContent) GoString() string {
	return s.String()
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContent) GetCategory() *string {
	return s.Category
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContent) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContent) GetName() *string {
	return s.Name
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContent) GetTemplateConfig() *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfig {
	return s.TemplateConfig
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContent) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContent) GetType() *string {
	return s.Type
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContent) SetCategory(v string) *GetLiveTranscodeTemplateResponseBodyTemplateContent {
	s.Category = &v
	return s
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContent) SetCreateTime(v string) *GetLiveTranscodeTemplateResponseBodyTemplateContent {
	s.CreateTime = &v
	return s
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContent) SetName(v string) *GetLiveTranscodeTemplateResponseBodyTemplateContent {
	s.Name = &v
	return s
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContent) SetTemplateConfig(v *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfig) *GetLiveTranscodeTemplateResponseBodyTemplateContent {
	s.TemplateConfig = v
	return s
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContent) SetTemplateId(v string) *GetLiveTranscodeTemplateResponseBodyTemplateContent {
	s.TemplateId = &v
	return s
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContent) SetType(v string) *GetLiveTranscodeTemplateResponseBodyTemplateContent {
	s.Type = &v
	return s
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContent) Validate() error {
	if s.TemplateConfig != nil {
		if err := s.TemplateConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfig struct {
	// The audio parameters.
	AudioParams *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigAudioParams `json:"AudioParams,omitempty" xml:"AudioParams,omitempty" type:"Struct"`
	// The video parameters.
	VideoParams *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigVideoParams `json:"VideoParams,omitempty" xml:"VideoParams,omitempty" type:"Struct"`
}

func (s GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfig) String() string {
	return dara.Prettify(s)
}

func (s GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfig) GoString() string {
	return s.String()
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfig) GetAudioParams() *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigAudioParams {
	return s.AudioParams
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfig) GetVideoParams() *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigVideoParams {
	return s.VideoParams
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfig) SetAudioParams(v *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigAudioParams) *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfig {
	s.AudioParams = v
	return s
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfig) SetVideoParams(v *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigVideoParams) *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfig {
	s.VideoParams = v
	return s
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfig) Validate() error {
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

type GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigAudioParams struct {
	// The bitrate of the output audio.
	//
	// example:
	//
	// 1000
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The number of sound channels.
	//
	// example:
	//
	// 2
	Channels *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	// The audio codec.
	//
	// example:
	//
	// AAC
	Codec *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	// The audio codec profile.
	//
	// example:
	//
	// 1
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The audio sampling rate.
	//
	// example:
	//
	// 44100
	Samplerate *string `json:"Samplerate,omitempty" xml:"Samplerate,omitempty"`
}

func (s GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigAudioParams) String() string {
	return dara.Prettify(s)
}

func (s GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigAudioParams) GoString() string {
	return s.String()
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigAudioParams) GetBitrate() *string {
	return s.Bitrate
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigAudioParams) GetChannels() *string {
	return s.Channels
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigAudioParams) GetCodec() *string {
	return s.Codec
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigAudioParams) GetProfile() *string {
	return s.Profile
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigAudioParams) GetSamplerate() *string {
	return s.Samplerate
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigAudioParams) SetBitrate(v string) *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigAudioParams {
	s.Bitrate = &v
	return s
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigAudioParams) SetChannels(v string) *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigAudioParams {
	s.Channels = &v
	return s
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigAudioParams) SetCodec(v string) *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigAudioParams {
	s.Codec = &v
	return s
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigAudioParams) SetProfile(v string) *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigAudioParams {
	s.Profile = &v
	return s
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigAudioParams) SetSamplerate(v string) *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigAudioParams {
	s.Samplerate = &v
	return s
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigAudioParams) Validate() error {
	return dara.Validate(s)
}

type GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigVideoParams struct {
	// The bitrate of the output video.
	//
	// example:
	//
	// 2500
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The encoding type.
	//
	// example:
	//
	// H.264
	Codec *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	// The frame rate of the output video.
	//
	// example:
	//
	// 30
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The group of pictures (GOP) of the output video.
	//
	// example:
	//
	// 1000
	Gop *string `json:"Gop,omitempty" xml:"Gop,omitempty"`
	// The height of the output video.
	//
	// example:
	//
	// 720
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The encoding profile.
	//
	// example:
	//
	// 2
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The width of the output video.
	//
	// example:
	//
	// 1280
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigVideoParams) String() string {
	return dara.Prettify(s)
}

func (s GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigVideoParams) GoString() string {
	return s.String()
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigVideoParams) GetBitrate() *string {
	return s.Bitrate
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigVideoParams) GetCodec() *string {
	return s.Codec
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigVideoParams) GetFps() *string {
	return s.Fps
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigVideoParams) GetGop() *string {
	return s.Gop
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigVideoParams) GetHeight() *string {
	return s.Height
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigVideoParams) GetProfile() *string {
	return s.Profile
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigVideoParams) GetWidth() *string {
	return s.Width
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigVideoParams) SetBitrate(v string) *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigVideoParams {
	s.Bitrate = &v
	return s
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigVideoParams) SetCodec(v string) *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigVideoParams {
	s.Codec = &v
	return s
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigVideoParams) SetFps(v string) *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigVideoParams {
	s.Fps = &v
	return s
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigVideoParams) SetGop(v string) *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigVideoParams {
	s.Gop = &v
	return s
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigVideoParams) SetHeight(v string) *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigVideoParams {
	s.Height = &v
	return s
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigVideoParams) SetProfile(v string) *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigVideoParams {
	s.Profile = &v
	return s
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigVideoParams) SetWidth(v string) *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigVideoParams {
	s.Width = &v
	return s
}

func (s *GetLiveTranscodeTemplateResponseBodyTemplateContentTemplateConfigVideoParams) Validate() error {
	return dara.Validate(s)
}
