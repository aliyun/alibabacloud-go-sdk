// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddTemplateResponseBody
	GetRequestId() *string
	SetTemplate(v *AddTemplateResponseBodyTemplate) *AddTemplateResponseBody
	GetTemplate() *AddTemplateResponseBodyTemplate
}

type AddTemplateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// FA258E67-09B8-4EAA-8F33-BA567834A2C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the transcoding template.
	Template *AddTemplateResponseBodyTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
}

func (s AddTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *AddTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddTemplateResponseBody) GetTemplate() *AddTemplateResponseBodyTemplate {
	return s.Template
}

func (s *AddTemplateResponseBody) SetRequestId(v string) *AddTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddTemplateResponseBody) SetTemplate(v *AddTemplateResponseBodyTemplate) *AddTemplateResponseBody {
	s.Template = v
	return s
}

func (s *AddTemplateResponseBody) Validate() error {
	if s.Template != nil {
		if err := s.Template.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddTemplateResponseBodyTemplate struct {
	// The audio codec configurations.
	Audio *AddTemplateResponseBodyTemplateAudio `json:"Audio,omitempty" xml:"Audio,omitempty" type:"Struct"`
	// The container format settings.
	Container *AddTemplateResponseBodyTemplateContainer `json:"Container,omitempty" xml:"Container,omitempty" type:"Struct"`
	// The ID of the transcoding template. We recommend that you keep this ID for subsequent operation calls.
	//
	// example:
	//
	// 16f01ad6175e4230ac42bb5182cd****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The transmuxing settings.
	MuxConfig *AddTemplateResponseBodyTemplateMuxConfig `json:"MuxConfig,omitempty" xml:"MuxConfig,omitempty" type:"Struct"`
	// The name of the transcoding template.
	//
	// example:
	//
	// mps-example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the template. Valid values:
	//
	// 	- **Normal**: The template is normal.
	//
	// 	- **Deleted**: The template is deleted.
	//
	// example:
	//
	// Normal
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The general transcoding settings.
	TransConfig *AddTemplateResponseBodyTemplateTransConfig `json:"TransConfig,omitempty" xml:"TransConfig,omitempty" type:"Struct"`
	// The video codec configurations.
	Video *AddTemplateResponseBodyTemplateVideo `json:"Video,omitempty" xml:"Video,omitempty" type:"Struct"`
}

func (s AddTemplateResponseBodyTemplate) String() string {
	return dara.Prettify(s)
}

func (s AddTemplateResponseBodyTemplate) GoString() string {
	return s.String()
}

func (s *AddTemplateResponseBodyTemplate) GetAudio() *AddTemplateResponseBodyTemplateAudio {
	return s.Audio
}

func (s *AddTemplateResponseBodyTemplate) GetContainer() *AddTemplateResponseBodyTemplateContainer {
	return s.Container
}

func (s *AddTemplateResponseBodyTemplate) GetId() *string {
	return s.Id
}

func (s *AddTemplateResponseBodyTemplate) GetMuxConfig() *AddTemplateResponseBodyTemplateMuxConfig {
	return s.MuxConfig
}

func (s *AddTemplateResponseBodyTemplate) GetName() *string {
	return s.Name
}

func (s *AddTemplateResponseBodyTemplate) GetState() *string {
	return s.State
}

func (s *AddTemplateResponseBodyTemplate) GetTransConfig() *AddTemplateResponseBodyTemplateTransConfig {
	return s.TransConfig
}

func (s *AddTemplateResponseBodyTemplate) GetVideo() *AddTemplateResponseBodyTemplateVideo {
	return s.Video
}

func (s *AddTemplateResponseBodyTemplate) SetAudio(v *AddTemplateResponseBodyTemplateAudio) *AddTemplateResponseBodyTemplate {
	s.Audio = v
	return s
}

func (s *AddTemplateResponseBodyTemplate) SetContainer(v *AddTemplateResponseBodyTemplateContainer) *AddTemplateResponseBodyTemplate {
	s.Container = v
	return s
}

func (s *AddTemplateResponseBodyTemplate) SetId(v string) *AddTemplateResponseBodyTemplate {
	s.Id = &v
	return s
}

func (s *AddTemplateResponseBodyTemplate) SetMuxConfig(v *AddTemplateResponseBodyTemplateMuxConfig) *AddTemplateResponseBodyTemplate {
	s.MuxConfig = v
	return s
}

func (s *AddTemplateResponseBodyTemplate) SetName(v string) *AddTemplateResponseBodyTemplate {
	s.Name = &v
	return s
}

func (s *AddTemplateResponseBodyTemplate) SetState(v string) *AddTemplateResponseBodyTemplate {
	s.State = &v
	return s
}

func (s *AddTemplateResponseBodyTemplate) SetTransConfig(v *AddTemplateResponseBodyTemplateTransConfig) *AddTemplateResponseBodyTemplate {
	s.TransConfig = v
	return s
}

func (s *AddTemplateResponseBodyTemplate) SetVideo(v *AddTemplateResponseBodyTemplateVideo) *AddTemplateResponseBodyTemplate {
	s.Video = v
	return s
}

func (s *AddTemplateResponseBodyTemplate) Validate() error {
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

type AddTemplateResponseBodyTemplateAudio struct {
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
	// The codec profile of the audio. Valid values if the **Codec*	- parameter is set to **AAC**:
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
	// The level of the independent denoising algorithm.
	//
	// example:
	//
	// 5
	Qscale *string `json:"Qscale,omitempty" xml:"Qscale,omitempty"`
	// Indicates whether the audio stream is deleted.
	//
	// 	- **true**: The audio stream is deleted.
	//
	// 	- **false**: The audio stream is retained.
	//
	// 	- Default value: **false**.
	//
	// example:
	//
	// true
	Remove *string `json:"Remove,omitempty" xml:"Remove,omitempty"`
	// The sampling rate.
	//
	// 	- Unit: Hz.
	//
	// 	- Default value: **44100**.
	//
	// example:
	//
	// 44100
	Samplerate *string `json:"Samplerate,omitempty" xml:"Samplerate,omitempty"`
	// The volume control configurations
	Volume *AddTemplateResponseBodyTemplateAudioVolume `json:"Volume,omitempty" xml:"Volume,omitempty" type:"Struct"`
}

func (s AddTemplateResponseBodyTemplateAudio) String() string {
	return dara.Prettify(s)
}

func (s AddTemplateResponseBodyTemplateAudio) GoString() string {
	return s.String()
}

func (s *AddTemplateResponseBodyTemplateAudio) GetBitrate() *string {
	return s.Bitrate
}

func (s *AddTemplateResponseBodyTemplateAudio) GetChannels() *string {
	return s.Channels
}

func (s *AddTemplateResponseBodyTemplateAudio) GetCodec() *string {
	return s.Codec
}

func (s *AddTemplateResponseBodyTemplateAudio) GetProfile() *string {
	return s.Profile
}

func (s *AddTemplateResponseBodyTemplateAudio) GetQscale() *string {
	return s.Qscale
}

func (s *AddTemplateResponseBodyTemplateAudio) GetRemove() *string {
	return s.Remove
}

func (s *AddTemplateResponseBodyTemplateAudio) GetSamplerate() *string {
	return s.Samplerate
}

func (s *AddTemplateResponseBodyTemplateAudio) GetVolume() *AddTemplateResponseBodyTemplateAudioVolume {
	return s.Volume
}

func (s *AddTemplateResponseBodyTemplateAudio) SetBitrate(v string) *AddTemplateResponseBodyTemplateAudio {
	s.Bitrate = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateAudio) SetChannels(v string) *AddTemplateResponseBodyTemplateAudio {
	s.Channels = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateAudio) SetCodec(v string) *AddTemplateResponseBodyTemplateAudio {
	s.Codec = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateAudio) SetProfile(v string) *AddTemplateResponseBodyTemplateAudio {
	s.Profile = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateAudio) SetQscale(v string) *AddTemplateResponseBodyTemplateAudio {
	s.Qscale = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateAudio) SetRemove(v string) *AddTemplateResponseBodyTemplateAudio {
	s.Remove = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateAudio) SetSamplerate(v string) *AddTemplateResponseBodyTemplateAudio {
	s.Samplerate = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateAudio) SetVolume(v *AddTemplateResponseBodyTemplateAudioVolume) *AddTemplateResponseBodyTemplateAudio {
	s.Volume = v
	return s
}

func (s *AddTemplateResponseBodyTemplateAudio) Validate() error {
	if s.Volume != nil {
		if err := s.Volume.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddTemplateResponseBodyTemplateAudioVolume struct {
	// The output volume.
	//
	// This parameter takes effect only when the value of Method is dynamic.
	//
	// Unit: dB.
	//
	// Valid values: [-70,-5].
	//
	// Default value: -6.
	//
	// example:
	//
	// -6
	IntegratedLoudnessTarget *string `json:"IntegratedLoudnessTarget,omitempty" xml:"IntegratedLoudnessTarget,omitempty"`
	// The volume adjustment range.
	//
	// 	- Default value: **-20**.
	//
	// 	- Unit: dB.
	//
	// example:
	//
	// -20
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The range of the volume relative to the output volume.
	//
	// This parameter takes effect only when the value of Method is dynamic.
	//
	// Unit: dB.
	//
	// Valid values: [1,20].
	//
	// Default value: 8.
	//
	// example:
	//
	// 8
	LoudnessRangeTarget *string `json:"LoudnessRangeTarget,omitempty" xml:"LoudnessRangeTarget,omitempty"`
	// The volume adjustment method. Valid values:
	//
	// 	- **auto**: The volume is automatically adjusted.
	//
	// 	- **dynamic**: The volume is dynamically adjusted.
	//
	// 	- **linear**: The volume is linearly adjusted.
	//
	// example:
	//
	// auto
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The volume adjustment coefficient.
	//
	// This parameter takes effect only when the value of Method is adaptive.
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
	// This parameter takes effect only when the value of Method is dynamic.
	//
	// Unit: dB.
	//
	// Valid values: [-9,0].
	//
	// Default value: -1.
	//
	// example:
	//
	// 0
	TruePeak *string `json:"TruePeak,omitempty" xml:"TruePeak,omitempty"`
}

func (s AddTemplateResponseBodyTemplateAudioVolume) String() string {
	return dara.Prettify(s)
}

func (s AddTemplateResponseBodyTemplateAudioVolume) GoString() string {
	return s.String()
}

func (s *AddTemplateResponseBodyTemplateAudioVolume) GetIntegratedLoudnessTarget() *string {
	return s.IntegratedLoudnessTarget
}

func (s *AddTemplateResponseBodyTemplateAudioVolume) GetLevel() *string {
	return s.Level
}

func (s *AddTemplateResponseBodyTemplateAudioVolume) GetLoudnessRangeTarget() *string {
	return s.LoudnessRangeTarget
}

func (s *AddTemplateResponseBodyTemplateAudioVolume) GetMethod() *string {
	return s.Method
}

func (s *AddTemplateResponseBodyTemplateAudioVolume) GetPeakLevel() *string {
	return s.PeakLevel
}

func (s *AddTemplateResponseBodyTemplateAudioVolume) GetTruePeak() *string {
	return s.TruePeak
}

func (s *AddTemplateResponseBodyTemplateAudioVolume) SetIntegratedLoudnessTarget(v string) *AddTemplateResponseBodyTemplateAudioVolume {
	s.IntegratedLoudnessTarget = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateAudioVolume) SetLevel(v string) *AddTemplateResponseBodyTemplateAudioVolume {
	s.Level = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateAudioVolume) SetLoudnessRangeTarget(v string) *AddTemplateResponseBodyTemplateAudioVolume {
	s.LoudnessRangeTarget = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateAudioVolume) SetMethod(v string) *AddTemplateResponseBodyTemplateAudioVolume {
	s.Method = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateAudioVolume) SetPeakLevel(v string) *AddTemplateResponseBodyTemplateAudioVolume {
	s.PeakLevel = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateAudioVolume) SetTruePeak(v string) *AddTemplateResponseBodyTemplateAudioVolume {
	s.TruePeak = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateAudioVolume) Validate() error {
	return dara.Validate(s)
}

type AddTemplateResponseBodyTemplateContainer struct {
	// The container format.
	//
	// example:
	//
	// mp4
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
}

func (s AddTemplateResponseBodyTemplateContainer) String() string {
	return dara.Prettify(s)
}

func (s AddTemplateResponseBodyTemplateContainer) GoString() string {
	return s.String()
}

func (s *AddTemplateResponseBodyTemplateContainer) GetFormat() *string {
	return s.Format
}

func (s *AddTemplateResponseBodyTemplateContainer) SetFormat(v string) *AddTemplateResponseBodyTemplateContainer {
	s.Format = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateContainer) Validate() error {
	return dara.Validate(s)
}

type AddTemplateResponseBodyTemplateMuxConfig struct {
	// The transmuxing settings for GIF.
	Gif *AddTemplateResponseBodyTemplateMuxConfigGif `json:"Gif,omitempty" xml:"Gif,omitempty" type:"Struct"`
	// The segment settings.
	Segment *AddTemplateResponseBodyTemplateMuxConfigSegment `json:"Segment,omitempty" xml:"Segment,omitempty" type:"Struct"`
	// The transmuxing settings for WebP.
	Webp *AddTemplateResponseBodyTemplateMuxConfigWebp `json:"Webp,omitempty" xml:"Webp,omitempty" type:"Struct"`
}

func (s AddTemplateResponseBodyTemplateMuxConfig) String() string {
	return dara.Prettify(s)
}

func (s AddTemplateResponseBodyTemplateMuxConfig) GoString() string {
	return s.String()
}

func (s *AddTemplateResponseBodyTemplateMuxConfig) GetGif() *AddTemplateResponseBodyTemplateMuxConfigGif {
	return s.Gif
}

func (s *AddTemplateResponseBodyTemplateMuxConfig) GetSegment() *AddTemplateResponseBodyTemplateMuxConfigSegment {
	return s.Segment
}

func (s *AddTemplateResponseBodyTemplateMuxConfig) GetWebp() *AddTemplateResponseBodyTemplateMuxConfigWebp {
	return s.Webp
}

func (s *AddTemplateResponseBodyTemplateMuxConfig) SetGif(v *AddTemplateResponseBodyTemplateMuxConfigGif) *AddTemplateResponseBodyTemplateMuxConfig {
	s.Gif = v
	return s
}

func (s *AddTemplateResponseBodyTemplateMuxConfig) SetSegment(v *AddTemplateResponseBodyTemplateMuxConfigSegment) *AddTemplateResponseBodyTemplateMuxConfig {
	s.Segment = v
	return s
}

func (s *AddTemplateResponseBodyTemplateMuxConfig) SetWebp(v *AddTemplateResponseBodyTemplateMuxConfigWebp) *AddTemplateResponseBodyTemplateMuxConfig {
	s.Webp = v
	return s
}

func (s *AddTemplateResponseBodyTemplateMuxConfig) Validate() error {
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

type AddTemplateResponseBodyTemplateMuxConfigGif struct {
	// The color dithering algorithm of the palette. Valid values: sierra and bayer.
	//
	// example:
	//
	// sierra
	DitherMode *string `json:"DitherMode,omitempty" xml:"DitherMode,omitempty"`
	// The duration for which the final frame is paused. Unit: centiseconds.
	//
	// example:
	//
	// 0
	FinalDelay *string `json:"FinalDelay,omitempty" xml:"FinalDelay,omitempty"`
	// Indicates whether the custom palette is used.
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

func (s AddTemplateResponseBodyTemplateMuxConfigGif) String() string {
	return dara.Prettify(s)
}

func (s AddTemplateResponseBodyTemplateMuxConfigGif) GoString() string {
	return s.String()
}

func (s *AddTemplateResponseBodyTemplateMuxConfigGif) GetDitherMode() *string {
	return s.DitherMode
}

func (s *AddTemplateResponseBodyTemplateMuxConfigGif) GetFinalDelay() *string {
	return s.FinalDelay
}

func (s *AddTemplateResponseBodyTemplateMuxConfigGif) GetIsCustomPalette() *string {
	return s.IsCustomPalette
}

func (s *AddTemplateResponseBodyTemplateMuxConfigGif) GetLoop() *string {
	return s.Loop
}

func (s *AddTemplateResponseBodyTemplateMuxConfigGif) SetDitherMode(v string) *AddTemplateResponseBodyTemplateMuxConfigGif {
	s.DitherMode = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateMuxConfigGif) SetFinalDelay(v string) *AddTemplateResponseBodyTemplateMuxConfigGif {
	s.FinalDelay = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateMuxConfigGif) SetIsCustomPalette(v string) *AddTemplateResponseBodyTemplateMuxConfigGif {
	s.IsCustomPalette = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateMuxConfigGif) SetLoop(v string) *AddTemplateResponseBodyTemplateMuxConfigGif {
	s.Loop = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateMuxConfigGif) Validate() error {
	return dara.Validate(s)
}

type AddTemplateResponseBodyTemplateMuxConfigSegment struct {
	// The length of the segment. Unit: seconds.
	//
	// example:
	//
	// 10
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s AddTemplateResponseBodyTemplateMuxConfigSegment) String() string {
	return dara.Prettify(s)
}

func (s AddTemplateResponseBodyTemplateMuxConfigSegment) GoString() string {
	return s.String()
}

func (s *AddTemplateResponseBodyTemplateMuxConfigSegment) GetDuration() *string {
	return s.Duration
}

func (s *AddTemplateResponseBodyTemplateMuxConfigSegment) SetDuration(v string) *AddTemplateResponseBodyTemplateMuxConfigSegment {
	s.Duration = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateMuxConfigSegment) Validate() error {
	return dara.Validate(s)
}

type AddTemplateResponseBodyTemplateMuxConfigWebp struct {
	// The loop count.
	//
	// example:
	//
	// 0
	Loop *string `json:"Loop,omitempty" xml:"Loop,omitempty"`
}

func (s AddTemplateResponseBodyTemplateMuxConfigWebp) String() string {
	return dara.Prettify(s)
}

func (s AddTemplateResponseBodyTemplateMuxConfigWebp) GoString() string {
	return s.String()
}

func (s *AddTemplateResponseBodyTemplateMuxConfigWebp) GetLoop() *string {
	return s.Loop
}

func (s *AddTemplateResponseBodyTemplateMuxConfigWebp) SetLoop(v string) *AddTemplateResponseBodyTemplateMuxConfigWebp {
	s.Loop = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateMuxConfigWebp) Validate() error {
	return dara.Validate(s)
}

type AddTemplateResponseBodyTemplateTransConfig struct {
	// The method of resolution adjustment. Default value: **none**. Valid values:
	//
	// 	- **rescale**: The input video is rescaled.
	//
	// 	- **crop**: The input video is cropped.
	//
	// 	- **none**: No change is made.
	//
	// example:
	//
	// rescale
	AdjDarMethod *string `json:"AdjDarMethod,omitempty" xml:"AdjDarMethod,omitempty"`
	// Indicates whether the audio bitrate is checked.
	//
	// If this feature is enabled and the system detects that the audio bitrate of the output file is greater than that of the input file, the audio bitrate of the input file is retained after transcoding.
	//
	// 	- **true**: The audio bitrate is checked.
	//
	// 	- **false**: The audio bitrate is not checked.
	//
	// 	- Default value: **false**.
	//
	// example:
	//
	// true
	IsCheckAudioBitrate *string `json:"IsCheckAudioBitrate,omitempty" xml:"IsCheckAudioBitrate,omitempty"`
	// Indicates whether the audio bitrate is checked. If this feature is enabled and the system detects that the audio bitrate of the output file is higher than that of the input file, the input file is not transcoded. This parameter has a higher priority than the **IsCheckAudioBitrate*	- parameter. Valid values:
	//
	// 	- **true**: The audio bitrate is checked. In this case, if the audio bitrate of the output file is higher than that of the input file, the input file is not transcoded.
	//
	// 	- **false**: The audio bitrate is not checked.
	//
	// 	- Default value: **false**.
	//
	// example:
	//
	// true
	IsCheckAudioBitrateFail *string `json:"IsCheckAudioBitrateFail,omitempty" xml:"IsCheckAudioBitrateFail,omitempty"`
	// Indicates whether the resolution is checked.
	//
	// 	- **true**: The resolution is checked.
	//
	// 	- **false**: The resolution is not checked.
	//
	// 	- Default value: **false**.
	//
	// > If this feature is enabled and the system detects that the resolution of the output file is higher than that of the input file based on the width or height, the resolution of the input file is retained after transcoding.
	//
	// example:
	//
	// true
	IsCheckReso *string `json:"IsCheckReso,omitempty" xml:"IsCheckReso,omitempty"`
	// Indicates whether the resolution is checked.
	//
	// 	- **true**: The resolution is checked.
	//
	// 	- **false**: The resolution is not checked.
	//
	// 	- Default value: **false**.
	//
	// > If this feature is enabled and the system detects that the resolution of the output file is higher than that of the input file based on the width or height, an error that indicates a transcoding failure is returned.
	//
	// example:
	//
	// true
	IsCheckResoFail *string `json:"IsCheckResoFail,omitempty" xml:"IsCheckResoFail,omitempty"`
	// Indicates whether the video bitrate is checked.
	//
	// 	- **true**: The video bitrate is checked.
	//
	// 	- **false**: The video bitrate is not checked.
	//
	// 	- Default value: **false**.
	//
	// > If this feature is enabled and the system detects that the video bitrate of the output file is greater than that of the input file, the video bitrate of the input file is retained after transcoding.
	//
	// example:
	//
	// true
	IsCheckVideoBitrate *string `json:"IsCheckVideoBitrate,omitempty" xml:"IsCheckVideoBitrate,omitempty"`
	// Indicates whether the video bitrate is checked. If this feature is enabled and the system detects that the video bitrate of the output file is higher than that of the input file, the input file is not transcoded. This parameter has a higher priority than the IsCheckVideoBitrate parameter.
	//
	// 	- **true**: The video bitrate is checked. In this case, if the video bitrate of the output file is higher than that of the input file, the input file is not transcoded.
	//
	// 	- **false**: The video bitrate is not checked.
	//
	// 	- Default value: **false**.
	//
	// example:
	//
	// true
	IsCheckVideoBitrateFail *string `json:"IsCheckVideoBitrateFail,omitempty" xml:"IsCheckVideoBitrateFail,omitempty"`
	// The transcoding mode. Valid values:
	//
	// 	- **onepass**
	//
	// 	- **twopass**
	//
	// 	- **CBR**
	//
	// 	- Default value: **onepass**.
	//
	// example:
	//
	// onepass
	TransMode *string `json:"TransMode,omitempty" xml:"TransMode,omitempty"`
}

func (s AddTemplateResponseBodyTemplateTransConfig) String() string {
	return dara.Prettify(s)
}

func (s AddTemplateResponseBodyTemplateTransConfig) GoString() string {
	return s.String()
}

func (s *AddTemplateResponseBodyTemplateTransConfig) GetAdjDarMethod() *string {
	return s.AdjDarMethod
}

func (s *AddTemplateResponseBodyTemplateTransConfig) GetIsCheckAudioBitrate() *string {
	return s.IsCheckAudioBitrate
}

func (s *AddTemplateResponseBodyTemplateTransConfig) GetIsCheckAudioBitrateFail() *string {
	return s.IsCheckAudioBitrateFail
}

func (s *AddTemplateResponseBodyTemplateTransConfig) GetIsCheckReso() *string {
	return s.IsCheckReso
}

func (s *AddTemplateResponseBodyTemplateTransConfig) GetIsCheckResoFail() *string {
	return s.IsCheckResoFail
}

func (s *AddTemplateResponseBodyTemplateTransConfig) GetIsCheckVideoBitrate() *string {
	return s.IsCheckVideoBitrate
}

func (s *AddTemplateResponseBodyTemplateTransConfig) GetIsCheckVideoBitrateFail() *string {
	return s.IsCheckVideoBitrateFail
}

func (s *AddTemplateResponseBodyTemplateTransConfig) GetTransMode() *string {
	return s.TransMode
}

func (s *AddTemplateResponseBodyTemplateTransConfig) SetAdjDarMethod(v string) *AddTemplateResponseBodyTemplateTransConfig {
	s.AdjDarMethod = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateTransConfig) SetIsCheckAudioBitrate(v string) *AddTemplateResponseBodyTemplateTransConfig {
	s.IsCheckAudioBitrate = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateTransConfig) SetIsCheckAudioBitrateFail(v string) *AddTemplateResponseBodyTemplateTransConfig {
	s.IsCheckAudioBitrateFail = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateTransConfig) SetIsCheckReso(v string) *AddTemplateResponseBodyTemplateTransConfig {
	s.IsCheckReso = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateTransConfig) SetIsCheckResoFail(v string) *AddTemplateResponseBodyTemplateTransConfig {
	s.IsCheckResoFail = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateTransConfig) SetIsCheckVideoBitrate(v string) *AddTemplateResponseBodyTemplateTransConfig {
	s.IsCheckVideoBitrate = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateTransConfig) SetIsCheckVideoBitrateFail(v string) *AddTemplateResponseBodyTemplateTransConfig {
	s.IsCheckVideoBitrateFail = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateTransConfig) SetTransMode(v string) *AddTemplateResponseBodyTemplateTransConfig {
	s.TransMode = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateTransConfig) Validate() error {
	return dara.Validate(s)
}

type AddTemplateResponseBodyTemplateVideo struct {
	// The bitrate of the output video. Unit: Kbit/s.
	//
	// example:
	//
	// 500
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The bitrate range of the video.
	BitrateBnd *AddTemplateResponseBodyTemplateVideoBitrateBnd `json:"BitrateBnd,omitempty" xml:"BitrateBnd,omitempty" type:"Struct"`
	// The size of the buffer.
	//
	// 	- Default value: **6000**.
	//
	// 	- Unit: KB.
	//
	// example:
	//
	// 6000
	Bufsize *string `json:"Bufsize,omitempty" xml:"Bufsize,omitempty"`
	// The video codec. Valid values: H.264, H.265, GIF, and WebP. Default value: **H.264**.
	//
	// example:
	//
	// H.264
	Codec *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	// The constant rate factor. Default value if the video codec is set to H.264: **23**. Default value if the video codec is set to H.265: **26**.
	//
	// > If this parameter is specified, the setting of the Bitrate parameter becomes invalid.
	//
	// example:
	//
	// 15
	Crf *string `json:"Crf,omitempty" xml:"Crf,omitempty"`
	// The method of video cropping. Valid values:
	//
	// 	- **border**: automatically detects and removes borders.
	//
	// 	- **Value in the format of width:height:left:top**: crops the video image based on the custom settings. Example: 1280:800:0:140.
	//
	// example:
	//
	// border
	Crop *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	// The level of quality control on the video.
	//
	// example:
	//
	// 10
	Degrain *string `json:"Degrain,omitempty" xml:"Degrain,omitempty"`
	// The frame rate. Default value: the frame rate of the input file. The value is 60 if the frame rate of the input file exceeds 60. Unit: frames per second.
	//
	// example:
	//
	// 25
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The GOP size. The GOP size can be the maximum interval of keyframes or the maximum number of frames in a frame group. If the maximum interval is specified, the value contains the unit (s). If the maximum number of frames is specified, the value does not contain a unit. Default value: **10s**.
	//
	// example:
	//
	// 10s
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
	// 	- Default value: the height of the input video.
	//
	// example:
	//
	// 800
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// Indicates whether the auto-rotate screen feature is enabled. Default value: **false**. Valid values:
	//
	// 	- **true**: The auto-rotate screen feature is enabled.
	//
	// 	- **false**: The auto-rotate screen feature is disabled.
	//
	// > If this feature is enabled, the width of the output video corresponds to the long side of the input video, which is the height of the input video in portrait mode. The height of the output video corresponds to the short side of the input video, which is the width of the input video in portrait mode.
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
	NarrowBand *AddTemplateResponseBodyTemplateVideoNarrowBand `json:"NarrowBand,omitempty" xml:"NarrowBand,omitempty" type:"Struct"`
	// The black borders to be added to the video. The value is in the width:height:left:top format.
	//
	// example:
	//
	// 1280:800:0:140
	Pad *string `json:"Pad,omitempty" xml:"Pad,omitempty"`
	// The pixel format. Standard pixel formats such as yuv420p and yuvj420p are supported. The default pixel format can be **yuv420p*	- or the pixel format of the input video.
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
	// > This parameter is valid only if the Codec parameter is set to H.264.
	//
	// example:
	//
	// fast
	Preset *string `json:"Preset,omitempty" xml:"Preset,omitempty"`
	// The codec profile.
	//
	// 	- **baseline**: suitable for mobile devices
	//
	// 	- **main**: suitable for standard-definition devices
	//
	// 	- **high**: suitable for high-definition devices
	//
	// 	- Default value: **high**.
	//
	// If multiple definitions are available, we recommend that you set this parameter to baseline for the lowest definition to ensure normal playback on low-end devices. Set this parameter to main or high for other definitions.
	//
	// > This parameter is valid only if the Codec parameter is set to H.264.
	//
	// example:
	//
	// high
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The level of the independent denoising algorithm.
	//
	// example:
	//
	// 1
	Qscale *string `json:"Qscale,omitempty" xml:"Qscale,omitempty"`
	// Indicates whether the video stream is deleted.
	//
	// 	- **true**: The video stream is deleted.
	//
	// 	- **false**: The video stream is retained.
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
	// 0
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
	// 	- Default value: the width of the input video.****
	//
	// 	- Unit: pixel.
	//
	// example:
	//
	// 256
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s AddTemplateResponseBodyTemplateVideo) String() string {
	return dara.Prettify(s)
}

func (s AddTemplateResponseBodyTemplateVideo) GoString() string {
	return s.String()
}

func (s *AddTemplateResponseBodyTemplateVideo) GetBitrate() *string {
	return s.Bitrate
}

func (s *AddTemplateResponseBodyTemplateVideo) GetBitrateBnd() *AddTemplateResponseBodyTemplateVideoBitrateBnd {
	return s.BitrateBnd
}

func (s *AddTemplateResponseBodyTemplateVideo) GetBufsize() *string {
	return s.Bufsize
}

func (s *AddTemplateResponseBodyTemplateVideo) GetCodec() *string {
	return s.Codec
}

func (s *AddTemplateResponseBodyTemplateVideo) GetCrf() *string {
	return s.Crf
}

func (s *AddTemplateResponseBodyTemplateVideo) GetCrop() *string {
	return s.Crop
}

func (s *AddTemplateResponseBodyTemplateVideo) GetDegrain() *string {
	return s.Degrain
}

func (s *AddTemplateResponseBodyTemplateVideo) GetFps() *string {
	return s.Fps
}

func (s *AddTemplateResponseBodyTemplateVideo) GetGop() *string {
	return s.Gop
}

func (s *AddTemplateResponseBodyTemplateVideo) GetHdr2sdr() *string {
	return s.Hdr2sdr
}

func (s *AddTemplateResponseBodyTemplateVideo) GetHeight() *string {
	return s.Height
}

func (s *AddTemplateResponseBodyTemplateVideo) GetLongShortMode() *string {
	return s.LongShortMode
}

func (s *AddTemplateResponseBodyTemplateVideo) GetMaxFps() *string {
	return s.MaxFps
}

func (s *AddTemplateResponseBodyTemplateVideo) GetMaxrate() *string {
	return s.Maxrate
}

func (s *AddTemplateResponseBodyTemplateVideo) GetNarrowBand() *AddTemplateResponseBodyTemplateVideoNarrowBand {
	return s.NarrowBand
}

func (s *AddTemplateResponseBodyTemplateVideo) GetPad() *string {
	return s.Pad
}

func (s *AddTemplateResponseBodyTemplateVideo) GetPixFmt() *string {
	return s.PixFmt
}

func (s *AddTemplateResponseBodyTemplateVideo) GetPreset() *string {
	return s.Preset
}

func (s *AddTemplateResponseBodyTemplateVideo) GetProfile() *string {
	return s.Profile
}

func (s *AddTemplateResponseBodyTemplateVideo) GetQscale() *string {
	return s.Qscale
}

func (s *AddTemplateResponseBodyTemplateVideo) GetRemove() *string {
	return s.Remove
}

func (s *AddTemplateResponseBodyTemplateVideo) GetResoPriority() *string {
	return s.ResoPriority
}

func (s *AddTemplateResponseBodyTemplateVideo) GetScanMode() *string {
	return s.ScanMode
}

func (s *AddTemplateResponseBodyTemplateVideo) GetWidth() *string {
	return s.Width
}

func (s *AddTemplateResponseBodyTemplateVideo) SetBitrate(v string) *AddTemplateResponseBodyTemplateVideo {
	s.Bitrate = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateVideo) SetBitrateBnd(v *AddTemplateResponseBodyTemplateVideoBitrateBnd) *AddTemplateResponseBodyTemplateVideo {
	s.BitrateBnd = v
	return s
}

func (s *AddTemplateResponseBodyTemplateVideo) SetBufsize(v string) *AddTemplateResponseBodyTemplateVideo {
	s.Bufsize = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateVideo) SetCodec(v string) *AddTemplateResponseBodyTemplateVideo {
	s.Codec = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateVideo) SetCrf(v string) *AddTemplateResponseBodyTemplateVideo {
	s.Crf = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateVideo) SetCrop(v string) *AddTemplateResponseBodyTemplateVideo {
	s.Crop = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateVideo) SetDegrain(v string) *AddTemplateResponseBodyTemplateVideo {
	s.Degrain = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateVideo) SetFps(v string) *AddTemplateResponseBodyTemplateVideo {
	s.Fps = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateVideo) SetGop(v string) *AddTemplateResponseBodyTemplateVideo {
	s.Gop = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateVideo) SetHdr2sdr(v string) *AddTemplateResponseBodyTemplateVideo {
	s.Hdr2sdr = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateVideo) SetHeight(v string) *AddTemplateResponseBodyTemplateVideo {
	s.Height = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateVideo) SetLongShortMode(v string) *AddTemplateResponseBodyTemplateVideo {
	s.LongShortMode = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateVideo) SetMaxFps(v string) *AddTemplateResponseBodyTemplateVideo {
	s.MaxFps = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateVideo) SetMaxrate(v string) *AddTemplateResponseBodyTemplateVideo {
	s.Maxrate = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateVideo) SetNarrowBand(v *AddTemplateResponseBodyTemplateVideoNarrowBand) *AddTemplateResponseBodyTemplateVideo {
	s.NarrowBand = v
	return s
}

func (s *AddTemplateResponseBodyTemplateVideo) SetPad(v string) *AddTemplateResponseBodyTemplateVideo {
	s.Pad = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateVideo) SetPixFmt(v string) *AddTemplateResponseBodyTemplateVideo {
	s.PixFmt = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateVideo) SetPreset(v string) *AddTemplateResponseBodyTemplateVideo {
	s.Preset = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateVideo) SetProfile(v string) *AddTemplateResponseBodyTemplateVideo {
	s.Profile = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateVideo) SetQscale(v string) *AddTemplateResponseBodyTemplateVideo {
	s.Qscale = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateVideo) SetRemove(v string) *AddTemplateResponseBodyTemplateVideo {
	s.Remove = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateVideo) SetResoPriority(v string) *AddTemplateResponseBodyTemplateVideo {
	s.ResoPriority = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateVideo) SetScanMode(v string) *AddTemplateResponseBodyTemplateVideo {
	s.ScanMode = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateVideo) SetWidth(v string) *AddTemplateResponseBodyTemplateVideo {
	s.Width = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateVideo) Validate() error {
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

type AddTemplateResponseBodyTemplateVideoBitrateBnd struct {
	// The maximum bitrate.
	//
	// example:
	//
	// 1500
	Max *string `json:"Max,omitempty" xml:"Max,omitempty"`
	// The minimum bitrate.
	//
	// example:
	//
	// 800
	Min *string `json:"Min,omitempty" xml:"Min,omitempty"`
}

func (s AddTemplateResponseBodyTemplateVideoBitrateBnd) String() string {
	return dara.Prettify(s)
}

func (s AddTemplateResponseBodyTemplateVideoBitrateBnd) GoString() string {
	return s.String()
}

func (s *AddTemplateResponseBodyTemplateVideoBitrateBnd) GetMax() *string {
	return s.Max
}

func (s *AddTemplateResponseBodyTemplateVideoBitrateBnd) GetMin() *string {
	return s.Min
}

func (s *AddTemplateResponseBodyTemplateVideoBitrateBnd) SetMax(v string) *AddTemplateResponseBodyTemplateVideoBitrateBnd {
	s.Max = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateVideoBitrateBnd) SetMin(v string) *AddTemplateResponseBodyTemplateVideoBitrateBnd {
	s.Min = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateVideoBitrateBnd) Validate() error {
	return dara.Validate(s)
}

type AddTemplateResponseBodyTemplateVideoNarrowBand struct {
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

func (s AddTemplateResponseBodyTemplateVideoNarrowBand) String() string {
	return dara.Prettify(s)
}

func (s AddTemplateResponseBodyTemplateVideoNarrowBand) GoString() string {
	return s.String()
}

func (s *AddTemplateResponseBodyTemplateVideoNarrowBand) GetAbrmax() *float32 {
	return s.Abrmax
}

func (s *AddTemplateResponseBodyTemplateVideoNarrowBand) GetMaxAbrRatio() *float32 {
	return s.MaxAbrRatio
}

func (s *AddTemplateResponseBodyTemplateVideoNarrowBand) GetVersion() *string {
	return s.Version
}

func (s *AddTemplateResponseBodyTemplateVideoNarrowBand) SetAbrmax(v float32) *AddTemplateResponseBodyTemplateVideoNarrowBand {
	s.Abrmax = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateVideoNarrowBand) SetMaxAbrRatio(v float32) *AddTemplateResponseBodyTemplateVideoNarrowBand {
	s.MaxAbrRatio = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateVideoNarrowBand) SetVersion(v string) *AddTemplateResponseBodyTemplateVideoNarrowBand {
	s.Version = &v
	return s
}

func (s *AddTemplateResponseBodyTemplateVideoNarrowBand) Validate() error {
	return dara.Validate(s)
}
