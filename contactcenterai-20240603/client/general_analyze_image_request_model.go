// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGeneralAnalyzeImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomPrompt(v string) *GeneralAnalyzeImageRequest
	GetCustomPrompt() *string
	SetImageUrls(v []*string) *GeneralAnalyzeImageRequest
	GetImageUrls() []*string
	SetStream(v bool) *GeneralAnalyzeImageRequest
	GetStream() *bool
	SetTemplateIds(v []*int64) *GeneralAnalyzeImageRequest
	GetTemplateIds() []*int64
}

type GeneralAnalyzeImageRequest struct {
	// example:
	//
	// Analyze the content in the image
	CustomPrompt *string `json:"customPrompt,omitempty" xml:"customPrompt,omitempty"`
	// This parameter is required.
	ImageUrls []*string `json:"imageUrls,omitempty" xml:"imageUrls,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Stream      *bool    `json:"stream,omitempty" xml:"stream,omitempty"`
	TemplateIds []*int64 `json:"templateIds,omitempty" xml:"templateIds,omitempty" type:"Repeated"`
}

func (s GeneralAnalyzeImageRequest) String() string {
	return dara.Prettify(s)
}

func (s GeneralAnalyzeImageRequest) GoString() string {
	return s.String()
}

func (s *GeneralAnalyzeImageRequest) GetCustomPrompt() *string {
	return s.CustomPrompt
}

func (s *GeneralAnalyzeImageRequest) GetImageUrls() []*string {
	return s.ImageUrls
}

func (s *GeneralAnalyzeImageRequest) GetStream() *bool {
	return s.Stream
}

func (s *GeneralAnalyzeImageRequest) GetTemplateIds() []*int64 {
	return s.TemplateIds
}

func (s *GeneralAnalyzeImageRequest) SetCustomPrompt(v string) *GeneralAnalyzeImageRequest {
	s.CustomPrompt = &v
	return s
}

func (s *GeneralAnalyzeImageRequest) SetImageUrls(v []*string) *GeneralAnalyzeImageRequest {
	s.ImageUrls = v
	return s
}

func (s *GeneralAnalyzeImageRequest) SetStream(v bool) *GeneralAnalyzeImageRequest {
	s.Stream = &v
	return s
}

func (s *GeneralAnalyzeImageRequest) SetTemplateIds(v []*int64) *GeneralAnalyzeImageRequest {
	s.TemplateIds = v
	return s
}

func (s *GeneralAnalyzeImageRequest) Validate() error {
	return dara.Validate(s)
}
