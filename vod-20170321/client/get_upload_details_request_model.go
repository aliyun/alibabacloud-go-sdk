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
	// The media IDs, which are the audio or video IDs (VideoId). Separate multiple IDs with commas (,). You can specify up to 20 IDs. You can obtain the IDs by using one of the following methods:
	//
	// - For audio or video files uploaded through the console, log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Media Library*	- > **Audio and Video*	- to view the audio or video IDs.
	//
	// - When you call the [Create upload URLs and credentials for audio and video files](https://help.aliyun.com/document_detail/55407.html) operation to obtain the upload URL and credential, the value of the VideoId response parameter is the media ID.
	//
	// - After an audio or video file is uploaded, you can call the [Search for media information](https://help.aliyun.com/document_detail/86044.html) operation to query the media ID. The value of the VideoId response parameter is the media ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 61ccbdb06fa83012be4d8083f6****,7d2fbc380b0e08e55f****
	MediaIds *string `json:"MediaIds,omitempty" xml:"MediaIds,omitempty"`
	// The media type. Set the value to **video**, which indicates audio and video files.
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
