// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrefixFilters interface {
	dara.Model
	String() string
	GoString() string
	SetIncludes(v *PrefixFiltersIncludeItem) *PrefixFilters
	GetIncludes() *PrefixFiltersIncludeItem
}

type PrefixFilters struct {
	// The collection of object prefixes to migrate, where multiple prefixes are matched using OR logic.
	Includes *PrefixFiltersIncludeItem `json:"Includes,omitempty" xml:"Includes,omitempty"`
}

func (s PrefixFilters) String() string {
	return dara.Prettify(s)
}

func (s PrefixFilters) GoString() string {
	return s.String()
}

func (s *PrefixFilters) GetIncludes() *PrefixFiltersIncludeItem {
	return s.Includes
}

func (s *PrefixFilters) SetIncludes(v *PrefixFiltersIncludeItem) *PrefixFilters {
	s.Includes = v
	return s
}

func (s *PrefixFilters) Validate() error {
	if s.Includes != nil {
		if err := s.Includes.Validate(); err != nil {
			return err
		}
	}
	return nil
}
