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
	// The list of translation features. Valid values: speech and visual. You can select both. At least one feature must be specified. Duplicate values are not allowed.
	//
	// This parameter is required.
	CapabilitiesShrink *string `json:"Capabilities,omitempty" xml:"Capabilities,omitempty"`
	// Specifies whether to enable the quality check for on-screen text processing. This parameter takes effect only when visual is included and explicitly set to true. If only speech is selected, this parameter is ignored. No quality check is performed, no quality check fields are returned, and no error is reported for this parameter. Enabling the quality check may increase the waiting time. Do not enable it unless necessary. Default value: false.
	//
	// example:
	//
	// false
	EnableQualityCheck *bool `json:"EnableQualityCheck,omitempty" xml:"EnableQualityCheck,omitempty"`
	// Valid only when visual is included. Specifies whether to only erase the on-screen text without translation. Valid values:
	//
	// - true: Erases the on-screen text without translation.
	//
	// - false: Translates the on-screen text.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EraseOnly *bool `json:"EraseOnly,omitempty" xml:"EraseOnly,omitempty"`
	// Valid only when visual is included. Specifies whether to erase brand logos. Valid values:
	//
	// - true: Erases brand logos.
	//
	// - false: Retains brand logos.
	//
	// Default value: false. This parameter is independent of EraseOnly.
	//
	// example:
	//
	// false
	RemoveLogo *bool `json:"RemoveLogo,omitempty" xml:"RemoveLogo,omitempty"`
	// The original language of the video. For supported languages, see Language support.
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// Valid only when speech is included. Valid values:
	//
	// - erase: Erases only the original subtitles.
	//
	// - replace: Replaces the original subtitles with translated subtitles. This erases the original subtitles and adds subtitles synchronized with the translated speech.
	//
	// - If not specified, the original subtitles are not processed. This means the original subtitles are neither erased nor replaced with new subtitles.
	//
	// example:
	//
	// replace
	SubtitleMode *string `json:"SubtitleMode,omitempty" xml:"SubtitleMode,omitempty"`
	// Required when speech is included. Optional when only visual is selected and EraseOnly is set to true. If provided in this case, it is ignored. Required in all other cases.
	//
	// example:
	//
	// en
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
