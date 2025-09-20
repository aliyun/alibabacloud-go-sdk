// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLicensePlate interface {
	dara.Model
	String() string
	GoString() string
	SetBoundary(v *Boundary) *LicensePlate
	GetBoundary() *Boundary
	SetConfidence(v float64) *LicensePlate
	GetConfidence() *float64
	SetContent(v string) *LicensePlate
	GetContent() *string
}

type LicensePlate struct {
	Boundary   *Boundary `json:"Boundary,omitempty" xml:"Boundary,omitempty"`
	Confidence *float64  `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// example:
	//
	// 川A0123
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s LicensePlate) String() string {
	return dara.Prettify(s)
}

func (s LicensePlate) GoString() string {
	return s.String()
}

func (s *LicensePlate) GetBoundary() *Boundary {
	return s.Boundary
}

func (s *LicensePlate) GetConfidence() *float64 {
	return s.Confidence
}

func (s *LicensePlate) GetContent() *string {
	return s.Content
}

func (s *LicensePlate) SetBoundary(v *Boundary) *LicensePlate {
	s.Boundary = v
	return s
}

func (s *LicensePlate) SetConfidence(v float64) *LicensePlate {
	s.Confidence = &v
	return s
}

func (s *LicensePlate) SetContent(v string) *LicensePlate {
	s.Content = &v
	return s
}

func (s *LicensePlate) Validate() error {
	return dara.Validate(s)
}
