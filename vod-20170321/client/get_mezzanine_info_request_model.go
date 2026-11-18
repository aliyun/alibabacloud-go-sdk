// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMezzanineInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionType(v string) *GetMezzanineInfoRequest
	GetAdditionType() *string
	SetAuthTimeout(v int64) *GetMezzanineInfoRequest
	GetAuthTimeout() *int64
	SetOutputType(v string) *GetMezzanineInfoRequest
	GetOutputType() *string
	SetReferenceId(v string) *GetMezzanineInfoRequest
	GetReferenceId() *string
	SetVideoId(v string) *GetMezzanineInfoRequest
	GetVideoId() *string
}

type GetMezzanineInfoRequest struct {
	// The type of additional information. Separate multiple values with commas (,). By default, only basic information is returned. Valid values:
	//
	// - **video**: the video stream information.
	//
	// - **audio**: the audio stream information.
	//
	// example:
	//
	// video
	AdditionType *string `json:"AdditionType,omitempty" xml:"AdditionType,omitempty"`
	// The expiration time of the signature for the FileURL (source file URL). Unit: seconds. Default value: **3600**. Minimum value: **1**.
	//
	//  - If OutputType is set to **cdn**:
	//
	//     - The FileURL expires periodically only if URL authentication is enabled. Otherwise, the FileURL is permanently valid.
	//
	//     - Minimum value: **1**.
	//
	//     - Maximum value: unlimited.
	//
	//     - Default value: **3600*	- if not specified.
	//
	// - If OutputType is set to **oss**:
	//
	//     - The FileURL expires periodically only if the storage permission is set to private. Otherwise, the FileURL is permanently valid.
	//
	//     - Minimum value: **1**.
	//
	//     - Maximum value: To reduce security risks of the origin server, the maximum value is **2592000*	- (30 days) when the audio or video file is stored in an ApsaraVideo VOD bucket, and **129600*	- (36 hours) when the audio or video file is stored in a user-owned OSS bucket.
	//
	//     - Default value: **3600*	- if not specified.
	//
	// example:
	//
	// 3600
	AuthTimeout *int64 `json:"AuthTimeout,omitempty" xml:"AuthTimeout,omitempty"`
	// The output URL type. Valid values:
	//
	// - **oss**: the back-to-origin URL.
	//
	// - **cdn*	- (default): the accelerated URL.
	//
	// > If the bucket where the source file is stored is of the in type, only the OSS URL is returned.
	//
	// example:
	//
	// oss
	OutputType *string `json:"OutputType,omitempty" xml:"OutputType,omitempty"`
	// The custom ID. The ID must be 6 to 64 characters in length and can contain only lowercase letters, uppercase letters, digits, hyphens (-), and underscores (_). The ID must be unique per user.
	//
	// example:
	//
	// 123-123
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
	// The audio or video ID. You can obtain the ID by using one of the following methods:
	//
	// - For an audio or video file uploaded in the console, log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Media Assets*	- > **Audio/Video*	- to view the video ID.
	//
	// - When you call the [Obtain the upload URL and credential for a video](https://help.aliyun.com/document_detail/55407.html) operation to obtain the upload URL and credential, the audio or video ID is the value of the VideoId response parameter.
	//
	// - After a video is uploaded, you can call the [Search media information](https://help.aliyun.com/document_detail/86044.html) operation to query the audio or video ID, which is the value of the VideoId response parameter.
	//
	// example:
	//
	// 1f1a6fc03ca04814031b8a6559e****
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetMezzanineInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMezzanineInfoRequest) GoString() string {
	return s.String()
}

func (s *GetMezzanineInfoRequest) GetAdditionType() *string {
	return s.AdditionType
}

func (s *GetMezzanineInfoRequest) GetAuthTimeout() *int64 {
	return s.AuthTimeout
}

func (s *GetMezzanineInfoRequest) GetOutputType() *string {
	return s.OutputType
}

func (s *GetMezzanineInfoRequest) GetReferenceId() *string {
	return s.ReferenceId
}

func (s *GetMezzanineInfoRequest) GetVideoId() *string {
	return s.VideoId
}

func (s *GetMezzanineInfoRequest) SetAdditionType(v string) *GetMezzanineInfoRequest {
	s.AdditionType = &v
	return s
}

func (s *GetMezzanineInfoRequest) SetAuthTimeout(v int64) *GetMezzanineInfoRequest {
	s.AuthTimeout = &v
	return s
}

func (s *GetMezzanineInfoRequest) SetOutputType(v string) *GetMezzanineInfoRequest {
	s.OutputType = &v
	return s
}

func (s *GetMezzanineInfoRequest) SetReferenceId(v string) *GetMezzanineInfoRequest {
	s.ReferenceId = &v
	return s
}

func (s *GetMezzanineInfoRequest) SetVideoId(v string) *GetMezzanineInfoRequest {
	s.VideoId = &v
	return s
}

func (s *GetMezzanineInfoRequest) Validate() error {
	return dara.Validate(s)
}
