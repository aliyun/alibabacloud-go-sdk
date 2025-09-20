// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAudioStream interface {
	dara.Model
	String() string
	GoString() string
	SetBitrate(v int64) *AudioStream
	GetBitrate() *int64
	SetChannelLayout(v string) *AudioStream
	GetChannelLayout() *string
	SetChannels(v int64) *AudioStream
	GetChannels() *int64
	SetCodecLongName(v string) *AudioStream
	GetCodecLongName() *string
	SetCodecName(v string) *AudioStream
	GetCodecName() *string
	SetCodecTag(v string) *AudioStream
	GetCodecTag() *string
	SetCodecTagString(v string) *AudioStream
	GetCodecTagString() *string
	SetCodecTimeBase(v string) *AudioStream
	GetCodecTimeBase() *string
	SetDuration(v float64) *AudioStream
	GetDuration() *float64
	SetFrameCount(v int64) *AudioStream
	GetFrameCount() *int64
	SetIndex(v int64) *AudioStream
	GetIndex() *int64
	SetLanguage(v string) *AudioStream
	GetLanguage() *string
	SetLyric(v string) *AudioStream
	GetLyric() *string
	SetSampleFormat(v string) *AudioStream
	GetSampleFormat() *string
	SetSampleRate(v int64) *AudioStream
	GetSampleRate() *int64
	SetStartTime(v float64) *AudioStream
	GetStartTime() *float64
	SetTimeBase(v string) *AudioStream
	GetTimeBase() *string
}

type AudioStream struct {
	Bitrate        *int64   `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	ChannelLayout  *string  `json:"ChannelLayout,omitempty" xml:"ChannelLayout,omitempty"`
	Channels       *int64   `json:"Channels,omitempty" xml:"Channels,omitempty"`
	CodecLongName  *string  `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	CodecName      *string  `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	CodecTag       *string  `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	CodecTagString *string  `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	CodecTimeBase  *string  `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	Duration       *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	FrameCount     *int64   `json:"FrameCount,omitempty" xml:"FrameCount,omitempty"`
	Index          *int64   `json:"Index,omitempty" xml:"Index,omitempty"`
	Language       *string  `json:"Language,omitempty" xml:"Language,omitempty"`
	Lyric          *string  `json:"Lyric,omitempty" xml:"Lyric,omitempty"`
	SampleFormat   *string  `json:"SampleFormat,omitempty" xml:"SampleFormat,omitempty"`
	SampleRate     *int64   `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	StartTime      *float64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TimeBase       *string  `json:"TimeBase,omitempty" xml:"TimeBase,omitempty"`
}

func (s AudioStream) String() string {
	return dara.Prettify(s)
}

func (s AudioStream) GoString() string {
	return s.String()
}

func (s *AudioStream) GetBitrate() *int64 {
	return s.Bitrate
}

func (s *AudioStream) GetChannelLayout() *string {
	return s.ChannelLayout
}

func (s *AudioStream) GetChannels() *int64 {
	return s.Channels
}

func (s *AudioStream) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *AudioStream) GetCodecName() *string {
	return s.CodecName
}

func (s *AudioStream) GetCodecTag() *string {
	return s.CodecTag
}

func (s *AudioStream) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *AudioStream) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *AudioStream) GetDuration() *float64 {
	return s.Duration
}

func (s *AudioStream) GetFrameCount() *int64 {
	return s.FrameCount
}

func (s *AudioStream) GetIndex() *int64 {
	return s.Index
}

func (s *AudioStream) GetLanguage() *string {
	return s.Language
}

func (s *AudioStream) GetLyric() *string {
	return s.Lyric
}

func (s *AudioStream) GetSampleFormat() *string {
	return s.SampleFormat
}

func (s *AudioStream) GetSampleRate() *int64 {
	return s.SampleRate
}

func (s *AudioStream) GetStartTime() *float64 {
	return s.StartTime
}

func (s *AudioStream) GetTimeBase() *string {
	return s.TimeBase
}

func (s *AudioStream) SetBitrate(v int64) *AudioStream {
	s.Bitrate = &v
	return s
}

func (s *AudioStream) SetChannelLayout(v string) *AudioStream {
	s.ChannelLayout = &v
	return s
}

func (s *AudioStream) SetChannels(v int64) *AudioStream {
	s.Channels = &v
	return s
}

func (s *AudioStream) SetCodecLongName(v string) *AudioStream {
	s.CodecLongName = &v
	return s
}

func (s *AudioStream) SetCodecName(v string) *AudioStream {
	s.CodecName = &v
	return s
}

func (s *AudioStream) SetCodecTag(v string) *AudioStream {
	s.CodecTag = &v
	return s
}

func (s *AudioStream) SetCodecTagString(v string) *AudioStream {
	s.CodecTagString = &v
	return s
}

func (s *AudioStream) SetCodecTimeBase(v string) *AudioStream {
	s.CodecTimeBase = &v
	return s
}

func (s *AudioStream) SetDuration(v float64) *AudioStream {
	s.Duration = &v
	return s
}

func (s *AudioStream) SetFrameCount(v int64) *AudioStream {
	s.FrameCount = &v
	return s
}

func (s *AudioStream) SetIndex(v int64) *AudioStream {
	s.Index = &v
	return s
}

func (s *AudioStream) SetLanguage(v string) *AudioStream {
	s.Language = &v
	return s
}

func (s *AudioStream) SetLyric(v string) *AudioStream {
	s.Lyric = &v
	return s
}

func (s *AudioStream) SetSampleFormat(v string) *AudioStream {
	s.SampleFormat = &v
	return s
}

func (s *AudioStream) SetSampleRate(v int64) *AudioStream {
	s.SampleRate = &v
	return s
}

func (s *AudioStream) SetStartTime(v float64) *AudioStream {
	s.StartTime = &v
	return s
}

func (s *AudioStream) SetTimeBase(v string) *AudioStream {
	s.TimeBase = &v
	return s
}

func (s *AudioStream) Validate() error {
	return dara.Validate(s)
}
