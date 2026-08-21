// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDynamicImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicImageIds(v string) *DeleteDynamicImageRequest
	GetDynamicImageIds() *string
	SetVideoId(v string) *DeleteDynamicImageRequest
	GetVideoId() *string
}

type DeleteDynamicImageRequest struct {
	// The list of animated sticker IDs. The animated sticker ID is the value of the DynamicImageId response parameter returned by the [ListDynamicImage](https://help.aliyun.com/document_detail/180958.html) operation.
	//
	// - Separate multiple IDs with commas (,). You can specify a maximum of 10 IDs.
	//
	// - **If you do not specify this parameter, all animated stickers associated with the specified VideoId are deleted. However, if the video has more than 10 animated stickers, the deletion request is rejected.**
	//
	// example:
	//
	// beafec3834a4e52ea52042a4****,8281c8519847fd8970e79e80b6****
	DynamicImageIds *string `json:"DynamicImageIds,omitempty" xml:"DynamicImageIds,omitempty"`
	// The ID of the video associated with the animated stickers that you want to delete. You can obtain the video ID by using one of the following methods:
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
	// 2321077d460b028700ef6c2f4d****
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s DeleteDynamicImageRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDynamicImageRequest) GoString() string {
	return s.String()
}

func (s *DeleteDynamicImageRequest) GetDynamicImageIds() *string {
	return s.DynamicImageIds
}

func (s *DeleteDynamicImageRequest) GetVideoId() *string {
	return s.VideoId
}

func (s *DeleteDynamicImageRequest) SetDynamicImageIds(v string) *DeleteDynamicImageRequest {
	s.DynamicImageIds = &v
	return s
}

func (s *DeleteDynamicImageRequest) SetVideoId(v string) *DeleteDynamicImageRequest {
	s.VideoId = &v
	return s
}

func (s *DeleteDynamicImageRequest) Validate() error {
	return dara.Validate(s)
}
