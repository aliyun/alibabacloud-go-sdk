// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnifiedCostCredits interface {
	dara.Model
	String() string
	GoString() string
	SetSearch(v *SearchCredits) *UnifiedCostCredits
	GetSearch() *SearchCredits
	SetValueAdded(v *ValueAddedCredits) *UnifiedCostCredits
	GetValueAdded() *ValueAddedCredits
}

type UnifiedCostCredits struct {
	Search     *SearchCredits     `json:"search,omitempty" xml:"search,omitempty"`
	ValueAdded *ValueAddedCredits `json:"valueAdded,omitempty" xml:"valueAdded,omitempty"`
}

func (s UnifiedCostCredits) String() string {
	return dara.Prettify(s)
}

func (s UnifiedCostCredits) GoString() string {
	return s.String()
}

func (s *UnifiedCostCredits) GetSearch() *SearchCredits {
	return s.Search
}

func (s *UnifiedCostCredits) GetValueAdded() *ValueAddedCredits {
	return s.ValueAdded
}

func (s *UnifiedCostCredits) SetSearch(v *SearchCredits) *UnifiedCostCredits {
	s.Search = v
	return s
}

func (s *UnifiedCostCredits) SetValueAdded(v *ValueAddedCredits) *UnifiedCostCredits {
	s.ValueAdded = v
	return s
}

func (s *UnifiedCostCredits) Validate() error {
	return dara.Validate(s)
}
