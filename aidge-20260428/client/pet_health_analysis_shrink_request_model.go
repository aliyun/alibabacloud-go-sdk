// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPetHealthAnalysisShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrlShrink(v string) *PetHealthAnalysisShrinkRequest
	GetImageUrlShrink() *string
	SetSystemPrompt(v string) *PetHealthAnalysisShrinkRequest
	GetSystemPrompt() *string
	SetUserPrompt(v string) *PetHealthAnalysisShrinkRequest
	GetUserPrompt() *string
}

type PetHealthAnalysisShrinkRequest struct {
	// This parameter is required.
	ImageUrlShrink *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// example:
	//
	// 你是一名专业的宠物医生
	SystemPrompt *string `json:"SystemPrompt,omitempty" xml:"SystemPrompt,omitempty"`
	// example:
	//
	// 请分析这只宠物的健康状况
	UserPrompt *string `json:"UserPrompt,omitempty" xml:"UserPrompt,omitempty"`
}

func (s PetHealthAnalysisShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PetHealthAnalysisShrinkRequest) GoString() string {
	return s.String()
}

func (s *PetHealthAnalysisShrinkRequest) GetImageUrlShrink() *string {
	return s.ImageUrlShrink
}

func (s *PetHealthAnalysisShrinkRequest) GetSystemPrompt() *string {
	return s.SystemPrompt
}

func (s *PetHealthAnalysisShrinkRequest) GetUserPrompt() *string {
	return s.UserPrompt
}

func (s *PetHealthAnalysisShrinkRequest) SetImageUrlShrink(v string) *PetHealthAnalysisShrinkRequest {
	s.ImageUrlShrink = &v
	return s
}

func (s *PetHealthAnalysisShrinkRequest) SetSystemPrompt(v string) *PetHealthAnalysisShrinkRequest {
	s.SystemPrompt = &v
	return s
}

func (s *PetHealthAnalysisShrinkRequest) SetUserPrompt(v string) *PetHealthAnalysisShrinkRequest {
	s.UserPrompt = &v
	return s
}

func (s *PetHealthAnalysisShrinkRequest) Validate() error {
	return dara.Validate(s)
}
