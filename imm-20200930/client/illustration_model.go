// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIllustration interface {
	dara.Model
	String() string
	GoString() string
	SetImageIndex(v int32) *Illustration
	GetImageIndex() *int32
	SetImagePath(v string) *Illustration
	GetImagePath() *string
	SetNormalizedBox(v []*float32) *Illustration
	GetNormalizedBox() []*float32
	SetPageNumber(v int32) *Illustration
	GetPageNumber() *int32
	SetText(v string) *Illustration
	GetText() *string
	SetType(v string) *Illustration
	GetType() *string
}

type Illustration struct {
	// The zero-based image index in a file that contains multiple images, such as a multi-page TIFF file.
	ImageIndex *int32 `json:"ImageIndex,omitempty" xml:"ImageIndex,omitempty"`
	// The path to the image file containing the illustration.
	ImagePath *string `json:"ImagePath,omitempty" xml:"ImagePath,omitempty"`
	// An array of four floating-point numbers that defines the normalized box for the illustration in [x_min, y_min, x_max, y_max] format. The coordinates are normalized to a range of [0, 1] relative to the page dimensions.
	NormalizedBox []*float32 `json:"NormalizedBox,omitempty" xml:"NormalizedBox,omitempty" type:"Repeated"`
	// The one-based page number where the illustration is located.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The text associated with the illustration.
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The type of the illustration, such as `figure` or `chart`.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s Illustration) String() string {
	return dara.Prettify(s)
}

func (s Illustration) GoString() string {
	return s.String()
}

func (s *Illustration) GetImageIndex() *int32 {
	return s.ImageIndex
}

func (s *Illustration) GetImagePath() *string {
	return s.ImagePath
}

func (s *Illustration) GetNormalizedBox() []*float32 {
	return s.NormalizedBox
}

func (s *Illustration) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *Illustration) GetText() *string {
	return s.Text
}

func (s *Illustration) GetType() *string {
	return s.Type
}

func (s *Illustration) SetImageIndex(v int32) *Illustration {
	s.ImageIndex = &v
	return s
}

func (s *Illustration) SetImagePath(v string) *Illustration {
	s.ImagePath = &v
	return s
}

func (s *Illustration) SetNormalizedBox(v []*float32) *Illustration {
	s.NormalizedBox = v
	return s
}

func (s *Illustration) SetPageNumber(v int32) *Illustration {
	s.PageNumber = &v
	return s
}

func (s *Illustration) SetText(v string) *Illustration {
	s.Text = &v
	return s
}

func (s *Illustration) SetType(v string) *Illustration {
	s.Type = &v
	return s
}

func (s *Illustration) Validate() error {
	return dara.Validate(s)
}
