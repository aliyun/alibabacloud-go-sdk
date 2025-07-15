// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLivePullToPushRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallbackUrl(v string) *CreateLivePullToPushRequest
	GetCallbackUrl() *string
	SetDstUrl(v string) *CreateLivePullToPushRequest
	GetDstUrl() *string
	SetEndTime(v string) *CreateLivePullToPushRequest
	GetEndTime() *string
	SetFileIndex(v int32) *CreateLivePullToPushRequest
	GetFileIndex() *int32
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
	SourceUrls []*string `json:"SourceUrls,omitempty" xml:"SourceUrls,omitempty" type:"Repeated"`
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

func (s CreateLivePullToPushRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLivePullToPushRequest) GoString() string {
	return s.String()
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
