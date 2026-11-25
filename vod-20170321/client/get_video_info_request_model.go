// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReferenceId(v string) *GetVideoInfoRequest
	GetReferenceId() *string
	SetVideoId(v string) *GetVideoInfoRequest
	GetVideoId() *string
}

type GetVideoInfoRequest struct {
	// The custom ID. The ID must be 6 to 64 characters in length and can contain only lowercase letters, uppercase letters, digits, hyphens, and underscores. The ID must be unique per user.
	//
	// example:
	//
	// 123-123
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
	// The audio or video ID. Only one audio or video ID is supported. You can obtain the ID by using one of the following methods:
	//
	// - For videos uploaded through the console, log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Media Assets*	- > **Audio/Video*	- to view the audio or video ID.
	//
	// - When you call the [CreateUploadVideo](https://help.aliyun.com/document_detail/55407.html) operation to obtain the upload URL and credential, the audio or video ID is the value of the VideoId response parameter.
	//
	// - After the audio or video file is uploaded, you can call the [SearchMedia](https://help.aliyun.com/document_detail/86044.html) operation to query the audio or video ID, which is the value of the VideoId response parameter.
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

func (s *GetVideoInfoRequest) GetReferenceId() *string {
	return s.ReferenceId
}

func (s *GetVideoInfoRequest) GetVideoId() *string {
	return s.VideoId
}

func (s *GetVideoInfoRequest) SetReferenceId(v string) *GetVideoInfoRequest {
	s.ReferenceId = &v
	return s
}

func (s *GetVideoInfoRequest) SetVideoId(v string) *GetVideoInfoRequest {
	s.VideoId = &v
	return s
}

func (s *GetVideoInfoRequest) Validate() error {
	return dara.Validate(s)
}
