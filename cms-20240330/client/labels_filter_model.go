// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLabelsFilter interface {
	dara.Model
	String() string
	GoString() string
	SetAllOf(v []*LabelMatcher) *LabelsFilter
	GetAllOf() []*LabelMatcher
	SetAnyOf(v []*LabelMatcher) *LabelsFilter
	GetAnyOf() []*LabelMatcher
}

type LabelsFilter struct {
	// An array of `LabelMatcher` requirements. An object is selected only if it satisfies all of the requirements in this list (a logical AND). If provided, the list cannot be empty.
	AllOf []*LabelMatcher `json:"allOf,omitempty" xml:"allOf,omitempty" type:"Repeated"`
	// An array of `LabelMatcher` requirements. An object is selected if it satisfies at least one of the requirements in this list (a logical OR). If provided, the list cannot be empty.
	AnyOf []*LabelMatcher `json:"anyOf,omitempty" xml:"anyOf,omitempty" type:"Repeated"`
}

func (s LabelsFilter) String() string {
	return dara.Prettify(s)
}

func (s LabelsFilter) GoString() string {
	return s.String()
}

func (s *LabelsFilter) GetAllOf() []*LabelMatcher {
	return s.AllOf
}

func (s *LabelsFilter) GetAnyOf() []*LabelMatcher {
	return s.AnyOf
}

func (s *LabelsFilter) SetAllOf(v []*LabelMatcher) *LabelsFilter {
	s.AllOf = v
	return s
}

func (s *LabelsFilter) SetAnyOf(v []*LabelMatcher) *LabelsFilter {
	s.AnyOf = v
	return s
}

func (s *LabelsFilter) Validate() error {
	if s.AllOf != nil {
		for _, item := range s.AllOf {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AnyOf != nil {
		for _, item := range s.AnyOf {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
