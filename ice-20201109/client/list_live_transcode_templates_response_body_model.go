// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveTranscodeTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListLiveTranscodeTemplatesResponseBody
	GetRequestId() *string
	SetTemplateContentList(v []*ListLiveTranscodeTemplatesResponseBodyTemplateContentList) *ListLiveTranscodeTemplatesResponseBody
	GetTemplateContentList() []*ListLiveTranscodeTemplatesResponseBodyTemplateContentList
	SetTotalCount(v int32) *ListLiveTranscodeTemplatesResponseBody
	GetTotalCount() *int32
}

type ListLiveTranscodeTemplatesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of transcoding templates.
	TemplateContentList []*ListLiveTranscodeTemplatesResponseBodyTemplateContentList `json:"TemplateContentList,omitempty" xml:"TemplateContentList,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLiveTranscodeTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLiveTranscodeTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListLiveTranscodeTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLiveTranscodeTemplatesResponseBody) GetTemplateContentList() []*ListLiveTranscodeTemplatesResponseBodyTemplateContentList {
	return s.TemplateContentList
}

func (s *ListLiveTranscodeTemplatesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListLiveTranscodeTemplatesResponseBody) SetRequestId(v string) *ListLiveTranscodeTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLiveTranscodeTemplatesResponseBody) SetTemplateContentList(v []*ListLiveTranscodeTemplatesResponseBodyTemplateContentList) *ListLiveTranscodeTemplatesResponseBody {
	s.TemplateContentList = v
	return s
}

func (s *ListLiveTranscodeTemplatesResponseBody) SetTotalCount(v int32) *ListLiveTranscodeTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListLiveTranscodeTemplatesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListLiveTranscodeTemplatesResponseBodyTemplateContentList struct {
	// The category of the template. Valid values:
	//
	// example:
	//
	// system
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The time when the job was created.
	//
	// example:
	//
	// 2022-07-20T03:26:36Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The template name.
	//
	// example:
	//
	// my_template
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The configuration of the template.
	TemplateConfig *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfig `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty" type:"Struct"`
	// The template ID.
	//
	// example:
	//
	// 9b1571b513cb44f7a1ba6ae561ff46f7
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The type of the template.
	//
	// example:
	//
	// normal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListLiveTranscodeTemplatesResponseBodyTemplateContentList) String() string {
	return dara.Prettify(s)
}

func (s ListLiveTranscodeTemplatesResponseBodyTemplateContentList) GoString() string {
	return s.String()
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentList) GetCategory() *string {
	return s.Category
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentList) GetName() *string {
	return s.Name
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentList) GetTemplateConfig() *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfig {
	return s.TemplateConfig
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentList) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentList) GetType() *string {
	return s.Type
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentList) SetCategory(v string) *ListLiveTranscodeTemplatesResponseBodyTemplateContentList {
	s.Category = &v
	return s
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentList) SetCreateTime(v string) *ListLiveTranscodeTemplatesResponseBodyTemplateContentList {
	s.CreateTime = &v
	return s
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentList) SetName(v string) *ListLiveTranscodeTemplatesResponseBodyTemplateContentList {
	s.Name = &v
	return s
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentList) SetTemplateConfig(v *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfig) *ListLiveTranscodeTemplatesResponseBodyTemplateContentList {
	s.TemplateConfig = v
	return s
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentList) SetTemplateId(v string) *ListLiveTranscodeTemplatesResponseBodyTemplateContentList {
	s.TemplateId = &v
	return s
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentList) SetType(v string) *ListLiveTranscodeTemplatesResponseBodyTemplateContentList {
	s.Type = &v
	return s
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentList) Validate() error {
	return dara.Validate(s)
}

type ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfig struct {
	// The audio parameters.
	AudioParams *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigAudioParams `json:"AudioParams,omitempty" xml:"AudioParams,omitempty" type:"Struct"`
	// The video parameters.
	VideoParams *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigVideoParams `json:"VideoParams,omitempty" xml:"VideoParams,omitempty" type:"Struct"`
}

func (s ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfig) String() string {
	return dara.Prettify(s)
}

func (s ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfig) GoString() string {
	return s.String()
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfig) GetAudioParams() *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigAudioParams {
	return s.AudioParams
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfig) GetVideoParams() *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigVideoParams {
	return s.VideoParams
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfig) SetAudioParams(v *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigAudioParams) *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfig {
	s.AudioParams = v
	return s
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfig) SetVideoParams(v *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigVideoParams) *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfig {
	s.VideoParams = v
	return s
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfig) Validate() error {
	return dara.Validate(s)
}

type ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigAudioParams struct {
	// The audio bitrate.
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
	// The encoding profile.
	//
	// example:
	//
	// aac_low
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The audio sampling rate.
	//
	// example:
	//
	// 44100
	Samplerate *string `json:"Samplerate,omitempty" xml:"Samplerate,omitempty"`
}

func (s ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigAudioParams) String() string {
	return dara.Prettify(s)
}

func (s ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigAudioParams) GoString() string {
	return s.String()
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigAudioParams) GetBitrate() *string {
	return s.Bitrate
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigAudioParams) GetChannels() *string {
	return s.Channels
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigAudioParams) GetCodec() *string {
	return s.Codec
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigAudioParams) GetProfile() *string {
	return s.Profile
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigAudioParams) GetSamplerate() *string {
	return s.Samplerate
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigAudioParams) SetBitrate(v string) *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigAudioParams {
	s.Bitrate = &v
	return s
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigAudioParams) SetChannels(v string) *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigAudioParams {
	s.Channels = &v
	return s
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigAudioParams) SetCodec(v string) *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigAudioParams {
	s.Codec = &v
	return s
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigAudioParams) SetProfile(v string) *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigAudioParams {
	s.Profile = &v
	return s
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigAudioParams) SetSamplerate(v string) *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigAudioParams {
	s.Samplerate = &v
	return s
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigAudioParams) Validate() error {
	return dara.Validate(s)
}

type ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigVideoParams struct {
	// The video bitrate.
	//
	// example:
	//
	// 2500
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The encoding format.
	//
	// example:
	//
	// 264
	Codec *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	// The video frame rate.
	//
	// example:
	//
	// 30
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The group of pictures (GOP) of the output video. Unit: frame.
	//
	// example:
	//
	// 1000
	Gop *string `json:"Gop,omitempty" xml:"Gop,omitempty"`
	// The vertical resolution of the video.
	//
	// example:
	//
	// 1280
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The encoding profile.
	//
	// example:
	//
	// 3
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The horizontal resolution of the video.
	//
	// example:
	//
	// 720
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigVideoParams) String() string {
	return dara.Prettify(s)
}

func (s ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigVideoParams) GoString() string {
	return s.String()
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigVideoParams) GetBitrate() *string {
	return s.Bitrate
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigVideoParams) GetCodec() *string {
	return s.Codec
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigVideoParams) GetFps() *string {
	return s.Fps
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigVideoParams) GetGop() *string {
	return s.Gop
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigVideoParams) GetHeight() *string {
	return s.Height
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigVideoParams) GetProfile() *string {
	return s.Profile
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigVideoParams) GetWidth() *string {
	return s.Width
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigVideoParams) SetBitrate(v string) *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigVideoParams {
	s.Bitrate = &v
	return s
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigVideoParams) SetCodec(v string) *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigVideoParams {
	s.Codec = &v
	return s
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigVideoParams) SetFps(v string) *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigVideoParams {
	s.Fps = &v
	return s
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigVideoParams) SetGop(v string) *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigVideoParams {
	s.Gop = &v
	return s
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigVideoParams) SetHeight(v string) *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigVideoParams {
	s.Height = &v
	return s
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigVideoParams) SetProfile(v string) *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigVideoParams {
	s.Profile = &v
	return s
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigVideoParams) SetWidth(v string) *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigVideoParams {
	s.Width = &v
	return s
}

func (s *ListLiveTranscodeTemplatesResponseBodyTemplateContentListTemplateConfigVideoParams) Validate() error {
	return dara.Validate(s)
}
