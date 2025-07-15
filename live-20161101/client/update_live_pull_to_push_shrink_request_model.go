// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLivePullToPushShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallbackUrl(v string) *UpdateLivePullToPushShrinkRequest
	GetCallbackUrl() *string
	SetEndTime(v string) *UpdateLivePullToPushShrinkRequest
	GetEndTime() *string
	SetFileIndex(v int32) *UpdateLivePullToPushShrinkRequest
	GetFileIndex() *int32
	SetOffset(v int32) *UpdateLivePullToPushShrinkRequest
	GetOffset() *int32
	SetOwnerId(v int64) *UpdateLivePullToPushShrinkRequest
	GetOwnerId() *int64
	SetRegion(v string) *UpdateLivePullToPushShrinkRequest
	GetRegion() *string
	SetRegionId(v string) *UpdateLivePullToPushShrinkRequest
	GetRegionId() *string
	SetRepeatNumber(v int32) *UpdateLivePullToPushShrinkRequest
	GetRepeatNumber() *int32
	SetSourceUrlsShrink(v string) *UpdateLivePullToPushShrinkRequest
	GetSourceUrlsShrink() *string
	SetStartTime(v string) *UpdateLivePullToPushShrinkRequest
	GetStartTime() *string
	SetTaskId(v string) *UpdateLivePullToPushShrinkRequest
	GetTaskId() *string
}

type UpdateLivePullToPushShrinkRequest struct {
	// The callback URL. By default, this parameter is left empty.
	//
	// >
	//
	// 	- The URL is used to receive callbacks related to the task.
	//
	// 	- The URL can be up to 2,000 characters in length.
	//
	// 	- If you do not specify this parameter, no callbacks are returned for events related to the task.
	//
	// 	- The update takes effect for subsequent events that occur.
	//
	// example:
	//
	// https://callback*****.com
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// The end time of the task.
	//
	// >
	//
	// 	- Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// 	- The time range specified by the StartTime and EndTime parameters cannot exceed seven days.
	//
	// 	- The end time must be later than the start time.
	//
	// 	- The end time must be later than the current time.
	//
	// 	- If the task has ended, the update does not take effect.
	//
	// example:
	//
	// 2024-08-27T14:30:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The file index. Default value: 0.
	//
	// >  You can modify this parameter only if the task is stopped. The update takes effect after you restart the task.
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
	// 	- The update takes effect only for the first video in a video list.
	//
	// 	- You can modify this parameter only if the task is stopped. The update takes effect immediately.
	//
	// example:
	//
	// 2
	Offset  *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region where the task is started. Valid values:
	//
	// 	- ap-southeast-1: Singapore
	//
	// 	- ap-southeast-5: Indonesia (Jakarta)
	//
	// 	- cn-beijing: China (Beijing)
	//
	// 	- cn-shanghai: China (Shanghai)
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	Region   *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of playbacks after the first playback is complete. Valid values:
	//
	// 	- 0 (default): specifies that the video list is played only once.
	//
	// 	- \\-1: specifies that the video list is played in loop mode.
	//
	// 	- Positive integer: specifies the number of times the video list repeats after the first playback is complete.
	//
	// >
	//
	// 	- This parameter is applicable to only video resources from ApsaraVideo VOD or a third party.
	//
	// 	- The update can take effect immediately.
	//
	// example:
	//
	// 0
	RepeatNumber *int32 `json:"RepeatNumber,omitempty" xml:"RepeatNumber,omitempty"`
	// The source URLs.
	//
	// >
	//
	// 	- If SourceType is set to live, you can specify only one streaming URL.
	//
	// 	- If SourceType is set to vod or url, you can specify up to 30 IDs or URLs.
	//
	// 	- If SourceType is set to live, the supported protocols for URLs are Real-Time Messaging Protocol (RTMP), Real-Time Streaming Protocol (RTSP), Secure Reliable Transport Protocol (SRT), and HTTP-FLV.
	//
	// 	- If SourceType is set to vod, specify the IDs of media assets from ApsaraVideo VOD.
	//
	// 	- If SourceType is set to url, the supported protocols for URLs are MP4 and HTTP-FLV.
	//
	// 	- If the source is a live stream, the update takes effect immediately. If the source is a list of video resources from ApsaraVideo VOD or a third party, the update does not take effect until the playback of the current video ends. After the update takes effect, the video list starts to play from the beginning.
	//
	// 	- You can modify this parameter only if the task is stopped. The update takes effect immediately.
	//
	// example:
	//
	// changedtesturl
	SourceUrlsShrink *string `json:"SourceUrls,omitempty" xml:"SourceUrls,omitempty"`
	// The start time of the task.
	//
	// >
	//
	// 	- Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// 	- The time range specified by the StartTime and EndTime parameters cannot exceed seven days.
	//
	// 	- If the task has already started, the update does not take effect.
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

func (s UpdateLivePullToPushShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLivePullToPushShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateLivePullToPushShrinkRequest) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *UpdateLivePullToPushShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateLivePullToPushShrinkRequest) GetFileIndex() *int32 {
	return s.FileIndex
}

func (s *UpdateLivePullToPushShrinkRequest) GetOffset() *int32 {
	return s.Offset
}

func (s *UpdateLivePullToPushShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateLivePullToPushShrinkRequest) GetRegion() *string {
	return s.Region
}

func (s *UpdateLivePullToPushShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateLivePullToPushShrinkRequest) GetRepeatNumber() *int32 {
	return s.RepeatNumber
}

func (s *UpdateLivePullToPushShrinkRequest) GetSourceUrlsShrink() *string {
	return s.SourceUrlsShrink
}

func (s *UpdateLivePullToPushShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateLivePullToPushShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateLivePullToPushShrinkRequest) SetCallbackUrl(v string) *UpdateLivePullToPushShrinkRequest {
	s.CallbackUrl = &v
	return s
}

func (s *UpdateLivePullToPushShrinkRequest) SetEndTime(v string) *UpdateLivePullToPushShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateLivePullToPushShrinkRequest) SetFileIndex(v int32) *UpdateLivePullToPushShrinkRequest {
	s.FileIndex = &v
	return s
}

func (s *UpdateLivePullToPushShrinkRequest) SetOffset(v int32) *UpdateLivePullToPushShrinkRequest {
	s.Offset = &v
	return s
}

func (s *UpdateLivePullToPushShrinkRequest) SetOwnerId(v int64) *UpdateLivePullToPushShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateLivePullToPushShrinkRequest) SetRegion(v string) *UpdateLivePullToPushShrinkRequest {
	s.Region = &v
	return s
}

func (s *UpdateLivePullToPushShrinkRequest) SetRegionId(v string) *UpdateLivePullToPushShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateLivePullToPushShrinkRequest) SetRepeatNumber(v int32) *UpdateLivePullToPushShrinkRequest {
	s.RepeatNumber = &v
	return s
}

func (s *UpdateLivePullToPushShrinkRequest) SetSourceUrlsShrink(v string) *UpdateLivePullToPushShrinkRequest {
	s.SourceUrlsShrink = &v
	return s
}

func (s *UpdateLivePullToPushShrinkRequest) SetStartTime(v string) *UpdateLivePullToPushShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateLivePullToPushShrinkRequest) SetTaskId(v string) *UpdateLivePullToPushShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateLivePullToPushShrinkRequest) Validate() error {
	return dara.Validate(s)
}
