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
	// The expiration time of the image access URL. Unit: seconds.
	//
	// - If OutputType is set to cdn:
	//
	//     - Only image URLs for which URL authentication is enabled expire. Otherwise, the URLs are permanently valid.
	//
	//     - Minimum value: 1.
	//
	//     - Maximum value: unlimited.
	//
	//     - Default value: If this parameter is not specified, the default validity period configured for URL authentication is used.
	//
	// - If OutputType is set to oss:
	//
	//     - Only playback URLs for images stored in private buckets expire. Otherwise, the URLs are permanently valid.
	//
	//     - Minimum value: 1.
	//
	//     - Maximum value: To reduce security risks to the origin server, the maximum value is **2592000*	- (30 days) if the image is stored in an ApsaraVideo VOD bucket, and **129600*	- (36 hours) if the image is stored in a custom OSS bucket.
	//
	//     - Default value: If this parameter is not specified, the value is 3600.
	//
	// example:
	//
	// 3600
	AuthTimeout *int64 `json:"AuthTimeout,omitempty" xml:"AuthTimeout,omitempty"`
	// The list of image IDs. Separate multiple IDs with commas (,). You can specify up to 20 IDs. You can obtain the image IDs by using one of the following methods:
	//
	// - Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com/) and choose **Media Files > Images*	- to view the IDs.
	//
	// - Call the [CreateUploadImage](~~CreateUploadImage~~) operation to obtain the upload URL and credential. The image ID is returned in the response.
	//
	// - Call the [SearchMedia](~~SearchMedia~~) operation to query images. The image ID is returned in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// bbc65bba53fed90de118a7849****,594228cdd14b4d069fc17a8c4a****
	ImageIds *string `json:"ImageIds,omitempty" xml:"ImageIds,omitempty"`
	// The type of the output image access URL. Valid values:
	//
	// - oss: back-to-origin URL.
	//
	// - cdn (default): accelerated URL.
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
