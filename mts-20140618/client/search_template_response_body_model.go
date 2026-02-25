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
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Audio        *SearchTemplateResponseBodyTemplateListTemplateAudio       `json:"Audio,omitempty" xml:"Audio,omitempty" type:"Struct"`
	Container    *SearchTemplateResponseBodyTemplateListTemplateContainer   `json:"Container,omitempty" xml:"Container,omitempty" type:"Struct"`
	CreationTime *string                                                    `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Id           *string                                                    `json:"Id,omitempty" xml:"Id,omitempty"`
	MuxConfig    *SearchTemplateResponseBodyTemplateListTemplateMuxConfig   `json:"MuxConfig,omitempty" xml:"MuxConfig,omitempty" type:"Struct"`
	Name         *string                                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	State        *string                                                    `json:"State,omitempty" xml:"State,omitempty"`
	TransConfig  *SearchTemplateResponseBodyTemplateListTemplateTransConfig `json:"TransConfig,omitempty" xml:"TransConfig,omitempty" type:"Struct"`
	Video        *SearchTemplateResponseBodyTemplateListTemplateVideo       `json:"Video,omitempty" xml:"Video,omitempty" type:"Struct"`
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
	Bitrate    *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	Channels   *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	Codec      *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	Profile    *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	Qscale     *string `json:"Qscale,omitempty" xml:"Qscale,omitempty"`
	Remove     *string `json:"Remove,omitempty" xml:"Remove,omitempty"`
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
	Gif     *SearchTemplateResponseBodyTemplateListTemplateMuxConfigGif     `json:"Gif,omitempty" xml:"Gif,omitempty" type:"Struct"`
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
	DitherMode      *string `json:"DitherMode,omitempty" xml:"DitherMode,omitempty"`
	FinalDelay      *string `json:"FinalDelay,omitempty" xml:"FinalDelay,omitempty"`
	IsCustomPalette *string `json:"IsCustomPalette,omitempty" xml:"IsCustomPalette,omitempty"`
	Loop            *string `json:"Loop,omitempty" xml:"Loop,omitempty"`
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
	AdjDarMethod            *string `json:"AdjDarMethod,omitempty" xml:"AdjDarMethod,omitempty"`
	IsCheckAudioBitrate     *string `json:"IsCheckAudioBitrate,omitempty" xml:"IsCheckAudioBitrate,omitempty"`
	IsCheckAudioBitrateFail *string `json:"IsCheckAudioBitrateFail,omitempty" xml:"IsCheckAudioBitrateFail,omitempty"`
	IsCheckReso             *string `json:"IsCheckReso,omitempty" xml:"IsCheckReso,omitempty"`
	IsCheckResoFail         *string `json:"IsCheckResoFail,omitempty" xml:"IsCheckResoFail,omitempty"`
	IsCheckVideoBitrate     *string `json:"IsCheckVideoBitrate,omitempty" xml:"IsCheckVideoBitrate,omitempty"`
	IsCheckVideoBitrateFail *string `json:"IsCheckVideoBitrateFail,omitempty" xml:"IsCheckVideoBitrateFail,omitempty"`
	TransMode               *string `json:"TransMode,omitempty" xml:"TransMode,omitempty"`
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
	Bitrate       *string                                                        `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	BitrateBnd    *SearchTemplateResponseBodyTemplateListTemplateVideoBitrateBnd `json:"BitrateBnd,omitempty" xml:"BitrateBnd,omitempty" type:"Struct"`
	Bufsize       *string                                                        `json:"Bufsize,omitempty" xml:"Bufsize,omitempty"`
	Codec         *string                                                        `json:"Codec,omitempty" xml:"Codec,omitempty"`
	Crf           *string                                                        `json:"Crf,omitempty" xml:"Crf,omitempty"`
	Crop          *string                                                        `json:"Crop,omitempty" xml:"Crop,omitempty"`
	Degrain       *string                                                        `json:"Degrain,omitempty" xml:"Degrain,omitempty"`
	Fps           *string                                                        `json:"Fps,omitempty" xml:"Fps,omitempty"`
	Gop           *string                                                        `json:"Gop,omitempty" xml:"Gop,omitempty"`
	Hdr2sdr       *string                                                        `json:"Hdr2sdr,omitempty" xml:"Hdr2sdr,omitempty"`
	Height        *string                                                        `json:"Height,omitempty" xml:"Height,omitempty"`
	LongShortMode *string                                                        `json:"LongShortMode,omitempty" xml:"LongShortMode,omitempty"`
	MaxFps        *string                                                        `json:"MaxFps,omitempty" xml:"MaxFps,omitempty"`
	Maxrate       *string                                                        `json:"Maxrate,omitempty" xml:"Maxrate,omitempty"`
	NarrowBand    *SearchTemplateResponseBodyTemplateListTemplateVideoNarrowBand `json:"NarrowBand,omitempty" xml:"NarrowBand,omitempty" type:"Struct"`
	Pad           *string                                                        `json:"Pad,omitempty" xml:"Pad,omitempty"`
	PixFmt        *string                                                        `json:"PixFmt,omitempty" xml:"PixFmt,omitempty"`
	Preset        *string                                                        `json:"Preset,omitempty" xml:"Preset,omitempty"`
	Profile       *string                                                        `json:"Profile,omitempty" xml:"Profile,omitempty"`
	Qscale        *string                                                        `json:"Qscale,omitempty" xml:"Qscale,omitempty"`
	Remove        *string                                                        `json:"Remove,omitempty" xml:"Remove,omitempty"`
	ResoPriority  *string                                                        `json:"ResoPriority,omitempty" xml:"ResoPriority,omitempty"`
	ScanMode      *string                                                        `json:"ScanMode,omitempty" xml:"ScanMode,omitempty"`
	Width         *string                                                        `json:"Width,omitempty" xml:"Width,omitempty"`
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
	Max *string `json:"Max,omitempty" xml:"Max,omitempty"`
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
	Abrmax      *float32 `json:"Abrmax,omitempty" xml:"Abrmax,omitempty"`
	MaxAbrRatio *float32 `json:"MaxAbrRatio,omitempty" xml:"MaxAbrRatio,omitempty"`
	Version     *string  `json:"Version,omitempty" xml:"Version,omitempty"`
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
