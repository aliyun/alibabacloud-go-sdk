// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisplayNameFilter interface {
	dara.Model
	String() string
	GoString() string
	SetContains(v string) *DisplayNameFilter
	GetContains() *string
	SetNotContains(v string) *DisplayNameFilter
	GetNotContains() *string
}

type DisplayNameFilter struct {
	Contains    *string `json:"contains,omitempty" xml:"contains,omitempty"`
	NotContains *string `json:"notContains,omitempty" xml:"notContains,omitempty"`
}

func (s DisplayNameFilter) String() string {
	return dara.Prettify(s)
}

func (s DisplayNameFilter) GoString() string {
	return s.String()
}

func (s *DisplayNameFilter) GetContains() *string {
	return s.Contains
}

func (s *DisplayNameFilter) GetNotContains() *string {
	return s.NotContains
}

func (s *DisplayNameFilter) SetContains(v string) *DisplayNameFilter {
	s.Contains = &v
	return s
}

func (s *DisplayNameFilter) SetNotContains(v string) *DisplayNameFilter {
	s.NotContains = &v
	return s
}

func (s *DisplayNameFilter) Validate() error {
	return dara.Validate(s)
}
