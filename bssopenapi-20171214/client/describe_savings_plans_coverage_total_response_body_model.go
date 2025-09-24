// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSavingsPlansCoverageTotalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSavingsPlansCoverageTotalResponseBody
	GetCode() *string
	SetData(v *DescribeSavingsPlansCoverageTotalResponseBodyData) *DescribeSavingsPlansCoverageTotalResponseBody
	GetData() *DescribeSavingsPlansCoverageTotalResponseBodyData
	SetMessage(v string) *DescribeSavingsPlansCoverageTotalResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSavingsPlansCoverageTotalResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeSavingsPlansCoverageTotalResponseBody
	GetSuccess() *bool
}

type DescribeSavingsPlansCoverageTotalResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The return data.
	Data *DescribeSavingsPlansCoverageTotalResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s DescribeSavingsPlansCoverageTotalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansCoverageTotalResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansCoverageTotalResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSavingsPlansCoverageTotalResponseBody) GetData() *DescribeSavingsPlansCoverageTotalResponseBodyData {
	return s.Data
}

func (s *DescribeSavingsPlansCoverageTotalResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSavingsPlansCoverageTotalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSavingsPlansCoverageTotalResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeSavingsPlansCoverageTotalResponseBody) SetCode(v string) *DescribeSavingsPlansCoverageTotalResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSavingsPlansCoverageTotalResponseBody) SetData(v *DescribeSavingsPlansCoverageTotalResponseBodyData) *DescribeSavingsPlansCoverageTotalResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSavingsPlansCoverageTotalResponseBody) SetMessage(v string) *DescribeSavingsPlansCoverageTotalResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSavingsPlansCoverageTotalResponseBody) SetRequestId(v string) *DescribeSavingsPlansCoverageTotalResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSavingsPlansCoverageTotalResponseBody) SetSuccess(v bool) *DescribeSavingsPlansCoverageTotalResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSavingsPlansCoverageTotalResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSavingsPlansCoverageTotalResponseBodyData struct {
	// The coverage in different periods.
	PeriodCoverage []*DescribeSavingsPlansCoverageTotalResponseBodyDataPeriodCoverage `json:"PeriodCoverage,omitempty" xml:"PeriodCoverage,omitempty" type:"Repeated"`
	// The coverage summary.
	TotalCoverage *DescribeSavingsPlansCoverageTotalResponseBodyDataTotalCoverage `json:"TotalCoverage,omitempty" xml:"TotalCoverage,omitempty" type:"Struct"`
}

func (s DescribeSavingsPlansCoverageTotalResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansCoverageTotalResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansCoverageTotalResponseBodyData) GetPeriodCoverage() []*DescribeSavingsPlansCoverageTotalResponseBodyDataPeriodCoverage {
	return s.PeriodCoverage
}

func (s *DescribeSavingsPlansCoverageTotalResponseBodyData) GetTotalCoverage() *DescribeSavingsPlansCoverageTotalResponseBodyDataTotalCoverage {
	return s.TotalCoverage
}

func (s *DescribeSavingsPlansCoverageTotalResponseBodyData) SetPeriodCoverage(v []*DescribeSavingsPlansCoverageTotalResponseBodyDataPeriodCoverage) *DescribeSavingsPlansCoverageTotalResponseBodyData {
	s.PeriodCoverage = v
	return s
}

func (s *DescribeSavingsPlansCoverageTotalResponseBodyData) SetTotalCoverage(v *DescribeSavingsPlansCoverageTotalResponseBodyDataTotalCoverage) *DescribeSavingsPlansCoverageTotalResponseBodyData {
	s.TotalCoverage = v
	return s
}

func (s *DescribeSavingsPlansCoverageTotalResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeSavingsPlansCoverageTotalResponseBodyDataPeriodCoverage struct {
	// The coverage.
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
	// 2021071500
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
}

func (s DescribeSavingsPlansCoverageTotalResponseBodyDataPeriodCoverage) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansCoverageTotalResponseBodyDataPeriodCoverage) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansCoverageTotalResponseBodyDataPeriodCoverage) GetPercentage() *float32 {
	return s.Percentage
}

func (s *DescribeSavingsPlansCoverageTotalResponseBodyDataPeriodCoverage) GetPeriod() *string {
	return s.Period
}

func (s *DescribeSavingsPlansCoverageTotalResponseBodyDataPeriodCoverage) SetPercentage(v float32) *DescribeSavingsPlansCoverageTotalResponseBodyDataPeriodCoverage {
	s.Percentage = &v
	return s
}

func (s *DescribeSavingsPlansCoverageTotalResponseBodyDataPeriodCoverage) SetPeriod(v string) *DescribeSavingsPlansCoverageTotalResponseBodyDataPeriodCoverage {
	s.Period = &v
	return s
}

func (s *DescribeSavingsPlansCoverageTotalResponseBodyDataPeriodCoverage) Validate() error {
	return dara.Validate(s)
}

type DescribeSavingsPlansCoverageTotalResponseBodyDataTotalCoverage struct {
	// The total coverage.
	//
	// example:
	//
	// 1
	CoveragePercentage *float32 `json:"CoveragePercentage,omitempty" xml:"CoveragePercentage,omitempty"`
	// The total deducted amount.
	//
	// example:
	//
	// 100
	DeductAmount *float32 `json:"DeductAmount,omitempty" xml:"DeductAmount,omitempty"`
}

func (s DescribeSavingsPlansCoverageTotalResponseBodyDataTotalCoverage) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansCoverageTotalResponseBodyDataTotalCoverage) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansCoverageTotalResponseBodyDataTotalCoverage) GetCoveragePercentage() *float32 {
	return s.CoveragePercentage
}

func (s *DescribeSavingsPlansCoverageTotalResponseBodyDataTotalCoverage) GetDeductAmount() *float32 {
	return s.DeductAmount
}

func (s *DescribeSavingsPlansCoverageTotalResponseBodyDataTotalCoverage) SetCoveragePercentage(v float32) *DescribeSavingsPlansCoverageTotalResponseBodyDataTotalCoverage {
	s.CoveragePercentage = &v
	return s
}

func (s *DescribeSavingsPlansCoverageTotalResponseBodyDataTotalCoverage) SetDeductAmount(v float32) *DescribeSavingsPlansCoverageTotalResponseBodyDataTotalCoverage {
	s.DeductAmount = &v
	return s
}

func (s *DescribeSavingsPlansCoverageTotalResponseBodyDataTotalCoverage) Validate() error {
	return dara.Validate(s)
}
