// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorHBaseRegionServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDoctorHBaseRegionServerResponseBodyData) *GetDoctorHBaseRegionServerResponseBody
	GetData() *GetDoctorHBaseRegionServerResponseBodyData
	SetRequestId(v string) *GetDoctorHBaseRegionServerResponseBody
	GetRequestId() *string
}

type GetDoctorHBaseRegionServerResponseBody struct {
	// The returned data.
	Data *GetDoctorHBaseRegionServerResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDoctorHBaseRegionServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseRegionServerResponseBody) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseRegionServerResponseBody) GetData() *GetDoctorHBaseRegionServerResponseBodyData {
	return s.Data
}

func (s *GetDoctorHBaseRegionServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDoctorHBaseRegionServerResponseBody) SetData(v *GetDoctorHBaseRegionServerResponseBodyData) *GetDoctorHBaseRegionServerResponseBody {
	s.Data = v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBody) SetRequestId(v string) *GetDoctorHBaseRegionServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseRegionServerResponseBodyData struct {
	// The metric information.
	Metrics *GetDoctorHBaseRegionServerResponseBodyDataMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Struct"`
}

func (s GetDoctorHBaseRegionServerResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseRegionServerResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseRegionServerResponseBodyData) GetMetrics() *GetDoctorHBaseRegionServerResponseBodyDataMetrics {
	return s.Metrics
}

func (s *GetDoctorHBaseRegionServerResponseBodyData) SetMetrics(v *GetDoctorHBaseRegionServerResponseBodyDataMetrics) *GetDoctorHBaseRegionServerResponseBodyData {
	s.Metrics = v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseRegionServerResponseBodyDataMetrics struct {
	// The average garbage collection (GC) duration.
	AvgGc *GetDoctorHBaseRegionServerResponseBodyDataMetricsAvgGc `json:"AvgGc,omitempty" xml:"AvgGc,omitempty" type:"Struct"`
	// The cache hit ratio.
	CacheRatio *GetDoctorHBaseRegionServerResponseBodyDataMetricsCacheRatio `json:"CacheRatio,omitempty" xml:"CacheRatio,omitempty" type:"Struct"`
	// The number of daily read requests.
	DailyReadRequest *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyReadRequest `json:"DailyReadRequest,omitempty" xml:"DailyReadRequest,omitempty" type:"Struct"`
	// The day-to-day increment rate of the number of daily read requests.
	DailyReadRequestDayGrowthRatio *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyReadRequestDayGrowthRatio `json:"DailyReadRequestDayGrowthRatio,omitempty" xml:"DailyReadRequestDayGrowthRatio,omitempty" type:"Struct"`
	// The number of daily write requests.
	DailyWriteRequest *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyWriteRequest `json:"DailyWriteRequest,omitempty" xml:"DailyWriteRequest,omitempty" type:"Struct"`
	// The day-to-day increment rate of the number of daily write requests.
	DailyWriteRequestDayGrowthRatio *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio `json:"DailyWriteRequestDayGrowthRatio,omitempty" xml:"DailyWriteRequestDayGrowthRatio,omitempty" type:"Struct"`
	// The number of regions.
	RegionCount *GetDoctorHBaseRegionServerResponseBodyDataMetricsRegionCount `json:"RegionCount,omitempty" xml:"RegionCount,omitempty" type:"Struct"`
	// The cumulative number of read requests.
	TotalReadRequest *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalReadRequest `json:"TotalReadRequest,omitempty" xml:"TotalReadRequest,omitempty" type:"Struct"`
	// The cumulative number of total requests.
	TotalRequest *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalRequest `json:"TotalRequest,omitempty" xml:"TotalRequest,omitempty" type:"Struct"`
	// The cumulative number of write requests.
	TotalWriteRequest *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalWriteRequest `json:"TotalWriteRequest,omitempty" xml:"TotalWriteRequest,omitempty" type:"Struct"`
}

func (s GetDoctorHBaseRegionServerResponseBodyDataMetrics) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseRegionServerResponseBodyDataMetrics) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetrics) GetAvgGc() *GetDoctorHBaseRegionServerResponseBodyDataMetricsAvgGc {
	return s.AvgGc
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetrics) GetCacheRatio() *GetDoctorHBaseRegionServerResponseBodyDataMetricsCacheRatio {
	return s.CacheRatio
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetrics) GetDailyReadRequest() *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyReadRequest {
	return s.DailyReadRequest
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetrics) GetDailyReadRequestDayGrowthRatio() *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyReadRequestDayGrowthRatio {
	return s.DailyReadRequestDayGrowthRatio
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetrics) GetDailyWriteRequest() *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyWriteRequest {
	return s.DailyWriteRequest
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetrics) GetDailyWriteRequestDayGrowthRatio() *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio {
	return s.DailyWriteRequestDayGrowthRatio
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetrics) GetRegionCount() *GetDoctorHBaseRegionServerResponseBodyDataMetricsRegionCount {
	return s.RegionCount
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetrics) GetTotalReadRequest() *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalReadRequest {
	return s.TotalReadRequest
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetrics) GetTotalRequest() *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalRequest {
	return s.TotalRequest
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetrics) GetTotalWriteRequest() *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalWriteRequest {
	return s.TotalWriteRequest
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetrics) SetAvgGc(v *GetDoctorHBaseRegionServerResponseBodyDataMetricsAvgGc) *GetDoctorHBaseRegionServerResponseBodyDataMetrics {
	s.AvgGc = v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetrics) SetCacheRatio(v *GetDoctorHBaseRegionServerResponseBodyDataMetricsCacheRatio) *GetDoctorHBaseRegionServerResponseBodyDataMetrics {
	s.CacheRatio = v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetrics) SetDailyReadRequest(v *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyReadRequest) *GetDoctorHBaseRegionServerResponseBodyDataMetrics {
	s.DailyReadRequest = v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetrics) SetDailyReadRequestDayGrowthRatio(v *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) *GetDoctorHBaseRegionServerResponseBodyDataMetrics {
	s.DailyReadRequestDayGrowthRatio = v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetrics) SetDailyWriteRequest(v *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyWriteRequest) *GetDoctorHBaseRegionServerResponseBodyDataMetrics {
	s.DailyWriteRequest = v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetrics) SetDailyWriteRequestDayGrowthRatio(v *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) *GetDoctorHBaseRegionServerResponseBodyDataMetrics {
	s.DailyWriteRequestDayGrowthRatio = v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetrics) SetRegionCount(v *GetDoctorHBaseRegionServerResponseBodyDataMetricsRegionCount) *GetDoctorHBaseRegionServerResponseBodyDataMetrics {
	s.RegionCount = v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetrics) SetTotalReadRequest(v *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalReadRequest) *GetDoctorHBaseRegionServerResponseBodyDataMetrics {
	s.TotalReadRequest = v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetrics) SetTotalRequest(v *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalRequest) *GetDoctorHBaseRegionServerResponseBodyDataMetrics {
	s.TotalRequest = v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetrics) SetTotalWriteRequest(v *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalWriteRequest) *GetDoctorHBaseRegionServerResponseBodyDataMetrics {
	s.TotalWriteRequest = v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetrics) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseRegionServerResponseBodyDataMetricsAvgGc struct {
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
	// 42.3
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseRegionServerResponseBodyDataMetricsAvgGc) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseRegionServerResponseBodyDataMetricsAvgGc) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsAvgGc) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsAvgGc) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsAvgGc) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsAvgGc) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsAvgGc) SetDescription(v string) *GetDoctorHBaseRegionServerResponseBodyDataMetricsAvgGc {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsAvgGc) SetName(v string) *GetDoctorHBaseRegionServerResponseBodyDataMetricsAvgGc {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsAvgGc) SetUnit(v string) *GetDoctorHBaseRegionServerResponseBodyDataMetricsAvgGc {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsAvgGc) SetValue(v float32) *GetDoctorHBaseRegionServerResponseBodyDataMetricsAvgGc {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsAvgGc) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseRegionServerResponseBodyDataMetricsCacheRatio struct {
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
	// 95.3
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseRegionServerResponseBodyDataMetricsCacheRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseRegionServerResponseBodyDataMetricsCacheRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsCacheRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsCacheRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsCacheRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsCacheRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsCacheRatio) SetDescription(v string) *GetDoctorHBaseRegionServerResponseBodyDataMetricsCacheRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsCacheRatio) SetName(v string) *GetDoctorHBaseRegionServerResponseBodyDataMetricsCacheRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsCacheRatio) SetUnit(v string) *GetDoctorHBaseRegionServerResponseBodyDataMetricsCacheRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsCacheRatio) SetValue(v float32) *GetDoctorHBaseRegionServerResponseBodyDataMetricsCacheRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsCacheRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyReadRequest struct {
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
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyReadRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyReadRequest) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyReadRequest) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyReadRequest) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyReadRequest) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyReadRequest) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyReadRequest) SetDescription(v string) *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyReadRequest {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyReadRequest) SetName(v string) *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyReadRequest {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyReadRequest) SetUnit(v string) *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyReadRequest {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyReadRequest) SetValue(v int64) *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyReadRequest {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyReadRequest) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyReadRequestDayGrowthRatio struct {
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

func (s GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) SetDescription(v string) *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyReadRequestDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) SetName(v string) *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyReadRequestDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) SetUnit(v string) *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyReadRequestDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) SetValue(v float32) *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyReadRequestDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyWriteRequest struct {
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
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyWriteRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyWriteRequest) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyWriteRequest) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyWriteRequest) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyWriteRequest) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyWriteRequest) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyWriteRequest) SetDescription(v string) *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyWriteRequest {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyWriteRequest) SetName(v string) *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyWriteRequest {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyWriteRequest) SetUnit(v string) *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyWriteRequest {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyWriteRequest) SetValue(v int64) *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyWriteRequest {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyWriteRequest) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio struct {
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

func (s GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) SetDescription(v string) *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) SetName(v string) *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) SetUnit(v string) *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) SetValue(v float32) *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseRegionServerResponseBodyDataMetricsRegionCount struct {
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
	// 15
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseRegionServerResponseBodyDataMetricsRegionCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseRegionServerResponseBodyDataMetricsRegionCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsRegionCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsRegionCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsRegionCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsRegionCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsRegionCount) SetDescription(v string) *GetDoctorHBaseRegionServerResponseBodyDataMetricsRegionCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsRegionCount) SetName(v string) *GetDoctorHBaseRegionServerResponseBodyDataMetricsRegionCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsRegionCount) SetUnit(v string) *GetDoctorHBaseRegionServerResponseBodyDataMetricsRegionCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsRegionCount) SetValue(v int64) *GetDoctorHBaseRegionServerResponseBodyDataMetricsRegionCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsRegionCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalReadRequest struct {
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
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalReadRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalReadRequest) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalReadRequest) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalReadRequest) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalReadRequest) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalReadRequest) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalReadRequest) SetDescription(v string) *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalReadRequest {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalReadRequest) SetName(v string) *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalReadRequest {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalReadRequest) SetUnit(v string) *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalReadRequest {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalReadRequest) SetValue(v int64) *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalReadRequest {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalReadRequest) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalRequest struct {
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
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalRequest) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalRequest) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalRequest) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalRequest) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalRequest) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalRequest) SetDescription(v string) *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalRequest {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalRequest) SetName(v string) *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalRequest {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalRequest) SetUnit(v string) *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalRequest {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalRequest) SetValue(v int64) *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalRequest {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalRequest) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalWriteRequest struct {
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
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalWriteRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalWriteRequest) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalWriteRequest) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalWriteRequest) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalWriteRequest) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalWriteRequest) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalWriteRequest) SetDescription(v string) *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalWriteRequest {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalWriteRequest) SetName(v string) *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalWriteRequest {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalWriteRequest) SetUnit(v string) *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalWriteRequest {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalWriteRequest) SetValue(v int64) *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalWriteRequest {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponseBodyDataMetricsTotalWriteRequest) Validate() error {
	return dara.Validate(s)
}
