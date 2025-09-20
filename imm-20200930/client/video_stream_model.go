// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoStream interface {
	dara.Model
	String() string
	GoString() string
	SetAverageFrameRate(v string) *VideoStream
	GetAverageFrameRate() *string
	SetBitDepth(v int64) *VideoStream
	GetBitDepth() *int64
	SetBitrate(v int64) *VideoStream
	GetBitrate() *int64
	SetCodecLongName(v string) *VideoStream
	GetCodecLongName() *string
	SetCodecName(v string) *VideoStream
	GetCodecName() *string
	SetCodecTag(v string) *VideoStream
	GetCodecTag() *string
	SetCodecTagString(v string) *VideoStream
	GetCodecTagString() *string
	SetCodecTimeBase(v string) *VideoStream
	GetCodecTimeBase() *string
	SetColorPrimaries(v string) *VideoStream
	GetColorPrimaries() *string
	SetColorRange(v string) *VideoStream
	GetColorRange() *string
	SetColorSpace(v string) *VideoStream
	GetColorSpace() *string
	SetColorTransfer(v string) *VideoStream
	GetColorTransfer() *string
	SetDisplayAspectRatio(v string) *VideoStream
	GetDisplayAspectRatio() *string
	SetDuration(v float64) *VideoStream
	GetDuration() *float64
	SetFrameCount(v int64) *VideoStream
	GetFrameCount() *int64
	SetFrameRate(v string) *VideoStream
	GetFrameRate() *string
	SetHasBFrames(v int64) *VideoStream
	GetHasBFrames() *int64
	SetHeight(v int64) *VideoStream
	GetHeight() *int64
	SetIndex(v int64) *VideoStream
	GetIndex() *int64
	SetLanguage(v string) *VideoStream
	GetLanguage() *string
	SetLevel(v int64) *VideoStream
	GetLevel() *int64
	SetPixelFormat(v string) *VideoStream
	GetPixelFormat() *string
	SetProfile(v string) *VideoStream
	GetProfile() *string
	SetRotate(v string) *VideoStream
	GetRotate() *string
	SetSampleAspectRatio(v string) *VideoStream
	GetSampleAspectRatio() *string
	SetStartTime(v float64) *VideoStream
	GetStartTime() *float64
	SetTimeBase(v string) *VideoStream
	GetTimeBase() *string
	SetWidth(v int64) *VideoStream
	GetWidth() *int64
}

type VideoStream struct {
	AverageFrameRate   *string  `json:"AverageFrameRate,omitempty" xml:"AverageFrameRate,omitempty"`
	BitDepth           *int64   `json:"BitDepth,omitempty" xml:"BitDepth,omitempty"`
	Bitrate            *int64   `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	CodecLongName      *string  `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	CodecName          *string  `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	CodecTag           *string  `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	CodecTagString     *string  `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	CodecTimeBase      *string  `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	ColorPrimaries     *string  `json:"ColorPrimaries,omitempty" xml:"ColorPrimaries,omitempty"`
	ColorRange         *string  `json:"ColorRange,omitempty" xml:"ColorRange,omitempty"`
	ColorSpace         *string  `json:"ColorSpace,omitempty" xml:"ColorSpace,omitempty"`
	ColorTransfer      *string  `json:"ColorTransfer,omitempty" xml:"ColorTransfer,omitempty"`
	DisplayAspectRatio *string  `json:"DisplayAspectRatio,omitempty" xml:"DisplayAspectRatio,omitempty"`
	Duration           *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	FrameCount         *int64   `json:"FrameCount,omitempty" xml:"FrameCount,omitempty"`
	FrameRate          *string  `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	HasBFrames         *int64   `json:"HasBFrames,omitempty" xml:"HasBFrames,omitempty"`
	Height             *int64   `json:"Height,omitempty" xml:"Height,omitempty"`
	Index              *int64   `json:"Index,omitempty" xml:"Index,omitempty"`
	Language           *string  `json:"Language,omitempty" xml:"Language,omitempty"`
	Level              *int64   `json:"Level,omitempty" xml:"Level,omitempty"`
	PixelFormat        *string  `json:"PixelFormat,omitempty" xml:"PixelFormat,omitempty"`
	Profile            *string  `json:"Profile,omitempty" xml:"Profile,omitempty"`
	Rotate             *string  `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	SampleAspectRatio  *string  `json:"SampleAspectRatio,omitempty" xml:"SampleAspectRatio,omitempty"`
	StartTime          *float64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TimeBase           *string  `json:"TimeBase,omitempty" xml:"TimeBase,omitempty"`
	Width              *int64   `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s VideoStream) String() string {
	return dara.Prettify(s)
}

func (s VideoStream) GoString() string {
	return s.String()
}

func (s *VideoStream) GetAverageFrameRate() *string {
	return s.AverageFrameRate
}

func (s *VideoStream) GetBitDepth() *int64 {
	return s.BitDepth
}

func (s *VideoStream) GetBitrate() *int64 {
	return s.Bitrate
}

func (s *VideoStream) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *VideoStream) GetCodecName() *string {
	return s.CodecName
}

func (s *VideoStream) GetCodecTag() *string {
	return s.CodecTag
}

func (s *VideoStream) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *VideoStream) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *VideoStream) GetColorPrimaries() *string {
	return s.ColorPrimaries
}

func (s *VideoStream) GetColorRange() *string {
	return s.ColorRange
}

func (s *VideoStream) GetColorSpace() *string {
	return s.ColorSpace
}

func (s *VideoStream) GetColorTransfer() *string {
	return s.ColorTransfer
}

func (s *VideoStream) GetDisplayAspectRatio() *string {
	return s.DisplayAspectRatio
}

func (s *VideoStream) GetDuration() *float64 {
	return s.Duration
}

func (s *VideoStream) GetFrameCount() *int64 {
	return s.FrameCount
}

func (s *VideoStream) GetFrameRate() *string {
	return s.FrameRate
}

func (s *VideoStream) GetHasBFrames() *int64 {
	return s.HasBFrames
}

func (s *VideoStream) GetHeight() *int64 {
	return s.Height
}

func (s *VideoStream) GetIndex() *int64 {
	return s.Index
}

func (s *VideoStream) GetLanguage() *string {
	return s.Language
}

func (s *VideoStream) GetLevel() *int64 {
	return s.Level
}

func (s *VideoStream) GetPixelFormat() *string {
	return s.PixelFormat
}

func (s *VideoStream) GetProfile() *string {
	return s.Profile
}

func (s *VideoStream) GetRotate() *string {
	return s.Rotate
}

func (s *VideoStream) GetSampleAspectRatio() *string {
	return s.SampleAspectRatio
}

func (s *VideoStream) GetStartTime() *float64 {
	return s.StartTime
}

func (s *VideoStream) GetTimeBase() *string {
	return s.TimeBase
}

func (s *VideoStream) GetWidth() *int64 {
	return s.Width
}

func (s *VideoStream) SetAverageFrameRate(v string) *VideoStream {
	s.AverageFrameRate = &v
	return s
}

func (s *VideoStream) SetBitDepth(v int64) *VideoStream {
	s.BitDepth = &v
	return s
}

func (s *VideoStream) SetBitrate(v int64) *VideoStream {
	s.Bitrate = &v
	return s
}

func (s *VideoStream) SetCodecLongName(v string) *VideoStream {
	s.CodecLongName = &v
	return s
}

func (s *VideoStream) SetCodecName(v string) *VideoStream {
	s.CodecName = &v
	return s
}

func (s *VideoStream) SetCodecTag(v string) *VideoStream {
	s.CodecTag = &v
	return s
}

func (s *VideoStream) SetCodecTagString(v string) *VideoStream {
	s.CodecTagString = &v
	return s
}

func (s *VideoStream) SetCodecTimeBase(v string) *VideoStream {
	s.CodecTimeBase = &v
	return s
}

func (s *VideoStream) SetColorPrimaries(v string) *VideoStream {
	s.ColorPrimaries = &v
	return s
}

func (s *VideoStream) SetColorRange(v string) *VideoStream {
	s.ColorRange = &v
	return s
}

func (s *VideoStream) SetColorSpace(v string) *VideoStream {
	s.ColorSpace = &v
	return s
}

func (s *VideoStream) SetColorTransfer(v string) *VideoStream {
	s.ColorTransfer = &v
	return s
}

func (s *VideoStream) SetDisplayAspectRatio(v string) *VideoStream {
	s.DisplayAspectRatio = &v
	return s
}

func (s *VideoStream) SetDuration(v float64) *VideoStream {
	s.Duration = &v
	return s
}

func (s *VideoStream) SetFrameCount(v int64) *VideoStream {
	s.FrameCount = &v
	return s
}

func (s *VideoStream) SetFrameRate(v string) *VideoStream {
	s.FrameRate = &v
	return s
}

func (s *VideoStream) SetHasBFrames(v int64) *VideoStream {
	s.HasBFrames = &v
	return s
}

func (s *VideoStream) SetHeight(v int64) *VideoStream {
	s.Height = &v
	return s
}

func (s *VideoStream) SetIndex(v int64) *VideoStream {
	s.Index = &v
	return s
}

func (s *VideoStream) SetLanguage(v string) *VideoStream {
	s.Language = &v
	return s
}

func (s *VideoStream) SetLevel(v int64) *VideoStream {
	s.Level = &v
	return s
}

func (s *VideoStream) SetPixelFormat(v string) *VideoStream {
	s.PixelFormat = &v
	return s
}

func (s *VideoStream) SetProfile(v string) *VideoStream {
	s.Profile = &v
	return s
}

func (s *VideoStream) SetRotate(v string) *VideoStream {
	s.Rotate = &v
	return s
}

func (s *VideoStream) SetSampleAspectRatio(v string) *VideoStream {
	s.SampleAspectRatio = &v
	return s
}

func (s *VideoStream) SetStartTime(v float64) *VideoStream {
	s.StartTime = &v
	return s
}

func (s *VideoStream) SetTimeBase(v string) *VideoStream {
	s.TimeBase = &v
	return s
}

func (s *VideoStream) SetWidth(v int64) *VideoStream {
	s.Width = &v
	return s
}

func (s *VideoStream) Validate() error {
	return dara.Validate(s)
}
