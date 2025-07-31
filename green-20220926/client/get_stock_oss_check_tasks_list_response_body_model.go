// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStockOssCheckTasksListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetStockOssCheckTasksListResponseBody
	GetCurrentPage() *int32
	SetItems(v []*GetStockOssCheckTasksListResponseBodyItems) *GetStockOssCheckTasksListResponseBody
	GetItems() []*GetStockOssCheckTasksListResponseBodyItems
	SetPageSize(v int32) *GetStockOssCheckTasksListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetStockOssCheckTasksListResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *GetStockOssCheckTasksListResponseBody
	GetTotalCount() *int64
}

type GetStockOssCheckTasksListResponseBody struct {
	// example:
	//
	// 1
	CurrentPage *int32                                        `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Items       []*GetStockOssCheckTasksListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetStockOssCheckTasksListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStockOssCheckTasksListResponseBody) GoString() string {
	return s.String()
}

func (s *GetStockOssCheckTasksListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetStockOssCheckTasksListResponseBody) GetItems() []*GetStockOssCheckTasksListResponseBodyItems {
	return s.Items
}

func (s *GetStockOssCheckTasksListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetStockOssCheckTasksListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStockOssCheckTasksListResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetStockOssCheckTasksListResponseBody) SetCurrentPage(v int32) *GetStockOssCheckTasksListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBody) SetItems(v []*GetStockOssCheckTasksListResponseBodyItems) *GetStockOssCheckTasksListResponseBody {
	s.Items = v
	return s
}

func (s *GetStockOssCheckTasksListResponseBody) SetPageSize(v int32) *GetStockOssCheckTasksListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBody) SetRequestId(v string) *GetStockOssCheckTasksListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBody) SetTotalCount(v int64) *GetStockOssCheckTasksListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetStockOssCheckTasksListResponseBodyItems struct {
	// example:
	//
	// tmp
	Buckets *string                                           `json:"Buckets,omitempty" xml:"Buckets,omitempty"`
	Config  *GetStockOssCheckTasksListResponseBodyItemsConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// example:
	//
	// 2024-01-10 11:42:31
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 2
	FinishNum *int64 `json:"FinishNum,omitempty" xml:"FinishNum,omitempty"`
	// example:
	//
	// false
	IsInc *bool `json:"IsInc,omitempty" xml:"IsInc,omitempty"`
	// example:
	//
	// 02:00:00
	LastExecuteDate *string `json:"LastExecuteDate,omitempty" xml:"LastExecuteDate,omitempty"`
	// example:
	//
	// video
	MediaType *int32 `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// example:
	//
	// 02:00:00
	NextExecuteDate *string `json:"NextExecuteDate,omitempty" xml:"NextExecuteDate,omitempty"`
	// example:
	//
	// 10
	ObjectNum *int64 `json:"ObjectNum,omitempty" xml:"ObjectNum,omitempty"`
	// example:
	//
	// 10
	SearchNum *int64 `json:"SearchNum,omitempty" xml:"SearchNum,omitempty"`
	// example:
	//
	// 2023-12-21 15:30:19
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 4
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// P_XHDUS
	TaskId   *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// batch
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetStockOssCheckTasksListResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s GetStockOssCheckTasksListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *GetStockOssCheckTasksListResponseBodyItems) GetBuckets() *string {
	return s.Buckets
}

func (s *GetStockOssCheckTasksListResponseBodyItems) GetConfig() *GetStockOssCheckTasksListResponseBodyItemsConfig {
	return s.Config
}

func (s *GetStockOssCheckTasksListResponseBodyItems) GetEndTime() *string {
	return s.EndTime
}

func (s *GetStockOssCheckTasksListResponseBodyItems) GetFinishNum() *int64 {
	return s.FinishNum
}

func (s *GetStockOssCheckTasksListResponseBodyItems) GetIsInc() *bool {
	return s.IsInc
}

func (s *GetStockOssCheckTasksListResponseBodyItems) GetLastExecuteDate() *string {
	return s.LastExecuteDate
}

func (s *GetStockOssCheckTasksListResponseBodyItems) GetMediaType() *int32 {
	return s.MediaType
}

func (s *GetStockOssCheckTasksListResponseBodyItems) GetNextExecuteDate() *string {
	return s.NextExecuteDate
}

func (s *GetStockOssCheckTasksListResponseBodyItems) GetObjectNum() *int64 {
	return s.ObjectNum
}

func (s *GetStockOssCheckTasksListResponseBodyItems) GetSearchNum() *int64 {
	return s.SearchNum
}

func (s *GetStockOssCheckTasksListResponseBodyItems) GetStartTime() *string {
	return s.StartTime
}

func (s *GetStockOssCheckTasksListResponseBodyItems) GetStatus() *int32 {
	return s.Status
}

func (s *GetStockOssCheckTasksListResponseBodyItems) GetTaskId() *string {
	return s.TaskId
}

func (s *GetStockOssCheckTasksListResponseBodyItems) GetTaskName() *string {
	return s.TaskName
}

func (s *GetStockOssCheckTasksListResponseBodyItems) GetTaskType() *string {
	return s.TaskType
}

func (s *GetStockOssCheckTasksListResponseBodyItems) SetBuckets(v string) *GetStockOssCheckTasksListResponseBodyItems {
	s.Buckets = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItems) SetConfig(v *GetStockOssCheckTasksListResponseBodyItemsConfig) *GetStockOssCheckTasksListResponseBodyItems {
	s.Config = v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItems) SetEndTime(v string) *GetStockOssCheckTasksListResponseBodyItems {
	s.EndTime = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItems) SetFinishNum(v int64) *GetStockOssCheckTasksListResponseBodyItems {
	s.FinishNum = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItems) SetIsInc(v bool) *GetStockOssCheckTasksListResponseBodyItems {
	s.IsInc = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItems) SetLastExecuteDate(v string) *GetStockOssCheckTasksListResponseBodyItems {
	s.LastExecuteDate = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItems) SetMediaType(v int32) *GetStockOssCheckTasksListResponseBodyItems {
	s.MediaType = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItems) SetNextExecuteDate(v string) *GetStockOssCheckTasksListResponseBodyItems {
	s.NextExecuteDate = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItems) SetObjectNum(v int64) *GetStockOssCheckTasksListResponseBodyItems {
	s.ObjectNum = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItems) SetSearchNum(v int64) *GetStockOssCheckTasksListResponseBodyItems {
	s.SearchNum = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItems) SetStartTime(v string) *GetStockOssCheckTasksListResponseBodyItems {
	s.StartTime = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItems) SetStatus(v int32) *GetStockOssCheckTasksListResponseBodyItems {
	s.Status = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItems) SetTaskId(v string) *GetStockOssCheckTasksListResponseBodyItems {
	s.TaskId = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItems) SetTaskName(v string) *GetStockOssCheckTasksListResponseBodyItems {
	s.TaskName = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItems) SetTaskType(v string) *GetStockOssCheckTasksListResponseBodyItems {
	s.TaskType = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItems) Validate() error {
	return dara.Validate(s)
}

type GetStockOssCheckTasksListResponseBodyItemsConfig struct {
	CallbackId *int64 `json:"CallbackId,omitempty" xml:"CallbackId,omitempty"`
	// example:
	//
	// false
	DistinctHistoryTasks *bool `json:"DistinctHistoryTasks,omitempty" xml:"DistinctHistoryTasks,omitempty"`
	// example:
	//
	// 2024-01-10 11:42:31
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	ExecuteDate *int32 `json:"ExecuteDate,omitempty" xml:"ExecuteDate,omitempty"`
	// example:
	//
	// 02:00:00
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
	// all
	PrefixFilterType *string   `json:"PrefixFilterType,omitempty" xml:"PrefixFilterType,omitempty"`
	PrefixFilters    []*string `json:"PrefixFilters,omitempty" xml:"PrefixFilters,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Priority *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Referer  *string `json:"Referer,omitempty" xml:"Referer,omitempty"`
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
	ScanResourceType *int32                                                              `json:"ScanResourceType,omitempty" xml:"ScanResourceType,omitempty"`
	ScanService      []*string                                                           `json:"ScanService,omitempty" xml:"ScanService,omitempty" type:"Repeated"`
	ScanServiceInfos []*GetStockOssCheckTasksListResponseBodyItemsConfigScanServiceInfos `json:"ScanServiceInfos,omitempty" xml:"ScanServiceInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 2023-12-21 15:30:19
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 0
	TaskCycle        *int32                                                            `json:"TaskCycle,omitempty" xml:"TaskCycle,omitempty"`
	UserFreezeConfig *GetStockOssCheckTasksListResponseBodyItemsConfigUserFreezeConfig `json:"UserFreezeConfig,omitempty" xml:"UserFreezeConfig,omitempty" type:"Struct"`
}

func (s GetStockOssCheckTasksListResponseBodyItemsConfig) String() string {
	return dara.Prettify(s)
}

func (s GetStockOssCheckTasksListResponseBodyItemsConfig) GoString() string {
	return s.String()
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) GetCallbackId() *int64 {
	return s.CallbackId
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) GetDistinctHistoryTasks() *bool {
	return s.DistinctHistoryTasks
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) GetEndTime() *string {
	return s.EndTime
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) GetExecuteDate() *int32 {
	return s.ExecuteDate
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) GetExecuteTime() *string {
	return s.ExecuteTime
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) GetFreeze() *bool {
	return s.Freeze
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) GetFreezeHighRisk1() *bool {
	return s.FreezeHighRisk1
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) GetFreezeHighRisk2() *bool {
	return s.FreezeHighRisk2
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) GetFreezeMediumRisk1() *bool {
	return s.FreezeMediumRisk1
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) GetFreezeMediumRisk2() *bool {
	return s.FreezeMediumRisk2
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) GetFreezeRestorePath() *string {
	return s.FreezeRestorePath
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) GetFreezeType() *string {
	return s.FreezeType
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) GetPrefixFilterType() *string {
	return s.PrefixFilterType
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) GetPrefixFilters() []*string {
	return s.PrefixFilters
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) GetPriority() *int32 {
	return s.Priority
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) GetReferer() *string {
	return s.Referer
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) GetScanLimit() *int64 {
	return s.ScanLimit
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) GetScanNoFileType() *bool {
	return s.ScanNoFileType
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) GetScanResourceType() *int32 {
	return s.ScanResourceType
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) GetScanService() []*string {
	return s.ScanService
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) GetScanServiceInfos() []*GetStockOssCheckTasksListResponseBodyItemsConfigScanServiceInfos {
	return s.ScanServiceInfos
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) GetStartTime() *string {
	return s.StartTime
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) GetTaskCycle() *int32 {
	return s.TaskCycle
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) GetUserFreezeConfig() *GetStockOssCheckTasksListResponseBodyItemsConfigUserFreezeConfig {
	return s.UserFreezeConfig
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetCallbackId(v int64) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.CallbackId = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetDistinctHistoryTasks(v bool) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.DistinctHistoryTasks = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetEndTime(v string) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.EndTime = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetExecuteDate(v int32) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.ExecuteDate = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetExecuteTime(v string) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.ExecuteTime = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetFreeze(v bool) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.Freeze = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetFreezeHighRisk1(v bool) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.FreezeHighRisk1 = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetFreezeHighRisk2(v bool) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.FreezeHighRisk2 = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetFreezeMediumRisk1(v bool) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.FreezeMediumRisk1 = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetFreezeMediumRisk2(v bool) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.FreezeMediumRisk2 = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetFreezeRestorePath(v string) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.FreezeRestorePath = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetFreezeType(v string) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.FreezeType = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetPrefixFilterType(v string) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.PrefixFilterType = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetPrefixFilters(v []*string) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.PrefixFilters = v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetPriority(v int32) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.Priority = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetReferer(v string) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.Referer = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetScanLimit(v int64) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.ScanLimit = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetScanNoFileType(v bool) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.ScanNoFileType = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetScanResourceType(v int32) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.ScanResourceType = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetScanService(v []*string) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.ScanService = v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetScanServiceInfos(v []*GetStockOssCheckTasksListResponseBodyItemsConfigScanServiceInfos) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.ScanServiceInfos = v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetStartTime(v string) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.StartTime = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetTaskCycle(v int32) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.TaskCycle = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetUserFreezeConfig(v *GetStockOssCheckTasksListResponseBodyItemsConfigUserFreezeConfig) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.UserFreezeConfig = v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) Validate() error {
	return dara.Validate(s)
}

type GetStockOssCheckTasksListResponseBodyItemsConfigScanServiceInfos struct {
	// example:
	//
	// baselineCheck
	CopyFrom *string `json:"CopyFrom,omitempty" xml:"CopyFrom,omitempty"`
	// example:
	//
	// false
	IsCopy *bool `json:"IsCopy,omitempty" xml:"IsCopy,omitempty"`
	// example:
	//
	// baselineCheck_01
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s GetStockOssCheckTasksListResponseBodyItemsConfigScanServiceInfos) String() string {
	return dara.Prettify(s)
}

func (s GetStockOssCheckTasksListResponseBodyItemsConfigScanServiceInfos) GoString() string {
	return s.String()
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfigScanServiceInfos) GetCopyFrom() *string {
	return s.CopyFrom
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfigScanServiceInfos) GetIsCopy() *bool {
	return s.IsCopy
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfigScanServiceInfos) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfigScanServiceInfos) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfigScanServiceInfos) SetCopyFrom(v string) *GetStockOssCheckTasksListResponseBodyItemsConfigScanServiceInfos {
	s.CopyFrom = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfigScanServiceInfos) SetIsCopy(v bool) *GetStockOssCheckTasksListResponseBodyItemsConfigScanServiceInfos {
	s.IsCopy = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfigScanServiceInfos) SetServiceCode(v string) *GetStockOssCheckTasksListResponseBodyItemsConfigScanServiceInfos {
	s.ServiceCode = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfigScanServiceInfos) SetServiceName(v string) *GetStockOssCheckTasksListResponseBodyItemsConfigScanServiceInfos {
	s.ServiceName = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfigScanServiceInfos) Validate() error {
	return dara.Validate(s)
}

type GetStockOssCheckTasksListResponseBodyItemsConfigUserFreezeConfig struct {
	FreezeRestorePath *string `json:"FreezeRestorePath,omitempty" xml:"FreezeRestorePath,omitempty"`
	FreezeType        *string `json:"FreezeType,omitempty" xml:"FreezeType,omitempty"`
}

func (s GetStockOssCheckTasksListResponseBodyItemsConfigUserFreezeConfig) String() string {
	return dara.Prettify(s)
}

func (s GetStockOssCheckTasksListResponseBodyItemsConfigUserFreezeConfig) GoString() string {
	return s.String()
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfigUserFreezeConfig) GetFreezeRestorePath() *string {
	return s.FreezeRestorePath
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfigUserFreezeConfig) GetFreezeType() *string {
	return s.FreezeType
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfigUserFreezeConfig) SetFreezeRestorePath(v string) *GetStockOssCheckTasksListResponseBodyItemsConfigUserFreezeConfig {
	s.FreezeRestorePath = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfigUserFreezeConfig) SetFreezeType(v string) *GetStockOssCheckTasksListResponseBodyItemsConfigUserFreezeConfig {
	s.FreezeType = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfigUserFreezeConfig) Validate() error {
	return dara.Validate(s)
}
