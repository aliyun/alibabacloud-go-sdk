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
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2019-01-23T12:40:12Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of the page to return. You can specify a page number to return data from the specified page. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2019-01-23T12:35:12Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the audio or video file. You can use one of the following methods to obtain the ID of the file:
	//
	// 	- Log on to the [ApsaraVideo VOD](https://vod.console.aliyun.com) console. In the left-side navigation pane, choose **Media Files*	- > **Audio/Video**. On the Video and Audio page, view the ID of the audio or video file. This method is applicable to files that are uploaded by using the ApsaraVideo VOD console.
	//
	// 	- Obtain the value of VideoId from the response to the [CreateUploadVideo](https://help.aliyun.com/document_detail/55407.html) operation that you call to obtain the upload URL and credential.
	//
	// 	- Obtain the value of VideoId by calling the [SearchMedia](https://help.aliyun.com/document_detail/86044.html) operation. This method is applicable to files that have been uploaded.
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
