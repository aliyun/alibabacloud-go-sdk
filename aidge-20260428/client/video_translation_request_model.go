// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoTranslationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCapabilities(v []*string) *VideoTranslationRequest
	GetCapabilities() []*string
	SetEnableQualityCheck(v bool) *VideoTranslationRequest
	GetEnableQualityCheck() *bool
	SetEraseOnly(v bool) *VideoTranslationRequest
	GetEraseOnly() *bool
	SetRemoveLogo(v bool) *VideoTranslationRequest
	GetRemoveLogo() *bool
	SetSourceLanguage(v string) *VideoTranslationRequest
	GetSourceLanguage() *string
	SetSubtitleMode(v string) *VideoTranslationRequest
	GetSubtitleMode() *string
	SetTargetLanguage(v string) *VideoTranslationRequest
	GetTargetLanguage() *string
	SetVideoUrl(v string) *VideoTranslationRequest
	GetVideoUrl() *string
}

type VideoTranslationRequest struct {
	// The list of translation features. Valid values: speech and visual. You can select both. At least one feature is required. Duplicate values are not accepted.
	//
	// This parameter is required.
	Capabilities []*string `json:"Capabilities,omitempty" xml:"Capabilities,omitempty" type:"Repeated"`
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

func (s VideoTranslationRequest) String() string {
	return dara.Prettify(s)
}

func (s VideoTranslationRequest) GoString() string {
	return s.String()
}

func (s *VideoTranslationRequest) GetCapabilities() []*string {
	return s.Capabilities
}

func (s *VideoTranslationRequest) GetEnableQualityCheck() *bool {
	return s.EnableQualityCheck
}

func (s *VideoTranslationRequest) GetEraseOnly() *bool {
	return s.EraseOnly
}

func (s *VideoTranslationRequest) GetRemoveLogo() *bool {
	return s.RemoveLogo
}

func (s *VideoTranslationRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *VideoTranslationRequest) GetSubtitleMode() *string {
	return s.SubtitleMode
}

func (s *VideoTranslationRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *VideoTranslationRequest) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *VideoTranslationRequest) SetCapabilities(v []*string) *VideoTranslationRequest {
	s.Capabilities = v
	return s
}

func (s *VideoTranslationRequest) SetEnableQualityCheck(v bool) *VideoTranslationRequest {
	s.EnableQualityCheck = &v
	return s
}

func (s *VideoTranslationRequest) SetEraseOnly(v bool) *VideoTranslationRequest {
	s.EraseOnly = &v
	return s
}

func (s *VideoTranslationRequest) SetRemoveLogo(v bool) *VideoTranslationRequest {
	s.RemoveLogo = &v
	return s
}

func (s *VideoTranslationRequest) SetSourceLanguage(v string) *VideoTranslationRequest {
	s.SourceLanguage = &v
	return s
}

func (s *VideoTranslationRequest) SetSubtitleMode(v string) *VideoTranslationRequest {
	s.SubtitleMode = &v
	return s
}

func (s *VideoTranslationRequest) SetTargetLanguage(v string) *VideoTranslationRequest {
	s.TargetLanguage = &v
	return s
}

func (s *VideoTranslationRequest) SetVideoUrl(v string) *VideoTranslationRequest {
	s.VideoUrl = &v
	return s
}

func (s *VideoTranslationRequest) Validate() error {
	return dara.Validate(s)
}
