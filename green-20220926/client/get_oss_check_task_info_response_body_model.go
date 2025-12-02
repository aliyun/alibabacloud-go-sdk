// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssCheckTaskInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBuckets(v string) *GetOssCheckTaskInfoResponseBody
	GetBuckets() *string
	SetConfig(v *GetOssCheckTaskInfoResponseBodyConfig) *GetOssCheckTaskInfoResponseBody
	GetConfig() *GetOssCheckTaskInfoResponseBodyConfig
	SetEndTime(v string) *GetOssCheckTaskInfoResponseBody
	GetEndTime() *string
	SetFinishNum(v int64) *GetOssCheckTaskInfoResponseBody
	GetFinishNum() *int64
	SetIsInc(v bool) *GetOssCheckTaskInfoResponseBody
	GetIsInc() *bool
	SetLastExecuteDate(v string) *GetOssCheckTaskInfoResponseBody
	GetLastExecuteDate() *string
	SetMediaType(v int32) *GetOssCheckTaskInfoResponseBody
	GetMediaType() *int32
	SetNextExecuteDate(v string) *GetOssCheckTaskInfoResponseBody
	GetNextExecuteDate() *string
	SetObjectNum(v int64) *GetOssCheckTaskInfoResponseBody
	GetObjectNum() *int64
	SetRequestId(v string) *GetOssCheckTaskInfoResponseBody
	GetRequestId() *string
	SetSearchNum(v int64) *GetOssCheckTaskInfoResponseBody
	GetSearchNum() *int64
	SetStartTime(v string) *GetOssCheckTaskInfoResponseBody
	GetStartTime() *string
	SetStatus(v int32) *GetOssCheckTaskInfoResponseBody
	GetStatus() *int32
	SetTaskId(v string) *GetOssCheckTaskInfoResponseBody
	GetTaskId() *string
	SetTaskName(v string) *GetOssCheckTaskInfoResponseBody
	GetTaskName() *string
	SetTaskType(v string) *GetOssCheckTaskInfoResponseBody
	GetTaskType() *string
}

type GetOssCheckTaskInfoResponseBody struct {
	// example:
	//
	// [{\\"Bucket\\":\\"aileshijie\\",\\"Region\\":\\"cn-hangzhou\\"}]
	Buckets *string                                `json:"Buckets,omitempty" xml:"Buckets,omitempty"`
	Config  *GetOssCheckTaskInfoResponseBodyConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// example:
	//
	// 2025-07-09 10:30:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 56
	FinishNum *int64 `json:"FinishNum,omitempty" xml:"FinishNum,omitempty"`
	// example:
	//
	// false
	IsInc *bool `json:"IsInc,omitempty" xml:"IsInc,omitempty"`
	// example:
	//
	// 2025-07-09 10:30:00
	LastExecuteDate *string `json:"LastExecuteDate,omitempty" xml:"LastExecuteDate,omitempty"`
	// example:
	//
	// 1
	MediaType *int32 `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// example:
	//
	// 2025-07-08 10:30:00
	NextExecuteDate *string `json:"NextExecuteDate,omitempty" xml:"NextExecuteDate,omitempty"`
	// example:
	//
	// 100
	ObjectNum *int64 `json:"ObjectNum,omitempty" xml:"ObjectNum,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	SearchNum *int64 `json:"SearchNum,omitempty" xml:"SearchNum,omitempty"`
	// example:
	//
	// 2023-08-21 16:08:38
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Success
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// xxxx-xxx
	TaskId   *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// increment
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetOssCheckTaskInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOssCheckTaskInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetOssCheckTaskInfoResponseBody) GetBuckets() *string {
	return s.Buckets
}

func (s *GetOssCheckTaskInfoResponseBody) GetConfig() *GetOssCheckTaskInfoResponseBodyConfig {
	return s.Config
}

func (s *GetOssCheckTaskInfoResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *GetOssCheckTaskInfoResponseBody) GetFinishNum() *int64 {
	return s.FinishNum
}

func (s *GetOssCheckTaskInfoResponseBody) GetIsInc() *bool {
	return s.IsInc
}

func (s *GetOssCheckTaskInfoResponseBody) GetLastExecuteDate() *string {
	return s.LastExecuteDate
}

func (s *GetOssCheckTaskInfoResponseBody) GetMediaType() *int32 {
	return s.MediaType
}

func (s *GetOssCheckTaskInfoResponseBody) GetNextExecuteDate() *string {
	return s.NextExecuteDate
}

func (s *GetOssCheckTaskInfoResponseBody) GetObjectNum() *int64 {
	return s.ObjectNum
}

func (s *GetOssCheckTaskInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOssCheckTaskInfoResponseBody) GetSearchNum() *int64 {
	return s.SearchNum
}

func (s *GetOssCheckTaskInfoResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *GetOssCheckTaskInfoResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *GetOssCheckTaskInfoResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *GetOssCheckTaskInfoResponseBody) GetTaskName() *string {
	return s.TaskName
}

func (s *GetOssCheckTaskInfoResponseBody) GetTaskType() *string {
	return s.TaskType
}

func (s *GetOssCheckTaskInfoResponseBody) SetBuckets(v string) *GetOssCheckTaskInfoResponseBody {
	s.Buckets = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBody) SetConfig(v *GetOssCheckTaskInfoResponseBodyConfig) *GetOssCheckTaskInfoResponseBody {
	s.Config = v
	return s
}

func (s *GetOssCheckTaskInfoResponseBody) SetEndTime(v string) *GetOssCheckTaskInfoResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBody) SetFinishNum(v int64) *GetOssCheckTaskInfoResponseBody {
	s.FinishNum = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBody) SetIsInc(v bool) *GetOssCheckTaskInfoResponseBody {
	s.IsInc = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBody) SetLastExecuteDate(v string) *GetOssCheckTaskInfoResponseBody {
	s.LastExecuteDate = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBody) SetMediaType(v int32) *GetOssCheckTaskInfoResponseBody {
	s.MediaType = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBody) SetNextExecuteDate(v string) *GetOssCheckTaskInfoResponseBody {
	s.NextExecuteDate = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBody) SetObjectNum(v int64) *GetOssCheckTaskInfoResponseBody {
	s.ObjectNum = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBody) SetRequestId(v string) *GetOssCheckTaskInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBody) SetSearchNum(v int64) *GetOssCheckTaskInfoResponseBody {
	s.SearchNum = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBody) SetStartTime(v string) *GetOssCheckTaskInfoResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBody) SetStatus(v int32) *GetOssCheckTaskInfoResponseBody {
	s.Status = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBody) SetTaskId(v string) *GetOssCheckTaskInfoResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBody) SetTaskName(v string) *GetOssCheckTaskInfoResponseBody {
	s.TaskName = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBody) SetTaskType(v string) *GetOssCheckTaskInfoResponseBody {
	s.TaskType = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBody) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOssCheckTaskInfoResponseBodyConfig struct {
	// example:
	//
	// 188
	CallbackId *int64 `json:"CallbackId,omitempty" xml:"CallbackId,omitempty"`
	// example:
	//
	// true
	DistinctHistoryTasks *bool `json:"DistinctHistoryTasks,omitempty" xml:"DistinctHistoryTasks,omitempty"`
	// example:
	//
	// 2025-07-09 10:30:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	ExecuteDate *int32 `json:"ExecuteDate,omitempty" xml:"ExecuteDate,omitempty"`
	// example:
	//
	// 2025-07-09 10:30:00
	ExecuteTime *string `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	// example:
	//
	// true
	Freeze *bool `json:"Freeze,omitempty" xml:"Freeze,omitempty"`
	// example:
	//
	// true
	FreezeHighRisk1 *bool `json:"FreezeHighRisk1,omitempty" xml:"FreezeHighRisk1,omitempty"`
	// example:
	//
	// true
	FreezeHighRisk2 *bool `json:"FreezeHighRisk2,omitempty" xml:"FreezeHighRisk2,omitempty"`
	// example:
	//
	// true
	FreezeMediumRisk1 *bool `json:"FreezeMediumRisk1,omitempty" xml:"FreezeMediumRisk1,omitempty"`
	// example:
	//
	// false
	FreezeMediumRisk2 *bool `json:"FreezeMediumRisk2,omitempty" xml:"FreezeMediumRisk2,omitempty"`
	// example:
	//
	// test
	FreezeRestorePath *string `json:"FreezeRestorePath,omitempty" xml:"FreezeRestorePath,omitempty"`
	// example:
	//
	// ACL
	FreezeType *string `json:"FreezeType,omitempty" xml:"FreezeType,omitempty"`
	// example:
	//
	// all
	PrefixFilterType *string   `json:"PrefixFilterType,omitempty" xml:"PrefixFilterType,omitempty"`
	PrefixFilters    []*string `json:"PrefixFilters,omitempty" xml:"PrefixFilters,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// Referer。
	//
	// example:
	//
	// https://www.aliyun.com/
	Referer *string `json:"Referer,omitempty" xml:"Referer,omitempty"`
	// example:
	//
	// 100
	ScanLimit *int64 `json:"ScanLimit,omitempty" xml:"ScanLimit,omitempty"`
	// example:
	//
	// true
	ScanNoFileType *bool `json:"ScanNoFileType,omitempty" xml:"ScanNoFileType,omitempty"`
	// example:
	//
	// 0
	ScanResourceType *int32                                                   `json:"ScanResourceType,omitempty" xml:"ScanResourceType,omitempty"`
	ScanService      []*string                                                `json:"ScanService,omitempty" xml:"ScanService,omitempty" type:"Repeated"`
	ScanServiceInfos []*GetOssCheckTaskInfoResponseBodyConfigScanServiceInfos `json:"ScanServiceInfos,omitempty" xml:"ScanServiceInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 2023-08-21 16:08:38
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 1
	TaskCycle        *int32                                                 `json:"TaskCycle,omitempty" xml:"TaskCycle,omitempty"`
	UserFreezeConfig *GetOssCheckTaskInfoResponseBodyConfigUserFreezeConfig `json:"UserFreezeConfig,omitempty" xml:"UserFreezeConfig,omitempty" type:"Struct"`
}

func (s GetOssCheckTaskInfoResponseBodyConfig) String() string {
	return dara.Prettify(s)
}

func (s GetOssCheckTaskInfoResponseBodyConfig) GoString() string {
	return s.String()
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) GetCallbackId() *int64 {
	return s.CallbackId
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) GetDistinctHistoryTasks() *bool {
	return s.DistinctHistoryTasks
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) GetEndTime() *string {
	return s.EndTime
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) GetExecuteDate() *int32 {
	return s.ExecuteDate
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) GetExecuteTime() *string {
	return s.ExecuteTime
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) GetFreeze() *bool {
	return s.Freeze
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) GetFreezeHighRisk1() *bool {
	return s.FreezeHighRisk1
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) GetFreezeHighRisk2() *bool {
	return s.FreezeHighRisk2
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) GetFreezeMediumRisk1() *bool {
	return s.FreezeMediumRisk1
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) GetFreezeMediumRisk2() *bool {
	return s.FreezeMediumRisk2
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) GetFreezeRestorePath() *string {
	return s.FreezeRestorePath
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) GetFreezeType() *string {
	return s.FreezeType
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) GetPrefixFilterType() *string {
	return s.PrefixFilterType
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) GetPrefixFilters() []*string {
	return s.PrefixFilters
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) GetPriority() *int32 {
	return s.Priority
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) GetReferer() *string {
	return s.Referer
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) GetScanLimit() *int64 {
	return s.ScanLimit
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) GetScanNoFileType() *bool {
	return s.ScanNoFileType
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) GetScanResourceType() *int32 {
	return s.ScanResourceType
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) GetScanService() []*string {
	return s.ScanService
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) GetScanServiceInfos() []*GetOssCheckTaskInfoResponseBodyConfigScanServiceInfos {
	return s.ScanServiceInfos
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) GetStartTime() *string {
	return s.StartTime
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) GetTaskCycle() *int32 {
	return s.TaskCycle
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) GetUserFreezeConfig() *GetOssCheckTaskInfoResponseBodyConfigUserFreezeConfig {
	return s.UserFreezeConfig
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) SetCallbackId(v int64) *GetOssCheckTaskInfoResponseBodyConfig {
	s.CallbackId = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) SetDistinctHistoryTasks(v bool) *GetOssCheckTaskInfoResponseBodyConfig {
	s.DistinctHistoryTasks = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) SetEndTime(v string) *GetOssCheckTaskInfoResponseBodyConfig {
	s.EndTime = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) SetExecuteDate(v int32) *GetOssCheckTaskInfoResponseBodyConfig {
	s.ExecuteDate = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) SetExecuteTime(v string) *GetOssCheckTaskInfoResponseBodyConfig {
	s.ExecuteTime = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) SetFreeze(v bool) *GetOssCheckTaskInfoResponseBodyConfig {
	s.Freeze = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) SetFreezeHighRisk1(v bool) *GetOssCheckTaskInfoResponseBodyConfig {
	s.FreezeHighRisk1 = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) SetFreezeHighRisk2(v bool) *GetOssCheckTaskInfoResponseBodyConfig {
	s.FreezeHighRisk2 = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) SetFreezeMediumRisk1(v bool) *GetOssCheckTaskInfoResponseBodyConfig {
	s.FreezeMediumRisk1 = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) SetFreezeMediumRisk2(v bool) *GetOssCheckTaskInfoResponseBodyConfig {
	s.FreezeMediumRisk2 = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) SetFreezeRestorePath(v string) *GetOssCheckTaskInfoResponseBodyConfig {
	s.FreezeRestorePath = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) SetFreezeType(v string) *GetOssCheckTaskInfoResponseBodyConfig {
	s.FreezeType = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) SetPrefixFilterType(v string) *GetOssCheckTaskInfoResponseBodyConfig {
	s.PrefixFilterType = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) SetPrefixFilters(v []*string) *GetOssCheckTaskInfoResponseBodyConfig {
	s.PrefixFilters = v
	return s
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) SetPriority(v int32) *GetOssCheckTaskInfoResponseBodyConfig {
	s.Priority = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) SetReferer(v string) *GetOssCheckTaskInfoResponseBodyConfig {
	s.Referer = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) SetScanLimit(v int64) *GetOssCheckTaskInfoResponseBodyConfig {
	s.ScanLimit = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) SetScanNoFileType(v bool) *GetOssCheckTaskInfoResponseBodyConfig {
	s.ScanNoFileType = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) SetScanResourceType(v int32) *GetOssCheckTaskInfoResponseBodyConfig {
	s.ScanResourceType = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) SetScanService(v []*string) *GetOssCheckTaskInfoResponseBodyConfig {
	s.ScanService = v
	return s
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) SetScanServiceInfos(v []*GetOssCheckTaskInfoResponseBodyConfigScanServiceInfos) *GetOssCheckTaskInfoResponseBodyConfig {
	s.ScanServiceInfos = v
	return s
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) SetStartTime(v string) *GetOssCheckTaskInfoResponseBodyConfig {
	s.StartTime = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) SetTaskCycle(v int32) *GetOssCheckTaskInfoResponseBodyConfig {
	s.TaskCycle = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) SetUserFreezeConfig(v *GetOssCheckTaskInfoResponseBodyConfigUserFreezeConfig) *GetOssCheckTaskInfoResponseBodyConfig {
	s.UserFreezeConfig = v
	return s
}

func (s *GetOssCheckTaskInfoResponseBodyConfig) Validate() error {
	if s.ScanServiceInfos != nil {
		for _, item := range s.ScanServiceInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UserFreezeConfig != nil {
		if err := s.UserFreezeConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOssCheckTaskInfoResponseBodyConfigScanServiceInfos struct {
	// example:
	//
	// oss_baselineCheck
	CopyFrom *string `json:"CopyFrom,omitempty" xml:"CopyFrom,omitempty"`
	// example:
	//
	// false
	IsCopy *bool `json:"IsCopy,omitempty" xml:"IsCopy,omitempty"`
	// example:
	//
	// oss_baselineCheck
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s GetOssCheckTaskInfoResponseBodyConfigScanServiceInfos) String() string {
	return dara.Prettify(s)
}

func (s GetOssCheckTaskInfoResponseBodyConfigScanServiceInfos) GoString() string {
	return s.String()
}

func (s *GetOssCheckTaskInfoResponseBodyConfigScanServiceInfos) GetCopyFrom() *string {
	return s.CopyFrom
}

func (s *GetOssCheckTaskInfoResponseBodyConfigScanServiceInfos) GetIsCopy() *bool {
	return s.IsCopy
}

func (s *GetOssCheckTaskInfoResponseBodyConfigScanServiceInfos) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *GetOssCheckTaskInfoResponseBodyConfigScanServiceInfos) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetOssCheckTaskInfoResponseBodyConfigScanServiceInfos) SetCopyFrom(v string) *GetOssCheckTaskInfoResponseBodyConfigScanServiceInfos {
	s.CopyFrom = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBodyConfigScanServiceInfos) SetIsCopy(v bool) *GetOssCheckTaskInfoResponseBodyConfigScanServiceInfos {
	s.IsCopy = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBodyConfigScanServiceInfos) SetServiceCode(v string) *GetOssCheckTaskInfoResponseBodyConfigScanServiceInfos {
	s.ServiceCode = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBodyConfigScanServiceInfos) SetServiceName(v string) *GetOssCheckTaskInfoResponseBodyConfigScanServiceInfos {
	s.ServiceName = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBodyConfigScanServiceInfos) Validate() error {
	return dara.Validate(s)
}

type GetOssCheckTaskInfoResponseBodyConfigUserFreezeConfig struct {
	// example:
	//
	// test
	FreezeRestorePath *string `json:"FreezeRestorePath,omitempty" xml:"FreezeRestorePath,omitempty"`
	// example:
	//
	// ACL
	FreezeType *string `json:"FreezeType,omitempty" xml:"FreezeType,omitempty"`
}

func (s GetOssCheckTaskInfoResponseBodyConfigUserFreezeConfig) String() string {
	return dara.Prettify(s)
}

func (s GetOssCheckTaskInfoResponseBodyConfigUserFreezeConfig) GoString() string {
	return s.String()
}

func (s *GetOssCheckTaskInfoResponseBodyConfigUserFreezeConfig) GetFreezeRestorePath() *string {
	return s.FreezeRestorePath
}

func (s *GetOssCheckTaskInfoResponseBodyConfigUserFreezeConfig) GetFreezeType() *string {
	return s.FreezeType
}

func (s *GetOssCheckTaskInfoResponseBodyConfigUserFreezeConfig) SetFreezeRestorePath(v string) *GetOssCheckTaskInfoResponseBodyConfigUserFreezeConfig {
	s.FreezeRestorePath = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBodyConfigUserFreezeConfig) SetFreezeType(v string) *GetOssCheckTaskInfoResponseBodyConfigUserFreezeConfig {
	s.FreezeType = &v
	return s
}

func (s *GetOssCheckTaskInfoResponseBodyConfigUserFreezeConfig) Validate() error {
	return dara.Validate(s)
}
