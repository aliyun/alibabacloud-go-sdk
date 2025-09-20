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
	DisableAudio   *bool                      `json:"DisableAudio,omitempty" xml:"DisableAudio,omitempty"`
	FilterAudio    *TargetAudioFilterAudio    `json:"FilterAudio,omitempty" xml:"FilterAudio,omitempty" type:"Struct"`
	Stream         []*int64                   `json:"Stream,omitempty" xml:"Stream,omitempty" type:"Repeated"`
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
	return dara.Validate(s)
}

type TargetAudioFilterAudio struct {
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
	Bitrate          *int32  `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	BitrateOption    *string `json:"BitrateOption,omitempty" xml:"BitrateOption,omitempty"`
	BitsPerSample    *int32  `json:"BitsPerSample,omitempty" xml:"BitsPerSample,omitempty"`
	Channel          *int32  `json:"Channel,omitempty" xml:"Channel,omitempty"`
	Codec            *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	Quality          *int32  `json:"Quality,omitempty" xml:"Quality,omitempty"`
	SampleRate       *int32  `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
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
