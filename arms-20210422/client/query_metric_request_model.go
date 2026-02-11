// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMetricRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsistencyDataKey(v string) *QueryMetricRequest
	GetConsistencyDataKey() *string
	SetConsistencyQueryStrategy(v string) *QueryMetricRequest
	GetConsistencyQueryStrategy() *string
	SetDimensions(v []*string) *QueryMetricRequest
	GetDimensions() []*string
	SetEndTime(v int64) *QueryMetricRequest
	GetEndTime() *int64
	SetFilters(v []*QueryMetricRequestFilters) *QueryMetricRequest
	GetFilters() []*QueryMetricRequestFilters
	SetIntervalInSec(v int32) *QueryMetricRequest
	GetIntervalInSec() *int32
	SetLimit(v int32) *QueryMetricRequest
	GetLimit() *int32
	SetMeasures(v []*string) *QueryMetricRequest
	GetMeasures() []*string
	SetMetric(v string) *QueryMetricRequest
	GetMetric() *string
	SetOrder(v string) *QueryMetricRequest
	GetOrder() *string
	SetOrderBy(v string) *QueryMetricRequest
	GetOrderBy() *string
	SetProxyUserId(v string) *QueryMetricRequest
	GetProxyUserId() *string
	SetStartTime(v int64) *QueryMetricRequest
	GetStartTime() *int64
}

type QueryMetricRequest struct {
	ConsistencyDataKey       *string   `json:"ConsistencyDataKey,omitempty" xml:"ConsistencyDataKey,omitempty"`
	ConsistencyQueryStrategy *string   `json:"ConsistencyQueryStrategy,omitempty" xml:"ConsistencyQueryStrategy,omitempty"`
	Dimensions               []*string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// This parameter is required.
	EndTime       *int64                       `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Filters       []*QueryMetricRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	IntervalInSec *int32                       `json:"IntervalInSec,omitempty" xml:"IntervalInSec,omitempty"`
	Limit         *int32                       `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// This parameter is required.
	Measures []*string `json:"Measures,omitempty" xml:"Measures,omitempty" type:"Repeated"`
	// This parameter is required.
	Metric      *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	Order       *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OrderBy     *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	ProxyUserId *string `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	// This parameter is required.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryMetricRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMetricRequest) GoString() string {
	return s.String()
}

func (s *QueryMetricRequest) GetConsistencyDataKey() *string {
	return s.ConsistencyDataKey
}

func (s *QueryMetricRequest) GetConsistencyQueryStrategy() *string {
	return s.ConsistencyQueryStrategy
}

func (s *QueryMetricRequest) GetDimensions() []*string {
	return s.Dimensions
}

func (s *QueryMetricRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryMetricRequest) GetFilters() []*QueryMetricRequestFilters {
	return s.Filters
}

func (s *QueryMetricRequest) GetIntervalInSec() *int32 {
	return s.IntervalInSec
}

func (s *QueryMetricRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *QueryMetricRequest) GetMeasures() []*string {
	return s.Measures
}

func (s *QueryMetricRequest) GetMetric() *string {
	return s.Metric
}

func (s *QueryMetricRequest) GetOrder() *string {
	return s.Order
}

func (s *QueryMetricRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *QueryMetricRequest) GetProxyUserId() *string {
	return s.ProxyUserId
}

func (s *QueryMetricRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *QueryMetricRequest) SetConsistencyDataKey(v string) *QueryMetricRequest {
	s.ConsistencyDataKey = &v
	return s
}

func (s *QueryMetricRequest) SetConsistencyQueryStrategy(v string) *QueryMetricRequest {
	s.ConsistencyQueryStrategy = &v
	return s
}

func (s *QueryMetricRequest) SetDimensions(v []*string) *QueryMetricRequest {
	s.Dimensions = v
	return s
}

func (s *QueryMetricRequest) SetEndTime(v int64) *QueryMetricRequest {
	s.EndTime = &v
	return s
}

func (s *QueryMetricRequest) SetFilters(v []*QueryMetricRequestFilters) *QueryMetricRequest {
	s.Filters = v
	return s
}

func (s *QueryMetricRequest) SetIntervalInSec(v int32) *QueryMetricRequest {
	s.IntervalInSec = &v
	return s
}

func (s *QueryMetricRequest) SetLimit(v int32) *QueryMetricRequest {
	s.Limit = &v
	return s
}

func (s *QueryMetricRequest) SetMeasures(v []*string) *QueryMetricRequest {
	s.Measures = v
	return s
}

func (s *QueryMetricRequest) SetMetric(v string) *QueryMetricRequest {
	s.Metric = &v
	return s
}

func (s *QueryMetricRequest) SetOrder(v string) *QueryMetricRequest {
	s.Order = &v
	return s
}

func (s *QueryMetricRequest) SetOrderBy(v string) *QueryMetricRequest {
	s.OrderBy = &v
	return s
}

func (s *QueryMetricRequest) SetProxyUserId(v string) *QueryMetricRequest {
	s.ProxyUserId = &v
	return s
}

func (s *QueryMetricRequest) SetStartTime(v int64) *QueryMetricRequest {
	s.StartTime = &v
	return s
}

func (s *QueryMetricRequest) Validate() error {
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

type QueryMetricRequestFilters struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryMetricRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s QueryMetricRequestFilters) GoString() string {
	return s.String()
}

func (s *QueryMetricRequestFilters) GetKey() *string {
	return s.Key
}

func (s *QueryMetricRequestFilters) GetValue() *string {
	return s.Value
}

func (s *QueryMetricRequestFilters) SetKey(v string) *QueryMetricRequestFilters {
	s.Key = &v
	return s
}

func (s *QueryMetricRequestFilters) SetValue(v string) *QueryMetricRequestFilters {
	s.Value = &v
	return s
}

func (s *QueryMetricRequestFilters) Validate() error {
	return dara.Validate(s)
}
