// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTranscodeSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVideoIds(v string) *GetTranscodeSummaryRequest
	GetVideoIds() *string
}

type GetTranscodeSummaryRequest struct {
	// The audio or video IDs. You can specify a maximum of 10 IDs, separated by commas (,). You can obtain the audio or video ID by using the following methods:
	//
	// - For audio or video files uploaded through the console, log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Media Files*	- > **Audio/Video*	- to view the audio or video ID.
	//
	// - Obtain the video ID from the value of the VideoId parameter returned by the [CreateUploadVideo](https://help.aliyun.com/document_detail/55407.html) operation when you request an upload URL and credential.
	//
	// - After the audio or video file is uploaded, call the [SearchMedia](https://help.aliyun.com/document_detail/86044.html) operation to query the audio or video ID, which is the value of the VideoId parameter in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// "d4860fcc6ae9fed52e8938244****,e1db68cc586644b83e562bcd94****,hhhhhhh"
	VideoIds *string `json:"VideoIds,omitempty" xml:"VideoIds,omitempty"`
}

func (s GetTranscodeSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetTranscodeSummaryRequest) GetVideoIds() *string {
	return s.VideoIds
}

func (s *GetTranscodeSummaryRequest) SetVideoIds(v string) *GetTranscodeSummaryRequest {
	s.VideoIds = &v
	return s
}

func (s *GetTranscodeSummaryRequest) Validate() error {
	return dara.Validate(s)
}
