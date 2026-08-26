// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLivePullToPushShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthKey(v string) *CreateLivePullToPushShrinkRequest
	GetAuthKey() *string
	SetCallbackUrl(v string) *CreateLivePullToPushShrinkRequest
	GetCallbackUrl() *string
	SetDstUrl(v string) *CreateLivePullToPushShrinkRequest
	GetDstUrl() *string
	SetEndTime(v string) *CreateLivePullToPushShrinkRequest
	GetEndTime() *string
	SetFileIndex(v int32) *CreateLivePullToPushShrinkRequest
	GetFileIndex() *int32
	SetNotifyItemSwitch(v string) *CreateLivePullToPushShrinkRequest
	GetNotifyItemSwitch() *string
	SetOffset(v int32) *CreateLivePullToPushShrinkRequest
	GetOffset() *int32
	SetOwnerId(v int64) *CreateLivePullToPushShrinkRequest
	GetOwnerId() *int64
	SetRegion(v string) *CreateLivePullToPushShrinkRequest
	GetRegion() *string
	SetRegionId(v string) *CreateLivePullToPushShrinkRequest
	GetRegionId() *string
	SetRepeatNumber(v int32) *CreateLivePullToPushShrinkRequest
	GetRepeatNumber() *int32
	SetReqAuth(v string) *CreateLivePullToPushShrinkRequest
	GetReqAuth() *string
	SetRetryCount(v int32) *CreateLivePullToPushShrinkRequest
	GetRetryCount() *int32
	SetRetryInterval(v int32) *CreateLivePullToPushShrinkRequest
	GetRetryInterval() *int32
	SetSourceProtocol(v string) *CreateLivePullToPushShrinkRequest
	GetSourceProtocol() *string
	SetSourceType(v string) *CreateLivePullToPushShrinkRequest
	GetSourceType() *string
	SetSourceUrlsShrink(v string) *CreateLivePullToPushShrinkRequest
	GetSourceUrlsShrink() *string
	SetStartTime(v string) *CreateLivePullToPushShrinkRequest
	GetStartTime() *string
	SetTaskName(v string) *CreateLivePullToPushShrinkRequest
	GetTaskName() *string
}

type CreateLivePullToPushShrinkRequest struct {
	AuthKey *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	// HTTP callback URL. Default value: empty.
	//
	// > - The URL that receives task-related callbacks.
	//
	// > - Maximum length is 2000 characters.
	//
	// > - If this parameter is not specified, no task event callbacks will be sent.
	//
	// example:
	//
	// https://callback*****.com
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// Destination URL address for pushing the stream.
	//
	// > - The rtmp protocol is supported.
	//
	// > - Maximum length is 2000 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// rtmp://pushtest.********.aliyunlive.com/pulltest493/pulltest-w434
	DstUrl *string `json:"DstUrl,omitempty" xml:"DstUrl,omitempty"`
	// Task end time.
	//
	// > - Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC time).
	//
	// > - EndTime must be later than StartTime.
	//
	// > - EndTime must be later than the current time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024-08-27T14:30:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// File index. Starts playback from the nth file.
	//
	// example:
	//
	// 0
	FileIndex        *int32  `json:"FileIndex,omitempty" xml:"FileIndex,omitempty"`
	NotifyItemSwitch *string `json:"NotifyItemSwitch,omitempty" xml:"NotifyItemSwitch,omitempty"`
	// Start offset. The offset value from the beginning of the video file. Unit: seconds. Valid values: greater than 0.
	//
	// > - Indicates the position to start reading from, relative to the first frame (applies to the first video).
	//
	// > - This parameter applies only to VOD or third-party video streams.
	//
	// example:
	//
	// 2
	Offset  *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies the region where the task is launched. Valid values:
	//
	// - ap-southeast-1 (Singapore)
	//
	// - ap-southeast-5 (Indonesia)
	//
	// - cn-beijing (Beijing)
	//
	// - cn-shanghai (Shanghai)
	//
	// - cn-shenzhen (Shenzhen)
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Number of times to repeat playback after the initial playback is complete. Valid values:
	//
	// - 0 (default): no repeat playback.
	//
	// - -1: loop indefinitely.
	//
	// - Other positive integers: number of times to repeat playback after the initial playback is complete.
	//
	// > This parameter applies only to VOD or third-party video streams.
	//
	// example:
	//
	// 0
	RepeatNumber *int32  `json:"RepeatNumber,omitempty" xml:"RepeatNumber,omitempty"`
	ReqAuth      *string `json:"ReqAuth,omitempty" xml:"ReqAuth,omitempty"`
	// Number of retries. Default value: 3.
	//
	// example:
	//
	// 3
	RetryCount *int32 `json:"RetryCount,omitempty" xml:"RetryCount,omitempty"`
	// Retry interval, in seconds. Valid values: [60, 300]. Default value: 60 seconds.
	//
	// example:
	//
	// 60
	RetryInterval *int32 `json:"RetryInterval,omitempty" xml:"RetryInterval,omitempty"`
	// Source stream protocol name.
	//
	// Valid values:
	//
	// - rtmp
	//
	// - srt
	//
	// - http-flv
	//
	// - hls
	//
	// > This parameter is **required only when the SourceType parameter is set to live**, and is invalid when the value is vod or url.
	//
	// example:
	//
	// rtmp
	SourceProtocol *string `json:"SourceProtocol,omitempty" xml:"SourceProtocol,omitempty"`
	// Source stream type. Valid values:
	//
	// - live: live stream.
	//
	// - vod: ApsaraVideo VOD resource.
	//
	// - url: third-party video file resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// live
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// List of source stream URL addresses.
	//
	// > - For the live type, only one complete live playback URL is supported.
	//
	// > - For the vod and url types, a maximum of 30 URLs can be specified.
	//
	// > - The live type supports: rtmp, srt, and http-flv protocols.
	//
	// > - For the vod type, specify ApsaraVideo VOD media asset IDs.
	//
	// > - The url type supports: mp4 and http-flv protocols.
	//
	// This parameter is required.
	SourceUrlsShrink *string `json:"SourceUrls,omitempty" xml:"SourceUrls,omitempty"`
	// Task start time.
	//
	// > - Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC time).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024-08-26T10:30:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Task name, used to support fuzzy query. Default value: "".
	//
	// example:
	//
	// test
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreateLivePullToPushShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLivePullToPushShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateLivePullToPushShrinkRequest) GetAuthKey() *string {
	return s.AuthKey
}

func (s *CreateLivePullToPushShrinkRequest) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *CreateLivePullToPushShrinkRequest) GetDstUrl() *string {
	return s.DstUrl
}

func (s *CreateLivePullToPushShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateLivePullToPushShrinkRequest) GetFileIndex() *int32 {
	return s.FileIndex
}

func (s *CreateLivePullToPushShrinkRequest) GetNotifyItemSwitch() *string {
	return s.NotifyItemSwitch
}

func (s *CreateLivePullToPushShrinkRequest) GetOffset() *int32 {
	return s.Offset
}

func (s *CreateLivePullToPushShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateLivePullToPushShrinkRequest) GetRegion() *string {
	return s.Region
}

func (s *CreateLivePullToPushShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateLivePullToPushShrinkRequest) GetRepeatNumber() *int32 {
	return s.RepeatNumber
}

func (s *CreateLivePullToPushShrinkRequest) GetReqAuth() *string {
	return s.ReqAuth
}

func (s *CreateLivePullToPushShrinkRequest) GetRetryCount() *int32 {
	return s.RetryCount
}

func (s *CreateLivePullToPushShrinkRequest) GetRetryInterval() *int32 {
	return s.RetryInterval
}

func (s *CreateLivePullToPushShrinkRequest) GetSourceProtocol() *string {
	return s.SourceProtocol
}

func (s *CreateLivePullToPushShrinkRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateLivePullToPushShrinkRequest) GetSourceUrlsShrink() *string {
	return s.SourceUrlsShrink
}

func (s *CreateLivePullToPushShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateLivePullToPushShrinkRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateLivePullToPushShrinkRequest) SetAuthKey(v string) *CreateLivePullToPushShrinkRequest {
	s.AuthKey = &v
	return s
}

func (s *CreateLivePullToPushShrinkRequest) SetCallbackUrl(v string) *CreateLivePullToPushShrinkRequest {
	s.CallbackUrl = &v
	return s
}

func (s *CreateLivePullToPushShrinkRequest) SetDstUrl(v string) *CreateLivePullToPushShrinkRequest {
	s.DstUrl = &v
	return s
}

func (s *CreateLivePullToPushShrinkRequest) SetEndTime(v string) *CreateLivePullToPushShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *CreateLivePullToPushShrinkRequest) SetFileIndex(v int32) *CreateLivePullToPushShrinkRequest {
	s.FileIndex = &v
	return s
}

func (s *CreateLivePullToPushShrinkRequest) SetNotifyItemSwitch(v string) *CreateLivePullToPushShrinkRequest {
	s.NotifyItemSwitch = &v
	return s
}

func (s *CreateLivePullToPushShrinkRequest) SetOffset(v int32) *CreateLivePullToPushShrinkRequest {
	s.Offset = &v
	return s
}

func (s *CreateLivePullToPushShrinkRequest) SetOwnerId(v int64) *CreateLivePullToPushShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateLivePullToPushShrinkRequest) SetRegion(v string) *CreateLivePullToPushShrinkRequest {
	s.Region = &v
	return s
}

func (s *CreateLivePullToPushShrinkRequest) SetRegionId(v string) *CreateLivePullToPushShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLivePullToPushShrinkRequest) SetRepeatNumber(v int32) *CreateLivePullToPushShrinkRequest {
	s.RepeatNumber = &v
	return s
}

func (s *CreateLivePullToPushShrinkRequest) SetReqAuth(v string) *CreateLivePullToPushShrinkRequest {
	s.ReqAuth = &v
	return s
}

func (s *CreateLivePullToPushShrinkRequest) SetRetryCount(v int32) *CreateLivePullToPushShrinkRequest {
	s.RetryCount = &v
	return s
}

func (s *CreateLivePullToPushShrinkRequest) SetRetryInterval(v int32) *CreateLivePullToPushShrinkRequest {
	s.RetryInterval = &v
	return s
}

func (s *CreateLivePullToPushShrinkRequest) SetSourceProtocol(v string) *CreateLivePullToPushShrinkRequest {
	s.SourceProtocol = &v
	return s
}

func (s *CreateLivePullToPushShrinkRequest) SetSourceType(v string) *CreateLivePullToPushShrinkRequest {
	s.SourceType = &v
	return s
}

func (s *CreateLivePullToPushShrinkRequest) SetSourceUrlsShrink(v string) *CreateLivePullToPushShrinkRequest {
	s.SourceUrlsShrink = &v
	return s
}

func (s *CreateLivePullToPushShrinkRequest) SetStartTime(v string) *CreateLivePullToPushShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *CreateLivePullToPushShrinkRequest) SetTaskName(v string) *CreateLivePullToPushShrinkRequest {
	s.TaskName = &v
	return s
}

func (s *CreateLivePullToPushShrinkRequest) Validate() error {
	return dara.Validate(s)
}
