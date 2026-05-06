// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryHistoryAvgMetricListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataDate(v string) *QueryHistoryAvgMetricListRequest
	GetDataDate() *string
	SetDesktopId(v []*string) *QueryHistoryAvgMetricListRequest
	GetDesktopId() []*string
	SetMetricName(v string) *QueryHistoryAvgMetricListRequest
	GetMetricName() *string
	SetPageNum(v int32) *QueryHistoryAvgMetricListRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryHistoryAvgMetricListRequest
	GetPageSize() *int32
	SetRange(v *QueryHistoryAvgMetricListRequestRange) *QueryHistoryAvgMetricListRequest
	GetRange() *QueryHistoryAvgMetricListRequestRange
	SetResourceRegionId(v string) *QueryHistoryAvgMetricListRequest
	GetResourceRegionId() *string
	SetSortType(v string) *QueryHistoryAvgMetricListRequest
	GetSortType() *string
}

type QueryHistoryAvgMetricListRequest struct {
	// example:
	//
	// 2026-04-01
	DataDate  *string   `json:"DataDate,omitempty" xml:"DataDate,omitempty"`
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// example:
	//
	// LOAD_SCORE
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 20
	PageSize *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Range    *QueryHistoryAvgMetricListRequestRange `json:"Range,omitempty" xml:"Range,omitempty" type:"Struct"`
	// example:
	//
	// cn-hangzhou
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// example:
	//
	// ASC
	SortType *string `json:"SortType,omitempty" xml:"SortType,omitempty"`
}

func (s QueryHistoryAvgMetricListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryHistoryAvgMetricListRequest) GoString() string {
	return s.String()
}

func (s *QueryHistoryAvgMetricListRequest) GetDataDate() *string {
	return s.DataDate
}

func (s *QueryHistoryAvgMetricListRequest) GetDesktopId() []*string {
	return s.DesktopId
}

func (s *QueryHistoryAvgMetricListRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *QueryHistoryAvgMetricListRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryHistoryAvgMetricListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryHistoryAvgMetricListRequest) GetRange() *QueryHistoryAvgMetricListRequestRange {
	return s.Range
}

func (s *QueryHistoryAvgMetricListRequest) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *QueryHistoryAvgMetricListRequest) GetSortType() *string {
	return s.SortType
}

func (s *QueryHistoryAvgMetricListRequest) SetDataDate(v string) *QueryHistoryAvgMetricListRequest {
	s.DataDate = &v
	return s
}

func (s *QueryHistoryAvgMetricListRequest) SetDesktopId(v []*string) *QueryHistoryAvgMetricListRequest {
	s.DesktopId = v
	return s
}

func (s *QueryHistoryAvgMetricListRequest) SetMetricName(v string) *QueryHistoryAvgMetricListRequest {
	s.MetricName = &v
	return s
}

func (s *QueryHistoryAvgMetricListRequest) SetPageNum(v int32) *QueryHistoryAvgMetricListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryHistoryAvgMetricListRequest) SetPageSize(v int32) *QueryHistoryAvgMetricListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryHistoryAvgMetricListRequest) SetRange(v *QueryHistoryAvgMetricListRequestRange) *QueryHistoryAvgMetricListRequest {
	s.Range = v
	return s
}

func (s *QueryHistoryAvgMetricListRequest) SetResourceRegionId(v string) *QueryHistoryAvgMetricListRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *QueryHistoryAvgMetricListRequest) SetSortType(v string) *QueryHistoryAvgMetricListRequest {
	s.SortType = &v
	return s
}

func (s *QueryHistoryAvgMetricListRequest) Validate() error {
	if s.Range != nil {
		if err := s.Range.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryHistoryAvgMetricListRequestRange struct {
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

func (s QueryHistoryAvgMetricListRequestRange) String() string {
	return dara.Prettify(s)
}

func (s QueryHistoryAvgMetricListRequestRange) GoString() string {
	return s.String()
}

func (s *QueryHistoryAvgMetricListRequestRange) GetIncludeMax() *bool {
	return s.IncludeMax
}

func (s *QueryHistoryAvgMetricListRequestRange) GetIncludeMin() *bool {
	return s.IncludeMin
}

func (s *QueryHistoryAvgMetricListRequestRange) GetLabel() *string {
	return s.Label
}

func (s *QueryHistoryAvgMetricListRequestRange) GetMax() *float32 {
	return s.Max
}

func (s *QueryHistoryAvgMetricListRequestRange) GetMin() *float32 {
	return s.Min
}

func (s *QueryHistoryAvgMetricListRequestRange) SetIncludeMax(v bool) *QueryHistoryAvgMetricListRequestRange {
	s.IncludeMax = &v
	return s
}

func (s *QueryHistoryAvgMetricListRequestRange) SetIncludeMin(v bool) *QueryHistoryAvgMetricListRequestRange {
	s.IncludeMin = &v
	return s
}

func (s *QueryHistoryAvgMetricListRequestRange) SetLabel(v string) *QueryHistoryAvgMetricListRequestRange {
	s.Label = &v
	return s
}

func (s *QueryHistoryAvgMetricListRequestRange) SetMax(v float32) *QueryHistoryAvgMetricListRequestRange {
	s.Max = &v
	return s
}

func (s *QueryHistoryAvgMetricListRequestRange) SetMin(v float32) *QueryHistoryAvgMetricListRequestRange {
	s.Min = &v
	return s
}

func (s *QueryHistoryAvgMetricListRequestRange) Validate() error {
	return dara.Validate(s)
}
