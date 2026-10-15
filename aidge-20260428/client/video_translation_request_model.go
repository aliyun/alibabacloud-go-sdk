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
	SetSourceLanguage(v string) *VideoTranslationRequest
	GetSourceLanguage() *string
	SetTargetLanguage(v string) *VideoTranslationRequest
	GetTargetLanguage() *string
	SetVideoUrl(v string) *VideoTranslationRequest
	GetVideoUrl() *string
}

type VideoTranslationRequest struct {
	// The list of translation features. Valid values: `speech` and `visual`. You can select both. At least one feature is required. Duplicate values are not accepted.
	//
	// This parameter is required.
	Capabilities []*string `json:"Capabilities,omitempty" xml:"Capabilities,omitempty" type:"Repeated"`
	// Default value: false. If `visual` is included and this parameter is explicitly set to true, the quality check for on-screen text processing is enabled. If only `speech` is selected, this parameter is ignored. No quality check is performed, no quality check fields are returned, and no errors are reported due to this parameter. Enabling the quality check may increase the waiting time. Do not enable it unless necessary.
	//
	// example:
	//
	// true
	EnableQualityCheck *bool `json:"EnableQualityCheck,omitempty" xml:"EnableQualityCheck,omitempty"`
	// The source language. For supported languages, see "Supported languages".
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

func (s *VideoTranslationRequest) SetEnableQualityCheck(v bool) *VideoTranslationRequest {
	s.EnableQualityCheck = &v
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
