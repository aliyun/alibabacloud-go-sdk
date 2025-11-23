// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAbacPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListAbacPoliciesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListAbacPoliciesRequest
	GetPageSize() *int64
	SetSearchKey(v string) *ListAbacPoliciesRequest
	GetSearchKey() *string
	SetTid(v int64) *ListAbacPoliciesRequest
	GetTid() *int64
}

type ListAbacPoliciesRequest struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The search keyword. Fuzzy match is supported.
	//
	// example:
	//
	// policy_test
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// The ID of the tenant.
	//
	// > To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListAbacPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAbacPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListAbacPoliciesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListAbacPoliciesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAbacPoliciesRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *ListAbacPoliciesRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListAbacPoliciesRequest) SetPageNumber(v int64) *ListAbacPoliciesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAbacPoliciesRequest) SetPageSize(v int64) *ListAbacPoliciesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAbacPoliciesRequest) SetSearchKey(v string) *ListAbacPoliciesRequest {
	s.SearchKey = &v
	return s
}

func (s *ListAbacPoliciesRequest) SetTid(v int64) *ListAbacPoliciesRequest {
	s.Tid = &v
	return s
}

func (s *ListAbacPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
