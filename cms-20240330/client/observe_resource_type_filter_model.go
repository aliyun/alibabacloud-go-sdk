// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObserveResourceTypeFilter interface {
	dara.Model
	String() string
	GoString() string
	SetEq(v string) *ObserveResourceTypeFilter
	GetEq() *string
}

type ObserveResourceTypeFilter struct {
	// The exact resource type to query. Returns only resources of this type.
	Eq *string `json:"eq,omitempty" xml:"eq,omitempty"`
}

func (s ObserveResourceTypeFilter) String() string {
	return dara.Prettify(s)
}

func (s ObserveResourceTypeFilter) GoString() string {
	return s.String()
}

func (s *ObserveResourceTypeFilter) GetEq() *string {
	return s.Eq
}

func (s *ObserveResourceTypeFilter) SetEq(v string) *ObserveResourceTypeFilter {
	s.Eq = &v
	return s
}

func (s *ObserveResourceTypeFilter) Validate() error {
	return dara.Validate(s)
}
