// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCodes interface {
	dara.Model
	String() string
	GoString() string
	SetBoundary(v *Boundary) *Codes
	GetBoundary() *Boundary
	SetConfidence(v float32) *Codes
	GetConfidence() *float32
	SetContent(v string) *Codes
	GetContent() *string
	SetType(v string) *Codes
	GetType() *string
}

type Codes struct {
	Boundary   *Boundary `json:"Boundary,omitempty" xml:"Boundary,omitempty"`
	Confidence *float32  `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	Content    *string   `json:"Content,omitempty" xml:"Content,omitempty"`
	Type       *string   `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s Codes) String() string {
	return dara.Prettify(s)
}

func (s Codes) GoString() string {
	return s.String()
}

func (s *Codes) GetBoundary() *Boundary {
	return s.Boundary
}

func (s *Codes) GetConfidence() *float32 {
	return s.Confidence
}

func (s *Codes) GetContent() *string {
	return s.Content
}

func (s *Codes) GetType() *string {
	return s.Type
}

func (s *Codes) SetBoundary(v *Boundary) *Codes {
	s.Boundary = v
	return s
}

func (s *Codes) SetConfidence(v float32) *Codes {
	s.Confidence = &v
	return s
}

func (s *Codes) SetContent(v string) *Codes {
	s.Content = &v
	return s
}

func (s *Codes) SetType(v string) *Codes {
	s.Type = &v
	return s
}

func (s *Codes) Validate() error {
	if s.Boundary != nil {
		if err := s.Boundary.Validate(); err != nil {
			return err
		}
	}
	return nil
}
