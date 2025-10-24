// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTemplateResponseBody
	GetRequestId() *string
	SetTemplate(v *UpdateTemplateResponseBodyTemplate) *UpdateTemplateResponseBody
	GetTemplate() *UpdateTemplateResponseBodyTemplate
}

type UpdateTemplateResponseBody struct {
	// The type of the transcoding template.
	//
	// example:
	//
	// 5E4FB22E-B9EA-4E24-8FFC-B407EA71QW21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The type of the transcoding template.
	Template *UpdateTemplateResponseBodyTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
}

func (s UpdateTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTemplateResponseBody) GetTemplate() *UpdateTemplateResponseBodyTemplate {
	return s.Template
}

func (s *UpdateTemplateResponseBody) SetRequestId(v string) *UpdateTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTemplateResponseBody) SetTemplate(v *UpdateTemplateResponseBodyTemplate) *UpdateTemplateResponseBody {
	s.Template = v
	return s
}

func (s *UpdateTemplateResponseBody) Validate() error {
	if s.Template != nil {
		if err := s.Template.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateTemplateResponseBodyTemplate struct {
	// The audio codec settings.
	Audio *UpdateTemplateResponseBodyTemplateAudio `json:"Audio,omitempty" xml:"Audio,omitempty" type:"Struct"`
	// The container format.
	Container *UpdateTemplateResponseBodyTemplateContainer `json:"Container,omitempty" xml:"Container,omitempty" type:"Struct"`
	// The container configurations.
	//
	// example:
	//
	// 16f01ad6175e4230ac42bb5182cd****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The transmuxing configurations for WebP.
	MuxConfig *UpdateTemplateResponseBodyTemplateMuxConfig `json:"MuxConfig,omitempty" xml:"MuxConfig,omitempty" type:"Struct"`
	// The audio codec configurations.
	//
	// example:
	//
	// MPS-example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The transmuxing configurations.
	//
	// example:
	//
	// Normal
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// Indicates whether the audio bitrate is checked. If the bitrate of the output audio is greater than the bitrate of the input audio, the bitrate of the input audio is retained after transcoding. In this case, the specified audio bitrate does not take effect. This parameter has a lower priority than the IsCheckAudioBitrateFail parameter. Valid values:
	//
	// 	- **true**: The audio bitrate is checked.
	//
	// 	- **false**: The audio bitrate is not checked.
	//
	// 	- Default value:
	//
	//     	- If the parameter is left empty and the codec of the output audio is different from that of the input audio, the default value is false.
	//
	//     	- If the parameter is left empty and the codec of the output audio is the same as that of the input audio, the default value is true.
	TransConfig *UpdateTemplateResponseBodyTemplateTransConfig `json:"TransConfig,omitempty" xml:"TransConfig,omitempty" type:"Struct"`
	// The video codec configurations.
	Video *UpdateTemplateResponseBodyTemplateVideo `json:"Video,omitempty" xml:"Video,omitempty" type:"Struct"`
}

func (s UpdateTemplateResponseBodyTemplate) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateResponseBodyTemplate) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponseBodyTemplate) GetAudio() *UpdateTemplateResponseBodyTemplateAudio {
	return s.Audio
}

func (s *UpdateTemplateResponseBodyTemplate) GetContainer() *UpdateTemplateResponseBodyTemplateContainer {
	return s.Container
}

func (s *UpdateTemplateResponseBodyTemplate) GetId() *string {
	return s.Id
}

func (s *UpdateTemplateResponseBodyTemplate) GetMuxConfig() *UpdateTemplateResponseBodyTemplateMuxConfig {
	return s.MuxConfig
}

func (s *UpdateTemplateResponseBodyTemplate) GetName() *string {
	return s.Name
}

func (s *UpdateTemplateResponseBodyTemplate) GetState() *string {
	return s.State
}

func (s *UpdateTemplateResponseBodyTemplate) GetTransConfig() *UpdateTemplateResponseBodyTemplateTransConfig {
	return s.TransConfig
}

func (s *UpdateTemplateResponseBodyTemplate) GetVideo() *UpdateTemplateResponseBodyTemplateVideo {
	return s.Video
}

func (s *UpdateTemplateResponseBodyTemplate) SetAudio(v *UpdateTemplateResponseBodyTemplateAudio) *UpdateTemplateResponseBodyTemplate {
	s.Audio = v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetContainer(v *UpdateTemplateResponseBodyTemplateContainer) *UpdateTemplateResponseBodyTemplate {
	s.Container = v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetId(v string) *UpdateTemplateResponseBodyTemplate {
	s.Id = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetMuxConfig(v *UpdateTemplateResponseBodyTemplateMuxConfig) *UpdateTemplateResponseBodyTemplate {
	s.MuxConfig = v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetName(v string) *UpdateTemplateResponseBodyTemplate {
	s.Name = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetState(v string) *UpdateTemplateResponseBodyTemplate {
	s.State = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetTransConfig(v *UpdateTemplateResponseBodyTemplateTransConfig) *UpdateTemplateResponseBodyTemplate {
	s.TransConfig = v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetVideo(v *UpdateTemplateResponseBodyTemplateVideo) *UpdateTemplateResponseBodyTemplate {
	s.Video = v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) Validate() error {
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

type UpdateTemplateResponseBodyTemplateAudio struct {
	// The ID of the transcoding template.
	//
	// example:
	//
	// 500
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The audio bitrate of the output file.
	//
	// 	- Valid values: 8 to 1000.****
	//
	// 	- Unit: Kbit/s.
	//
	// 	- Default value: **128**.
	//
	// example:
	//
	// 2
	Channels *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	// The sampling rate.
	//
	// 	- Unit: Hz.
	//
	// 	- Default value: **44100**.
	//
	// example:
	//
	// aac
	Codec *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
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
	// aac_low
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The number of sound channels. Default value: **2**.
	//
	// example:
	//
	// 1
	Qscale *string `json:"Qscale,omitempty" xml:"Qscale,omitempty"`
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
	// false
	Remove *string `json:"Remove,omitempty" xml:"Remove,omitempty"`
	// The level of the independent denoising algorithm.
	//
	// example:
	//
	// 44100
	Samplerate *string `json:"Samplerate,omitempty" xml:"Samplerate,omitempty"`
	// The volume control configurations.
	Volume *UpdateTemplateResponseBodyTemplateAudioVolume `json:"Volume,omitempty" xml:"Volume,omitempty" type:"Struct"`
}

func (s UpdateTemplateResponseBodyTemplateAudio) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateResponseBodyTemplateAudio) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponseBodyTemplateAudio) GetBitrate() *string {
	return s.Bitrate
}

func (s *UpdateTemplateResponseBodyTemplateAudio) GetChannels() *string {
	return s.Channels
}

func (s *UpdateTemplateResponseBodyTemplateAudio) GetCodec() *string {
	return s.Codec
}

func (s *UpdateTemplateResponseBodyTemplateAudio) GetProfile() *string {
	return s.Profile
}

func (s *UpdateTemplateResponseBodyTemplateAudio) GetQscale() *string {
	return s.Qscale
}

func (s *UpdateTemplateResponseBodyTemplateAudio) GetRemove() *string {
	return s.Remove
}

func (s *UpdateTemplateResponseBodyTemplateAudio) GetSamplerate() *string {
	return s.Samplerate
}

func (s *UpdateTemplateResponseBodyTemplateAudio) GetVolume() *UpdateTemplateResponseBodyTemplateAudioVolume {
	return s.Volume
}

func (s *UpdateTemplateResponseBodyTemplateAudio) SetBitrate(v string) *UpdateTemplateResponseBodyTemplateAudio {
	s.Bitrate = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateAudio) SetChannels(v string) *UpdateTemplateResponseBodyTemplateAudio {
	s.Channels = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateAudio) SetCodec(v string) *UpdateTemplateResponseBodyTemplateAudio {
	s.Codec = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateAudio) SetProfile(v string) *UpdateTemplateResponseBodyTemplateAudio {
	s.Profile = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateAudio) SetQscale(v string) *UpdateTemplateResponseBodyTemplateAudio {
	s.Qscale = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateAudio) SetRemove(v string) *UpdateTemplateResponseBodyTemplateAudio {
	s.Remove = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateAudio) SetSamplerate(v string) *UpdateTemplateResponseBodyTemplateAudio {
	s.Samplerate = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateAudio) SetVolume(v *UpdateTemplateResponseBodyTemplateAudioVolume) *UpdateTemplateResponseBodyTemplateAudio {
	s.Volume = v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateAudio) Validate() error {
	if s.Volume != nil {
		if err := s.Volume.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateTemplateResponseBodyTemplateAudioVolume struct {
	// The expected volume.
	//
	// 	- This parameter takes effect only if the value of Method is dynamic.
	//
	// 	- Unit: decibel.
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
	// 	- Unit: decibel.
	//
	// 	- Valid values: less than or equal to 20.
	//
	// 	- Default value: -20.
	//
	// example:
	//
	// -20
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The range of the volume relative to the expected volume.
	//
	// 	- This parameter takes effect only if the value of Method is dynamic.
	//
	// 	- Unit: decibel.
	//
	// 	- Valid values: [1,20].
	//
	// 	- Default value: 8.
	//
	// example:
	//
	// 8
	LoudnessRangeTarget *string `json:"LoudnessRangeTarget,omitempty" xml:"LoudnessRangeTarget,omitempty"`
	// The volume adjustment method. Valid values:
	//
	// 	- **auto**
	//
	// 	- **dynamic**
	//
	// 	- **linear**
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
	// 	- Unit: decibel.
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

func (s UpdateTemplateResponseBodyTemplateAudioVolume) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateResponseBodyTemplateAudioVolume) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponseBodyTemplateAudioVolume) GetIntegratedLoudnessTarget() *string {
	return s.IntegratedLoudnessTarget
}

func (s *UpdateTemplateResponseBodyTemplateAudioVolume) GetLevel() *string {
	return s.Level
}

func (s *UpdateTemplateResponseBodyTemplateAudioVolume) GetLoudnessRangeTarget() *string {
	return s.LoudnessRangeTarget
}

func (s *UpdateTemplateResponseBodyTemplateAudioVolume) GetMethod() *string {
	return s.Method
}

func (s *UpdateTemplateResponseBodyTemplateAudioVolume) GetPeakLevel() *string {
	return s.PeakLevel
}

func (s *UpdateTemplateResponseBodyTemplateAudioVolume) GetTruePeak() *string {
	return s.TruePeak
}

func (s *UpdateTemplateResponseBodyTemplateAudioVolume) SetIntegratedLoudnessTarget(v string) *UpdateTemplateResponseBodyTemplateAudioVolume {
	s.IntegratedLoudnessTarget = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateAudioVolume) SetLevel(v string) *UpdateTemplateResponseBodyTemplateAudioVolume {
	s.Level = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateAudioVolume) SetLoudnessRangeTarget(v string) *UpdateTemplateResponseBodyTemplateAudioVolume {
	s.LoudnessRangeTarget = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateAudioVolume) SetMethod(v string) *UpdateTemplateResponseBodyTemplateAudioVolume {
	s.Method = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateAudioVolume) SetPeakLevel(v string) *UpdateTemplateResponseBodyTemplateAudioVolume {
	s.PeakLevel = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateAudioVolume) SetTruePeak(v string) *UpdateTemplateResponseBodyTemplateAudioVolume {
	s.TruePeak = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateAudioVolume) Validate() error {
	return dara.Validate(s)
}

type UpdateTemplateResponseBodyTemplateContainer struct {
	// The container format.
	//
	// example:
	//
	// mp4
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
}

func (s UpdateTemplateResponseBodyTemplateContainer) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateResponseBodyTemplateContainer) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponseBodyTemplateContainer) GetFormat() *string {
	return s.Format
}

func (s *UpdateTemplateResponseBodyTemplateContainer) SetFormat(v string) *UpdateTemplateResponseBodyTemplateContainer {
	s.Format = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateContainer) Validate() error {
	return dara.Validate(s)
}

type UpdateTemplateResponseBodyTemplateMuxConfig struct {
	// The duration for which the final frame is paused. Unit: milliseconds.
	Gif *UpdateTemplateResponseBodyTemplateMuxConfigGif `json:"Gif,omitempty" xml:"Gif,omitempty" type:"Struct"`
	// The length of the segment. Unit: seconds.
	Segment *UpdateTemplateResponseBodyTemplateMuxConfigSegment `json:"Segment,omitempty" xml:"Segment,omitempty" type:"Struct"`
	// The loop count.
	Webp *UpdateTemplateResponseBodyTemplateMuxConfigWebp `json:"Webp,omitempty" xml:"Webp,omitempty" type:"Struct"`
}

func (s UpdateTemplateResponseBodyTemplateMuxConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateResponseBodyTemplateMuxConfig) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponseBodyTemplateMuxConfig) GetGif() *UpdateTemplateResponseBodyTemplateMuxConfigGif {
	return s.Gif
}

func (s *UpdateTemplateResponseBodyTemplateMuxConfig) GetSegment() *UpdateTemplateResponseBodyTemplateMuxConfigSegment {
	return s.Segment
}

func (s *UpdateTemplateResponseBodyTemplateMuxConfig) GetWebp() *UpdateTemplateResponseBodyTemplateMuxConfigWebp {
	return s.Webp
}

func (s *UpdateTemplateResponseBodyTemplateMuxConfig) SetGif(v *UpdateTemplateResponseBodyTemplateMuxConfigGif) *UpdateTemplateResponseBodyTemplateMuxConfig {
	s.Gif = v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateMuxConfig) SetSegment(v *UpdateTemplateResponseBodyTemplateMuxConfigSegment) *UpdateTemplateResponseBodyTemplateMuxConfig {
	s.Segment = v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateMuxConfig) SetWebp(v *UpdateTemplateResponseBodyTemplateMuxConfigWebp) *UpdateTemplateResponseBodyTemplateMuxConfig {
	s.Webp = v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateMuxConfig) Validate() error {
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

type UpdateTemplateResponseBodyTemplateMuxConfigGif struct {
	// The loop count.
	//
	// example:
	//
	// bayer
	DitherMode *string `json:"DitherMode,omitempty" xml:"DitherMode,omitempty"`
	// The color dithering algorithm of the palette. Valid values: sierra and bayer.
	//
	// example:
	//
	// false
	FinalDelay *string `json:"FinalDelay,omitempty" xml:"FinalDelay,omitempty"`
	// The segment configurations.
	//
	// example:
	//
	// 0
	IsCustomPalette *string `json:"IsCustomPalette,omitempty" xml:"IsCustomPalette,omitempty"`
	// Indicates whether the custom palette is used.
	//
	// example:
	//
	// 0
	Loop *string `json:"Loop,omitempty" xml:"Loop,omitempty"`
}

func (s UpdateTemplateResponseBodyTemplateMuxConfigGif) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateResponseBodyTemplateMuxConfigGif) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponseBodyTemplateMuxConfigGif) GetDitherMode() *string {
	return s.DitherMode
}

func (s *UpdateTemplateResponseBodyTemplateMuxConfigGif) GetFinalDelay() *string {
	return s.FinalDelay
}

func (s *UpdateTemplateResponseBodyTemplateMuxConfigGif) GetIsCustomPalette() *string {
	return s.IsCustomPalette
}

func (s *UpdateTemplateResponseBodyTemplateMuxConfigGif) GetLoop() *string {
	return s.Loop
}

func (s *UpdateTemplateResponseBodyTemplateMuxConfigGif) SetDitherMode(v string) *UpdateTemplateResponseBodyTemplateMuxConfigGif {
	s.DitherMode = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateMuxConfigGif) SetFinalDelay(v string) *UpdateTemplateResponseBodyTemplateMuxConfigGif {
	s.FinalDelay = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateMuxConfigGif) SetIsCustomPalette(v string) *UpdateTemplateResponseBodyTemplateMuxConfigGif {
	s.IsCustomPalette = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateMuxConfigGif) SetLoop(v string) *UpdateTemplateResponseBodyTemplateMuxConfigGif {
	s.Loop = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateMuxConfigGif) Validate() error {
	return dara.Validate(s)
}

type UpdateTemplateResponseBodyTemplateMuxConfigSegment struct {
	// The name of the template.
	//
	// example:
	//
	// 10
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s UpdateTemplateResponseBodyTemplateMuxConfigSegment) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateResponseBodyTemplateMuxConfigSegment) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponseBodyTemplateMuxConfigSegment) GetDuration() *string {
	return s.Duration
}

func (s *UpdateTemplateResponseBodyTemplateMuxConfigSegment) SetDuration(v string) *UpdateTemplateResponseBodyTemplateMuxConfigSegment {
	s.Duration = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateMuxConfigSegment) Validate() error {
	return dara.Validate(s)
}

type UpdateTemplateResponseBodyTemplateMuxConfigWebp struct {
	// The transmuxing configurations for GIF.
	//
	// example:
	//
	// 0
	Loop *string `json:"Loop,omitempty" xml:"Loop,omitempty"`
}

func (s UpdateTemplateResponseBodyTemplateMuxConfigWebp) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateResponseBodyTemplateMuxConfigWebp) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponseBodyTemplateMuxConfigWebp) GetLoop() *string {
	return s.Loop
}

func (s *UpdateTemplateResponseBodyTemplateMuxConfigWebp) SetLoop(v string) *UpdateTemplateResponseBodyTemplateMuxConfigWebp {
	s.Loop = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateMuxConfigWebp) Validate() error {
	return dara.Validate(s)
}

type UpdateTemplateResponseBodyTemplateTransConfig struct {
	// Indicates whether the video bitrate is checked. If this parameter is set to true and the system detects that the video bitrate of the output file is greater than that of the input file, the video bitrate of the input file is retained after transcoding. Valid values:
	//
	// 	- **true**: The video bitrate is checked.
	//
	// 	- **false**: The video bitrate is not checked.
	//
	// 	- Default value: **false**.
	//
	// example:
	//
	// none
	AdjDarMethod *string `json:"AdjDarMethod,omitempty" xml:"AdjDarMethod,omitempty"`
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
	// false
	IsCheckAudioBitrate *string `json:"IsCheckAudioBitrate,omitempty" xml:"IsCheckAudioBitrate,omitempty"`
	// The status of the template. Valid values:
	//
	// 	- **Normal**: The template is normal.
	//
	// 	- **Deleted**: The template is deleted.
	//
	// example:
	//
	// false
	IsCheckAudioBitrateFail *string `json:"IsCheckAudioBitrateFail,omitempty" xml:"IsCheckAudioBitrateFail,omitempty"`
	// Indicates whether the video bitrate is checked. This parameter has a higher priority than the IsCheckVideoBitrate parameter. Valid values:
	//
	// 	- **true**: The video bitrate is checked
	//
	// 	- **false**: The video bitrate is not checked.
	//
	// 	- Default value: **false**.
	//
	// example:
	//
	// false
	IsCheckReso *string `json:"IsCheckReso,omitempty" xml:"IsCheckReso,omitempty"`
	// Indicates whether the audio bitrate is checked. This parameter has a higher priority than the IsCheckAudioBitrate parameter. Valid values:
	//
	// 	- **true**: The audio bitrate is checked.
	//
	// 	- **false**: The audio bitrate is not checked.
	//
	// 	- Default value: **false**.
	//
	// example:
	//
	// false
	IsCheckResoFail *string `json:"IsCheckResoFail,omitempty" xml:"IsCheckResoFail,omitempty"`
	// Indicates whether the resolution is checked. If this parameter is set to true and the system detects that the resolution of the output file is higher than that of the input file based on the width or height, an error that indicates a transcoding failure is returned. Valid values:
	//
	// 	- **true**: The resolution is checked.
	//
	// 	- **false**: The resolution is not checked.
	//
	// 	- Default value: **false**.
	//
	// example:
	//
	// false
	IsCheckVideoBitrate *string `json:"IsCheckVideoBitrate,omitempty" xml:"IsCheckVideoBitrate,omitempty"`
	// The method of resolution adjustment. Default value: **none**. Valid values:
	//
	// 	- rescale: The input video is rescaled.
	//
	// 	- crop: The input video is cropped.
	//
	// 	- none: No change is made.
	//
	// example:
	//
	// false
	IsCheckVideoBitrateFail *string `json:"IsCheckVideoBitrateFail,omitempty" xml:"IsCheckVideoBitrateFail,omitempty"`
	// Indicates whether the resolution is checked. If the output resolution is higher than the input resolution based on the width or height, the input resolution is retained after transcoding. Valid values:
	//
	// 	- **true**: The resolution is checked.
	//
	// 	- **false**: The resolution is not checked.
	//
	// 	- Default value: **false**.
	//
	// example:
	//
	// onepass
	TransMode *string `json:"TransMode,omitempty" xml:"TransMode,omitempty"`
}

func (s UpdateTemplateResponseBodyTemplateTransConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateResponseBodyTemplateTransConfig) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponseBodyTemplateTransConfig) GetAdjDarMethod() *string {
	return s.AdjDarMethod
}

func (s *UpdateTemplateResponseBodyTemplateTransConfig) GetIsCheckAudioBitrate() *string {
	return s.IsCheckAudioBitrate
}

func (s *UpdateTemplateResponseBodyTemplateTransConfig) GetIsCheckAudioBitrateFail() *string {
	return s.IsCheckAudioBitrateFail
}

func (s *UpdateTemplateResponseBodyTemplateTransConfig) GetIsCheckReso() *string {
	return s.IsCheckReso
}

func (s *UpdateTemplateResponseBodyTemplateTransConfig) GetIsCheckResoFail() *string {
	return s.IsCheckResoFail
}

func (s *UpdateTemplateResponseBodyTemplateTransConfig) GetIsCheckVideoBitrate() *string {
	return s.IsCheckVideoBitrate
}

func (s *UpdateTemplateResponseBodyTemplateTransConfig) GetIsCheckVideoBitrateFail() *string {
	return s.IsCheckVideoBitrateFail
}

func (s *UpdateTemplateResponseBodyTemplateTransConfig) GetTransMode() *string {
	return s.TransMode
}

func (s *UpdateTemplateResponseBodyTemplateTransConfig) SetAdjDarMethod(v string) *UpdateTemplateResponseBodyTemplateTransConfig {
	s.AdjDarMethod = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateTransConfig) SetIsCheckAudioBitrate(v string) *UpdateTemplateResponseBodyTemplateTransConfig {
	s.IsCheckAudioBitrate = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateTransConfig) SetIsCheckAudioBitrateFail(v string) *UpdateTemplateResponseBodyTemplateTransConfig {
	s.IsCheckAudioBitrateFail = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateTransConfig) SetIsCheckReso(v string) *UpdateTemplateResponseBodyTemplateTransConfig {
	s.IsCheckReso = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateTransConfig) SetIsCheckResoFail(v string) *UpdateTemplateResponseBodyTemplateTransConfig {
	s.IsCheckResoFail = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateTransConfig) SetIsCheckVideoBitrate(v string) *UpdateTemplateResponseBodyTemplateTransConfig {
	s.IsCheckVideoBitrate = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateTransConfig) SetIsCheckVideoBitrateFail(v string) *UpdateTemplateResponseBodyTemplateTransConfig {
	s.IsCheckVideoBitrateFail = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateTransConfig) SetTransMode(v string) *UpdateTemplateResponseBodyTemplateTransConfig {
	s.TransMode = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateTransConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateTemplateResponseBodyTemplateVideo struct {
	// The maximum bitrate of the video. Unit: Kbit/s.
	//
	// example:
	//
	// 200
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The upper limit of the total bitrate. Unit: Kbit/s.
	BitrateBnd *UpdateTemplateResponseBodyTemplateVideoBitrateBnd `json:"BitrateBnd,omitempty" xml:"BitrateBnd,omitempty" type:"Struct"`
	// The level of quality control on the video.
	//
	// example:
	//
	// 6000
	Bufsize *string `json:"Bufsize,omitempty" xml:"Bufsize,omitempty"`
	// The height of the output video.
	//
	// 	- Unit: pixel.
	//
	// 	- Default value: the height of the input video.
	//
	// example:
	//
	// H.264
	Codec *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	// Indicates whether the video stream is deleted. Valid values:
	//
	// 	- **true**: The video stream is deleted.
	//
	// 	- **false**: The video stream is retained.
	//
	// 	- Default value: **false**.
	//
	// example:
	//
	// 15
	Crf *string `json:"Crf,omitempty" xml:"Crf,omitempty"`
	// The average bitrate of the video. Unit: Kbit/s.
	//
	// example:
	//
	// border
	Crop *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	// The average bitrate range of the video.
	//
	// example:
	//
	// 10
	Degrain *string `json:"Degrain,omitempty" xml:"Degrain,omitempty"`
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
	// 25
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The width of the video.
	//
	// 	- Unit: pixel.
	//
	// 	- Default value: **the width of the input video**.
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
	// The level of the independent denoising algorithm.
	//
	// example:
	//
	// 800
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The size of the buffer.
	//
	// 	- Unit: KB.
	//
	// 	- Default value: **6000**.
	//
	// example:
	//
	// false
	LongShortMode *string `json:"LongShortMode,omitempty" xml:"LongShortMode,omitempty"`
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
	// 60
	MaxFps *string `json:"MaxFps,omitempty" xml:"MaxFps,omitempty"`
	// The maximum frame rate.
	//
	// example:
	//
	// 500
	Maxrate *string `json:"Maxrate,omitempty" xml:"Maxrate,omitempty"`
	// The Narrowband HD settings.
	NarrowBand *UpdateTemplateResponseBodyTemplateVideoNarrowBand `json:"NarrowBand,omitempty" xml:"NarrowBand,omitempty" type:"Struct"`
	// The video codec. Default value: **H.264**.
	//
	// example:
	//
	// 1280:800:0:140
	Pad *string `json:"Pad,omitempty" xml:"Pad,omitempty"`
	// The black borders added to the video.
	//
	// 	- Format: width:height:left:top.
	//
	// 	- Example: 1280:800:0:140.
	//
	// example:
	//
	// yuv420p
	PixFmt *string `json:"PixFmt,omitempty" xml:"PixFmt,omitempty"`
	// The scan mode. Valid values:
	//
	// 	- **interlaced**: An interlaced scan is performed.
	//
	// 	- **progressive**: A progressive scan is performed.
	//
	// example:
	//
	// medium
	Preset *string `json:"Preset,omitempty" xml:"Preset,omitempty"`
	// The bitrate quality control factor.
	//
	// 	- Default value if the Codec parameter is set to H.264: **23**. Default value if the Codec parameter is set to H.265: **26**.
	//
	// 	- If this parameter is returned, the setting of the Bitrate parameter is invalid.
	//
	// example:
	//
	// high
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The method used to crop the video.
	//
	// 	- **border**: automatically detects and removes borders.
	//
	// 	- Value in the width:height:left:top format: crops the video based on custom settings.***	- Example: 1280:800:0:140.
	//
	// example:
	//
	// 1
	Qscale *string `json:"Qscale,omitempty" xml:"Qscale,omitempty"`
	// The maximum number of frames between two keyframes. Default value: **250**.
	//
	// example:
	//
	// false
	Remove *string `json:"Remove,omitempty" xml:"Remove,omitempty"`
	// The general transcoding configurations.
	//
	// example:
	//
	// 1
	ResoPriority *string `json:"ResoPriority,omitempty" xml:"ResoPriority,omitempty"`
	// The policy of resolution adjustment.
	//
	// example:
	//
	// interlaced
	ScanMode *string `json:"ScanMode,omitempty" xml:"ScanMode,omitempty"`
	// The frame rate.
	//
	// 	- A value of 60 is returned if the frame rate of the input video exceeds 60.
	//
	// 	- Default value: the frame rate of the input video.
	//
	// example:
	//
	// 256
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s UpdateTemplateResponseBodyTemplateVideo) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateResponseBodyTemplateVideo) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponseBodyTemplateVideo) GetBitrate() *string {
	return s.Bitrate
}

func (s *UpdateTemplateResponseBodyTemplateVideo) GetBitrateBnd() *UpdateTemplateResponseBodyTemplateVideoBitrateBnd {
	return s.BitrateBnd
}

func (s *UpdateTemplateResponseBodyTemplateVideo) GetBufsize() *string {
	return s.Bufsize
}

func (s *UpdateTemplateResponseBodyTemplateVideo) GetCodec() *string {
	return s.Codec
}

func (s *UpdateTemplateResponseBodyTemplateVideo) GetCrf() *string {
	return s.Crf
}

func (s *UpdateTemplateResponseBodyTemplateVideo) GetCrop() *string {
	return s.Crop
}

func (s *UpdateTemplateResponseBodyTemplateVideo) GetDegrain() *string {
	return s.Degrain
}

func (s *UpdateTemplateResponseBodyTemplateVideo) GetFps() *string {
	return s.Fps
}

func (s *UpdateTemplateResponseBodyTemplateVideo) GetGop() *string {
	return s.Gop
}

func (s *UpdateTemplateResponseBodyTemplateVideo) GetHdr2sdr() *string {
	return s.Hdr2sdr
}

func (s *UpdateTemplateResponseBodyTemplateVideo) GetHeight() *string {
	return s.Height
}

func (s *UpdateTemplateResponseBodyTemplateVideo) GetLongShortMode() *string {
	return s.LongShortMode
}

func (s *UpdateTemplateResponseBodyTemplateVideo) GetMaxFps() *string {
	return s.MaxFps
}

func (s *UpdateTemplateResponseBodyTemplateVideo) GetMaxrate() *string {
	return s.Maxrate
}

func (s *UpdateTemplateResponseBodyTemplateVideo) GetNarrowBand() *UpdateTemplateResponseBodyTemplateVideoNarrowBand {
	return s.NarrowBand
}

func (s *UpdateTemplateResponseBodyTemplateVideo) GetPad() *string {
	return s.Pad
}

func (s *UpdateTemplateResponseBodyTemplateVideo) GetPixFmt() *string {
	return s.PixFmt
}

func (s *UpdateTemplateResponseBodyTemplateVideo) GetPreset() *string {
	return s.Preset
}

func (s *UpdateTemplateResponseBodyTemplateVideo) GetProfile() *string {
	return s.Profile
}

func (s *UpdateTemplateResponseBodyTemplateVideo) GetQscale() *string {
	return s.Qscale
}

func (s *UpdateTemplateResponseBodyTemplateVideo) GetRemove() *string {
	return s.Remove
}

func (s *UpdateTemplateResponseBodyTemplateVideo) GetResoPriority() *string {
	return s.ResoPriority
}

func (s *UpdateTemplateResponseBodyTemplateVideo) GetScanMode() *string {
	return s.ScanMode
}

func (s *UpdateTemplateResponseBodyTemplateVideo) GetWidth() *string {
	return s.Width
}

func (s *UpdateTemplateResponseBodyTemplateVideo) SetBitrate(v string) *UpdateTemplateResponseBodyTemplateVideo {
	s.Bitrate = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateVideo) SetBitrateBnd(v *UpdateTemplateResponseBodyTemplateVideoBitrateBnd) *UpdateTemplateResponseBodyTemplateVideo {
	s.BitrateBnd = v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateVideo) SetBufsize(v string) *UpdateTemplateResponseBodyTemplateVideo {
	s.Bufsize = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateVideo) SetCodec(v string) *UpdateTemplateResponseBodyTemplateVideo {
	s.Codec = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateVideo) SetCrf(v string) *UpdateTemplateResponseBodyTemplateVideo {
	s.Crf = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateVideo) SetCrop(v string) *UpdateTemplateResponseBodyTemplateVideo {
	s.Crop = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateVideo) SetDegrain(v string) *UpdateTemplateResponseBodyTemplateVideo {
	s.Degrain = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateVideo) SetFps(v string) *UpdateTemplateResponseBodyTemplateVideo {
	s.Fps = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateVideo) SetGop(v string) *UpdateTemplateResponseBodyTemplateVideo {
	s.Gop = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateVideo) SetHdr2sdr(v string) *UpdateTemplateResponseBodyTemplateVideo {
	s.Hdr2sdr = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateVideo) SetHeight(v string) *UpdateTemplateResponseBodyTemplateVideo {
	s.Height = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateVideo) SetLongShortMode(v string) *UpdateTemplateResponseBodyTemplateVideo {
	s.LongShortMode = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateVideo) SetMaxFps(v string) *UpdateTemplateResponseBodyTemplateVideo {
	s.MaxFps = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateVideo) SetMaxrate(v string) *UpdateTemplateResponseBodyTemplateVideo {
	s.Maxrate = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateVideo) SetNarrowBand(v *UpdateTemplateResponseBodyTemplateVideoNarrowBand) *UpdateTemplateResponseBodyTemplateVideo {
	s.NarrowBand = v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateVideo) SetPad(v string) *UpdateTemplateResponseBodyTemplateVideo {
	s.Pad = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateVideo) SetPixFmt(v string) *UpdateTemplateResponseBodyTemplateVideo {
	s.PixFmt = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateVideo) SetPreset(v string) *UpdateTemplateResponseBodyTemplateVideo {
	s.Preset = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateVideo) SetProfile(v string) *UpdateTemplateResponseBodyTemplateVideo {
	s.Profile = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateVideo) SetQscale(v string) *UpdateTemplateResponseBodyTemplateVideo {
	s.Qscale = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateVideo) SetRemove(v string) *UpdateTemplateResponseBodyTemplateVideo {
	s.Remove = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateVideo) SetResoPriority(v string) *UpdateTemplateResponseBodyTemplateVideo {
	s.ResoPriority = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateVideo) SetScanMode(v string) *UpdateTemplateResponseBodyTemplateVideo {
	s.ScanMode = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateVideo) SetWidth(v string) *UpdateTemplateResponseBodyTemplateVideo {
	s.Width = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateVideo) Validate() error {
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

type UpdateTemplateResponseBodyTemplateVideoBitrateBnd struct {
	// The lower limit of the total bitrate. Unit: Kbit/s.
	//
	// example:
	//
	// 500
	Max *string `json:"Max,omitempty" xml:"Max,omitempty"`
	// The pixel format. Valid values: standard pixel formats such as yuv420p and yuvj420p.
	//
	// example:
	//
	// 100
	Min *string `json:"Min,omitempty" xml:"Min,omitempty"`
}

func (s UpdateTemplateResponseBodyTemplateVideoBitrateBnd) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateResponseBodyTemplateVideoBitrateBnd) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponseBodyTemplateVideoBitrateBnd) GetMax() *string {
	return s.Max
}

func (s *UpdateTemplateResponseBodyTemplateVideoBitrateBnd) GetMin() *string {
	return s.Min
}

func (s *UpdateTemplateResponseBodyTemplateVideoBitrateBnd) SetMax(v string) *UpdateTemplateResponseBodyTemplateVideoBitrateBnd {
	s.Max = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateVideoBitrateBnd) SetMin(v string) *UpdateTemplateResponseBodyTemplateVideoBitrateBnd {
	s.Min = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateVideoBitrateBnd) Validate() error {
	return dara.Validate(s)
}

type UpdateTemplateResponseBodyTemplateVideoNarrowBand struct {
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

func (s UpdateTemplateResponseBodyTemplateVideoNarrowBand) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateResponseBodyTemplateVideoNarrowBand) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponseBodyTemplateVideoNarrowBand) GetAbrmax() *float32 {
	return s.Abrmax
}

func (s *UpdateTemplateResponseBodyTemplateVideoNarrowBand) GetMaxAbrRatio() *float32 {
	return s.MaxAbrRatio
}

func (s *UpdateTemplateResponseBodyTemplateVideoNarrowBand) GetVersion() *string {
	return s.Version
}

func (s *UpdateTemplateResponseBodyTemplateVideoNarrowBand) SetAbrmax(v float32) *UpdateTemplateResponseBodyTemplateVideoNarrowBand {
	s.Abrmax = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateVideoNarrowBand) SetMaxAbrRatio(v float32) *UpdateTemplateResponseBodyTemplateVideoNarrowBand {
	s.MaxAbrRatio = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateVideoNarrowBand) SetVersion(v string) *UpdateTemplateResponseBodyTemplateVideoNarrowBand {
	s.Version = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplateVideoNarrowBand) Validate() error {
	return dara.Validate(s)
}
