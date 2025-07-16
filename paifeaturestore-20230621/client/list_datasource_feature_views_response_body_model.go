// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasourceFeatureViewsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureViews(v []*ListDatasourceFeatureViewsResponseBodyFeatureViews) *ListDatasourceFeatureViewsResponseBody
	GetFeatureViews() []*ListDatasourceFeatureViewsResponseBodyFeatureViews
	SetTotalCount(v int64) *ListDatasourceFeatureViewsResponseBody
	GetTotalCount() *int64
	SetTotalUsageStatistics(v *ListDatasourceFeatureViewsResponseBodyTotalUsageStatistics) *ListDatasourceFeatureViewsResponseBody
	GetTotalUsageStatistics() *ListDatasourceFeatureViewsResponseBodyTotalUsageStatistics
	SetRequestId(v string) *ListDatasourceFeatureViewsResponseBody
	GetRequestId() *string
}

type ListDatasourceFeatureViewsResponseBody struct {
	FeatureViews []*ListDatasourceFeatureViewsResponseBodyFeatureViews `json:"FeatureViews,omitempty" xml:"FeatureViews,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalCount           *int64                                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalUsageStatistics *ListDatasourceFeatureViewsResponseBodyTotalUsageStatistics `json:"TotalUsageStatistics,omitempty" xml:"TotalUsageStatistics,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 514F82AF-3C04-5C3D-8F38-A11261BF37B0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListDatasourceFeatureViewsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDatasourceFeatureViewsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatasourceFeatureViewsResponseBody) GetFeatureViews() []*ListDatasourceFeatureViewsResponseBodyFeatureViews {
	return s.FeatureViews
}

func (s *ListDatasourceFeatureViewsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDatasourceFeatureViewsResponseBody) GetTotalUsageStatistics() *ListDatasourceFeatureViewsResponseBodyTotalUsageStatistics {
	return s.TotalUsageStatistics
}

func (s *ListDatasourceFeatureViewsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDatasourceFeatureViewsResponseBody) SetFeatureViews(v []*ListDatasourceFeatureViewsResponseBodyFeatureViews) *ListDatasourceFeatureViewsResponseBody {
	s.FeatureViews = v
	return s
}

func (s *ListDatasourceFeatureViewsResponseBody) SetTotalCount(v int64) *ListDatasourceFeatureViewsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDatasourceFeatureViewsResponseBody) SetTotalUsageStatistics(v *ListDatasourceFeatureViewsResponseBodyTotalUsageStatistics) *ListDatasourceFeatureViewsResponseBody {
	s.TotalUsageStatistics = v
	return s
}

func (s *ListDatasourceFeatureViewsResponseBody) SetRequestId(v string) *ListDatasourceFeatureViewsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDatasourceFeatureViewsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDatasourceFeatureViewsResponseBodyFeatureViews struct {
	// example:
	//
	// {"shard_count":5,"replication_count":1}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// user
	FeatureEntityName *string `json:"FeatureEntityName,omitempty" xml:"FeatureEntityName,omitempty"`
	// example:
	//
	// 3
	FeatureViewId *string `json:"FeatureViewId,omitempty" xml:"FeatureViewId,omitempty"`
	// example:
	//
	// fv1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// p1
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// 86400
	TTL *int32 `json:"TTL,omitempty" xml:"TTL,omitempty"`
	// example:
	//
	// Batch
	Type            *string                                                            `json:"Type,omitempty" xml:"Type,omitempty"`
	UsageStatistics *ListDatasourceFeatureViewsResponseBodyFeatureViewsUsageStatistics `json:"UsageStatistics,omitempty" xml:"UsageStatistics,omitempty" type:"Struct"`
}

func (s ListDatasourceFeatureViewsResponseBodyFeatureViews) String() string {
	return dara.Prettify(s)
}

func (s ListDatasourceFeatureViewsResponseBodyFeatureViews) GoString() string {
	return s.String()
}

func (s *ListDatasourceFeatureViewsResponseBodyFeatureViews) GetConfig() *string {
	return s.Config
}

func (s *ListDatasourceFeatureViewsResponseBodyFeatureViews) GetFeatureEntityName() *string {
	return s.FeatureEntityName
}

func (s *ListDatasourceFeatureViewsResponseBodyFeatureViews) GetFeatureViewId() *string {
	return s.FeatureViewId
}

func (s *ListDatasourceFeatureViewsResponseBodyFeatureViews) GetName() *string {
	return s.Name
}

func (s *ListDatasourceFeatureViewsResponseBodyFeatureViews) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListDatasourceFeatureViewsResponseBodyFeatureViews) GetTTL() *int32 {
	return s.TTL
}

func (s *ListDatasourceFeatureViewsResponseBodyFeatureViews) GetType() *string {
	return s.Type
}

func (s *ListDatasourceFeatureViewsResponseBodyFeatureViews) GetUsageStatistics() *ListDatasourceFeatureViewsResponseBodyFeatureViewsUsageStatistics {
	return s.UsageStatistics
}

func (s *ListDatasourceFeatureViewsResponseBodyFeatureViews) SetConfig(v string) *ListDatasourceFeatureViewsResponseBodyFeatureViews {
	s.Config = &v
	return s
}

func (s *ListDatasourceFeatureViewsResponseBodyFeatureViews) SetFeatureEntityName(v string) *ListDatasourceFeatureViewsResponseBodyFeatureViews {
	s.FeatureEntityName = &v
	return s
}

func (s *ListDatasourceFeatureViewsResponseBodyFeatureViews) SetFeatureViewId(v string) *ListDatasourceFeatureViewsResponseBodyFeatureViews {
	s.FeatureViewId = &v
	return s
}

func (s *ListDatasourceFeatureViewsResponseBodyFeatureViews) SetName(v string) *ListDatasourceFeatureViewsResponseBodyFeatureViews {
	s.Name = &v
	return s
}

func (s *ListDatasourceFeatureViewsResponseBodyFeatureViews) SetProjectName(v string) *ListDatasourceFeatureViewsResponseBodyFeatureViews {
	s.ProjectName = &v
	return s
}

func (s *ListDatasourceFeatureViewsResponseBodyFeatureViews) SetTTL(v int32) *ListDatasourceFeatureViewsResponseBodyFeatureViews {
	s.TTL = &v
	return s
}

func (s *ListDatasourceFeatureViewsResponseBodyFeatureViews) SetType(v string) *ListDatasourceFeatureViewsResponseBodyFeatureViews {
	s.Type = &v
	return s
}

func (s *ListDatasourceFeatureViewsResponseBodyFeatureViews) SetUsageStatistics(v *ListDatasourceFeatureViewsResponseBodyFeatureViewsUsageStatistics) *ListDatasourceFeatureViewsResponseBodyFeatureViews {
	s.UsageStatistics = v
	return s
}

func (s *ListDatasourceFeatureViewsResponseBodyFeatureViews) Validate() error {
	return dara.Validate(s)
}

type ListDatasourceFeatureViewsResponseBodyFeatureViewsUsageStatistics struct {
	// example:
	//
	// 1.23
	DiskUsage *float64 `json:"DiskUsage,omitempty" xml:"DiskUsage,omitempty"`
	// example:
	//
	// 0.12
	MemoryUsage    *float64                                                                           `json:"MemoryUsage,omitempty" xml:"MemoryUsage,omitempty"`
	ReadWriteCount []*ListDatasourceFeatureViewsResponseBodyFeatureViewsUsageStatisticsReadWriteCount `json:"ReadWriteCount,omitempty" xml:"ReadWriteCount,omitempty" type:"Repeated"`
	// example:
	//
	// 10000
	RowCount *int64 `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
}

func (s ListDatasourceFeatureViewsResponseBodyFeatureViewsUsageStatistics) String() string {
	return dara.Prettify(s)
}

func (s ListDatasourceFeatureViewsResponseBodyFeatureViewsUsageStatistics) GoString() string {
	return s.String()
}

func (s *ListDatasourceFeatureViewsResponseBodyFeatureViewsUsageStatistics) GetDiskUsage() *float64 {
	return s.DiskUsage
}

func (s *ListDatasourceFeatureViewsResponseBodyFeatureViewsUsageStatistics) GetMemoryUsage() *float64 {
	return s.MemoryUsage
}

func (s *ListDatasourceFeatureViewsResponseBodyFeatureViewsUsageStatistics) GetReadWriteCount() []*ListDatasourceFeatureViewsResponseBodyFeatureViewsUsageStatisticsReadWriteCount {
	return s.ReadWriteCount
}

func (s *ListDatasourceFeatureViewsResponseBodyFeatureViewsUsageStatistics) GetRowCount() *int64 {
	return s.RowCount
}

func (s *ListDatasourceFeatureViewsResponseBodyFeatureViewsUsageStatistics) SetDiskUsage(v float64) *ListDatasourceFeatureViewsResponseBodyFeatureViewsUsageStatistics {
	s.DiskUsage = &v
	return s
}

func (s *ListDatasourceFeatureViewsResponseBodyFeatureViewsUsageStatistics) SetMemoryUsage(v float64) *ListDatasourceFeatureViewsResponseBodyFeatureViewsUsageStatistics {
	s.MemoryUsage = &v
	return s
}

func (s *ListDatasourceFeatureViewsResponseBodyFeatureViewsUsageStatistics) SetReadWriteCount(v []*ListDatasourceFeatureViewsResponseBodyFeatureViewsUsageStatisticsReadWriteCount) *ListDatasourceFeatureViewsResponseBodyFeatureViewsUsageStatistics {
	s.ReadWriteCount = v
	return s
}

func (s *ListDatasourceFeatureViewsResponseBodyFeatureViewsUsageStatistics) SetRowCount(v int64) *ListDatasourceFeatureViewsResponseBodyFeatureViewsUsageStatistics {
	s.RowCount = &v
	return s
}

func (s *ListDatasourceFeatureViewsResponseBodyFeatureViewsUsageStatistics) Validate() error {
	return dara.Validate(s)
}

type ListDatasourceFeatureViewsResponseBodyFeatureViewsUsageStatisticsReadWriteCount struct {
	// example:
	//
	// 2025-03-18T00:00:00+08:00
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// example:
	//
	// 200
	ReadCount *int64 `json:"ReadCount,omitempty" xml:"ReadCount,omitempty"`
	// example:
	//
	// 100
	WriteCount *int64 `json:"WriteCount,omitempty" xml:"WriteCount,omitempty"`
}

func (s ListDatasourceFeatureViewsResponseBodyFeatureViewsUsageStatisticsReadWriteCount) String() string {
	return dara.Prettify(s)
}

func (s ListDatasourceFeatureViewsResponseBodyFeatureViewsUsageStatisticsReadWriteCount) GoString() string {
	return s.String()
}

func (s *ListDatasourceFeatureViewsResponseBodyFeatureViewsUsageStatisticsReadWriteCount) GetDate() *string {
	return s.Date
}

func (s *ListDatasourceFeatureViewsResponseBodyFeatureViewsUsageStatisticsReadWriteCount) GetReadCount() *int64 {
	return s.ReadCount
}

func (s *ListDatasourceFeatureViewsResponseBodyFeatureViewsUsageStatisticsReadWriteCount) GetWriteCount() *int64 {
	return s.WriteCount
}

func (s *ListDatasourceFeatureViewsResponseBodyFeatureViewsUsageStatisticsReadWriteCount) SetDate(v string) *ListDatasourceFeatureViewsResponseBodyFeatureViewsUsageStatisticsReadWriteCount {
	s.Date = &v
	return s
}

func (s *ListDatasourceFeatureViewsResponseBodyFeatureViewsUsageStatisticsReadWriteCount) SetReadCount(v int64) *ListDatasourceFeatureViewsResponseBodyFeatureViewsUsageStatisticsReadWriteCount {
	s.ReadCount = &v
	return s
}

func (s *ListDatasourceFeatureViewsResponseBodyFeatureViewsUsageStatisticsReadWriteCount) SetWriteCount(v int64) *ListDatasourceFeatureViewsResponseBodyFeatureViewsUsageStatisticsReadWriteCount {
	s.WriteCount = &v
	return s
}

func (s *ListDatasourceFeatureViewsResponseBodyFeatureViewsUsageStatisticsReadWriteCount) Validate() error {
	return dara.Validate(s)
}

type ListDatasourceFeatureViewsResponseBodyTotalUsageStatistics struct {
	// example:
	//
	// 12.3
	TotalDiskUsage *float64 `json:"TotalDiskUsage,omitempty" xml:"TotalDiskUsage,omitempty"`
	// example:
	//
	// 1.23
	TotalMemoryUsage    *float64                                                                         `json:"TotalMemoryUsage,omitempty" xml:"TotalMemoryUsage,omitempty"`
	TotalReadWriteCount []*ListDatasourceFeatureViewsResponseBodyTotalUsageStatisticsTotalReadWriteCount `json:"TotalReadWriteCount,omitempty" xml:"TotalReadWriteCount,omitempty" type:"Repeated"`
}

func (s ListDatasourceFeatureViewsResponseBodyTotalUsageStatistics) String() string {
	return dara.Prettify(s)
}

func (s ListDatasourceFeatureViewsResponseBodyTotalUsageStatistics) GoString() string {
	return s.String()
}

func (s *ListDatasourceFeatureViewsResponseBodyTotalUsageStatistics) GetTotalDiskUsage() *float64 {
	return s.TotalDiskUsage
}

func (s *ListDatasourceFeatureViewsResponseBodyTotalUsageStatistics) GetTotalMemoryUsage() *float64 {
	return s.TotalMemoryUsage
}

func (s *ListDatasourceFeatureViewsResponseBodyTotalUsageStatistics) GetTotalReadWriteCount() []*ListDatasourceFeatureViewsResponseBodyTotalUsageStatisticsTotalReadWriteCount {
	return s.TotalReadWriteCount
}

func (s *ListDatasourceFeatureViewsResponseBodyTotalUsageStatistics) SetTotalDiskUsage(v float64) *ListDatasourceFeatureViewsResponseBodyTotalUsageStatistics {
	s.TotalDiskUsage = &v
	return s
}

func (s *ListDatasourceFeatureViewsResponseBodyTotalUsageStatistics) SetTotalMemoryUsage(v float64) *ListDatasourceFeatureViewsResponseBodyTotalUsageStatistics {
	s.TotalMemoryUsage = &v
	return s
}

func (s *ListDatasourceFeatureViewsResponseBodyTotalUsageStatistics) SetTotalReadWriteCount(v []*ListDatasourceFeatureViewsResponseBodyTotalUsageStatisticsTotalReadWriteCount) *ListDatasourceFeatureViewsResponseBodyTotalUsageStatistics {
	s.TotalReadWriteCount = v
	return s
}

func (s *ListDatasourceFeatureViewsResponseBodyTotalUsageStatistics) Validate() error {
	return dara.Validate(s)
}

type ListDatasourceFeatureViewsResponseBodyTotalUsageStatisticsTotalReadWriteCount struct {
	// example:
	//
	// 2025-03-18T00:00:00+08:00
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// example:
	//
	// 456
	TotalReadCount *int64 `json:"TotalReadCount,omitempty" xml:"TotalReadCount,omitempty"`
	// example:
	//
	// 123
	TotalWriteCount *int64 `json:"TotalWriteCount,omitempty" xml:"TotalWriteCount,omitempty"`
}

func (s ListDatasourceFeatureViewsResponseBodyTotalUsageStatisticsTotalReadWriteCount) String() string {
	return dara.Prettify(s)
}

func (s ListDatasourceFeatureViewsResponseBodyTotalUsageStatisticsTotalReadWriteCount) GoString() string {
	return s.String()
}

func (s *ListDatasourceFeatureViewsResponseBodyTotalUsageStatisticsTotalReadWriteCount) GetDate() *string {
	return s.Date
}

func (s *ListDatasourceFeatureViewsResponseBodyTotalUsageStatisticsTotalReadWriteCount) GetTotalReadCount() *int64 {
	return s.TotalReadCount
}

func (s *ListDatasourceFeatureViewsResponseBodyTotalUsageStatisticsTotalReadWriteCount) GetTotalWriteCount() *int64 {
	return s.TotalWriteCount
}

func (s *ListDatasourceFeatureViewsResponseBodyTotalUsageStatisticsTotalReadWriteCount) SetDate(v string) *ListDatasourceFeatureViewsResponseBodyTotalUsageStatisticsTotalReadWriteCount {
	s.Date = &v
	return s
}

func (s *ListDatasourceFeatureViewsResponseBodyTotalUsageStatisticsTotalReadWriteCount) SetTotalReadCount(v int64) *ListDatasourceFeatureViewsResponseBodyTotalUsageStatisticsTotalReadWriteCount {
	s.TotalReadCount = &v
	return s
}

func (s *ListDatasourceFeatureViewsResponseBodyTotalUsageStatisticsTotalReadWriteCount) SetTotalWriteCount(v int64) *ListDatasourceFeatureViewsResponseBodyTotalUsageStatisticsTotalReadWriteCount {
	s.TotalWriteCount = &v
	return s
}

func (s *ListDatasourceFeatureViewsResponseBodyTotalUsageStatisticsTotalReadWriteCount) Validate() error {
	return dara.Validate(s)
}
