// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorHiveDatabaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDoctorHiveDatabaseResponseBodyData) *GetDoctorHiveDatabaseResponseBody
	GetData() *GetDoctorHiveDatabaseResponseBodyData
	SetRequestId(v string) *GetDoctorHiveDatabaseResponseBody
	GetRequestId() *string
}

type GetDoctorHiveDatabaseResponseBody struct {
	// The analysis results of the Hive database.
	Data *GetDoctorHiveDatabaseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDoctorHiveDatabaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBody) GetData() *GetDoctorHiveDatabaseResponseBodyData {
	return s.Data
}

func (s *GetDoctorHiveDatabaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDoctorHiveDatabaseResponseBody) SetData(v *GetDoctorHiveDatabaseResponseBodyData) *GetDoctorHiveDatabaseResponseBody {
	s.Data = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBody) SetRequestId(v string) *GetDoctorHiveDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyData struct {
	// The analysis results.
	Analysis *GetDoctorHiveDatabaseResponseBodyDataAnalysis `json:"Analysis,omitempty" xml:"Analysis,omitempty" type:"Struct"`
	// The information from the perspective of storage formats.
	Formats []*GetDoctorHiveDatabaseResponseBodyDataFormats `json:"Formats,omitempty" xml:"Formats,omitempty" type:"Repeated"`
	// The metric information.
	Metrics *GetDoctorHiveDatabaseResponseBodyDataMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Struct"`
}

func (s GetDoctorHiveDatabaseResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyData) GetAnalysis() *GetDoctorHiveDatabaseResponseBodyDataAnalysis {
	return s.Analysis
}

func (s *GetDoctorHiveDatabaseResponseBodyData) GetFormats() []*GetDoctorHiveDatabaseResponseBodyDataFormats {
	return s.Formats
}

func (s *GetDoctorHiveDatabaseResponseBodyData) GetMetrics() *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	return s.Metrics
}

func (s *GetDoctorHiveDatabaseResponseBodyData) SetAnalysis(v *GetDoctorHiveDatabaseResponseBodyDataAnalysis) *GetDoctorHiveDatabaseResponseBodyData {
	s.Analysis = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyData) SetFormats(v []*GetDoctorHiveDatabaseResponseBodyDataFormats) *GetDoctorHiveDatabaseResponseBodyData {
	s.Formats = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyData) SetMetrics(v *GetDoctorHiveDatabaseResponseBodyDataMetrics) *GetDoctorHiveDatabaseResponseBodyData {
	s.Metrics = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataAnalysis struct {
	// The score for the file sizes of the Hive database.
	//
	// example:
	//
	// 85
	HiveDistributionScore *int32 `json:"HiveDistributionScore,omitempty" xml:"HiveDistributionScore,omitempty"`
	// The score for the data formats of the Hive database.
	//
	// example:
	//
	// 85
	HiveFormatScore *int32 `json:"HiveFormatScore,omitempty" xml:"HiveFormatScore,omitempty"`
	// The score for the access frequency of the Hive database.
	//
	// example:
	//
	// 85
	HiveFrequencyScore *int32 `json:"HiveFrequencyScore,omitempty" xml:"HiveFrequencyScore,omitempty"`
	// The overall score of the Hive database.
	//
	// example:
	//
	// 85
	HiveScore *int32 `json:"HiveScore,omitempty" xml:"HiveScore,omitempty"`
}

func (s GetDoctorHiveDatabaseResponseBodyDataAnalysis) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataAnalysis) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataAnalysis) GetHiveDistributionScore() *int32 {
	return s.HiveDistributionScore
}

func (s *GetDoctorHiveDatabaseResponseBodyDataAnalysis) GetHiveFormatScore() *int32 {
	return s.HiveFormatScore
}

func (s *GetDoctorHiveDatabaseResponseBodyDataAnalysis) GetHiveFrequencyScore() *int32 {
	return s.HiveFrequencyScore
}

func (s *GetDoctorHiveDatabaseResponseBodyDataAnalysis) GetHiveScore() *int32 {
	return s.HiveScore
}

func (s *GetDoctorHiveDatabaseResponseBodyDataAnalysis) SetHiveDistributionScore(v int32) *GetDoctorHiveDatabaseResponseBodyDataAnalysis {
	s.HiveDistributionScore = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataAnalysis) SetHiveFormatScore(v int32) *GetDoctorHiveDatabaseResponseBodyDataAnalysis {
	s.HiveFormatScore = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataAnalysis) SetHiveFrequencyScore(v int32) *GetDoctorHiveDatabaseResponseBodyDataAnalysis {
	s.HiveFrequencyScore = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataAnalysis) SetHiveScore(v int32) *GetDoctorHiveDatabaseResponseBodyDataAnalysis {
	s.HiveScore = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataAnalysis) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataFormats struct {
	// The daily increment of data in the format.
	//
	// example:
	//
	// 1000
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
	// 0.5
	FormatRatio *float32 `json:"FormatRatio,omitempty" xml:"FormatRatio,omitempty"`
	// The amount of data in the format.
	//
	// example:
	//
	// 1000
	FormatSize *int64 `json:"FormatSize,omitempty" xml:"FormatSize,omitempty"`
	// The day-to-day growth rate of data in the format.
	//
	// example:
	//
	// 0.5
	FormatSizeDayGrowthRatio *float32 `json:"FormatSizeDayGrowthRatio,omitempty" xml:"FormatSizeDayGrowthRatio,omitempty"`
	// The unit of the amount of data in the format.
	//
	// example:
	//
	// MB
	FormatSizeUnit *string `json:"FormatSizeUnit,omitempty" xml:"FormatSizeUnit,omitempty"`
}

func (s GetDoctorHiveDatabaseResponseBodyDataFormats) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataFormats) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataFormats) GetFormatDayGrowthSize() *int64 {
	return s.FormatDayGrowthSize
}

func (s *GetDoctorHiveDatabaseResponseBodyDataFormats) GetFormatName() *string {
	return s.FormatName
}

func (s *GetDoctorHiveDatabaseResponseBodyDataFormats) GetFormatRatio() *float32 {
	return s.FormatRatio
}

func (s *GetDoctorHiveDatabaseResponseBodyDataFormats) GetFormatSize() *int64 {
	return s.FormatSize
}

func (s *GetDoctorHiveDatabaseResponseBodyDataFormats) GetFormatSizeDayGrowthRatio() *float32 {
	return s.FormatSizeDayGrowthRatio
}

func (s *GetDoctorHiveDatabaseResponseBodyDataFormats) GetFormatSizeUnit() *string {
	return s.FormatSizeUnit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataFormats) SetFormatDayGrowthSize(v int64) *GetDoctorHiveDatabaseResponseBodyDataFormats {
	s.FormatDayGrowthSize = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataFormats) SetFormatName(v string) *GetDoctorHiveDatabaseResponseBodyDataFormats {
	s.FormatName = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataFormats) SetFormatRatio(v float32) *GetDoctorHiveDatabaseResponseBodyDataFormats {
	s.FormatRatio = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataFormats) SetFormatSize(v int64) *GetDoctorHiveDatabaseResponseBodyDataFormats {
	s.FormatSize = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataFormats) SetFormatSizeDayGrowthRatio(v float32) *GetDoctorHiveDatabaseResponseBodyDataFormats {
	s.FormatSizeDayGrowthRatio = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataFormats) SetFormatSizeUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataFormats {
	s.FormatSizeUnit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataFormats) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetrics struct {
	// The daily increment of the amount of cold data. Cold data refers to data that is not accessed for more than 30 days but is accessed in previous 90 days.
	ColdDataDayGrowthSize *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataDayGrowthSize `json:"ColdDataDayGrowthSize,omitempty" xml:"ColdDataDayGrowthSize,omitempty" type:"Struct"`
	// The proportion of cold data. Cold data refers to data that is not accessed for more than 30 days but is accessed in previous 90 days.
	ColdDataRatio *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataRatio `json:"ColdDataRatio,omitempty" xml:"ColdDataRatio,omitempty" type:"Struct"`
	// The amount of cold data. Cold data refers to data that is not accessed for more than 30 days but is accessed in previous 90 days.
	ColdDataSize *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataSize `json:"ColdDataSize,omitempty" xml:"ColdDataSize,omitempty" type:"Struct"`
	// The day-to-day growth rate of the amount of cold data. Cold data refers to data that is not accessed for more than 30 days but is accessed in previous 90 days.
	ColdDataSizeDayGrowthRatio *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataSizeDayGrowthRatio `json:"ColdDataSizeDayGrowthRatio,omitempty" xml:"ColdDataSizeDayGrowthRatio,omitempty" type:"Struct"`
	// The number of empty files. Empty files are those with a size of 0 MB.
	EmptyFileCount *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileCount `json:"EmptyFileCount,omitempty" xml:"EmptyFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the number of empty files. Empty files are those with a size of 0 MB.
	EmptyFileCountDayGrowthRatio *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileCountDayGrowthRatio `json:"EmptyFileCountDayGrowthRatio,omitempty" xml:"EmptyFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the number of empty files. Empty files are those with a size of 0 MB.
	EmptyFileDayGrowthCount *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileDayGrowthCount `json:"EmptyFileDayGrowthCount,omitempty" xml:"EmptyFileDayGrowthCount,omitempty" type:"Struct"`
	// The proportion of empty files. Empty files are those with a size of 0 MB.
	EmptyFileRatio *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileRatio `json:"EmptyFileRatio,omitempty" xml:"EmptyFileRatio,omitempty" type:"Struct"`
	// The daily increment of the amount of very cold data. Very cold data refers to data that is not accessed for more than 90 days.
	FreezeDataDayGrowthSize *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataDayGrowthSize `json:"FreezeDataDayGrowthSize,omitempty" xml:"FreezeDataDayGrowthSize,omitempty" type:"Struct"`
	// The proportion of very cold data. Very cold data refers to data that is not accessed for more than 90 days.
	FreezeDataRatio *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataRatio `json:"FreezeDataRatio,omitempty" xml:"FreezeDataRatio,omitempty" type:"Struct"`
	// The amount of very cold data. Very cold data refers to data that is not accessed for more than 90 days.
	FreezeDataSize *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataSize `json:"FreezeDataSize,omitempty" xml:"FreezeDataSize,omitempty" type:"Struct"`
	// The day-to-day growth rate of the amount of very cold data. Very cold data refers to data that is not accessed for more than 90 days.
	FreezeDataSizeDayGrowthRatio *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio `json:"FreezeDataSizeDayGrowthRatio,omitempty" xml:"FreezeDataSizeDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the amount of hot data. Hot data refers to data that is accessed in previous seven days.
	HotDataDayGrowthSize *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataDayGrowthSize `json:"HotDataDayGrowthSize,omitempty" xml:"HotDataDayGrowthSize,omitempty" type:"Struct"`
	// The proportion of hot data. Hot data refers to data that is accessed in previous seven days.
	HotDataRatio *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataRatio `json:"HotDataRatio,omitempty" xml:"HotDataRatio,omitempty" type:"Struct"`
	// The amount of hot data. Hot data refers to data that is accessed in previous seven days.
	HotDataSize *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataSize `json:"HotDataSize,omitempty" xml:"HotDataSize,omitempty" type:"Struct"`
	// The day-to-day growth rate of the amount of hot data. Hot data refers to data that is accessed in previous seven days.
	HotDataSizeDayGrowthRatio *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataSizeDayGrowthRatio `json:"HotDataSizeDayGrowthRatio,omitempty" xml:"HotDataSizeDayGrowthRatio,omitempty" type:"Struct"`
	// The number of large files. Large files are those with a size greater than 1 GB.
	LargeFileCount *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileCount `json:"LargeFileCount,omitempty" xml:"LargeFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the number of large files. Large files are those with a size greater than 1 GB.
	LargeFileCountDayGrowthRatio *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileCountDayGrowthRatio `json:"LargeFileCountDayGrowthRatio,omitempty" xml:"LargeFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the number of large files. Large files are those with a size greater than 1 GB.
	LargeFileDayGrowthCount *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileDayGrowthCount `json:"LargeFileDayGrowthCount,omitempty" xml:"LargeFileDayGrowthCount,omitempty" type:"Struct"`
	// The proportion of large files. Large files are those with a size greater than 1 GB.
	LargeFileRatio *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileRatio `json:"LargeFileRatio,omitempty" xml:"LargeFileRatio,omitempty" type:"Struct"`
	// The number of medium files. Medium files are those with a size greater than or equal to 128 MB and less than or equal to 1 GB.
	MediumFileCount *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileCount `json:"MediumFileCount,omitempty" xml:"MediumFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the number of medium files. Medium files are those with a size greater than or equal to 128 MB and less than or equal to 1 GB.
	MediumFileCountDayGrowthRatio *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileCountDayGrowthRatio `json:"MediumFileCountDayGrowthRatio,omitempty" xml:"MediumFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the number of medium files. Medium files are those with a size greater than or equal to 128 MB and less than or equal to 1 GB.
	MediumFileDayGrowthCount *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileDayGrowthCount `json:"MediumFileDayGrowthCount,omitempty" xml:"MediumFileDayGrowthCount,omitempty" type:"Struct"`
	// The proportion of medium files. Medium files are those with a size greater than or equal to 128 MB and less than or equal to 1 GB.
	MediumFileRatio *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileRatio `json:"MediumFileRatio,omitempty" xml:"MediumFileRatio,omitempty" type:"Struct"`
	// The number of partitions.
	PartitionNum *GetDoctorHiveDatabaseResponseBodyDataMetricsPartitionNum `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty" type:"Struct"`
	// The number of small files. Small files are those with a size greater than or equal to 10 MB and less than 128 MB.
	SmallFileCount *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileCount `json:"SmallFileCount,omitempty" xml:"SmallFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the number of small files. Small files are those with a size greater than or equal to 10 MB and less than 128 MB.
	SmallFileCountDayGrowthRatio *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileCountDayGrowthRatio `json:"SmallFileCountDayGrowthRatio,omitempty" xml:"SmallFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the number of small files. Small files are those with a size greater than or equal to 10 MB and less than 128 MB.
	SmallFileDayGrowthCount *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileDayGrowthCount `json:"SmallFileDayGrowthCount,omitempty" xml:"SmallFileDayGrowthCount,omitempty" type:"Struct"`
	// The proportion of small files. Small files are those with a size greater than or equal to 10 MB and less than 128 MB.
	SmallFileRatio *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileRatio `json:"SmallFileRatio,omitempty" xml:"SmallFileRatio,omitempty" type:"Struct"`
	// The number of tables.
	TableCount *GetDoctorHiveDatabaseResponseBodyDataMetricsTableCount `json:"TableCount,omitempty" xml:"TableCount,omitempty" type:"Struct"`
	// The number of very small files. Very small files are those with a size greater than 0 MB and less than 10 MB.
	TinyFileCount *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileCount `json:"TinyFileCount,omitempty" xml:"TinyFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the number of very small files. Very small files are those with a size greater than 0 MB and less than 10 MB.
	TinyFileCountDayGrowthRatio *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileCountDayGrowthRatio `json:"TinyFileCountDayGrowthRatio,omitempty" xml:"TinyFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the number of very small files. Very small files are those with a size greater than 0 MB and less than 10 MB.
	TinyFileDayGrowthCount *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileDayGrowthCount `json:"TinyFileDayGrowthCount,omitempty" xml:"TinyFileDayGrowthCount,omitempty" type:"Struct"`
	// The proportion of very small files. Very small files are those with a size greater than 0 MB and less than 10 MB.
	TinyFileRatio *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileRatio `json:"TinyFileRatio,omitempty" xml:"TinyFileRatio,omitempty" type:"Struct"`
	// The daily incremental of the total amount of data.
	TotalDataDayGrowthSize *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataDayGrowthSize `json:"TotalDataDayGrowthSize,omitempty" xml:"TotalDataDayGrowthSize,omitempty" type:"Struct"`
	// The total amount of data.
	TotalDataSize *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataSize `json:"TotalDataSize,omitempty" xml:"TotalDataSize,omitempty" type:"Struct"`
	// The day-to-day growth rate of the total data volume.
	TotalDataSizeDayGrowthRatio *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataSizeDayGrowthRatio `json:"TotalDataSizeDayGrowthRatio,omitempty" xml:"TotalDataSizeDayGrowthRatio,omitempty" type:"Struct"`
	// The total number of files.
	TotalFileCount *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileCount `json:"TotalFileCount,omitempty" xml:"TotalFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the total number of files.
	TotalFileCountDayGrowthRatio *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileCountDayGrowthRatio `json:"TotalFileCountDayGrowthRatio,omitempty" xml:"TotalFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the total number of files.
	TotalFileDayGrowthCount *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileDayGrowthCount `json:"TotalFileDayGrowthCount,omitempty" xml:"TotalFileDayGrowthCount,omitempty" type:"Struct"`
	// The daily increment of the amount of warm data. Warm data refers to data that is not accessed for more than 7 days but is accessed in previous 30 days.
	WarmDataDayGrowthSize *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataDayGrowthSize `json:"WarmDataDayGrowthSize,omitempty" xml:"WarmDataDayGrowthSize,omitempty" type:"Struct"`
	// The proportion of warm data. Warm data refers to data that is not accessed for more than 7 days but is accessed in previous 30 days.
	WarmDataRatio *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataRatio `json:"WarmDataRatio,omitempty" xml:"WarmDataRatio,omitempty" type:"Struct"`
	// The amount of warm data. Warm data refers to data that is not accessed for more than 7 days but is accessed in previous 30 days.
	WarmDataSize *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataSize `json:"WarmDataSize,omitempty" xml:"WarmDataSize,omitempty" type:"Struct"`
	// The day-to-day growth rate of the amount of warm data. Warm data refers to data that is not accessed for more than 7 days but is accessed in previous 30 days.
	WarmDataSizeDayGrowthRatio *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataSizeDayGrowthRatio `json:"WarmDataSizeDayGrowthRatio,omitempty" xml:"WarmDataSizeDayGrowthRatio,omitempty" type:"Struct"`
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetrics) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetrics) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetColdDataDayGrowthSize() *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataDayGrowthSize {
	return s.ColdDataDayGrowthSize
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetColdDataRatio() *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataRatio {
	return s.ColdDataRatio
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetColdDataSize() *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataSize {
	return s.ColdDataSize
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetColdDataSizeDayGrowthRatio() *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	return s.ColdDataSizeDayGrowthRatio
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetEmptyFileCount() *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileCount {
	return s.EmptyFileCount
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetEmptyFileCountDayGrowthRatio() *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	return s.EmptyFileCountDayGrowthRatio
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetEmptyFileDayGrowthCount() *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileDayGrowthCount {
	return s.EmptyFileDayGrowthCount
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetEmptyFileRatio() *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileRatio {
	return s.EmptyFileRatio
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetFreezeDataDayGrowthSize() *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataDayGrowthSize {
	return s.FreezeDataDayGrowthSize
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetFreezeDataRatio() *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataRatio {
	return s.FreezeDataRatio
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetFreezeDataSize() *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataSize {
	return s.FreezeDataSize
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetFreezeDataSizeDayGrowthRatio() *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	return s.FreezeDataSizeDayGrowthRatio
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetHotDataDayGrowthSize() *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataDayGrowthSize {
	return s.HotDataDayGrowthSize
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetHotDataRatio() *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataRatio {
	return s.HotDataRatio
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetHotDataSize() *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataSize {
	return s.HotDataSize
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetHotDataSizeDayGrowthRatio() *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	return s.HotDataSizeDayGrowthRatio
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetLargeFileCount() *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileCount {
	return s.LargeFileCount
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetLargeFileCountDayGrowthRatio() *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	return s.LargeFileCountDayGrowthRatio
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetLargeFileDayGrowthCount() *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileDayGrowthCount {
	return s.LargeFileDayGrowthCount
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetLargeFileRatio() *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileRatio {
	return s.LargeFileRatio
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetMediumFileCount() *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileCount {
	return s.MediumFileCount
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetMediumFileCountDayGrowthRatio() *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	return s.MediumFileCountDayGrowthRatio
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetMediumFileDayGrowthCount() *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileDayGrowthCount {
	return s.MediumFileDayGrowthCount
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetMediumFileRatio() *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileRatio {
	return s.MediumFileRatio
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetPartitionNum() *GetDoctorHiveDatabaseResponseBodyDataMetricsPartitionNum {
	return s.PartitionNum
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetSmallFileCount() *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileCount {
	return s.SmallFileCount
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetSmallFileCountDayGrowthRatio() *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	return s.SmallFileCountDayGrowthRatio
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetSmallFileDayGrowthCount() *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileDayGrowthCount {
	return s.SmallFileDayGrowthCount
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetSmallFileRatio() *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileRatio {
	return s.SmallFileRatio
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetTableCount() *GetDoctorHiveDatabaseResponseBodyDataMetricsTableCount {
	return s.TableCount
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetTinyFileCount() *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileCount {
	return s.TinyFileCount
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetTinyFileCountDayGrowthRatio() *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	return s.TinyFileCountDayGrowthRatio
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetTinyFileDayGrowthCount() *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileDayGrowthCount {
	return s.TinyFileDayGrowthCount
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetTinyFileRatio() *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileRatio {
	return s.TinyFileRatio
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetTotalDataDayGrowthSize() *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataDayGrowthSize {
	return s.TotalDataDayGrowthSize
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetTotalDataSize() *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataSize {
	return s.TotalDataSize
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetTotalDataSizeDayGrowthRatio() *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	return s.TotalDataSizeDayGrowthRatio
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetTotalFileCount() *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileCount {
	return s.TotalFileCount
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetTotalFileCountDayGrowthRatio() *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	return s.TotalFileCountDayGrowthRatio
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetTotalFileDayGrowthCount() *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileDayGrowthCount {
	return s.TotalFileDayGrowthCount
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetWarmDataDayGrowthSize() *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataDayGrowthSize {
	return s.WarmDataDayGrowthSize
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetWarmDataRatio() *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataRatio {
	return s.WarmDataRatio
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetWarmDataSize() *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataSize {
	return s.WarmDataSize
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) GetWarmDataSizeDayGrowthRatio() *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	return s.WarmDataSizeDayGrowthRatio
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetColdDataDayGrowthSize(v *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataDayGrowthSize) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.ColdDataDayGrowthSize = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetColdDataRatio(v *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataRatio) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.ColdDataRatio = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetColdDataSize(v *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataSize) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.ColdDataSize = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetColdDataSizeDayGrowthRatio(v *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataSizeDayGrowthRatio) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.ColdDataSizeDayGrowthRatio = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetEmptyFileCount(v *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileCount) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.EmptyFileCount = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetEmptyFileCountDayGrowthRatio(v *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.EmptyFileCountDayGrowthRatio = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetEmptyFileDayGrowthCount(v *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileDayGrowthCount) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.EmptyFileDayGrowthCount = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetEmptyFileRatio(v *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileRatio) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.EmptyFileRatio = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetFreezeDataDayGrowthSize(v *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataDayGrowthSize) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.FreezeDataDayGrowthSize = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetFreezeDataRatio(v *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataRatio) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.FreezeDataRatio = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetFreezeDataSize(v *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataSize) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.FreezeDataSize = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetFreezeDataSizeDayGrowthRatio(v *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.FreezeDataSizeDayGrowthRatio = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetHotDataDayGrowthSize(v *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataDayGrowthSize) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.HotDataDayGrowthSize = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetHotDataRatio(v *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataRatio) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.HotDataRatio = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetHotDataSize(v *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataSize) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.HotDataSize = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetHotDataSizeDayGrowthRatio(v *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataSizeDayGrowthRatio) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.HotDataSizeDayGrowthRatio = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetLargeFileCount(v *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileCount) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.LargeFileCount = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetLargeFileCountDayGrowthRatio(v *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileCountDayGrowthRatio) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.LargeFileCountDayGrowthRatio = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetLargeFileDayGrowthCount(v *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileDayGrowthCount) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.LargeFileDayGrowthCount = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetLargeFileRatio(v *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileRatio) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.LargeFileRatio = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetMediumFileCount(v *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileCount) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.MediumFileCount = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetMediumFileCountDayGrowthRatio(v *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileCountDayGrowthRatio) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.MediumFileCountDayGrowthRatio = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetMediumFileDayGrowthCount(v *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileDayGrowthCount) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.MediumFileDayGrowthCount = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetMediumFileRatio(v *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileRatio) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.MediumFileRatio = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetPartitionNum(v *GetDoctorHiveDatabaseResponseBodyDataMetricsPartitionNum) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.PartitionNum = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetSmallFileCount(v *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileCount) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.SmallFileCount = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetSmallFileCountDayGrowthRatio(v *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileCountDayGrowthRatio) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.SmallFileCountDayGrowthRatio = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetSmallFileDayGrowthCount(v *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileDayGrowthCount) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.SmallFileDayGrowthCount = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetSmallFileRatio(v *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileRatio) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.SmallFileRatio = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetTableCount(v *GetDoctorHiveDatabaseResponseBodyDataMetricsTableCount) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.TableCount = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetTinyFileCount(v *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileCount) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.TinyFileCount = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetTinyFileCountDayGrowthRatio(v *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileCountDayGrowthRatio) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.TinyFileCountDayGrowthRatio = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetTinyFileDayGrowthCount(v *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileDayGrowthCount) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.TinyFileDayGrowthCount = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetTinyFileRatio(v *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileRatio) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.TinyFileRatio = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetTotalDataDayGrowthSize(v *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataDayGrowthSize) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.TotalDataDayGrowthSize = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetTotalDataSize(v *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataSize) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.TotalDataSize = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetTotalDataSizeDayGrowthRatio(v *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.TotalDataSizeDayGrowthRatio = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetTotalFileCount(v *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileCount) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.TotalFileCount = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetTotalFileCountDayGrowthRatio(v *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileCountDayGrowthRatio) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.TotalFileCountDayGrowthRatio = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetTotalFileDayGrowthCount(v *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileDayGrowthCount) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.TotalFileDayGrowthCount = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetWarmDataDayGrowthSize(v *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataDayGrowthSize) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.WarmDataDayGrowthSize = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetWarmDataRatio(v *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataRatio) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.WarmDataRatio = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetWarmDataSize(v *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataSize) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.WarmDataSize = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) SetWarmDataSizeDayGrowthRatio(v *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) *GetDoctorHiveDatabaseResponseBodyDataMetrics {
	s.WarmDataSizeDayGrowthRatio = v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetrics) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataDayGrowthSize struct {
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
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataDayGrowthSize) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataDayGrowthSize) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataDayGrowthSize) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataDayGrowthSize) SetValue(v int64) *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataRatio struct {
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

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataRatio) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataRatio) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataRatio) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataRatio) SetValue(v float32) *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataSize struct {
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

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataSize) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataSize) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataSize) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataSize) SetValue(v int64) *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataSizeDayGrowthRatio struct {
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

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataSizeDayGrowthRatio) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataSizeDayGrowthRatio) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataSizeDayGrowthRatio) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataSizeDayGrowthRatio) SetValue(v float32) *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsColdDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileCount struct {
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

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileCount) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileCount) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileCount) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileCount) SetValue(v int64) *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileCountDayGrowthRatio struct {
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

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) SetValue(v float32) *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileDayGrowthCount struct {
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

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileDayGrowthCount) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileDayGrowthCount) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileDayGrowthCount) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileDayGrowthCount) SetValue(v int64) *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileRatio struct {
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

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileRatio) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileRatio) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileRatio) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileRatio) SetValue(v float32) *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsEmptyFileRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataDayGrowthSize struct {
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

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataDayGrowthSize) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataDayGrowthSize) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataDayGrowthSize) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataDayGrowthSize) SetValue(v int64) *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataRatio struct {
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
	// 0.5
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataRatio) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataRatio) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataRatio) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataRatio) SetValue(v float32) *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataSize struct {
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

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataSize) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataSize) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataSize) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataSize) SetValue(v int64) *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio struct {
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

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) SetValue(v float32) *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataDayGrowthSize struct {
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
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataDayGrowthSize) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataDayGrowthSize) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataDayGrowthSize) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataDayGrowthSize) SetValue(v int64) *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataRatio struct {
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

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataRatio) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataRatio) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataRatio) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataRatio) SetValue(v float32) *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataSize struct {
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

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataSize) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataSize) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataSize) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataSize) SetValue(v int64) *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataSizeDayGrowthRatio struct {
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

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataSizeDayGrowthRatio) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataSizeDayGrowthRatio) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataSizeDayGrowthRatio) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataSizeDayGrowthRatio) SetValue(v float32) *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsHotDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileCount struct {
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

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileCount) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileCount) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileCount) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileCount) SetValue(v int64) *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileCountDayGrowthRatio struct {
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

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileCountDayGrowthRatio) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileCountDayGrowthRatio) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileCountDayGrowthRatio) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileCountDayGrowthRatio) SetValue(v float32) *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileDayGrowthCount struct {
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

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileDayGrowthCount) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileDayGrowthCount) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileDayGrowthCount) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileDayGrowthCount) SetValue(v int64) *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileRatio struct {
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

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileRatio) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileRatio) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileRatio) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileRatio) SetValue(v float32) *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsLargeFileRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileCount struct {
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

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileCount) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileCount) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileCount) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileCount) SetValue(v int64) *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileCountDayGrowthRatio struct {
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

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileCountDayGrowthRatio) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileCountDayGrowthRatio) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileCountDayGrowthRatio) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileCountDayGrowthRatio) SetValue(v float32) *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileDayGrowthCount struct {
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
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileDayGrowthCount) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileDayGrowthCount) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileDayGrowthCount) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileDayGrowthCount) SetValue(v int64) *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileRatio struct {
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

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileRatio) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileRatio) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileRatio) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileRatio) SetValue(v float32) *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsMediumFileRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsPartitionNum struct {
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

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsPartitionNum) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsPartitionNum) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsPartitionNum) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsPartitionNum) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsPartitionNum) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsPartitionNum) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsPartitionNum) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsPartitionNum {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsPartitionNum) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsPartitionNum {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsPartitionNum) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsPartitionNum {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsPartitionNum) SetValue(v int64) *GetDoctorHiveDatabaseResponseBodyDataMetricsPartitionNum {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsPartitionNum) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileCount struct {
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

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileCount) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileCount) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileCount) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileCount) SetValue(v int64) *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileCountDayGrowthRatio struct {
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

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileCountDayGrowthRatio) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileCountDayGrowthRatio) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileCountDayGrowthRatio) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileCountDayGrowthRatio) SetValue(v float32) *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileDayGrowthCount struct {
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
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileDayGrowthCount) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileDayGrowthCount) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileDayGrowthCount) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileDayGrowthCount) SetValue(v int64) *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileRatio struct {
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

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileRatio) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileRatio) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileRatio) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileRatio) SetValue(v float32) *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsSmallFileRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsTableCount struct {
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

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsTableCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsTableCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTableCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTableCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTableCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTableCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTableCount) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsTableCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTableCount) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsTableCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTableCount) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsTableCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTableCount) SetValue(v int64) *GetDoctorHiveDatabaseResponseBodyDataMetricsTableCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTableCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileCount struct {
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

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileCount) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileCount) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileCount) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileCount) SetValue(v int64) *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileCountDayGrowthRatio struct {
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

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileCountDayGrowthRatio) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileCountDayGrowthRatio) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileCountDayGrowthRatio) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileCountDayGrowthRatio) SetValue(v float32) *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileDayGrowthCount struct {
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
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileDayGrowthCount) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileDayGrowthCount) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileDayGrowthCount) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileDayGrowthCount) SetValue(v int64) *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileRatio struct {
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

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileRatio) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileRatio) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileRatio) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileRatio) SetValue(v float32) *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTinyFileRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataDayGrowthSize struct {
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
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataDayGrowthSize) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataDayGrowthSize) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataDayGrowthSize) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataDayGrowthSize) SetValue(v int64) *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataSize struct {
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

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataSize) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataSize) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataSize) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataSize) SetValue(v int64) *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataSizeDayGrowthRatio struct {
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

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) SetValue(v float32) *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileCount struct {
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
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileCount) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileCount) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileCount) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileCount) SetValue(v int64) *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileCountDayGrowthRatio struct {
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

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileCountDayGrowthRatio) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileCountDayGrowthRatio) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileCountDayGrowthRatio) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileCountDayGrowthRatio) SetValue(v float32) *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileDayGrowthCount struct {
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
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileDayGrowthCount) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileDayGrowthCount) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileDayGrowthCount) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileDayGrowthCount) SetValue(v int64) *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsTotalFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataDayGrowthSize struct {
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
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataDayGrowthSize) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataDayGrowthSize) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataDayGrowthSize) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataDayGrowthSize) SetValue(v int64) *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataRatio struct {
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

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataRatio) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataRatio) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataRatio) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataRatio) SetValue(v float32) *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataSize struct {
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

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataSize) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataSize) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataSize) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataSize) SetValue(v int64) *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataSizeDayGrowthRatio struct {
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

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) SetDescription(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) SetName(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) SetUnit(v string) *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) SetValue(v float32) *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}
