// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoInsightsCaptionConfig interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v bool) *VideoInsightsCaptionConfig
	GetEnable() *bool
	SetPersonReference(v *PersonReferenceConfig) *VideoInsightsCaptionConfig
	GetPersonReference() *PersonReferenceConfig
	SetPrompt(v string) *VideoInsightsCaptionConfig
	GetPrompt() *string
}

type VideoInsightsCaptionConfig struct {
	Enable          *bool                  `json:"Enable,omitempty" xml:"Enable,omitempty"`
	PersonReference *PersonReferenceConfig `json:"PersonReference,omitempty" xml:"PersonReference,omitempty"`
	Prompt          *string                `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
}

func (s VideoInsightsCaptionConfig) String() string {
	return dara.Prettify(s)
}

func (s VideoInsightsCaptionConfig) GoString() string {
	return s.String()
}

func (s *VideoInsightsCaptionConfig) GetEnable() *bool {
	return s.Enable
}

func (s *VideoInsightsCaptionConfig) GetPersonReference() *PersonReferenceConfig {
	return s.PersonReference
}

func (s *VideoInsightsCaptionConfig) GetPrompt() *string {
	return s.Prompt
}

func (s *VideoInsightsCaptionConfig) SetEnable(v bool) *VideoInsightsCaptionConfig {
	s.Enable = &v
	return s
}

func (s *VideoInsightsCaptionConfig) SetPersonReference(v *PersonReferenceConfig) *VideoInsightsCaptionConfig {
	s.PersonReference = v
	return s
}

func (s *VideoInsightsCaptionConfig) SetPrompt(v string) *VideoInsightsCaptionConfig {
	s.Prompt = &v
	return s
}

func (s *VideoInsightsCaptionConfig) Validate() error {
	if s.PersonReference != nil {
		if err := s.PersonReference.Validate(); err != nil {
			return err
		}
	}
	return nil
}
