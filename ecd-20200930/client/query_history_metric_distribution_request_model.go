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
	// example:
	//
	// 2026-04-01
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// LOAD_SCORE
	MetricName *string                                        `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Ranges     []*QueryHistoryMetricDistributionRequestRanges `json:"Ranges,omitempty" xml:"Ranges,omitempty" type:"Repeated"`
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
	// example:
	//
	// false
	IncludeMax *bool `json:"IncludeMax,omitempty" xml:"IncludeMax,omitempty"`
	// example:
	//
	// true
	IncludeMin *bool `json:"IncludeMin,omitempty" xml:"IncludeMin,omitempty"`
	// example:
	//
	// label-02\\"
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// 20
	Max *float32 `json:"Max,omitempty" xml:"Max,omitempty"`
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
