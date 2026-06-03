// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobGroupsAsyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListJobGroupsAsyncResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListJobGroupsAsyncResponseBody
	GetHttpStatusCode() *int32
	SetJobGroups(v []*ListJobGroupsAsyncResponseBodyJobGroups) *ListJobGroupsAsyncResponseBody
	GetJobGroups() []*ListJobGroupsAsyncResponseBodyJobGroups
	SetMessage(v string) *ListJobGroupsAsyncResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListJobGroupsAsyncResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListJobGroupsAsyncResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListJobGroupsAsyncResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListJobGroupsAsyncResponseBody
	GetSuccess() *bool
	SetTimeout(v bool) *ListJobGroupsAsyncResponseBody
	GetTimeout() *bool
	SetTotalCount(v int32) *ListJobGroupsAsyncResponseBody
	GetTotalCount() *int32
	SetVaild(v bool) *ListJobGroupsAsyncResponseBody
	GetVaild() *bool
}

type ListJobGroupsAsyncResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32                                     `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	JobGroups      []*ListJobGroupsAsyncResponseBodyJobGroups `json:"JobGroups,omitempty" xml:"JobGroups,omitempty" type:"Repeated"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// true
	Timeout *bool `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// true
	Vaild *bool `json:"Vaild,omitempty" xml:"Vaild,omitempty"`
}

func (s ListJobGroupsAsyncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListJobGroupsAsyncResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobGroupsAsyncResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListJobGroupsAsyncResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListJobGroupsAsyncResponseBody) GetJobGroups() []*ListJobGroupsAsyncResponseBodyJobGroups {
	return s.JobGroups
}

func (s *ListJobGroupsAsyncResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListJobGroupsAsyncResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListJobGroupsAsyncResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListJobGroupsAsyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListJobGroupsAsyncResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListJobGroupsAsyncResponseBody) GetTimeout() *bool {
	return s.Timeout
}

func (s *ListJobGroupsAsyncResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListJobGroupsAsyncResponseBody) GetVaild() *bool {
	return s.Vaild
}

func (s *ListJobGroupsAsyncResponseBody) SetCode(v string) *ListJobGroupsAsyncResponseBody {
	s.Code = &v
	return s
}

func (s *ListJobGroupsAsyncResponseBody) SetHttpStatusCode(v int32) *ListJobGroupsAsyncResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListJobGroupsAsyncResponseBody) SetJobGroups(v []*ListJobGroupsAsyncResponseBodyJobGroups) *ListJobGroupsAsyncResponseBody {
	s.JobGroups = v
	return s
}

func (s *ListJobGroupsAsyncResponseBody) SetMessage(v string) *ListJobGroupsAsyncResponseBody {
	s.Message = &v
	return s
}

func (s *ListJobGroupsAsyncResponseBody) SetPageNumber(v int32) *ListJobGroupsAsyncResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListJobGroupsAsyncResponseBody) SetPageSize(v int32) *ListJobGroupsAsyncResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListJobGroupsAsyncResponseBody) SetRequestId(v string) *ListJobGroupsAsyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobGroupsAsyncResponseBody) SetSuccess(v bool) *ListJobGroupsAsyncResponseBody {
	s.Success = &v
	return s
}

func (s *ListJobGroupsAsyncResponseBody) SetTimeout(v bool) *ListJobGroupsAsyncResponseBody {
	s.Timeout = &v
	return s
}

func (s *ListJobGroupsAsyncResponseBody) SetTotalCount(v int32) *ListJobGroupsAsyncResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListJobGroupsAsyncResponseBody) SetVaild(v bool) *ListJobGroupsAsyncResponseBody {
	s.Vaild = &v
	return s
}

func (s *ListJobGroupsAsyncResponseBody) Validate() error {
	if s.JobGroups != nil {
		for _, item := range s.JobGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListJobGroupsAsyncResponseBodyJobGroups struct {
	// example:
	//
	// 1640316786259
	CreationTime *int64 `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// example:
	//
	// {}
	ExportProgress *ListJobGroupsAsyncResponseBodyJobGroupsExportProgress `json:"ExportProgress,omitempty" xml:"ExportProgress,omitempty" type:"Struct"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	JobDataParsingTaskId *string `json:"JobDataParsingTaskId,omitempty" xml:"JobDataParsingTaskId,omitempty"`
	// example:
	//
	// xxx
	JobGroupDescription *string `json:"JobGroupDescription,omitempty" xml:"JobGroupDescription,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	// example:
	//
	// xxx
	JobGroupName   *string `json:"JobGroupName,omitempty" xml:"JobGroupName,omitempty"`
	MinConcurrency *int32  `json:"MinConcurrency,omitempty" xml:"MinConcurrency,omitempty"`
	// example:
	//
	// 1640316786259
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// {}
	Progress *ListJobGroupsAsyncResponseBodyJobGroupsProgress `json:"Progress,omitempty" xml:"Progress,omitempty" type:"Struct"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// example:
	//
	// xxxx
	ScriptName *string `json:"ScriptName,omitempty" xml:"ScriptName,omitempty"`
	// example:
	//
	// 111
	ScriptVersion *string `json:"ScriptVersion,omitempty" xml:"ScriptVersion,omitempty"`
	// example:
	//
	// Scheduling
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// {}
	Strategy *ListJobGroupsAsyncResponseBodyJobGroupsStrategy `json:"Strategy,omitempty" xml:"Strategy,omitempty" type:"Struct"`
	// example:
	//
	// 100
	TotalCallNum *int32 `json:"TotalCallNum,omitempty" xml:"TotalCallNum,omitempty"`
}

func (s ListJobGroupsAsyncResponseBodyJobGroups) String() string {
	return dara.Prettify(s)
}

func (s ListJobGroupsAsyncResponseBodyJobGroups) GoString() string {
	return s.String()
}

func (s *ListJobGroupsAsyncResponseBodyJobGroups) GetCreationTime() *int64 {
	return s.CreationTime
}

func (s *ListJobGroupsAsyncResponseBodyJobGroups) GetExportProgress() *ListJobGroupsAsyncResponseBodyJobGroupsExportProgress {
	return s.ExportProgress
}

func (s *ListJobGroupsAsyncResponseBodyJobGroups) GetJobDataParsingTaskId() *string {
	return s.JobDataParsingTaskId
}

func (s *ListJobGroupsAsyncResponseBodyJobGroups) GetJobGroupDescription() *string {
	return s.JobGroupDescription
}

func (s *ListJobGroupsAsyncResponseBodyJobGroups) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *ListJobGroupsAsyncResponseBodyJobGroups) GetJobGroupName() *string {
	return s.JobGroupName
}

func (s *ListJobGroupsAsyncResponseBodyJobGroups) GetMinConcurrency() *int32 {
	return s.MinConcurrency
}

func (s *ListJobGroupsAsyncResponseBodyJobGroups) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListJobGroupsAsyncResponseBodyJobGroups) GetProgress() *ListJobGroupsAsyncResponseBodyJobGroupsProgress {
	return s.Progress
}

func (s *ListJobGroupsAsyncResponseBodyJobGroups) GetScriptId() *string {
	return s.ScriptId
}

func (s *ListJobGroupsAsyncResponseBodyJobGroups) GetScriptName() *string {
	return s.ScriptName
}

func (s *ListJobGroupsAsyncResponseBodyJobGroups) GetScriptVersion() *string {
	return s.ScriptVersion
}

func (s *ListJobGroupsAsyncResponseBodyJobGroups) GetStatus() *string {
	return s.Status
}

func (s *ListJobGroupsAsyncResponseBodyJobGroups) GetStrategy() *ListJobGroupsAsyncResponseBodyJobGroupsStrategy {
	return s.Strategy
}

func (s *ListJobGroupsAsyncResponseBodyJobGroups) GetTotalCallNum() *int32 {
	return s.TotalCallNum
}

func (s *ListJobGroupsAsyncResponseBodyJobGroups) SetCreationTime(v int64) *ListJobGroupsAsyncResponseBodyJobGroups {
	s.CreationTime = &v
	return s
}

func (s *ListJobGroupsAsyncResponseBodyJobGroups) SetExportProgress(v *ListJobGroupsAsyncResponseBodyJobGroupsExportProgress) *ListJobGroupsAsyncResponseBodyJobGroups {
	s.ExportProgress = v
	return s
}

func (s *ListJobGroupsAsyncResponseBodyJobGroups) SetJobDataParsingTaskId(v string) *ListJobGroupsAsyncResponseBodyJobGroups {
	s.JobDataParsingTaskId = &v
	return s
}

func (s *ListJobGroupsAsyncResponseBodyJobGroups) SetJobGroupDescription(v string) *ListJobGroupsAsyncResponseBodyJobGroups {
	s.JobGroupDescription = &v
	return s
}

func (s *ListJobGroupsAsyncResponseBodyJobGroups) SetJobGroupId(v string) *ListJobGroupsAsyncResponseBodyJobGroups {
	s.JobGroupId = &v
	return s
}

func (s *ListJobGroupsAsyncResponseBodyJobGroups) SetJobGroupName(v string) *ListJobGroupsAsyncResponseBodyJobGroups {
	s.JobGroupName = &v
	return s
}

func (s *ListJobGroupsAsyncResponseBodyJobGroups) SetMinConcurrency(v int32) *ListJobGroupsAsyncResponseBodyJobGroups {
	s.MinConcurrency = &v
	return s
}

func (s *ListJobGroupsAsyncResponseBodyJobGroups) SetModifyTime(v string) *ListJobGroupsAsyncResponseBodyJobGroups {
	s.ModifyTime = &v
	return s
}

func (s *ListJobGroupsAsyncResponseBodyJobGroups) SetProgress(v *ListJobGroupsAsyncResponseBodyJobGroupsProgress) *ListJobGroupsAsyncResponseBodyJobGroups {
	s.Progress = v
	return s
}

func (s *ListJobGroupsAsyncResponseBodyJobGroups) SetScriptId(v string) *ListJobGroupsAsyncResponseBodyJobGroups {
	s.ScriptId = &v
	return s
}

func (s *ListJobGroupsAsyncResponseBodyJobGroups) SetScriptName(v string) *ListJobGroupsAsyncResponseBodyJobGroups {
	s.ScriptName = &v
	return s
}

func (s *ListJobGroupsAsyncResponseBodyJobGroups) SetScriptVersion(v string) *ListJobGroupsAsyncResponseBodyJobGroups {
	s.ScriptVersion = &v
	return s
}

func (s *ListJobGroupsAsyncResponseBodyJobGroups) SetStatus(v string) *ListJobGroupsAsyncResponseBodyJobGroups {
	s.Status = &v
	return s
}

func (s *ListJobGroupsAsyncResponseBodyJobGroups) SetStrategy(v *ListJobGroupsAsyncResponseBodyJobGroupsStrategy) *ListJobGroupsAsyncResponseBodyJobGroups {
	s.Strategy = v
	return s
}

func (s *ListJobGroupsAsyncResponseBodyJobGroups) SetTotalCallNum(v int32) *ListJobGroupsAsyncResponseBodyJobGroups {
	s.TotalCallNum = &v
	return s
}

func (s *ListJobGroupsAsyncResponseBodyJobGroups) Validate() error {
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

type ListJobGroupsAsyncResponseBodyJobGroupsExportProgress struct {
	// example:
	//
	// http://
	FileHttpUrl *string `json:"FileHttpUrl,omitempty" xml:"FileHttpUrl,omitempty"`
	// example:
	//
	// 50
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// PENDING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListJobGroupsAsyncResponseBodyJobGroupsExportProgress) String() string {
	return dara.Prettify(s)
}

func (s ListJobGroupsAsyncResponseBodyJobGroupsExportProgress) GoString() string {
	return s.String()
}

func (s *ListJobGroupsAsyncResponseBodyJobGroupsExportProgress) GetFileHttpUrl() *string {
	return s.FileHttpUrl
}

func (s *ListJobGroupsAsyncResponseBodyJobGroupsExportProgress) GetProgress() *string {
	return s.Progress
}

func (s *ListJobGroupsAsyncResponseBodyJobGroupsExportProgress) GetStatus() *string {
	return s.Status
}

func (s *ListJobGroupsAsyncResponseBodyJobGroupsExportProgress) SetFileHttpUrl(v string) *ListJobGroupsAsyncResponseBodyJobGroupsExportProgress {
	s.FileHttpUrl = &v
	return s
}

func (s *ListJobGroupsAsyncResponseBodyJobGroupsExportProgress) SetProgress(v string) *ListJobGroupsAsyncResponseBodyJobGroupsExportProgress {
	s.Progress = &v
	return s
}

func (s *ListJobGroupsAsyncResponseBodyJobGroupsExportProgress) SetStatus(v string) *ListJobGroupsAsyncResponseBodyJobGroupsExportProgress {
	s.Status = &v
	return s
}

func (s *ListJobGroupsAsyncResponseBodyJobGroupsExportProgress) Validate() error {
	return dara.Validate(s)
}

type ListJobGroupsAsyncResponseBodyJobGroupsProgress struct {
	// example:
	//
	// 10
	CancelledNum *int32 `json:"CancelledNum,omitempty" xml:"CancelledNum,omitempty"`
	// example:
	//
	// 1000
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 20
	ExecutingNum *int32 `json:"ExecutingNum,omitempty" xml:"ExecutingNum,omitempty"`
	// example:
	//
	// 10
	FailedNum *int32 `json:"FailedNum,omitempty" xml:"FailedNum,omitempty"`
	// example:
	//
	// 10
	PausedNum *int32 `json:"PausedNum,omitempty" xml:"PausedNum,omitempty"`
	// example:
	//
	// 20
	Scheduling *int32 `json:"Scheduling,omitempty" xml:"Scheduling,omitempty"`
	// example:
	//
	// 1640316786259
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 10
	TotalCompleted *int32 `json:"TotalCompleted,omitempty" xml:"TotalCompleted,omitempty"`
	// example:
	//
	// 100
	TotalJobs *int32 `json:"TotalJobs,omitempty" xml:"TotalJobs,omitempty"`
	// example:
	//
	// 1
	TotalNotAnswered *int32 `json:"TotalNotAnswered,omitempty" xml:"TotalNotAnswered,omitempty"`
}

func (s ListJobGroupsAsyncResponseBodyJobGroupsProgress) String() string {
	return dara.Prettify(s)
}

func (s ListJobGroupsAsyncResponseBodyJobGroupsProgress) GoString() string {
	return s.String()
}

func (s *ListJobGroupsAsyncResponseBodyJobGroupsProgress) GetCancelledNum() *int32 {
	return s.CancelledNum
}

func (s *ListJobGroupsAsyncResponseBodyJobGroupsProgress) GetDuration() *int32 {
	return s.Duration
}

func (s *ListJobGroupsAsyncResponseBodyJobGroupsProgress) GetExecutingNum() *int32 {
	return s.ExecutingNum
}

func (s *ListJobGroupsAsyncResponseBodyJobGroupsProgress) GetFailedNum() *int32 {
	return s.FailedNum
}

func (s *ListJobGroupsAsyncResponseBodyJobGroupsProgress) GetPausedNum() *int32 {
	return s.PausedNum
}

func (s *ListJobGroupsAsyncResponseBodyJobGroupsProgress) GetScheduling() *int32 {
	return s.Scheduling
}

func (s *ListJobGroupsAsyncResponseBodyJobGroupsProgress) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListJobGroupsAsyncResponseBodyJobGroupsProgress) GetStatus() *string {
	return s.Status
}

func (s *ListJobGroupsAsyncResponseBodyJobGroupsProgress) GetTotalCompleted() *int32 {
	return s.TotalCompleted
}

func (s *ListJobGroupsAsyncResponseBodyJobGroupsProgress) GetTotalJobs() *int32 {
	return s.TotalJobs
}

func (s *ListJobGroupsAsyncResponseBodyJobGroupsProgress) GetTotalNotAnswered() *int32 {
	return s.TotalNotAnswered
}

func (s *ListJobGroupsAsyncResponseBodyJobGroupsProgress) SetCancelledNum(v int32) *ListJobGroupsAsyncResponseBodyJobGroupsProgress {
	s.CancelledNum = &v
	return s
}

func (s *ListJobGroupsAsyncResponseBodyJobGroupsProgress) SetDuration(v int32) *ListJobGroupsAsyncResponseBodyJobGroupsProgress {
	s.Duration = &v
	return s
}

func (s *ListJobGroupsAsyncResponseBodyJobGroupsProgress) SetExecutingNum(v int32) *ListJobGroupsAsyncResponseBodyJobGroupsProgress {
	s.ExecutingNum = &v
	return s
}

func (s *ListJobGroupsAsyncResponseBodyJobGroupsProgress) SetFailedNum(v int32) *ListJobGroupsAsyncResponseBodyJobGroupsProgress {
	s.FailedNum = &v
	return s
}

func (s *ListJobGroupsAsyncResponseBodyJobGroupsProgress) SetPausedNum(v int32) *ListJobGroupsAsyncResponseBodyJobGroupsProgress {
	s.PausedNum = &v
	return s
}

func (s *ListJobGroupsAsyncResponseBodyJobGroupsProgress) SetScheduling(v int32) *ListJobGroupsAsyncResponseBodyJobGroupsProgress {
	s.Scheduling = &v
	return s
}

func (s *ListJobGroupsAsyncResponseBodyJobGroupsProgress) SetStartTime(v int64) *ListJobGroupsAsyncResponseBodyJobGroupsProgress {
	s.StartTime = &v
	return s
}

func (s *ListJobGroupsAsyncResponseBodyJobGroupsProgress) SetStatus(v string) *ListJobGroupsAsyncResponseBodyJobGroupsProgress {
	s.Status = &v
	return s
}

func (s *ListJobGroupsAsyncResponseBodyJobGroupsProgress) SetTotalCompleted(v int32) *ListJobGroupsAsyncResponseBodyJobGroupsProgress {
	s.TotalCompleted = &v
	return s
}

func (s *ListJobGroupsAsyncResponseBodyJobGroupsProgress) SetTotalJobs(v int32) *ListJobGroupsAsyncResponseBodyJobGroupsProgress {
	s.TotalJobs = &v
	return s
}

func (s *ListJobGroupsAsyncResponseBodyJobGroupsProgress) SetTotalNotAnswered(v int32) *ListJobGroupsAsyncResponseBodyJobGroupsProgress {
	s.TotalNotAnswered = &v
	return s
}

func (s *ListJobGroupsAsyncResponseBodyJobGroupsProgress) Validate() error {
	return dara.Validate(s)
}

type ListJobGroupsAsyncResponseBodyJobGroupsStrategy struct {
	// example:
	//
	// 1640316786259
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1640316786259
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListJobGroupsAsyncResponseBodyJobGroupsStrategy) String() string {
	return dara.Prettify(s)
}

func (s ListJobGroupsAsyncResponseBodyJobGroupsStrategy) GoString() string {
	return s.String()
}

func (s *ListJobGroupsAsyncResponseBodyJobGroupsStrategy) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListJobGroupsAsyncResponseBodyJobGroupsStrategy) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListJobGroupsAsyncResponseBodyJobGroupsStrategy) SetEndTime(v int64) *ListJobGroupsAsyncResponseBodyJobGroupsStrategy {
	s.EndTime = &v
	return s
}

func (s *ListJobGroupsAsyncResponseBodyJobGroupsStrategy) SetStartTime(v int64) *ListJobGroupsAsyncResponseBodyJobGroupsStrategy {
	s.StartTime = &v
	return s
}

func (s *ListJobGroupsAsyncResponseBodyJobGroupsStrategy) Validate() error {
	return dara.Validate(s)
}
