// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMetricByPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *QueryMetricByPageRequest
	GetCurrentPage() *int32
	SetCustomFilters(v []*string) *QueryMetricByPageRequest
	GetCustomFilters() []*string
	SetDimensions(v []*string) *QueryMetricByPageRequest
	GetDimensions() []*string
	SetEndTime(v int64) *QueryMetricByPageRequest
	GetEndTime() *int64
	SetFilters(v []*QueryMetricByPageRequestFilters) *QueryMetricByPageRequest
	GetFilters() []*QueryMetricByPageRequestFilters
	SetIntervalInSec(v int32) *QueryMetricByPageRequest
	GetIntervalInSec() *int32
	SetMeasures(v []*string) *QueryMetricByPageRequest
	GetMeasures() []*string
	SetMetric(v string) *QueryMetricByPageRequest
	GetMetric() *string
	SetOrder(v string) *QueryMetricByPageRequest
	GetOrder() *string
	SetOrderBy(v string) *QueryMetricByPageRequest
	GetOrderBy() *string
	SetPageSize(v int32) *QueryMetricByPageRequest
	GetPageSize() *int32
	SetStartTime(v int64) *QueryMetricByPageRequest
	GetStartTime() *int64
}

type QueryMetricByPageRequest struct {
	CurrentPage   *int32    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	CustomFilters []*string `json:"CustomFilters,omitempty" xml:"CustomFilters,omitempty" type:"Repeated"`
	Dimensions    []*string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// This parameter is required.
	EndTime       *int64                             `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Filters       []*QueryMetricByPageRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	IntervalInSec *int32                             `json:"IntervalInSec,omitempty" xml:"IntervalInSec,omitempty"`
	// This parameter is required.
	Measures []*string `json:"Measures,omitempty" xml:"Measures,omitempty" type:"Repeated"`
	// This parameter is required.
	Metric   *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	Order    *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OrderBy  *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryMetricByPageRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMetricByPageRequest) GoString() string {
	return s.String()
}

func (s *QueryMetricByPageRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryMetricByPageRequest) GetCustomFilters() []*string {
	return s.CustomFilters
}

func (s *QueryMetricByPageRequest) GetDimensions() []*string {
	return s.Dimensions
}

func (s *QueryMetricByPageRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryMetricByPageRequest) GetFilters() []*QueryMetricByPageRequestFilters {
	return s.Filters
}

func (s *QueryMetricByPageRequest) GetIntervalInSec() *int32 {
	return s.IntervalInSec
}

func (s *QueryMetricByPageRequest) GetMeasures() []*string {
	return s.Measures
}

func (s *QueryMetricByPageRequest) GetMetric() *string {
	return s.Metric
}

func (s *QueryMetricByPageRequest) GetOrder() *string {
	return s.Order
}

func (s *QueryMetricByPageRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *QueryMetricByPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryMetricByPageRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *QueryMetricByPageRequest) SetCurrentPage(v int32) *QueryMetricByPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryMetricByPageRequest) SetCustomFilters(v []*string) *QueryMetricByPageRequest {
	s.CustomFilters = v
	return s
}

func (s *QueryMetricByPageRequest) SetDimensions(v []*string) *QueryMetricByPageRequest {
	s.Dimensions = v
	return s
}

func (s *QueryMetricByPageRequest) SetEndTime(v int64) *QueryMetricByPageRequest {
	s.EndTime = &v
	return s
}

func (s *QueryMetricByPageRequest) SetFilters(v []*QueryMetricByPageRequestFilters) *QueryMetricByPageRequest {
	s.Filters = v
	return s
}

func (s *QueryMetricByPageRequest) SetIntervalInSec(v int32) *QueryMetricByPageRequest {
	s.IntervalInSec = &v
	return s
}

func (s *QueryMetricByPageRequest) SetMeasures(v []*string) *QueryMetricByPageRequest {
	s.Measures = v
	return s
}

func (s *QueryMetricByPageRequest) SetMetric(v string) *QueryMetricByPageRequest {
	s.Metric = &v
	return s
}

func (s *QueryMetricByPageRequest) SetOrder(v string) *QueryMetricByPageRequest {
	s.Order = &v
	return s
}

func (s *QueryMetricByPageRequest) SetOrderBy(v string) *QueryMetricByPageRequest {
	s.OrderBy = &v
	return s
}

func (s *QueryMetricByPageRequest) SetPageSize(v int32) *QueryMetricByPageRequest {
	s.PageSize = &v
	return s
}

func (s *QueryMetricByPageRequest) SetStartTime(v int64) *QueryMetricByPageRequest {
	s.StartTime = &v
	return s
}

func (s *QueryMetricByPageRequest) Validate() error {
	if s.Filters != nil {
		for _, item := range s.Filters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMetricByPageRequestFilters struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryMetricByPageRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s QueryMetricByPageRequestFilters) GoString() string {
	return s.String()
}

func (s *QueryMetricByPageRequestFilters) GetKey() *string {
	return s.Key
}

func (s *QueryMetricByPageRequestFilters) GetValue() *string {
	return s.Value
}

func (s *QueryMetricByPageRequestFilters) SetKey(v string) *QueryMetricByPageRequestFilters {
	s.Key = &v
	return s
}

func (s *QueryMetricByPageRequestFilters) SetValue(v string) *QueryMetricByPageRequestFilters {
	s.Value = &v
	return s
}

func (s *QueryMetricByPageRequestFilters) Validate() error {
	return dara.Validate(s)
}
