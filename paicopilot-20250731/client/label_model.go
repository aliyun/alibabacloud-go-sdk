// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLabel interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *Label
	GetKey() *string
	SetValue(v string) *Label
	GetValue() *string
}

type Label struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s Label) String() string {
	return dara.Prettify(s)
}

func (s Label) GoString() string {
	return s.String()
}

func (s *Label) GetKey() *string {
	return s.Key
}

func (s *Label) GetValue() *string {
	return s.Value
}

func (s *Label) SetKey(v string) *Label {
	s.Key = &v
	return s
}

func (s *Label) SetValue(v string) *Label {
	s.Value = &v
	return s
}

func (s *Label) Validate() error {
	return dara.Validate(s)
}
