// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInterceptionRulePageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListInterceptionRulePageRequest
	GetClusterId() *string
	SetCriteria(v string) *ListInterceptionRulePageRequest
	GetCriteria() *string
	SetCriteriaType(v string) *ListInterceptionRulePageRequest
	GetCriteriaType() *string
	SetCurrentPage(v int32) *ListInterceptionRulePageRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *ListInterceptionRulePageRequest
	GetPageSize() *int32
}

type ListInterceptionRulePageRequest struct {
	// The ID of the container cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc41de13ab5474210bc0ce772a009****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The query condition.
	//
	// example:
	//
	// 80
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// The type of the query condition. Valid values:
	//
	// 	- **ID**
	//
	// 	- **RULE_NAME**
	//
	// 	- **SRC_TARGET**
	//
	// 	- **DST_TARGET**
	//
	// 	- **DST_PORT**
	//
	// 	- **RULE_SWITCH**
	//
	// 	- **INTERCEPTOR_TYPE**
	//
	// example:
	//
	// DST_PORT
	CriteriaType *string `json:"CriteriaType,omitempty" xml:"CriteriaType,omitempty"`
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListInterceptionRulePageRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInterceptionRulePageRequest) GoString() string {
	return s.String()
}

func (s *ListInterceptionRulePageRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListInterceptionRulePageRequest) GetCriteria() *string {
	return s.Criteria
}

func (s *ListInterceptionRulePageRequest) GetCriteriaType() *string {
	return s.CriteriaType
}

func (s *ListInterceptionRulePageRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListInterceptionRulePageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInterceptionRulePageRequest) SetClusterId(v string) *ListInterceptionRulePageRequest {
	s.ClusterId = &v
	return s
}

func (s *ListInterceptionRulePageRequest) SetCriteria(v string) *ListInterceptionRulePageRequest {
	s.Criteria = &v
	return s
}

func (s *ListInterceptionRulePageRequest) SetCriteriaType(v string) *ListInterceptionRulePageRequest {
	s.CriteriaType = &v
	return s
}

func (s *ListInterceptionRulePageRequest) SetCurrentPage(v int32) *ListInterceptionRulePageRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListInterceptionRulePageRequest) SetPageSize(v int32) *ListInterceptionRulePageRequest {
	s.PageSize = &v
	return s
}

func (s *ListInterceptionRulePageRequest) Validate() error {
	return dara.Validate(s)
}
