// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunParam interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *RunParam
	GetKey() *string
	SetValue(v string) *RunParam
	GetValue() *string
}

type RunParam struct {
	// This parameter is required.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s RunParam) String() string {
	return dara.Prettify(s)
}

func (s RunParam) GoString() string {
	return s.String()
}

func (s *RunParam) GetKey() *string {
	return s.Key
}

func (s *RunParam) GetValue() *string {
	return s.Value
}

func (s *RunParam) SetKey(v string) *RunParam {
	s.Key = &v
	return s
}

func (s *RunParam) SetValue(v string) *RunParam {
	s.Value = &v
	return s
}

func (s *RunParam) Validate() error {
	return dara.Validate(s)
}
