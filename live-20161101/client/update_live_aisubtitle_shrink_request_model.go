// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveAISubtitleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBgColor(v string) *UpdateLiveAISubtitleShrinkRequest
	GetBgColor() *string
	SetBgWidthNormalized(v float32) *UpdateLiveAISubtitleShrinkRequest
	GetBgWidthNormalized() *float32
	SetBorderWidthNormalized(v float32) *UpdateLiveAISubtitleShrinkRequest
	GetBorderWidthNormalized() *float32
	SetDescription(v string) *UpdateLiveAISubtitleShrinkRequest
	GetDescription() *string
	SetDstLanguage(v string) *UpdateLiveAISubtitleShrinkRequest
	GetDstLanguage() *string
	SetFontColor(v string) *UpdateLiveAISubtitleShrinkRequest
	GetFontColor() *string
	SetFontName(v string) *UpdateLiveAISubtitleShrinkRequest
	GetFontName() *string
	SetFontSizeNormalized(v float32) *UpdateLiveAISubtitleShrinkRequest
	GetFontSizeNormalized() *float32
	SetHeight(v string) *UpdateLiveAISubtitleShrinkRequest
	GetHeight() *string
	SetMaxLines(v int32) *UpdateLiveAISubtitleShrinkRequest
	GetMaxLines() *int32
	SetOwnerId(v int64) *UpdateLiveAISubtitleShrinkRequest
	GetOwnerId() *int64
	SetPositionNormalizedShrink(v string) *UpdateLiveAISubtitleShrinkRequest
	GetPositionNormalizedShrink() *string
	SetRegionId(v string) *UpdateLiveAISubtitleShrinkRequest
	GetRegionId() *string
	SetShowSourceLan(v bool) *UpdateLiveAISubtitleShrinkRequest
	GetShowSourceLan() *bool
	SetSrcLanguage(v string) *UpdateLiveAISubtitleShrinkRequest
	GetSrcLanguage() *string
	SetSubtitleId(v string) *UpdateLiveAISubtitleShrinkRequest
	GetSubtitleId() *string
	SetSubtitleName(v string) *UpdateLiveAISubtitleShrinkRequest
	GetSubtitleName() *string
	SetWidth(v string) *UpdateLiveAISubtitleShrinkRequest
	GetWidth() *string
	SetWordPerLine(v int32) *UpdateLiveAISubtitleShrinkRequest
	GetWordPerLine() *int32
}

type UpdateLiveAISubtitleShrinkRequest struct {
	// The background color of the subtitles. Color format: RGBA.
	//
	// example:
	//
	// 0xFF0000
	BgColor *string `json:"BgColor,omitempty" xml:"BgColor,omitempty"`
	// The size of the background box. Valid values: [0,1].
	//
	// example:
	//
	// 0.09
	BgWidthNormalized *float32 `json:"BgWidthNormalized,omitempty" xml:"BgWidthNormalized,omitempty"`
	// The font weight. Valid values: [0,1].
	//
	// example:
	//
	// 0.05
	BorderWidthNormalized *float32 `json:"BorderWidthNormalized,omitempty" xml:"BorderWidthNormalized,omitempty"`
	// The description of the subtitle template. The description can be up to 128 characters in length and can contain letters, digits, and special characters.
	//
	// example:
	//
	// live AI subtitle template
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The target language. Valid values:
	//
	// 	- en-US: English
	//
	// 	- zh-CN: Chinese
	//
	// 	- es-ES: Spanish
	//
	// 	- ru-RU: Russian
	//
	// example:
	//
	// zh-CN
	DstLanguage *string `json:"DstLanguage,omitempty" xml:"DstLanguage,omitempty"`
	// The font color. Color format: RGBA.
	//
	// example:
	//
	// 0xFFFFFF
	FontColor *string `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	// The font. Valid values:
	//
	// 	- KaiTi (default)
	//
	// 	- AlibabaPuHuiTi-Regular
	//
	// 	- AlibabaPuHuiTi-Bold
	//
	// 	- AlibabaPuHuiTi-Light
	//
	// 	- NotoSansHans-Regular
	//
	// 	- NotoSansHans-Bold
	//
	// 	- NotoSansHans-Light
	//
	// example:
	//
	// KaiTi
	FontName *string `json:"FontName,omitempty" xml:"FontName,omitempty"`
	// The font size. Valid values: [0,1].
	//
	// example:
	//
	// 0.037
	FontSizeNormalized *float32 `json:"FontSizeNormalized,omitempty" xml:"FontSizeNormalized,omitempty"`
	// The height of the preview. Unit: pixels.
	//
	// The following preview specifications (width x height) are supported:
	//
	// 	- 360p (640 x 360)
	//
	// 	- 360p (360 x 640)
	//
	// 	- 480p (854 x 480)
	//
	// 	- 480p (480 x 854)
	//
	// 	- 720p (1280 x 720)
	//
	// 	- 720p (720 x 1280)
	//
	// 	- 1080p (1920 x 1080)
	//
	// 	- 1080p (1080 x 1920)
	//
	// example:
	//
	// 720
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The number of displayed lines.
	//
	// example:
	//
	// 2
	MaxLines *int32 `json:"MaxLines,omitempty" xml:"MaxLines,omitempty"`
	OwnerId  *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The position of the subtitles relative to the lower-left corner of the screen. The value is a pair of coordinates.
	PositionNormalizedShrink *string `json:"PositionNormalized,omitempty" xml:"PositionNormalized,omitempty"`
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to display the source language. Valid values: true and false. Default value: false.
	//
	// example:
	//
	// true
	ShowSourceLan *bool `json:"ShowSourceLan,omitempty" xml:"ShowSourceLan,omitempty"`
	// The source language. Valid values:
	//
	// 	- en-US: English
	//
	// 	- zh-CN: Chinese
	//
	// 	- ru-RU: Russian
	//
	// example:
	//
	// zh-CN
	SrcLanguage *string `json:"SrcLanguage,omitempty" xml:"SrcLanguage,omitempty"`
	// The ID of the subtitle template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 597991f3-6ef9-4100-9238-82951de1****
	SubtitleId *string `json:"SubtitleId,omitempty" xml:"SubtitleId,omitempty"`
	// The name of the subtitle template. The name can contain digits, letters, and hyphens (-) but cannot start with a hyphen (-).
	//
	// example:
	//
	// live AI subtitle template
	SubtitleName *string `json:"SubtitleName,omitempty" xml:"SubtitleName,omitempty"`
	// The width of the preview. Unit: pixels.
	//
	// example:
	//
	// 1280
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
	// The number of characters per line. Valid values: 1 to 500.
	//
	// example:
	//
	// 20
	WordPerLine *int32 `json:"WordPerLine,omitempty" xml:"WordPerLine,omitempty"`
}

func (s UpdateLiveAISubtitleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveAISubtitleShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveAISubtitleShrinkRequest) GetBgColor() *string {
	return s.BgColor
}

func (s *UpdateLiveAISubtitleShrinkRequest) GetBgWidthNormalized() *float32 {
	return s.BgWidthNormalized
}

func (s *UpdateLiveAISubtitleShrinkRequest) GetBorderWidthNormalized() *float32 {
	return s.BorderWidthNormalized
}

func (s *UpdateLiveAISubtitleShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateLiveAISubtitleShrinkRequest) GetDstLanguage() *string {
	return s.DstLanguage
}

func (s *UpdateLiveAISubtitleShrinkRequest) GetFontColor() *string {
	return s.FontColor
}

func (s *UpdateLiveAISubtitleShrinkRequest) GetFontName() *string {
	return s.FontName
}

func (s *UpdateLiveAISubtitleShrinkRequest) GetFontSizeNormalized() *float32 {
	return s.FontSizeNormalized
}

func (s *UpdateLiveAISubtitleShrinkRequest) GetHeight() *string {
	return s.Height
}

func (s *UpdateLiveAISubtitleShrinkRequest) GetMaxLines() *int32 {
	return s.MaxLines
}

func (s *UpdateLiveAISubtitleShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateLiveAISubtitleShrinkRequest) GetPositionNormalizedShrink() *string {
	return s.PositionNormalizedShrink
}

func (s *UpdateLiveAISubtitleShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateLiveAISubtitleShrinkRequest) GetShowSourceLan() *bool {
	return s.ShowSourceLan
}

func (s *UpdateLiveAISubtitleShrinkRequest) GetSrcLanguage() *string {
	return s.SrcLanguage
}

func (s *UpdateLiveAISubtitleShrinkRequest) GetSubtitleId() *string {
	return s.SubtitleId
}

func (s *UpdateLiveAISubtitleShrinkRequest) GetSubtitleName() *string {
	return s.SubtitleName
}

func (s *UpdateLiveAISubtitleShrinkRequest) GetWidth() *string {
	return s.Width
}

func (s *UpdateLiveAISubtitleShrinkRequest) GetWordPerLine() *int32 {
	return s.WordPerLine
}

func (s *UpdateLiveAISubtitleShrinkRequest) SetBgColor(v string) *UpdateLiveAISubtitleShrinkRequest {
	s.BgColor = &v
	return s
}

func (s *UpdateLiveAISubtitleShrinkRequest) SetBgWidthNormalized(v float32) *UpdateLiveAISubtitleShrinkRequest {
	s.BgWidthNormalized = &v
	return s
}

func (s *UpdateLiveAISubtitleShrinkRequest) SetBorderWidthNormalized(v float32) *UpdateLiveAISubtitleShrinkRequest {
	s.BorderWidthNormalized = &v
	return s
}

func (s *UpdateLiveAISubtitleShrinkRequest) SetDescription(v string) *UpdateLiveAISubtitleShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateLiveAISubtitleShrinkRequest) SetDstLanguage(v string) *UpdateLiveAISubtitleShrinkRequest {
	s.DstLanguage = &v
	return s
}

func (s *UpdateLiveAISubtitleShrinkRequest) SetFontColor(v string) *UpdateLiveAISubtitleShrinkRequest {
	s.FontColor = &v
	return s
}

func (s *UpdateLiveAISubtitleShrinkRequest) SetFontName(v string) *UpdateLiveAISubtitleShrinkRequest {
	s.FontName = &v
	return s
}

func (s *UpdateLiveAISubtitleShrinkRequest) SetFontSizeNormalized(v float32) *UpdateLiveAISubtitleShrinkRequest {
	s.FontSizeNormalized = &v
	return s
}

func (s *UpdateLiveAISubtitleShrinkRequest) SetHeight(v string) *UpdateLiveAISubtitleShrinkRequest {
	s.Height = &v
	return s
}

func (s *UpdateLiveAISubtitleShrinkRequest) SetMaxLines(v int32) *UpdateLiveAISubtitleShrinkRequest {
	s.MaxLines = &v
	return s
}

func (s *UpdateLiveAISubtitleShrinkRequest) SetOwnerId(v int64) *UpdateLiveAISubtitleShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateLiveAISubtitleShrinkRequest) SetPositionNormalizedShrink(v string) *UpdateLiveAISubtitleShrinkRequest {
	s.PositionNormalizedShrink = &v
	return s
}

func (s *UpdateLiveAISubtitleShrinkRequest) SetRegionId(v string) *UpdateLiveAISubtitleShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateLiveAISubtitleShrinkRequest) SetShowSourceLan(v bool) *UpdateLiveAISubtitleShrinkRequest {
	s.ShowSourceLan = &v
	return s
}

func (s *UpdateLiveAISubtitleShrinkRequest) SetSrcLanguage(v string) *UpdateLiveAISubtitleShrinkRequest {
	s.SrcLanguage = &v
	return s
}

func (s *UpdateLiveAISubtitleShrinkRequest) SetSubtitleId(v string) *UpdateLiveAISubtitleShrinkRequest {
	s.SubtitleId = &v
	return s
}

func (s *UpdateLiveAISubtitleShrinkRequest) SetSubtitleName(v string) *UpdateLiveAISubtitleShrinkRequest {
	s.SubtitleName = &v
	return s
}

func (s *UpdateLiveAISubtitleShrinkRequest) SetWidth(v string) *UpdateLiveAISubtitleShrinkRequest {
	s.Width = &v
	return s
}

func (s *UpdateLiveAISubtitleShrinkRequest) SetWordPerLine(v int32) *UpdateLiveAISubtitleShrinkRequest {
	s.WordPerLine = &v
	return s
}

func (s *UpdateLiveAISubtitleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
