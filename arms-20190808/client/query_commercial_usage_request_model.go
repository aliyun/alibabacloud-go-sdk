// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCommercialUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdvancedFilters(v []*QueryCommercialUsageRequestAdvancedFilters) *QueryCommercialUsageRequest
	GetAdvancedFilters() []*QueryCommercialUsageRequestAdvancedFilters
	SetDimensions(v []*string) *QueryCommercialUsageRequest
	GetDimensions() []*string
	SetEndTime(v int64) *QueryCommercialUsageRequest
	GetEndTime() *int64
	SetIntervalInSec(v int32) *QueryCommercialUsageRequest
	GetIntervalInSec() *int32
	SetMeasures(v []*string) *QueryCommercialUsageRequest
	GetMeasures() []*string
	SetMetric(v string) *QueryCommercialUsageRequest
	GetMetric() *string
	SetOrder(v string) *QueryCommercialUsageRequest
	GetOrder() *string
	SetOrderBy(v string) *QueryCommercialUsageRequest
	GetOrderBy() *string
	SetQueryType(v string) *QueryCommercialUsageRequest
	GetQueryType() *string
	SetStartTime(v int64) *QueryCommercialUsageRequest
	GetStartTime() *int64
}

type QueryCommercialUsageRequest struct {
	// The filter conditions.
	AdvancedFilters []*QueryCommercialUsageRequestAdvancedFilters `json:"AdvancedFilters,omitempty" xml:"AdvancedFilters,omitempty" type:"Repeated"`
	// The dimensions of the metric that you want to query. Valid values:
	//
	// 	- dataType: data type
	//
	// 	- productType: product type
	//
	// 	- instanceId: instance ID
	//
	// 	- instanceName: instance name
	//
	// 	- instanceType: instance type
	Dimensions []*string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The end of the time range to query. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1699286400000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time interval between data slices. Unit: seconds. Minimum value: 3600.
	//
	// Valid values:
	//
	// 	- 3600: 1 hour
	//
	// 	- 86400: 1 day
	//
	// example:
	//
	// 3600
	IntervalInSec *int32 `json:"IntervalInSec,omitempty" xml:"IntervalInSec,omitempty"`
	// The measures of the metric that you want to query.
	Measures []*string `json:"Measures,omitempty" xml:"Measures,omitempty" type:"Repeated"`
	// The name of the metric. Valid value: USAGEFEE.STAT.
	//
	// This parameter is required.
	//
	// example:
	//
	// USAGEFEE.STAT
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// The order in which data is sorted. Valid value:
	//
	// 	- `ASC`: ascending order
	//
	// 	- `DESC`: descending order
	//
	// example:
	//
	// ASC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The dimension by which data is sorted.
	//
	// Valid value:
	//
	// 	- dataType
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// dataType
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The data type. Valid values:
	//
	// 	- instantQuery: non-time series
	//
	// 	- timeSeriesQuery: time series
	//
	// This parameter is required.
	//
	// example:
	//
	// instantQuery
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	// The start of the time range to query. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1699200000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryCommercialUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCommercialUsageRequest) GoString() string {
	return s.String()
}

func (s *QueryCommercialUsageRequest) GetAdvancedFilters() []*QueryCommercialUsageRequestAdvancedFilters {
	return s.AdvancedFilters
}

func (s *QueryCommercialUsageRequest) GetDimensions() []*string {
	return s.Dimensions
}

func (s *QueryCommercialUsageRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryCommercialUsageRequest) GetIntervalInSec() *int32 {
	return s.IntervalInSec
}

func (s *QueryCommercialUsageRequest) GetMeasures() []*string {
	return s.Measures
}

func (s *QueryCommercialUsageRequest) GetMetric() *string {
	return s.Metric
}

func (s *QueryCommercialUsageRequest) GetOrder() *string {
	return s.Order
}

func (s *QueryCommercialUsageRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *QueryCommercialUsageRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *QueryCommercialUsageRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *QueryCommercialUsageRequest) SetAdvancedFilters(v []*QueryCommercialUsageRequestAdvancedFilters) *QueryCommercialUsageRequest {
	s.AdvancedFilters = v
	return s
}

func (s *QueryCommercialUsageRequest) SetDimensions(v []*string) *QueryCommercialUsageRequest {
	s.Dimensions = v
	return s
}

func (s *QueryCommercialUsageRequest) SetEndTime(v int64) *QueryCommercialUsageRequest {
	s.EndTime = &v
	return s
}

func (s *QueryCommercialUsageRequest) SetIntervalInSec(v int32) *QueryCommercialUsageRequest {
	s.IntervalInSec = &v
	return s
}

func (s *QueryCommercialUsageRequest) SetMeasures(v []*string) *QueryCommercialUsageRequest {
	s.Measures = v
	return s
}

func (s *QueryCommercialUsageRequest) SetMetric(v string) *QueryCommercialUsageRequest {
	s.Metric = &v
	return s
}

func (s *QueryCommercialUsageRequest) SetOrder(v string) *QueryCommercialUsageRequest {
	s.Order = &v
	return s
}

func (s *QueryCommercialUsageRequest) SetOrderBy(v string) *QueryCommercialUsageRequest {
	s.OrderBy = &v
	return s
}

func (s *QueryCommercialUsageRequest) SetQueryType(v string) *QueryCommercialUsageRequest {
	s.QueryType = &v
	return s
}

func (s *QueryCommercialUsageRequest) SetStartTime(v int64) *QueryCommercialUsageRequest {
	s.StartTime = &v
	return s
}

func (s *QueryCommercialUsageRequest) Validate() error {
	return dara.Validate(s)
}

type QueryCommercialUsageRequestAdvancedFilters struct {
	// The key of the filter condition.
	//
	// example:
	//
	// regionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The operator. Valid values: eq and in.
	//
	// example:
	//
	// eq
	OpType *string `json:"OpType,omitempty" xml:"OpType,omitempty"`
	// The value of the filter condition.
	//
	// example:
	//
	// cn-hangzhou
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryCommercialUsageRequestAdvancedFilters) String() string {
	return dara.Prettify(s)
}

func (s QueryCommercialUsageRequestAdvancedFilters) GoString() string {
	return s.String()
}

func (s *QueryCommercialUsageRequestAdvancedFilters) GetKey() *string {
	return s.Key
}

func (s *QueryCommercialUsageRequestAdvancedFilters) GetOpType() *string {
	return s.OpType
}

func (s *QueryCommercialUsageRequestAdvancedFilters) GetValue() *string {
	return s.Value
}

func (s *QueryCommercialUsageRequestAdvancedFilters) SetKey(v string) *QueryCommercialUsageRequestAdvancedFilters {
	s.Key = &v
	return s
}

func (s *QueryCommercialUsageRequestAdvancedFilters) SetOpType(v string) *QueryCommercialUsageRequestAdvancedFilters {
	s.OpType = &v
	return s
}

func (s *QueryCommercialUsageRequestAdvancedFilters) SetValue(v string) *QueryCommercialUsageRequestAdvancedFilters {
	s.Value = &v
	return s
}

func (s *QueryCommercialUsageRequestAdvancedFilters) Validate() error {
	return dara.Validate(s)
}
