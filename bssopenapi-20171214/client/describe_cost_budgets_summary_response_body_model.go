// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCostBudgetsSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeCostBudgetsSummaryResponseBody
	GetCode() *string
	SetData(v *DescribeCostBudgetsSummaryResponseBodyData) *DescribeCostBudgetsSummaryResponseBody
	GetData() *DescribeCostBudgetsSummaryResponseBodyData
	SetMessage(v string) *DescribeCostBudgetsSummaryResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeCostBudgetsSummaryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeCostBudgetsSummaryResponseBody
	GetSuccess() *bool
}

type DescribeCostBudgetsSummaryResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data that is returned.
	Data *DescribeCostBudgetsSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 3BFC23FE-A182-4D96-A1E4-7521B30B8E43
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCostBudgetsSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCostBudgetsSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCostBudgetsSummaryResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeCostBudgetsSummaryResponseBody) GetData() *DescribeCostBudgetsSummaryResponseBodyData {
	return s.Data
}

func (s *DescribeCostBudgetsSummaryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeCostBudgetsSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCostBudgetsSummaryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeCostBudgetsSummaryResponseBody) SetCode(v string) *DescribeCostBudgetsSummaryResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCostBudgetsSummaryResponseBody) SetData(v *DescribeCostBudgetsSummaryResponseBodyData) *DescribeCostBudgetsSummaryResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCostBudgetsSummaryResponseBody) SetMessage(v string) *DescribeCostBudgetsSummaryResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCostBudgetsSummaryResponseBody) SetRequestId(v string) *DescribeCostBudgetsSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCostBudgetsSummaryResponseBody) SetSuccess(v bool) *DescribeCostBudgetsSummaryResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeCostBudgetsSummaryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCostBudgetsSummaryResponseBodyData struct {
	// The website.
	//
	// example:
	//
	// cn
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The list of items.
	Items []*DescribeCostBudgetsSummaryResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The maximum number of entries returned on the current page.
	//
	// example:
	//
	// 5
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next page of results.
	//
	// example:
	//
	// eyJwYWdlTnVtIjoyLCJwYWdlU2l6ZSI6NH0=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCostBudgetsSummaryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCostBudgetsSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCostBudgetsSummaryResponseBodyData) GetHostId() *string {
	return s.HostId
}

func (s *DescribeCostBudgetsSummaryResponseBodyData) GetItems() []*DescribeCostBudgetsSummaryResponseBodyDataItems {
	return s.Items
}

func (s *DescribeCostBudgetsSummaryResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeCostBudgetsSummaryResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeCostBudgetsSummaryResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCostBudgetsSummaryResponseBodyData) SetHostId(v string) *DescribeCostBudgetsSummaryResponseBodyData {
	s.HostId = &v
	return s
}

func (s *DescribeCostBudgetsSummaryResponseBodyData) SetItems(v []*DescribeCostBudgetsSummaryResponseBodyDataItems) *DescribeCostBudgetsSummaryResponseBodyData {
	s.Items = v
	return s
}

func (s *DescribeCostBudgetsSummaryResponseBodyData) SetMaxResults(v int32) *DescribeCostBudgetsSummaryResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *DescribeCostBudgetsSummaryResponseBodyData) SetNextToken(v string) *DescribeCostBudgetsSummaryResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *DescribeCostBudgetsSummaryResponseBodyData) SetTotalCount(v int32) *DescribeCostBudgetsSummaryResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeCostBudgetsSummaryResponseBodyData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCostBudgetsSummaryResponseBodyDataItems struct {
	// Information about the budget.
	//
	// - BudgetCycleType: The budget period. Valid values: daily, monthly, quarterly, and yearly.
	//
	// - TotalBudgetAmount: The total budget amount.
	//
	// - BudgetMemo: The notes for the budget.
	//
	// example:
	//
	// {           "budgetCycleType": "monthly",           "budgetMemo": "",           "totalBudgetAmount": 220         }
	Budget map[string]interface{} `json:"Budget,omitempty" xml:"Budget,omitempty"`
	// The name of the budget.
	//
	// example:
	//
	// 全年预算
	BudgetName *string `json:"BudgetName,omitempty" xml:"BudgetName,omitempty"`
	// The status of the budget.
	//
	// example:
	//
	// overdue
	BudgetStatus *string `json:"BudgetStatus,omitempty" xml:"BudgetStatus,omitempty"`
	// The type of the budget.
	//
	// example:
	//
	// cost
	BudgetType *string `json:"BudgetType,omitempty" xml:"BudgetType,omitempty"`
	// Information about the estimate-to-actual comparison.
	//
	// - ActualConsumeSum: The accumulated actual consumption.
	//
	// - ActualAddForecastedAmount: The sum of the accumulated actual value and the predicted value. For a cost budget, this value represents the actual spending to date plus the forecasted spending until the budget period ends.
	//
	// - ActualAndBudgetComparison: The comparison of the actual consumption to the budget. This value is calculated using the following formula: (Accumulated actual value / Total budget amount) × 100%.
	//
	// example:
	//
	// {           "actualConsumeSum": 88.6,           "actualAddForecastedAmount": 89.6,           "actualAndBudgetComparison": "20.73%"         }
	CalculatedValues map[string]interface{} `json:"CalculatedValues,omitempty" xml:"CalculatedValues,omitempty"`
	// Information about the billing cycle.
	//
	// - ConsumePeriodBegin: The start date of the budget.
	//
	// - ConsumePeriodEnd: The end date of the budget.
	//
	// example:
	//
	// {           "consumePeriodBegin": "2022-10",           "consumePeriodEnd": "2022-11"         }
	ConsumePeriod map[string]interface{} `json:"ConsumePeriod,omitempty" xml:"ConsumePeriod,omitempty"`
}

func (s DescribeCostBudgetsSummaryResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeCostBudgetsSummaryResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *DescribeCostBudgetsSummaryResponseBodyDataItems) GetBudget() map[string]interface{} {
	return s.Budget
}

func (s *DescribeCostBudgetsSummaryResponseBodyDataItems) GetBudgetName() *string {
	return s.BudgetName
}

func (s *DescribeCostBudgetsSummaryResponseBodyDataItems) GetBudgetStatus() *string {
	return s.BudgetStatus
}

func (s *DescribeCostBudgetsSummaryResponseBodyDataItems) GetBudgetType() *string {
	return s.BudgetType
}

func (s *DescribeCostBudgetsSummaryResponseBodyDataItems) GetCalculatedValues() map[string]interface{} {
	return s.CalculatedValues
}

func (s *DescribeCostBudgetsSummaryResponseBodyDataItems) GetConsumePeriod() map[string]interface{} {
	return s.ConsumePeriod
}

func (s *DescribeCostBudgetsSummaryResponseBodyDataItems) SetBudget(v map[string]interface{}) *DescribeCostBudgetsSummaryResponseBodyDataItems {
	s.Budget = v
	return s
}

func (s *DescribeCostBudgetsSummaryResponseBodyDataItems) SetBudgetName(v string) *DescribeCostBudgetsSummaryResponseBodyDataItems {
	s.BudgetName = &v
	return s
}

func (s *DescribeCostBudgetsSummaryResponseBodyDataItems) SetBudgetStatus(v string) *DescribeCostBudgetsSummaryResponseBodyDataItems {
	s.BudgetStatus = &v
	return s
}

func (s *DescribeCostBudgetsSummaryResponseBodyDataItems) SetBudgetType(v string) *DescribeCostBudgetsSummaryResponseBodyDataItems {
	s.BudgetType = &v
	return s
}

func (s *DescribeCostBudgetsSummaryResponseBodyDataItems) SetCalculatedValues(v map[string]interface{}) *DescribeCostBudgetsSummaryResponseBodyDataItems {
	s.CalculatedValues = v
	return s
}

func (s *DescribeCostBudgetsSummaryResponseBodyDataItems) SetConsumePeriod(v map[string]interface{}) *DescribeCostBudgetsSummaryResponseBodyDataItems {
	s.ConsumePeriod = v
	return s
}

func (s *DescribeCostBudgetsSummaryResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
