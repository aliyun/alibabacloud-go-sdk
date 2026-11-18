// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthTimeout(v int64) *GetImageInfoRequest
	GetAuthTimeout() *int64
	SetImageId(v string) *GetImageInfoRequest
	GetImageId() *string
	SetOutputType(v string) *GetImageInfoRequest
	GetOutputType() *string
}

type GetImageInfoRequest struct {
	// The expiration time of the image access URL. Unit: seconds.
	//
	// - If OutputType is set to cdn:
	//
	//     - The image URL expires only if URL authentication is enabled. Otherwise, the URL is permanently valid.
	//
	//     - Minimum value: 1.
	//
	//     - Maximum value: unlimited.
	//
	//     - Default value: If this parameter is not specified, the default validity period configured for URL authentication is used.
	//
	// - If OutputType is set to oss:
	//
	//     - The playback URL expires only if the storage permission is set to private. Otherwise, the URL is permanently valid.
	//
	//     - Minimum value: 1.
	//
	//     - Maximum value: To reduce security risks of the origin server, the maximum value is **2592000*	- (30 days) if the image is stored in an ApsaraVideo VOD bucket, and **129600*	- (36 hours) if the image is stored in a custom OSS bucket.
	//
	//     - Default value: If this parameter is not specified, the value is 3600.
	//
	// example:
	//
	// 3600
	AuthTimeout *int64 `json:"AuthTimeout,omitempty" xml:"AuthTimeout,omitempty"`
	// The image ID. You can obtain the image ID by using one of the following methods:
	//
	// - Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com/) and choose **Media Assets*	- > **Images*	- to view the ID.
	//
	// - Call the [CreateUploadImage](~~CreateUploadImage~~) operation to obtain the upload URL and credential. The image ID is returned in the response.
	//
	// - Call the [SearchMedia](~~SearchMedia~~) operation to query images. The image ID is returned in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3e34733b40b9a96ccf5c1ff6f69****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The type of the output image access URL. Valid values:
	//
	// - oss: the back-to-origin URL.
	//
	// - cdn: the accelerated URL. This is the default value.
	//
	// example:
	//
	// cdn
	OutputType *string `json:"OutputType,omitempty" xml:"OutputType,omitempty"`
}

func (s GetImageInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetImageInfoRequest) GoString() string {
	return s.String()
}

func (s *GetImageInfoRequest) GetAuthTimeout() *int64 {
	return s.AuthTimeout
}

func (s *GetImageInfoRequest) GetImageId() *string {
	return s.ImageId
}

func (s *GetImageInfoRequest) GetOutputType() *string {
	return s.OutputType
}

func (s *GetImageInfoRequest) SetAuthTimeout(v int64) *GetImageInfoRequest {
	s.AuthTimeout = &v
	return s
}

func (s *GetImageInfoRequest) SetImageId(v string) *GetImageInfoRequest {
	s.ImageId = &v
	return s
}

func (s *GetImageInfoRequest) SetOutputType(v string) *GetImageInfoRequest {
	s.OutputType = &v
	return s
}

func (s *GetImageInfoRequest) Validate() error {
	return dara.Validate(s)
}
