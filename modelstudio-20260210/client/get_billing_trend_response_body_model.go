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
	// example:
	//
	// 200
	Code *string                          `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetBillingTrendResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 099A671E-FA21-5A36-8A73-918572DDEF53
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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
	CostTotals   *GetBillingTrendResponseBodyDataCostTotals     `json:"costTotals,omitempty" xml:"costTotals,omitempty" type:"Struct"`
	GroupByTotal []*GetBillingTrendResponseBodyDataGroupByTotal `json:"groupByTotal,omitempty" xml:"groupByTotal,omitempty" type:"Repeated"`
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
	// example:
	//
	// 100
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// example:
	//
	// CNY
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty"`
	// example:
	//
	// 94.34
	PretaxAmount *string `json:"pretaxAmount,omitempty" xml:"pretaxAmount,omitempty"`
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
	// example:
	//
	// 60
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// example:
	//
	// qwen-plus
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// qwen-plus
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 56.60
	PretaxAmount *string `json:"pretaxAmount,omitempty" xml:"pretaxAmount,omitempty"`
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

type GetBillingTrendResponseBodyDataResultByTime struct {
	// example:
	//
	// 20260801
	Period        *string                                                     `json:"period,omitempty" xml:"period,omitempty"`
	PeriodDetails []*GetBillingTrendResponseBodyDataResultByTimePeriodDetails `json:"periodDetails,omitempty" xml:"periodDetails,omitempty" type:"Repeated"`
	Total         *GetBillingTrendResponseBodyDataResultByTimeTotal           `json:"total,omitempty" xml:"total,omitempty" type:"Struct"`
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
	// example:
	//
	// 20
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// example:
	//
	// qwen-plus
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// qwen-plus
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 0.6667
	Percentage *string `json:"percentage,omitempty" xml:"percentage,omitempty"`
	// example:
	//
	// 18.87
	PretaxAmount *string `json:"pretaxAmount,omitempty" xml:"pretaxAmount,omitempty"`
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
	// example:
	//
	// 30
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// example:
	//
	// CNY
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty"`
	// example:
	//
	// 28.30
	PretaxAmount *string `json:"pretaxAmount,omitempty" xml:"pretaxAmount,omitempty"`
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
