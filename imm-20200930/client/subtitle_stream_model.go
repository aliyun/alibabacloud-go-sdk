// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubtitleStream interface {
	dara.Model
	String() string
	GoString() string
	SetBitrate(v int64) *SubtitleStream
	GetBitrate() *int64
	SetCodecLongName(v string) *SubtitleStream
	GetCodecLongName() *string
	SetCodecName(v string) *SubtitleStream
	GetCodecName() *string
	SetCodecTag(v string) *SubtitleStream
	GetCodecTag() *string
	SetCodecTagString(v string) *SubtitleStream
	GetCodecTagString() *string
	SetContent(v string) *SubtitleStream
	GetContent() *string
	SetDuration(v float64) *SubtitleStream
	GetDuration() *float64
	SetHeight(v int64) *SubtitleStream
	GetHeight() *int64
	SetIndex(v int64) *SubtitleStream
	GetIndex() *int64
	SetLanguage(v string) *SubtitleStream
	GetLanguage() *string
	SetStartTime(v float64) *SubtitleStream
	GetStartTime() *float64
	SetWidth(v int64) *SubtitleStream
	GetWidth() *int64
}

type SubtitleStream struct {
	Bitrate        *int64   `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	CodecLongName  *string  `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	CodecName      *string  `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	CodecTag       *string  `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	CodecTagString *string  `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	Content        *string  `json:"Content,omitempty" xml:"Content,omitempty"`
	Duration       *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Height         *int64   `json:"Height,omitempty" xml:"Height,omitempty"`
	Index          *int64   `json:"Index,omitempty" xml:"Index,omitempty"`
	Language       *string  `json:"Language,omitempty" xml:"Language,omitempty"`
	StartTime      *float64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Width          *int64   `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s SubtitleStream) String() string {
	return dara.Prettify(s)
}

func (s SubtitleStream) GoString() string {
	return s.String()
}

func (s *SubtitleStream) GetBitrate() *int64 {
	return s.Bitrate
}

func (s *SubtitleStream) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *SubtitleStream) GetCodecName() *string {
	return s.CodecName
}

func (s *SubtitleStream) GetCodecTag() *string {
	return s.CodecTag
}

func (s *SubtitleStream) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *SubtitleStream) GetContent() *string {
	return s.Content
}

func (s *SubtitleStream) GetDuration() *float64 {
	return s.Duration
}

func (s *SubtitleStream) GetHeight() *int64 {
	return s.Height
}

func (s *SubtitleStream) GetIndex() *int64 {
	return s.Index
}

func (s *SubtitleStream) GetLanguage() *string {
	return s.Language
}

func (s *SubtitleStream) GetStartTime() *float64 {
	return s.StartTime
}

func (s *SubtitleStream) GetWidth() *int64 {
	return s.Width
}

func (s *SubtitleStream) SetBitrate(v int64) *SubtitleStream {
	s.Bitrate = &v
	return s
}

func (s *SubtitleStream) SetCodecLongName(v string) *SubtitleStream {
	s.CodecLongName = &v
	return s
}

func (s *SubtitleStream) SetCodecName(v string) *SubtitleStream {
	s.CodecName = &v
	return s
}

func (s *SubtitleStream) SetCodecTag(v string) *SubtitleStream {
	s.CodecTag = &v
	return s
}

func (s *SubtitleStream) SetCodecTagString(v string) *SubtitleStream {
	s.CodecTagString = &v
	return s
}

func (s *SubtitleStream) SetContent(v string) *SubtitleStream {
	s.Content = &v
	return s
}

func (s *SubtitleStream) SetDuration(v float64) *SubtitleStream {
	s.Duration = &v
	return s
}

func (s *SubtitleStream) SetHeight(v int64) *SubtitleStream {
	s.Height = &v
	return s
}

func (s *SubtitleStream) SetIndex(v int64) *SubtitleStream {
	s.Index = &v
	return s
}

func (s *SubtitleStream) SetLanguage(v string) *SubtitleStream {
	s.Language = &v
	return s
}

func (s *SubtitleStream) SetStartTime(v float64) *SubtitleStream {
	s.StartTime = &v
	return s
}

func (s *SubtitleStream) SetWidth(v int64) *SubtitleStream {
	s.Width = &v
	return s
}

func (s *SubtitleStream) Validate() error {
	return dara.Validate(s)
}
