// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRequestTag interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *RequestTag
	GetKey() *string
	SetValue(v string) *RequestTag
	GetValue() *string
}

type RequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s RequestTag) String() string {
	return dara.Prettify(s)
}

func (s RequestTag) GoString() string {
	return s.String()
}

func (s *RequestTag) GetKey() *string {
	return s.Key
}

func (s *RequestTag) GetValue() *string {
	return s.Value
}

func (s *RequestTag) SetKey(v string) *RequestTag {
	s.Key = &v
	return s
}

func (s *RequestTag) SetValue(v string) *RequestTag {
	s.Value = &v
	return s
}

func (s *RequestTag) Validate() error {
	return dara.Validate(s)
}
