// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaRefreshJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *GetMediaRefreshJobsRequest
	GetMediaId() *string
	SetMediaRefreshJobId(v string) *GetMediaRefreshJobsRequest
	GetMediaRefreshJobId() *string
}

type GetMediaRefreshJobsRequest struct {
	// The audio or video ID. The value is the `MediaIds` parameter passed when you call the [Submit audio and video refresh or prefetch tasks](https://help.aliyun.com/document_detail/431095.html) API operation. You can specify only one audio or video ID.
	//
	// If you do not specify this parameter, the task information of all audio and video files under the `MediaRefreshJobId` is queried. If you specify this parameter, only the task information of the specified audio or video file under the `MediaRefreshJobId` is queried.
	//
	// example:
	//
	// ca3a8f6e4957b658067095869****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The task ID of the audio or video refresh or prefetch task. The value is the MediaRefreshJobId parameter returned by the [Submit audio and video refresh or prefetch tasks](https://help.aliyun.com/document_detail/431095.html) API operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 41d465e31957****
	MediaRefreshJobId *string `json:"MediaRefreshJobId,omitempty" xml:"MediaRefreshJobId,omitempty"`
}

func (s GetMediaRefreshJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMediaRefreshJobsRequest) GoString() string {
	return s.String()
}

func (s *GetMediaRefreshJobsRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *GetMediaRefreshJobsRequest) GetMediaRefreshJobId() *string {
	return s.MediaRefreshJobId
}

func (s *GetMediaRefreshJobsRequest) SetMediaId(v string) *GetMediaRefreshJobsRequest {
	s.MediaId = &v
	return s
}

func (s *GetMediaRefreshJobsRequest) SetMediaRefreshJobId(v string) *GetMediaRefreshJobsRequest {
	s.MediaRefreshJobId = &v
	return s
}

func (s *GetMediaRefreshJobsRequest) Validate() error {
	return dara.Validate(s)
}
