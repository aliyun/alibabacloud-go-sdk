// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iGenerateImageWithTextAndImageAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAspectRatioMode(v string) *GenerateImageWithTextAndImageAdvanceRequest
	GetAspectRatioMode() *string
	SetNumber(v int32) *GenerateImageWithTextAndImageAdvanceRequest
	GetNumber() *int32
	SetRefImageUrlObject(v io.Reader) *GenerateImageWithTextAndImageAdvanceRequest
	GetRefImageUrlObject() io.Reader
	SetResolution(v string) *GenerateImageWithTextAndImageAdvanceRequest
	GetResolution() *string
	SetSimilarity(v float32) *GenerateImageWithTextAndImageAdvanceRequest
	GetSimilarity() *float32
	SetText(v string) *GenerateImageWithTextAndImageAdvanceRequest
	GetText() *string
}

type GenerateImageWithTextAndImageAdvanceRequest struct {
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
	RefImageUrlObject io.Reader `json:"RefImageUrl,omitempty" xml:"RefImageUrl,omitempty"`
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

func (s GenerateImageWithTextAndImageAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateImageWithTextAndImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *GenerateImageWithTextAndImageAdvanceRequest) GetAspectRatioMode() *string {
	return s.AspectRatioMode
}

func (s *GenerateImageWithTextAndImageAdvanceRequest) GetNumber() *int32 {
	return s.Number
}

func (s *GenerateImageWithTextAndImageAdvanceRequest) GetRefImageUrlObject() io.Reader {
	return s.RefImageUrlObject
}

func (s *GenerateImageWithTextAndImageAdvanceRequest) GetResolution() *string {
	return s.Resolution
}

func (s *GenerateImageWithTextAndImageAdvanceRequest) GetSimilarity() *float32 {
	return s.Similarity
}

func (s *GenerateImageWithTextAndImageAdvanceRequest) GetText() *string {
	return s.Text
}

func (s *GenerateImageWithTextAndImageAdvanceRequest) SetAspectRatioMode(v string) *GenerateImageWithTextAndImageAdvanceRequest {
	s.AspectRatioMode = &v
	return s
}

func (s *GenerateImageWithTextAndImageAdvanceRequest) SetNumber(v int32) *GenerateImageWithTextAndImageAdvanceRequest {
	s.Number = &v
	return s
}

func (s *GenerateImageWithTextAndImageAdvanceRequest) SetRefImageUrlObject(v io.Reader) *GenerateImageWithTextAndImageAdvanceRequest {
	s.RefImageUrlObject = v
	return s
}

func (s *GenerateImageWithTextAndImageAdvanceRequest) SetResolution(v string) *GenerateImageWithTextAndImageAdvanceRequest {
	s.Resolution = &v
	return s
}

func (s *GenerateImageWithTextAndImageAdvanceRequest) SetSimilarity(v float32) *GenerateImageWithTextAndImageAdvanceRequest {
	s.Similarity = &v
	return s
}

func (s *GenerateImageWithTextAndImageAdvanceRequest) SetText(v string) *GenerateImageWithTextAndImageAdvanceRequest {
	s.Text = &v
	return s
}

func (s *GenerateImageWithTextAndImageAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
