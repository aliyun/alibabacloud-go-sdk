// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorHBaseTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDoctorHBaseTableResponseBodyData) *GetDoctorHBaseTableResponseBody
	GetData() *GetDoctorHBaseTableResponseBodyData
	SetRequestId(v string) *GetDoctorHBaseTableResponseBody
	GetRequestId() *string
}

type GetDoctorHBaseTableResponseBody struct {
	// Returned data.
	Data *GetDoctorHBaseTableResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDoctorHBaseTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseTableResponseBody) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseTableResponseBody) GetData() *GetDoctorHBaseTableResponseBodyData {
	return s.Data
}

func (s *GetDoctorHBaseTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDoctorHBaseTableResponseBody) SetData(v *GetDoctorHBaseTableResponseBodyData) *GetDoctorHBaseTableResponseBody {
	s.Data = v
	return s
}

func (s *GetDoctorHBaseTableResponseBody) SetRequestId(v string) *GetDoctorHBaseTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseTableResponseBodyData struct {
	// Diagnostic results.
	Analysis *GetDoctorHBaseTableResponseBodyDataAnalysis `json:"Analysis,omitempty" xml:"Analysis,omitempty" type:"Struct"`
	// Metrics information.
	Metrics *GetDoctorHBaseTableResponseBodyDataMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Struct"`
}

func (s GetDoctorHBaseTableResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseTableResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseTableResponseBodyData) GetAnalysis() *GetDoctorHBaseTableResponseBodyDataAnalysis {
	return s.Analysis
}

func (s *GetDoctorHBaseTableResponseBodyData) GetMetrics() *GetDoctorHBaseTableResponseBodyDataMetrics {
	return s.Metrics
}

func (s *GetDoctorHBaseTableResponseBodyData) SetAnalysis(v *GetDoctorHBaseTableResponseBodyDataAnalysis) *GetDoctorHBaseTableResponseBodyData {
	s.Analysis = v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyData) SetMetrics(v *GetDoctorHBaseTableResponseBodyDataMetrics) *GetDoctorHBaseTableResponseBodyData {
	s.Metrics = v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseTableResponseBodyDataAnalysis struct {
	// List of read hotspot regions.
	//
	// example:
	//
	// null
	ReadRequestHotspotRegionList []*string `json:"ReadRequestHotspotRegionList,omitempty" xml:"ReadRequestHotspotRegionList,omitempty" type:"Repeated"`
	// Description of read imbalance.
	ReadRequestUnbalanceSuggestion *string `json:"ReadRequestUnbalanceSuggestion,omitempty" xml:"ReadRequestUnbalanceSuggestion,omitempty"`
	// List of read/write hotspot regions.
	//
	// example:
	//
	// null
	RequestHotspotRegionList []*string `json:"RequestHotspotRegionList,omitempty" xml:"RequestHotspotRegionList,omitempty" type:"Repeated"`
	// Description of read/write imbalance.
	RequestUnbalanceSuggestion *string `json:"RequestUnbalanceSuggestion,omitempty" xml:"RequestUnbalanceSuggestion,omitempty"`
	// Table score.
	//
	// example:
	//
	// 85
	TableScore *int32 `json:"TableScore,omitempty" xml:"TableScore,omitempty"`
	// List of write hotspot regions.
	//
	// example:
	//
	// null
	WriteRequestHotspotRegionList []*string `json:"WriteRequestHotspotRegionList,omitempty" xml:"WriteRequestHotspotRegionList,omitempty" type:"Repeated"`
	// Description of write imbalance.
	WriteRequestUnbalanceSuggestion *string `json:"WriteRequestUnbalanceSuggestion,omitempty" xml:"WriteRequestUnbalanceSuggestion,omitempty"`
}

func (s GetDoctorHBaseTableResponseBodyDataAnalysis) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseTableResponseBodyDataAnalysis) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseTableResponseBodyDataAnalysis) GetReadRequestHotspotRegionList() []*string {
	return s.ReadRequestHotspotRegionList
}

func (s *GetDoctorHBaseTableResponseBodyDataAnalysis) GetReadRequestUnbalanceSuggestion() *string {
	return s.ReadRequestUnbalanceSuggestion
}

func (s *GetDoctorHBaseTableResponseBodyDataAnalysis) GetRequestHotspotRegionList() []*string {
	return s.RequestHotspotRegionList
}

func (s *GetDoctorHBaseTableResponseBodyDataAnalysis) GetRequestUnbalanceSuggestion() *string {
	return s.RequestUnbalanceSuggestion
}

func (s *GetDoctorHBaseTableResponseBodyDataAnalysis) GetTableScore() *int32 {
	return s.TableScore
}

func (s *GetDoctorHBaseTableResponseBodyDataAnalysis) GetWriteRequestHotspotRegionList() []*string {
	return s.WriteRequestHotspotRegionList
}

func (s *GetDoctorHBaseTableResponseBodyDataAnalysis) GetWriteRequestUnbalanceSuggestion() *string {
	return s.WriteRequestUnbalanceSuggestion
}

func (s *GetDoctorHBaseTableResponseBodyDataAnalysis) SetReadRequestHotspotRegionList(v []*string) *GetDoctorHBaseTableResponseBodyDataAnalysis {
	s.ReadRequestHotspotRegionList = v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataAnalysis) SetReadRequestUnbalanceSuggestion(v string) *GetDoctorHBaseTableResponseBodyDataAnalysis {
	s.ReadRequestUnbalanceSuggestion = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataAnalysis) SetRequestHotspotRegionList(v []*string) *GetDoctorHBaseTableResponseBodyDataAnalysis {
	s.RequestHotspotRegionList = v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataAnalysis) SetRequestUnbalanceSuggestion(v string) *GetDoctorHBaseTableResponseBodyDataAnalysis {
	s.RequestUnbalanceSuggestion = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataAnalysis) SetTableScore(v int32) *GetDoctorHBaseTableResponseBodyDataAnalysis {
	s.TableScore = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataAnalysis) SetWriteRequestHotspotRegionList(v []*string) *GetDoctorHBaseTableResponseBodyDataAnalysis {
	s.WriteRequestHotspotRegionList = v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataAnalysis) SetWriteRequestUnbalanceSuggestion(v string) *GetDoctorHBaseTableResponseBodyDataAnalysis {
	s.WriteRequestUnbalanceSuggestion = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataAnalysis) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseTableResponseBodyDataMetrics struct {
	// Number of days the table has not been accessed.
	ColdAccessDay *GetDoctorHBaseTableResponseBodyDataMetricsColdAccessDay `json:"ColdAccessDay,omitempty" xml:"ColdAccessDay,omitempty" type:"Struct"`
	// Cold data access days configuration.
	ColdConfigDay *GetDoctorHBaseTableResponseBodyDataMetricsColdConfigDay `json:"ColdConfigDay,omitempty" xml:"ColdConfigDay,omitempty" type:"Struct"`
	// Cold data size.
	ColdDataSize *GetDoctorHBaseTableResponseBodyDataMetricsColdDataSize `json:"ColdDataSize,omitempty" xml:"ColdDataSize,omitempty" type:"Struct"`
	// Number of read requests per day.
	DailyReadRequest *GetDoctorHBaseTableResponseBodyDataMetricsDailyReadRequest `json:"DailyReadRequest,omitempty" xml:"DailyReadRequest,omitempty" type:"Struct"`
	// Daily growth ratio of daily read requests.
	DailyReadRequestDayGrowthRatio *GetDoctorHBaseTableResponseBodyDataMetricsDailyReadRequestDayGrowthRatio `json:"DailyReadRequestDayGrowthRatio,omitempty" xml:"DailyReadRequestDayGrowthRatio,omitempty" type:"Struct"`
	// Number of write requests per day.
	DailyWriteRequest *GetDoctorHBaseTableResponseBodyDataMetricsDailyWriteRequest `json:"DailyWriteRequest,omitempty" xml:"DailyWriteRequest,omitempty" type:"Struct"`
	// Daily write request growth ratio.
	DailyWriteRequestDayGrowthRatio *GetDoctorHBaseTableResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio `json:"DailyWriteRequestDayGrowthRatio,omitempty" xml:"DailyWriteRequestDayGrowthRatio,omitempty" type:"Struct"`
	// Configuration for the number of days cold data is accessed.
	FreezeConfigDay *GetDoctorHBaseTableResponseBodyDataMetricsFreezeConfigDay `json:"FreezeConfigDay,omitempty" xml:"FreezeConfigDay,omitempty" type:"Struct"`
	// Frozen data size.
	FreezeDataSize *GetDoctorHBaseTableResponseBodyDataMetricsFreezeDataSize `json:"FreezeDataSize,omitempty" xml:"FreezeDataSize,omitempty" type:"Struct"`
	// Hot data size.
	HotDataSize *GetDoctorHBaseTableResponseBodyDataMetricsHotDataSize `json:"HotDataSize,omitempty" xml:"HotDataSize,omitempty" type:"Struct"`
	// Locality rate.
	Locality *GetDoctorHBaseTableResponseBodyDataMetricsLocality `json:"Locality,omitempty" xml:"Locality,omitempty" type:"Struct"`
	// Read request balance.
	ReadRequestBalance *GetDoctorHBaseTableResponseBodyDataMetricsReadRequestBalance `json:"ReadRequestBalance,omitempty" xml:"ReadRequestBalance,omitempty" type:"Struct"`
	// Region balance.
	RegionBalance *GetDoctorHBaseTableResponseBodyDataMetricsRegionBalance `json:"RegionBalance,omitempty" xml:"RegionBalance,omitempty" type:"Struct"`
	// Number of regions.
	RegionCount *GetDoctorHBaseTableResponseBodyDataMetricsRegionCount `json:"RegionCount,omitempty" xml:"RegionCount,omitempty" type:"Struct"`
	// Daily incremental ratio of regions
	RegionCountDayGrowthRatio *GetDoctorHBaseTableResponseBodyDataMetricsRegionCountDayGrowthRatio `json:"RegionCountDayGrowthRatio,omitempty" xml:"RegionCountDayGrowthRatio,omitempty" type:"Struct"`
	// Number of RegionServers.
	RegionServerCount *GetDoctorHBaseTableResponseBodyDataMetricsRegionServerCount `json:"RegionServerCount,omitempty" xml:"RegionServerCount,omitempty" type:"Struct"`
	// Request balance.
	RequestBalance *GetDoctorHBaseTableResponseBodyDataMetricsRequestBalance `json:"RequestBalance,omitempty" xml:"RequestBalance,omitempty" type:"Struct"`
	// Number of store files.
	StoreFileCount *GetDoctorHBaseTableResponseBodyDataMetricsStoreFileCount `json:"StoreFileCount,omitempty" xml:"StoreFileCount,omitempty" type:"Struct"`
	// Daily growth ratio of store file count.
	StoreFileCountDayGrowthRatio *GetDoctorHBaseTableResponseBodyDataMetricsStoreFileCountDayGrowthRatio `json:"StoreFileCountDayGrowthRatio,omitempty" xml:"StoreFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// Table size.
	TableSize *GetDoctorHBaseTableResponseBodyDataMetricsTableSize `json:"TableSize,omitempty" xml:"TableSize,omitempty" type:"Struct"`
	// Daily growth ratio of table size.
	TableSizeDayGrowthRatio *GetDoctorHBaseTableResponseBodyDataMetricsTableSizeDayGrowthRatio `json:"TableSizeDayGrowthRatio,omitempty" xml:"TableSizeDayGrowthRatio,omitempty" type:"Struct"`
	// Warm data access days configuration.
	WarmConfigDay *GetDoctorHBaseTableResponseBodyDataMetricsWarmConfigDay `json:"WarmConfigDay,omitempty" xml:"WarmConfigDay,omitempty" type:"Struct"`
	// Warm data size.
	WarmDataSize *GetDoctorHBaseTableResponseBodyDataMetricsWarmDataSize `json:"WarmDataSize,omitempty" xml:"WarmDataSize,omitempty" type:"Struct"`
	// Write request balance.
	WriteRequestBalance *GetDoctorHBaseTableResponseBodyDataMetricsWriteRequestBalance `json:"WriteRequestBalance,omitempty" xml:"WriteRequestBalance,omitempty" type:"Struct"`
}

func (s GetDoctorHBaseTableResponseBodyDataMetrics) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseTableResponseBodyDataMetrics) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) GetColdAccessDay() *GetDoctorHBaseTableResponseBodyDataMetricsColdAccessDay {
	return s.ColdAccessDay
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) GetColdConfigDay() *GetDoctorHBaseTableResponseBodyDataMetricsColdConfigDay {
	return s.ColdConfigDay
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) GetColdDataSize() *GetDoctorHBaseTableResponseBodyDataMetricsColdDataSize {
	return s.ColdDataSize
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) GetDailyReadRequest() *GetDoctorHBaseTableResponseBodyDataMetricsDailyReadRequest {
	return s.DailyReadRequest
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) GetDailyReadRequestDayGrowthRatio() *GetDoctorHBaseTableResponseBodyDataMetricsDailyReadRequestDayGrowthRatio {
	return s.DailyReadRequestDayGrowthRatio
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) GetDailyWriteRequest() *GetDoctorHBaseTableResponseBodyDataMetricsDailyWriteRequest {
	return s.DailyWriteRequest
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) GetDailyWriteRequestDayGrowthRatio() *GetDoctorHBaseTableResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio {
	return s.DailyWriteRequestDayGrowthRatio
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) GetFreezeConfigDay() *GetDoctorHBaseTableResponseBodyDataMetricsFreezeConfigDay {
	return s.FreezeConfigDay
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) GetFreezeDataSize() *GetDoctorHBaseTableResponseBodyDataMetricsFreezeDataSize {
	return s.FreezeDataSize
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) GetHotDataSize() *GetDoctorHBaseTableResponseBodyDataMetricsHotDataSize {
	return s.HotDataSize
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) GetLocality() *GetDoctorHBaseTableResponseBodyDataMetricsLocality {
	return s.Locality
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) GetReadRequestBalance() *GetDoctorHBaseTableResponseBodyDataMetricsReadRequestBalance {
	return s.ReadRequestBalance
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) GetRegionBalance() *GetDoctorHBaseTableResponseBodyDataMetricsRegionBalance {
	return s.RegionBalance
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) GetRegionCount() *GetDoctorHBaseTableResponseBodyDataMetricsRegionCount {
	return s.RegionCount
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) GetRegionCountDayGrowthRatio() *GetDoctorHBaseTableResponseBodyDataMetricsRegionCountDayGrowthRatio {
	return s.RegionCountDayGrowthRatio
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) GetRegionServerCount() *GetDoctorHBaseTableResponseBodyDataMetricsRegionServerCount {
	return s.RegionServerCount
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) GetRequestBalance() *GetDoctorHBaseTableResponseBodyDataMetricsRequestBalance {
	return s.RequestBalance
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) GetStoreFileCount() *GetDoctorHBaseTableResponseBodyDataMetricsStoreFileCount {
	return s.StoreFileCount
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) GetStoreFileCountDayGrowthRatio() *GetDoctorHBaseTableResponseBodyDataMetricsStoreFileCountDayGrowthRatio {
	return s.StoreFileCountDayGrowthRatio
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) GetTableSize() *GetDoctorHBaseTableResponseBodyDataMetricsTableSize {
	return s.TableSize
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) GetTableSizeDayGrowthRatio() *GetDoctorHBaseTableResponseBodyDataMetricsTableSizeDayGrowthRatio {
	return s.TableSizeDayGrowthRatio
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) GetWarmConfigDay() *GetDoctorHBaseTableResponseBodyDataMetricsWarmConfigDay {
	return s.WarmConfigDay
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) GetWarmDataSize() *GetDoctorHBaseTableResponseBodyDataMetricsWarmDataSize {
	return s.WarmDataSize
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) GetWriteRequestBalance() *GetDoctorHBaseTableResponseBodyDataMetricsWriteRequestBalance {
	return s.WriteRequestBalance
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) SetColdAccessDay(v *GetDoctorHBaseTableResponseBodyDataMetricsColdAccessDay) *GetDoctorHBaseTableResponseBodyDataMetrics {
	s.ColdAccessDay = v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) SetColdConfigDay(v *GetDoctorHBaseTableResponseBodyDataMetricsColdConfigDay) *GetDoctorHBaseTableResponseBodyDataMetrics {
	s.ColdConfigDay = v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) SetColdDataSize(v *GetDoctorHBaseTableResponseBodyDataMetricsColdDataSize) *GetDoctorHBaseTableResponseBodyDataMetrics {
	s.ColdDataSize = v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) SetDailyReadRequest(v *GetDoctorHBaseTableResponseBodyDataMetricsDailyReadRequest) *GetDoctorHBaseTableResponseBodyDataMetrics {
	s.DailyReadRequest = v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) SetDailyReadRequestDayGrowthRatio(v *GetDoctorHBaseTableResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) *GetDoctorHBaseTableResponseBodyDataMetrics {
	s.DailyReadRequestDayGrowthRatio = v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) SetDailyWriteRequest(v *GetDoctorHBaseTableResponseBodyDataMetricsDailyWriteRequest) *GetDoctorHBaseTableResponseBodyDataMetrics {
	s.DailyWriteRequest = v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) SetDailyWriteRequestDayGrowthRatio(v *GetDoctorHBaseTableResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) *GetDoctorHBaseTableResponseBodyDataMetrics {
	s.DailyWriteRequestDayGrowthRatio = v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) SetFreezeConfigDay(v *GetDoctorHBaseTableResponseBodyDataMetricsFreezeConfigDay) *GetDoctorHBaseTableResponseBodyDataMetrics {
	s.FreezeConfigDay = v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) SetFreezeDataSize(v *GetDoctorHBaseTableResponseBodyDataMetricsFreezeDataSize) *GetDoctorHBaseTableResponseBodyDataMetrics {
	s.FreezeDataSize = v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) SetHotDataSize(v *GetDoctorHBaseTableResponseBodyDataMetricsHotDataSize) *GetDoctorHBaseTableResponseBodyDataMetrics {
	s.HotDataSize = v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) SetLocality(v *GetDoctorHBaseTableResponseBodyDataMetricsLocality) *GetDoctorHBaseTableResponseBodyDataMetrics {
	s.Locality = v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) SetReadRequestBalance(v *GetDoctorHBaseTableResponseBodyDataMetricsReadRequestBalance) *GetDoctorHBaseTableResponseBodyDataMetrics {
	s.ReadRequestBalance = v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) SetRegionBalance(v *GetDoctorHBaseTableResponseBodyDataMetricsRegionBalance) *GetDoctorHBaseTableResponseBodyDataMetrics {
	s.RegionBalance = v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) SetRegionCount(v *GetDoctorHBaseTableResponseBodyDataMetricsRegionCount) *GetDoctorHBaseTableResponseBodyDataMetrics {
	s.RegionCount = v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) SetRegionCountDayGrowthRatio(v *GetDoctorHBaseTableResponseBodyDataMetricsRegionCountDayGrowthRatio) *GetDoctorHBaseTableResponseBodyDataMetrics {
	s.RegionCountDayGrowthRatio = v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) SetRegionServerCount(v *GetDoctorHBaseTableResponseBodyDataMetricsRegionServerCount) *GetDoctorHBaseTableResponseBodyDataMetrics {
	s.RegionServerCount = v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) SetRequestBalance(v *GetDoctorHBaseTableResponseBodyDataMetricsRequestBalance) *GetDoctorHBaseTableResponseBodyDataMetrics {
	s.RequestBalance = v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) SetStoreFileCount(v *GetDoctorHBaseTableResponseBodyDataMetricsStoreFileCount) *GetDoctorHBaseTableResponseBodyDataMetrics {
	s.StoreFileCount = v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) SetStoreFileCountDayGrowthRatio(v *GetDoctorHBaseTableResponseBodyDataMetricsStoreFileCountDayGrowthRatio) *GetDoctorHBaseTableResponseBodyDataMetrics {
	s.StoreFileCountDayGrowthRatio = v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) SetTableSize(v *GetDoctorHBaseTableResponseBodyDataMetricsTableSize) *GetDoctorHBaseTableResponseBodyDataMetrics {
	s.TableSize = v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) SetTableSizeDayGrowthRatio(v *GetDoctorHBaseTableResponseBodyDataMetricsTableSizeDayGrowthRatio) *GetDoctorHBaseTableResponseBodyDataMetrics {
	s.TableSizeDayGrowthRatio = v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) SetWarmConfigDay(v *GetDoctorHBaseTableResponseBodyDataMetricsWarmConfigDay) *GetDoctorHBaseTableResponseBodyDataMetrics {
	s.WarmConfigDay = v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) SetWarmDataSize(v *GetDoctorHBaseTableResponseBodyDataMetricsWarmDataSize) *GetDoctorHBaseTableResponseBodyDataMetrics {
	s.WarmDataSize = v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) SetWriteRequestBalance(v *GetDoctorHBaseTableResponseBodyDataMetricsWriteRequestBalance) *GetDoctorHBaseTableResponseBodyDataMetrics {
	s.WriteRequestBalance = v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetrics) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseTableResponseBodyDataMetricsColdAccessDay struct {
	// Description of the metric.
	//
	// example:
	//
	// Cold access day
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Name of the metric.
	//
	// example:
	//
	// coldAccessDay
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Unit of the metric.
	//
	// example:
	//
	// day
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// Value of the metric.
	//
	// example:
	//
	// 3
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsColdAccessDay) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsColdAccessDay) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsColdAccessDay) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsColdAccessDay) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsColdAccessDay) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsColdAccessDay) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsColdAccessDay) SetDescription(v string) *GetDoctorHBaseTableResponseBodyDataMetricsColdAccessDay {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsColdAccessDay) SetName(v string) *GetDoctorHBaseTableResponseBodyDataMetricsColdAccessDay {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsColdAccessDay) SetUnit(v string) *GetDoctorHBaseTableResponseBodyDataMetricsColdAccessDay {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsColdAccessDay) SetValue(v int64) *GetDoctorHBaseTableResponseBodyDataMetricsColdAccessDay {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsColdAccessDay) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseTableResponseBodyDataMetricsColdConfigDay struct {
	// Description of the metric.
	//
	// example:
	//
	// Cold config day
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Metric name.
	//
	// example:
	//
	// coldConfigDay
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Unit of the metric.
	//
	// example:
	//
	// day
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// Metric value.
	//
	// example:
	//
	// 10
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsColdConfigDay) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsColdConfigDay) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsColdConfigDay) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsColdConfigDay) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsColdConfigDay) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsColdConfigDay) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsColdConfigDay) SetDescription(v string) *GetDoctorHBaseTableResponseBodyDataMetricsColdConfigDay {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsColdConfigDay) SetName(v string) *GetDoctorHBaseTableResponseBodyDataMetricsColdConfigDay {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsColdConfigDay) SetUnit(v string) *GetDoctorHBaseTableResponseBodyDataMetricsColdConfigDay {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsColdConfigDay) SetValue(v int64) *GetDoctorHBaseTableResponseBodyDataMetricsColdConfigDay {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsColdConfigDay) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseTableResponseBodyDataMetricsColdDataSize struct {
	// Metric description.
	//
	// example:
	//
	// Size of the cold data size
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Metric name.
	//
	// example:
	//
	// coldDataSize
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Metric unit.
	//
	// example:
	//
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// Metric value.
	//
	// example:
	//
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsColdDataSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsColdDataSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsColdDataSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsColdDataSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsColdDataSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsColdDataSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsColdDataSize) SetDescription(v string) *GetDoctorHBaseTableResponseBodyDataMetricsColdDataSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsColdDataSize) SetName(v string) *GetDoctorHBaseTableResponseBodyDataMetricsColdDataSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsColdDataSize) SetUnit(v string) *GetDoctorHBaseTableResponseBodyDataMetricsColdDataSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsColdDataSize) SetValue(v int64) *GetDoctorHBaseTableResponseBodyDataMetricsColdDataSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsColdDataSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseTableResponseBodyDataMetricsDailyReadRequest struct {
	// Description of the metric.
	//
	// example:
	//
	// test-update
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Name of the item.
	//
	// example:
	//
	// dailyReadRequest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Unit of the metric.
	//
	// example:
	//
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// Value of the metric.
	//
	// example:
	//
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsDailyReadRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsDailyReadRequest) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsDailyReadRequest) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsDailyReadRequest) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsDailyReadRequest) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsDailyReadRequest) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsDailyReadRequest) SetDescription(v string) *GetDoctorHBaseTableResponseBodyDataMetricsDailyReadRequest {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsDailyReadRequest) SetName(v string) *GetDoctorHBaseTableResponseBodyDataMetricsDailyReadRequest {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsDailyReadRequest) SetUnit(v string) *GetDoctorHBaseTableResponseBodyDataMetricsDailyReadRequest {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsDailyReadRequest) SetValue(v int64) *GetDoctorHBaseTableResponseBodyDataMetricsDailyReadRequest {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsDailyReadRequest) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseTableResponseBodyDataMetricsDailyReadRequestDayGrowthRatio struct {
	// Description of the metric.
	//
	// example:
	//
	// Day growth ratio of table size
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Metric name.
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
	// 1.5
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) SetDescription(v string) *GetDoctorHBaseTableResponseBodyDataMetricsDailyReadRequestDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) SetName(v string) *GetDoctorHBaseTableResponseBodyDataMetricsDailyReadRequestDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) SetUnit(v string) *GetDoctorHBaseTableResponseBodyDataMetricsDailyReadRequestDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) SetValue(v float32) *GetDoctorHBaseTableResponseBodyDataMetricsDailyReadRequestDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseTableResponseBodyDataMetricsDailyWriteRequest struct {
	// Description of the metric.
	//
	// example:
	//
	// Number of write requests per day
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Name of the metric.
	//
	// example:
	//
	// dailyWriteRequest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Unit of the metric.
	//
	// example:
	//
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// Value of the metric.
	//
	// example:
	//
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsDailyWriteRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsDailyWriteRequest) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsDailyWriteRequest) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsDailyWriteRequest) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsDailyWriteRequest) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsDailyWriteRequest) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsDailyWriteRequest) SetDescription(v string) *GetDoctorHBaseTableResponseBodyDataMetricsDailyWriteRequest {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsDailyWriteRequest) SetName(v string) *GetDoctorHBaseTableResponseBodyDataMetricsDailyWriteRequest {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsDailyWriteRequest) SetUnit(v string) *GetDoctorHBaseTableResponseBodyDataMetricsDailyWriteRequest {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsDailyWriteRequest) SetValue(v int64) *GetDoctorHBaseTableResponseBodyDataMetricsDailyWriteRequest {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsDailyWriteRequest) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseTableResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio struct {
	// Description of the metric.
	//
	// example:
	//
	// The balance of distributing requests
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Name of the metric.
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
	// 0.5
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) SetDescription(v string) *GetDoctorHBaseTableResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) SetName(v string) *GetDoctorHBaseTableResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) SetUnit(v string) *GetDoctorHBaseTableResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) SetValue(v float32) *GetDoctorHBaseTableResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseTableResponseBodyDataMetricsFreezeConfigDay struct {
	// Description of the metric.
	//
	// example:
	//
	// Freeze config day
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Name of the metric.
	//
	// example:
	//
	// freezeConfigDay
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Unit of the metric.
	//
	// example:
	//
	// day
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// Value of the metric.
	//
	// example:
	//
	// 10
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsFreezeConfigDay) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsFreezeConfigDay) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsFreezeConfigDay) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsFreezeConfigDay) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsFreezeConfigDay) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsFreezeConfigDay) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsFreezeConfigDay) SetDescription(v string) *GetDoctorHBaseTableResponseBodyDataMetricsFreezeConfigDay {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsFreezeConfigDay) SetName(v string) *GetDoctorHBaseTableResponseBodyDataMetricsFreezeConfigDay {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsFreezeConfigDay) SetUnit(v string) *GetDoctorHBaseTableResponseBodyDataMetricsFreezeConfigDay {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsFreezeConfigDay) SetValue(v int64) *GetDoctorHBaseTableResponseBodyDataMetricsFreezeConfigDay {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsFreezeConfigDay) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseTableResponseBodyDataMetricsFreezeDataSize struct {
	// Metric description.
	//
	// example:
	//
	// Size of the freeze data size
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Metric name.
	//
	// example:
	//
	// freezeDataSize
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Metric unit.
	//
	// example:
	//
	// MB
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// Metric value.
	//
	// example:
	//
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsFreezeDataSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsFreezeDataSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsFreezeDataSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsFreezeDataSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsFreezeDataSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsFreezeDataSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsFreezeDataSize) SetDescription(v string) *GetDoctorHBaseTableResponseBodyDataMetricsFreezeDataSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsFreezeDataSize) SetName(v string) *GetDoctorHBaseTableResponseBodyDataMetricsFreezeDataSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsFreezeDataSize) SetUnit(v string) *GetDoctorHBaseTableResponseBodyDataMetricsFreezeDataSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsFreezeDataSize) SetValue(v int64) *GetDoctorHBaseTableResponseBodyDataMetricsFreezeDataSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsFreezeDataSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseTableResponseBodyDataMetricsHotDataSize struct {
	// Description of the metric.
	//
	// example:
	//
	// Size of the hot data size
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Name of the metric.
	//
	// example:
	//
	// hotDataSize
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// MB
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The metric value.
	//
	// example:
	//
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsHotDataSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsHotDataSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsHotDataSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsHotDataSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsHotDataSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsHotDataSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsHotDataSize) SetDescription(v string) *GetDoctorHBaseTableResponseBodyDataMetricsHotDataSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsHotDataSize) SetName(v string) *GetDoctorHBaseTableResponseBodyDataMetricsHotDataSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsHotDataSize) SetUnit(v string) *GetDoctorHBaseTableResponseBodyDataMetricsHotDataSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsHotDataSize) SetValue(v int64) *GetDoctorHBaseTableResponseBodyDataMetricsHotDataSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsHotDataSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseTableResponseBodyDataMetricsLocality struct {
	// Description of the metric.
	//
	// example:
	//
	// The locality of data
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Name of the metric.
	//
	// example:
	//
	// locality
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Unit of the metric.
	//
	// example:
	//
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// Value of the metric.
	//
	// example:
	//
	// 0.5
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsLocality) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsLocality) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsLocality) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsLocality) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsLocality) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsLocality) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsLocality) SetDescription(v string) *GetDoctorHBaseTableResponseBodyDataMetricsLocality {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsLocality) SetName(v string) *GetDoctorHBaseTableResponseBodyDataMetricsLocality {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsLocality) SetUnit(v string) *GetDoctorHBaseTableResponseBodyDataMetricsLocality {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsLocality) SetValue(v float32) *GetDoctorHBaseTableResponseBodyDataMetricsLocality {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsLocality) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseTableResponseBodyDataMetricsReadRequestBalance struct {
	// Description of the metric.
	//
	// example:
	//
	// The balance of distributing read requests
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Name of the metric.
	//
	// example:
	//
	// readRequestBalance
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
	// 0.5
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsReadRequestBalance) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsReadRequestBalance) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsReadRequestBalance) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsReadRequestBalance) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsReadRequestBalance) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsReadRequestBalance) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsReadRequestBalance) SetDescription(v string) *GetDoctorHBaseTableResponseBodyDataMetricsReadRequestBalance {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsReadRequestBalance) SetName(v string) *GetDoctorHBaseTableResponseBodyDataMetricsReadRequestBalance {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsReadRequestBalance) SetUnit(v string) *GetDoctorHBaseTableResponseBodyDataMetricsReadRequestBalance {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsReadRequestBalance) SetValue(v float32) *GetDoctorHBaseTableResponseBodyDataMetricsReadRequestBalance {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsReadRequestBalance) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseTableResponseBodyDataMetricsRegionBalance struct {
	// Description of the metric.
	//
	// example:
	//
	// The ability to evenly distribute Regions on different RegionServer nodes
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Metric name.
	//
	// example:
	//
	// regionBalance
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The metric value.
	//
	// example:
	//
	// 1.0
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsRegionBalance) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsRegionBalance) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRegionBalance) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRegionBalance) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRegionBalance) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRegionBalance) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRegionBalance) SetDescription(v string) *GetDoctorHBaseTableResponseBodyDataMetricsRegionBalance {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRegionBalance) SetName(v string) *GetDoctorHBaseTableResponseBodyDataMetricsRegionBalance {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRegionBalance) SetUnit(v string) *GetDoctorHBaseTableResponseBodyDataMetricsRegionBalance {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRegionBalance) SetValue(v float32) *GetDoctorHBaseTableResponseBodyDataMetricsRegionBalance {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRegionBalance) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseTableResponseBodyDataMetricsRegionCount struct {
	// Metric description.
	//
	// example:
	//
	// Number of regions count
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Metric name.
	//
	// example:
	//
	// regionCount
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Metric unit.
	//
	// example:
	//
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// Metric value.
	//
	// example:
	//
	// 10
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsRegionCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsRegionCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRegionCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRegionCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRegionCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRegionCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRegionCount) SetDescription(v string) *GetDoctorHBaseTableResponseBodyDataMetricsRegionCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRegionCount) SetName(v string) *GetDoctorHBaseTableResponseBodyDataMetricsRegionCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRegionCount) SetUnit(v string) *GetDoctorHBaseTableResponseBodyDataMetricsRegionCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRegionCount) SetValue(v int64) *GetDoctorHBaseTableResponseBodyDataMetricsRegionCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRegionCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseTableResponseBodyDataMetricsRegionCountDayGrowthRatio struct {
	// Metric description.
	//
	// example:
	//
	// Day growth ratio of region count
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Metric name.
	//
	// example:
	//
	// regionCountDayGrowthRatio
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Metric unit.
	//
	// example:
	//
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// Metric value.
	//
	// example:
	//
	// 0.8
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsRegionCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsRegionCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRegionCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRegionCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRegionCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRegionCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRegionCountDayGrowthRatio) SetDescription(v string) *GetDoctorHBaseTableResponseBodyDataMetricsRegionCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRegionCountDayGrowthRatio) SetName(v string) *GetDoctorHBaseTableResponseBodyDataMetricsRegionCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRegionCountDayGrowthRatio) SetUnit(v string) *GetDoctorHBaseTableResponseBodyDataMetricsRegionCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRegionCountDayGrowthRatio) SetValue(v float32) *GetDoctorHBaseTableResponseBodyDataMetricsRegionCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRegionCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseTableResponseBodyDataMetricsRegionServerCount struct {
	// Metric description.
	//
	// example:
	//
	// Number of region servers count
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Metric name.
	//
	// example:
	//
	// regionServerCount
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Metric unit.
	//
	// example:
	//
	// “”
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// Usage.
	//
	// example:
	//
	// 10
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsRegionServerCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsRegionServerCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRegionServerCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRegionServerCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRegionServerCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRegionServerCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRegionServerCount) SetDescription(v string) *GetDoctorHBaseTableResponseBodyDataMetricsRegionServerCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRegionServerCount) SetName(v string) *GetDoctorHBaseTableResponseBodyDataMetricsRegionServerCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRegionServerCount) SetUnit(v string) *GetDoctorHBaseTableResponseBodyDataMetricsRegionServerCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRegionServerCount) SetValue(v int64) *GetDoctorHBaseTableResponseBodyDataMetricsRegionServerCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRegionServerCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseTableResponseBodyDataMetricsRequestBalance struct {
	// Description of the metric.
	//
	// example:
	//
	// The balance of distributing requests
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Name of the metric.
	//
	// example:
	//
	// requestBalance
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The metric value.
	//
	// example:
	//
	// 1.0
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsRequestBalance) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsRequestBalance) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRequestBalance) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRequestBalance) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRequestBalance) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRequestBalance) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRequestBalance) SetDescription(v string) *GetDoctorHBaseTableResponseBodyDataMetricsRequestBalance {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRequestBalance) SetName(v string) *GetDoctorHBaseTableResponseBodyDataMetricsRequestBalance {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRequestBalance) SetUnit(v string) *GetDoctorHBaseTableResponseBodyDataMetricsRequestBalance {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRequestBalance) SetValue(v float32) *GetDoctorHBaseTableResponseBodyDataMetricsRequestBalance {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsRequestBalance) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseTableResponseBodyDataMetricsStoreFileCount struct {
	// Description of the metric.
	//
	// example:
	//
	// Number of store files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Name of the metric.
	//
	// example:
	//
	// storeFileCount
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Unit of the metric.
	//
	// example:
	//
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// Value of the metric.
	//
	// example:
	//
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsStoreFileCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsStoreFileCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsStoreFileCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsStoreFileCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsStoreFileCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsStoreFileCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsStoreFileCount) SetDescription(v string) *GetDoctorHBaseTableResponseBodyDataMetricsStoreFileCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsStoreFileCount) SetName(v string) *GetDoctorHBaseTableResponseBodyDataMetricsStoreFileCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsStoreFileCount) SetUnit(v string) *GetDoctorHBaseTableResponseBodyDataMetricsStoreFileCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsStoreFileCount) SetValue(v int64) *GetDoctorHBaseTableResponseBodyDataMetricsStoreFileCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsStoreFileCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseTableResponseBodyDataMetricsStoreFileCountDayGrowthRatio struct {
	// Metric description.
	//
	// example:
	//
	// Day growth ratio of store file count
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Metric name.
	//
	// example:
	//
	// storeFileCountDayGrowthRatio
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Metric unit.
	//
	// example:
	//
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// Metric value.
	//
	// example:
	//
	// 1.5
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsStoreFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsStoreFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsStoreFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsStoreFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsStoreFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsStoreFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsStoreFileCountDayGrowthRatio) SetDescription(v string) *GetDoctorHBaseTableResponseBodyDataMetricsStoreFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsStoreFileCountDayGrowthRatio) SetName(v string) *GetDoctorHBaseTableResponseBodyDataMetricsStoreFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsStoreFileCountDayGrowthRatio) SetUnit(v string) *GetDoctorHBaseTableResponseBodyDataMetricsStoreFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsStoreFileCountDayGrowthRatio) SetValue(v float32) *GetDoctorHBaseTableResponseBodyDataMetricsStoreFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsStoreFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseTableResponseBodyDataMetricsTableSize struct {
	// Description of the metric.
	//
	// example:
	//
	// Size of the table
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Metric name.
	//
	// example:
	//
	// tableSize
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Unit of the metric
	//
	// example:
	//
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// Metric value.
	//
	// example:
	//
	// tb_item
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsTableSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsTableSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsTableSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsTableSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsTableSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsTableSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsTableSize) SetDescription(v string) *GetDoctorHBaseTableResponseBodyDataMetricsTableSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsTableSize) SetName(v string) *GetDoctorHBaseTableResponseBodyDataMetricsTableSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsTableSize) SetUnit(v string) *GetDoctorHBaseTableResponseBodyDataMetricsTableSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsTableSize) SetValue(v int64) *GetDoctorHBaseTableResponseBodyDataMetricsTableSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsTableSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseTableResponseBodyDataMetricsTableSizeDayGrowthRatio struct {
	// Metric description.
	//
	// example:
	//
	// Day growth ratio of table size
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Metric name.
	//
	// example:
	//
	// tableSizeDayGrowthRatio
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Metric unit.
	//
	// example:
	//
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// Metric value.
	//
	// example:
	//
	// 1.5
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsTableSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsTableSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsTableSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsTableSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsTableSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsTableSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsTableSizeDayGrowthRatio) SetDescription(v string) *GetDoctorHBaseTableResponseBodyDataMetricsTableSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsTableSizeDayGrowthRatio) SetName(v string) *GetDoctorHBaseTableResponseBodyDataMetricsTableSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsTableSizeDayGrowthRatio) SetUnit(v string) *GetDoctorHBaseTableResponseBodyDataMetricsTableSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsTableSizeDayGrowthRatio) SetValue(v float32) *GetDoctorHBaseTableResponseBodyDataMetricsTableSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsTableSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseTableResponseBodyDataMetricsWarmConfigDay struct {
	// Metric description.
	//
	// example:
	//
	// Size of the warm data size
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Metric name.
	//
	// example:
	//
	// warmConfigDay
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Metric unit.
	//
	// example:
	//
	// MB
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// Metric value.
	//
	// example:
	//
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsWarmConfigDay) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsWarmConfigDay) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsWarmConfigDay) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsWarmConfigDay) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsWarmConfigDay) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsWarmConfigDay) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsWarmConfigDay) SetDescription(v string) *GetDoctorHBaseTableResponseBodyDataMetricsWarmConfigDay {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsWarmConfigDay) SetName(v string) *GetDoctorHBaseTableResponseBodyDataMetricsWarmConfigDay {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsWarmConfigDay) SetUnit(v string) *GetDoctorHBaseTableResponseBodyDataMetricsWarmConfigDay {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsWarmConfigDay) SetValue(v int64) *GetDoctorHBaseTableResponseBodyDataMetricsWarmConfigDay {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsWarmConfigDay) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseTableResponseBodyDataMetricsWarmDataSize struct {
	// Description of the metric.
	//
	// example:
	//
	// Size of the warm data size
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Name of the metric.
	//
	// example:
	//
	// warmDataSize
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// MB
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// Usage rate.
	//
	// example:
	//
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsWarmDataSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsWarmDataSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsWarmDataSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsWarmDataSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsWarmDataSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsWarmDataSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsWarmDataSize) SetDescription(v string) *GetDoctorHBaseTableResponseBodyDataMetricsWarmDataSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsWarmDataSize) SetName(v string) *GetDoctorHBaseTableResponseBodyDataMetricsWarmDataSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsWarmDataSize) SetUnit(v string) *GetDoctorHBaseTableResponseBodyDataMetricsWarmDataSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsWarmDataSize) SetValue(v int64) *GetDoctorHBaseTableResponseBodyDataMetricsWarmDataSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsWarmDataSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseTableResponseBodyDataMetricsWriteRequestBalance struct {
	// Description of the metric.
	//
	// example:
	//
	// The balance of distributing write requests
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Name of the metric.
	//
	// example:
	//
	// writeRequestBalance
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Unit of the metric.
	//
	// example:
	//
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 0.5
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsWriteRequestBalance) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseTableResponseBodyDataMetricsWriteRequestBalance) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsWriteRequestBalance) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsWriteRequestBalance) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsWriteRequestBalance) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsWriteRequestBalance) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsWriteRequestBalance) SetDescription(v string) *GetDoctorHBaseTableResponseBodyDataMetricsWriteRequestBalance {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsWriteRequestBalance) SetName(v string) *GetDoctorHBaseTableResponseBodyDataMetricsWriteRequestBalance {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsWriteRequestBalance) SetUnit(v string) *GetDoctorHBaseTableResponseBodyDataMetricsWriteRequestBalance {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsWriteRequestBalance) SetValue(v float32) *GetDoctorHBaseTableResponseBodyDataMetricsWriteRequestBalance {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseTableResponseBodyDataMetricsWriteRequestBalance) Validate() error {
	return dara.Validate(s)
}
