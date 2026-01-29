// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSavingsPlansUsageTotalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSavingsPlansUsageTotalResponseBody
	GetCode() *string
	SetData(v *DescribeSavingsPlansUsageTotalResponseBodyData) *DescribeSavingsPlansUsageTotalResponseBody
	GetData() *DescribeSavingsPlansUsageTotalResponseBodyData
	SetMessage(v string) *DescribeSavingsPlansUsageTotalResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSavingsPlansUsageTotalResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeSavingsPlansUsageTotalResponseBody
	GetSuccess() *bool
}

type DescribeSavingsPlansUsageTotalResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The return data.
	Data *DescribeSavingsPlansUsageTotalResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 79EE7556-0CFD-44EB-9CD6-B3B526E3A85F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSavingsPlansUsageTotalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansUsageTotalResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansUsageTotalResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSavingsPlansUsageTotalResponseBody) GetData() *DescribeSavingsPlansUsageTotalResponseBodyData {
	return s.Data
}

func (s *DescribeSavingsPlansUsageTotalResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSavingsPlansUsageTotalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSavingsPlansUsageTotalResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeSavingsPlansUsageTotalResponseBody) SetCode(v string) *DescribeSavingsPlansUsageTotalResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSavingsPlansUsageTotalResponseBody) SetData(v *DescribeSavingsPlansUsageTotalResponseBodyData) *DescribeSavingsPlansUsageTotalResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSavingsPlansUsageTotalResponseBody) SetMessage(v string) *DescribeSavingsPlansUsageTotalResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSavingsPlansUsageTotalResponseBody) SetRequestId(v string) *DescribeSavingsPlansUsageTotalResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSavingsPlansUsageTotalResponseBody) SetSuccess(v bool) *DescribeSavingsPlansUsageTotalResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSavingsPlansUsageTotalResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSavingsPlansUsageTotalResponseBodyData struct {
	// The usage in different periods.
	PeriodCoverage []*DescribeSavingsPlansUsageTotalResponseBodyDataPeriodCoverage `json:"PeriodCoverage,omitempty" xml:"PeriodCoverage,omitempty" type:"Repeated"`
	// The usage summary.
	TotalUsage *DescribeSavingsPlansUsageTotalResponseBodyDataTotalUsage `json:"TotalUsage,omitempty" xml:"TotalUsage,omitempty" type:"Struct"`
}

func (s DescribeSavingsPlansUsageTotalResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansUsageTotalResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansUsageTotalResponseBodyData) GetPeriodCoverage() []*DescribeSavingsPlansUsageTotalResponseBodyDataPeriodCoverage {
	return s.PeriodCoverage
}

func (s *DescribeSavingsPlansUsageTotalResponseBodyData) GetTotalUsage() *DescribeSavingsPlansUsageTotalResponseBodyDataTotalUsage {
	return s.TotalUsage
}

func (s *DescribeSavingsPlansUsageTotalResponseBodyData) SetPeriodCoverage(v []*DescribeSavingsPlansUsageTotalResponseBodyDataPeriodCoverage) *DescribeSavingsPlansUsageTotalResponseBodyData {
	s.PeriodCoverage = v
	return s
}

func (s *DescribeSavingsPlansUsageTotalResponseBodyData) SetTotalUsage(v *DescribeSavingsPlansUsageTotalResponseBodyDataTotalUsage) *DescribeSavingsPlansUsageTotalResponseBodyData {
	s.TotalUsage = v
	return s
}

func (s *DescribeSavingsPlansUsageTotalResponseBodyData) Validate() error {
	if s.PeriodCoverage != nil {
		for _, item := range s.PeriodCoverage {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TotalUsage != nil {
		if err := s.TotalUsage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSavingsPlansUsageTotalResponseBodyDataPeriodCoverage struct {
	// The usage.
	//
	// example:
	//
	// 1
	Percentage *float32 `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	// The period.
	//
	// The value is in the format of yyyyMMddHH.
	//
	// example:
	//
	// 2021041500
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
}

func (s DescribeSavingsPlansUsageTotalResponseBodyDataPeriodCoverage) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansUsageTotalResponseBodyDataPeriodCoverage) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansUsageTotalResponseBodyDataPeriodCoverage) GetPercentage() *float32 {
	return s.Percentage
}

func (s *DescribeSavingsPlansUsageTotalResponseBodyDataPeriodCoverage) GetPeriod() *string {
	return s.Period
}

func (s *DescribeSavingsPlansUsageTotalResponseBodyDataPeriodCoverage) SetPercentage(v float32) *DescribeSavingsPlansUsageTotalResponseBodyDataPeriodCoverage {
	s.Percentage = &v
	return s
}

func (s *DescribeSavingsPlansUsageTotalResponseBodyDataPeriodCoverage) SetPeriod(v string) *DescribeSavingsPlansUsageTotalResponseBodyDataPeriodCoverage {
	s.Period = &v
	return s
}

func (s *DescribeSavingsPlansUsageTotalResponseBodyDataPeriodCoverage) Validate() error {
	return dara.Validate(s)
}

type DescribeSavingsPlansUsageTotalResponseBodyDataTotalUsage struct {
	// The total amount of the savings plan.
	//
	// example:
	//
	// 100
	PoolValue *float32 `json:"PoolValue,omitempty" xml:"PoolValue,omitempty"`
	// The pay-as-you-go cost.
	//
	// example:
	//
	// 200
	PostpaidCost *float32 `json:"PostpaidCost,omitempty" xml:"PostpaidCost,omitempty"`
	// The amount that is saved.
	//
	// example:
	//
	// 100
	SavedCost *float32 `json:"SavedCost,omitempty" xml:"SavedCost,omitempty"`
	// The total usage.
	//
	// example:
	//
	// 1
	UsagePercentage *float32 `json:"UsagePercentage,omitempty" xml:"UsagePercentage,omitempty"`
}

func (s DescribeSavingsPlansUsageTotalResponseBodyDataTotalUsage) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansUsageTotalResponseBodyDataTotalUsage) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansUsageTotalResponseBodyDataTotalUsage) GetPoolValue() *float32 {
	return s.PoolValue
}

func (s *DescribeSavingsPlansUsageTotalResponseBodyDataTotalUsage) GetPostpaidCost() *float32 {
	return s.PostpaidCost
}

func (s *DescribeSavingsPlansUsageTotalResponseBodyDataTotalUsage) GetSavedCost() *float32 {
	return s.SavedCost
}

func (s *DescribeSavingsPlansUsageTotalResponseBodyDataTotalUsage) GetUsagePercentage() *float32 {
	return s.UsagePercentage
}

func (s *DescribeSavingsPlansUsageTotalResponseBodyDataTotalUsage) SetPoolValue(v float32) *DescribeSavingsPlansUsageTotalResponseBodyDataTotalUsage {
	s.PoolValue = &v
	return s
}

func (s *DescribeSavingsPlansUsageTotalResponseBodyDataTotalUsage) SetPostpaidCost(v float32) *DescribeSavingsPlansUsageTotalResponseBodyDataTotalUsage {
	s.PostpaidCost = &v
	return s
}

func (s *DescribeSavingsPlansUsageTotalResponseBodyDataTotalUsage) SetSavedCost(v float32) *DescribeSavingsPlansUsageTotalResponseBodyDataTotalUsage {
	s.SavedCost = &v
	return s
}

func (s *DescribeSavingsPlansUsageTotalResponseBodyDataTotalUsage) SetUsagePercentage(v float32) *DescribeSavingsPlansUsageTotalResponseBodyDataTotalUsage {
	s.UsagePercentage = &v
	return s
}

func (s *DescribeSavingsPlansUsageTotalResponseBodyDataTotalUsage) Validate() error {
	return dara.Validate(s)
}
