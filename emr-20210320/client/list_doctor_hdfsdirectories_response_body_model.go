// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDoctorHDFSDirectoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListDoctorHDFSDirectoriesResponseBodyData) *ListDoctorHDFSDirectoriesResponseBody
	GetData() []*ListDoctorHDFSDirectoriesResponseBodyData
	SetMaxResults(v int32) *ListDoctorHDFSDirectoriesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDoctorHDFSDirectoriesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDoctorHDFSDirectoriesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListDoctorHDFSDirectoriesResponseBody
	GetTotalCount() *int32
}

type ListDoctorHDFSDirectoriesResponseBody struct {
	Data []*ListDoctorHDFSDirectoriesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// 本次请求所返回的最大记录条数。
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 返回读取到的数据位置，空代表数据已经读取完毕。
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 请求ID。
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 本次请求条件下的数据总量。
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDoctorHDFSDirectoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSDirectoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSDirectoriesResponseBody) GetData() []*ListDoctorHDFSDirectoriesResponseBodyData {
	return s.Data
}

func (s *ListDoctorHDFSDirectoriesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDoctorHDFSDirectoriesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDoctorHDFSDirectoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDoctorHDFSDirectoriesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDoctorHDFSDirectoriesResponseBody) SetData(v []*ListDoctorHDFSDirectoriesResponseBodyData) *ListDoctorHDFSDirectoriesResponseBody {
	s.Data = v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBody) SetMaxResults(v int32) *ListDoctorHDFSDirectoriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBody) SetNextToken(v string) *ListDoctorHDFSDirectoriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBody) SetRequestId(v string) *ListDoctorHDFSDirectoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBody) SetTotalCount(v int32) *ListDoctorHDFSDirectoriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSDirectoriesResponseBodyData struct {
	// example:
	//
	// 2
	Depth *int32 `json:"Depth,omitempty" xml:"Depth,omitempty"`
	// example:
	//
	// /tmp/test
	DirPath *string `json:"DirPath,omitempty" xml:"DirPath,omitempty"`
	// example:
	//
	// DW
	Group   *string                                           `json:"Group,omitempty" xml:"Group,omitempty"`
	Metrics *ListDoctorHDFSDirectoriesResponseBodyDataMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Struct"`
	// example:
	//
	// DW
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s ListDoctorHDFSDirectoriesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSDirectoriesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSDirectoriesResponseBodyData) GetDepth() *int32 {
	return s.Depth
}

func (s *ListDoctorHDFSDirectoriesResponseBodyData) GetDirPath() *string {
	return s.DirPath
}

func (s *ListDoctorHDFSDirectoriesResponseBodyData) GetGroup() *string {
	return s.Group
}

func (s *ListDoctorHDFSDirectoriesResponseBodyData) GetMetrics() *ListDoctorHDFSDirectoriesResponseBodyDataMetrics {
	return s.Metrics
}

func (s *ListDoctorHDFSDirectoriesResponseBodyData) GetUser() *string {
	return s.User
}

func (s *ListDoctorHDFSDirectoriesResponseBodyData) SetDepth(v int32) *ListDoctorHDFSDirectoriesResponseBodyData {
	s.Depth = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyData) SetDirPath(v string) *ListDoctorHDFSDirectoriesResponseBodyData {
	s.DirPath = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyData) SetGroup(v string) *ListDoctorHDFSDirectoriesResponseBodyData {
	s.Group = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyData) SetMetrics(v *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) *ListDoctorHDFSDirectoriesResponseBodyData {
	s.Metrics = v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyData) SetUser(v string) *ListDoctorHDFSDirectoriesResponseBodyData {
	s.User = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSDirectoriesResponseBodyDataMetrics struct {
	ColdDataDayGrowthSize         *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataDayGrowthSize         `json:"ColdDataDayGrowthSize,omitempty" xml:"ColdDataDayGrowthSize,omitempty" type:"Struct"`
	ColdDataSize                  *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataSize                  `json:"ColdDataSize,omitempty" xml:"ColdDataSize,omitempty" type:"Struct"`
	ColdDataSizeDayGrowthRatio    *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataSizeDayGrowthRatio    `json:"ColdDataSizeDayGrowthRatio,omitempty" xml:"ColdDataSizeDayGrowthRatio,omitempty" type:"Struct"`
	EmptyFileCount                *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileCount                `json:"EmptyFileCount,omitempty" xml:"EmptyFileCount,omitempty" type:"Struct"`
	EmptyFileCountDayGrowthRatio  *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio  `json:"EmptyFileCountDayGrowthRatio,omitempty" xml:"EmptyFileCountDayGrowthRatio,omitempty" type:"Struct"`
	EmptyFileDayGrowthCount       *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileDayGrowthCount       `json:"EmptyFileDayGrowthCount,omitempty" xml:"EmptyFileDayGrowthCount,omitempty" type:"Struct"`
	FreezeDataDayGrowthSize       *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataDayGrowthSize       `json:"FreezeDataDayGrowthSize,omitempty" xml:"FreezeDataDayGrowthSize,omitempty" type:"Struct"`
	FreezeDataSize                *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataSize                `json:"FreezeDataSize,omitempty" xml:"FreezeDataSize,omitempty" type:"Struct"`
	FreezeDataSizeDayGrowthRatio  *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio  `json:"FreezeDataSizeDayGrowthRatio,omitempty" xml:"FreezeDataSizeDayGrowthRatio,omitempty" type:"Struct"`
	HotDataDayGrowthSize          *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataDayGrowthSize          `json:"HotDataDayGrowthSize,omitempty" xml:"HotDataDayGrowthSize,omitempty" type:"Struct"`
	HotDataSize                   *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataSize                   `json:"HotDataSize,omitempty" xml:"HotDataSize,omitempty" type:"Struct"`
	HotDataSizeDayGrowthRatio     *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataSizeDayGrowthRatio     `json:"HotDataSizeDayGrowthRatio,omitempty" xml:"HotDataSizeDayGrowthRatio,omitempty" type:"Struct"`
	LargeFileCount                *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileCount                `json:"LargeFileCount,omitempty" xml:"LargeFileCount,omitempty" type:"Struct"`
	LargeFileCountDayGrowthRatio  *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileCountDayGrowthRatio  `json:"LargeFileCountDayGrowthRatio,omitempty" xml:"LargeFileCountDayGrowthRatio,omitempty" type:"Struct"`
	LargeFileDayGrowthCount       *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileDayGrowthCount       `json:"LargeFileDayGrowthCount,omitempty" xml:"LargeFileDayGrowthCount,omitempty" type:"Struct"`
	MediumFileCount               *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileCount               `json:"MediumFileCount,omitempty" xml:"MediumFileCount,omitempty" type:"Struct"`
	MediumFileCountDayGrowthRatio *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileCountDayGrowthRatio `json:"MediumFileCountDayGrowthRatio,omitempty" xml:"MediumFileCountDayGrowthRatio,omitempty" type:"Struct"`
	MediumFileDayGrowthCount      *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileDayGrowthCount      `json:"MediumFileDayGrowthCount,omitempty" xml:"MediumFileDayGrowthCount,omitempty" type:"Struct"`
	SmallFileCount                *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileCount                `json:"SmallFileCount,omitempty" xml:"SmallFileCount,omitempty" type:"Struct"`
	SmallFileCountDayGrowthRatio  *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileCountDayGrowthRatio  `json:"SmallFileCountDayGrowthRatio,omitempty" xml:"SmallFileCountDayGrowthRatio,omitempty" type:"Struct"`
	SmallFileDayGrowthCount       *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileDayGrowthCount       `json:"SmallFileDayGrowthCount,omitempty" xml:"SmallFileDayGrowthCount,omitempty" type:"Struct"`
	TinyFileCount                 *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileCount                 `json:"TinyFileCount,omitempty" xml:"TinyFileCount,omitempty" type:"Struct"`
	TinyFileCountDayGrowthRatio   *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileCountDayGrowthRatio   `json:"TinyFileCountDayGrowthRatio,omitempty" xml:"TinyFileCountDayGrowthRatio,omitempty" type:"Struct"`
	TinyFileDayGrowthCount        *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileDayGrowthCount        `json:"TinyFileDayGrowthCount,omitempty" xml:"TinyFileDayGrowthCount,omitempty" type:"Struct"`
	TotalDataDayGrowthSize        *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataDayGrowthSize        `json:"TotalDataDayGrowthSize,omitempty" xml:"TotalDataDayGrowthSize,omitempty" type:"Struct"`
	TotalDataSize                 *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataSize                 `json:"TotalDataSize,omitempty" xml:"TotalDataSize,omitempty" type:"Struct"`
	TotalDataSizeDayGrowthRatio   *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio   `json:"TotalDataSizeDayGrowthRatio,omitempty" xml:"TotalDataSizeDayGrowthRatio,omitempty" type:"Struct"`
	TotalFileCount                *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileCount                `json:"TotalFileCount,omitempty" xml:"TotalFileCount,omitempty" type:"Struct"`
	TotalFileCountDayGrowthRatio  *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileCountDayGrowthRatio  `json:"TotalFileCountDayGrowthRatio,omitempty" xml:"TotalFileCountDayGrowthRatio,omitempty" type:"Struct"`
	TotalFileDayGrowthCount       *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileDayGrowthCount       `json:"TotalFileDayGrowthCount,omitempty" xml:"TotalFileDayGrowthCount,omitempty" type:"Struct"`
	WarmDataDayGrowthSize         *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataDayGrowthSize         `json:"WarmDataDayGrowthSize,omitempty" xml:"WarmDataDayGrowthSize,omitempty" type:"Struct"`
	WarmDataSize                  *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataSize                  `json:"WarmDataSize,omitempty" xml:"WarmDataSize,omitempty" type:"Struct"`
	WarmDataSizeDayGrowthRatio    *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio    `json:"WarmDataSizeDayGrowthRatio,omitempty" xml:"WarmDataSizeDayGrowthRatio,omitempty" type:"Struct"`
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetrics) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetrics) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) GetColdDataDayGrowthSize() *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataDayGrowthSize {
	return s.ColdDataDayGrowthSize
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) GetColdDataSize() *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataSize {
	return s.ColdDataSize
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) GetColdDataSizeDayGrowthRatio() *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	return s.ColdDataSizeDayGrowthRatio
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) GetEmptyFileCount() *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileCount {
	return s.EmptyFileCount
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) GetEmptyFileCountDayGrowthRatio() *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	return s.EmptyFileCountDayGrowthRatio
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) GetEmptyFileDayGrowthCount() *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileDayGrowthCount {
	return s.EmptyFileDayGrowthCount
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) GetFreezeDataDayGrowthSize() *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataDayGrowthSize {
	return s.FreezeDataDayGrowthSize
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) GetFreezeDataSize() *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataSize {
	return s.FreezeDataSize
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) GetFreezeDataSizeDayGrowthRatio() *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	return s.FreezeDataSizeDayGrowthRatio
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) GetHotDataDayGrowthSize() *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataDayGrowthSize {
	return s.HotDataDayGrowthSize
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) GetHotDataSize() *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataSize {
	return s.HotDataSize
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) GetHotDataSizeDayGrowthRatio() *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	return s.HotDataSizeDayGrowthRatio
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) GetLargeFileCount() *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileCount {
	return s.LargeFileCount
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) GetLargeFileCountDayGrowthRatio() *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	return s.LargeFileCountDayGrowthRatio
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) GetLargeFileDayGrowthCount() *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileDayGrowthCount {
	return s.LargeFileDayGrowthCount
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) GetMediumFileCount() *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileCount {
	return s.MediumFileCount
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) GetMediumFileCountDayGrowthRatio() *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	return s.MediumFileCountDayGrowthRatio
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) GetMediumFileDayGrowthCount() *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileDayGrowthCount {
	return s.MediumFileDayGrowthCount
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) GetSmallFileCount() *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileCount {
	return s.SmallFileCount
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) GetSmallFileCountDayGrowthRatio() *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	return s.SmallFileCountDayGrowthRatio
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) GetSmallFileDayGrowthCount() *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileDayGrowthCount {
	return s.SmallFileDayGrowthCount
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) GetTinyFileCount() *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileCount {
	return s.TinyFileCount
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) GetTinyFileCountDayGrowthRatio() *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	return s.TinyFileCountDayGrowthRatio
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) GetTinyFileDayGrowthCount() *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileDayGrowthCount {
	return s.TinyFileDayGrowthCount
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) GetTotalDataDayGrowthSize() *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataDayGrowthSize {
	return s.TotalDataDayGrowthSize
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) GetTotalDataSize() *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataSize {
	return s.TotalDataSize
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) GetTotalDataSizeDayGrowthRatio() *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	return s.TotalDataSizeDayGrowthRatio
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) GetTotalFileCount() *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileCount {
	return s.TotalFileCount
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) GetTotalFileCountDayGrowthRatio() *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	return s.TotalFileCountDayGrowthRatio
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) GetTotalFileDayGrowthCount() *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileDayGrowthCount {
	return s.TotalFileDayGrowthCount
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) GetWarmDataDayGrowthSize() *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataDayGrowthSize {
	return s.WarmDataDayGrowthSize
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) GetWarmDataSize() *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataSize {
	return s.WarmDataSize
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) GetWarmDataSizeDayGrowthRatio() *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	return s.WarmDataSizeDayGrowthRatio
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) SetColdDataDayGrowthSize(v *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataDayGrowthSize) *ListDoctorHDFSDirectoriesResponseBodyDataMetrics {
	s.ColdDataDayGrowthSize = v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) SetColdDataSize(v *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataSize) *ListDoctorHDFSDirectoriesResponseBodyDataMetrics {
	s.ColdDataSize = v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) SetColdDataSizeDayGrowthRatio(v *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataSizeDayGrowthRatio) *ListDoctorHDFSDirectoriesResponseBodyDataMetrics {
	s.ColdDataSizeDayGrowthRatio = v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) SetEmptyFileCount(v *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileCount) *ListDoctorHDFSDirectoriesResponseBodyDataMetrics {
	s.EmptyFileCount = v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) SetEmptyFileCountDayGrowthRatio(v *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) *ListDoctorHDFSDirectoriesResponseBodyDataMetrics {
	s.EmptyFileCountDayGrowthRatio = v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) SetEmptyFileDayGrowthCount(v *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileDayGrowthCount) *ListDoctorHDFSDirectoriesResponseBodyDataMetrics {
	s.EmptyFileDayGrowthCount = v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) SetFreezeDataDayGrowthSize(v *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataDayGrowthSize) *ListDoctorHDFSDirectoriesResponseBodyDataMetrics {
	s.FreezeDataDayGrowthSize = v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) SetFreezeDataSize(v *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataSize) *ListDoctorHDFSDirectoriesResponseBodyDataMetrics {
	s.FreezeDataSize = v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) SetFreezeDataSizeDayGrowthRatio(v *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) *ListDoctorHDFSDirectoriesResponseBodyDataMetrics {
	s.FreezeDataSizeDayGrowthRatio = v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) SetHotDataDayGrowthSize(v *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataDayGrowthSize) *ListDoctorHDFSDirectoriesResponseBodyDataMetrics {
	s.HotDataDayGrowthSize = v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) SetHotDataSize(v *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataSize) *ListDoctorHDFSDirectoriesResponseBodyDataMetrics {
	s.HotDataSize = v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) SetHotDataSizeDayGrowthRatio(v *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataSizeDayGrowthRatio) *ListDoctorHDFSDirectoriesResponseBodyDataMetrics {
	s.HotDataSizeDayGrowthRatio = v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) SetLargeFileCount(v *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileCount) *ListDoctorHDFSDirectoriesResponseBodyDataMetrics {
	s.LargeFileCount = v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) SetLargeFileCountDayGrowthRatio(v *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileCountDayGrowthRatio) *ListDoctorHDFSDirectoriesResponseBodyDataMetrics {
	s.LargeFileCountDayGrowthRatio = v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) SetLargeFileDayGrowthCount(v *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileDayGrowthCount) *ListDoctorHDFSDirectoriesResponseBodyDataMetrics {
	s.LargeFileDayGrowthCount = v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) SetMediumFileCount(v *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileCount) *ListDoctorHDFSDirectoriesResponseBodyDataMetrics {
	s.MediumFileCount = v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) SetMediumFileCountDayGrowthRatio(v *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileCountDayGrowthRatio) *ListDoctorHDFSDirectoriesResponseBodyDataMetrics {
	s.MediumFileCountDayGrowthRatio = v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) SetMediumFileDayGrowthCount(v *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileDayGrowthCount) *ListDoctorHDFSDirectoriesResponseBodyDataMetrics {
	s.MediumFileDayGrowthCount = v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) SetSmallFileCount(v *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileCount) *ListDoctorHDFSDirectoriesResponseBodyDataMetrics {
	s.SmallFileCount = v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) SetSmallFileCountDayGrowthRatio(v *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileCountDayGrowthRatio) *ListDoctorHDFSDirectoriesResponseBodyDataMetrics {
	s.SmallFileCountDayGrowthRatio = v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) SetSmallFileDayGrowthCount(v *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileDayGrowthCount) *ListDoctorHDFSDirectoriesResponseBodyDataMetrics {
	s.SmallFileDayGrowthCount = v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) SetTinyFileCount(v *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileCount) *ListDoctorHDFSDirectoriesResponseBodyDataMetrics {
	s.TinyFileCount = v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) SetTinyFileCountDayGrowthRatio(v *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileCountDayGrowthRatio) *ListDoctorHDFSDirectoriesResponseBodyDataMetrics {
	s.TinyFileCountDayGrowthRatio = v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) SetTinyFileDayGrowthCount(v *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileDayGrowthCount) *ListDoctorHDFSDirectoriesResponseBodyDataMetrics {
	s.TinyFileDayGrowthCount = v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) SetTotalDataDayGrowthSize(v *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataDayGrowthSize) *ListDoctorHDFSDirectoriesResponseBodyDataMetrics {
	s.TotalDataDayGrowthSize = v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) SetTotalDataSize(v *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataSize) *ListDoctorHDFSDirectoriesResponseBodyDataMetrics {
	s.TotalDataSize = v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) SetTotalDataSizeDayGrowthRatio(v *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) *ListDoctorHDFSDirectoriesResponseBodyDataMetrics {
	s.TotalDataSizeDayGrowthRatio = v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) SetTotalFileCount(v *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileCount) *ListDoctorHDFSDirectoriesResponseBodyDataMetrics {
	s.TotalFileCount = v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) SetTotalFileCountDayGrowthRatio(v *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileCountDayGrowthRatio) *ListDoctorHDFSDirectoriesResponseBodyDataMetrics {
	s.TotalFileCountDayGrowthRatio = v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) SetTotalFileDayGrowthCount(v *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileDayGrowthCount) *ListDoctorHDFSDirectoriesResponseBodyDataMetrics {
	s.TotalFileDayGrowthCount = v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) SetWarmDataDayGrowthSize(v *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataDayGrowthSize) *ListDoctorHDFSDirectoriesResponseBodyDataMetrics {
	s.WarmDataDayGrowthSize = v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) SetWarmDataSize(v *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataSize) *ListDoctorHDFSDirectoriesResponseBodyDataMetrics {
	s.WarmDataSize = v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) SetWarmDataSizeDayGrowthRatio(v *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) *ListDoctorHDFSDirectoriesResponseBodyDataMetrics {
	s.WarmDataSizeDayGrowthRatio = v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetrics) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataDayGrowthSize struct {
	// example:
	//
	// Day growth size of cold data
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// coldDataDayGrowthSize
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// MB
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// example:
	//
	// -182636577752
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataDayGrowthSize) SetDescription(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataDayGrowthSize) SetName(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataDayGrowthSize) SetUnit(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataDayGrowthSize) SetValue(v int64) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataSize struct {
	// example:
	//
	// Size of the cold data size
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// coldDataSize
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// MB
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// example:
	//
	// 5570958082267
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataSize) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataSize) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataSize) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataSize) GetName() *string {
	return s.Name
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataSize) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataSize) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataSize) SetDescription(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataSize {
	s.Description = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataSize) SetName(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataSize {
	s.Name = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataSize) SetUnit(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataSize {
	s.Unit = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataSize) SetValue(v int64) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataSize {
	s.Value = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataSize) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataSizeDayGrowthRatio struct {
	// example:
	//
	// Day growth ratio of cold data size
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// coldDataSizeDayGrowthRatio
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// example:
	//
	// -0.03
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataSizeDayGrowthRatio) SetDescription(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataSizeDayGrowthRatio) SetName(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataSizeDayGrowthRatio) SetUnit(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataSizeDayGrowthRatio) SetValue(v float32) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsColdDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileCount struct {
	// example:
	//
	// Number of empty files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// emptyFileCount
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// example:
	//
	// 15595897
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileCount) SetDescription(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileCount) SetName(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileCount) SetUnit(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileCount) SetValue(v int64) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio struct {
	// example:
	//
	// Day growth ratio of empty files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// emptyFileCountDayGrowthRatio
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// example:
	//
	// 0.005
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) SetDescription(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) SetName(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) SetUnit(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) SetValue(v float32) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileDayGrowthCount struct {
	// example:
	//
	// Day growth count of empty files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// emptyFileDayGrowthCount
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// example:
	//
	// 114
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileDayGrowthCount) SetDescription(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileDayGrowthCount) SetName(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileDayGrowthCount) SetUnit(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileDayGrowthCount) SetValue(v int64) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsEmptyFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataDayGrowthSize struct {
	// example:
	//
	// Day growth size of freeze data
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// freezeDataDayGrowthSize
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// MB
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// example:
	//
	// -167683929450
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataDayGrowthSize) SetDescription(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataDayGrowthSize) SetName(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataDayGrowthSize) SetUnit(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataDayGrowthSize) SetValue(v int64) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataSize struct {
	// example:
	//
	// Size of the freeze data size
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// freezeDataSize
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// MB
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// example:
	//
	// 1231312431
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataSize) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataSize) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataSize) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataSize) GetName() *string {
	return s.Name
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataSize) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataSize) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataSize) SetDescription(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataSize {
	s.Description = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataSize) SetName(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataSize {
	s.Name = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataSize) SetUnit(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataSize {
	s.Unit = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataSize) SetValue(v int64) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataSize {
	s.Value = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataSize) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio struct {
	// example:
	//
	// Day growth ratio of freeze data size
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// freezeDataSizeDayGrowthRatio
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// example:
	//
	// -0.09
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) SetDescription(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) SetName(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) SetUnit(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) SetValue(v float32) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsFreezeDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataDayGrowthSize struct {
	// example:
	//
	// Day growth size of hot data
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// hotDataDayGrowthSize
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// MB
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// example:
	//
	// 123154
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataDayGrowthSize) SetDescription(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataDayGrowthSize) SetName(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataDayGrowthSize) SetUnit(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataDayGrowthSize) SetValue(v int64) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataSize struct {
	// example:
	//
	// Size of the hot data size
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// hotDataSize
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// MB
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// example:
	//
	// 6701531944206
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataSize) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataSize) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataSize) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataSize) GetName() *string {
	return s.Name
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataSize) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataSize) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataSize) SetDescription(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataSize {
	s.Description = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataSize) SetName(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataSize {
	s.Name = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataSize) SetUnit(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataSize {
	s.Unit = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataSize) SetValue(v int64) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataSize {
	s.Value = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataSize) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataSizeDayGrowthRatio struct {
	// example:
	//
	// Day growth ratio of hot data size
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// hotDataSizeDayGrowthRatio
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// example:
	//
	// 0.1114
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataSizeDayGrowthRatio) SetDescription(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataSizeDayGrowthRatio) SetName(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataSizeDayGrowthRatio) SetUnit(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataSizeDayGrowthRatio) SetValue(v float32) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsHotDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileCount struct {
	// example:
	//
	// Number of large files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// largeFileCount
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// example:
	//
	// 123
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileCount) SetDescription(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileCount) SetName(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileCount) SetUnit(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileCount) SetValue(v int64) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileCountDayGrowthRatio struct {
	// example:
	//
	// Day growth ratio of large files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// largeFileCountDayGrowthRatio
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// example:
	//
	// 0.39
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileCountDayGrowthRatio) SetDescription(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileCountDayGrowthRatio) SetName(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileCountDayGrowthRatio) SetUnit(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileCountDayGrowthRatio) SetValue(v float32) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileDayGrowthCount struct {
	// example:
	//
	// Day growth count of large files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// largeFileDayGrowthCount
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// example:
	//
	// 2
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileDayGrowthCount) SetDescription(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileDayGrowthCount) SetName(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileDayGrowthCount) SetUnit(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileDayGrowthCount) SetValue(v int64) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsLargeFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileCount struct {
	// example:
	//
	// Number of medium files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// mediumFileCount
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// example:
	//
	// 323
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileCount) SetDescription(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileCount) SetName(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileCount) SetUnit(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileCount) SetValue(v int64) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileCountDayGrowthRatio struct {
	// example:
	//
	// Day growth ratio of medium files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// mediumFileCountDayGrowthRatio
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// example:
	//
	// 0.19
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileCountDayGrowthRatio) SetDescription(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileCountDayGrowthRatio) SetName(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileCountDayGrowthRatio) SetUnit(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileCountDayGrowthRatio) SetValue(v float32) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileDayGrowthCount struct {
	// example:
	//
	// Day growth count of medium files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// mediumFileDayGrowthCount
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// example:
	//
	// 176
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileDayGrowthCount) SetDescription(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileDayGrowthCount) SetName(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileDayGrowthCount) SetUnit(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileDayGrowthCount) SetValue(v int64) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsMediumFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileCount struct {
	// example:
	//
	// Number of small files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// smallFileCount
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// example:
	//
	// 12345
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileCount) SetDescription(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileCount) SetName(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileCount) SetUnit(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileCount) SetValue(v int64) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileCountDayGrowthRatio struct {
	// example:
	//
	// Day growth ratio of small files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// smallFileCountDayGrowthRatio
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// example:
	//
	// 0.02
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileCountDayGrowthRatio) SetDescription(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileCountDayGrowthRatio) SetName(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileCountDayGrowthRatio) SetUnit(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileCountDayGrowthRatio) SetValue(v float32) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileDayGrowthCount struct {
	// example:
	//
	// Day growth count of small files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// smallFileDayGrowthCount
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// example:
	//
	// 12345
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileDayGrowthCount) SetDescription(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileDayGrowthCount) SetName(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileDayGrowthCount) SetUnit(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileDayGrowthCount) SetValue(v int64) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsSmallFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileCount struct {
	// example:
	//
	// Number of tiny files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// tinyFileCount
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// example:
	//
	// 232131
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileCount) SetDescription(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileCount) SetName(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileCount) SetUnit(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileCount) SetValue(v int64) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileCountDayGrowthRatio struct {
	// example:
	//
	// Day growth ratio of tiny files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// tinyFileCountDayGrowthRatio
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// example:
	//
	// 0.003
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileCountDayGrowthRatio) SetDescription(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileCountDayGrowthRatio) SetName(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileCountDayGrowthRatio) SetUnit(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileCountDayGrowthRatio) SetValue(v float32) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileDayGrowthCount struct {
	// example:
	//
	// Day growth count of tiny files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// tinyFileDayGrowthCount
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// example:
	//
	// -123
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileDayGrowthCount) SetDescription(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileDayGrowthCount) SetName(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileDayGrowthCount) SetUnit(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileDayGrowthCount) SetValue(v int64) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTinyFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataDayGrowthSize struct {
	// example:
	//
	// Day growth size of total data
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// totalDataDayGrowthSize
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// MB
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// example:
	//
	// 256482228248
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataDayGrowthSize) SetDescription(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataDayGrowthSize) SetName(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataDayGrowthSize) SetUnit(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataDayGrowthSize) SetValue(v int64) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataSize struct {
	// example:
	//
	// Total data size in megabytes (MB)
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// totalDataSize
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// MB
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// example:
	//
	// 62086342083623
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataSize) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataSize) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataSize) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataSize) GetName() *string {
	return s.Name
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataSize) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataSize) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataSize) SetDescription(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataSize {
	s.Description = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataSize) SetName(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataSize {
	s.Name = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataSize) SetUnit(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataSize {
	s.Unit = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataSize) SetValue(v int64) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataSize {
	s.Value = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataSize) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio struct {
	// example:
	//
	// Day growth ratio of total data size
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// totalDataSizeDayGrowthRatio
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// example:
	//
	// 0.14
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) SetDescription(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) SetName(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) SetUnit(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) SetValue(v float32) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileCount struct {
	// example:
	//
	// Number of total files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// totalFileCount
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// example:
	//
	// 51683279
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileCount) SetDescription(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileCount) SetName(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileCount) SetUnit(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileCount) SetValue(v int64) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileCountDayGrowthRatio struct {
	// example:
	//
	// Day growth ratio of total files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// totalFileCountDayGrowthRatio
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// example:
	//
	// 0.02
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileCountDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileCountDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileCountDayGrowthRatio) SetDescription(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileCountDayGrowthRatio) SetName(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileCountDayGrowthRatio) SetUnit(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileCountDayGrowthRatio) SetValue(v float32) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileCountDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileCountDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileDayGrowthCount struct {
	// example:
	//
	// Day growth count of total files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// totalFileDayGrowthCount
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// example:
	//
	// 27809
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileDayGrowthCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileDayGrowthCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileDayGrowthCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileDayGrowthCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileDayGrowthCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileDayGrowthCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileDayGrowthCount) SetDescription(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileDayGrowthCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileDayGrowthCount) SetName(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileDayGrowthCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileDayGrowthCount) SetUnit(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileDayGrowthCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileDayGrowthCount) SetValue(v int64) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileDayGrowthCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsTotalFileDayGrowthCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataDayGrowthSize struct {
	// example:
	//
	// Day growth size of warm data
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// warmDataDayGrowthSize
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// MB
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// example:
	//
	// -64806998319
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataDayGrowthSize) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataDayGrowthSize) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataDayGrowthSize) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataDayGrowthSize) GetName() *string {
	return s.Name
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataDayGrowthSize) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataDayGrowthSize) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataDayGrowthSize) SetDescription(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataDayGrowthSize {
	s.Description = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataDayGrowthSize) SetName(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataDayGrowthSize {
	s.Name = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataDayGrowthSize) SetUnit(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataDayGrowthSize {
	s.Unit = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataDayGrowthSize) SetValue(v int64) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataDayGrowthSize {
	s.Value = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataDayGrowthSize) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataSize struct {
	// example:
	//
	// Size of the warm data size
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// warmDataSize
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// MB
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// example:
	//
	// 4062349775577
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataSize) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataSize) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataSize) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataSize) GetName() *string {
	return s.Name
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataSize) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataSize) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataSize) SetDescription(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataSize {
	s.Description = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataSize) SetName(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataSize {
	s.Name = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataSize) SetUnit(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataSize {
	s.Unit = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataSize) SetValue(v int64) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataSize {
	s.Value = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataSize) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio struct {
	// example:
	//
	// Day growth ratio of warm data size
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// warmDataSizeDayGrowthRatio
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// example:
	//
	// -0.015
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GetName() *string {
	return s.Name
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) SetDescription(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	s.Description = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) SetName(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	s.Name = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) SetUnit(v string) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	s.Unit = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) SetValue(v float32) *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio {
	s.Value = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponseBodyDataMetricsWarmDataSizeDayGrowthRatio) Validate() error {
	return dara.Validate(s)
}
