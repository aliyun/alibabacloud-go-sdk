// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPreManagedRulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListPreManagedRulesShrinkRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListPreManagedRulesShrinkRequest
	GetPageSize() *int64
	SetResourceTypesShrink(v string) *ListPreManagedRulesShrinkRequest
	GetResourceTypesShrink() *string
}

type ListPreManagedRulesShrinkRequest struct {
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Default value: 10.
	//
	// example:
	//
	// 1
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the resource.
	ResourceTypesShrink *string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty"`
}

func (s ListPreManagedRulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPreManagedRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListPreManagedRulesShrinkRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListPreManagedRulesShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListPreManagedRulesShrinkRequest) GetResourceTypesShrink() *string {
	return s.ResourceTypesShrink
}

func (s *ListPreManagedRulesShrinkRequest) SetPageNumber(v int64) *ListPreManagedRulesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPreManagedRulesShrinkRequest) SetPageSize(v int64) *ListPreManagedRulesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListPreManagedRulesShrinkRequest) SetResourceTypesShrink(v string) *ListPreManagedRulesShrinkRequest {
	s.ResourceTypesShrink = &v
	return s
}

func (s *ListPreManagedRulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
