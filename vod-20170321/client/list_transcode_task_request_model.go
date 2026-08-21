// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTranscodeTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListTranscodeTaskRequest
	GetEndTime() *string
	SetPageNo(v int32) *ListTranscodeTaskRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListTranscodeTaskRequest
	GetPageSize() *int32
	SetStartTime(v string) *ListTranscodeTaskRequest
	GetStartTime() *string
	SetVideoId(v string) *ListTranscodeTaskRequest
	GetVideoId() *string
}

type ListTranscodeTaskRequest struct {
	// The end time, which must be later than the start time. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2019-01-23T12:40:12Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number. Specify this parameter to return data starting from the specified page. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The start time. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2019-01-23T12:35:12Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The audio or video ID. You can obtain the ID by using one of the following methods:
	//
	// - For audio or video files uploaded in the console, log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Media Files*	- > **Audio/Video*	- to view the ID.
	//
	// - Obtain the video ID from the VideoId response parameter when you call the [CreateUploadVideo](https://help.aliyun.com/document_detail/55407.html) operation to obtain the upload URL and credential.
	//
	// - After the audio or video file is uploaded, call the [SearchMedia](https://help.aliyun.com/document_detail/86044.html) operation to query the audio or video ID, which is the VideoId value in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// d4860fcc6a5*****bce9fed52e893824
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s ListTranscodeTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTranscodeTaskRequest) GoString() string {
	return s.String()
}

func (s *ListTranscodeTaskRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListTranscodeTaskRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListTranscodeTaskRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTranscodeTaskRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListTranscodeTaskRequest) GetVideoId() *string {
	return s.VideoId
}

func (s *ListTranscodeTaskRequest) SetEndTime(v string) *ListTranscodeTaskRequest {
	s.EndTime = &v
	return s
}

func (s *ListTranscodeTaskRequest) SetPageNo(v int32) *ListTranscodeTaskRequest {
	s.PageNo = &v
	return s
}

func (s *ListTranscodeTaskRequest) SetPageSize(v int32) *ListTranscodeTaskRequest {
	s.PageSize = &v
	return s
}

func (s *ListTranscodeTaskRequest) SetStartTime(v string) *ListTranscodeTaskRequest {
	s.StartTime = &v
	return s
}

func (s *ListTranscodeTaskRequest) SetVideoId(v string) *ListTranscodeTaskRequest {
	s.VideoId = &v
	return s
}

func (s *ListTranscodeTaskRequest) Validate() error {
	return dara.Validate(s)
}
