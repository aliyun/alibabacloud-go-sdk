// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBody interface {
	dara.Model
	String() string
	GoString() string
	SetBoundary(v *Boundary) *Body
	GetBoundary() *Boundary
	SetConfidence(v float32) *Body
	GetConfidence() *float32
}

type Body struct {
	Boundary   *Boundary `json:"Boundary,omitempty" xml:"Boundary,omitempty"`
	Confidence *float32  `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
}

func (s Body) String() string {
	return dara.Prettify(s)
}

func (s Body) GoString() string {
	return s.String()
}

func (s *Body) GetBoundary() *Boundary {
	return s.Boundary
}

func (s *Body) GetConfidence() *float32 {
	return s.Confidence
}

func (s *Body) SetBoundary(v *Boundary) *Body {
	s.Boundary = v
	return s
}

func (s *Body) SetConfidence(v float32) *Body {
	s.Confidence = &v
	return s
}

func (s *Body) Validate() error {
	if s.Boundary != nil {
		if err := s.Boundary.Validate(); err != nil {
			return err
		}
	}
	return nil
}
