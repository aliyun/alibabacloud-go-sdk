// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOCRContents interface {
	dara.Model
	String() string
	GoString() string
	SetBoundary(v *Boundary) *OCRContents
	GetBoundary() *Boundary
	SetConfidence(v float32) *OCRContents
	GetConfidence() *float32
	SetContents(v string) *OCRContents
	GetContents() *string
	SetLanguage(v string) *OCRContents
	GetLanguage() *string
}

type OCRContents struct {
	// The boundary information.
	Boundary *Boundary `json:"Boundary,omitempty" xml:"Boundary,omitempty"`
	// The confidence level of the content. Valid values: 0 to 1. The value 0 indicates the lowest confidence level. The value 1 indicates the highest confidence level.
	//
	// example:
	//
	// 0.8254936695098877
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// The content.
	Contents *string `json:"Contents,omitempty" xml:"Contents,omitempty"`
	// The BCP 47 language code. This parameter is not supported in the current version.
	//
	// example:
	//
	// zh-hans
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s OCRContents) String() string {
	return dara.Prettify(s)
}

func (s OCRContents) GoString() string {
	return s.String()
}

func (s *OCRContents) GetBoundary() *Boundary {
	return s.Boundary
}

func (s *OCRContents) GetConfidence() *float32 {
	return s.Confidence
}

func (s *OCRContents) GetContents() *string {
	return s.Contents
}

func (s *OCRContents) GetLanguage() *string {
	return s.Language
}

func (s *OCRContents) SetBoundary(v *Boundary) *OCRContents {
	s.Boundary = v
	return s
}

func (s *OCRContents) SetConfidence(v float32) *OCRContents {
	s.Confidence = &v
	return s
}

func (s *OCRContents) SetContents(v string) *OCRContents {
	s.Contents = &v
	return s
}

func (s *OCRContents) SetLanguage(v string) *OCRContents {
	s.Language = &v
	return s
}

func (s *OCRContents) Validate() error {
	if s.Boundary != nil {
		if err := s.Boundary.Validate(); err != nil {
			return err
		}
	}
	return nil
}
