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
	// Specifies whether to enable the quality check for on-screen text processing. Default value: false. The quality check is enabled when visual is included and this parameter is explicitly set to true. If only speech is selected, this parameter is ignored, no quality check is performed, no quality check fields are returned, and no errors are reported due to this parameter. Enabling the quality check may increase the waiting time. Do not enable it unless necessary.
	//
	// example:
	//
	// true
	EnableQualityCheck *bool `json:"EnableQualityCheck,omitempty" xml:"EnableQualityCheck,omitempty"`
	// Specifies whether to only erase the source language text in the video frames without translating it. This parameter is valid only when visual is included. Valid values: true (erases only) and false (translates the on-screen text). Default value: false.
	//
	// example:
	//
	// true
	EraseOnly *bool `json:"EraseOnly,omitempty" xml:"EraseOnly,omitempty"`
	// Specifies whether to erase brand logos. This parameter is valid only when visual is included. Valid values: true (erases brand logos) and false (retains logos). Default value: false. This parameter is independent of eraseOnly.
	//
	// example:
	//
	// true
	RemoveLogo *bool `json:"RemoveLogo,omitempty" xml:"RemoveLogo,omitempty"`
	// The source language. For supported languages, see Language support.
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// The subtitle processing mode. This parameter is valid only when speech is included. Valid values: erase (erases only the original subtitles) and replace (erases the original subtitles and adds new subtitles). If this parameter is omitted, the original subtitles are not processed and no new subtitles are added.
	//
	// example:
	//
	// erase
	SubtitleMode *string `json:"SubtitleMode,omitempty" xml:"SubtitleMode,omitempty"`
	// The target language. This parameter is required.
	//
	// example:
	//
	// ru
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	// The URL of the video. Supported formats: MP4 and MOV. Maximum size: 200 MB.
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
