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
	SetSourceLanguage(v string) *VideoTranslationShrinkRequest
	GetSourceLanguage() *string
	SetTargetLanguage(v string) *VideoTranslationShrinkRequest
	GetTargetLanguage() *string
	SetVideoUrl(v string) *VideoTranslationShrinkRequest
	GetVideoUrl() *string
}

type VideoTranslationShrinkRequest struct {
	// This parameter is required.
	CapabilitiesShrink *string `json:"Capabilities,omitempty" xml:"Capabilities,omitempty"`
	// example:
	//
	// zh
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ru
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/video.mp4
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

func (s *VideoTranslationShrinkRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
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

func (s *VideoTranslationShrinkRequest) SetSourceLanguage(v string) *VideoTranslationShrinkRequest {
	s.SourceLanguage = &v
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
