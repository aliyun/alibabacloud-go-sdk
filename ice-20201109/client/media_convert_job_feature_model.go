// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMediaConvertJobFeature interface {
	dara.Model
	String() string
	GoString() string
	SetClip(v *MediaConvertJobFeatureClip) *MediaConvertJobFeature
	GetClip() *MediaConvertJobFeatureClip
	SetMetadata(v map[string]*string) *MediaConvertJobFeature
	GetMetadata() map[string]*string
	SetWatermarks(v []*MediaConvertJobFeatureWatermarks) *MediaConvertJobFeature
	GetWatermarks() []*MediaConvertJobFeatureWatermarks
}

type MediaConvertJobFeature struct {
	Clip       *MediaConvertJobFeatureClip         `json:"Clip,omitempty" xml:"Clip,omitempty" type:"Struct"`
	Metadata   map[string]*string                  `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	Watermarks []*MediaConvertJobFeatureWatermarks `json:"Watermarks,omitempty" xml:"Watermarks,omitempty" type:"Repeated"`
}

func (s MediaConvertJobFeature) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertJobFeature) GoString() string {
	return s.String()
}

func (s *MediaConvertJobFeature) GetClip() *MediaConvertJobFeatureClip {
	return s.Clip
}

func (s *MediaConvertJobFeature) GetMetadata() map[string]*string {
	return s.Metadata
}

func (s *MediaConvertJobFeature) GetWatermarks() []*MediaConvertJobFeatureWatermarks {
	return s.Watermarks
}

func (s *MediaConvertJobFeature) SetClip(v *MediaConvertJobFeatureClip) *MediaConvertJobFeature {
	s.Clip = v
	return s
}

func (s *MediaConvertJobFeature) SetMetadata(v map[string]*string) *MediaConvertJobFeature {
	s.Metadata = v
	return s
}

func (s *MediaConvertJobFeature) SetWatermarks(v []*MediaConvertJobFeatureWatermarks) *MediaConvertJobFeature {
	s.Watermarks = v
	return s
}

func (s *MediaConvertJobFeature) Validate() error {
	if s.Clip != nil {
		if err := s.Clip.Validate(); err != nil {
			return err
		}
	}
	if s.Watermarks != nil {
		for _, item := range s.Watermarks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MediaConvertJobFeatureClip struct {
	ConfigToClipFirstPart *string                             `json:"ConfigToClipFirstPart,omitempty" xml:"ConfigToClipFirstPart,omitempty"`
	TimeSpan              *MediaConvertJobFeatureClipTimeSpan `json:"TimeSpan,omitempty" xml:"TimeSpan,omitempty" type:"Struct"`
}

func (s MediaConvertJobFeatureClip) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertJobFeatureClip) GoString() string {
	return s.String()
}

func (s *MediaConvertJobFeatureClip) GetConfigToClipFirstPart() *string {
	return s.ConfigToClipFirstPart
}

func (s *MediaConvertJobFeatureClip) GetTimeSpan() *MediaConvertJobFeatureClipTimeSpan {
	return s.TimeSpan
}

func (s *MediaConvertJobFeatureClip) SetConfigToClipFirstPart(v string) *MediaConvertJobFeatureClip {
	s.ConfigToClipFirstPart = &v
	return s
}

func (s *MediaConvertJobFeatureClip) SetTimeSpan(v *MediaConvertJobFeatureClipTimeSpan) *MediaConvertJobFeatureClip {
	s.TimeSpan = v
	return s
}

func (s *MediaConvertJobFeatureClip) Validate() error {
	if s.TimeSpan != nil {
		if err := s.TimeSpan.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MediaConvertJobFeatureClipTimeSpan struct {
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	End      *string `json:"End,omitempty" xml:"End,omitempty"`
	Seek     *string `json:"Seek,omitempty" xml:"Seek,omitempty"`
}

func (s MediaConvertJobFeatureClipTimeSpan) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertJobFeatureClipTimeSpan) GoString() string {
	return s.String()
}

func (s *MediaConvertJobFeatureClipTimeSpan) GetDuration() *string {
	return s.Duration
}

func (s *MediaConvertJobFeatureClipTimeSpan) GetEnd() *string {
	return s.End
}

func (s *MediaConvertJobFeatureClipTimeSpan) GetSeek() *string {
	return s.Seek
}

func (s *MediaConvertJobFeatureClipTimeSpan) SetDuration(v string) *MediaConvertJobFeatureClipTimeSpan {
	s.Duration = &v
	return s
}

func (s *MediaConvertJobFeatureClipTimeSpan) SetEnd(v string) *MediaConvertJobFeatureClipTimeSpan {
	s.End = &v
	return s
}

func (s *MediaConvertJobFeatureClipTimeSpan) SetSeek(v string) *MediaConvertJobFeatureClipTimeSpan {
	s.Seek = &v
	return s
}

func (s *MediaConvertJobFeatureClipTimeSpan) Validate() error {
	return dara.Validate(s)
}

type MediaConvertJobFeatureWatermarks struct {
	Adaptive    *string `json:"Adaptive,omitempty" xml:"Adaptive,omitempty"`
	BorderColor *string `json:"BorderColor,omitempty" xml:"BorderColor,omitempty"`
	BorderWidth *string `json:"BorderWidth,omitempty" xml:"BorderWidth,omitempty"`
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	FontAlpha   *string `json:"FontAlpha,omitempty" xml:"FontAlpha,omitempty"`
	FontColor   *string `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	FontName    *string `json:"FontName,omitempty" xml:"FontName,omitempty"`
	FontSize    *string `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	Height      *string `json:"Height,omitempty" xml:"Height,omitempty"`
	TemplateId  *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Width       *string `json:"Width,omitempty" xml:"Width,omitempty"`
	X           *string `json:"X,omitempty" xml:"X,omitempty"`
	Y           *string `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s MediaConvertJobFeatureWatermarks) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertJobFeatureWatermarks) GoString() string {
	return s.String()
}

func (s *MediaConvertJobFeatureWatermarks) GetAdaptive() *string {
	return s.Adaptive
}

func (s *MediaConvertJobFeatureWatermarks) GetBorderColor() *string {
	return s.BorderColor
}

func (s *MediaConvertJobFeatureWatermarks) GetBorderWidth() *string {
	return s.BorderWidth
}

func (s *MediaConvertJobFeatureWatermarks) GetContent() *string {
	return s.Content
}

func (s *MediaConvertJobFeatureWatermarks) GetFontAlpha() *string {
	return s.FontAlpha
}

func (s *MediaConvertJobFeatureWatermarks) GetFontColor() *string {
	return s.FontColor
}

func (s *MediaConvertJobFeatureWatermarks) GetFontName() *string {
	return s.FontName
}

func (s *MediaConvertJobFeatureWatermarks) GetFontSize() *string {
	return s.FontSize
}

func (s *MediaConvertJobFeatureWatermarks) GetHeight() *string {
	return s.Height
}

func (s *MediaConvertJobFeatureWatermarks) GetTemplateId() *string {
	return s.TemplateId
}

func (s *MediaConvertJobFeatureWatermarks) GetType() *string {
	return s.Type
}

func (s *MediaConvertJobFeatureWatermarks) GetWidth() *string {
	return s.Width
}

func (s *MediaConvertJobFeatureWatermarks) GetX() *string {
	return s.X
}

func (s *MediaConvertJobFeatureWatermarks) GetY() *string {
	return s.Y
}

func (s *MediaConvertJobFeatureWatermarks) SetAdaptive(v string) *MediaConvertJobFeatureWatermarks {
	s.Adaptive = &v
	return s
}

func (s *MediaConvertJobFeatureWatermarks) SetBorderColor(v string) *MediaConvertJobFeatureWatermarks {
	s.BorderColor = &v
	return s
}

func (s *MediaConvertJobFeatureWatermarks) SetBorderWidth(v string) *MediaConvertJobFeatureWatermarks {
	s.BorderWidth = &v
	return s
}

func (s *MediaConvertJobFeatureWatermarks) SetContent(v string) *MediaConvertJobFeatureWatermarks {
	s.Content = &v
	return s
}

func (s *MediaConvertJobFeatureWatermarks) SetFontAlpha(v string) *MediaConvertJobFeatureWatermarks {
	s.FontAlpha = &v
	return s
}

func (s *MediaConvertJobFeatureWatermarks) SetFontColor(v string) *MediaConvertJobFeatureWatermarks {
	s.FontColor = &v
	return s
}

func (s *MediaConvertJobFeatureWatermarks) SetFontName(v string) *MediaConvertJobFeatureWatermarks {
	s.FontName = &v
	return s
}

func (s *MediaConvertJobFeatureWatermarks) SetFontSize(v string) *MediaConvertJobFeatureWatermarks {
	s.FontSize = &v
	return s
}

func (s *MediaConvertJobFeatureWatermarks) SetHeight(v string) *MediaConvertJobFeatureWatermarks {
	s.Height = &v
	return s
}

func (s *MediaConvertJobFeatureWatermarks) SetTemplateId(v string) *MediaConvertJobFeatureWatermarks {
	s.TemplateId = &v
	return s
}

func (s *MediaConvertJobFeatureWatermarks) SetType(v string) *MediaConvertJobFeatureWatermarks {
	s.Type = &v
	return s
}

func (s *MediaConvertJobFeatureWatermarks) SetWidth(v string) *MediaConvertJobFeatureWatermarks {
	s.Width = &v
	return s
}

func (s *MediaConvertJobFeatureWatermarks) SetX(v string) *MediaConvertJobFeatureWatermarks {
	s.X = &v
	return s
}

func (s *MediaConvertJobFeatureWatermarks) SetY(v string) *MediaConvertJobFeatureWatermarks {
	s.Y = &v
	return s
}

func (s *MediaConvertJobFeatureWatermarks) Validate() error {
	return dara.Validate(s)
}
