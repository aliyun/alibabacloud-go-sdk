// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDoctorHBaseTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListDoctorHBaseTablesResponseBodyData) *ListDoctorHBaseTablesResponseBody
	GetData() []*ListDoctorHBaseTablesResponseBodyData
	SetMaxResults(v int32) *ListDoctorHBaseTablesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDoctorHBaseTablesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDoctorHBaseTablesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListDoctorHBaseTablesResponseBody
	GetTotalCount() *int32
}

type ListDoctorHBaseTablesResponseBody struct {
	// The data returned.
	Data []*ListDoctorHBaseTablesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The maximum number of entries returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The page number of the next page returned.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
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

func (s ListDoctorHBaseTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseTablesResponseBody) GetData() []*ListDoctorHBaseTablesResponseBodyData {
	return s.Data
}

func (s *ListDoctorHBaseTablesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDoctorHBaseTablesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDoctorHBaseTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDoctorHBaseTablesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDoctorHBaseTablesResponseBody) SetData(v []*ListDoctorHBaseTablesResponseBodyData) *ListDoctorHBaseTablesResponseBody {
	s.Data = v
	return s
}

func (s *ListDoctorHBaseTablesResponseBody) SetMaxResults(v int32) *ListDoctorHBaseTablesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBody) SetNextToken(v string) *ListDoctorHBaseTablesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBody) SetRequestId(v string) *ListDoctorHBaseTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBody) SetTotalCount(v int32) *ListDoctorHBaseTablesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHBaseTablesResponseBodyData struct {
	// The diagnosis result.
	Analysis *ListDoctorHBaseTablesResponseBodyDataAnalysis `json:"Analysis,omitempty" xml:"Analysis,omitempty" type:"Struct"`
	// The metric information.
	Metrics *ListDoctorHBaseTablesResponseBodyDataMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Struct"`
	// The name of the table.
	//
	// example:
	//
	// tb_item
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ListDoctorHBaseTablesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseTablesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseTablesResponseBodyData) GetAnalysis() *ListDoctorHBaseTablesResponseBodyDataAnalysis {
	return s.Analysis
}

func (s *ListDoctorHBaseTablesResponseBodyData) GetMetrics() *ListDoctorHBaseTablesResponseBodyDataMetrics {
	return s.Metrics
}

func (s *ListDoctorHBaseTablesResponseBodyData) GetTableName() *string {
	return s.TableName
}

func (s *ListDoctorHBaseTablesResponseBodyData) SetAnalysis(v *ListDoctorHBaseTablesResponseBodyDataAnalysis) *ListDoctorHBaseTablesResponseBodyData {
	s.Analysis = v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyData) SetMetrics(v *ListDoctorHBaseTablesResponseBodyDataMetrics) *ListDoctorHBaseTablesResponseBodyData {
	s.Metrics = v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyData) SetTableName(v string) *ListDoctorHBaseTablesResponseBodyData {
	s.TableName = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHBaseTablesResponseBodyDataAnalysis struct {
	// The regions that have read hotspot issues.
	//
	// example:
	//
	// null
	ReadRequestHotspotRegionList []*string `json:"ReadRequestHotspotRegionList,omitempty" xml:"ReadRequestHotspotRegionList,omitempty" type:"Repeated"`
	// The description of read imbalance.
	//
	// example:
	//
	// read request unbalance is <p class=\\"report-detail-topic\\">
	ReadRequestUnbalanceSuggestion *string `json:"ReadRequestUnbalanceSuggestion,omitempty" xml:"ReadRequestUnbalanceSuggestion,omitempty"`
	// The regions that have read/write hotspot issues.
	//
	// example:
	//
	// null
	RequestHotspotRegionList []*string `json:"RequestHotspotRegionList,omitempty" xml:"RequestHotspotRegionList,omitempty" type:"Repeated"`
	// The description of read/write imbalance.
	//
	// example:
	//
	// read request unbalance is <p class=\\"report-detail-topic\\">
	RequestUnbalanceSuggestion *string `json:"RequestUnbalanceSuggestion,omitempty" xml:"RequestUnbalanceSuggestion,omitempty"`
	// The score of the table.
	//
	// example:
	//
	// 67
	TableScore *int32 `json:"TableScore,omitempty" xml:"TableScore,omitempty"`
	// The regions that have write hotspot issues.
	//
	// example:
	//
	// null
	WriteRequestHotspotRegionList []*string `json:"WriteRequestHotspotRegionList,omitempty" xml:"WriteRequestHotspotRegionList,omitempty" type:"Repeated"`
	// The description of write imbalance.
	//
	// example:
	//
	// write request unbalance is <p class=\\"report-detail-topic\\">
	WriteRequestUnbalanceSuggestion *string `json:"WriteRequestUnbalanceSuggestion,omitempty" xml:"WriteRequestUnbalanceSuggestion,omitempty"`
}

func (s ListDoctorHBaseTablesResponseBodyDataAnalysis) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseTablesResponseBodyDataAnalysis) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseTablesResponseBodyDataAnalysis) GetReadRequestHotspotRegionList() []*string {
	return s.ReadRequestHotspotRegionList
}

func (s *ListDoctorHBaseTablesResponseBodyDataAnalysis) GetReadRequestUnbalanceSuggestion() *string {
	return s.ReadRequestUnbalanceSuggestion
}

func (s *ListDoctorHBaseTablesResponseBodyDataAnalysis) GetRequestHotspotRegionList() []*string {
	return s.RequestHotspotRegionList
}

func (s *ListDoctorHBaseTablesResponseBodyDataAnalysis) GetRequestUnbalanceSuggestion() *string {
	return s.RequestUnbalanceSuggestion
}

func (s *ListDoctorHBaseTablesResponseBodyDataAnalysis) GetTableScore() *int32 {
	return s.TableScore
}

func (s *ListDoctorHBaseTablesResponseBodyDataAnalysis) GetWriteRequestHotspotRegionList() []*string {
	return s.WriteRequestHotspotRegionList
}

func (s *ListDoctorHBaseTablesResponseBodyDataAnalysis) GetWriteRequestUnbalanceSuggestion() *string {
	return s.WriteRequestUnbalanceSuggestion
}

func (s *ListDoctorHBaseTablesResponseBodyDataAnalysis) SetReadRequestHotspotRegionList(v []*string) *ListDoctorHBaseTablesResponseBodyDataAnalysis {
	s.ReadRequestHotspotRegionList = v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataAnalysis) SetReadRequestUnbalanceSuggestion(v string) *ListDoctorHBaseTablesResponseBodyDataAnalysis {
	s.ReadRequestUnbalanceSuggestion = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataAnalysis) SetRequestHotspotRegionList(v []*string) *ListDoctorHBaseTablesResponseBodyDataAnalysis {
	s.RequestHotspotRegionList = v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataAnalysis) SetRequestUnbalanceSuggestion(v string) *ListDoctorHBaseTablesResponseBodyDataAnalysis {
	s.RequestUnbalanceSuggestion = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataAnalysis) SetTableScore(v int32) *ListDoctorHBaseTablesResponseBodyDataAnalysis {
	s.TableScore = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataAnalysis) SetWriteRequestHotspotRegionList(v []*string) *ListDoctorHBaseTablesResponseBodyDataAnalysis {
	s.WriteRequestHotspotRegionList = v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataAnalysis) SetWriteRequestUnbalanceSuggestion(v string) *ListDoctorHBaseTablesResponseBodyDataAnalysis {
	s.WriteRequestUnbalanceSuggestion = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataAnalysis) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHBaseTablesResponseBodyDataMetrics struct {
	// The number of days during which the table was not accessed.
	ColdAccessDay *ListDoctorHBaseTablesResponseBodyDataMetricsColdAccessDay `json:"ColdAccessDay,omitempty" xml:"ColdAccessDay,omitempty" type:"Struct"`
	// The number of consecutive days without access to data before the data is considered as very cold data.
	ColdConfigDay *ListDoctorHBaseTablesResponseBodyDataMetricsColdConfigDay `json:"ColdConfigDay,omitempty" xml:"ColdConfigDay,omitempty" type:"Struct"`
	// The size of cold data.
	ColdDataSize *ListDoctorHBaseTablesResponseBodyDataMetricsColdDataSize `json:"ColdDataSize,omitempty" xml:"ColdDataSize,omitempty" type:"Struct"`
	// The total number of read requests for the table in a day.
	DailyReadRequest *ListDoctorHBaseTablesResponseBodyDataMetricsDailyReadRequest `json:"DailyReadRequest,omitempty" xml:"DailyReadRequest,omitempty" type:"Struct"`
	// The daily increment ratio of the number of read requests in a day.
	DailyReadRequestDayGrowthRatio *ListDoctorHBaseTablesResponseBodyDataMetricsDailyReadRequestDayGrowthRatio `json:"DailyReadRequestDayGrowthRatio,omitempty" xml:"DailyReadRequestDayGrowthRatio,omitempty" type:"Struct"`
	// The total number of write requests for the table in a day.
	DailyWriteRequest *ListDoctorHBaseTablesResponseBodyDataMetricsDailyWriteRequest `json:"DailyWriteRequest,omitempty" xml:"DailyWriteRequest,omitempty" type:"Struct"`
	// The daily increment ratio of the number of write requests in a day.
	DailyWriteRequestDayGrowthRatio *ListDoctorHBaseTablesResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio `json:"DailyWriteRequestDayGrowthRatio,omitempty" xml:"DailyWriteRequestDayGrowthRatio,omitempty" type:"Struct"`
	// The number of consecutive days without access to data before the data was considered as very cold data.
	FreezeConfigDay *ListDoctorHBaseTablesResponseBodyDataMetricsFreezeConfigDay `json:"FreezeConfigDay,omitempty" xml:"FreezeConfigDay,omitempty" type:"Struct"`
	// The size of very cold data.
	FreezeDataSize *ListDoctorHBaseTablesResponseBodyDataMetricsFreezeDataSize `json:"FreezeDataSize,omitempty" xml:"FreezeDataSize,omitempty" type:"Struct"`
	// The size of hot data.
	HotDataSize *ListDoctorHBaseTablesResponseBodyDataMetricsHotDataSize `json:"HotDataSize,omitempty" xml:"HotDataSize,omitempty" type:"Struct"`
	// The localization rate.
	Locality *ListDoctorHBaseTablesResponseBodyDataMetricsLocality `json:"Locality,omitempty" xml:"Locality,omitempty" type:"Struct"`
	// The read balancing degree.
	ReadRequestBalance *ListDoctorHBaseTablesResponseBodyDataMetricsReadRequestBalance `json:"ReadRequestBalance,omitempty" xml:"ReadRequestBalance,omitempty" type:"Struct"`
	// The balancing degree.
	RegionBalance *ListDoctorHBaseTablesResponseBodyDataMetricsRegionBalance `json:"RegionBalance,omitempty" xml:"RegionBalance,omitempty" type:"Struct"`
	// The number of regions that host the table.
	RegionCount *ListDoctorHBaseTablesResponseBodyDataMetricsRegionCount `json:"RegionCount,omitempty" xml:"RegionCount,omitempty" type:"Struct"`
	// The daily increment ratio of the number of regions.
	RegionCountDayGrowthRatio *ListDoctorHBaseTablesResponseBodyDataMetricsRegionCountDayGrowthRatio `json:"RegionCountDayGrowthRatio,omitempty" xml:"RegionCountDayGrowthRatio,omitempty" type:"Struct"`
	// The number of region servers that host the table.
	RegionServerCount *ListDoctorHBaseTablesResponseBodyDataMetricsRegionServerCount `json:"RegionServerCount,omitempty" xml:"RegionServerCount,omitempty" type:"Struct"`
	// The request balancing degree.
	RequestBalance *ListDoctorHBaseTablesResponseBodyDataMetricsRequestBalance `json:"RequestBalance,omitempty" xml:"RequestBalance,omitempty" type:"Struct"`
	// The number of StoreFiles.
	StoreFileCount *ListDoctorHBaseTablesResponseBodyDataMetricsStoreFileCount `json:"StoreFileCount,omitempty" xml:"StoreFileCount,omitempty" type:"Struct"`
	// The daily increment ratio of the number of StoreFiles.
	StoreFileCountDayGrowthRatio *ListDoctorHBaseTablesResponseBodyDataMetricsStoreFileCountDayGrowthRatio `json:"StoreFileCountDayGrowthRatio,omitempty" xml:"StoreFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The size of the table.
	TableSize *ListDoctorHBaseTablesResponseBodyDataMetricsTableSize `json:"TableSize,omitempty" xml:"TableSize,omitempty" type:"Struct"`
	// The daily increment ratio of the table size.
	TableSizeDayGrowthRatio *ListDoctorHBaseTablesResponseBodyDataMetricsTableSizeDayGrowthRatio `json:"TableSizeDayGrowthRatio,omitempty" xml:"TableSizeDayGrowthRatio,omitempty" type:"Struct"`
	// The number of consecutive days without access to data before the data is considered as cold data.
	WarmConfigDay *ListDoctorHBaseTablesResponseBodyDataMetricsWarmConfigDay `json:"WarmConfigDay,omitempty" xml:"WarmConfigDay,omitempty" type:"Struct"`
	// The size of warm data.
	WarmDataSize *ListDoctorHBaseTablesResponseBodyDataMetricsWarmDataSize `json:"WarmDataSize,omitempty" xml:"WarmDataSize,omitempty" type:"Struct"`
	// The write balancing degree.
	WriteRequestBalance *ListDoctorHBaseTablesResponseBodyDataMetricsWriteRequestBalance `json:"WriteRequestBalance,omitempty" xml:"WriteRequestBalance,omitempty" type:"Struct"`
}

func (s ListDoctorHBaseTablesResponseBodyDataMetrics) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseTablesResponseBodyDataMetrics) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) GetColdAccessDay() *ListDoctorHBaseTablesResponseBodyDataMetricsColdAccessDay {
	return s.ColdAccessDay
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) GetColdConfigDay() *ListDoctorHBaseTablesResponseBodyDataMetricsColdConfigDay {
	return s.ColdConfigDay
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) GetColdDataSize() *ListDoctorHBaseTablesResponseBodyDataMetricsColdDataSize {
	return s.ColdDataSize
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) GetDailyReadRequest() *ListDoctorHBaseTablesResponseBodyDataMetricsDailyReadRequest {
	return s.DailyReadRequest
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) GetDailyReadRequestDayGrowthRatio() *ListDoctorHBaseTablesResponseBodyDataMetricsDailyReadRequestDayGrowthRatio {
	return s.DailyReadRequestDayGrowthRatio
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) GetDailyWriteRequest() *ListDoctorHBaseTablesResponseBodyDataMetricsDailyWriteRequest {
	return s.DailyWriteRequest
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) GetDailyWriteRequestDayGrowthRatio() *ListDoctorHBaseTablesResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio {
	return s.DailyWriteRequestDayGrowthRatio
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) GetFreezeConfigDay() *ListDoctorHBaseTablesResponseBodyDataMetricsFreezeConfigDay {
	return s.FreezeConfigDay
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) GetFreezeDataSize() *ListDoctorHBaseTablesResponseBodyDataMetricsFreezeDataSize {
	return s.FreezeDataSize
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) GetHotDataSize() *ListDoctorHBaseTablesResponseBodyDataMetricsHotDataSize {
	return s.HotDataSize
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) GetLocality() *ListDoctorHBaseTablesResponseBodyDataMetricsLocality {
	return s.Locality
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) GetReadRequestBalance() *ListDoctorHBaseTablesResponseBodyDataMetricsReadRequestBalance {
	return s.ReadRequestBalance
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) GetRegionBalance() *ListDoctorHBaseTablesResponseBodyDataMetricsRegionBalance {
	return s.RegionBalance
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) GetRegionCount() *ListDoctorHBaseTablesResponseBodyDataMetricsRegionCount {
	return s.RegionCount
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) GetRegionCountDayGrowthRatio() *ListDoctorHBaseTablesResponseBodyDataMetricsRegionCountDayGrowthRatio {
	return s.RegionCountDayGrowthRatio
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) GetRegionServerCount() *ListDoctorHBaseTablesResponseBodyDataMetricsRegionServerCount {
	return s.RegionServerCount
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) GetRequestBalance() *ListDoctorHBaseTablesResponseBodyDataMetricsRequestBalance {
	return s.RequestBalance
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) GetStoreFileCount() *ListDoctorHBaseTablesResponseBodyDataMetricsStoreFileCount {
	return s.StoreFileCount
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) GetStoreFileCountDayGrowthRatio() *ListDoctorHBaseTablesResponseBodyDataMetricsStoreFileCountDayGrowthRatio {
	return s.StoreFileCountDayGrowthRatio
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) GetTableSize() *ListDoctorHBaseTablesResponseBodyDataMetricsTableSize {
	return s.TableSize
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) GetTableSizeDayGrowthRatio() *ListDoctorHBaseTablesResponseBodyDataMetricsTableSizeDayGrowthRatio {
	return s.TableSizeDayGrowthRatio
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) GetWarmConfigDay() *ListDoctorHBaseTablesResponseBodyDataMetricsWarmConfigDay {
	return s.WarmConfigDay
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) GetWarmDataSize() *ListDoctorHBaseTablesResponseBodyDataMetricsWarmDataSize {
	return s.WarmDataSize
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) GetWriteRequestBalance() *ListDoctorHBaseTablesResponseBodyDataMetricsWriteRequestBalance {
	return s.WriteRequestBalance
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) SetColdAccessDay(v *ListDoctorHBaseTablesResponseBodyDataMetricsColdAccessDay) *ListDoctorHBaseTablesResponseBodyDataMetrics {
	s.ColdAccessDay = v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) SetColdConfigDay(v *ListDoctorHBaseTablesResponseBodyDataMetricsColdConfigDay) *ListDoctorHBaseTablesResponseBodyDataMetrics {
	s.ColdConfigDay = v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) SetColdDataSize(v *ListDoctorHBaseTablesResponseBodyDataMetricsColdDataSize) *ListDoctorHBaseTablesResponseBodyDataMetrics {
	s.ColdDataSize = v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) SetDailyReadRequest(v *ListDoctorHBaseTablesResponseBodyDataMetricsDailyReadRequest) *ListDoctorHBaseTablesResponseBodyDataMetrics {
	s.DailyReadRequest = v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) SetDailyReadRequestDayGrowthRatio(v *ListDoctorHBaseTablesResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) *ListDoctorHBaseTablesResponseBodyDataMetrics {
	s.DailyReadRequestDayGrowthRatio = v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) SetDailyWriteRequest(v *ListDoctorHBaseTablesResponseBodyDataMetricsDailyWriteRequest) *ListDoctorHBaseTablesResponseBodyDataMetrics {
	s.DailyWriteRequest = v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) SetDailyWriteRequestDayGrowthRatio(v *ListDoctorHBaseTablesResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) *ListDoctorHBaseTablesResponseBodyDataMetrics {
	s.DailyWriteRequestDayGrowthRatio = v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) SetFreezeConfigDay(v *ListDoctorHBaseTablesResponseBodyDataMetricsFreezeConfigDay) *ListDoctorHBaseTablesResponseBodyDataMetrics {
	s.FreezeConfigDay = v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) SetFreezeDataSize(v *ListDoctorHBaseTablesResponseBodyDataMetricsFreezeDataSize) *ListDoctorHBaseTablesResponseBodyDataMetrics {
	s.FreezeDataSize = v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) SetHotDataSize(v *ListDoctorHBaseTablesResponseBodyDataMetricsHotDataSize) *ListDoctorHBaseTablesResponseBodyDataMetrics {
	s.HotDataSize = v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) SetLocality(v *ListDoctorHBaseTablesResponseBodyDataMetricsLocality) *ListDoctorHBaseTablesResponseBodyDataMetrics {
	s.Locality = v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) SetReadRequestBalance(v *ListDoctorHBaseTablesResponseBodyDataMetricsReadRequestBalance) *ListDoctorHBaseTablesResponseBodyDataMetrics {
	s.ReadRequestBalance = v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) SetRegionBalance(v *ListDoctorHBaseTablesResponseBodyDataMetricsRegionBalance) *ListDoctorHBaseTablesResponseBodyDataMetrics {
	s.RegionBalance = v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) SetRegionCount(v *ListDoctorHBaseTablesResponseBodyDataMetricsRegionCount) *ListDoctorHBaseTablesResponseBodyDataMetrics {
	s.RegionCount = v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) SetRegionCountDayGrowthRatio(v *ListDoctorHBaseTablesResponseBodyDataMetricsRegionCountDayGrowthRatio) *ListDoctorHBaseTablesResponseBodyDataMetrics {
	s.RegionCountDayGrowthRatio = v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) SetRegionServerCount(v *ListDoctorHBaseTablesResponseBodyDataMetricsRegionServerCount) *ListDoctorHBaseTablesResponseBodyDataMetrics {
	s.RegionServerCount = v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) SetRequestBalance(v *ListDoctorHBaseTablesResponseBodyDataMetricsRequestBalance) *ListDoctorHBaseTablesResponseBodyDataMetrics {
	s.RequestBalance = v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) SetStoreFileCount(v *ListDoctorHBaseTablesResponseBodyDataMetricsStoreFileCount) *ListDoctorHBaseTablesResponseBodyDataMetrics {
	s.StoreFileCount = v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) SetStoreFileCountDayGrowthRatio(v *ListDoctorHBaseTablesResponseBodyDataMetricsStoreFileCountDayGrowthRatio) *ListDoctorHBaseTablesResponseBodyDataMetrics {
	s.StoreFileCountDayGrowthRatio = v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) SetTableSize(v *ListDoctorHBaseTablesResponseBodyDataMetricsTableSize) *ListDoctorHBaseTablesResponseBodyDataMetrics {
	s.TableSize = v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) SetTableSizeDayGrowthRatio(v *ListDoctorHBaseTablesResponseBodyDataMetricsTableSizeDayGrowthRatio) *ListDoctorHBaseTablesResponseBodyDataMetrics {
	s.TableSizeDayGrowthRatio = v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) SetWarmConfigDay(v *ListDoctorHBaseTablesResponseBodyDataMetricsWarmConfigDay) *ListDoctorHBaseTablesResponseBodyDataMetrics {
	s.WarmConfigDay = v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) SetWarmDataSize(v *ListDoctorHBaseTablesResponseBodyDataMetricsWarmDataSize) *ListDoctorHBaseTablesResponseBodyDataMetrics {
	s.WarmDataSize = v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) SetWriteRequestBalance(v *ListDoctorHBaseTablesResponseBodyDataMetricsWriteRequestBalance) *ListDoctorHBaseTablesResponseBodyDataMetrics {
	s.WriteRequestBalance = v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetrics) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHBaseTablesResponseBodyDataMetricsColdAccessDay struct {
	// The description of the metric.
	//
	// example:
	//
	// Cold access day
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// coldAccessDay
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// day
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 1
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsColdAccessDay) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsColdAccessDay) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsColdAccessDay) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsColdAccessDay) GetName() *string {
	return s.Name
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsColdAccessDay) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsColdAccessDay) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsColdAccessDay) SetDescription(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsColdAccessDay {
	s.Description = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsColdAccessDay) SetName(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsColdAccessDay {
	s.Name = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsColdAccessDay) SetUnit(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsColdAccessDay {
	s.Unit = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsColdAccessDay) SetValue(v int64) *ListDoctorHBaseTablesResponseBodyDataMetricsColdAccessDay {
	s.Value = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsColdAccessDay) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHBaseTablesResponseBodyDataMetricsColdConfigDay struct {
	// The description of the metric.
	//
	// example:
	//
	// Cold config day
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// coldConfigDay
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// day
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 1
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsColdConfigDay) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsColdConfigDay) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsColdConfigDay) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsColdConfigDay) GetName() *string {
	return s.Name
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsColdConfigDay) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsColdConfigDay) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsColdConfigDay) SetDescription(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsColdConfigDay {
	s.Description = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsColdConfigDay) SetName(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsColdConfigDay {
	s.Name = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsColdConfigDay) SetUnit(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsColdConfigDay {
	s.Unit = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsColdConfigDay) SetValue(v int64) *ListDoctorHBaseTablesResponseBodyDataMetricsColdConfigDay {
	s.Value = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsColdConfigDay) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHBaseTablesResponseBodyDataMetricsColdDataSize struct {
	// The description of the metric.
	//
	// example:
	//
	// Size of the cold data size
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// coldDataSize
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// MB
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 100
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsColdDataSize) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsColdDataSize) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsColdDataSize) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsColdDataSize) GetName() *string {
	return s.Name
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsColdDataSize) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsColdDataSize) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsColdDataSize) SetDescription(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsColdDataSize {
	s.Description = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsColdDataSize) SetName(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsColdDataSize {
	s.Name = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsColdDataSize) SetUnit(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsColdDataSize {
	s.Unit = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsColdDataSize) SetValue(v int64) *ListDoctorHBaseTablesResponseBodyDataMetricsColdDataSize {
	s.Value = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsColdDataSize) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHBaseTablesResponseBodyDataMetricsDailyReadRequest struct {
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
	// “”
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsDailyReadRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsDailyReadRequest) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsDailyReadRequest) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsDailyReadRequest) GetName() *string {
	return s.Name
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsDailyReadRequest) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsDailyReadRequest) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsDailyReadRequest) SetDescription(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsDailyReadRequest {
	s.Description = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsDailyReadRequest) SetName(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsDailyReadRequest {
	s.Name = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsDailyReadRequest) SetUnit(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsDailyReadRequest {
	s.Unit = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsDailyReadRequest) SetValue(v int64) *ListDoctorHBaseTablesResponseBodyDataMetricsDailyReadRequest {
	s.Value = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsDailyReadRequest) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHBaseTablesResponseBodyDataMetricsDailyReadRequestDayGrowthRatio struct {
	// The description of the metric.
	//
	// example:
	//
	// Day growth ratio of read requests
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
	// 0.8
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) SetDescription(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsDailyReadRequestDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) SetName(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsDailyReadRequestDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) SetUnit(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsDailyReadRequestDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) SetValue(v float32) *ListDoctorHBaseTablesResponseBodyDataMetricsDailyReadRequestDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsDailyReadRequestDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHBaseTablesResponseBodyDataMetricsDailyWriteRequest struct {
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
	// “”
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsDailyWriteRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsDailyWriteRequest) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsDailyWriteRequest) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsDailyWriteRequest) GetName() *string {
	return s.Name
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsDailyWriteRequest) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsDailyWriteRequest) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsDailyWriteRequest) SetDescription(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsDailyWriteRequest {
	s.Description = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsDailyWriteRequest) SetName(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsDailyWriteRequest {
	s.Name = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsDailyWriteRequest) SetUnit(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsDailyWriteRequest {
	s.Unit = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsDailyWriteRequest) SetValue(v int64) *ListDoctorHBaseTablesResponseBodyDataMetricsDailyWriteRequest {
	s.Value = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsDailyWriteRequest) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHBaseTablesResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio struct {
	// The description of the metric.
	//
	// example:
	//
	// The balance of distributing requests
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
	// 0.8
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) SetDescription(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) SetName(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) SetUnit(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) SetValue(v float32) *ListDoctorHBaseTablesResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsDailyWriteRequestDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHBaseTablesResponseBodyDataMetricsFreezeConfigDay struct {
	// The description of the metric.
	//
	// example:
	//
	// Freeze config day
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// freezeConfigDay
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// day
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 1
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsFreezeConfigDay) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsFreezeConfigDay) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsFreezeConfigDay) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsFreezeConfigDay) GetName() *string {
	return s.Name
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsFreezeConfigDay) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsFreezeConfigDay) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsFreezeConfigDay) SetDescription(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsFreezeConfigDay {
	s.Description = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsFreezeConfigDay) SetName(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsFreezeConfigDay {
	s.Name = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsFreezeConfigDay) SetUnit(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsFreezeConfigDay {
	s.Unit = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsFreezeConfigDay) SetValue(v int64) *ListDoctorHBaseTablesResponseBodyDataMetricsFreezeConfigDay {
	s.Value = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsFreezeConfigDay) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHBaseTablesResponseBodyDataMetricsFreezeDataSize struct {
	// The description of the metric.
	//
	// example:
	//
	// Size of the freeze data size
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// freezeDataSize
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// MB
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 100
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsFreezeDataSize) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsFreezeDataSize) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsFreezeDataSize) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsFreezeDataSize) GetName() *string {
	return s.Name
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsFreezeDataSize) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsFreezeDataSize) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsFreezeDataSize) SetDescription(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsFreezeDataSize {
	s.Description = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsFreezeDataSize) SetName(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsFreezeDataSize {
	s.Name = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsFreezeDataSize) SetUnit(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsFreezeDataSize {
	s.Unit = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsFreezeDataSize) SetValue(v int64) *ListDoctorHBaseTablesResponseBodyDataMetricsFreezeDataSize {
	s.Value = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsFreezeDataSize) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHBaseTablesResponseBodyDataMetricsHotDataSize struct {
	// The description of the metric.
	//
	// example:
	//
	// Size of the hot data size
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
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
	// The value of the metric.
	//
	// example:
	//
	// 100
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsHotDataSize) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsHotDataSize) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsHotDataSize) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsHotDataSize) GetName() *string {
	return s.Name
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsHotDataSize) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsHotDataSize) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsHotDataSize) SetDescription(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsHotDataSize {
	s.Description = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsHotDataSize) SetName(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsHotDataSize {
	s.Name = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsHotDataSize) SetUnit(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsHotDataSize {
	s.Unit = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsHotDataSize) SetValue(v int64) *ListDoctorHBaseTablesResponseBodyDataMetricsHotDataSize {
	s.Value = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsHotDataSize) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHBaseTablesResponseBodyDataMetricsLocality struct {
	// The description of the metric.
	//
	// example:
	//
	// The locality of data
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// locality
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
	// 0
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsLocality) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsLocality) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsLocality) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsLocality) GetName() *string {
	return s.Name
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsLocality) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsLocality) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsLocality) SetDescription(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsLocality {
	s.Description = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsLocality) SetName(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsLocality {
	s.Name = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsLocality) SetUnit(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsLocality {
	s.Unit = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsLocality) SetValue(v float32) *ListDoctorHBaseTablesResponseBodyDataMetricsLocality {
	s.Value = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsLocality) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHBaseTablesResponseBodyDataMetricsReadRequestBalance struct {
	// The description of the metric.
	//
	// example:
	//
	// The balance of distributing read requests
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
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
	// 1.0
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsReadRequestBalance) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsReadRequestBalance) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsReadRequestBalance) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsReadRequestBalance) GetName() *string {
	return s.Name
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsReadRequestBalance) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsReadRequestBalance) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsReadRequestBalance) SetDescription(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsReadRequestBalance {
	s.Description = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsReadRequestBalance) SetName(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsReadRequestBalance {
	s.Name = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsReadRequestBalance) SetUnit(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsReadRequestBalance {
	s.Unit = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsReadRequestBalance) SetValue(v float32) *ListDoctorHBaseTablesResponseBodyDataMetricsReadRequestBalance {
	s.Value = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsReadRequestBalance) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHBaseTablesResponseBodyDataMetricsRegionBalance struct {
	// The description of the metric.
	//
	// example:
	//
	// The ability to evenly distribute Regions on different RegionServer nodes
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// regionBalance
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// “”
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 1.0
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsRegionBalance) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsRegionBalance) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRegionBalance) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRegionBalance) GetName() *string {
	return s.Name
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRegionBalance) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRegionBalance) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRegionBalance) SetDescription(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsRegionBalance {
	s.Description = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRegionBalance) SetName(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsRegionBalance {
	s.Name = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRegionBalance) SetUnit(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsRegionBalance {
	s.Unit = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRegionBalance) SetValue(v float32) *ListDoctorHBaseTablesResponseBodyDataMetricsRegionBalance {
	s.Value = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRegionBalance) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHBaseTablesResponseBodyDataMetricsRegionCount struct {
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
	// 3
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsRegionCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsRegionCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRegionCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRegionCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRegionCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRegionCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRegionCount) SetDescription(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsRegionCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRegionCount) SetName(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsRegionCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRegionCount) SetUnit(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsRegionCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRegionCount) SetValue(v int64) *ListDoctorHBaseTablesResponseBodyDataMetricsRegionCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRegionCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHBaseTablesResponseBodyDataMetricsRegionCountDayGrowthRatio struct {
	// The description of the metric.
	//
	// example:
	//
	// Day growth ratio of region count
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// regionCountDayGrowthRatio
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

func (s ListDoctorHBaseTablesResponseBodyDataMetricsRegionCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsRegionCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRegionCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRegionCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRegionCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRegionCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRegionCountDayGrowthRatio) SetDescription(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsRegionCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRegionCountDayGrowthRatio) SetName(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsRegionCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRegionCountDayGrowthRatio) SetUnit(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsRegionCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRegionCountDayGrowthRatio) SetValue(v float32) *ListDoctorHBaseTablesResponseBodyDataMetricsRegionCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRegionCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHBaseTablesResponseBodyDataMetricsRegionServerCount struct {
	// The description of the metric.
	//
	// example:
	//
	// Number of region servers count
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// regionServerCount
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
	// 2
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsRegionServerCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsRegionServerCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRegionServerCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRegionServerCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRegionServerCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRegionServerCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRegionServerCount) SetDescription(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsRegionServerCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRegionServerCount) SetName(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsRegionServerCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRegionServerCount) SetUnit(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsRegionServerCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRegionServerCount) SetValue(v int64) *ListDoctorHBaseTablesResponseBodyDataMetricsRegionServerCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRegionServerCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHBaseTablesResponseBodyDataMetricsRequestBalance struct {
	// The description of the metric.
	//
	// example:
	//
	// The balance of distributing requests
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// requestBalance
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// “”
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 0.9
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsRequestBalance) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsRequestBalance) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRequestBalance) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRequestBalance) GetName() *string {
	return s.Name
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRequestBalance) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRequestBalance) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRequestBalance) SetDescription(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsRequestBalance {
	s.Description = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRequestBalance) SetName(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsRequestBalance {
	s.Name = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRequestBalance) SetUnit(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsRequestBalance {
	s.Unit = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRequestBalance) SetValue(v float32) *ListDoctorHBaseTablesResponseBodyDataMetricsRequestBalance {
	s.Value = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsRequestBalance) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHBaseTablesResponseBodyDataMetricsStoreFileCount struct {
	// The description of the metric.
	//
	// example:
	//
	// Number of store files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// storeFileCount
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
	// 36
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsStoreFileCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsStoreFileCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsStoreFileCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsStoreFileCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsStoreFileCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsStoreFileCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsStoreFileCount) SetDescription(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsStoreFileCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsStoreFileCount) SetName(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsStoreFileCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsStoreFileCount) SetUnit(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsStoreFileCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsStoreFileCount) SetValue(v int64) *ListDoctorHBaseTablesResponseBodyDataMetricsStoreFileCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsStoreFileCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHBaseTablesResponseBodyDataMetricsStoreFileCountDayGrowthRatio struct {
	// The description of the metric.
	//
	// example:
	//
	// Day growth ratio of store file count
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// storeFileCountDayGrowthRatio
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
	// 0.7
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsStoreFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsStoreFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsStoreFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsStoreFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsStoreFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsStoreFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsStoreFileCountDayGrowthRatio) SetDescription(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsStoreFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsStoreFileCountDayGrowthRatio) SetName(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsStoreFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsStoreFileCountDayGrowthRatio) SetUnit(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsStoreFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsStoreFileCountDayGrowthRatio) SetValue(v float32) *ListDoctorHBaseTablesResponseBodyDataMetricsStoreFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsStoreFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHBaseTablesResponseBodyDataMetricsTableSize struct {
	// The description of the metric.
	//
	// example:
	//
	// Size of the table
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// tableSize
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// MB
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 678
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsTableSize) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsTableSize) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsTableSize) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsTableSize) GetName() *string {
	return s.Name
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsTableSize) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsTableSize) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsTableSize) SetDescription(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsTableSize {
	s.Description = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsTableSize) SetName(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsTableSize {
	s.Name = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsTableSize) SetUnit(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsTableSize {
	s.Unit = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsTableSize) SetValue(v int64) *ListDoctorHBaseTablesResponseBodyDataMetricsTableSize {
	s.Value = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsTableSize) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHBaseTablesResponseBodyDataMetricsTableSizeDayGrowthRatio struct {
	// The description of the metric.
	//
	// example:
	//
	// Day growth ratio of table size
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// tableSizeDayGrowthRatio
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// \\""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 1.0
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsTableSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsTableSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsTableSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsTableSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsTableSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsTableSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsTableSizeDayGrowthRatio) SetDescription(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsTableSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsTableSizeDayGrowthRatio) SetName(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsTableSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsTableSizeDayGrowthRatio) SetUnit(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsTableSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsTableSizeDayGrowthRatio) SetValue(v float32) *ListDoctorHBaseTablesResponseBodyDataMetricsTableSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsTableSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHBaseTablesResponseBodyDataMetricsWarmConfigDay struct {
	// The description of the metric.
	//
	// example:
	//
	// Warm config day
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// warmConfigDay
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// day
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 1
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsWarmConfigDay) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsWarmConfigDay) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsWarmConfigDay) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsWarmConfigDay) GetName() *string {
	return s.Name
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsWarmConfigDay) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsWarmConfigDay) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsWarmConfigDay) SetDescription(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsWarmConfigDay {
	s.Description = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsWarmConfigDay) SetName(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsWarmConfigDay {
	s.Name = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsWarmConfigDay) SetUnit(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsWarmConfigDay {
	s.Unit = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsWarmConfigDay) SetValue(v int64) *ListDoctorHBaseTablesResponseBodyDataMetricsWarmConfigDay {
	s.Value = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsWarmConfigDay) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHBaseTablesResponseBodyDataMetricsWarmDataSize struct {
	// The description of the metric.
	//
	// example:
	//
	// Size of the warm data size
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
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
	// The value of the metric.
	//
	// example:
	//
	// 100
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsWarmDataSize) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsWarmDataSize) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsWarmDataSize) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsWarmDataSize) GetName() *string {
	return s.Name
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsWarmDataSize) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsWarmDataSize) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsWarmDataSize) SetDescription(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsWarmDataSize {
	s.Description = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsWarmDataSize) SetName(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsWarmDataSize {
	s.Name = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsWarmDataSize) SetUnit(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsWarmDataSize {
	s.Unit = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsWarmDataSize) SetValue(v int64) *ListDoctorHBaseTablesResponseBodyDataMetricsWarmDataSize {
	s.Value = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsWarmDataSize) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHBaseTablesResponseBodyDataMetricsWriteRequestBalance struct {
	// The description of the metric.
	//
	// example:
	//
	// The balance of distributing write requests
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// writeRequestBalance
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

func (s ListDoctorHBaseTablesResponseBodyDataMetricsWriteRequestBalance) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseTablesResponseBodyDataMetricsWriteRequestBalance) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsWriteRequestBalance) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsWriteRequestBalance) GetName() *string {
	return s.Name
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsWriteRequestBalance) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsWriteRequestBalance) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsWriteRequestBalance) SetDescription(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsWriteRequestBalance {
	s.Description = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsWriteRequestBalance) SetName(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsWriteRequestBalance {
	s.Name = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsWriteRequestBalance) SetUnit(v string) *ListDoctorHBaseTablesResponseBodyDataMetricsWriteRequestBalance {
	s.Unit = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsWriteRequestBalance) SetValue(v float32) *ListDoctorHBaseTablesResponseBodyDataMetricsWriteRequestBalance {
	s.Value = &v
	return s
}

func (s *ListDoctorHBaseTablesResponseBodyDataMetricsWriteRequestBalance) Validate() error {
	return dara.Validate(s)
}
