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
	// Configuration for clipping from the source video.
	Clip *MediaConvertJobFeatureClip `json:"Clip,omitempty" xml:"Clip,omitempty" type:"Struct"`
	// A map of key-value pairs to be embedded as container-level metadata in the output file. Provided as a JSON string. Example: {"key1":"value1","key2":"value2"}.
	//
	// 	- Max key length: 64 characters.
	//
	// 	- Max value length: 512 characters.
	//
	// Max 4 key-value pairs.
	Metadata map[string]*string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// Image or text watermarks to add to the video. These parameters override the corresponding settings from a specified watermark template.
	//
	// 	- You can add up to four watermarks to a transcoding task.
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
	// Specifies the order of operations when concatenating multiple files and clipping.
	//
	// 	- true: Clips the first input file before it is concatenated.
	//
	// 	- false: Concatenates all input files first, then applies clipping.
	//
	// 	- Default value: false.
	//
	// example:
	//
	// false
	ConfigToClipFirstPart *string `json:"ConfigToClipFirstPart,omitempty" xml:"ConfigToClipFirstPart,omitempty"`
	// The time range for the clip.
	TimeSpan *MediaConvertJobFeatureClipTimeSpan `json:"TimeSpan,omitempty" xml:"TimeSpan,omitempty" type:"Struct"`
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
	// The duration of the clip, starting from the Seek time. The default duration is from the Seek time to the end of the video. Duration and End are mutually exclusive. If End is set, Duration is ignored.
	//
	// 	- Format: hh:mm:ss[.SSS] or sssss[.SSS].
	//
	// 	- Valid values: [00:00:00.000,23:59:59.999] or [0.000,86399.999].
	//
	// 	- Example: 00:01:59.99 or 180.30.
	//
	// example:
	//
	// 60.0
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Specifies a duration to trim from the end of the video. Duration and End are mutually exclusive. If End is set, Duration is ignored.
	//
	// 	- Format: hh:mm:ss[.SSS] or sssss[.SSS].
	//
	// 	- Valid values: [00:00:00.000,23:59:59.999] or [0.000,86399.999].
	//
	// 	- Example: 00:01:59.99 or 180.30.
	//
	// example:
	//
	// 50
	End *string `json:"End,omitempty" xml:"End,omitempty"`
	// The start time of the clip. It defaults to the beginning of the video.
	//
	// 	- Format: hh:mm:ss[.SSS] or sssss[.SSS].
	//
	// 	- Valid values: [00:00:00.000,23:59:59.999] or [0.000,86399.999].
	//
	// 	- Example: 00:01:59.99 or 180.30.
	//
	// example:
	//
	// 180.30
	Seek *string `json:"Seek,omitempty" xml:"Seek,omitempty"`
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
	// Specifies if the font size adapts to the output resolution. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// 	- Default value: false.
	//
	// example:
	//
	// true
	Adaptive *string `json:"Adaptive,omitempty" xml:"Adaptive,omitempty"`
	// The color of the font border.
	//
	// 	- Default value: Black.
	//
	// example:
	//
	// Black
	BorderColor *string `json:"BorderColor,omitempty" xml:"BorderColor,omitempty"`
	// The width of the font border.
	//
	// 	- Unit: pixels.
	//
	// 	- Valid values: [0,4096].
	//
	// 	- Default value: 0.
	//
	// example:
	//
	// 0
	BorderWidth *string `json:"BorderWidth,omitempty" xml:"BorderWidth,omitempty"`
	// The text to be displayed as the watermark.
	//
	// example:
	//
	// TextWatarmark
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The font opacity.
	//
	// 	- Valid values: (0,1].
	//
	// 	- Default value: 1.0.
	//
	// example:
	//
	// 1.0
	FontAlpha *string `json:"FontAlpha,omitempty" xml:"FontAlpha,omitempty"`
	// The font color of the text watermark.
	//
	// 	- Default value: black.
	//
	// example:
	//
	// black
	FontColor *string `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	// The font of the text watermark.
	//
	// 	- Default value: SimSum.
	//
	// example:
	//
	// SimSum
	FontName *string `json:"FontName,omitempty" xml:"FontName,omitempty"`
	// The font size of the text watermark.
	//
	// 	- Valid values: (4,120).
	//
	// 	- Default value: 16.
	FontSize *string `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	// The height of the image watermark. This parameter overrides the corresponding setting from a specified watermark template. The following value types are supported:
	//
	// 	- Integer: the pixel value of the watermark height.
	//
	//     *
	//
	//     	- Valid values: [8,4096].
	//
	// 	- Decimal: A decimal of the output video\\"s height.
	//
	//     	- Valid values: (0,1).
	//
	//     	- The decimal number can be accurate to four decimal places, such as 0.9999. Excessive digits are automatically discarded.
	//
	// example:
	//
	// 0.1
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The ID of the watermark template.
	//
	// example:
	//
	// 962e1332fa2d4e12bdfb76dd1402fcfa
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The watermark type.
	//
	// 	- Text: a text watermark. In this case, you must specify the parameters related to the text watermark.
	//
	// 	- Image: an image watermark. In this case, you must specify the parameters related to the image watermark.
	//
	// If not specified, the type is inferred from the TemplateId.
	//
	// example:
	//
	// Image
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The width of the image watermark. This parameter overrides the corresponding setting from a specified watermark template. The following value types are supported:
	//
	// 	- Integer: the pixel value of the watermark width.
	//
	//     *
	//
	//     	- Valid values: [8,4096].
	//
	// 	- Decimal: A decimal of the output video\\"s width.
	//
	//     	- Valid values: (0,1).
	//
	//     	- The decimal number can be accurate to four decimal places, such as 0.9999. Excessive digits are automatically discarded.
	//
	// example:
	//
	// 0.1
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
	// The horizontal offset of the image watermark relative to the output video. This parameter overrides the corresponding setting from a specified watermark template. The following value types are supported:
	//
	// 	- Integer: the pixel value of the horizontal offset.
	//
	//     	- Unit: pixels.
	//
	//     	- Valid values: [8,4096].
	//
	// 	- Decimal: the ratio of the horizontal offset to the width of the output video.
	//
	//     	- Valid values: (0,1).
	//
	//     	- The decimal number can be accurate to four decimal places, such as 0.9999. Excessive digits are automatically discarded.
	//
	// example:
	//
	// 0.08
	X *string `json:"X,omitempty" xml:"X,omitempty"`
	// The vertical offset of the image watermark relative to the output video. This parameter overrides the corresponding setting from a specified watermark template. The following value types are supported:
	//
	// 	- Integer: the pixel value of the vertical offset.
	//
	//     	- Unit: pixels.
	//
	//     	- Valid values: [8,4096].
	//
	// 	- Decimal: the ratio of the vertical offset to the height of the output video.
	//
	//     	- Valid values: (0,1).
	//
	//     	- The decimal number can be accurate to four decimal places, such as 0.9999. Excessive digits are automatically discarded.
	//
	// example:
	//
	// 0.08
	Y *string `json:"Y,omitempty" xml:"Y,omitempty"`
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
