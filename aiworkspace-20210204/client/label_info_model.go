// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLabelInfo interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *LabelInfo
	GetKey() *string
	SetValue(v string) *LabelInfo
	GetValue() *string
}

type LabelInfo struct {
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s LabelInfo) String() string {
	return dara.Prettify(s)
}

func (s LabelInfo) GoString() string {
	return s.String()
}

func (s *LabelInfo) GetKey() *string {
	return s.Key
}

func (s *LabelInfo) GetValue() *string {
	return s.Value
}

func (s *LabelInfo) SetKey(v string) *LabelInfo {
	s.Key = &v
	return s
}

func (s *LabelInfo) SetValue(v string) *LabelInfo {
	s.Value = &v
	return s
}

func (s *LabelInfo) Validate() error {
	return dara.Validate(s)
}
