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
