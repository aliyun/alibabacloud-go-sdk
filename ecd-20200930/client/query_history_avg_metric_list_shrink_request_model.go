// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryHistoryAvgMetricListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataDate(v string) *QueryHistoryAvgMetricListShrinkRequest
	GetDataDate() *string
	SetDesktopId(v []*string) *QueryHistoryAvgMetricListShrinkRequest
	GetDesktopId() []*string
	SetMetricName(v string) *QueryHistoryAvgMetricListShrinkRequest
	GetMetricName() *string
	SetPageNum(v int32) *QueryHistoryAvgMetricListShrinkRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryHistoryAvgMetricListShrinkRequest
	GetPageSize() *int32
	SetRangeShrink(v string) *QueryHistoryAvgMetricListShrinkRequest
	GetRangeShrink() *string
	SetResourceRegionId(v string) *QueryHistoryAvgMetricListShrinkRequest
	GetResourceRegionId() *string
	SetSortType(v string) *QueryHistoryAvgMetricListShrinkRequest
	GetSortType() *string
}

type QueryHistoryAvgMetricListShrinkRequest struct {
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
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RangeShrink *string `json:"Range,omitempty" xml:"Range,omitempty"`
	// example:
	//
	// cn-hangzhou
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// example:
	//
	// ASC
	SortType *string `json:"SortType,omitempty" xml:"SortType,omitempty"`
}

func (s QueryHistoryAvgMetricListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryHistoryAvgMetricListShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryHistoryAvgMetricListShrinkRequest) GetDataDate() *string {
	return s.DataDate
}

func (s *QueryHistoryAvgMetricListShrinkRequest) GetDesktopId() []*string {
	return s.DesktopId
}

func (s *QueryHistoryAvgMetricListShrinkRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *QueryHistoryAvgMetricListShrinkRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryHistoryAvgMetricListShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryHistoryAvgMetricListShrinkRequest) GetRangeShrink() *string {
	return s.RangeShrink
}

func (s *QueryHistoryAvgMetricListShrinkRequest) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *QueryHistoryAvgMetricListShrinkRequest) GetSortType() *string {
	return s.SortType
}

func (s *QueryHistoryAvgMetricListShrinkRequest) SetDataDate(v string) *QueryHistoryAvgMetricListShrinkRequest {
	s.DataDate = &v
	return s
}

func (s *QueryHistoryAvgMetricListShrinkRequest) SetDesktopId(v []*string) *QueryHistoryAvgMetricListShrinkRequest {
	s.DesktopId = v
	return s
}

func (s *QueryHistoryAvgMetricListShrinkRequest) SetMetricName(v string) *QueryHistoryAvgMetricListShrinkRequest {
	s.MetricName = &v
	return s
}

func (s *QueryHistoryAvgMetricListShrinkRequest) SetPageNum(v int32) *QueryHistoryAvgMetricListShrinkRequest {
	s.PageNum = &v
	return s
}

func (s *QueryHistoryAvgMetricListShrinkRequest) SetPageSize(v int32) *QueryHistoryAvgMetricListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *QueryHistoryAvgMetricListShrinkRequest) SetRangeShrink(v string) *QueryHistoryAvgMetricListShrinkRequest {
	s.RangeShrink = &v
	return s
}

func (s *QueryHistoryAvgMetricListShrinkRequest) SetResourceRegionId(v string) *QueryHistoryAvgMetricListShrinkRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *QueryHistoryAvgMetricListShrinkRequest) SetSortType(v string) *QueryHistoryAvgMetricListShrinkRequest {
	s.SortType = &v
	return s
}

func (s *QueryHistoryAvgMetricListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
