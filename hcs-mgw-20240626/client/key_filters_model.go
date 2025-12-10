// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKeyFilters interface {
	dara.Model
	String() string
	GoString() string
	SetExcludes(v *KeyFilterItem) *KeyFilters
	GetExcludes() *KeyFilterItem
	SetIncludes(v *KeyFilterItem) *KeyFilters
	GetIncludes() *KeyFilterItem
}

type KeyFilters struct {
	Excludes *KeyFilterItem `json:"Excludes,omitempty" xml:"Excludes,omitempty"`
	Includes *KeyFilterItem `json:"Includes,omitempty" xml:"Includes,omitempty"`
}

func (s KeyFilters) String() string {
	return dara.Prettify(s)
}

func (s KeyFilters) GoString() string {
	return s.String()
}

func (s *KeyFilters) GetExcludes() *KeyFilterItem {
	return s.Excludes
}

func (s *KeyFilters) GetIncludes() *KeyFilterItem {
	return s.Includes
}

func (s *KeyFilters) SetExcludes(v *KeyFilterItem) *KeyFilters {
	s.Excludes = v
	return s
}

func (s *KeyFilters) SetIncludes(v *KeyFilterItem) *KeyFilters {
	s.Includes = v
	return s
}

func (s *KeyFilters) Validate() error {
	if s.Excludes != nil {
		if err := s.Excludes.Validate(); err != nil {
			return err
		}
	}
	if s.Includes != nil {
		if err := s.Includes.Validate(); err != nil {
			return err
		}
	}
	return nil
}
