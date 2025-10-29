// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLivePullToPushListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeLivePullToPushListResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLivePullToPushListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeLivePullToPushListResponseBody
	GetRequestId() *string
	SetTaskList(v []*DescribeLivePullToPushListResponseBodyTaskList) *DescribeLivePullToPushListResponseBody
	GetTaskList() []*DescribeLivePullToPushListResponseBodyTaskList
	SetTotal(v int32) *DescribeLivePullToPushListResponseBody
	GetTotal() *int32
}

type DescribeLivePullToPushListResponseBody struct {
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
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// a05e6b15-15af-405b-a4a2-0152245*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tasks.
	TaskList []*DescribeLivePullToPushListResponseBodyTaskList `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 3
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeLivePullToPushListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLivePullToPushListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLivePullToPushListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLivePullToPushListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLivePullToPushListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLivePullToPushListResponseBody) GetTaskList() []*DescribeLivePullToPushListResponseBodyTaskList {
	return s.TaskList
}

func (s *DescribeLivePullToPushListResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeLivePullToPushListResponseBody) SetPageNumber(v int32) *DescribeLivePullToPushListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeLivePullToPushListResponseBody) SetPageSize(v int32) *DescribeLivePullToPushListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLivePullToPushListResponseBody) SetRequestId(v string) *DescribeLivePullToPushListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLivePullToPushListResponseBody) SetTaskList(v []*DescribeLivePullToPushListResponseBodyTaskList) *DescribeLivePullToPushListResponseBody {
	s.TaskList = v
	return s
}

func (s *DescribeLivePullToPushListResponseBody) SetTotal(v int32) *DescribeLivePullToPushListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeLivePullToPushListResponseBody) Validate() error {
	if s.TaskList != nil {
		for _, item := range s.TaskList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLivePullToPushListResponseBodyTaskList struct {
	// The current file index.
	//
	// example:
	//
	// 0
	CurrentFileIndex *int32 `json:"CurrentFileIndex,omitempty" xml:"CurrentFileIndex,omitempty"`
	// The current offset for video playback.
	//
	// example:
	//
	// 0
	CurrentOffset *int32 `json:"CurrentOffset,omitempty" xml:"CurrentOffset,omitempty"`
	// The reason why the task was exited. Valid values:
	//
	// 	- TriggerByUser: You proactively ended the task.
	//
	// 	- OverEndTime: The specified end time was exceeded.
	//
	// >  This parameter is returned only if the task status is exited.
	//
	// example:
	//
	// TriggerByUser
	TaskExitReason *string `json:"TaskExitReason,omitempty" xml:"TaskExitReason,omitempty"`
	// The time when the task was exited. The value is a Unix timestamp in seconds.
	//
	// >  This parameter is returned only if the task status is exited.
	//
	// example:
	//
	// 1726354625
	TaskExitTime *int32 `json:"TaskExitTime,omitempty" xml:"TaskExitTime,omitempty"`
	// The task ID.
	//
	// example:
	//
	// fb0d4ac7-c7e3-4978-9743-0bf2f6e8****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The information about the task.
	TaskInfo *DescribeLivePullToPushListResponseBodyTaskListTaskInfo `json:"TaskInfo,omitempty" xml:"TaskInfo,omitempty" type:"Struct"`
	// The reason why the task was stopped.
	//
	// 	- PullStreamFailed: An exception occurred while pulling the source stream. A retry is in progress.
	//
	// 	- PushStreamFailed: An exception occurred while ingesting the stream. A retry is in progress.
	//
	// 	- UnknownError: An unknown exception occurred.
	//
	// >  This parameter is returned only if the task status is stopped.
	//
	// example:
	//
	// PullStreamFailed
	TaskInvalidReason *string `json:"TaskInvalidReason,omitempty" xml:"TaskInvalidReason,omitempty"`
	// The task status. Valid values:
	//
	// 	- 0: not started.
	//
	// 	- 1: running. Stream pulling and stream relay are normal.
	//
	// 	- 2: abnormal.
	//
	// 	- 3: stopped. It may be because exceptions occur during stream pulling or stream relay or you proactively call the StopLivePullToPush operation.
	//
	// 	- \\-1: exited.
	//
	// example:
	//
	// 0
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s DescribeLivePullToPushListResponseBodyTaskList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLivePullToPushListResponseBodyTaskList) GoString() string {
	return s.String()
}

func (s *DescribeLivePullToPushListResponseBodyTaskList) GetCurrentFileIndex() *int32 {
	return s.CurrentFileIndex
}

func (s *DescribeLivePullToPushListResponseBodyTaskList) GetCurrentOffset() *int32 {
	return s.CurrentOffset
}

func (s *DescribeLivePullToPushListResponseBodyTaskList) GetTaskExitReason() *string {
	return s.TaskExitReason
}

func (s *DescribeLivePullToPushListResponseBodyTaskList) GetTaskExitTime() *int32 {
	return s.TaskExitTime
}

func (s *DescribeLivePullToPushListResponseBodyTaskList) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeLivePullToPushListResponseBodyTaskList) GetTaskInfo() *DescribeLivePullToPushListResponseBodyTaskListTaskInfo {
	return s.TaskInfo
}

func (s *DescribeLivePullToPushListResponseBodyTaskList) GetTaskInvalidReason() *string {
	return s.TaskInvalidReason
}

func (s *DescribeLivePullToPushListResponseBodyTaskList) GetTaskStatus() *int32 {
	return s.TaskStatus
}

func (s *DescribeLivePullToPushListResponseBodyTaskList) SetCurrentFileIndex(v int32) *DescribeLivePullToPushListResponseBodyTaskList {
	s.CurrentFileIndex = &v
	return s
}

func (s *DescribeLivePullToPushListResponseBodyTaskList) SetCurrentOffset(v int32) *DescribeLivePullToPushListResponseBodyTaskList {
	s.CurrentOffset = &v
	return s
}

func (s *DescribeLivePullToPushListResponseBodyTaskList) SetTaskExitReason(v string) *DescribeLivePullToPushListResponseBodyTaskList {
	s.TaskExitReason = &v
	return s
}

func (s *DescribeLivePullToPushListResponseBodyTaskList) SetTaskExitTime(v int32) *DescribeLivePullToPushListResponseBodyTaskList {
	s.TaskExitTime = &v
	return s
}

func (s *DescribeLivePullToPushListResponseBodyTaskList) SetTaskId(v string) *DescribeLivePullToPushListResponseBodyTaskList {
	s.TaskId = &v
	return s
}

func (s *DescribeLivePullToPushListResponseBodyTaskList) SetTaskInfo(v *DescribeLivePullToPushListResponseBodyTaskListTaskInfo) *DescribeLivePullToPushListResponseBodyTaskList {
	s.TaskInfo = v
	return s
}

func (s *DescribeLivePullToPushListResponseBodyTaskList) SetTaskInvalidReason(v string) *DescribeLivePullToPushListResponseBodyTaskList {
	s.TaskInvalidReason = &v
	return s
}

func (s *DescribeLivePullToPushListResponseBodyTaskList) SetTaskStatus(v int32) *DescribeLivePullToPushListResponseBodyTaskList {
	s.TaskStatus = &v
	return s
}

func (s *DescribeLivePullToPushListResponseBodyTaskList) Validate() error {
	if s.TaskInfo != nil {
		if err := s.TaskInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLivePullToPushListResponseBodyTaskListTaskInfo struct {
	// The HTTP callback URL.
	//
	// example:
	//
	// hahaha.com
	CallbackURL *string `json:"CallbackURL,omitempty" xml:"CallbackURL,omitempty"`
	// The destination URL to which the stream is relayed.
	//
	// example:
	//
	// rtmp://qd.push.lgg.alivecdn.com/testhsc/streamhsc?live_rtmp_*******
	DstUrl *string `json:"DstUrl,omitempty" xml:"DstUrl,omitempty"`
	// The end time of the task. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-12-30T14:30:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The file index, which indicates the sequence of the file where the playback starts.
	//
	// example:
	//
	// 0
	FileIndex *int32 `json:"FileIndex,omitempty" xml:"FileIndex,omitempty"`
	// The offset of the position where the system starts to read the video resource. Unit: seconds. Valid values: positive numbers.
	//
	// >
	//
	// 	- This parameter indicates an offset from the first frame.
	//
	// 	- This parameter is applicable to only video resources from ApsaraVideo VOD or a third party.
	//
	// example:
	//
	// 0
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The number of playbacks after the first playback is complete. Valid values:
	//
	// 	- 0 (default): specifies that the video list is played only once.
	//
	// 	- \\-1: specifies that the video list is played in loop mode.
	//
	// 	- Positive integer: specifies the number of times the video list repeats after the first playback is complete.
	//
	// >  This parameter is applicable to only video resources from ApsaraVideo VOD or a third party.
	//
	// example:
	//
	// 0
	RepeatNumber *int32 `json:"RepeatNumber,omitempty" xml:"RepeatNumber,omitempty"`
	// The number of retries allowed.
	//
	// example:
	//
	// 3
	RetryCount *int32 `json:"RetryCount,omitempty" xml:"RetryCount,omitempty"`
	// The retry interval. Unit: seconds.
	//
	// example:
	//
	// 60
	RetryInterval *int32 `json:"RetryInterval,omitempty" xml:"RetryInterval,omitempty"`
	// The protocol of the source stream.
	//
	// example:
	//
	// flv
	SourceProtocol *string `json:"SourceProtocol,omitempty" xml:"SourceProtocol,omitempty"`
	// The type of the source stream. Valid values:
	//
	// 	- live: a live stream
	//
	// 	- vod: a list of ApsaraVideo VOD resources
	//
	// 	- url: a list of video resources from a third party
	//
	// example:
	//
	// vod
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The source URLs.
	SourceUrls []*string `json:"SourceUrls,omitempty" xml:"SourceUrls,omitempty" type:"Repeated"`
	// The start time of the task. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-12-04T09:16:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task ID.
	//
	// example:
	//
	// fb0d4ac7-c7e3-4978-9743-0bf2f6e8****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task name.
	//
	// example:
	//
	// taskname
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s DescribeLivePullToPushListResponseBodyTaskListTaskInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLivePullToPushListResponseBodyTaskListTaskInfo) GoString() string {
	return s.String()
}

func (s *DescribeLivePullToPushListResponseBodyTaskListTaskInfo) GetCallbackURL() *string {
	return s.CallbackURL
}

func (s *DescribeLivePullToPushListResponseBodyTaskListTaskInfo) GetDstUrl() *string {
	return s.DstUrl
}

func (s *DescribeLivePullToPushListResponseBodyTaskListTaskInfo) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLivePullToPushListResponseBodyTaskListTaskInfo) GetFileIndex() *int32 {
	return s.FileIndex
}

func (s *DescribeLivePullToPushListResponseBodyTaskListTaskInfo) GetOffset() *int32 {
	return s.Offset
}

func (s *DescribeLivePullToPushListResponseBodyTaskListTaskInfo) GetRepeatNumber() *int32 {
	return s.RepeatNumber
}

func (s *DescribeLivePullToPushListResponseBodyTaskListTaskInfo) GetRetryCount() *int32 {
	return s.RetryCount
}

func (s *DescribeLivePullToPushListResponseBodyTaskListTaskInfo) GetRetryInterval() *int32 {
	return s.RetryInterval
}

func (s *DescribeLivePullToPushListResponseBodyTaskListTaskInfo) GetSourceProtocol() *string {
	return s.SourceProtocol
}

func (s *DescribeLivePullToPushListResponseBodyTaskListTaskInfo) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeLivePullToPushListResponseBodyTaskListTaskInfo) GetSourceUrls() []*string {
	return s.SourceUrls
}

func (s *DescribeLivePullToPushListResponseBodyTaskListTaskInfo) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLivePullToPushListResponseBodyTaskListTaskInfo) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeLivePullToPushListResponseBodyTaskListTaskInfo) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeLivePullToPushListResponseBodyTaskListTaskInfo) SetCallbackURL(v string) *DescribeLivePullToPushListResponseBodyTaskListTaskInfo {
	s.CallbackURL = &v
	return s
}

func (s *DescribeLivePullToPushListResponseBodyTaskListTaskInfo) SetDstUrl(v string) *DescribeLivePullToPushListResponseBodyTaskListTaskInfo {
	s.DstUrl = &v
	return s
}

func (s *DescribeLivePullToPushListResponseBodyTaskListTaskInfo) SetEndTime(v string) *DescribeLivePullToPushListResponseBodyTaskListTaskInfo {
	s.EndTime = &v
	return s
}

func (s *DescribeLivePullToPushListResponseBodyTaskListTaskInfo) SetFileIndex(v int32) *DescribeLivePullToPushListResponseBodyTaskListTaskInfo {
	s.FileIndex = &v
	return s
}

func (s *DescribeLivePullToPushListResponseBodyTaskListTaskInfo) SetOffset(v int32) *DescribeLivePullToPushListResponseBodyTaskListTaskInfo {
	s.Offset = &v
	return s
}

func (s *DescribeLivePullToPushListResponseBodyTaskListTaskInfo) SetRepeatNumber(v int32) *DescribeLivePullToPushListResponseBodyTaskListTaskInfo {
	s.RepeatNumber = &v
	return s
}

func (s *DescribeLivePullToPushListResponseBodyTaskListTaskInfo) SetRetryCount(v int32) *DescribeLivePullToPushListResponseBodyTaskListTaskInfo {
	s.RetryCount = &v
	return s
}

func (s *DescribeLivePullToPushListResponseBodyTaskListTaskInfo) SetRetryInterval(v int32) *DescribeLivePullToPushListResponseBodyTaskListTaskInfo {
	s.RetryInterval = &v
	return s
}

func (s *DescribeLivePullToPushListResponseBodyTaskListTaskInfo) SetSourceProtocol(v string) *DescribeLivePullToPushListResponseBodyTaskListTaskInfo {
	s.SourceProtocol = &v
	return s
}

func (s *DescribeLivePullToPushListResponseBodyTaskListTaskInfo) SetSourceType(v string) *DescribeLivePullToPushListResponseBodyTaskListTaskInfo {
	s.SourceType = &v
	return s
}

func (s *DescribeLivePullToPushListResponseBodyTaskListTaskInfo) SetSourceUrls(v []*string) *DescribeLivePullToPushListResponseBodyTaskListTaskInfo {
	s.SourceUrls = v
	return s
}

func (s *DescribeLivePullToPushListResponseBodyTaskListTaskInfo) SetStartTime(v string) *DescribeLivePullToPushListResponseBodyTaskListTaskInfo {
	s.StartTime = &v
	return s
}

func (s *DescribeLivePullToPushListResponseBodyTaskListTaskInfo) SetTaskId(v string) *DescribeLivePullToPushListResponseBodyTaskListTaskInfo {
	s.TaskId = &v
	return s
}

func (s *DescribeLivePullToPushListResponseBodyTaskListTaskInfo) SetTaskName(v string) *DescribeLivePullToPushListResponseBodyTaskListTaskInfo {
	s.TaskName = &v
	return s
}

func (s *DescribeLivePullToPushListResponseBodyTaskListTaskInfo) Validate() error {
	return dara.Validate(s)
}
