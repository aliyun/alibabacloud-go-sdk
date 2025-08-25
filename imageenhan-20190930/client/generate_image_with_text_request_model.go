// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateImageWithTextRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNumber(v int32) *GenerateImageWithTextRequest
	GetNumber() *int32
	SetResolution(v string) *GenerateImageWithTextRequest
	GetResolution() *string
	SetText(v string) *GenerateImageWithTextRequest
	GetText() *string
}

type GenerateImageWithTextRequest struct {
	// example:
	//
	// 1
	Number *int32 `json:"Number,omitempty" xml:"Number,omitempty"`
	// example:
	//
	// 1024*1024
	Resolution *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	// This parameter is required.
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GenerateImageWithTextRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateImageWithTextRequest) GoString() string {
	return s.String()
}

func (s *GenerateImageWithTextRequest) GetNumber() *int32 {
	return s.Number
}

func (s *GenerateImageWithTextRequest) GetResolution() *string {
	return s.Resolution
}

func (s *GenerateImageWithTextRequest) GetText() *string {
	return s.Text
}

func (s *GenerateImageWithTextRequest) SetNumber(v int32) *GenerateImageWithTextRequest {
	s.Number = &v
	return s
}

func (s *GenerateImageWithTextRequest) SetResolution(v string) *GenerateImageWithTextRequest {
	s.Resolution = &v
	return s
}

func (s *GenerateImageWithTextRequest) SetText(v string) *GenerateImageWithTextRequest {
	s.Text = &v
	return s
}

func (s *GenerateImageWithTextRequest) Validate() error {
	return dara.Validate(s)
}
