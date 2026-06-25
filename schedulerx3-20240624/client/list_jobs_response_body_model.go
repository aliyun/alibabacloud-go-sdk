// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListJobsResponseBody
	GetCode() *int32
	SetData(v *ListJobsResponseBodyData) *ListJobsResponseBody
	GetData() *ListJobsResponseBodyData
	SetMessage(v string) *ListJobsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListJobsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListJobsResponseBody
	GetSuccess() *bool
}

type ListJobsResponseBody struct {
	// The HTTP status code. A value of `200` indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// - The returned data.
	Data *ListJobsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// A unique ID that Alibaba Cloud generates for each request. Use this ID to troubleshoot issues.
	//
	// example:
	//
	// 1DF6732E-15D8-5E1F-95E3-C10077F556B5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - `true`: The request was successful.
	//
	// - `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListJobsResponseBody) GetData() *ListJobsResponseBodyData {
	return s.Data
}

func (s *ListJobsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListJobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListJobsResponseBody) SetCode(v int32) *ListJobsResponseBody {
	s.Code = &v
	return s
}

func (s *ListJobsResponseBody) SetData(v *ListJobsResponseBodyData) *ListJobsResponseBody {
	s.Data = v
	return s
}

func (s *ListJobsResponseBody) SetMessage(v string) *ListJobsResponseBody {
	s.Message = &v
	return s
}

func (s *ListJobsResponseBody) SetRequestId(v string) *ListJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobsResponseBody) SetSuccess(v bool) *ListJobsResponseBody {
	s.Success = &v
	return s
}

func (s *ListJobsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListJobsResponseBodyData struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// - A list of jobs.
	Records []*ListJobsResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// The total number of jobs returned.
	//
	// example:
	//
	// 65
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListJobsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListJobsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListJobsResponseBodyData) GetRecords() []*ListJobsResponseBodyDataRecords {
	return s.Records
}

func (s *ListJobsResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListJobsResponseBodyData) SetPageNumber(v int32) *ListJobsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListJobsResponseBodyData) SetPageSize(v int32) *ListJobsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListJobsResponseBodyData) SetRecords(v []*ListJobsResponseBodyDataRecords) *ListJobsResponseBodyData {
	s.Records = v
	return s
}

func (s *ListJobsResponseBodyData) SetTotal(v int32) *ListJobsResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListJobsResponseBodyData) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListJobsResponseBodyDataRecords struct {
	// The application name.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The retry interval, in seconds.
	//
	// example:
	//
	// 30
	AttemptInterval *int32 `json:"AttemptInterval,omitempty" xml:"AttemptInterval,omitempty"`
	// The calendar.
	//
	// example:
	//
	// work-day
	Calendar *string `json:"Calendar,omitempty" xml:"Calendar,omitempty"`
	// The child job ID.
	//
	// example:
	//
	// 1246
	ChildJobId *string `json:"ChildJobId,omitempty" xml:"ChildJobId,omitempty"`
	// The cleanup mode.
	//
	// example:
	//
	// {"cleanMode":"NUM_ONLY","totalRemain":300}
	CleanMode *string `json:"CleanMode,omitempty" xml:"CleanMode,omitempty"`
	// The ID of the user who created the job.
	//
	// example:
	//
	// 1963096506470832
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The current execution status. Valid values:
	//
	// - `0`: Not Started
	//
	// - `1`: Running
	//
	// - `2`: Queued
	//
	// - `3`: Waiting
	//
	// example:
	//
	// 3
	CurrentExecuteStatus *int32 `json:"CurrentExecuteStatus,omitempty" xml:"CurrentExecuteStatus,omitempty"`
	// The data offset.
	//
	// example:
	//
	// 3
	DataOffset *int32 `json:"DataOffset,omitempty" xml:"DataOffset,omitempty"`
	// The dependency check strategy.
	//
	// example:
	//
	// 1
	DependentStrategy *int32 `json:"DependentStrategy,omitempty" xml:"DependentStrategy,omitempty"`
	// The job description.
	//
	// example:
	//
	// job01单机任务
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The executor blocking strategy. Valid values:
	//
	// - `1`: Serial Execution
	//
	// - `2`: Discard Later
	//
	// - `3`: Cover Earlier
	//
	// example:
	//
	// 1
	ExecutorBlockStrategy *string `json:"ExecutorBlockStrategy,omitempty" xml:"ExecutorBlockStrategy,omitempty"`
	// The name of the job handler.
	//
	// example:
	//
	// jobDemoHandler
	JobHandler *string `json:"JobHandler,omitempty" xml:"JobHandler,omitempty"`
	// The job ID.
	//
	// example:
	//
	// 74
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The job type.
	//
	// example:
	//
	// xxljob
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The time when the last execution ended.
	//
	// example:
	//
	// 2025-06-29 15:56:36
	LastExecuteEndTime *string `json:"LastExecuteEndTime,omitempty" xml:"LastExecuteEndTime,omitempty"`
	// The status of the last execution. Valid values:
	//
	// - `4`: Success
	//
	// - `5`: Failure
	//
	// example:
	//
	// 4
	LastExecuteStatus *int32 `json:"LastExecuteStatus,omitempty" xml:"LastExecuteStatus,omitempty"`
	// The maximum number of retries for a failed job.
	//
	// example:
	//
	// 5
	MaxAttempt *int32 `json:"MaxAttempt,omitempty" xml:"MaxAttempt,omitempty"`
	// The maximum concurrency.
	//
	// example:
	//
	// 100
	MaxConcurrency *int32 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// The job name.
	//
	// example:
	//
	// job01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The node type.
	//
	// example:
	//
	// 1
	NodeType *int32 `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The notification configuration.
	//
	// example:
	//
	// {"failLimitTimes":1,"failEnable":true,"timeoutKillEnable":false,"missWorkerEnable":true,"timeoutEnable":true,"sendChannel":"","timeout":300,"successNotice":false}
	NoticeConfig *string `json:"NoticeConfig,omitempty" xml:"NoticeConfig,omitempty"`
	// The notification contacts.
	//
	// example:
	//
	// [{"contactType":1,"name":"恰橙"}]
	NoticeContacts *string `json:"NoticeContacts,omitempty" xml:"NoticeContacts,omitempty"`
	// The job parameters.
	//
	// example:
	//
	// name=10
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The job priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The routing strategy. Valid values:
	//
	// - `1`: Round-robin
	//
	// - `2`: Random
	//
	// - `3`: First
	//
	// - `4`: Last
	//
	// - `5`: Least Frequently Used
	//
	// - `6`: Least Recently Used
	//
	// - `7`: Consistent Hashing
	//
	// - `8`: Sharded Broadcast
	//
	// example:
	//
	// 1
	RouteStrategy *int32 `json:"RouteStrategy,omitempty" xml:"RouteStrategy,omitempty"`
	// The script content.
	//
	// example:
	//
	// #!/bin/bash
	//
	// echo "xxl-job: hello shell"
	//
	// echo "Good bye!"
	//
	// exit 0
	Script *string `json:"Script,omitempty" xml:"Script,omitempty"`
	// The type of the start time.
	//
	// example:
	//
	// 1
	StartTimeType *int32 `json:"StartTimeType,omitempty" xml:"StartTimeType,omitempty"`
	// The job status. Valid values:
	//
	// - `0`: Disabled
	//
	// - `1`: Enabled
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time expression.
	//
	// example:
	//
	// 0 0 12 	- 	- ?
	TimeExpression *string `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	// The time type. Valid values:
	//
	// - `-1`: none
	//
	// - `1`: cron
	//
	// - `3`: fix_rate
	//
	// - `5`: one_time
	//
	// - `100`: api
	//
	// example:
	//
	// 1
	TimeType *int32 `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// The time zone.
	//
	// example:
	//
	// Hongkong
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
	// The time zone.
	//
	// example:
	//
	// HangKong
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	// The ID of the user who last updated the job.
	//
	// example:
	//
	// 1963096506470832
	Updater *string `json:"Updater,omitempty" xml:"Updater,omitempty"`
	// The job weight.
	//
	// example:
	//
	// 1
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
	// The workflow ID.
	//
	// example:
	//
	// 10
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
	// The extended attributes.
	//
	// > This parameter is not currently supported.
	//
	// example:
	//
	// 暂无
	Xattrs *string `json:"Xattrs,omitempty" xml:"Xattrs,omitempty"`
}

func (s ListJobsResponseBodyDataRecords) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyDataRecords) GetAppName() *string {
	return s.AppName
}

func (s *ListJobsResponseBodyDataRecords) GetAttemptInterval() *int32 {
	return s.AttemptInterval
}

func (s *ListJobsResponseBodyDataRecords) GetCalendar() *string {
	return s.Calendar
}

func (s *ListJobsResponseBodyDataRecords) GetChildJobId() *string {
	return s.ChildJobId
}

func (s *ListJobsResponseBodyDataRecords) GetCleanMode() *string {
	return s.CleanMode
}

func (s *ListJobsResponseBodyDataRecords) GetCreator() *string {
	return s.Creator
}

func (s *ListJobsResponseBodyDataRecords) GetCurrentExecuteStatus() *int32 {
	return s.CurrentExecuteStatus
}

func (s *ListJobsResponseBodyDataRecords) GetDataOffset() *int32 {
	return s.DataOffset
}

func (s *ListJobsResponseBodyDataRecords) GetDependentStrategy() *int32 {
	return s.DependentStrategy
}

func (s *ListJobsResponseBodyDataRecords) GetDescription() *string {
	return s.Description
}

func (s *ListJobsResponseBodyDataRecords) GetExecutorBlockStrategy() *string {
	return s.ExecutorBlockStrategy
}

func (s *ListJobsResponseBodyDataRecords) GetJobHandler() *string {
	return s.JobHandler
}

func (s *ListJobsResponseBodyDataRecords) GetJobId() *int64 {
	return s.JobId
}

func (s *ListJobsResponseBodyDataRecords) GetJobType() *string {
	return s.JobType
}

func (s *ListJobsResponseBodyDataRecords) GetLastExecuteEndTime() *string {
	return s.LastExecuteEndTime
}

func (s *ListJobsResponseBodyDataRecords) GetLastExecuteStatus() *int32 {
	return s.LastExecuteStatus
}

func (s *ListJobsResponseBodyDataRecords) GetMaxAttempt() *int32 {
	return s.MaxAttempt
}

func (s *ListJobsResponseBodyDataRecords) GetMaxConcurrency() *int32 {
	return s.MaxConcurrency
}

func (s *ListJobsResponseBodyDataRecords) GetName() *string {
	return s.Name
}

func (s *ListJobsResponseBodyDataRecords) GetNodeType() *int32 {
	return s.NodeType
}

func (s *ListJobsResponseBodyDataRecords) GetNoticeConfig() *string {
	return s.NoticeConfig
}

func (s *ListJobsResponseBodyDataRecords) GetNoticeContacts() *string {
	return s.NoticeContacts
}

func (s *ListJobsResponseBodyDataRecords) GetParameters() *string {
	return s.Parameters
}

func (s *ListJobsResponseBodyDataRecords) GetPriority() *int32 {
	return s.Priority
}

func (s *ListJobsResponseBodyDataRecords) GetRouteStrategy() *int32 {
	return s.RouteStrategy
}

func (s *ListJobsResponseBodyDataRecords) GetScript() *string {
	return s.Script
}

func (s *ListJobsResponseBodyDataRecords) GetStartTimeType() *int32 {
	return s.StartTimeType
}

func (s *ListJobsResponseBodyDataRecords) GetStatus() *int32 {
	return s.Status
}

func (s *ListJobsResponseBodyDataRecords) GetTimeExpression() *string {
	return s.TimeExpression
}

func (s *ListJobsResponseBodyDataRecords) GetTimeType() *int32 {
	return s.TimeType
}

func (s *ListJobsResponseBodyDataRecords) GetTimeZone() *string {
	return s.TimeZone
}

func (s *ListJobsResponseBodyDataRecords) GetTimezone() *string {
	return s.Timezone
}

func (s *ListJobsResponseBodyDataRecords) GetUpdater() *string {
	return s.Updater
}

func (s *ListJobsResponseBodyDataRecords) GetWeight() *int32 {
	return s.Weight
}

func (s *ListJobsResponseBodyDataRecords) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *ListJobsResponseBodyDataRecords) GetXattrs() *string {
	return s.Xattrs
}

func (s *ListJobsResponseBodyDataRecords) SetAppName(v string) *ListJobsResponseBodyDataRecords {
	s.AppName = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetAttemptInterval(v int32) *ListJobsResponseBodyDataRecords {
	s.AttemptInterval = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetCalendar(v string) *ListJobsResponseBodyDataRecords {
	s.Calendar = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetChildJobId(v string) *ListJobsResponseBodyDataRecords {
	s.ChildJobId = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetCleanMode(v string) *ListJobsResponseBodyDataRecords {
	s.CleanMode = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetCreator(v string) *ListJobsResponseBodyDataRecords {
	s.Creator = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetCurrentExecuteStatus(v int32) *ListJobsResponseBodyDataRecords {
	s.CurrentExecuteStatus = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetDataOffset(v int32) *ListJobsResponseBodyDataRecords {
	s.DataOffset = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetDependentStrategy(v int32) *ListJobsResponseBodyDataRecords {
	s.DependentStrategy = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetDescription(v string) *ListJobsResponseBodyDataRecords {
	s.Description = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetExecutorBlockStrategy(v string) *ListJobsResponseBodyDataRecords {
	s.ExecutorBlockStrategy = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetJobHandler(v string) *ListJobsResponseBodyDataRecords {
	s.JobHandler = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetJobId(v int64) *ListJobsResponseBodyDataRecords {
	s.JobId = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetJobType(v string) *ListJobsResponseBodyDataRecords {
	s.JobType = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetLastExecuteEndTime(v string) *ListJobsResponseBodyDataRecords {
	s.LastExecuteEndTime = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetLastExecuteStatus(v int32) *ListJobsResponseBodyDataRecords {
	s.LastExecuteStatus = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetMaxAttempt(v int32) *ListJobsResponseBodyDataRecords {
	s.MaxAttempt = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetMaxConcurrency(v int32) *ListJobsResponseBodyDataRecords {
	s.MaxConcurrency = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetName(v string) *ListJobsResponseBodyDataRecords {
	s.Name = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetNodeType(v int32) *ListJobsResponseBodyDataRecords {
	s.NodeType = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetNoticeConfig(v string) *ListJobsResponseBodyDataRecords {
	s.NoticeConfig = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetNoticeContacts(v string) *ListJobsResponseBodyDataRecords {
	s.NoticeContacts = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetParameters(v string) *ListJobsResponseBodyDataRecords {
	s.Parameters = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetPriority(v int32) *ListJobsResponseBodyDataRecords {
	s.Priority = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetRouteStrategy(v int32) *ListJobsResponseBodyDataRecords {
	s.RouteStrategy = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetScript(v string) *ListJobsResponseBodyDataRecords {
	s.Script = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetStartTimeType(v int32) *ListJobsResponseBodyDataRecords {
	s.StartTimeType = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetStatus(v int32) *ListJobsResponseBodyDataRecords {
	s.Status = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetTimeExpression(v string) *ListJobsResponseBodyDataRecords {
	s.TimeExpression = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetTimeType(v int32) *ListJobsResponseBodyDataRecords {
	s.TimeType = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetTimeZone(v string) *ListJobsResponseBodyDataRecords {
	s.TimeZone = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetTimezone(v string) *ListJobsResponseBodyDataRecords {
	s.Timezone = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetUpdater(v string) *ListJobsResponseBodyDataRecords {
	s.Updater = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetWeight(v int32) *ListJobsResponseBodyDataRecords {
	s.Weight = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetWorkflowId(v int64) *ListJobsResponseBodyDataRecords {
	s.WorkflowId = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) SetXattrs(v string) *ListJobsResponseBodyDataRecords {
	s.Xattrs = &v
	return s
}

func (s *ListJobsResponseBodyDataRecords) Validate() error {
	return dara.Validate(s)
}
