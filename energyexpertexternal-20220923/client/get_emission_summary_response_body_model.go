// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEmissionSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetEmissionSummaryResponseBodyData) *GetEmissionSummaryResponseBody
	GetData() *GetEmissionSummaryResponseBodyData
	SetRequestId(v string) *GetEmissionSummaryResponseBody
	GetRequestId() *string
}

type GetEmissionSummaryResponseBody struct {
	// Details of summarized data
	Data *GetEmissionSummaryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetEmissionSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEmissionSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetEmissionSummaryResponseBody) GetData() *GetEmissionSummaryResponseBodyData {
	return s.Data
}

func (s *GetEmissionSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEmissionSummaryResponseBody) SetData(v *GetEmissionSummaryResponseBodyData) *GetEmissionSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetEmissionSummaryResponseBody) SetRequestId(v string) *GetEmissionSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEmissionSummaryResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetEmissionSummaryResponseBodyData struct {
	// Percentage of scheduled.
	//
	// example:
	//
	// 2.7657
	ActualEmissionRatio *float64 `json:"actualEmissionRatio,omitempty" xml:"actualEmissionRatio,omitempty"`
	// Carbon emissions reduction.
	//
	// example:
	//
	// 0.0
	CarbonSaveConversion *float64 `json:"carbonSaveConversion,omitempty" xml:"carbonSaveConversion,omitempty"`
	// Statistical indicators for this cycle.
	//
	// example:
	//
	// 276.4
	CurrentPeriodCarbonEmissionData *float64 `json:"currentPeriodCarbonEmissionData,omitempty" xml:"currentPeriodCarbonEmissionData,omitempty"`
	// Whether to show the weighting ratio carbon emission.
	//
	// example:
	//
	// true
	IsWeighting *bool `json:"isWeighting,omitempty" xml:"isWeighting,omitempty"`
	// Last period statistical indicators.
	//
	// example:
	//
	// 9.40100
	LastPeriodCarbonEmissionData *float64 `json:"lastPeriodCarbonEmissionData,omitempty" xml:"lastPeriodCarbonEmissionData,omitempty"`
	// Calculation of carbon emissions based on shareholding ratio in the last cycle.
	//
	// example:
	//
	// 8.4609
	LastPeriodWeightingCarbonEmissionData *float64 `json:"lastPeriodWeightingCarbonEmissionData,omitempty" xml:"lastPeriodWeightingCarbonEmissionData,omitempty"`
	// Year-on-year.
	//
	// example:
	//
	// 28.410
	Ratio *float64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// Total carbon emissions.
	//
	// example:
	//
	// 276.46
	TotalCarbonEmissionData *float64 `json:"totalCarbonEmissionData,omitempty" xml:"totalCarbonEmissionData,omitempty"`
	// Calculate carbon emissions by share ratio
	//
	// example:
	//
	// 248.81400
	WeightingCarbonEmissionData *float64 `json:"weightingCarbonEmissionData,omitempty" xml:"weightingCarbonEmissionData,omitempty"`
	// Year-on-year of weighting ratio carbon emissions.
	//
	// example:
	//
	// 28.4102
	WeightingRatio *float64 `json:"weightingRatio,omitempty" xml:"weightingRatio,omitempty"`
}

func (s GetEmissionSummaryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetEmissionSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEmissionSummaryResponseBodyData) GetActualEmissionRatio() *float64 {
	return s.ActualEmissionRatio
}

func (s *GetEmissionSummaryResponseBodyData) GetCarbonSaveConversion() *float64 {
	return s.CarbonSaveConversion
}

func (s *GetEmissionSummaryResponseBodyData) GetCurrentPeriodCarbonEmissionData() *float64 {
	return s.CurrentPeriodCarbonEmissionData
}

func (s *GetEmissionSummaryResponseBodyData) GetIsWeighting() *bool {
	return s.IsWeighting
}

func (s *GetEmissionSummaryResponseBodyData) GetLastPeriodCarbonEmissionData() *float64 {
	return s.LastPeriodCarbonEmissionData
}

func (s *GetEmissionSummaryResponseBodyData) GetLastPeriodWeightingCarbonEmissionData() *float64 {
	return s.LastPeriodWeightingCarbonEmissionData
}

func (s *GetEmissionSummaryResponseBodyData) GetRatio() *float64 {
	return s.Ratio
}

func (s *GetEmissionSummaryResponseBodyData) GetTotalCarbonEmissionData() *float64 {
	return s.TotalCarbonEmissionData
}

func (s *GetEmissionSummaryResponseBodyData) GetWeightingCarbonEmissionData() *float64 {
	return s.WeightingCarbonEmissionData
}

func (s *GetEmissionSummaryResponseBodyData) GetWeightingRatio() *float64 {
	return s.WeightingRatio
}

func (s *GetEmissionSummaryResponseBodyData) SetActualEmissionRatio(v float64) *GetEmissionSummaryResponseBodyData {
	s.ActualEmissionRatio = &v
	return s
}

func (s *GetEmissionSummaryResponseBodyData) SetCarbonSaveConversion(v float64) *GetEmissionSummaryResponseBodyData {
	s.CarbonSaveConversion = &v
	return s
}

func (s *GetEmissionSummaryResponseBodyData) SetCurrentPeriodCarbonEmissionData(v float64) *GetEmissionSummaryResponseBodyData {
	s.CurrentPeriodCarbonEmissionData = &v
	return s
}

func (s *GetEmissionSummaryResponseBodyData) SetIsWeighting(v bool) *GetEmissionSummaryResponseBodyData {
	s.IsWeighting = &v
	return s
}

func (s *GetEmissionSummaryResponseBodyData) SetLastPeriodCarbonEmissionData(v float64) *GetEmissionSummaryResponseBodyData {
	s.LastPeriodCarbonEmissionData = &v
	return s
}

func (s *GetEmissionSummaryResponseBodyData) SetLastPeriodWeightingCarbonEmissionData(v float64) *GetEmissionSummaryResponseBodyData {
	s.LastPeriodWeightingCarbonEmissionData = &v
	return s
}

func (s *GetEmissionSummaryResponseBodyData) SetRatio(v float64) *GetEmissionSummaryResponseBodyData {
	s.Ratio = &v
	return s
}

func (s *GetEmissionSummaryResponseBodyData) SetTotalCarbonEmissionData(v float64) *GetEmissionSummaryResponseBodyData {
	s.TotalCarbonEmissionData = &v
	return s
}

func (s *GetEmissionSummaryResponseBodyData) SetWeightingCarbonEmissionData(v float64) *GetEmissionSummaryResponseBodyData {
	s.WeightingCarbonEmissionData = &v
	return s
}

func (s *GetEmissionSummaryResponseBodyData) SetWeightingRatio(v float64) *GetEmissionSummaryResponseBodyData {
	s.WeightingRatio = &v
	return s
}

func (s *GetEmissionSummaryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
