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
	// The HTTP status code.
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
	// Indicates whether the request is successful.
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
	return dara.Validate(s)
}

type DescribeCostBudgetsSummaryResponseBodyData struct {
	// The site of the host.
	//
	// example:
	//
	// cn
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The data that is returned.
	Items []*DescribeCostBudgetsSummaryResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The maximum number of entries that are returned.
	//
	// example:
	//
	// 5
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page
	//
	// example:
	//
	// eyJwYWdlTnVtIjoyLCJwYWdlU2l6ZSI6NH0=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of entries that are returned.
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
	return dara.Validate(s)
}

type DescribeCostBudgetsSummaryResponseBodyDataItems struct {
	// The information about the budget. The BudgetCycleType parameter indicates the cycle of the budget. Valid values: daily, monthly, quarterly, and yearly. The TotalBudgetAmount parameter indicates the total budget. The BudgetMemo parameter indicates the remarks of the budget.
	//
	// example:
	//
	// {           "budgetCycleType": "monthly",           "budgetMemo": "",           "totalBudgetAmount": 220         }
	Budget map[string]interface{} `json:"Budget,omitempty" xml:"Budget,omitempty"`
	// The name of the budget.
	//
	// example:
	//
	// Annual budget
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
	// The information about the estimate-to-actual analysis. The ActualConsumeSum parameter indicates the accumulated actual value. The ActualAddForecastedAmount parameter indicates the sum of accumulated actual value and predicted value. If the BudgetType parameter is set to cost, the sum of actual value and predicted value includes the actual cost incurred from the budget start date to the current date and the predicted cost from the current date to the budget end date. If the BudgetType parameter is set to asset, the sum of actual value and predicted value includes the actual usage or coverage from the budget start date to the budget end date. If the budget end date minus the current date is more than one year, the part that exceeds one year is not included. The ActualAndBudgetComparison parameter indicates the comparison between the actual value and the predicted value. The value of the ActualAndBudgetComparison parameter is calculated based on the following formula: Accumulated actual value/Total budget × 100%.
	//
	// example:
	//
	// {           "actualConsumeSum": 88.6,           "actualAddForecastedAmount": 89.6,           "actualAndBudgetComparison": "20.73%"         }
	CalculatedValues map[string]interface{} `json:"CalculatedValues,omitempty" xml:"CalculatedValues,omitempty"`
	// The information about the billing cycle. The ConsumePeriodBegin parameter indicates the start date of the budget. The ConsumePeriodEnd parameter indicates the end date of the budget.
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
