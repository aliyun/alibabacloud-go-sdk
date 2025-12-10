// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLastModifiedFilters interface {
	dara.Model
	String() string
	GoString() string
	SetExcludes(v *LastModifyFilterItem) *LastModifiedFilters
	GetExcludes() *LastModifyFilterItem
	SetIncludes(v *LastModifyFilterItem) *LastModifiedFilters
	GetIncludes() *LastModifyFilterItem
}

type LastModifiedFilters struct {
	Excludes *LastModifyFilterItem `json:"Excludes,omitempty" xml:"Excludes,omitempty"`
	Includes *LastModifyFilterItem `json:"Includes,omitempty" xml:"Includes,omitempty"`
}

func (s LastModifiedFilters) String() string {
	return dara.Prettify(s)
}

func (s LastModifiedFilters) GoString() string {
	return s.String()
}

func (s *LastModifiedFilters) GetExcludes() *LastModifyFilterItem {
	return s.Excludes
}

func (s *LastModifiedFilters) GetIncludes() *LastModifyFilterItem {
	return s.Includes
}

func (s *LastModifiedFilters) SetExcludes(v *LastModifyFilterItem) *LastModifiedFilters {
	s.Excludes = v
	return s
}

func (s *LastModifiedFilters) SetIncludes(v *LastModifyFilterItem) *LastModifiedFilters {
	s.Includes = v
	return s
}

func (s *LastModifiedFilters) Validate() error {
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
