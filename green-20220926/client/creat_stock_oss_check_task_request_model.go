// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatStockOssCheckTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBuckets(v string) *CreatStockOssCheckTaskRequest
	GetBuckets() *string
	SetCallbackId(v string) *CreatStockOssCheckTaskRequest
	GetCallbackId() *string
	SetDistinctHistoryTasks(v bool) *CreatStockOssCheckTaskRequest
	GetDistinctHistoryTasks() *bool
	SetEndTime(v string) *CreatStockOssCheckTaskRequest
	GetEndTime() *string
	SetExecuteDate(v int32) *CreatStockOssCheckTaskRequest
	GetExecuteDate() *int32
	SetExecuteTime(v string) *CreatStockOssCheckTaskRequest
	GetExecuteTime() *string
	SetFreeze(v bool) *CreatStockOssCheckTaskRequest
	GetFreeze() *bool
	SetFreezeHighRisk1(v bool) *CreatStockOssCheckTaskRequest
	GetFreezeHighRisk1() *bool
	SetFreezeHighRisk2(v bool) *CreatStockOssCheckTaskRequest
	GetFreezeHighRisk2() *bool
	SetFreezeMediumRisk1(v bool) *CreatStockOssCheckTaskRequest
	GetFreezeMediumRisk1() *bool
	SetFreezeMediumRisk2(v bool) *CreatStockOssCheckTaskRequest
	GetFreezeMediumRisk2() *bool
	SetFreezeRestorePath(v string) *CreatStockOssCheckTaskRequest
	GetFreezeRestorePath() *string
	SetFreezeType(v string) *CreatStockOssCheckTaskRequest
	GetFreezeType() *string
	SetIsInc(v bool) *CreatStockOssCheckTaskRequest
	GetIsInc() *bool
	SetMediaType(v int32) *CreatStockOssCheckTaskRequest
	GetMediaType() *int32
	SetPrefixFilterType(v string) *CreatStockOssCheckTaskRequest
	GetPrefixFilterType() *string
	SetPrefixFilters(v string) *CreatStockOssCheckTaskRequest
	GetPrefixFilters() *string
	SetPriority(v int32) *CreatStockOssCheckTaskRequest
	GetPriority() *int32
	SetReferer(v string) *CreatStockOssCheckTaskRequest
	GetReferer() *string
	SetRegionId(v string) *CreatStockOssCheckTaskRequest
	GetRegionId() *string
	SetScanLimit(v int64) *CreatStockOssCheckTaskRequest
	GetScanLimit() *int64
	SetScanNoFileType(v bool) *CreatStockOssCheckTaskRequest
	GetScanNoFileType() *bool
	SetScanResourceType(v string) *CreatStockOssCheckTaskRequest
	GetScanResourceType() *string
	SetScanService(v string) *CreatStockOssCheckTaskRequest
	GetScanService() *string
	SetStartTime(v string) *CreatStockOssCheckTaskRequest
	GetStartTime() *string
	SetTaskCycle(v int32) *CreatStockOssCheckTaskRequest
	GetTaskCycle() *int32
	SetTaskName(v string) *CreatStockOssCheckTaskRequest
	GetTaskName() *string
	SetTaskType(v string) *CreatStockOssCheckTaskRequest
	GetTaskType() *string
}

type CreatStockOssCheckTaskRequest struct {
	// example:
	//
	// [{\\"Bucket\\":\\"bucket01-test\\",\\"Region\\":\\"cn-beijing\\"}]
	Buckets *string `json:"Buckets,omitempty" xml:"Buckets,omitempty"`
	// example:
	//
	// 1751
	CallbackId *string `json:"CallbackId,omitempty" xml:"CallbackId,omitempty"`
	// example:
	//
	// true
	DistinctHistoryTasks *bool `json:"DistinctHistoryTasks,omitempty" xml:"DistinctHistoryTasks,omitempty"`
	// example:
	//
	// 2023-12-18 10:08:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	ExecuteDate *int32 `json:"ExecuteDate,omitempty" xml:"ExecuteDate,omitempty"`
	// example:
	//
	// 01:09:30-01:19:30
	ExecuteTime       *string `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	Freeze            *bool   `json:"Freeze,omitempty" xml:"Freeze,omitempty"`
	FreezeHighRisk1   *bool   `json:"FreezeHighRisk1,omitempty" xml:"FreezeHighRisk1,omitempty"`
	FreezeHighRisk2   *bool   `json:"FreezeHighRisk2,omitempty" xml:"FreezeHighRisk2,omitempty"`
	FreezeMediumRisk1 *bool   `json:"FreezeMediumRisk1,omitempty" xml:"FreezeMediumRisk1,omitempty"`
	FreezeMediumRisk2 *bool   `json:"FreezeMediumRisk2,omitempty" xml:"FreezeMediumRisk2,omitempty"`
	FreezeRestorePath *string `json:"FreezeRestorePath,omitempty" xml:"FreezeRestorePath,omitempty"`
	FreezeType        *string `json:"FreezeType,omitempty" xml:"FreezeType,omitempty"`
	// example:
	//
	// false
	IsInc *bool `json:"IsInc,omitempty" xml:"IsInc,omitempty"`
	// example:
	//
	// 1
	MediaType *int32 `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// example:
	//
	// all
	PrefixFilterType *string `json:"PrefixFilterType,omitempty" xml:"PrefixFilterType,omitempty"`
	// example:
	//
	// dir1,dir2
	PrefixFilters *string `json:"PrefixFilters,omitempty" xml:"PrefixFilters,omitempty"`
	// example:
	//
	// 0
	Priority *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Referer  *string `json:"Referer,omitempty" xml:"Referer,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 10
	ScanLimit *int64 `json:"ScanLimit,omitempty" xml:"ScanLimit,omitempty"`
	// example:
	//
	// true
	ScanNoFileType *bool `json:"ScanNoFileType,omitempty" xml:"ScanNoFileType,omitempty"`
	// example:
	//
	// 0
	ScanResourceType *string `json:"ScanResourceType,omitempty" xml:"ScanResourceType,omitempty"`
	// example:
	//
	// baselineCheck
	ScanService *string `json:"ScanService,omitempty" xml:"ScanService,omitempty"`
	// example:
	//
	// 2023-12-17 10:08:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 0
	TaskCycle *int32  `json:"TaskCycle,omitempty" xml:"TaskCycle,omitempty"`
	TaskName  *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// batch
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s CreatStockOssCheckTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatStockOssCheckTaskRequest) GoString() string {
	return s.String()
}

func (s *CreatStockOssCheckTaskRequest) GetBuckets() *string {
	return s.Buckets
}

func (s *CreatStockOssCheckTaskRequest) GetCallbackId() *string {
	return s.CallbackId
}

func (s *CreatStockOssCheckTaskRequest) GetDistinctHistoryTasks() *bool {
	return s.DistinctHistoryTasks
}

func (s *CreatStockOssCheckTaskRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CreatStockOssCheckTaskRequest) GetExecuteDate() *int32 {
	return s.ExecuteDate
}

func (s *CreatStockOssCheckTaskRequest) GetExecuteTime() *string {
	return s.ExecuteTime
}

func (s *CreatStockOssCheckTaskRequest) GetFreeze() *bool {
	return s.Freeze
}

func (s *CreatStockOssCheckTaskRequest) GetFreezeHighRisk1() *bool {
	return s.FreezeHighRisk1
}

func (s *CreatStockOssCheckTaskRequest) GetFreezeHighRisk2() *bool {
	return s.FreezeHighRisk2
}

func (s *CreatStockOssCheckTaskRequest) GetFreezeMediumRisk1() *bool {
	return s.FreezeMediumRisk1
}

func (s *CreatStockOssCheckTaskRequest) GetFreezeMediumRisk2() *bool {
	return s.FreezeMediumRisk2
}

func (s *CreatStockOssCheckTaskRequest) GetFreezeRestorePath() *string {
	return s.FreezeRestorePath
}

func (s *CreatStockOssCheckTaskRequest) GetFreezeType() *string {
	return s.FreezeType
}

func (s *CreatStockOssCheckTaskRequest) GetIsInc() *bool {
	return s.IsInc
}

func (s *CreatStockOssCheckTaskRequest) GetMediaType() *int32 {
	return s.MediaType
}

func (s *CreatStockOssCheckTaskRequest) GetPrefixFilterType() *string {
	return s.PrefixFilterType
}

func (s *CreatStockOssCheckTaskRequest) GetPrefixFilters() *string {
	return s.PrefixFilters
}

func (s *CreatStockOssCheckTaskRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreatStockOssCheckTaskRequest) GetReferer() *string {
	return s.Referer
}

func (s *CreatStockOssCheckTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatStockOssCheckTaskRequest) GetScanLimit() *int64 {
	return s.ScanLimit
}

func (s *CreatStockOssCheckTaskRequest) GetScanNoFileType() *bool {
	return s.ScanNoFileType
}

func (s *CreatStockOssCheckTaskRequest) GetScanResourceType() *string {
	return s.ScanResourceType
}

func (s *CreatStockOssCheckTaskRequest) GetScanService() *string {
	return s.ScanService
}

func (s *CreatStockOssCheckTaskRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreatStockOssCheckTaskRequest) GetTaskCycle() *int32 {
	return s.TaskCycle
}

func (s *CreatStockOssCheckTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreatStockOssCheckTaskRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *CreatStockOssCheckTaskRequest) SetBuckets(v string) *CreatStockOssCheckTaskRequest {
	s.Buckets = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetCallbackId(v string) *CreatStockOssCheckTaskRequest {
	s.CallbackId = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetDistinctHistoryTasks(v bool) *CreatStockOssCheckTaskRequest {
	s.DistinctHistoryTasks = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetEndTime(v string) *CreatStockOssCheckTaskRequest {
	s.EndTime = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetExecuteDate(v int32) *CreatStockOssCheckTaskRequest {
	s.ExecuteDate = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetExecuteTime(v string) *CreatStockOssCheckTaskRequest {
	s.ExecuteTime = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetFreeze(v bool) *CreatStockOssCheckTaskRequest {
	s.Freeze = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetFreezeHighRisk1(v bool) *CreatStockOssCheckTaskRequest {
	s.FreezeHighRisk1 = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetFreezeHighRisk2(v bool) *CreatStockOssCheckTaskRequest {
	s.FreezeHighRisk2 = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetFreezeMediumRisk1(v bool) *CreatStockOssCheckTaskRequest {
	s.FreezeMediumRisk1 = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetFreezeMediumRisk2(v bool) *CreatStockOssCheckTaskRequest {
	s.FreezeMediumRisk2 = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetFreezeRestorePath(v string) *CreatStockOssCheckTaskRequest {
	s.FreezeRestorePath = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetFreezeType(v string) *CreatStockOssCheckTaskRequest {
	s.FreezeType = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetIsInc(v bool) *CreatStockOssCheckTaskRequest {
	s.IsInc = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetMediaType(v int32) *CreatStockOssCheckTaskRequest {
	s.MediaType = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetPrefixFilterType(v string) *CreatStockOssCheckTaskRequest {
	s.PrefixFilterType = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetPrefixFilters(v string) *CreatStockOssCheckTaskRequest {
	s.PrefixFilters = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetPriority(v int32) *CreatStockOssCheckTaskRequest {
	s.Priority = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetReferer(v string) *CreatStockOssCheckTaskRequest {
	s.Referer = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetRegionId(v string) *CreatStockOssCheckTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetScanLimit(v int64) *CreatStockOssCheckTaskRequest {
	s.ScanLimit = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetScanNoFileType(v bool) *CreatStockOssCheckTaskRequest {
	s.ScanNoFileType = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetScanResourceType(v string) *CreatStockOssCheckTaskRequest {
	s.ScanResourceType = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetScanService(v string) *CreatStockOssCheckTaskRequest {
	s.ScanService = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetStartTime(v string) *CreatStockOssCheckTaskRequest {
	s.StartTime = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetTaskCycle(v int32) *CreatStockOssCheckTaskRequest {
	s.TaskCycle = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetTaskName(v string) *CreatStockOssCheckTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetTaskType(v string) *CreatStockOssCheckTaskRequest {
	s.TaskType = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) Validate() error {
	return dara.Validate(s)
}
