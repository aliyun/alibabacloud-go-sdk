// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFilterList interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *FilterList
	GetKey() *string
	SetType(v string) *FilterList
	GetType() *string
	SetValue(v string) *FilterList
	GetValue() *string
}

type FilterList struct {
	// This parameter is required.
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// This parameter is required.
	Type  *string `json:"type,omitempty" xml:"type,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s FilterList) String() string {
	return dara.Prettify(s)
}

func (s FilterList) GoString() string {
	return s.String()
}

func (s *FilterList) GetKey() *string {
	return s.Key
}

func (s *FilterList) GetType() *string {
	return s.Type
}

func (s *FilterList) GetValue() *string {
	return s.Value
}

func (s *FilterList) SetKey(v string) *FilterList {
	s.Key = &v
	return s
}

func (s *FilterList) SetType(v string) *FilterList {
	s.Type = &v
	return s
}

func (s *FilterList) SetValue(v string) *FilterList {
	s.Value = &v
	return s
}

func (s *FilterList) Validate() error {
	return dara.Validate(s)
}
