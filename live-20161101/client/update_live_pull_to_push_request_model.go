// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLivePullToPushRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthKey(v string) *UpdateLivePullToPushRequest
	GetAuthKey() *string
	SetCallbackUrl(v string) *UpdateLivePullToPushRequest
	GetCallbackUrl() *string
	SetEndTime(v string) *UpdateLivePullToPushRequest
	GetEndTime() *string
	SetFileIndex(v int32) *UpdateLivePullToPushRequest
	GetFileIndex() *int32
	SetNotifyItemSwitch(v string) *UpdateLivePullToPushRequest
	GetNotifyItemSwitch() *string
	SetOffset(v int32) *UpdateLivePullToPushRequest
	GetOffset() *int32
	SetOwnerId(v int64) *UpdateLivePullToPushRequest
	GetOwnerId() *int64
	SetRegion(v string) *UpdateLivePullToPushRequest
	GetRegion() *string
	SetRegionId(v string) *UpdateLivePullToPushRequest
	GetRegionId() *string
	SetRepeatNumber(v int32) *UpdateLivePullToPushRequest
	GetRepeatNumber() *int32
	SetReqAuth(v string) *UpdateLivePullToPushRequest
	GetReqAuth() *string
	SetSourceUrls(v []*string) *UpdateLivePullToPushRequest
	GetSourceUrls() []*string
	SetStartTime(v string) *UpdateLivePullToPushRequest
	GetStartTime() *string
	SetTaskId(v string) *UpdateLivePullToPushRequest
	GetTaskId() *string
}

type UpdateLivePullToPushRequest struct {
	AuthKey *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	// The callback URL. Default value: empty.
	//
	// > - The URL that receives task-related callbacks.
	//
	// > - Maximum length: 2000 characters.
	//
	// > - If this parameter is not specified, task events are not sent as callbacks.
	//
	// > - The update takes effect only when the next event is triggered.
	//
	// example:
	//
	// https://callback*****.com
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// The end time of the task.
	//
	// > - Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC).
	//
	// > - EndTime must be later than StartTime.
	//
	// > - EndTime must be later than the current time.
	//
	// > - If the task has ended, the update does not take effect.
	//
	// example:
	//
	// 2024-08-27T14:30:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The video index. Default value: 0.
	//
	// > The update must be performed when the task is stopped and takes effect after the task is restarted.
	//
	// example:
	//
	// 0
	FileIndex        *int32  `json:"FileIndex,omitempty" xml:"FileIndex,omitempty"`
	NotifyItemSwitch *string `json:"NotifyItemSwitch,omitempty" xml:"NotifyItemSwitch,omitempty"`
	// The start offset of the video file, in seconds. Valid values: greater than 0.
	//
	// > - Specifies the position to start reading from, relative to the first frame.
	//
	// > - This parameter applies only to video-on-demand or third-party video streams.
	//
	// > - This parameter takes effect only when the first video in the playlist is played.
	//
	// > - The update must be performed when the task is stopped and takes effect after the task is restarted.
	//
	// example:
	//
	// 2
	Offset  *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region where the task is started. Valid values:
	//
	// - ap-southeast-1 (Singapore)
	//
	// - ap-southeast-5 (Indonesia)
	//
	// - cn-beijing (Beijing)
	//
	// - cn-shanghai (Shanghai)
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of times playback repeats after the playlist finishes. Valid values:
	//
	// - 0 (default): No repeat playback.
	//
	// - -1: Loops indefinitely.
	//
	// - Other positive integers: The number of times playback repeats after the playlist finishes.
	//
	// > - This parameter applies only to video-on-demand or third-party video streams.
	//
	// > - The update takes effect immediately.
	//
	// example:
	//
	// 0
	RepeatNumber *int32  `json:"RepeatNumber,omitempty" xml:"RepeatNumber,omitempty"`
	ReqAuth      *string `json:"ReqAuth,omitempty" xml:"ReqAuth,omitempty"`
	// The list of source stream URLs.
	//
	// > - For the live type, only one complete live streaming URL is supported.
	//
	// > - For the vod and url types, up to 30 URLs can be specified.
	//
	// > - The live type supports RTMP, SRT, and HTTP-FLV protocols.
	//
	// > - For the vod type, specify ApsaraVideo VOD media asset IDs.
	//
	// > - The url type supports MP4 and HTTP-FLV protocols.
	//
	// > - For live source streams, the update takes effect immediately. For video file source streams, the update takes effect after the currently playing video ends, and playback restarts from the beginning of the updated video list.
	//
	// > - The update must be performed when the task is stopped and takes effect after the task is restarted.
	//
	// example:
	//
	// changedtesturl
	SourceUrls []*string `json:"SourceUrls,omitempty" xml:"SourceUrls,omitempty" type:"Repeated"`
	// The start time of the task.
	//
	// > - Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC).
	//
	// > - If the task has already started running, the update does not take effect.
	//
	// example:
	//
	// 2024-08-23T15:30:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// fd245384-4067-4f91-9d75-9666a6bc****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateLivePullToPushRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLivePullToPushRequest) GoString() string {
	return s.String()
}

func (s *UpdateLivePullToPushRequest) GetAuthKey() *string {
	return s.AuthKey
}

func (s *UpdateLivePullToPushRequest) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *UpdateLivePullToPushRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateLivePullToPushRequest) GetFileIndex() *int32 {
	return s.FileIndex
}

func (s *UpdateLivePullToPushRequest) GetNotifyItemSwitch() *string {
	return s.NotifyItemSwitch
}

func (s *UpdateLivePullToPushRequest) GetOffset() *int32 {
	return s.Offset
}

func (s *UpdateLivePullToPushRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateLivePullToPushRequest) GetRegion() *string {
	return s.Region
}

func (s *UpdateLivePullToPushRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateLivePullToPushRequest) GetRepeatNumber() *int32 {
	return s.RepeatNumber
}

func (s *UpdateLivePullToPushRequest) GetReqAuth() *string {
	return s.ReqAuth
}

func (s *UpdateLivePullToPushRequest) GetSourceUrls() []*string {
	return s.SourceUrls
}

func (s *UpdateLivePullToPushRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateLivePullToPushRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateLivePullToPushRequest) SetAuthKey(v string) *UpdateLivePullToPushRequest {
	s.AuthKey = &v
	return s
}

func (s *UpdateLivePullToPushRequest) SetCallbackUrl(v string) *UpdateLivePullToPushRequest {
	s.CallbackUrl = &v
	return s
}

func (s *UpdateLivePullToPushRequest) SetEndTime(v string) *UpdateLivePullToPushRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateLivePullToPushRequest) SetFileIndex(v int32) *UpdateLivePullToPushRequest {
	s.FileIndex = &v
	return s
}

func (s *UpdateLivePullToPushRequest) SetNotifyItemSwitch(v string) *UpdateLivePullToPushRequest {
	s.NotifyItemSwitch = &v
	return s
}

func (s *UpdateLivePullToPushRequest) SetOffset(v int32) *UpdateLivePullToPushRequest {
	s.Offset = &v
	return s
}

func (s *UpdateLivePullToPushRequest) SetOwnerId(v int64) *UpdateLivePullToPushRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateLivePullToPushRequest) SetRegion(v string) *UpdateLivePullToPushRequest {
	s.Region = &v
	return s
}

func (s *UpdateLivePullToPushRequest) SetRegionId(v string) *UpdateLivePullToPushRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateLivePullToPushRequest) SetRepeatNumber(v int32) *UpdateLivePullToPushRequest {
	s.RepeatNumber = &v
	return s
}

func (s *UpdateLivePullToPushRequest) SetReqAuth(v string) *UpdateLivePullToPushRequest {
	s.ReqAuth = &v
	return s
}

func (s *UpdateLivePullToPushRequest) SetSourceUrls(v []*string) *UpdateLivePullToPushRequest {
	s.SourceUrls = v
	return s
}

func (s *UpdateLivePullToPushRequest) SetStartTime(v string) *UpdateLivePullToPushRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateLivePullToPushRequest) SetTaskId(v string) *UpdateLivePullToPushRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateLivePullToPushRequest) Validate() error {
	return dara.Validate(s)
}
