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
	// The validity period of the image access URL. Unit: seconds.
	//
	// - If OutputType is set to cdn:
	//
	//     - The image URL expires only if URL signing is enabled. Otherwise, the URL is permanently valid.
	//
	//     - Minimum value: 1.
	//
	//     - Maximum value: unlimited.
	//
	//     - Default value: If this parameter is not specified, the default validity period specified in URL signing is used.
	//
	// - If OutputType is set to oss:
	//
	//     - The playback URL expires only if the storage permission is set to private. Otherwise, the URL is permanently valid.
	//
	//     - Minimum value: 1.
	//
	//     - Maximum value: To reduce security risks to the origin server, the maximum value is **2592000*	- (30 days) if the image is stored in a VOD system bucket, and **129600*	- (36 hours) if the image is stored in your own OSS bucket.
	//
	//     - Default value: If this parameter is not specified, the value is 3600.
	//
	// example:
	//
	// 3600
	AuthTimeout *int64 `json:"AuthTimeout,omitempty" xml:"AuthTimeout,omitempty"`
	// The image ID. You can obtain the image ID by using one of the following methods:
	//
	// - Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com/) and choose **Media Files > Images*	- to view the ID.
	//
	// - Obtain the ID from the response of the [CreateUploadImage](~~CreateUploadImage~~) operation when you retrieve the upload URL and credential.
	//
	// - Obtain the ID from the response of the [SearchMedia](~~SearchMedia~~) operation when you query images.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3e34733b40b9a96ccf5c1ff6f69****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The type of the image access URL to return. Valid values:
	//
	// - oss: the origin URL.
	//
	// - cdn (default): the accelerated URL.
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
