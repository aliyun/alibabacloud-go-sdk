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
	SetSourceLanguage(v string) *VideoTranslationRequest
	GetSourceLanguage() *string
	SetTargetLanguage(v string) *VideoTranslationRequest
	GetTargetLanguage() *string
	SetVideoUrl(v string) *VideoTranslationRequest
	GetVideoUrl() *string
}

type VideoTranslationRequest struct {
	// The array of translation capabilities. Valid values: ["visual"].
	//
	// This parameter is required.
	Capabilities []*string `json:"Capabilities,omitempty" xml:"Capabilities,omitempty" type:"Repeated"`
	// The source language. This parameter is optional. Default value: auto (automatic detection).
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// The target language. This parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// ru
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	// The video URL (MP4/MOV, ≤ 200 MB).
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/video.mp4
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

func (s *VideoTranslationRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
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

func (s *VideoTranslationRequest) SetSourceLanguage(v string) *VideoTranslationRequest {
	s.SourceLanguage = &v
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
