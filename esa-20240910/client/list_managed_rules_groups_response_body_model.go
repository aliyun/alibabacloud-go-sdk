// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListManagedRulesGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetManagedRulesGroups(v []*ListManagedRulesGroupsResponseBodyManagedRulesGroups) *ListManagedRulesGroupsResponseBody
	GetManagedRulesGroups() []*ListManagedRulesGroupsResponseBodyManagedRulesGroups
	SetPageNumber(v int32) *ListManagedRulesGroupsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListManagedRulesGroupsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListManagedRulesGroupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListManagedRulesGroupsResponseBody
	GetTotalCount() *int32
}

type ListManagedRulesGroupsResponseBody struct {
	// List of managed rule group information.
	ManagedRulesGroups []*ListManagedRulesGroupsResponseBodyManagedRulesGroups `json:"ManagedRulesGroups,omitempty" xml:"ManagedRulesGroups,omitempty" type:"Repeated"`
	// Current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of records after filtering.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListManagedRulesGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListManagedRulesGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListManagedRulesGroupsResponseBody) GetManagedRulesGroups() []*ListManagedRulesGroupsResponseBodyManagedRulesGroups {
	return s.ManagedRulesGroups
}

func (s *ListManagedRulesGroupsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListManagedRulesGroupsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListManagedRulesGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListManagedRulesGroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListManagedRulesGroupsResponseBody) SetManagedRulesGroups(v []*ListManagedRulesGroupsResponseBodyManagedRulesGroups) *ListManagedRulesGroupsResponseBody {
	s.ManagedRulesGroups = v
	return s
}

func (s *ListManagedRulesGroupsResponseBody) SetPageNumber(v int32) *ListManagedRulesGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListManagedRulesGroupsResponseBody) SetPageSize(v int32) *ListManagedRulesGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListManagedRulesGroupsResponseBody) SetRequestId(v string) *ListManagedRulesGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListManagedRulesGroupsResponseBody) SetTotalCount(v int32) *ListManagedRulesGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListManagedRulesGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListManagedRulesGroupsResponseBodyManagedRulesGroups struct {
	// Name of the managed rule group.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Number of rules within the managed rule group.
	//
	// example:
	//
	// 1000
	RuleCount *int64 `json:"RuleCount,omitempty" xml:"RuleCount,omitempty"`
}

func (s ListManagedRulesGroupsResponseBodyManagedRulesGroups) String() string {
	return dara.Prettify(s)
}

func (s ListManagedRulesGroupsResponseBodyManagedRulesGroups) GoString() string {
	return s.String()
}

func (s *ListManagedRulesGroupsResponseBodyManagedRulesGroups) GetName() *string {
	return s.Name
}

func (s *ListManagedRulesGroupsResponseBodyManagedRulesGroups) GetRuleCount() *int64 {
	return s.RuleCount
}

func (s *ListManagedRulesGroupsResponseBodyManagedRulesGroups) SetName(v string) *ListManagedRulesGroupsResponseBodyManagedRulesGroups {
	s.Name = &v
	return s
}

func (s *ListManagedRulesGroupsResponseBodyManagedRulesGroups) SetRuleCount(v int64) *ListManagedRulesGroupsResponseBodyManagedRulesGroups {
	s.RuleCount = &v
	return s
}

func (s *ListManagedRulesGroupsResponseBodyManagedRulesGroups) Validate() error {
	return dara.Validate(s)
}
