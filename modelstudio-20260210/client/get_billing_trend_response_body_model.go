// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBillingTrendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetBillingTrendResponseBody
	GetCode() *string
	SetData(v *GetBillingTrendResponseBodyData) *GetBillingTrendResponseBody
	GetData() *GetBillingTrendResponseBodyData
	SetMessage(v string) *GetBillingTrendResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetBillingTrendResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetBillingTrendResponseBody
	GetSuccess() *bool
}

type GetBillingTrendResponseBody struct {
	// The request result code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *GetBillingTrendResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The description of the request result.
	//
	// example:
	//
	// null
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 099A671E-FA21-5A36-8A73-918572DDEF53
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetBillingTrendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBillingTrendResponseBody) GoString() string {
	return s.String()
}

func (s *GetBillingTrendResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetBillingTrendResponseBody) GetData() *GetBillingTrendResponseBodyData {
	return s.Data
}

func (s *GetBillingTrendResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetBillingTrendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBillingTrendResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetBillingTrendResponseBody) SetCode(v string) *GetBillingTrendResponseBody {
	s.Code = &v
	return s
}

func (s *GetBillingTrendResponseBody) SetData(v *GetBillingTrendResponseBodyData) *GetBillingTrendResponseBody {
	s.Data = v
	return s
}

func (s *GetBillingTrendResponseBody) SetMessage(v string) *GetBillingTrendResponseBody {
	s.Message = &v
	return s
}

func (s *GetBillingTrendResponseBody) SetRequestId(v string) *GetBillingTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBillingTrendResponseBody) SetSuccess(v bool) *GetBillingTrendResponseBody {
	s.Success = &v
	return s
}

func (s *GetBillingTrendResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBillingTrendResponseBodyData struct {
	// The total costs within the entire query time range, including the top N items and Other.
	CostTotals *GetBillingTrendResponseBodyDataCostTotals `json:"costTotals,omitempty" xml:"costTotals,omitempty" type:"Struct"`
	// The total costs of the top N groups and the optional Other group within the period.
	GroupByTotal []*GetBillingTrendResponseBodyDataGroupByTotal `json:"groupByTotal,omitempty" xml:"groupByTotal,omitempty" type:"Repeated"`
	// The historical period data, which is returned when the includeHistoryPeriod input parameter is set to true.
	HistoryPeriod *GetBillingTrendResponseBodyDataHistoryPeriod `json:"historyPeriod,omitempty" xml:"historyPeriod,omitempty" type:"Struct"`
	// The list of cost trends sorted in ascending order by time.
	ResultByTime []*GetBillingTrendResponseBodyDataResultByTime `json:"resultByTime,omitempty" xml:"resultByTime,omitempty" type:"Repeated"`
}

func (s GetBillingTrendResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetBillingTrendResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBillingTrendResponseBodyData) GetCostTotals() *GetBillingTrendResponseBodyDataCostTotals {
	return s.CostTotals
}

func (s *GetBillingTrendResponseBodyData) GetGroupByTotal() []*GetBillingTrendResponseBodyDataGroupByTotal {
	return s.GroupByTotal
}

func (s *GetBillingTrendResponseBodyData) GetHistoryPeriod() *GetBillingTrendResponseBodyDataHistoryPeriod {
	return s.HistoryPeriod
}

func (s *GetBillingTrendResponseBodyData) GetResultByTime() []*GetBillingTrendResponseBodyDataResultByTime {
	return s.ResultByTime
}

func (s *GetBillingTrendResponseBodyData) SetCostTotals(v *GetBillingTrendResponseBodyDataCostTotals) *GetBillingTrendResponseBodyData {
	s.CostTotals = v
	return s
}

func (s *GetBillingTrendResponseBodyData) SetGroupByTotal(v []*GetBillingTrendResponseBodyDataGroupByTotal) *GetBillingTrendResponseBodyData {
	s.GroupByTotal = v
	return s
}

func (s *GetBillingTrendResponseBodyData) SetHistoryPeriod(v *GetBillingTrendResponseBodyDataHistoryPeriod) *GetBillingTrendResponseBodyData {
	s.HistoryPeriod = v
	return s
}

func (s *GetBillingTrendResponseBodyData) SetResultByTime(v []*GetBillingTrendResponseBodyDataResultByTime) *GetBillingTrendResponseBodyData {
	s.ResultByTime = v
	return s
}

func (s *GetBillingTrendResponseBodyData) Validate() error {
	if s.CostTotals != nil {
		if err := s.CostTotals.Validate(); err != nil {
			return err
		}
	}
	if s.GroupByTotal != nil {
		for _, item := range s.GroupByTotal {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.HistoryPeriod != nil {
		if err := s.HistoryPeriod.Validate(); err != nil {
			return err
		}
	}
	if s.ResultByTime != nil {
		for _, item := range s.ResultByTime {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetBillingTrendResponseBodyDataCostTotals struct {
	// The total cost amount.
	//
	// example:
	//
	// 100
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// The currency of the amount.
	//
	// example:
	//
	// CNY
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty"`
	// The pretax amount.
	//
	// example:
	//
	// 94.34
	PretaxAmount *string `json:"pretaxAmount,omitempty" xml:"pretaxAmount,omitempty"`
	// The tax amount.
	//
	// example:
	//
	// 5.66
	TaxAmount *string `json:"taxAmount,omitempty" xml:"taxAmount,omitempty"`
}

func (s GetBillingTrendResponseBodyDataCostTotals) String() string {
	return dara.Prettify(s)
}

func (s GetBillingTrendResponseBodyDataCostTotals) GoString() string {
	return s.String()
}

func (s *GetBillingTrendResponseBodyDataCostTotals) GetAmount() *string {
	return s.Amount
}

func (s *GetBillingTrendResponseBodyDataCostTotals) GetCurrency() *string {
	return s.Currency
}

func (s *GetBillingTrendResponseBodyDataCostTotals) GetPretaxAmount() *string {
	return s.PretaxAmount
}

func (s *GetBillingTrendResponseBodyDataCostTotals) GetTaxAmount() *string {
	return s.TaxAmount
}

func (s *GetBillingTrendResponseBodyDataCostTotals) SetAmount(v string) *GetBillingTrendResponseBodyDataCostTotals {
	s.Amount = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataCostTotals) SetCurrency(v string) *GetBillingTrendResponseBodyDataCostTotals {
	s.Currency = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataCostTotals) SetPretaxAmount(v string) *GetBillingTrendResponseBodyDataCostTotals {
	s.PretaxAmount = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataCostTotals) SetTaxAmount(v string) *GetBillingTrendResponseBodyDataCostTotals {
	s.TaxAmount = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataCostTotals) Validate() error {
	return dara.Validate(s)
}

type GetBillingTrendResponseBodyDataGroupByTotal struct {
	// The total amount of the current group.
	//
	// example:
	//
	// 60
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// The dimension value of the group.
	//
	// example:
	//
	// qwen-plus
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The display name of the group, which is affected by the locale input parameter.
	//
	// example:
	//
	// qwen-plus
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The pretax amount of the current group.
	//
	// example:
	//
	// 56.60
	PretaxAmount *string `json:"pretaxAmount,omitempty" xml:"pretaxAmount,omitempty"`
	// The tax amount of the current group.
	//
	// example:
	//
	// 3.40
	TaxAmount *string `json:"taxAmount,omitempty" xml:"taxAmount,omitempty"`
}

func (s GetBillingTrendResponseBodyDataGroupByTotal) String() string {
	return dara.Prettify(s)
}

func (s GetBillingTrendResponseBodyDataGroupByTotal) GoString() string {
	return s.String()
}

func (s *GetBillingTrendResponseBodyDataGroupByTotal) GetAmount() *string {
	return s.Amount
}

func (s *GetBillingTrendResponseBodyDataGroupByTotal) GetKey() *string {
	return s.Key
}

func (s *GetBillingTrendResponseBodyDataGroupByTotal) GetName() *string {
	return s.Name
}

func (s *GetBillingTrendResponseBodyDataGroupByTotal) GetPretaxAmount() *string {
	return s.PretaxAmount
}

func (s *GetBillingTrendResponseBodyDataGroupByTotal) GetTaxAmount() *string {
	return s.TaxAmount
}

func (s *GetBillingTrendResponseBodyDataGroupByTotal) SetAmount(v string) *GetBillingTrendResponseBodyDataGroupByTotal {
	s.Amount = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataGroupByTotal) SetKey(v string) *GetBillingTrendResponseBodyDataGroupByTotal {
	s.Key = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataGroupByTotal) SetName(v string) *GetBillingTrendResponseBodyDataGroupByTotal {
	s.Name = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataGroupByTotal) SetPretaxAmount(v string) *GetBillingTrendResponseBodyDataGroupByTotal {
	s.PretaxAmount = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataGroupByTotal) SetTaxAmount(v string) *GetBillingTrendResponseBodyDataGroupByTotal {
	s.TaxAmount = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataGroupByTotal) Validate() error {
	return dara.Validate(s)
}

type GetBillingTrendResponseBodyDataHistoryPeriod struct {
	// The total costs within the historical range, including the top N items and Other.
	CostTotals *GetBillingTrendResponseBodyDataHistoryPeriodCostTotals `json:"costTotals,omitempty" xml:"costTotals,omitempty" type:"Struct"`
	// The total costs of the top N groups and the optional Other group within the historical period.
	GroupByTotal []*GetBillingTrendResponseBodyDataHistoryPeriodGroupByTotal `json:"groupByTotal,omitempty" xml:"groupByTotal,omitempty" type:"Repeated"`
	// The list of cost trends sorted in ascending order by time.
	ResultByTime []*GetBillingTrendResponseBodyDataHistoryPeriodResultByTime `json:"resultByTime,omitempty" xml:"resultByTime,omitempty" type:"Repeated"`
}

func (s GetBillingTrendResponseBodyDataHistoryPeriod) String() string {
	return dara.Prettify(s)
}

func (s GetBillingTrendResponseBodyDataHistoryPeriod) GoString() string {
	return s.String()
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriod) GetCostTotals() *GetBillingTrendResponseBodyDataHistoryPeriodCostTotals {
	return s.CostTotals
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriod) GetGroupByTotal() []*GetBillingTrendResponseBodyDataHistoryPeriodGroupByTotal {
	return s.GroupByTotal
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriod) GetResultByTime() []*GetBillingTrendResponseBodyDataHistoryPeriodResultByTime {
	return s.ResultByTime
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriod) SetCostTotals(v *GetBillingTrendResponseBodyDataHistoryPeriodCostTotals) *GetBillingTrendResponseBodyDataHistoryPeriod {
	s.CostTotals = v
	return s
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriod) SetGroupByTotal(v []*GetBillingTrendResponseBodyDataHistoryPeriodGroupByTotal) *GetBillingTrendResponseBodyDataHistoryPeriod {
	s.GroupByTotal = v
	return s
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriod) SetResultByTime(v []*GetBillingTrendResponseBodyDataHistoryPeriodResultByTime) *GetBillingTrendResponseBodyDataHistoryPeriod {
	s.ResultByTime = v
	return s
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriod) Validate() error {
	if s.CostTotals != nil {
		if err := s.CostTotals.Validate(); err != nil {
			return err
		}
	}
	if s.GroupByTotal != nil {
		for _, item := range s.GroupByTotal {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ResultByTime != nil {
		for _, item := range s.ResultByTime {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetBillingTrendResponseBodyDataHistoryPeriodCostTotals struct {
	// The total cost amount.
	//
	// example:
	//
	// 1000
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// The currency of the amount.
	//
	// example:
	//
	// CNY
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty"`
	// The pretax amount.
	//
	// example:
	//
	// 800
	PretaxAmount *string `json:"pretaxAmount,omitempty" xml:"pretaxAmount,omitempty"`
	// The tax amount.
	//
	// example:
	//
	// 200
	TaxAmount *string `json:"taxAmount,omitempty" xml:"taxAmount,omitempty"`
}

func (s GetBillingTrendResponseBodyDataHistoryPeriodCostTotals) String() string {
	return dara.Prettify(s)
}

func (s GetBillingTrendResponseBodyDataHistoryPeriodCostTotals) GoString() string {
	return s.String()
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodCostTotals) GetAmount() *string {
	return s.Amount
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodCostTotals) GetCurrency() *string {
	return s.Currency
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodCostTotals) GetPretaxAmount() *string {
	return s.PretaxAmount
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodCostTotals) GetTaxAmount() *string {
	return s.TaxAmount
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodCostTotals) SetAmount(v string) *GetBillingTrendResponseBodyDataHistoryPeriodCostTotals {
	s.Amount = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodCostTotals) SetCurrency(v string) *GetBillingTrendResponseBodyDataHistoryPeriodCostTotals {
	s.Currency = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodCostTotals) SetPretaxAmount(v string) *GetBillingTrendResponseBodyDataHistoryPeriodCostTotals {
	s.PretaxAmount = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodCostTotals) SetTaxAmount(v string) *GetBillingTrendResponseBodyDataHistoryPeriodCostTotals {
	s.TaxAmount = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodCostTotals) Validate() error {
	return dara.Validate(s)
}

type GetBillingTrendResponseBodyDataHistoryPeriodGroupByTotal struct {
	// The total amount of the current group.
	//
	// example:
	//
	// 246
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// The dimension value of the group. If the value is null, DIMENSION_FILTER_NULL_VALUE is returned.
	//
	// example:
	//
	// qwen-plus
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The display name of the group, which is affected by the locale parameter. If the value is null, a hyphen (-) is displayed.
	//
	// example:
	//
	// qwen-plus
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The total pretax amount.
	//
	// example:
	//
	// 100
	PretaxAmount *string `json:"pretaxAmount,omitempty" xml:"pretaxAmount,omitempty"`
	// The total tax amount.
	//
	// example:
	//
	// 146
	TaxAmount *string `json:"taxAmount,omitempty" xml:"taxAmount,omitempty"`
}

func (s GetBillingTrendResponseBodyDataHistoryPeriodGroupByTotal) String() string {
	return dara.Prettify(s)
}

func (s GetBillingTrendResponseBodyDataHistoryPeriodGroupByTotal) GoString() string {
	return s.String()
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodGroupByTotal) GetAmount() *string {
	return s.Amount
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodGroupByTotal) GetKey() *string {
	return s.Key
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodGroupByTotal) GetName() *string {
	return s.Name
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodGroupByTotal) GetPretaxAmount() *string {
	return s.PretaxAmount
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodGroupByTotal) GetTaxAmount() *string {
	return s.TaxAmount
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodGroupByTotal) SetAmount(v string) *GetBillingTrendResponseBodyDataHistoryPeriodGroupByTotal {
	s.Amount = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodGroupByTotal) SetKey(v string) *GetBillingTrendResponseBodyDataHistoryPeriodGroupByTotal {
	s.Key = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodGroupByTotal) SetName(v string) *GetBillingTrendResponseBodyDataHistoryPeriodGroupByTotal {
	s.Name = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodGroupByTotal) SetPretaxAmount(v string) *GetBillingTrendResponseBodyDataHistoryPeriodGroupByTotal {
	s.PretaxAmount = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodGroupByTotal) SetTaxAmount(v string) *GetBillingTrendResponseBodyDataHistoryPeriodGroupByTotal {
	s.TaxAmount = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodGroupByTotal) Validate() error {
	return dara.Validate(s)
}

type GetBillingTrendResponseBodyDataHistoryPeriodResultByTime struct {
	// The statistical period.
	//
	// example:
	//
	// 20260801
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// The cost groups that actually exist in the current period.
	PeriodDetails []*GetBillingTrendResponseBodyDataHistoryPeriodResultByTimePeriodDetails `json:"periodDetails,omitempty" xml:"periodDetails,omitempty" type:"Repeated"`
	// The total costs of the current period.
	Total *GetBillingTrendResponseBodyDataHistoryPeriodResultByTimeTotal `json:"total,omitempty" xml:"total,omitempty" type:"Struct"`
}

func (s GetBillingTrendResponseBodyDataHistoryPeriodResultByTime) String() string {
	return dara.Prettify(s)
}

func (s GetBillingTrendResponseBodyDataHistoryPeriodResultByTime) GoString() string {
	return s.String()
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodResultByTime) GetPeriod() *string {
	return s.Period
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodResultByTime) GetPeriodDetails() []*GetBillingTrendResponseBodyDataHistoryPeriodResultByTimePeriodDetails {
	return s.PeriodDetails
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodResultByTime) GetTotal() *GetBillingTrendResponseBodyDataHistoryPeriodResultByTimeTotal {
	return s.Total
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodResultByTime) SetPeriod(v string) *GetBillingTrendResponseBodyDataHistoryPeriodResultByTime {
	s.Period = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodResultByTime) SetPeriodDetails(v []*GetBillingTrendResponseBodyDataHistoryPeriodResultByTimePeriodDetails) *GetBillingTrendResponseBodyDataHistoryPeriodResultByTime {
	s.PeriodDetails = v
	return s
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodResultByTime) SetTotal(v *GetBillingTrendResponseBodyDataHistoryPeriodResultByTimeTotal) *GetBillingTrendResponseBodyDataHistoryPeriodResultByTime {
	s.Total = v
	return s
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodResultByTime) Validate() error {
	if s.PeriodDetails != nil {
		for _, item := range s.PeriodDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Total != nil {
		if err := s.Total.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBillingTrendResponseBodyDataHistoryPeriodResultByTimePeriodDetails struct {
	// The amount of the current group.
	//
	// example:
	//
	// 526
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// The dimension value of the group. If the value is null, DIMENSION_FILTER_NULL_VALUE is returned.
	//
	// example:
	//
	// qwen-plus
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The display name of the group, which is affected by the locale parameter. If the value is null, a hyphen (-) is displayed.
	//
	// example:
	//
	// qwen-plus
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The percentage of the group amount.
	//
	// example:
	//
	// 0.5
	Percentage *string `json:"percentage,omitempty" xml:"percentage,omitempty"`
	// The total pretax amount.
	//
	// example:
	//
	// 100
	PretaxAmount *string `json:"pretaxAmount,omitempty" xml:"pretaxAmount,omitempty"`
	// The total tax amount.
	//
	// example:
	//
	// 426
	TaxAmount *string `json:"taxAmount,omitempty" xml:"taxAmount,omitempty"`
}

func (s GetBillingTrendResponseBodyDataHistoryPeriodResultByTimePeriodDetails) String() string {
	return dara.Prettify(s)
}

func (s GetBillingTrendResponseBodyDataHistoryPeriodResultByTimePeriodDetails) GoString() string {
	return s.String()
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodResultByTimePeriodDetails) GetAmount() *string {
	return s.Amount
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodResultByTimePeriodDetails) GetKey() *string {
	return s.Key
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodResultByTimePeriodDetails) GetName() *string {
	return s.Name
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodResultByTimePeriodDetails) GetPercentage() *string {
	return s.Percentage
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodResultByTimePeriodDetails) GetPretaxAmount() *string {
	return s.PretaxAmount
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodResultByTimePeriodDetails) GetTaxAmount() *string {
	return s.TaxAmount
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodResultByTimePeriodDetails) SetAmount(v string) *GetBillingTrendResponseBodyDataHistoryPeriodResultByTimePeriodDetails {
	s.Amount = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodResultByTimePeriodDetails) SetKey(v string) *GetBillingTrendResponseBodyDataHistoryPeriodResultByTimePeriodDetails {
	s.Key = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodResultByTimePeriodDetails) SetName(v string) *GetBillingTrendResponseBodyDataHistoryPeriodResultByTimePeriodDetails {
	s.Name = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodResultByTimePeriodDetails) SetPercentage(v string) *GetBillingTrendResponseBodyDataHistoryPeriodResultByTimePeriodDetails {
	s.Percentage = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodResultByTimePeriodDetails) SetPretaxAmount(v string) *GetBillingTrendResponseBodyDataHistoryPeriodResultByTimePeriodDetails {
	s.PretaxAmount = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodResultByTimePeriodDetails) SetTaxAmount(v string) *GetBillingTrendResponseBodyDataHistoryPeriodResultByTimePeriodDetails {
	s.TaxAmount = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodResultByTimePeriodDetails) Validate() error {
	return dara.Validate(s)
}

type GetBillingTrendResponseBodyDataHistoryPeriodResultByTimeTotal struct {
	// The total cost amount of the current period.
	//
	// example:
	//
	// 600
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// The currency of the amount.
	//
	// example:
	//
	// CNY
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty"`
	// The total pretax amount.
	//
	// example:
	//
	// 100
	PretaxAmount *string `json:"pretaxAmount,omitempty" xml:"pretaxAmount,omitempty"`
	// The total tax amount.
	//
	// example:
	//
	// 500
	TaxAmount *string `json:"taxAmount,omitempty" xml:"taxAmount,omitempty"`
}

func (s GetBillingTrendResponseBodyDataHistoryPeriodResultByTimeTotal) String() string {
	return dara.Prettify(s)
}

func (s GetBillingTrendResponseBodyDataHistoryPeriodResultByTimeTotal) GoString() string {
	return s.String()
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodResultByTimeTotal) GetAmount() *string {
	return s.Amount
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodResultByTimeTotal) GetCurrency() *string {
	return s.Currency
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodResultByTimeTotal) GetPretaxAmount() *string {
	return s.PretaxAmount
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodResultByTimeTotal) GetTaxAmount() *string {
	return s.TaxAmount
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodResultByTimeTotal) SetAmount(v string) *GetBillingTrendResponseBodyDataHistoryPeriodResultByTimeTotal {
	s.Amount = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodResultByTimeTotal) SetCurrency(v string) *GetBillingTrendResponseBodyDataHistoryPeriodResultByTimeTotal {
	s.Currency = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodResultByTimeTotal) SetPretaxAmount(v string) *GetBillingTrendResponseBodyDataHistoryPeriodResultByTimeTotal {
	s.PretaxAmount = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodResultByTimeTotal) SetTaxAmount(v string) *GetBillingTrendResponseBodyDataHistoryPeriodResultByTimeTotal {
	s.TaxAmount = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataHistoryPeriodResultByTimeTotal) Validate() error {
	return dara.Validate(s)
}

type GetBillingTrendResponseBodyDataResultByTime struct {
	// The statistical period. DAY returns yyyyMMdd, and MONTH returns yyyyMM.
	//
	// example:
	//
	// 20260801
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// The cost groups that actually exist in the current period.
	PeriodDetails []*GetBillingTrendResponseBodyDataResultByTimePeriodDetails `json:"periodDetails,omitempty" xml:"periodDetails,omitempty" type:"Repeated"`
	// The total costs of the current period.
	Total *GetBillingTrendResponseBodyDataResultByTimeTotal `json:"total,omitempty" xml:"total,omitempty" type:"Struct"`
}

func (s GetBillingTrendResponseBodyDataResultByTime) String() string {
	return dara.Prettify(s)
}

func (s GetBillingTrendResponseBodyDataResultByTime) GoString() string {
	return s.String()
}

func (s *GetBillingTrendResponseBodyDataResultByTime) GetPeriod() *string {
	return s.Period
}

func (s *GetBillingTrendResponseBodyDataResultByTime) GetPeriodDetails() []*GetBillingTrendResponseBodyDataResultByTimePeriodDetails {
	return s.PeriodDetails
}

func (s *GetBillingTrendResponseBodyDataResultByTime) GetTotal() *GetBillingTrendResponseBodyDataResultByTimeTotal {
	return s.Total
}

func (s *GetBillingTrendResponseBodyDataResultByTime) SetPeriod(v string) *GetBillingTrendResponseBodyDataResultByTime {
	s.Period = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataResultByTime) SetPeriodDetails(v []*GetBillingTrendResponseBodyDataResultByTimePeriodDetails) *GetBillingTrendResponseBodyDataResultByTime {
	s.PeriodDetails = v
	return s
}

func (s *GetBillingTrendResponseBodyDataResultByTime) SetTotal(v *GetBillingTrendResponseBodyDataResultByTimeTotal) *GetBillingTrendResponseBodyDataResultByTime {
	s.Total = v
	return s
}

func (s *GetBillingTrendResponseBodyDataResultByTime) Validate() error {
	if s.PeriodDetails != nil {
		for _, item := range s.PeriodDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Total != nil {
		if err := s.Total.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBillingTrendResponseBodyDataResultByTimePeriodDetails struct {
	// The amount of the group in the current period.
	//
	// example:
	//
	// 20
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// The dimension value of the group. For data other than the top N items, DIMENSION_GROUP_OTHERS_VALUE is used.
	//
	// example:
	//
	// qwen-plus
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The display name of the group, which is affected by the locale input parameter.
	//
	// example:
	//
	// qwen-plus
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ratio of the current group amount to the total amount of the current period.
	//
	// example:
	//
	// 0.6667
	Percentage *string `json:"percentage,omitempty" xml:"percentage,omitempty"`
	// The pretax amount of the group in the current period.
	//
	// example:
	//
	// 18.87
	PretaxAmount *string `json:"pretaxAmount,omitempty" xml:"pretaxAmount,omitempty"`
	// The tax amount of the group in the current period.
	//
	// example:
	//
	// 1.13
	TaxAmount *string `json:"taxAmount,omitempty" xml:"taxAmount,omitempty"`
}

func (s GetBillingTrendResponseBodyDataResultByTimePeriodDetails) String() string {
	return dara.Prettify(s)
}

func (s GetBillingTrendResponseBodyDataResultByTimePeriodDetails) GoString() string {
	return s.String()
}

func (s *GetBillingTrendResponseBodyDataResultByTimePeriodDetails) GetAmount() *string {
	return s.Amount
}

func (s *GetBillingTrendResponseBodyDataResultByTimePeriodDetails) GetKey() *string {
	return s.Key
}

func (s *GetBillingTrendResponseBodyDataResultByTimePeriodDetails) GetName() *string {
	return s.Name
}

func (s *GetBillingTrendResponseBodyDataResultByTimePeriodDetails) GetPercentage() *string {
	return s.Percentage
}

func (s *GetBillingTrendResponseBodyDataResultByTimePeriodDetails) GetPretaxAmount() *string {
	return s.PretaxAmount
}

func (s *GetBillingTrendResponseBodyDataResultByTimePeriodDetails) GetTaxAmount() *string {
	return s.TaxAmount
}

func (s *GetBillingTrendResponseBodyDataResultByTimePeriodDetails) SetAmount(v string) *GetBillingTrendResponseBodyDataResultByTimePeriodDetails {
	s.Amount = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataResultByTimePeriodDetails) SetKey(v string) *GetBillingTrendResponseBodyDataResultByTimePeriodDetails {
	s.Key = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataResultByTimePeriodDetails) SetName(v string) *GetBillingTrendResponseBodyDataResultByTimePeriodDetails {
	s.Name = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataResultByTimePeriodDetails) SetPercentage(v string) *GetBillingTrendResponseBodyDataResultByTimePeriodDetails {
	s.Percentage = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataResultByTimePeriodDetails) SetPretaxAmount(v string) *GetBillingTrendResponseBodyDataResultByTimePeriodDetails {
	s.PretaxAmount = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataResultByTimePeriodDetails) SetTaxAmount(v string) *GetBillingTrendResponseBodyDataResultByTimePeriodDetails {
	s.TaxAmount = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataResultByTimePeriodDetails) Validate() error {
	return dara.Validate(s)
}

type GetBillingTrendResponseBodyDataResultByTimeTotal struct {
	// The total cost amount of the current period.
	//
	// example:
	//
	// 30
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// The currency of the amount in the current period.
	//
	// example:
	//
	// CNY
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty"`
	// The pretax amount of the current period.
	//
	// example:
	//
	// 28.30
	PretaxAmount *string `json:"pretaxAmount,omitempty" xml:"pretaxAmount,omitempty"`
	// The tax amount of the current period.
	//
	// example:
	//
	// 1.70
	TaxAmount *string `json:"taxAmount,omitempty" xml:"taxAmount,omitempty"`
}

func (s GetBillingTrendResponseBodyDataResultByTimeTotal) String() string {
	return dara.Prettify(s)
}

func (s GetBillingTrendResponseBodyDataResultByTimeTotal) GoString() string {
	return s.String()
}

func (s *GetBillingTrendResponseBodyDataResultByTimeTotal) GetAmount() *string {
	return s.Amount
}

func (s *GetBillingTrendResponseBodyDataResultByTimeTotal) GetCurrency() *string {
	return s.Currency
}

func (s *GetBillingTrendResponseBodyDataResultByTimeTotal) GetPretaxAmount() *string {
	return s.PretaxAmount
}

func (s *GetBillingTrendResponseBodyDataResultByTimeTotal) GetTaxAmount() *string {
	return s.TaxAmount
}

func (s *GetBillingTrendResponseBodyDataResultByTimeTotal) SetAmount(v string) *GetBillingTrendResponseBodyDataResultByTimeTotal {
	s.Amount = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataResultByTimeTotal) SetCurrency(v string) *GetBillingTrendResponseBodyDataResultByTimeTotal {
	s.Currency = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataResultByTimeTotal) SetPretaxAmount(v string) *GetBillingTrendResponseBodyDataResultByTimeTotal {
	s.PretaxAmount = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataResultByTimeTotal) SetTaxAmount(v string) *GetBillingTrendResponseBodyDataResultByTimeTotal {
	s.TaxAmount = &v
	return s
}

func (s *GetBillingTrendResponseBodyDataResultByTimeTotal) Validate() error {
	return dara.Validate(s)
}
