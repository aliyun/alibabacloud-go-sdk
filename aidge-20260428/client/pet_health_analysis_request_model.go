// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPetHealthAnalysisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrl(v []*string) *PetHealthAnalysisRequest
	GetImageUrl() []*string
	SetSystemPrompt(v string) *PetHealthAnalysisRequest
	GetSystemPrompt() *string
	SetUserPrompt(v string) *PetHealthAnalysisRequest
	GetUserPrompt() *string
}

type PetHealthAnalysisRequest struct {
	// The list of HTTPS URLs of images to analyze. At least one accessible image must be provided.
	//
	// This parameter is required.
	ImageUrl []*string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty" type:"Repeated"`
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

func (s PetHealthAnalysisRequest) String() string {
	return dara.Prettify(s)
}

func (s PetHealthAnalysisRequest) GoString() string {
	return s.String()
}

func (s *PetHealthAnalysisRequest) GetImageUrl() []*string {
	return s.ImageUrl
}

func (s *PetHealthAnalysisRequest) GetSystemPrompt() *string {
	return s.SystemPrompt
}

func (s *PetHealthAnalysisRequest) GetUserPrompt() *string {
	return s.UserPrompt
}

func (s *PetHealthAnalysisRequest) SetImageUrl(v []*string) *PetHealthAnalysisRequest {
	s.ImageUrl = v
	return s
}

func (s *PetHealthAnalysisRequest) SetSystemPrompt(v string) *PetHealthAnalysisRequest {
	s.SystemPrompt = &v
	return s
}

func (s *PetHealthAnalysisRequest) SetUserPrompt(v string) *PetHealthAnalysisRequest {
	s.UserPrompt = &v
	return s
}

func (s *PetHealthAnalysisRequest) Validate() error {
	return dara.Validate(s)
}
