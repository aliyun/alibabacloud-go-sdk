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
	// The current effective playlist sequence offset.
	//
	// example:
	//
	// 0
	CurrentFileIndex *int32 `json:"CurrentFileIndex,omitempty" xml:"CurrentFileIndex,omitempty"`
	// The current effective video playback offset.
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
	// a05e6b15-15af-405b-a4a2-0152245d****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The return code.
	//
	// > - "0" is returned in normal cases.
	//
	// > - For error cases, refer to the error code list below.
	//
	// example:
	//
	// 0
	RetCode *int32 `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
	// The reason why the task exited. Valid values:
	//
	// - TriggerByUser: The task was actively ended by the user.
	//
	// - OverEndTime: The preset end time was exceeded.
	//
	// > This parameter is returned only when the task is in the exited state.
	//
	// example:
	//
	// TriggerByUser
	TaskExitReason *string `json:"TaskExitReason,omitempty" xml:"TaskExitReason,omitempty"`
	// The time when the task exited. The value is a UNIX timestamp in seconds.
	//
	// > This parameter is returned only when the task is in the exited state.
	//
	// example:
	//
	// 1726354625
	TaskExitTime *int32 `json:"TaskExitTime,omitempty" xml:"TaskExitTime,omitempty"`
	// The ID of the node returned when you create task.
	//
	// example:
	//
	// fb0d4ac7-c7e3-4978-9743-0bf2f6e8****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task information.
	TaskInfo *DescribeLivePullToPushResponseBodyTaskInfo `json:"TaskInfo,omitempty" xml:"TaskInfo,omitempty" type:"Struct"`
	// The reason why the task stopped running. Valid values:
	//
	// - PullStreamFailed: Source stream pulling is abnormal. Retrying.
	//
	// - PushStreamFailed: Destination stream pushing is abnormal. Retrying.
	//
	// - UnknownError: Unknown error.
	//
	// > This parameter is returned only when the task is in the stopped state.
	//
	// example:
	//
	// PullStreamFailed
	TaskInvalidReason *string `json:"TaskInvalidReason,omitempty" xml:"TaskInvalidReason,omitempty"`
	// The current status of the task. Valid values:
	//
	// - 0: Not started (the start time has not been reached).
	//
	// - 1: Running normally (stream pulling and pushing are both normal).
	//
	// - 2: Running abnormally.
	//
	// - 3: Stopped (stream pulling or pushing is abnormal, or the task was actively stopped by calling an API operation).
	//
	// - -1: Exited.
	//
	// example:
	//
	// 0
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
	if s.TaskInfo != nil {
		if err := s.TaskInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLivePullToPushResponseBodyTaskInfo struct {
	AuthKey *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	// The HTTP callback URL.
	//
	// example:
	//
	// https://callback*****.com
	CallbackURL *string `json:"CallbackURL,omitempty" xml:"CallbackURL,omitempty"`
	// The destination ingest URL.
	//
	// example:
	//
	// rtmp://pushtest.********.aliyunlive.com/pulltest493/pulltest-w434
	DstUrl *string `json:"DstUrl,omitempty" xml:"DstUrl,omitempty"`
	// The end time of the task. Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC).
	//
	// example:
	//
	// 2024-08-27T14:30:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The file index. Playback starts from the nth file.
	//
	// example:
	//
	// 0
	FileIndex        *int32  `json:"FileIndex,omitempty" xml:"FileIndex,omitempty"`
	NotifyItemSwitch *string `json:"NotifyItemSwitch,omitempty" xml:"NotifyItemSwitch,omitempty"`
	// The start offset of the video file. Unit: seconds. The value must be greater than 0.
	//
	// > - Indicates the position from which reading starts, relative to the first frame.
	//
	// > - This parameter is valid only for video-on-demand resources or video files.
	//
	// example:
	//
	// 0
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The number of times playback repeats after completion. Valid values:
	//
	// - 0 (default): No repeat playback.
	//
	// - -1: Infinite loop.
	//
	// - Other positive integers: the number of times playback repeats after completion.
	//
	// > This parameter applies only to video-on-demand or third-party video streams.
	//
	// example:
	//
	// 0
	RepeatNumber *int32  `json:"RepeatNumber,omitempty" xml:"RepeatNumber,omitempty"`
	ReqAuth      *string `json:"ReqAuth,omitempty" xml:"ReqAuth,omitempty"`
	// The number of retries.
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
	// The source stream protocol name.
	//
	// example:
	//
	// RTMP
	SourceProtocol *string `json:"SourceProtocol,omitempty" xml:"SourceProtocol,omitempty"`
	// The source stream type. Valid values:
	//
	// - live: live stream.
	//
	// - vod: ApsaraVideo VOD resource.
	//
	// - url: third-party video file resource.
	//
	// example:
	//
	// vod
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The source stream URL.
	//
	// example:
	//
	// rtmp://pulltest.****.aliyunlive.com/pulltest493/pulltest-w434
	SourceUrls []*string `json:"SourceUrls,omitempty" xml:"SourceUrls,omitempty" type:"Repeated"`
	// The start time of the task. Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC).
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

func (s *DescribeLivePullToPushResponseBodyTaskInfo) GetAuthKey() *string {
	return s.AuthKey
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

func (s *DescribeLivePullToPushResponseBodyTaskInfo) GetNotifyItemSwitch() *string {
	return s.NotifyItemSwitch
}

func (s *DescribeLivePullToPushResponseBodyTaskInfo) GetOffset() *int32 {
	return s.Offset
}

func (s *DescribeLivePullToPushResponseBodyTaskInfo) GetRepeatNumber() *int32 {
	return s.RepeatNumber
}

func (s *DescribeLivePullToPushResponseBodyTaskInfo) GetReqAuth() *string {
	return s.ReqAuth
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

func (s *DescribeLivePullToPushResponseBodyTaskInfo) SetAuthKey(v string) *DescribeLivePullToPushResponseBodyTaskInfo {
	s.AuthKey = &v
	return s
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

func (s *DescribeLivePullToPushResponseBodyTaskInfo) SetNotifyItemSwitch(v string) *DescribeLivePullToPushResponseBodyTaskInfo {
	s.NotifyItemSwitch = &v
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

func (s *DescribeLivePullToPushResponseBodyTaskInfo) SetReqAuth(v string) *DescribeLivePullToPushResponseBodyTaskInfo {
	s.ReqAuth = &v
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
