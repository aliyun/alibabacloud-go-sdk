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
	// The time when the image URL expires. Unit: seconds.
	//
	// 	- If you set OutputType to cdn:
	//
	//     	- This parameter takes effect only if URL authentication is enabled. Otherwise, the image URL does not expire.
	//
	//     	- Minimum value: 1.
	//
	//     	- Maximum value: unlimited.
	//
	//     	- Default value: If you leave this parameter empty, the default validity period that is specified in URL signing is used.
	//
	// 	- If you set OutputType to oss:
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
	// The ID of the image. You can use one of the following methods to obtain the ID:
	//
	// 	- Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com/). In the left-side navigation pane, choose Media Files > Image. On the Image page, view the image ID.
	//
	// 	- Obtain the image ID from the response to the [CreateUploadImage](~~CreateUploadImage~~) operation that you call to obtain the upload URL and credential.
	//
	// 	- Obtain the image ID from the response to the [SearchMedia](~~SearchMedia~~) operation that you call to query the image.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3e34733b40b9a96ccf5c1ff6f69****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
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
