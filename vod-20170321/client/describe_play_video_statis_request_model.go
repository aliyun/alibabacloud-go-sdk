// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlayVideoStatisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribePlayVideoStatisRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribePlayVideoStatisRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribePlayVideoStatisRequest
	GetStartTime() *string
	SetVideoId(v string) *DescribePlayVideoStatisRequest
	GetVideoId() *string
}

type DescribePlayVideoStatisRequest struct {
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// >  The end time must be later than the start time. The interval between the start time and the end time cannot exceed 180 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2016-06-30T13:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2016-06-29T13:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the video. You can specify only one ID. You can use one of the following methods to obtain the ID:
	//
	// 	- Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com). In the left-side navigation pane, choose **Media Files*	- > **Audio/Video**. On the page that appears, view the video ID.
	//
	// 	- Obtain the video ID from the response to the [CreateUploadVideo](~~CreateUploadVideo~~) operation that you call to obtain the upload URL and credential.
	//
	// 	- Obtain the video ID from the response to the [SearchMedia](~~SearchMedia~~) operation that you call to query the audio or video file.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2a8d4cb9ecbb487681473****aba8fda
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s DescribePlayVideoStatisRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayVideoStatisRequest) GoString() string {
	return s.String()
}

func (s *DescribePlayVideoStatisRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribePlayVideoStatisRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePlayVideoStatisRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribePlayVideoStatisRequest) GetVideoId() *string {
	return s.VideoId
}

func (s *DescribePlayVideoStatisRequest) SetEndTime(v string) *DescribePlayVideoStatisRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePlayVideoStatisRequest) SetOwnerId(v int64) *DescribePlayVideoStatisRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePlayVideoStatisRequest) SetStartTime(v string) *DescribePlayVideoStatisRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePlayVideoStatisRequest) SetVideoId(v string) *DescribePlayVideoStatisRequest {
	s.VideoId = &v
	return s
}

func (s *DescribePlayVideoStatisRequest) Validate() error {
	return dara.Validate(s)
}
