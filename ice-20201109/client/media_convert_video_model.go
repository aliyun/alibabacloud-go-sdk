// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMediaConvertVideo interface {
	dara.Model
	String() string
	GoString() string
	SetBitrate(v int32) *MediaConvertVideo
	GetBitrate() *int32
	SetBufsize(v int32) *MediaConvertVideo
	GetBufsize() *int32
	SetCodec(v string) *MediaConvertVideo
	GetCodec() *string
	SetCrf(v interface{}) *MediaConvertVideo
	GetCrf() interface{}
	SetCrop(v string) *MediaConvertVideo
	GetCrop() *string
	SetFps(v interface{}) *MediaConvertVideo
	GetFps() interface{}
	SetGop(v interface{}) *MediaConvertVideo
	GetGop() interface{}
	SetHeight(v int32) *MediaConvertVideo
	GetHeight() *int32
	SetLongShortMode(v bool) *MediaConvertVideo
	GetLongShortMode() *bool
	SetMaxFps(v interface{}) *MediaConvertVideo
	GetMaxFps() interface{}
	SetMaxrate(v int32) *MediaConvertVideo
	GetMaxrate() *int32
	SetPad(v string) *MediaConvertVideo
	GetPad() *string
	SetProfile(v string) *MediaConvertVideo
	GetProfile() *string
	SetQscale(v int32) *MediaConvertVideo
	GetQscale() *int32
	SetRemove(v bool) *MediaConvertVideo
	GetRemove() *bool
	SetScanMode(v string) *MediaConvertVideo
	GetScanMode() *string
	SetWidth(v int32) *MediaConvertVideo
	GetWidth() *int32
}

type MediaConvertVideo struct {
	Bitrate       *int32      `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	Bufsize       *int32      `json:"Bufsize,omitempty" xml:"Bufsize,omitempty"`
	Codec         *string     `json:"Codec,omitempty" xml:"Codec,omitempty"`
	Crf           interface{} `json:"Crf,omitempty" xml:"Crf,omitempty"`
	Crop          *string     `json:"Crop,omitempty" xml:"Crop,omitempty"`
	Fps           interface{} `json:"Fps,omitempty" xml:"Fps,omitempty"`
	Gop           interface{} `json:"Gop,omitempty" xml:"Gop,omitempty"`
	Height        *int32      `json:"Height,omitempty" xml:"Height,omitempty"`
	LongShortMode *bool       `json:"LongShortMode,omitempty" xml:"LongShortMode,omitempty"`
	MaxFps        interface{} `json:"MaxFps,omitempty" xml:"MaxFps,omitempty"`
	Maxrate       *int32      `json:"Maxrate,omitempty" xml:"Maxrate,omitempty"`
	Pad           *string     `json:"Pad,omitempty" xml:"Pad,omitempty"`
	Profile       *string     `json:"Profile,omitempty" xml:"Profile,omitempty"`
	Qscale        *int32      `json:"Qscale,omitempty" xml:"Qscale,omitempty"`
	Remove        *bool       `json:"Remove,omitempty" xml:"Remove,omitempty"`
	ScanMode      *string     `json:"ScanMode,omitempty" xml:"ScanMode,omitempty"`
	Width         *int32      `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s MediaConvertVideo) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertVideo) GoString() string {
	return s.String()
}

func (s *MediaConvertVideo) GetBitrate() *int32 {
	return s.Bitrate
}

func (s *MediaConvertVideo) GetBufsize() *int32 {
	return s.Bufsize
}

func (s *MediaConvertVideo) GetCodec() *string {
	return s.Codec
}

func (s *MediaConvertVideo) GetCrf() interface{} {
	return s.Crf
}

func (s *MediaConvertVideo) GetCrop() *string {
	return s.Crop
}

func (s *MediaConvertVideo) GetFps() interface{} {
	return s.Fps
}

func (s *MediaConvertVideo) GetGop() interface{} {
	return s.Gop
}

func (s *MediaConvertVideo) GetHeight() *int32 {
	return s.Height
}

func (s *MediaConvertVideo) GetLongShortMode() *bool {
	return s.LongShortMode
}

func (s *MediaConvertVideo) GetMaxFps() interface{} {
	return s.MaxFps
}

func (s *MediaConvertVideo) GetMaxrate() *int32 {
	return s.Maxrate
}

func (s *MediaConvertVideo) GetPad() *string {
	return s.Pad
}

func (s *MediaConvertVideo) GetProfile() *string {
	return s.Profile
}

func (s *MediaConvertVideo) GetQscale() *int32 {
	return s.Qscale
}

func (s *MediaConvertVideo) GetRemove() *bool {
	return s.Remove
}

func (s *MediaConvertVideo) GetScanMode() *string {
	return s.ScanMode
}

func (s *MediaConvertVideo) GetWidth() *int32 {
	return s.Width
}

func (s *MediaConvertVideo) SetBitrate(v int32) *MediaConvertVideo {
	s.Bitrate = &v
	return s
}

func (s *MediaConvertVideo) SetBufsize(v int32) *MediaConvertVideo {
	s.Bufsize = &v
	return s
}

func (s *MediaConvertVideo) SetCodec(v string) *MediaConvertVideo {
	s.Codec = &v
	return s
}

func (s *MediaConvertVideo) SetCrf(v interface{}) *MediaConvertVideo {
	s.Crf = v
	return s
}

func (s *MediaConvertVideo) SetCrop(v string) *MediaConvertVideo {
	s.Crop = &v
	return s
}

func (s *MediaConvertVideo) SetFps(v interface{}) *MediaConvertVideo {
	s.Fps = v
	return s
}

func (s *MediaConvertVideo) SetGop(v interface{}) *MediaConvertVideo {
	s.Gop = v
	return s
}

func (s *MediaConvertVideo) SetHeight(v int32) *MediaConvertVideo {
	s.Height = &v
	return s
}

func (s *MediaConvertVideo) SetLongShortMode(v bool) *MediaConvertVideo {
	s.LongShortMode = &v
	return s
}

func (s *MediaConvertVideo) SetMaxFps(v interface{}) *MediaConvertVideo {
	s.MaxFps = v
	return s
}

func (s *MediaConvertVideo) SetMaxrate(v int32) *MediaConvertVideo {
	s.Maxrate = &v
	return s
}

func (s *MediaConvertVideo) SetPad(v string) *MediaConvertVideo {
	s.Pad = &v
	return s
}

func (s *MediaConvertVideo) SetProfile(v string) *MediaConvertVideo {
	s.Profile = &v
	return s
}

func (s *MediaConvertVideo) SetQscale(v int32) *MediaConvertVideo {
	s.Qscale = &v
	return s
}

func (s *MediaConvertVideo) SetRemove(v bool) *MediaConvertVideo {
	s.Remove = &v
	return s
}

func (s *MediaConvertVideo) SetScanMode(v string) *MediaConvertVideo {
	s.ScanMode = &v
	return s
}

func (s *MediaConvertVideo) SetWidth(v int32) *MediaConvertVideo {
	s.Width = &v
	return s
}

func (s *MediaConvertVideo) Validate() error {
	return dara.Validate(s)
}
