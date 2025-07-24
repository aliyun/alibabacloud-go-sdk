// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorHDFSClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDoctorHDFSClusterResponseBodyData) *GetDoctorHDFSClusterResponseBody
	GetData() *GetDoctorHDFSClusterResponseBodyData
	SetRequestId(v string) *GetDoctorHDFSClusterResponseBody
	GetRequestId() *string
}

type GetDoctorHDFSClusterResponseBody struct {
	// The HDFS analysis results.
	Data *GetDoctorHDFSClusterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBody) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBody) GetData() *GetDoctorHDFSClusterResponseBodyData {
	return s.Data
}

func (s *GetDoctorHDFSClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDoctorHDFSClusterResponseBody) SetData(v *GetDoctorHDFSClusterResponseBodyData) *GetDoctorHDFSClusterResponseBody {
	s.Data = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBody) SetRequestId(v string) *GetDoctorHDFSClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyData struct {
	// The analysis results.
	Analysis *GetDoctorHDFSClusterResponseBodyDataAnalysis `json:"Analysis,omitempty" xml:"Analysis,omitempty" type:"Struct"`
	// The metric information.
	Metrics *GetDoctorHDFSClusterResponseBodyDataMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Struct"`
}

func (s GetDoctorHDFSClusterResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyData) GetAnalysis() *GetDoctorHDFSClusterResponseBodyDataAnalysis {
	return s.Analysis
}

func (s *GetDoctorHDFSClusterResponseBodyData) GetMetrics() *GetDoctorHDFSClusterResponseBodyDataMetrics {
	return s.Metrics
}

func (s *GetDoctorHDFSClusterResponseBodyData) SetAnalysis(v *GetDoctorHDFSClusterResponseBodyDataAnalysis) *GetDoctorHDFSClusterResponseBodyData {
	s.Analysis = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyData) SetMetrics(v *GetDoctorHDFSClusterResponseBodyDataMetrics) *GetDoctorHDFSClusterResponseBodyData {
	s.Metrics = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataAnalysis struct {
	// The overall score of HDFS storage resources.
	//
	// example:
	//
	// 55
	HdfsScore *int32 `json:"HdfsScore,omitempty" xml:"HdfsScore,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataAnalysis) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataAnalysis) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataAnalysis) GetHdfsScore() *int32 {
	return s.HdfsScore
}

func (s *GetDoctorHDFSClusterResponseBodyDataAnalysis) SetHdfsScore(v int32) *GetDoctorHDFSClusterResponseBodyDataAnalysis {
	s.HdfsScore = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataAnalysis) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetrics struct {
	// The daily increment of the amount of cold data. Cold data refers to data that is not accessed for more than 30 days but is accessed in 90 days.
	ColdDataDayGrowthSize *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataDayGrowthSize `json:"ColdDataDayGrowthSize,omitempty" xml:"ColdDataDayGrowthSize,omitempty" type:"Struct"`
	// The proportion of cold data. Cold data refers to data that is not accessed for more than 30 days but is accessed in 90 days.
	ColdDataRatio *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataRatio `json:"ColdDataRatio,omitempty" xml:"ColdDataRatio,omitempty" type:"Struct"`
	// The amount of cold data. Cold data refers to data that is not accessed for more than 30 days but is accessed in 90 days.
	ColdDataSize *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataSize `json:"ColdDataSize,omitempty" xml:"ColdDataSize,omitempty" type:"Struct"`
	// The day-to-day growth rate of the amount of cold data. Cold data refers to data that is not accessed for more than 30 days but is accessed in 90 days.
	ColdDataSizeDayGrowthRatio *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataSizeDayGrowthRatio `json:"ColdDataSizeDayGrowthRatio,omitempty" xml:"ColdDataSizeDayGrowthRatio,omitempty" type:"Struct"`
	// The number of empty files. Empty files are those with a size of 0 MB.
	EmptyFileCount *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileCount `json:"EmptyFileCount,omitempty" xml:"EmptyFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the number of empty files. Empty files are those with a size of 0 MB.
	EmptyFileCountDayGrowthRatio *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileCountDayGrowthRatio `json:"EmptyFileCountDayGrowthRatio,omitempty" xml:"EmptyFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the number of empty files. Empty files are those with a size of 0 MB.
	EmptyFileDayGrowthCount *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileDayGrowthCount `json:"EmptyFileDayGrowthCount,omitempty" xml:"EmptyFileDayGrowthCount,omitempty" type:"Struct"`
	// The proportion of empty files. Empty files are those with a size of 0 MB.
	EmptyFileRatio *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileRatio `json:"EmptyFileRatio,omitempty" xml:"EmptyFileRatio,omitempty" type:"Struct"`
	// The daily increment of the amount of very cold data. Very cold data refers to data that is not accessed for more than 90 days.
	FreezeDataDayGrowthSize *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataDayGrowthSize `json:"FreezeDataDayGrowthSize,omitempty" xml:"FreezeDataDayGrowthSize,omitempty" type:"Struct"`
	// The proportion of very cold data. Very cold data refers to data that is not accessed for more than 90 days.
	FreezeDataRatio *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataRatio `json:"FreezeDataRatio,omitempty" xml:"FreezeDataRatio,omitempty" type:"Struct"`
	// The amount of very cold data. Very cold data refers to data that is not accessed for more than 90 days.
	FreezeDataSize *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataSize `json:"FreezeDataSize,omitempty" xml:"FreezeDataSize,omitempty" type:"Struct"`
	// The day-to-day growth rate of the amount of very cold data. Very cold data refers to data that is not accessed for more than 90 days.
	FreezeDataSizeDayGrowthRatio *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio `json:"FreezeDataSizeDayGrowthRatio,omitempty" xml:"FreezeDataSizeDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the amount of hot data. Hot data refers to data that is accessed in recent seven days.
	HotDataDayGrowthSize *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataDayGrowthSize `json:"HotDataDayGrowthSize,omitempty" xml:"HotDataDayGrowthSize,omitempty" type:"Struct"`
	// The proportion of hot data. Hot data refers to data that is accessed in recent seven days.
	HotDataRatio *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataRatio `json:"HotDataRatio,omitempty" xml:"HotDataRatio,omitempty" type:"Struct"`
	// The amount of hot data. Hot data refers to data that is accessed in recent seven days.
	HotDataSize *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataSize `json:"HotDataSize,omitempty" xml:"HotDataSize,omitempty" type:"Struct"`
	// The day-to-day growth rate of the amount of hot data. Hot data refers to data that is accessed in recent seven days.
	HotDataSizeDayGrowthRatio *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataSizeDayGrowthRatio `json:"HotDataSizeDayGrowthRatio,omitempty" xml:"HotDataSizeDayGrowthRatio,omitempty" type:"Struct"`
	// The number of large files. Large files are those with a size greater than 1 GB.
	LargeFileCount *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileCount `json:"LargeFileCount,omitempty" xml:"LargeFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the number of large files. Large files are those with a size greater than 1 GB.
	LargeFileCountDayGrowthRatio *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileCountDayGrowthRatio `json:"LargeFileCountDayGrowthRatio,omitempty" xml:"LargeFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the number of large files. Large files are those with a size greater than 1 GB.
	LargeFileDayGrowthCount *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileDayGrowthCount `json:"LargeFileDayGrowthCount,omitempty" xml:"LargeFileDayGrowthCount,omitempty" type:"Struct"`
	// The proportion of large files. Large files are those with a size greater than 1 GB.
	LargeFileRatio *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileRatio `json:"LargeFileRatio,omitempty" xml:"LargeFileRatio,omitempty" type:"Struct"`
	// The number of medium files. Medium files are those with a size greater than or equal to 128 MB and less than or equal to 1 GB.
	MediumFileCount *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileCount `json:"MediumFileCount,omitempty" xml:"MediumFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the number of medium files. Medium files are those with a size greater than or equal to 128 MB and less than or equal to 1 GB.
	MediumFileCountDayGrowthRatio *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileCountDayGrowthRatio `json:"MediumFileCountDayGrowthRatio,omitempty" xml:"MediumFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the number of medium files. Medium files are those with a size greater than or equal to 128 MB and less than or equal to 1 GB.
	MediumFileDayGrowthCount *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileDayGrowthCount `json:"MediumFileDayGrowthCount,omitempty" xml:"MediumFileDayGrowthCount,omitempty" type:"Struct"`
	// The proportion of medium files. Medium files are those with a size greater than or equal to 128 MB and less than or equal to 1 GB.
	MediumFileRatio *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileRatio `json:"MediumFileRatio,omitempty" xml:"MediumFileRatio,omitempty" type:"Struct"`
	// The number of small files. Small files are those with a size greater than or equal to 10 MB and less than 128 MB.
	SmallFileCount *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileCount `json:"SmallFileCount,omitempty" xml:"SmallFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the number of small files. Small files are those with a size greater than or equal to 10 MB and less than 128 MB.
	SmallFileCountDayGrowthRatio *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileCountDayGrowthRatio `json:"SmallFileCountDayGrowthRatio,omitempty" xml:"SmallFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the number of small files. Small files are those with a size greater than or equal to 10 MB and less than 128 MB.
	SmallFileDayGrowthCount *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileDayGrowthCount `json:"SmallFileDayGrowthCount,omitempty" xml:"SmallFileDayGrowthCount,omitempty" type:"Struct"`
	// The proportion of small files. Small files are those with a size greater than or equal to 10 MB and less than 128 MB.
	SmallFileRatio *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileRatio `json:"SmallFileRatio,omitempty" xml:"SmallFileRatio,omitempty" type:"Struct"`
	// The number of very small files. Very small files are those with a size greater than 0 MB and less than 10 MB.
	TinyFileCount *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileCount `json:"TinyFileCount,omitempty" xml:"TinyFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the number of very small files. Very small files are those with a size greater than 0 MB and less than 10 MB.
	TinyFileCountDayGrowthRatio *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileCountDayGrowthRatio `json:"TinyFileCountDayGrowthRatio,omitempty" xml:"TinyFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the number of very small files. Very small files are those with a size greater than 0 MB and less than 10 MB.
	TinyFileDayGrowthCount *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileDayGrowthCount `json:"TinyFileDayGrowthCount,omitempty" xml:"TinyFileDayGrowthCount,omitempty" type:"Struct"`
	// The proportion of very small files. Very small files are those with a size greater than 0 MB and less than 10 MB.
	TinyFileRatio *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileRatio `json:"TinyFileRatio,omitempty" xml:"TinyFileRatio,omitempty" type:"Struct"`
	// The daily incremental of the total data volume.
	TotalDataDayGrowthSize *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataDayGrowthSize `json:"TotalDataDayGrowthSize,omitempty" xml:"TotalDataDayGrowthSize,omitempty" type:"Struct"`
	// The total amount of data.
	TotalDataSize *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataSize `json:"TotalDataSize,omitempty" xml:"TotalDataSize,omitempty" type:"Struct"`
	// The day-to-day growth rate of the total data volume.
	TotalDataSizeDayGrowthRatio *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataSizeDayGrowthRatio `json:"TotalDataSizeDayGrowthRatio,omitempty" xml:"TotalDataSizeDayGrowthRatio,omitempty" type:"Struct"`
	// The total number of files.
	TotalFileCount *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileCount `json:"TotalFileCount,omitempty" xml:"TotalFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the total number of files.
	TotalFileCountDayGrowthRatio *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileCountDayGrowthRatio `json:"TotalFileCountDayGrowthRatio,omitempty" xml:"TotalFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the total number of files.
	TotalFileDayGrowthCount *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileDayGrowthCount `json:"TotalFileDayGrowthCount,omitempty" xml:"TotalFileDayGrowthCount,omitempty" type:"Struct"`
	// The daily increment of the amount of warm data. Warm data refers to data that is not accessed for more than 7 days but is accessed in 30 days.
	WarmDataDayGrowthSize *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataDayGrowthSize `json:"WarmDataDayGrowthSize,omitempty" xml:"WarmDataDayGrowthSize,omitempty" type:"Struct"`
	// The proportion of warm data. Warm data refers to data that is not accessed for more than 7 days but is accessed in 30 days.
	WarmDataRatio *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataRatio `json:"WarmDataRatio,omitempty" xml:"WarmDataRatio,omitempty" type:"Struct"`
	// The amount of warm data. Warm data refers to data that is not accessed for more than 7 days but is accessed in 30 days.
	WarmDataSize *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataSize `json:"WarmDataSize,omitempty" xml:"WarmDataSize,omitempty" type:"Struct"`
	// The day-to-day growth rate of the amount of warm data. Warm data refers to data that is not accessed for more than 7 days but is accessed in 30 days.
	WarmDataSizeDayGrowthRatio *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataSizeDayGrowthRatio `json:"WarmDataSizeDayGrowthRatio,omitempty" xml:"WarmDataSizeDayGrowthRatio,omitempty" type:"Struct"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetrics) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetrics) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetColdDataDayGrowthSize() *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataDayGrowthSize {
	return s.ColdDataDayGrowthSize
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetColdDataRatio() *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataRatio {
	return s.ColdDataRatio
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetColdDataSize() *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataSize {
	return s.ColdDataSize
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetColdDataSizeDayGrowthRatio() *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	return s.ColdDataSizeDayGrowthRatio
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetEmptyFileCount() *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileCount {
	return s.EmptyFileCount
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetEmptyFileCountDayGrowthRatio() *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	return s.EmptyFileCountDayGrowthRatio
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetEmptyFileDayGrowthCount() *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileDayGrowthCount {
	return s.EmptyFileDayGrowthCount
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetEmptyFileRatio() *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileRatio {
	return s.EmptyFileRatio
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetFreezeDataDayGrowthSize() *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataDayGrowthSize {
	return s.FreezeDataDayGrowthSize
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetFreezeDataRatio() *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataRatio {
	return s.FreezeDataRatio
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetFreezeDataSize() *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataSize {
	return s.FreezeDataSize
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetFreezeDataSizeDayGrowthRatio() *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	return s.FreezeDataSizeDayGrowthRatio
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetHotDataDayGrowthSize() *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataDayGrowthSize {
	return s.HotDataDayGrowthSize
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetHotDataRatio() *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataRatio {
	return s.HotDataRatio
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetHotDataSize() *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataSize {
	return s.HotDataSize
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetHotDataSizeDayGrowthRatio() *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	return s.HotDataSizeDayGrowthRatio
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetLargeFileCount() *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileCount {
	return s.LargeFileCount
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetLargeFileCountDayGrowthRatio() *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	return s.LargeFileCountDayGrowthRatio
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetLargeFileDayGrowthCount() *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileDayGrowthCount {
	return s.LargeFileDayGrowthCount
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetLargeFileRatio() *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileRatio {
	return s.LargeFileRatio
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetMediumFileCount() *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileCount {
	return s.MediumFileCount
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetMediumFileCountDayGrowthRatio() *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	return s.MediumFileCountDayGrowthRatio
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetMediumFileDayGrowthCount() *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileDayGrowthCount {
	return s.MediumFileDayGrowthCount
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetMediumFileRatio() *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileRatio {
	return s.MediumFileRatio
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetSmallFileCount() *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileCount {
	return s.SmallFileCount
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetSmallFileCountDayGrowthRatio() *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	return s.SmallFileCountDayGrowthRatio
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetSmallFileDayGrowthCount() *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileDayGrowthCount {
	return s.SmallFileDayGrowthCount
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetSmallFileRatio() *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileRatio {
	return s.SmallFileRatio
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetTinyFileCount() *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileCount {
	return s.TinyFileCount
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetTinyFileCountDayGrowthRatio() *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	return s.TinyFileCountDayGrowthRatio
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetTinyFileDayGrowthCount() *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileDayGrowthCount {
	return s.TinyFileDayGrowthCount
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetTinyFileRatio() *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileRatio {
	return s.TinyFileRatio
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetTotalDataDayGrowthSize() *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataDayGrowthSize {
	return s.TotalDataDayGrowthSize
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetTotalDataSize() *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataSize {
	return s.TotalDataSize
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetTotalDataSizeDayGrowthRatio() *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	return s.TotalDataSizeDayGrowthRatio
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetTotalFileCount() *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileCount {
	return s.TotalFileCount
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetTotalFileCountDayGrowthRatio() *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	return s.TotalFileCountDayGrowthRatio
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetTotalFileDayGrowthCount() *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileDayGrowthCount {
	return s.TotalFileDayGrowthCount
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetWarmDataDayGrowthSize() *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataDayGrowthSize {
	return s.WarmDataDayGrowthSize
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetWarmDataRatio() *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataRatio {
	return s.WarmDataRatio
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetWarmDataSize() *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataSize {
	return s.WarmDataSize
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) GetWarmDataSizeDayGrowthRatio() *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	return s.WarmDataSizeDayGrowthRatio
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetColdDataDayGrowthSize(v *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataDayGrowthSize) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.ColdDataDayGrowthSize = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetColdDataRatio(v *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataRatio) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.ColdDataRatio = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetColdDataSize(v *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataSize) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.ColdDataSize = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetColdDataSizeDayGrowthRatio(v *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataSizeDayGrowthRatio) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.ColdDataSizeDayGrowthRatio = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetEmptyFileCount(v *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileCount) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.EmptyFileCount = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetEmptyFileCountDayGrowthRatio(v *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.EmptyFileCountDayGrowthRatio = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetEmptyFileDayGrowthCount(v *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileDayGrowthCount) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.EmptyFileDayGrowthCount = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetEmptyFileRatio(v *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileRatio) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.EmptyFileRatio = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetFreezeDataDayGrowthSize(v *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataDayGrowthSize) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.FreezeDataDayGrowthSize = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetFreezeDataRatio(v *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataRatio) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.FreezeDataRatio = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetFreezeDataSize(v *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataSize) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.FreezeDataSize = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetFreezeDataSizeDayGrowthRatio(v *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.FreezeDataSizeDayGrowthRatio = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetHotDataDayGrowthSize(v *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataDayGrowthSize) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.HotDataDayGrowthSize = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetHotDataRatio(v *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataRatio) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.HotDataRatio = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetHotDataSize(v *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataSize) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.HotDataSize = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetHotDataSizeDayGrowthRatio(v *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataSizeDayGrowthRatio) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.HotDataSizeDayGrowthRatio = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetLargeFileCount(v *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileCount) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.LargeFileCount = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetLargeFileCountDayGrowthRatio(v *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileCountDayGrowthRatio) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.LargeFileCountDayGrowthRatio = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetLargeFileDayGrowthCount(v *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileDayGrowthCount) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.LargeFileDayGrowthCount = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetLargeFileRatio(v *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileRatio) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.LargeFileRatio = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetMediumFileCount(v *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileCount) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.MediumFileCount = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetMediumFileCountDayGrowthRatio(v *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileCountDayGrowthRatio) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.MediumFileCountDayGrowthRatio = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetMediumFileDayGrowthCount(v *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileDayGrowthCount) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.MediumFileDayGrowthCount = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetMediumFileRatio(v *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileRatio) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.MediumFileRatio = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetSmallFileCount(v *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileCount) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.SmallFileCount = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetSmallFileCountDayGrowthRatio(v *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileCountDayGrowthRatio) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.SmallFileCountDayGrowthRatio = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetSmallFileDayGrowthCount(v *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileDayGrowthCount) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.SmallFileDayGrowthCount = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetSmallFileRatio(v *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileRatio) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.SmallFileRatio = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetTinyFileCount(v *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileCount) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.TinyFileCount = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetTinyFileCountDayGrowthRatio(v *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileCountDayGrowthRatio) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.TinyFileCountDayGrowthRatio = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetTinyFileDayGrowthCount(v *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileDayGrowthCount) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.TinyFileDayGrowthCount = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetTinyFileRatio(v *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileRatio) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.TinyFileRatio = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetTotalDataDayGrowthSize(v *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataDayGrowthSize) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.TotalDataDayGrowthSize = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetTotalDataSize(v *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataSize) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.TotalDataSize = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetTotalDataSizeDayGrowthRatio(v *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.TotalDataSizeDayGrowthRatio = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetTotalFileCount(v *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileCount) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.TotalFileCount = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetTotalFileCountDayGrowthRatio(v *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileCountDayGrowthRatio) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.TotalFileCountDayGrowthRatio = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetTotalFileDayGrowthCount(v *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileDayGrowthCount) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.TotalFileDayGrowthCount = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetWarmDataDayGrowthSize(v *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataDayGrowthSize) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.WarmDataDayGrowthSize = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetWarmDataRatio(v *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataRatio) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.WarmDataRatio = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetWarmDataSize(v *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataSize) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.WarmDataSize = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) SetWarmDataSizeDayGrowthRatio(v *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) *GetDoctorHDFSClusterResponseBodyDataMetrics {
	s.WarmDataSizeDayGrowthRatio = v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetrics) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsColdDataDayGrowthSize struct {
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
	// -182636577752
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsColdDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsColdDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataDayGrowthSize) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataDayGrowthSize) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataDayGrowthSize) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataDayGrowthSize) SetValue(v int64) *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsColdDataRatio struct {
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
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 0.01
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsColdDataRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsColdDataRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataRatio) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataRatio) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataRatio) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataRatio) SetValue(v float32) *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsColdDataSize struct {
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
	// 5570958082267
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsColdDataSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsColdDataSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataSize) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataSize) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataSize) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataSize) SetValue(v int64) *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsColdDataSizeDayGrowthRatio struct {
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
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// -0.03
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsColdDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataSizeDayGrowthRatio) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataSizeDayGrowthRatio) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataSizeDayGrowthRatio) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataSizeDayGrowthRatio) SetValue(v float32) *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsColdDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileCount struct {
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
	// 15595897
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileCount) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileCount) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileCount) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileCount) SetValue(v int64) *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileCountDayGrowthRatio struct {
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
	// 0.005
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) SetValue(v float32) *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileDayGrowthCount struct {
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
	// 114
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileDayGrowthCount) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileDayGrowthCount) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileDayGrowthCount) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileDayGrowthCount) SetValue(v int64) *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileRatio struct {
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
	// 0.3
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileRatio) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileRatio) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileRatio) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileRatio) SetValue(v float32) *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsEmptyFileRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataDayGrowthSize struct {
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
	// -167683929450
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataDayGrowthSize) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataDayGrowthSize) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataDayGrowthSize) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataDayGrowthSize) SetValue(v int64) *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataRatio struct {
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
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 0.12
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataRatio) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataRatio) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataRatio) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataRatio) SetValue(v float32) *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataSize struct {
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
	// 1231312431
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataSize) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataSize) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataSize) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataSize) SetValue(v int64) *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio struct {
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
	// -0.09
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) SetValue(v float32) *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsHotDataDayGrowthSize struct {
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
	// 123154
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsHotDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsHotDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataDayGrowthSize) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataDayGrowthSize) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataDayGrowthSize) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataDayGrowthSize) SetValue(v int64) *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsHotDataRatio struct {
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
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 0.22
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsHotDataRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsHotDataRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataRatio) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataRatio) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataRatio) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataRatio) SetValue(v float32) *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsHotDataSize struct {
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
	// 6701531944206
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsHotDataSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsHotDataSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataSize) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataSize) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataSize) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataSize) SetValue(v int64) *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsHotDataSizeDayGrowthRatio struct {
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
	// 0.1114
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsHotDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataSizeDayGrowthRatio) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataSizeDayGrowthRatio) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataSizeDayGrowthRatio) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataSizeDayGrowthRatio) SetValue(v float32) *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsHotDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileCount struct {
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
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 123
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileCount) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileCount) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileCount) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileCount) SetValue(v int64) *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileCountDayGrowthRatio struct {
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
	// 0.39
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileCountDayGrowthRatio) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileCountDayGrowthRatio) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileCountDayGrowthRatio) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileCountDayGrowthRatio) SetValue(v float32) *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileDayGrowthCount struct {
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
	// “”
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 2
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileDayGrowthCount) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileDayGrowthCount) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileDayGrowthCount) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileDayGrowthCount) SetValue(v int64) *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileRatio struct {
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
	// 0.22
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileRatio) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileRatio) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileRatio) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileRatio) SetValue(v float32) *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsLargeFileRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileCount struct {
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
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 234
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileCount) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileCount) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileCount) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileCount) SetValue(v int64) *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileCountDayGrowthRatio struct {
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
	// 0.19
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileCountDayGrowthRatio) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileCountDayGrowthRatio) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileCountDayGrowthRatio) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileCountDayGrowthRatio) SetValue(v float32) *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileDayGrowthCount struct {
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
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 176
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileDayGrowthCount) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileDayGrowthCount) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileDayGrowthCount) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileDayGrowthCount) SetValue(v int64) *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileRatio struct {
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
	// 0.21
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileRatio) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileRatio) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileRatio) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileRatio) SetValue(v float32) *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsMediumFileRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileCount struct {
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
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 12345
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileCount) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileCount) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileCount) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileCount) SetValue(v int64) *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileCountDayGrowthRatio struct {
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
	// 0.02
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileCountDayGrowthRatio) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileCountDayGrowthRatio) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileCountDayGrowthRatio) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileCountDayGrowthRatio) SetValue(v float32) *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileDayGrowthCount struct {
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
	// 12321
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileDayGrowthCount) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileDayGrowthCount) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileDayGrowthCount) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileDayGrowthCount) SetValue(v int64) *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileRatio struct {
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
	// 0.19
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileRatio) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileRatio) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileRatio) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileRatio) SetValue(v float32) *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsSmallFileRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileCount struct {
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
	// 232131
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileCount) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileCount) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileCount) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileCount) SetValue(v int64) *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileCountDayGrowthRatio struct {
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
	// 0.003
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileCountDayGrowthRatio) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileCountDayGrowthRatio) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileCountDayGrowthRatio) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileCountDayGrowthRatio) SetValue(v float32) *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileDayGrowthCount struct {
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
	// -123
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileDayGrowthCount) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileDayGrowthCount) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileDayGrowthCount) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileDayGrowthCount) SetValue(v int64) *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileRatio struct {
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
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 0.19
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileRatio) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileRatio) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileRatio) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileRatio) SetValue(v float32) *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTinyFileRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataDayGrowthSize struct {
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
	// 256482228248
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataDayGrowthSize) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataDayGrowthSize) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataDayGrowthSize) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataDayGrowthSize) SetValue(v int64) *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataSize struct {
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
	// 62086342083623
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataSize) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataSize) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataSize) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataSize) SetValue(v int64) *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataSizeDayGrowthRatio struct {
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
	// 0.14
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) SetValue(v float32) *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileCount struct {
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
	// 51683279
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileCount) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileCount) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileCount) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileCount) SetValue(v int64) *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileCountDayGrowthRatio struct {
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
	// 0.02
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileCountDayGrowthRatio) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileCountDayGrowthRatio) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileCountDayGrowthRatio) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileCountDayGrowthRatio) SetValue(v float32) *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileDayGrowthCount struct {
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
	// “”
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 27809
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileDayGrowthCount) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileDayGrowthCount) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileDayGrowthCount) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileDayGrowthCount) SetValue(v int64) *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsTotalFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataDayGrowthSize struct {
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
	// -64806998319
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataDayGrowthSize) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataDayGrowthSize) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataDayGrowthSize) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataDayGrowthSize) SetValue(v int64) *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataRatio struct {
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
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 0.12
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataRatio) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataRatio) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataRatio) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataRatio) SetValue(v float32) *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataSize struct {
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
	// 4062349775577
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataSize) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataSize) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataSize) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataSize) SetValue(v int64) *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataSizeDayGrowthRatio struct {
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
	// -0.015
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) SetDescription(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) SetName(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) SetUnit(v string) *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) SetValue(v float32) *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSClusterResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}
