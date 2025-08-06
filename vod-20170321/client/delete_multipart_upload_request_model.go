// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMultipartUploadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *DeleteMultipartUploadRequest
	GetMediaId() *string
	SetMediaType(v string) *DeleteMultipartUploadRequest
	GetMediaType() *string
	SetOwnerAccount(v string) *DeleteMultipartUploadRequest
	GetOwnerAccount() *string
}

type DeleteMultipartUploadRequest struct {
	// The ID of the media file (VideoId). You can use one of the following methods to obtain the ID:
	//
	// 	- After you upload a video in the [ApsaraVideo VOD console](https://vod.console.aliyun.com), you can log on to the ApsaraVideo VOD console and choose **Media Files*	- > **Audio/Video*	- to view the ID of the video.
	//
	// 	- Obtain the value of VideoId from the response to the [CreateUploadVideo](https://help.aliyun.com/document_detail/55407.html) operation that you called to obtain the upload URL and credential.
	//
	// 	- Obtain the value of VideoId from the response to the [SearchMedia](https://help.aliyun.com/document_detail/86044.html) operation that you called to query media information after the audio or video file is uploaded.
	//
	// This parameter is required.
	//
	// example:
	//
	// 61ccbdb06fa3012be4d8083f6****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The type of the media file. Set the value to **video**. video indicates audio and video files.
	//
	// This parameter is required.
	//
	// example:
	//
	// video
	MediaType    *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
}

func (s DeleteMultipartUploadRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultipartUploadRequest) GoString() string {
	return s.String()
}

func (s *DeleteMultipartUploadRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *DeleteMultipartUploadRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *DeleteMultipartUploadRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteMultipartUploadRequest) SetMediaId(v string) *DeleteMultipartUploadRequest {
	s.MediaId = &v
	return s
}

func (s *DeleteMultipartUploadRequest) SetMediaType(v string) *DeleteMultipartUploadRequest {
	s.MediaType = &v
	return s
}

func (s *DeleteMultipartUploadRequest) SetOwnerAccount(v string) *DeleteMultipartUploadRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteMultipartUploadRequest) Validate() error {
	return dara.Validate(s)
}
