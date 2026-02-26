// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProductProperty interface {
	dara.Model
	String() string
	GoString() string
	SetText(v string) *ProductProperty
	GetText() *string
	SetValues(v []*string) *ProductProperty
	GetValues() []*string
}

type ProductProperty struct {
	Text   *string   `json:"text,omitempty" xml:"text,omitempty"`
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s ProductProperty) String() string {
	return dara.Prettify(s)
}

func (s ProductProperty) GoString() string {
	return s.String()
}

func (s *ProductProperty) GetText() *string {
	return s.Text
}

func (s *ProductProperty) GetValues() []*string {
	return s.Values
}

func (s *ProductProperty) SetText(v string) *ProductProperty {
	s.Text = &v
	return s
}

func (s *ProductProperty) SetValues(v []*string) *ProductProperty {
	s.Values = v
	return s
}

func (s *ProductProperty) Validate() error {
	return dara.Validate(s)
}
