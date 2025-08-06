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
	// The ID of the audio or video file. You can specify up to 10 IDs. Separate the IDs with commas (,). You can use one of the following methods to obtain the ID:
	//
	// 	- After you upload a video in the [ApsaraVideo VOD console](https://vod.console.aliyun.com), you can log on to the ApsaraVideo VOD console and choose **Media Files*	- > **Audio/Video*	- to view the ID of the video.
	//
	// 	- Obtain the value of VideoId from the response to the [CreateUploadVideo](https://help.aliyun.com/document_detail/55407.html) operation that you call to obtain the upload URL and credential.
	//
	// 	- Obtain the value of VideoId by calling the [SearchMedia](https://help.aliyun.com/document_detail/86044.html) operation. This method is applicable to files that have been uploaded.
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
