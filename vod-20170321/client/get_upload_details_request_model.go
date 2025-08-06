// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUploadDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaIds(v string) *GetUploadDetailsRequest
	GetMediaIds() *string
	SetMediaType(v string) *GetUploadDetailsRequest
	GetMediaType() *string
}

type GetUploadDetailsRequest struct {
	// The ID of the media file, namely, the audio or video ID. You can specify a maximum of 20 IDs. Separate multiple IDs with commas (,). You can use one of the following methods to obtain the audio or video ID:
	//
	// 	- Log on to the [ApsaraVideo VOD](https://vod.console.aliyun.com) console. In the left-side navigation pane, choose **Media Files*	- > **Audio/Video**. On the Video and Audio page, you can view the audio or video ID. Use this method if the audio or video file is uploaded by using the ApsaraVideo VOD console.
	//
	// 	- View the value of the VideoId parameter returned by the [CreateUploadVideo](https://help.aliyun.com/document_detail/55407.html) operation that you called to obtain an upload URL and credential.
	//
	// 	- View the value of the VideoId parameter returned by the [SearchMedia](https://help.aliyun.com/document_detail/86044.html) operation that you called to query media information after the audio or video file is uploaded.
	//
	// This parameter is required.
	//
	// example:
	//
	// 61ccbdb06fa83012be4d8083f6****,7d2fbc380b0e08e55f****
	MediaIds *string `json:"MediaIds,omitempty" xml:"MediaIds,omitempty"`
	// The type of the media file. Set the value to **video**, which indicates audio and video files.
	//
	// example:
	//
	// video
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
}

func (s GetUploadDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUploadDetailsRequest) GoString() string {
	return s.String()
}

func (s *GetUploadDetailsRequest) GetMediaIds() *string {
	return s.MediaIds
}

func (s *GetUploadDetailsRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *GetUploadDetailsRequest) SetMediaIds(v string) *GetUploadDetailsRequest {
	s.MediaIds = &v
	return s
}

func (s *GetUploadDetailsRequest) SetMediaType(v string) *GetUploadDetailsRequest {
	s.MediaType = &v
	return s
}

func (s *GetUploadDetailsRequest) Validate() error {
	return dara.Validate(s)
}
