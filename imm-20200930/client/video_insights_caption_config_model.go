// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoInsightsCaptionConfig interface {
	dara.Model
	String() string
	GoString() string
	SetPersonReference(v *PersonReferenceConfig) *VideoInsightsCaptionConfig
	GetPersonReference() *PersonReferenceConfig
}

type VideoInsightsCaptionConfig struct {
	PersonReference *PersonReferenceConfig `json:"PersonReference,omitempty" xml:"PersonReference,omitempty"`
}

func (s VideoInsightsCaptionConfig) String() string {
	return dara.Prettify(s)
}

func (s VideoInsightsCaptionConfig) GoString() string {
	return s.String()
}

func (s *VideoInsightsCaptionConfig) GetPersonReference() *PersonReferenceConfig {
	return s.PersonReference
}

func (s *VideoInsightsCaptionConfig) SetPersonReference(v *PersonReferenceConfig) *VideoInsightsCaptionConfig {
	s.PersonReference = v
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
