// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComponentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComponentName(v string) *ListComponentsRequest
	GetComponentName() *string
	SetLang(v string) *ListComponentsRequest
	GetLang() *string
	SetMaxResults(v int32) *ListComponentsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListComponentsRequest
	GetNextToken() *string
	SetPageNumber(v int32) *ListComponentsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListComponentsRequest
	GetPageSize() *int32
	SetRoleFor(v int64) *ListComponentsRequest
	GetRoleFor() *int64
}

type ListComponentsRequest struct {
	// The name of the component.
	//
	// example:
	//
	// SLS
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// The language type for the request and response. Values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The size of the page. Range: 1~100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token to start the next page query.
	//
	// example:
	//
	// kt0BLbenY2XCyRfsmoEcVg==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number for pagination, with a default value of 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of items per page for pagination. Range: 1~100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Resource directory member account ID.
	//
	// example:
	//
	// 13760*****718726
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s ListComponentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListComponentsRequest) GoString() string {
	return s.String()
}

func (s *ListComponentsRequest) GetComponentName() *string {
	return s.ComponentName
}

func (s *ListComponentsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListComponentsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListComponentsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListComponentsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListComponentsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListComponentsRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListComponentsRequest) SetComponentName(v string) *ListComponentsRequest {
	s.ComponentName = &v
	return s
}

func (s *ListComponentsRequest) SetLang(v string) *ListComponentsRequest {
	s.Lang = &v
	return s
}

func (s *ListComponentsRequest) SetMaxResults(v int32) *ListComponentsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListComponentsRequest) SetNextToken(v string) *ListComponentsRequest {
	s.NextToken = &v
	return s
}

func (s *ListComponentsRequest) SetPageNumber(v int32) *ListComponentsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListComponentsRequest) SetPageSize(v int32) *ListComponentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListComponentsRequest) SetRoleFor(v int64) *ListComponentsRequest {
	s.RoleFor = &v
	return s
}

func (s *ListComponentsRequest) Validate() error {
	return dara.Validate(s)
}
