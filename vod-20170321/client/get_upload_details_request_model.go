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
	// The media IDs, which are audio or video IDs (VideoId). Separate multiple IDs with commas (,). A maximum of 20 IDs are supported. You can obtain the IDs by using the following methods:
	//
	// - For audio or video files uploaded through the console, log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Media Files*	- > **Audio/Video*	- to view the audio or video ID.
	//
	// - When you call the [CreateUploadVideo](https://help.aliyun.com/document_detail/55407.html) operation to obtain the upload URL and credential, the video ID is the value of the VideoId response parameter.
	//
	// - After the audio or video file is uploaded, you can call the [SearchMedia](https://help.aliyun.com/document_detail/86044.html) operation to query the video ID, which is the value of the VideoId response parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 61ccbdb06fa83012be4d8083f6****,7d2fbc380b0e08e55f****
	MediaIds *string `json:"MediaIds,omitempty" xml:"MediaIds,omitempty"`
	// The media type. Set the value to **video*	- (audio/video).
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
