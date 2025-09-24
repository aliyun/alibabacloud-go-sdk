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
	// The name of the budget. Fuzzy match is supported.
	//
	// example:
	//
	// Annual budget
	BudgetName *string `json:"BudgetName,omitempty" xml:"BudgetName,omitempty"`
	// The status of the budget. Valid values: overdue and notOverdue. A value of overdue specifies to filter expired budgets. A value of notOverdue specifies to filter budgets that do not expire. By default, if you do not specify this parameter, information about all budgets is to be returned.
	//
	// example:
	//
	// notOverdue
	BudgetStatus *string `json:"BudgetStatus,omitempty" xml:"BudgetStatus,omitempty"`
	// The type of the budget. Valid values: cost, byquantity, and asset. A value of cost specifies to filter expense budgets. A value of byquantity specifies to filter budgets calculated based on the resource usage. A value of asset specifies to filter usage or coverage budgets. By default, information about all budgets is returned if you do not specify this parameter.
	//
	// example:
	//
	// cost
	BudgetType *string `json:"BudgetType,omitempty" xml:"BudgetType,omitempty"`
	// The number of entries to return on each page. Default value: 10. Maximum value: 10. Minimum value: 1.
	//
	// example:
	//
	// 5
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The position in which the query starts. You must set this parameter to null or the token that is obtained from the previous query. Otherwise, an error is returned. If you set the NextToken parameter to null, the query starts from the beginning. The default value is null.
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
