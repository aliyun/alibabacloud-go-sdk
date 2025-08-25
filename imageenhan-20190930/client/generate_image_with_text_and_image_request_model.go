// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateImageWithTextAndImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAspectRatioMode(v string) *GenerateImageWithTextAndImageRequest
	GetAspectRatioMode() *string
	SetNumber(v int32) *GenerateImageWithTextAndImageRequest
	GetNumber() *int32
	SetRefImageUrl(v string) *GenerateImageWithTextAndImageRequest
	GetRefImageUrl() *string
	SetResolution(v string) *GenerateImageWithTextAndImageRequest
	GetResolution() *string
	SetSimilarity(v float32) *GenerateImageWithTextAndImageRequest
	GetSimilarity() *float32
	SetText(v string) *GenerateImageWithTextAndImageRequest
	GetText() *string
}

type GenerateImageWithTextAndImageRequest struct {
	// example:
	//
	// center_crop
	AspectRatioMode *string `json:"AspectRatioMode,omitempty" xml:"AspectRatioMode,omitempty"`
	// example:
	//
	// 1
	Number *int32 `json:"Number,omitempty" xml:"Number,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://viapi-test.oss-cn-shanghai.aliyuncs.com/yxrtest/%E6%96%87%E7%94%9F%E5%9B%BE/%E7%9D%A1%E8%8E%B2.jpg"
	RefImageUrl *string `json:"RefImageUrl,omitempty" xml:"RefImageUrl,omitempty"`
	// example:
	//
	// 1024*1024
	Resolution *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	// example:
	//
	// 0.2
	Similarity *float32 `json:"Similarity,omitempty" xml:"Similarity,omitempty"`
	// This parameter is required.
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GenerateImageWithTextAndImageRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateImageWithTextAndImageRequest) GoString() string {
	return s.String()
}

func (s *GenerateImageWithTextAndImageRequest) GetAspectRatioMode() *string {
	return s.AspectRatioMode
}

func (s *GenerateImageWithTextAndImageRequest) GetNumber() *int32 {
	return s.Number
}

func (s *GenerateImageWithTextAndImageRequest) GetRefImageUrl() *string {
	return s.RefImageUrl
}

func (s *GenerateImageWithTextAndImageRequest) GetResolution() *string {
	return s.Resolution
}

func (s *GenerateImageWithTextAndImageRequest) GetSimilarity() *float32 {
	return s.Similarity
}

func (s *GenerateImageWithTextAndImageRequest) GetText() *string {
	return s.Text
}

func (s *GenerateImageWithTextAndImageRequest) SetAspectRatioMode(v string) *GenerateImageWithTextAndImageRequest {
	s.AspectRatioMode = &v
	return s
}

func (s *GenerateImageWithTextAndImageRequest) SetNumber(v int32) *GenerateImageWithTextAndImageRequest {
	s.Number = &v
	return s
}

func (s *GenerateImageWithTextAndImageRequest) SetRefImageUrl(v string) *GenerateImageWithTextAndImageRequest {
	s.RefImageUrl = &v
	return s
}

func (s *GenerateImageWithTextAndImageRequest) SetResolution(v string) *GenerateImageWithTextAndImageRequest {
	s.Resolution = &v
	return s
}

func (s *GenerateImageWithTextAndImageRequest) SetSimilarity(v float32) *GenerateImageWithTextAndImageRequest {
	s.Similarity = &v
	return s
}

func (s *GenerateImageWithTextAndImageRequest) SetText(v string) *GenerateImageWithTextAndImageRequest {
	s.Text = &v
	return s
}

func (s *GenerateImageWithTextAndImageRequest) Validate() error {
	return dara.Validate(s)
}
