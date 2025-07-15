// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveAISubtitleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBgColor(v string) *UpdateLiveAISubtitleRequest
	GetBgColor() *string
	SetBgWidthNormalized(v float32) *UpdateLiveAISubtitleRequest
	GetBgWidthNormalized() *float32
	SetBorderWidthNormalized(v float32) *UpdateLiveAISubtitleRequest
	GetBorderWidthNormalized() *float32
	SetDescription(v string) *UpdateLiveAISubtitleRequest
	GetDescription() *string
	SetDstLanguage(v string) *UpdateLiveAISubtitleRequest
	GetDstLanguage() *string
	SetFontColor(v string) *UpdateLiveAISubtitleRequest
	GetFontColor() *string
	SetFontName(v string) *UpdateLiveAISubtitleRequest
	GetFontName() *string
	SetFontSizeNormalized(v float32) *UpdateLiveAISubtitleRequest
	GetFontSizeNormalized() *float32
	SetHeight(v string) *UpdateLiveAISubtitleRequest
	GetHeight() *string
	SetMaxLines(v int32) *UpdateLiveAISubtitleRequest
	GetMaxLines() *int32
	SetOwnerId(v int64) *UpdateLiveAISubtitleRequest
	GetOwnerId() *int64
	SetPositionNormalized(v []*float32) *UpdateLiveAISubtitleRequest
	GetPositionNormalized() []*float32
	SetRegionId(v string) *UpdateLiveAISubtitleRequest
	GetRegionId() *string
	SetShowSourceLan(v bool) *UpdateLiveAISubtitleRequest
	GetShowSourceLan() *bool
	SetSrcLanguage(v string) *UpdateLiveAISubtitleRequest
	GetSrcLanguage() *string
	SetSubtitleId(v string) *UpdateLiveAISubtitleRequest
	GetSubtitleId() *string
	SetSubtitleName(v string) *UpdateLiveAISubtitleRequest
	GetSubtitleName() *string
	SetWidth(v string) *UpdateLiveAISubtitleRequest
	GetWidth() *string
	SetWordPerLine(v int32) *UpdateLiveAISubtitleRequest
	GetWordPerLine() *int32
}

type UpdateLiveAISubtitleRequest struct {
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
	PositionNormalized []*float32 `json:"PositionNormalized,omitempty" xml:"PositionNormalized,omitempty" type:"Repeated"`
	RegionId           *string    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s UpdateLiveAISubtitleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveAISubtitleRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveAISubtitleRequest) GetBgColor() *string {
	return s.BgColor
}

func (s *UpdateLiveAISubtitleRequest) GetBgWidthNormalized() *float32 {
	return s.BgWidthNormalized
}

func (s *UpdateLiveAISubtitleRequest) GetBorderWidthNormalized() *float32 {
	return s.BorderWidthNormalized
}

func (s *UpdateLiveAISubtitleRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateLiveAISubtitleRequest) GetDstLanguage() *string {
	return s.DstLanguage
}

func (s *UpdateLiveAISubtitleRequest) GetFontColor() *string {
	return s.FontColor
}

func (s *UpdateLiveAISubtitleRequest) GetFontName() *string {
	return s.FontName
}

func (s *UpdateLiveAISubtitleRequest) GetFontSizeNormalized() *float32 {
	return s.FontSizeNormalized
}

func (s *UpdateLiveAISubtitleRequest) GetHeight() *string {
	return s.Height
}

func (s *UpdateLiveAISubtitleRequest) GetMaxLines() *int32 {
	return s.MaxLines
}

func (s *UpdateLiveAISubtitleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateLiveAISubtitleRequest) GetPositionNormalized() []*float32 {
	return s.PositionNormalized
}

func (s *UpdateLiveAISubtitleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateLiveAISubtitleRequest) GetShowSourceLan() *bool {
	return s.ShowSourceLan
}

func (s *UpdateLiveAISubtitleRequest) GetSrcLanguage() *string {
	return s.SrcLanguage
}

func (s *UpdateLiveAISubtitleRequest) GetSubtitleId() *string {
	return s.SubtitleId
}

func (s *UpdateLiveAISubtitleRequest) GetSubtitleName() *string {
	return s.SubtitleName
}

func (s *UpdateLiveAISubtitleRequest) GetWidth() *string {
	return s.Width
}

func (s *UpdateLiveAISubtitleRequest) GetWordPerLine() *int32 {
	return s.WordPerLine
}

func (s *UpdateLiveAISubtitleRequest) SetBgColor(v string) *UpdateLiveAISubtitleRequest {
	s.BgColor = &v
	return s
}

func (s *UpdateLiveAISubtitleRequest) SetBgWidthNormalized(v float32) *UpdateLiveAISubtitleRequest {
	s.BgWidthNormalized = &v
	return s
}

func (s *UpdateLiveAISubtitleRequest) SetBorderWidthNormalized(v float32) *UpdateLiveAISubtitleRequest {
	s.BorderWidthNormalized = &v
	return s
}

func (s *UpdateLiveAISubtitleRequest) SetDescription(v string) *UpdateLiveAISubtitleRequest {
	s.Description = &v
	return s
}

func (s *UpdateLiveAISubtitleRequest) SetDstLanguage(v string) *UpdateLiveAISubtitleRequest {
	s.DstLanguage = &v
	return s
}

func (s *UpdateLiveAISubtitleRequest) SetFontColor(v string) *UpdateLiveAISubtitleRequest {
	s.FontColor = &v
	return s
}

func (s *UpdateLiveAISubtitleRequest) SetFontName(v string) *UpdateLiveAISubtitleRequest {
	s.FontName = &v
	return s
}

func (s *UpdateLiveAISubtitleRequest) SetFontSizeNormalized(v float32) *UpdateLiveAISubtitleRequest {
	s.FontSizeNormalized = &v
	return s
}

func (s *UpdateLiveAISubtitleRequest) SetHeight(v string) *UpdateLiveAISubtitleRequest {
	s.Height = &v
	return s
}

func (s *UpdateLiveAISubtitleRequest) SetMaxLines(v int32) *UpdateLiveAISubtitleRequest {
	s.MaxLines = &v
	return s
}

func (s *UpdateLiveAISubtitleRequest) SetOwnerId(v int64) *UpdateLiveAISubtitleRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateLiveAISubtitleRequest) SetPositionNormalized(v []*float32) *UpdateLiveAISubtitleRequest {
	s.PositionNormalized = v
	return s
}

func (s *UpdateLiveAISubtitleRequest) SetRegionId(v string) *UpdateLiveAISubtitleRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateLiveAISubtitleRequest) SetShowSourceLan(v bool) *UpdateLiveAISubtitleRequest {
	s.ShowSourceLan = &v
	return s
}

func (s *UpdateLiveAISubtitleRequest) SetSrcLanguage(v string) *UpdateLiveAISubtitleRequest {
	s.SrcLanguage = &v
	return s
}

func (s *UpdateLiveAISubtitleRequest) SetSubtitleId(v string) *UpdateLiveAISubtitleRequest {
	s.SubtitleId = &v
	return s
}

func (s *UpdateLiveAISubtitleRequest) SetSubtitleName(v string) *UpdateLiveAISubtitleRequest {
	s.SubtitleName = &v
	return s
}

func (s *UpdateLiveAISubtitleRequest) SetWidth(v string) *UpdateLiveAISubtitleRequest {
	s.Width = &v
	return s
}

func (s *UpdateLiveAISubtitleRequest) SetWordPerLine(v int32) *UpdateLiveAISubtitleRequest {
	s.WordPerLine = &v
	return s
}

func (s *UpdateLiveAISubtitleRequest) Validate() error {
	return dara.Validate(s)
}
