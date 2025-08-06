// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVideoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVideoIds(v string) *DeleteVideoRequest
	GetVideoIds() *string
}

type DeleteVideoRequest struct {
	// The list of video IDs. Separate multiple IDs with commas (,). A maximum of 20 IDs can be specified. You can obtain a video ID in one of the following ways:
	//
	// 	- If the video is uploaded by using the [ApsaraVideo VOD console](https://vod.console.aliyun.com), log on to the console and choose **Media Files*	- > **Audio/Video*	- to view the ID of the video.
	//
	// 	- If the video is uploaded by calling the [CreateUploadVideo](https://help.aliyun.com/document_detail/55407.html) operation, the video ID is the VideoId value in the response.
	//
	// 	- You can also call the [SearchMedia](https://help.aliyun.com/document_detail/86044.html) operation to obtain the video ID, which is the VideoId value in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// e44ebf1147hdsa2d2adbea8b****,e44ebf11oj984adbea8****,73ab850b4f6f4b6e91d24d81d54****
	VideoIds *string `json:"VideoIds,omitempty" xml:"VideoIds,omitempty"`
}

func (s DeleteVideoRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVideoRequest) GoString() string {
	return s.String()
}

func (s *DeleteVideoRequest) GetVideoIds() *string {
	return s.VideoIds
}

func (s *DeleteVideoRequest) SetVideoIds(v string) *DeleteVideoRequest {
	s.VideoIds = &v
	return s
}

func (s *DeleteVideoRequest) Validate() error {
	return dara.Validate(s)
}
