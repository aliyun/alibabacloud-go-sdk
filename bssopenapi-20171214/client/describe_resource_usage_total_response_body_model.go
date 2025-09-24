// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceUsageTotalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeResourceUsageTotalResponseBody
	GetCode() *string
	SetData(v *DescribeResourceUsageTotalResponseBodyData) *DescribeResourceUsageTotalResponseBody
	GetData() *DescribeResourceUsageTotalResponseBodyData
	SetMessage(v string) *DescribeResourceUsageTotalResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeResourceUsageTotalResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeResourceUsageTotalResponseBody
	GetSuccess() *bool
}

type DescribeResourceUsageTotalResponseBody struct {
	// The response code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The queried data.
	Data *DescribeResourceUsageTotalResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message. If the request was successful, a success message is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 364C7C81-5E5E-51A0-B738-1969D2671B05
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeResourceUsageTotalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceUsageTotalResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceUsageTotalResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeResourceUsageTotalResponseBody) GetData() *DescribeResourceUsageTotalResponseBodyData {
	return s.Data
}

func (s *DescribeResourceUsageTotalResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeResourceUsageTotalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeResourceUsageTotalResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeResourceUsageTotalResponseBody) SetCode(v string) *DescribeResourceUsageTotalResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeResourceUsageTotalResponseBody) SetData(v *DescribeResourceUsageTotalResponseBodyData) *DescribeResourceUsageTotalResponseBody {
	s.Data = v
	return s
}

func (s *DescribeResourceUsageTotalResponseBody) SetMessage(v string) *DescribeResourceUsageTotalResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeResourceUsageTotalResponseBody) SetRequestId(v string) *DescribeResourceUsageTotalResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceUsageTotalResponseBody) SetSuccess(v bool) *DescribeResourceUsageTotalResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeResourceUsageTotalResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeResourceUsageTotalResponseBodyData struct {
	// The usage of resource plans in the specified period.
	PeriodCoverage []*DescribeResourceUsageTotalResponseBodyDataPeriodCoverage `json:"PeriodCoverage,omitempty" xml:"PeriodCoverage,omitempty" type:"Repeated"`
	// The total usage of resource plans.
	TotalUsage *DescribeResourceUsageTotalResponseBodyDataTotalUsage `json:"TotalUsage,omitempty" xml:"TotalUsage,omitempty" type:"Struct"`
}

func (s DescribeResourceUsageTotalResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceUsageTotalResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeResourceUsageTotalResponseBodyData) GetPeriodCoverage() []*DescribeResourceUsageTotalResponseBodyDataPeriodCoverage {
	return s.PeriodCoverage
}

func (s *DescribeResourceUsageTotalResponseBodyData) GetTotalUsage() *DescribeResourceUsageTotalResponseBodyDataTotalUsage {
	return s.TotalUsage
}

func (s *DescribeResourceUsageTotalResponseBodyData) SetPeriodCoverage(v []*DescribeResourceUsageTotalResponseBodyDataPeriodCoverage) *DescribeResourceUsageTotalResponseBodyData {
	s.PeriodCoverage = v
	return s
}

func (s *DescribeResourceUsageTotalResponseBodyData) SetTotalUsage(v *DescribeResourceUsageTotalResponseBodyDataTotalUsage) *DescribeResourceUsageTotalResponseBodyData {
	s.TotalUsage = v
	return s
}

func (s *DescribeResourceUsageTotalResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeResourceUsageTotalResponseBodyDataPeriodCoverage struct {
	// The period.
	//
	// example:
	//
	// 2021071500
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The usage of resource plans.
	//
	// example:
	//
	// 0.1
	UsagePercentage *float32 `json:"UsagePercentage,omitempty" xml:"UsagePercentage,omitempty"`
}

func (s DescribeResourceUsageTotalResponseBodyDataPeriodCoverage) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceUsageTotalResponseBodyDataPeriodCoverage) GoString() string {
	return s.String()
}

func (s *DescribeResourceUsageTotalResponseBodyDataPeriodCoverage) GetPeriod() *string {
	return s.Period
}

func (s *DescribeResourceUsageTotalResponseBodyDataPeriodCoverage) GetUsagePercentage() *float32 {
	return s.UsagePercentage
}

func (s *DescribeResourceUsageTotalResponseBodyDataPeriodCoverage) SetPeriod(v string) *DescribeResourceUsageTotalResponseBodyDataPeriodCoverage {
	s.Period = &v
	return s
}

func (s *DescribeResourceUsageTotalResponseBodyDataPeriodCoverage) SetUsagePercentage(v float32) *DescribeResourceUsageTotalResponseBodyDataPeriodCoverage {
	s.UsagePercentage = &v
	return s
}

func (s *DescribeResourceUsageTotalResponseBodyDataPeriodCoverage) Validate() error {
	return dara.Validate(s)
}

type DescribeResourceUsageTotalResponseBodyDataTotalUsage struct {
	// The total costs of pay-as-you-go instances.
	//
	// example:
	//
	// 200
	PostpaidCost *float32 `json:"PostpaidCost,omitempty" xml:"PostpaidCost,omitempty"`
	// The total potential savings.
	//
	// example:
	//
	// 100
	PotentialSavedCost *float32 `json:"PotentialSavedCost,omitempty" xml:"PotentialSavedCost,omitempty"`
	// The fee of purchased resource plans.
	//
	// example:
	//
	// 10
	ReservationCost *float32 `json:"ReservationCost,omitempty" xml:"ReservationCost,omitempty"`
	// The total savings.
	//
	// example:
	//
	// 100
	SavedCost *float32 `json:"SavedCost,omitempty" xml:"SavedCost,omitempty"`
	// The total usage of resource plans.
	//
	// example:
	//
	// 1
	UsagePercentage *float32 `json:"UsagePercentage,omitempty" xml:"UsagePercentage,omitempty"`
}

func (s DescribeResourceUsageTotalResponseBodyDataTotalUsage) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceUsageTotalResponseBodyDataTotalUsage) GoString() string {
	return s.String()
}

func (s *DescribeResourceUsageTotalResponseBodyDataTotalUsage) GetPostpaidCost() *float32 {
	return s.PostpaidCost
}

func (s *DescribeResourceUsageTotalResponseBodyDataTotalUsage) GetPotentialSavedCost() *float32 {
	return s.PotentialSavedCost
}

func (s *DescribeResourceUsageTotalResponseBodyDataTotalUsage) GetReservationCost() *float32 {
	return s.ReservationCost
}

func (s *DescribeResourceUsageTotalResponseBodyDataTotalUsage) GetSavedCost() *float32 {
	return s.SavedCost
}

func (s *DescribeResourceUsageTotalResponseBodyDataTotalUsage) GetUsagePercentage() *float32 {
	return s.UsagePercentage
}

func (s *DescribeResourceUsageTotalResponseBodyDataTotalUsage) SetPostpaidCost(v float32) *DescribeResourceUsageTotalResponseBodyDataTotalUsage {
	s.PostpaidCost = &v
	return s
}

func (s *DescribeResourceUsageTotalResponseBodyDataTotalUsage) SetPotentialSavedCost(v float32) *DescribeResourceUsageTotalResponseBodyDataTotalUsage {
	s.PotentialSavedCost = &v
	return s
}

func (s *DescribeResourceUsageTotalResponseBodyDataTotalUsage) SetReservationCost(v float32) *DescribeResourceUsageTotalResponseBodyDataTotalUsage {
	s.ReservationCost = &v
	return s
}

func (s *DescribeResourceUsageTotalResponseBodyDataTotalUsage) SetSavedCost(v float32) *DescribeResourceUsageTotalResponseBodyDataTotalUsage {
	s.SavedCost = &v
	return s
}

func (s *DescribeResourceUsageTotalResponseBodyDataTotalUsage) SetUsagePercentage(v float32) *DescribeResourceUsageTotalResponseBodyDataTotalUsage {
	s.UsagePercentage = &v
	return s
}

func (s *DescribeResourceUsageTotalResponseBodyDataTotalUsage) Validate() error {
	return dara.Validate(s)
}
