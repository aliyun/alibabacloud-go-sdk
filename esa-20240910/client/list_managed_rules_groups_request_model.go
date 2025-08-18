// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListManagedRulesGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListManagedRulesGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListManagedRulesGroupsRequest
	GetPageSize() *int32
}

type ListManagedRulesGroupsRequest struct {
	// Page number, used to specify the page number for pagination queries.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size, used to specify the number of items per page for pagination queries.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListManagedRulesGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListManagedRulesGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListManagedRulesGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListManagedRulesGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListManagedRulesGroupsRequest) SetPageNumber(v int32) *ListManagedRulesGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListManagedRulesGroupsRequest) SetPageSize(v int32) *ListManagedRulesGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListManagedRulesGroupsRequest) Validate() error {
	return dara.Validate(s)
}
