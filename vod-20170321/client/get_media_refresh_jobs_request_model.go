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
	// The ID of the media file. It is the value of the `MediaIds` parameter that you specify when you call the [RefreshMediaPlayUrls](~~RefreshMediaPlayUrls~~) operation. You can specify only one media ID.
	//
	// If you leave this parameter empty, information about all media files in the refresh or prefetch job specified by `MediaRefreshJobId` is returned. If you set this parameter, only the information about the specified media file is returned.``
	//
	// example:
	//
	// ca3a8f6e4957b658067095869****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The ID of the refresh or prefetch job. It is the value of the MediaRefreshJobId parameter that is returned from the call to the [RefreshMediaPlayUrls](~~RefreshMediaPlayUrls~~) operation.
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
