// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIImageInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVideoId(v string) *ListAIImageInfoRequest
	GetVideoId() *string
}

type ListAIImageInfoRequest struct {
	// The video ID. You can obtain the video ID by using one of the following methods:
	//
	// - For videos uploaded through the console, log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Media Files*	- > **Audio/Video*	- to view the video ID.
	//
	// - Obtain the video ID from the value of the VideoId response parameter when you call the [CreateUploadVideo](https://help.aliyun.com/document_detail/55407.html) operation to obtain the upload URL and credential.
	//
	// - After the video is uploaded, call the [SearchMedia](https://help.aliyun.com/document_detail/86044.html) operation to query the video ID, which is the value of the VideoId response parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 357a8748c5789d2726e6436aa****
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s ListAIImageInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAIImageInfoRequest) GoString() string {
	return s.String()
}

func (s *ListAIImageInfoRequest) GetVideoId() *string {
	return s.VideoId
}

func (s *ListAIImageInfoRequest) SetVideoId(v string) *ListAIImageInfoRequest {
	s.VideoId = &v
	return s
}

func (s *ListAIImageInfoRequest) Validate() error {
	return dara.Validate(s)
}
