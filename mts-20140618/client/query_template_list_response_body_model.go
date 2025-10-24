// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTemplateListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNonExistTids(v *QueryTemplateListResponseBodyNonExistTids) *QueryTemplateListResponseBody
	GetNonExistTids() *QueryTemplateListResponseBodyNonExistTids
	SetRequestId(v string) *QueryTemplateListResponseBody
	GetRequestId() *string
	SetTemplateList(v *QueryTemplateListResponseBodyTemplateList) *QueryTemplateListResponseBody
	GetTemplateList() *QueryTemplateListResponseBodyTemplateList
}

type QueryTemplateListResponseBody struct {
	// The IDs of the templates that do not exist. This parameter is not returned if all specified transcoding templates are found.
	NonExistTids *QueryTemplateListResponseBodyNonExistTids `json:"NonExistTids,omitempty" xml:"NonExistTids,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// BC860F04-778A-472F-AB39-E1BF329C1EA8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The transcoding templates.
	TemplateList *QueryTemplateListResponseBodyTemplateList `json:"TemplateList,omitempty" xml:"TemplateList,omitempty" type:"Struct"`
}

func (s QueryTemplateListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTemplateListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTemplateListResponseBody) GetNonExistTids() *QueryTemplateListResponseBodyNonExistTids {
	return s.NonExistTids
}

func (s *QueryTemplateListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTemplateListResponseBody) GetTemplateList() *QueryTemplateListResponseBodyTemplateList {
	return s.TemplateList
}

func (s *QueryTemplateListResponseBody) SetNonExistTids(v *QueryTemplateListResponseBodyNonExistTids) *QueryTemplateListResponseBody {
	s.NonExistTids = v
	return s
}

func (s *QueryTemplateListResponseBody) SetRequestId(v string) *QueryTemplateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTemplateListResponseBody) SetTemplateList(v *QueryTemplateListResponseBodyTemplateList) *QueryTemplateListResponseBody {
	s.TemplateList = v
	return s
}

func (s *QueryTemplateListResponseBody) Validate() error {
	if s.NonExistTids != nil {
		if err := s.NonExistTids.Validate(); err != nil {
			return err
		}
	}
	if s.TemplateList != nil {
		if err := s.TemplateList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryTemplateListResponseBodyNonExistTids struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s QueryTemplateListResponseBodyNonExistTids) String() string {
	return dara.Prettify(s)
}

func (s QueryTemplateListResponseBodyNonExistTids) GoString() string {
	return s.String()
}

func (s *QueryTemplateListResponseBodyNonExistTids) GetString_() []*string {
	return s.String_
}

func (s *QueryTemplateListResponseBodyNonExistTids) SetString_(v []*string) *QueryTemplateListResponseBodyNonExistTids {
	s.String_ = v
	return s
}

func (s *QueryTemplateListResponseBodyNonExistTids) Validate() error {
	return dara.Validate(s)
}

type QueryTemplateListResponseBodyTemplateList struct {
	Template []*QueryTemplateListResponseBodyTemplateListTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Repeated"`
}

func (s QueryTemplateListResponseBodyTemplateList) String() string {
	return dara.Prettify(s)
}

func (s QueryTemplateListResponseBodyTemplateList) GoString() string {
	return s.String()
}

func (s *QueryTemplateListResponseBodyTemplateList) GetTemplate() []*QueryTemplateListResponseBodyTemplateListTemplate {
	return s.Template
}

func (s *QueryTemplateListResponseBodyTemplateList) SetTemplate(v []*QueryTemplateListResponseBodyTemplateListTemplate) *QueryTemplateListResponseBodyTemplateList {
	s.Template = v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateList) Validate() error {
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

type QueryTemplateListResponseBodyTemplateListTemplate struct {
	// The audio codec configurations.
	Audio *QueryTemplateListResponseBodyTemplateListTemplateAudio `json:"Audio,omitempty" xml:"Audio,omitempty" type:"Struct"`
	// The container format configurations.
	Container *QueryTemplateListResponseBodyTemplateListTemplateContainer `json:"Container,omitempty" xml:"Container,omitempty" type:"Struct"`
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
	MuxConfig *QueryTemplateListResponseBodyTemplateListTemplateMuxConfig `json:"MuxConfig,omitempty" xml:"MuxConfig,omitempty" type:"Struct"`
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
	TransConfig *QueryTemplateListResponseBodyTemplateListTemplateTransConfig `json:"TransConfig,omitempty" xml:"TransConfig,omitempty" type:"Struct"`
	// The video codec configurations.
	Video *QueryTemplateListResponseBodyTemplateListTemplateVideo `json:"Video,omitempty" xml:"Video,omitempty" type:"Struct"`
}

func (s QueryTemplateListResponseBodyTemplateListTemplate) String() string {
	return dara.Prettify(s)
}

func (s QueryTemplateListResponseBodyTemplateListTemplate) GoString() string {
	return s.String()
}

func (s *QueryTemplateListResponseBodyTemplateListTemplate) GetAudio() *QueryTemplateListResponseBodyTemplateListTemplateAudio {
	return s.Audio
}

func (s *QueryTemplateListResponseBodyTemplateListTemplate) GetContainer() *QueryTemplateListResponseBodyTemplateListTemplateContainer {
	return s.Container
}

func (s *QueryTemplateListResponseBodyTemplateListTemplate) GetCreationTime() *string {
	return s.CreationTime
}

func (s *QueryTemplateListResponseBodyTemplateListTemplate) GetId() *string {
	return s.Id
}

func (s *QueryTemplateListResponseBodyTemplateListTemplate) GetMuxConfig() *QueryTemplateListResponseBodyTemplateListTemplateMuxConfig {
	return s.MuxConfig
}

func (s *QueryTemplateListResponseBodyTemplateListTemplate) GetName() *string {
	return s.Name
}

func (s *QueryTemplateListResponseBodyTemplateListTemplate) GetState() *string {
	return s.State
}

func (s *QueryTemplateListResponseBodyTemplateListTemplate) GetTransConfig() *QueryTemplateListResponseBodyTemplateListTemplateTransConfig {
	return s.TransConfig
}

func (s *QueryTemplateListResponseBodyTemplateListTemplate) GetVideo() *QueryTemplateListResponseBodyTemplateListTemplateVideo {
	return s.Video
}

func (s *QueryTemplateListResponseBodyTemplateListTemplate) SetAudio(v *QueryTemplateListResponseBodyTemplateListTemplateAudio) *QueryTemplateListResponseBodyTemplateListTemplate {
	s.Audio = v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplate) SetContainer(v *QueryTemplateListResponseBodyTemplateListTemplateContainer) *QueryTemplateListResponseBodyTemplateListTemplate {
	s.Container = v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplate) SetCreationTime(v string) *QueryTemplateListResponseBodyTemplateListTemplate {
	s.CreationTime = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplate) SetId(v string) *QueryTemplateListResponseBodyTemplateListTemplate {
	s.Id = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplate) SetMuxConfig(v *QueryTemplateListResponseBodyTemplateListTemplateMuxConfig) *QueryTemplateListResponseBodyTemplateListTemplate {
	s.MuxConfig = v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplate) SetName(v string) *QueryTemplateListResponseBodyTemplateListTemplate {
	s.Name = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplate) SetState(v string) *QueryTemplateListResponseBodyTemplateListTemplate {
	s.State = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplate) SetTransConfig(v *QueryTemplateListResponseBodyTemplateListTemplateTransConfig) *QueryTemplateListResponseBodyTemplateListTemplate {
	s.TransConfig = v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplate) SetVideo(v *QueryTemplateListResponseBodyTemplateListTemplateVideo) *QueryTemplateListResponseBodyTemplateListTemplate {
	s.Video = v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplate) Validate() error {
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

type QueryTemplateListResponseBodyTemplateListTemplateAudio struct {
	// The audio bitrate of the output file.
	//
	// 	- Valid values: **[8,1000]**.
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
	// The strength of the independent denoising algorithm.
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
	// The volume control configurations.
	Volume *QueryTemplateListResponseBodyTemplateListTemplateAudioVolume `json:"Volume,omitempty" xml:"Volume,omitempty" type:"Struct"`
}

func (s QueryTemplateListResponseBodyTemplateListTemplateAudio) String() string {
	return dara.Prettify(s)
}

func (s QueryTemplateListResponseBodyTemplateListTemplateAudio) GoString() string {
	return s.String()
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateAudio) GetBitrate() *string {
	return s.Bitrate
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateAudio) GetChannels() *string {
	return s.Channels
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateAudio) GetCodec() *string {
	return s.Codec
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateAudio) GetProfile() *string {
	return s.Profile
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateAudio) GetQscale() *string {
	return s.Qscale
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateAudio) GetRemove() *string {
	return s.Remove
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateAudio) GetSamplerate() *string {
	return s.Samplerate
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateAudio) GetVolume() *QueryTemplateListResponseBodyTemplateListTemplateAudioVolume {
	return s.Volume
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateAudio) SetBitrate(v string) *QueryTemplateListResponseBodyTemplateListTemplateAudio {
	s.Bitrate = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateAudio) SetChannels(v string) *QueryTemplateListResponseBodyTemplateListTemplateAudio {
	s.Channels = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateAudio) SetCodec(v string) *QueryTemplateListResponseBodyTemplateListTemplateAudio {
	s.Codec = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateAudio) SetProfile(v string) *QueryTemplateListResponseBodyTemplateListTemplateAudio {
	s.Profile = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateAudio) SetQscale(v string) *QueryTemplateListResponseBodyTemplateListTemplateAudio {
	s.Qscale = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateAudio) SetRemove(v string) *QueryTemplateListResponseBodyTemplateListTemplateAudio {
	s.Remove = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateAudio) SetSamplerate(v string) *QueryTemplateListResponseBodyTemplateListTemplateAudio {
	s.Samplerate = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateAudio) SetVolume(v *QueryTemplateListResponseBodyTemplateListTemplateAudioVolume) *QueryTemplateListResponseBodyTemplateListTemplateAudio {
	s.Volume = v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateAudio) Validate() error {
	if s.Volume != nil {
		if err := s.Volume.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryTemplateListResponseBodyTemplateListTemplateAudioVolume struct {
	// The output volume.
	//
	// 	- This parameter takes effect only if the value of Method is dynamic.
	//
	// 	- Unit: dB.
	//
	// 	- Valid values: [-70,-5].
	//
	// 	- Default value: -6.
	//
	// example:
	//
	// -6
	IntegratedLoudnessTarget *string `json:"IntegratedLoudnessTarget,omitempty" xml:"IntegratedLoudnessTarget,omitempty"`
	// The increased volume relative to the volume of the input audio.
	//
	// 	- This parameter takes effect only if the value of Method is linear.
	//
	// 	- Unit: dB.
	//
	// 	- Valid values: less than or equal to 20.
	//
	// 	- Default value: -20.
	//
	// example:
	//
	// -20
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The range of the volume relative to the output volume.
	//
	// 	- This parameter takes effect only if the value of Method is dynamic.
	//
	// 	- Unit: dB.
	//
	// 	- Valid values: [1,20].
	//
	// 	- Default value: 8.
	//
	// example:
	//
	// 8
	LoudnessRangeTarget *string `json:"LoudnessRangeTarget,omitempty" xml:"LoudnessRangeTarget,omitempty"`
	// The method that is used to adjust the volume. Valid values:
	//
	// 	- **auto**
	//
	// 	- **dynamic**
	//
	// 	- **linear**
	//
	// 	- Default value: dynamic.
	//
	// example:
	//
	// auto
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The volume adjustment coefficient.
	//
	// This parameter takes effect only if the value of Method is adaptive.
	//
	// Valid values: [0,1].
	//
	// Default value: 0.9.
	//
	// example:
	//
	// 0.9
	PeakLevel *string `json:"PeakLevel,omitempty" xml:"PeakLevel,omitempty"`
	// The peak volume.
	//
	// 	- This parameter takes effect only if the value of Method is dynamic.
	//
	// 	- Unit: dB.
	//
	// 	- Valid values: [-9,0].
	//
	// 	- Default value: -1.
	//
	// example:
	//
	// -1
	TruePeak *string `json:"TruePeak,omitempty" xml:"TruePeak,omitempty"`
}

func (s QueryTemplateListResponseBodyTemplateListTemplateAudioVolume) String() string {
	return dara.Prettify(s)
}

func (s QueryTemplateListResponseBodyTemplateListTemplateAudioVolume) GoString() string {
	return s.String()
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateAudioVolume) GetIntegratedLoudnessTarget() *string {
	return s.IntegratedLoudnessTarget
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateAudioVolume) GetLevel() *string {
	return s.Level
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateAudioVolume) GetLoudnessRangeTarget() *string {
	return s.LoudnessRangeTarget
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateAudioVolume) GetMethod() *string {
	return s.Method
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateAudioVolume) GetPeakLevel() *string {
	return s.PeakLevel
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateAudioVolume) GetTruePeak() *string {
	return s.TruePeak
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateAudioVolume) SetIntegratedLoudnessTarget(v string) *QueryTemplateListResponseBodyTemplateListTemplateAudioVolume {
	s.IntegratedLoudnessTarget = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateAudioVolume) SetLevel(v string) *QueryTemplateListResponseBodyTemplateListTemplateAudioVolume {
	s.Level = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateAudioVolume) SetLoudnessRangeTarget(v string) *QueryTemplateListResponseBodyTemplateListTemplateAudioVolume {
	s.LoudnessRangeTarget = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateAudioVolume) SetMethod(v string) *QueryTemplateListResponseBodyTemplateListTemplateAudioVolume {
	s.Method = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateAudioVolume) SetPeakLevel(v string) *QueryTemplateListResponseBodyTemplateListTemplateAudioVolume {
	s.PeakLevel = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateAudioVolume) SetTruePeak(v string) *QueryTemplateListResponseBodyTemplateListTemplateAudioVolume {
	s.TruePeak = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateAudioVolume) Validate() error {
	return dara.Validate(s)
}

type QueryTemplateListResponseBodyTemplateListTemplateContainer struct {
	// The format of the container. Valid values: flv, mp4, ts, m3u8, gif, mp3, ogg, and flac.
	//
	// example:
	//
	// flv
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
}

func (s QueryTemplateListResponseBodyTemplateListTemplateContainer) String() string {
	return dara.Prettify(s)
}

func (s QueryTemplateListResponseBodyTemplateListTemplateContainer) GoString() string {
	return s.String()
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateContainer) GetFormat() *string {
	return s.Format
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateContainer) SetFormat(v string) *QueryTemplateListResponseBodyTemplateListTemplateContainer {
	s.Format = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateContainer) Validate() error {
	return dara.Validate(s)
}

type QueryTemplateListResponseBodyTemplateListTemplateMuxConfig struct {
	// The transmuxing configurations for GIF.
	Gif *QueryTemplateListResponseBodyTemplateListTemplateMuxConfigGif `json:"Gif,omitempty" xml:"Gif,omitempty" type:"Struct"`
	// The segment configurations.
	Segment *QueryTemplateListResponseBodyTemplateListTemplateMuxConfigSegment `json:"Segment,omitempty" xml:"Segment,omitempty" type:"Struct"`
	// The transmuxing configurations for WebP.
	Webp *QueryTemplateListResponseBodyTemplateListTemplateMuxConfigWebp `json:"Webp,omitempty" xml:"Webp,omitempty" type:"Struct"`
}

func (s QueryTemplateListResponseBodyTemplateListTemplateMuxConfig) String() string {
	return dara.Prettify(s)
}

func (s QueryTemplateListResponseBodyTemplateListTemplateMuxConfig) GoString() string {
	return s.String()
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateMuxConfig) GetGif() *QueryTemplateListResponseBodyTemplateListTemplateMuxConfigGif {
	return s.Gif
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateMuxConfig) GetSegment() *QueryTemplateListResponseBodyTemplateListTemplateMuxConfigSegment {
	return s.Segment
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateMuxConfig) GetWebp() *QueryTemplateListResponseBodyTemplateListTemplateMuxConfigWebp {
	return s.Webp
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateMuxConfig) SetGif(v *QueryTemplateListResponseBodyTemplateListTemplateMuxConfigGif) *QueryTemplateListResponseBodyTemplateListTemplateMuxConfig {
	s.Gif = v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateMuxConfig) SetSegment(v *QueryTemplateListResponseBodyTemplateListTemplateMuxConfigSegment) *QueryTemplateListResponseBodyTemplateListTemplateMuxConfig {
	s.Segment = v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateMuxConfig) SetWebp(v *QueryTemplateListResponseBodyTemplateListTemplateMuxConfigWebp) *QueryTemplateListResponseBodyTemplateListTemplateMuxConfig {
	s.Webp = v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateMuxConfig) Validate() error {
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
	if s.Webp != nil {
		if err := s.Webp.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryTemplateListResponseBodyTemplateListTemplateMuxConfigGif struct {
	// The color dithering algorithm of the palette. Valid values: sierra and bayer.
	//
	// example:
	//
	// sierra
	DitherMode *string `json:"DitherMode,omitempty" xml:"DitherMode,omitempty"`
	// The duration for which the final frame is paused.
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

func (s QueryTemplateListResponseBodyTemplateListTemplateMuxConfigGif) String() string {
	return dara.Prettify(s)
}

func (s QueryTemplateListResponseBodyTemplateListTemplateMuxConfigGif) GoString() string {
	return s.String()
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateMuxConfigGif) GetDitherMode() *string {
	return s.DitherMode
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateMuxConfigGif) GetFinalDelay() *string {
	return s.FinalDelay
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateMuxConfigGif) GetIsCustomPalette() *string {
	return s.IsCustomPalette
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateMuxConfigGif) GetLoop() *string {
	return s.Loop
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateMuxConfigGif) SetDitherMode(v string) *QueryTemplateListResponseBodyTemplateListTemplateMuxConfigGif {
	s.DitherMode = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateMuxConfigGif) SetFinalDelay(v string) *QueryTemplateListResponseBodyTemplateListTemplateMuxConfigGif {
	s.FinalDelay = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateMuxConfigGif) SetIsCustomPalette(v string) *QueryTemplateListResponseBodyTemplateListTemplateMuxConfigGif {
	s.IsCustomPalette = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateMuxConfigGif) SetLoop(v string) *QueryTemplateListResponseBodyTemplateListTemplateMuxConfigGif {
	s.Loop = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateMuxConfigGif) Validate() error {
	return dara.Validate(s)
}

type QueryTemplateListResponseBodyTemplateListTemplateMuxConfigSegment struct {
	// The length of the segment. Unit: seconds.
	//
	// example:
	//
	// 10
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s QueryTemplateListResponseBodyTemplateListTemplateMuxConfigSegment) String() string {
	return dara.Prettify(s)
}

func (s QueryTemplateListResponseBodyTemplateListTemplateMuxConfigSegment) GoString() string {
	return s.String()
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateMuxConfigSegment) GetDuration() *string {
	return s.Duration
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateMuxConfigSegment) SetDuration(v string) *QueryTemplateListResponseBodyTemplateListTemplateMuxConfigSegment {
	s.Duration = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateMuxConfigSegment) Validate() error {
	return dara.Validate(s)
}

type QueryTemplateListResponseBodyTemplateListTemplateMuxConfigWebp struct {
	// The loop count.
	//
	// example:
	//
	// 0
	Loop *string `json:"Loop,omitempty" xml:"Loop,omitempty"`
}

func (s QueryTemplateListResponseBodyTemplateListTemplateMuxConfigWebp) String() string {
	return dara.Prettify(s)
}

func (s QueryTemplateListResponseBodyTemplateListTemplateMuxConfigWebp) GoString() string {
	return s.String()
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateMuxConfigWebp) GetLoop() *string {
	return s.Loop
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateMuxConfigWebp) SetLoop(v string) *QueryTemplateListResponseBodyTemplateListTemplateMuxConfigWebp {
	s.Loop = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateMuxConfigWebp) Validate() error {
	return dara.Validate(s)
}

type QueryTemplateListResponseBodyTemplateListTemplateTransConfig struct {
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
	//     	- If this parameter is not specified and the codec of the output audio is different from that of the input audio, the default value is false.
	//
	//     	- If this parameter is not specified and the codec of the output audio is the same as that of the input audio, the default value is true.
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

func (s QueryTemplateListResponseBodyTemplateListTemplateTransConfig) String() string {
	return dara.Prettify(s)
}

func (s QueryTemplateListResponseBodyTemplateListTemplateTransConfig) GoString() string {
	return s.String()
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateTransConfig) GetAdjDarMethod() *string {
	return s.AdjDarMethod
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateTransConfig) GetIsCheckAudioBitrate() *string {
	return s.IsCheckAudioBitrate
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateTransConfig) GetIsCheckAudioBitrateFail() *string {
	return s.IsCheckAudioBitrateFail
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateTransConfig) GetIsCheckReso() *string {
	return s.IsCheckReso
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateTransConfig) GetIsCheckResoFail() *string {
	return s.IsCheckResoFail
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateTransConfig) GetIsCheckVideoBitrate() *string {
	return s.IsCheckVideoBitrate
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateTransConfig) GetIsCheckVideoBitrateFail() *string {
	return s.IsCheckVideoBitrateFail
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateTransConfig) GetTransMode() *string {
	return s.TransMode
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateTransConfig) SetAdjDarMethod(v string) *QueryTemplateListResponseBodyTemplateListTemplateTransConfig {
	s.AdjDarMethod = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateTransConfig) SetIsCheckAudioBitrate(v string) *QueryTemplateListResponseBodyTemplateListTemplateTransConfig {
	s.IsCheckAudioBitrate = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateTransConfig) SetIsCheckAudioBitrateFail(v string) *QueryTemplateListResponseBodyTemplateListTemplateTransConfig {
	s.IsCheckAudioBitrateFail = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateTransConfig) SetIsCheckReso(v string) *QueryTemplateListResponseBodyTemplateListTemplateTransConfig {
	s.IsCheckReso = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateTransConfig) SetIsCheckResoFail(v string) *QueryTemplateListResponseBodyTemplateListTemplateTransConfig {
	s.IsCheckResoFail = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateTransConfig) SetIsCheckVideoBitrate(v string) *QueryTemplateListResponseBodyTemplateListTemplateTransConfig {
	s.IsCheckVideoBitrate = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateTransConfig) SetIsCheckVideoBitrateFail(v string) *QueryTemplateListResponseBodyTemplateListTemplateTransConfig {
	s.IsCheckVideoBitrateFail = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateTransConfig) SetTransMode(v string) *QueryTemplateListResponseBodyTemplateListTemplateTransConfig {
	s.TransMode = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateTransConfig) Validate() error {
	return dara.Validate(s)
}

type QueryTemplateListResponseBodyTemplateListTemplateVideo struct {
	// The average bitrate of the video. Unit: Kbit/s.
	//
	// example:
	//
	// 200
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The average bitrate range of the video.
	BitrateBnd *QueryTemplateListResponseBodyTemplateListTemplateVideoBitrateBnd `json:"BitrateBnd,omitempty" xml:"BitrateBnd,omitempty" type:"Struct"`
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
	// The codec. Default value: **H.264**.
	//
	// example:
	//
	// H.264
	Codec *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	// The constant rate factor.
	//
	// 	- Default value when the encoding format is H.264: **23**. Default value when the encoding format is H.265: **26**.
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
	// 	- **Value in the width:height:left:top format**: crops the video based on custom settings. Format: width:height:left:top. Example: 1280:800:0:140.
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
	// The frame rate.
	//
	// 	- The value is 60 if the frame rate of the input file exceeds 60.
	//
	// 	- Default value: **the frame rate of the input file**.
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
	NarrowBand *QueryTemplateListResponseBodyTemplateListTemplateVideoNarrowBand `json:"NarrowBand,omitempty" xml:"NarrowBand,omitempty" type:"Struct"`
	// The black bars that are added to the video. Format: width:height:left:top. Example: 1280:800:0:140.
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
	// The preset video algorithm. Valid values: veryfast, fast, medium, slow, and slower. Default value: **medium**.
	//
	// example:
	//
	// medium
	Preset *string `json:"Preset,omitempty" xml:"Preset,omitempty"`
	// The encoding profile. Valid values:
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
	// The policy of resolution adjustment. Valid values: cropFirst, widthFirst, and heightFirst.
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
	// 	- Unit: pixel.
	//
	// 	- Default value: **the width of the input video**.
	//
	// example:
	//
	// 256
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s QueryTemplateListResponseBodyTemplateListTemplateVideo) String() string {
	return dara.Prettify(s)
}

func (s QueryTemplateListResponseBodyTemplateListTemplateVideo) GoString() string {
	return s.String()
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) GetBitrate() *string {
	return s.Bitrate
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) GetBitrateBnd() *QueryTemplateListResponseBodyTemplateListTemplateVideoBitrateBnd {
	return s.BitrateBnd
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) GetBufsize() *string {
	return s.Bufsize
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) GetCodec() *string {
	return s.Codec
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) GetCrf() *string {
	return s.Crf
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) GetCrop() *string {
	return s.Crop
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) GetDegrain() *string {
	return s.Degrain
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) GetFps() *string {
	return s.Fps
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) GetGop() *string {
	return s.Gop
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) GetHdr2sdr() *string {
	return s.Hdr2sdr
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) GetHeight() *string {
	return s.Height
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) GetLongShortMode() *string {
	return s.LongShortMode
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) GetMaxFps() *string {
	return s.MaxFps
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) GetMaxrate() *string {
	return s.Maxrate
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) GetNarrowBand() *QueryTemplateListResponseBodyTemplateListTemplateVideoNarrowBand {
	return s.NarrowBand
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) GetPad() *string {
	return s.Pad
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) GetPixFmt() *string {
	return s.PixFmt
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) GetPreset() *string {
	return s.Preset
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) GetProfile() *string {
	return s.Profile
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) GetQscale() *string {
	return s.Qscale
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) GetRemove() *string {
	return s.Remove
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) GetResoPriority() *string {
	return s.ResoPriority
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) GetScanMode() *string {
	return s.ScanMode
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) GetWidth() *string {
	return s.Width
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) SetBitrate(v string) *QueryTemplateListResponseBodyTemplateListTemplateVideo {
	s.Bitrate = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) SetBitrateBnd(v *QueryTemplateListResponseBodyTemplateListTemplateVideoBitrateBnd) *QueryTemplateListResponseBodyTemplateListTemplateVideo {
	s.BitrateBnd = v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) SetBufsize(v string) *QueryTemplateListResponseBodyTemplateListTemplateVideo {
	s.Bufsize = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) SetCodec(v string) *QueryTemplateListResponseBodyTemplateListTemplateVideo {
	s.Codec = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) SetCrf(v string) *QueryTemplateListResponseBodyTemplateListTemplateVideo {
	s.Crf = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) SetCrop(v string) *QueryTemplateListResponseBodyTemplateListTemplateVideo {
	s.Crop = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) SetDegrain(v string) *QueryTemplateListResponseBodyTemplateListTemplateVideo {
	s.Degrain = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) SetFps(v string) *QueryTemplateListResponseBodyTemplateListTemplateVideo {
	s.Fps = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) SetGop(v string) *QueryTemplateListResponseBodyTemplateListTemplateVideo {
	s.Gop = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) SetHdr2sdr(v string) *QueryTemplateListResponseBodyTemplateListTemplateVideo {
	s.Hdr2sdr = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) SetHeight(v string) *QueryTemplateListResponseBodyTemplateListTemplateVideo {
	s.Height = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) SetLongShortMode(v string) *QueryTemplateListResponseBodyTemplateListTemplateVideo {
	s.LongShortMode = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) SetMaxFps(v string) *QueryTemplateListResponseBodyTemplateListTemplateVideo {
	s.MaxFps = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) SetMaxrate(v string) *QueryTemplateListResponseBodyTemplateListTemplateVideo {
	s.Maxrate = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) SetNarrowBand(v *QueryTemplateListResponseBodyTemplateListTemplateVideoNarrowBand) *QueryTemplateListResponseBodyTemplateListTemplateVideo {
	s.NarrowBand = v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) SetPad(v string) *QueryTemplateListResponseBodyTemplateListTemplateVideo {
	s.Pad = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) SetPixFmt(v string) *QueryTemplateListResponseBodyTemplateListTemplateVideo {
	s.PixFmt = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) SetPreset(v string) *QueryTemplateListResponseBodyTemplateListTemplateVideo {
	s.Preset = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) SetProfile(v string) *QueryTemplateListResponseBodyTemplateListTemplateVideo {
	s.Profile = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) SetQscale(v string) *QueryTemplateListResponseBodyTemplateListTemplateVideo {
	s.Qscale = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) SetRemove(v string) *QueryTemplateListResponseBodyTemplateListTemplateVideo {
	s.Remove = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) SetResoPriority(v string) *QueryTemplateListResponseBodyTemplateListTemplateVideo {
	s.ResoPriority = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) SetScanMode(v string) *QueryTemplateListResponseBodyTemplateListTemplateVideo {
	s.ScanMode = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) SetWidth(v string) *QueryTemplateListResponseBodyTemplateListTemplateVideo {
	s.Width = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideo) Validate() error {
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

type QueryTemplateListResponseBodyTemplateListTemplateVideoBitrateBnd struct {
	// The upper limit of the total bitrate. Unit: Kbit/s.
	//
	// example:
	//
	// 100
	Max *string `json:"Max,omitempty" xml:"Max,omitempty"`
	// The lower limit of the total bitrate. Unit: Kbit/s.
	//
	// example:
	//
	// 500
	Min *string `json:"Min,omitempty" xml:"Min,omitempty"`
}

func (s QueryTemplateListResponseBodyTemplateListTemplateVideoBitrateBnd) String() string {
	return dara.Prettify(s)
}

func (s QueryTemplateListResponseBodyTemplateListTemplateVideoBitrateBnd) GoString() string {
	return s.String()
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideoBitrateBnd) GetMax() *string {
	return s.Max
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideoBitrateBnd) GetMin() *string {
	return s.Min
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideoBitrateBnd) SetMax(v string) *QueryTemplateListResponseBodyTemplateListTemplateVideoBitrateBnd {
	s.Max = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideoBitrateBnd) SetMin(v string) *QueryTemplateListResponseBodyTemplateListTemplateVideoBitrateBnd {
	s.Min = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideoBitrateBnd) Validate() error {
	return dara.Validate(s)
}

type QueryTemplateListResponseBodyTemplateListTemplateVideoNarrowBand struct {
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

func (s QueryTemplateListResponseBodyTemplateListTemplateVideoNarrowBand) String() string {
	return dara.Prettify(s)
}

func (s QueryTemplateListResponseBodyTemplateListTemplateVideoNarrowBand) GoString() string {
	return s.String()
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideoNarrowBand) GetAbrmax() *float32 {
	return s.Abrmax
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideoNarrowBand) GetMaxAbrRatio() *float32 {
	return s.MaxAbrRatio
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideoNarrowBand) GetVersion() *string {
	return s.Version
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideoNarrowBand) SetAbrmax(v float32) *QueryTemplateListResponseBodyTemplateListTemplateVideoNarrowBand {
	s.Abrmax = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideoNarrowBand) SetMaxAbrRatio(v float32) *QueryTemplateListResponseBodyTemplateListTemplateVideoNarrowBand {
	s.MaxAbrRatio = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideoNarrowBand) SetVersion(v string) *QueryTemplateListResponseBodyTemplateListTemplateVideoNarrowBand {
	s.Version = &v
	return s
}

func (s *QueryTemplateListResponseBodyTemplateListTemplateVideoNarrowBand) Validate() error {
	return dara.Validate(s)
}
