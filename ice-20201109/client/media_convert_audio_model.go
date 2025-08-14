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
	Bitrate    *int64  `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	Channels   *int64  `json:"Channels,omitempty" xml:"Channels,omitempty"`
	Codec      *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	Profile    *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	Remove     *bool   `json:"Remove,omitempty" xml:"Remove,omitempty"`
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
