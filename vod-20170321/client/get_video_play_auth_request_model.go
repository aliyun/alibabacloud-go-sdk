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
	SetVideoId(v string) *GetVideoPlayAuthRequest
	GetVideoId() *string
}

type GetVideoPlayAuthRequest struct {
	// The API version. Set the value to **1.0.0**.
	//
	// example:
	//
	// 1.0.0
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	// The validity period of the playback credential. Unit: **seconds**. You cannot obtain the playback URL of a video by using a credential that has expired. A new credential is required.
	//
	// 	- Default value: **100**.
	//
	// 	- Valid values: `[100,3000]`.
	//
	// example:
	//
	// 100
	AuthInfoTimeout *int64 `json:"AuthInfoTimeout,omitempty" xml:"AuthInfoTimeout,omitempty"`
	// The ID of the media file. You can specify only one ID. You can use one of the following methods to obtain the ID of the file:
	//
	// 	- Log on to the [ApsaraVideo VOD](https://vod.console.aliyun.com) console. In the left-side navigation pane, choose **Media Files*	- > **Audio/Video**. On the Video and Audio page, view the ID of the media file. This method is applicable to files that are uploaded by using the ApsaraVideo VOD console.
	//
	// 	- Obtain the value of the VideoId parameter from the response to the [CreateUploadVideo](https://help.aliyun.com/document_detail/55407.html) operation.
	//
	// 	- Obtain the value of the VideoId parameter from the response to the [SearchMedia](https://help.aliyun.com/document_detail/86044.html) operation. This method is applicable to files that have been uploaded.
	//
	// This parameter is required.
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

func (s *GetVideoPlayAuthRequest) SetVideoId(v string) *GetVideoPlayAuthRequest {
	s.VideoId = &v
	return s
}

func (s *GetVideoPlayAuthRequest) Validate() error {
	return dara.Validate(s)
}
