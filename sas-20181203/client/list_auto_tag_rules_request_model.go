// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAutoTagRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListAutoTagRulesRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *ListAutoTagRulesRequest
	GetPageSize() *int32
	SetRuleName(v string) *ListAutoTagRulesRequest
	GetRuleName() *string
}

type ListAutoTagRulesRequest struct {
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// text-001
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s ListAutoTagRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAutoTagRulesRequest) GoString() string {
	return s.String()
}

func (s *ListAutoTagRulesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAutoTagRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAutoTagRulesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListAutoTagRulesRequest) SetCurrentPage(v int32) *ListAutoTagRulesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListAutoTagRulesRequest) SetPageSize(v int32) *ListAutoTagRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAutoTagRulesRequest) SetRuleName(v string) *ListAutoTagRulesRequest {
	s.RuleName = &v
	return s
}

func (s *ListAutoTagRulesRequest) Validate() error {
	return dara.Validate(s)
}
