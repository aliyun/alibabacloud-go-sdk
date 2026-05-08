// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTextTheme interface {
	dara.Model
	String() string
	GoString() string
	SetDesc(v string) *TextTheme
	GetDesc() *string
	SetName(v string) *TextTheme
	GetName() *string
}

type TextTheme struct {
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s TextTheme) String() string {
	return dara.Prettify(s)
}

func (s TextTheme) GoString() string {
	return s.String()
}

func (s *TextTheme) GetDesc() *string {
	return s.Desc
}

func (s *TextTheme) GetName() *string {
	return s.Name
}

func (s *TextTheme) SetDesc(v string) *TextTheme {
	s.Desc = &v
	return s
}

func (s *TextTheme) SetName(v string) *TextTheme {
	s.Name = &v
	return s
}

func (s *TextTheme) Validate() error {
	return dara.Validate(s)
}
