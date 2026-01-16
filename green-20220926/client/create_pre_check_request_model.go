// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePreCheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucketPrefixFilterConfig(v string) *CreatePreCheckRequest
	GetBucketPrefixFilterConfig() *string
	SetBuckets(v string) *CreatePreCheckRequest
	GetBuckets() *string
	SetDistinctHistoryTasks(v bool) *CreatePreCheckRequest
	GetDistinctHistoryTasks() *bool
	SetEndTime(v string) *CreatePreCheckRequest
	GetEndTime() *string
	SetIsInc(v bool) *CreatePreCheckRequest
	GetIsInc() *bool
	SetMediaType(v int32) *CreatePreCheckRequest
	GetMediaType() *int32
	SetPrefixFilterType(v string) *CreatePreCheckRequest
	GetPrefixFilterType() *string
	SetPrefixFilters(v string) *CreatePreCheckRequest
	GetPrefixFilters() *string
	SetPriority(v int32) *CreatePreCheckRequest
	GetPriority() *int32
	SetRegionId(v string) *CreatePreCheckRequest
	GetRegionId() *string
	SetScanLimit(v int64) *CreatePreCheckRequest
	GetScanLimit() *int64
	SetScanNoFileType(v bool) *CreatePreCheckRequest
	GetScanNoFileType() *bool
	SetScanService(v string) *CreatePreCheckRequest
	GetScanService() *string
	SetStartTime(v string) *CreatePreCheckRequest
	GetStartTime() *string
	SetTaskName(v string) *CreatePreCheckRequest
	GetTaskName() *string
}

type CreatePreCheckRequest struct {
	BucketPrefixFilterConfig *string `json:"BucketPrefixFilterConfig,omitempty" xml:"BucketPrefixFilterConfig,omitempty"`
	// Buckets.
	//
	// example:
	//
	// [{\\"Bucket\\":\\"bucket01-test\\",\\"Region\\":\\"cn-beijing\\"}]
	Buckets *string `json:"Buckets,omitempty" xml:"Buckets,omitempty"`
	// Whether to deduplicate historical detected tasks.
	//
	// example:
	//
	// true
	DistinctHistoryTasks *bool `json:"DistinctHistoryTasks,omitempty" xml:"DistinctHistoryTasks,omitempty"`
	// Task end time.
	//
	// example:
	//
	// 2023-12-18 10:08:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Whether it is a scheduled scan task.
	//
	// example:
	//
	// false
	IsInc *bool `json:"IsInc,omitempty" xml:"IsInc,omitempty"`
	// Media type.
	//
	// example:
	//
	// 1
	MediaType *int32 `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// Prefix filter type.
	//
	// example:
	//
	// all
	PrefixFilterType *string `json:"PrefixFilterType,omitempty" xml:"PrefixFilterType,omitempty"`
	// Prefixes.
	//
	// example:
	//
	// dir1,dir2
	PrefixFilters *string `json:"PrefixFilters,omitempty" xml:"PrefixFilters,omitempty"`
	// Priority.
	//
	// example:
	//
	// 0
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Scan limit count.
	//
	// example:
	//
	// 10
	ScanLimit *int64 `json:"ScanLimit,omitempty" xml:"ScanLimit,omitempty"`
	// Whether to scan images without file extensions.
	//
	// example:
	//
	// true
	ScanNoFileType *bool `json:"ScanNoFileType,omitempty" xml:"ScanNoFileType,omitempty"`
	// Scan service code.
	//
	// example:
	//
	// baselineCheck
	ScanService *string `json:"ScanService,omitempty" xml:"ScanService,omitempty"`
	// Task start time.
	//
	// example:
	//
	// 2023-12-17 10:08:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Task name.
	//
	// example:
	//
	// 图片任务 20240709101602004
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreatePreCheckRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePreCheckRequest) GoString() string {
	return s.String()
}

func (s *CreatePreCheckRequest) GetBucketPrefixFilterConfig() *string {
	return s.BucketPrefixFilterConfig
}

func (s *CreatePreCheckRequest) GetBuckets() *string {
	return s.Buckets
}

func (s *CreatePreCheckRequest) GetDistinctHistoryTasks() *bool {
	return s.DistinctHistoryTasks
}

func (s *CreatePreCheckRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CreatePreCheckRequest) GetIsInc() *bool {
	return s.IsInc
}

func (s *CreatePreCheckRequest) GetMediaType() *int32 {
	return s.MediaType
}

func (s *CreatePreCheckRequest) GetPrefixFilterType() *string {
	return s.PrefixFilterType
}

func (s *CreatePreCheckRequest) GetPrefixFilters() *string {
	return s.PrefixFilters
}

func (s *CreatePreCheckRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreatePreCheckRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePreCheckRequest) GetScanLimit() *int64 {
	return s.ScanLimit
}

func (s *CreatePreCheckRequest) GetScanNoFileType() *bool {
	return s.ScanNoFileType
}

func (s *CreatePreCheckRequest) GetScanService() *string {
	return s.ScanService
}

func (s *CreatePreCheckRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreatePreCheckRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreatePreCheckRequest) SetBucketPrefixFilterConfig(v string) *CreatePreCheckRequest {
	s.BucketPrefixFilterConfig = &v
	return s
}

func (s *CreatePreCheckRequest) SetBuckets(v string) *CreatePreCheckRequest {
	s.Buckets = &v
	return s
}

func (s *CreatePreCheckRequest) SetDistinctHistoryTasks(v bool) *CreatePreCheckRequest {
	s.DistinctHistoryTasks = &v
	return s
}

func (s *CreatePreCheckRequest) SetEndTime(v string) *CreatePreCheckRequest {
	s.EndTime = &v
	return s
}

func (s *CreatePreCheckRequest) SetIsInc(v bool) *CreatePreCheckRequest {
	s.IsInc = &v
	return s
}

func (s *CreatePreCheckRequest) SetMediaType(v int32) *CreatePreCheckRequest {
	s.MediaType = &v
	return s
}

func (s *CreatePreCheckRequest) SetPrefixFilterType(v string) *CreatePreCheckRequest {
	s.PrefixFilterType = &v
	return s
}

func (s *CreatePreCheckRequest) SetPrefixFilters(v string) *CreatePreCheckRequest {
	s.PrefixFilters = &v
	return s
}

func (s *CreatePreCheckRequest) SetPriority(v int32) *CreatePreCheckRequest {
	s.Priority = &v
	return s
}

func (s *CreatePreCheckRequest) SetRegionId(v string) *CreatePreCheckRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePreCheckRequest) SetScanLimit(v int64) *CreatePreCheckRequest {
	s.ScanLimit = &v
	return s
}

func (s *CreatePreCheckRequest) SetScanNoFileType(v bool) *CreatePreCheckRequest {
	s.ScanNoFileType = &v
	return s
}

func (s *CreatePreCheckRequest) SetScanService(v string) *CreatePreCheckRequest {
	s.ScanService = &v
	return s
}

func (s *CreatePreCheckRequest) SetStartTime(v string) *CreatePreCheckRequest {
	s.StartTime = &v
	return s
}

func (s *CreatePreCheckRequest) SetTaskName(v string) *CreatePreCheckRequest {
	s.TaskName = &v
	return s
}

func (s *CreatePreCheckRequest) Validate() error {
	return dara.Validate(s)
}
