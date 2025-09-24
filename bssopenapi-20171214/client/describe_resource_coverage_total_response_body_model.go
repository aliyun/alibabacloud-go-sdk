// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceCoverageTotalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeResourceCoverageTotalResponseBody
	GetCode() *string
	SetData(v *DescribeResourceCoverageTotalResponseBodyData) *DescribeResourceCoverageTotalResponseBody
	GetData() *DescribeResourceCoverageTotalResponseBodyData
	SetMessage(v string) *DescribeResourceCoverageTotalResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeResourceCoverageTotalResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeResourceCoverageTotalResponseBody
	GetSuccess() *bool
}

type DescribeResourceCoverageTotalResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *DescribeResourceCoverageTotalResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s DescribeResourceCoverageTotalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceCoverageTotalResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceCoverageTotalResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeResourceCoverageTotalResponseBody) GetData() *DescribeResourceCoverageTotalResponseBodyData {
	return s.Data
}

func (s *DescribeResourceCoverageTotalResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeResourceCoverageTotalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeResourceCoverageTotalResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeResourceCoverageTotalResponseBody) SetCode(v string) *DescribeResourceCoverageTotalResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeResourceCoverageTotalResponseBody) SetData(v *DescribeResourceCoverageTotalResponseBodyData) *DescribeResourceCoverageTotalResponseBody {
	s.Data = v
	return s
}

func (s *DescribeResourceCoverageTotalResponseBody) SetMessage(v string) *DescribeResourceCoverageTotalResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeResourceCoverageTotalResponseBody) SetRequestId(v string) *DescribeResourceCoverageTotalResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceCoverageTotalResponseBody) SetSuccess(v bool) *DescribeResourceCoverageTotalResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeResourceCoverageTotalResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeResourceCoverageTotalResponseBodyData struct {
	// The information about the coverage rate of deduction plans within a period.
	PeriodCoverage []*DescribeResourceCoverageTotalResponseBodyDataPeriodCoverage `json:"PeriodCoverage,omitempty" xml:"PeriodCoverage,omitempty" type:"Repeated"`
	// The information about the total coverage data of deduction plans.
	TotalCoverage *DescribeResourceCoverageTotalResponseBodyDataTotalCoverage `json:"TotalCoverage,omitempty" xml:"TotalCoverage,omitempty" type:"Struct"`
}

func (s DescribeResourceCoverageTotalResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceCoverageTotalResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeResourceCoverageTotalResponseBodyData) GetPeriodCoverage() []*DescribeResourceCoverageTotalResponseBodyDataPeriodCoverage {
	return s.PeriodCoverage
}

func (s *DescribeResourceCoverageTotalResponseBodyData) GetTotalCoverage() *DescribeResourceCoverageTotalResponseBodyDataTotalCoverage {
	return s.TotalCoverage
}

func (s *DescribeResourceCoverageTotalResponseBodyData) SetPeriodCoverage(v []*DescribeResourceCoverageTotalResponseBodyDataPeriodCoverage) *DescribeResourceCoverageTotalResponseBodyData {
	s.PeriodCoverage = v
	return s
}

func (s *DescribeResourceCoverageTotalResponseBodyData) SetTotalCoverage(v *DescribeResourceCoverageTotalResponseBodyDataTotalCoverage) *DescribeResourceCoverageTotalResponseBodyData {
	s.TotalCoverage = v
	return s
}

func (s *DescribeResourceCoverageTotalResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeResourceCoverageTotalResponseBodyDataPeriodCoverage struct {
	// The coverage rate of deduction plans within the specified period.
	//
	// example:
	//
	// 0.1
	CoveragePercentage *float32 `json:"CoveragePercentage,omitempty" xml:"CoveragePercentage,omitempty"`
	// The period.
	//
	// The value is in the format of yyyyMMddHH.
	//
	// example:
	//
	// 2020110100
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
}

func (s DescribeResourceCoverageTotalResponseBodyDataPeriodCoverage) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceCoverageTotalResponseBodyDataPeriodCoverage) GoString() string {
	return s.String()
}

func (s *DescribeResourceCoverageTotalResponseBodyDataPeriodCoverage) GetCoveragePercentage() *float32 {
	return s.CoveragePercentage
}

func (s *DescribeResourceCoverageTotalResponseBodyDataPeriodCoverage) GetPeriod() *string {
	return s.Period
}

func (s *DescribeResourceCoverageTotalResponseBodyDataPeriodCoverage) SetCoveragePercentage(v float32) *DescribeResourceCoverageTotalResponseBodyDataPeriodCoverage {
	s.CoveragePercentage = &v
	return s
}

func (s *DescribeResourceCoverageTotalResponseBodyDataPeriodCoverage) SetPeriod(v string) *DescribeResourceCoverageTotalResponseBodyDataPeriodCoverage {
	s.Period = &v
	return s
}

func (s *DescribeResourceCoverageTotalResponseBodyDataPeriodCoverage) Validate() error {
	return dara.Validate(s)
}

type DescribeResourceCoverageTotalResponseBodyDataTotalCoverage struct {
	// The unit that is used to measure the resources deducted from deduction plans.
	CapacityUnit *string `json:"CapacityUnit,omitempty" xml:"CapacityUnit,omitempty"`
	// The total coverage rate of deduction plans.
	//
	// example:
	//
	// 1
	CoveragePercentage *float32 `json:"CoveragePercentage,omitempty" xml:"CoveragePercentage,omitempty"`
	// The total amount of the resources deducted from deduction plans.
	//
	// example:
	//
	// 1
	DeductQuantity *float32 `json:"DeductQuantity,omitempty" xml:"DeductQuantity,omitempty"`
	// The total amount of resources consumed.
	//
	// example:
	//
	// 1
	TotalQuantity *float32 `json:"TotalQuantity,omitempty" xml:"TotalQuantity,omitempty"`
}

func (s DescribeResourceCoverageTotalResponseBodyDataTotalCoverage) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceCoverageTotalResponseBodyDataTotalCoverage) GoString() string {
	return s.String()
}

func (s *DescribeResourceCoverageTotalResponseBodyDataTotalCoverage) GetCapacityUnit() *string {
	return s.CapacityUnit
}

func (s *DescribeResourceCoverageTotalResponseBodyDataTotalCoverage) GetCoveragePercentage() *float32 {
	return s.CoveragePercentage
}

func (s *DescribeResourceCoverageTotalResponseBodyDataTotalCoverage) GetDeductQuantity() *float32 {
	return s.DeductQuantity
}

func (s *DescribeResourceCoverageTotalResponseBodyDataTotalCoverage) GetTotalQuantity() *float32 {
	return s.TotalQuantity
}

func (s *DescribeResourceCoverageTotalResponseBodyDataTotalCoverage) SetCapacityUnit(v string) *DescribeResourceCoverageTotalResponseBodyDataTotalCoverage {
	s.CapacityUnit = &v
	return s
}

func (s *DescribeResourceCoverageTotalResponseBodyDataTotalCoverage) SetCoveragePercentage(v float32) *DescribeResourceCoverageTotalResponseBodyDataTotalCoverage {
	s.CoveragePercentage = &v
	return s
}

func (s *DescribeResourceCoverageTotalResponseBodyDataTotalCoverage) SetDeductQuantity(v float32) *DescribeResourceCoverageTotalResponseBodyDataTotalCoverage {
	s.DeductQuantity = &v
	return s
}

func (s *DescribeResourceCoverageTotalResponseBodyDataTotalCoverage) SetTotalQuantity(v float32) *DescribeResourceCoverageTotalResponseBodyDataTotalCoverage {
	s.TotalQuantity = &v
	return s
}

func (s *DescribeResourceCoverageTotalResponseBodyDataTotalCoverage) Validate() error {
	return dara.Validate(s)
}
