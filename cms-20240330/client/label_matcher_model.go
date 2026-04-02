// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLabelMatcher interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *LabelMatcher
	GetKey() *string
	SetValue(v string) *LabelMatcher
	GetValue() *string
}

type LabelMatcher struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s LabelMatcher) String() string {
	return dara.Prettify(s)
}

func (s LabelMatcher) GoString() string {
	return s.String()
}

func (s *LabelMatcher) GetKey() *string {
	return s.Key
}

func (s *LabelMatcher) GetValue() *string {
	return s.Value
}

func (s *LabelMatcher) SetKey(v string) *LabelMatcher {
	s.Key = &v
	return s
}

func (s *LabelMatcher) SetValue(v string) *LabelMatcher {
	s.Value = &v
	return s
}

func (s *LabelMatcher) Validate() error {
	return dara.Validate(s)
}
