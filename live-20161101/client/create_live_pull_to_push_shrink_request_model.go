// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLivePullToPushShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallbackUrl(v string) *CreateLivePullToPushShrinkRequest
	GetCallbackUrl() *string
	SetDstUrl(v string) *CreateLivePullToPushShrinkRequest
	GetDstUrl() *string
	SetEndTime(v string) *CreateLivePullToPushShrinkRequest
	GetEndTime() *string
	SetFileIndex(v int32) *CreateLivePullToPushShrinkRequest
	GetFileIndex() *int32
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
	// The HTTP callback URL. By default, this parameter is left empty.
	//
	// >
	//
	// 	- The URL is used to receive callbacks related to the task.
	//
	// 	- The URL can be up to 2,000 characters in length.
	//
	// 	- If you do not specify this parameter, no callbacks are returned for events related to the task.
	//
	// example:
	//
	// https://callback*****.com
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// The destination URL to which the stream is relayed.
	//
	// >
	//
	// 	- The supported protocol for the URL is RTMP.
	//
	// 	- The URL can be up to 2,000 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// rtmp://pushtest.********.aliyunlive.com/pulltest493/pulltest-w434
	DstUrl *string `json:"DstUrl,omitempty" xml:"DstUrl,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 2024-08-27T14:30:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The file index, which specifies the sequence of the file where the playback starts.
	//
	// example:
	//
	// 0
	FileIndex *int32 `json:"FileIndex,omitempty" xml:"FileIndex,omitempty"`
	// The offset of the position where the system starts to read the video resource. Unit: seconds. Valid values: positive numbers.
	//
	// >
	//
	// 	- This parameter indicates an offset from the first frame of the first video resource in the list.
	//
	// 	- This parameter is applicable to only video resources from ApsaraVideo VOD or a third party.
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
	// >  This parameter is applicable to only video resources from ApsaraVideo VOD or a third party.
	//
	// example:
	//
	// 0
	RepeatNumber *int32 `json:"RepeatNumber,omitempty" xml:"RepeatNumber,omitempty"`
	// The number of retries allowed. Default value: 3.
	//
	// example:
	//
	// 3
	RetryCount *int32 `json:"RetryCount,omitempty" xml:"RetryCount,omitempty"`
	// The retry interval. Unit: seconds. Valid values: [60,300]. Default value: 60.
	//
	// example:
	//
	// 60
	RetryInterval *int32 `json:"RetryInterval,omitempty" xml:"RetryInterval,omitempty"`
	// The protocol of the source stream.
	//
	// Valid values:
	//
	// 	- rtmp
	//
	// 	- rtsp
	//
	// 	- srt
	//
	// 	- http-flv
	//
	// 	- flv
	//
	// >  This parameter is required if you set the **SourceType*	- parameter to live, but does not take effect if you set the SourceType parameter to vod or url.
	//
	// example:
	//
	// rtmp
	SourceProtocol *string `json:"SourceProtocol,omitempty" xml:"SourceProtocol,omitempty"`
	// The type of the source stream. Valid values:
	//
	// 	- live: a live stream
	//
	// 	- vod: a list of ApsaraVideo VOD resources
	//
	// 	- url: a list of video resources from a third party
	//
	// This parameter is required.
	//
	// example:
	//
	// live
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// testurls
	SourceUrlsShrink *string `json:"SourceUrls,omitempty" xml:"SourceUrls,omitempty"`
	// The start time of the task.
	//
	// >
	//
	// 	- Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// 	- The time range specified by the StartTime and EndTime parameters cannot exceed seven days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024-08-26T10:30:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the task. Default value: "". Fuzzy search for task names is supported.
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
