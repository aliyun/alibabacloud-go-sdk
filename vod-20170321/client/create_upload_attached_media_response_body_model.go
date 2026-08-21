// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUploadAttachedMediaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFileURL(v string) *CreateUploadAttachedMediaResponseBody
	GetFileURL() *string
	SetMediaId(v string) *CreateUploadAttachedMediaResponseBody
	GetMediaId() *string
	SetMediaURL(v string) *CreateUploadAttachedMediaResponseBody
	GetMediaURL() *string
	SetRequestId(v string) *CreateUploadAttachedMediaResponseBody
	GetRequestId() *string
	SetUploadAddress(v string) *CreateUploadAttachedMediaResponseBody
	GetUploadAddress() *string
	SetUploadAuth(v string) *CreateUploadAttachedMediaResponseBody
	GetUploadAuth() *string
}

type CreateUploadAttachedMediaResponseBody struct {
	// The OSS URL of the auxiliary media asset file (without authentication).
	//
	// When you add an image watermark template, this URL can be used as the request parameter `FileUrl` of the [AddWatermark](~~AddWatermark~~) operation.
	//
	// example:
	//
	// https://****.oss-cn-shanghai.aliyuncs.com/watermark/****.mov
	FileURL *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	// The media asset ID.
	//
	// example:
	//
	// 97dc17a5abc3668489b84ce9****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The access URL of the media asset.
	//
	// If a CDN domain name is configured, a CDN URL is returned. Otherwise, an OSS URL is returned.
	//
	// > If the returned MediaURL is inaccessible in a browser (403), you have enabled URL authentication for the VOD domain name. You can disable [URL authentication](https://help.aliyun.com/document_detail/86090.html) or [generate an authentication signature](https://help.aliyun.com/document_detail/57007.html) yourself.
	//
	// example:
	//
	// http://example.aliyundoc.com/watermark/****.mov?auth_key=****
	MediaURL *string `json:"MediaURL,omitempty" xml:"MediaURL,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 73254DE5-F260-4720-D06856B63C01****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The upload URL.
	//
	// > The upload URL returned by the operation is a Base64-encoded value. When you use the SDK or API to upload media assets, you must Base64-decode the value before use. Only uploads by using the OSS native SDK or OSS API require you to parse UploadAddress yourself.
	//
	// example:
	//
	// LWNuLXNoYW5naGFpLmFsaXl1b****
	UploadAddress *string `json:"UploadAddress,omitempty" xml:"UploadAddress,omitempty"`
	// The upload credential.
	//
	// > The upload credential returned by the operation is a Base64-encoded value. When you use the SDK or API to upload media assets, you must Base64-decode the value before use. Only uploads by using the OSS native SDK or OSS API require you to parse UploadAuth yourself.
	//
	// example:
	//
	// UzFnUjFxNkZ0NUIZTaklyNWJoQ00zdHF****
	UploadAuth *string `json:"UploadAuth,omitempty" xml:"UploadAuth,omitempty"`
}

func (s CreateUploadAttachedMediaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUploadAttachedMediaResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUploadAttachedMediaResponseBody) GetFileURL() *string {
	return s.FileURL
}

func (s *CreateUploadAttachedMediaResponseBody) GetMediaId() *string {
	return s.MediaId
}

func (s *CreateUploadAttachedMediaResponseBody) GetMediaURL() *string {
	return s.MediaURL
}

func (s *CreateUploadAttachedMediaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUploadAttachedMediaResponseBody) GetUploadAddress() *string {
	return s.UploadAddress
}

func (s *CreateUploadAttachedMediaResponseBody) GetUploadAuth() *string {
	return s.UploadAuth
}

func (s *CreateUploadAttachedMediaResponseBody) SetFileURL(v string) *CreateUploadAttachedMediaResponseBody {
	s.FileURL = &v
	return s
}

func (s *CreateUploadAttachedMediaResponseBody) SetMediaId(v string) *CreateUploadAttachedMediaResponseBody {
	s.MediaId = &v
	return s
}

func (s *CreateUploadAttachedMediaResponseBody) SetMediaURL(v string) *CreateUploadAttachedMediaResponseBody {
	s.MediaURL = &v
	return s
}

func (s *CreateUploadAttachedMediaResponseBody) SetRequestId(v string) *CreateUploadAttachedMediaResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUploadAttachedMediaResponseBody) SetUploadAddress(v string) *CreateUploadAttachedMediaResponseBody {
	s.UploadAddress = &v
	return s
}

func (s *CreateUploadAttachedMediaResponseBody) SetUploadAuth(v string) *CreateUploadAttachedMediaResponseBody {
	s.UploadAuth = &v
	return s
}

func (s *CreateUploadAttachedMediaResponseBody) Validate() error {
	return dara.Validate(s)
}
