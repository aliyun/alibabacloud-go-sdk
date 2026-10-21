// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoTranslationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCapabilitiesShrink(v string) *VideoTranslationShrinkRequest
	GetCapabilitiesShrink() *string
	SetEnableQualityCheck(v bool) *VideoTranslationShrinkRequest
	GetEnableQualityCheck() *bool
	SetEraseOnly(v bool) *VideoTranslationShrinkRequest
	GetEraseOnly() *bool
	SetRemoveLogo(v bool) *VideoTranslationShrinkRequest
	GetRemoveLogo() *bool
	SetSourceLanguage(v string) *VideoTranslationShrinkRequest
	GetSourceLanguage() *string
	SetSubtitleMode(v string) *VideoTranslationShrinkRequest
	GetSubtitleMode() *string
	SetTargetLanguage(v string) *VideoTranslationShrinkRequest
	GetTargetLanguage() *string
	SetVideoUrl(v string) *VideoTranslationShrinkRequest
	GetVideoUrl() *string
}

type VideoTranslationShrinkRequest struct {
	// The list of translation features. Valid values: speech and visual. You can select both. At least one feature is required. Duplicate values are not accepted.
	//
	// This parameter is required.
	CapabilitiesShrink *string `json:"Capabilities,omitempty" xml:"Capabilities,omitempty"`
	// Specifies whether to enable quality inspection for on-screen text processing. Effective only when visual is included and this parameter is explicitly set to true. If only speech is selected, this parameter is ignored, no quality inspection is performed, no quality inspection fields are returned, and no error is reported due to this parameter. Enabling quality inspection may increase the processing time. Do not enable it unless necessary. Default value: false.
	//
	// example:
	//
	// false
	EnableQualityCheck *bool `json:"EnableQualityCheck,omitempty" xml:"EnableQualityCheck,omitempty"`
	// Applicable only when visual is included. Specifies whether to erase on-screen text without translation. Valid values: true: erases on-screen text without translation. false: translates on-screen text. Default value: false.
	//
	// example:
	//
	// false
	EraseOnly *bool `json:"EraseOnly,omitempty" xml:"EraseOnly,omitempty"`
	// Applicable only when visual is included. Specifies whether to erase brand logos. Valid values: true: erases brand logos. false: retains brand logos. Default value: false. This parameter is independent of EraseOnly.
	//
	// example:
	//
	// false
	RemoveLogo *bool `json:"RemoveLogo,omitempty" xml:"RemoveLogo,omitempty"`
	// The original language of the video (source language). For supported languages, see Language support.
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// Applicable only when speech is included. Valid values: erase: erases the original subtitles only. replace: replaces the original subtitles with translated subtitles, which means erasing the original subtitles and adding subtitles synchronized with the translated speech. If not specified, the original subtitles are not processed, which means neither erasing the original subtitles nor adding new subtitles.
	//
	// example:
	//
	// replace
	SubtitleMode *string `json:"SubtitleMode,omitempty" xml:"SubtitleMode,omitempty"`
	// Required when speech is included. Optional when only visual is included and EraseOnly is set to true. If specified in this case, it is ignored. Required in other cases.
	//
	// example:
	//
	// en
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	// The video URL. Supported formats: MP4 and MOV. Maximum size: 200 MB.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://cdn.example.com/product-demo.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s VideoTranslationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s VideoTranslationShrinkRequest) GoString() string {
	return s.String()
}

func (s *VideoTranslationShrinkRequest) GetCapabilitiesShrink() *string {
	return s.CapabilitiesShrink
}

func (s *VideoTranslationShrinkRequest) GetEnableQualityCheck() *bool {
	return s.EnableQualityCheck
}

func (s *VideoTranslationShrinkRequest) GetEraseOnly() *bool {
	return s.EraseOnly
}

func (s *VideoTranslationShrinkRequest) GetRemoveLogo() *bool {
	return s.RemoveLogo
}

func (s *VideoTranslationShrinkRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *VideoTranslationShrinkRequest) GetSubtitleMode() *string {
	return s.SubtitleMode
}

func (s *VideoTranslationShrinkRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *VideoTranslationShrinkRequest) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *VideoTranslationShrinkRequest) SetCapabilitiesShrink(v string) *VideoTranslationShrinkRequest {
	s.CapabilitiesShrink = &v
	return s
}

func (s *VideoTranslationShrinkRequest) SetEnableQualityCheck(v bool) *VideoTranslationShrinkRequest {
	s.EnableQualityCheck = &v
	return s
}

func (s *VideoTranslationShrinkRequest) SetEraseOnly(v bool) *VideoTranslationShrinkRequest {
	s.EraseOnly = &v
	return s
}

func (s *VideoTranslationShrinkRequest) SetRemoveLogo(v bool) *VideoTranslationShrinkRequest {
	s.RemoveLogo = &v
	return s
}

func (s *VideoTranslationShrinkRequest) SetSourceLanguage(v string) *VideoTranslationShrinkRequest {
	s.SourceLanguage = &v
	return s
}

func (s *VideoTranslationShrinkRequest) SetSubtitleMode(v string) *VideoTranslationShrinkRequest {
	s.SubtitleMode = &v
	return s
}

func (s *VideoTranslationShrinkRequest) SetTargetLanguage(v string) *VideoTranslationShrinkRequest {
	s.TargetLanguage = &v
	return s
}

func (s *VideoTranslationShrinkRequest) SetVideoUrl(v string) *VideoTranslationShrinkRequest {
	s.VideoUrl = &v
	return s
}

func (s *VideoTranslationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
