// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDoctorHBaseRegionServersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListDoctorHBaseRegionServersResponseBodyData) *ListDoctorHBaseRegionServersResponseBody
	GetData() []*ListDoctorHBaseRegionServersResponseBodyData
	SetMaxResults(v int32) *ListDoctorHBaseRegionServersResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDoctorHBaseRegionServersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDoctorHBaseRegionServersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListDoctorHBaseRegionServersResponseBody
	GetTotalCount() *int32
}

type ListDoctorHBaseRegionServersResponseBody struct {
	// The returned data.
	Data []*ListDoctorHBaseRegionServersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The maximum number of entries that are returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDoctorHBaseRegionServersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseRegionServersResponseBody) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseRegionServersResponseBody) GetData() []*ListDoctorHBaseRegionServersResponseBodyData {
	return s.Data
}

func (s *ListDoctorHBaseRegionServersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDoctorHBaseRegionServersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDoctorHBaseRegionServersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDoctorHBaseRegionServersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDoctorHBaseRegionServersResponseBody) SetData(v []*ListDoctorHBaseRegionServersResponseBodyData) *ListDoctorHBaseRegionServersResponseBody {
	s.Data = v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBody) SetMaxResults(v int32) *ListDoctorHBaseRegionServersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBody) SetNextToken(v string) *ListDoctorHBaseRegionServersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBody) SetRequestId(v string) *ListDoctorHBaseRegionServersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBody) SetTotalCount(v int32) *ListDoctorHBaseRegionServersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHBaseRegionServersResponseBodyData struct {
	// The metric information.
	Metrics *ListDoctorHBaseRegionServersResponseBodyDataMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Struct"`
	// The RegionServer host.
	//
	// example:
	//
	// emr-worker-4.cluster-20****
	RegionServerHost *string `json:"RegionServerHost,omitempty" xml:"RegionServerHost,omitempty"`
}

func (s ListDoctorHBaseRegionServersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseRegionServersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseRegionServersResponseBodyData) GetMetrics() *ListDoctorHBaseRegionServersResponseBodyDataMetrics {
	return s.Metrics
}

func (s *ListDoctorHBaseRegionServersResponseBodyData) GetRegionServerHost() *string {
	return s.RegionServerHost
}

func (s *ListDoctorHBaseRegionServersResponseBodyData) SetMetrics(v *ListDoctorHBaseRegionServersResponseBodyDataMetrics) *ListDoctorHBaseRegionServersResponseBodyData {
	s.Metrics = v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyData) SetRegionServerHost(v string) *ListDoctorHBaseRegionServersResponseBodyData {
	s.RegionServerHost = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHBaseRegionServersResponseBodyDataMetrics struct {
	// The average garbage collection (GC) duration.
	AvgGc *ListDoctorHBaseRegionServersResponseBodyDataMetricsAvgGc `json:"AvgGc,omitempty" xml:"AvgGc,omitempty" type:"Struct"`
	// The cache hit ratio.
	CacheRatio *ListDoctorHBaseRegionServersResponseBodyDataMetricsCacheRatio `json:"CacheRatio,omitempty" xml:"CacheRatio,omitempty" type:"Struct"`
	// The number of daily read requests.
	DailyReadRequest *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyReadRequest `json:"DailyReadRequest,omitempty" xml:"DailyReadRequest,omitempty" type:"Struct"`
	// The growth rate of the number of daily read requests.
	DailyReadRequestDayGrowthRatio *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyReadRequestDayGrowthRatio `json:"DailyReadRequestDayGrowthRatio,omitempty" xml:"DailyReadRequestDayGrowthRatio,omitempty" type:"Struct"`
	// The number of daily write requests.
	DailyWriteRequest *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyWriteRequest `json:"DailyWriteRequest,omitempty" xml:"DailyWriteRequest,omitempty" type:"Struct"`
	// The growth rate of the number of daily write requests.
	DailyWriteRequestDayGrowthRatio *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio `json:"DailyWriteRequestDayGrowthRatio,omitempty" xml:"DailyWriteRequestDayGrowthRatio,omitempty" type:"Struct"`
	// The number of regions.
	RegionCount *ListDoctorHBaseRegionServersResponseBodyDataMetricsRegionCount `json:"RegionCount,omitempty" xml:"RegionCount,omitempty" type:"Struct"`
	// The cumulative number of read requests.
	TotalReadRequest *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalReadRequest `json:"TotalReadRequest,omitempty" xml:"TotalReadRequest,omitempty" type:"Struct"`
	// The cumulative number of all requests.
	TotalRequest *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalRequest `json:"TotalRequest,omitempty" xml:"TotalRequest,omitempty" type:"Struct"`
	// The cumulative number of write requests.
	TotalWriteRequest *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalWriteRequest `json:"TotalWriteRequest,omitempty" xml:"TotalWriteRequest,omitempty" type:"Struct"`
}

func (s ListDoctorHBaseRegionServersResponseBodyDataMetrics) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseRegionServersResponseBodyDataMetrics) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetrics) GetAvgGc() *ListDoctorHBaseRegionServersResponseBodyDataMetricsAvgGc {
	return s.AvgGc
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetrics) GetCacheRatio() *ListDoctorHBaseRegionServersResponseBodyDataMetricsCacheRatio {
	return s.CacheRatio
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetrics) GetDailyReadRequest() *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyReadRequest {
	return s.DailyReadRequest
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetrics) GetDailyReadRequestDayGrowthRatio() *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyReadRequestDayGrowthRatio {
	return s.DailyReadRequestDayGrowthRatio
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetrics) GetDailyWriteRequest() *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyWriteRequest {
	return s.DailyWriteRequest
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetrics) GetDailyWriteRequestDayGrowthRatio() *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio {
	return s.DailyWriteRequestDayGrowthRatio
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetrics) GetRegionCount() *ListDoctorHBaseRegionServersResponseBodyDataMetricsRegionCount {
	return s.RegionCount
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetrics) GetTotalReadRequest() *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalReadRequest {
	return s.TotalReadRequest
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetrics) GetTotalRequest() *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalRequest {
	return s.TotalRequest
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetrics) GetTotalWriteRequest() *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalWriteRequest {
	return s.TotalWriteRequest
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetrics) SetAvgGc(v *ListDoctorHBaseRegionServersResponseBodyDataMetricsAvgGc) *ListDoctorHBaseRegionServersResponseBodyDataMetrics {
	s.AvgGc = v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetrics) SetCacheRatio(v *ListDoctorHBaseRegionServersResponseBodyDataMetricsCacheRatio) *ListDoctorHBaseRegionServersResponseBodyDataMetrics {
	s.CacheRatio = v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetrics) SetDailyReadRequest(v *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyReadRequest) *ListDoctorHBaseRegionServersResponseBodyDataMetrics {
	s.DailyReadRequest = v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetrics) SetDailyReadRequestDayGrowthRatio(v *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) *ListDoctorHBaseRegionServersResponseBodyDataMetrics {
	s.DailyReadRequestDayGrowthRatio = v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetrics) SetDailyWriteRequest(v *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyWriteRequest) *ListDoctorHBaseRegionServersResponseBodyDataMetrics {
	s.DailyWriteRequest = v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetrics) SetDailyWriteRequestDayGrowthRatio(v *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) *ListDoctorHBaseRegionServersResponseBodyDataMetrics {
	s.DailyWriteRequestDayGrowthRatio = v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetrics) SetRegionCount(v *ListDoctorHBaseRegionServersResponseBodyDataMetricsRegionCount) *ListDoctorHBaseRegionServersResponseBodyDataMetrics {
	s.RegionCount = v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetrics) SetTotalReadRequest(v *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalReadRequest) *ListDoctorHBaseRegionServersResponseBodyDataMetrics {
	s.TotalReadRequest = v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetrics) SetTotalRequest(v *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalRequest) *ListDoctorHBaseRegionServersResponseBodyDataMetrics {
	s.TotalRequest = v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetrics) SetTotalWriteRequest(v *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalWriteRequest) *ListDoctorHBaseRegionServersResponseBodyDataMetrics {
	s.TotalWriteRequest = v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetrics) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHBaseRegionServersResponseBodyDataMetricsAvgGc struct {
	// The description of the metric.
	//
	// example:
	//
	// The efficiency of garbage collection in the system
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// avgGc
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 37.9
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHBaseRegionServersResponseBodyDataMetricsAvgGc) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseRegionServersResponseBodyDataMetricsAvgGc) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsAvgGc) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsAvgGc) GetName() *string {
	return s.Name
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsAvgGc) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsAvgGc) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsAvgGc) SetDescription(v string) *ListDoctorHBaseRegionServersResponseBodyDataMetricsAvgGc {
	s.Description = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsAvgGc) SetName(v string) *ListDoctorHBaseRegionServersResponseBodyDataMetricsAvgGc {
	s.Name = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsAvgGc) SetUnit(v string) *ListDoctorHBaseRegionServersResponseBodyDataMetricsAvgGc {
	s.Unit = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsAvgGc) SetValue(v float32) *ListDoctorHBaseRegionServersResponseBodyDataMetricsAvgGc {
	s.Value = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsAvgGc) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHBaseRegionServersResponseBodyDataMetricsCacheRatio struct {
	// The description of the metric.
	//
	// example:
	//
	// Ratio of the BlockCache memory size
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// cacheRatio
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 96.7
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHBaseRegionServersResponseBodyDataMetricsCacheRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseRegionServersResponseBodyDataMetricsCacheRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsCacheRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsCacheRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsCacheRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsCacheRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsCacheRatio) SetDescription(v string) *ListDoctorHBaseRegionServersResponseBodyDataMetricsCacheRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsCacheRatio) SetName(v string) *ListDoctorHBaseRegionServersResponseBodyDataMetricsCacheRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsCacheRatio) SetUnit(v string) *ListDoctorHBaseRegionServersResponseBodyDataMetricsCacheRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsCacheRatio) SetValue(v float32) *ListDoctorHBaseRegionServersResponseBodyDataMetricsCacheRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsCacheRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyReadRequest struct {
	// The description of the metric.
	//
	// example:
	//
	// Number of read requests per day
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// dailyReadRequest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 42571
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyReadRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyReadRequest) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyReadRequest) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyReadRequest) GetName() *string {
	return s.Name
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyReadRequest) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyReadRequest) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyReadRequest) SetDescription(v string) *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyReadRequest {
	s.Description = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyReadRequest) SetName(v string) *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyReadRequest {
	s.Name = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyReadRequest) SetUnit(v string) *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyReadRequest {
	s.Unit = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyReadRequest) SetValue(v int64) *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyReadRequest {
	s.Value = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyReadRequest) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyReadRequestDayGrowthRatio struct {
	// The description of the metric.
	//
	// example:
	//
	// The growth rate of daily read request quantity.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// dailyReadRequestDayGrowthRatio
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 1.0
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) SetDescription(v string) *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyReadRequestDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) SetName(v string) *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyReadRequestDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) SetUnit(v string) *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyReadRequestDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) SetValue(v float32) *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyReadRequestDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyWriteRequest struct {
	// The description of the metric.
	//
	// example:
	//
	// Number of write requests per day
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// dailyWriteRequest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 23124
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyWriteRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyWriteRequest) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyWriteRequest) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyWriteRequest) GetName() *string {
	return s.Name
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyWriteRequest) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyWriteRequest) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyWriteRequest) SetDescription(v string) *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyWriteRequest {
	s.Description = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyWriteRequest) SetName(v string) *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyWriteRequest {
	s.Name = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyWriteRequest) SetUnit(v string) *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyWriteRequest {
	s.Unit = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyWriteRequest) SetValue(v int64) *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyWriteRequest {
	s.Value = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyWriteRequest) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio struct {
	// The description of the metric.
	//
	// example:
	//
	// The growth rate of daily write request quantity.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// dailyWriteRequestDayGrowthRatio
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 1.0
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) SetDescription(v string) *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) SetName(v string) *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) SetUnit(v string) *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) SetValue(v float32) *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHBaseRegionServersResponseBodyDataMetricsRegionCount struct {
	// The description of the metric.
	//
	// example:
	//
	// Number of regions count
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// regionCount
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 81
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHBaseRegionServersResponseBodyDataMetricsRegionCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseRegionServersResponseBodyDataMetricsRegionCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsRegionCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsRegionCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsRegionCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsRegionCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsRegionCount) SetDescription(v string) *ListDoctorHBaseRegionServersResponseBodyDataMetricsRegionCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsRegionCount) SetName(v string) *ListDoctorHBaseRegionServersResponseBodyDataMetricsRegionCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsRegionCount) SetUnit(v string) *ListDoctorHBaseRegionServersResponseBodyDataMetricsRegionCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsRegionCount) SetValue(v int64) *ListDoctorHBaseRegionServersResponseBodyDataMetricsRegionCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsRegionCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalReadRequest struct {
	// The description of the metric.
	//
	// example:
	//
	// Total number of read requests
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// totalReadRequest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 170500567
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalReadRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalReadRequest) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalReadRequest) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalReadRequest) GetName() *string {
	return s.Name
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalReadRequest) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalReadRequest) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalReadRequest) SetDescription(v string) *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalReadRequest {
	s.Description = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalReadRequest) SetName(v string) *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalReadRequest {
	s.Name = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalReadRequest) SetUnit(v string) *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalReadRequest {
	s.Unit = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalReadRequest) SetValue(v int64) *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalReadRequest {
	s.Value = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalReadRequest) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalRequest struct {
	// The description of the metric.
	//
	// example:
	//
	// Total number of requests
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// totalRequest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 89499511
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalRequest) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalRequest) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalRequest) GetName() *string {
	return s.Name
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalRequest) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalRequest) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalRequest) SetDescription(v string) *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalRequest {
	s.Description = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalRequest) SetName(v string) *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalRequest {
	s.Name = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalRequest) SetUnit(v string) *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalRequest {
	s.Unit = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalRequest) SetValue(v int64) *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalRequest {
	s.Value = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalRequest) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalWriteRequest struct {
	// The description of the metric.
	//
	// example:
	//
	// Total number of write requests
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// totalWriteRequest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 30109837
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalWriteRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalWriteRequest) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalWriteRequest) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalWriteRequest) GetName() *string {
	return s.Name
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalWriteRequest) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalWriteRequest) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalWriteRequest) SetDescription(v string) *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalWriteRequest {
	s.Description = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalWriteRequest) SetName(v string) *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalWriteRequest {
	s.Name = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalWriteRequest) SetUnit(v string) *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalWriteRequest {
	s.Unit = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalWriteRequest) SetValue(v int64) *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalWriteRequest {
	s.Value = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponseBodyDataMetricsTotalWriteRequest) Validate() error {
	return dara.Validate(s)
}
