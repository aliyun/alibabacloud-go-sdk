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
	// The time when the image URL expires. Unit: seconds.
	//
	// 	- If the OutputType parameter is set to cdn:
	//
	//     	- This parameter takes effect only if URL authentication is enabled. Otherwise, the image URL does not expire.
	//
	//     	- Minimum value: 1.
	//
	//     	- Maximum value: unlimited.
	//
	//     	- Default value: The default validity period that is specified in URL authentication is used.
	//
	// 	- If the OutputType parameter is set to oss:
	//
	//     	- This parameter takes effect only when the ACL of the Object Storage Service (OSS) bucket is private. Otherwise, the image URL does not expire.
	//
	//     	- Minimum value: 1.
	//
	//     	- If you store the image in the VOD bucket, the maximum value of this parameter is **2592000*	- (30 days). If you store the image in an OSS bucket, the maximum value of this parameter is **129600*	- (36 hours). The maximum value is limited to reduce security risks of the origin.
	//
	//     	- Default value: 3600.
	//
	// example:
	//
	// 3600
	AuthTimeout *int64 `json:"AuthTimeout,omitempty" xml:"AuthTimeout,omitempty"`
	// The image IDs. Separate multiple IDs with commas (,). You can specify up to 20 image IDs. You can use one of the following methods to obtain the ID:
	//
	// 	- Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com/) and choose **Media Files > Images*	- in the left-side navigation pane.
	//
	// 	- Obtain the value of ImageId from the response to the CreateUploadImage operation that you call to obtain the upload URL and credential.
	//
	// 	- Obtain the value of ImageId from the response to the [SearchMedia](~~SearchMedia~~) operation after you upload images.
	//
	// This parameter is required.
	//
	// example:
	//
	// bbc65bba53fed90de118a7849****,594228cdd14b4d069fc17a8c4a****
	ImageIds *string `json:"ImageIds,omitempty" xml:"ImageIds,omitempty"`
	// The type of the output image URL. Valid values:
	//
	// 	- oss: OSS URL
	//
	// 	- cdn: CDN URL
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
