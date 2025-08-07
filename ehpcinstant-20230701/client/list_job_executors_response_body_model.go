// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobExecutorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExecutorStatus(v *ListJobExecutorsResponseBodyExecutorStatus) *ListJobExecutorsResponseBody
	GetExecutorStatus() *ListJobExecutorsResponseBodyExecutorStatus
	SetExecutors(v []*ListJobExecutorsResponseBodyExecutors) *ListJobExecutorsResponseBody
	GetExecutors() []*ListJobExecutorsResponseBodyExecutors
	SetJobId(v string) *ListJobExecutorsResponseBody
	GetJobId() *string
	SetPageNumber(v int32) *ListJobExecutorsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListJobExecutorsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListJobExecutorsResponseBody
	GetRequestId() *string
	SetTaskName(v string) *ListJobExecutorsResponseBody
	GetTaskName() *string
	SetTotalCount(v string) *ListJobExecutorsResponseBody
	GetTotalCount() *string
}

type ListJobExecutorsResponseBody struct {
	ExecutorStatus *ListJobExecutorsResponseBodyExecutorStatus `json:"ExecutorStatus,omitempty" xml:"ExecutorStatus,omitempty" type:"Struct"`
	Executors      []*ListJobExecutorsResponseBodyExecutors    `json:"Executors,omitempty" xml:"Executors,omitempty" type:"Repeated"`
	// example:
	//
	// job-xxxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
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
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// task0
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// 50
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListJobExecutorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListJobExecutorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobExecutorsResponseBody) GetExecutorStatus() *ListJobExecutorsResponseBodyExecutorStatus {
	return s.ExecutorStatus
}

func (s *ListJobExecutorsResponseBody) GetExecutors() []*ListJobExecutorsResponseBodyExecutors {
	return s.Executors
}

func (s *ListJobExecutorsResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *ListJobExecutorsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListJobExecutorsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListJobExecutorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListJobExecutorsResponseBody) GetTaskName() *string {
	return s.TaskName
}

func (s *ListJobExecutorsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListJobExecutorsResponseBody) SetExecutorStatus(v *ListJobExecutorsResponseBodyExecutorStatus) *ListJobExecutorsResponseBody {
	s.ExecutorStatus = v
	return s
}

func (s *ListJobExecutorsResponseBody) SetExecutors(v []*ListJobExecutorsResponseBodyExecutors) *ListJobExecutorsResponseBody {
	s.Executors = v
	return s
}

func (s *ListJobExecutorsResponseBody) SetJobId(v string) *ListJobExecutorsResponseBody {
	s.JobId = &v
	return s
}

func (s *ListJobExecutorsResponseBody) SetPageNumber(v int32) *ListJobExecutorsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListJobExecutorsResponseBody) SetPageSize(v int32) *ListJobExecutorsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListJobExecutorsResponseBody) SetRequestId(v string) *ListJobExecutorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobExecutorsResponseBody) SetTaskName(v string) *ListJobExecutorsResponseBody {
	s.TaskName = &v
	return s
}

func (s *ListJobExecutorsResponseBody) SetTotalCount(v string) *ListJobExecutorsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListJobExecutorsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListJobExecutorsResponseBodyExecutorStatus struct {
	Deleted    *int32 `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Exception  *int32 `json:"Exception,omitempty" xml:"Exception,omitempty"`
	Failed     *int32 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	Initing    *int32 `json:"Initing,omitempty" xml:"Initing,omitempty"`
	Pending    *int32 `json:"Pending,omitempty" xml:"Pending,omitempty"`
	Restarting *int32 `json:"Restarting,omitempty" xml:"Restarting,omitempty"`
	Running    *int32 `json:"Running,omitempty" xml:"Running,omitempty"`
	Succeeded  *int32 `json:"Succeeded,omitempty" xml:"Succeeded,omitempty"`
	Suspended  *int32 `json:"Suspended,omitempty" xml:"Suspended,omitempty"`
}

func (s ListJobExecutorsResponseBodyExecutorStatus) String() string {
	return dara.Prettify(s)
}

func (s ListJobExecutorsResponseBodyExecutorStatus) GoString() string {
	return s.String()
}

func (s *ListJobExecutorsResponseBodyExecutorStatus) GetDeleted() *int32 {
	return s.Deleted
}

func (s *ListJobExecutorsResponseBodyExecutorStatus) GetException() *int32 {
	return s.Exception
}

func (s *ListJobExecutorsResponseBodyExecutorStatus) GetFailed() *int32 {
	return s.Failed
}

func (s *ListJobExecutorsResponseBodyExecutorStatus) GetIniting() *int32 {
	return s.Initing
}

func (s *ListJobExecutorsResponseBodyExecutorStatus) GetPending() *int32 {
	return s.Pending
}

func (s *ListJobExecutorsResponseBodyExecutorStatus) GetRestarting() *int32 {
	return s.Restarting
}

func (s *ListJobExecutorsResponseBodyExecutorStatus) GetRunning() *int32 {
	return s.Running
}

func (s *ListJobExecutorsResponseBodyExecutorStatus) GetSucceeded() *int32 {
	return s.Succeeded
}

func (s *ListJobExecutorsResponseBodyExecutorStatus) GetSuspended() *int32 {
	return s.Suspended
}

func (s *ListJobExecutorsResponseBodyExecutorStatus) SetDeleted(v int32) *ListJobExecutorsResponseBodyExecutorStatus {
	s.Deleted = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutorStatus) SetException(v int32) *ListJobExecutorsResponseBodyExecutorStatus {
	s.Exception = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutorStatus) SetFailed(v int32) *ListJobExecutorsResponseBodyExecutorStatus {
	s.Failed = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutorStatus) SetIniting(v int32) *ListJobExecutorsResponseBodyExecutorStatus {
	s.Initing = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutorStatus) SetPending(v int32) *ListJobExecutorsResponseBodyExecutorStatus {
	s.Pending = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutorStatus) SetRestarting(v int32) *ListJobExecutorsResponseBodyExecutorStatus {
	s.Restarting = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutorStatus) SetRunning(v int32) *ListJobExecutorsResponseBodyExecutorStatus {
	s.Running = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutorStatus) SetSucceeded(v int32) *ListJobExecutorsResponseBodyExecutorStatus {
	s.Succeeded = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutorStatus) SetSuspended(v int32) *ListJobExecutorsResponseBodyExecutorStatus {
	s.Suspended = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutorStatus) Validate() error {
	return dara.Validate(s)
}

type ListJobExecutorsResponseBodyExecutors struct {
	// example:
	//
	// 0
	ArrayIndex    *int32 `json:"ArrayIndex,omitempty" xml:"ArrayIndex,omitempty"`
	BlockDuration *int32 `json:"BlockDuration,omitempty" xml:"BlockDuration,omitempty"`
	// example:
	//
	// 2024-02-20 10:04:10
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2024-02-20 10:04:18
	EndTime           *string   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExecutorId        *string   `json:"ExecutorId,omitempty" xml:"ExecutorId,omitempty"`
	ExpirationTime    *string   `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	ExternalIpAddress []*string `json:"ExternalIpAddress,omitempty" xml:"ExternalIpAddress,omitempty" type:"Repeated"`
	HostName          []*string `json:"HostName,omitempty" xml:"HostName,omitempty" type:"Repeated"`
	IpAddress         []*string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty" type:"Repeated"`
	Preemptible       *bool     `json:"Preemptible,omitempty" xml:"Preemptible,omitempty"`
	StartTime         *string   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// Creating executor
	StatusReason *string                                      `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	Tags         []*ListJobExecutorsResponseBodyExecutorsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListJobExecutorsResponseBodyExecutors) String() string {
	return dara.Prettify(s)
}

func (s ListJobExecutorsResponseBodyExecutors) GoString() string {
	return s.String()
}

func (s *ListJobExecutorsResponseBodyExecutors) GetArrayIndex() *int32 {
	return s.ArrayIndex
}

func (s *ListJobExecutorsResponseBodyExecutors) GetBlockDuration() *int32 {
	return s.BlockDuration
}

func (s *ListJobExecutorsResponseBodyExecutors) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListJobExecutorsResponseBodyExecutors) GetEndTime() *string {
	return s.EndTime
}

func (s *ListJobExecutorsResponseBodyExecutors) GetExecutorId() *string {
	return s.ExecutorId
}

func (s *ListJobExecutorsResponseBodyExecutors) GetExpirationTime() *string {
	return s.ExpirationTime
}

func (s *ListJobExecutorsResponseBodyExecutors) GetExternalIpAddress() []*string {
	return s.ExternalIpAddress
}

func (s *ListJobExecutorsResponseBodyExecutors) GetHostName() []*string {
	return s.HostName
}

func (s *ListJobExecutorsResponseBodyExecutors) GetIpAddress() []*string {
	return s.IpAddress
}

func (s *ListJobExecutorsResponseBodyExecutors) GetPreemptible() *bool {
	return s.Preemptible
}

func (s *ListJobExecutorsResponseBodyExecutors) GetStartTime() *string {
	return s.StartTime
}

func (s *ListJobExecutorsResponseBodyExecutors) GetStatus() *string {
	return s.Status
}

func (s *ListJobExecutorsResponseBodyExecutors) GetStatusReason() *string {
	return s.StatusReason
}

func (s *ListJobExecutorsResponseBodyExecutors) GetTags() []*ListJobExecutorsResponseBodyExecutorsTags {
	return s.Tags
}

func (s *ListJobExecutorsResponseBodyExecutors) SetArrayIndex(v int32) *ListJobExecutorsResponseBodyExecutors {
	s.ArrayIndex = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutors) SetBlockDuration(v int32) *ListJobExecutorsResponseBodyExecutors {
	s.BlockDuration = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutors) SetCreateTime(v string) *ListJobExecutorsResponseBodyExecutors {
	s.CreateTime = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutors) SetEndTime(v string) *ListJobExecutorsResponseBodyExecutors {
	s.EndTime = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutors) SetExecutorId(v string) *ListJobExecutorsResponseBodyExecutors {
	s.ExecutorId = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutors) SetExpirationTime(v string) *ListJobExecutorsResponseBodyExecutors {
	s.ExpirationTime = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutors) SetExternalIpAddress(v []*string) *ListJobExecutorsResponseBodyExecutors {
	s.ExternalIpAddress = v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutors) SetHostName(v []*string) *ListJobExecutorsResponseBodyExecutors {
	s.HostName = v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutors) SetIpAddress(v []*string) *ListJobExecutorsResponseBodyExecutors {
	s.IpAddress = v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutors) SetPreemptible(v bool) *ListJobExecutorsResponseBodyExecutors {
	s.Preemptible = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutors) SetStartTime(v string) *ListJobExecutorsResponseBodyExecutors {
	s.StartTime = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutors) SetStatus(v string) *ListJobExecutorsResponseBodyExecutors {
	s.Status = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutors) SetStatusReason(v string) *ListJobExecutorsResponseBodyExecutors {
	s.StatusReason = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutors) SetTags(v []*ListJobExecutorsResponseBodyExecutorsTags) *ListJobExecutorsResponseBodyExecutors {
	s.Tags = v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutors) Validate() error {
	return dara.Validate(s)
}

type ListJobExecutorsResponseBodyExecutorsTags struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListJobExecutorsResponseBodyExecutorsTags) String() string {
	return dara.Prettify(s)
}

func (s ListJobExecutorsResponseBodyExecutorsTags) GoString() string {
	return s.String()
}

func (s *ListJobExecutorsResponseBodyExecutorsTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ListJobExecutorsResponseBodyExecutorsTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ListJobExecutorsResponseBodyExecutorsTags) SetTagKey(v string) *ListJobExecutorsResponseBodyExecutorsTags {
	s.TagKey = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutorsTags) SetTagValue(v string) *ListJobExecutorsResponseBodyExecutorsTags {
	s.TagValue = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutorsTags) Validate() error {
	return dara.Validate(s)
}
