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
	// The audio or video ID, which is the `MediaIds` value specified when calling the [SubmitMediaRefreshJob](https://help.aliyun.com/document_detail/431095.html) operation. Only one audio or video ID can be specified.
	//
	// If this parameter is not specified, task information for all audio or video files under the specified `MediaRefreshJobId` is returned. If this parameter is specified, only the task information for the specified audio or video ID under the `MediaRefreshJobId` is returned.
	//
	// example:
	//
	// ca3a8f6e4957b658067095869****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The ID of the audio or video purge or prefetch task. This is the value of MediaRefreshJobId returned by the [SubmitMediaRefreshJob](https://help.aliyun.com/document_detail/431095.html) operation.
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
