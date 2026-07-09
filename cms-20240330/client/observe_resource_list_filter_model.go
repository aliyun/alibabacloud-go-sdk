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
	// Returns resources if a specified property contains any string in this array.
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
