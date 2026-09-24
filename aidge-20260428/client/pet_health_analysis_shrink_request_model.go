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
	// The list of HTTPS URLs of images to analyze. At least one accessible image must be provided.
	//
	// This parameter is required.
	ImageUrlShrink *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The system prompt used to specify the response role or requirements. The value must comply with JSON string escaping rules.
	//
	// example:
	//
	// You are a professional veterinarian
	SystemPrompt *string `json:"SystemPrompt,omitempty" xml:"SystemPrompt,omitempty"`
	// The custom analysis requirement. If not specified or set to an empty string, excrement analysis is performed by default. The value must comply with JSON string escaping rules.
	//
	// example:
	//
	// Please analyze the health condition of this pet
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
