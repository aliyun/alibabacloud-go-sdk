// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVideoId(v string) *GetVideoInfoRequest
	GetVideoId() *string
}

type GetVideoInfoRequest struct {
	// The ID of the audio or video file. You can specify only one ID in each call. You can use one of the following methods to obtain the ID:
	//
	// 	- Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com). In the left-side navigation pane, choose **Media Files*	- > **Audio/Video**. On the Video and Audio page, view the ID of the audio or video file. This method is applicable to files that are uploaded by using the ApsaraVideo VOD console.
	//
	// 	- Obtain the value of VideoId from the response to the [CreateUploadVideo](https://help.aliyun.com/document_detail/55407.html) operation that you called to obtain the upload URL and credential.
	//
	// 	- Obtain the value of VideoId from the response to the [SearchMedia](https://help.aliyun.com/document_detail/86044.html) operation that you called to query the media ID after the media file is uploaded.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9b73864d75f1d231e9001cd5f8****
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetVideoInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVideoInfoRequest) GoString() string {
	return s.String()
}

func (s *GetVideoInfoRequest) GetVideoId() *string {
	return s.VideoId
}

func (s *GetVideoInfoRequest) SetVideoId(v string) *GetVideoInfoRequest {
	s.VideoId = &v
	return s
}

func (s *GetVideoInfoRequest) Validate() error {
	return dara.Validate(s)
}
