// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMediaConvertAudio interface {
	dara.Model
	String() string
	GoString() string
	SetBitrate(v int64) *MediaConvertAudio
	GetBitrate() *int64
	SetChannels(v int64) *MediaConvertAudio
	GetChannels() *int64
	SetCodec(v string) *MediaConvertAudio
	GetCodec() *string
	SetProfile(v string) *MediaConvertAudio
	GetProfile() *string
	SetRemove(v bool) *MediaConvertAudio
	GetRemove() *bool
	SetSamplerate(v string) *MediaConvertAudio
	GetSamplerate() *string
}

type MediaConvertAudio struct {
	// The audio bitrate of the output file.
	//
	// 	- Unit: Kbps.
	//
	// 	- Valid values: [8,1000].
	//
	// 	- Default value: 128.
	//
	// 	- Common values: 64, 128, and 256.
	//
	// example:
	//
	// 128
	Bitrate *int64 `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The number of audio channels.
	//
	// 	- Valid values: 0, 1, 2, 4, 5, 6, and 8.
	//
	//     	- If the Codec parameter is set to MP3 or OPUS, you can set this parameter to 0, 1, or 2.
	//
	//     	- If the Codec parameter is set to AAC or FLAC, you can set this parameter to 0, 1, 2, 4, 5, 6, or 8.
	//
	//     	- If the Codec parameter is set to VORBIS, you can set this parameter to 2.
	//
	//     	- If the Format parameter is set to MPD, you cannot set this parameter to 8.
	//
	// 	- Default value: 2.
	//
	// 	- Set the value to 0 to preserve the original number of channels from the source file.
	//
	// example:
	//
	// 2
	Channels *int64 `json:"Channels,omitempty" xml:"Channels,omitempty"`
	// The audio codec.
	//
	// 	- Valid values: AAC, AC3, EAC3, MP2, MP3, FLAC, OPUS, VORBIS, WMA-V1, WMA-V2, and pcm_s16le.
	//
	// 	- Default value: AAC.
	//
	// example:
	//
	// AAC
	Codec *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	// The audio codec profile.
	//
	// 	- This parameter takes effect only if the Codec parameter is set to AAC.
	//
	// 	- Valid values: aac_low, aac_he, aac_he_v2, aac_ld, and aac_eld.
	//
	// 	- Default value: aac_low.
	//
	// example:
	//
	// aac_low
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// Specifies whether to remove the audio stream from the output.
	//
	// 	- true: deletes the audio stream. All other parameters in the Audio object are ignored.
	//
	// 	- false: retains the audio stream.
	//
	// 	- Default value: false.
	Remove *bool `json:"Remove,omitempty" xml:"Remove,omitempty"`
	// The sample rate.
	//
	// 	- Unit: Hz
	//
	// 	- Valid values: 22050, 32000, 44100, 48000, and 96000.
	//
	// 	- Default value: 44100.
	//
	// The supported sample rates vary depending on the container and codec format. For example, when the audio codec is MP3, a sample rate of 96000 is not supported. If the container format is FLV, only sample rates of 22050 and 44100 are supported.
	//
	// example:
	//
	// 44100
	Samplerate *string `json:"Samplerate,omitempty" xml:"Samplerate,omitempty"`
}

func (s MediaConvertAudio) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertAudio) GoString() string {
	return s.String()
}

func (s *MediaConvertAudio) GetBitrate() *int64 {
	return s.Bitrate
}

func (s *MediaConvertAudio) GetChannels() *int64 {
	return s.Channels
}

func (s *MediaConvertAudio) GetCodec() *string {
	return s.Codec
}

func (s *MediaConvertAudio) GetProfile() *string {
	return s.Profile
}

func (s *MediaConvertAudio) GetRemove() *bool {
	return s.Remove
}

func (s *MediaConvertAudio) GetSamplerate() *string {
	return s.Samplerate
}

func (s *MediaConvertAudio) SetBitrate(v int64) *MediaConvertAudio {
	s.Bitrate = &v
	return s
}

func (s *MediaConvertAudio) SetChannels(v int64) *MediaConvertAudio {
	s.Channels = &v
	return s
}

func (s *MediaConvertAudio) SetCodec(v string) *MediaConvertAudio {
	s.Codec = &v
	return s
}

func (s *MediaConvertAudio) SetProfile(v string) *MediaConvertAudio {
	s.Profile = &v
	return s
}

func (s *MediaConvertAudio) SetRemove(v bool) *MediaConvertAudio {
	s.Remove = &v
	return s
}

func (s *MediaConvertAudio) SetSamplerate(v string) *MediaConvertAudio {
	s.Samplerate = &v
	return s
}

func (s *MediaConvertAudio) Validate() error {
	return dara.Validate(s)
}
