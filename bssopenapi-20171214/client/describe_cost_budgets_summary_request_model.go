// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCostBudgetsSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBudgetName(v string) *DescribeCostBudgetsSummaryRequest
	GetBudgetName() *string
	SetBudgetStatus(v string) *DescribeCostBudgetsSummaryRequest
	GetBudgetStatus() *string
	SetBudgetType(v string) *DescribeCostBudgetsSummaryRequest
	GetBudgetType() *string
	SetMaxResults(v int32) *DescribeCostBudgetsSummaryRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeCostBudgetsSummaryRequest
	GetNextToken() *string
}

type DescribeCostBudgetsSummaryRequest struct {
	// The name of the budget. Supports fuzzy search.
	//
	// example:
	//
	// 全年预算
	BudgetName *string `json:"BudgetName,omitempty" xml:"BudgetName,omitempty"`
	// The status of the budget. If you do not specify this parameter, budgets in all statuses are returned.
	//
	// example:
	//
	// notOverdue
	BudgetStatus *string `json:"BudgetStatus,omitempty" xml:"BudgetStatus,omitempty"`
	// The type of the budget. If you do not specify this parameter, budgets of all types are returned.
	//
	// example:
	//
	// cost
	BudgetType *string `json:"BudgetType,omitempty" xml:"BudgetType,omitempty"`
	// The maximum number of items to return. Valid values: 1 to 10. Default value: 10.
	//
	// example:
	//
	// 5
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results. Set this parameter to the NextToken value that is returned in the response to the previous request. If you do not specify this parameter, the query starts from the first page.
	//
	// example:
	//
	// eyJwYWdlTnVtIjoyLCJwYWdlU2l6ZSI6NH0=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeCostBudgetsSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCostBudgetsSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeCostBudgetsSummaryRequest) GetBudgetName() *string {
	return s.BudgetName
}

func (s *DescribeCostBudgetsSummaryRequest) GetBudgetStatus() *string {
	return s.BudgetStatus
}

func (s *DescribeCostBudgetsSummaryRequest) GetBudgetType() *string {
	return s.BudgetType
}

func (s *DescribeCostBudgetsSummaryRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeCostBudgetsSummaryRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeCostBudgetsSummaryRequest) SetBudgetName(v string) *DescribeCostBudgetsSummaryRequest {
	s.BudgetName = &v
	return s
}

func (s *DescribeCostBudgetsSummaryRequest) SetBudgetStatus(v string) *DescribeCostBudgetsSummaryRequest {
	s.BudgetStatus = &v
	return s
}

func (s *DescribeCostBudgetsSummaryRequest) SetBudgetType(v string) *DescribeCostBudgetsSummaryRequest {
	s.BudgetType = &v
	return s
}

func (s *DescribeCostBudgetsSummaryRequest) SetMaxResults(v int32) *DescribeCostBudgetsSummaryRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeCostBudgetsSummaryRequest) SetNextToken(v string) *DescribeCostBudgetsSummaryRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeCostBudgetsSummaryRequest) Validate() error {
	return dara.Validate(s)
}
