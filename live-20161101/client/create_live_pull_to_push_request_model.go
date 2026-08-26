// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLivePullToPushRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthKey(v string) *CreateLivePullToPushRequest
	GetAuthKey() *string
	SetCallbackUrl(v string) *CreateLivePullToPushRequest
	GetCallbackUrl() *string
	SetDstUrl(v string) *CreateLivePullToPushRequest
	GetDstUrl() *string
	SetEndTime(v string) *CreateLivePullToPushRequest
	GetEndTime() *string
	SetFileIndex(v int32) *CreateLivePullToPushRequest
	GetFileIndex() *int32
	SetNotifyItemSwitch(v string) *CreateLivePullToPushRequest
	GetNotifyItemSwitch() *string
	SetOffset(v int32) *CreateLivePullToPushRequest
	GetOffset() *int32
	SetOwnerId(v int64) *CreateLivePullToPushRequest
	GetOwnerId() *int64
	SetRegion(v string) *CreateLivePullToPushRequest
	GetRegion() *string
	SetRegionId(v string) *CreateLivePullToPushRequest
	GetRegionId() *string
	SetRepeatNumber(v int32) *CreateLivePullToPushRequest
	GetRepeatNumber() *int32
	SetReqAuth(v string) *CreateLivePullToPushRequest
	GetReqAuth() *string
	SetRetryCount(v int32) *CreateLivePullToPushRequest
	GetRetryCount() *int32
	SetRetryInterval(v int32) *CreateLivePullToPushRequest
	GetRetryInterval() *int32
	SetSourceProtocol(v string) *CreateLivePullToPushRequest
	GetSourceProtocol() *string
	SetSourceType(v string) *CreateLivePullToPushRequest
	GetSourceType() *string
	SetSourceUrls(v []*string) *CreateLivePullToPushRequest
	GetSourceUrls() []*string
	SetStartTime(v string) *CreateLivePullToPushRequest
	GetStartTime() *string
	SetTaskName(v string) *CreateLivePullToPushRequest
	GetTaskName() *string
}

type CreateLivePullToPushRequest struct {
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
	SourceUrls []*string `json:"SourceUrls,omitempty" xml:"SourceUrls,omitempty" type:"Repeated"`
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

func (s CreateLivePullToPushRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLivePullToPushRequest) GoString() string {
	return s.String()
}

func (s *CreateLivePullToPushRequest) GetAuthKey() *string {
	return s.AuthKey
}

func (s *CreateLivePullToPushRequest) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *CreateLivePullToPushRequest) GetDstUrl() *string {
	return s.DstUrl
}

func (s *CreateLivePullToPushRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateLivePullToPushRequest) GetFileIndex() *int32 {
	return s.FileIndex
}

func (s *CreateLivePullToPushRequest) GetNotifyItemSwitch() *string {
	return s.NotifyItemSwitch
}

func (s *CreateLivePullToPushRequest) GetOffset() *int32 {
	return s.Offset
}

func (s *CreateLivePullToPushRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateLivePullToPushRequest) GetRegion() *string {
	return s.Region
}

func (s *CreateLivePullToPushRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateLivePullToPushRequest) GetRepeatNumber() *int32 {
	return s.RepeatNumber
}

func (s *CreateLivePullToPushRequest) GetReqAuth() *string {
	return s.ReqAuth
}

func (s *CreateLivePullToPushRequest) GetRetryCount() *int32 {
	return s.RetryCount
}

func (s *CreateLivePullToPushRequest) GetRetryInterval() *int32 {
	return s.RetryInterval
}

func (s *CreateLivePullToPushRequest) GetSourceProtocol() *string {
	return s.SourceProtocol
}

func (s *CreateLivePullToPushRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateLivePullToPushRequest) GetSourceUrls() []*string {
	return s.SourceUrls
}

func (s *CreateLivePullToPushRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateLivePullToPushRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateLivePullToPushRequest) SetAuthKey(v string) *CreateLivePullToPushRequest {
	s.AuthKey = &v
	return s
}

func (s *CreateLivePullToPushRequest) SetCallbackUrl(v string) *CreateLivePullToPushRequest {
	s.CallbackUrl = &v
	return s
}

func (s *CreateLivePullToPushRequest) SetDstUrl(v string) *CreateLivePullToPushRequest {
	s.DstUrl = &v
	return s
}

func (s *CreateLivePullToPushRequest) SetEndTime(v string) *CreateLivePullToPushRequest {
	s.EndTime = &v
	return s
}

func (s *CreateLivePullToPushRequest) SetFileIndex(v int32) *CreateLivePullToPushRequest {
	s.FileIndex = &v
	return s
}

func (s *CreateLivePullToPushRequest) SetNotifyItemSwitch(v string) *CreateLivePullToPushRequest {
	s.NotifyItemSwitch = &v
	return s
}

func (s *CreateLivePullToPushRequest) SetOffset(v int32) *CreateLivePullToPushRequest {
	s.Offset = &v
	return s
}

func (s *CreateLivePullToPushRequest) SetOwnerId(v int64) *CreateLivePullToPushRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateLivePullToPushRequest) SetRegion(v string) *CreateLivePullToPushRequest {
	s.Region = &v
	return s
}

func (s *CreateLivePullToPushRequest) SetRegionId(v string) *CreateLivePullToPushRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLivePullToPushRequest) SetRepeatNumber(v int32) *CreateLivePullToPushRequest {
	s.RepeatNumber = &v
	return s
}

func (s *CreateLivePullToPushRequest) SetReqAuth(v string) *CreateLivePullToPushRequest {
	s.ReqAuth = &v
	return s
}

func (s *CreateLivePullToPushRequest) SetRetryCount(v int32) *CreateLivePullToPushRequest {
	s.RetryCount = &v
	return s
}

func (s *CreateLivePullToPushRequest) SetRetryInterval(v int32) *CreateLivePullToPushRequest {
	s.RetryInterval = &v
	return s
}

func (s *CreateLivePullToPushRequest) SetSourceProtocol(v string) *CreateLivePullToPushRequest {
	s.SourceProtocol = &v
	return s
}

func (s *CreateLivePullToPushRequest) SetSourceType(v string) *CreateLivePullToPushRequest {
	s.SourceType = &v
	return s
}

func (s *CreateLivePullToPushRequest) SetSourceUrls(v []*string) *CreateLivePullToPushRequest {
	s.SourceUrls = v
	return s
}

func (s *CreateLivePullToPushRequest) SetStartTime(v string) *CreateLivePullToPushRequest {
	s.StartTime = &v
	return s
}

func (s *CreateLivePullToPushRequest) SetTaskName(v string) *CreateLivePullToPushRequest {
	s.TaskName = &v
	return s
}

func (s *CreateLivePullToPushRequest) Validate() error {
	return dara.Validate(s)
}
