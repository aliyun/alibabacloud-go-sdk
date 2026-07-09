// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObserveResourceGlobalScopeFilter interface {
	dara.Model
	String() string
	GoString() string
	SetEq(v bool) *ObserveResourceGlobalScopeFilter
	GetEq() *bool
}

type ObserveResourceGlobalScopeFilter struct {
	// Specifies whether to retrieve only resources that are exclusive to the global scope.
	Eq *bool `json:"eq,omitempty" xml:"eq,omitempty"`
}

func (s ObserveResourceGlobalScopeFilter) String() string {
	return dara.Prettify(s)
}

func (s ObserveResourceGlobalScopeFilter) GoString() string {
	return s.String()
}

func (s *ObserveResourceGlobalScopeFilter) GetEq() *bool {
	return s.Eq
}

func (s *ObserveResourceGlobalScopeFilter) SetEq(v bool) *ObserveResourceGlobalScopeFilter {
	s.Eq = &v
	return s
}

func (s *ObserveResourceGlobalScopeFilter) Validate() error {
	return dara.Validate(s)
}
