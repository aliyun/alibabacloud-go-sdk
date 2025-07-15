// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLivePullToPushResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentFileIndex(v int32) *DescribeLivePullToPushResponseBody
	GetCurrentFileIndex() *int32
	SetCurrentOffset(v int32) *DescribeLivePullToPushResponseBody
	GetCurrentOffset() *int32
	SetDescription(v string) *DescribeLivePullToPushResponseBody
	GetDescription() *string
	SetRequestId(v string) *DescribeLivePullToPushResponseBody
	GetRequestId() *string
	SetRetCode(v int32) *DescribeLivePullToPushResponseBody
	GetRetCode() *int32
	SetTaskExitReason(v string) *DescribeLivePullToPushResponseBody
	GetTaskExitReason() *string
	SetTaskExitTime(v int32) *DescribeLivePullToPushResponseBody
	GetTaskExitTime() *int32
	SetTaskId(v string) *DescribeLivePullToPushResponseBody
	GetTaskId() *string
	SetTaskInfo(v *DescribeLivePullToPushResponseBodyTaskInfo) *DescribeLivePullToPushResponseBody
	GetTaskInfo() *DescribeLivePullToPushResponseBodyTaskInfo
	SetTaskInvalidReason(v string) *DescribeLivePullToPushResponseBody
	GetTaskInvalidReason() *string
	SetTaskStatus(v int32) *DescribeLivePullToPushResponseBody
	GetTaskStatus() *int32
}

type DescribeLivePullToPushResponseBody struct {
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
	// The error description.
	//
	// example:
	//
	// OK
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3271ACD2-F143-1204-AFDB-9A87C131****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The code that is returned for the request.
	//
	// >
	//
	// 	- 0 is returned if the request is normal.
	//
	// 	- For information about codes that are returned when exceptions occur, see the following Error codes table.
	//
	// example:
	//
	// 0
	RetCode *int32 `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
	// The reason why the task is stopped.
	//
	// 	- TriggerByUser: You proactively stopped the task.
	//
	// 	- OverEndTime: The specified end time was exceeded.
	//
	// >  This parameter is returned only if the task is stopped.
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
	// 1724740200
	TaskExitTime *int32 `json:"TaskExitTime,omitempty" xml:"TaskExitTime,omitempty"`
	// The task ID.
	//
	// example:
	//
	// fd245384-4067-4f91-9d75-9666a6bc9****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The information about the task.
	TaskInfo *DescribeLivePullToPushResponseBodyTaskInfo `json:"TaskInfo,omitempty" xml:"TaskInfo,omitempty" type:"Struct"`
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
	// UnknownError
	TaskInvalidReason *string `json:"TaskInvalidReason,omitempty" xml:"TaskInvalidReason,omitempty"`
	// The current status of the task.
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
	// 1
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s DescribeLivePullToPushResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLivePullToPushResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLivePullToPushResponseBody) GetCurrentFileIndex() *int32 {
	return s.CurrentFileIndex
}

func (s *DescribeLivePullToPushResponseBody) GetCurrentOffset() *int32 {
	return s.CurrentOffset
}

func (s *DescribeLivePullToPushResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeLivePullToPushResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLivePullToPushResponseBody) GetRetCode() *int32 {
	return s.RetCode
}

func (s *DescribeLivePullToPushResponseBody) GetTaskExitReason() *string {
	return s.TaskExitReason
}

func (s *DescribeLivePullToPushResponseBody) GetTaskExitTime() *int32 {
	return s.TaskExitTime
}

func (s *DescribeLivePullToPushResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeLivePullToPushResponseBody) GetTaskInfo() *DescribeLivePullToPushResponseBodyTaskInfo {
	return s.TaskInfo
}

func (s *DescribeLivePullToPushResponseBody) GetTaskInvalidReason() *string {
	return s.TaskInvalidReason
}

func (s *DescribeLivePullToPushResponseBody) GetTaskStatus() *int32 {
	return s.TaskStatus
}

func (s *DescribeLivePullToPushResponseBody) SetCurrentFileIndex(v int32) *DescribeLivePullToPushResponseBody {
	s.CurrentFileIndex = &v
	return s
}

func (s *DescribeLivePullToPushResponseBody) SetCurrentOffset(v int32) *DescribeLivePullToPushResponseBody {
	s.CurrentOffset = &v
	return s
}

func (s *DescribeLivePullToPushResponseBody) SetDescription(v string) *DescribeLivePullToPushResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeLivePullToPushResponseBody) SetRequestId(v string) *DescribeLivePullToPushResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLivePullToPushResponseBody) SetRetCode(v int32) *DescribeLivePullToPushResponseBody {
	s.RetCode = &v
	return s
}

func (s *DescribeLivePullToPushResponseBody) SetTaskExitReason(v string) *DescribeLivePullToPushResponseBody {
	s.TaskExitReason = &v
	return s
}

func (s *DescribeLivePullToPushResponseBody) SetTaskExitTime(v int32) *DescribeLivePullToPushResponseBody {
	s.TaskExitTime = &v
	return s
}

func (s *DescribeLivePullToPushResponseBody) SetTaskId(v string) *DescribeLivePullToPushResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribeLivePullToPushResponseBody) SetTaskInfo(v *DescribeLivePullToPushResponseBodyTaskInfo) *DescribeLivePullToPushResponseBody {
	s.TaskInfo = v
	return s
}

func (s *DescribeLivePullToPushResponseBody) SetTaskInvalidReason(v string) *DescribeLivePullToPushResponseBody {
	s.TaskInvalidReason = &v
	return s
}

func (s *DescribeLivePullToPushResponseBody) SetTaskStatus(v int32) *DescribeLivePullToPushResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *DescribeLivePullToPushResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLivePullToPushResponseBodyTaskInfo struct {
	// The HTTP callback URL.
	//
	// example:
	//
	// https://callback*****.com
	CallbackURL *string `json:"CallbackURL,omitempty" xml:"CallbackURL,omitempty"`
	// The destination URL to which the stream is relayed.
	//
	// example:
	//
	// rtmp://pushtest.********.aliyunlive.com/pulltest493/pulltest-w434
	DstUrl *string `json:"DstUrl,omitempty" xml:"DstUrl,omitempty"`
	// The end time of the task. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-08-27T14:30:00Z
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
	// 2
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
	// RTMP
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
	// live
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The source URLs.
	//
	// example:
	//
	// rtmp://pulltest.****.aliyunlive.com/pulltest493/pulltest-w434
	SourceUrls []*string `json:"SourceUrls,omitempty" xml:"SourceUrls,omitempty" type:"Repeated"`
	// The start time of the task. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-08-26T10:30:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task ID.
	//
	// example:
	//
	// fd245384-4067-4f91-9d75-9666a6bc9****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task name.
	//
	// example:
	//
	// test
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s DescribeLivePullToPushResponseBodyTaskInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLivePullToPushResponseBodyTaskInfo) GoString() string {
	return s.String()
}

func (s *DescribeLivePullToPushResponseBodyTaskInfo) GetCallbackURL() *string {
	return s.CallbackURL
}

func (s *DescribeLivePullToPushResponseBodyTaskInfo) GetDstUrl() *string {
	return s.DstUrl
}

func (s *DescribeLivePullToPushResponseBodyTaskInfo) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLivePullToPushResponseBodyTaskInfo) GetFileIndex() *int32 {
	return s.FileIndex
}

func (s *DescribeLivePullToPushResponseBodyTaskInfo) GetOffset() *int32 {
	return s.Offset
}

func (s *DescribeLivePullToPushResponseBodyTaskInfo) GetRepeatNumber() *int32 {
	return s.RepeatNumber
}

func (s *DescribeLivePullToPushResponseBodyTaskInfo) GetRetryCount() *int32 {
	return s.RetryCount
}

func (s *DescribeLivePullToPushResponseBodyTaskInfo) GetRetryInterval() *int32 {
	return s.RetryInterval
}

func (s *DescribeLivePullToPushResponseBodyTaskInfo) GetSourceProtocol() *string {
	return s.SourceProtocol
}

func (s *DescribeLivePullToPushResponseBodyTaskInfo) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeLivePullToPushResponseBodyTaskInfo) GetSourceUrls() []*string {
	return s.SourceUrls
}

func (s *DescribeLivePullToPushResponseBodyTaskInfo) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLivePullToPushResponseBodyTaskInfo) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeLivePullToPushResponseBodyTaskInfo) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeLivePullToPushResponseBodyTaskInfo) SetCallbackURL(v string) *DescribeLivePullToPushResponseBodyTaskInfo {
	s.CallbackURL = &v
	return s
}

func (s *DescribeLivePullToPushResponseBodyTaskInfo) SetDstUrl(v string) *DescribeLivePullToPushResponseBodyTaskInfo {
	s.DstUrl = &v
	return s
}

func (s *DescribeLivePullToPushResponseBodyTaskInfo) SetEndTime(v string) *DescribeLivePullToPushResponseBodyTaskInfo {
	s.EndTime = &v
	return s
}

func (s *DescribeLivePullToPushResponseBodyTaskInfo) SetFileIndex(v int32) *DescribeLivePullToPushResponseBodyTaskInfo {
	s.FileIndex = &v
	return s
}

func (s *DescribeLivePullToPushResponseBodyTaskInfo) SetOffset(v int32) *DescribeLivePullToPushResponseBodyTaskInfo {
	s.Offset = &v
	return s
}

func (s *DescribeLivePullToPushResponseBodyTaskInfo) SetRepeatNumber(v int32) *DescribeLivePullToPushResponseBodyTaskInfo {
	s.RepeatNumber = &v
	return s
}

func (s *DescribeLivePullToPushResponseBodyTaskInfo) SetRetryCount(v int32) *DescribeLivePullToPushResponseBodyTaskInfo {
	s.RetryCount = &v
	return s
}

func (s *DescribeLivePullToPushResponseBodyTaskInfo) SetRetryInterval(v int32) *DescribeLivePullToPushResponseBodyTaskInfo {
	s.RetryInterval = &v
	return s
}

func (s *DescribeLivePullToPushResponseBodyTaskInfo) SetSourceProtocol(v string) *DescribeLivePullToPushResponseBodyTaskInfo {
	s.SourceProtocol = &v
	return s
}

func (s *DescribeLivePullToPushResponseBodyTaskInfo) SetSourceType(v string) *DescribeLivePullToPushResponseBodyTaskInfo {
	s.SourceType = &v
	return s
}

func (s *DescribeLivePullToPushResponseBodyTaskInfo) SetSourceUrls(v []*string) *DescribeLivePullToPushResponseBodyTaskInfo {
	s.SourceUrls = v
	return s
}

func (s *DescribeLivePullToPushResponseBodyTaskInfo) SetStartTime(v string) *DescribeLivePullToPushResponseBodyTaskInfo {
	s.StartTime = &v
	return s
}

func (s *DescribeLivePullToPushResponseBodyTaskInfo) SetTaskId(v string) *DescribeLivePullToPushResponseBodyTaskInfo {
	s.TaskId = &v
	return s
}

func (s *DescribeLivePullToPushResponseBodyTaskInfo) SetTaskName(v string) *DescribeLivePullToPushResponseBodyTaskInfo {
	s.TaskName = &v
	return s
}

func (s *DescribeLivePullToPushResponseBodyTaskInfo) Validate() error {
	return dara.Validate(s)
}
