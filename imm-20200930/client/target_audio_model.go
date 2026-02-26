// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTargetAudio interface {
	dara.Model
	String() string
	GoString() string
	SetDisableAudio(v bool) *TargetAudio
	GetDisableAudio() *bool
	SetFilterAudio(v *TargetAudioFilterAudio) *TargetAudio
	GetFilterAudio() *TargetAudioFilterAudio
	SetStream(v []*int64) *TargetAudio
	GetStream() []*int64
	SetTranscodeAudio(v *TargetAudioTranscodeAudio) *TargetAudio
	GetTranscodeAudio() *TargetAudioTranscodeAudio
}

type TargetAudio struct {
	// Specifies whether to disable audio stream generation. Valid values:
	//
	// 	- true: disables audio stream generation. No audio stream is included in the output file.
	//
	// 	- false: does not disable audio stream generation. This is the default value.
	//
	// example:
	//
	// false
	DisableAudio *bool `json:"DisableAudio,omitempty" xml:"DisableAudio,omitempty"`
	// The audio processing settings. This parameter is invalid if **TranscodeAudio*	- is left empty or **TranscodeAudio.Codec*	- is set to copy.
	//
	// >  This parameter is not available to the GenerateVideoPlaylist operation.
	FilterAudio *TargetAudioFilterAudio `json:"FilterAudio,omitempty" xml:"FilterAudio,omitempty" type:"Struct"`
	// The index numbers of audio streams. If you do not specify this parameter, the first audio stream (the one with the smallest index number) is processed. If the array contains an element greater than 100, all audio streams are processed.
	//
	// 	- For example, you can set the parameter to `[0,1]` to process audio streams with index numbers 0 and 1, `[1]` to process only the audio stream with the index number 1, or `[101]` to process all audio streams.
	//
	// >  If you specify an index number but no audio stream with the index number is found, the index number is ignored.
	Stream []*int64 `json:"Stream,omitempty" xml:"Stream,omitempty" type:"Repeated"`
	// The audio transcoding settings. If you do not specify this parameter, no audio streams are included in the output file.
	//
	// >  We recommend that you do not use this parameter to disable audio stream generation.
	TranscodeAudio *TargetAudioTranscodeAudio `json:"TranscodeAudio,omitempty" xml:"TranscodeAudio,omitempty" type:"Struct"`
}

func (s TargetAudio) String() string {
	return dara.Prettify(s)
}

func (s TargetAudio) GoString() string {
	return s.String()
}

func (s *TargetAudio) GetDisableAudio() *bool {
	return s.DisableAudio
}

func (s *TargetAudio) GetFilterAudio() *TargetAudioFilterAudio {
	return s.FilterAudio
}

func (s *TargetAudio) GetStream() []*int64 {
	return s.Stream
}

func (s *TargetAudio) GetTranscodeAudio() *TargetAudioTranscodeAudio {
	return s.TranscodeAudio
}

func (s *TargetAudio) SetDisableAudio(v bool) *TargetAudio {
	s.DisableAudio = &v
	return s
}

func (s *TargetAudio) SetFilterAudio(v *TargetAudioFilterAudio) *TargetAudio {
	s.FilterAudio = v
	return s
}

func (s *TargetAudio) SetStream(v []*int64) *TargetAudio {
	s.Stream = v
	return s
}

func (s *TargetAudio) SetTranscodeAudio(v *TargetAudioTranscodeAudio) *TargetAudio {
	s.TranscodeAudio = v
	return s
}

func (s *TargetAudio) Validate() error {
	if s.FilterAudio != nil {
		if err := s.FilterAudio.Validate(); err != nil {
			return err
		}
	}
	if s.TranscodeAudio != nil {
		if err := s.TranscodeAudio.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TargetAudioFilterAudio struct {
	// Specifies whether to mix all sound tracks into a single track. Valid values:
	//
	// 	- false (default)
	//
	// 	- true
	//
	// example:
	//
	// false
	Mixing *bool `json:"Mixing,omitempty" xml:"Mixing,omitempty"`
}

func (s TargetAudioFilterAudio) String() string {
	return dara.Prettify(s)
}

func (s TargetAudioFilterAudio) GoString() string {
	return s.String()
}

func (s *TargetAudioFilterAudio) GetMixing() *bool {
	return s.Mixing
}

func (s *TargetAudioFilterAudio) SetMixing(v bool) *TargetAudioFilterAudio {
	s.Mixing = &v
	return s
}

func (s *TargetAudioFilterAudio) Validate() error {
	return dara.Validate(s)
}

type TargetAudioTranscodeAudio struct {
	// The bitrate of the audio stream. Unit: bit/s. This parameter and the **Quality*	- parameter are mutually exclusive. Valid values: 1000 to 10000000.
	//
	// example:
	//
	// 64000
	Bitrate *int32 `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The audio bitrate option. Valid values:
	//
	// 	- fixed: always uses the target bitrate.
	//
	// 	- adaptive: uses the source bitrate when the source bitrate is smaller than the target bitrate.
	//
	// 	- fall: returns a failure when the source bitrate is smaller than the target bitrate.
	//
	// Default values:
	//
	// 	- fixed for the CreateMediaConvert operation.
	//
	// 	- adaptive for the GenerateVideoPlaylist operation.
	//
	// >  This parameter must be used in conjunction with the **Bitrate*	- parameter.
	//
	// example:
	//
	// fixed
	BitrateOption *string `json:"BitrateOption,omitempty" xml:"BitrateOption,omitempty"`
	// The audio bit depth. Valid values: 16 and 24.
	//
	// >  This parameter takes effect only when Codec is set to flac.
	//
	// example:
	//
	// 16
	BitsPerSample *int32 `json:"BitsPerSample,omitempty" xml:"BitsPerSample,omitempty"`
	// The number of sound channels. By default, the audio stream has the same number of sound channels as the source audio. Valid values: [1,8].
	//
	// >  The number of sound channels varies with audio formats: one or two for MP3, up to six for AC3 5.1, and one for AMR.
	//
	// example:
	//
	// 2
	Channel *int32 `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// The codec. Valid values:
	//
	// 	- copy, mp3, vorbis, aac, flac, ac3, opus, and amr for the CreateMediaConvert operation. The default value is copy.
	//
	// 	- aac for the GenerateVideoPlaylist operation. The default value is aac.
	//
	// >  When you set the parameter to copy, the audio stream is directly copied into the output file and all other parameters in **TranscodeAudio*	- do not take effect. The copy value is commonly used in container format conversion scenarios. You cannot use this value in audio merging scenarios.
	//
	// example:
	//
	// aac
	Codec *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	// The audio quality. Valid values: 0 to 100. The greater the value, the higher the quality. This parameter and the **Bitrate*	- parameter are mutually exclusive.
	//
	// example:
	//
	// 6
	Quality *int32 `json:"Quality,omitempty" xml:"Quality,omitempty"`
	// The sampling rate option. Unit: Hz. By default, the source sampling rate is used. Valid values: 8000, 12025, 12000, 16000, 22050, 24000, 32000, 44100, 48000, 64000, 88200, and 96000.
	//
	// >  Supported sampling rates vary with formats: 48 kHz and lower for MP3, 8 kHz, 12 kHz, 16 kHz, 24 kHz, and 48 kHz for Opus, 32 kHz, 44.1 kHz, and 48 kHz for AC3, and 8 kHz and 16 kHz for AMR.
	//
	// example:
	//
	// 12050
	SampleRate *int32 `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	// The sampling rate option. Valid values:
	//
	// 	- fixed: always uses the target sampling rate.
	//
	// 	- adaptive: uses the source sampling rate when the source sampling rate is smaller than the target sampling rate.
	//
	// 	- fall: returns a failure when the source sampling rate is smaller than the target sampling rate.
	//
	// Default values:
	//
	// 	- fixed for the CreateMediaConvert operation.
	//
	// 	- adaptive for the GenerateVideoPlaylist operation.
	//
	// >  This parameter must be used in conjunction with the **SampleRate*	- parameter.
	//
	// example:
	//
	// fixed
	SampleRateOption *string `json:"SampleRateOption,omitempty" xml:"SampleRateOption,omitempty"`
}

func (s TargetAudioTranscodeAudio) String() string {
	return dara.Prettify(s)
}

func (s TargetAudioTranscodeAudio) GoString() string {
	return s.String()
}

func (s *TargetAudioTranscodeAudio) GetBitrate() *int32 {
	return s.Bitrate
}

func (s *TargetAudioTranscodeAudio) GetBitrateOption() *string {
	return s.BitrateOption
}

func (s *TargetAudioTranscodeAudio) GetBitsPerSample() *int32 {
	return s.BitsPerSample
}

func (s *TargetAudioTranscodeAudio) GetChannel() *int32 {
	return s.Channel
}

func (s *TargetAudioTranscodeAudio) GetCodec() *string {
	return s.Codec
}

func (s *TargetAudioTranscodeAudio) GetQuality() *int32 {
	return s.Quality
}

func (s *TargetAudioTranscodeAudio) GetSampleRate() *int32 {
	return s.SampleRate
}

func (s *TargetAudioTranscodeAudio) GetSampleRateOption() *string {
	return s.SampleRateOption
}

func (s *TargetAudioTranscodeAudio) SetBitrate(v int32) *TargetAudioTranscodeAudio {
	s.Bitrate = &v
	return s
}

func (s *TargetAudioTranscodeAudio) SetBitrateOption(v string) *TargetAudioTranscodeAudio {
	s.BitrateOption = &v
	return s
}

func (s *TargetAudioTranscodeAudio) SetBitsPerSample(v int32) *TargetAudioTranscodeAudio {
	s.BitsPerSample = &v
	return s
}

func (s *TargetAudioTranscodeAudio) SetChannel(v int32) *TargetAudioTranscodeAudio {
	s.Channel = &v
	return s
}

func (s *TargetAudioTranscodeAudio) SetCodec(v string) *TargetAudioTranscodeAudio {
	s.Codec = &v
	return s
}

func (s *TargetAudioTranscodeAudio) SetQuality(v int32) *TargetAudioTranscodeAudio {
	s.Quality = &v
	return s
}

func (s *TargetAudioTranscodeAudio) SetSampleRate(v int32) *TargetAudioTranscodeAudio {
	s.SampleRate = &v
	return s
}

func (s *TargetAudioTranscodeAudio) SetSampleRateOption(v string) *TargetAudioTranscodeAudio {
	s.SampleRateOption = &v
	return s
}

func (s *TargetAudioTranscodeAudio) Validate() error {
	return dara.Validate(s)
}
