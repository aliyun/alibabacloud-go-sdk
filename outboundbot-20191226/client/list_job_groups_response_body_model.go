// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncTaskId(v string) *ListJobGroupsResponseBody
	GetAsyncTaskId() *string
	SetCode(v string) *ListJobGroupsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListJobGroupsResponseBody
	GetHttpStatusCode() *int32
	SetJobGroups(v *ListJobGroupsResponseBodyJobGroups) *ListJobGroupsResponseBody
	GetJobGroups() *ListJobGroupsResponseBodyJobGroups
	SetMessage(v string) *ListJobGroupsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListJobGroupsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListJobGroupsResponseBody
	GetSuccess() *bool
}

type ListJobGroupsResponseBody struct {
	// The ID of the asynchronous task. You can use this ID to query the status of the task.
	//
	// example:
	//
	// 6243d904-939d-42ce-a8e4-886a139e77a3
	AsyncTaskId *string `json:"AsyncTaskId,omitempty" xml:"AsyncTaskId,omitempty"`
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The list of job groups.
	JobGroups *ListJobGroupsResponseBodyJobGroups `json:"JobGroups,omitempty" xml:"JobGroups,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: `true` and `false`.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListJobGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListJobGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobGroupsResponseBody) GetAsyncTaskId() *string {
	return s.AsyncTaskId
}

func (s *ListJobGroupsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListJobGroupsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListJobGroupsResponseBody) GetJobGroups() *ListJobGroupsResponseBodyJobGroups {
	return s.JobGroups
}

func (s *ListJobGroupsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListJobGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListJobGroupsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListJobGroupsResponseBody) SetAsyncTaskId(v string) *ListJobGroupsResponseBody {
	s.AsyncTaskId = &v
	return s
}

func (s *ListJobGroupsResponseBody) SetCode(v string) *ListJobGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *ListJobGroupsResponseBody) SetHttpStatusCode(v int32) *ListJobGroupsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListJobGroupsResponseBody) SetJobGroups(v *ListJobGroupsResponseBodyJobGroups) *ListJobGroupsResponseBody {
	s.JobGroups = v
	return s
}

func (s *ListJobGroupsResponseBody) SetMessage(v string) *ListJobGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *ListJobGroupsResponseBody) SetRequestId(v string) *ListJobGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobGroupsResponseBody) SetSuccess(v bool) *ListJobGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *ListJobGroupsResponseBody) Validate() error {
	if s.JobGroups != nil {
		if err := s.JobGroups.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListJobGroupsResponseBodyJobGroups struct {
	// The list of job groups.
	List []*ListJobGroupsResponseBodyJobGroupsList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListJobGroupsResponseBodyJobGroups) String() string {
	return dara.Prettify(s)
}

func (s ListJobGroupsResponseBodyJobGroups) GoString() string {
	return s.String()
}

func (s *ListJobGroupsResponseBodyJobGroups) GetList() []*ListJobGroupsResponseBodyJobGroupsList {
	return s.List
}

func (s *ListJobGroupsResponseBodyJobGroups) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListJobGroupsResponseBodyJobGroups) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListJobGroupsResponseBodyJobGroups) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListJobGroupsResponseBodyJobGroups) SetList(v []*ListJobGroupsResponseBodyJobGroupsList) *ListJobGroupsResponseBodyJobGroups {
	s.List = v
	return s
}

func (s *ListJobGroupsResponseBodyJobGroups) SetPageNumber(v int32) *ListJobGroupsResponseBodyJobGroups {
	s.PageNumber = &v
	return s
}

func (s *ListJobGroupsResponseBodyJobGroups) SetPageSize(v int32) *ListJobGroupsResponseBodyJobGroups {
	s.PageSize = &v
	return s
}

func (s *ListJobGroupsResponseBodyJobGroups) SetTotalCount(v int32) *ListJobGroupsResponseBodyJobGroups {
	s.TotalCount = &v
	return s
}

func (s *ListJobGroupsResponseBodyJobGroups) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListJobGroupsResponseBodyJobGroupsList struct {
	// The time when the job group was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1578550074361
	CreationTime *int64 `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The progress of the export task.
	//
	// example:
	//
	// {}
	ExportProgress *ListJobGroupsResponseBodyJobGroupsListExportProgress `json:"ExportProgress,omitempty" xml:"ExportProgress,omitempty" type:"Struct"`
	// The ID of the task that parses the job data file. This parameter is deprecated.
	//
	// example:
	//
	// c62e6789-28a8-41db-941e-171a01d3b3b9
	JobDataParsingTaskId *string `json:"JobDataParsingTaskId,omitempty" xml:"JobDataParsingTaskId,omitempty"`
	// The description of the job group.
	//
	// example:
	//
	// 催收的作业组
	JobGroupDescription *string `json:"JobGroupDescription,omitempty" xml:"JobGroupDescription,omitempty"`
	// The ID of the job group.
	//
	// example:
	//
	// c62e6789-28a8-41db-941e-171a01d3b3b9
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	// The name of the job group.
	//
	// example:
	//
	// 催收作业组
	JobGroupName *string `json:"JobGroupName,omitempty" xml:"JobGroupName,omitempty"`
	// The minimum number of concurrent calls.
	//
	// example:
	//
	// 1
	MinConcurrency *int32 `json:"MinConcurrency,omitempty" xml:"MinConcurrency,omitempty"`
	// The time when the job group was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1578550074361
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The progress of the job group.
	//
	// example:
	//
	// {}
	Progress *ListJobGroupsResponseBodyJobGroupsListProgress `json:"Progress,omitempty" xml:"Progress,omitempty" type:"Struct"`
	// The ID of the script.
	//
	// example:
	//
	// c62e6789-28a8-41db-941e-171a01d3b3b9
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// The name of the script.
	//
	// example:
	//
	// 话术名称
	ScriptName *string `json:"ScriptName,omitempty" xml:"ScriptName,omitempty"`
	// The script version.
	//
	// example:
	//
	// d9e828ac-744b-4dd3-848a-17a3da9167b8
	ScriptVersion *string `json:"ScriptVersion,omitempty" xml:"ScriptVersion,omitempty"`
	// The status of the job group. Valid values:
	//
	// - **Draft**: The job group is a draft.
	//
	// - **Scheduling**: The job group is being scheduled.
	//
	// - **Executing**: The job group is running.
	//
	// - **Completed**: The job group is completed.
	//
	// - **Paused**: The job group is paused.
	//
	// - **Failed**: The job group failed.
	//
	// - **Cancelled**: The job group is canceled.
	//
	// - **Initializing**: The job group is being initialized.
	//
	// example:
	//
	// Draft
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The calling strategy. This parameter is deprecated.
	//
	// > To view the strategy for a job group, call the `DescribeJobGroup` operation.
	//
	// example:
	//
	// {}
	Strategy *ListJobGroupsResponseBodyJobGroupsListStrategy `json:"Strategy,omitempty" xml:"Strategy,omitempty" type:"Struct"`
	// The total number of calls.
	//
	// example:
	//
	// 10
	TotalCallNum *int32 `json:"TotalCallNum,omitempty" xml:"TotalCallNum,omitempty"`
}

func (s ListJobGroupsResponseBodyJobGroupsList) String() string {
	return dara.Prettify(s)
}

func (s ListJobGroupsResponseBodyJobGroupsList) GoString() string {
	return s.String()
}

func (s *ListJobGroupsResponseBodyJobGroupsList) GetCreationTime() *int64 {
	return s.CreationTime
}

func (s *ListJobGroupsResponseBodyJobGroupsList) GetExportProgress() *ListJobGroupsResponseBodyJobGroupsListExportProgress {
	return s.ExportProgress
}

func (s *ListJobGroupsResponseBodyJobGroupsList) GetJobDataParsingTaskId() *string {
	return s.JobDataParsingTaskId
}

func (s *ListJobGroupsResponseBodyJobGroupsList) GetJobGroupDescription() *string {
	return s.JobGroupDescription
}

func (s *ListJobGroupsResponseBodyJobGroupsList) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *ListJobGroupsResponseBodyJobGroupsList) GetJobGroupName() *string {
	return s.JobGroupName
}

func (s *ListJobGroupsResponseBodyJobGroupsList) GetMinConcurrency() *int32 {
	return s.MinConcurrency
}

func (s *ListJobGroupsResponseBodyJobGroupsList) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListJobGroupsResponseBodyJobGroupsList) GetProgress() *ListJobGroupsResponseBodyJobGroupsListProgress {
	return s.Progress
}

func (s *ListJobGroupsResponseBodyJobGroupsList) GetScriptId() *string {
	return s.ScriptId
}

func (s *ListJobGroupsResponseBodyJobGroupsList) GetScriptName() *string {
	return s.ScriptName
}

func (s *ListJobGroupsResponseBodyJobGroupsList) GetScriptVersion() *string {
	return s.ScriptVersion
}

func (s *ListJobGroupsResponseBodyJobGroupsList) GetStatus() *string {
	return s.Status
}

func (s *ListJobGroupsResponseBodyJobGroupsList) GetStrategy() *ListJobGroupsResponseBodyJobGroupsListStrategy {
	return s.Strategy
}

func (s *ListJobGroupsResponseBodyJobGroupsList) GetTotalCallNum() *int32 {
	return s.TotalCallNum
}

func (s *ListJobGroupsResponseBodyJobGroupsList) SetCreationTime(v int64) *ListJobGroupsResponseBodyJobGroupsList {
	s.CreationTime = &v
	return s
}

func (s *ListJobGroupsResponseBodyJobGroupsList) SetExportProgress(v *ListJobGroupsResponseBodyJobGroupsListExportProgress) *ListJobGroupsResponseBodyJobGroupsList {
	s.ExportProgress = v
	return s
}

func (s *ListJobGroupsResponseBodyJobGroupsList) SetJobDataParsingTaskId(v string) *ListJobGroupsResponseBodyJobGroupsList {
	s.JobDataParsingTaskId = &v
	return s
}

func (s *ListJobGroupsResponseBodyJobGroupsList) SetJobGroupDescription(v string) *ListJobGroupsResponseBodyJobGroupsList {
	s.JobGroupDescription = &v
	return s
}

func (s *ListJobGroupsResponseBodyJobGroupsList) SetJobGroupId(v string) *ListJobGroupsResponseBodyJobGroupsList {
	s.JobGroupId = &v
	return s
}

func (s *ListJobGroupsResponseBodyJobGroupsList) SetJobGroupName(v string) *ListJobGroupsResponseBodyJobGroupsList {
	s.JobGroupName = &v
	return s
}

func (s *ListJobGroupsResponseBodyJobGroupsList) SetMinConcurrency(v int32) *ListJobGroupsResponseBodyJobGroupsList {
	s.MinConcurrency = &v
	return s
}

func (s *ListJobGroupsResponseBodyJobGroupsList) SetModifyTime(v string) *ListJobGroupsResponseBodyJobGroupsList {
	s.ModifyTime = &v
	return s
}

func (s *ListJobGroupsResponseBodyJobGroupsList) SetProgress(v *ListJobGroupsResponseBodyJobGroupsListProgress) *ListJobGroupsResponseBodyJobGroupsList {
	s.Progress = v
	return s
}

func (s *ListJobGroupsResponseBodyJobGroupsList) SetScriptId(v string) *ListJobGroupsResponseBodyJobGroupsList {
	s.ScriptId = &v
	return s
}

func (s *ListJobGroupsResponseBodyJobGroupsList) SetScriptName(v string) *ListJobGroupsResponseBodyJobGroupsList {
	s.ScriptName = &v
	return s
}

func (s *ListJobGroupsResponseBodyJobGroupsList) SetScriptVersion(v string) *ListJobGroupsResponseBodyJobGroupsList {
	s.ScriptVersion = &v
	return s
}

func (s *ListJobGroupsResponseBodyJobGroupsList) SetStatus(v string) *ListJobGroupsResponseBodyJobGroupsList {
	s.Status = &v
	return s
}

func (s *ListJobGroupsResponseBodyJobGroupsList) SetStrategy(v *ListJobGroupsResponseBodyJobGroupsListStrategy) *ListJobGroupsResponseBodyJobGroupsList {
	s.Strategy = v
	return s
}

func (s *ListJobGroupsResponseBodyJobGroupsList) SetTotalCallNum(v int32) *ListJobGroupsResponseBodyJobGroupsList {
	s.TotalCallNum = &v
	return s
}

func (s *ListJobGroupsResponseBodyJobGroupsList) Validate() error {
	if s.ExportProgress != nil {
		if err := s.ExportProgress.Validate(); err != nil {
			return err
		}
	}
	if s.Progress != nil {
		if err := s.Progress.Validate(); err != nil {
			return err
		}
	}
	if s.Strategy != nil {
		if err := s.Strategy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListJobGroupsResponseBodyJobGroupsListExportProgress struct {
	// This parameter is deprecated.
	//
	// example:
	//
	// http://www.xxx.com/xxx
	FileHttpUrl *string `json:"FileHttpUrl,omitempty" xml:"FileHttpUrl,omitempty"`
	// The progress of the export task.
	//
	// example:
	//
	// 1
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The status of the export task. Valid values:
	//
	// - **PENDING**: The task is pending.
	//
	// - **IN_PROGRESS**: The task is in progress.
	//
	// - **FINISHED**: The task is finished.
	//
	// - **FAILED**: The task failed.
	//
	// example:
	//
	// PENDING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListJobGroupsResponseBodyJobGroupsListExportProgress) String() string {
	return dara.Prettify(s)
}

func (s ListJobGroupsResponseBodyJobGroupsListExportProgress) GoString() string {
	return s.String()
}

func (s *ListJobGroupsResponseBodyJobGroupsListExportProgress) GetFileHttpUrl() *string {
	return s.FileHttpUrl
}

func (s *ListJobGroupsResponseBodyJobGroupsListExportProgress) GetProgress() *string {
	return s.Progress
}

func (s *ListJobGroupsResponseBodyJobGroupsListExportProgress) GetStatus() *string {
	return s.Status
}

func (s *ListJobGroupsResponseBodyJobGroupsListExportProgress) SetFileHttpUrl(v string) *ListJobGroupsResponseBodyJobGroupsListExportProgress {
	s.FileHttpUrl = &v
	return s
}

func (s *ListJobGroupsResponseBodyJobGroupsListExportProgress) SetProgress(v string) *ListJobGroupsResponseBodyJobGroupsListExportProgress {
	s.Progress = &v
	return s
}

func (s *ListJobGroupsResponseBodyJobGroupsListExportProgress) SetStatus(v string) *ListJobGroupsResponseBodyJobGroupsListExportProgress {
	s.Status = &v
	return s
}

func (s *ListJobGroupsResponseBodyJobGroupsListExportProgress) Validate() error {
	return dara.Validate(s)
}

type ListJobGroupsResponseBodyJobGroupsListProgress struct {
	// The number of canceled jobs.
	//
	// example:
	//
	// 5
	CancelledNum *int32 `json:"CancelledNum,omitempty" xml:"CancelledNum,omitempty"`
	// The total runtime. This parameter is deprecated.
	//
	// example:
	//
	// 1578550074361
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The number of running jobs.
	//
	// example:
	//
	// 10
	ExecutingNum *int32 `json:"ExecutingNum,omitempty" xml:"ExecutingNum,omitempty"`
	// The number of failed jobs.
	//
	// example:
	//
	// 5
	FailedNum *int32 `json:"FailedNum,omitempty" xml:"FailedNum,omitempty"`
	// The number of paused jobs.
	//
	// example:
	//
	// 5
	PausedNum *int32 `json:"PausedNum,omitempty" xml:"PausedNum,omitempty"`
	// The number of jobs that are being scheduled.
	//
	// example:
	//
	// 10
	Scheduling *int32 `json:"Scheduling,omitempty" xml:"Scheduling,omitempty"`
	// The start time. This parameter is deprecated.
	//
	// example:
	//
	// 1578550074361
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// > This parameter is no longer returned.
	//
	// The status of the job. Valid values:
	//
	// - **Draft**: The job is a draft.
	//
	// - **Scheduling**: The job is being scheduled.
	//
	// - **Executing**: The job is running.
	//
	// - **Completed**: The job is completed.
	//
	// - **Paused**: The job is paused.
	//
	// - **Failed**: The job failed.
	//
	// - **Cancelled**: The job is canceled.
	//
	// - **Initializing**: The job is being initialized.
	//
	// example:
	//
	// Scheduling
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of completed jobs.
	//
	// example:
	//
	// 3
	TotalCompleted *int32 `json:"TotalCompleted,omitempty" xml:"TotalCompleted,omitempty"`
	// The total number of jobs.
	//
	// example:
	//
	// 20
	TotalJobs *int32 `json:"TotalJobs,omitempty" xml:"TotalJobs,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	TotalNotAnswered *int32 `json:"TotalNotAnswered,omitempty" xml:"TotalNotAnswered,omitempty"`
}

func (s ListJobGroupsResponseBodyJobGroupsListProgress) String() string {
	return dara.Prettify(s)
}

func (s ListJobGroupsResponseBodyJobGroupsListProgress) GoString() string {
	return s.String()
}

func (s *ListJobGroupsResponseBodyJobGroupsListProgress) GetCancelledNum() *int32 {
	return s.CancelledNum
}

func (s *ListJobGroupsResponseBodyJobGroupsListProgress) GetDuration() *int32 {
	return s.Duration
}

func (s *ListJobGroupsResponseBodyJobGroupsListProgress) GetExecutingNum() *int32 {
	return s.ExecutingNum
}

func (s *ListJobGroupsResponseBodyJobGroupsListProgress) GetFailedNum() *int32 {
	return s.FailedNum
}

func (s *ListJobGroupsResponseBodyJobGroupsListProgress) GetPausedNum() *int32 {
	return s.PausedNum
}

func (s *ListJobGroupsResponseBodyJobGroupsListProgress) GetScheduling() *int32 {
	return s.Scheduling
}

func (s *ListJobGroupsResponseBodyJobGroupsListProgress) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListJobGroupsResponseBodyJobGroupsListProgress) GetStatus() *string {
	return s.Status
}

func (s *ListJobGroupsResponseBodyJobGroupsListProgress) GetTotalCompleted() *int32 {
	return s.TotalCompleted
}

func (s *ListJobGroupsResponseBodyJobGroupsListProgress) GetTotalJobs() *int32 {
	return s.TotalJobs
}

func (s *ListJobGroupsResponseBodyJobGroupsListProgress) GetTotalNotAnswered() *int32 {
	return s.TotalNotAnswered
}

func (s *ListJobGroupsResponseBodyJobGroupsListProgress) SetCancelledNum(v int32) *ListJobGroupsResponseBodyJobGroupsListProgress {
	s.CancelledNum = &v
	return s
}

func (s *ListJobGroupsResponseBodyJobGroupsListProgress) SetDuration(v int32) *ListJobGroupsResponseBodyJobGroupsListProgress {
	s.Duration = &v
	return s
}

func (s *ListJobGroupsResponseBodyJobGroupsListProgress) SetExecutingNum(v int32) *ListJobGroupsResponseBodyJobGroupsListProgress {
	s.ExecutingNum = &v
	return s
}

func (s *ListJobGroupsResponseBodyJobGroupsListProgress) SetFailedNum(v int32) *ListJobGroupsResponseBodyJobGroupsListProgress {
	s.FailedNum = &v
	return s
}

func (s *ListJobGroupsResponseBodyJobGroupsListProgress) SetPausedNum(v int32) *ListJobGroupsResponseBodyJobGroupsListProgress {
	s.PausedNum = &v
	return s
}

func (s *ListJobGroupsResponseBodyJobGroupsListProgress) SetScheduling(v int32) *ListJobGroupsResponseBodyJobGroupsListProgress {
	s.Scheduling = &v
	return s
}

func (s *ListJobGroupsResponseBodyJobGroupsListProgress) SetStartTime(v int64) *ListJobGroupsResponseBodyJobGroupsListProgress {
	s.StartTime = &v
	return s
}

func (s *ListJobGroupsResponseBodyJobGroupsListProgress) SetStatus(v string) *ListJobGroupsResponseBodyJobGroupsListProgress {
	s.Status = &v
	return s
}

func (s *ListJobGroupsResponseBodyJobGroupsListProgress) SetTotalCompleted(v int32) *ListJobGroupsResponseBodyJobGroupsListProgress {
	s.TotalCompleted = &v
	return s
}

func (s *ListJobGroupsResponseBodyJobGroupsListProgress) SetTotalJobs(v int32) *ListJobGroupsResponseBodyJobGroupsListProgress {
	s.TotalJobs = &v
	return s
}

func (s *ListJobGroupsResponseBodyJobGroupsListProgress) SetTotalNotAnswered(v int32) *ListJobGroupsResponseBodyJobGroupsListProgress {
	s.TotalNotAnswered = &v
	return s
}

func (s *ListJobGroupsResponseBodyJobGroupsListProgress) Validate() error {
	return dara.Validate(s)
}

type ListJobGroupsResponseBodyJobGroupsListStrategy struct {
	// The end time of the calling window.
	//
	// example:
	//
	// 2209702074000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time of the calling window.
	//
	// example:
	//
	// 1578550074000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListJobGroupsResponseBodyJobGroupsListStrategy) String() string {
	return dara.Prettify(s)
}

func (s ListJobGroupsResponseBodyJobGroupsListStrategy) GoString() string {
	return s.String()
}

func (s *ListJobGroupsResponseBodyJobGroupsListStrategy) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListJobGroupsResponseBodyJobGroupsListStrategy) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListJobGroupsResponseBodyJobGroupsListStrategy) SetEndTime(v int64) *ListJobGroupsResponseBodyJobGroupsListStrategy {
	s.EndTime = &v
	return s
}

func (s *ListJobGroupsResponseBodyJobGroupsListStrategy) SetStartTime(v int64) *ListJobGroupsResponseBodyJobGroupsListStrategy {
	s.StartTime = &v
	return s
}

func (s *ListJobGroupsResponseBodyJobGroupsListStrategy) Validate() error {
	return dara.Validate(s)
}
