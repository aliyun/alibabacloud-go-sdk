// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSummary interface {
	dara.Model
	String() string
	GoString() string
	SetImage(v *Illustration) *Summary
	GetImage() *Illustration
	SetText(v string) *Summary
	GetText() *string
}

type Summary struct {
	Image *Illustration `json:"Image,omitempty" xml:"Image,omitempty"`
	Text  *string       `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s Summary) String() string {
	return dara.Prettify(s)
}

func (s Summary) GoString() string {
	return s.String()
}

func (s *Summary) GetImage() *Illustration {
	return s.Image
}

func (s *Summary) GetText() *string {
	return s.Text
}

func (s *Summary) SetImage(v *Illustration) *Summary {
	s.Image = v
	return s
}

func (s *Summary) SetText(v string) *Summary {
	s.Text = &v
	return s
}

func (s *Summary) Validate() error {
	if s.Image != nil {
		if err := s.Image.Validate(); err != nil {
			return err
		}
	}
	return nil
}
