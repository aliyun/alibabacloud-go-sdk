// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObserveResourceListFilter interface {
	dara.Model
	String() string
	GoString() string
	SetContains(v []*string) *ObserveResourceListFilter
	GetContains() []*string
}

type ObserveResourceListFilter struct {
	// The match condition that requires the observeResourceList of a rule to contain at least one instance ID from the array (OR semantics).
	//
	// example:
	//
	// ["i-bp1a2b3c4d5e6f7g8h9i"]
	Contains []*string `json:"contains,omitempty" xml:"contains,omitempty" type:"Repeated"`
}

func (s ObserveResourceListFilter) String() string {
	return dara.Prettify(s)
}

func (s ObserveResourceListFilter) GoString() string {
	return s.String()
}

func (s *ObserveResourceListFilter) GetContains() []*string {
	return s.Contains
}

func (s *ObserveResourceListFilter) SetContains(v []*string) *ObserveResourceListFilter {
	s.Contains = v
	return s
}

func (s *ObserveResourceListFilter) Validate() error {
	return dara.Validate(s)
}
