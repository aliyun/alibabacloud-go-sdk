// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorHiveClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDoctorHiveClusterResponseBodyData) *GetDoctorHiveClusterResponseBody
	GetData() *GetDoctorHiveClusterResponseBodyData
	SetRequestId(v string) *GetDoctorHiveClusterResponseBody
	GetRequestId() *string
}

type GetDoctorHiveClusterResponseBody struct {
	// The analysis results of the Hive cluster.
	Data *GetDoctorHiveClusterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDoctorHiveClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBody) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBody) GetData() *GetDoctorHiveClusterResponseBodyData {
	return s.Data
}

func (s *GetDoctorHiveClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDoctorHiveClusterResponseBody) SetData(v *GetDoctorHiveClusterResponseBodyData) *GetDoctorHiveClusterResponseBody {
	s.Data = v
	return s
}

func (s *GetDoctorHiveClusterResponseBody) SetRequestId(v string) *GetDoctorHiveClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyData struct {
	// The analysis results.
	Analysis *GetDoctorHiveClusterResponseBodyDataAnalysis `json:"Analysis,omitempty" xml:"Analysis,omitempty" type:"Struct"`
	// The information from the perspective of storage formats.
	Formats []*GetDoctorHiveClusterResponseBodyDataFormats `json:"Formats,omitempty" xml:"Formats,omitempty" type:"Repeated"`
	// The metric information.
	Metrics *GetDoctorHiveClusterResponseBodyDataMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Struct"`
}

func (s GetDoctorHiveClusterResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyData) GetAnalysis() *GetDoctorHiveClusterResponseBodyDataAnalysis {
	return s.Analysis
}

func (s *GetDoctorHiveClusterResponseBodyData) GetFormats() []*GetDoctorHiveClusterResponseBodyDataFormats {
	return s.Formats
}

func (s *GetDoctorHiveClusterResponseBodyData) GetMetrics() *GetDoctorHiveClusterResponseBodyDataMetrics {
	return s.Metrics
}

func (s *GetDoctorHiveClusterResponseBodyData) SetAnalysis(v *GetDoctorHiveClusterResponseBodyDataAnalysis) *GetDoctorHiveClusterResponseBodyData {
	s.Analysis = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyData) SetFormats(v []*GetDoctorHiveClusterResponseBodyDataFormats) *GetDoctorHiveClusterResponseBodyData {
	s.Formats = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyData) SetMetrics(v *GetDoctorHiveClusterResponseBodyDataMetrics) *GetDoctorHiveClusterResponseBodyData {
	s.Metrics = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataAnalysis struct {
	// The score for the distribution of files of different sizes stored in the Hive cluster.
	//
	// example:
	//
	// 80
	HiveDistributionScore *int32 `json:"HiveDistributionScore,omitempty" xml:"HiveDistributionScore,omitempty"`
	// The score for the distribution of files stored in different formats in the Hive cluster.
	//
	// example:
	//
	// 80
	HiveFormatScore *int32 `json:"HiveFormatScore,omitempty" xml:"HiveFormatScore,omitempty"`
	// The score for the access frequency of the Hive cluster.
	//
	// example:
	//
	// 80
	HiveFrequencyScore *int32 `json:"HiveFrequencyScore,omitempty" xml:"HiveFrequencyScore,omitempty"`
	// The overall score of the Hive cluster.
	//
	// example:
	//
	// 80
	HiveScore *int32 `json:"HiveScore,omitempty" xml:"HiveScore,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataAnalysis) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataAnalysis) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataAnalysis) GetHiveDistributionScore() *int32 {
	return s.HiveDistributionScore
}

func (s *GetDoctorHiveClusterResponseBodyDataAnalysis) GetHiveFormatScore() *int32 {
	return s.HiveFormatScore
}

func (s *GetDoctorHiveClusterResponseBodyDataAnalysis) GetHiveFrequencyScore() *int32 {
	return s.HiveFrequencyScore
}

func (s *GetDoctorHiveClusterResponseBodyDataAnalysis) GetHiveScore() *int32 {
	return s.HiveScore
}

func (s *GetDoctorHiveClusterResponseBodyDataAnalysis) SetHiveDistributionScore(v int32) *GetDoctorHiveClusterResponseBodyDataAnalysis {
	s.HiveDistributionScore = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataAnalysis) SetHiveFormatScore(v int32) *GetDoctorHiveClusterResponseBodyDataAnalysis {
	s.HiveFormatScore = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataAnalysis) SetHiveFrequencyScore(v int32) *GetDoctorHiveClusterResponseBodyDataAnalysis {
	s.HiveFrequencyScore = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataAnalysis) SetHiveScore(v int32) *GetDoctorHiveClusterResponseBodyDataAnalysis {
	s.HiveScore = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataAnalysis) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataFormats struct {
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
	// 0.5
	FormatRatio *float32 `json:"FormatRatio,omitempty" xml:"FormatRatio,omitempty"`
	// The amount of data in the format.
	//
	// example:
	//
	// 100
	FormatSize *int64 `json:"FormatSize,omitempty" xml:"FormatSize,omitempty"`
	// The unit of the amount of data in the format.
	//
	// example:
	//
	// MB
	FormatSizeUnit *string `json:"FormatSizeUnit,omitempty" xml:"FormatSizeUnit,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataFormats) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataFormats) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataFormats) GetFormatName() *string {
	return s.FormatName
}

func (s *GetDoctorHiveClusterResponseBodyDataFormats) GetFormatRatio() *float32 {
	return s.FormatRatio
}

func (s *GetDoctorHiveClusterResponseBodyDataFormats) GetFormatSize() *int64 {
	return s.FormatSize
}

func (s *GetDoctorHiveClusterResponseBodyDataFormats) GetFormatSizeUnit() *string {
	return s.FormatSizeUnit
}

func (s *GetDoctorHiveClusterResponseBodyDataFormats) SetFormatName(v string) *GetDoctorHiveClusterResponseBodyDataFormats {
	s.FormatName = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataFormats) SetFormatRatio(v float32) *GetDoctorHiveClusterResponseBodyDataFormats {
	s.FormatRatio = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataFormats) SetFormatSize(v int64) *GetDoctorHiveClusterResponseBodyDataFormats {
	s.FormatSize = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataFormats) SetFormatSizeUnit(v string) *GetDoctorHiveClusterResponseBodyDataFormats {
	s.FormatSizeUnit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataFormats) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetrics struct {
	// The daily increment of the amount of cold data. Cold data refers to data that is not accessed for more than 30 days but is accessed in previous 90 days.
	ColdDataDayGrowthSize *GetDoctorHiveClusterResponseBodyDataMetricsColdDataDayGrowthSize `json:"ColdDataDayGrowthSize,omitempty" xml:"ColdDataDayGrowthSize,omitempty" type:"Struct"`
	// The proportion of cold data. Cold data refers to data that is not accessed for more than 30 days but is accessed in previous 90 days.
	ColdDataRatio *GetDoctorHiveClusterResponseBodyDataMetricsColdDataRatio `json:"ColdDataRatio,omitempty" xml:"ColdDataRatio,omitempty" type:"Struct"`
	// The amount of cold data. Cold data refers to data that is not accessed for more than 30 days but is accessed in previous 90 days.
	ColdDataSize *GetDoctorHiveClusterResponseBodyDataMetricsColdDataSize `json:"ColdDataSize,omitempty" xml:"ColdDataSize,omitempty" type:"Struct"`
	// The day-to-day growth rate of the amount of cold data. Cold data refers to data that is not accessed for more than 30 days but is accessed in previous 90 days.
	ColdDataSizeDayGrowthRatio *GetDoctorHiveClusterResponseBodyDataMetricsColdDataSizeDayGrowthRatio `json:"ColdDataSizeDayGrowthRatio,omitempty" xml:"ColdDataSizeDayGrowthRatio,omitempty" type:"Struct"`
	// The number of databases.
	DatabaseCount *GetDoctorHiveClusterResponseBodyDataMetricsDatabaseCount `json:"DatabaseCount,omitempty" xml:"DatabaseCount,omitempty" type:"Struct"`
	// The number of empty files. Empty files are those with a size of 0 MB.
	EmptyFileCount *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileCount `json:"EmptyFileCount,omitempty" xml:"EmptyFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the number of empty files. Empty files are those with a size of 0 MB.
	EmptyFileCountDayGrowthRatio *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileCountDayGrowthRatio `json:"EmptyFileCountDayGrowthRatio,omitempty" xml:"EmptyFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the number of empty files. Empty files are those with a size of 0 MB.
	EmptyFileDayGrowthCount *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileDayGrowthCount `json:"EmptyFileDayGrowthCount,omitempty" xml:"EmptyFileDayGrowthCount,omitempty" type:"Struct"`
	// The proportion of empty files. Empty files are those with a size of 0 MB.
	EmptyFileRatio *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileRatio `json:"EmptyFileRatio,omitempty" xml:"EmptyFileRatio,omitempty" type:"Struct"`
	// The daily increment of the amount of very cold data. Very cold data refers to data that is not accessed for more than 90 days.
	FreezeDataDayGrowthSize *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataDayGrowthSize `json:"FreezeDataDayGrowthSize,omitempty" xml:"FreezeDataDayGrowthSize,omitempty" type:"Struct"`
	// The proportion of very cold data. Very cold data refers to data that is not accessed for more than 90 days.
	FreezeDataRatio *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataRatio `json:"FreezeDataRatio,omitempty" xml:"FreezeDataRatio,omitempty" type:"Struct"`
	// The amount of very cold data. Very cold data refers to data that is not accessed for more than 90 days.
	FreezeDataSize *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataSize `json:"FreezeDataSize,omitempty" xml:"FreezeDataSize,omitempty" type:"Struct"`
	// The day-to-day growth rate of the amount of very cold data. Very cold data refers to data that is not accessed for more than 90 days.
	FreezeDataSizeDayGrowthRatio *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio `json:"FreezeDataSizeDayGrowthRatio,omitempty" xml:"FreezeDataSizeDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the amount of hot data. Hot data refers to data that is accessed in previous seven days.
	HotDataDayGrowthSize *GetDoctorHiveClusterResponseBodyDataMetricsHotDataDayGrowthSize `json:"HotDataDayGrowthSize,omitempty" xml:"HotDataDayGrowthSize,omitempty" type:"Struct"`
	// The proportion of hot data. Hot data refers to data that is accessed in previous seven days.
	HotDataRatio *GetDoctorHiveClusterResponseBodyDataMetricsHotDataRatio `json:"HotDataRatio,omitempty" xml:"HotDataRatio,omitempty" type:"Struct"`
	// The amount of hot data. Hot data refers to data that is accessed in previous seven days.
	HotDataSize *GetDoctorHiveClusterResponseBodyDataMetricsHotDataSize `json:"HotDataSize,omitempty" xml:"HotDataSize,omitempty" type:"Struct"`
	// The day-to-day growth rate of the amount of hot data. Hot data refers to data that is accessed in previous seven days.
	HotDataSizeDayGrowthRatio *GetDoctorHiveClusterResponseBodyDataMetricsHotDataSizeDayGrowthRatio `json:"HotDataSizeDayGrowthRatio,omitempty" xml:"HotDataSizeDayGrowthRatio,omitempty" type:"Struct"`
	// The number of large files. Large files are those with a size greater than 1 GB.
	LargeFileCount *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileCount `json:"LargeFileCount,omitempty" xml:"LargeFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the number of large files. Large files are those with a size greater than 1 GB.
	LargeFileCountDayGrowthRatio *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileCountDayGrowthRatio `json:"LargeFileCountDayGrowthRatio,omitempty" xml:"LargeFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the number of large files. Large files are those with a size greater than 1 GB.
	LargeFileDayGrowthCount *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileDayGrowthCount `json:"LargeFileDayGrowthCount,omitempty" xml:"LargeFileDayGrowthCount,omitempty" type:"Struct"`
	// The proportion of large files. Large files are those with a size greater than 1 GB.
	LargeFileRatio *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileRatio `json:"LargeFileRatio,omitempty" xml:"LargeFileRatio,omitempty" type:"Struct"`
	// The number of medium files. Medium files are those with a size greater than or equal to 128 MB and less than or equal to 1 GB.
	MediumFileCount *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileCount `json:"MediumFileCount,omitempty" xml:"MediumFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the number of medium files. Medium files are those with a size greater than or equal to 128 MB and less than or equal to 1 GB.
	MediumFileCountDayGrowthRatio *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileCountDayGrowthRatio `json:"MediumFileCountDayGrowthRatio,omitempty" xml:"MediumFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the number of medium files. Medium files are those with a size greater than or equal to 128 MB and less than or equal to 1 GB.
	MediumFileDayGrowthCount *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileDayGrowthCount `json:"MediumFileDayGrowthCount,omitempty" xml:"MediumFileDayGrowthCount,omitempty" type:"Struct"`
	// The proportion of medium files. Medium files are those with a size greater than or equal to 128 MB and less than or equal to 1 GB.
	MediumFileRatio *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileRatio `json:"MediumFileRatio,omitempty" xml:"MediumFileRatio,omitempty" type:"Struct"`
	// The number of partitions.
	PartitionNum *GetDoctorHiveClusterResponseBodyDataMetricsPartitionNum `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty" type:"Struct"`
	// The number of small files. Small files are those with a size greater than or equal to 10 MB and less than 128 MB.
	SmallFileCount *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileCount `json:"SmallFileCount,omitempty" xml:"SmallFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the number of small files. Small files are those with a size greater than or equal to 10 MB and less than 128 MB.
	SmallFileCountDayGrowthRatio *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileCountDayGrowthRatio `json:"SmallFileCountDayGrowthRatio,omitempty" xml:"SmallFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the number of small files. Small files are those with a size greater than or equal to 10 MB and less than 128 MB.
	SmallFileDayGrowthCount *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileDayGrowthCount `json:"SmallFileDayGrowthCount,omitempty" xml:"SmallFileDayGrowthCount,omitempty" type:"Struct"`
	// The proportion of small files. Small files are those with a size greater than or equal to 10 MB and less than 128 MB.
	SmallFileRatio *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileRatio `json:"SmallFileRatio,omitempty" xml:"SmallFileRatio,omitempty" type:"Struct"`
	// The number of tables.
	TableCount *GetDoctorHiveClusterResponseBodyDataMetricsTableCount `json:"TableCount,omitempty" xml:"TableCount,omitempty" type:"Struct"`
	// The number of very small files. Very small files are those with a size greater than 0 MB and less than 10 MB.
	TinyFileCount *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileCount `json:"TinyFileCount,omitempty" xml:"TinyFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the number of very small files. Very small files are those with a size greater than 0 MB and less than 10 MB.
	TinyFileCountDayGrowthRatio *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileCountDayGrowthRatio `json:"TinyFileCountDayGrowthRatio,omitempty" xml:"TinyFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the number of very small files. Very small files are those with a size greater than 0 MB and less than 10 MB.
	TinyFileDayGrowthCount *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileDayGrowthCount `json:"TinyFileDayGrowthCount,omitempty" xml:"TinyFileDayGrowthCount,omitempty" type:"Struct"`
	// The proportion of very small files. Very small files are those with a size greater than 0 MB and less than 10 MB.
	TinyFileRatio *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileRatio `json:"TinyFileRatio,omitempty" xml:"TinyFileRatio,omitempty" type:"Struct"`
	// The daily incremental of the amount of total data.
	TotalDataDayGrowthSize *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataDayGrowthSize `json:"TotalDataDayGrowthSize,omitempty" xml:"TotalDataDayGrowthSize,omitempty" type:"Struct"`
	// The total amount of data.
	TotalDataSize *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataSize `json:"TotalDataSize,omitempty" xml:"TotalDataSize,omitempty" type:"Struct"`
	// The day-to-day growth rate of the total data volume.
	TotalDataSizeDayGrowthRatio *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataSizeDayGrowthRatio `json:"TotalDataSizeDayGrowthRatio,omitempty" xml:"TotalDataSizeDayGrowthRatio,omitempty" type:"Struct"`
	// The total number of files.
	TotalFileCount *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileCount `json:"TotalFileCount,omitempty" xml:"TotalFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the total number of files.
	TotalFileCountDayGrowthRatio *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileCountDayGrowthRatio `json:"TotalFileCountDayGrowthRatio,omitempty" xml:"TotalFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the total number of files.
	TotalFileDayGrowthCount *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileDayGrowthCount `json:"TotalFileDayGrowthCount,omitempty" xml:"TotalFileDayGrowthCount,omitempty" type:"Struct"`
	// The daily increment of the amount of warm data. Warm data refers to data that is not accessed for more than 7 days but is accessed in previous 30 days.
	WarmDataDayGrowthSize *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataDayGrowthSize `json:"WarmDataDayGrowthSize,omitempty" xml:"WarmDataDayGrowthSize,omitempty" type:"Struct"`
	// The proportion of warm data. Warm data refers to data that is not accessed for more than 7 days but is accessed in previous 30 days.
	WarmDataRatio *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataRatio `json:"WarmDataRatio,omitempty" xml:"WarmDataRatio,omitempty" type:"Struct"`
	// The amount of warm data. Warm data refers to data that is not accessed for more than 7 days but is accessed in previous 30 days.
	WarmDataSize *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataSize `json:"WarmDataSize,omitempty" xml:"WarmDataSize,omitempty" type:"Struct"`
	// The day-to-day growth rate of the amount of warm data. Warm data refers to data that is not accessed for more than 7 days but is accessed in previous 30 days.
	WarmDataSizeDayGrowthRatio *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataSizeDayGrowthRatio `json:"WarmDataSizeDayGrowthRatio,omitempty" xml:"WarmDataSizeDayGrowthRatio,omitempty" type:"Struct"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetrics) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetrics) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetColdDataDayGrowthSize() *GetDoctorHiveClusterResponseBodyDataMetricsColdDataDayGrowthSize {
	return s.ColdDataDayGrowthSize
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetColdDataRatio() *GetDoctorHiveClusterResponseBodyDataMetricsColdDataRatio {
	return s.ColdDataRatio
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetColdDataSize() *GetDoctorHiveClusterResponseBodyDataMetricsColdDataSize {
	return s.ColdDataSize
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetColdDataSizeDayGrowthRatio() *GetDoctorHiveClusterResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	return s.ColdDataSizeDayGrowthRatio
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetDatabaseCount() *GetDoctorHiveClusterResponseBodyDataMetricsDatabaseCount {
	return s.DatabaseCount
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetEmptyFileCount() *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileCount {
	return s.EmptyFileCount
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetEmptyFileCountDayGrowthRatio() *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	return s.EmptyFileCountDayGrowthRatio
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetEmptyFileDayGrowthCount() *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileDayGrowthCount {
	return s.EmptyFileDayGrowthCount
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetEmptyFileRatio() *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileRatio {
	return s.EmptyFileRatio
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetFreezeDataDayGrowthSize() *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataDayGrowthSize {
	return s.FreezeDataDayGrowthSize
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetFreezeDataRatio() *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataRatio {
	return s.FreezeDataRatio
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetFreezeDataSize() *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataSize {
	return s.FreezeDataSize
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetFreezeDataSizeDayGrowthRatio() *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	return s.FreezeDataSizeDayGrowthRatio
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetHotDataDayGrowthSize() *GetDoctorHiveClusterResponseBodyDataMetricsHotDataDayGrowthSize {
	return s.HotDataDayGrowthSize
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetHotDataRatio() *GetDoctorHiveClusterResponseBodyDataMetricsHotDataRatio {
	return s.HotDataRatio
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetHotDataSize() *GetDoctorHiveClusterResponseBodyDataMetricsHotDataSize {
	return s.HotDataSize
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetHotDataSizeDayGrowthRatio() *GetDoctorHiveClusterResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	return s.HotDataSizeDayGrowthRatio
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetLargeFileCount() *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileCount {
	return s.LargeFileCount
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetLargeFileCountDayGrowthRatio() *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	return s.LargeFileCountDayGrowthRatio
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetLargeFileDayGrowthCount() *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileDayGrowthCount {
	return s.LargeFileDayGrowthCount
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetLargeFileRatio() *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileRatio {
	return s.LargeFileRatio
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetMediumFileCount() *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileCount {
	return s.MediumFileCount
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetMediumFileCountDayGrowthRatio() *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	return s.MediumFileCountDayGrowthRatio
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetMediumFileDayGrowthCount() *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileDayGrowthCount {
	return s.MediumFileDayGrowthCount
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetMediumFileRatio() *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileRatio {
	return s.MediumFileRatio
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetPartitionNum() *GetDoctorHiveClusterResponseBodyDataMetricsPartitionNum {
	return s.PartitionNum
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetSmallFileCount() *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileCount {
	return s.SmallFileCount
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetSmallFileCountDayGrowthRatio() *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	return s.SmallFileCountDayGrowthRatio
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetSmallFileDayGrowthCount() *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileDayGrowthCount {
	return s.SmallFileDayGrowthCount
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetSmallFileRatio() *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileRatio {
	return s.SmallFileRatio
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetTableCount() *GetDoctorHiveClusterResponseBodyDataMetricsTableCount {
	return s.TableCount
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetTinyFileCount() *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileCount {
	return s.TinyFileCount
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetTinyFileCountDayGrowthRatio() *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	return s.TinyFileCountDayGrowthRatio
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetTinyFileDayGrowthCount() *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileDayGrowthCount {
	return s.TinyFileDayGrowthCount
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetTinyFileRatio() *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileRatio {
	return s.TinyFileRatio
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetTotalDataDayGrowthSize() *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataDayGrowthSize {
	return s.TotalDataDayGrowthSize
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetTotalDataSize() *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataSize {
	return s.TotalDataSize
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetTotalDataSizeDayGrowthRatio() *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	return s.TotalDataSizeDayGrowthRatio
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetTotalFileCount() *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileCount {
	return s.TotalFileCount
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetTotalFileCountDayGrowthRatio() *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	return s.TotalFileCountDayGrowthRatio
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetTotalFileDayGrowthCount() *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileDayGrowthCount {
	return s.TotalFileDayGrowthCount
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetWarmDataDayGrowthSize() *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataDayGrowthSize {
	return s.WarmDataDayGrowthSize
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetWarmDataRatio() *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataRatio {
	return s.WarmDataRatio
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetWarmDataSize() *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataSize {
	return s.WarmDataSize
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) GetWarmDataSizeDayGrowthRatio() *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	return s.WarmDataSizeDayGrowthRatio
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetColdDataDayGrowthSize(v *GetDoctorHiveClusterResponseBodyDataMetricsColdDataDayGrowthSize) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.ColdDataDayGrowthSize = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetColdDataRatio(v *GetDoctorHiveClusterResponseBodyDataMetricsColdDataRatio) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.ColdDataRatio = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetColdDataSize(v *GetDoctorHiveClusterResponseBodyDataMetricsColdDataSize) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.ColdDataSize = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetColdDataSizeDayGrowthRatio(v *GetDoctorHiveClusterResponseBodyDataMetricsColdDataSizeDayGrowthRatio) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.ColdDataSizeDayGrowthRatio = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetDatabaseCount(v *GetDoctorHiveClusterResponseBodyDataMetricsDatabaseCount) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.DatabaseCount = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetEmptyFileCount(v *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileCount) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.EmptyFileCount = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetEmptyFileCountDayGrowthRatio(v *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.EmptyFileCountDayGrowthRatio = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetEmptyFileDayGrowthCount(v *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileDayGrowthCount) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.EmptyFileDayGrowthCount = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetEmptyFileRatio(v *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileRatio) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.EmptyFileRatio = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetFreezeDataDayGrowthSize(v *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataDayGrowthSize) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.FreezeDataDayGrowthSize = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetFreezeDataRatio(v *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataRatio) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.FreezeDataRatio = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetFreezeDataSize(v *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataSize) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.FreezeDataSize = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetFreezeDataSizeDayGrowthRatio(v *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.FreezeDataSizeDayGrowthRatio = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetHotDataDayGrowthSize(v *GetDoctorHiveClusterResponseBodyDataMetricsHotDataDayGrowthSize) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.HotDataDayGrowthSize = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetHotDataRatio(v *GetDoctorHiveClusterResponseBodyDataMetricsHotDataRatio) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.HotDataRatio = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetHotDataSize(v *GetDoctorHiveClusterResponseBodyDataMetricsHotDataSize) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.HotDataSize = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetHotDataSizeDayGrowthRatio(v *GetDoctorHiveClusterResponseBodyDataMetricsHotDataSizeDayGrowthRatio) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.HotDataSizeDayGrowthRatio = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetLargeFileCount(v *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileCount) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.LargeFileCount = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetLargeFileCountDayGrowthRatio(v *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileCountDayGrowthRatio) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.LargeFileCountDayGrowthRatio = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetLargeFileDayGrowthCount(v *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileDayGrowthCount) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.LargeFileDayGrowthCount = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetLargeFileRatio(v *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileRatio) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.LargeFileRatio = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetMediumFileCount(v *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileCount) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.MediumFileCount = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetMediumFileCountDayGrowthRatio(v *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileCountDayGrowthRatio) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.MediumFileCountDayGrowthRatio = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetMediumFileDayGrowthCount(v *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileDayGrowthCount) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.MediumFileDayGrowthCount = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetMediumFileRatio(v *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileRatio) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.MediumFileRatio = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetPartitionNum(v *GetDoctorHiveClusterResponseBodyDataMetricsPartitionNum) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.PartitionNum = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetSmallFileCount(v *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileCount) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.SmallFileCount = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetSmallFileCountDayGrowthRatio(v *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileCountDayGrowthRatio) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.SmallFileCountDayGrowthRatio = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetSmallFileDayGrowthCount(v *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileDayGrowthCount) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.SmallFileDayGrowthCount = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetSmallFileRatio(v *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileRatio) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.SmallFileRatio = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetTableCount(v *GetDoctorHiveClusterResponseBodyDataMetricsTableCount) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.TableCount = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetTinyFileCount(v *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileCount) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.TinyFileCount = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetTinyFileCountDayGrowthRatio(v *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileCountDayGrowthRatio) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.TinyFileCountDayGrowthRatio = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetTinyFileDayGrowthCount(v *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileDayGrowthCount) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.TinyFileDayGrowthCount = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetTinyFileRatio(v *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileRatio) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.TinyFileRatio = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetTotalDataDayGrowthSize(v *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataDayGrowthSize) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.TotalDataDayGrowthSize = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetTotalDataSize(v *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataSize) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.TotalDataSize = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetTotalDataSizeDayGrowthRatio(v *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.TotalDataSizeDayGrowthRatio = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetTotalFileCount(v *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileCount) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.TotalFileCount = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetTotalFileCountDayGrowthRatio(v *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileCountDayGrowthRatio) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.TotalFileCountDayGrowthRatio = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetTotalFileDayGrowthCount(v *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileDayGrowthCount) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.TotalFileDayGrowthCount = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetWarmDataDayGrowthSize(v *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataDayGrowthSize) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.WarmDataDayGrowthSize = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetWarmDataRatio(v *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataRatio) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.WarmDataRatio = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetWarmDataSize(v *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataSize) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.WarmDataSize = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) SetWarmDataSizeDayGrowthRatio(v *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) *GetDoctorHiveClusterResponseBodyDataMetrics {
	s.WarmDataSizeDayGrowthRatio = v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetrics) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsColdDataDayGrowthSize struct {
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
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsColdDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsColdDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsColdDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsColdDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsColdDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsColdDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsColdDataDayGrowthSize) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsColdDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsColdDataDayGrowthSize) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsColdDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsColdDataDayGrowthSize) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsColdDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsColdDataDayGrowthSize) SetValue(v int64) *GetDoctorHiveClusterResponseBodyDataMetricsColdDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsColdDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsColdDataRatio struct {
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
	// 0.5
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsColdDataRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsColdDataRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsColdDataRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsColdDataRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsColdDataRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsColdDataRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsColdDataRatio) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsColdDataRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsColdDataRatio) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsColdDataRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsColdDataRatio) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsColdDataRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsColdDataRatio) SetValue(v float32) *GetDoctorHiveClusterResponseBodyDataMetricsColdDataRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsColdDataRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsColdDataSize struct {
	// The description of the metric.
	//
	// example:
	//
	// Ratio of cold files
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
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsColdDataSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsColdDataSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsColdDataSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsColdDataSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsColdDataSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsColdDataSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsColdDataSize) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsColdDataSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsColdDataSize) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsColdDataSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsColdDataSize) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsColdDataSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsColdDataSize) SetValue(v int64) *GetDoctorHiveClusterResponseBodyDataMetricsColdDataSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsColdDataSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsColdDataSizeDayGrowthRatio struct {
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
	// 0.5
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsColdDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsColdDataSizeDayGrowthRatio) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsColdDataSizeDayGrowthRatio) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsColdDataSizeDayGrowthRatio) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsColdDataSizeDayGrowthRatio) SetValue(v float32) *GetDoctorHiveClusterResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsColdDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsDatabaseCount struct {
	// The description of the metric.
	//
	// example:
	//
	// Number of databases
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// databaseCount
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

func (s GetDoctorHiveClusterResponseBodyDataMetricsDatabaseCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsDatabaseCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsDatabaseCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsDatabaseCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsDatabaseCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsDatabaseCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsDatabaseCount) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsDatabaseCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsDatabaseCount) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsDatabaseCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsDatabaseCount) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsDatabaseCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsDatabaseCount) SetValue(v int64) *GetDoctorHiveClusterResponseBodyDataMetricsDatabaseCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsDatabaseCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileCount struct {
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
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileCount) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileCount) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileCount) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileCount) SetValue(v int64) *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileCountDayGrowthRatio struct {
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
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 0.5
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) SetValue(v float32) *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileDayGrowthCount struct {
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
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileDayGrowthCount) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileDayGrowthCount) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileDayGrowthCount) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileDayGrowthCount) SetValue(v int64) *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileRatio struct {
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
	// 0.5
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileRatio) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileRatio) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileRatio) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileRatio) SetValue(v float32) *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsEmptyFileRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataDayGrowthSize struct {
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
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataDayGrowthSize) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataDayGrowthSize) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataDayGrowthSize) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataDayGrowthSize) SetValue(v int64) *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataRatio struct {
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
	// 0.12
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataRatio) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataRatio) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataRatio) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataRatio) SetValue(v float32) *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataSize struct {
	// The description of the metric.
	//
	// example:
	//
	// Ratio of freeze files
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
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataSize) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataSize) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataSize) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataSize) SetValue(v int64) *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio struct {
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
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 0.5
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) SetValue(v float32) *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsHotDataDayGrowthSize struct {
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
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsHotDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsHotDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsHotDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsHotDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsHotDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsHotDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsHotDataDayGrowthSize) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsHotDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsHotDataDayGrowthSize) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsHotDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsHotDataDayGrowthSize) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsHotDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsHotDataDayGrowthSize) SetValue(v int64) *GetDoctorHiveClusterResponseBodyDataMetricsHotDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsHotDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsHotDataRatio struct {
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
	// 0.5
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsHotDataRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsHotDataRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsHotDataRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsHotDataRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsHotDataRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsHotDataRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsHotDataRatio) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsHotDataRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsHotDataRatio) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsHotDataRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsHotDataRatio) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsHotDataRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsHotDataRatio) SetValue(v float32) *GetDoctorHiveClusterResponseBodyDataMetricsHotDataRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsHotDataRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsHotDataSize struct {
	// The description of the metric.
	//
	// example:
	//
	// Ratio of hot files
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
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsHotDataSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsHotDataSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsHotDataSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsHotDataSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsHotDataSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsHotDataSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsHotDataSize) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsHotDataSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsHotDataSize) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsHotDataSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsHotDataSize) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsHotDataSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsHotDataSize) SetValue(v int64) *GetDoctorHiveClusterResponseBodyDataMetricsHotDataSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsHotDataSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsHotDataSizeDayGrowthRatio struct {
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
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 0.5
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsHotDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsHotDataSizeDayGrowthRatio) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsHotDataSizeDayGrowthRatio) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsHotDataSizeDayGrowthRatio) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsHotDataSizeDayGrowthRatio) SetValue(v float32) *GetDoctorHiveClusterResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsHotDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsLargeFileCount struct {
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
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsLargeFileCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsLargeFileCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileCount) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileCount) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileCount) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileCount) SetValue(v int64) *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsLargeFileCountDayGrowthRatio struct {
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
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 0.5
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsLargeFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileCountDayGrowthRatio) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileCountDayGrowthRatio) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileCountDayGrowthRatio) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileCountDayGrowthRatio) SetValue(v float32) *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsLargeFileDayGrowthCount struct {
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
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsLargeFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsLargeFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileDayGrowthCount) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileDayGrowthCount) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileDayGrowthCount) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileDayGrowthCount) SetValue(v int64) *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsLargeFileRatio struct {
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
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 0.5
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsLargeFileRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsLargeFileRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileRatio) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileRatio) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileRatio) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileRatio) SetValue(v float32) *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsLargeFileRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsMediumFileCount struct {
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
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsMediumFileCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsMediumFileCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileCount) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileCount) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileCount) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileCount) SetValue(v int64) *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsMediumFileCountDayGrowthRatio struct {
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
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 0.5
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsMediumFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileCountDayGrowthRatio) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileCountDayGrowthRatio) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileCountDayGrowthRatio) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileCountDayGrowthRatio) SetValue(v float32) *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsMediumFileDayGrowthCount struct {
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
	// “”
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 178
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsMediumFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsMediumFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileDayGrowthCount) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileDayGrowthCount) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileDayGrowthCount) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileDayGrowthCount) SetValue(v int64) *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsMediumFileRatio struct {
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
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 0.5
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsMediumFileRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsMediumFileRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileRatio) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileRatio) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileRatio) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileRatio) SetValue(v float32) *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsMediumFileRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsPartitionNum struct {
	// The description of the metric.
	//
	// example:
	//
	// Number of partitions
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
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsPartitionNum) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsPartitionNum) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsPartitionNum) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsPartitionNum) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsPartitionNum) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsPartitionNum) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsPartitionNum) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsPartitionNum {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsPartitionNum) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsPartitionNum {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsPartitionNum) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsPartitionNum {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsPartitionNum) SetValue(v int64) *GetDoctorHiveClusterResponseBodyDataMetricsPartitionNum {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsPartitionNum) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsSmallFileCount struct {
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
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsSmallFileCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsSmallFileCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileCount) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileCount) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileCount) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileCount) SetValue(v int64) *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsSmallFileCountDayGrowthRatio struct {
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
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 0.5
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsSmallFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileCountDayGrowthRatio) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileCountDayGrowthRatio) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileCountDayGrowthRatio) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileCountDayGrowthRatio) SetValue(v float32) *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsSmallFileDayGrowthCount struct {
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
	// “”
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsSmallFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsSmallFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileDayGrowthCount) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileDayGrowthCount) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileDayGrowthCount) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileDayGrowthCount) SetValue(v int64) *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsSmallFileRatio struct {
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
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 0.5
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsSmallFileRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsSmallFileRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileRatio) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileRatio) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileRatio) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileRatio) SetValue(v float32) *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsSmallFileRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsTableCount struct {
	// The description of the metric.
	//
	// example:
	//
	// Number of tables
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// tableCount
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

func (s GetDoctorHiveClusterResponseBodyDataMetricsTableCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsTableCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTableCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTableCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTableCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTableCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTableCount) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsTableCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTableCount) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsTableCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTableCount) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsTableCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTableCount) SetValue(v int64) *GetDoctorHiveClusterResponseBodyDataMetricsTableCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTableCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsTinyFileCount struct {
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
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsTinyFileCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsTinyFileCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileCount) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileCount) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileCount) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileCount) SetValue(v int64) *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsTinyFileCountDayGrowthRatio struct {
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
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 0.5
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsTinyFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileCountDayGrowthRatio) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileCountDayGrowthRatio) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileCountDayGrowthRatio) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileCountDayGrowthRatio) SetValue(v float32) *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsTinyFileDayGrowthCount struct {
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
	// day growth count of tiny files
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

func (s GetDoctorHiveClusterResponseBodyDataMetricsTinyFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsTinyFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileDayGrowthCount) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileDayGrowthCount) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileDayGrowthCount) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileDayGrowthCount) SetValue(v int64) *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsTinyFileRatio struct {
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
	// 0.5
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsTinyFileRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsTinyFileRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileRatio) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileRatio) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileRatio) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileRatio) SetValue(v float32) *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTinyFileRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsTotalDataDayGrowthSize struct {
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
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsTotalDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsTotalDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataDayGrowthSize) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataDayGrowthSize) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataDayGrowthSize) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataDayGrowthSize) SetValue(v int64) *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsTotalDataSize struct {
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
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsTotalDataSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsTotalDataSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataSize) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataSize) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataSize) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataSize) SetValue(v int64) *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsTotalDataSizeDayGrowthRatio struct {
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
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 0.5
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) SetValue(v float32) *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsTotalFileCount struct {
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
	// totalFileCount
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

func (s GetDoctorHiveClusterResponseBodyDataMetricsTotalFileCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsTotalFileCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileCount) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileCount) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileCount) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileCount) SetValue(v int64) *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsTotalFileCountDayGrowthRatio struct {
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
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 0.5
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsTotalFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileCountDayGrowthRatio) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileCountDayGrowthRatio) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileCountDayGrowthRatio) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileCountDayGrowthRatio) SetValue(v float32) *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsTotalFileDayGrowthCount struct {
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
	// 27800
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsTotalFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsTotalFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileDayGrowthCount) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileDayGrowthCount) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileDayGrowthCount) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileDayGrowthCount) SetValue(v int64) *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsTotalFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsWarmDataDayGrowthSize struct {
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
	// -100
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsWarmDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsWarmDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataDayGrowthSize) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataDayGrowthSize) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataDayGrowthSize) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataDayGrowthSize) SetValue(v int64) *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsWarmDataRatio struct {
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
	// 0.5
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsWarmDataRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsWarmDataRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataRatio) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataRatio) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataRatio) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataRatio) SetValue(v float32) *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsWarmDataSize struct {
	// The description of the metric.
	//
	// example:
	//
	// Ratio of warm files
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
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsWarmDataSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsWarmDataSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataSize) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataSize) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataSize) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataSize) SetValue(v int64) *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveClusterResponseBodyDataMetricsWarmDataSizeDayGrowthRatio struct {
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
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 0.5
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) SetDescription(v string) *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) SetName(v string) *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) SetUnit(v string) *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) SetValue(v float32) *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveClusterResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}
