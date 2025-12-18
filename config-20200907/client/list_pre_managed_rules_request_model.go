// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPreManagedRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListPreManagedRulesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListPreManagedRulesRequest
	GetPageSize() *int64
	SetResourceTypes(v []*string) *ListPreManagedRulesRequest
	GetResourceTypes() []*string
}

type ListPreManagedRulesRequest struct {
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
	ResourceTypes []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
}

func (s ListPreManagedRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPreManagedRulesRequest) GoString() string {
	return s.String()
}

func (s *ListPreManagedRulesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListPreManagedRulesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListPreManagedRulesRequest) GetResourceTypes() []*string {
	return s.ResourceTypes
}

func (s *ListPreManagedRulesRequest) SetPageNumber(v int64) *ListPreManagedRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPreManagedRulesRequest) SetPageSize(v int64) *ListPreManagedRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListPreManagedRulesRequest) SetResourceTypes(v []*string) *ListPreManagedRulesRequest {
	s.ResourceTypes = v
	return s
}

func (s *ListPreManagedRulesRequest) Validate() error {
	return dara.Validate(s)
}
