// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveAISubtitleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBgColor(v string) *AddLiveAISubtitleShrinkRequest
	GetBgColor() *string
	SetBgWidthNormalized(v float32) *AddLiveAISubtitleShrinkRequest
	GetBgWidthNormalized() *float32
	SetBorderWidthNormalized(v float32) *AddLiveAISubtitleShrinkRequest
	GetBorderWidthNormalized() *float32
	SetCopyFrom(v string) *AddLiveAISubtitleShrinkRequest
	GetCopyFrom() *string
	SetDescription(v string) *AddLiveAISubtitleShrinkRequest
	GetDescription() *string
	SetDstLanguage(v string) *AddLiveAISubtitleShrinkRequest
	GetDstLanguage() *string
	SetFontColor(v string) *AddLiveAISubtitleShrinkRequest
	GetFontColor() *string
	SetFontName(v string) *AddLiveAISubtitleShrinkRequest
	GetFontName() *string
	SetFontSizeNormalized(v float32) *AddLiveAISubtitleShrinkRequest
	GetFontSizeNormalized() *float32
	SetHeight(v string) *AddLiveAISubtitleShrinkRequest
	GetHeight() *string
	SetMaxLines(v int32) *AddLiveAISubtitleShrinkRequest
	GetMaxLines() *int32
	SetOwnerId(v int64) *AddLiveAISubtitleShrinkRequest
	GetOwnerId() *int64
	SetPositionNormalizedShrink(v string) *AddLiveAISubtitleShrinkRequest
	GetPositionNormalizedShrink() *string
	SetRegionId(v string) *AddLiveAISubtitleShrinkRequest
	GetRegionId() *string
	SetShowSourceLan(v bool) *AddLiveAISubtitleShrinkRequest
	GetShowSourceLan() *bool
	SetSrcLanguage(v string) *AddLiveAISubtitleShrinkRequest
	GetSrcLanguage() *string
	SetSubtitleName(v string) *AddLiveAISubtitleShrinkRequest
	GetSubtitleName() *string
	SetWidth(v string) *AddLiveAISubtitleShrinkRequest
	GetWidth() *string
	SetWordPerLine(v int32) *AddLiveAISubtitleShrinkRequest
	GetWordPerLine() *int32
}

type AddLiveAISubtitleShrinkRequest struct {
	// The background color of the subtitles, which is an RGBA value.
	//
	// example:
	//
	// 0xFF0000
	BgColor *string `json:"BgColor,omitempty" xml:"BgColor,omitempty"`
	// The background size of the subtitles. Valid values: [0,1].
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
	// The subtitle template that you copy. Set the value to the name of the subtitle template.
	//
	// example:
	//
	// sub01
	CopyFrom *string `json:"CopyFrom,omitempty" xml:"CopyFrom,omitempty"`
	// The custom description of the subtitle template. The description can be up to 128 characters in length and can contain letters, digits, and special characters.
	//
	// example:
	//
	// live AI subtitle template
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The target language. Valid values:
	//
	//  - en-US: English
	//
	// - zh-CN: Chinese
	//
	// - es-ES: Spanish
	//
	// - ru-RU: Russian
	//
	// example:
	//
	// zh-CN
	DstLanguage *string `json:"DstLanguage,omitempty" xml:"DstLanguage,omitempty"`
	// The font color, which is an RGBA value.
	//
	// example:
	//
	// 0xFFFFFF
	FontColor *string `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	// The font. Valid values:
	//
	// - KaiTi (default)
	//
	// - AlibabaPuHuiTi-Regular
	//
	// - AlibabaPuHuiTi-Bold
	//
	// - AlibabaPuHuiTi-Light
	//
	// - NotoSansHans-Regular
	//
	// - NotoSansHans-Bold
	//
	// - NotoSansHans-Light
	//
	// example:
	//
	// KaiTi
	FontName *string `json:"FontName,omitempty" xml:"FontName,omitempty"`
	// The font size. Valid values: [0,1].
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.037
	FontSizeNormalized *float32 `json:"FontSizeNormalized,omitempty" xml:"FontSizeNormalized,omitempty"`
	// The preview height. Unit: pixels.
	//
	// The following specifications of preview width × preview height are supported:
	//
	// - Landscape low definition 360p (640×360)
	//
	// - Portrait low definition 360p (360×640)
	//
	// - Landscape standard definition 480p (854×480)
	//
	// - Portrait standard definition 480p (480×854)
	//
	// - Landscape high definition 720p (1280×720)
	//
	// - Portrait high definition 720p (720×1280)
	//
	// - Landscape ultra-high definition 1080p (1920×1080)
	//
	// - Portrait ultra-high definition 1080p (1080×1920)
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
	// The position of the subtitles. The value is a pair of coordinates for which the origin of the x and y axes is the lower-left corner of the screen.
	//
	// This parameter is required.
	PositionNormalizedShrink *string `json:"PositionNormalized,omitempty" xml:"PositionNormalized,omitempty"`
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to display the source language. Default value: false.
	//
	// example:
	//
	// true
	ShowSourceLan *bool `json:"ShowSourceLan,omitempty" xml:"ShowSourceLan,omitempty"`
	// The source language. Valid values:
	//
	//  - en-US: English
	//
	// - zh-CN: Chinese
	//
	// - ru-RU: Russian
	//
	// This parameter is required.
	//
	// example:
	//
	// zh-CN
	SrcLanguage *string `json:"SrcLanguage,omitempty" xml:"SrcLanguage,omitempty"`
	// The name of the subtitle template. The name can contain only digits, letters, and hyphens (-). The name cannot start with a hyphen.
	//
	// This parameter is required.
	//
	// example:
	//
	// sub01
	SubtitleName *string `json:"SubtitleName,omitempty" xml:"SubtitleName,omitempty"`
	// The preview width. Unit: pixels.
	//
	// example:
	//
	// 1280
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
	// The number of words displayed per line. Valid values: integers from 1 to 500.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	WordPerLine *int32 `json:"WordPerLine,omitempty" xml:"WordPerLine,omitempty"`
}

func (s AddLiveAISubtitleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddLiveAISubtitleShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddLiveAISubtitleShrinkRequest) GetBgColor() *string {
	return s.BgColor
}

func (s *AddLiveAISubtitleShrinkRequest) GetBgWidthNormalized() *float32 {
	return s.BgWidthNormalized
}

func (s *AddLiveAISubtitleShrinkRequest) GetBorderWidthNormalized() *float32 {
	return s.BorderWidthNormalized
}

func (s *AddLiveAISubtitleShrinkRequest) GetCopyFrom() *string {
	return s.CopyFrom
}

func (s *AddLiveAISubtitleShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *AddLiveAISubtitleShrinkRequest) GetDstLanguage() *string {
	return s.DstLanguage
}

func (s *AddLiveAISubtitleShrinkRequest) GetFontColor() *string {
	return s.FontColor
}

func (s *AddLiveAISubtitleShrinkRequest) GetFontName() *string {
	return s.FontName
}

func (s *AddLiveAISubtitleShrinkRequest) GetFontSizeNormalized() *float32 {
	return s.FontSizeNormalized
}

func (s *AddLiveAISubtitleShrinkRequest) GetHeight() *string {
	return s.Height
}

func (s *AddLiveAISubtitleShrinkRequest) GetMaxLines() *int32 {
	return s.MaxLines
}

func (s *AddLiveAISubtitleShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddLiveAISubtitleShrinkRequest) GetPositionNormalizedShrink() *string {
	return s.PositionNormalizedShrink
}

func (s *AddLiveAISubtitleShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddLiveAISubtitleShrinkRequest) GetShowSourceLan() *bool {
	return s.ShowSourceLan
}

func (s *AddLiveAISubtitleShrinkRequest) GetSrcLanguage() *string {
	return s.SrcLanguage
}

func (s *AddLiveAISubtitleShrinkRequest) GetSubtitleName() *string {
	return s.SubtitleName
}

func (s *AddLiveAISubtitleShrinkRequest) GetWidth() *string {
	return s.Width
}

func (s *AddLiveAISubtitleShrinkRequest) GetWordPerLine() *int32 {
	return s.WordPerLine
}

func (s *AddLiveAISubtitleShrinkRequest) SetBgColor(v string) *AddLiveAISubtitleShrinkRequest {
	s.BgColor = &v
	return s
}

func (s *AddLiveAISubtitleShrinkRequest) SetBgWidthNormalized(v float32) *AddLiveAISubtitleShrinkRequest {
	s.BgWidthNormalized = &v
	return s
}

func (s *AddLiveAISubtitleShrinkRequest) SetBorderWidthNormalized(v float32) *AddLiveAISubtitleShrinkRequest {
	s.BorderWidthNormalized = &v
	return s
}

func (s *AddLiveAISubtitleShrinkRequest) SetCopyFrom(v string) *AddLiveAISubtitleShrinkRequest {
	s.CopyFrom = &v
	return s
}

func (s *AddLiveAISubtitleShrinkRequest) SetDescription(v string) *AddLiveAISubtitleShrinkRequest {
	s.Description = &v
	return s
}

func (s *AddLiveAISubtitleShrinkRequest) SetDstLanguage(v string) *AddLiveAISubtitleShrinkRequest {
	s.DstLanguage = &v
	return s
}

func (s *AddLiveAISubtitleShrinkRequest) SetFontColor(v string) *AddLiveAISubtitleShrinkRequest {
	s.FontColor = &v
	return s
}

func (s *AddLiveAISubtitleShrinkRequest) SetFontName(v string) *AddLiveAISubtitleShrinkRequest {
	s.FontName = &v
	return s
}

func (s *AddLiveAISubtitleShrinkRequest) SetFontSizeNormalized(v float32) *AddLiveAISubtitleShrinkRequest {
	s.FontSizeNormalized = &v
	return s
}

func (s *AddLiveAISubtitleShrinkRequest) SetHeight(v string) *AddLiveAISubtitleShrinkRequest {
	s.Height = &v
	return s
}

func (s *AddLiveAISubtitleShrinkRequest) SetMaxLines(v int32) *AddLiveAISubtitleShrinkRequest {
	s.MaxLines = &v
	return s
}

func (s *AddLiveAISubtitleShrinkRequest) SetOwnerId(v int64) *AddLiveAISubtitleShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *AddLiveAISubtitleShrinkRequest) SetPositionNormalizedShrink(v string) *AddLiveAISubtitleShrinkRequest {
	s.PositionNormalizedShrink = &v
	return s
}

func (s *AddLiveAISubtitleShrinkRequest) SetRegionId(v string) *AddLiveAISubtitleShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *AddLiveAISubtitleShrinkRequest) SetShowSourceLan(v bool) *AddLiveAISubtitleShrinkRequest {
	s.ShowSourceLan = &v
	return s
}

func (s *AddLiveAISubtitleShrinkRequest) SetSrcLanguage(v string) *AddLiveAISubtitleShrinkRequest {
	s.SrcLanguage = &v
	return s
}

func (s *AddLiveAISubtitleShrinkRequest) SetSubtitleName(v string) *AddLiveAISubtitleShrinkRequest {
	s.SubtitleName = &v
	return s
}

func (s *AddLiveAISubtitleShrinkRequest) SetWidth(v string) *AddLiveAISubtitleShrinkRequest {
	s.Width = &v
	return s
}

func (s *AddLiveAISubtitleShrinkRequest) SetWordPerLine(v int32) *AddLiveAISubtitleShrinkRequest {
	s.WordPerLine = &v
	return s
}

func (s *AddLiveAISubtitleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
