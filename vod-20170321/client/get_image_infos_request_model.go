// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthTimeout(v int64) *GetImageInfosRequest
	GetAuthTimeout() *int64
	SetImageIds(v string) *GetImageInfosRequest
	GetImageIds() *string
	SetOutputType(v string) *GetImageInfosRequest
	GetOutputType() *string
}

type GetImageInfosRequest struct {
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
	//     - The image URL expires only if the storage permission is set to private. Otherwise, the URL is permanently valid.
	//
	//     - Minimum value: 1.
	//
	//     - Maximum value: To reduce security risks to the origin server, the maximum value is **2592000*	- (30 days) if the image is stored in a bucket managed by ApsaraVideo VOD, and **129600*	- (36 hours) if the image is stored in your own OSS bucket.
	//
	//     - Default value: If this parameter is not specified, the value is 3600.
	//
	// example:
	//
	// 3600
	AuthTimeout *int64 `json:"AuthTimeout,omitempty" xml:"AuthTimeout,omitempty"`
	// The list of image IDs. Separate multiple IDs with commas (,). A maximum of 20 IDs are supported. You can obtain image IDs by using the following methods:
	//
	// - Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com/) and choose **Media Files > Images*	- to view the IDs.
	//
	// - Obtain the IDs from the response when you call [CreateUploadImage](~~CreateUploadImage~~) to obtain the upload URL and credential.
	//
	// - Obtain the IDs from the response when you call [SearchMedia](~~SearchMedia~~) to query images.
	//
	// This parameter is required.
	//
	// example:
	//
	// bbc65bba53fed90de118a7849****,594228cdd14b4d069fc17a8c4a****
	ImageIds *string `json:"ImageIds,omitempty" xml:"ImageIds,omitempty"`
	// The type of the image access URL to return. Valid values:
	//
	// - oss: the storage address.
	//
	// - cdn (default): the CDN-accelerated URL.
	//
	// example:
	//
	// cdn
	OutputType *string `json:"OutputType,omitempty" xml:"OutputType,omitempty"`
}

func (s GetImageInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s GetImageInfosRequest) GoString() string {
	return s.String()
}

func (s *GetImageInfosRequest) GetAuthTimeout() *int64 {
	return s.AuthTimeout
}

func (s *GetImageInfosRequest) GetImageIds() *string {
	return s.ImageIds
}

func (s *GetImageInfosRequest) GetOutputType() *string {
	return s.OutputType
}

func (s *GetImageInfosRequest) SetAuthTimeout(v int64) *GetImageInfosRequest {
	s.AuthTimeout = &v
	return s
}

func (s *GetImageInfosRequest) SetImageIds(v string) *GetImageInfosRequest {
	s.ImageIds = &v
	return s
}

func (s *GetImageInfosRequest) SetOutputType(v string) *GetImageInfosRequest {
	s.OutputType = &v
	return s
}

func (s *GetImageInfosRequest) Validate() error {
	return dara.Validate(s)
}
