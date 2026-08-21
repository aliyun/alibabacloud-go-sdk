// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoPlayAuthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiVersion(v string) *GetVideoPlayAuthRequest
	GetApiVersion() *string
	SetAuthInfoTimeout(v int64) *GetVideoPlayAuthRequest
	GetAuthInfoTimeout() *int64
	SetReferenceId(v string) *GetVideoPlayAuthRequest
	GetReferenceId() *string
	SetVideoId(v string) *GetVideoPlayAuthRequest
	GetVideoId() *string
}

type GetVideoPlayAuthRequest struct {
	// The API version number. Set the value to **1.0.0**.
	//
	// example:
	//
	// 1.0.0
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	// The expiration time of the playback credential. Unit: **seconds**. If the credential expires, the playback URL cannot be obtained. You must obtain a new credential.
	//
	// - Default value: **100**.
	//
	// - Valid values: `[100,3000]`.
	//
	// example:
	//
	// 100
	AuthInfoTimeout *int64 `json:"AuthInfoTimeout,omitempty" xml:"AuthInfoTimeout,omitempty"`
	// The custom ID. Only lowercase letters, uppercase letters, digits, hyphens, and underscores are supported. Length: 6 to 64 characters. The ID is unique per user.
	//
	// example:
	//
	// 123-123
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
	// The audio or video ID. Only a single audio or video ID is supported. You can obtain the ID by using the following methods:
	//
	// - For videos uploaded through the console, log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Media Files*	- > **Audio/Video*	- to view the audio or video ID.
	//
	// - When uploading audio or video files by calling the [CreateUploadVideo](https://help.aliyun.com/document_detail/55407.html) operation, the audio or video ID is the value of the VideoId response parameter.
	//
	// - After the audio or video file is uploaded, you can call the [SearchMedia](https://help.aliyun.com/document_detail/86044.html) operation to query the audio or video ID, which is the value of the VideoId response parameter.
	//
	// example:
	//
	// dfde02284a5c46622a097adaf44a****
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetVideoPlayAuthRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVideoPlayAuthRequest) GoString() string {
	return s.String()
}

func (s *GetVideoPlayAuthRequest) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *GetVideoPlayAuthRequest) GetAuthInfoTimeout() *int64 {
	return s.AuthInfoTimeout
}

func (s *GetVideoPlayAuthRequest) GetReferenceId() *string {
	return s.ReferenceId
}

func (s *GetVideoPlayAuthRequest) GetVideoId() *string {
	return s.VideoId
}

func (s *GetVideoPlayAuthRequest) SetApiVersion(v string) *GetVideoPlayAuthRequest {
	s.ApiVersion = &v
	return s
}

func (s *GetVideoPlayAuthRequest) SetAuthInfoTimeout(v int64) *GetVideoPlayAuthRequest {
	s.AuthInfoTimeout = &v
	return s
}

func (s *GetVideoPlayAuthRequest) SetReferenceId(v string) *GetVideoPlayAuthRequest {
	s.ReferenceId = &v
	return s
}

func (s *GetVideoPlayAuthRequest) SetVideoId(v string) *GetVideoPlayAuthRequest {
	s.VideoId = &v
	return s
}

func (s *GetVideoPlayAuthRequest) Validate() error {
	return dara.Validate(s)
}
