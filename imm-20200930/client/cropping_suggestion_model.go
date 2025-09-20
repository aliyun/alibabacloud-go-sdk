// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCroppingSuggestion interface {
	dara.Model
	String() string
	GoString() string
	SetAspectRatio(v string) *CroppingSuggestion
	GetAspectRatio() *string
	SetBoundary(v *Boundary) *CroppingSuggestion
	GetBoundary() *Boundary
	SetConfidence(v float32) *CroppingSuggestion
	GetConfidence() *float32
}

type CroppingSuggestion struct {
	AspectRatio *string   `json:"AspectRatio,omitempty" xml:"AspectRatio,omitempty"`
	Boundary    *Boundary `json:"Boundary,omitempty" xml:"Boundary,omitempty"`
	Confidence  *float32  `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
}

func (s CroppingSuggestion) String() string {
	return dara.Prettify(s)
}

func (s CroppingSuggestion) GoString() string {
	return s.String()
}

func (s *CroppingSuggestion) GetAspectRatio() *string {
	return s.AspectRatio
}

func (s *CroppingSuggestion) GetBoundary() *Boundary {
	return s.Boundary
}

func (s *CroppingSuggestion) GetConfidence() *float32 {
	return s.Confidence
}

func (s *CroppingSuggestion) SetAspectRatio(v string) *CroppingSuggestion {
	s.AspectRatio = &v
	return s
}

func (s *CroppingSuggestion) SetBoundary(v *Boundary) *CroppingSuggestion {
	s.Boundary = v
	return s
}

func (s *CroppingSuggestion) SetConfidence(v float32) *CroppingSuggestion {
	s.Confidence = &v
	return s
}

func (s *CroppingSuggestion) Validate() error {
	return dara.Validate(s)
}
