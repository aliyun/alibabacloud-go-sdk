// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAgentThemeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *ListDataAgentThemeRequest
	GetCategory() *string
	SetMaxResults(v int32) *ListDataAgentThemeRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataAgentThemeRequest
	GetNextToken() *string
	SetPageNumber(v int32) *ListDataAgentThemeRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataAgentThemeRequest
	GetPageSize() *int32
	SetThemeFrom(v string) *ListDataAgentThemeRequest
	GetThemeFrom() *string
	SetThemeType(v string) *ListDataAgentThemeRequest
	GetThemeType() *string
}

type ListDataAgentThemeRequest struct {
	// The common scenarios. Valid values: report, infographic, and others.
	//
	// example:
	//
	// report
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// **[Not supported]*	- The page size. Maximum value: 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// **[Not supported]*	- The pagination token for the next query. Valid values:
	//
	// - If **NextToken*	- is empty, no next query exists.
	//
	// - If **NextToken*	- has a return value, the value indicates the token for the next query.
	//
	// example:
	//
	// f056501ada12****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The current page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The source of the theme. Valid values:
	//
	// - system
	//
	// - custom
	//
	// - derived
	//
	// example:
	//
	// custom
	ThemeFrom *string `json:"ThemeFrom,omitempty" xml:"ThemeFrom,omitempty"`
	// The theme stage. Valid values:
	//
	// - design: contains only design.md.
	//
	// - template: complete and renderable.
	//
	// example:
	//
	// template
	ThemeType *string `json:"ThemeType,omitempty" xml:"ThemeType,omitempty"`
}

func (s ListDataAgentThemeRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentThemeRequest) GoString() string {
	return s.String()
}

func (s *ListDataAgentThemeRequest) GetCategory() *string {
	return s.Category
}

func (s *ListDataAgentThemeRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataAgentThemeRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataAgentThemeRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataAgentThemeRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataAgentThemeRequest) GetThemeFrom() *string {
	return s.ThemeFrom
}

func (s *ListDataAgentThemeRequest) GetThemeType() *string {
	return s.ThemeType
}

func (s *ListDataAgentThemeRequest) SetCategory(v string) *ListDataAgentThemeRequest {
	s.Category = &v
	return s
}

func (s *ListDataAgentThemeRequest) SetMaxResults(v int32) *ListDataAgentThemeRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDataAgentThemeRequest) SetNextToken(v string) *ListDataAgentThemeRequest {
	s.NextToken = &v
	return s
}

func (s *ListDataAgentThemeRequest) SetPageNumber(v int32) *ListDataAgentThemeRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataAgentThemeRequest) SetPageSize(v int32) *ListDataAgentThemeRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataAgentThemeRequest) SetThemeFrom(v string) *ListDataAgentThemeRequest {
	s.ThemeFrom = &v
	return s
}

func (s *ListDataAgentThemeRequest) SetThemeType(v string) *ListDataAgentThemeRequest {
	s.ThemeType = &v
	return s
}

func (s *ListDataAgentThemeRequest) Validate() error {
	return dara.Validate(s)
}
