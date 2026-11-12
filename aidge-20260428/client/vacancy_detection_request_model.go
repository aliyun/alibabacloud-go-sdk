// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVacancyDetectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrl(v string) *VacancyDetectionRequest
	GetImageUrl() *string
	SetSystemPrompt(v string) *VacancyDetectionRequest
	GetSystemPrompt() *string
	SetUserPrompt(v string) *VacancyDetectionRequest
	GetUserPrompt() *string
}

type VacancyDetectionRequest struct {
	// The HTTPS URL of the single image to be detected. Arrays or other image encodings are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/shelf.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The optional system prompt.
	//
	// example:
	//
	// You are a helpful assistant.
	SystemPrompt *string `json:"SystemPrompt,omitempty" xml:"SystemPrompt,omitempty"`
	// The optional user prompt. If this parameter is empty, the complete vacancy detection prompt is used. Plain text and JSON containing only text in messages or content are supported.
	//
	// example:
	//
	// Detect vacancies and underfilled areas on the shelf
	UserPrompt *string `json:"UserPrompt,omitempty" xml:"UserPrompt,omitempty"`
}

func (s VacancyDetectionRequest) String() string {
	return dara.Prettify(s)
}

func (s VacancyDetectionRequest) GoString() string {
	return s.String()
}

func (s *VacancyDetectionRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *VacancyDetectionRequest) GetSystemPrompt() *string {
	return s.SystemPrompt
}

func (s *VacancyDetectionRequest) GetUserPrompt() *string {
	return s.UserPrompt
}

func (s *VacancyDetectionRequest) SetImageUrl(v string) *VacancyDetectionRequest {
	s.ImageUrl = &v
	return s
}

func (s *VacancyDetectionRequest) SetSystemPrompt(v string) *VacancyDetectionRequest {
	s.SystemPrompt = &v
	return s
}

func (s *VacancyDetectionRequest) SetUserPrompt(v string) *VacancyDetectionRequest {
	s.UserPrompt = &v
	return s
}

func (s *VacancyDetectionRequest) Validate() error {
	return dara.Validate(s)
}
