// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *SearchTemplateResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *SearchTemplateResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *SearchTemplateResponseBody
	GetRequestId() *string
	SetTemplateList(v *SearchTemplateResponseBodyTemplateList) *SearchTemplateResponseBody
	GetTemplateList() *SearchTemplateResponseBodyTemplateList
	SetTotalCount(v int64) *SearchTemplateResponseBody
	GetTotalCount() *int64
}

type SearchTemplateResponseBody struct {
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// BC860F04-778A-472F-AB39-E1BF329C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The transcoding templates.
	TemplateList *SearchTemplateResponseBodyTemplateList `json:"TemplateList,omitempty" xml:"TemplateList,omitempty" type:"Struct"`
	// The total number of search results.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTemplateResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *SearchTemplateResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *SearchTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchTemplateResponseBody) GetTemplateList() *SearchTemplateResponseBodyTemplateList {
	return s.TemplateList
}

func (s *SearchTemplateResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *SearchTemplateResponseBody) SetPageNumber(v int64) *SearchTemplateResponseBody {
	s.PageNumber = &v
	return s
}

func (s *SearchTemplateResponseBody) SetPageSize(v int64) *SearchTemplateResponseBody {
	s.PageSize = &v
	return s
}

func (s *SearchTemplateResponseBody) SetRequestId(v string) *SearchTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchTemplateResponseBody) SetTemplateList(v *SearchTemplateResponseBodyTemplateList) *SearchTemplateResponseBody {
	s.TemplateList = v
	return s
}

func (s *SearchTemplateResponseBody) SetTotalCount(v int64) *SearchTemplateResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SearchTemplateResponseBody) Validate() error {
	if s.TemplateList != nil {
		if err := s.TemplateList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchTemplateResponseBodyTemplateList struct {
	Template []*SearchTemplateResponseBodyTemplateListTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Repeated"`
}

func (s SearchTemplateResponseBodyTemplateList) String() string {
	return dara.Prettify(s)
}

func (s SearchTemplateResponseBodyTemplateList) GoString() string {
	return s.String()
}

func (s *SearchTemplateResponseBodyTemplateList) GetTemplate() []*SearchTemplateResponseBodyTemplateListTemplate {
	return s.Template
}

func (s *SearchTemplateResponseBodyTemplateList) SetTemplate(v []*SearchTemplateResponseBodyTemplateListTemplate) *SearchTemplateResponseBodyTemplateList {
	s.Template = v
	return s
}

func (s *SearchTemplateResponseBodyTemplateList) Validate() error {
	if s.Template != nil {
		for _, item := range s.Template {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchTemplateResponseBodyTemplateListTemplate struct {
	// The audio codec configurations.
	Audio *SearchTemplateResponseBodyTemplateListTemplateAudio `json:"Audio,omitempty" xml:"Audio,omitempty" type:"Struct"`
	// The container format configurations.
	Container *SearchTemplateResponseBodyTemplateListTemplateContainer `json:"Container,omitempty" xml:"Container,omitempty" type:"Struct"`
	// The time when the template was created.
	//
	// example:
	//
	// 2021-03-04T06:44:43Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The transcoding template ID.
	//
	// example:
	//
	// 16f01ad6175e4230ac42bb5182cd****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The transmuxing configurations.
	MuxConfig *SearchTemplateResponseBodyTemplateListTemplateMuxConfig `json:"MuxConfig,omitempty" xml:"MuxConfig,omitempty" type:"Struct"`
	// The name of the template.
	//
	// example:
	//
	// MPS-example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the template. Valid values:
	//
	// 	- **Normal**
	//
	// 	- **Deleted**
	//
	// example:
	//
	// Normal
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The general transcoding configurations.
	TransConfig *SearchTemplateResponseBodyTemplateListTemplateTransConfig `json:"TransConfig,omitempty" xml:"TransConfig,omitempty" type:"Struct"`
	// The video codec configurations.
	Video *SearchTemplateResponseBodyTemplateListTemplateVideo `json:"Video,omitempty" xml:"Video,omitempty" type:"Struct"`
}

func (s SearchTemplateResponseBodyTemplateListTemplate) String() string {
	return dara.Prettify(s)
}

func (s SearchTemplateResponseBodyTemplateListTemplate) GoString() string {
	return s.String()
}

func (s *SearchTemplateResponseBodyTemplateListTemplate) GetAudio() *SearchTemplateResponseBodyTemplateListTemplateAudio {
	return s.Audio
}

func (s *SearchTemplateResponseBodyTemplateListTemplate) GetContainer() *SearchTemplateResponseBodyTemplateListTemplateContainer {
	return s.Container
}

func (s *SearchTemplateResponseBodyTemplateListTemplate) GetCreationTime() *string {
	return s.CreationTime
}

func (s *SearchTemplateResponseBodyTemplateListTemplate) GetId() *string {
	return s.Id
}

func (s *SearchTemplateResponseBodyTemplateListTemplate) GetMuxConfig() *SearchTemplateResponseBodyTemplateListTemplateMuxConfig {
	return s.MuxConfig
}

func (s *SearchTemplateResponseBodyTemplateListTemplate) GetName() *string {
	return s.Name
}

func (s *SearchTemplateResponseBodyTemplateListTemplate) GetState() *string {
	return s.State
}

func (s *SearchTemplateResponseBodyTemplateListTemplate) GetTransConfig() *SearchTemplateResponseBodyTemplateListTemplateTransConfig {
	return s.TransConfig
}

func (s *SearchTemplateResponseBodyTemplateListTemplate) GetVideo() *SearchTemplateResponseBodyTemplateListTemplateVideo {
	return s.Video
}

func (s *SearchTemplateResponseBodyTemplateListTemplate) SetAudio(v *SearchTemplateResponseBodyTemplateListTemplateAudio) *SearchTemplateResponseBodyTemplateListTemplate {
	s.Audio = v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplate) SetContainer(v *SearchTemplateResponseBodyTemplateListTemplateContainer) *SearchTemplateResponseBodyTemplateListTemplate {
	s.Container = v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplate) SetCreationTime(v string) *SearchTemplateResponseBodyTemplateListTemplate {
	s.CreationTime = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplate) SetId(v string) *SearchTemplateResponseBodyTemplateListTemplate {
	s.Id = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplate) SetMuxConfig(v *SearchTemplateResponseBodyTemplateListTemplateMuxConfig) *SearchTemplateResponseBodyTemplateListTemplate {
	s.MuxConfig = v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplate) SetName(v string) *SearchTemplateResponseBodyTemplateListTemplate {
	s.Name = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplate) SetState(v string) *SearchTemplateResponseBodyTemplateListTemplate {
	s.State = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplate) SetTransConfig(v *SearchTemplateResponseBodyTemplateListTemplateTransConfig) *SearchTemplateResponseBodyTemplateListTemplate {
	s.TransConfig = v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplate) SetVideo(v *SearchTemplateResponseBodyTemplateListTemplateVideo) *SearchTemplateResponseBodyTemplateListTemplate {
	s.Video = v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplate) Validate() error {
	if s.Audio != nil {
		if err := s.Audio.Validate(); err != nil {
			return err
		}
	}
	if s.Container != nil {
		if err := s.Container.Validate(); err != nil {
			return err
		}
	}
	if s.MuxConfig != nil {
		if err := s.MuxConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TransConfig != nil {
		if err := s.TransConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Video != nil {
		if err := s.Video.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchTemplateResponseBodyTemplateListTemplateAudio struct {
	// The audio bitrate of the output file.
	//
	// 	- Unit: Kbit/s.
	//
	// 	- Default value: **128**.
	//
	// example:
	//
	// 500
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The number of sound channels. Default value: **2**.
	//
	// example:
	//
	// 2
	Channels *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	// The audio codec format. Default value: **aac**. Valid values:
	//
	// 	- **aac**
	//
	// 	- **mp3**
	//
	// 	- **vorbis**
	//
	// 	- **flac**
	//
	// example:
	//
	// aac
	Codec *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	// The codec profile of the audio. Valid values when the value of Codec is aac:
	//
	// 	- **aac_low**
	//
	// 	- **aac_he**
	//
	// 	- **aac_he_v2**
	//
	// 	- **aac_ld**
	//
	// 	- **aac_eld**
	//
	// example:
	//
	// aac_low
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The strength of the independent denoising algorithm. Valid values: **[1,9]**.
	//
	// example:
	//
	// 1
	Qscale *string `json:"Qscale,omitempty" xml:"Qscale,omitempty"`
	// Indicates whether the audio stream is deleted. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// 	- Default value: **false**.
	//
	// example:
	//
	// false
	Remove *string `json:"Remove,omitempty" xml:"Remove,omitempty"`
	// The sampling rate.
	//
	// 	- Unit: Hz
	//
	// 	- Default value: **44100**.
	//
	// example:
	//
	// 44100
	Samplerate *string `json:"Samplerate,omitempty" xml:"Samplerate,omitempty"`
}

func (s SearchTemplateResponseBodyTemplateListTemplateAudio) String() string {
	return dara.Prettify(s)
}

func (s SearchTemplateResponseBodyTemplateListTemplateAudio) GoString() string {
	return s.String()
}

func (s *SearchTemplateResponseBodyTemplateListTemplateAudio) GetBitrate() *string {
	return s.Bitrate
}

func (s *SearchTemplateResponseBodyTemplateListTemplateAudio) GetChannels() *string {
	return s.Channels
}

func (s *SearchTemplateResponseBodyTemplateListTemplateAudio) GetCodec() *string {
	return s.Codec
}

func (s *SearchTemplateResponseBodyTemplateListTemplateAudio) GetProfile() *string {
	return s.Profile
}

func (s *SearchTemplateResponseBodyTemplateListTemplateAudio) GetQscale() *string {
	return s.Qscale
}

func (s *SearchTemplateResponseBodyTemplateListTemplateAudio) GetRemove() *string {
	return s.Remove
}

func (s *SearchTemplateResponseBodyTemplateListTemplateAudio) GetSamplerate() *string {
	return s.Samplerate
}

func (s *SearchTemplateResponseBodyTemplateListTemplateAudio) SetBitrate(v string) *SearchTemplateResponseBodyTemplateListTemplateAudio {
	s.Bitrate = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateAudio) SetChannels(v string) *SearchTemplateResponseBodyTemplateListTemplateAudio {
	s.Channels = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateAudio) SetCodec(v string) *SearchTemplateResponseBodyTemplateListTemplateAudio {
	s.Codec = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateAudio) SetProfile(v string) *SearchTemplateResponseBodyTemplateListTemplateAudio {
	s.Profile = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateAudio) SetQscale(v string) *SearchTemplateResponseBodyTemplateListTemplateAudio {
	s.Qscale = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateAudio) SetRemove(v string) *SearchTemplateResponseBodyTemplateListTemplateAudio {
	s.Remove = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateAudio) SetSamplerate(v string) *SearchTemplateResponseBodyTemplateListTemplateAudio {
	s.Samplerate = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateAudio) Validate() error {
	return dara.Validate(s)
}

type SearchTemplateResponseBodyTemplateListTemplateContainer struct {
	// The format of the container. Valid values:
	//
	// 	- **flv**
	//
	// 	- **mp4**
	//
	// 	- **ts**
	//
	// 	- **m3u8**
	//
	// 	- **gif**
	//
	// 	- **mp3**
	//
	// 	- **ogg**
	//
	// 	- **flac**
	//
	// example:
	//
	// mp4
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
}

func (s SearchTemplateResponseBodyTemplateListTemplateContainer) String() string {
	return dara.Prettify(s)
}

func (s SearchTemplateResponseBodyTemplateListTemplateContainer) GoString() string {
	return s.String()
}

func (s *SearchTemplateResponseBodyTemplateListTemplateContainer) GetFormat() *string {
	return s.Format
}

func (s *SearchTemplateResponseBodyTemplateListTemplateContainer) SetFormat(v string) *SearchTemplateResponseBodyTemplateListTemplateContainer {
	s.Format = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateContainer) Validate() error {
	return dara.Validate(s)
}

type SearchTemplateResponseBodyTemplateListTemplateMuxConfig struct {
	// The transmuxing configurations for GIF.
	Gif *SearchTemplateResponseBodyTemplateListTemplateMuxConfigGif `json:"Gif,omitempty" xml:"Gif,omitempty" type:"Struct"`
	// The segment configurations.
	Segment *SearchTemplateResponseBodyTemplateListTemplateMuxConfigSegment `json:"Segment,omitempty" xml:"Segment,omitempty" type:"Struct"`
}

func (s SearchTemplateResponseBodyTemplateListTemplateMuxConfig) String() string {
	return dara.Prettify(s)
}

func (s SearchTemplateResponseBodyTemplateListTemplateMuxConfig) GoString() string {
	return s.String()
}

func (s *SearchTemplateResponseBodyTemplateListTemplateMuxConfig) GetGif() *SearchTemplateResponseBodyTemplateListTemplateMuxConfigGif {
	return s.Gif
}

func (s *SearchTemplateResponseBodyTemplateListTemplateMuxConfig) GetSegment() *SearchTemplateResponseBodyTemplateListTemplateMuxConfigSegment {
	return s.Segment
}

func (s *SearchTemplateResponseBodyTemplateListTemplateMuxConfig) SetGif(v *SearchTemplateResponseBodyTemplateListTemplateMuxConfigGif) *SearchTemplateResponseBodyTemplateListTemplateMuxConfig {
	s.Gif = v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateMuxConfig) SetSegment(v *SearchTemplateResponseBodyTemplateListTemplateMuxConfigSegment) *SearchTemplateResponseBodyTemplateListTemplateMuxConfig {
	s.Segment = v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateMuxConfig) Validate() error {
	if s.Gif != nil {
		if err := s.Gif.Validate(); err != nil {
			return err
		}
	}
	if s.Segment != nil {
		if err := s.Segment.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchTemplateResponseBodyTemplateListTemplateMuxConfigGif struct {
	// The color dithering algorithm of the palette. Valid values: sierra and bayer.
	//
	// example:
	//
	// sierra
	DitherMode *string `json:"DitherMode,omitempty" xml:"DitherMode,omitempty"`
	// The duration for which the final frame is paused. Unit: centisecond.
	//
	// example:
	//
	// 0
	FinalDelay *string `json:"FinalDelay,omitempty" xml:"FinalDelay,omitempty"`
	// Indicates whether a custom palette is used.
	//
	// example:
	//
	// false
	IsCustomPalette *string `json:"IsCustomPalette,omitempty" xml:"IsCustomPalette,omitempty"`
	// The loop count.
	//
	// example:
	//
	// 0
	Loop *string `json:"Loop,omitempty" xml:"Loop,omitempty"`
}

func (s SearchTemplateResponseBodyTemplateListTemplateMuxConfigGif) String() string {
	return dara.Prettify(s)
}

func (s SearchTemplateResponseBodyTemplateListTemplateMuxConfigGif) GoString() string {
	return s.String()
}

func (s *SearchTemplateResponseBodyTemplateListTemplateMuxConfigGif) GetDitherMode() *string {
	return s.DitherMode
}

func (s *SearchTemplateResponseBodyTemplateListTemplateMuxConfigGif) GetFinalDelay() *string {
	return s.FinalDelay
}

func (s *SearchTemplateResponseBodyTemplateListTemplateMuxConfigGif) GetIsCustomPalette() *string {
	return s.IsCustomPalette
}

func (s *SearchTemplateResponseBodyTemplateListTemplateMuxConfigGif) GetLoop() *string {
	return s.Loop
}

func (s *SearchTemplateResponseBodyTemplateListTemplateMuxConfigGif) SetDitherMode(v string) *SearchTemplateResponseBodyTemplateListTemplateMuxConfigGif {
	s.DitherMode = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateMuxConfigGif) SetFinalDelay(v string) *SearchTemplateResponseBodyTemplateListTemplateMuxConfigGif {
	s.FinalDelay = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateMuxConfigGif) SetIsCustomPalette(v string) *SearchTemplateResponseBodyTemplateListTemplateMuxConfigGif {
	s.IsCustomPalette = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateMuxConfigGif) SetLoop(v string) *SearchTemplateResponseBodyTemplateListTemplateMuxConfigGif {
	s.Loop = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateMuxConfigGif) Validate() error {
	return dara.Validate(s)
}

type SearchTemplateResponseBodyTemplateListTemplateMuxConfigSegment struct {
	// The length of the segment. Unit: seconds.
	//
	// example:
	//
	// 10
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s SearchTemplateResponseBodyTemplateListTemplateMuxConfigSegment) String() string {
	return dara.Prettify(s)
}

func (s SearchTemplateResponseBodyTemplateListTemplateMuxConfigSegment) GoString() string {
	return s.String()
}

func (s *SearchTemplateResponseBodyTemplateListTemplateMuxConfigSegment) GetDuration() *string {
	return s.Duration
}

func (s *SearchTemplateResponseBodyTemplateListTemplateMuxConfigSegment) SetDuration(v string) *SearchTemplateResponseBodyTemplateListTemplateMuxConfigSegment {
	s.Duration = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateMuxConfigSegment) Validate() error {
	return dara.Validate(s)
}

type SearchTemplateResponseBodyTemplateListTemplateTransConfig struct {
	// The method of resolution adjustment. Default value: **none**. Valid values:
	//
	// 	- rescale
	//
	// 	- crop
	//
	// 	- none
	//
	// example:
	//
	// none
	AdjDarMethod *string `json:"AdjDarMethod,omitempty" xml:"AdjDarMethod,omitempty"`
	// Indicates whether the audio bitrate is checked. If the bitrate of the output audio is higher than that of the input audio, the input bitrate is retained and the specified audio bitrate does not take effect. This parameter has a lower priority than IsCheckAudioBitrateFail. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// 	- Default value:
	//
	//     	- If this parameter is empty and the codec of the output audio is different from that of the input audio, the default value is false.
	//
	//     	- If this parameter is empty and the codec of the output audio is the same as that of the input audio, the default value is true.
	//
	// example:
	//
	// false
	IsCheckAudioBitrate *string `json:"IsCheckAudioBitrate,omitempty" xml:"IsCheckAudioBitrate,omitempty"`
	// Indicates whether audio bitrate check errors are allowed. This parameter has a greater priority than IsCheckAudioBitrate. Valid values:
	//
	// 	- **true**: If the audio bitrate check fails, the input file is not transcoded.
	//
	// 	- **false**: The audio bitrate is not checked.
	//
	// 	- Default value: **false**.
	//
	// example:
	//
	// false
	IsCheckAudioBitrateFail *string `json:"IsCheckAudioBitrateFail,omitempty" xml:"IsCheckAudioBitrateFail,omitempty"`
	// Indicates whether the resolution is checked. If the output resolution is higher than the input resolution based on the width or height, the input resolution is retained. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// 	- Default value: **false**.
	//
	// example:
	//
	// false
	IsCheckReso *string `json:"IsCheckReso,omitempty" xml:"IsCheckReso,omitempty"`
	// Indicates whether the resolution is checked. If the output resolution is higher than the input resolution based on the width or height, a transcoding failure is returned. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// 	- Default value: **false**.
	//
	// example:
	//
	// false
	IsCheckResoFail *string `json:"IsCheckResoFail,omitempty" xml:"IsCheckResoFail,omitempty"`
	// Indicates whether the video bitrate is checked. If the bitrate of the output video is higher than that of the input video, the input bitrate is retained. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// 	- Default value: **false**.
	//
	// example:
	//
	// false
	IsCheckVideoBitrate *string `json:"IsCheckVideoBitrate,omitempty" xml:"IsCheckVideoBitrate,omitempty"`
	// Indicates whether video bitrate check errors are allowed. This parameter has a higher priority than IsCheckVideoBitrate. Valid values:
	//
	// 	- **true**: If the video bitrate check fails, the input file is not transcoded.
	//
	// 	- **false**: The video bitrate is not checked.
	//
	// 	- Default value: **false**.
	//
	// example:
	//
	// false
	IsCheckVideoBitrateFail *string `json:"IsCheckVideoBitrateFail,omitempty" xml:"IsCheckVideoBitrateFail,omitempty"`
	// The transcoding mode. Default value: **onepass**. Valid values:
	//
	// 	- **onepass**
	//
	// 	- **twopass**
	//
	// 	- **CBR**
	//
	// example:
	//
	// onepass
	TransMode *string `json:"TransMode,omitempty" xml:"TransMode,omitempty"`
}

func (s SearchTemplateResponseBodyTemplateListTemplateTransConfig) String() string {
	return dara.Prettify(s)
}

func (s SearchTemplateResponseBodyTemplateListTemplateTransConfig) GoString() string {
	return s.String()
}

func (s *SearchTemplateResponseBodyTemplateListTemplateTransConfig) GetAdjDarMethod() *string {
	return s.AdjDarMethod
}

func (s *SearchTemplateResponseBodyTemplateListTemplateTransConfig) GetIsCheckAudioBitrate() *string {
	return s.IsCheckAudioBitrate
}

func (s *SearchTemplateResponseBodyTemplateListTemplateTransConfig) GetIsCheckAudioBitrateFail() *string {
	return s.IsCheckAudioBitrateFail
}

func (s *SearchTemplateResponseBodyTemplateListTemplateTransConfig) GetIsCheckReso() *string {
	return s.IsCheckReso
}

func (s *SearchTemplateResponseBodyTemplateListTemplateTransConfig) GetIsCheckResoFail() *string {
	return s.IsCheckResoFail
}

func (s *SearchTemplateResponseBodyTemplateListTemplateTransConfig) GetIsCheckVideoBitrate() *string {
	return s.IsCheckVideoBitrate
}

func (s *SearchTemplateResponseBodyTemplateListTemplateTransConfig) GetIsCheckVideoBitrateFail() *string {
	return s.IsCheckVideoBitrateFail
}

func (s *SearchTemplateResponseBodyTemplateListTemplateTransConfig) GetTransMode() *string {
	return s.TransMode
}

func (s *SearchTemplateResponseBodyTemplateListTemplateTransConfig) SetAdjDarMethod(v string) *SearchTemplateResponseBodyTemplateListTemplateTransConfig {
	s.AdjDarMethod = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateTransConfig) SetIsCheckAudioBitrate(v string) *SearchTemplateResponseBodyTemplateListTemplateTransConfig {
	s.IsCheckAudioBitrate = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateTransConfig) SetIsCheckAudioBitrateFail(v string) *SearchTemplateResponseBodyTemplateListTemplateTransConfig {
	s.IsCheckAudioBitrateFail = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateTransConfig) SetIsCheckReso(v string) *SearchTemplateResponseBodyTemplateListTemplateTransConfig {
	s.IsCheckReso = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateTransConfig) SetIsCheckResoFail(v string) *SearchTemplateResponseBodyTemplateListTemplateTransConfig {
	s.IsCheckResoFail = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateTransConfig) SetIsCheckVideoBitrate(v string) *SearchTemplateResponseBodyTemplateListTemplateTransConfig {
	s.IsCheckVideoBitrate = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateTransConfig) SetIsCheckVideoBitrateFail(v string) *SearchTemplateResponseBodyTemplateListTemplateTransConfig {
	s.IsCheckVideoBitrateFail = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateTransConfig) SetTransMode(v string) *SearchTemplateResponseBodyTemplateListTemplateTransConfig {
	s.TransMode = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateTransConfig) Validate() error {
	return dara.Validate(s)
}

type SearchTemplateResponseBodyTemplateListTemplateVideo struct {
	// The average bitrate of the video. Unit: Kbit/s.
	//
	// example:
	//
	// 200
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The average bitrate range of the video.
	BitrateBnd *SearchTemplateResponseBodyTemplateListTemplateVideoBitrateBnd `json:"BitrateBnd,omitempty" xml:"BitrateBnd,omitempty" type:"Struct"`
	// The buffer size.
	//
	// 	- Unit: KB.
	//
	// 	- Default value: **6000**.
	//
	// example:
	//
	// 6000
	Bufsize *string `json:"Bufsize,omitempty" xml:"Bufsize,omitempty"`
	// The codec.
	//
	// 	- Valid values: H.264 and H.265.
	//
	// 	- Default value: **H.264**.
	//
	// example:
	//
	// H.264
	Codec *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	// The constant rate factor.
	//
	// 	- Default value when the value of Codec is H.264: **23**, default value when the value of Codec is H.265: **26**.
	//
	// 	- If this parameter is set, the value of Bitrate becomes invalid.
	//
	// example:
	//
	// 15
	Crf *string `json:"Crf,omitempty" xml:"Crf,omitempty"`
	// The method of video cropping. Valid values:
	//
	// 	- **border**: automatically detects and removes black bars.
	//
	// 	- **Value in the width:height:left:top format**: crops the video image based on the custom settings. Format: width:height:left:top. Example: 1280:800:0:140.
	//
	// example:
	//
	// border
	Crop *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	// The level of video quality control.
	//
	// example:
	//
	// 10
	Degrain *string `json:"Degrain,omitempty" xml:"Degrain,omitempty"`
	// The frame rate of the video.
	//
	// 	- The value is 60 if the frame rate of the input video exceeds 60.
	//
	// 	- Default value: **the frame rate of the input video**.
	//
	// example:
	//
	// 25
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The maximum number of frames between two keyframes. Default value: **250**.
	//
	// example:
	//
	// 10
	Gop *string `json:"Gop,omitempty" xml:"Gop,omitempty"`
	// Indicates whether the HDR2SDR conversion feature is enabled. If this feature is enabled, high dynamic range (HDR) videos are transcoded to standard dynamic range (SDR) videos.
	//
	// example:
	//
	// true
	Hdr2sdr *string `json:"Hdr2sdr,omitempty" xml:"Hdr2sdr,omitempty"`
	// The height of the video.
	//
	// 	- Unit: pixel.
	//
	// 	- Default value: **the height of the input video**.
	//
	// example:
	//
	// 800
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// Indicates whether the auto-rotate screen feature is enabled.
	//
	// 	- If this feature is enabled, the width of the output video corresponds to the long side of the input video, which is the height of the input video in portrait mode. The height of the output video corresponds to the short side of the input video, which is the width of the input video in portrait mode. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// 	- Default value: **false**.
	//
	// example:
	//
	// false
	LongShortMode *string `json:"LongShortMode,omitempty" xml:"LongShortMode,omitempty"`
	// The maximum frame rate.
	//
	// example:
	//
	// 60
	MaxFps *string `json:"MaxFps,omitempty" xml:"MaxFps,omitempty"`
	// The maximum bitrate of the video. Unit: Kbit/s.
	//
	// example:
	//
	// 500
	Maxrate *string `json:"Maxrate,omitempty" xml:"Maxrate,omitempty"`
	// The Narrowband HD settings.
	NarrowBand *SearchTemplateResponseBodyTemplateListTemplateVideoNarrowBand `json:"NarrowBand,omitempty" xml:"NarrowBand,omitempty" type:"Struct"`
	// The black bars that are added to the video.
	//
	// 	- Format: width:height:left:top.
	//
	// 	- Example: 1280:800:0:140.
	//
	// example:
	//
	// 1280:800:0:140
	Pad *string `json:"Pad,omitempty" xml:"Pad,omitempty"`
	// The pixel format of the video. Valid values: standard pixel formats such as yuv420p and yuvj420p.
	//
	// example:
	//
	// yuv420p
	PixFmt *string `json:"PixFmt,omitempty" xml:"PixFmt,omitempty"`
	// The preset video algorithm. Default value: **medium**. Valid values:
	//
	// 	- **veryfast**
	//
	// 	- **fast**
	//
	// 	- **medium**
	//
	// 	- **slow**
	//
	// 	- **slower**
	//
	// example:
	//
	// medium
	Preset *string `json:"Preset,omitempty" xml:"Preset,omitempty"`
	// The codec profile. Valid values:
	//
	// 	- **baseline**: applicable to mobile devices.
	//
	// 	- **main**: applicable to standard-definition devices.
	//
	// 	- **high**: applicable to high-definition devices.
	//
	// 	- Default value: **high**.
	//
	// example:
	//
	// high
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The strength of the independent denoising algorithm.
	//
	// example:
	//
	// 1
	Qscale *string `json:"Qscale,omitempty" xml:"Qscale,omitempty"`
	// Indicates whether the video stream is deleted. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// 	- Default value: **false**.
	//
	// example:
	//
	// false
	Remove *string `json:"Remove,omitempty" xml:"Remove,omitempty"`
	// The policy of resolution adjustment.
	//
	// example:
	//
	// heightFirst
	ResoPriority *string `json:"ResoPriority,omitempty" xml:"ResoPriority,omitempty"`
	// The scan mode. Valid values:
	//
	// 	- **interlaced**
	//
	// 	- **progressive**
	//
	// example:
	//
	// interlaced
	ScanMode *string `json:"ScanMode,omitempty" xml:"ScanMode,omitempty"`
	// The width of the video.
	//
	// 	- Valid values: **[128,4096]**.
	//
	// 	- Unit: pixel.
	//
	// 	- Default value: **the width of the input video**.
	//
	// example:
	//
	// 256
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s SearchTemplateResponseBodyTemplateListTemplateVideo) String() string {
	return dara.Prettify(s)
}

func (s SearchTemplateResponseBodyTemplateListTemplateVideo) GoString() string {
	return s.String()
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) GetBitrate() *string {
	return s.Bitrate
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) GetBitrateBnd() *SearchTemplateResponseBodyTemplateListTemplateVideoBitrateBnd {
	return s.BitrateBnd
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) GetBufsize() *string {
	return s.Bufsize
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) GetCodec() *string {
	return s.Codec
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) GetCrf() *string {
	return s.Crf
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) GetCrop() *string {
	return s.Crop
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) GetDegrain() *string {
	return s.Degrain
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) GetFps() *string {
	return s.Fps
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) GetGop() *string {
	return s.Gop
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) GetHdr2sdr() *string {
	return s.Hdr2sdr
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) GetHeight() *string {
	return s.Height
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) GetLongShortMode() *string {
	return s.LongShortMode
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) GetMaxFps() *string {
	return s.MaxFps
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) GetMaxrate() *string {
	return s.Maxrate
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) GetNarrowBand() *SearchTemplateResponseBodyTemplateListTemplateVideoNarrowBand {
	return s.NarrowBand
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) GetPad() *string {
	return s.Pad
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) GetPixFmt() *string {
	return s.PixFmt
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) GetPreset() *string {
	return s.Preset
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) GetProfile() *string {
	return s.Profile
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) GetQscale() *string {
	return s.Qscale
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) GetRemove() *string {
	return s.Remove
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) GetResoPriority() *string {
	return s.ResoPriority
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) GetScanMode() *string {
	return s.ScanMode
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) GetWidth() *string {
	return s.Width
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) SetBitrate(v string) *SearchTemplateResponseBodyTemplateListTemplateVideo {
	s.Bitrate = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) SetBitrateBnd(v *SearchTemplateResponseBodyTemplateListTemplateVideoBitrateBnd) *SearchTemplateResponseBodyTemplateListTemplateVideo {
	s.BitrateBnd = v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) SetBufsize(v string) *SearchTemplateResponseBodyTemplateListTemplateVideo {
	s.Bufsize = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) SetCodec(v string) *SearchTemplateResponseBodyTemplateListTemplateVideo {
	s.Codec = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) SetCrf(v string) *SearchTemplateResponseBodyTemplateListTemplateVideo {
	s.Crf = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) SetCrop(v string) *SearchTemplateResponseBodyTemplateListTemplateVideo {
	s.Crop = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) SetDegrain(v string) *SearchTemplateResponseBodyTemplateListTemplateVideo {
	s.Degrain = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) SetFps(v string) *SearchTemplateResponseBodyTemplateListTemplateVideo {
	s.Fps = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) SetGop(v string) *SearchTemplateResponseBodyTemplateListTemplateVideo {
	s.Gop = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) SetHdr2sdr(v string) *SearchTemplateResponseBodyTemplateListTemplateVideo {
	s.Hdr2sdr = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) SetHeight(v string) *SearchTemplateResponseBodyTemplateListTemplateVideo {
	s.Height = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) SetLongShortMode(v string) *SearchTemplateResponseBodyTemplateListTemplateVideo {
	s.LongShortMode = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) SetMaxFps(v string) *SearchTemplateResponseBodyTemplateListTemplateVideo {
	s.MaxFps = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) SetMaxrate(v string) *SearchTemplateResponseBodyTemplateListTemplateVideo {
	s.Maxrate = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) SetNarrowBand(v *SearchTemplateResponseBodyTemplateListTemplateVideoNarrowBand) *SearchTemplateResponseBodyTemplateListTemplateVideo {
	s.NarrowBand = v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) SetPad(v string) *SearchTemplateResponseBodyTemplateListTemplateVideo {
	s.Pad = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) SetPixFmt(v string) *SearchTemplateResponseBodyTemplateListTemplateVideo {
	s.PixFmt = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) SetPreset(v string) *SearchTemplateResponseBodyTemplateListTemplateVideo {
	s.Preset = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) SetProfile(v string) *SearchTemplateResponseBodyTemplateListTemplateVideo {
	s.Profile = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) SetQscale(v string) *SearchTemplateResponseBodyTemplateListTemplateVideo {
	s.Qscale = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) SetRemove(v string) *SearchTemplateResponseBodyTemplateListTemplateVideo {
	s.Remove = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) SetResoPriority(v string) *SearchTemplateResponseBodyTemplateListTemplateVideo {
	s.ResoPriority = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) SetScanMode(v string) *SearchTemplateResponseBodyTemplateListTemplateVideo {
	s.ScanMode = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) SetWidth(v string) *SearchTemplateResponseBodyTemplateListTemplateVideo {
	s.Width = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideo) Validate() error {
	if s.BitrateBnd != nil {
		if err := s.BitrateBnd.Validate(); err != nil {
			return err
		}
	}
	if s.NarrowBand != nil {
		if err := s.NarrowBand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchTemplateResponseBodyTemplateListTemplateVideoBitrateBnd struct {
	// The upper limit of the total bitrate. Unit: Kbit/s.
	//
	// example:
	//
	// 500
	Max *string `json:"Max,omitempty" xml:"Max,omitempty"`
	// The lower limit of the total bitrate. Unit: Kbit/s.
	//
	// example:
	//
	// 100
	Min *string `json:"Min,omitempty" xml:"Min,omitempty"`
}

func (s SearchTemplateResponseBodyTemplateListTemplateVideoBitrateBnd) String() string {
	return dara.Prettify(s)
}

func (s SearchTemplateResponseBodyTemplateListTemplateVideoBitrateBnd) GoString() string {
	return s.String()
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideoBitrateBnd) GetMax() *string {
	return s.Max
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideoBitrateBnd) GetMin() *string {
	return s.Min
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideoBitrateBnd) SetMax(v string) *SearchTemplateResponseBodyTemplateListTemplateVideoBitrateBnd {
	s.Max = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideoBitrateBnd) SetMin(v string) *SearchTemplateResponseBodyTemplateListTemplateVideoBitrateBnd {
	s.Min = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideoBitrateBnd) Validate() error {
	return dara.Validate(s)
}

type SearchTemplateResponseBodyTemplateListTemplateVideoNarrowBand struct {
	// The upper limit of the dynamic bitrate. If this parameter is set, the average bitrate is in the range of (0, 1000000].
	//
	// example:
	//
	// 3000
	Abrmax *float32 `json:"Abrmax,omitempty" xml:"Abrmax,omitempty"`
	// The maximum ratio of the upper limit of dynamic bitrate. If this parameter is set, the value of Abrmax does not exceed x times of the source video bitrate. Valid values: (0,1.0].
	//
	// example:
	//
	// 1.0
	MaxAbrRatio *float32 `json:"MaxAbrRatio,omitempty" xml:"MaxAbrRatio,omitempty"`
	// The Narrowband HD version. Only 1.0 may be returned.
	//
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s SearchTemplateResponseBodyTemplateListTemplateVideoNarrowBand) String() string {
	return dara.Prettify(s)
}

func (s SearchTemplateResponseBodyTemplateListTemplateVideoNarrowBand) GoString() string {
	return s.String()
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideoNarrowBand) GetAbrmax() *float32 {
	return s.Abrmax
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideoNarrowBand) GetMaxAbrRatio() *float32 {
	return s.MaxAbrRatio
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideoNarrowBand) GetVersion() *string {
	return s.Version
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideoNarrowBand) SetAbrmax(v float32) *SearchTemplateResponseBodyTemplateListTemplateVideoNarrowBand {
	s.Abrmax = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideoNarrowBand) SetMaxAbrRatio(v float32) *SearchTemplateResponseBodyTemplateListTemplateVideoNarrowBand {
	s.MaxAbrRatio = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideoNarrowBand) SetVersion(v string) *SearchTemplateResponseBodyTemplateListTemplateVideoNarrowBand {
	s.Version = &v
	return s
}

func (s *SearchTemplateResponseBodyTemplateListTemplateVideoNarrowBand) Validate() error {
	return dara.Validate(s)
}
