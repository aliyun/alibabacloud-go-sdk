// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDoctorHiveDatabasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListDoctorHiveDatabasesResponseBodyData) *ListDoctorHiveDatabasesResponseBody
	GetData() []*ListDoctorHiveDatabasesResponseBodyData
	SetMaxResults(v int32) *ListDoctorHiveDatabasesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDoctorHiveDatabasesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDoctorHiveDatabasesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListDoctorHiveDatabasesResponseBody
	GetTotalCount() *int32
}

type ListDoctorHiveDatabasesResponseBody struct {
	// The analysis results of Hive databases.
	Data []*ListDoctorHiveDatabasesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s ListDoctorHiveDatabasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBody) GetData() []*ListDoctorHiveDatabasesResponseBodyData {
	return s.Data
}

func (s *ListDoctorHiveDatabasesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDoctorHiveDatabasesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDoctorHiveDatabasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDoctorHiveDatabasesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDoctorHiveDatabasesResponseBody) SetData(v []*ListDoctorHiveDatabasesResponseBodyData) *ListDoctorHiveDatabasesResponseBody {
	s.Data = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBody) SetMaxResults(v int32) *ListDoctorHiveDatabasesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBody) SetNextToken(v string) *ListDoctorHiveDatabasesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBody) SetRequestId(v string) *ListDoctorHiveDatabasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBody) SetTotalCount(v int32) *ListDoctorHiveDatabasesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDoctorHiveDatabasesResponseBodyData struct {
	// The analysis results.
	Analysis *ListDoctorHiveDatabasesResponseBodyDataAnalysis `json:"Analysis,omitempty" xml:"Analysis,omitempty" type:"Struct"`
	// The database name.
	//
	// example:
	//
	// db1
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The information from the perspective of storage formats.
	Formats []*ListDoctorHiveDatabasesResponseBodyDataFormats `json:"Formats,omitempty" xml:"Formats,omitempty" type:"Repeated"`
	// The metric information.
	Metrics *ListDoctorHiveDatabasesResponseBodyDataMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Struct"`
}

func (s ListDoctorHiveDatabasesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyData) GetAnalysis() *ListDoctorHiveDatabasesResponseBodyDataAnalysis {
	return s.Analysis
}

func (s *ListDoctorHiveDatabasesResponseBodyData) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *ListDoctorHiveDatabasesResponseBodyData) GetFormats() []*ListDoctorHiveDatabasesResponseBodyDataFormats {
	return s.Formats
}

func (s *ListDoctorHiveDatabasesResponseBodyData) GetMetrics() *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	return s.Metrics
}

func (s *ListDoctorHiveDatabasesResponseBodyData) SetAnalysis(v *ListDoctorHiveDatabasesResponseBodyDataAnalysis) *ListDoctorHiveDatabasesResponseBodyData {
	s.Analysis = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyData) SetDatabaseName(v string) *ListDoctorHiveDatabasesResponseBodyData {
	s.DatabaseName = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyData) SetFormats(v []*ListDoctorHiveDatabasesResponseBodyDataFormats) *ListDoctorHiveDatabasesResponseBodyData {
	s.Formats = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyData) SetMetrics(v *ListDoctorHiveDatabasesResponseBodyDataMetrics) *ListDoctorHiveDatabasesResponseBodyData {
	s.Metrics = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyData) Validate() error {
	if s.Analysis != nil {
		if err := s.Analysis.Validate(); err != nil {
			return err
		}
	}
	if s.Formats != nil {
		for _, item := range s.Formats {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Metrics != nil {
		if err := s.Metrics.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDoctorHiveDatabasesResponseBodyDataAnalysis struct {
	// The score for the distribution of files of different sizes stored in the Hive database.
	//
	// example:
	//
	// 85
	HiveDistributionScore *int32 `json:"HiveDistributionScore,omitempty" xml:"HiveDistributionScore,omitempty"`
	// The score for the distribution of files stored in different formats in the Hive database.
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

func (s ListDoctorHiveDatabasesResponseBodyDataAnalysis) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataAnalysis) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataAnalysis) GetHiveDistributionScore() *int32 {
	return s.HiveDistributionScore
}

func (s *ListDoctorHiveDatabasesResponseBodyDataAnalysis) GetHiveFormatScore() *int32 {
	return s.HiveFormatScore
}

func (s *ListDoctorHiveDatabasesResponseBodyDataAnalysis) GetHiveFrequencyScore() *int32 {
	return s.HiveFrequencyScore
}

func (s *ListDoctorHiveDatabasesResponseBodyDataAnalysis) GetHiveScore() *int32 {
	return s.HiveScore
}

func (s *ListDoctorHiveDatabasesResponseBodyDataAnalysis) SetHiveDistributionScore(v int32) *ListDoctorHiveDatabasesResponseBodyDataAnalysis {
	s.HiveDistributionScore = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataAnalysis) SetHiveFormatScore(v int32) *ListDoctorHiveDatabasesResponseBodyDataAnalysis {
	s.HiveFormatScore = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataAnalysis) SetHiveFrequencyScore(v int32) *ListDoctorHiveDatabasesResponseBodyDataAnalysis {
	s.HiveFrequencyScore = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataAnalysis) SetHiveScore(v int32) *ListDoctorHiveDatabasesResponseBodyDataAnalysis {
	s.HiveScore = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataAnalysis) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataFormats struct {
	// The daily increment of storage format-specific data.
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
	// The proportion of data in a specific storage format.
	//
	// example:
	//
	// 0.5
	FormatRatio *float32 `json:"FormatRatio,omitempty" xml:"FormatRatio,omitempty"`
	// The amount of storage format-specific data.
	//
	// example:
	//
	// 1000
	FormatSize *int64 `json:"FormatSize,omitempty" xml:"FormatSize,omitempty"`
	// The day-to-day growth rate of storage format-specific data.
	//
	// example:
	//
	// 0.5
	FormatSizeDayGrowthRatio *float32 `json:"FormatSizeDayGrowthRatio,omitempty" xml:"FormatSizeDayGrowthRatio,omitempty"`
	// The unit of the amount of storage format-specific data.
	//
	// example:
	//
	// MB
	FormatSizeUnit *string `json:"FormatSizeUnit,omitempty" xml:"FormatSizeUnit,omitempty"`
}

func (s ListDoctorHiveDatabasesResponseBodyDataFormats) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataFormats) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataFormats) GetFormatDayGrowthSize() *int64 {
	return s.FormatDayGrowthSize
}

func (s *ListDoctorHiveDatabasesResponseBodyDataFormats) GetFormatName() *string {
	return s.FormatName
}

func (s *ListDoctorHiveDatabasesResponseBodyDataFormats) GetFormatRatio() *float32 {
	return s.FormatRatio
}

func (s *ListDoctorHiveDatabasesResponseBodyDataFormats) GetFormatSize() *int64 {
	return s.FormatSize
}

func (s *ListDoctorHiveDatabasesResponseBodyDataFormats) GetFormatSizeDayGrowthRatio() *float32 {
	return s.FormatSizeDayGrowthRatio
}

func (s *ListDoctorHiveDatabasesResponseBodyDataFormats) GetFormatSizeUnit() *string {
	return s.FormatSizeUnit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataFormats) SetFormatDayGrowthSize(v int64) *ListDoctorHiveDatabasesResponseBodyDataFormats {
	s.FormatDayGrowthSize = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataFormats) SetFormatName(v string) *ListDoctorHiveDatabasesResponseBodyDataFormats {
	s.FormatName = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataFormats) SetFormatRatio(v float32) *ListDoctorHiveDatabasesResponseBodyDataFormats {
	s.FormatRatio = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataFormats) SetFormatSize(v int64) *ListDoctorHiveDatabasesResponseBodyDataFormats {
	s.FormatSize = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataFormats) SetFormatSizeDayGrowthRatio(v float32) *ListDoctorHiveDatabasesResponseBodyDataFormats {
	s.FormatSizeDayGrowthRatio = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataFormats) SetFormatSizeUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataFormats {
	s.FormatSizeUnit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataFormats) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetrics struct {
	// The daily increment of the amount of cold data. Cold data refers to data that is not accessed for more than 30 days but is accessed in 90 days.
	ColdDataDayGrowthSize *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataDayGrowthSize `json:"ColdDataDayGrowthSize,omitempty" xml:"ColdDataDayGrowthSize,omitempty" type:"Struct"`
	// The proportion of cold data. Cold data refers to data that is not accessed for more than 30 days but is accessed in 90 days.
	ColdDataRatio *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataRatio `json:"ColdDataRatio,omitempty" xml:"ColdDataRatio,omitempty" type:"Struct"`
	// The amount of cold data. Cold data refers to data that is not accessed for more than 30 days but is accessed in 90 days.
	ColdDataSize *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataSize `json:"ColdDataSize,omitempty" xml:"ColdDataSize,omitempty" type:"Struct"`
	// The day-to-day growth rate of the amount of cold data. Cold data refers to data that is not accessed for more than 30 days but is accessed in 90 days.
	ColdDataSizeDayGrowthRatio *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataSizeDayGrowthRatio `json:"ColdDataSizeDayGrowthRatio,omitempty" xml:"ColdDataSizeDayGrowthRatio,omitempty" type:"Struct"`
	// The number of empty files. Empty files are those with a size of 0 MB.
	EmptyFileCount *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileCount `json:"EmptyFileCount,omitempty" xml:"EmptyFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the number of empty files. Empty files are those with a size of 0 MB.
	EmptyFileCountDayGrowthRatio *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio `json:"EmptyFileCountDayGrowthRatio,omitempty" xml:"EmptyFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the number of empty files. Empty files are those with a size of 0 MB.
	EmptyFileDayGrowthCount *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileDayGrowthCount `json:"EmptyFileDayGrowthCount,omitempty" xml:"EmptyFileDayGrowthCount,omitempty" type:"Struct"`
	// The proportion of empty files. Empty files are those with a size of 0 MB.
	EmptyFileRatio *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileRatio `json:"EmptyFileRatio,omitempty" xml:"EmptyFileRatio,omitempty" type:"Struct"`
	// The daily increment of the amount of very cold data. Very cold data refers to data that is not accessed for more than 90 days.
	FreezeDataDayGrowthSize *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataDayGrowthSize `json:"FreezeDataDayGrowthSize,omitempty" xml:"FreezeDataDayGrowthSize,omitempty" type:"Struct"`
	// The proportion of very cold data. Very cold data refers to data that is not accessed for more than 90 days.
	FreezeDataRatio *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataRatio `json:"FreezeDataRatio,omitempty" xml:"FreezeDataRatio,omitempty" type:"Struct"`
	// The amount of very cold data. Very cold data refers to data that is not accessed for more than 90 days.
	FreezeDataSize *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataSize `json:"FreezeDataSize,omitempty" xml:"FreezeDataSize,omitempty" type:"Struct"`
	// The day-to-day growth rate of the amount of very cold data. Very cold data refers to data that is not accessed for more than 90 days.
	FreezeDataSizeDayGrowthRatio *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio `json:"FreezeDataSizeDayGrowthRatio,omitempty" xml:"FreezeDataSizeDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the amount of hot data. Hot data refers to data that is accessed in recent seven days.
	HotDataDayGrowthSize *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataDayGrowthSize `json:"HotDataDayGrowthSize,omitempty" xml:"HotDataDayGrowthSize,omitempty" type:"Struct"`
	// The proportion of hot data. Hot data refers to data that is accessed in recent seven days.
	HotDataRatio *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataRatio `json:"HotDataRatio,omitempty" xml:"HotDataRatio,omitempty" type:"Struct"`
	// The amount of hot data. Hot data refers to data that is accessed in recent seven days.
	HotDataSize *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataSize `json:"HotDataSize,omitempty" xml:"HotDataSize,omitempty" type:"Struct"`
	// The day-to-day growth rate of the amount of hot data. Hot data refers to data that is accessed in recent seven days.
	HotDataSizeDayGrowthRatio *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataSizeDayGrowthRatio `json:"HotDataSizeDayGrowthRatio,omitempty" xml:"HotDataSizeDayGrowthRatio,omitempty" type:"Struct"`
	// The number of large files. Large files are those with a size greater than 1 GB.
	LargeFileCount *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileCount `json:"LargeFileCount,omitempty" xml:"LargeFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the number of large files. Large files are those with a size greater than 1 GB.
	LargeFileCountDayGrowthRatio *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileCountDayGrowthRatio `json:"LargeFileCountDayGrowthRatio,omitempty" xml:"LargeFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the number of large files. Large files are those with a size greater than 1 GB.
	LargeFileDayGrowthCount *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileDayGrowthCount `json:"LargeFileDayGrowthCount,omitempty" xml:"LargeFileDayGrowthCount,omitempty" type:"Struct"`
	// The proportion of large files. Large files are those with a size greater than 1 GB.
	LargeFileRatio *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileRatio `json:"LargeFileRatio,omitempty" xml:"LargeFileRatio,omitempty" type:"Struct"`
	// The number of medium files. Medium files are those with a size greater than or equal to 128 MB and less than or equal to 1 GB.
	MediumFileCount *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileCount `json:"MediumFileCount,omitempty" xml:"MediumFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the number of medium files. Medium files are those with a size greater than or equal to 128 MB and less than or equal to 1 GB.
	MediumFileCountDayGrowthRatio *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileCountDayGrowthRatio `json:"MediumFileCountDayGrowthRatio,omitempty" xml:"MediumFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the number of medium files. Medium files are those with a size greater than or equal to 128 MB and less than or equal to 1 GB.
	MediumFileDayGrowthCount *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileDayGrowthCount `json:"MediumFileDayGrowthCount,omitempty" xml:"MediumFileDayGrowthCount,omitempty" type:"Struct"`
	// The proportion of medium files. Medium files are those with a size greater than or equal to 128 MB and less than or equal to 1 GB.
	MediumFileRatio *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileRatio `json:"MediumFileRatio,omitempty" xml:"MediumFileRatio,omitempty" type:"Struct"`
	// The number of partitions.
	PartitionNum *ListDoctorHiveDatabasesResponseBodyDataMetricsPartitionNum `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty" type:"Struct"`
	// The number of small files. Small files are those with a size greater than or equal to 10 MB and less than 128 MB.
	SmallFileCount *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileCount `json:"SmallFileCount,omitempty" xml:"SmallFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the number of small files. Small files are those with a size greater than or equal to 10 MB and less than 128 MB.
	SmallFileCountDayGrowthRatio *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileCountDayGrowthRatio `json:"SmallFileCountDayGrowthRatio,omitempty" xml:"SmallFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the number of small files. Small files are those with a size greater than or equal to 10 MB and less than 128 MB.
	SmallFileDayGrowthCount *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileDayGrowthCount `json:"SmallFileDayGrowthCount,omitempty" xml:"SmallFileDayGrowthCount,omitempty" type:"Struct"`
	// The proportion of small files. Small files are those with a size greater than or equal to 10 MB and less than 128 MB.
	SmallFileRatio *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileRatio `json:"SmallFileRatio,omitempty" xml:"SmallFileRatio,omitempty" type:"Struct"`
	// The number of tables.
	TableCount *ListDoctorHiveDatabasesResponseBodyDataMetricsTableCount `json:"TableCount,omitempty" xml:"TableCount,omitempty" type:"Struct"`
	// The number of very small files. Very small files are those with a size greater than 0 MB and less than 10 MB.
	TinyFileCount *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileCount `json:"TinyFileCount,omitempty" xml:"TinyFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the number of very small files. Very small files are those with a size greater than 0 MB and less than 10 MB.
	TinyFileCountDayGrowthRatio *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileCountDayGrowthRatio `json:"TinyFileCountDayGrowthRatio,omitempty" xml:"TinyFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the number of very small files. Very small files are those with a size greater than 0 MB and less than 10 MB.
	TinyFileDayGrowthCount *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileDayGrowthCount `json:"TinyFileDayGrowthCount,omitempty" xml:"TinyFileDayGrowthCount,omitempty" type:"Struct"`
	// The proportion of very small files. Very small files are those with a size greater than 0 MB and less than 10 MB.
	TinyFileRatio *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileRatio `json:"TinyFileRatio,omitempty" xml:"TinyFileRatio,omitempty" type:"Struct"`
	// The daily increment of the total data volume.
	TotalDataDayGrowthSize *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataDayGrowthSize `json:"TotalDataDayGrowthSize,omitempty" xml:"TotalDataDayGrowthSize,omitempty" type:"Struct"`
	// The total amount of data.
	TotalDataSize *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataSize `json:"TotalDataSize,omitempty" xml:"TotalDataSize,omitempty" type:"Struct"`
	// The day-to-day growth rate of the total data volume.
	TotalDataSizeDayGrowthRatio *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio `json:"TotalDataSizeDayGrowthRatio,omitempty" xml:"TotalDataSizeDayGrowthRatio,omitempty" type:"Struct"`
	// The total number of files.
	TotalFileCount *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileCount `json:"TotalFileCount,omitempty" xml:"TotalFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the total number of files.
	TotalFileCountDayGrowthRatio *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileCountDayGrowthRatio `json:"TotalFileCountDayGrowthRatio,omitempty" xml:"TotalFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the total number of files.
	TotalFileDayGrowthCount *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileDayGrowthCount `json:"TotalFileDayGrowthCount,omitempty" xml:"TotalFileDayGrowthCount,omitempty" type:"Struct"`
	// The daily increment of the amount of warm data. Warm data refers to data that is not accessed for more than 7 days but is accessed in 30 days.
	WarmDataDayGrowthSize *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataDayGrowthSize `json:"WarmDataDayGrowthSize,omitempty" xml:"WarmDataDayGrowthSize,omitempty" type:"Struct"`
	// The proportion of warm data. Warm data refers to data that is not accessed for more than 7 days but is accessed in 30 days.
	WarmDataRatio *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataRatio `json:"WarmDataRatio,omitempty" xml:"WarmDataRatio,omitempty" type:"Struct"`
	// The amount of warm data. Warm data refers to data that is not accessed for more than 7 days but is accessed in 30 days.
	WarmDataSize *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataSize `json:"WarmDataSize,omitempty" xml:"WarmDataSize,omitempty" type:"Struct"`
	// The day-to-day growth rate of the amount of warm data. Warm data refers to data that is not accessed for more than 7 days but is accessed in 30 days.
	WarmDataSizeDayGrowthRatio *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio `json:"WarmDataSizeDayGrowthRatio,omitempty" xml:"WarmDataSizeDayGrowthRatio,omitempty" type:"Struct"`
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetrics) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetrics) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetColdDataDayGrowthSize() *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataDayGrowthSize {
	return s.ColdDataDayGrowthSize
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetColdDataRatio() *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataRatio {
	return s.ColdDataRatio
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetColdDataSize() *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataSize {
	return s.ColdDataSize
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetColdDataSizeDayGrowthRatio() *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	return s.ColdDataSizeDayGrowthRatio
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetEmptyFileCount() *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileCount {
	return s.EmptyFileCount
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetEmptyFileCountDayGrowthRatio() *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	return s.EmptyFileCountDayGrowthRatio
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetEmptyFileDayGrowthCount() *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileDayGrowthCount {
	return s.EmptyFileDayGrowthCount
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetEmptyFileRatio() *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileRatio {
	return s.EmptyFileRatio
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetFreezeDataDayGrowthSize() *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataDayGrowthSize {
	return s.FreezeDataDayGrowthSize
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetFreezeDataRatio() *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataRatio {
	return s.FreezeDataRatio
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetFreezeDataSize() *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataSize {
	return s.FreezeDataSize
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetFreezeDataSizeDayGrowthRatio() *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	return s.FreezeDataSizeDayGrowthRatio
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetHotDataDayGrowthSize() *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataDayGrowthSize {
	return s.HotDataDayGrowthSize
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetHotDataRatio() *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataRatio {
	return s.HotDataRatio
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetHotDataSize() *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataSize {
	return s.HotDataSize
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetHotDataSizeDayGrowthRatio() *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	return s.HotDataSizeDayGrowthRatio
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetLargeFileCount() *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileCount {
	return s.LargeFileCount
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetLargeFileCountDayGrowthRatio() *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	return s.LargeFileCountDayGrowthRatio
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetLargeFileDayGrowthCount() *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileDayGrowthCount {
	return s.LargeFileDayGrowthCount
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetLargeFileRatio() *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileRatio {
	return s.LargeFileRatio
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetMediumFileCount() *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileCount {
	return s.MediumFileCount
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetMediumFileCountDayGrowthRatio() *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	return s.MediumFileCountDayGrowthRatio
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetMediumFileDayGrowthCount() *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileDayGrowthCount {
	return s.MediumFileDayGrowthCount
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetMediumFileRatio() *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileRatio {
	return s.MediumFileRatio
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetPartitionNum() *ListDoctorHiveDatabasesResponseBodyDataMetricsPartitionNum {
	return s.PartitionNum
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetSmallFileCount() *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileCount {
	return s.SmallFileCount
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetSmallFileCountDayGrowthRatio() *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	return s.SmallFileCountDayGrowthRatio
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetSmallFileDayGrowthCount() *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileDayGrowthCount {
	return s.SmallFileDayGrowthCount
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetSmallFileRatio() *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileRatio {
	return s.SmallFileRatio
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetTableCount() *ListDoctorHiveDatabasesResponseBodyDataMetricsTableCount {
	return s.TableCount
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetTinyFileCount() *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileCount {
	return s.TinyFileCount
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetTinyFileCountDayGrowthRatio() *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	return s.TinyFileCountDayGrowthRatio
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetTinyFileDayGrowthCount() *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileDayGrowthCount {
	return s.TinyFileDayGrowthCount
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetTinyFileRatio() *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileRatio {
	return s.TinyFileRatio
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetTotalDataDayGrowthSize() *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataDayGrowthSize {
	return s.TotalDataDayGrowthSize
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetTotalDataSize() *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataSize {
	return s.TotalDataSize
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetTotalDataSizeDayGrowthRatio() *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	return s.TotalDataSizeDayGrowthRatio
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetTotalFileCount() *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileCount {
	return s.TotalFileCount
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetTotalFileCountDayGrowthRatio() *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	return s.TotalFileCountDayGrowthRatio
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetTotalFileDayGrowthCount() *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileDayGrowthCount {
	return s.TotalFileDayGrowthCount
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetWarmDataDayGrowthSize() *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataDayGrowthSize {
	return s.WarmDataDayGrowthSize
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetWarmDataRatio() *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataRatio {
	return s.WarmDataRatio
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetWarmDataSize() *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataSize {
	return s.WarmDataSize
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) GetWarmDataSizeDayGrowthRatio() *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	return s.WarmDataSizeDayGrowthRatio
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetColdDataDayGrowthSize(v *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataDayGrowthSize) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.ColdDataDayGrowthSize = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetColdDataRatio(v *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataRatio) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.ColdDataRatio = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetColdDataSize(v *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataSize) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.ColdDataSize = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetColdDataSizeDayGrowthRatio(v *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataSizeDayGrowthRatio) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.ColdDataSizeDayGrowthRatio = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetEmptyFileCount(v *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileCount) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.EmptyFileCount = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetEmptyFileCountDayGrowthRatio(v *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.EmptyFileCountDayGrowthRatio = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetEmptyFileDayGrowthCount(v *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileDayGrowthCount) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.EmptyFileDayGrowthCount = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetEmptyFileRatio(v *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileRatio) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.EmptyFileRatio = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetFreezeDataDayGrowthSize(v *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataDayGrowthSize) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.FreezeDataDayGrowthSize = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetFreezeDataRatio(v *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataRatio) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.FreezeDataRatio = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetFreezeDataSize(v *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataSize) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.FreezeDataSize = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetFreezeDataSizeDayGrowthRatio(v *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.FreezeDataSizeDayGrowthRatio = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetHotDataDayGrowthSize(v *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataDayGrowthSize) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.HotDataDayGrowthSize = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetHotDataRatio(v *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataRatio) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.HotDataRatio = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetHotDataSize(v *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataSize) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.HotDataSize = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetHotDataSizeDayGrowthRatio(v *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataSizeDayGrowthRatio) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.HotDataSizeDayGrowthRatio = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetLargeFileCount(v *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileCount) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.LargeFileCount = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetLargeFileCountDayGrowthRatio(v *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileCountDayGrowthRatio) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.LargeFileCountDayGrowthRatio = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetLargeFileDayGrowthCount(v *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileDayGrowthCount) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.LargeFileDayGrowthCount = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetLargeFileRatio(v *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileRatio) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.LargeFileRatio = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetMediumFileCount(v *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileCount) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.MediumFileCount = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetMediumFileCountDayGrowthRatio(v *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileCountDayGrowthRatio) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.MediumFileCountDayGrowthRatio = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetMediumFileDayGrowthCount(v *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileDayGrowthCount) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.MediumFileDayGrowthCount = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetMediumFileRatio(v *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileRatio) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.MediumFileRatio = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetPartitionNum(v *ListDoctorHiveDatabasesResponseBodyDataMetricsPartitionNum) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.PartitionNum = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetSmallFileCount(v *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileCount) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.SmallFileCount = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetSmallFileCountDayGrowthRatio(v *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileCountDayGrowthRatio) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.SmallFileCountDayGrowthRatio = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetSmallFileDayGrowthCount(v *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileDayGrowthCount) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.SmallFileDayGrowthCount = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetSmallFileRatio(v *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileRatio) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.SmallFileRatio = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetTableCount(v *ListDoctorHiveDatabasesResponseBodyDataMetricsTableCount) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.TableCount = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetTinyFileCount(v *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileCount) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.TinyFileCount = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetTinyFileCountDayGrowthRatio(v *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileCountDayGrowthRatio) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.TinyFileCountDayGrowthRatio = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetTinyFileDayGrowthCount(v *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileDayGrowthCount) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.TinyFileDayGrowthCount = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetTinyFileRatio(v *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileRatio) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.TinyFileRatio = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetTotalDataDayGrowthSize(v *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataDayGrowthSize) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.TotalDataDayGrowthSize = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetTotalDataSize(v *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataSize) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.TotalDataSize = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetTotalDataSizeDayGrowthRatio(v *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.TotalDataSizeDayGrowthRatio = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetTotalFileCount(v *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileCount) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.TotalFileCount = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetTotalFileCountDayGrowthRatio(v *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileCountDayGrowthRatio) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.TotalFileCountDayGrowthRatio = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetTotalFileDayGrowthCount(v *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileDayGrowthCount) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.TotalFileDayGrowthCount = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetWarmDataDayGrowthSize(v *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataDayGrowthSize) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.WarmDataDayGrowthSize = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetWarmDataRatio(v *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataRatio) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.WarmDataRatio = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetWarmDataSize(v *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataSize) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.WarmDataSize = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) SetWarmDataSizeDayGrowthRatio(v *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) *ListDoctorHiveDatabasesResponseBodyDataMetrics {
	s.WarmDataSizeDayGrowthRatio = v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetrics) Validate() error {
	if s.ColdDataDayGrowthSize != nil {
		if err := s.ColdDataDayGrowthSize.Validate(); err != nil {
			return err
		}
	}
	if s.ColdDataRatio != nil {
		if err := s.ColdDataRatio.Validate(); err != nil {
			return err
		}
	}
	if s.ColdDataSize != nil {
		if err := s.ColdDataSize.Validate(); err != nil {
			return err
		}
	}
	if s.ColdDataSizeDayGrowthRatio != nil {
		if err := s.ColdDataSizeDayGrowthRatio.Validate(); err != nil {
			return err
		}
	}
	if s.EmptyFileCount != nil {
		if err := s.EmptyFileCount.Validate(); err != nil {
			return err
		}
	}
	if s.EmptyFileCountDayGrowthRatio != nil {
		if err := s.EmptyFileCountDayGrowthRatio.Validate(); err != nil {
			return err
		}
	}
	if s.EmptyFileDayGrowthCount != nil {
		if err := s.EmptyFileDayGrowthCount.Validate(); err != nil {
			return err
		}
	}
	if s.EmptyFileRatio != nil {
		if err := s.EmptyFileRatio.Validate(); err != nil {
			return err
		}
	}
	if s.FreezeDataDayGrowthSize != nil {
		if err := s.FreezeDataDayGrowthSize.Validate(); err != nil {
			return err
		}
	}
	if s.FreezeDataRatio != nil {
		if err := s.FreezeDataRatio.Validate(); err != nil {
			return err
		}
	}
	if s.FreezeDataSize != nil {
		if err := s.FreezeDataSize.Validate(); err != nil {
			return err
		}
	}
	if s.FreezeDataSizeDayGrowthRatio != nil {
		if err := s.FreezeDataSizeDayGrowthRatio.Validate(); err != nil {
			return err
		}
	}
	if s.HotDataDayGrowthSize != nil {
		if err := s.HotDataDayGrowthSize.Validate(); err != nil {
			return err
		}
	}
	if s.HotDataRatio != nil {
		if err := s.HotDataRatio.Validate(); err != nil {
			return err
		}
	}
	if s.HotDataSize != nil {
		if err := s.HotDataSize.Validate(); err != nil {
			return err
		}
	}
	if s.HotDataSizeDayGrowthRatio != nil {
		if err := s.HotDataSizeDayGrowthRatio.Validate(); err != nil {
			return err
		}
	}
	if s.LargeFileCount != nil {
		if err := s.LargeFileCount.Validate(); err != nil {
			return err
		}
	}
	if s.LargeFileCountDayGrowthRatio != nil {
		if err := s.LargeFileCountDayGrowthRatio.Validate(); err != nil {
			return err
		}
	}
	if s.LargeFileDayGrowthCount != nil {
		if err := s.LargeFileDayGrowthCount.Validate(); err != nil {
			return err
		}
	}
	if s.LargeFileRatio != nil {
		if err := s.LargeFileRatio.Validate(); err != nil {
			return err
		}
	}
	if s.MediumFileCount != nil {
		if err := s.MediumFileCount.Validate(); err != nil {
			return err
		}
	}
	if s.MediumFileCountDayGrowthRatio != nil {
		if err := s.MediumFileCountDayGrowthRatio.Validate(); err != nil {
			return err
		}
	}
	if s.MediumFileDayGrowthCount != nil {
		if err := s.MediumFileDayGrowthCount.Validate(); err != nil {
			return err
		}
	}
	if s.MediumFileRatio != nil {
		if err := s.MediumFileRatio.Validate(); err != nil {
			return err
		}
	}
	if s.PartitionNum != nil {
		if err := s.PartitionNum.Validate(); err != nil {
			return err
		}
	}
	if s.SmallFileCount != nil {
		if err := s.SmallFileCount.Validate(); err != nil {
			return err
		}
	}
	if s.SmallFileCountDayGrowthRatio != nil {
		if err := s.SmallFileCountDayGrowthRatio.Validate(); err != nil {
			return err
		}
	}
	if s.SmallFileDayGrowthCount != nil {
		if err := s.SmallFileDayGrowthCount.Validate(); err != nil {
			return err
		}
	}
	if s.SmallFileRatio != nil {
		if err := s.SmallFileRatio.Validate(); err != nil {
			return err
		}
	}
	if s.TableCount != nil {
		if err := s.TableCount.Validate(); err != nil {
			return err
		}
	}
	if s.TinyFileCount != nil {
		if err := s.TinyFileCount.Validate(); err != nil {
			return err
		}
	}
	if s.TinyFileCountDayGrowthRatio != nil {
		if err := s.TinyFileCountDayGrowthRatio.Validate(); err != nil {
			return err
		}
	}
	if s.TinyFileDayGrowthCount != nil {
		if err := s.TinyFileDayGrowthCount.Validate(); err != nil {
			return err
		}
	}
	if s.TinyFileRatio != nil {
		if err := s.TinyFileRatio.Validate(); err != nil {
			return err
		}
	}
	if s.TotalDataDayGrowthSize != nil {
		if err := s.TotalDataDayGrowthSize.Validate(); err != nil {
			return err
		}
	}
	if s.TotalDataSize != nil {
		if err := s.TotalDataSize.Validate(); err != nil {
			return err
		}
	}
	if s.TotalDataSizeDayGrowthRatio != nil {
		if err := s.TotalDataSizeDayGrowthRatio.Validate(); err != nil {
			return err
		}
	}
	if s.TotalFileCount != nil {
		if err := s.TotalFileCount.Validate(); err != nil {
			return err
		}
	}
	if s.TotalFileCountDayGrowthRatio != nil {
		if err := s.TotalFileCountDayGrowthRatio.Validate(); err != nil {
			return err
		}
	}
	if s.TotalFileDayGrowthCount != nil {
		if err := s.TotalFileDayGrowthCount.Validate(); err != nil {
			return err
		}
	}
	if s.WarmDataDayGrowthSize != nil {
		if err := s.WarmDataDayGrowthSize.Validate(); err != nil {
			return err
		}
	}
	if s.WarmDataRatio != nil {
		if err := s.WarmDataRatio.Validate(); err != nil {
			return err
		}
	}
	if s.WarmDataSize != nil {
		if err := s.WarmDataSize.Validate(); err != nil {
			return err
		}
	}
	if s.WarmDataSizeDayGrowthRatio != nil {
		if err := s.WarmDataSizeDayGrowthRatio.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataDayGrowthSize struct {
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataDayGrowthSize) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataDayGrowthSize) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataDayGrowthSize) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataDayGrowthSize) SetValue(v int64) *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataRatio struct {
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataRatio) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataRatio) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataRatio) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataRatio) SetValue(v float32) *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataSize struct {
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataSize) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataSize) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataSize) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataSize) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataSize) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataSize) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataSize) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataSize {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataSize) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataSize {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataSize) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataSize {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataSize) SetValue(v int64) *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataSize {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataSize) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataSizeDayGrowthRatio struct {
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataSizeDayGrowthRatio) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataSizeDayGrowthRatio) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataSizeDayGrowthRatio) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataSizeDayGrowthRatio) SetValue(v float32) *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsColdDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileCount struct {
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileCount) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileCount) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileCount) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileCount) SetValue(v int64) *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio struct {
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) SetValue(v float32) *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileDayGrowthCount struct {
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileDayGrowthCount) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileDayGrowthCount) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileDayGrowthCount) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileDayGrowthCount) SetValue(v int64) *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileRatio struct {
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileRatio) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileRatio) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileRatio) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileRatio) SetValue(v float32) *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsEmptyFileRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataDayGrowthSize struct {
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
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataDayGrowthSize) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataDayGrowthSize) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataDayGrowthSize) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataDayGrowthSize) SetValue(v int64) *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataRatio struct {
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataRatio) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataRatio) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataRatio) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataRatio) SetValue(v float32) *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataSize struct {
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataSize) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataSize) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataSize) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataSize) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataSize) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataSize) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataSize) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataSize {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataSize) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataSize {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataSize) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataSize {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataSize) SetValue(v int64) *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataSize {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataSize) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio struct {
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) SetValue(v float32) *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataDayGrowthSize struct {
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataDayGrowthSize) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataDayGrowthSize) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataDayGrowthSize) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataDayGrowthSize) SetValue(v int64) *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataRatio struct {
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataRatio) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataRatio) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataRatio) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataRatio) SetValue(v float32) *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataSize struct {
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataSize) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataSize) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataSize) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataSize) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataSize) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataSize) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataSize) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataSize {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataSize) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataSize {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataSize) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataSize {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataSize) SetValue(v int64) *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataSize {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataSize) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataSizeDayGrowthRatio struct {
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataSizeDayGrowthRatio) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataSizeDayGrowthRatio) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataSizeDayGrowthRatio) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataSizeDayGrowthRatio) SetValue(v float32) *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsHotDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileCount struct {
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
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileCount) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileCount) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileCount) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileCount) SetValue(v int64) *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileCountDayGrowthRatio struct {
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileCountDayGrowthRatio) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileCountDayGrowthRatio) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileCountDayGrowthRatio) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileCountDayGrowthRatio) SetValue(v float32) *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileDayGrowthCount struct {
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileDayGrowthCount) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileDayGrowthCount) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileDayGrowthCount) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileDayGrowthCount) SetValue(v int64) *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileRatio struct {
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileRatio) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileRatio) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileRatio) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileRatio) SetValue(v float32) *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsLargeFileRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileCount struct {
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileCount) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileCount) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileCount) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileCount) SetValue(v int64) *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileCountDayGrowthRatio struct {
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileCountDayGrowthRatio) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileCountDayGrowthRatio) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileCountDayGrowthRatio) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileCountDayGrowthRatio) SetValue(v float32) *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileDayGrowthCount struct {
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileDayGrowthCount) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileDayGrowthCount) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileDayGrowthCount) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileDayGrowthCount) SetValue(v int64) *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileRatio struct {
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileRatio) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileRatio) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileRatio) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileRatio) SetValue(v float32) *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsMediumFileRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsPartitionNum struct {
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
	// ppartitionNum
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsPartitionNum) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsPartitionNum) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsPartitionNum) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsPartitionNum) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsPartitionNum) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsPartitionNum) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsPartitionNum) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsPartitionNum {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsPartitionNum) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsPartitionNum {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsPartitionNum) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsPartitionNum {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsPartitionNum) SetValue(v int64) *ListDoctorHiveDatabasesResponseBodyDataMetricsPartitionNum {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsPartitionNum) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileCount struct {
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileCount) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileCount) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileCount) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileCount) SetValue(v int64) *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileCountDayGrowthRatio struct {
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileCountDayGrowthRatio) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileCountDayGrowthRatio) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileCountDayGrowthRatio) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileCountDayGrowthRatio) SetValue(v float32) *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileDayGrowthCount struct {
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileDayGrowthCount) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileDayGrowthCount) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileDayGrowthCount) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileDayGrowthCount) SetValue(v int64) *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileRatio struct {
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileRatio) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileRatio) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileRatio) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileRatio) SetValue(v float32) *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsSmallFileRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsTableCount struct {
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
	// TableCount
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsTableCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsTableCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTableCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTableCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTableCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTableCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTableCount) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsTableCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTableCount) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsTableCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTableCount) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsTableCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTableCount) SetValue(v int64) *ListDoctorHiveDatabasesResponseBodyDataMetricsTableCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTableCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileCount struct {
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileCount) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileCount) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileCount) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileCount) SetValue(v int64) *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileCountDayGrowthRatio struct {
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileCountDayGrowthRatio) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileCountDayGrowthRatio) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileCountDayGrowthRatio) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileCountDayGrowthRatio) SetValue(v float32) *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileDayGrowthCount struct {
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileDayGrowthCount) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileDayGrowthCount) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileDayGrowthCount) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileDayGrowthCount) SetValue(v int64) *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileRatio struct {
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileRatio) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileRatio) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileRatio) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileRatio) SetValue(v float32) *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTinyFileRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataDayGrowthSize struct {
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataDayGrowthSize) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataDayGrowthSize) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataDayGrowthSize) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataDayGrowthSize) SetValue(v int64) *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataSize struct {
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataSize) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataSize) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataSize) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataSize) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataSize) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataSize) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataSize) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataSize {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataSize) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataSize {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataSize) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataSize {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataSize) SetValue(v int64) *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataSize {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataSize) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio struct {
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) SetValue(v float32) *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileCount struct {
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileCount) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileCount) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileCount) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileCount) SetValue(v int64) *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileCountDayGrowthRatio struct {
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileCountDayGrowthRatio) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileCountDayGrowthRatio) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileCountDayGrowthRatio) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileCountDayGrowthRatio) SetValue(v float32) *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileDayGrowthCount struct {
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileDayGrowthCount) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileDayGrowthCount) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileDayGrowthCount) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileDayGrowthCount) SetValue(v int64) *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsTotalFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataDayGrowthSize struct {
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
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataDayGrowthSize) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataDayGrowthSize) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataDayGrowthSize) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataDayGrowthSize) SetValue(v int64) *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataRatio struct {
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataRatio) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataRatio) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataRatio) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataRatio) SetValue(v float32) *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataSize struct {
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataSize) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataSize) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataSize) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataSize) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataSize) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataSize) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataSize) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataSize {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataSize) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataSize {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataSize) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataSize {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataSize) SetValue(v int64) *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataSize {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataSize) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio struct {
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
	// WarmDataSizeDayGrowthRatio
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

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) SetDescription(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) SetName(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) SetUnit(v string) *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) SetValue(v float32) *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}
