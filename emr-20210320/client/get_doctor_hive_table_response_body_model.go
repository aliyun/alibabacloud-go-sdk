// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorHiveTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDoctorHiveTableResponseBodyData) *GetDoctorHiveTableResponseBody
	GetData() *GetDoctorHiveTableResponseBodyData
	SetRequestId(v string) *GetDoctorHiveTableResponseBody
	GetRequestId() *string
}

type GetDoctorHiveTableResponseBody struct {
	// The analysis results of the Hive table.
	Data *GetDoctorHiveTableResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDoctorHiveTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBody) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBody) GetData() *GetDoctorHiveTableResponseBodyData {
	return s.Data
}

func (s *GetDoctorHiveTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDoctorHiveTableResponseBody) SetData(v *GetDoctorHiveTableResponseBodyData) *GetDoctorHiveTableResponseBody {
	s.Data = v
	return s
}

func (s *GetDoctorHiveTableResponseBody) SetRequestId(v string) *GetDoctorHiveTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDoctorHiveTableResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyData struct {
	// The analysis results.
	Analysis *GetDoctorHiveTableResponseBodyDataAnalysis `json:"Analysis,omitempty" xml:"Analysis,omitempty" type:"Struct"`
	// The information from the perspective of formats.
	Formats []*GetDoctorHiveTableResponseBodyDataFormats `json:"Formats,omitempty" xml:"Formats,omitempty" type:"Repeated"`
	// The metric information.
	Metrics *GetDoctorHiveTableResponseBodyDataMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Struct"`
	// The owner.
	//
	// example:
	//
	// DW
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
}

func (s GetDoctorHiveTableResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyData) GetAnalysis() *GetDoctorHiveTableResponseBodyDataAnalysis {
	return s.Analysis
}

func (s *GetDoctorHiveTableResponseBodyData) GetFormats() []*GetDoctorHiveTableResponseBodyDataFormats {
	return s.Formats
}

func (s *GetDoctorHiveTableResponseBodyData) GetMetrics() *GetDoctorHiveTableResponseBodyDataMetrics {
	return s.Metrics
}

func (s *GetDoctorHiveTableResponseBodyData) GetOwner() *string {
	return s.Owner
}

func (s *GetDoctorHiveTableResponseBodyData) SetAnalysis(v *GetDoctorHiveTableResponseBodyDataAnalysis) *GetDoctorHiveTableResponseBodyData {
	s.Analysis = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyData) SetFormats(v []*GetDoctorHiveTableResponseBodyDataFormats) *GetDoctorHiveTableResponseBodyData {
	s.Formats = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyData) SetMetrics(v *GetDoctorHiveTableResponseBodyDataMetrics) *GetDoctorHiveTableResponseBodyData {
	s.Metrics = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyData) SetOwner(v string) *GetDoctorHiveTableResponseBodyData {
	s.Owner = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataAnalysis struct {
	// The score for the distribution of files of different sizes stored in the Hive table.
	//
	// example:
	//
	// 80
	HiveDistributionScore *int32 `json:"HiveDistributionScore,omitempty" xml:"HiveDistributionScore,omitempty"`
	// The score for the distribution of files stored in different formats in the Hive table.
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

func (s GetDoctorHiveTableResponseBodyDataAnalysis) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataAnalysis) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataAnalysis) GetHiveDistributionScore() *int32 {
	return s.HiveDistributionScore
}

func (s *GetDoctorHiveTableResponseBodyDataAnalysis) GetHiveFormatScore() *int32 {
	return s.HiveFormatScore
}

func (s *GetDoctorHiveTableResponseBodyDataAnalysis) GetHiveFrequencyScore() *int32 {
	return s.HiveFrequencyScore
}

func (s *GetDoctorHiveTableResponseBodyDataAnalysis) GetHiveScore() *int32 {
	return s.HiveScore
}

func (s *GetDoctorHiveTableResponseBodyDataAnalysis) SetHiveDistributionScore(v int32) *GetDoctorHiveTableResponseBodyDataAnalysis {
	s.HiveDistributionScore = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataAnalysis) SetHiveFormatScore(v int32) *GetDoctorHiveTableResponseBodyDataAnalysis {
	s.HiveFormatScore = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataAnalysis) SetHiveFrequencyScore(v int32) *GetDoctorHiveTableResponseBodyDataAnalysis {
	s.HiveFrequencyScore = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataAnalysis) SetHiveScore(v int32) *GetDoctorHiveTableResponseBodyDataAnalysis {
	s.HiveScore = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataAnalysis) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataFormats struct {
	// The daily amount increment of the data in a specific storage format.
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
	// The ratio of the data in a specific storage format.
	//
	// example:
	//
	// 0.23
	FormatRatio *float32 `json:"FormatRatio,omitempty" xml:"FormatRatio,omitempty"`
	// The size of storage format-specific data.
	//
	// example:
	//
	// 506930200
	FormatSize *int64 `json:"FormatSize,omitempty" xml:"FormatSize,omitempty"`
	// The day-to-day growth rate of the amount of the data in a specific storage format.
	//
	// example:
	//
	// 0.04
	FormatSizeDayGrowthRatio *float32 `json:"FormatSizeDayGrowthRatio,omitempty" xml:"FormatSizeDayGrowthRatio,omitempty"`
	// The unit of the data size.
	//
	// example:
	//
	// MB
	FormatSizeUnit *string `json:"FormatSizeUnit,omitempty" xml:"FormatSizeUnit,omitempty"`
}

func (s GetDoctorHiveTableResponseBodyDataFormats) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataFormats) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataFormats) GetFormatDayGrowthSize() *int64 {
	return s.FormatDayGrowthSize
}

func (s *GetDoctorHiveTableResponseBodyDataFormats) GetFormatName() *string {
	return s.FormatName
}

func (s *GetDoctorHiveTableResponseBodyDataFormats) GetFormatRatio() *float32 {
	return s.FormatRatio
}

func (s *GetDoctorHiveTableResponseBodyDataFormats) GetFormatSize() *int64 {
	return s.FormatSize
}

func (s *GetDoctorHiveTableResponseBodyDataFormats) GetFormatSizeDayGrowthRatio() *float32 {
	return s.FormatSizeDayGrowthRatio
}

func (s *GetDoctorHiveTableResponseBodyDataFormats) GetFormatSizeUnit() *string {
	return s.FormatSizeUnit
}

func (s *GetDoctorHiveTableResponseBodyDataFormats) SetFormatDayGrowthSize(v int64) *GetDoctorHiveTableResponseBodyDataFormats {
	s.FormatDayGrowthSize = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataFormats) SetFormatName(v string) *GetDoctorHiveTableResponseBodyDataFormats {
	s.FormatName = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataFormats) SetFormatRatio(v float32) *GetDoctorHiveTableResponseBodyDataFormats {
	s.FormatRatio = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataFormats) SetFormatSize(v int64) *GetDoctorHiveTableResponseBodyDataFormats {
	s.FormatSize = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataFormats) SetFormatSizeDayGrowthRatio(v float32) *GetDoctorHiveTableResponseBodyDataFormats {
	s.FormatSizeDayGrowthRatio = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataFormats) SetFormatSizeUnit(v string) *GetDoctorHiveTableResponseBodyDataFormats {
	s.FormatSizeUnit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataFormats) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetrics struct {
	// The daily increment of the amount of cold data. Cold data refers to data that is not accessed for more than 30 days but is accessed in 90 days.
	ColdDataDayGrowthSize *GetDoctorHiveTableResponseBodyDataMetricsColdDataDayGrowthSize `json:"ColdDataDayGrowthSize,omitempty" xml:"ColdDataDayGrowthSize,omitempty" type:"Struct"`
	// The proportion of cold data. Cold data refers to data that is not accessed for more than 30 days but is accessed in 90 days.
	ColdDataRatio *GetDoctorHiveTableResponseBodyDataMetricsColdDataRatio `json:"ColdDataRatio,omitempty" xml:"ColdDataRatio,omitempty" type:"Struct"`
	// The amount of cold data. Cold data refers to data that is not accessed for more than 30 days but is accessed in 90 days.
	ColdDataSize *GetDoctorHiveTableResponseBodyDataMetricsColdDataSize `json:"ColdDataSize,omitempty" xml:"ColdDataSize,omitempty" type:"Struct"`
	// The day-to-day growth rate of the amount of cold data. Cold data refers to data that is not accessed for more than 30 days but is accessed in 90 days.
	ColdDataSizeDayGrowthRatio *GetDoctorHiveTableResponseBodyDataMetricsColdDataSizeDayGrowthRatio `json:"ColdDataSizeDayGrowthRatio,omitempty" xml:"ColdDataSizeDayGrowthRatio,omitempty" type:"Struct"`
	// The number of empty files. Empty files are those with a size of 0 MB.
	EmptyFileCount *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileCount `json:"EmptyFileCount,omitempty" xml:"EmptyFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the number of empty files. Empty files are those with a size of 0 MB.
	EmptyFileCountDayGrowthRatio *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileCountDayGrowthRatio `json:"EmptyFileCountDayGrowthRatio,omitempty" xml:"EmptyFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the number of empty files. Empty files are those with a size of 0 MB.
	EmptyFileDayGrowthCount *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileDayGrowthCount `json:"EmptyFileDayGrowthCount,omitempty" xml:"EmptyFileDayGrowthCount,omitempty" type:"Struct"`
	// The proportion of empty files. Empty files are those with a size of 0 MB.
	EmptyFileRatio *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileRatio `json:"EmptyFileRatio,omitempty" xml:"EmptyFileRatio,omitempty" type:"Struct"`
	// The daily increment of the amount of very cold data. Very cold data refers to data that is not accessed for more than 90 days.
	FreezeDataDayGrowthSize *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataDayGrowthSize `json:"FreezeDataDayGrowthSize,omitempty" xml:"FreezeDataDayGrowthSize,omitempty" type:"Struct"`
	// The proportion of very cold data. Very cold data refers to data that is not accessed for more than 90 days.
	FreezeDataRatio *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataRatio `json:"FreezeDataRatio,omitempty" xml:"FreezeDataRatio,omitempty" type:"Struct"`
	// The amount of very cold data. Very cold data refers to data that is not accessed for more than 90 days.
	FreezeDataSize *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataSize `json:"FreezeDataSize,omitempty" xml:"FreezeDataSize,omitempty" type:"Struct"`
	// The day-to-day growth rate of the amount of very cold data. Very cold data refers to data that is not accessed for more than 90 days.
	FreezeDataSizeDayGrowthRatio *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio `json:"FreezeDataSizeDayGrowthRatio,omitempty" xml:"FreezeDataSizeDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the amount of hot data. Hot data refers to data that is accessed in recent seven days.
	HotDataDayGrowthSize *GetDoctorHiveTableResponseBodyDataMetricsHotDataDayGrowthSize `json:"HotDataDayGrowthSize,omitempty" xml:"HotDataDayGrowthSize,omitempty" type:"Struct"`
	// The proportion of hot data. Hot data refers to data that is accessed in recent seven days.
	HotDataRatio *GetDoctorHiveTableResponseBodyDataMetricsHotDataRatio `json:"HotDataRatio,omitempty" xml:"HotDataRatio,omitempty" type:"Struct"`
	// The amount of hot data. Hot data refers to data that is accessed in recent seven days.
	HotDataSize *GetDoctorHiveTableResponseBodyDataMetricsHotDataSize `json:"HotDataSize,omitempty" xml:"HotDataSize,omitempty" type:"Struct"`
	// The day-to-day growth rate of the amount of hot data. Hot data refers to data that is accessed in recent seven days.
	HotDataSizeDayGrowthRatio *GetDoctorHiveTableResponseBodyDataMetricsHotDataSizeDayGrowthRatio `json:"HotDataSizeDayGrowthRatio,omitempty" xml:"HotDataSizeDayGrowthRatio,omitempty" type:"Struct"`
	// The number of large files. Large files are those with a size greater than 1 GB.
	LargeFileCount *GetDoctorHiveTableResponseBodyDataMetricsLargeFileCount `json:"LargeFileCount,omitempty" xml:"LargeFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the number of large files. Large files are those with a size greater than 1 GB.
	LargeFileCountDayGrowthRatio *GetDoctorHiveTableResponseBodyDataMetricsLargeFileCountDayGrowthRatio `json:"LargeFileCountDayGrowthRatio,omitempty" xml:"LargeFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the number of large files. Large files are those with a size greater than 1 GB.
	LargeFileDayGrowthCount *GetDoctorHiveTableResponseBodyDataMetricsLargeFileDayGrowthCount `json:"LargeFileDayGrowthCount,omitempty" xml:"LargeFileDayGrowthCount,omitempty" type:"Struct"`
	// The proportion of large files. Large files are those with a size greater than 1 GB.
	LargeFileRatio *GetDoctorHiveTableResponseBodyDataMetricsLargeFileRatio `json:"LargeFileRatio,omitempty" xml:"LargeFileRatio,omitempty" type:"Struct"`
	// The number of medium files. Medium files are those with a size greater than or equal to 128 MB and less than or equal to 1 GB.
	MediumFileCount *GetDoctorHiveTableResponseBodyDataMetricsMediumFileCount `json:"MediumFileCount,omitempty" xml:"MediumFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the number of medium files. Medium files are those with a size greater than or equal to 128 MB and less than or equal to 1 GB.
	MediumFileCountDayGrowthRatio *GetDoctorHiveTableResponseBodyDataMetricsMediumFileCountDayGrowthRatio `json:"MediumFileCountDayGrowthRatio,omitempty" xml:"MediumFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the number of medium files. Medium files are those with a size greater than or equal to 128 MB and less than or equal to 1 GB.
	MediumFileDayGrowthCount *GetDoctorHiveTableResponseBodyDataMetricsMediumFileDayGrowthCount `json:"MediumFileDayGrowthCount,omitempty" xml:"MediumFileDayGrowthCount,omitempty" type:"Struct"`
	// The proportion of medium files. Medium files are those with a size greater than or equal to 128 MB and less than or equal to 1 GB.
	MediumFileRatio *GetDoctorHiveTableResponseBodyDataMetricsMediumFileRatio `json:"MediumFileRatio,omitempty" xml:"MediumFileRatio,omitempty" type:"Struct"`
	// The number of partitions.
	PartitionNum *GetDoctorHiveTableResponseBodyDataMetricsPartitionNum `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty" type:"Struct"`
	// The number of small files. Small files are those with a size greater than or equal to 10 MB and less than 128 MB.
	SmallFileCount *GetDoctorHiveTableResponseBodyDataMetricsSmallFileCount `json:"SmallFileCount,omitempty" xml:"SmallFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the number of small files. Small files are those with a size greater than or equal to 10 MB and less than 128 MB.
	SmallFileCountDayGrowthRatio *GetDoctorHiveTableResponseBodyDataMetricsSmallFileCountDayGrowthRatio `json:"SmallFileCountDayGrowthRatio,omitempty" xml:"SmallFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the number of small files. Small files are those with a size greater than or equal to 10 MB and less than 128 MB.
	SmallFileDayGrowthCount *GetDoctorHiveTableResponseBodyDataMetricsSmallFileDayGrowthCount `json:"SmallFileDayGrowthCount,omitempty" xml:"SmallFileDayGrowthCount,omitempty" type:"Struct"`
	// The proportion of small files. Small files are those with a size greater than or equal to 10 MB and less than 128 MB.
	SmallFileRatio *GetDoctorHiveTableResponseBodyDataMetricsSmallFileRatio `json:"SmallFileRatio,omitempty" xml:"SmallFileRatio,omitempty" type:"Struct"`
	// The number of very small files. Very small files are those with a size greater than 0 MB and less than 10 MB.
	TinyFileCount *GetDoctorHiveTableResponseBodyDataMetricsTinyFileCount `json:"TinyFileCount,omitempty" xml:"TinyFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the number of very small files. Very small files are those with a size greater than 0 MB and less than 10 MB.
	TinyFileCountDayGrowthRatio *GetDoctorHiveTableResponseBodyDataMetricsTinyFileCountDayGrowthRatio `json:"TinyFileCountDayGrowthRatio,omitempty" xml:"TinyFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the number of very small files. Very small files are those with a size greater than 0 MB and less than 10 MB.
	TinyFileDayGrowthCount *GetDoctorHiveTableResponseBodyDataMetricsTinyFileDayGrowthCount `json:"TinyFileDayGrowthCount,omitempty" xml:"TinyFileDayGrowthCount,omitempty" type:"Struct"`
	// The proportion of very small files. Very small files are those with a size greater than 0 MB and less than 10 MB.
	TinyFileRatio *GetDoctorHiveTableResponseBodyDataMetricsTinyFileRatio `json:"TinyFileRatio,omitempty" xml:"TinyFileRatio,omitempty" type:"Struct"`
	// The daily incremental of the total data volume.
	TotalDataDayGrowthSize *GetDoctorHiveTableResponseBodyDataMetricsTotalDataDayGrowthSize `json:"TotalDataDayGrowthSize,omitempty" xml:"TotalDataDayGrowthSize,omitempty" type:"Struct"`
	// The total amount of data.
	TotalDataSize *GetDoctorHiveTableResponseBodyDataMetricsTotalDataSize `json:"TotalDataSize,omitempty" xml:"TotalDataSize,omitempty" type:"Struct"`
	// The day-to-day growth rate of the total data volume.
	TotalDataSizeDayGrowthRatio *GetDoctorHiveTableResponseBodyDataMetricsTotalDataSizeDayGrowthRatio `json:"TotalDataSizeDayGrowthRatio,omitempty" xml:"TotalDataSizeDayGrowthRatio,omitempty" type:"Struct"`
	// The total number of files.
	TotalFileCount *GetDoctorHiveTableResponseBodyDataMetricsTotalFileCount `json:"TotalFileCount,omitempty" xml:"TotalFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the total number of files.
	TotalFileCountDayGrowthRatio *GetDoctorHiveTableResponseBodyDataMetricsTotalFileCountDayGrowthRatio `json:"TotalFileCountDayGrowthRatio,omitempty" xml:"TotalFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the total number of files.
	TotalFileDayGrowthCount *GetDoctorHiveTableResponseBodyDataMetricsTotalFileDayGrowthCount `json:"TotalFileDayGrowthCount,omitempty" xml:"TotalFileDayGrowthCount,omitempty" type:"Struct"`
	// The daily increment of the amount of warm data. Warm data refers to data that is not accessed for more than 7 days but is accessed in 30 days.
	WarmDataDayGrowthSize *GetDoctorHiveTableResponseBodyDataMetricsWarmDataDayGrowthSize `json:"WarmDataDayGrowthSize,omitempty" xml:"WarmDataDayGrowthSize,omitempty" type:"Struct"`
	// The proportion of warm data. Warm data refers to data that is not accessed for more than 7 days but is accessed in 30 days.
	WarmDataRatio *GetDoctorHiveTableResponseBodyDataMetricsWarmDataRatio `json:"WarmDataRatio,omitempty" xml:"WarmDataRatio,omitempty" type:"Struct"`
	// The amount of warm data. Warm data refers to data that is not accessed for more than 7 days but is accessed in 30 days.
	WarmDataSize *GetDoctorHiveTableResponseBodyDataMetricsWarmDataSize `json:"WarmDataSize,omitempty" xml:"WarmDataSize,omitempty" type:"Struct"`
	// The day-to-day growth rate of the amount of warm data. Warm data refers to data that is not accessed for more than 7 days but is accessed in 30 days.
	WarmDataSizeDayGrowthRatio *GetDoctorHiveTableResponseBodyDataMetricsWarmDataSizeDayGrowthRatio `json:"WarmDataSizeDayGrowthRatio,omitempty" xml:"WarmDataSizeDayGrowthRatio,omitempty" type:"Struct"`
}

func (s GetDoctorHiveTableResponseBodyDataMetrics) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetrics) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetColdDataDayGrowthSize() *GetDoctorHiveTableResponseBodyDataMetricsColdDataDayGrowthSize {
	return s.ColdDataDayGrowthSize
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetColdDataRatio() *GetDoctorHiveTableResponseBodyDataMetricsColdDataRatio {
	return s.ColdDataRatio
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetColdDataSize() *GetDoctorHiveTableResponseBodyDataMetricsColdDataSize {
	return s.ColdDataSize
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetColdDataSizeDayGrowthRatio() *GetDoctorHiveTableResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	return s.ColdDataSizeDayGrowthRatio
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetEmptyFileCount() *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileCount {
	return s.EmptyFileCount
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetEmptyFileCountDayGrowthRatio() *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	return s.EmptyFileCountDayGrowthRatio
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetEmptyFileDayGrowthCount() *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileDayGrowthCount {
	return s.EmptyFileDayGrowthCount
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetEmptyFileRatio() *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileRatio {
	return s.EmptyFileRatio
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetFreezeDataDayGrowthSize() *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataDayGrowthSize {
	return s.FreezeDataDayGrowthSize
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetFreezeDataRatio() *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataRatio {
	return s.FreezeDataRatio
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetFreezeDataSize() *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataSize {
	return s.FreezeDataSize
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetFreezeDataSizeDayGrowthRatio() *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	return s.FreezeDataSizeDayGrowthRatio
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetHotDataDayGrowthSize() *GetDoctorHiveTableResponseBodyDataMetricsHotDataDayGrowthSize {
	return s.HotDataDayGrowthSize
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetHotDataRatio() *GetDoctorHiveTableResponseBodyDataMetricsHotDataRatio {
	return s.HotDataRatio
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetHotDataSize() *GetDoctorHiveTableResponseBodyDataMetricsHotDataSize {
	return s.HotDataSize
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetHotDataSizeDayGrowthRatio() *GetDoctorHiveTableResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	return s.HotDataSizeDayGrowthRatio
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetLargeFileCount() *GetDoctorHiveTableResponseBodyDataMetricsLargeFileCount {
	return s.LargeFileCount
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetLargeFileCountDayGrowthRatio() *GetDoctorHiveTableResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	return s.LargeFileCountDayGrowthRatio
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetLargeFileDayGrowthCount() *GetDoctorHiveTableResponseBodyDataMetricsLargeFileDayGrowthCount {
	return s.LargeFileDayGrowthCount
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetLargeFileRatio() *GetDoctorHiveTableResponseBodyDataMetricsLargeFileRatio {
	return s.LargeFileRatio
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetMediumFileCount() *GetDoctorHiveTableResponseBodyDataMetricsMediumFileCount {
	return s.MediumFileCount
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetMediumFileCountDayGrowthRatio() *GetDoctorHiveTableResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	return s.MediumFileCountDayGrowthRatio
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetMediumFileDayGrowthCount() *GetDoctorHiveTableResponseBodyDataMetricsMediumFileDayGrowthCount {
	return s.MediumFileDayGrowthCount
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetMediumFileRatio() *GetDoctorHiveTableResponseBodyDataMetricsMediumFileRatio {
	return s.MediumFileRatio
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetPartitionNum() *GetDoctorHiveTableResponseBodyDataMetricsPartitionNum {
	return s.PartitionNum
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetSmallFileCount() *GetDoctorHiveTableResponseBodyDataMetricsSmallFileCount {
	return s.SmallFileCount
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetSmallFileCountDayGrowthRatio() *GetDoctorHiveTableResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	return s.SmallFileCountDayGrowthRatio
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetSmallFileDayGrowthCount() *GetDoctorHiveTableResponseBodyDataMetricsSmallFileDayGrowthCount {
	return s.SmallFileDayGrowthCount
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetSmallFileRatio() *GetDoctorHiveTableResponseBodyDataMetricsSmallFileRatio {
	return s.SmallFileRatio
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetTinyFileCount() *GetDoctorHiveTableResponseBodyDataMetricsTinyFileCount {
	return s.TinyFileCount
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetTinyFileCountDayGrowthRatio() *GetDoctorHiveTableResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	return s.TinyFileCountDayGrowthRatio
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetTinyFileDayGrowthCount() *GetDoctorHiveTableResponseBodyDataMetricsTinyFileDayGrowthCount {
	return s.TinyFileDayGrowthCount
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetTinyFileRatio() *GetDoctorHiveTableResponseBodyDataMetricsTinyFileRatio {
	return s.TinyFileRatio
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetTotalDataDayGrowthSize() *GetDoctorHiveTableResponseBodyDataMetricsTotalDataDayGrowthSize {
	return s.TotalDataDayGrowthSize
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetTotalDataSize() *GetDoctorHiveTableResponseBodyDataMetricsTotalDataSize {
	return s.TotalDataSize
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetTotalDataSizeDayGrowthRatio() *GetDoctorHiveTableResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	return s.TotalDataSizeDayGrowthRatio
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetTotalFileCount() *GetDoctorHiveTableResponseBodyDataMetricsTotalFileCount {
	return s.TotalFileCount
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetTotalFileCountDayGrowthRatio() *GetDoctorHiveTableResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	return s.TotalFileCountDayGrowthRatio
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetTotalFileDayGrowthCount() *GetDoctorHiveTableResponseBodyDataMetricsTotalFileDayGrowthCount {
	return s.TotalFileDayGrowthCount
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetWarmDataDayGrowthSize() *GetDoctorHiveTableResponseBodyDataMetricsWarmDataDayGrowthSize {
	return s.WarmDataDayGrowthSize
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetWarmDataRatio() *GetDoctorHiveTableResponseBodyDataMetricsWarmDataRatio {
	return s.WarmDataRatio
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetWarmDataSize() *GetDoctorHiveTableResponseBodyDataMetricsWarmDataSize {
	return s.WarmDataSize
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) GetWarmDataSizeDayGrowthRatio() *GetDoctorHiveTableResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	return s.WarmDataSizeDayGrowthRatio
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetColdDataDayGrowthSize(v *GetDoctorHiveTableResponseBodyDataMetricsColdDataDayGrowthSize) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.ColdDataDayGrowthSize = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetColdDataRatio(v *GetDoctorHiveTableResponseBodyDataMetricsColdDataRatio) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.ColdDataRatio = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetColdDataSize(v *GetDoctorHiveTableResponseBodyDataMetricsColdDataSize) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.ColdDataSize = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetColdDataSizeDayGrowthRatio(v *GetDoctorHiveTableResponseBodyDataMetricsColdDataSizeDayGrowthRatio) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.ColdDataSizeDayGrowthRatio = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetEmptyFileCount(v *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileCount) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.EmptyFileCount = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetEmptyFileCountDayGrowthRatio(v *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.EmptyFileCountDayGrowthRatio = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetEmptyFileDayGrowthCount(v *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileDayGrowthCount) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.EmptyFileDayGrowthCount = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetEmptyFileRatio(v *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileRatio) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.EmptyFileRatio = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetFreezeDataDayGrowthSize(v *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataDayGrowthSize) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.FreezeDataDayGrowthSize = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetFreezeDataRatio(v *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataRatio) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.FreezeDataRatio = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetFreezeDataSize(v *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataSize) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.FreezeDataSize = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetFreezeDataSizeDayGrowthRatio(v *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.FreezeDataSizeDayGrowthRatio = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetHotDataDayGrowthSize(v *GetDoctorHiveTableResponseBodyDataMetricsHotDataDayGrowthSize) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.HotDataDayGrowthSize = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetHotDataRatio(v *GetDoctorHiveTableResponseBodyDataMetricsHotDataRatio) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.HotDataRatio = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetHotDataSize(v *GetDoctorHiveTableResponseBodyDataMetricsHotDataSize) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.HotDataSize = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetHotDataSizeDayGrowthRatio(v *GetDoctorHiveTableResponseBodyDataMetricsHotDataSizeDayGrowthRatio) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.HotDataSizeDayGrowthRatio = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetLargeFileCount(v *GetDoctorHiveTableResponseBodyDataMetricsLargeFileCount) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.LargeFileCount = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetLargeFileCountDayGrowthRatio(v *GetDoctorHiveTableResponseBodyDataMetricsLargeFileCountDayGrowthRatio) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.LargeFileCountDayGrowthRatio = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetLargeFileDayGrowthCount(v *GetDoctorHiveTableResponseBodyDataMetricsLargeFileDayGrowthCount) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.LargeFileDayGrowthCount = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetLargeFileRatio(v *GetDoctorHiveTableResponseBodyDataMetricsLargeFileRatio) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.LargeFileRatio = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetMediumFileCount(v *GetDoctorHiveTableResponseBodyDataMetricsMediumFileCount) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.MediumFileCount = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetMediumFileCountDayGrowthRatio(v *GetDoctorHiveTableResponseBodyDataMetricsMediumFileCountDayGrowthRatio) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.MediumFileCountDayGrowthRatio = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetMediumFileDayGrowthCount(v *GetDoctorHiveTableResponseBodyDataMetricsMediumFileDayGrowthCount) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.MediumFileDayGrowthCount = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetMediumFileRatio(v *GetDoctorHiveTableResponseBodyDataMetricsMediumFileRatio) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.MediumFileRatio = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetPartitionNum(v *GetDoctorHiveTableResponseBodyDataMetricsPartitionNum) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.PartitionNum = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetSmallFileCount(v *GetDoctorHiveTableResponseBodyDataMetricsSmallFileCount) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.SmallFileCount = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetSmallFileCountDayGrowthRatio(v *GetDoctorHiveTableResponseBodyDataMetricsSmallFileCountDayGrowthRatio) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.SmallFileCountDayGrowthRatio = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetSmallFileDayGrowthCount(v *GetDoctorHiveTableResponseBodyDataMetricsSmallFileDayGrowthCount) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.SmallFileDayGrowthCount = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetSmallFileRatio(v *GetDoctorHiveTableResponseBodyDataMetricsSmallFileRatio) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.SmallFileRatio = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetTinyFileCount(v *GetDoctorHiveTableResponseBodyDataMetricsTinyFileCount) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.TinyFileCount = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetTinyFileCountDayGrowthRatio(v *GetDoctorHiveTableResponseBodyDataMetricsTinyFileCountDayGrowthRatio) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.TinyFileCountDayGrowthRatio = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetTinyFileDayGrowthCount(v *GetDoctorHiveTableResponseBodyDataMetricsTinyFileDayGrowthCount) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.TinyFileDayGrowthCount = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetTinyFileRatio(v *GetDoctorHiveTableResponseBodyDataMetricsTinyFileRatio) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.TinyFileRatio = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetTotalDataDayGrowthSize(v *GetDoctorHiveTableResponseBodyDataMetricsTotalDataDayGrowthSize) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.TotalDataDayGrowthSize = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetTotalDataSize(v *GetDoctorHiveTableResponseBodyDataMetricsTotalDataSize) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.TotalDataSize = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetTotalDataSizeDayGrowthRatio(v *GetDoctorHiveTableResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.TotalDataSizeDayGrowthRatio = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetTotalFileCount(v *GetDoctorHiveTableResponseBodyDataMetricsTotalFileCount) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.TotalFileCount = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetTotalFileCountDayGrowthRatio(v *GetDoctorHiveTableResponseBodyDataMetricsTotalFileCountDayGrowthRatio) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.TotalFileCountDayGrowthRatio = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetTotalFileDayGrowthCount(v *GetDoctorHiveTableResponseBodyDataMetricsTotalFileDayGrowthCount) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.TotalFileDayGrowthCount = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetWarmDataDayGrowthSize(v *GetDoctorHiveTableResponseBodyDataMetricsWarmDataDayGrowthSize) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.WarmDataDayGrowthSize = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetWarmDataRatio(v *GetDoctorHiveTableResponseBodyDataMetricsWarmDataRatio) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.WarmDataRatio = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetWarmDataSize(v *GetDoctorHiveTableResponseBodyDataMetricsWarmDataSize) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.WarmDataSize = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) SetWarmDataSizeDayGrowthRatio(v *GetDoctorHiveTableResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) *GetDoctorHiveTableResponseBodyDataMetrics {
	s.WarmDataSizeDayGrowthRatio = v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetrics) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsColdDataDayGrowthSize struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsColdDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsColdDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsColdDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsColdDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsColdDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsColdDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsColdDataDayGrowthSize) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsColdDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsColdDataDayGrowthSize) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsColdDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsColdDataDayGrowthSize) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsColdDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsColdDataDayGrowthSize) SetValue(v int64) *GetDoctorHiveTableResponseBodyDataMetricsColdDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsColdDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsColdDataRatio struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsColdDataRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsColdDataRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsColdDataRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsColdDataRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsColdDataRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsColdDataRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsColdDataRatio) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsColdDataRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsColdDataRatio) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsColdDataRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsColdDataRatio) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsColdDataRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsColdDataRatio) SetValue(v float32) *GetDoctorHiveTableResponseBodyDataMetricsColdDataRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsColdDataRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsColdDataSize struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsColdDataSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsColdDataSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsColdDataSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsColdDataSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsColdDataSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsColdDataSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsColdDataSize) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsColdDataSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsColdDataSize) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsColdDataSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsColdDataSize) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsColdDataSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsColdDataSize) SetValue(v int64) *GetDoctorHiveTableResponseBodyDataMetricsColdDataSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsColdDataSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsColdDataSizeDayGrowthRatio struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsColdDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsColdDataSizeDayGrowthRatio) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsColdDataSizeDayGrowthRatio) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsColdDataSizeDayGrowthRatio) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsColdDataSizeDayGrowthRatio) SetValue(v float32) *GetDoctorHiveTableResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsColdDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsEmptyFileCount struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsEmptyFileCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsEmptyFileCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileCount) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileCount) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileCount) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileCount) SetValue(v int64) *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsEmptyFileCountDayGrowthRatio struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) SetValue(v float32) *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsEmptyFileDayGrowthCount struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsEmptyFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsEmptyFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileDayGrowthCount) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileDayGrowthCount) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileDayGrowthCount) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileDayGrowthCount) SetValue(v int64) *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsEmptyFileRatio struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsEmptyFileRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsEmptyFileRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileRatio) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileRatio) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileRatio) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileRatio) SetValue(v float32) *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsEmptyFileRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsFreezeDataDayGrowthSize struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsFreezeDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsFreezeDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataDayGrowthSize) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataDayGrowthSize) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataDayGrowthSize) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataDayGrowthSize) SetValue(v int64) *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsFreezeDataRatio struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsFreezeDataRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsFreezeDataRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataRatio) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataRatio) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataRatio) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataRatio) SetValue(v float32) *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsFreezeDataSize struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsFreezeDataSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsFreezeDataSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataSize) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataSize) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataSize) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataSize) SetValue(v int64) *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) SetValue(v float32) *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsHotDataDayGrowthSize struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsHotDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsHotDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsHotDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsHotDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsHotDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsHotDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsHotDataDayGrowthSize) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsHotDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsHotDataDayGrowthSize) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsHotDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsHotDataDayGrowthSize) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsHotDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsHotDataDayGrowthSize) SetValue(v int64) *GetDoctorHiveTableResponseBodyDataMetricsHotDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsHotDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsHotDataRatio struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsHotDataRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsHotDataRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsHotDataRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsHotDataRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsHotDataRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsHotDataRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsHotDataRatio) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsHotDataRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsHotDataRatio) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsHotDataRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsHotDataRatio) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsHotDataRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsHotDataRatio) SetValue(v float32) *GetDoctorHiveTableResponseBodyDataMetricsHotDataRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsHotDataRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsHotDataSize struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsHotDataSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsHotDataSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsHotDataSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsHotDataSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsHotDataSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsHotDataSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsHotDataSize) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsHotDataSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsHotDataSize) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsHotDataSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsHotDataSize) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsHotDataSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsHotDataSize) SetValue(v int64) *GetDoctorHiveTableResponseBodyDataMetricsHotDataSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsHotDataSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsHotDataSizeDayGrowthRatio struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsHotDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsHotDataSizeDayGrowthRatio) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsHotDataSizeDayGrowthRatio) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsHotDataSizeDayGrowthRatio) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsHotDataSizeDayGrowthRatio) SetValue(v float32) *GetDoctorHiveTableResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsHotDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsLargeFileCount struct {
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
	// 123
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveTableResponseBodyDataMetricsLargeFileCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsLargeFileCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsLargeFileCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsLargeFileCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsLargeFileCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsLargeFileCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsLargeFileCount) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsLargeFileCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsLargeFileCount) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsLargeFileCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsLargeFileCount) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsLargeFileCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsLargeFileCount) SetValue(v int64) *GetDoctorHiveTableResponseBodyDataMetricsLargeFileCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsLargeFileCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsLargeFileCountDayGrowthRatio struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsLargeFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsLargeFileCountDayGrowthRatio) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsLargeFileCountDayGrowthRatio) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsLargeFileCountDayGrowthRatio) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsLargeFileCountDayGrowthRatio) SetValue(v float32) *GetDoctorHiveTableResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsLargeFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsLargeFileDayGrowthCount struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsLargeFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsLargeFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsLargeFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsLargeFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsLargeFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsLargeFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsLargeFileDayGrowthCount) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsLargeFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsLargeFileDayGrowthCount) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsLargeFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsLargeFileDayGrowthCount) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsLargeFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsLargeFileDayGrowthCount) SetValue(v int64) *GetDoctorHiveTableResponseBodyDataMetricsLargeFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsLargeFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsLargeFileRatio struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsLargeFileRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsLargeFileRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsLargeFileRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsLargeFileRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsLargeFileRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsLargeFileRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsLargeFileRatio) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsLargeFileRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsLargeFileRatio) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsLargeFileRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsLargeFileRatio) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsLargeFileRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsLargeFileRatio) SetValue(v float32) *GetDoctorHiveTableResponseBodyDataMetricsLargeFileRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsLargeFileRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsMediumFileCount struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsMediumFileCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsMediumFileCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsMediumFileCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsMediumFileCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsMediumFileCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsMediumFileCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsMediumFileCount) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsMediumFileCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsMediumFileCount) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsMediumFileCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsMediumFileCount) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsMediumFileCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsMediumFileCount) SetValue(v int64) *GetDoctorHiveTableResponseBodyDataMetricsMediumFileCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsMediumFileCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsMediumFileCountDayGrowthRatio struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsMediumFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsMediumFileCountDayGrowthRatio) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsMediumFileCountDayGrowthRatio) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsMediumFileCountDayGrowthRatio) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsMediumFileCountDayGrowthRatio) SetValue(v float32) *GetDoctorHiveTableResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsMediumFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsMediumFileDayGrowthCount struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsMediumFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsMediumFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsMediumFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsMediumFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsMediumFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsMediumFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsMediumFileDayGrowthCount) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsMediumFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsMediumFileDayGrowthCount) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsMediumFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsMediumFileDayGrowthCount) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsMediumFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsMediumFileDayGrowthCount) SetValue(v int64) *GetDoctorHiveTableResponseBodyDataMetricsMediumFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsMediumFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsMediumFileRatio struct {
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
	// 0.80
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveTableResponseBodyDataMetricsMediumFileRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsMediumFileRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsMediumFileRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsMediumFileRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsMediumFileRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsMediumFileRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsMediumFileRatio) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsMediumFileRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsMediumFileRatio) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsMediumFileRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsMediumFileRatio) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsMediumFileRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsMediumFileRatio) SetValue(v float32) *GetDoctorHiveTableResponseBodyDataMetricsMediumFileRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsMediumFileRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsPartitionNum struct {
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
	// 441
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHiveTableResponseBodyDataMetricsPartitionNum) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsPartitionNum) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsPartitionNum) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsPartitionNum) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsPartitionNum) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsPartitionNum) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsPartitionNum) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsPartitionNum {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsPartitionNum) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsPartitionNum {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsPartitionNum) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsPartitionNum {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsPartitionNum) SetValue(v int64) *GetDoctorHiveTableResponseBodyDataMetricsPartitionNum {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsPartitionNum) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsSmallFileCount struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsSmallFileCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsSmallFileCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsSmallFileCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsSmallFileCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsSmallFileCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsSmallFileCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsSmallFileCount) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsSmallFileCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsSmallFileCount) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsSmallFileCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsSmallFileCount) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsSmallFileCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsSmallFileCount) SetValue(v int64) *GetDoctorHiveTableResponseBodyDataMetricsSmallFileCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsSmallFileCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsSmallFileCountDayGrowthRatio struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsSmallFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsSmallFileCountDayGrowthRatio) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsSmallFileCountDayGrowthRatio) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsSmallFileCountDayGrowthRatio) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsSmallFileCountDayGrowthRatio) SetValue(v float32) *GetDoctorHiveTableResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsSmallFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsSmallFileDayGrowthCount struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsSmallFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsSmallFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsSmallFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsSmallFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsSmallFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsSmallFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsSmallFileDayGrowthCount) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsSmallFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsSmallFileDayGrowthCount) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsSmallFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsSmallFileDayGrowthCount) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsSmallFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsSmallFileDayGrowthCount) SetValue(v int64) *GetDoctorHiveTableResponseBodyDataMetricsSmallFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsSmallFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsSmallFileRatio struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsSmallFileRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsSmallFileRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsSmallFileRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsSmallFileRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsSmallFileRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsSmallFileRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsSmallFileRatio) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsSmallFileRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsSmallFileRatio) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsSmallFileRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsSmallFileRatio) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsSmallFileRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsSmallFileRatio) SetValue(v float32) *GetDoctorHiveTableResponseBodyDataMetricsSmallFileRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsSmallFileRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsTinyFileCount struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsTinyFileCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsTinyFileCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTinyFileCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTinyFileCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTinyFileCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTinyFileCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTinyFileCount) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsTinyFileCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTinyFileCount) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsTinyFileCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTinyFileCount) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsTinyFileCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTinyFileCount) SetValue(v int64) *GetDoctorHiveTableResponseBodyDataMetricsTinyFileCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTinyFileCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsTinyFileCountDayGrowthRatio struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsTinyFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTinyFileCountDayGrowthRatio) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTinyFileCountDayGrowthRatio) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTinyFileCountDayGrowthRatio) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTinyFileCountDayGrowthRatio) SetValue(v float32) *GetDoctorHiveTableResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTinyFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsTinyFileDayGrowthCount struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsTinyFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsTinyFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTinyFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTinyFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTinyFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTinyFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTinyFileDayGrowthCount) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsTinyFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTinyFileDayGrowthCount) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsTinyFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTinyFileDayGrowthCount) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsTinyFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTinyFileDayGrowthCount) SetValue(v int64) *GetDoctorHiveTableResponseBodyDataMetricsTinyFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTinyFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsTinyFileRatio struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsTinyFileRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsTinyFileRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTinyFileRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTinyFileRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTinyFileRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTinyFileRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTinyFileRatio) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsTinyFileRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTinyFileRatio) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsTinyFileRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTinyFileRatio) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsTinyFileRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTinyFileRatio) SetValue(v float32) *GetDoctorHiveTableResponseBodyDataMetricsTinyFileRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTinyFileRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsTotalDataDayGrowthSize struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsTotalDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsTotalDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalDataDayGrowthSize) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsTotalDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalDataDayGrowthSize) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsTotalDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalDataDayGrowthSize) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsTotalDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalDataDayGrowthSize) SetValue(v int64) *GetDoctorHiveTableResponseBodyDataMetricsTotalDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsTotalDataSize struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsTotalDataSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsTotalDataSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalDataSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalDataSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalDataSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalDataSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalDataSize) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsTotalDataSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalDataSize) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsTotalDataSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalDataSize) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsTotalDataSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalDataSize) SetValue(v int64) *GetDoctorHiveTableResponseBodyDataMetricsTotalDataSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalDataSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsTotalDataSizeDayGrowthRatio struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) SetValue(v float32) *GetDoctorHiveTableResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsTotalFileCount struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsTotalFileCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsTotalFileCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalFileCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalFileCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalFileCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalFileCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalFileCount) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsTotalFileCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalFileCount) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsTotalFileCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalFileCount) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsTotalFileCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalFileCount) SetValue(v int64) *GetDoctorHiveTableResponseBodyDataMetricsTotalFileCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalFileCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsTotalFileCountDayGrowthRatio struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsTotalFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalFileCountDayGrowthRatio) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalFileCountDayGrowthRatio) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalFileCountDayGrowthRatio) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalFileCountDayGrowthRatio) SetValue(v float32) *GetDoctorHiveTableResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsTotalFileDayGrowthCount struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsTotalFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsTotalFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalFileDayGrowthCount) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsTotalFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalFileDayGrowthCount) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsTotalFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalFileDayGrowthCount) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsTotalFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalFileDayGrowthCount) SetValue(v int64) *GetDoctorHiveTableResponseBodyDataMetricsTotalFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsTotalFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsWarmDataDayGrowthSize struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsWarmDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsWarmDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsWarmDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsWarmDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsWarmDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsWarmDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsWarmDataDayGrowthSize) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsWarmDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsWarmDataDayGrowthSize) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsWarmDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsWarmDataDayGrowthSize) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsWarmDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsWarmDataDayGrowthSize) SetValue(v int64) *GetDoctorHiveTableResponseBodyDataMetricsWarmDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsWarmDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsWarmDataRatio struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsWarmDataRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsWarmDataRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsWarmDataRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsWarmDataRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsWarmDataRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsWarmDataRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsWarmDataRatio) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsWarmDataRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsWarmDataRatio) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsWarmDataRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsWarmDataRatio) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsWarmDataRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsWarmDataRatio) SetValue(v float32) *GetDoctorHiveTableResponseBodyDataMetricsWarmDataRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsWarmDataRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsWarmDataSize struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsWarmDataSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsWarmDataSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsWarmDataSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsWarmDataSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsWarmDataSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsWarmDataSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsWarmDataSize) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsWarmDataSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsWarmDataSize) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsWarmDataSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsWarmDataSize) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsWarmDataSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsWarmDataSize) SetValue(v int64) *GetDoctorHiveTableResponseBodyDataMetricsWarmDataSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsWarmDataSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHiveTableResponseBodyDataMetricsWarmDataSizeDayGrowthRatio struct {
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

func (s GetDoctorHiveTableResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) SetDescription(v string) *GetDoctorHiveTableResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) SetName(v string) *GetDoctorHiveTableResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) SetUnit(v string) *GetDoctorHiveTableResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) SetValue(v float32) *GetDoctorHiveTableResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHiveTableResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}
