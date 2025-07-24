// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDoctorHiveTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListDoctorHiveTablesResponseBodyData) *ListDoctorHiveTablesResponseBody
	GetData() []*ListDoctorHiveTablesResponseBodyData
	SetMaxResults(v int32) *ListDoctorHiveTablesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDoctorHiveTablesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDoctorHiveTablesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListDoctorHiveTablesResponseBody
	GetTotalCount() *int32
}

type ListDoctorHiveTablesResponseBody struct {
	// The analysis results of Hive tables.
	Data []*ListDoctorHiveTablesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s ListDoctorHiveTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBody) GetData() []*ListDoctorHiveTablesResponseBodyData {
	return s.Data
}

func (s *ListDoctorHiveTablesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDoctorHiveTablesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDoctorHiveTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDoctorHiveTablesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDoctorHiveTablesResponseBody) SetData(v []*ListDoctorHiveTablesResponseBodyData) *ListDoctorHiveTablesResponseBody {
	s.Data = v
	return s
}

func (s *ListDoctorHiveTablesResponseBody) SetMaxResults(v int32) *ListDoctorHiveTablesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBody) SetNextToken(v string) *ListDoctorHiveTablesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBody) SetRequestId(v string) *ListDoctorHiveTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBody) SetTotalCount(v int32) *ListDoctorHiveTablesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyData struct {
	// The analysis results.
	Analysis *ListDoctorHiveTablesResponseBodyDataAnalysis `json:"Analysis,omitempty" xml:"Analysis,omitempty" type:"Struct"`
	// The table format information.
	Formats []*ListDoctorHiveTablesResponseBodyDataFormats `json:"Formats,omitempty" xml:"Formats,omitempty" type:"Repeated"`
	// The metric information.
	Metrics *ListDoctorHiveTablesResponseBodyDataMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Struct"`
	// The owner.
	//
	// example:
	//
	// DW
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The table name. The table name must follow the naming rule in Hive. A name in the {Database name.Table name} format uniquely identifies a table.
	//
	// example:
	//
	// dw.dwd_creta_service_order_long_renew_long_da
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyData) GetAnalysis() *ListDoctorHiveTablesResponseBodyDataAnalysis {
	return s.Analysis
}

func (s *ListDoctorHiveTablesResponseBodyData) GetFormats() []*ListDoctorHiveTablesResponseBodyDataFormats {
	return s.Formats
}

func (s *ListDoctorHiveTablesResponseBodyData) GetMetrics() *ListDoctorHiveTablesResponseBodyDataMetrics {
	return s.Metrics
}

func (s *ListDoctorHiveTablesResponseBodyData) GetOwner() *string {
	return s.Owner
}

func (s *ListDoctorHiveTablesResponseBodyData) GetTableName() *string {
	return s.TableName
}

func (s *ListDoctorHiveTablesResponseBodyData) SetAnalysis(v *ListDoctorHiveTablesResponseBodyDataAnalysis) *ListDoctorHiveTablesResponseBodyData {
	s.Analysis = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyData) SetFormats(v []*ListDoctorHiveTablesResponseBodyDataFormats) *ListDoctorHiveTablesResponseBodyData {
	s.Formats = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyData) SetMetrics(v *ListDoctorHiveTablesResponseBodyDataMetrics) *ListDoctorHiveTablesResponseBodyData {
	s.Metrics = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyData) SetOwner(v string) *ListDoctorHiveTablesResponseBodyData {
	s.Owner = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyData) SetTableName(v string) *ListDoctorHiveTablesResponseBodyData {
	s.TableName = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataAnalysis struct {
	// The score for the file sizes of the Hive table.
	//
	// example:
	//
	// 80
	HiveDistributionScore *int32 `json:"HiveDistributionScore,omitempty" xml:"HiveDistributionScore,omitempty"`
	// The score for the data formats of the Hive table.
	//
	// example:
	//
	// 60
	HiveFormatScore *int32 `json:"HiveFormatScore,omitempty" xml:"HiveFormatScore,omitempty"`
	// The score for the access frequency of the Hive table.
	//
	// example:
	//
	// 70
	HiveFrequencyScore *int32 `json:"HiveFrequencyScore,omitempty" xml:"HiveFrequencyScore,omitempty"`
	// The overall score of the Hive table.
	//
	// example:
	//
	// 80
	HiveScore *int32 `json:"HiveScore,omitempty" xml:"HiveScore,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataAnalysis) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataAnalysis) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataAnalysis) GetHiveDistributionScore() *int32 {
	return s.HiveDistributionScore
}

func (s *ListDoctorHiveTablesResponseBodyDataAnalysis) GetHiveFormatScore() *int32 {
	return s.HiveFormatScore
}

func (s *ListDoctorHiveTablesResponseBodyDataAnalysis) GetHiveFrequencyScore() *int32 {
	return s.HiveFrequencyScore
}

func (s *ListDoctorHiveTablesResponseBodyDataAnalysis) GetHiveScore() *int32 {
	return s.HiveScore
}

func (s *ListDoctorHiveTablesResponseBodyDataAnalysis) SetHiveDistributionScore(v int32) *ListDoctorHiveTablesResponseBodyDataAnalysis {
	s.HiveDistributionScore = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataAnalysis) SetHiveFormatScore(v int32) *ListDoctorHiveTablesResponseBodyDataAnalysis {
	s.HiveFormatScore = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataAnalysis) SetHiveFrequencyScore(v int32) *ListDoctorHiveTablesResponseBodyDataAnalysis {
	s.HiveFrequencyScore = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataAnalysis) SetHiveScore(v int32) *ListDoctorHiveTablesResponseBodyDataAnalysis {
	s.HiveScore = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataAnalysis) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataFormats struct {
	// The daily increment of data in the format.
	//
	// example:
	//
	// 1232124
	FormatDayGrowthSize *int64 `json:"FormatDayGrowthSize,omitempty" xml:"FormatDayGrowthSize,omitempty"`
	// The name of the storage format.
	//
	// example:
	//
	// TextInputFormat
	FormatName *string `json:"FormatName,omitempty" xml:"FormatName,omitempty"`
	// The proportion of the data in the format.
	//
	// example:
	//
	// 0.23
	FormatRatio *float32 `json:"FormatRatio,omitempty" xml:"FormatRatio,omitempty"`
	// The amount of data in the format.
	//
	// example:
	//
	// 506930200
	FormatSize *int64 `json:"FormatSize,omitempty" xml:"FormatSize,omitempty"`
	// The day-to-day growth rate of data in the format.
	//
	// example:
	//
	// 0.04
	FormatSizeDayGrowthRatio *float32 `json:"FormatSizeDayGrowthRatio,omitempty" xml:"FormatSizeDayGrowthRatio,omitempty"`
	// The unit of the amount of data in the format.
	//
	// example:
	//
	// MB
	FormatSizeUnit *string `json:"FormatSizeUnit,omitempty" xml:"FormatSizeUnit,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataFormats) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataFormats) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataFormats) GetFormatDayGrowthSize() *int64 {
	return s.FormatDayGrowthSize
}

func (s *ListDoctorHiveTablesResponseBodyDataFormats) GetFormatName() *string {
	return s.FormatName
}

func (s *ListDoctorHiveTablesResponseBodyDataFormats) GetFormatRatio() *float32 {
	return s.FormatRatio
}

func (s *ListDoctorHiveTablesResponseBodyDataFormats) GetFormatSize() *int64 {
	return s.FormatSize
}

func (s *ListDoctorHiveTablesResponseBodyDataFormats) GetFormatSizeDayGrowthRatio() *float32 {
	return s.FormatSizeDayGrowthRatio
}

func (s *ListDoctorHiveTablesResponseBodyDataFormats) GetFormatSizeUnit() *string {
	return s.FormatSizeUnit
}

func (s *ListDoctorHiveTablesResponseBodyDataFormats) SetFormatDayGrowthSize(v int64) *ListDoctorHiveTablesResponseBodyDataFormats {
	s.FormatDayGrowthSize = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataFormats) SetFormatName(v string) *ListDoctorHiveTablesResponseBodyDataFormats {
	s.FormatName = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataFormats) SetFormatRatio(v float32) *ListDoctorHiveTablesResponseBodyDataFormats {
	s.FormatRatio = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataFormats) SetFormatSize(v int64) *ListDoctorHiveTablesResponseBodyDataFormats {
	s.FormatSize = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataFormats) SetFormatSizeDayGrowthRatio(v float32) *ListDoctorHiveTablesResponseBodyDataFormats {
	s.FormatSizeDayGrowthRatio = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataFormats) SetFormatSizeUnit(v string) *ListDoctorHiveTablesResponseBodyDataFormats {
	s.FormatSizeUnit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataFormats) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetrics struct {
	// The daily increment of the amount of cold data. Cold data refers to data that is not accessed for more than 30 days but is accessed in previous 90 days.
	ColdDataDayGrowthSize *ListDoctorHiveTablesResponseBodyDataMetricsColdDataDayGrowthSize `json:"ColdDataDayGrowthSize,omitempty" xml:"ColdDataDayGrowthSize,omitempty" type:"Struct"`
	// The proportion of cold data. Cold data refers to data that is not accessed for more than 30 days but is accessed in previous 90 days.
	ColdDataRatio *ListDoctorHiveTablesResponseBodyDataMetricsColdDataRatio `json:"ColdDataRatio,omitempty" xml:"ColdDataRatio,omitempty" type:"Struct"`
	// The amount of cold data. Cold data refers to data that is not accessed for more than 30 days but is accessed in previous 90 days.
	ColdDataSize *ListDoctorHiveTablesResponseBodyDataMetricsColdDataSize `json:"ColdDataSize,omitempty" xml:"ColdDataSize,omitempty" type:"Struct"`
	// The day-to-day growth rate of the amount of cold data. Cold data refers to data that is not accessed for more than 30 days but is accessed in previous 90 days.
	ColdDataSizeDayGrowthRatio *ListDoctorHiveTablesResponseBodyDataMetricsColdDataSizeDayGrowthRatio `json:"ColdDataSizeDayGrowthRatio,omitempty" xml:"ColdDataSizeDayGrowthRatio,omitempty" type:"Struct"`
	// The number of empty files. Empty files are those with a size of 0 MB.
	EmptyFileCount *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileCount `json:"EmptyFileCount,omitempty" xml:"EmptyFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the number of empty files. Empty files are those with a size of 0 MB.
	EmptyFileCountDayGrowthRatio *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio `json:"EmptyFileCountDayGrowthRatio,omitempty" xml:"EmptyFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the number of empty files. Empty files are those with a size of 0 MB.
	EmptyFileDayGrowthCount *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileDayGrowthCount `json:"EmptyFileDayGrowthCount,omitempty" xml:"EmptyFileDayGrowthCount,omitempty" type:"Struct"`
	// The proportion of empty files. Empty files are those with a size of 0 MB.
	EmptyFileRatio *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileRatio `json:"EmptyFileRatio,omitempty" xml:"EmptyFileRatio,omitempty" type:"Struct"`
	// The daily increment of the amount of very cold data. Very cold data refers to data that is not accessed for more than 90 days.
	FreezeDataDayGrowthSize *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataDayGrowthSize `json:"FreezeDataDayGrowthSize,omitempty" xml:"FreezeDataDayGrowthSize,omitempty" type:"Struct"`
	// The proportion of very cold data. Very cold data refers to data that is not accessed for more than 90 days.
	FreezeDataRatio *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataRatio `json:"FreezeDataRatio,omitempty" xml:"FreezeDataRatio,omitempty" type:"Struct"`
	// The amount of very cold data. Very cold data refers to data that is not accessed for more than 90 days.
	FreezeDataSize *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataSize `json:"FreezeDataSize,omitempty" xml:"FreezeDataSize,omitempty" type:"Struct"`
	// The day-to-day growth rate of the amount of very cold data. Very cold data refers to data that is not accessed for more than 90 days.
	FreezeDataSizeDayGrowthRatio *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio `json:"FreezeDataSizeDayGrowthRatio,omitempty" xml:"FreezeDataSizeDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the amount of hot data. Hot data refers to data that is accessed in previous seven days.
	HotDataDayGrowthSize *ListDoctorHiveTablesResponseBodyDataMetricsHotDataDayGrowthSize `json:"HotDataDayGrowthSize,omitempty" xml:"HotDataDayGrowthSize,omitempty" type:"Struct"`
	// The proportion of hot data. Hot data refers to data that is accessed in previous seven days.
	HotDataRatio *ListDoctorHiveTablesResponseBodyDataMetricsHotDataRatio `json:"HotDataRatio,omitempty" xml:"HotDataRatio,omitempty" type:"Struct"`
	// The amount of hot data. Hot data refers to data that is accessed in previous seven days.
	HotDataSize *ListDoctorHiveTablesResponseBodyDataMetricsHotDataSize `json:"HotDataSize,omitempty" xml:"HotDataSize,omitempty" type:"Struct"`
	// The day-to-day growth rate of the amount of hot data. Hot data refers to data that is accessed in previous seven days.
	HotDataSizeDayGrowthRatio *ListDoctorHiveTablesResponseBodyDataMetricsHotDataSizeDayGrowthRatio `json:"HotDataSizeDayGrowthRatio,omitempty" xml:"HotDataSizeDayGrowthRatio,omitempty" type:"Struct"`
	// The number of large files. Large files are those with a size greater than 1 GB.
	LargeFileCount *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileCount `json:"LargeFileCount,omitempty" xml:"LargeFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the number of large files. Large files are those with a size greater than 1 GB.
	LargeFileCountDayGrowthRatio *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileCountDayGrowthRatio `json:"LargeFileCountDayGrowthRatio,omitempty" xml:"LargeFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the number of large files. Large files are those with a size greater than 1 GB.
	LargeFileDayGrowthCount *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileDayGrowthCount `json:"LargeFileDayGrowthCount,omitempty" xml:"LargeFileDayGrowthCount,omitempty" type:"Struct"`
	// The proportion of large files. Large files are those with a size greater than 1 GB.
	LargeFileRatio *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileRatio `json:"LargeFileRatio,omitempty" xml:"LargeFileRatio,omitempty" type:"Struct"`
	// The number of medium files. Medium files are those with a size greater than or equal to 128 MB and less than or equal to 1 GB.
	MediumFileCount *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileCount `json:"MediumFileCount,omitempty" xml:"MediumFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the number of medium files. Medium files are those with a size greater than or equal to 128 MB and less than or equal to 1 GB.
	MediumFileCountDayGrowthRatio *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileCountDayGrowthRatio `json:"MediumFileCountDayGrowthRatio,omitempty" xml:"MediumFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the number of medium files. Medium files are those with a size greater than or equal to 128 MB and less than or equal to 1 GB.
	MediumFileDayGrowthCount *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileDayGrowthCount `json:"MediumFileDayGrowthCount,omitempty" xml:"MediumFileDayGrowthCount,omitempty" type:"Struct"`
	// The proportion of medium files. Medium files are those with a size greater than or equal to 128 MB and less than or equal to 1 GB.
	MediumFileRatio *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileRatio `json:"MediumFileRatio,omitempty" xml:"MediumFileRatio,omitempty" type:"Struct"`
	// The number of partitions.
	PartitionNum *ListDoctorHiveTablesResponseBodyDataMetricsPartitionNum `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty" type:"Struct"`
	// The number of small files. Small files are those with a size greater than or equal to 10 MB and less than 128 MB.
	SmallFileCount *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileCount `json:"SmallFileCount,omitempty" xml:"SmallFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the number of small files. Small files are those with a size greater than or equal to 10 MB and less than 128 MB.
	SmallFileCountDayGrowthRatio *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileCountDayGrowthRatio `json:"SmallFileCountDayGrowthRatio,omitempty" xml:"SmallFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the number of small files. Small files are those with a size greater than or equal to 10 MB and less than 128 MB.
	SmallFileDayGrowthCount *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileDayGrowthCount `json:"SmallFileDayGrowthCount,omitempty" xml:"SmallFileDayGrowthCount,omitempty" type:"Struct"`
	// The proportion of small files. Small files are those with a size greater than or equal to 10 MB and less than 128 MB.
	SmallFileRatio *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileRatio `json:"SmallFileRatio,omitempty" xml:"SmallFileRatio,omitempty" type:"Struct"`
	// The number of very small files. Very small files are those with a size greater than 0 MB and less than 10 MB.
	TinyFileCount *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileCount `json:"TinyFileCount,omitempty" xml:"TinyFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the number of very small files. Very small files are those with a size greater than 0 MB and less than 10 MB.
	TinyFileCountDayGrowthRatio *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileCountDayGrowthRatio `json:"TinyFileCountDayGrowthRatio,omitempty" xml:"TinyFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the number of very small files. Very small files are those with a size greater than 0 MB and less than 10 MB.
	TinyFileDayGrowthCount *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileDayGrowthCount `json:"TinyFileDayGrowthCount,omitempty" xml:"TinyFileDayGrowthCount,omitempty" type:"Struct"`
	// The proportion of very small files. Very small files are those with a size greater than 0 MB and less than 10 MB.
	TinyFileRatio *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileRatio `json:"TinyFileRatio,omitempty" xml:"TinyFileRatio,omitempty" type:"Struct"`
	// The daily increment of the total amount of data.
	TotalDataDayGrowthSize *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataDayGrowthSize `json:"TotalDataDayGrowthSize,omitempty" xml:"TotalDataDayGrowthSize,omitempty" type:"Struct"`
	// The total amount of data.
	TotalDataSize *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataSize `json:"TotalDataSize,omitempty" xml:"TotalDataSize,omitempty" type:"Struct"`
	// The day-to-day growth rate of the total amount of data.
	TotalDataSizeDayGrowthRatio *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio `json:"TotalDataSizeDayGrowthRatio,omitempty" xml:"TotalDataSizeDayGrowthRatio,omitempty" type:"Struct"`
	// The total number of files.
	TotalFileCount *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileCount `json:"TotalFileCount,omitempty" xml:"TotalFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the total number of files.
	TotalFileCountDayGrowthRatio *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileCountDayGrowthRatio `json:"TotalFileCountDayGrowthRatio,omitempty" xml:"TotalFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the total number of files.
	TotalFileDayGrowthCount *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileDayGrowthCount `json:"TotalFileDayGrowthCount,omitempty" xml:"TotalFileDayGrowthCount,omitempty" type:"Struct"`
	// The daily increment of the amount of warm data. Warm data refers to data that is not accessed for more than 7 days but is accessed in previous 30 days.
	WarmDataDayGrowthSize *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataDayGrowthSize `json:"WarmDataDayGrowthSize,omitempty" xml:"WarmDataDayGrowthSize,omitempty" type:"Struct"`
	// The proportion of warm data. Warm data refers to data that is not accessed for more than 7 days but is accessed in previous 30 days.
	WarmDataRatio *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataRatio `json:"WarmDataRatio,omitempty" xml:"WarmDataRatio,omitempty" type:"Struct"`
	// The amount of warm data. Warm data refers to data that is not accessed for more than 7 days but is accessed in previous 30 days.
	WarmDataSize *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataSize `json:"WarmDataSize,omitempty" xml:"WarmDataSize,omitempty" type:"Struct"`
	// The day-to-day growth rate of the amount of warm data. Warm data refers to data that is not accessed for more than 7 days but is accessed in previous 30 days.
	WarmDataSizeDayGrowthRatio *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio `json:"WarmDataSizeDayGrowthRatio,omitempty" xml:"WarmDataSizeDayGrowthRatio,omitempty" type:"Struct"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetrics) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetrics) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetColdDataDayGrowthSize() *ListDoctorHiveTablesResponseBodyDataMetricsColdDataDayGrowthSize {
	return s.ColdDataDayGrowthSize
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetColdDataRatio() *ListDoctorHiveTablesResponseBodyDataMetricsColdDataRatio {
	return s.ColdDataRatio
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetColdDataSize() *ListDoctorHiveTablesResponseBodyDataMetricsColdDataSize {
	return s.ColdDataSize
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetColdDataSizeDayGrowthRatio() *ListDoctorHiveTablesResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	return s.ColdDataSizeDayGrowthRatio
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetEmptyFileCount() *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileCount {
	return s.EmptyFileCount
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetEmptyFileCountDayGrowthRatio() *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	return s.EmptyFileCountDayGrowthRatio
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetEmptyFileDayGrowthCount() *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileDayGrowthCount {
	return s.EmptyFileDayGrowthCount
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetEmptyFileRatio() *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileRatio {
	return s.EmptyFileRatio
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetFreezeDataDayGrowthSize() *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataDayGrowthSize {
	return s.FreezeDataDayGrowthSize
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetFreezeDataRatio() *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataRatio {
	return s.FreezeDataRatio
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetFreezeDataSize() *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataSize {
	return s.FreezeDataSize
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetFreezeDataSizeDayGrowthRatio() *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	return s.FreezeDataSizeDayGrowthRatio
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetHotDataDayGrowthSize() *ListDoctorHiveTablesResponseBodyDataMetricsHotDataDayGrowthSize {
	return s.HotDataDayGrowthSize
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetHotDataRatio() *ListDoctorHiveTablesResponseBodyDataMetricsHotDataRatio {
	return s.HotDataRatio
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetHotDataSize() *ListDoctorHiveTablesResponseBodyDataMetricsHotDataSize {
	return s.HotDataSize
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetHotDataSizeDayGrowthRatio() *ListDoctorHiveTablesResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	return s.HotDataSizeDayGrowthRatio
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetLargeFileCount() *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileCount {
	return s.LargeFileCount
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetLargeFileCountDayGrowthRatio() *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	return s.LargeFileCountDayGrowthRatio
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetLargeFileDayGrowthCount() *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileDayGrowthCount {
	return s.LargeFileDayGrowthCount
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetLargeFileRatio() *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileRatio {
	return s.LargeFileRatio
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetMediumFileCount() *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileCount {
	return s.MediumFileCount
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetMediumFileCountDayGrowthRatio() *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	return s.MediumFileCountDayGrowthRatio
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetMediumFileDayGrowthCount() *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileDayGrowthCount {
	return s.MediumFileDayGrowthCount
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetMediumFileRatio() *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileRatio {
	return s.MediumFileRatio
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetPartitionNum() *ListDoctorHiveTablesResponseBodyDataMetricsPartitionNum {
	return s.PartitionNum
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetSmallFileCount() *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileCount {
	return s.SmallFileCount
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetSmallFileCountDayGrowthRatio() *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	return s.SmallFileCountDayGrowthRatio
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetSmallFileDayGrowthCount() *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileDayGrowthCount {
	return s.SmallFileDayGrowthCount
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetSmallFileRatio() *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileRatio {
	return s.SmallFileRatio
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetTinyFileCount() *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileCount {
	return s.TinyFileCount
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetTinyFileCountDayGrowthRatio() *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	return s.TinyFileCountDayGrowthRatio
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetTinyFileDayGrowthCount() *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileDayGrowthCount {
	return s.TinyFileDayGrowthCount
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetTinyFileRatio() *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileRatio {
	return s.TinyFileRatio
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetTotalDataDayGrowthSize() *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataDayGrowthSize {
	return s.TotalDataDayGrowthSize
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetTotalDataSize() *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataSize {
	return s.TotalDataSize
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetTotalDataSizeDayGrowthRatio() *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	return s.TotalDataSizeDayGrowthRatio
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetTotalFileCount() *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileCount {
	return s.TotalFileCount
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetTotalFileCountDayGrowthRatio() *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	return s.TotalFileCountDayGrowthRatio
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetTotalFileDayGrowthCount() *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileDayGrowthCount {
	return s.TotalFileDayGrowthCount
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetWarmDataDayGrowthSize() *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataDayGrowthSize {
	return s.WarmDataDayGrowthSize
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetWarmDataRatio() *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataRatio {
	return s.WarmDataRatio
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetWarmDataSize() *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataSize {
	return s.WarmDataSize
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) GetWarmDataSizeDayGrowthRatio() *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	return s.WarmDataSizeDayGrowthRatio
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetColdDataDayGrowthSize(v *ListDoctorHiveTablesResponseBodyDataMetricsColdDataDayGrowthSize) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.ColdDataDayGrowthSize = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetColdDataRatio(v *ListDoctorHiveTablesResponseBodyDataMetricsColdDataRatio) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.ColdDataRatio = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetColdDataSize(v *ListDoctorHiveTablesResponseBodyDataMetricsColdDataSize) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.ColdDataSize = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetColdDataSizeDayGrowthRatio(v *ListDoctorHiveTablesResponseBodyDataMetricsColdDataSizeDayGrowthRatio) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.ColdDataSizeDayGrowthRatio = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetEmptyFileCount(v *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileCount) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.EmptyFileCount = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetEmptyFileCountDayGrowthRatio(v *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.EmptyFileCountDayGrowthRatio = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetEmptyFileDayGrowthCount(v *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileDayGrowthCount) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.EmptyFileDayGrowthCount = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetEmptyFileRatio(v *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileRatio) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.EmptyFileRatio = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetFreezeDataDayGrowthSize(v *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataDayGrowthSize) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.FreezeDataDayGrowthSize = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetFreezeDataRatio(v *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataRatio) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.FreezeDataRatio = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetFreezeDataSize(v *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataSize) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.FreezeDataSize = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetFreezeDataSizeDayGrowthRatio(v *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.FreezeDataSizeDayGrowthRatio = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetHotDataDayGrowthSize(v *ListDoctorHiveTablesResponseBodyDataMetricsHotDataDayGrowthSize) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.HotDataDayGrowthSize = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetHotDataRatio(v *ListDoctorHiveTablesResponseBodyDataMetricsHotDataRatio) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.HotDataRatio = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetHotDataSize(v *ListDoctorHiveTablesResponseBodyDataMetricsHotDataSize) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.HotDataSize = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetHotDataSizeDayGrowthRatio(v *ListDoctorHiveTablesResponseBodyDataMetricsHotDataSizeDayGrowthRatio) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.HotDataSizeDayGrowthRatio = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetLargeFileCount(v *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileCount) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.LargeFileCount = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetLargeFileCountDayGrowthRatio(v *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileCountDayGrowthRatio) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.LargeFileCountDayGrowthRatio = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetLargeFileDayGrowthCount(v *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileDayGrowthCount) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.LargeFileDayGrowthCount = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetLargeFileRatio(v *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileRatio) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.LargeFileRatio = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetMediumFileCount(v *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileCount) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.MediumFileCount = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetMediumFileCountDayGrowthRatio(v *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileCountDayGrowthRatio) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.MediumFileCountDayGrowthRatio = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetMediumFileDayGrowthCount(v *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileDayGrowthCount) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.MediumFileDayGrowthCount = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetMediumFileRatio(v *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileRatio) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.MediumFileRatio = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetPartitionNum(v *ListDoctorHiveTablesResponseBodyDataMetricsPartitionNum) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.PartitionNum = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetSmallFileCount(v *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileCount) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.SmallFileCount = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetSmallFileCountDayGrowthRatio(v *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileCountDayGrowthRatio) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.SmallFileCountDayGrowthRatio = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetSmallFileDayGrowthCount(v *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileDayGrowthCount) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.SmallFileDayGrowthCount = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetSmallFileRatio(v *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileRatio) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.SmallFileRatio = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetTinyFileCount(v *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileCount) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.TinyFileCount = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetTinyFileCountDayGrowthRatio(v *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileCountDayGrowthRatio) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.TinyFileCountDayGrowthRatio = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetTinyFileDayGrowthCount(v *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileDayGrowthCount) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.TinyFileDayGrowthCount = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetTinyFileRatio(v *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileRatio) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.TinyFileRatio = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetTotalDataDayGrowthSize(v *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataDayGrowthSize) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.TotalDataDayGrowthSize = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetTotalDataSize(v *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataSize) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.TotalDataSize = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetTotalDataSizeDayGrowthRatio(v *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.TotalDataSizeDayGrowthRatio = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetTotalFileCount(v *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileCount) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.TotalFileCount = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetTotalFileCountDayGrowthRatio(v *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileCountDayGrowthRatio) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.TotalFileCountDayGrowthRatio = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetTotalFileDayGrowthCount(v *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileDayGrowthCount) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.TotalFileDayGrowthCount = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetWarmDataDayGrowthSize(v *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataDayGrowthSize) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.WarmDataDayGrowthSize = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetWarmDataRatio(v *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataRatio) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.WarmDataRatio = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetWarmDataSize(v *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataSize) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.WarmDataSize = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) SetWarmDataSizeDayGrowthRatio(v *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) *ListDoctorHiveTablesResponseBodyDataMetrics {
	s.WarmDataSizeDayGrowthRatio = v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetrics) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsColdDataDayGrowthSize struct {
	// The description of the metric.
	//
	// example:
	//
	// Day growth size of cold data
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// coldDataDayGrowthSize
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
	// 217715
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsColdDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsColdDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsColdDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsColdDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsColdDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsColdDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsColdDataDayGrowthSize) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsColdDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsColdDataDayGrowthSize) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsColdDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsColdDataDayGrowthSize) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsColdDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsColdDataDayGrowthSize) SetValue(v int64) *ListDoctorHiveTablesResponseBodyDataMetricsColdDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsColdDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsColdDataRatio struct {
	// The description of the metric.
	//
	// example:
	//
	// Cold data ratio
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// coldDataRatio
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
	// 0.01
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsColdDataRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsColdDataRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsColdDataRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsColdDataRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsColdDataRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsColdDataRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsColdDataRatio) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsColdDataRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsColdDataRatio) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsColdDataRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsColdDataRatio) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsColdDataRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsColdDataRatio) SetValue(v float32) *ListDoctorHiveTablesResponseBodyDataMetricsColdDataRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsColdDataRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsColdDataSize struct {
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
	// 217715
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsColdDataSize) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsColdDataSize) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsColdDataSize) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsColdDataSize) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsColdDataSize) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsColdDataSize) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsColdDataSize) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsColdDataSize {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsColdDataSize) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsColdDataSize {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsColdDataSize) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsColdDataSize {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsColdDataSize) SetValue(v int64) *ListDoctorHiveTablesResponseBodyDataMetricsColdDataSize {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsColdDataSize) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsColdDataSizeDayGrowthRatio struct {
	// The description of the metric.
	//
	// example:
	//
	// Day growth ratio of cold data size
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// coldDataSizeDayGrowthRatio
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
	// 0.01
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsColdDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsColdDataSizeDayGrowthRatio) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsColdDataSizeDayGrowthRatio) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsColdDataSizeDayGrowthRatio) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsColdDataSizeDayGrowthRatio) SetValue(v float32) *ListDoctorHiveTablesResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsColdDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileCount struct {
	// The description of the metric.
	//
	// example:
	//
	// Number of empty files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// emptyFileCount
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 3123
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileCount) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileCount) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileCount) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileCount) SetValue(v int64) *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio struct {
	// The description of the metric.
	//
	// example:
	//
	// Day growth ratio of empty files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// emptyFileCountDayGrowthRatio
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 0.01
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) SetValue(v float32) *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileDayGrowthCount struct {
	// The description of the metric.
	//
	// example:
	//
	// Day growth count of empty files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// emptyFileDayGrowthCount
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// -20
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileDayGrowthCount) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileDayGrowthCount) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileDayGrowthCount) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileDayGrowthCount) SetValue(v int64) *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileRatio struct {
	// The description of the metric.
	//
	// example:
	//
	// Ratio of empty files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// emptyFileRatio
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
	// 0.01
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileRatio) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileRatio) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileRatio) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileRatio) SetValue(v float32) *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsEmptyFileRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataDayGrowthSize struct {
	// The description of the metric.
	//
	// example:
	//
	// Day growth size of freeze data
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// freezeDataDayGrowthSize
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
	// 33229309
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataDayGrowthSize) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataDayGrowthSize) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataDayGrowthSize) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataDayGrowthSize) SetValue(v int64) *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataRatio struct {
	// The description of the metric.
	//
	// example:
	//
	// Ratio of freeze data
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// freezeDataRatio
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
	// 0.98
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataRatio) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataRatio) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataRatio) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataRatio) SetValue(v float32) *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataSize struct {
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
	// 33229309
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataSize) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataSize) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataSize) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataSize) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataSize) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataSize) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataSize) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataSize {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataSize) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataSize {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataSize) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataSize {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataSize) SetValue(v int64) *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataSize {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataSize) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio struct {
	// The description of the metric.
	//
	// example:
	//
	// Day growth ratio of freeze data size
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// freezeDataSizeDayGrowthRatio
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 0.01
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) SetValue(v float32) *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsHotDataDayGrowthSize struct {
	// The description of the metric.
	//
	// example:
	//
	// Day growth size of hot data
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// hotDataDayGrowthSize
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
	// 203431
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsHotDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsHotDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsHotDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsHotDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsHotDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsHotDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsHotDataDayGrowthSize) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsHotDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsHotDataDayGrowthSize) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsHotDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsHotDataDayGrowthSize) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsHotDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsHotDataDayGrowthSize) SetValue(v int64) *ListDoctorHiveTablesResponseBodyDataMetricsHotDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsHotDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsHotDataRatio struct {
	// The description of the metric.
	//
	// example:
	//
	// Hot data ratio
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// hotDataRatio
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
	// 0.01
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsHotDataRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsHotDataRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsHotDataRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsHotDataRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsHotDataRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsHotDataRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsHotDataRatio) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsHotDataRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsHotDataRatio) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsHotDataRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsHotDataRatio) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsHotDataRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsHotDataRatio) SetValue(v float32) *ListDoctorHiveTablesResponseBodyDataMetricsHotDataRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsHotDataRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsHotDataSize struct {
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
	// 203431
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsHotDataSize) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsHotDataSize) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsHotDataSize) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsHotDataSize) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsHotDataSize) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsHotDataSize) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsHotDataSize) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsHotDataSize {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsHotDataSize) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsHotDataSize {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsHotDataSize) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsHotDataSize {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsHotDataSize) SetValue(v int64) *ListDoctorHiveTablesResponseBodyDataMetricsHotDataSize {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsHotDataSize) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsHotDataSizeDayGrowthRatio struct {
	// The description of the metric.
	//
	// example:
	//
	// Day growth ratio of hot data size
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// hotDataSizeDayGrowthRatio
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 0.01
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsHotDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsHotDataSizeDayGrowthRatio) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsHotDataSizeDayGrowthRatio) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsHotDataSizeDayGrowthRatio) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsHotDataSizeDayGrowthRatio) SetValue(v float32) *ListDoctorHiveTablesResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsHotDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsLargeFileCount struct {
	// The description of the metric.
	//
	// example:
	//
	// Number of large files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// largeFileCount
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
	// 132
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsLargeFileCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsLargeFileCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileCount) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileCount) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileCount) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileCount) SetValue(v int64) *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsLargeFileCountDayGrowthRatio struct {
	// The description of the metric.
	//
	// example:
	//
	// Day growth ratio of large files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// largeFileCountDayGrowthRatio
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 0.01
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsLargeFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileCountDayGrowthRatio) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileCountDayGrowthRatio) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileCountDayGrowthRatio) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileCountDayGrowthRatio) SetValue(v float32) *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsLargeFileDayGrowthCount struct {
	// The description of the metric.
	//
	// example:
	//
	// Day growth count of large files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// largeFileDayGrowthCount
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
	// 40
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsLargeFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsLargeFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileDayGrowthCount) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileDayGrowthCount) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileDayGrowthCount) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileDayGrowthCount) SetValue(v int64) *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsLargeFileRatio struct {
	// The description of the metric.
	//
	// example:
	//
	// Ratio of large files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// largeFileRatio
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 0.02
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsLargeFileRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsLargeFileRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileRatio) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileRatio) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileRatio) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileRatio) SetValue(v float32) *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsLargeFileRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsMediumFileCount struct {
	// The description of the metric.
	//
	// example:
	//
	// Number of medium files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// mediumFileCount
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
	// 5
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsMediumFileCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsMediumFileCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileCount) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileCount) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileCount) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileCount) SetValue(v int64) *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsMediumFileCountDayGrowthRatio struct {
	// The description of the metric.
	//
	// example:
	//
	// Day growth ratio of medium files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// mediumFileCountDayGrowthRatio
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 0.01
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsMediumFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileCountDayGrowthRatio) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileCountDayGrowthRatio) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileCountDayGrowthRatio) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileCountDayGrowthRatio) SetValue(v float32) *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsMediumFileDayGrowthCount struct {
	// The description of the metric.
	//
	// example:
	//
	// Day growth count of medium files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// mediumFileDayGrowthCount
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
	// 20
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsMediumFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsMediumFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileDayGrowthCount) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileDayGrowthCount) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileDayGrowthCount) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileDayGrowthCount) SetValue(v int64) *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsMediumFileRatio struct {
	// The description of the metric.
	//
	// example:
	//
	// Ratio of medium files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// mediumFileRatio
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 0.8
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsMediumFileRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsMediumFileRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileRatio) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileRatio) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileRatio) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileRatio) SetValue(v float32) *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsMediumFileRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsPartitionNum struct {
	// The description of the metric.
	//
	// example:
	//
	// number of partitions
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// partitionNum
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
	// 331
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsPartitionNum) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsPartitionNum) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsPartitionNum) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsPartitionNum) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsPartitionNum) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsPartitionNum) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsPartitionNum) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsPartitionNum {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsPartitionNum) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsPartitionNum {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsPartitionNum) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsPartitionNum {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsPartitionNum) SetValue(v int64) *ListDoctorHiveTablesResponseBodyDataMetricsPartitionNum {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsPartitionNum) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsSmallFileCount struct {
	// The description of the metric.
	//
	// example:
	//
	// Number of small files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// smallFileCount
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// "“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 18
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsSmallFileCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsSmallFileCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileCount) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileCount) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileCount) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileCount) SetValue(v int64) *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsSmallFileCountDayGrowthRatio struct {
	// The description of the metric.
	//
	// example:
	//
	// Day growth ratio of small files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// smallFileCountDayGrowthRatio
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 0.01
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsSmallFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileCountDayGrowthRatio) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileCountDayGrowthRatio) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileCountDayGrowthRatio) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileCountDayGrowthRatio) SetValue(v float32) *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsSmallFileDayGrowthCount struct {
	// The description of the metric.
	//
	// example:
	//
	// Day growth count of small files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// smallFileDayGrowthCount
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 18
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsSmallFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsSmallFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileDayGrowthCount) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileDayGrowthCount) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileDayGrowthCount) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileDayGrowthCount) SetValue(v int64) *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsSmallFileRatio struct {
	// The description of the metric.
	//
	// example:
	//
	// Ratio of small files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// smallFileRatio
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 0.04
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsSmallFileRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsSmallFileRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileRatio) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileRatio) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileRatio) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileRatio) SetValue(v float32) *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsSmallFileRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsTinyFileCount struct {
	// The description of the metric.
	//
	// example:
	//
	// Number of tiny files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// tinyFileCount
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 451
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsTinyFileCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsTinyFileCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileCount) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileCount) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileCount) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileCount) SetValue(v int64) *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsTinyFileCountDayGrowthRatio struct {
	// The description of the metric.
	//
	// example:
	//
	// Day growth ratio of tiny files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// tinyFileCountDayGrowthRatio
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 0.04
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsTinyFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileCountDayGrowthRatio) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileCountDayGrowthRatio) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileCountDayGrowthRatio) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileCountDayGrowthRatio) SetValue(v float32) *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsTinyFileDayGrowthCount struct {
	// The description of the metric.
	//
	// example:
	//
	// Day growth count of tiny files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// tinyFileDayGrowthCount
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 482
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsTinyFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsTinyFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileDayGrowthCount) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileDayGrowthCount) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileDayGrowthCount) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileDayGrowthCount) SetValue(v int64) *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsTinyFileRatio struct {
	// The description of the metric.
	//
	// example:
	//
	// Ratio of tiny files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// tinyFileRatio
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
	// 0.96
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsTinyFileRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsTinyFileRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileRatio) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileRatio) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileRatio) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileRatio) SetValue(v float32) *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTinyFileRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsTotalDataDayGrowthSize struct {
	// The description of the metric.
	//
	// example:
	//
	// Day growth size of total data
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// totalDataDayGrowthSize
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
	// 33800296
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsTotalDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsTotalDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataDayGrowthSize) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataDayGrowthSize) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataDayGrowthSize) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataDayGrowthSize) SetValue(v int64) *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsTotalDataSize struct {
	// The description of the metric.
	//
	// example:
	//
	// Total data size in megabytes (MB)
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// totalDataSize
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
	// 33800296
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsTotalDataSize) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsTotalDataSize) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataSize) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataSize) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataSize) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataSize) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataSize) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataSize {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataSize) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataSize {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataSize) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataSize {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataSize) SetValue(v int64) *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataSize {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataSize) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio struct {
	// The description of the metric.
	//
	// example:
	//
	// Day growth ratio of total data size
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// totalDataSizeDayGrowthRatio
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 0.01
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) SetValue(v float32) *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsTotalFileCount struct {
	// The description of the metric.
	//
	// example:
	//
	// Number of total files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// totalFileCount
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 123
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsTotalFileCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsTotalFileCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileCount) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileCount) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileCount) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileCount) SetValue(v int64) *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsTotalFileCountDayGrowthRatio struct {
	// The description of the metric.
	//
	// example:
	//
	// Day growth ratio of total files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// totalFileCountDayGrowthRatio
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 0.01
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsTotalFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileCountDayGrowthRatio) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileCountDayGrowthRatio) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileCountDayGrowthRatio) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileCountDayGrowthRatio) SetValue(v float32) *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsTotalFileDayGrowthCount struct {
	// The description of the metric.
	//
	// example:
	//
	// Day growth count of total files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// totalFileDayGrowthCount
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
	// 100
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsTotalFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsTotalFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileDayGrowthCount) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileDayGrowthCount) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileDayGrowthCount) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileDayGrowthCount) SetValue(v int64) *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsTotalFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsWarmDataDayGrowthSize struct {
	// The description of the metric.
	//
	// example:
	//
	// Day growth size of warm data
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// warmDataDayGrowthSize
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
	// 149841
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsWarmDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsWarmDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataDayGrowthSize) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataDayGrowthSize) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataDayGrowthSize) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataDayGrowthSize) SetValue(v int64) *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsWarmDataRatio struct {
	// The description of the metric.
	//
	// example:
	//
	// Ratio of warm data
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// warmDataRatio
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
	// 0.1
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsWarmDataRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsWarmDataRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataRatio) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataRatio) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataRatio) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataRatio) SetValue(v float32) *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsWarmDataSize struct {
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
	// 14981
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsWarmDataSize) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsWarmDataSize) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataSize) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataSize) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataSize) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataSize) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataSize) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataSize {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataSize) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataSize {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataSize) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataSize {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataSize) SetValue(v int64) *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataSize {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataSize) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveTablesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio struct {
	// The description of the metric.
	//
	// example:
	//
	// Day growth ratio of warm data size
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// warmDataSizeDayGrowthRatio
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 0.01
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) SetDescription(v string) *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) SetName(v string) *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) SetUnit(v string) *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) SetValue(v float32) *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveTablesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}
