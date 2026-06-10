// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryHistoryMetricDistributionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v string) *QueryHistoryMetricDistributionRequest
	GetEndDate() *string
	SetMetricName(v string) *QueryHistoryMetricDistributionRequest
	GetMetricName() *string
	SetRanges(v []*QueryHistoryMetricDistributionRequestRanges) *QueryHistoryMetricDistributionRequest
	GetRanges() []*QueryHistoryMetricDistributionRequestRanges
	SetStartDate(v string) *QueryHistoryMetricDistributionRequest
	GetStartDate() *string
}

type QueryHistoryMetricDistributionRequest struct {
	// The end date of the query period. The date must be in the `YYYY-MM-DD` format. The default value is T-1.
	//
	// example:
	//
	// 2026-04-01
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The name of the metric to query.
	//
	// example:
	//
	// LOAD_SCORE
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// An array of custom value ranges.
	Ranges []*QueryHistoryMetricDistributionRequestRanges `json:"Ranges,omitempty" xml:"Ranges,omitempty" type:"Repeated"`
	// The start date of the query period. The date must be in the `YYYY-MM-DD` format. The default value is T-1.
	//
	// example:
	//
	// 2026-04-01
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s QueryHistoryMetricDistributionRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryHistoryMetricDistributionRequest) GoString() string {
	return s.String()
}

func (s *QueryHistoryMetricDistributionRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *QueryHistoryMetricDistributionRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *QueryHistoryMetricDistributionRequest) GetRanges() []*QueryHistoryMetricDistributionRequestRanges {
	return s.Ranges
}

func (s *QueryHistoryMetricDistributionRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *QueryHistoryMetricDistributionRequest) SetEndDate(v string) *QueryHistoryMetricDistributionRequest {
	s.EndDate = &v
	return s
}

func (s *QueryHistoryMetricDistributionRequest) SetMetricName(v string) *QueryHistoryMetricDistributionRequest {
	s.MetricName = &v
	return s
}

func (s *QueryHistoryMetricDistributionRequest) SetRanges(v []*QueryHistoryMetricDistributionRequestRanges) *QueryHistoryMetricDistributionRequest {
	s.Ranges = v
	return s
}

func (s *QueryHistoryMetricDistributionRequest) SetStartDate(v string) *QueryHistoryMetricDistributionRequest {
	s.StartDate = &v
	return s
}

func (s *QueryHistoryMetricDistributionRequest) Validate() error {
	if s.Ranges != nil {
		for _, item := range s.Ranges {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryHistoryMetricDistributionRequestRanges struct {
	// Specifies whether to include the maximum value in the range. The default value is `false`.
	//
	// example:
	//
	// false
	IncludeMax *bool `json:"IncludeMax,omitempty" xml:"IncludeMax,omitempty"`
	// Specifies whether to include the minimum value in the range. The default value is `true`.
	//
	// example:
	//
	// true
	IncludeMin *bool `json:"IncludeMin,omitempty" xml:"IncludeMin,omitempty"`
	// The label for the value range. This label is returned in the response.
	//
	// example:
	//
	// label-02\\"
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The maximum value of the value range.
	//
	// example:
	//
	// 20
	Max *float32 `json:"Max,omitempty" xml:"Max,omitempty"`
	// The minimum value of the value range.
	//
	// example:
	//
	// 0
	Min *float32 `json:"Min,omitempty" xml:"Min,omitempty"`
}

func (s QueryHistoryMetricDistributionRequestRanges) String() string {
	return dara.Prettify(s)
}

func (s QueryHistoryMetricDistributionRequestRanges) GoString() string {
	return s.String()
}

func (s *QueryHistoryMetricDistributionRequestRanges) GetIncludeMax() *bool {
	return s.IncludeMax
}

func (s *QueryHistoryMetricDistributionRequestRanges) GetIncludeMin() *bool {
	return s.IncludeMin
}

func (s *QueryHistoryMetricDistributionRequestRanges) GetLabel() *string {
	return s.Label
}

func (s *QueryHistoryMetricDistributionRequestRanges) GetMax() *float32 {
	return s.Max
}

func (s *QueryHistoryMetricDistributionRequestRanges) GetMin() *float32 {
	return s.Min
}

func (s *QueryHistoryMetricDistributionRequestRanges) SetIncludeMax(v bool) *QueryHistoryMetricDistributionRequestRanges {
	s.IncludeMax = &v
	return s
}

func (s *QueryHistoryMetricDistributionRequestRanges) SetIncludeMin(v bool) *QueryHistoryMetricDistributionRequestRanges {
	s.IncludeMin = &v
	return s
}

func (s *QueryHistoryMetricDistributionRequestRanges) SetLabel(v string) *QueryHistoryMetricDistributionRequestRanges {
	s.Label = &v
	return s
}

func (s *QueryHistoryMetricDistributionRequestRanges) SetMax(v float32) *QueryHistoryMetricDistributionRequestRanges {
	s.Max = &v
	return s
}

func (s *QueryHistoryMetricDistributionRequestRanges) SetMin(v float32) *QueryHistoryMetricDistributionRequestRanges {
	s.Min = &v
	return s
}

func (s *QueryHistoryMetricDistributionRequestRanges) Validate() error {
	return dara.Validate(s)
}
