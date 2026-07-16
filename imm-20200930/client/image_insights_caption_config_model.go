// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageInsightsCaptionConfig interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v bool) *ImageInsightsCaptionConfig
	GetEnable() *bool
	SetPrompt(v string) *ImageInsightsCaptionConfig
	GetPrompt() *string
}

type ImageInsightsCaptionConfig struct {
	// Specifies whether to enable this feature.
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The prompt.
	//
	// example:
	//
	// Provide a concise title for this monitoring section, capturing the core subject and key event. Keep the title within 10 characters.
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
}

func (s ImageInsightsCaptionConfig) String() string {
	return dara.Prettify(s)
}

func (s ImageInsightsCaptionConfig) GoString() string {
	return s.String()
}

func (s *ImageInsightsCaptionConfig) GetEnable() *bool {
	return s.Enable
}

func (s *ImageInsightsCaptionConfig) GetPrompt() *string {
	return s.Prompt
}

func (s *ImageInsightsCaptionConfig) SetEnable(v bool) *ImageInsightsCaptionConfig {
	s.Enable = &v
	return s
}

func (s *ImageInsightsCaptionConfig) SetPrompt(v string) *ImageInsightsCaptionConfig {
	s.Prompt = &v
	return s
}

func (s *ImageInsightsCaptionConfig) Validate() error {
	return dara.Validate(s)
}
