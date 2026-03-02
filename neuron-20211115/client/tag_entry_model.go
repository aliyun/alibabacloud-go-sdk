// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagEntry interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *TagEntry
	GetKey() *string
	SetValue(v string) *TagEntry
	GetValue() *string
}

type TagEntry struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s TagEntry) String() string {
	return dara.Prettify(s)
}

func (s TagEntry) GoString() string {
	return s.String()
}

func (s *TagEntry) GetKey() *string {
	return s.Key
}

func (s *TagEntry) GetValue() *string {
	return s.Value
}

func (s *TagEntry) SetKey(v string) *TagEntry {
	s.Key = &v
	return s
}

func (s *TagEntry) SetValue(v string) *TagEntry {
	s.Value = &v
	return s
}

func (s *TagEntry) Validate() error {
	return dara.Validate(s)
}
