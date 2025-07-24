// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorHDFSDirectoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDoctorHDFSDirectoryResponseBodyData) *GetDoctorHDFSDirectoryResponseBody
	GetData() *GetDoctorHDFSDirectoryResponseBodyData
	SetRequestId(v string) *GetDoctorHDFSDirectoryResponseBody
	GetRequestId() *string
}

type GetDoctorHDFSDirectoryResponseBody struct {
	// The analysis results of the HDFS directory.
	Data *GetDoctorHDFSDirectoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDoctorHDFSDirectoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSDirectoryResponseBody) GetData() *GetDoctorHDFSDirectoryResponseBodyData {
	return s.Data
}

func (s *GetDoctorHDFSDirectoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDoctorHDFSDirectoryResponseBody) SetData(v *GetDoctorHDFSDirectoryResponseBodyData) *GetDoctorHDFSDirectoryResponseBody {
	s.Data = v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBody) SetRequestId(v string) *GetDoctorHDFSDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSDirectoryResponseBodyData struct {
	// The directory level.
	//
	// example:
	//
	// 2
	Depth *int32 `json:"Depth,omitempty" xml:"Depth,omitempty"`
	// The group to which the directory belongs.
	//
	// example:
	//
	// DW
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The metric information.
	Metrics *GetDoctorHDFSDirectoryResponseBodyDataMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Struct"`
	// The directory owner.
	//
	// example:
	//
	// DW
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s GetDoctorHDFSDirectoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSDirectoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSDirectoryResponseBodyData) GetDepth() *int32 {
	return s.Depth
}

func (s *GetDoctorHDFSDirectoryResponseBodyData) GetGroup() *string {
	return s.Group
}

func (s *GetDoctorHDFSDirectoryResponseBodyData) GetMetrics() *GetDoctorHDFSDirectoryResponseBodyDataMetrics {
	return s.Metrics
}

func (s *GetDoctorHDFSDirectoryResponseBodyData) GetUser() *string {
	return s.User
}

func (s *GetDoctorHDFSDirectoryResponseBodyData) SetDepth(v int32) *GetDoctorHDFSDirectoryResponseBodyData {
	s.Depth = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyData) SetGroup(v string) *GetDoctorHDFSDirectoryResponseBodyData {
	s.Group = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyData) SetMetrics(v *GetDoctorHDFSDirectoryResponseBodyDataMetrics) *GetDoctorHDFSDirectoryResponseBodyData {
	s.Metrics = v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyData) SetUser(v string) *GetDoctorHDFSDirectoryResponseBodyData {
	s.User = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSDirectoryResponseBodyDataMetrics struct {
	// The daily increment of the amount of cold data. Cold data refers to data that is not accessed for more than 30 days but is accessed in previous 90 days.
	ColdDataDayGrowthSize *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataDayGrowthSize `json:"ColdDataDayGrowthSize,omitempty" xml:"ColdDataDayGrowthSize,omitempty" type:"Struct"`
	// The amount of cold data. Cold data refers to data that is not accessed for more than 30 days but is accessed in previous 90 days.
	ColdDataSize *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataSize `json:"ColdDataSize,omitempty" xml:"ColdDataSize,omitempty" type:"Struct"`
	// The day-to-day growth rate of the amount of cold data. Cold data refers to data that is not accessed for more than 30 days but is accessed in previous 90 days.
	ColdDataSizeDayGrowthRatio *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataSizeDayGrowthRatio `json:"ColdDataSizeDayGrowthRatio,omitempty" xml:"ColdDataSizeDayGrowthRatio,omitempty" type:"Struct"`
	// The number of empty files. Empty files are those with a size of 0 MB.
	EmptyFileCount *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileCount `json:"EmptyFileCount,omitempty" xml:"EmptyFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the number of empty files. Empty files are those with a size of 0 MB.
	EmptyFileCountDayGrowthRatio *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileCountDayGrowthRatio `json:"EmptyFileCountDayGrowthRatio,omitempty" xml:"EmptyFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the number of empty files. Empty files are those with a size of 0 MB.
	EmptyFileDayGrowthCount *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileDayGrowthCount `json:"EmptyFileDayGrowthCount,omitempty" xml:"EmptyFileDayGrowthCount,omitempty" type:"Struct"`
	// The daily increment of the amount of very cold data. Very cold data refers to data that is not accessed for more than 90 days.
	FreezeDataDayGrowthSize *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataDayGrowthSize `json:"FreezeDataDayGrowthSize,omitempty" xml:"FreezeDataDayGrowthSize,omitempty" type:"Struct"`
	// The amount of very cold data. Very cold data refers to data that is not accessed for more than 90 days.
	FreezeDataSize *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataSize `json:"FreezeDataSize,omitempty" xml:"FreezeDataSize,omitempty" type:"Struct"`
	// The day-to-day growth rate of the amount of very cold data. Very cold data refers to data that is not accessed for more than 90 days.
	FreezeDataSizeDayGrowthRatio *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio `json:"FreezeDataSizeDayGrowthRatio,omitempty" xml:"FreezeDataSizeDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the amount of hot data. Hot data refers to data that is accessed in previous seven days.
	HotDataDayGrowthSize *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataDayGrowthSize `json:"HotDataDayGrowthSize,omitempty" xml:"HotDataDayGrowthSize,omitempty" type:"Struct"`
	// The amount of hot data. Hot data refers to data that is accessed in previous seven days.
	HotDataSize *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataSize `json:"HotDataSize,omitempty" xml:"HotDataSize,omitempty" type:"Struct"`
	// The day-to-day growth rate of the amount of hot data. Hot data refers to data that is accessed in previous seven days.
	HotDataSizeDayGrowthRatio *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataSizeDayGrowthRatio `json:"HotDataSizeDayGrowthRatio,omitempty" xml:"HotDataSizeDayGrowthRatio,omitempty" type:"Struct"`
	// The number of large files. Large files are those with a size greater than 1 GB.
	LargeFileCount *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileCount `json:"LargeFileCount,omitempty" xml:"LargeFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the number of large files. Large files are those with a size greater than 1 GB.
	LargeFileCountDayGrowthRatio *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileCountDayGrowthRatio `json:"LargeFileCountDayGrowthRatio,omitempty" xml:"LargeFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the number of large files. Large files are those with a size greater than 1 GB.
	LargeFileDayGrowthCount *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileDayGrowthCount `json:"LargeFileDayGrowthCount,omitempty" xml:"LargeFileDayGrowthCount,omitempty" type:"Struct"`
	// The number of medium files. Medium files are those with a size greater than or equal to 128 MB and less than or equal to 1 GB.
	MediumFileCount *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileCount `json:"MediumFileCount,omitempty" xml:"MediumFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the number of medium files. Medium files are those with a size greater than or equal to 128 MB and less than or equal to 1 GB.
	MediumFileCountDayGrowthRatio *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileCountDayGrowthRatio `json:"MediumFileCountDayGrowthRatio,omitempty" xml:"MediumFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the number of medium files. Medium files are those with a size greater than or equal to 128 MB and less than or equal to 1 GB.
	MediumFileDayGrowthCount *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileDayGrowthCount `json:"MediumFileDayGrowthCount,omitempty" xml:"MediumFileDayGrowthCount,omitempty" type:"Struct"`
	// The number of small files. Small files are those with a size greater than or equal to 10 MB and less than 128 MB.
	SmallFileCount *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileCount `json:"SmallFileCount,omitempty" xml:"SmallFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the number of small files. Small files are those with a size greater than or equal to 10 MB and less than 128 MB.
	SmallFileCountDayGrowthRatio *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileCountDayGrowthRatio `json:"SmallFileCountDayGrowthRatio,omitempty" xml:"SmallFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the number of small files. Small files are those with a size greater than or equal to 10 MB and less than 128 MB.
	SmallFileDayGrowthCount *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileDayGrowthCount `json:"SmallFileDayGrowthCount,omitempty" xml:"SmallFileDayGrowthCount,omitempty" type:"Struct"`
	// The number of very small files. Very small files are those with a size greater than 0 MB and less than 10 MB.
	TinyFileCount *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileCount `json:"TinyFileCount,omitempty" xml:"TinyFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the number of very small files. Very small files are those with a size greater than 0 MB and less than 10 MB.
	TinyFileCountDayGrowthRatio *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileCountDayGrowthRatio `json:"TinyFileCountDayGrowthRatio,omitempty" xml:"TinyFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the number of very small files. Very small files are those with a size greater than 0 MB and less than 10 MB.
	TinyFileDayGrowthCount *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileDayGrowthCount `json:"TinyFileDayGrowthCount,omitempty" xml:"TinyFileDayGrowthCount,omitempty" type:"Struct"`
	// The daily incremental of the total amount of data.
	TotalDataDayGrowthSize *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataDayGrowthSize `json:"TotalDataDayGrowthSize,omitempty" xml:"TotalDataDayGrowthSize,omitempty" type:"Struct"`
	// The total amount of data.
	TotalDataSize *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataSize `json:"TotalDataSize,omitempty" xml:"TotalDataSize,omitempty" type:"Struct"`
	// The day-to-day growth rate of the total data volume.
	TotalDataSizeDayGrowthRatio *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataSizeDayGrowthRatio `json:"TotalDataSizeDayGrowthRatio,omitempty" xml:"TotalDataSizeDayGrowthRatio,omitempty" type:"Struct"`
	// The total number of files.
	TotalFileCount *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileCount `json:"TotalFileCount,omitempty" xml:"TotalFileCount,omitempty" type:"Struct"`
	// The day-to-day growth rate of the total number of files.
	TotalFileCountDayGrowthRatio *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileCountDayGrowthRatio `json:"TotalFileCountDayGrowthRatio,omitempty" xml:"TotalFileCountDayGrowthRatio,omitempty" type:"Struct"`
	// The daily increment of the total number of files.
	TotalFileDayGrowthCount *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileDayGrowthCount `json:"TotalFileDayGrowthCount,omitempty" xml:"TotalFileDayGrowthCount,omitempty" type:"Struct"`
	// The daily increment of the amount of warm data. Warm data refers to data that is not accessed for more than 7 days but is accessed in previous 30 days.
	WarmDataDayGrowthSize *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataDayGrowthSize `json:"WarmDataDayGrowthSize,omitempty" xml:"WarmDataDayGrowthSize,omitempty" type:"Struct"`
	// The amount of warm data. Warm data refers to data that is not accessed for more than 7 days but is accessed in previous 30 days.
	WarmDataSize *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataSize `json:"WarmDataSize,omitempty" xml:"WarmDataSize,omitempty" type:"Struct"`
	// The day-to-day growth rate of the amount of warm data. Warm data refers to data that is not accessed for more than 7 days but is accessed in previous 30 days.
	WarmDataSizeDayGrowthRatio *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataSizeDayGrowthRatio `json:"WarmDataSizeDayGrowthRatio,omitempty" xml:"WarmDataSizeDayGrowthRatio,omitempty" type:"Struct"`
}

func (s GetDoctorHDFSDirectoryResponseBodyDataMetrics) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSDirectoryResponseBodyDataMetrics) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) GetColdDataDayGrowthSize() *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataDayGrowthSize {
	return s.ColdDataDayGrowthSize
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) GetColdDataSize() *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataSize {
	return s.ColdDataSize
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) GetColdDataSizeDayGrowthRatio() *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	return s.ColdDataSizeDayGrowthRatio
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) GetEmptyFileCount() *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileCount {
	return s.EmptyFileCount
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) GetEmptyFileCountDayGrowthRatio() *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	return s.EmptyFileCountDayGrowthRatio
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) GetEmptyFileDayGrowthCount() *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileDayGrowthCount {
	return s.EmptyFileDayGrowthCount
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) GetFreezeDataDayGrowthSize() *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataDayGrowthSize {
	return s.FreezeDataDayGrowthSize
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) GetFreezeDataSize() *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataSize {
	return s.FreezeDataSize
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) GetFreezeDataSizeDayGrowthRatio() *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	return s.FreezeDataSizeDayGrowthRatio
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) GetHotDataDayGrowthSize() *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataDayGrowthSize {
	return s.HotDataDayGrowthSize
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) GetHotDataSize() *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataSize {
	return s.HotDataSize
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) GetHotDataSizeDayGrowthRatio() *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	return s.HotDataSizeDayGrowthRatio
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) GetLargeFileCount() *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileCount {
	return s.LargeFileCount
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) GetLargeFileCountDayGrowthRatio() *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	return s.LargeFileCountDayGrowthRatio
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) GetLargeFileDayGrowthCount() *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileDayGrowthCount {
	return s.LargeFileDayGrowthCount
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) GetMediumFileCount() *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileCount {
	return s.MediumFileCount
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) GetMediumFileCountDayGrowthRatio() *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	return s.MediumFileCountDayGrowthRatio
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) GetMediumFileDayGrowthCount() *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileDayGrowthCount {
	return s.MediumFileDayGrowthCount
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) GetSmallFileCount() *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileCount {
	return s.SmallFileCount
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) GetSmallFileCountDayGrowthRatio() *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	return s.SmallFileCountDayGrowthRatio
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) GetSmallFileDayGrowthCount() *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileDayGrowthCount {
	return s.SmallFileDayGrowthCount
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) GetTinyFileCount() *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileCount {
	return s.TinyFileCount
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) GetTinyFileCountDayGrowthRatio() *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	return s.TinyFileCountDayGrowthRatio
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) GetTinyFileDayGrowthCount() *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileDayGrowthCount {
	return s.TinyFileDayGrowthCount
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) GetTotalDataDayGrowthSize() *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataDayGrowthSize {
	return s.TotalDataDayGrowthSize
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) GetTotalDataSize() *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataSize {
	return s.TotalDataSize
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) GetTotalDataSizeDayGrowthRatio() *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	return s.TotalDataSizeDayGrowthRatio
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) GetTotalFileCount() *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileCount {
	return s.TotalFileCount
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) GetTotalFileCountDayGrowthRatio() *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	return s.TotalFileCountDayGrowthRatio
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) GetTotalFileDayGrowthCount() *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileDayGrowthCount {
	return s.TotalFileDayGrowthCount
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) GetWarmDataDayGrowthSize() *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataDayGrowthSize {
	return s.WarmDataDayGrowthSize
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) GetWarmDataSize() *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataSize {
	return s.WarmDataSize
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) GetWarmDataSizeDayGrowthRatio() *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	return s.WarmDataSizeDayGrowthRatio
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) SetColdDataDayGrowthSize(v *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataDayGrowthSize) *GetDoctorHDFSDirectoryResponseBodyDataMetrics {
	s.ColdDataDayGrowthSize = v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) SetColdDataSize(v *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataSize) *GetDoctorHDFSDirectoryResponseBodyDataMetrics {
	s.ColdDataSize = v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) SetColdDataSizeDayGrowthRatio(v *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataSizeDayGrowthRatio) *GetDoctorHDFSDirectoryResponseBodyDataMetrics {
	s.ColdDataSizeDayGrowthRatio = v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) SetEmptyFileCount(v *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileCount) *GetDoctorHDFSDirectoryResponseBodyDataMetrics {
	s.EmptyFileCount = v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) SetEmptyFileCountDayGrowthRatio(v *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) *GetDoctorHDFSDirectoryResponseBodyDataMetrics {
	s.EmptyFileCountDayGrowthRatio = v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) SetEmptyFileDayGrowthCount(v *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileDayGrowthCount) *GetDoctorHDFSDirectoryResponseBodyDataMetrics {
	s.EmptyFileDayGrowthCount = v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) SetFreezeDataDayGrowthSize(v *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataDayGrowthSize) *GetDoctorHDFSDirectoryResponseBodyDataMetrics {
	s.FreezeDataDayGrowthSize = v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) SetFreezeDataSize(v *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataSize) *GetDoctorHDFSDirectoryResponseBodyDataMetrics {
	s.FreezeDataSize = v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) SetFreezeDataSizeDayGrowthRatio(v *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) *GetDoctorHDFSDirectoryResponseBodyDataMetrics {
	s.FreezeDataSizeDayGrowthRatio = v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) SetHotDataDayGrowthSize(v *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataDayGrowthSize) *GetDoctorHDFSDirectoryResponseBodyDataMetrics {
	s.HotDataDayGrowthSize = v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) SetHotDataSize(v *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataSize) *GetDoctorHDFSDirectoryResponseBodyDataMetrics {
	s.HotDataSize = v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) SetHotDataSizeDayGrowthRatio(v *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataSizeDayGrowthRatio) *GetDoctorHDFSDirectoryResponseBodyDataMetrics {
	s.HotDataSizeDayGrowthRatio = v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) SetLargeFileCount(v *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileCount) *GetDoctorHDFSDirectoryResponseBodyDataMetrics {
	s.LargeFileCount = v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) SetLargeFileCountDayGrowthRatio(v *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileCountDayGrowthRatio) *GetDoctorHDFSDirectoryResponseBodyDataMetrics {
	s.LargeFileCountDayGrowthRatio = v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) SetLargeFileDayGrowthCount(v *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileDayGrowthCount) *GetDoctorHDFSDirectoryResponseBodyDataMetrics {
	s.LargeFileDayGrowthCount = v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) SetMediumFileCount(v *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileCount) *GetDoctorHDFSDirectoryResponseBodyDataMetrics {
	s.MediumFileCount = v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) SetMediumFileCountDayGrowthRatio(v *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileCountDayGrowthRatio) *GetDoctorHDFSDirectoryResponseBodyDataMetrics {
	s.MediumFileCountDayGrowthRatio = v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) SetMediumFileDayGrowthCount(v *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileDayGrowthCount) *GetDoctorHDFSDirectoryResponseBodyDataMetrics {
	s.MediumFileDayGrowthCount = v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) SetSmallFileCount(v *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileCount) *GetDoctorHDFSDirectoryResponseBodyDataMetrics {
	s.SmallFileCount = v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) SetSmallFileCountDayGrowthRatio(v *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileCountDayGrowthRatio) *GetDoctorHDFSDirectoryResponseBodyDataMetrics {
	s.SmallFileCountDayGrowthRatio = v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) SetSmallFileDayGrowthCount(v *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileDayGrowthCount) *GetDoctorHDFSDirectoryResponseBodyDataMetrics {
	s.SmallFileDayGrowthCount = v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) SetTinyFileCount(v *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileCount) *GetDoctorHDFSDirectoryResponseBodyDataMetrics {
	s.TinyFileCount = v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) SetTinyFileCountDayGrowthRatio(v *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileCountDayGrowthRatio) *GetDoctorHDFSDirectoryResponseBodyDataMetrics {
	s.TinyFileCountDayGrowthRatio = v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) SetTinyFileDayGrowthCount(v *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileDayGrowthCount) *GetDoctorHDFSDirectoryResponseBodyDataMetrics {
	s.TinyFileDayGrowthCount = v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) SetTotalDataDayGrowthSize(v *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataDayGrowthSize) *GetDoctorHDFSDirectoryResponseBodyDataMetrics {
	s.TotalDataDayGrowthSize = v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) SetTotalDataSize(v *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataSize) *GetDoctorHDFSDirectoryResponseBodyDataMetrics {
	s.TotalDataSize = v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) SetTotalDataSizeDayGrowthRatio(v *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) *GetDoctorHDFSDirectoryResponseBodyDataMetrics {
	s.TotalDataSizeDayGrowthRatio = v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) SetTotalFileCount(v *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileCount) *GetDoctorHDFSDirectoryResponseBodyDataMetrics {
	s.TotalFileCount = v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) SetTotalFileCountDayGrowthRatio(v *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileCountDayGrowthRatio) *GetDoctorHDFSDirectoryResponseBodyDataMetrics {
	s.TotalFileCountDayGrowthRatio = v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) SetTotalFileDayGrowthCount(v *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileDayGrowthCount) *GetDoctorHDFSDirectoryResponseBodyDataMetrics {
	s.TotalFileDayGrowthCount = v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) SetWarmDataDayGrowthSize(v *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataDayGrowthSize) *GetDoctorHDFSDirectoryResponseBodyDataMetrics {
	s.WarmDataDayGrowthSize = v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) SetWarmDataSize(v *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataSize) *GetDoctorHDFSDirectoryResponseBodyDataMetrics {
	s.WarmDataSize = v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) SetWarmDataSizeDayGrowthRatio(v *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) *GetDoctorHDFSDirectoryResponseBodyDataMetrics {
	s.WarmDataSizeDayGrowthRatio = v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetrics) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataDayGrowthSize struct {
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

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataDayGrowthSize) SetDescription(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataDayGrowthSize) SetName(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataDayGrowthSize) SetUnit(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataDayGrowthSize) SetValue(v int64) *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataSize struct {
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

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataSize) SetDescription(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataSize) SetName(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataSize) SetUnit(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataSize) SetValue(v int64) *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataSizeDayGrowthRatio struct {
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

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataSizeDayGrowthRatio) SetDescription(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataSizeDayGrowthRatio) SetName(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataSizeDayGrowthRatio) SetUnit(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataSizeDayGrowthRatio) SetValue(v float32) *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsColdDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileCount struct {
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

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileCount) SetDescription(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileCount) SetName(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileCount) SetUnit(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileCount) SetValue(v int64) *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileCountDayGrowthRatio struct {
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

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) SetDescription(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) SetName(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) SetUnit(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) SetValue(v float32) *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileDayGrowthCount struct {
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

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileDayGrowthCount) SetDescription(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileDayGrowthCount) SetName(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileDayGrowthCount) SetUnit(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileDayGrowthCount) SetValue(v int64) *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsEmptyFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataDayGrowthSize struct {
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

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataDayGrowthSize) SetDescription(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataDayGrowthSize) SetName(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataDayGrowthSize) SetUnit(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataDayGrowthSize) SetValue(v int64) *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataSize struct {
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

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataSize) SetDescription(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataSize) SetName(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataSize) SetUnit(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataSize) SetValue(v int64) *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio struct {
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

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) SetDescription(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) SetName(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) SetUnit(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) SetValue(v float32) *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataDayGrowthSize struct {
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

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataDayGrowthSize) SetDescription(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataDayGrowthSize) SetName(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataDayGrowthSize) SetUnit(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataDayGrowthSize) SetValue(v int64) *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataSize struct {
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

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataSize) SetDescription(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataSize) SetName(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataSize) SetUnit(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataSize) SetValue(v int64) *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataSizeDayGrowthRatio struct {
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

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataSizeDayGrowthRatio) SetDescription(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataSizeDayGrowthRatio) SetName(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataSizeDayGrowthRatio) SetUnit(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataSizeDayGrowthRatio) SetValue(v float32) *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsHotDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileCount struct {
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

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileCount) SetDescription(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileCount) SetName(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileCount) SetUnit(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileCount) SetValue(v int64) *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileCountDayGrowthRatio struct {
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

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileCountDayGrowthRatio) SetDescription(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileCountDayGrowthRatio) SetName(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileCountDayGrowthRatio) SetUnit(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileCountDayGrowthRatio) SetValue(v float32) *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileDayGrowthCount struct {
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
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 2
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileDayGrowthCount) SetDescription(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileDayGrowthCount) SetName(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileDayGrowthCount) SetUnit(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileDayGrowthCount) SetValue(v int64) *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsLargeFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileCount struct {
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

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileCount) SetDescription(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileCount) SetName(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileCount) SetUnit(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileCount) SetValue(v int64) *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileCountDayGrowthRatio struct {
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

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileCountDayGrowthRatio) SetDescription(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileCountDayGrowthRatio) SetName(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileCountDayGrowthRatio) SetUnit(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileCountDayGrowthRatio) SetValue(v float32) *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileDayGrowthCount struct {
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

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileDayGrowthCount) SetDescription(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileDayGrowthCount) SetName(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileDayGrowthCount) SetUnit(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileDayGrowthCount) SetValue(v int64) *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsMediumFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileCount struct {
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

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileCount) SetDescription(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileCount) SetName(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileCount) SetUnit(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileCount) SetValue(v int64) *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileCountDayGrowthRatio struct {
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

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileCountDayGrowthRatio) SetDescription(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileCountDayGrowthRatio) SetName(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileCountDayGrowthRatio) SetUnit(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileCountDayGrowthRatio) SetValue(v float32) *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileDayGrowthCount struct {
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
	// 12345
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileDayGrowthCount) SetDescription(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileDayGrowthCount) SetName(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileDayGrowthCount) SetUnit(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileDayGrowthCount) SetValue(v int64) *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsSmallFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileCount struct {
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

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileCount) SetDescription(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileCount) SetName(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileCount) SetUnit(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileCount) SetValue(v int64) *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileCountDayGrowthRatio struct {
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

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileCountDayGrowthRatio) SetDescription(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileCountDayGrowthRatio) SetName(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileCountDayGrowthRatio) SetUnit(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileCountDayGrowthRatio) SetValue(v float32) *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileDayGrowthCount struct {
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

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileDayGrowthCount) SetDescription(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileDayGrowthCount) SetName(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileDayGrowthCount) SetUnit(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileDayGrowthCount) SetValue(v int64) *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTinyFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataDayGrowthSize struct {
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

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataDayGrowthSize) SetDescription(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataDayGrowthSize) SetName(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataDayGrowthSize) SetUnit(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataDayGrowthSize) SetValue(v int64) *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataSize struct {
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

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataSize) SetDescription(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataSize) SetName(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataSize) SetUnit(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataSize) SetValue(v int64) *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataSizeDayGrowthRatio struct {
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

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) SetDescription(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) SetName(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) SetUnit(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) SetValue(v float32) *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileCount struct {
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

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileCount) SetDescription(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileCount) SetName(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileCount) SetUnit(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileCount) SetValue(v int64) *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileCountDayGrowthRatio struct {
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

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileCountDayGrowthRatio) SetDescription(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileCountDayGrowthRatio) SetName(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileCountDayGrowthRatio) SetUnit(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileCountDayGrowthRatio) SetValue(v float32) *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileDayGrowthCount struct {
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
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 27809
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileDayGrowthCount) SetDescription(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileDayGrowthCount) SetName(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileDayGrowthCount) SetUnit(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileDayGrowthCount) SetValue(v int64) *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsTotalFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataDayGrowthSize struct {
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

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataDayGrowthSize) SetDescription(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataDayGrowthSize) SetName(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataDayGrowthSize) SetUnit(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataDayGrowthSize) SetValue(v int64) *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataSize struct {
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

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataSize) SetDescription(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataSize) SetName(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataSize) SetUnit(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataSize) SetValue(v int64) *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataSizeDayGrowthRatio struct {
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

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) SetDescription(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) SetName(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) SetUnit(v string) *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) SetValue(v float32) *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}
